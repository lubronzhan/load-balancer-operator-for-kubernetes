// +build !ignore_autogenerated

// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKODeploymentConfig) DeepCopyInto(out *AKODeploymentConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKODeploymentConfig.
func (in *AKODeploymentConfig) DeepCopy() *AKODeploymentConfig {
	if in == nil {
		return nil
	}
	out := new(AKODeploymentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKODeploymentConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKODeploymentConfigList) DeepCopyInto(out *AKODeploymentConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AKODeploymentConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKODeploymentConfigList.
func (in *AKODeploymentConfigList) DeepCopy() *AKODeploymentConfigList {
	if in == nil {
		return nil
	}
	out := new(AKODeploymentConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKODeploymentConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKODeploymentConfigSpec) DeepCopyInto(out *AKODeploymentConfigSpec) {
	*out = *in
	in.ClusterSelector.DeepCopyInto(&out.ClusterSelector)
	if in.WorkloadCredentialRef != nil {
		in, out := &in.WorkloadCredentialRef, &out.WorkloadCredentialRef
		*out = new(SecretRef)
		**out = **in
	}
	if in.AdminCredentialRef != nil {
		in, out := &in.AdminCredentialRef, &out.AdminCredentialRef
		*out = new(SecretRef)
		**out = **in
	}
	if in.CertificateAuthorityRef != nil {
		in, out := &in.CertificateAuthorityRef, &out.CertificateAuthorityRef
		*out = new(SecretRef)
		**out = **in
	}
	out.Tenant = in.Tenant
	in.DataNetwork.DeepCopyInto(&out.DataNetwork)
	out.ControlPlaneNetwork = in.ControlPlaneNetwork
	in.ExtraConfigs.DeepCopyInto(&out.ExtraConfigs)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKODeploymentConfigSpec.
func (in *AKODeploymentConfigSpec) DeepCopy() *AKODeploymentConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AKODeploymentConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKODeploymentConfigStatus) DeepCopyInto(out *AKODeploymentConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKODeploymentConfigStatus.
func (in *AKODeploymentConfigStatus) DeepCopy() *AKODeploymentConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AKODeploymentConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKOIngressConfig) DeepCopyInto(out *AKOIngressConfig) {
	*out = *in
	if in.DisableIngressClass != nil {
		in, out := &in.DisableIngressClass, &out.DisableIngressClass
		*out = new(bool)
		**out = **in
	}
	if in.DefaultIngressController != nil {
		in, out := &in.DefaultIngressController, &out.DefaultIngressController
		*out = new(bool)
		**out = **in
	}
	if in.NodeNetworkList != nil {
		in, out := &in.NodeNetworkList, &out.NodeNetworkList
		*out = make([]NodeNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NoPGForSNI != nil {
		in, out := &in.NoPGForSNI, &out.NoPGForSNI
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOIngressConfig.
func (in *AKOIngressConfig) DeepCopy() *AKOIngressConfig {
	if in == nil {
		return nil
	}
	out := new(AKOIngressConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKOL4Config) DeepCopyInto(out *AKOL4Config) {
	*out = *in
	if in.AdvancedL4 != nil {
		in, out := &in.AdvancedL4, &out.AdvancedL4
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOL4Config.
func (in *AKOL4Config) DeepCopy() *AKOL4Config {
	if in == nil {
		return nil
	}
	out := new(AKOL4Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKOLogConfig) DeepCopyInto(out *AKOLogConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKOLogConfig.
func (in *AKOLogConfig) DeepCopy() *AKOLogConfig {
	if in == nil {
		return nil
	}
	out := new(AKOLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKORbacConfig) DeepCopyInto(out *AKORbacConfig) {
	*out = *in
	if in.PspEnabled != nil {
		in, out := &in.PspEnabled, &out.PspEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKORbacConfig.
func (in *AKORbacConfig) DeepCopy() *AKORbacConfig {
	if in == nil {
		return nil
	}
	out := new(AKORbacConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AVITenant) DeepCopyInto(out *AVITenant) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AVITenant.
func (in *AVITenant) DeepCopy() *AVITenant {
	if in == nil {
		return nil
	}
	out := new(AVITenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneNetwork) DeepCopyInto(out *ControlPlaneNetwork) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneNetwork.
func (in *ControlPlaneNetwork) DeepCopy() *ControlPlaneNetwork {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNetwork) DeepCopyInto(out *DataNetwork) {
	*out = *in
	if in.IPPools != nil {
		in, out := &in.IPPools, &out.IPPools
		*out = make([]IPPool, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNetwork.
func (in *DataNetwork) DeepCopy() *DataNetwork {
	if in == nil {
		return nil
	}
	out := new(DataNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtraConfigs) DeepCopyInto(out *ExtraConfigs) {
	*out = *in
	out.Log = in.Log
	if in.ApiServerPort != nil {
		in, out := &in.ApiServerPort, &out.ApiServerPort
		*out = new(int)
		**out = **in
	}
	if in.DisableStaticRouteSync != nil {
		in, out := &in.DisableStaticRouteSync, &out.DisableStaticRouteSync
		*out = new(bool)
		**out = **in
	}
	if in.EnableEVH != nil {
		in, out := &in.EnableEVH, &out.EnableEVH
		*out = new(bool)
		**out = **in
	}
	if in.Layer7Only != nil {
		in, out := &in.Layer7Only, &out.Layer7Only
		*out = new(bool)
		**out = **in
	}
	out.NamespaceSelector = in.NamespaceSelector
	if in.ServicesAPI != nil {
		in, out := &in.ServicesAPI, &out.ServicesAPI
		*out = new(bool)
		**out = **in
	}
	if in.IstioEnabled != nil {
		in, out := &in.IstioEnabled, &out.IstioEnabled
		*out = new(bool)
		**out = **in
	}
	if in.VIPPerNamespace != nil {
		in, out := &in.VIPPerNamespace, &out.VIPPerNamespace
		*out = new(bool)
		**out = **in
	}
	in.NetworksConfig.DeepCopyInto(&out.NetworksConfig)
	in.IngressConfigs.DeepCopyInto(&out.IngressConfigs)
	in.L4Configs.DeepCopyInto(&out.L4Configs)
	out.NodePortSelector = in.NodePortSelector
	in.Rbac.DeepCopyInto(&out.Rbac)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraConfigs.
func (in *ExtraConfigs) DeepCopy() *ExtraConfigs {
	if in == nil {
		return nil
	}
	out := new(ExtraConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPool) DeepCopyInto(out *IPPool) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPool.
func (in *IPPool) DeepCopy() *IPPool {
	if in == nil {
		return nil
	}
	out := new(IPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSelector) DeepCopyInto(out *NamespaceSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSelector.
func (in *NamespaceSelector) DeepCopy() *NamespaceSelector {
	if in == nil {
		return nil
	}
	out := new(NamespaceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworksConfig) DeepCopyInto(out *NetworksConfig) {
	*out = *in
	if in.EnableRHI != nil {
		in, out := &in.EnableRHI, &out.EnableRHI
		*out = new(bool)
		**out = **in
	}
	if in.BGPPeerLabels != nil {
		in, out := &in.BGPPeerLabels, &out.BGPPeerLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworksConfig.
func (in *NetworksConfig) DeepCopy() *NetworksConfig {
	if in == nil {
		return nil
	}
	out := new(NetworksConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeNetwork) DeepCopyInto(out *NodeNetwork) {
	*out = *in
	if in.Cidrs != nil {
		in, out := &in.Cidrs, &out.Cidrs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeNetwork.
func (in *NodeNetwork) DeepCopy() *NodeNetwork {
	if in == nil {
		return nil
	}
	out := new(NodeNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePortSelector) DeepCopyInto(out *NodePortSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePortSelector.
func (in *NodePortSelector) DeepCopy() *NodePortSelector {
	if in == nil {
		return nil
	}
	out := new(NodePortSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VIPNetwork) DeepCopyInto(out *VIPNetwork) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VIPNetwork.
func (in *VIPNetwork) DeepCopy() *VIPNetwork {
	if in == nil {
		return nil
	}
	out := new(VIPNetwork)
	in.DeepCopyInto(out)
	return out
}
