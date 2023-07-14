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

type GlobalNetworkEndpointObservation struct {

	// Fully qualified domain name of network endpoint.
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The global network endpoint group this endpoint is part of.
	GlobalNetworkEndpointGroup *string `json:"globalNetworkEndpointGroup,omitempty" tf:"global_network_endpoint_group,omitempty"`

	// an identifier for the resource with format {{project}}/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IPv4 address external endpoint.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Port number of the external endpoint.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type GlobalNetworkEndpointParameters struct {

	// Fully qualified domain name of network endpoint.
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The global network endpoint group this endpoint is part of.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.GlobalNetworkEndpointGroup
	// +kubebuilder:validation:Optional
	GlobalNetworkEndpointGroup *string `json:"globalNetworkEndpointGroup,omitempty" tf:"global_network_endpoint_group,omitempty"`

	// Reference to a GlobalNetworkEndpointGroup in compute to populate globalNetworkEndpointGroup.
	// +kubebuilder:validation:Optional
	GlobalNetworkEndpointGroupRef *v1.Reference `json:"globalNetworkEndpointGroupRef,omitempty" tf:"-"`

	// Selector for a GlobalNetworkEndpointGroup in compute to populate globalNetworkEndpointGroup.
	// +kubebuilder:validation:Optional
	GlobalNetworkEndpointGroupSelector *v1.Selector `json:"globalNetworkEndpointGroupSelector,omitempty" tf:"-"`

	// IPv4 address external endpoint.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Port number of the external endpoint.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// GlobalNetworkEndpointSpec defines the desired state of GlobalNetworkEndpoint
type GlobalNetworkEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GlobalNetworkEndpointParameters `json:"forProvider"`
}

// GlobalNetworkEndpointStatus defines the observed state of GlobalNetworkEndpoint.
type GlobalNetworkEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GlobalNetworkEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalNetworkEndpoint is the Schema for the GlobalNetworkEndpoints API. A Global Network endpoint represents a IP address and port combination that exists outside of GCP.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type GlobalNetworkEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.port)",message="port is a required parameter"
	Spec   GlobalNetworkEndpointSpec   `json:"spec"`
	Status GlobalNetworkEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlobalNetworkEndpointList contains a list of GlobalNetworkEndpoints
type GlobalNetworkEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalNetworkEndpoint `json:"items"`
}

// Repository type metadata.
var (
	GlobalNetworkEndpoint_Kind             = "GlobalNetworkEndpoint"
	GlobalNetworkEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GlobalNetworkEndpoint_Kind}.String()
	GlobalNetworkEndpoint_KindAPIVersion   = GlobalNetworkEndpoint_Kind + "." + CRDGroupVersion.String()
	GlobalNetworkEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(GlobalNetworkEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&GlobalNetworkEndpoint{}, &GlobalNetworkEndpointList{})
}
