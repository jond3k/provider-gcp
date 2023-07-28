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

type DockerConfigInitParameters struct {

	// The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created.
	ImmutableTags *bool `json:"immutableTags,omitempty" tf:"immutable_tags,omitempty"`
}

type DockerConfigObservation struct {

	// The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created.
	ImmutableTags *bool `json:"immutableTags,omitempty" tf:"immutable_tags,omitempty"`
}

type DockerConfigParameters struct {

	// The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created.
	// +kubebuilder:validation:Optional
	ImmutableTags *bool `json:"immutableTags,omitempty" tf:"immutable_tags,omitempty"`
}

type DockerRepositoryInitParameters struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type DockerRepositoryObservation struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type DockerRepositoryParameters struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	// +kubebuilder:validation:Optional
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type MavenConfigInitParameters struct {

	// The repository with this flag will allow publishing the same
	// snapshot versions.
	AllowSnapshotOverwrites *bool `json:"allowSnapshotOverwrites,omitempty" tf:"allow_snapshot_overwrites,omitempty"`

	// Version policy defines the versions that the registry will accept.
	// Default value is VERSION_POLICY_UNSPECIFIED.
	// Possible values are: VERSION_POLICY_UNSPECIFIED, RELEASE, SNAPSHOT.
	VersionPolicy *string `json:"versionPolicy,omitempty" tf:"version_policy,omitempty"`
}

type MavenConfigObservation struct {

	// The repository with this flag will allow publishing the same
	// snapshot versions.
	AllowSnapshotOverwrites *bool `json:"allowSnapshotOverwrites,omitempty" tf:"allow_snapshot_overwrites,omitempty"`

	// Version policy defines the versions that the registry will accept.
	// Default value is VERSION_POLICY_UNSPECIFIED.
	// Possible values are: VERSION_POLICY_UNSPECIFIED, RELEASE, SNAPSHOT.
	VersionPolicy *string `json:"versionPolicy,omitempty" tf:"version_policy,omitempty"`
}

type MavenConfigParameters struct {

	// The repository with this flag will allow publishing the same
	// snapshot versions.
	// +kubebuilder:validation:Optional
	AllowSnapshotOverwrites *bool `json:"allowSnapshotOverwrites,omitempty" tf:"allow_snapshot_overwrites,omitempty"`

	// Version policy defines the versions that the registry will accept.
	// Default value is VERSION_POLICY_UNSPECIFIED.
	// Possible values are: VERSION_POLICY_UNSPECIFIED, RELEASE, SNAPSHOT.
	// +kubebuilder:validation:Optional
	VersionPolicy *string `json:"versionPolicy,omitempty" tf:"version_policy,omitempty"`
}

type MavenRepositoryInitParameters struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type MavenRepositoryObservation struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type MavenRepositoryParameters struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	// +kubebuilder:validation:Optional
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type NpmRepositoryInitParameters struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type NpmRepositoryObservation struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type NpmRepositoryParameters struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	// +kubebuilder:validation:Optional
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type PythonRepositoryInitParameters struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type PythonRepositoryObservation struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type PythonRepositoryParameters struct {

	// Address of the remote repository.
	// Default value is PYPI.
	// Possible values are: PYPI.
	// +kubebuilder:validation:Optional
	PublicRepository *string `json:"publicRepository,omitempty" tf:"public_repository,omitempty"`
}

type RegistryRepositoryInitParameters struct {

	// The user-provided description of the repository.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Docker repository config contains repository level configuration for the repositories of docker type.
	// Structure is documented below.
	DockerConfig []DockerConfigInitParameters `json:"dockerConfig,omitempty" tf:"docker_config,omitempty"`

	// The format of packages that are stored in the repository. Supported formats
	// can be found here.
	// You can only create alpha formats if you are a member of the
	// alpha user group.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key.
	// This value may not be changed after the Repository has been created.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// MavenRepositoryConfig is maven related repository details.
	// Provides additional configuration details for repositories of the maven
	// format type.
	// Structure is documented below.
	MavenConfig []MavenConfigInitParameters `json:"mavenConfig,omitempty" tf:"maven_config,omitempty"`

	// The mode configures the repository to serve artifacts from different sources.
	// Default value is STANDARD_REPOSITORY.
	// Possible values are: STANDARD_REPOSITORY, VIRTUAL_REPOSITORY, REMOTE_REPOSITORY.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Configuration specific for a Remote Repository.
	// Structure is documented below.
	RemoteRepositoryConfig []RemoteRepositoryConfigInitParameters `json:"remoteRepositoryConfig,omitempty" tf:"remote_repository_config,omitempty"`

	// Configuration specific for a Virtual Repository.
	// Structure is documented below.
	VirtualRepositoryConfig []VirtualRepositoryConfigInitParameters `json:"virtualRepositoryConfig,omitempty" tf:"virtual_repository_config,omitempty"`
}

