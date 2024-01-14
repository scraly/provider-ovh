// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AdmissionpluginsInitParameters struct {
	Disabled []*string `json:"disabled,omitempty" tf:"disabled,omitempty"`

	Enabled []*string `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AdmissionpluginsObservation struct {
	Disabled []*string `json:"disabled,omitempty" tf:"disabled,omitempty"`

	Enabled []*string `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type AdmissionpluginsParameters struct {

	// +kubebuilder:validation:Optional
	Disabled []*string `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled []*string `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ApiserverInitParameters struct {
	Admissionplugins []AdmissionpluginsInitParameters `json:"admissionplugins,omitempty" tf:"admissionplugins,omitempty"`
}

type ApiserverObservation struct {
	Admissionplugins []AdmissionpluginsObservation `json:"admissionplugins,omitempty" tf:"admissionplugins,omitempty"`
}

type ApiserverParameters struct {

	// +kubebuilder:validation:Optional
	Admissionplugins []AdmissionpluginsParameters `json:"admissionplugins,omitempty" tf:"admissionplugins,omitempty"`
}

type ClusterInitParameters struct {
	Customization []CustomizationInitParameters `json:"customization,omitempty" tf:"customization,omitempty"`

	CustomizationApiserver []CustomizationApiserverInitParameters `json:"customizationApiserver,omitempty" tf:"customization_apiserver,omitempty"`

	CustomizationKubeProxy []CustomizationKubeProxyInitParameters `json:"customizationKubeProxy,omitempty" tf:"customization_kube_proxy,omitempty"`

	KubeProxyMode *string `json:"kubeProxyMode,omitempty" tf:"kube_proxy_mode,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	PrivateNetworkConfiguration []PrivateNetworkConfigurationInitParameters `json:"privateNetworkConfiguration,omitempty" tf:"private_network_configuration,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	UpdatePolicy *string `json:"updatePolicy,omitempty" tf:"update_policy,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ClusterObservation struct {
	ControlPlaneIsUpToDate *bool `json:"controlPlaneIsUpToDate,omitempty" tf:"control_plane_is_up_to_date,omitempty"`

	Customization []CustomizationObservation `json:"customization,omitempty" tf:"customization,omitempty"`

	CustomizationApiserver []CustomizationApiserverObservation `json:"customizationApiserver,omitempty" tf:"customization_apiserver,omitempty"`

	CustomizationKubeProxy []CustomizationKubeProxyObservation `json:"customizationKubeProxy,omitempty" tf:"customization_kube_proxy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IsUpToDate *bool `json:"isUpToDate,omitempty" tf:"is_up_to_date,omitempty"`

	KubeProxyMode *string `json:"kubeProxyMode,omitempty" tf:"kube_proxy_mode,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NextUpgradeVersions []*string `json:"nextUpgradeVersions,omitempty" tf:"next_upgrade_versions,omitempty"`

	NodesURL *string `json:"nodesUrl,omitempty" tf:"nodes_url,omitempty"`

	PrivateNetworkConfiguration []PrivateNetworkConfigurationObservation `json:"privateNetworkConfiguration,omitempty" tf:"private_network_configuration,omitempty"`

	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`

	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	UpdatePolicy *string `json:"updatePolicy,omitempty" tf:"update_policy,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	Customization []CustomizationParameters `json:"customization,omitempty" tf:"customization,omitempty"`

	// +kubebuilder:validation:Optional
	CustomizationApiserver []CustomizationApiserverParameters `json:"customizationApiserver,omitempty" tf:"customization_apiserver,omitempty"`

	// +kubebuilder:validation:Optional
	CustomizationKubeProxy []CustomizationKubeProxyParameters `json:"customizationKubeProxy,omitempty" tf:"customization_kube_proxy,omitempty"`

	// +kubebuilder:validation:Optional
	KubeProxyMode *string `json:"kubeProxyMode,omitempty" tf:"kube_proxy_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateNetworkConfiguration []PrivateNetworkConfigurationParameters `json:"privateNetworkConfiguration,omitempty" tf:"private_network_configuration,omitempty"`

	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/network/v1alpha1.PrivateNetwork
	// +kubebuilder:validation:Optional
	PrivateNetworkID *string `json:"privateNetworkId,omitempty" tf:"private_network_id,omitempty"`

	// Reference to a PrivateNetwork in network to populate privateNetworkId.
	// +kubebuilder:validation:Optional
	PrivateNetworkIDRef *v1.Reference `json:"privateNetworkIdRef,omitempty" tf:"-"`

	// Selector for a PrivateNetwork in network to populate privateNetworkId.
	// +kubebuilder:validation:Optional
	PrivateNetworkIDSelector *v1.Selector `json:"privateNetworkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	UpdatePolicy *string `json:"updatePolicy,omitempty" tf:"update_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type CustomizationApiserverAdmissionpluginsInitParameters struct {
	Disabled []*string `json:"disabled,omitempty" tf:"disabled,omitempty"`

	Enabled []*string `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type CustomizationApiserverAdmissionpluginsObservation struct {
	Disabled []*string `json:"disabled,omitempty" tf:"disabled,omitempty"`

	Enabled []*string `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type CustomizationApiserverAdmissionpluginsParameters struct {

	// +kubebuilder:validation:Optional
	Disabled []*string `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled []*string `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type CustomizationApiserverInitParameters struct {
	Admissionplugins []CustomizationApiserverAdmissionpluginsInitParameters `json:"admissionplugins,omitempty" tf:"admissionplugins,omitempty"`
}

type CustomizationApiserverObservation struct {
	Admissionplugins []CustomizationApiserverAdmissionpluginsObservation `json:"admissionplugins,omitempty" tf:"admissionplugins,omitempty"`
}

type CustomizationApiserverParameters struct {

	// +kubebuilder:validation:Optional
	Admissionplugins []CustomizationApiserverAdmissionpluginsParameters `json:"admissionplugins,omitempty" tf:"admissionplugins,omitempty"`
}

type CustomizationInitParameters struct {
	Apiserver []ApiserverInitParameters `json:"apiserver,omitempty" tf:"apiserver,omitempty"`
}

type CustomizationKubeProxyInitParameters struct {
	Iptables []IptablesInitParameters `json:"iptables,omitempty" tf:"iptables,omitempty"`

	Ipvs []IpvsInitParameters `json:"ipvs,omitempty" tf:"ipvs,omitempty"`
}

type CustomizationKubeProxyObservation struct {
	Iptables []IptablesObservation `json:"iptables,omitempty" tf:"iptables,omitempty"`

	Ipvs []IpvsObservation `json:"ipvs,omitempty" tf:"ipvs,omitempty"`
}

type CustomizationKubeProxyParameters struct {

	// +kubebuilder:validation:Optional
	Iptables []IptablesParameters `json:"iptables,omitempty" tf:"iptables,omitempty"`

	// +kubebuilder:validation:Optional
	Ipvs []IpvsParameters `json:"ipvs,omitempty" tf:"ipvs,omitempty"`
}

type CustomizationObservation struct {
	Apiserver []ApiserverObservation `json:"apiserver,omitempty" tf:"apiserver,omitempty"`
}

type CustomizationParameters struct {

	// +kubebuilder:validation:Optional
	Apiserver []ApiserverParameters `json:"apiserver,omitempty" tf:"apiserver,omitempty"`
}

type IptablesInitParameters struct {
	MinSyncPeriod *string `json:"minSyncPeriod,omitempty" tf:"min_sync_period,omitempty"`

	SyncPeriod *string `json:"syncPeriod,omitempty" tf:"sync_period,omitempty"`
}

type IptablesObservation struct {
	MinSyncPeriod *string `json:"minSyncPeriod,omitempty" tf:"min_sync_period,omitempty"`

	SyncPeriod *string `json:"syncPeriod,omitempty" tf:"sync_period,omitempty"`
}

type IptablesParameters struct {

	// +kubebuilder:validation:Optional
	MinSyncPeriod *string `json:"minSyncPeriod,omitempty" tf:"min_sync_period,omitempty"`

	// +kubebuilder:validation:Optional
	SyncPeriod *string `json:"syncPeriod,omitempty" tf:"sync_period,omitempty"`
}

type IpvsInitParameters struct {
	MinSyncPeriod *string `json:"minSyncPeriod,omitempty" tf:"min_sync_period,omitempty"`

	Scheduler *string `json:"scheduler,omitempty" tf:"scheduler,omitempty"`

	SyncPeriod *string `json:"syncPeriod,omitempty" tf:"sync_period,omitempty"`

	TCPFinTimeout *string `json:"tcpFinTimeout,omitempty" tf:"tcp_fin_timeout,omitempty"`

	TCPTimeout *string `json:"tcpTimeout,omitempty" tf:"tcp_timeout,omitempty"`

	UDPTimeout *string `json:"udpTimeout,omitempty" tf:"udp_timeout,omitempty"`
}

type IpvsObservation struct {
	MinSyncPeriod *string `json:"minSyncPeriod,omitempty" tf:"min_sync_period,omitempty"`

	Scheduler *string `json:"scheduler,omitempty" tf:"scheduler,omitempty"`

	SyncPeriod *string `json:"syncPeriod,omitempty" tf:"sync_period,omitempty"`

	TCPFinTimeout *string `json:"tcpFinTimeout,omitempty" tf:"tcp_fin_timeout,omitempty"`

	TCPTimeout *string `json:"tcpTimeout,omitempty" tf:"tcp_timeout,omitempty"`

	UDPTimeout *string `json:"udpTimeout,omitempty" tf:"udp_timeout,omitempty"`
}

type IpvsParameters struct {

	// +kubebuilder:validation:Optional
	MinSyncPeriod *string `json:"minSyncPeriod,omitempty" tf:"min_sync_period,omitempty"`

	// +kubebuilder:validation:Optional
	Scheduler *string `json:"scheduler,omitempty" tf:"scheduler,omitempty"`

	// +kubebuilder:validation:Optional
	SyncPeriod *string `json:"syncPeriod,omitempty" tf:"sync_period,omitempty"`

	// +kubebuilder:validation:Optional
	TCPFinTimeout *string `json:"tcpFinTimeout,omitempty" tf:"tcp_fin_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	TCPTimeout *string `json:"tcpTimeout,omitempty" tf:"tcp_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	UDPTimeout *string `json:"udpTimeout,omitempty" tf:"udp_timeout,omitempty"`
}

type KubeconfigAttributesInitParameters struct {
}

type KubeconfigAttributesObservation struct {
	ClientCertificate *string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`

	ClientKey *string `json:"clientKey,omitempty" tf:"client_key,omitempty"`

	ClusterCACertificate *string `json:"clusterCaCertificate,omitempty" tf:"cluster_ca_certificate,omitempty"`

	Host *string `json:"host,omitempty" tf:"host,omitempty"`
}

type KubeconfigAttributesParameters struct {
}

type PrivateNetworkConfigurationInitParameters struct {

	// If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
	DefaultVrackGateway *string `json:"defaultVrackGateway,omitempty" tf:"default_vrack_gateway,omitempty"`

	// Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
	PrivateNetworkRoutingAsDefault *bool `json:"privateNetworkRoutingAsDefault,omitempty" tf:"private_network_routing_as_default,omitempty"`
}

type PrivateNetworkConfigurationObservation struct {

	// If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
	DefaultVrackGateway *string `json:"defaultVrackGateway,omitempty" tf:"default_vrack_gateway,omitempty"`

	// Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
	PrivateNetworkRoutingAsDefault *bool `json:"privateNetworkRoutingAsDefault,omitempty" tf:"private_network_routing_as_default,omitempty"`
}

type PrivateNetworkConfigurationParameters struct {

	// If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
	// +kubebuilder:validation:Optional
	DefaultVrackGateway *string `json:"defaultVrackGateway" tf:"default_vrack_gateway,omitempty"`

	// Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
	// +kubebuilder:validation:Optional
	PrivateNetworkRoutingAsDefault *bool `json:"privateNetworkRoutingAsDefault" tf:"private_network_routing_as_default,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
