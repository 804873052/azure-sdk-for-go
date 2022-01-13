//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

const (
	moduleName    = "armeventgrid"
	moduleVersion = "v0.2.0"
)

// AdvancedFilterOperatorType - The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
type AdvancedFilterOperatorType string

const (
	AdvancedFilterOperatorTypeBoolEquals                AdvancedFilterOperatorType = "BoolEquals"
	AdvancedFilterOperatorTypeIsNotNull                 AdvancedFilterOperatorType = "IsNotNull"
	AdvancedFilterOperatorTypeIsNullOrUndefined         AdvancedFilterOperatorType = "IsNullOrUndefined"
	AdvancedFilterOperatorTypeNumberGreaterThan         AdvancedFilterOperatorType = "NumberGreaterThan"
	AdvancedFilterOperatorTypeNumberGreaterThanOrEquals AdvancedFilterOperatorType = "NumberGreaterThanOrEquals"
	AdvancedFilterOperatorTypeNumberIn                  AdvancedFilterOperatorType = "NumberIn"
	AdvancedFilterOperatorTypeNumberInRange             AdvancedFilterOperatorType = "NumberInRange"
	AdvancedFilterOperatorTypeNumberLessThan            AdvancedFilterOperatorType = "NumberLessThan"
	AdvancedFilterOperatorTypeNumberLessThanOrEquals    AdvancedFilterOperatorType = "NumberLessThanOrEquals"
	AdvancedFilterOperatorTypeNumberNotIn               AdvancedFilterOperatorType = "NumberNotIn"
	AdvancedFilterOperatorTypeNumberNotInRange          AdvancedFilterOperatorType = "NumberNotInRange"
	AdvancedFilterOperatorTypeStringBeginsWith          AdvancedFilterOperatorType = "StringBeginsWith"
	AdvancedFilterOperatorTypeStringContains            AdvancedFilterOperatorType = "StringContains"
	AdvancedFilterOperatorTypeStringEndsWith            AdvancedFilterOperatorType = "StringEndsWith"
	AdvancedFilterOperatorTypeStringIn                  AdvancedFilterOperatorType = "StringIn"
	AdvancedFilterOperatorTypeStringNotBeginsWith       AdvancedFilterOperatorType = "StringNotBeginsWith"
	AdvancedFilterOperatorTypeStringNotContains         AdvancedFilterOperatorType = "StringNotContains"
	AdvancedFilterOperatorTypeStringNotEndsWith         AdvancedFilterOperatorType = "StringNotEndsWith"
	AdvancedFilterOperatorTypeStringNotIn               AdvancedFilterOperatorType = "StringNotIn"
)

// PossibleAdvancedFilterOperatorTypeValues returns the possible values for the AdvancedFilterOperatorType const type.
func PossibleAdvancedFilterOperatorTypeValues() []AdvancedFilterOperatorType {
	return []AdvancedFilterOperatorType{
		AdvancedFilterOperatorTypeBoolEquals,
		AdvancedFilterOperatorTypeIsNotNull,
		AdvancedFilterOperatorTypeIsNullOrUndefined,
		AdvancedFilterOperatorTypeNumberGreaterThan,
		AdvancedFilterOperatorTypeNumberGreaterThanOrEquals,
		AdvancedFilterOperatorTypeNumberIn,
		AdvancedFilterOperatorTypeNumberInRange,
		AdvancedFilterOperatorTypeNumberLessThan,
		AdvancedFilterOperatorTypeNumberLessThanOrEquals,
		AdvancedFilterOperatorTypeNumberNotIn,
		AdvancedFilterOperatorTypeNumberNotInRange,
		AdvancedFilterOperatorTypeStringBeginsWith,
		AdvancedFilterOperatorTypeStringContains,
		AdvancedFilterOperatorTypeStringEndsWith,
		AdvancedFilterOperatorTypeStringIn,
		AdvancedFilterOperatorTypeStringNotBeginsWith,
		AdvancedFilterOperatorTypeStringNotContains,
		AdvancedFilterOperatorTypeStringNotEndsWith,
		AdvancedFilterOperatorTypeStringNotIn,
	}
}

