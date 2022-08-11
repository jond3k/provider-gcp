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

type AllowObservation struct {
}

type AllowParameters struct {

	// An optional list of ports to which this rule applies. This field
	// is only applicable for UDP or TCP protocol. Each entry must be
	// either an integer or a range. If not specified, this rule
	// applies to connections through any port.
	//
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	// +kubebuilder:validation:Optional
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The IP protocol to which this rule applies. The protocol type is
	// required when creating a firewall rule. This value can either be
	// one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp, ipip, all), or the IP protocol number.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type DenyObservation struct {
}

type DenyParameters struct {

	// An optional list of ports to which this rule applies. This field
	// is only applicable for UDP or TCP protocol. Each entry must be
	// either an integer or a range. If not specified, this rule
	// applies to connections through any port.
	//
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	// +kubebuilder:validation:Optional
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The IP protocol to which this rule applies. The protocol type is
	// required when creating a firewall rule. This value can either be
	// one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp, ipip, all), or the IP protocol number.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type FirewallLogConfigObservation struct {
}

type FirewallLogConfigParameters struct {

	// This field denotes whether to include or exclude metadata for firewall logs. Possible values: ["EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA"]
	// +kubebuilder:validation:Required
	Metadata *string `json:"metadata" tf:"metadata,omitempty"`
}

type FirewallObservation struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type FirewallParameters struct {

	// The list of ALLOW rules specified by this firewall. Each rule
	// specifies a protocol and port-range tuple that describes a permitted
	// connection.
	// +kubebuilder:validation:Optional
	Allow []AllowParameters `json:"allow,omitempty" tf:"allow,omitempty"`

	// The list of DENY rules specified by this firewall. Each rule specifies
	// a protocol and port-range tuple that describes a denied connection.
	// +kubebuilder:validation:Optional
	Deny []DenyParameters `json:"deny,omitempty" tf:"deny,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If destination ranges are specified, the firewall will apply only to
	// traffic that has destination IP address in these ranges. These ranges
	// must be expressed in CIDR format. Only IPv4 is supported.
	// +kubebuilder:validation:Optional
	DestinationRanges []*string `json:"destinationRanges,omitempty" tf:"destination_ranges,omitempty"`

	// Direction of traffic to which this firewall applies; default is
	// INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
	// destinationRanges; For EGRESS traffic, it is NOT supported to specify
	// 'source_ranges' OR 'source_tags'. For INGRESS traffic, one of 'source_ranges',
	// 'source_tags' or 'source_service_accounts' is required. Possible values: ["INGRESS", "EGRESS"]
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Denotes whether the firewall rule is disabled, i.e not applied to the
	// network it is associated with. When set to true, the firewall rule is
	// not enforced and the network behaves as if it did not exist. If this
	// is unspecified, the firewall rule will be enabled.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported to Stackdriver.
	// +kubebuilder:validation:Optional
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// This field denotes the logging options for a particular firewall rule.
	// If defined, logging is enabled, and logs will be exported to Cloud Logging.
	// +kubebuilder:validation:Optional
	LogConfig []FirewallLogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The name or self_link of the network to attach this firewall to.
	// +crossplane:generate:reference:type=Network
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// Priority for this rule. This is an integer between 0 and 65535, both
	// inclusive. When not specified, the value assumed is 1000. Relative
	// priorities determine precedence of conflicting rules. Lower value of
	// priority implies higher precedence (eg, a rule with priority 0 has
	// higher precedence than a rule with priority 1). DENY rules take
	// precedence over ALLOW rules having equal priority.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// If source ranges are specified, the firewall will apply only to
	// traffic that has source IP address in these ranges. These ranges must
	// be expressed in CIDR format. One or both of sourceRanges and
	// sourceTags may be set. If both properties are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP that belongs to a tag listed in the sourceTags property. The
	// connection does not need to match both properties for the firewall to
	// apply. Only IPv4 is supported. For INGRESS traffic, one of 'source_ranges',
	// 'source_tags' or 'source_service_accounts' is required.
	// +kubebuilder:validation:Optional
	SourceRanges []*string `json:"sourceRanges,omitempty" tf:"source_ranges,omitempty"`

	// If source service accounts are specified, the firewall will apply only
	// to traffic originating from an instance with a service account in this
	// list. Source service accounts cannot be used to control traffic to an
	// instance's external IP address because service accounts are associated
	// with an instance, not an IP address. sourceRanges can be set at the
	// same time as sourceServiceAccounts. If both are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP belongs to an instance with service account listed in
	// sourceServiceAccount. The connection does not need to match both
	// properties for the firewall to apply. sourceServiceAccounts cannot be
	// used at the same time as sourceTags or targetTags. For INGRESS traffic,
	// one of 'source_ranges', 'source_tags' or 'source_service_accounts' is required.
	// +kubebuilder:validation:Optional
	SourceServiceAccounts []*string `json:"sourceServiceAccounts,omitempty" tf:"source_service_accounts,omitempty"`

	// If source tags are specified, the firewall will apply only to traffic
	// with source IP that belongs to a tag listed in source tags. Source
	// tags cannot be used to control traffic to an instance's external IP
	// address. Because tags are associated with an instance, not an IP
	// address. One or both of sourceRanges and sourceTags may be set. If
	// both properties are set, the firewall will apply to traffic that has
	// source IP address within sourceRanges OR the source IP that belongs to
	// a tag listed in the sourceTags property. The connection does not need
	// to match both properties for the firewall to apply. For INGRESS traffic,
	// one of 'source_ranges', 'source_tags' or 'source_service_accounts' is required.
	// +kubebuilder:validation:Optional
	SourceTags []*string `json:"sourceTags,omitempty" tf:"source_tags,omitempty"`

	// A list of service accounts indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// targetServiceAccounts cannot be used at the same time as targetTags or
	// sourceTags. If neither targetServiceAccounts nor targetTags are
	// specified, the firewall rule applies to all instances on the specified
	// network.
	// +kubebuilder:validation:Optional
	TargetServiceAccounts []*string `json:"targetServiceAccounts,omitempty" tf:"target_service_accounts,omitempty"`

	// A list of instance tags indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// If no targetTags are specified, the firewall rule applies to all
	// instances on the specified network.
	// +kubebuilder:validation:Optional
	TargetTags []*string `json:"targetTags,omitempty" tf:"target_tags,omitempty"`
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Firewall is the Schema for the Firewalls API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec"`
	Status            FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
