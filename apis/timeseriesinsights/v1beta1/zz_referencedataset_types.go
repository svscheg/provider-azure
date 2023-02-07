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

type KeyPropertyObservation struct {
}

type KeyPropertyParameters struct {

	// The name of the key property. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The data type of the key property. Valid values include Bool, DateTime, Double, String. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ReferenceDataSetObservation struct {

	// The ID of the IoT Time Series Insights Reference Data Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ReferenceDataSetParameters struct {

	// The comparison behavior that will be used to compare keys. Valid values include Ordinal and OrdinalIgnoreCase. Defaults to Ordinal. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataStringComparisonBehavior *string `json:"dataStringComparisonBehavior,omitempty" tf:"data_string_comparison_behavior,omitempty"`

	// A key_property block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	KeyProperty []KeyPropertyParameters `json:"keyProperty" tf:"key_property,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The resource ID of the Azure IoT Time Series Insights Environment in which to create the Azure IoT Time Series Insights Reference Data Set. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/timeseriesinsights/v1beta1.StandardEnvironment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TimeSeriesInsightsEnvironmentID *string `json:"timeSeriesInsightsEnvironmentId,omitempty" tf:"time_series_insights_environment_id,omitempty"`

	// Reference to a StandardEnvironment in timeseriesinsights to populate timeSeriesInsightsEnvironmentId.
	// +kubebuilder:validation:Optional
	TimeSeriesInsightsEnvironmentIDRef *v1.Reference `json:"timeSeriesInsightsEnvironmentIdRef,omitempty" tf:"-"`

	// Selector for a StandardEnvironment in timeseriesinsights to populate timeSeriesInsightsEnvironmentId.
	// +kubebuilder:validation:Optional
	TimeSeriesInsightsEnvironmentIDSelector *v1.Selector `json:"timeSeriesInsightsEnvironmentIdSelector,omitempty" tf:"-"`
}

// ReferenceDataSetSpec defines the desired state of ReferenceDataSet
type ReferenceDataSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReferenceDataSetParameters `json:"forProvider"`
}

// ReferenceDataSetStatus defines the observed state of ReferenceDataSet.
type ReferenceDataSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReferenceDataSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReferenceDataSet is the Schema for the ReferenceDataSets API. Manages an Azure IoT Time Series Insights Reference Data Set.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ReferenceDataSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReferenceDataSetSpec   `json:"spec"`
	Status            ReferenceDataSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReferenceDataSetList contains a list of ReferenceDataSets
type ReferenceDataSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReferenceDataSet `json:"items"`
}

// Repository type metadata.
var (
	ReferenceDataSet_Kind             = "ReferenceDataSet"
	ReferenceDataSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReferenceDataSet_Kind}.String()
	ReferenceDataSet_KindAPIVersion   = ReferenceDataSet_Kind + "." + CRDGroupVersion.String()
	ReferenceDataSet_GroupVersionKind = CRDGroupVersion.WithKind(ReferenceDataSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ReferenceDataSet{}, &ReferenceDataSetList{})
}