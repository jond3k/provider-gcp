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

type NetworkConfigInitParameters struct {

	// Immutable. Subnet IP range within the peered network. This is specified in CIDR notation with a slash and the subnet prefix size. You can optionally specify an IP address before the subnet prefix value. e.g. 192.168.0.0/29 would specify an IP range starting at 192.168.0.0 with a prefix size of 29 bits. /16 would specify a prefix size of 16 bits, with an automatically determined IP within the peered VPC. If unspecified, a value of /24 will be used.
	PeeredNetworkIPRange *string `json:"peeredNetworkIpRange,omitempty" tf:"peered_network_ip_range,omitempty"`
}

type NetworkConfigObservation struct {

	// Immutable. The network definition that the workers are peered to. If this section is left empty, the workers will be peered to WorkerPool.project_id on the service producer network. Must be in the format projects/{project}/global/networks/{network}, where {project} is a project number, such as 12345, and {network} is the name of a VPC network in the project. See (https://cloud.google.com/cloud-build/docs/custom-workers/set-up-custom-worker-pool-environment#understanding_the_network_configuration_options)
	PeeredNetwork *string `json:"peeredNetwork,omitempty" tf:"peered_network,omitempty"`

	// Immutable. Subnet IP range within the peered network. This is specified in CIDR notation with a slash and the subnet prefix size. You can optionally specify an IP address before the subnet prefix value. e.g. 192.168.0.0/29 would specify an IP range starting at 192.168.0.0 with a prefix size of 29 bits. /16 would specify a prefix size of 16 bits, with an automatically determined IP within the peered VPC. If unspecified, a value of /24 will be used.
	PeeredNetworkIPRange *string `json:"peeredNetworkIpRange,omitempty" tf:"peered_network_ip_range,omitempty"`
}

type NetworkConfigParameters struct {

	// Immutable. The network definition that the workers are peered to. If this section is left empty, the workers will be peered to WorkerPool.project_id on the service producer network. Must be in the format projects/{project}/global/networks/{network}, where {project} is a project number, such as 12345, and {network} is the name of a VPC network in the project. See (https://cloud.google.com/cloud-build/docs/custom-workers/set-up-custom-worker-pool-environment#understanding_the_network_configuration_options)
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PeeredNetwork *string `json:"peeredNetwork,omitempty" tf:"peered_network,omitempty"`

	// Immutable. Subnet IP range within the peered network. This is specified in CIDR notation with a slash and the subnet prefix size. You can optionally specify an IP address before the subnet prefix value. e.g. 192.168.0.0/29 would specify an IP range starting at 192.168.0.0 with a prefix size of 29 bits. /16 would specify a prefix size of 16 bits, with an automatically determined IP within the peered VPC. If unspecified, a value of /24 will be used.
	// +kubebuilder:validation:Optional
	PeeredNetworkIPRange *string `json:"peeredNetworkIpRange,omitempty" tf:"peered_network_ip_range,omitempty"`

	// Reference to a Network in compute to populate peeredNetwork.
	// +kubebuilder:validation:Optional
	PeeredNetworkRef *v1.Reference `json:"peeredNetworkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate peeredNetwork.
	// +kubebuilder:validation:Optional
	PeeredNetworkSelector *v1.Selector `json:"peeredNetworkSelector,omitempty" tf:"-"`
}

type WorkerConfigInitParameters struct {

	// Size of the disk attached to the worker, in GB. See (https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). Specify a value of up to 1000. If 0 is specified, Cloud Build will use a standard disk size.
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// Machine type of a worker, such as n1-standard-1. See (https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). If left blank, Cloud Build will use n1-standard-1.
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// If true, workers are created without any public address, which prevents network egress to public IPs.
	NoExternalIP *bool `json:"noExternalIp,omitempty" tf:"no_external_ip,omitempty"`
}

type WorkerConfigObservation struct {

	// Size of the disk attached to the worker, in GB. See (https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). Specify a value of up to 1000. If 0 is specified, Cloud Build will use a standard disk size.
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// Machine type of a worker, such as n1-standard-1. See (https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). If left blank, Cloud Build will use n1-standard-1.
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// If true, workers are created without any public address, which prevents network egress to public IPs.
	NoExternalIP *bool `json:"noExternalIp,omitempty" tf:"no_external_ip,omitempty"`
}

type WorkerConfigParameters struct {

	// Size of the disk attached to the worker, in GB. See (https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). Specify a value of up to 1000. If 0 is specified, Cloud Build will use a standard disk size.
	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// Machine type of a worker, such as n1-standard-1. See (https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). If left blank, Cloud Build will use n1-standard-1.
	// +kubebuilder:validation:Optional
	MachineType *string `json:"machineType,omitempty" tf:"machine_type,omitempty"`

	// If true, workers are created without any public address, which prevents network egress to public IPs.
	// +kubebuilder:validation:Optional
	NoExternalIP *bool `json:"noExternalIp,omitempty" tf:"no_external_ip,omitempty"`
}

type WorkerPoolInitParameters struct {
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// User-defined name of the WorkerPool.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Network configuration for the WorkerPool. Structure is documented below.
	NetworkConfig []NetworkConfigInitParameters `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Configuration to be used for a creating workers in the WorkerPool. Structure is documented below.
	WorkerConfig []WorkerConfigInitParameters `json:"workerConfig,omitempty" tf:"worker_config,omitempty"`
}