// ToPtr returns a *AdvancedFilterOperatorType pointing to the current value.
func (c AdvancedFilterOperatorType) ToPtr() *AdvancedFilterOperatorType {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// DeadLetterEndPointType - Type of the endpoint for the dead letter destination
type DeadLetterEndPointType string

const (
	DeadLetterEndPointTypeStorageBlob DeadLetterEndPointType = "StorageBlob"
)

// PossibleDeadLetterEndPointTypeValues returns the possible values for the DeadLetterEndPointType const type.
func PossibleDeadLetterEndPointTypeValues() []DeadLetterEndPointType {
	return []DeadLetterEndPointType{
		DeadLetterEndPointTypeStorageBlob,
	}
}

// ToPtr returns a *DeadLetterEndPointType pointing to the current value.
func (c DeadLetterEndPointType) ToPtr() *DeadLetterEndPointType {
	return &c
}

// DeliveryAttributeMappingType - Type of the delivery attribute or header name.
type DeliveryAttributeMappingType string

const (
	DeliveryAttributeMappingTypeDynamic DeliveryAttributeMappingType = "Dynamic"
	DeliveryAttributeMappingTypeStatic  DeliveryAttributeMappingType = "Static"
)

// PossibleDeliveryAttributeMappingTypeValues returns the possible values for the DeliveryAttributeMappingType const type.
func PossibleDeliveryAttributeMappingTypeValues() []DeliveryAttributeMappingType {
	return []DeliveryAttributeMappingType{
		DeliveryAttributeMappingTypeDynamic,
		DeliveryAttributeMappingTypeStatic,
	}
}

// ToPtr returns a *DeliveryAttributeMappingType pointing to the current value.
func (c DeliveryAttributeMappingType) ToPtr() *DeliveryAttributeMappingType {
	return &c
}

// DomainProvisioningState - Provisioning state of the Event Grid Domain Resource.
type DomainProvisioningState string

const (
	DomainProvisioningStateCanceled  DomainProvisioningState = "Canceled"
	DomainProvisioningStateCreating  DomainProvisioningState = "Creating"
	DomainProvisioningStateDeleting  DomainProvisioningState = "Deleting"
	DomainProvisioningStateFailed    DomainProvisioningState = "Failed"
	DomainProvisioningStateSucceeded DomainProvisioningState = "Succeeded"
	DomainProvisioningStateUpdating  DomainProvisioningState = "Updating"
)

// PossibleDomainProvisioningStateValues returns the possible values for the DomainProvisioningState const type.
func PossibleDomainProvisioningStateValues() []DomainProvisioningState {
	return []DomainProvisioningState{
		DomainProvisioningStateCanceled,
		DomainProvisioningStateCreating,
		DomainProvisioningStateDeleting,
		DomainProvisioningStateFailed,
		DomainProvisioningStateSucceeded,
		DomainProvisioningStateUpdating,
	}
}

// ToPtr returns a *DomainProvisioningState pointing to the current value.
func (c DomainProvisioningState) ToPtr() *DomainProvisioningState {
	return &c
}

// DomainTopicProvisioningState - Provisioning state of the domain topic.
type DomainTopicProvisioningState string

const (
	DomainTopicProvisioningStateCanceled  DomainTopicProvisioningState = "Canceled"
	DomainTopicProvisioningStateCreating  DomainTopicProvisioningState = "Creating"
	DomainTopicProvisioningStateDeleting  DomainTopicProvisioningState = "Deleting"
	DomainTopicProvisioningStateFailed    DomainTopicProvisioningState = "Failed"
	DomainTopicProvisioningStateSucceeded DomainTopicProvisioningState = "Succeeded"
	DomainTopicProvisioningStateUpdating  DomainTopicProvisioningState = "Updating"
)

// PossibleDomainTopicProvisioningStateValues returns the possible values for the DomainTopicProvisioningState const type.
func PossibleDomainTopicProvisioningStateValues() []DomainTopicProvisioningState {
	return []DomainTopicProvisioningState{
		DomainTopicProvisioningStateCanceled,
		DomainTopicProvisioningStateCreating,
		DomainTopicProvisioningStateDeleting,
		DomainTopicProvisioningStateFailed,
		DomainTopicProvisioningStateSucceeded,
		DomainTopicProvisioningStateUpdating,
	}
}

// ToPtr returns a *DomainTopicProvisioningState pointing to the current value.
func (c DomainTopicProvisioningState) ToPtr() *DomainTopicProvisioningState {
	return &c
}

// EndpointType - Type of the endpoint for the event subscription destination.
type EndpointType string

const (
	EndpointTypeAzureFunction    EndpointType = "AzureFunction"
	EndpointTypeEventHub         EndpointType = "EventHub"
	EndpointTypeHybridConnection EndpointType = "HybridConnection"
	EndpointTypeServiceBusQueue  EndpointType = "ServiceBusQueue"
	EndpointTypeServiceBusTopic  EndpointType = "ServiceBusTopic"
	EndpointTypeStorageQueue     EndpointType = "StorageQueue"
	EndpointTypeWebHook          EndpointType = "WebHook"
)

// PossibleEndpointTypeValues returns the possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{
		EndpointTypeAzureFunction,
		EndpointTypeEventHub,
		EndpointTypeHybridConnection,
		EndpointTypeServiceBusQueue,
		EndpointTypeServiceBusTopic,
		EndpointTypeStorageQueue,
		EndpointTypeWebHook,
	}
}