type RegistryRepositoryObservation struct {

	// The time when the repository was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The user-provided description of the repository.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Docker repository config contains repository level configuration for the repositories of docker type.
	// Structure is documented below.
	DockerConfig []DockerConfigObservation `json:"dockerConfig,omitempty" tf:"docker_config,omitempty"`

	// The format of packages that are stored in the repository. Supported formats
	// can be found here.
	// You can only create alpha formats if you are a member of the
	// alpha user group.
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key.
	// This value may not be changed after the Repository has been created.
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the location this repository is located in.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// MavenRepositoryConfig is maven related repository details.
	// Provides additional configuration details for repositories of the maven
	// format type.
	// Structure is documented below.
	MavenConfig []MavenConfigObservation `json:"mavenConfig,omitempty" tf:"maven_config,omitempty"`

	// The mode configures the repository to serve artifacts from different sources.
	// Default value is STANDARD_REPOSITORY.
	// Possible values are: STANDARD_REPOSITORY, VIRTUAL_REPOSITORY, REMOTE_REPOSITORY.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The name of the repository, for example:
	// "repo1"
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Configuration specific for a Remote Repository.
	// Structure is documented below.
	RemoteRepositoryConfig []RemoteRepositoryConfigObservation `json:"remoteRepositoryConfig,omitempty" tf:"remote_repository_config,omitempty"`

	// The time when the repository was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`

	// Configuration specific for a Virtual Repository.
	// Structure is documented below.
	VirtualRepositoryConfig []VirtualRepositoryConfigObservation `json:"virtualRepositoryConfig,omitempty" tf:"virtual_repository_config,omitempty"`
}

type RegistryRepositoryParameters struct {

	// The user-provided description of the repository.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Docker repository config contains repository level configuration for the repositories of docker type.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DockerConfig []DockerConfigParameters `json:"dockerConfig,omitempty" tf:"docker_config,omitempty"`

	// The format of packages that are stored in the repository. Supported formats
	// can be found here.
	// You can only create alpha formats if you are a member of the
	// alpha user group.
	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// The Cloud KMS resource name of the customer managed encryption key that’s
	// used to encrypt the contents of the Repository. Has the form:
	// projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key.
	// This value may not be changed after the Repository has been created.
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// Labels with user-defined metadata.
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The name of the location this repository is located in.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// MavenRepositoryConfig is maven related repository details.
	// Provides additional configuration details for repositories of the maven
	// format type.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MavenConfig []MavenConfigParameters `json:"mavenConfig,omitempty" tf:"maven_config,omitempty"`

	// The mode configures the repository to serve artifacts from different sources.
	// Default value is STANDARD_REPOSITORY.
	// Possible values are: STANDARD_REPOSITORY, VIRTUAL_REPOSITORY, REMOTE_REPOSITORY.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Configuration specific for a Remote Repository.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RemoteRepositoryConfig []RemoteRepositoryConfigParameters `json:"remoteRepositoryConfig,omitempty" tf:"remote_repository_config,omitempty"`

	// Configuration specific for a Virtual Repository.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VirtualRepositoryConfig []VirtualRepositoryConfigParameters `json:"virtualRepositoryConfig,omitempty" tf:"virtual_repository_config,omitempty"`
}

type RemoteRepositoryConfigInitParameters struct {

	// The description of the remote source.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specific settings for a Docker remote repository.
	// Structure is documented below.
	DockerRepository []DockerRepositoryInitParameters `json:"dockerRepository,omitempty" tf:"docker_repository,omitempty"`

	// Specific settings for a Maven remote repository.
	// Structure is documented below.
	MavenRepository []MavenRepositoryInitParameters `json:"mavenRepository,omitempty" tf:"maven_repository,omitempty"`

	// Specific settings for an Npm remote repository.
	// Structure is documented below.
	NpmRepository []NpmRepositoryInitParameters `json:"npmRepository,omitempty" tf:"npm_repository,omitempty"`

	// Specific settings for a Python remote repository.
	// Structure is documented below.
	PythonRepository []PythonRepositoryInitParameters `json:"pythonRepository,omitempty" tf:"python_repository,omitempty"`
}

