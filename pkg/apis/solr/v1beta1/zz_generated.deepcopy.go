// +build !ignore_autogenerated

/*
Copyright 2019 Bloomberg Finance LP.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerImage) DeepCopyInto(out *ContainerImage) {
	*out = *in
	return
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
func (in *EtcdSpec) DeepCopyInto(out *EtcdSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ContainerImage)
		**out = **in
	}
	if in.PersistentVolumeClaimSpec != nil {
		in, out := &in.PersistentVolumeClaimSpec, &out.PersistentVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdSpec.
func (in *EtcdSpec) DeepCopy() *EtcdSpec {
	if in == nil {
		return nil
	}
	out := new(EtcdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FullZetcdSpec) DeepCopyInto(out *FullZetcdSpec) {
	*out = *in
	if in.EtcdSpec != nil {
		in, out := &in.EtcdSpec, &out.EtcdSpec
		*out = new(EtcdSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ZetcdSpec != nil {
		in, out := &in.ZetcdSpec, &out.ZetcdSpec
		*out = new(ZetcdSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FullZetcdSpec.
func (in *FullZetcdSpec) DeepCopy() *FullZetcdSpec {
	if in == nil {
		return nil
	}
	out := new(FullZetcdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvidedZookeeper) DeepCopyInto(out *ProvidedZookeeper) {
	*out = *in
	if in.Zookeeper != nil {
		in, out := &in.Zookeeper, &out.Zookeeper
		*out = new(ZookeeperSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Zetcd != nil {
		in, out := &in.Zetcd, &out.Zetcd
		*out = new(FullZetcdSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvidedZookeeper.
func (in *ProvidedZookeeper) DeepCopy() *ProvidedZookeeper {
	if in == nil {
		return nil
	}
	out := new(ProvidedZookeeper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrBackup) DeepCopyInto(out *SolrBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SolrBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
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
func (in *SolrBackupSpec) DeepCopyInto(out *SolrBackupSpec) {
	*out = *in
	return
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
	return
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
	return
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
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SolrCloud, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
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
	if in.PersistentVolumeClaimSpec != nil {
		in, out := &in.PersistentVolumeClaimSpec, &out.PersistentVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.BusyBoxImage != nil {
		in, out := &in.BusyBoxImage, &out.BusyBoxImage
		*out = new(ContainerImage)
		**out = **in
	}
	return
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
	return
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
func (in *SolrNodeStatus) DeepCopyInto(out *SolrNodeStatus) {
	*out = *in
	return
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
func (in *ZetcdSpec) DeepCopyInto(out *ZetcdSpec) {
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZetcdSpec.
func (in *ZetcdSpec) DeepCopy() *ZetcdSpec {
	if in == nil {
		return nil
	}
	out := new(ZetcdSpec)
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
	return
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
func (in *ZookeeperRef) DeepCopyInto(out *ZookeeperRef) {
	*out = *in
	if in.ConnectionInfo != nil {
		in, out := &in.ConnectionInfo, &out.ConnectionInfo
		*out = new(ZookeeperConnectionInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvidedZookeeper != nil {
		in, out := &in.ProvidedZookeeper, &out.ProvidedZookeeper
		*out = new(ProvidedZookeeper)
		(*in).DeepCopyInto(*out)
	}
	return
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
	if in.PersistentVolumeClaimSpec != nil {
		in, out := &in.PersistentVolumeClaimSpec, &out.PersistentVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
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
