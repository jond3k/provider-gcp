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

type BackupGeoHealthCheckedTargetsObservation struct {

	// The list of internal load balancers to health check.
	// Structure is document below.
	InternalLoadBalancers []HealthCheckedTargetsInternalLoadBalancersObservation `json:"internalLoadBalancers,omitempty" tf:"internal_load_balancers,omitempty"`
}

type BackupGeoHealthCheckedTargetsParameters struct {

	// The list of internal load balancers to health check.
	// Structure is document below.
	// +kubebuilder:validation:Required
	InternalLoadBalancers []HealthCheckedTargetsInternalLoadBalancersParameters `json:"internalLoadBalancers" tf:"internal_load_balancers,omitempty"`
}

type BackupGeoObservation struct {

	// The list of targets to be health checked. Note that if DNSSEC is enabled for this zone, only one of rrdatas or health_checked_targets can be set.
	// Structure is document below.
	HealthCheckedTargets []BackupGeoHealthCheckedTargetsObservation `json:"healthCheckedTargets,omitempty" tf:"health_checked_targets,omitempty"`

	// The location name defined in Google Cloud.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Same as rrdatas above.
	Rrdatas []*string `json:"rrdatas,omitempty" tf:"rrdatas,omitempty"`
}

type BackupGeoParameters struct {

	// The list of targets to be health checked. Note that if DNSSEC is enabled for this zone, only one of rrdatas or health_checked_targets can be set.
	// Structure is document below.
	// +kubebuilder:validation:Optional
	HealthCheckedTargets []BackupGeoHealthCheckedTargetsParameters `json:"healthCheckedTargets,omitempty" tf:"health_checked_targets,omitempty"`

	// The location name defined in Google Cloud.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Same as rrdatas above.
	// +kubebuilder:validation:Optional
	Rrdatas []*string `json:"rrdatas,omitempty" tf:"rrdatas,omitempty"`
}

type GeoObservation struct {

	// The list of targets to be health checked. Note that if DNSSEC is enabled for this zone, only one of rrdatas or health_checked_targets can be set.
	// Structure is document below.
	HealthCheckedTargets []HealthCheckedTargetsObservation `json:"healthCheckedTargets,omitempty" tf:"health_checked_targets,omitempty"`

	// The location name defined in Google Cloud.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Same as rrdatas above.
	Rrdatas []*string `json:"rrdatas,omitempty" tf:"rrdatas,omitempty"`
}

type GeoParameters struct {

	// The list of targets to be health checked. Note that if DNSSEC is enabled for this zone, only one of rrdatas or health_checked_targets can be set.
	// Structure is document below.
	// +kubebuilder:validation:Optional
	HealthCheckedTargets []HealthCheckedTargetsParameters `json:"healthCheckedTargets,omitempty" tf:"health_checked_targets,omitempty"`

	// The location name defined in Google Cloud.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Same as rrdatas above.
	// +kubebuilder:validation:Optional
	Rrdatas []*string `json:"rrdatas,omitempty" tf:"rrdatas,omitempty"`
}