type RemoteRepositoryConfigObservation struct {

	// The description of the remote source.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specific settings for a Docker remote repository.
	// Structure is documented below.
	DockerRepository []DockerRepositoryObservation `json:"dockerRepository,omitempty" tf:"docker_repository,omitempty"`

	// Specific settings for a Maven remote repository.
	// Structure is documented below.
	MavenRepository []MavenRepositoryObservation `json:"mavenRepository,omitempty" tf:"maven_repository,omitempty"`

	// Specific settings for an Npm remote repository.
	// Structure is documented below.
	NpmRepository []NpmRepositoryObservation `json:"npmRepository,omitempty" tf:"npm_repository,omitempty"`

	// Specific settings for a Python remote repository.
	// Structure is documented below.
	PythonRepository []PythonRepositoryObservation `json:"pythonRepository,omitempty" tf:"python_repository,omitempty"`
}

type RemoteRepositoryConfigParameters struct {

	// The description of the remote source.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specific settings for a Docker remote repository.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DockerRepository []DockerRepositoryParameters `json:"dockerRepository,omitempty" tf:"docker_repository,omitempty"`

	// Specific settings for a Maven remote repository.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MavenRepository []MavenRepositoryParameters `json:"mavenRepository,omitempty" tf:"maven_repository,omitempty"`

	// Specific settings for an Npm remote repository.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NpmRepository []NpmRepositoryParameters `json:"npmRepository,omitempty" tf:"npm_repository,omitempty"`

	// Specific settings for a Python remote repository.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PythonRepository []PythonRepositoryParameters `json:"pythonRepository,omitempty" tf:"python_repository,omitempty"`
}

type UpstreamPoliciesInitParameters struct {

	// The user-provided ID of the upstream policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Entries with a greater priority value take precedence in the pull order.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`
}

type UpstreamPoliciesObservation struct {

	// The user-provided ID of the upstream policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Entries with a greater priority value take precedence in the pull order.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// A reference to the repository resource, for example:
	// "projects/p1/locations/us-central1/repository/repo1".
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`
}

type UpstreamPoliciesParameters struct {

	// The user-provided ID of the upstream policy.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Entries with a greater priority value take precedence in the pull order.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// A reference to the repository resource, for example:
	// "projects/p1/locations/us-central1/repository/repo1".
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/artifact/v1beta1.RegistryRepository
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a RegistryRepository in artifact to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a RegistryRepository in artifact to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`
}

type VirtualRepositoryConfigInitParameters struct {

	// Policies that configure the upstream artifacts distributed by the Virtual
	// Repository. Upstream policies cannot be set on a standard repository.
	// Structure is documented below.
	UpstreamPolicies []UpstreamPoliciesInitParameters `json:"upstreamPolicies,omitempty" tf:"upstream_policies,omitempty"`
}

type VirtualRepositoryConfigObservation struct {

	// Policies that configure the upstream artifacts distributed by the Virtual
	// Repository. Upstream policies cannot be set on a standard repository.
	// Structure is documented below.
	UpstreamPolicies []UpstreamPoliciesObservation `json:"upstreamPolicies,omitempty" tf:"upstream_policies,omitempty"`
}

type VirtualRepositoryConfigParameters struct {

	// Policies that configure the upstream artifacts distributed by the Virtual
	// Repository. Upstream policies cannot be set on a standard repository.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	UpstreamPolicies []UpstreamPoliciesParameters `json:"upstreamPolicies,omitempty" tf:"upstream_policies,omitempty"`
}

// RegistryRepositorySpec defines the desired state of RegistryRepository
type RegistryRepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryRepositoryParameters `json:"forProvider"`
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
	InitProvider RegistryRepositoryInitParameters `json:"initProvider,omitempty"`
}

// RegistryRepositoryStatus defines the observed state of RegistryRepository.
type RegistryRepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryRepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryRepository is the Schema for the RegistryRepositorys API. A repository for storing artifacts
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RegistryRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.format) || has(self.initProvider.format)",message="format is a required parameter"
	Spec   RegistryRepositorySpec   `json:"spec"`
	Status RegistryRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryRepositoryList contains a list of RegistryRepositorys
type RegistryRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryRepository `json:"items"`
}

// Repository type metadata.
var (
	RegistryRepository_Kind             = "RegistryRepository"
	RegistryRepository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegistryRepository_Kind}.String()
	RegistryRepository_KindAPIVersion   = RegistryRepository_Kind + "." + CRDGroupVersion.String()
	RegistryRepository_GroupVersionKind = CRDGroupVersion.WithKind(RegistryRepository_Kind)
)

func init() {
	SchemeBuilder.Register(&RegistryRepository{}, &RegistryRepositoryList{})
}
