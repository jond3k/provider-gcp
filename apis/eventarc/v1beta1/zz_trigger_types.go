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

type CloudRunServiceObservation struct {

	// Optional. The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Required. The region the Cloud Run service is deployed in.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Required. Name of the GKE service.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type CloudRunServiceParameters struct {

	// Optional. The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Required. The region the Cloud Run service is deployed in.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Required. Name of the GKE service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/cloudrun/v1beta1.Service
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// Reference to a Service in cloudrun to populate service.
	// +kubebuilder:validation:Optional
	ServiceRef *v1.Reference `json:"serviceRef,omitempty" tf:"-"`

	// Selector for a Service in cloudrun to populate service.
	// +kubebuilder:validation:Optional
	ServiceSelector *v1.Selector `json:"serviceSelector,omitempty" tf:"-"`
}

type DestinationObservation struct {

	// [WARNING] Configuring a Cloud Function in Trigger is not supported as of today. The Cloud Function resource name. Format: projects/{project}/locations/{location}/functions/{function}
	CloudFunction *string `json:"cloudFunction,omitempty" tf:"cloud_function,omitempty"`

	// Cloud Run fully-managed service that receives the events. The service should be running in the same project of the trigger.
	CloudRunService []CloudRunServiceObservation `json:"cloudRunService,omitempty" tf:"cloud_run_service,omitempty"`

	// A GKE service capable of receiving events. The service should be running in the same project as the trigger.
	Gke []GkeObservation `json:"gke,omitempty" tf:"gke,omitempty"`

	// The resource name of the Workflow whose Executions are triggered by the events. The Workflow resource should be deployed in the same project as the trigger. Format: projects/{project}/locations/{location}/workflows/{workflow}
	Workflow *string `json:"workflow,omitempty" tf:"workflow,omitempty"`
}

type DestinationParameters struct {

	// [WARNING] Configuring a Cloud Function in Trigger is not supported as of today. The Cloud Function resource name. Format: projects/{project}/locations/{location}/functions/{function}
	// +kubebuilder:validation:Optional
	CloudFunction *string `json:"cloudFunction,omitempty" tf:"cloud_function,omitempty"`

	// Cloud Run fully-managed service that receives the events. The service should be running in the same project of the trigger.
	// +kubebuilder:validation:Optional
	CloudRunService []CloudRunServiceParameters `json:"cloudRunService,omitempty" tf:"cloud_run_service,omitempty"`

	// A GKE service capable of receiving events. The service should be running in the same project as the trigger.
	// +kubebuilder:validation:Optional
	Gke []GkeParameters `json:"gke,omitempty" tf:"gke,omitempty"`

	// The resource name of the Workflow whose Executions are triggered by the events. The Workflow resource should be deployed in the same project as the trigger. Format: projects/{project}/locations/{location}/workflows/{workflow}
	// +kubebuilder:validation:Optional
	Workflow *string `json:"workflow,omitempty" tf:"workflow,omitempty"`
}