type HealthCheckedTargetsInternalLoadBalancersObservation struct {

	// The frontend IP address of the load balancer.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb"]
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// The fully qualified url of the network in which the load balancer belongs. This should be formatted like projects/{project}/global/networks/{network} or https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}.
	NetworkURL *string `json:"networkUrl,omitempty" tf:"network_url,omitempty"`

	// The configured port of the load balancer.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the load balancer. Only needed for regional load balancers.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type HealthCheckedTargetsInternalLoadBalancersParameters struct {

	// The frontend IP address of the load balancer.
	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]
	// +kubebuilder:validation:Required
	IPProtocol *string `json:"ipProtocol" tf:"ip_protocol,omitempty"`

	// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb"]
	// +kubebuilder:validation:Required
	LoadBalancerType *string `json:"loadBalancerType" tf:"load_balancer_type,omitempty"`

	// The fully qualified url of the network in which the load balancer belongs. This should be formatted like projects/{project}/global/networks/{network} or https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}.
	// +kubebuilder:validation:Required
	NetworkURL *string `json:"networkUrl" tf:"network_url,omitempty"`

	// The configured port of the load balancer.
	// +kubebuilder:validation:Required
	Port *string `json:"port" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// The region of the load balancer. Only needed for regional load balancers.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type HealthCheckedTargetsObservation struct {

	// The list of internal load balancers to health check.
	// Structure is document below.
	InternalLoadBalancers []InternalLoadBalancersObservation `json:"internalLoadBalancers,omitempty" tf:"internal_load_balancers,omitempty"`
}

type HealthCheckedTargetsParameters struct {

	// The list of internal load balancers to health check.
	// Structure is document below.
	// +kubebuilder:validation:Required
	InternalLoadBalancers []InternalLoadBalancersParameters `json:"internalLoadBalancers" tf:"internal_load_balancers,omitempty"`
}

type InternalLoadBalancersObservation struct {

	// The frontend IP address of the load balancer.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb"]
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// The fully qualified url of the network in which the load balancer belongs. This should be formatted like projects/{project}/global/networks/{network} or https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}.
	NetworkURL *string `json:"networkUrl,omitempty" tf:"network_url,omitempty"`

	// The configured port of the load balancer.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the load balancer. Only needed for regional load balancers.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type InternalLoadBalancersParameters struct {

	// The frontend IP address of the load balancer.
	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]
	// +kubebuilder:validation:Required
	IPProtocol *string `json:"ipProtocol" tf:"ip_protocol,omitempty"`

	// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb"]
	// +kubebuilder:validation:Required
	LoadBalancerType *string `json:"loadBalancerType" tf:"load_balancer_type,omitempty"`

	// The fully qualified url of the network in which the load balancer belongs. This should be formatted like projects/{project}/global/networks/{network} or https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}.
	// +kubebuilder:validation:Required
	NetworkURL *string `json:"networkUrl" tf:"network_url,omitempty"`

	// The configured port of the load balancer.
	// +kubebuilder:validation:Required
	Port *string `json:"port" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// The region of the load balancer. Only needed for regional load balancers.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type PrimaryBackupObservation struct {

	// The backup geo targets, which provide a regional failover policy for the otherwise global primary targets.
	// Structure is document above.
	BackupGeo []BackupGeoObservation `json:"backupGeo,omitempty" tf:"backup_geo,omitempty"`

	// Specifies whether to enable fencing for backup geo queries.
	EnableGeoFencingForBackups *bool `json:"enableGeoFencingForBackups,omitempty" tf:"enable_geo_fencing_for_backups,omitempty"`

	// The list of global primary targets to be health checked.
	// Structure is document below.
	Primary []PrimaryObservation `json:"primary,omitempty" tf:"primary,omitempty"`

	// Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy.
	TrickleRatio *float64 `json:"trickleRatio,omitempty" tf:"trickle_ratio,omitempty"`
}

type PrimaryBackupParameters struct {

	// The backup geo targets, which provide a regional failover policy for the otherwise global primary targets.
	// Structure is document above.
	// +kubebuilder:validation:Required
	BackupGeo []BackupGeoParameters `json:"backupGeo" tf:"backup_geo,omitempty"`

	// Specifies whether to enable fencing for backup geo queries.
	// +kubebuilder:validation:Optional
	EnableGeoFencingForBackups *bool `json:"enableGeoFencingForBackups,omitempty" tf:"enable_geo_fencing_for_backups,omitempty"`

	// The list of global primary targets to be health checked.
	// Structure is document below.
	// +kubebuilder:validation:Required
	Primary []PrimaryParameters `json:"primary" tf:"primary,omitempty"`

	// Specifies the percentage of traffic to send to the backup targets even when the primary targets are healthy.
	// +kubebuilder:validation:Optional
	TrickleRatio *float64 `json:"trickleRatio,omitempty" tf:"trickle_ratio,omitempty"`
}