type WorkerPoolObservation struct {
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Output only. Time at which the request to create the WorkerPool was received.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Output only. Time at which the request to delete the WorkerPool was received.
	DeleteTime *string `json:"deleteTime,omitempty" tf:"delete_time,omitempty"`

	// User-defined name of the WorkerPool.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/workerPools/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Network configuration for the WorkerPool. Structure is documented below.
	NetworkConfig []NetworkConfigObservation `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Output only. WorkerPool state. Possible values: STATE_UNSPECIFIED, PENDING, APPROVED, REJECTED, CANCELLED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/workerPools/{{name}}
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Output only. Time at which the request to update the WorkerPool was received.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Configuration to be used for a creating workers in the WorkerPool. Structure is documented below.
	WorkerConfig []WorkerConfigObservation `json:"workerConfig,omitempty" tf:"worker_config,omitempty"`
}

type WorkerPoolParameters struct {

	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// User-defined name of the WorkerPool.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Network configuration for the WorkerPool. Structure is documented below.
	// +kubebuilder:validation:Optional
	NetworkConfig []NetworkConfigParameters `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Configuration to be used for a creating workers in the WorkerPool. Structure is documented below.
	// +kubebuilder:validation:Optional
	WorkerConfig []WorkerConfigParameters `json:"workerConfig,omitempty" tf:"worker_config,omitempty"`
}

// WorkerPoolSpec defines the desired state of WorkerPool
type WorkerPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkerPoolParameters `json:"forProvider"`
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
	InitProvider WorkerPoolInitParameters `json:"initProvider,omitempty"`
}

// WorkerPoolStatus defines the observed state of WorkerPool.
type WorkerPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkerPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorkerPool is the Schema for the WorkerPools API. Configuration for custom WorkerPool to run builds
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type WorkerPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkerPoolSpec   `json:"spec"`
	Status            WorkerPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkerPoolList contains a list of WorkerPools
type WorkerPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkerPool `json:"items"`
}

// Repository type metadata.
var (
	WorkerPool_Kind             = "WorkerPool"
	WorkerPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkerPool_Kind}.String()
	WorkerPool_KindAPIVersion   = WorkerPool_Kind + "." + CRDGroupVersion.String()
	WorkerPool_GroupVersionKind = CRDGroupVersion.WithKind(WorkerPool_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkerPool{}, &WorkerPoolList{})
}
