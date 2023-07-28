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

type BigQueryFieldFieldInitParameters struct {

	// The resource name of the info type. Set by the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BigQueryFieldFieldObservation struct {

	// The resource name of the info type. Set by the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BigQueryFieldFieldParameters struct {

	// The resource name of the info type. Set by the server.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type BigQueryFieldInitParameters struct {

	// Designated field in the BigQuery table.
	// Structure is documented below.
	Field []BigQueryFieldFieldInitParameters `json:"field,omitempty" tf:"field,omitempty"`

	// Field in a BigQuery table where each cell represents a dictionary phrase.
	// Structure is documented below.
	Table []BigQueryFieldTableInitParameters `json:"table,omitempty" tf:"table,omitempty"`
}

type BigQueryFieldObservation struct {

	// Designated field in the BigQuery table.
	// Structure is documented below.
	Field []BigQueryFieldFieldObservation `json:"field,omitempty" tf:"field,omitempty"`

	// Field in a BigQuery table where each cell represents a dictionary phrase.
	// Structure is documented below.
	Table []BigQueryFieldTableObservation `json:"table,omitempty" tf:"table,omitempty"`
}

type BigQueryFieldParameters struct {

	// Designated field in the BigQuery table.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Field []BigQueryFieldFieldParameters `json:"field,omitempty" tf:"field,omitempty"`

	// Field in a BigQuery table where each cell represents a dictionary phrase.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Table []BigQueryFieldTableParameters `json:"table,omitempty" tf:"table,omitempty"`
}

type BigQueryFieldTableInitParameters struct {

	// The dataset ID of the table.
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// The Google Cloud Platform project ID of the project containing the table.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The name of the table.
	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`
}

type BigQueryFieldTableObservation struct {

	// The dataset ID of the table.
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// The Google Cloud Platform project ID of the project containing the table.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The name of the table.
	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`
}

