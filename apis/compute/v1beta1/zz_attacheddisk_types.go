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

type AttachedDiskInitParameters struct {

	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The project that the referenced compute instance is a part of. If instance is referenced by its
	// self_link the project defined in the link will take precedence.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The zone that the referenced compute instance is located within. If instance is referenced by its
	// self_link the zone defined in the link will take precedence.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type AttachedDiskObservation struct {

	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// name or self_link of the disk that will be attached.
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// an identifier for the resource with format projects/{{project}}/zones/{{zone}}/disks/{{disk.name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// name or self_link of the compute instance that the disk will be attached to.
	// If the self_link is provided then zone and project are extracted from the
	// self link. If only the name is used then zone and project must be defined
	// as properties on the resource or provider.
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The project that the referenced compute instance is a part of. If instance is referenced by its
	// self_link the project defined in the link will take precedence.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The zone that the referenced compute instance is located within. If instance is referenced by its
	// self_link the zone defined in the link will take precedence.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type AttachedDiskParameters struct {

	// Specifies a unique device name of your choice that is
	// reflected into the /dev/disk/by-id/google-* tree of a Linux operating
	// system running within the instance. This name can be used to
	// reference the device for mounting, resizing, and so on, from within
	// the instance.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// name or self_link of the disk that will be attached.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Disk
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Disk *string `json:"disk,omitempty" tf:"disk,omitempty"`

	// Reference to a Disk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskRef *v1.Reference `json:"diskRef,omitempty" tf:"-"`

	// Selector for a Disk in compute to populate disk.
	// +kubebuilder:validation:Optional
	DiskSelector *v1.Selector `json:"diskSelector,omitempty" tf:"-"`

	// name or self_link of the compute instance that the disk will be attached to.
	// If the self_link is provided then zone and project are extracted from the
	// self link. If only the name is used then zone and project must be defined
	// as properties on the resource or provider.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// Reference to a Instance in compute to populate instance.
	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// Selector for a Instance in compute to populate instance.
	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`

	// The mode in which to attach this disk, either READ_WRITE or
	// READ_ONLY. If not specified, the default is to attach the disk in
	// READ_WRITE mode.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The project that the referenced compute instance is a part of. If instance is referenced by its
	// self_link the project defined in the link will take precedence.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The zone that the referenced compute instance is located within. If instance is referenced by its
	// self_link the zone defined in the link will take precedence.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// AttachedDiskSpec defines the desired state of AttachedDisk
type AttachedDiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttachedDiskParameters `json:"forProvider"`
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
	InitProvider AttachedDiskInitParameters `json:"initProvider,omitempty"`
}

// AttachedDiskStatus defines the observed state of AttachedDisk.
type AttachedDiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttachedDiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AttachedDisk is the Schema for the AttachedDisks API. Resource that allows attaching existing persistent disks to compute instances.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AttachedDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AttachedDiskSpec   `json:"spec"`
	Status            AttachedDiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachedDiskList contains a list of AttachedDisks
type AttachedDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AttachedDisk `json:"items"`
}

// Repository type metadata.
var (
	AttachedDisk_Kind             = "AttachedDisk"
	AttachedDisk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AttachedDisk_Kind}.String()
	AttachedDisk_KindAPIVersion   = AttachedDisk_Kind + "." + CRDGroupVersion.String()
	AttachedDisk_GroupVersionKind = CRDGroupVersion.WithKind(AttachedDisk_Kind)
)

func init() {
	SchemeBuilder.Register(&AttachedDisk{}, &AttachedDiskList{})
}
