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

type AppEngineServiceIAMMemberObservation struct {
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type AppEngineServiceIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

// AppEngineServiceIAMMemberSpec defines the desired state of AppEngineServiceIAMMember
type AppEngineServiceIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppEngineServiceIAMMemberParameters `json:"forProvider"`
}

// AppEngineServiceIAMMemberStatus defines the observed state of AppEngineServiceIAMMember.
type AppEngineServiceIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppEngineServiceIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppEngineServiceIAMMember is the Schema for the AppEngineServiceIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AppEngineServiceIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.appId)",message="appId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member)",message="member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role)",message="role is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.service)",message="service is a required parameter"
	Spec   AppEngineServiceIAMMemberSpec   `json:"spec"`
	Status AppEngineServiceIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppEngineServiceIAMMemberList contains a list of AppEngineServiceIAMMembers
type AppEngineServiceIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppEngineServiceIAMMember `json:"items"`
}

// Repository type metadata.
var (
	AppEngineServiceIAMMember_Kind             = "AppEngineServiceIAMMember"
	AppEngineServiceIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppEngineServiceIAMMember_Kind}.String()
	AppEngineServiceIAMMember_KindAPIVersion   = AppEngineServiceIAMMember_Kind + "." + CRDGroupVersion.String()
	AppEngineServiceIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(AppEngineServiceIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&AppEngineServiceIAMMember{}, &AppEngineServiceIAMMemberList{})
}
