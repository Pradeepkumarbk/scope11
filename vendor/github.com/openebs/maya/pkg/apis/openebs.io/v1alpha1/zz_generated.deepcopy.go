// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASTemplate) DeepCopyInto(out *CASTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASTemplate.
func (in *CASTemplate) DeepCopy() *CASTemplate {
	if in == nil {
		return nil
	}
	out := new(CASTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CASTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASTemplateList) DeepCopyInto(out *CASTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CASTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASTemplateList.
func (in *CASTemplateList) DeepCopy() *CASTemplateList {
	if in == nil {
		return nil
	}
	out := new(CASTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CASTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASTemplateSpec) DeepCopyInto(out *CASTemplateSpec) {
	*out = *in
	if in.Defaults != nil {
		in, out := &in.Defaults, &out.Defaults
		*out = make([]Config, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.RunTasks.DeepCopyInto(&out.RunTasks)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASTemplateSpec.
func (in *CASTemplateSpec) DeepCopy() *CASTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(CASTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASVolume) DeepCopyInto(out *CASVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASVolume.
func (in *CASVolume) DeepCopy() *CASVolume {
	if in == nil {
		return nil
	}
	out := new(CASVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASVolumeList) DeepCopyInto(out *CASVolumeList) {
	*out = *in
	in.ListOptions.DeepCopyInto(&out.ListOptions)
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CASVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASVolumeList.
func (in *CASVolumeList) DeepCopy() *CASVolumeList {
	if in == nil {
		return nil
	}
	out := new(CASVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASVolumeSpec) DeepCopyInto(out *CASVolumeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASVolumeSpec.
func (in *CASVolumeSpec) DeepCopy() *CASVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(CASVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASVolumeStatus) DeepCopyInto(out *CASVolumeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASVolumeStatus.
func (in *CASVolumeStatus) DeepCopy() *CASVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(CASVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPool) DeepCopyInto(out *CStorPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPool.
func (in *CStorPool) DeepCopy() *CStorPool {
	if in == nil {
		return nil
	}
	out := new(CStorPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolAttr) DeepCopyInto(out *CStorPoolAttr) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolAttr.
func (in *CStorPoolAttr) DeepCopy() *CStorPoolAttr {
	if in == nil {
		return nil
	}
	out := new(CStorPoolAttr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolList) DeepCopyInto(out *CStorPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolList.
func (in *CStorPoolList) DeepCopy() *CStorPoolList {
	if in == nil {
		return nil
	}
	out := new(CStorPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolSpec) DeepCopyInto(out *CStorPoolSpec) {
	*out = *in
	in.Disks.DeepCopyInto(&out.Disks)
	out.PoolSpec = in.PoolSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolSpec.
func (in *CStorPoolSpec) DeepCopy() *CStorPoolSpec {
	if in == nil {
		return nil
	}
	out := new(CStorPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolStatus) DeepCopyInto(out *CStorPoolStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolStatus.
func (in *CStorPoolStatus) DeepCopy() *CStorPoolStatus {
	if in == nil {
		return nil
	}
	out := new(CStorPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolume) DeepCopyInto(out *CStorVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolume.
func (in *CStorVolume) DeepCopy() *CStorVolume {
	if in == nil {
		return nil
	}
	out := new(CStorVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeList) DeepCopyInto(out *CStorVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeList.
func (in *CStorVolumeList) DeepCopy() *CStorVolumeList {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplica) DeepCopyInto(out *CStorVolumeReplica) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplica.
func (in *CStorVolumeReplica) DeepCopy() *CStorVolumeReplica {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumeReplica) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplicaList) DeepCopyInto(out *CStorVolumeReplicaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorVolumeReplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplicaList.
func (in *CStorVolumeReplicaList) DeepCopy() *CStorVolumeReplicaList {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplicaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumeReplicaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplicaSpec) DeepCopyInto(out *CStorVolumeReplicaSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplicaSpec.
func (in *CStorVolumeReplicaSpec) DeepCopy() *CStorVolumeReplicaSpec {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplicaStatus) DeepCopyInto(out *CStorVolumeReplicaStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplicaStatus.
func (in *CStorVolumeReplicaStatus) DeepCopy() *CStorVolumeReplicaStatus {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeSpec) DeepCopyInto(out *CStorVolumeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeSpec.
func (in *CStorVolumeSpec) DeepCopy() *CStorVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeStatus) DeepCopyInto(out *CStorVolumeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeStatus.
func (in *CStorVolumeStatus) DeepCopy() *CStorVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CasPool) DeepCopyInto(out *CasPool) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CasPool.
func (in *CasPool) DeepCopy() *CasPool {
	if in == nil {
		return nil
	}
	out := new(CasPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Disk) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskAttr) DeepCopyInto(out *DiskAttr) {
	*out = *in
	if in.DiskList != nil {
		in, out := &in.DiskList, &out.DiskList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskAttr.
func (in *DiskAttr) DeepCopy() *DiskAttr {
	if in == nil {
		return nil
	}
	out := new(DiskAttr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskCapacity) DeepCopyInto(out *DiskCapacity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskCapacity.
func (in *DiskCapacity) DeepCopy() *DiskCapacity {
	if in == nil {
		return nil
	}
	out := new(DiskCapacity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskDetails) DeepCopyInto(out *DiskDetails) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskDetails.
func (in *DiskDetails) DeepCopy() *DiskDetails {
	if in == nil {
		return nil
	}
	out := new(DiskDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskDevLink) DeepCopyInto(out *DiskDevLink) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskDevLink.
func (in *DiskDevLink) DeepCopy() *DiskDevLink {
	if in == nil {
		return nil
	}
	out := new(DiskDevLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskList) DeepCopyInto(out *DiskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Disk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskList.
func (in *DiskList) DeepCopy() *DiskList {
	if in == nil {
		return nil
	}
	out := new(DiskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSpec) DeepCopyInto(out *DiskSpec) {
	*out = *in
	out.Capacity = in.Capacity
	out.Details = in.Details
	if in.DevLinks != nil {
		in, out := &in.DevLinks, &out.DevLinks
		*out = make([]DiskDevLink, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSpec.
func (in *DiskSpec) DeepCopy() *DiskSpec {
	if in == nil {
		return nil
	}
	out := new(DiskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskStatus) DeepCopyInto(out *DiskStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskStatus.
func (in *DiskStatus) DeepCopy() *DiskStatus {
	if in == nil {
		return nil
	}
	out := new(DiskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunTask) DeepCopyInto(out *RunTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunTask.
func (in *RunTask) DeepCopy() *RunTask {
	if in == nil {
		return nil
	}
	out := new(RunTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunTaskSpec) DeepCopyInto(out *RunTaskSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunTaskSpec.
func (in *RunTaskSpec) DeepCopy() *RunTaskSpec {
	if in == nil {
		return nil
	}
	out := new(RunTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunTasks) DeepCopyInto(out *RunTasks) {
	*out = *in
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunTasks.
func (in *RunTasks) DeepCopy() *RunTasks {
	if in == nil {
		return nil
	}
	out := new(RunTasks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePool) DeepCopyInto(out *StoragePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePool.
func (in *StoragePool) DeepCopy() *StoragePool {
	if in == nil {
		return nil
	}
	out := new(StoragePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolClaim) DeepCopyInto(out *StoragePoolClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolClaim.
func (in *StoragePoolClaim) DeepCopy() *StoragePoolClaim {
	if in == nil {
		return nil
	}
	out := new(StoragePoolClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePoolClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolClaimList) DeepCopyInto(out *StoragePoolClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoragePoolClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolClaimList.
func (in *StoragePoolClaimList) DeepCopy() *StoragePoolClaimList {
	if in == nil {
		return nil
	}
	out := new(StoragePoolClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePoolClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolClaimSpec) DeepCopyInto(out *StoragePoolClaimSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Disks.DeepCopyInto(&out.Disks)
	out.PoolSpec = in.PoolSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolClaimSpec.
func (in *StoragePoolClaimSpec) DeepCopy() *StoragePoolClaimSpec {
	if in == nil {
		return nil
	}
	out := new(StoragePoolClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolClaimStatus) DeepCopyInto(out *StoragePoolClaimStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolClaimStatus.
func (in *StoragePoolClaimStatus) DeepCopy() *StoragePoolClaimStatus {
	if in == nil {
		return nil
	}
	out := new(StoragePoolClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolList) DeepCopyInto(out *StoragePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoragePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolList.
func (in *StoragePoolList) DeepCopy() *StoragePoolList {
	if in == nil {
		return nil
	}
	out := new(StoragePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolSpec) DeepCopyInto(out *StoragePoolSpec) {
	*out = *in
	in.Disks.DeepCopyInto(&out.Disks)
	out.PoolSpec = in.PoolSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolSpec.
func (in *StoragePoolSpec) DeepCopy() *StoragePoolSpec {
	if in == nil {
		return nil
	}
	out := new(StoragePoolSpec)
	in.DeepCopyInto(out)
	return out
}
