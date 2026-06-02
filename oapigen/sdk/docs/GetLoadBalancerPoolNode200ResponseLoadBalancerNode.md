# GetLoadBalancerPoolNode200ResponseLoadBalancerNode

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
**Server** | Pointer to [**GetLoadBalancerPoolNode200ResponseLoadBalancerNodeServer**](GetLoadBalancerPoolNode200ResponseLoadBalancerNodeServer.md) |  | [optional] 
**InstanceId** | Pointer to **NullableInt64** |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**NodeSource** | Pointer to **NullableString** |  | [optional] 
**Monitor** | Pointer to [**GetLoadBalancerPoolNode200ResponseLoadBalancerNodeMonitor**](GetLoadBalancerPoolNode200ResponseLoadBalancerNodeMonitor.md) |  | [optional] 
**MaxConnections** | Pointer to **NullableInt64** |  | [optional] 
**ExternalRefType** | Pointer to **NullableString** |  | [optional] 
**ExternalRefId** | Pointer to **NullableString** |  | [optional] 
**ExternalRefName** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**GetLoadBalancerPoolNode200ResponseLoadBalancerNodeCreatedBy**](GetLoadBalancerPoolNode200ResponseLoadBalancerNodeCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetLoadBalancerPoolNode200ResponseLoadBalancerNode

`func NewGetLoadBalancerPoolNode200ResponseLoadBalancerNode() *GetLoadBalancerPoolNode200ResponseLoadBalancerNode`

NewGetLoadBalancerPoolNode200ResponseLoadBalancerNode instantiates a new GetLoadBalancerPoolNode200ResponseLoadBalancerNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisibility

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpAddress

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPort

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortType

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### SetPortTypeNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetPortTypeNil(b bool)`

 SetPortTypeNil sets the value for PortType to be an explicit nil

### UnsetPortType
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetPortType()`

UnsetPortType ensures that no value is present for PortType, not even an explicit nil
### GetMonitorPort

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetMonitorPort() string`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetMonitorPortOk() (*string, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetMonitorPort(v string)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil
### GetWeight

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetNodeState

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetNodeState() string`

GetNodeState returns the NodeState field if non-nil, zero value otherwise.

### GetNodeStateOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetNodeStateOk() (*string, bool)`

GetNodeStateOk returns a tuple with the NodeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeState

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetNodeState(v string)`

SetNodeState sets NodeState field to given value.

### HasNodeState

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasNodeState() bool`

HasNodeState returns a boolean if a field has been set.

### SetNodeStateNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetNodeStateNil(b bool)`

 SetNodeStateNil sets the value for NodeState to be an explicit nil

### UnsetNodeState
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetNodeState()`

UnsetNodeState ensures that no value is present for NodeState, not even an explicit nil
### GetInternalId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetEnabled

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetServer

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetServer() GetLoadBalancerPoolNode200ResponseLoadBalancerNodeServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetServerOk() (*GetLoadBalancerPoolNode200ResponseLoadBalancerNodeServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetServer(v GetLoadBalancerPoolNode200ResponseLoadBalancerNodeServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetInstanceId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetContainerId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetNodeSource

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetNodeSource() string`

GetNodeSource returns the NodeSource field if non-nil, zero value otherwise.

### GetNodeSourceOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetNodeSourceOk() (*string, bool)`

GetNodeSourceOk returns a tuple with the NodeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSource

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetNodeSource(v string)`

SetNodeSource sets NodeSource field to given value.

### HasNodeSource

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasNodeSource() bool`

HasNodeSource returns a boolean if a field has been set.

### SetNodeSourceNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetNodeSourceNil(b bool)`

 SetNodeSourceNil sets the value for NodeSource to be an explicit nil

### UnsetNodeSource
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetNodeSource()`

UnsetNodeSource ensures that no value is present for NodeSource, not even an explicit nil
### GetMonitor

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetMonitor() GetLoadBalancerPoolNode200ResponseLoadBalancerNodeMonitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetMonitorOk() (*GetLoadBalancerPoolNode200ResponseLoadBalancerNodeMonitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetMonitor(v GetLoadBalancerPoolNode200ResponseLoadBalancerNodeMonitor)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetMaxConnections

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### SetMaxConnectionsNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetMaxConnectionsNil(b bool)`

 SetMaxConnectionsNil sets the value for MaxConnections to be an explicit nil

### UnsetMaxConnections
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetMaxConnections()`

UnsetMaxConnections ensures that no value is present for MaxConnections, not even an explicit nil
### GetExternalRefType

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetExternalRefType() string`

GetExternalRefType returns the ExternalRefType field if non-nil, zero value otherwise.

### GetExternalRefTypeOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetExternalRefTypeOk() (*string, bool)`

GetExternalRefTypeOk returns a tuple with the ExternalRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefType

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetExternalRefType(v string)`

SetExternalRefType sets ExternalRefType field to given value.

### HasExternalRefType

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasExternalRefType() bool`

HasExternalRefType returns a boolean if a field has been set.

### SetExternalRefTypeNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetExternalRefTypeNil(b bool)`

 SetExternalRefTypeNil sets the value for ExternalRefType to be an explicit nil

### UnsetExternalRefType
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetExternalRefType()`

UnsetExternalRefType ensures that no value is present for ExternalRefType, not even an explicit nil
### GetExternalRefId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetExternalRefId() string`

GetExternalRefId returns the ExternalRefId field if non-nil, zero value otherwise.

### GetExternalRefIdOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetExternalRefIdOk() (*string, bool)`

GetExternalRefIdOk returns a tuple with the ExternalRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetExternalRefId(v string)`

SetExternalRefId sets ExternalRefId field to given value.

### HasExternalRefId

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasExternalRefId() bool`

HasExternalRefId returns a boolean if a field has been set.

### SetExternalRefIdNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetExternalRefIdNil(b bool)`

 SetExternalRefIdNil sets the value for ExternalRefId to be an explicit nil

### UnsetExternalRefId
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetExternalRefId()`

UnsetExternalRefId ensures that no value is present for ExternalRefId, not even an explicit nil
### GetExternalRefName

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetExternalRefName() string`

GetExternalRefName returns the ExternalRefName field if non-nil, zero value otherwise.

### GetExternalRefNameOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetExternalRefNameOk() (*string, bool)`

GetExternalRefNameOk returns a tuple with the ExternalRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefName

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetExternalRefName(v string)`

SetExternalRefName sets ExternalRefName field to given value.

### HasExternalRefName

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasExternalRefName() bool`

HasExternalRefName returns a boolean if a field has been set.

### SetExternalRefNameNil

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetExternalRefNameNil(b bool)`

 SetExternalRefNameNil sets the value for ExternalRefName to be an explicit nil

### UnsetExternalRefName
`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) UnsetExternalRefName()`

UnsetExternalRefName ensures that no value is present for ExternalRefName, not even an explicit nil
### GetCreatedBy

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetCreatedBy() GetLoadBalancerPoolNode200ResponseLoadBalancerNodeCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetCreatedByOk() (*GetLoadBalancerPoolNode200ResponseLoadBalancerNodeCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetCreatedBy(v GetLoadBalancerPoolNode200ResponseLoadBalancerNodeCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetLoadBalancerPoolNode200ResponseLoadBalancerNode) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


