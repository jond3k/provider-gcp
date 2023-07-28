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

type TensorboardEncryptionSpecInitParameters struct {

	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type TensorboardEncryptionSpecObservation struct {

	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type TensorboardEncryptionSpecParameters struct {

	// The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

type TensorboardInitParameters struct {

	// Description of this Tensorboard.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
	// Structure is documented below.
	EncryptionSpec []TensorboardEncryptionSpecInitParameters `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// The labels with user-defined metadata to organize your Tensorboards.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type TensorboardObservation struct {

	// Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not end with a '/'.
	BlobStoragePathPrefix *string `json:"blobStoragePathPrefix,omitempty" tf:"blob_storage_path_prefix,omitempty"`

	// The timestamp of when the Tensorboard was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Description of this Tensorboard.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
	// Structure is documented below.
	EncryptionSpec []TensorboardEncryptionSpecObservation `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The labels with user-defined metadata to organize your Tensorboards.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the Tensorboard.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the tensorboard. eg us-central1
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The number of Runs stored in this Tensorboard.
	RunCount *string `json:"runCount,omitempty" tf:"run_count,omitempty"`

	// The timestamp of when the Tensorboard was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type TensorboardParameters struct {

	// Description of this Tensorboard.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this Tensorboard will be secured by this key.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EncryptionSpec []TensorboardEncryptionSpecParameters `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// The labels with user-defined metadata to organize your Tensorboards.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the tensorboard. eg us-central1
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

// TensorboardSpec defines the desired state of Tensorboard
type TensorboardSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TensorboardParameters `json:"forProvider"`
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
	InitProvider TensorboardInitParameters `json:"initProvider,omitempty"`
}

// TensorboardStatus defines the observed state of Tensorboard.
type TensorboardStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TensorboardObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Tensorboard is the Schema for the Tensorboards API. Tensorboard is a physical database that stores users' training metrics.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Tensorboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TensorboardSpec   `json:"spec"`
	Status            TensorboardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TensorboardList contains a list of Tensorboards
type TensorboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tensorboard `json:"items"`
}

// Repository type metadata.
var (
	Tensorboard_Kind             = "Tensorboard"
	Tensorboard_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tensorboard_Kind}.String()
	Tensorboard_KindAPIVersion   = Tensorboard_Kind + "." + CRDGroupVersion.String()
	Tensorboard_GroupVersionKind = CRDGroupVersion.WithKind(Tensorboard_Kind)
)

func init() {
	SchemeBuilder.Register(&Tensorboard{}, &TensorboardList{})
}
