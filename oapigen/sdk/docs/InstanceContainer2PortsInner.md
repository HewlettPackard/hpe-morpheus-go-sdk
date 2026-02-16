# InstanceContainer2PortsInner

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

### NewInstanceContainer2PortsInner

`func NewInstanceContainer2PortsInner() *InstanceContainer2PortsInner`

NewInstanceContainer2PortsInner instantiates a new InstanceContainer2PortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainer2PortsInnerWithDefaults

`func NewInstanceContainer2PortsInnerWithDefaults() *InstanceContainer2PortsInner`

NewInstanceContainer2PortsInnerWithDefaults instantiates a new InstanceContainer2PortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainer2PortsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainer2PortsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainer2PortsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainer2PortsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternal

`func (o *InstanceContainer2PortsInner) GetExternal() int64`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *InstanceContainer2PortsInner) GetExternalOk() (*int64, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *InstanceContainer2PortsInner) SetExternal(v int64)`

SetExternal sets External field to given value.

### HasExternal

`func (o *InstanceContainer2PortsInner) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetInternal

`func (o *InstanceContainer2PortsInner) GetInternal() int64`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *InstanceContainer2PortsInner) GetInternalOk() (*int64, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *InstanceContainer2PortsInner) SetInternal(v int64)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *InstanceContainer2PortsInner) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetDisplayName

`func (o *InstanceContainer2PortsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InstanceContainer2PortsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InstanceContainer2PortsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InstanceContainer2PortsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPrimaryPort

`func (o *InstanceContainer2PortsInner) GetPrimaryPort() bool`

GetPrimaryPort returns the PrimaryPort field if non-nil, zero value otherwise.

### GetPrimaryPortOk

`func (o *InstanceContainer2PortsInner) GetPrimaryPortOk() (*bool, bool)`

GetPrimaryPortOk returns a tuple with the PrimaryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPort

`func (o *InstanceContainer2PortsInner) SetPrimaryPort(v bool)`

SetPrimaryPort sets PrimaryPort field to given value.

### HasPrimaryPort

`func (o *InstanceContainer2PortsInner) HasPrimaryPort() bool`

HasPrimaryPort returns a boolean if a field has been set.

### GetExport

`func (o *InstanceContainer2PortsInner) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *InstanceContainer2PortsInner) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *InstanceContainer2PortsInner) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *InstanceContainer2PortsInner) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetVisible

`func (o *InstanceContainer2PortsInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InstanceContainer2PortsInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InstanceContainer2PortsInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *InstanceContainer2PortsInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetExportName

`func (o *InstanceContainer2PortsInner) GetExportName() string`

GetExportName returns the ExportName field if non-nil, zero value otherwise.

### GetExportNameOk

`func (o *InstanceContainer2PortsInner) GetExportNameOk() (*string, bool)`

GetExportNameOk returns a tuple with the ExportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportName

`func (o *InstanceContainer2PortsInner) SetExportName(v string)`

SetExportName sets ExportName field to given value.

### HasExportName

`func (o *InstanceContainer2PortsInner) HasExportName() bool`

HasExportName returns a boolean if a field has been set.

### GetLoadBalanceProtocol

`func (o *InstanceContainer2PortsInner) GetLoadBalanceProtocol() string`

GetLoadBalanceProtocol returns the LoadBalanceProtocol field if non-nil, zero value otherwise.

### GetLoadBalanceProtocolOk

`func (o *InstanceContainer2PortsInner) GetLoadBalanceProtocolOk() (*string, bool)`

GetLoadBalanceProtocolOk returns a tuple with the LoadBalanceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceProtocol

`func (o *InstanceContainer2PortsInner) SetLoadBalanceProtocol(v string)`

SetLoadBalanceProtocol sets LoadBalanceProtocol field to given value.

### HasLoadBalanceProtocol

`func (o *InstanceContainer2PortsInner) HasLoadBalanceProtocol() bool`

HasLoadBalanceProtocol returns a boolean if a field has been set.

### SetLoadBalanceProtocolNil

`func (o *InstanceContainer2PortsInner) SetLoadBalanceProtocolNil(b bool)`

 SetLoadBalanceProtocolNil sets the value for LoadBalanceProtocol to be an explicit nil

### UnsetLoadBalanceProtocol
`func (o *InstanceContainer2PortsInner) UnsetLoadBalanceProtocol()`

UnsetLoadBalanceProtocol ensures that no value is present for LoadBalanceProtocol, not even an explicit nil
### GetLoadBalance

`func (o *InstanceContainer2PortsInner) GetLoadBalance() bool`

GetLoadBalance returns the LoadBalance field if non-nil, zero value otherwise.

### GetLoadBalanceOk

`func (o *InstanceContainer2PortsInner) GetLoadBalanceOk() (*bool, bool)`

GetLoadBalanceOk returns a tuple with the LoadBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalance

`func (o *InstanceContainer2PortsInner) SetLoadBalance(v bool)`

SetLoadBalance sets LoadBalance field to given value.

### HasLoadBalance

`func (o *InstanceContainer2PortsInner) HasLoadBalance() bool`

HasLoadBalance returns a boolean if a field has been set.

### GetProtocol

`func (o *InstanceContainer2PortsInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InstanceContainer2PortsInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InstanceContainer2PortsInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InstanceContainer2PortsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *InstanceContainer2PortsInner) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *InstanceContainer2PortsInner) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetLink

`func (o *InstanceContainer2PortsInner) GetLink() bool`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *InstanceContainer2PortsInner) GetLinkOk() (*bool, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *InstanceContainer2PortsInner) SetLink(v bool)`

SetLink sets Link field to given value.

### HasLink

`func (o *InstanceContainer2PortsInner) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetExternalIp

`func (o *InstanceContainer2PortsInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *InstanceContainer2PortsInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *InstanceContainer2PortsInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *InstanceContainer2PortsInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *InstanceContainer2PortsInner) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *InstanceContainer2PortsInner) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetInternalIp

`func (o *InstanceContainer2PortsInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *InstanceContainer2PortsInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *InstanceContainer2PortsInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *InstanceContainer2PortsInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *InstanceContainer2PortsInner) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *InstanceContainer2PortsInner) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


