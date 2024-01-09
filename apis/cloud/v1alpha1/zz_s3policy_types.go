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

type S3PolicyInitParameters struct {

	// The policy document. This is a JSON formatted string.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Service name of the resource representing the ID of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type S3PolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The policy document. This is a JSON formatted string.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Service name of the resource representing the ID of the cloud project.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The user ID
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type S3PolicyParameters struct {

	// The policy document. This is a JSON formatted string.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Service name of the resource representing the ID of the cloud project.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The user ID
	// +crossplane:generate:reference:type=github.com/edixos/provider-ovh/apis/cloud/v1alpha1.User
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in cloud to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in cloud to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

// S3PolicySpec defines the desired state of S3Policy
type S3PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     S3PolicyParameters `json:"forProvider"`
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
	InitProvider S3PolicyInitParameters `json:"initProvider,omitempty"`
}

// S3PolicyStatus defines the observed state of S3Policy.
type S3PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        S3PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3Policy is the Schema for the S3Policys API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type S3Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   S3PolicySpec   `json:"spec"`
	Status S3PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3PolicyList contains a list of S3Policys
type S3PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3Policy `json:"items"`
}

// Repository type metadata.
var (
	S3Policy_Kind             = "S3Policy"
	S3Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: S3Policy_Kind}.String()
	S3Policy_KindAPIVersion   = S3Policy_Kind + "." + CRDGroupVersion.String()
	S3Policy_GroupVersionKind = CRDGroupVersion.WithKind(S3Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&S3Policy{}, &S3PolicyList{})
}
