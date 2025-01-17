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

type TCPFarmInitParameters struct {
	Balance *string `json:"balance,omitempty" tf:"balance,omitempty"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	Probe []TCPFarmProbeInitParameters `json:"probe,omitempty" tf:"probe,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Stickiness *string `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	VrackNetworkID *float64 `json:"vrackNetworkId,omitempty" tf:"vrack_network_id,omitempty"`

	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type TCPFarmObservation struct {
	Balance *string `json:"balance,omitempty" tf:"balance,omitempty"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	Probe []TCPFarmProbeObservation `json:"probe,omitempty" tf:"probe,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	Stickiness *string `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	VrackNetworkID *float64 `json:"vrackNetworkId,omitempty" tf:"vrack_network_id,omitempty"`

	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type TCPFarmParameters struct {

	// +kubebuilder:validation:Optional
	Balance *string `json:"balance,omitempty" tf:"balance,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Probe []TCPFarmProbeParameters `json:"probe,omitempty" tf:"probe,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// +kubebuilder:validation:Optional
	Stickiness *string `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// +kubebuilder:validation:Optional
	VrackNetworkID *float64 `json:"vrackNetworkId,omitempty" tf:"vrack_network_id,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type TCPFarmProbeInitParameters struct {
	ForceSSL *bool `json:"forceSsl,omitempty" tf:"force_ssl,omitempty"`

	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	Match *string `json:"match,omitempty" tf:"match,omitempty"`

	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type TCPFarmProbeObservation struct {
	ForceSSL *bool `json:"forceSsl,omitempty" tf:"force_ssl,omitempty"`

	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	Match *string `json:"match,omitempty" tf:"match,omitempty"`

	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type TCPFarmProbeParameters struct {

	// +kubebuilder:validation:Optional
	ForceSSL *bool `json:"forceSsl,omitempty" tf:"force_ssl,omitempty"`

	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// +kubebuilder:validation:Optional
	Match *string `json:"match,omitempty" tf:"match,omitempty"`

	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// +kubebuilder:validation:Optional
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// TCPFarmSpec defines the desired state of TCPFarm
type TCPFarmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TCPFarmParameters `json:"forProvider"`
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
	InitProvider TCPFarmInitParameters `json:"initProvider,omitempty"`
}

// TCPFarmStatus defines the observed state of TCPFarm.
type TCPFarmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TCPFarmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TCPFarm is the Schema for the TCPFarms API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type TCPFarm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zone) || (has(self.initProvider) && has(self.initProvider.zone))",message="spec.forProvider.zone is a required parameter"
	Spec   TCPFarmSpec   `json:"spec"`
	Status TCPFarmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TCPFarmList contains a list of TCPFarms
type TCPFarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCPFarm `json:"items"`
}

// Repository type metadata.
var (
	TCPFarm_Kind             = "TCPFarm"
	TCPFarm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TCPFarm_Kind}.String()
	TCPFarm_KindAPIVersion   = TCPFarm_Kind + "." + CRDGroupVersion.String()
	TCPFarm_GroupVersionKind = CRDGroupVersion.WithKind(TCPFarm_Kind)
)

func init() {
	SchemeBuilder.Register(&TCPFarm{}, &TCPFarmList{})
}
