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

type GlobalAddressInitParameters struct {

	// The IP address or beginning of the address range represented by this
	// resource. This can be supplied as an input to reserve a specific
	// address or omitted to allow GCP to choose a valid one for you.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of the address to reserve.
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP Version that will be used by this address. The default value is IPV4.
	// Possible values are: IPV4, IPV6.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The prefix length of the IP range. If not present, it means the
	// address field is a single IP address.
	// This field is not applicable to addresses with addressType=EXTERNAL,
	// or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of the resource. Possible values include:
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`
}

type GlobalAddressObservation struct {

	// The IP address or beginning of the address range represented by this
	// resource. This can be supplied as an input to reserve a specific
	// address or omitted to allow GCP to choose a valid one for you.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of the address to reserve.
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/addresses/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP Version that will be used by this address. The default value is IPV4.
	// Possible values are: IPV4, IPV6.
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The URL of the network in which to reserve the IP range. The IP range
	// must be in RFC1918 space. The network cannot be deleted if there are
	// any reserved IP ranges referring to it.
	// This should only be set when using an Internal address.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The prefix length of the IP range. If not present, it means the
	// address field is a single IP address.
	// This field is not applicable to addresses with addressType=EXTERNAL,
	// or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of the resource. Possible values include:
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type GlobalAddressParameters struct {

	// The IP address or beginning of the address range represented by this
	// resource. This can be supplied as an input to reserve a specific
	// address or omitted to allow GCP to choose a valid one for you.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of the address to reserve.
	// +kubebuilder:validation:Optional
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The IP Version that will be used by this address. The default value is IPV4.
	// Possible values are: IPV4, IPV6.
	// +kubebuilder:validation:Optional
	IPVersion *string `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The URL of the network in which to reserve the IP range. The IP range
	// must be in RFC1918 space. The network cannot be deleted if there are
	// any reserved IP ranges referring to it.
	// This should only be set when using an Internal address.
	// +crossplane:generate:reference:type=Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The prefix length of the IP range. If not present, it means the
	// address field is a single IP address.
	// This field is not applicable to addresses with addressType=EXTERNAL,
	// or addressType=INTERNAL when purpose=PRIVATE_SERVICE_CONNECT
	// +kubebuilder:validation:Optional
	PrefixLength *float64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of the resource. Possible values include:
	// +kubebuilder:validation:Optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`
}

// GlobalAddressSpec defines the desired state of GlobalAddress
type GlobalAddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalAddressParameters `json:"forProvider"`
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
	InitProvider GlobalAddressInitParameters `json:"initProvider,omitempty"`
}

// GlobalAddressStatus defines the observed state of GlobalAddress.
type GlobalAddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalAddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalAddress is the Schema for the GlobalAddresss API. Represents a Global Address resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type GlobalAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalAddressSpec   `json:"spec"`
	Status            GlobalAddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalAddressList contains a list of GlobalAddresss
type GlobalAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalAddress `json:"items"`
}

// Repository type metadata.
var (
	GlobalAddress_Kind             = "GlobalAddress"
	GlobalAddress_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalAddress_Kind}.String()
	GlobalAddress_KindAPIVersion   = GlobalAddress_Kind + "." + CRDGroupVersion.String()
	GlobalAddress_GroupVersionKind = CRDGroupVersion.WithKind(GlobalAddress_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalAddress{}, &GlobalAddressList{})
}