// ToPtr returns a *EndpointType pointing to the current value.
func (c EndpointType) ToPtr() *EndpointType {
	return &c
}

type Enum18 string

const (
	Enum18Domains Enum18 = "domains"
	Enum18Topics  Enum18 = "topics"
)

// PossibleEnum18Values returns the possible values for the Enum18 const type.
func PossibleEnum18Values() []Enum18 {
	return []Enum18{
		Enum18Domains,
		Enum18Topics,
	}
}

// ToPtr returns a *Enum18 pointing to the current value.
func (c Enum18) ToPtr() *Enum18 {
	return &c
}

type Enum19 string

const (
	Enum19Domains Enum19 = "domains"
	Enum19Topics  Enum19 = "topics"
)

// PossibleEnum19Values returns the possible values for the Enum19 const type.
func PossibleEnum19Values() []Enum19 {
	return []Enum19{
		Enum19Domains,
		Enum19Topics,
	}
}

// ToPtr returns a *Enum19 pointing to the current value.
func (c Enum19) ToPtr() *Enum19 {
	return &c
}

type Enum20 string

const (
	Enum20Domains Enum20 = "domains"
	Enum20Topics  Enum20 = "topics"
)

// PossibleEnum20Values returns the possible values for the Enum20 const type.
func PossibleEnum20Values() []Enum20 {
	return []Enum20{
		Enum20Domains,
		Enum20Topics,
	}
}

// ToPtr returns a *Enum20 pointing to the current value.
func (c Enum20) ToPtr() *Enum20 {
	return &c
}

type Enum21 string

const (
	Enum21Domains Enum21 = "domains"
	Enum21Topics  Enum21 = "topics"
)

// PossibleEnum21Values returns the possible values for the Enum21 const type.
func PossibleEnum21Values() []Enum21 {
	return []Enum21{
		Enum21Domains,
		Enum21Topics,
	}
}

// ToPtr returns a *Enum21 pointing to the current value.
func (c Enum21) ToPtr() *Enum21 {
	return &c
}

// EventDeliverySchema - The event delivery schema for the event subscription.
type EventDeliverySchema string

const (
	EventDeliverySchemaCloudEventSchemaV10 EventDeliverySchema = "CloudEventSchemaV1_0"
	EventDeliverySchemaCustomInputSchema   EventDeliverySchema = "CustomInputSchema"
	EventDeliverySchemaEventGridSchema     EventDeliverySchema = "EventGridSchema"
)

// PossibleEventDeliverySchemaValues returns the possible values for the EventDeliverySchema const type.
func PossibleEventDeliverySchemaValues() []EventDeliverySchema {
	return []EventDeliverySchema{
		EventDeliverySchemaCloudEventSchemaV10,
		EventDeliverySchemaCustomInputSchema,
		EventDeliverySchemaEventGridSchema,
	}
}

// ToPtr returns a *EventDeliverySchema pointing to the current value.
func (c EventDeliverySchema) ToPtr() *EventDeliverySchema {
	return &c
}

// EventSubscriptionIdentityType - The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both
// an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identity.
type EventSubscriptionIdentityType string

