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

type AutoscalingObservation struct {

	// Maximum number of nodes in the node pool. Must be >= min_node_count.
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// Minimum number of nodes in the node pool. Must be >= 1 and <= max_node_count.
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`
}

type AutoscalingParameters struct {

	// Maximum number of nodes in the node pool. Must be >= min_node_count.
	// +kubebuilder:validation:Required
	MaxNodeCount *float64 `json:"maxNodeCount" tf:"max_node_count,omitempty"`

	// Minimum number of nodes in the node pool. Must be >= 1 and <= max_node_count.
	// +kubebuilder:validation:Required
	MinNodeCount *float64 `json:"minNodeCount" tf:"min_node_count,omitempty"`
}

type ConfigObservation struct {

	// Proxy configuration for outbound HTTP(S) traffic.
	ProxyConfig []ConfigProxyConfigObservation `json:"proxyConfig,omitempty" tf:"proxy_config,omitempty"`

	// Optional. Configuration related to the root volume provisioned for each node pool machine. When unspecified, it defaults to a 32-GiB Azure Disk.
	RootVolume []ConfigRootVolumeObservation `json:"rootVolume,omitempty" tf:"root_volume,omitempty"`

	// SSH configuration for how to access the node pool machines.
	SSHConfig []ConfigSSHConfigObservation `json:"sshConfig,omitempty" tf:"ssh_config,omitempty"`

	// Optional. A set of tags to apply to all underlying Azure resources for this node pool. This currently only includes Virtual Machine Scale Sets. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Optional. The Azure VM size name. Example: Standard_DS2_v2. See (/anthos/clusters/docs/azure/reference/supported-vms) for options. When unspecified, it defaults to Standard_DS2_v2.
	VMSize *string `json:"vmSize,omitempty" tf:"vm_size,omitempty"`
}

type ConfigParameters struct {

	// Proxy configuration for outbound HTTP(S) traffic.
	// +kubebuilder:validation:Optional
	ProxyConfig []ConfigProxyConfigParameters `json:"proxyConfig,omitempty" tf:"proxy_config,omitempty"`

	// Optional. Configuration related to the root volume provisioned for each node pool machine. When unspecified, it defaults to a 32-GiB Azure Disk.
	// +kubebuilder:validation:Optional
	RootVolume []ConfigRootVolumeParameters `json:"rootVolume,omitempty" tf:"root_volume,omitempty"`

	// SSH configuration for how to access the node pool machines.
	// +kubebuilder:validation:Required
	SSHConfig []ConfigSSHConfigParameters `json:"sshConfig" tf:"ssh_config,omitempty"`

	// Optional. A set of tags to apply to all underlying Azure resources for this node pool. This currently only includes Virtual Machine Scale Sets. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Optional. The Azure VM size name. Example: Standard_DS2_v2. See (/anthos/clusters/docs/azure/reference/supported-vms) for options. When unspecified, it defaults to Standard_DS2_v2.
	// +kubebuilder:validation:Optional
	VMSize *string `json:"vmSize,omitempty" tf:"vm_size,omitempty"`
}

type ConfigProxyConfigObservation struct {

	// The ARM ID the of the resource group containing proxy keyvault. Resource group ids are formatted as /subscriptions/<subscription-id>/resourceGroups/<resource-group-name>
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// The URL the of the proxy setting secret with its version. Secret ids are formatted as https:<key-vault-name>.vault.azure.net/secrets/<secret-name>/<secret-version>.
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`
}

type ConfigProxyConfigParameters struct {

	// The ARM ID the of the resource group containing proxy keyvault. Resource group ids are formatted as /subscriptions/<subscription-id>/resourceGroups/<resource-group-name>
	// +kubebuilder:validation:Required
	ResourceGroupID *string `json:"resourceGroupId" tf:"resource_group_id,omitempty"`

	// The URL the of the proxy setting secret with its version. Secret ids are formatted as https:<key-vault-name>.vault.azure.net/secrets/<secret-name>/<secret-version>.
	// +kubebuilder:validation:Required
	SecretID *string `json:"secretId" tf:"secret_id,omitempty"`
}

type ConfigRootVolumeObservation struct {

	// Optional. The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	SizeGib *float64 `json:"sizeGib,omitempty" tf:"size_gib,omitempty"`
}

type ConfigRootVolumeParameters struct {

	// Optional. The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	// +kubebuilder:validation:Optional
	SizeGib *float64 `json:"sizeGib,omitempty" tf:"size_gib,omitempty"`
}

