# InstanceContainer3PortsInner

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

### NewInstanceContainer3PortsInner

`func NewInstanceContainer3PortsInner() *InstanceContainer3PortsInner`

NewInstanceContainer3PortsInner instantiates a new InstanceContainer3PortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainer3PortsInnerWithDefaults

`func NewInstanceContainer3PortsInnerWithDefaults() *InstanceContainer3PortsInner`

NewInstanceContainer3PortsInnerWithDefaults instantiates a new InstanceContainer3PortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainer3PortsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainer3PortsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainer3PortsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainer3PortsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternal

`func (o *InstanceContainer3PortsInner) GetExternal() int64`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InstanceContainer3PortsInner) GetExternalOk() (*int64, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InstanceContainer3PortsInner) SetExternal(v int64)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InstanceContainer3PortsInner) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetInternal

`func (o *InstanceContainer3PortsInner) GetInternal() int64`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *InstanceContainer3PortsInner) GetInternalOk() (*int64, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *InstanceContainer3PortsInner) SetInternal(v int64)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *InstanceContainer3PortsInner) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetDisplayName

`func (o *InstanceContainer3PortsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InstanceContainer3PortsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InstanceContainer3PortsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InstanceContainer3PortsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPrimaryPort

`func (o *InstanceContainer3PortsInner) GetPrimaryPort() bool`

GetPrimaryPort returns the PrimaryPort field if non-nil, zero value otherwise.

### GetPrimaryPortOk

`func (o *InstanceContainer3PortsInner) GetPrimaryPortOk() (*bool, bool)`

GetPrimaryPortOk returns a tuple with the PrimaryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPort

`func (o *InstanceContainer3PortsInner) SetPrimaryPort(v bool)`

SetPrimaryPort sets PrimaryPort field to given value.

### HasPrimaryPort

`func (o *InstanceContainer3PortsInner) HasPrimaryPort() bool`

HasPrimaryPort returns a boolean if a field has been set.

### GetExport

`func (o *InstanceContainer3PortsInner) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *InstanceContainer3PortsInner) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *InstanceContainer3PortsInner) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *InstanceContainer3PortsInner) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetVisible

`func (o *InstanceContainer3PortsInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InstanceContainer3PortsInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InstanceContainer3PortsInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *InstanceContainer3PortsInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetExportName

`func (o *InstanceContainer3PortsInner) GetExportName() string`

GetExportName returns the ExportName field if non-nil, zero value otherwise.

### GetExportNameOk

`func (o *InstanceContainer3PortsInner) GetExportNameOk() (*string, bool)`

GetExportNameOk returns a tuple with the ExportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportName

`func (o *InstanceContainer3PortsInner) SetExportName(v string)`

SetExportName sets ExportName field to given value.

### HasExportName

`func (o *InstanceContainer3PortsInner) HasExportName() bool`

HasExportName returns a boolean if a field has been set.

### GetLoadBalanceProtocol

`func (o *InstanceContainer3PortsInner) GetLoadBalanceProtocol() string`

GetLoadBalanceProtocol returns the LoadBalanceProtocol field if non-nil, zero value otherwise.

### GetLoadBalanceProtocolOk

`func (o *InstanceContainer3PortsInner) GetLoadBalanceProtocolOk() (*string, bool)`

GetLoadBalanceProtocolOk returns a tuple with the LoadBalanceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceProtocol

`func (o *InstanceContainer3PortsInner) SetLoadBalanceProtocol(v string)`

SetLoadBalanceProtocol sets LoadBalanceProtocol field to given value.

### HasLoadBalanceProtocol

`func (o *InstanceContainer3PortsInner) HasLoadBalanceProtocol() bool`

HasLoadBalanceProtocol returns a boolean if a field has been set.

### SetLoadBalanceProtocolNil

`func (o *InstanceContainer3PortsInner) SetLoadBalanceProtocolNil(b bool)`

 SetLoadBalanceProtocolNil sets the value for LoadBalanceProtocol to be an explicit nil

### UnsetLoadBalanceProtocol
`func (o *InstanceContainer3PortsInner) UnsetLoadBalanceProtocol()`

UnsetLoadBalanceProtocol ensures that no value is present for LoadBalanceProtocol, not even an explicit nil
### GetLoadBalance

`func (o *InstanceContainer3PortsInner) GetLoadBalance() bool`

GetLoadBalance returns the LoadBalance field if non-nil, zero value otherwise.

### GetLoadBalanceOk

`func (o *InstanceContainer3PortsInner) GetLoadBalanceOk() (*bool, bool)`

GetLoadBalanceOk returns a tuple with the LoadBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalance

`func (o *InstanceContainer3PortsInner) SetLoadBalance(v bool)`

SetLoadBalance sets LoadBalance field to given value.

### HasLoadBalance

`func (o *InstanceContainer3PortsInner) HasLoadBalance() bool`

HasLoadBalance returns a boolean if a field has been set.

### GetProtocol

`func (o *InstanceContainer3PortsInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InstanceContainer3PortsInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InstanceContainer3PortsInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InstanceContainer3PortsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *InstanceContainer3PortsInner) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *InstanceContainer3PortsInner) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetLink

`func (o *InstanceContainer3PortsInner) GetLink() bool`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *InstanceContainer3PortsInner) GetLinkOk() (*bool, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *InstanceContainer3PortsInner) SetLink(v bool)`

SetLink sets Link field to given value.

### HasLink

`func (o *InstanceContainer3PortsInner) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetExternalIp

`func (o *InstanceContainer3PortsInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *InstanceContainer3PortsInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *InstanceContainer3PortsInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *InstanceContainer3PortsInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *InstanceContainer3PortsInner) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *InstanceContainer3PortsInner) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetInternalIp

`func (o *InstanceContainer3PortsInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *InstanceContainer3PortsInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *InstanceContainer3PortsInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *InstanceContainer3PortsInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *InstanceContainer3PortsInner) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *InstanceContainer3PortsInner) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


