//go:build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/kluctl/kluctl/v2/pkg/types"
	"github.com/kluctl/kluctl/v2/pkg/types/result"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Decryption) DeepCopyInto(out *Decryption) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Decryption.
func (in *Decryption) DeepCopy() *Decryption {
	if in == nil {
		return nil
	}
	out := new(Decryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmCredentials) DeepCopyInto(out *HelmCredentials) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmCredentials.
func (in *HelmCredentials) DeepCopy() *HelmCredentials {
	if in == nil {
		return nil
	}
	out := new(HelmCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeployment) DeepCopyInto(out *KluctlDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeployment.
func (in *KluctlDeployment) DeepCopy() *KluctlDeployment {
	if in == nil {
		return nil
	}
	out := new(KluctlDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KluctlDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeploymentList) DeepCopyInto(out *KluctlDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KluctlDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeploymentList.
func (in *KluctlDeploymentList) DeepCopy() *KluctlDeploymentList {
	if in == nil {
		return nil
	}
	out := new(KluctlDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KluctlDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeploymentSpec) DeepCopyInto(out *KluctlDeploymentSpec) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Credentials.DeepCopyInto(&out.Credentials)
	if in.Decryption != nil {
		in, out := &in.Decryption, &out.Decryption
		*out = new(Decryption)
		(*in).DeepCopyInto(*out)
	}
	out.Interval = in.Interval
	if in.RetryInterval != nil {
		in, out := &in.RetryInterval, &out.RetryInterval
		*out = new(v1.Duration)
		**out = **in
	}
	if in.DeployInterval != nil {
		in, out := &in.DeployInterval, &out.DeployInterval
		*out = new(SafeDuration)
		**out = **in
	}
	if in.ValidateInterval != nil {
		in, out := &in.ValidateInterval, &out.ValidateInterval
		*out = new(SafeDuration)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.HelmCredentials != nil {
		in, out := &in.HelmCredentials, &out.HelmCredentials
		*out = make([]HelmCredentials, len(*in))
		copy(*out, *in)
	}
	if in.KubeConfig != nil {
		in, out := &in.KubeConfig, &out.KubeConfig
		*out = new(KubeConfig)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.TargetNameOverride != nil {
		in, out := &in.TargetNameOverride, &out.TargetNameOverride
		*out = new(string)
		**out = **in
	}
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(string)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]types.FixedImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludeTags != nil {
		in, out := &in.IncludeTags, &out.IncludeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeTags != nil {
		in, out := &in.ExcludeTags, &out.ExcludeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludeDeploymentDirs != nil {
		in, out := &in.IncludeDeploymentDirs, &out.IncludeDeploymentDirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeDeploymentDirs != nil {
		in, out := &in.ExcludeDeploymentDirs, &out.ExcludeDeploymentDirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ManualObjectsHash != nil {
		in, out := &in.ManualObjectsHash, &out.ManualObjectsHash
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeploymentSpec.
func (in *KluctlDeploymentSpec) DeepCopy() *KluctlDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(KluctlDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeploymentStatus) DeepCopyInto(out *KluctlDeploymentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProjectKey != nil {
		in, out := &in.ProjectKey, &out.ProjectKey
		*out = new(result.ProjectKey)
		**out = **in
	}
	if in.TargetKey != nil {
		in, out := &in.TargetKey, &out.TargetKey
		*out = new(result.TargetKey)
		**out = **in
	}
	if in.LastManualObjectsHash != nil {
		in, out := &in.LastManualObjectsHash, &out.LastManualObjectsHash
		*out = new(string)
		**out = **in
	}
	if in.LastDeployResult != nil {
		in, out := &in.LastDeployResult, &out.LastDeployResult
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.LastValidateResult != nil {
		in, out := &in.LastValidateResult, &out.LastValidateResult
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.LastDriftDetectionResult != nil {
		in, out := &in.LastDriftDetectionResult, &out.LastDriftDetectionResult
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeploymentStatus.
func (in *KluctlDeploymentStatus) DeepCopy() *KluctlDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(KluctlDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeConfig) DeepCopyInto(out *KubeConfig) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeConfig.
func (in *KubeConfig) DeepCopy() *KubeConfig {
	if in == nil {
		return nil
	}
	out := new(KubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectCredentials) DeepCopyInto(out *ProjectCredentials) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = make([]ProjectCredentialsGit, len(*in))
		copy(*out, *in)
	}
	if in.Oci != nil {
		in, out := &in.Oci, &out.Oci
		*out = make([]ProjectCredentialsOci, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectCredentials.
func (in *ProjectCredentials) DeepCopy() *ProjectCredentials {
	if in == nil {
		return nil
	}
	out := new(ProjectCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectCredentialsGit) DeepCopyInto(out *ProjectCredentialsGit) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectCredentialsGit.
func (in *ProjectCredentialsGit) DeepCopy() *ProjectCredentialsGit {
	if in == nil {
		return nil
	}
	out := new(ProjectCredentialsGit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectCredentialsOci) DeepCopyInto(out *ProjectCredentialsOci) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectCredentialsOci.
func (in *ProjectCredentialsOci) DeepCopy() *ProjectCredentialsOci {
	if in == nil {
		return nil
	}
	out := new(ProjectCredentialsOci)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSource) DeepCopyInto(out *ProjectSource) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(ProjectSourceGit)
		(*in).DeepCopyInto(*out)
	}
	if in.Oci != nil {
		in, out := &in.Oci, &out.Oci
		*out = new(ProjectSourceOci)
		(*in).DeepCopyInto(*out)
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = (*in).DeepCopy()
	}
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(types.GitRef)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalObjectReference)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]ProjectCredentialsGit, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSource.
func (in *ProjectSource) DeepCopy() *ProjectSource {
	if in == nil {
		return nil
	}
	out := new(ProjectSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSourceGit) DeepCopyInto(out *ProjectSourceGit) {
	*out = *in
	in.URL.DeepCopyInto(&out.URL)
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(types.GitRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSourceGit.
func (in *ProjectSourceGit) DeepCopy() *ProjectSourceGit {
	if in == nil {
		return nil
	}
	out := new(ProjectSourceGit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSourceOci) DeepCopyInto(out *ProjectSourceOci) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(types.OciRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSourceOci.
func (in *ProjectSourceOci) DeepCopy() *ProjectSourceOci {
	if in == nil {
		return nil
	}
	out := new(ProjectSourceOci)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SafeDuration) DeepCopyInto(out *SafeDuration) {
	*out = *in
	out.Duration = in.Duration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SafeDuration.
func (in *SafeDuration) DeepCopy() *SafeDuration {
	if in == nil {
		return nil
	}
	out := new(SafeDuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyReference) DeepCopyInto(out *SecretKeyReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyReference.
func (in *SecretKeyReference) DeepCopy() *SecretKeyReference {
	if in == nil {
		return nil
	}
	out := new(SecretKeyReference)
	in.DeepCopyInto(out)
	return out
}
