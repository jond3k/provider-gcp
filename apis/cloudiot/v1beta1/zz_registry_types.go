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

type EventNotificationConfigsObservation struct {

	// PubSub topic name to publish device events.
	PubsubTopicName *string `json:"pubsubTopicName,omitempty" tf:"pubsub_topic_name,omitempty"`

	// If the subfolder name matches this string exactly, this
	// configuration will be used. The string must not include the
	// leading '/' character. If empty, all strings are matched. Empty
	// value can only be used for the last event_notification_configs
	// item.
	SubfolderMatches *string `json:"subfolderMatches,omitempty" tf:"subfolder_matches,omitempty"`
}

type EventNotificationConfigsParameters struct {

	// PubSub topic name to publish device events.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/pubsub/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PubsubTopicName *string `json:"pubsubTopicName,omitempty" tf:"pubsub_topic_name,omitempty"`

	// Reference to a Topic in pubsub to populate pubsubTopicName.
	// +kubebuilder:validation:Optional
	PubsubTopicNameRef *v1.Reference `json:"pubsubTopicNameRef,omitempty" tf:"-"`

	// Selector for a Topic in pubsub to populate pubsubTopicName.
	// +kubebuilder:validation:Optional
	PubsubTopicNameSelector *v1.Selector `json:"pubsubTopicNameSelector,omitempty" tf:"-"`

	// If the subfolder name matches this string exactly, this
	// configuration will be used. The string must not include the
	// leading '/' character. If empty, all strings are matched. Empty
	// value can only be used for the last event_notification_configs
	// item.
	// +kubebuilder:validation:Optional
	SubfolderMatches *string `json:"subfolderMatches,omitempty" tf:"subfolder_matches,omitempty"`
}

type RegistryCredentialsObservation struct {

	// A public key certificate format and data.
	PublicKeyCertificate map[string]string `json:"publicKeyCertificate,omitempty" tf:"public_key_certificate,omitempty"`
}

type RegistryCredentialsParameters struct {

	// A public key certificate format and data.
	// +kubebuilder:validation:Required
	PublicKeyCertificate map[string]string `json:"publicKeyCertificate" tf:"public_key_certificate,omitempty"`
}

type RegistryObservation struct {

	// List of public key certificates to authenticate devices.
	// The structure is documented below.
	Credentials []RegistryCredentialsObservation `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// List of configurations for event notifications, such as PubSub topics
	// to publish device events to.
	// Structure is documented below.
	EventNotificationConfigs []EventNotificationConfigsObservation `json:"eventNotificationConfigs,omitempty" tf:"event_notification_configs,omitempty"`

	// Activate or deactivate HTTP.
	// The structure is documented below.
	HTTPConfig map[string]string `json:"httpConfig,omitempty" tf:"http_config,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/registries/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The default logging verbosity for activity from devices in this
	// registry. Specifies which events should be written to logs. For
	// example, if the LogLevel is ERROR, only events that terminate in
	// errors will be logged. LogLevel is inclusive; enabling INFO logging
	// will also enable ERROR logging.
	// Default value is NONE.
	// Possible values are: NONE, ERROR, INFO, DEBUG.
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`

	// Activate or deactivate MQTT.
	// The structure is documented below.
	MqttConfig map[string]string `json:"mqttConfig,omitempty" tf:"mqtt_config,omitempty"`

	// A unique name for the resource, required by device registry.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region in which the created registry should reside.
	// If it is not provided, the provider region is used.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A PubSub topic to publish device state updates.
	// The structure is documented below.
	StateNotificationConfig map[string]string `json:"stateNotificationConfig,omitempty" tf:"state_notification_config,omitempty"`
}

type RegistryParameters struct {

	// List of public key certificates to authenticate devices.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	Credentials []RegistryCredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// List of configurations for event notifications, such as PubSub topics
	// to publish device events to.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EventNotificationConfigs []EventNotificationConfigsParameters `json:"eventNotificationConfigs,omitempty" tf:"event_notification_configs,omitempty"`

	// Activate or deactivate HTTP.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPConfig map[string]string `json:"httpConfig,omitempty" tf:"http_config,omitempty"`

	// The default logging verbosity for activity from devices in this
	// registry. Specifies which events should be written to logs. For
	// example, if the LogLevel is ERROR, only events that terminate in
	// errors will be logged. LogLevel is inclusive; enabling INFO logging
	// will also enable ERROR logging.
	// Default value is NONE.
	// Possible values are: NONE, ERROR, INFO, DEBUG.
	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`

	// Activate or deactivate MQTT.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	MqttConfig map[string]string `json:"mqttConfig,omitempty" tf:"mqtt_config,omitempty"`

	// A unique name for the resource, required by device registry.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region in which the created registry should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A PubSub topic to publish device state updates.
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	StateNotificationConfig map[string]string `json:"stateNotificationConfig,omitempty" tf:"state_notification_config,omitempty"`
}

// RegistrySpec defines the desired state of Registry
type RegistrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryParameters `json:"forProvider"`
}

// RegistryStatus defines the observed state of Registry.
type RegistryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Registry is the Schema for the Registrys API. A Google Cloud IoT Core device registry.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Registry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name)",message="name is a required parameter"
	Spec   RegistrySpec   `json:"spec"`
	Status RegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryList contains a list of Registrys
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registry `json:"items"`
}

// Repository type metadata.
var (
	Registry_Kind             = "Registry"
	Registry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Registry_Kind}.String()
	Registry_KindAPIVersion   = Registry_Kind + "." + CRDGroupVersion.String()
	Registry_GroupVersionKind = CRDGroupVersion.WithKind(Registry_Kind)
)

func init() {
	SchemeBuilder.Register(&Registry{}, &RegistryList{})
}
