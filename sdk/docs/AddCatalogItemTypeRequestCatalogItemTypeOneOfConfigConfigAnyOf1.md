# AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoAgent** | Pointer to **bool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**ResourcePoolId** | Pointer to **string** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional] 
**HostId** | Pointer to **string** | Specific host to deploy to if so desired. | [optional] 
**SmbiosAssetTag** | Pointer to **string** | Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used. | [optional] 
**NestedVirtualization** | Pointer to **string** | Enable Nested Virtualization | [optional] [default to "off"]
**VmwareFolderId** | Pointer to **string** | VMWare Folder External ID (as a String) or ID (as an Integer or String) | [optional] 

## Methods

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1WithDefaults

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1WithDefaults() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1WithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoAgent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetHostId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf1) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


