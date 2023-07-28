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

type SecretVersionInitParameters struct {

	// The current state of the SecretVersion.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SecretVersionObservation struct {

	// The time at which the Secret was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The time at which the Secret was destroyed. Only present if state is DESTROYED.
	DestroyTime *string `json:"destroyTime,omitempty" tf:"destroy_time,omitempty"`

	// The current state of the SecretVersion.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource name of the SecretVersion. Format:
	// projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Secret Manager secret resource
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// The version of the Secret.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SecretVersionParameters struct {

	// The current state of the SecretVersion.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Secret Manager secret resource
	// +crossplane:generate:reference:type=Secret
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Secret *string `json:"secret,omitempty" tf:"secret,omitempty"`

	// The secret data. Must be no larger than 64KiB.
	// +kubebuilder:validation:Optional
	SecretDataSecretRef v1.SecretKeySelector `json:"secretDataSecretRef" tf:"-"`

	// Reference to a Secret to populate secret.
	// +kubebuilder:validation:Optional
	SecretRef *v1.Reference `json:"secretRef,omitempty" tf:"-"`

	// Selector for a Secret to populate secret.
	// +kubebuilder:validation:Optional
	SecretSelector *v1.Selector `json:"secretSelector,omitempty" tf:"-"`
}

// SecretVersionSpec defines the desired state of SecretVersion
type SecretVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretVersionParameters `json:"forProvider"`
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
	InitProvider SecretVersionInitParameters `json:"initProvider,omitempty"`
}

// SecretVersionStatus defines the observed state of SecretVersion.
type SecretVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretVersion is the Schema for the SecretVersions API. A secret version resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type SecretVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.secretDataSecretRef)",message="secretDataSecretRef is a required parameter"
	Spec   SecretVersionSpec   `json:"spec"`
	Status SecretVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretVersionList contains a list of SecretVersions
type SecretVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretVersion `json:"items"`
}

// Repository type metadata.
var (
	SecretVersion_Kind             = "SecretVersion"
	SecretVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretVersion_Kind}.String()
	SecretVersion_KindAPIVersion   = SecretVersion_Kind + "." + CRDGroupVersion.String()
	SecretVersion_GroupVersionKind = CRDGroupVersion.WithKind(SecretVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretVersion{}, &SecretVersionList{})
}
