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

type ConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type CryptoKeyIAMMemberInitParameters struct {
	Condition []ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type CryptoKeyIAMMemberObservation struct {
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	CryptoKeyID *string `json:"cryptoKeyId,omitempty" tf:"crypto_key_id,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type CryptoKeyIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +crossplane:generate:reference:type=CryptoKey
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CryptoKeyID *string `json:"cryptoKeyId,omitempty" tf:"crypto_key_id,omitempty"`

	// Reference to a CryptoKey to populate cryptoKeyId.
	// +kubebuilder:validation:Optional
	CryptoKeyIDRef *v1.Reference `json:"cryptoKeyIdRef,omitempty" tf:"-"`

	// Selector for a CryptoKey to populate cryptoKeyId.
	// +kubebuilder:validation:Optional
	CryptoKeyIDSelector *v1.Selector `json:"cryptoKeyIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

// CryptoKeyIAMMemberSpec defines the desired state of CryptoKeyIAMMember
type CryptoKeyIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CryptoKeyIAMMemberParameters `json:"forProvider"`
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
	InitProvider CryptoKeyIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// CryptoKeyIAMMemberStatus defines the observed state of CryptoKeyIAMMember.
type CryptoKeyIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CryptoKeyIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CryptoKeyIAMMember is the Schema for the CryptoKeyIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type CryptoKeyIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member) || has(self.initProvider.member)",message="member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || has(self.initProvider.role)",message="role is a required parameter"
	Spec   CryptoKeyIAMMemberSpec   `json:"spec"`
	Status CryptoKeyIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CryptoKeyIAMMemberList contains a list of CryptoKeyIAMMembers
type CryptoKeyIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CryptoKeyIAMMember `json:"items"`
}

// Repository type metadata.
var (
	CryptoKeyIAMMember_Kind             = "CryptoKeyIAMMember"
	CryptoKeyIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CryptoKeyIAMMember_Kind}.String()
	CryptoKeyIAMMember_KindAPIVersion   = CryptoKeyIAMMember_Kind + "." + CRDGroupVersion.String()
	CryptoKeyIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(CryptoKeyIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&CryptoKeyIAMMember{}, &CryptoKeyIAMMemberList{})
}