const (
	EventSubscriptionIdentityTypeSystemAssigned EventSubscriptionIdentityType = "SystemAssigned"
	EventSubscriptionIdentityTypeUserAssigned   EventSubscriptionIdentityType = "UserAssigned"
)

// PossibleEventSubscriptionIdentityTypeValues returns the possible values for the EventSubscriptionIdentityType const type.
func PossibleEventSubscriptionIdentityTypeValues() []EventSubscriptionIdentityType {
	return []EventSubscriptionIdentityType{
		EventSubscriptionIdentityTypeSystemAssigned,
		EventSubscriptionIdentityTypeUserAssigned,
	}
}

// ToPtr returns a *EventSubscriptionIdentityType pointing to the current value.
func (c EventSubscriptionIdentityType) ToPtr() *EventSubscriptionIdentityType {
	return &c
}

// EventSubscriptionProvisioningState - Provisioning state of the event subscription.
type EventSubscriptionProvisioningState string

const (
	EventSubscriptionProvisioningStateAwaitingManualAction EventSubscriptionProvisioningState = "AwaitingManualAction"
	EventSubscriptionProvisioningStateCanceled             EventSubscriptionProvisioningState = "Canceled"
	EventSubscriptionProvisioningStateCreating             EventSubscriptionProvisioningState = "Creating"
	EventSubscriptionProvisioningStateDeleting             EventSubscriptionProvisioningState = "Deleting"
	EventSubscriptionProvisioningStateFailed               EventSubscriptionProvisioningState = "Failed"
	EventSubscriptionProvisioningStateSucceeded            EventSubscriptionProvisioningState = "Succeeded"
	EventSubscriptionProvisioningStateUpdating             EventSubscriptionProvisioningState = "Updating"
)

// PossibleEventSubscriptionProvisioningStateValues returns the possible values for the EventSubscriptionProvisioningState const type.
func PossibleEventSubscriptionProvisioningStateValues() []EventSubscriptionProvisioningState {
	return []EventSubscriptionProvisioningState{
		EventSubscriptionProvisioningStateAwaitingManualAction,
		EventSubscriptionProvisioningStateCanceled,
		EventSubscriptionProvisioningStateCreating,
		EventSubscriptionProvisioningStateDeleting,
		EventSubscriptionProvisioningStateFailed,
		EventSubscriptionProvisioningStateSucceeded,
		EventSubscriptionProvisioningStateUpdating,
	}
}

// ToPtr returns a *EventSubscriptionProvisioningState pointing to the current value.
func (c EventSubscriptionProvisioningState) ToPtr() *EventSubscriptionProvisioningState {
	return &c
}

// IPActionType - Action to perform based on the match or no match of the IpMask.
type IPActionType string

const (
	IPActionTypeAllow IPActionType = "Allow"
)

// PossibleIPActionTypeValues returns the possible values for the IPActionType const type.
func PossibleIPActionTypeValues() []IPActionType {
	return []IPActionType{
		IPActionTypeAllow,
	}
}

// ToPtr returns a *IPActionType pointing to the current value.
func (c IPActionType) ToPtr() *IPActionType {
	return &c
}

// IdentityType - The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created
// identity and a set of user-assigned identities. The type 'None' will remove any identity.
type IdentityType string

const (
	IdentityTypeNone                       IdentityType = "None"
	IdentityTypeSystemAssigned             IdentityType = "SystemAssigned"
	IdentityTypeSystemAssignedUserAssigned IdentityType = "SystemAssigned, UserAssigned"
	IdentityTypeUserAssigned               IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeSystemAssignedUserAssigned,
		IdentityTypeUserAssigned,
	}
}

// ToPtr returns a *IdentityType pointing to the current value.
func (c IdentityType) ToPtr() *IdentityType {
	return &c
}

// InputSchema - This determines the format that Event Grid should expect for incoming events published to the domain.
type InputSchema string

const (
	InputSchemaCloudEventSchemaV10 InputSchema = "CloudEventSchemaV1_0"
	InputSchemaCustomEventSchema   InputSchema = "CustomEventSchema"
	InputSchemaEventGridSchema     InputSchema = "EventGridSchema"
)

