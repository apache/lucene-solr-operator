// +build !ignore_autogenerated

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	zookeeperv1beta1 "github.com/pravega/zookeeper-operator/pkg/apis/zookeeper/v1beta1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalVolume) DeepCopyInto(out *AdditionalVolume) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.DefaultContainerMount != nil {
		in, out := &in.DefaultContainerMount, &out.DefaultContainerMount
		*out = new(v1.VolumeMount)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalVolume.
func (in *AdditionalVolume) DeepCopy() *AdditionalVolume {
	if in == nil {
		return nil
	}
	out := new(AdditionalVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPersistenceStatus) DeepCopyInto(out *BackupPersistenceStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.FinishTime != nil {
		in, out := &in.FinishTime, &out.FinishTime
		*out = (*in).DeepCopy()
	}
	if in.Successful != nil {
		in, out := &in.Successful, &out.Successful
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPersistenceStatus.
func (in *BackupPersistenceStatus) DeepCopy() *BackupPersistenceStatus {
	if in == nil {
		return nil
	}
	out := new(BackupPersistenceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionBackupStatus) DeepCopyInto(out *CollectionBackupStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.FinishTime != nil {
		in, out := &in.FinishTime, &out.FinishTime
		*out = (*in).DeepCopy()
	}
	if in.Successful != nil {
		in, out := &in.Successful, &out.Successful
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionBackupStatus.
func (in *CollectionBackupStatus) DeepCopy() *CollectionBackupStatus {
	if in == nil {
		return nil
	}
	out := new(CollectionBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapOptions) DeepCopyInto(out *ConfigMapOptions) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapOptions.
func (in *ConfigMapOptions) DeepCopy() *ConfigMapOptions {
	if in == nil {
		return nil
	}
	out := new(ConfigMapOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerImage) DeepCopyInto(out *ContainerImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerImage.
func (in *ContainerImage) DeepCopy() *ContainerImage {
	if in == nil {
		return nil
	}
	out := new(ContainerImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomExporterKubeOptions) DeepCopyInto(out *CustomExporterKubeOptions) {
	*out = *in
	if in.PodOptions != nil {
		in, out := &in.PodOptions, &out.PodOptions
		*out = new(PodOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.DeploymentOptions != nil {
		in, out := &in.DeploymentOptions, &out.DeploymentOptions
		*out = new(DeploymentOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceOptions != nil {
		in, out := &in.ServiceOptions, &out.ServiceOptions
		*out = new(ServiceOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMapOptions != nil {
		in, out := &in.ConfigMapOptions, &out.ConfigMapOptions
		*out = new(ConfigMapOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomExporterKubeOptions.
func (in *CustomExporterKubeOptions) DeepCopy() *CustomExporterKubeOptions {
	if in == nil {
		return nil
	}
	out := new(CustomExporterKubeOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSolrKubeOptions) DeepCopyInto(out *CustomSolrKubeOptions) {
	*out = *in
	if in.PodOptions != nil {
		in, out := &in.PodOptions, &out.PodOptions
		*out = new(PodOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.StatefulSetOptions != nil {
		in, out := &in.StatefulSetOptions, &out.StatefulSetOptions
		*out = new(StatefulSetOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.CommonServiceOptions != nil {
		in, out := &in.CommonServiceOptions, &out.CommonServiceOptions
		*out = new(ServiceOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.HeadlessServiceOptions != nil {
		in, out := &in.HeadlessServiceOptions, &out.HeadlessServiceOptions
		*out = new(ServiceOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeServiceOptions != nil {
		in, out := &in.NodeServiceOptions, &out.NodeServiceOptions
		*out = new(ServiceOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMapOptions != nil {
		in, out := &in.ConfigMapOptions, &out.ConfigMapOptions
		*out = new(ConfigMapOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.IngressOptions != nil {
		in, out := &in.IngressOptions, &out.IngressOptions
		*out = new(IngressOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSolrKubeOptions.
func (in *CustomSolrKubeOptions) DeepCopy() *CustomSolrKubeOptions {
	if in == nil {
		return nil
	}
	out := new(CustomSolrKubeOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentOptions) DeepCopyInto(out *DeploymentOptions) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentOptions.
func (in *DeploymentOptions) DeepCopy() *DeploymentOptions {
	if in == nil {
		return nil
	}
	out := new(DeploymentOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAddressability) DeepCopyInto(out *ExternalAddressability) {
	*out = *in
	if in.AdditionalDomainNames != nil {
		in, out := &in.AdditionalDomainNames, &out.AdditionalDomainNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAddressability.
func (in *ExternalAddressability) DeepCopy() *ExternalAddressability {
	if in == nil {
		return nil
	}
	out := new(ExternalAddressability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressOptions) DeepCopyInto(out *IngressOptions) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressOptions.
func (in *IngressOptions) DeepCopy() *IngressOptions {
	if in == nil {
		return nil
	}
	out := new(IngressOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedUpdateOptions) DeepCopyInto(out *ManagedUpdateOptions) {
	*out = *in
	if in.MaxPodsUnavailable != nil {
		in, out := &in.MaxPodsUnavailable, &out.MaxPodsUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxShardReplicasUnavailable != nil {
		in, out := &in.MaxShardReplicasUnavailable, &out.MaxShardReplicasUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedUpdateOptions.
func (in *ManagedUpdateOptions) DeepCopy() *ManagedUpdateOptions {
	if in == nil {
		return nil
	}
	out := new(ManagedUpdateOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceSource) DeepCopyInto(out *PersistenceSource) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3PersistenceSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(VolumePersistenceSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceSource.
func (in *PersistenceSource) DeepCopy() *PersistenceSource {
	if in == nil {
		return nil
	}
	out := new(PersistenceSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimTemplate) DeepCopyInto(out *PersistentVolumeClaimTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimTemplate.
func (in *PersistentVolumeClaimTemplate) DeepCopy() *PersistentVolumeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodOptions) DeepCopyInto(out *PodOptions) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]AdditionalVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvVariables != nil {
		in, out := &in.EnvVariables, &out.EnvVariables
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.StartupProbe != nil {
		in, out := &in.StartupProbe, &out.StartupProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.SidecarContainers != nil {
		in, out := &in.SidecarContainers, &out.SidecarContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodOptions.
func (in *PodOptions) DeepCopy() *PodOptions {
	if in == nil {
		return nil
	}
	out := new(PodOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3PersistenceSource) DeepCopyInto(out *S3PersistenceSource) {
	*out = *in
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(int32)
		**out = **in
	}
	out.Secrets = in.Secrets
	out.AWSCliImage = in.AWSCliImage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3PersistenceSource.
func (in *S3PersistenceSource) DeepCopy() *S3PersistenceSource {
	if in == nil {
		return nil
	}
	out := new(S3PersistenceSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Secrets) DeepCopyInto(out *S3Secrets) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Secrets.
func (in *S3Secrets) DeepCopy() *S3Secrets {
	if in == nil {
		return nil
	}
	out := new(S3Secrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceOptions) DeepCopyInto(out *ServiceOptions) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceOptions.
func (in *ServiceOptions) DeepCopy() *ServiceOptions {
	if in == nil {
		return nil
	}
	out := new(ServiceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrAddressabilityOptions) DeepCopyInto(out *SolrAddressabilityOptions) {
	*out = *in
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = new(ExternalAddressability)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrAddressabilityOptions.
func (in *SolrAddressabilityOptions) DeepCopy() *SolrAddressabilityOptions {
	if in == nil {
		return nil
	}
	out := new(SolrAddressabilityOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrBackup) DeepCopyInto(out *SolrBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrBackup.
func (in *SolrBackup) DeepCopy() *SolrBackup {
	if in == nil {
		return nil
	}
	out := new(SolrBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrBackupList) DeepCopyInto(out *SolrBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SolrBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrBackupList.
func (in *SolrBackupList) DeepCopy() *SolrBackupList {
	if in == nil {
		return nil
	}
	out := new(SolrBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrBackupRestoreOptions) DeepCopyInto(out *SolrBackupRestoreOptions) {
	*out = *in
	in.Volume.DeepCopyInto(&out.Volume)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrBackupRestoreOptions.
func (in *SolrBackupRestoreOptions) DeepCopy() *SolrBackupRestoreOptions {
	if in == nil {
		return nil
	}
	out := new(SolrBackupRestoreOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrBackupSpec) DeepCopyInto(out *SolrBackupSpec) {
	*out = *in
	if in.Collections != nil {
		in, out := &in.Collections, &out.Collections
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Persistence.DeepCopyInto(&out.Persistence)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrBackupSpec.
func (in *SolrBackupSpec) DeepCopy() *SolrBackupSpec {
	if in == nil {
		return nil
	}
	out := new(SolrBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrBackupStatus) DeepCopyInto(out *SolrBackupStatus) {
	*out = *in
	if in.CollectionBackupStatuses != nil {
		in, out := &in.CollectionBackupStatuses, &out.CollectionBackupStatuses
		*out = make([]CollectionBackupStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.PersistenceStatus.DeepCopyInto(&out.PersistenceStatus)
	if in.FinishTime != nil {
		in, out := &in.FinishTime, &out.FinishTime
		*out = (*in).DeepCopy()
	}
	if in.Successful != nil {
		in, out := &in.Successful, &out.Successful
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrBackupStatus.
func (in *SolrBackupStatus) DeepCopy() *SolrBackupStatus {
	if in == nil {
		return nil
	}
	out := new(SolrBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCloud) DeepCopyInto(out *SolrCloud) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCloud.
func (in *SolrCloud) DeepCopy() *SolrCloud {
	if in == nil {
		return nil
	}
	out := new(SolrCloud)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrCloud) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCloudList) DeepCopyInto(out *SolrCloudList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SolrCloud, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCloudList.
func (in *SolrCloudList) DeepCopy() *SolrCloudList {
	if in == nil {
		return nil
	}
	out := new(SolrCloudList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrCloudList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCloudReference) DeepCopyInto(out *SolrCloudReference) {
	*out = *in
	if in.ZookeeperConnectionInfo != nil {
		in, out := &in.ZookeeperConnectionInfo, &out.ZookeeperConnectionInfo
		*out = new(ZookeeperConnectionInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCloudReference.
func (in *SolrCloudReference) DeepCopy() *SolrCloudReference {
	if in == nil {
		return nil
	}
	out := new(SolrCloudReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCloudSpec) DeepCopyInto(out *SolrCloudSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ZookeeperRef != nil {
		in, out := &in.ZookeeperRef, &out.ZookeeperRef
		*out = new(ZookeeperRef)
		(*in).DeepCopyInto(*out)
	}
	if in.SolrImage != nil {
		in, out := &in.SolrImage, &out.SolrImage
		*out = new(ContainerImage)
		**out = **in
	}
	in.StorageOptions.DeepCopyInto(&out.StorageOptions)
	in.CustomSolrKubeOptions.DeepCopyInto(&out.CustomSolrKubeOptions)
	in.SolrAddressability.DeepCopyInto(&out.SolrAddressability)
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	if in.BusyBoxImage != nil {
		in, out := &in.BusyBoxImage, &out.BusyBoxImage
		*out = new(ContainerImage)
		**out = **in
	}
	if in.SolrTLS != nil {
		in, out := &in.SolrTLS, &out.SolrTLS
		*out = new(SolrTLSOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCloudSpec.
func (in *SolrCloudSpec) DeepCopy() *SolrCloudSpec {
	if in == nil {
		return nil
	}
	out := new(SolrCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCloudStatus) DeepCopyInto(out *SolrCloudStatus) {
	*out = *in
	if in.SolrNodes != nil {
		in, out := &in.SolrNodes, &out.SolrNodes
		*out = make([]SolrNodeStatus, len(*in))
		copy(*out, *in)
	}
	if in.ExternalCommonAddress != nil {
		in, out := &in.ExternalCommonAddress, &out.ExternalCommonAddress
		*out = new(string)
		**out = **in
	}
	in.ZookeeperConnectionInfo.DeepCopyInto(&out.ZookeeperConnectionInfo)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCloudStatus.
func (in *SolrCloudStatus) DeepCopy() *SolrCloudStatus {
	if in == nil {
		return nil
	}
	out := new(SolrCloudStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrDataStorageOptions) DeepCopyInto(out *SolrDataStorageOptions) {
	*out = *in
	if in.PersistentStorage != nil {
		in, out := &in.PersistentStorage, &out.PersistentStorage
		*out = new(SolrPersistentDataStorageOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.EphemeralStorage != nil {
		in, out := &in.EphemeralStorage, &out.EphemeralStorage
		*out = new(SolrEphemeralDataStorageOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupRestoreOptions != nil {
		in, out := &in.BackupRestoreOptions, &out.BackupRestoreOptions
		*out = new(SolrBackupRestoreOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrDataStorageOptions.
func (in *SolrDataStorageOptions) DeepCopy() *SolrDataStorageOptions {
	if in == nil {
		return nil
	}
	out := new(SolrDataStorageOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrEphemeralDataStorageOptions) DeepCopyInto(out *SolrEphemeralDataStorageOptions) {
	*out = *in
	in.EmptyDir.DeepCopyInto(&out.EmptyDir)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrEphemeralDataStorageOptions.
func (in *SolrEphemeralDataStorageOptions) DeepCopy() *SolrEphemeralDataStorageOptions {
	if in == nil {
		return nil
	}
	out := new(SolrEphemeralDataStorageOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrNodeStatus) DeepCopyInto(out *SolrNodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrNodeStatus.
func (in *SolrNodeStatus) DeepCopy() *SolrNodeStatus {
	if in == nil {
		return nil
	}
	out := new(SolrNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrPersistentDataStorageOptions) DeepCopyInto(out *SolrPersistentDataStorageOptions) {
	*out = *in
	in.PersistentVolumeClaimTemplate.DeepCopyInto(&out.PersistentVolumeClaimTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrPersistentDataStorageOptions.
func (in *SolrPersistentDataStorageOptions) DeepCopy() *SolrPersistentDataStorageOptions {
	if in == nil {
		return nil
	}
	out := new(SolrPersistentDataStorageOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrPrometheusExporter) DeepCopyInto(out *SolrPrometheusExporter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrPrometheusExporter.
func (in *SolrPrometheusExporter) DeepCopy() *SolrPrometheusExporter {
	if in == nil {
		return nil
	}
	out := new(SolrPrometheusExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrPrometheusExporter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrPrometheusExporterList) DeepCopyInto(out *SolrPrometheusExporterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SolrPrometheusExporter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrPrometheusExporterList.
func (in *SolrPrometheusExporterList) DeepCopy() *SolrPrometheusExporterList {
	if in == nil {
		return nil
	}
	out := new(SolrPrometheusExporterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrPrometheusExporterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrPrometheusExporterSpec) DeepCopyInto(out *SolrPrometheusExporterSpec) {
	*out = *in
	in.SolrReference.DeepCopyInto(&out.SolrReference)
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ContainerImage)
		**out = **in
	}
	in.CustomKubeOptions.DeepCopyInto(&out.CustomKubeOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrPrometheusExporterSpec.
func (in *SolrPrometheusExporterSpec) DeepCopy() *SolrPrometheusExporterSpec {
	if in == nil {
		return nil
	}
	out := new(SolrPrometheusExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrPrometheusExporterStatus) DeepCopyInto(out *SolrPrometheusExporterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrPrometheusExporterStatus.
func (in *SolrPrometheusExporterStatus) DeepCopy() *SolrPrometheusExporterStatus {
	if in == nil {
		return nil
	}
	out := new(SolrPrometheusExporterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrReference) DeepCopyInto(out *SolrReference) {
	*out = *in
	if in.Cloud != nil {
		in, out := &in.Cloud, &out.Cloud
		*out = new(SolrCloudReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Standalone != nil {
		in, out := &in.Standalone, &out.Standalone
		*out = new(StandaloneSolrReference)
		**out = **in
	}
	if in.SolrTLS != nil {
		in, out := &in.SolrTLS, &out.SolrTLS
		*out = new(SolrTLSOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrReference.
func (in *SolrReference) DeepCopy() *SolrReference {
	if in == nil {
		return nil
	}
	out := new(SolrReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrTLSOptions) DeepCopyInto(out *SolrTLSOptions) {
	*out = *in
	if in.PKCS12Secret != nil {
		in, out := &in.PKCS12Secret, &out.PKCS12Secret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyStorePasswordSecret != nil {
		in, out := &in.KeyStorePasswordSecret, &out.KeyStorePasswordSecret
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrTLSOptions.
func (in *SolrTLSOptions) DeepCopy() *SolrTLSOptions {
	if in == nil {
		return nil
	}
	out := new(SolrTLSOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrUpdateStrategy) DeepCopyInto(out *SolrUpdateStrategy) {
	*out = *in
	in.ManagedUpdateOptions.DeepCopyInto(&out.ManagedUpdateOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrUpdateStrategy.
func (in *SolrUpdateStrategy) DeepCopy() *SolrUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(SolrUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneSolrReference) DeepCopyInto(out *StandaloneSolrReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneSolrReference.
func (in *StandaloneSolrReference) DeepCopy() *StandaloneSolrReference {
	if in == nil {
		return nil
	}
	out := new(StandaloneSolrReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetOptions) DeepCopyInto(out *StatefulSetOptions) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetOptions.
func (in *StatefulSetOptions) DeepCopy() *StatefulSetOptions {
	if in == nil {
		return nil
	}
	out := new(StatefulSetOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateMeta) DeepCopyInto(out *TemplateMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateMeta.
func (in *TemplateMeta) DeepCopy() *TemplateMeta {
	if in == nil {
		return nil
	}
	out := new(TemplateMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumePersistenceSource) DeepCopyInto(out *VolumePersistenceSource) {
	*out = *in
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
	out.BusyBoxImage = in.BusyBoxImage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumePersistenceSource.
func (in *VolumePersistenceSource) DeepCopy() *VolumePersistenceSource {
	if in == nil {
		return nil
	}
	out := new(VolumePersistenceSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperACL) DeepCopyInto(out *ZookeeperACL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperACL.
func (in *ZookeeperACL) DeepCopy() *ZookeeperACL {
	if in == nil {
		return nil
	}
	out := new(ZookeeperACL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperConnectionInfo) DeepCopyInto(out *ZookeeperConnectionInfo) {
	*out = *in
	if in.ExternalConnectionString != nil {
		in, out := &in.ExternalConnectionString, &out.ExternalConnectionString
		*out = new(string)
		**out = **in
	}
	if in.AllACL != nil {
		in, out := &in.AllACL, &out.AllACL
		*out = new(ZookeeperACL)
		**out = **in
	}
	if in.ReadOnlyACL != nil {
		in, out := &in.ReadOnlyACL, &out.ReadOnlyACL
		*out = new(ZookeeperACL)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperConnectionInfo.
func (in *ZookeeperConnectionInfo) DeepCopy() *ZookeeperConnectionInfo {
	if in == nil {
		return nil
	}
	out := new(ZookeeperConnectionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperPodPolicy) DeepCopyInto(out *ZookeeperPodPolicy) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperPodPolicy.
func (in *ZookeeperPodPolicy) DeepCopy() *ZookeeperPodPolicy {
	if in == nil {
		return nil
	}
	out := new(ZookeeperPodPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperRef) DeepCopyInto(out *ZookeeperRef) {
	*out = *in
	if in.ConnectionInfo != nil {
		in, out := &in.ConnectionInfo, &out.ConnectionInfo
		*out = new(ZookeeperConnectionInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvidedZookeeper != nil {
		in, out := &in.ProvidedZookeeper, &out.ProvidedZookeeper
		*out = new(ZookeeperSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperRef.
func (in *ZookeeperRef) DeepCopy() *ZookeeperRef {
	if in == nil {
		return nil
	}
	out := new(ZookeeperRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperSpec) DeepCopyInto(out *ZookeeperSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ContainerImage)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(zookeeperv1beta1.Persistence)
		(*in).DeepCopyInto(*out)
	}
	in.ZookeeperPod.DeepCopyInto(&out.ZookeeperPod)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperSpec.
func (in *ZookeeperSpec) DeepCopy() *ZookeeperSpec {
	if in == nil {
		return nil
	}
	out := new(ZookeeperSpec)
	in.DeepCopyInto(out)
	return out
}
