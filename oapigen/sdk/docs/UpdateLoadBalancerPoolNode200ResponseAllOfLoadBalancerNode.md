# UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode

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
**Server** | Pointer to [**UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeServer**](UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeServer.md) |  | [optional] 
**InstanceId** | Pointer to **NullableInt64** |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**NodeSource** | Pointer to **NullableString** |  | [optional] 
**Monitor** | Pointer to [**UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeMonitor**](UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeMonitor.md) |  | [optional] 
**MaxConnections** | Pointer to **NullableInt64** |  | [optional] 
**ExternalRefType** | Pointer to **NullableString** |  | [optional] 
**ExternalRefId** | Pointer to **NullableString** |  | [optional] 
**ExternalRefName** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeCreatedBy**](UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode

`func NewUpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode() *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode`

NewUpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode instantiates a new UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpAddress

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortType

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### SetPortTypeNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetPortTypeNil(b bool)`

 SetPortTypeNil sets the value for PortType to be an explicit nil

### UnsetPortType
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetPortType()`

UnsetPortType ensures that no value is present for PortType, not even an explicit nil
### GetMonitorPort

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetMonitorPort() string`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetMonitorPortOk() (*string, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetMonitorPort(v string)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil
### GetWeight

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetNodeState

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetNodeState() string`

GetNodeState returns the NodeState field if non-nil, zero value otherwise.

### GetNodeStateOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetNodeStateOk() (*string, bool)`

GetNodeStateOk returns a tuple with the NodeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeState

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetNodeState(v string)`

SetNodeState sets NodeState field to given value.

### HasNodeState

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasNodeState() bool`

HasNodeState returns a boolean if a field has been set.

### SetNodeStateNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetNodeStateNil(b bool)`

 SetNodeStateNil sets the value for NodeState to be an explicit nil

### UnsetNodeState
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetNodeState()`

UnsetNodeState ensures that no value is present for NodeState, not even an explicit nil
### GetInternalId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetEnabled

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetServer

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetServer() UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetServerOk() (*UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetServer(v UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetInstanceId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetContainerId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetNodeSource

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetNodeSource() string`

GetNodeSource returns the NodeSource field if non-nil, zero value otherwise.

### GetNodeSourceOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetNodeSourceOk() (*string, bool)`

GetNodeSourceOk returns a tuple with the NodeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSource

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetNodeSource(v string)`

SetNodeSource sets NodeSource field to given value.

### HasNodeSource

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasNodeSource() bool`

HasNodeSource returns a boolean if a field has been set.

### SetNodeSourceNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetNodeSourceNil(b bool)`

 SetNodeSourceNil sets the value for NodeSource to be an explicit nil

### UnsetNodeSource
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetNodeSource()`

UnsetNodeSource ensures that no value is present for NodeSource, not even an explicit nil
### GetMonitor

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetMonitor() UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeMonitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetMonitorOk() (*UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeMonitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetMonitor(v UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeMonitor)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetMaxConnections

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### SetMaxConnectionsNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetMaxConnectionsNil(b bool)`

 SetMaxConnectionsNil sets the value for MaxConnections to be an explicit nil

### UnsetMaxConnections
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetMaxConnections()`

UnsetMaxConnections ensures that no value is present for MaxConnections, not even an explicit nil
### GetExternalRefType

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetExternalRefType() string`

GetExternalRefType returns the ExternalRefType field if non-nil, zero value otherwise.

### GetExternalRefTypeOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetExternalRefTypeOk() (*string, bool)`

GetExternalRefTypeOk returns a tuple with the ExternalRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefType

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetExternalRefType(v string)`

SetExternalRefType sets ExternalRefType field to given value.

### HasExternalRefType

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasExternalRefType() bool`

HasExternalRefType returns a boolean if a field has been set.

### SetExternalRefTypeNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetExternalRefTypeNil(b bool)`

 SetExternalRefTypeNil sets the value for ExternalRefType to be an explicit nil

### UnsetExternalRefType
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetExternalRefType()`

UnsetExternalRefType ensures that no value is present for ExternalRefType, not even an explicit nil
### GetExternalRefId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetExternalRefId() string`

GetExternalRefId returns the ExternalRefId field if non-nil, zero value otherwise.

### GetExternalRefIdOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetExternalRefIdOk() (*string, bool)`

GetExternalRefIdOk returns a tuple with the ExternalRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetExternalRefId(v string)`

SetExternalRefId sets ExternalRefId field to given value.

### HasExternalRefId

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasExternalRefId() bool`

HasExternalRefId returns a boolean if a field has been set.

### SetExternalRefIdNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetExternalRefIdNil(b bool)`

 SetExternalRefIdNil sets the value for ExternalRefId to be an explicit nil

### UnsetExternalRefId
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetExternalRefId()`

UnsetExternalRefId ensures that no value is present for ExternalRefId, not even an explicit nil
### GetExternalRefName

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetExternalRefName() string`

GetExternalRefName returns the ExternalRefName field if non-nil, zero value otherwise.

### GetExternalRefNameOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetExternalRefNameOk() (*string, bool)`

GetExternalRefNameOk returns a tuple with the ExternalRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefName

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetExternalRefName(v string)`

SetExternalRefName sets ExternalRefName field to given value.

### HasExternalRefName

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasExternalRefName() bool`

HasExternalRefName returns a boolean if a field has been set.

### SetExternalRefNameNil

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetExternalRefNameNil(b bool)`

 SetExternalRefNameNil sets the value for ExternalRefName to be an explicit nil

### UnsetExternalRefName
`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) UnsetExternalRefName()`

UnsetExternalRefName ensures that no value is present for ExternalRefName, not even an explicit nil
### GetCreatedBy

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetCreatedBy() UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetCreatedByOk() (*UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetCreatedBy(v UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNodeCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateLoadBalancerPoolNode200ResponseAllOfLoadBalancerNode) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