// PossibleInputSchemaValues returns the possible values for the InputSchema const type.
func PossibleInputSchemaValues() []InputSchema {
	return []InputSchema{
		InputSchemaCloudEventSchemaV10,
		InputSchemaCustomEventSchema,
		InputSchemaEventGridSchema,
	}
}

// ToPtr returns a *InputSchema pointing to the current value.
func (c InputSchema) ToPtr() *InputSchema {
	return &c
}

// InputSchemaMappingType - Type of the custom mapping
type InputSchemaMappingType string

const (
	InputSchemaMappingTypeJSON InputSchemaMappingType = "Json"
)

// PossibleInputSchemaMappingTypeValues returns the possible values for the InputSchemaMappingType const type.
func PossibleInputSchemaMappingTypeValues() []InputSchemaMappingType {
	return []InputSchemaMappingType{
		InputSchemaMappingTypeJSON,
	}
}

// ToPtr returns a *InputSchemaMappingType pointing to the current value.
func (c InputSchemaMappingType) ToPtr() *InputSchemaMappingType {
	return &c
}

// PersistedConnectionStatus - Status of the connection.
type PersistedConnectionStatus string

const (
	PersistedConnectionStatusApproved     PersistedConnectionStatus = "Approved"
	PersistedConnectionStatusDisconnected PersistedConnectionStatus = "Disconnected"
	PersistedConnectionStatusPending      PersistedConnectionStatus = "Pending"
	PersistedConnectionStatusRejected     PersistedConnectionStatus = "Rejected"
)

// PossiblePersistedConnectionStatusValues returns the possible values for the PersistedConnectionStatus const type.
func PossiblePersistedConnectionStatusValues() []PersistedConnectionStatus {
	return []PersistedConnectionStatus{
		PersistedConnectionStatusApproved,
		PersistedConnectionStatusDisconnected,
		PersistedConnectionStatusPending,
		PersistedConnectionStatusRejected,
	}
}

// ToPtr returns a *PersistedConnectionStatus pointing to the current value.
func (c PersistedConnectionStatus) ToPtr() *PersistedConnectionStatus {
	return &c
}

// PublicNetworkAccess - This determines if traffic is allowed over public network. By default it is enabled. You can further
// restrict to specific IPs by configuring
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ToPtr returns a *PublicNetworkAccess pointing to the current value.
func (c PublicNetworkAccess) ToPtr() *PublicNetworkAccess {
	return &c
}

// ResourceProvisioningState - Provisioning state of the Private Endpoint Connection.
type ResourceProvisioningState string

const (
	ResourceProvisioningStateCanceled  ResourceProvisioningState = "Canceled"
	ResourceProvisioningStateCreating  ResourceProvisioningState = "Creating"
	ResourceProvisioningStateDeleting  ResourceProvisioningState = "Deleting"
	ResourceProvisioningStateFailed    ResourceProvisioningState = "Failed"
	ResourceProvisioningStateSucceeded ResourceProvisioningState = "Succeeded"
	ResourceProvisioningStateUpdating  ResourceProvisioningState = "Updating"
)

// PossibleResourceProvisioningStateValues returns the possible values for the ResourceProvisioningState const type.
func PossibleResourceProvisioningStateValues() []ResourceProvisioningState {
	return []ResourceProvisioningState{
		ResourceProvisioningStateCanceled,
		ResourceProvisioningStateCreating,
		ResourceProvisioningStateDeleting,
		ResourceProvisioningStateFailed,
		ResourceProvisioningStateSucceeded,
		ResourceProvisioningStateUpdating,
	}
}

// ToPtr returns a *ResourceProvisioningState pointing to the current value.
func (c ResourceProvisioningState) ToPtr() *ResourceProvisioningState {
	return &c
}

// ResourceRegionType - Region type of the resource.
type ResourceRegionType string

const (
	ResourceRegionTypeGlobalResource   ResourceRegionType = "GlobalResource"
	ResourceRegionTypeRegionalResource ResourceRegionType = "RegionalResource"
)

// PossibleResourceRegionTypeValues returns the possible values for the ResourceRegionType const type.
func PossibleResourceRegionTypeValues() []ResourceRegionType {
	return []ResourceRegionType{
		ResourceRegionTypeGlobalResource,
		ResourceRegionTypeRegionalResource,
	}
}

