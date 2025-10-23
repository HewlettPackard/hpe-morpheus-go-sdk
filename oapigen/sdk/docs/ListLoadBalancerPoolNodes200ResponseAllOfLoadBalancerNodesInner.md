# ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**PortType** | Pointer to **NullableString** |  | [optional] 
**MonitorPort** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **NullableInt64** |  | [optional] 
**NodeState** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**Server** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**InstanceId** | Pointer to **NullableInt64** |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**NodeSource** | Pointer to **NullableString** |  | [optional] 
**Monitor** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**MaxConnections** | Pointer to **NullableInt64** |  | [optional] 
**ExternalRefType** | Pointer to **NullableString** |  | [optional] 
**ExternalRefId** | Pointer to **NullableString** |  | [optional] 
**ExternalRefName** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner

`func NewListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner() *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner`

NewListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner instantiates a new ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInnerWithDefaults

`func NewListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInnerWithDefaults() *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner`

NewListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInnerWithDefaults instantiates a new ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisibility

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpAddress

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortType

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### SetPortTypeNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetPortTypeNil(b bool)`

 SetPortTypeNil sets the value for PortType to be an explicit nil

### UnsetPortType
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetPortType()`

UnsetPortType ensures that no value is present for PortType, not even an explicit nil
### GetMonitorPort

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetMonitorPort() string`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetMonitorPortOk() (*string, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetMonitorPort(v string)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil
### GetWeight

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetNodeState

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetNodeState() string`

GetNodeState returns the NodeState field if non-nil, zero value otherwise.

### GetNodeStateOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetNodeStateOk() (*string, bool)`

GetNodeStateOk returns a tuple with the NodeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeState

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetNodeState(v string)`

SetNodeState sets NodeState field to given value.

### HasNodeState

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasNodeState() bool`

HasNodeState returns a boolean if a field has been set.

### SetNodeStateNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetNodeStateNil(b bool)`

 SetNodeStateNil sets the value for NodeState to be an explicit nil

### UnsetNodeState
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetNodeState()`

UnsetNodeState ensures that no value is present for NodeState, not even an explicit nil
### GetInternalId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetEnabled

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetServer

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetServer() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetServerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetServer(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetServer sets Server field to given value.

### HasServer

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetInstanceId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetContainerId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetNodeSource

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetNodeSource() string`

GetNodeSource returns the NodeSource field if non-nil, zero value otherwise.

### GetNodeSourceOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetNodeSourceOk() (*string, bool)`

GetNodeSourceOk returns a tuple with the NodeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSource

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetNodeSource(v string)`

SetNodeSource sets NodeSource field to given value.

### HasNodeSource

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasNodeSource() bool`

HasNodeSource returns a boolean if a field has been set.

### SetNodeSourceNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetNodeSourceNil(b bool)`

 SetNodeSourceNil sets the value for NodeSource to be an explicit nil

### UnsetNodeSource
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetNodeSource()`

UnsetNodeSource ensures that no value is present for NodeSource, not even an explicit nil
### GetMonitor

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetMonitor() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetMonitorOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetMonitor(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetMaxConnections

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### SetMaxConnectionsNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetMaxConnectionsNil(b bool)`

 SetMaxConnectionsNil sets the value for MaxConnections to be an explicit nil

### UnsetMaxConnections
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetMaxConnections()`

UnsetMaxConnections ensures that no value is present for MaxConnections, not even an explicit nil
### GetExternalRefType

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetExternalRefType() string`

GetExternalRefType returns the ExternalRefType field if non-nil, zero value otherwise.

### GetExternalRefTypeOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetExternalRefTypeOk() (*string, bool)`

GetExternalRefTypeOk returns a tuple with the ExternalRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefType

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetExternalRefType(v string)`

SetExternalRefType sets ExternalRefType field to given value.

### HasExternalRefType

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasExternalRefType() bool`

HasExternalRefType returns a boolean if a field has been set.

### SetExternalRefTypeNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetExternalRefTypeNil(b bool)`

 SetExternalRefTypeNil sets the value for ExternalRefType to be an explicit nil

### UnsetExternalRefType
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetExternalRefType()`

UnsetExternalRefType ensures that no value is present for ExternalRefType, not even an explicit nil
### GetExternalRefId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetExternalRefId() string`

GetExternalRefId returns the ExternalRefId field if non-nil, zero value otherwise.

### GetExternalRefIdOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetExternalRefIdOk() (*string, bool)`

GetExternalRefIdOk returns a tuple with the ExternalRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetExternalRefId(v string)`

SetExternalRefId sets ExternalRefId field to given value.

### HasExternalRefId

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasExternalRefId() bool`

HasExternalRefId returns a boolean if a field has been set.

### SetExternalRefIdNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetExternalRefIdNil(b bool)`

 SetExternalRefIdNil sets the value for ExternalRefId to be an explicit nil

### UnsetExternalRefId
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetExternalRefId()`

UnsetExternalRefId ensures that no value is present for ExternalRefId, not even an explicit nil
### GetExternalRefName

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetExternalRefName() string`

GetExternalRefName returns the ExternalRefName field if non-nil, zero value otherwise.

### GetExternalRefNameOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetExternalRefNameOk() (*string, bool)`

GetExternalRefNameOk returns a tuple with the ExternalRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefName

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetExternalRefName(v string)`

SetExternalRefName sets ExternalRefName field to given value.

### HasExternalRefName

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasExternalRefName() bool`

HasExternalRefName returns a boolean if a field has been set.

### SetExternalRefNameNil

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetExternalRefNameNil(b bool)`

 SetExternalRefNameNil sets the value for ExternalRefName to be an explicit nil

### UnsetExternalRefName
`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) UnsetExternalRefName()`

UnsetExternalRefName ensures that no value is present for ExternalRefName, not even an explicit nil
### GetCreatedBy

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


