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

type ProjectWorkflowBackupInitParameters struct {
	BackupName *string `json:"backupName,omitempty" tf:"backup_name,omitempty"`

	Cron *string `json:"cron,omitempty" tf:"cron,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	MaxExecutionCount *float64 `json:"maxExecutionCount,omitempty" tf:"max_execution_count,omitempty"`

	// Region name.
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	Rotation *float64 `json:"rotation,omitempty" tf:"rotation,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectWorkflowBackupObservation struct {
	BackupName *string `json:"backupName,omitempty" tf:"backup_name,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	Cron *string `json:"cron,omitempty" tf:"cron,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	MaxExecutionCount *float64 `json:"maxExecutionCount,omitempty" tf:"max_execution_count,omitempty"`

	// Region name.
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	Rotation *float64 `json:"rotation,omitempty" tf:"rotation,omitempty"`

	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type ProjectWorkflowBackupParameters struct {

	// +kubebuilder:validation:Optional
	BackupName *string `json:"backupName,omitempty" tf:"backup_name,omitempty"`

	// +kubebuilder:validation:Optional
	Cron *string `json:"cron,omitempty" tf:"cron,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	MaxExecutionCount *float64 `json:"maxExecutionCount,omitempty" tf:"max_execution_count,omitempty"`

	// Region name.
	// +kubebuilder:validation:Optional
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`

	// +kubebuilder:validation:Optional
	Rotation *float64 `json:"rotation,omitempty" tf:"rotation,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// ProjectWorkflowBackupSpec defines the desired state of ProjectWorkflowBackup
type ProjectWorkflowBackupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectWorkflowBackupParameters `json:"forProvider"`
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
	InitProvider ProjectWorkflowBackupInitParameters `json:"initProvider,omitempty"`
}

// ProjectWorkflowBackupStatus defines the observed state of ProjectWorkflowBackup.
type ProjectWorkflowBackupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectWorkflowBackupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectWorkflowBackup is the Schema for the ProjectWorkflowBackups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ovh}
type ProjectWorkflowBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cron) || (has(self.initProvider) && has(self.initProvider.cron))",message="spec.forProvider.cron is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instanceId) || (has(self.initProvider) && has(self.initProvider.instanceId))",message="spec.forProvider.instanceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.regionName) || (has(self.initProvider) && has(self.initProvider.regionName))",message="spec.forProvider.regionName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rotation) || (has(self.initProvider) && has(self.initProvider.rotation))",message="spec.forProvider.rotation is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   ProjectWorkflowBackupSpec   `json:"spec"`
	Status ProjectWorkflowBackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectWorkflowBackupList contains a list of ProjectWorkflowBackups
type ProjectWorkflowBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectWorkflowBackup `json:"items"`
}

// Repository type metadata.
var (
	ProjectWorkflowBackup_Kind             = "ProjectWorkflowBackup"
	ProjectWorkflowBackup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectWorkflowBackup_Kind}.String()
	ProjectWorkflowBackup_KindAPIVersion   = ProjectWorkflowBackup_Kind + "." + CRDGroupVersion.String()
	ProjectWorkflowBackup_GroupVersionKind = CRDGroupVersion.WithKind(ProjectWorkflowBackup_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectWorkflowBackup{}, &ProjectWorkflowBackupList{})
}