type PrimaryInternalLoadBalancersObservation struct {

	// The frontend IP address of the load balancer.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb"]
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// The fully qualified url of the network in which the load balancer belongs. This should be formatted like projects/{project}/global/networks/{network} or https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}.
	NetworkURL *string `json:"networkUrl,omitempty" tf:"network_url,omitempty"`

	// The configured port of the load balancer.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the load balancer. Only needed for regional load balancers.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type PrimaryInternalLoadBalancersParameters struct {

	// The frontend IP address of the load balancer.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.ForwardingRule
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("ip_address",false)
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Reference to a ForwardingRule in compute to populate ipAddress.
	// +kubebuilder:validation:Optional
	IPAddressRef *v1.Reference `json:"ipAddressRef,omitempty" tf:"-"`

	// Selector for a ForwardingRule in compute to populate ipAddress.
	// +kubebuilder:validation:Optional
	IPAddressSelector *v1.Selector `json:"ipAddressSelector,omitempty" tf:"-"`

	// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]
	// +kubebuilder:validation:Required
	IPProtocol *string `json:"ipProtocol" tf:"ip_protocol,omitempty"`

	// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb"]
	// +kubebuilder:validation:Required
	LoadBalancerType *string `json:"loadBalancerType" tf:"load_balancer_type,omitempty"`

	// The fully qualified url of the network in which the load balancer belongs. This should be formatted like projects/{project}/global/networks/{network} or https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkURL *string `json:"networkUrl,omitempty" tf:"network_url,omitempty"`

	// Reference to a Network in compute to populate networkUrl.
	// +kubebuilder:validation:Optional
	NetworkURLRef *v1.Reference `json:"networkUrlRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate networkUrl.
	// +kubebuilder:validation:Optional
	NetworkURLSelector *v1.Selector `json:"networkUrlSelector,omitempty" tf:"-"`

	// The configured port of the load balancer.
	// +kubebuilder:validation:Required
	Port *string `json:"port" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.ForwardingRule
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("project",false)
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Reference to a ForwardingRule in compute to populate project.
	// +kubebuilder:validation:Optional
	ProjectRef *v1.Reference `json:"projectRef,omitempty" tf:"-"`

	// Selector for a ForwardingRule in compute to populate project.
	// +kubebuilder:validation:Optional
	ProjectSelector *v1.Selector `json:"projectSelector,omitempty" tf:"-"`

	// The region of the load balancer. Only needed for regional load balancers.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.ForwardingRule
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("region",false)
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Reference to a ForwardingRule in compute to populate region.
	// +kubebuilder:validation:Optional
	RegionRef *v1.Reference `json:"regionRef,omitempty" tf:"-"`

	// Selector for a ForwardingRule in compute to populate region.
	// +kubebuilder:validation:Optional
	RegionSelector *v1.Selector `json:"regionSelector,omitempty" tf:"-"`
}

type PrimaryObservation struct {

	// The list of internal load balancers to health check.
	// Structure is document below.
	InternalLoadBalancers []PrimaryInternalLoadBalancersObservation `json:"internalLoadBalancers,omitempty" tf:"internal_load_balancers,omitempty"`
}

type PrimaryParameters struct {

	// The list of internal load balancers to health check.
	// Structure is document below.
	// +kubebuilder:validation:Required
	InternalLoadBalancers []PrimaryInternalLoadBalancersParameters `json:"internalLoadBalancers" tf:"internal_load_balancers,omitempty"`
}