// ToPtr returns a *ResourceRegionType pointing to the current value.
func (c ResourceRegionType) ToPtr() *ResourceRegionType {
	return &c
}

// TopicProvisioningState - Provisioning state of the topic.
type TopicProvisioningState string

const (
	TopicProvisioningStateCanceled  TopicProvisioningState = "Canceled"
	TopicProvisioningStateCreating  TopicProvisioningState = "Creating"
	TopicProvisioningStateDeleting  TopicProvisioningState = "Deleting"
	TopicProvisioningStateFailed    TopicProvisioningState = "Failed"
	TopicProvisioningStateSucceeded TopicProvisioningState = "Succeeded"
	TopicProvisioningStateUpdating  TopicProvisioningState = "Updating"
)

// PossibleTopicProvisioningStateValues returns the possible values for the TopicProvisioningState const type.
func PossibleTopicProvisioningStateValues() []TopicProvisioningState {
	return []TopicProvisioningState{
		TopicProvisioningStateCanceled,
		TopicProvisioningStateCreating,
		TopicProvisioningStateDeleting,
		TopicProvisioningStateFailed,
		TopicProvisioningStateSucceeded,
		TopicProvisioningStateUpdating,
	}
}

// ToPtr returns a *TopicProvisioningState pointing to the current value.
func (c TopicProvisioningState) ToPtr() *TopicProvisioningState {
	return &c
}

type TopicTypePropertiesSupportedScopesForSourceItem string

const (
	TopicTypePropertiesSupportedScopesForSourceItemAzureSubscription TopicTypePropertiesSupportedScopesForSourceItem = "AzureSubscription"
	TopicTypePropertiesSupportedScopesForSourceItemResource          TopicTypePropertiesSupportedScopesForSourceItem = "Resource"
	TopicTypePropertiesSupportedScopesForSourceItemResourceGroup     TopicTypePropertiesSupportedScopesForSourceItem = "ResourceGroup"
)

// PossibleTopicTypePropertiesSupportedScopesForSourceItemValues returns the possible values for the TopicTypePropertiesSupportedScopesForSourceItem const type.
func PossibleTopicTypePropertiesSupportedScopesForSourceItemValues() []TopicTypePropertiesSupportedScopesForSourceItem {
	return []TopicTypePropertiesSupportedScopesForSourceItem{
		TopicTypePropertiesSupportedScopesForSourceItemAzureSubscription,
		TopicTypePropertiesSupportedScopesForSourceItemResource,
		TopicTypePropertiesSupportedScopesForSourceItemResourceGroup,
	}
}

// ToPtr returns a *TopicTypePropertiesSupportedScopesForSourceItem pointing to the current value.
func (c TopicTypePropertiesSupportedScopesForSourceItem) ToPtr() *TopicTypePropertiesSupportedScopesForSourceItem {
	return &c
}

// TopicTypeProvisioningState - Provisioning state of the topic type
type TopicTypeProvisioningState string

const (
	TopicTypeProvisioningStateCanceled  TopicTypeProvisioningState = "Canceled"
	TopicTypeProvisioningStateCreating  TopicTypeProvisioningState = "Creating"
	TopicTypeProvisioningStateDeleting  TopicTypeProvisioningState = "Deleting"
	TopicTypeProvisioningStateFailed    TopicTypeProvisioningState = "Failed"
	TopicTypeProvisioningStateSucceeded TopicTypeProvisioningState = "Succeeded"
	TopicTypeProvisioningStateUpdating  TopicTypeProvisioningState = "Updating"
)

// PossibleTopicTypeProvisioningStateValues returns the possible values for the TopicTypeProvisioningState const type.
func PossibleTopicTypeProvisioningStateValues() []TopicTypeProvisioningState {
	return []TopicTypeProvisioningState{
		TopicTypeProvisioningStateCanceled,
		TopicTypeProvisioningStateCreating,
		TopicTypeProvisioningStateDeleting,
		TopicTypeProvisioningStateFailed,
		TopicTypeProvisioningStateSucceeded,
		TopicTypeProvisioningStateUpdating,
	}
}

// ToPtr returns a *TopicTypeProvisioningState pointing to the current value.
func (c TopicTypeProvisioningState) ToPtr() *TopicTypeProvisioningState {
	return &c
}
