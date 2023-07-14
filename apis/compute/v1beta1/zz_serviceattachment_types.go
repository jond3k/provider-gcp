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

type ConnectedEndpointsObservation struct {

	// (Output)
	// The URL of the consumer forwarding rule.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// (Output)
	// The status of the connection from the consumer forwarding rule to
	// this service attachment.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ConnectedEndpointsParameters struct {
}

type ConsumerAcceptListsObservation struct {

	// The number of consumer forwarding rules the consumer project can
	// create.
	ConnectionLimit *float64 `json:"connectionLimit,omitempty" tf:"connection_limit,omitempty"`

	// A project that is allowed to connect to this service attachment.
	ProjectIDOrNum *string `json:"projectIdOrNum,omitempty" tf:"project_id_or_num,omitempty"`
}

type ConsumerAcceptListsParameters struct {

	// The number of consumer forwarding rules the consumer project can
	// create.
	// +kubebuilder:validation:Required
	ConnectionLimit *float64 `json:"connectionLimit" tf:"connection_limit,omitempty"`

	// A project that is allowed to connect to this service attachment.
	// +kubebuilder:validation:Required
	ProjectIDOrNum *string `json:"projectIdOrNum" tf:"project_id_or_num,omitempty"`
}

type ServiceAttachmentObservation struct {

	// An array of the consumer forwarding rules connected to this service
	// attachment.
	// Structure is documented below.
	ConnectedEndpoints []ConnectedEndpointsObservation `json:"connectedEndpoints,omitempty" tf:"connected_endpoints,omitempty"`

	// The connection preference to use for this service attachment. Valid
	// values include "ACCEPT_AUTOMATIC", "ACCEPT_MANUAL".
	ConnectionPreference *string `json:"connectionPreference,omitempty" tf:"connection_preference,omitempty"`

	// An array of projects that are allowed to connect to this service
	// attachment.
	// Structure is documented below.
	ConsumerAcceptLists []ConsumerAcceptListsObservation `json:"consumerAcceptLists,omitempty" tf:"consumer_accept_lists,omitempty"`

	// An array of projects that are not allowed to connect to this service
	// attachment.
	ConsumerRejectLists []*string `json:"consumerRejectLists,omitempty" tf:"consumer_reject_lists,omitempty"`

	// An optional description of this resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If specified, the domain name will be used during the integration between
	// the PSC connected endpoints and the Cloud DNS. For example, this is a
	// valid domain name: "p.mycompany.com.". Current max number of domain names
	// supported is 1.
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// If true, enable the proxy protocol which is for supplying client TCP/IP
	// address data in TCP connections that traverse proxies on their way to
	// destination servers.
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol,omitempty"`

	// Fingerprint of this resource. This field is used internally during
	// updates of this resource.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An array of subnets that is provided for NAT in this service attachment.
	NATSubnets []*string `json:"natSubnets,omitempty" tf:"nat_subnets,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// URL of the region where the resource resides.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The URL of a forwarding rule that represents the service identified by
	// this service attachment.
	TargetService *string `json:"targetService,omitempty" tf:"target_service,omitempty"`
}

type ServiceAttachmentParameters struct {

	// The connection preference to use for this service attachment. Valid
	// values include "ACCEPT_AUTOMATIC", "ACCEPT_MANUAL".
	// +kubebuilder:validation:Optional
	ConnectionPreference *string `json:"connectionPreference,omitempty" tf:"connection_preference,omitempty"`

	// An array of projects that are allowed to connect to this service
	// attachment.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ConsumerAcceptLists []ConsumerAcceptListsParameters `json:"consumerAcceptLists,omitempty" tf:"consumer_accept_lists,omitempty"`

	// An array of projects that are not allowed to connect to this service
	// attachment.
	// +kubebuilder:validation:Optional
	ConsumerRejectLists []*string `json:"consumerRejectLists,omitempty" tf:"consumer_reject_lists,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If specified, the domain name will be used during the integration between
	// the PSC connected endpoints and the Cloud DNS. For example, this is a
	// valid domain name: "p.mycompany.com.". Current max number of domain names
	// supported is 1.
	// +kubebuilder:validation:Optional
	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`

	// If true, enable the proxy protocol which is for supplying client TCP/IP
	// address data in TCP connections that traverse proxies on their way to
	// destination servers.
	// +kubebuilder:validation:Optional
	EnableProxyProtocol *bool `json:"enableProxyProtocol,omitempty" tf:"enable_proxy_protocol,omitempty"`

	// An array of subnets that is provided for NAT in this service attachment.
	// +crossplane:generate:reference:type=Subnetwork
	// +kubebuilder:validation:Optional
	NATSubnets []*string `json:"natSubnets,omitempty" tf:"nat_subnets,omitempty"`

	// References to Subnetwork to populate natSubnets.
	// +kubebuilder:validation:Optional
	NATSubnetsRefs []v1.Reference `json:"natSubnetsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnetwork to populate natSubnets.
	// +kubebuilder:validation:Optional
	NATSubnetsSelector *v1.Selector `json:"natSubnetsSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// URL of the region where the resource resides.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// The URL of a forwarding rule that represents the service identified by
	// this service attachment.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.ForwardingRule
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetService *string `json:"targetService,omitempty" tf:"target_service,omitempty"`

	// Reference to a ForwardingRule in compute to populate targetService.
	// +kubebuilder:validation:Optional
	TargetServiceRef *v1.Reference `json:"targetServiceRef,omitempty" tf:"-"`

	// Selector for a ForwardingRule in compute to populate targetService.
	// +kubebuilder:validation:Optional
	TargetServiceSelector *v1.Selector `json:"targetServiceSelector,omitempty" tf:"-"`
}

// ServiceAttachmentSpec defines the desired state of ServiceAttachment
type ServiceAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAttachmentParameters `json:"forProvider"`
}

// ServiceAttachmentStatus defines the observed state of ServiceAttachment.
type ServiceAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAttachment is the Schema for the ServiceAttachments API. Represents a ServiceAttachment resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServiceAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionPreference)",message="connectionPreference is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enableProxyProtocol)",message="enableProxyProtocol is a required parameter"
	Spec   ServiceAttachmentSpec   `json:"spec"`
	Status ServiceAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAttachmentList contains a list of ServiceAttachments
type ServiceAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAttachment `json:"items"`
}

// Repository type metadata.
var (
	ServiceAttachment_Kind             = "ServiceAttachment"
	ServiceAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAttachment_Kind}.String()
	ServiceAttachment_KindAPIVersion   = ServiceAttachment_Kind + "." + CRDGroupVersion.String()
	ServiceAttachment_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAttachment{}, &ServiceAttachmentList{})
}
