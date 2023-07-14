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

type FirewallPolicyAssociationObservation struct {

	// The target that the firewall policy is attached to.
	AttachmentTarget *string `json:"attachmentTarget,omitempty" tf:"attachment_target,omitempty"`

	// The firewall policy ID of the association.
	FirewallPolicy *string `json:"firewallPolicy,omitempty" tf:"firewall_policy,omitempty"`

	// an identifier for the resource with format locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name for an association.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The short name of the firewall policy of the association.
	ShortName *string `json:"shortName,omitempty" tf:"short_name,omitempty"`
}

type FirewallPolicyAssociationParameters struct {

	// The target that the firewall policy is attached to.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudplatform/v1beta1.Folder
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",true)
	// +kubebuilder:validation:Optional
	AttachmentTarget *string `json:"attachmentTarget,omitempty" tf:"attachment_target,omitempty"`

	// Reference to a Folder in cloudplatform to populate attachmentTarget.
	// +kubebuilder:validation:Optional
	AttachmentTargetRef *v1.Reference `json:"attachmentTargetRef,omitempty" tf:"-"`

	// Selector for a Folder in cloudplatform to populate attachmentTarget.
	// +kubebuilder:validation:Optional
	AttachmentTargetSelector *v1.Selector `json:"attachmentTargetSelector,omitempty" tf:"-"`

	// The firewall policy ID of the association.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.FirewallPolicy
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	FirewallPolicy *string `json:"firewallPolicy,omitempty" tf:"firewall_policy,omitempty"`

	// Reference to a FirewallPolicy in compute to populate firewallPolicy.
	// +kubebuilder:validation:Optional
	FirewallPolicyRef *v1.Reference `json:"firewallPolicyRef,omitempty" tf:"-"`

	// Selector for a FirewallPolicy in compute to populate firewallPolicy.
	// +kubebuilder:validation:Optional
	FirewallPolicySelector *v1.Selector `json:"firewallPolicySelector,omitempty" tf:"-"`

	// The name for an association.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// FirewallPolicyAssociationSpec defines the desired state of FirewallPolicyAssociation
type FirewallPolicyAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallPolicyAssociationParameters `json:"forProvider"`
}

// FirewallPolicyAssociationStatus defines the observed state of FirewallPolicyAssociation.
type FirewallPolicyAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallPolicyAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyAssociation is the Schema for the FirewallPolicyAssociations API. Applies a hierarchical firewall policy to a target resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type FirewallPolicyAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	Spec   FirewallPolicyAssociationSpec   `json:"spec"`
	Status FirewallPolicyAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallPolicyAssociationList contains a list of FirewallPolicyAssociations
type FirewallPolicyAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallPolicyAssociation `json:"items"`
}

// Repository type metadata.
var (
	FirewallPolicyAssociation_Kind             = "FirewallPolicyAssociation"
	FirewallPolicyAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallPolicyAssociation_Kind}.String()
	FirewallPolicyAssociation_KindAPIVersion   = FirewallPolicyAssociation_Kind + "." + CRDGroupVersion.String()
	FirewallPolicyAssociation_GroupVersionKind = CRDGroupVersion.WithKind(FirewallPolicyAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallPolicyAssociation{}, &FirewallPolicyAssociationList{})
}
