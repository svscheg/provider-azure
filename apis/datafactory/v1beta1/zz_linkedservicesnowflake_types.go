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

type LinkedServiceSnowflakeKeyVaultPasswordObservation struct {
}

type LinkedServiceSnowflakeKeyVaultPasswordParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceKeyVault
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceKeyVault in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// Specifies the secret name in Azure Key Vault that stores Snowflake password.
	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type LinkedServiceSnowflakeObservation struct {

	// The ID of the Data Factory Snowflake Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LinkedServiceSnowflakeParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The connection string in which to authenticate with Snowflake.
	// +kubebuilder:validation:Required
	ConnectionString *string `json:"connectionString" tf:"connection_string,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_password block as defined below. Use this argument to store Snowflake password in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	KeyVaultPassword []LinkedServiceSnowflakeKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// LinkedServiceSnowflakeSpec defines the desired state of LinkedServiceSnowflake
type LinkedServiceSnowflakeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceSnowflakeParameters `json:"forProvider"`
}

// LinkedServiceSnowflakeStatus defines the observed state of LinkedServiceSnowflake.
type LinkedServiceSnowflakeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceSnowflakeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceSnowflake is the Schema for the LinkedServiceSnowflakes API. Manages a Linked Service (connection) between Snowflake and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceSnowflake struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkedServiceSnowflakeSpec   `json:"spec"`
	Status            LinkedServiceSnowflakeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceSnowflakeList contains a list of LinkedServiceSnowflakes
type LinkedServiceSnowflakeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceSnowflake `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceSnowflake_Kind             = "LinkedServiceSnowflake"
	LinkedServiceSnowflake_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceSnowflake_Kind}.String()
	LinkedServiceSnowflake_KindAPIVersion   = LinkedServiceSnowflake_Kind + "." + CRDGroupVersion.String()
	LinkedServiceSnowflake_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceSnowflake_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceSnowflake{}, &LinkedServiceSnowflakeList{})
}