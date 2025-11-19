# InstanceContainerPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**External** | Pointer to **int64** |  | [optional] 
**Internal** | Pointer to **int64** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PrimaryPort** | Pointer to **bool** |  | [optional] 
**Export** | Pointer to **bool** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 
**ExportName** | Pointer to **string** |  | [optional] 
**LoadBalanceProtocol** | Pointer to **NullableString** |  | [optional] 
**LoadBalance** | Pointer to **bool** |  | [optional] 
**Protocol** | Pointer to **NullableString** |  | [optional] 
**Link** | Pointer to **bool** |  | [optional] 
**ExternalIp** | Pointer to **NullableString** |  | [optional] 
**InternalIp** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstanceContainerPortsInner

`func NewInstanceContainerPortsInner() *InstanceContainerPortsInner`

NewInstanceContainerPortsInner instantiates a new InstanceContainerPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainerPortsInnerWithDefaults

`func NewInstanceContainerPortsInnerWithDefaults() *InstanceContainerPortsInner`

NewInstanceContainerPortsInnerWithDefaults instantiates a new InstanceContainerPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainerPortsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainerPortsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainerPortsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainerPortsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternal

`func (o *InstanceContainerPortsInner) GetExternal() int64`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InstanceContainerPortsInner) GetExternalOk() (*int64, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InstanceContainerPortsInner) SetExternal(v int64)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InstanceContainerPortsInner) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetInternal

`func (o *InstanceContainerPortsInner) GetInternal() int64`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *InstanceContainerPortsInner) GetInternalOk() (*int64, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *InstanceContainerPortsInner) SetInternal(v int64)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *InstanceContainerPortsInner) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetDisplayName

`func (o *InstanceContainerPortsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InstanceContainerPortsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InstanceContainerPortsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InstanceContainerPortsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPrimaryPort

`func (o *InstanceContainerPortsInner) GetPrimaryPort() bool`

GetPrimaryPort returns the PrimaryPort field if non-nil, zero value otherwise.

### GetPrimaryPortOk

`func (o *InstanceContainerPortsInner) GetPrimaryPortOk() (*bool, bool)`

GetPrimaryPortOk returns a tuple with the PrimaryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPort

`func (o *InstanceContainerPortsInner) SetPrimaryPort(v bool)`

SetPrimaryPort sets PrimaryPort field to given value.

### HasPrimaryPort

`func (o *InstanceContainerPortsInner) HasPrimaryPort() bool`

HasPrimaryPort returns a boolean if a field has been set.

### GetExport

`func (o *InstanceContainerPortsInner) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *InstanceContainerPortsInner) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *InstanceContainerPortsInner) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *InstanceContainerPortsInner) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetVisible

`func (o *InstanceContainerPortsInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InstanceContainerPortsInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InstanceContainerPortsInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *InstanceContainerPortsInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetExportName

`func (o *InstanceContainerPortsInner) GetExportName() string`

GetExportName returns the ExportName field if non-nil, zero value otherwise.

### GetExportNameOk

`func (o *InstanceContainerPortsInner) GetExportNameOk() (*string, bool)`

GetExportNameOk returns a tuple with the ExportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportName

`func (o *InstanceContainerPortsInner) SetExportName(v string)`

SetExportName sets ExportName field to given value.

### HasExportName

`func (o *InstanceContainerPortsInner) HasExportName() bool`

HasExportName returns a boolean if a field has been set.

### GetLoadBalanceProtocol

`func (o *InstanceContainerPortsInner) GetLoadBalanceProtocol() string`

GetLoadBalanceProtocol returns the LoadBalanceProtocol field if non-nil, zero value otherwise.

### GetLoadBalanceProtocolOk

`func (o *InstanceContainerPortsInner) GetLoadBalanceProtocolOk() (*string, bool)`

GetLoadBalanceProtocolOk returns a tuple with the LoadBalanceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceProtocol

`func (o *InstanceContainerPortsInner) SetLoadBalanceProtocol(v string)`

SetLoadBalanceProtocol sets LoadBalanceProtocol field to given value.

### HasLoadBalanceProtocol

`func (o *InstanceContainerPortsInner) HasLoadBalanceProtocol() bool`

HasLoadBalanceProtocol returns a boolean if a field has been set.

### SetLoadBalanceProtocolNil

`func (o *InstanceContainerPortsInner) SetLoadBalanceProtocolNil(b bool)`

 SetLoadBalanceProtocolNil sets the value for LoadBalanceProtocol to be an explicit nil

### UnsetLoadBalanceProtocol
`func (o *InstanceContainerPortsInner) UnsetLoadBalanceProtocol()`

UnsetLoadBalanceProtocol ensures that no value is present for LoadBalanceProtocol, not even an explicit nil
### GetLoadBalance

`func (o *InstanceContainerPortsInner) GetLoadBalance() bool`

GetLoadBalance returns the LoadBalance field if non-nil, zero value otherwise.

### GetLoadBalanceOk

`func (o *InstanceContainerPortsInner) GetLoadBalanceOk() (*bool, bool)`

GetLoadBalanceOk returns a tuple with the LoadBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalance

`func (o *InstanceContainerPortsInner) SetLoadBalance(v bool)`

SetLoadBalance sets LoadBalance field to given value.

### HasLoadBalance

`func (o *InstanceContainerPortsInner) HasLoadBalance() bool`

HasLoadBalance returns a boolean if a field has been set.

### GetProtocol

`func (o *InstanceContainerPortsInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InstanceContainerPortsInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InstanceContainerPortsInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InstanceContainerPortsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *InstanceContainerPortsInner) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *InstanceContainerPortsInner) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetLink

`func (o *InstanceContainerPortsInner) GetLink() bool`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *InstanceContainerPortsInner) GetLinkOk() (*bool, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *InstanceContainerPortsInner) SetLink(v bool)`

SetLink sets Link field to given value.

### HasLink

`func (o *InstanceContainerPortsInner) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetExternalIp

`func (o *InstanceContainerPortsInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *InstanceContainerPortsInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *InstanceContainerPortsInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *InstanceContainerPortsInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *InstanceContainerPortsInner) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *InstanceContainerPortsInner) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetInternalIp

`func (o *InstanceContainerPortsInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *InstanceContainerPortsInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *InstanceContainerPortsInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *InstanceContainerPortsInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *InstanceContainerPortsInner) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *InstanceContainerPortsInner) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


