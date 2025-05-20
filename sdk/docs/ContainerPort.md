# ContainerPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Index** | Pointer to **int64** |  | [optional] 
**External** | Pointer to **int64** |  | [optional] 
**Internal** | Pointer to **int64** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**PrimaryPort** | Pointer to **bool** |  | [optional] 
**Export** | Pointer to **bool** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 
**ExportName** | Pointer to **string** |  | [optional] 
**LoadBalanceProtocol** | Pointer to **string** |  | [optional] 
**LoadBalance** | Pointer to **bool** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **bool** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 

## Methods

### NewContainerPort

`func NewContainerPort() *ContainerPort`

NewContainerPort instantiates a new ContainerPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerPortWithDefaults

`func NewContainerPortWithDefaults() *ContainerPort`

NewContainerPortWithDefaults instantiates a new ContainerPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerPort) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerPort) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerPort) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ContainerPort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *ContainerPort) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ContainerPort) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ContainerPort) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ContainerPort) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetExternal

`func (o *ContainerPort) GetExternal() int64`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *ContainerPort) GetExternalOk() (*int64, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *ContainerPort) SetExternal(v int64)`

SetExternal sets External field to given value.

### HasExternal

`func (o *ContainerPort) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetInternal

`func (o *ContainerPort) GetInternal() int64`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ContainerPort) GetInternalOk() (*int64, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ContainerPort) SetInternal(v int64)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ContainerPort) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetDisplayName

`func (o *ContainerPort) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ContainerPort) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ContainerPort) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ContainerPort) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPrimaryPort

`func (o *ContainerPort) GetPrimaryPort() bool`

GetPrimaryPort returns the PrimaryPort field if non-nil, zero value otherwise.

### GetPrimaryPortOk

`func (o *ContainerPort) GetPrimaryPortOk() (*bool, bool)`

GetPrimaryPortOk returns a tuple with the PrimaryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPort

`func (o *ContainerPort) SetPrimaryPort(v bool)`

SetPrimaryPort sets PrimaryPort field to given value.

### HasPrimaryPort

`func (o *ContainerPort) HasPrimaryPort() bool`

HasPrimaryPort returns a boolean if a field has been set.

### GetExport

`func (o *ContainerPort) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *ContainerPort) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *ContainerPort) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *ContainerPort) HasExport() bool`

HasExport returns a boolean if a field has been set.

### GetVisible

`func (o *ContainerPort) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ContainerPort) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ContainerPort) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ContainerPort) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetExportName

`func (o *ContainerPort) GetExportName() string`

GetExportName returns the ExportName field if non-nil, zero value otherwise.

### GetExportNameOk

`func (o *ContainerPort) GetExportNameOk() (*string, bool)`

GetExportNameOk returns a tuple with the ExportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportName

`func (o *ContainerPort) SetExportName(v string)`

SetExportName sets ExportName field to given value.

### HasExportName

`func (o *ContainerPort) HasExportName() bool`

HasExportName returns a boolean if a field has been set.

### GetLoadBalanceProtocol

`func (o *ContainerPort) GetLoadBalanceProtocol() string`

GetLoadBalanceProtocol returns the LoadBalanceProtocol field if non-nil, zero value otherwise.

### GetLoadBalanceProtocolOk

`func (o *ContainerPort) GetLoadBalanceProtocolOk() (*string, bool)`

GetLoadBalanceProtocolOk returns a tuple with the LoadBalanceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceProtocol

`func (o *ContainerPort) SetLoadBalanceProtocol(v string)`

SetLoadBalanceProtocol sets LoadBalanceProtocol field to given value.

### HasLoadBalanceProtocol

`func (o *ContainerPort) HasLoadBalanceProtocol() bool`

HasLoadBalanceProtocol returns a boolean if a field has been set.

### GetLoadBalance

`func (o *ContainerPort) GetLoadBalance() bool`

GetLoadBalance returns the LoadBalance field if non-nil, zero value otherwise.

### GetLoadBalanceOk

`func (o *ContainerPort) GetLoadBalanceOk() (*bool, bool)`

GetLoadBalanceOk returns a tuple with the LoadBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalance

`func (o *ContainerPort) SetLoadBalance(v bool)`

SetLoadBalance sets LoadBalance field to given value.

### HasLoadBalance

`func (o *ContainerPort) HasLoadBalance() bool`

HasLoadBalance returns a boolean if a field has been set.

### GetProtocol

`func (o *ContainerPort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ContainerPort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ContainerPort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ContainerPort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetLink

`func (o *ContainerPort) GetLink() bool`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ContainerPort) GetLinkOk() (*bool, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ContainerPort) SetLink(v bool)`

SetLink sets Link field to given value.

### HasLink

`func (o *ContainerPort) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetExternalIp

`func (o *ContainerPort) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ContainerPort) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ContainerPort) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ContainerPort) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *ContainerPort) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ContainerPort) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ContainerPort) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ContainerPort) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