type ConfigSSHConfigObservation struct {

	// The SSH public key data for VMs managed by Anthos. This accepts the authorized_keys file format used in OpenSSH according to the sshd(8) manual page.
	AuthorizedKey *string `json:"authorizedKey,omitempty" tf:"authorized_key,omitempty"`
}

type ConfigSSHConfigParameters struct {

	// The SSH public key data for VMs managed by Anthos. This accepts the authorized_keys file format used in OpenSSH according to the sshd(8) manual page.
	// +kubebuilder:validation:Required
	AuthorizedKey *string `json:"authorizedKey" tf:"authorized_key,omitempty"`
}

type MaxPodsConstraintObservation struct {

	// The maximum number of pods to schedule on a single node.
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`
}

type MaxPodsConstraintParameters struct {

	// The maximum number of pods to schedule on a single node.
	// +kubebuilder:validation:Required
	MaxPodsPerNode *float64 `json:"maxPodsPerNode" tf:"max_pods_per_node,omitempty"`
}

type NodePoolObservation struct {

	// Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix  and name , separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Autoscaler configuration for this node pool.
	Autoscaling []AutoscalingObservation `json:"autoscaling,omitempty" tf:"autoscaling,omitempty"`

	// Optional. The Azure availability zone of the nodes in this nodepool. When unspecified, it defaults to 1.
	AzureAvailabilityZone *string `json:"azureAvailabilityZone,omitempty" tf:"azure_availability_zone,omitempty"`

	// The azureCluster for the resource
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The node configuration of the node pool.
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// Output only. The time at which this node pool was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/azureClusters/{{cluster}}/azureNodePools/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
	MaxPodsConstraint []MaxPodsConstraintObservation `json:"maxPodsConstraint,omitempty" tf:"max_pods_constraint,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Output only. If set, there are currently pending changes to the node pool.
	Reconciling *bool `json:"reconciling,omitempty" tf:"reconciling,omitempty"`

	// Output only. The current state of the node pool. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The ARM ID of the subnet where the node pool VMs run. Make sure it's a subnet under the virtual network in the cluster configuration.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Output only. A globally unique identifier for the node pool.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Output only. The time at which this node pool was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// The Kubernetes version (e.g. 1.19.10-gke.1000) running on this node pool.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NodePoolParameters struct {

	// Optional. Annotations on the node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix  and name , separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Autoscaler configuration for this node pool.
	// +kubebuilder:validation:Optional
	Autoscaling []AutoscalingParameters `json:"autoscaling,omitempty" tf:"autoscaling,omitempty"`

	// Optional. The Azure availability zone of the nodes in this nodepool. When unspecified, it defaults to 1.
	// +kubebuilder:validation:Optional
	AzureAvailabilityZone *string `json:"azureAvailabilityZone,omitempty" tf:"azure_availability_zone,omitempty"`

	// The azureCluster for the resource
	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// Reference to a Cluster to populate cluster.
	// +kubebuilder:validation:Optional
	ClusterRef *v1.Reference `json:"clusterRef,omitempty" tf:"-"`

	// Selector for a Cluster to populate cluster.
	// +kubebuilder:validation:Optional
	ClusterSelector *v1.Selector `json:"clusterSelector,omitempty" tf:"-"`

	// The node configuration of the node pool.
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool.
	// +kubebuilder:validation:Optional
	MaxPodsConstraint []MaxPodsConstraintParameters `json:"maxPodsConstraint,omitempty" tf:"max_pods_constraint,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The ARM ID of the subnet where the node pool VMs run. Make sure it's a subnet under the virtual network in the cluster configuration.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The Kubernetes version (e.g. 1.19.10-gke.1000) running on this node pool.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// NodePoolSpec defines the desired state of NodePool
type NodePoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodePoolParameters `json:"forProvider"`
}

// NodePoolStatus defines the observed state of NodePool.
type NodePoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodePoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodePool is the Schema for the NodePools API. An Anthos node pool running on Azure.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type NodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.autoscaling)",message="autoscaling is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config)",message="config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.maxPodsConstraint)",message="maxPodsConstraint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subnetId)",message="subnetId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version)",message="version is a required parameter"
	Spec   NodePoolSpec   `json:"spec"`
	Status NodePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodePoolList contains a list of NodePools
type NodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodePool `json:"items"`
}

// Repository type metadata.
var (
	NodePool_Kind             = "NodePool"
	NodePool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodePool_Kind}.String()
	NodePool_KindAPIVersion   = NodePool_Kind + "." + CRDGroupVersion.String()
	NodePool_GroupVersionKind = CRDGroupVersion.WithKind(NodePool_Kind)
)

func init() {
	SchemeBuilder.Register(&NodePool{}, &NodePoolList{})
}