type RecordSetObservation struct {

	// an identifier for the resource with format projects/{{project}}/managedZones/{{zone}}/rrsets/{{name}}/{{type}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the zone in which this record set will
	// reside.
	ManagedZone *string `json:"managedZone,omitempty" tf:"managed_zone,omitempty"`

	// The DNS name this record set will apply to.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The configuration for steering traffic based on query.
	// Now you can specify either Weighted Round Robin(WRR) type or Geolocation(GEO) type.
	// Structure is documented below.
	RoutingPolicy []RoutingPolicyObservation `json:"routingPolicy,omitempty" tf:"routing_policy,omitempty"`

	// The string data for the records in this record set
	// whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding \" if you don't want your string to get split on spaces.g. "first255characters\" \"morecharacters").
	Rrdatas []*string `json:"rrdatas,omitempty" tf:"rrdatas,omitempty"`

	// The time-to-live of this record set (seconds).
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The DNS record set type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RecordSetParameters struct {

	// The name of the zone in which this record set will
	// reside.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/dns/v1beta1.ManagedZone
	// +kubebuilder:validation:Optional
	ManagedZone *string `json:"managedZone,omitempty" tf:"managed_zone,omitempty"`

	// Reference to a ManagedZone in dns to populate managedZone.
	// +kubebuilder:validation:Optional
	ManagedZoneRef *v1.Reference `json:"managedZoneRef,omitempty" tf:"-"`

	// Selector for a ManagedZone in dns to populate managedZone.
	// +kubebuilder:validation:Optional
	ManagedZoneSelector *v1.Selector `json:"managedZoneSelector,omitempty" tf:"-"`

	// The DNS name this record set will apply to.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The configuration for steering traffic based on query.
	// Now you can specify either Weighted Round Robin(WRR) type or Geolocation(GEO) type.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RoutingPolicy []RoutingPolicyParameters `json:"routingPolicy,omitempty" tf:"routing_policy,omitempty"`

	// The string data for the records in this record set
	// whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding \" if you don't want your string to get split on spaces.g. "first255characters\" \"morecharacters").
	// +kubebuilder:validation:Optional
	Rrdatas []*string `json:"rrdatas,omitempty" tf:"rrdatas,omitempty"`

	// The time-to-live of this record set (seconds).
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The DNS record set type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RoutingPolicyObservation struct {

	// Specifies whether to enable fencing for geo queries.
	EnableGeoFencing *bool `json:"enableGeoFencing,omitempty" tf:"enable_geo_fencing,omitempty"`

	// The configuration for Geolocation based routing policy.
	// Structure is document below.
	Geo []GeoObservation `json:"geo,omitempty" tf:"geo,omitempty"`

	// The configuration for a primary-backup policy with global to regional failover. Queries are responded to with the global primary targets, but if none of the primary targets are healthy, then we fallback to a regional failover policy.
	// Structure is document below.
	PrimaryBackup []PrimaryBackupObservation `json:"primaryBackup,omitempty" tf:"primary_backup,omitempty"`

	// The configuration for Weighted Round Robin based routing policy.
	// Structure is document below.
	Wrr []WrrObservation `json:"wrr,omitempty" tf:"wrr,omitempty"`
}

type RoutingPolicyParameters struct {

	// Specifies whether to enable fencing for geo queries.
	// +kubebuilder:validation:Optional
	EnableGeoFencing *bool `json:"enableGeoFencing,omitempty" tf:"enable_geo_fencing,omitempty"`

	// The configuration for Geolocation based routing policy.
	// Structure is document below.
	// +kubebuilder:validation:Optional
	Geo []GeoParameters `json:"geo,omitempty" tf:"geo,omitempty"`

	// The configuration for a primary-backup policy with global to regional failover. Queries are responded to with the global primary targets, but if none of the primary targets are healthy, then we fallback to a regional failover policy.
	// Structure is document below.
	// +kubebuilder:validation:Optional
	PrimaryBackup []PrimaryBackupParameters `json:"primaryBackup,omitempty" tf:"primary_backup,omitempty"`

	// The configuration for Weighted Round Robin based routing policy.
	// Structure is document below.
	// +kubebuilder:validation:Optional
	Wrr []WrrParameters `json:"wrr,omitempty" tf:"wrr,omitempty"`
}

type WrrHealthCheckedTargetsInternalLoadBalancersObservation struct {

	// The frontend IP address of the load balancer.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]
	IPProtocol *string `json:"ipProtocol,omitempty" tf:"ip_protocol,omitempty"`

	// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb"]
	LoadBalancerType *string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`

	// The fully qualified url of the network in which the load balancer belongs. This should be formatted like projects/{project}/global/networks/{network} or https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}.
	NetworkURL *string `json:"networkUrl,omitempty" tf:"network_url,omitempty"`

	// The configured port of the load balancer.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the load balancer. Only needed for regional load balancers.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type WrrHealthCheckedTargetsInternalLoadBalancersParameters struct {

	// The frontend IP address of the load balancer.
	// +kubebuilder:validation:Required
	IPAddress *string `json:"ipAddress" tf:"ip_address,omitempty"`

	// The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: ["tcp", "udp"]
	// +kubebuilder:validation:Required
	IPProtocol *string `json:"ipProtocol" tf:"ip_protocol,omitempty"`

	// The type of load balancer. This value is case-sensitive. Possible values: ["regionalL4ilb"]
	// +kubebuilder:validation:Required
	LoadBalancerType *string `json:"loadBalancerType" tf:"load_balancer_type,omitempty"`

	// The fully qualified url of the network in which the load balancer belongs. This should be formatted like projects/{project}/global/networks/{network} or https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}.
	// +kubebuilder:validation:Required
	NetworkURL *string `json:"networkUrl" tf:"network_url,omitempty"`

	// The configured port of the load balancer.
	// +kubebuilder:validation:Required
	Port *string `json:"port" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// The region of the load balancer. Only needed for regional load balancers.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type WrrHealthCheckedTargetsObservation struct {

	// The list of internal load balancers to health check.
	// Structure is document below.
	InternalLoadBalancers []WrrHealthCheckedTargetsInternalLoadBalancersObservation `json:"internalLoadBalancers,omitempty" tf:"internal_load_balancers,omitempty"`
}

type WrrHealthCheckedTargetsParameters struct {

	// The list of internal load balancers to health check.
	// Structure is document below.
	// +kubebuilder:validation:Required
	InternalLoadBalancers []WrrHealthCheckedTargetsInternalLoadBalancersParameters `json:"internalLoadBalancers" tf:"internal_load_balancers,omitempty"`
}

type WrrObservation struct {

	// The list of targets to be health checked. Note that if DNSSEC is enabled for this zone, only one of rrdatas or health_checked_targets can be set.
	// Structure is document below.
	HealthCheckedTargets []WrrHealthCheckedTargetsObservation `json:"healthCheckedTargets,omitempty" tf:"health_checked_targets,omitempty"`

	// Same as rrdatas above.
	Rrdatas []*string `json:"rrdatas,omitempty" tf:"rrdatas,omitempty"`

	// The ratio of traffic routed to the target.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type WrrParameters struct {

	// The list of targets to be health checked. Note that if DNSSEC is enabled for this zone, only one of rrdatas or health_checked_targets can be set.
	// Structure is document below.
	// +kubebuilder:validation:Optional
	HealthCheckedTargets []WrrHealthCheckedTargetsParameters `json:"healthCheckedTargets,omitempty" tf:"health_checked_targets,omitempty"`

	// Same as rrdatas above.
	// +kubebuilder:validation:Optional
	Rrdatas []*string `json:"rrdatas,omitempty" tf:"rrdatas,omitempty"`

	// The ratio of traffic routed to the target.
	// +kubebuilder:validation:Required
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// RecordSetSpec defines the desired state of RecordSet
type RecordSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordSetParameters `json:"forProvider"`
}

// RecordSetStatus defines the observed state of RecordSet.
type RecordSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RecordSet is the Schema for the RecordSets API. Manages a set of DNS records within Google Cloud DNS.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RecordSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type)",message="type is a required parameter"
	Spec   RecordSetSpec   `json:"spec"`
	Status RecordSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordSetList contains a list of RecordSets
type RecordSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecordSet `json:"items"`
}

// Repository type metadata.
var (
	RecordSet_Kind             = "RecordSet"
	RecordSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RecordSet_Kind}.String()
	RecordSet_KindAPIVersion   = RecordSet_Kind + "." + CRDGroupVersion.String()
	RecordSet_GroupVersionKind = CRDGroupVersion.WithKind(RecordSet_Kind)
)

func init() {
	SchemeBuilder.Register(&RecordSet{}, &RecordSetList{})
}
