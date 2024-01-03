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

type ProjectDatabaseKafkaACLInitParameters struct {

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Permission to give to this username on this topic
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Topic affected by this acl
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// Username affected by this acl
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ProjectDatabaseKafkaACLObservation struct {

	// Id of the database cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Permission to give to this username on this topic
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Topic affected by this acl
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// Username affected by this acl
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ProjectDatabaseKafkaACLParameters struct {

	// Id of the database cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Permission to give to this username on this topic
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Topic affected by this acl
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// Username affected by this acl
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// ProjectDatabaseKafkaACLSpec defines the desired state of ProjectDatabaseKafkaACL
type ProjectDatabaseKafkaACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDatabaseKafkaACLParameters `json:"forProvider"`
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
	InitProvider ProjectDatabaseKafkaACLInitParameters `json:"initProvider,omitempty"`
}

// ProjectDatabaseKafkaACLStatus defines the observed state of ProjectDatabaseKafkaACL.
type ProjectDatabaseKafkaACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDatabaseKafkaACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDatabaseKafkaACL is the Schema for the ProjectDatabaseKafkaACLs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ProjectDatabaseKafkaACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clusterId) || (has(self.initProvider) && has(self.initProvider.clusterId))",message="spec.forProvider.clusterId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permission) || (has(self.initProvider) && has(self.initProvider.permission))",message="spec.forProvider.permission is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.topic) || (has(self.initProvider) && has(self.initProvider.topic))",message="spec.forProvider.topic is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   ProjectDatabaseKafkaACLSpec   `json:"spec"`
	Status ProjectDatabaseKafkaACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDatabaseKafkaACLList contains a list of ProjectDatabaseKafkaACLs
type ProjectDatabaseKafkaACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDatabaseKafkaACL `json:"items"`
}

// Repository type metadata.
var (
	ProjectDatabaseKafkaACL_Kind             = "ProjectDatabaseKafkaACL"
	ProjectDatabaseKafkaACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDatabaseKafkaACL_Kind}.String()
	ProjectDatabaseKafkaACL_KindAPIVersion   = ProjectDatabaseKafkaACL_Kind + "." + CRDGroupVersion.String()
	ProjectDatabaseKafkaACL_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDatabaseKafkaACL_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDatabaseKafkaACL{}, &ProjectDatabaseKafkaACLList{})
}