type BigQueryFieldTableParameters struct {

	// The dataset ID of the table.
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// The Google Cloud Platform project ID of the project containing the table.
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The name of the table.
	// +kubebuilder:validation:Optional
	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`
}

type CloudStorageFileSetInitParameters struct {

	// The url, in the format gs://<bucket>/<path>. Trailing wildcard in the path is allowed.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type CloudStorageFileSetObservation struct {

	// The url, in the format gs://<bucket>/<path>. Trailing wildcard in the path is allowed.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type CloudStorageFileSetParameters struct {

	// The url, in the format gs://<bucket>/<path>. Trailing wildcard in the path is allowed.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type LargeCustomDictionaryInitParameters struct {

	// Field in a BigQuery table where each cell represents a dictionary phrase.
	// Structure is documented below.
	BigQueryField []BigQueryFieldInitParameters `json:"bigQueryField,omitempty" tf:"big_query_field,omitempty"`

	// Set of files containing newline-delimited lists of dictionary phrases.
	// Structure is documented below.
	CloudStorageFileSet []CloudStorageFileSetInitParameters `json:"cloudStorageFileSet,omitempty" tf:"cloud_storage_file_set,omitempty"`

	// Location to store dictionary artifacts in Google Cloud Storage. These files will only be accessible by project owners and the DLP API.
	// If any of these artifacts are modified, the dictionary is considered invalid and can no longer be used.
	// Structure is documented below.
	OutputPath []OutputPathInitParameters `json:"outputPath,omitempty" tf:"output_path,omitempty"`
}

type LargeCustomDictionaryObservation struct {

	// Field in a BigQuery table where each cell represents a dictionary phrase.
	// Structure is documented below.
	BigQueryField []BigQueryFieldObservation `json:"bigQueryField,omitempty" tf:"big_query_field,omitempty"`

	// Set of files containing newline-delimited lists of dictionary phrases.
	// Structure is documented below.
	CloudStorageFileSet []CloudStorageFileSetObservation `json:"cloudStorageFileSet,omitempty" tf:"cloud_storage_file_set,omitempty"`

	// Location to store dictionary artifacts in Google Cloud Storage. These files will only be accessible by project owners and the DLP API.
	// If any of these artifacts are modified, the dictionary is considered invalid and can no longer be used.
	// Structure is documented below.
	OutputPath []OutputPathObservation `json:"outputPath,omitempty" tf:"output_path,omitempty"`
}

type LargeCustomDictionaryParameters struct {

	// Field in a BigQuery table where each cell represents a dictionary phrase.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	BigQueryField []BigQueryFieldParameters `json:"bigQueryField,omitempty" tf:"big_query_field,omitempty"`

	// Set of files containing newline-delimited lists of dictionary phrases.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CloudStorageFileSet []CloudStorageFileSetParameters `json:"cloudStorageFileSet,omitempty" tf:"cloud_storage_file_set,omitempty"`

	// Location to store dictionary artifacts in Google Cloud Storage. These files will only be accessible by project owners and the DLP API.
	// If any of these artifacts are modified, the dictionary is considered invalid and can no longer be used.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	OutputPath []OutputPathParameters `json:"outputPath,omitempty" tf:"output_path,omitempty"`
}

type OutputPathInitParameters struct {

	// A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type OutputPathObservation struct {

	// A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type OutputPathParameters struct {

	// A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type StoredInfoTypeDictionaryCloudStoragePathInitParameters struct {

	// A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type StoredInfoTypeDictionaryCloudStoragePathObservation struct {

	// A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type StoredInfoTypeDictionaryCloudStoragePathParameters struct {

	// A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type StoredInfoTypeDictionaryInitParameters struct {

	// Newline-delimited file of words in Cloud Storage. Only a single file is accepted.
	// Structure is documented below.
	CloudStoragePath []StoredInfoTypeDictionaryCloudStoragePathInitParameters `json:"cloudStoragePath,omitempty" tf:"cloud_storage_path,omitempty"`

	// List of words or phrases to search for.
	// Structure is documented below.
	WordList []StoredInfoTypeDictionaryWordListInitParameters `json:"wordList,omitempty" tf:"word_list,omitempty"`
}

type StoredInfoTypeDictionaryObservation struct {

	// Newline-delimited file of words in Cloud Storage. Only a single file is accepted.
	// Structure is documented below.
	CloudStoragePath []StoredInfoTypeDictionaryCloudStoragePathObservation `json:"cloudStoragePath,omitempty" tf:"cloud_storage_path,omitempty"`

	// List of words or phrases to search for.
	// Structure is documented below.
	WordList []StoredInfoTypeDictionaryWordListObservation `json:"wordList,omitempty" tf:"word_list,omitempty"`
}

type StoredInfoTypeDictionaryParameters struct {

	// Newline-delimited file of words in Cloud Storage. Only a single file is accepted.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	CloudStoragePath []StoredInfoTypeDictionaryCloudStoragePathParameters `json:"cloudStoragePath,omitempty" tf:"cloud_storage_path,omitempty"`

	// List of words or phrases to search for.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	WordList []StoredInfoTypeDictionaryWordListParameters `json:"wordList,omitempty" tf:"word_list,omitempty"`
}

type StoredInfoTypeDictionaryWordListInitParameters struct {

	// Words or phrases defining the dictionary. The dictionary must contain at least one
	// phrase and every phrase must contain at least 2 characters that are letters or digits.
	Words []*string `json:"words,omitempty" tf:"words,omitempty"`
}

type StoredInfoTypeDictionaryWordListObservation struct {

	// Words or phrases defining the dictionary. The dictionary must contain at least one
	// phrase and every phrase must contain at least 2 characters that are letters or digits.
	Words []*string `json:"words,omitempty" tf:"words,omitempty"`
}

type StoredInfoTypeDictionaryWordListParameters struct {

	// Words or phrases defining the dictionary. The dictionary must contain at least one
	// phrase and every phrase must contain at least 2 characters that are letters or digits.
	// +kubebuilder:validation:Optional
	Words []*string `json:"words,omitempty" tf:"words,omitempty"`
}

type StoredInfoTypeInitParameters struct {

	// A description of the info type.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary []StoredInfoTypeDictionaryInitParameters `json:"dictionary,omitempty" tf:"dictionary,omitempty"`

	// User set display name of the info type.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary []LargeCustomDictionaryInitParameters `json:"largeCustomDictionary,omitempty" tf:"large_custom_dictionary,omitempty"`

	// The parent of the info type in any of the following formats:
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex []StoredInfoTypeRegexInitParameters `json:"regex,omitempty" tf:"regex,omitempty"`
}

type StoredInfoTypeObservation struct {

	// A description of the info type.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Dictionary which defines the rule.
	// Structure is documented below.
	Dictionary []StoredInfoTypeDictionaryObservation `json:"dictionary,omitempty" tf:"dictionary,omitempty"`

	// User set display name of the info type.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format {{parent}}/storedInfoTypes/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Dictionary which defines the rule.
	// Structure is documented below.
	LargeCustomDictionary []LargeCustomDictionaryObservation `json:"largeCustomDictionary,omitempty" tf:"large_custom_dictionary,omitempty"`

	// The resource name of the info type. Set by the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The parent of the info type in any of the following formats:
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Regular expression which defines the rule.
	// Structure is documented below.
	Regex []StoredInfoTypeRegexObservation `json:"regex,omitempty" tf:"regex,omitempty"`
}

type StoredInfoTypeParameters struct {

	// A description of the info type.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Dictionary which defines the rule.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Dictionary []StoredInfoTypeDictionaryParameters `json:"dictionary,omitempty" tf:"dictionary,omitempty"`

	// User set display name of the info type.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Dictionary which defines the rule.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	LargeCustomDictionary []LargeCustomDictionaryParameters `json:"largeCustomDictionary,omitempty" tf:"large_custom_dictionary,omitempty"`

	// The parent of the info type in any of the following formats:
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`

	// Regular expression which defines the rule.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Regex []StoredInfoTypeRegexParameters `json:"regex,omitempty" tf:"regex,omitempty"`
}