type GkeObservation struct {

	// Required. The name of the cluster the GKE service is running in. The cluster must be running in the same project as the trigger being created.
	Cluster *string `json:"cluster,omitempty" tf:"cluster,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Required. The namespace the GKE service is running in.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// Optional. The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Required. Name of the GKE service.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type GkeParameters struct {

	// Required. The name of the cluster the GKE service is running in. The cluster must be running in the same project as the trigger being created.
	// +kubebuilder:validation:Required
	Cluster *string `json:"cluster" tf:"cluster,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Required. The namespace the GKE service is running in.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`

	// Optional. The relative path on the GKE service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Required. Name of the GKE service.
	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

type MatchingCriteriaObservation struct {

	// Required. The name of a CloudEvents attribute. Currently, only a subset of attributes are supported for filtering. All triggers MUST provide a filter for the 'type' attribute.
	Attribute *string `json:"attribute,omitempty" tf:"attribute,omitempty"`

	// Optional. The operator used for matching the events with the value of the filter. If not specified, only events that have an exact key-value pair specified in the filter are matched. The only allowed value is match-path-pattern.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Required. The value for the attribute. See https://cloud.google.com/eventarc/docs/creating-triggers#trigger-gcloud for available values.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MatchingCriteriaParameters struct {

	// Required. The name of a CloudEvents attribute. Currently, only a subset of attributes are supported for filtering. All triggers MUST provide a filter for the 'type' attribute.
	// +kubebuilder:validation:Required
	Attribute *string `json:"attribute" tf:"attribute,omitempty"`

	// Optional. The operator used for matching the events with the value of the filter. If not specified, only events that have an exact key-value pair specified in the filter are matched. The only allowed value is match-path-pattern.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Required. The value for the attribute. See https://cloud.google.com/eventarc/docs/creating-triggers#trigger-gcloud for available values.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type PubsubObservation struct {

	// Output only. The name of the Pub/Sub subscription created and managed by Eventarc system as a transport for the event delivery. Format: projects/{PROJECT_ID}/subscriptions/{SUBSCRIPTION_NAME}.
	Subscription *string `json:"subscription,omitempty" tf:"subscription,omitempty"`

	// Optional. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: projects/{PROJECT_ID}/topics/{TOPIC_NAME}. You may set an existing topic for triggers of the type google.cloud.pubsub.topic.v1.messagePublished only. The topic you provide here will not be deleted by Eventarc at trigger deletion.
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type PubsubParameters struct {

	// Optional. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: projects/{PROJECT_ID}/topics/{TOPIC_NAME}. You may set an existing topic for triggers of the type google.cloud.pubsub.topic.v1.messagePublished only. The topic you provide here will not be deleted by Eventarc at trigger deletion.
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type TransportObservation struct {

	// The Pub/Sub topic and subscription used by Eventarc as delivery intermediary.
	Pubsub []PubsubObservation `json:"pubsub,omitempty" tf:"pubsub,omitempty"`
}

type TransportParameters struct {

	// The Pub/Sub topic and subscription used by Eventarc as delivery intermediary.
	// +kubebuilder:validation:Optional
	Pubsub []PubsubParameters `json:"pubsub,omitempty" tf:"pubsub,omitempty"`
}

type TriggerObservation struct {

	// Optional. The name of the channel associated with the trigger in projects/{project}/locations/{location}/channels/{channel} format. You must provide a channel to receive events from Eventarc SaaS partners.
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// Output only. The reason(s) why a trigger is in FAILED state.
	Conditions map[string]*string `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Output only. The creation time.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// Required. Destination specifies where the events should be sent to.
	Destination []DestinationObservation `json:"destination,omitempty" tf:"destination,omitempty"`

	// Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/triggers/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional. User labels attached to the triggers that can be used to group resources.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location for the resource
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	MatchingCriteria []MatchingCriteriaObservation `json:"matchingCriteria,omitempty" tf:"matching_criteria,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have iam.serviceAccounts.actAs permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have roles/eventarc.eventReceiver IAM role.
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	Transport []TransportObservation `json:"transport,omitempty" tf:"transport,omitempty"`

	// Output only. Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// Output only. The last-modified time.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type TriggerParameters struct {

	// Optional. The name of the channel associated with the trigger in projects/{project}/locations/{location}/channels/{channel} format. You must provide a channel to receive events from Eventarc SaaS partners.
	// +kubebuilder:validation:Optional
	Channel *string `json:"channel,omitempty" tf:"channel,omitempty"`

	// Required. Destination specifies where the events should be sent to.
	// +kubebuilder:validation:Optional
	Destination []DestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// Optional. User labels attached to the triggers that can be used to group resources.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location for the resource
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	// +kubebuilder:validation:Optional
	MatchingCriteria []MatchingCriteriaParameters `json:"matchingCriteria,omitempty" tf:"matching_criteria,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have iam.serviceAccounts.actAs permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have roles/eventarc.eventReceiver IAM role.
	// +kubebuilder:validation:Optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`

	// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	// +kubebuilder:validation:Optional
	Transport []TransportParameters `json:"transport,omitempty" tf:"transport,omitempty"`
}

// TriggerSpec defines the desired state of Trigger
type TriggerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerParameters `json:"forProvider"`
}

// TriggerStatus defines the observed state of Trigger.
type TriggerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Trigger is the Schema for the Triggers API. The Eventarc Trigger resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Trigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.destination)",message="destination is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.matchingCriteria)",message="matchingCriteria is a required parameter"
	Spec   TriggerSpec   `json:"spec"`
	Status TriggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerList contains a list of Triggers
type TriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trigger `json:"items"`
}

// Repository type metadata.
var (
	Trigger_Kind             = "Trigger"
	Trigger_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Trigger_Kind}.String()
	Trigger_KindAPIVersion   = Trigger_Kind + "." + CRDGroupVersion.String()
	Trigger_GroupVersionKind = CRDGroupVersion.WithKind(Trigger_Kind)
)

func init() {
	SchemeBuilder.Register(&Trigger{}, &TriggerList{})
}
