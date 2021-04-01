#!/usr/bin/env bash
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# exit immediately when a command fails
set -e
# only exit with zero if all commands of the pipeline exit successfully
set -o pipefail
# error on unset variables
set -u

show_help() {
cat << EOF
Usage: ./hack/release/smoke_test/verify_all.sh [-h] -v VERSION -l LOCATION

Verify checksums and signatures of all release artifacts.
Check that the docker image contains the necessary LICENSE and NOTICE.

    -h  Display this help and exit
    -v  Version of the Solr Operator
    -l  Base location of the staged artifacts. Can be a URL or relative or absolute file path.
EOF
}

OPTIND=1
while getopts hv:l: opt; do
    case $opt in
        h)
            show_help
            exit 0
            ;;
        v)  VERSION=$OPTARG
            ;;
        l)  LOCATION=$OPTARG
            ;;
        *)
            show_help >&2
            exit 1
            ;;
    esac
done
shift "$((OPTIND-1))"   # Discard the options and sentinel --

if [[ -z "${VERSION:-}" ]]; then
  error "Specify a project version through -v, or through the VERSION env var"; exit 1
fi
if [[ -z "${LOCATION:-}" ]]; then
  error "Specify an base artifact location -l, or through the LOCATION env var"; exit 1
fi

TMP_DIR=$(mktemp -d --tmpdir "solr-operator-smoke-test-source-XXXXXX")

# If LOCATION is not a URL, then get the absolute path
if ! (echo "${LOCATION}" | grep -E "http://"); then
  LOCATION=$(cd "${LOCATION}"; pwd)
fi

echo "Import Solr Keys"
curl -sL0 "https://dist.apache.org/repos/dist/release/solr/KEYS" | gpg2 --import --quiet

# First generate the temporary public key ring
gpg --export >~/.gnupg/pubring.gpg

echo "Download all artifacts and verify signatures"
# Do all logic in temporary directory
(
  cd "${TMP_DIR}"

  if (echo "${LOCATION}" | grep -E "http"); then
    # Download Source files from the staged location
    wget -r -np -nH -nd --level=1 "${LOCATION}/"

    # Download Helm files from the staged location
    wget -r -np -nH -nd --level=1 -P "helm-charts" "${LOCATION}/helm-charts/"

    # Download CRD files from the staged location
    wget -r -np -nH -nd --level=1 -P "crds" "${LOCATION}/crds/"
  else
    cp -r "${LOCATION}/"* .
  fi

  for artifact_directory in $(find * -type d); do
    (
      cd "${artifact_directory}"

      for artifact in $(find * -type f ! \( -name '*.asc' -o -name '*.sha512' -o -name '*.prov' \) ); do
        sha512sum -c "${artifact}.sha512" \
          || { echo "Invalid sha512 for ${artifact}. Aborting!"; exit 1; }
        gpg2 --verify "${artifact}.asc" "${artifact}" \
          || { echo "Invalid signature for ${artifact}. Aborting!"; exit 1; }
      done

      # If a helm chart has a provenance file, verify it
      if [[ -f "${artifact}.prov" ]]; then
        helm verify "${artifact}"
      fi
    )
  done
)

# Delete temporary source download directory
rm -rf "${TMP_DIR}"

echo "Successfully verified all artifacts!"