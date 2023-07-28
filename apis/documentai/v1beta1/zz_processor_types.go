/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProcessorInitParameters struct {

	// The display name. Must be unique.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// The location of the resource.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The type of processor. For possible types see the official list
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProcessorObservation struct {

	// The display name. Must be unique.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/processors/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// The location of the resource.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The resource name of the processor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The type of processor. For possible types see the official list
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ProcessorParameters struct {

	// The display name. Must be unique.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// The location of the resource.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The type of processor. For possible types see the official list
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ProcessorSpec defines the desired state of Processor
type ProcessorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProcessorParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ProcessorInitParameters `json:"initProvider,omitempty"`
}

// ProcessorStatus defines the observed state of Processor.
type ProcessorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProcessorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Processor is the Schema for the Processors API. The first-class citizen for Document AI.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Processor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || has(self.initProvider.displayName)",message="displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || has(self.initProvider.type)",message="type is a required parameter"
	Spec   ProcessorSpec   `json:"spec"`
	Status ProcessorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProcessorList contains a list of Processors
type ProcessorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Processor `json:"items"`
}

// Repository type metadata.
var (
	Processor_Kind             = "Processor"
	Processor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Processor_Kind}.String()
	Processor_KindAPIVersion   = Processor_Kind + "." + CRDGroupVersion.String()
	Processor_GroupVersionKind = CRDGroupVersion.WithKind(Processor_Kind)
)

func init() {
	SchemeBuilder.Register(&Processor{}, &ProcessorList{})
}
