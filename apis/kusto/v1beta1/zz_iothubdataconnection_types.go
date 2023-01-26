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

type IOTHubDataConnectionObservation struct {

	// The ID of the Kusto IotHub Data Connection.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTHubDataConnectionParameters struct {

	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name,omitempty"`

	// Reference to a Cluster in kusto to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// Selector for a Cluster in kusto to populate clusterName.
	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// Specifies the IotHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHubConsumerGroup
	// +kubebuilder:validation:Optional
	ConsumerGroup *string `json:"consumerGroup,omitempty" tf:"consumer_group,omitempty"`

	// Reference to a IOTHubConsumerGroup in devices to populate consumerGroup.
	// +kubebuilder:validation:Optional
	ConsumerGroupRef *v1.Reference `json:"consumerGroupRef,omitempty" tf:"-"`

	// Selector for a IOTHubConsumerGroup in devices to populate consumerGroup.
	// +kubebuilder:validation:Optional
	ConsumerGroupSelector *v1.Selector `json:"consumerGroupSelector,omitempty" tf:"-"`

	// Specifies the data format of the IoTHub messages. Allowed values: APACHEAVRO, AVRO, CSV, JSON, MULTIJSON, ORC, PARQUET, PSV, RAW, SCSV, SINGLEJSON, SOHSV, TSV, TSVE, TXT and W3CLOGFILE. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/kusto/v1beta1.Database
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Reference to a Database in kusto to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameRef *v1.Reference `json:"databaseNameRef,omitempty" tf:"-"`

	// Selector for a Database in kusto to populate databaseName.
	// +kubebuilder:validation:Optional
	DatabaseNameSelector *v1.Selector `json:"databaseNameSelector,omitempty" tf:"-"`

	// Indication for database routing information from the data connection, by default only database routing information is allowed. Allowed values: Single, Multi. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DatabaseRoutingType *string `json:"databaseRoutingType,omitempty" tf:"database_routing_type,omitempty"`

	// Specifies the System Properties that each IoT Hub message should contain. Changing this forces a new resource to be created. Possible values are message-id, sequence-number, to, absolute-expiry-time, iothub-enqueuedtime, correlation-id, user-id, iothub-ack, iothub-connection-device-id, iothub-connection-auth-generation-id and iothub-connection-auth-method.
	// +kubebuilder:validation:Optional
	EventSystemProperties []*string `json:"eventSystemProperties,omitempty" tf:"event_system_properties,omitempty"`

	// Specifies the resource id of the IotHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHub
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IOTHubID *string `json:"iothubId,omitempty" tf:"iothub_id,omitempty"`

	// Reference to a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDRef *v1.Reference `json:"iothubIdRef,omitempty" tf:"-"`

	// Selector for a IOTHub in devices to populate iothubId.
	// +kubebuilder:validation:Optional
	IOTHubIDSelector *v1.Selector `json:"iothubIdSelector,omitempty" tf:"-"`

	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MappingRuleName *string `json:"mappingRuleName,omitempty" tf:"mapping_rule_name,omitempty"`

	// The name of the Kusto IotHub Data Connection to create. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the IotHub Shared Access Policy this data connection will use for ingestion, which must have read permission. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devices/v1beta1.IOTHubSharedAccessPolicy
	// +kubebuilder:validation:Optional
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName,omitempty" tf:"shared_access_policy_name,omitempty"`

	// Reference to a IOTHubSharedAccessPolicy in devices to populate sharedAccessPolicyName.
	// +kubebuilder:validation:Optional
	SharedAccessPolicyNameRef *v1.Reference `json:"sharedAccessPolicyNameRef,omitempty" tf:"-"`

	// Selector for a IOTHubSharedAccessPolicy in devices to populate sharedAccessPolicyName.
	// +kubebuilder:validation:Optional
	SharedAccessPolicyNameSelector *v1.Selector `json:"sharedAccessPolicyNameSelector,omitempty" tf:"-"`

	// Specifies the target table name used for the message ingestion. Table must exist before resource is created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

// IOTHubDataConnectionSpec defines the desired state of IOTHubDataConnection
type IOTHubDataConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubDataConnectionParameters `json:"forProvider"`
}

// IOTHubDataConnectionStatus defines the observed state of IOTHubDataConnection.
type IOTHubDataConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubDataConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubDataConnection is the Schema for the IOTHubDataConnections API. Manages Kusto / Data Explorer IotHub Data Connection
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubDataConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubDataConnectionSpec   `json:"spec"`
	Status            IOTHubDataConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubDataConnectionList contains a list of IOTHubDataConnections
type IOTHubDataConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubDataConnection `json:"items"`
}

// Repository type metadata.
var (
	IOTHubDataConnection_Kind             = "IOTHubDataConnection"
	IOTHubDataConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubDataConnection_Kind}.String()
	IOTHubDataConnection_KindAPIVersion   = IOTHubDataConnection_Kind + "." + CRDGroupVersion.String()
	IOTHubDataConnection_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubDataConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubDataConnection{}, &IOTHubDataConnectionList{})
}