type StoredInfoTypeRegexInitParameters struct {

	// The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.
	GroupIndexes []*float64 `json:"groupIndexes,omitempty" tf:"group_indexes,omitempty"`

	// Pattern defining the regular expression.
	// Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type StoredInfoTypeRegexObservation struct {

	// The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.
	GroupIndexes []*float64 `json:"groupIndexes,omitempty" tf:"group_indexes,omitempty"`

	// Pattern defining the regular expression.
	// Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type StoredInfoTypeRegexParameters struct {

	// The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.
	// +kubebuilder:validation:Optional
	GroupIndexes []*float64 `json:"groupIndexes,omitempty" tf:"group_indexes,omitempty"`

	// Pattern defining the regular expression.
	// Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

// StoredInfoTypeSpec defines the desired state of StoredInfoType
type StoredInfoTypeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StoredInfoTypeParameters `json:"forProvider"`
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
	InitProvider StoredInfoTypeInitParameters `json:"initProvider,omitempty"`
}

// StoredInfoTypeStatus defines the observed state of StoredInfoType.
type StoredInfoTypeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StoredInfoTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StoredInfoType is the Schema for the StoredInfoTypes API. Allows creation of custom info types.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type StoredInfoType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.parent) || has(self.initProvider.parent)",message="parent is a required parameter"
	Spec   StoredInfoTypeSpec   `json:"spec"`
	Status StoredInfoTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoredInfoTypeList contains a list of StoredInfoTypes
type StoredInfoTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoredInfoType `json:"items"`
}

// Repository type metadata.
var (
	StoredInfoType_Kind             = "StoredInfoType"
	StoredInfoType_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StoredInfoType_Kind}.String()
	StoredInfoType_KindAPIVersion   = StoredInfoType_Kind + "." + CRDGroupVersion.String()
	StoredInfoType_GroupVersionKind = CRDGroupVersion.WithKind(StoredInfoType_Kind)
)

func init() {
	SchemeBuilder.Register(&StoredInfoType{}, &StoredInfoTypeList{})
}
