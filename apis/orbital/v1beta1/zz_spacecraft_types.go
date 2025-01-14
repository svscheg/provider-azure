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

type SpacecraftLinksObservation struct {

	// Bandwidth in Mhz.
	BandwidthMhz *float64 `json:"bandwidthMhz,omitempty" tf:"bandwidth_mhz,omitempty"`

	// Center frequency in Mhz.
	CenterFrequencyMhz *float64 `json:"centerFrequencyMhz,omitempty" tf:"center_frequency_mhz,omitempty"`

	// Direction if the communication. Possible values are Uplink and Downlink.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Name of the link.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Polarization. Possible values are RHCP, LHCP, linearVertical and linearHorizontal.
	Polarization *string `json:"polarization,omitempty" tf:"polarization,omitempty"`
}

type SpacecraftLinksParameters struct {

	// Bandwidth in Mhz.
	// +kubebuilder:validation:Required
	BandwidthMhz *float64 `json:"bandwidthMhz" tf:"bandwidth_mhz,omitempty"`

	// Center frequency in Mhz.
	// +kubebuilder:validation:Required
	CenterFrequencyMhz *float64 `json:"centerFrequencyMhz" tf:"center_frequency_mhz,omitempty"`

	// Direction if the communication. Possible values are Uplink and Downlink.
	// +kubebuilder:validation:Required
	Direction *string `json:"direction" tf:"direction,omitempty"`

	// Name of the link.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Polarization. Possible values are RHCP, LHCP, linearVertical and linearHorizontal.
	// +kubebuilder:validation:Required
	Polarization *string `json:"polarization" tf:"polarization,omitempty"`
}

type SpacecraftObservation struct {

	// The ID of the Spacecraft.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A links block as defined below. Changing this forces a new resource to be created.
	Links []SpacecraftLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// The location where the Spacecraft exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// NORAD ID of the Spacecraft.
	NoradID *string `json:"noradId,omitempty" tf:"norad_id,omitempty"`

	// The name of the Resource Group where the Spacecraft exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Title of the two line elements (TLE).
	TitleLine *string `json:"titleLine,omitempty" tf:"title_line,omitempty"`

	// A list of the two line elements (TLE), the first string being the first of the TLE, the second string being the second line of the TLE. Changing this forces a new resource to be created.
	TwoLineElements []*string `json:"twoLineElements,omitempty" tf:"two_line_elements,omitempty"`
}

type SpacecraftParameters struct {

	// A links block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Links []SpacecraftLinksParameters `json:"links,omitempty" tf:"links,omitempty"`

	// The location where the Spacecraft exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// NORAD ID of the Spacecraft.
	// +kubebuilder:validation:Optional
	NoradID *string `json:"noradId,omitempty" tf:"norad_id,omitempty"`

	// The name of the Resource Group where the Spacecraft exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Title of the two line elements (TLE).
	// +kubebuilder:validation:Optional
	TitleLine *string `json:"titleLine,omitempty" tf:"title_line,omitempty"`

	// A list of the two line elements (TLE), the first string being the first of the TLE, the second string being the second line of the TLE. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TwoLineElements []*string `json:"twoLineElements,omitempty" tf:"two_line_elements,omitempty"`
}

// SpacecraftSpec defines the desired state of Spacecraft
type SpacecraftSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpacecraftParameters `json:"forProvider"`
}

// SpacecraftStatus defines the observed state of Spacecraft.
type SpacecraftStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpacecraftObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Spacecraft is the Schema for the Spacecrafts API. Manages a Spacecraft resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Spacecraft struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.links)",message="links is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.noradId)",message="noradId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.titleLine)",message="titleLine is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.twoLineElements)",message="twoLineElements is a required parameter"
	Spec   SpacecraftSpec   `json:"spec"`
	Status SpacecraftStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpacecraftList contains a list of Spacecrafts
type SpacecraftList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Spacecraft `json:"items"`
}

// Repository type metadata.
var (
	Spacecraft_Kind             = "Spacecraft"
	Spacecraft_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Spacecraft_Kind}.String()
	Spacecraft_KindAPIVersion   = Spacecraft_Kind + "." + CRDGroupVersion.String()
	Spacecraft_GroupVersionKind = CRDGroupVersion.WithKind(Spacecraft_Kind)
)

func init() {
	SchemeBuilder.Register(&Spacecraft{}, &SpacecraftList{})
}
