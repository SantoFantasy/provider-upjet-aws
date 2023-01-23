/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SamplingRuleObservation struct {

	// The ARN of the sampling rule.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the sampling rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SamplingRuleParameters struct {

	// Matches attributes derived from the request.
	// +kubebuilder:validation:Optional
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// The percentage of matching requests to instrument, after the reservoir is exhausted.
	// +kubebuilder:validation:Required
	FixedRate *float64 `json:"fixedRate" tf:"fixed_rate,omitempty"`

	// Matches the HTTP method of a request.
	// +kubebuilder:validation:Required
	HTTPMethod *string `json:"httpMethod" tf:"http_method,omitempty"`

	// Matches the hostname from a request URL.
	// +kubebuilder:validation:Required
	Host *string `json:"host" tf:"host,omitempty"`

	// The priority of the sampling rule.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.
	// +kubebuilder:validation:Required
	ReservoirSize *float64 `json:"reservoirSize" tf:"reservoir_size,omitempty"`

	// Matches the ARN of the AWS resource on which the service runs.
	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`

	// Matches the name that the service uses to identify itself in segments.
	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`

	// Matches the origin that the service uses to identify its type in segments.
	// +kubebuilder:validation:Required
	ServiceType *string `json:"serviceType" tf:"service_type,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Matches the path from a request URL.
	// +kubebuilder:validation:Required
	URLPath *string `json:"urlPath" tf:"url_path,omitempty"`

	// The version of the sampling rule format (1 )
	// +kubebuilder:validation:Required
	Version *float64 `json:"version" tf:"version,omitempty"`
}

// SamplingRuleSpec defines the desired state of SamplingRule
type SamplingRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SamplingRuleParameters `json:"forProvider"`
}

// SamplingRuleStatus defines the observed state of SamplingRule.
type SamplingRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SamplingRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SamplingRule is the Schema for the SamplingRules API. Creates and manages an AWS XRay Sampling Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SamplingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SamplingRuleSpec   `json:"spec"`
	Status            SamplingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SamplingRuleList contains a list of SamplingRules
type SamplingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SamplingRule `json:"items"`
}

// Repository type metadata.
var (
	SamplingRule_Kind             = "SamplingRule"
	SamplingRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SamplingRule_Kind}.String()
	SamplingRule_KindAPIVersion   = SamplingRule_Kind + "." + CRDGroupVersion.String()
	SamplingRule_GroupVersionKind = CRDGroupVersion.WithKind(SamplingRule_Kind)
)

func init() {
	SchemeBuilder.Register(&SamplingRule{}, &SamplingRuleList{})
}