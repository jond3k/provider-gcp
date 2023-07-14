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

type TableIAMPolicyObservation struct {

	// (Computed) The etag of the tables's IAM policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or relative resource id of the instance that owns the table.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// The policy data generated by a google_iam_policy data source.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The project in which the table belongs.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name or relative resource id of the table to manage IAM policies for.
	Table *string `json:"table,omitempty" tf:"table,omitempty"`
}

type TableIAMPolicyParameters struct {

	// The name or relative resource id of the instance that owns the table.
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// The policy data generated by a google_iam_policy data source.
	// +kubebuilder:validation:Optional
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The project in which the table belongs.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name or relative resource id of the table to manage IAM policies for.
	// +kubebuilder:validation:Required
	Table *string `json:"table" tf:"table,omitempty"`
}

// TableIAMPolicySpec defines the desired state of TableIAMPolicy
type TableIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableIAMPolicyParameters `json:"forProvider"`
}

// TableIAMPolicyStatus defines the observed state of TableIAMPolicy.
type TableIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMPolicy is the Schema for the TableIAMPolicys API. Collection of resources to manage IAM policy for a Bigtable Table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TableIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instance)",message="instance is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyData)",message="policyData is a required parameter"
	Spec   TableIAMPolicySpec   `json:"spec"`
	Status TableIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMPolicyList contains a list of TableIAMPolicys
type TableIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	TableIAMPolicy_Kind             = "TableIAMPolicy"
	TableIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableIAMPolicy_Kind}.String()
	TableIAMPolicy_KindAPIVersion   = TableIAMPolicy_Kind + "." + CRDGroupVersion.String()
	TableIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(TableIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&TableIAMPolicy{}, &TableIAMPolicyList{})
}
