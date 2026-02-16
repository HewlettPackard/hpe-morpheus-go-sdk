# VMWareInstanceConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoAgent** | Pointer to **NullableBool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**ResourcePoolId** | Pointer to **string** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional] 
**HostId** | Pointer to **string** | Specific host to deploy to if so desired. | [optional] 
**SmbiosAssetTag** | Pointer to **string** | Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used. | [optional] 
**NestedVirtualization** | Pointer to **string** | Enable Nested Virtualization | [optional] [default to "off"]
**VmwareFolderId** | Pointer to **string** | VMWare Folder External ID (as a String) or ID (as an Integer or String) | [optional] 
**CreateUser** | Pointer to **NullableBool** | Create user | [optional] [default to false]
**Template** | Pointer to **int64** | Image ID. This is the ID of a Virtual Image. | [optional] 

## Methods

### NewVMWareInstanceConfiguration1

`func NewVMWareInstanceConfiguration1() *VMWareInstanceConfiguration1`

NewVMWareInstanceConfiguration1 instantiates a new VMWareInstanceConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMWareInstanceConfiguration1WithDefaults

`func NewVMWareInstanceConfiguration1WithDefaults() *VMWareInstanceConfiguration1`

NewVMWareInstanceConfiguration1WithDefaults instantiates a new VMWareInstanceConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoAgent

`func (o *VMWareInstanceConfiguration1) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *VMWareInstanceConfiguration1) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *VMWareInstanceConfiguration1) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *VMWareInstanceConfiguration1) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### SetNoAgentNil

`func (o *VMWareInstanceConfiguration1) SetNoAgentNil(b bool)`

 SetNoAgentNil sets the value for NoAgent to be an explicit nil

### UnsetNoAgent
`func (o *VMWareInstanceConfiguration1) UnsetNoAgent()`

UnsetNoAgent ensures that no value is present for NoAgent, not even an explicit nil
### GetResourcePoolId

`func (o *VMWareInstanceConfiguration1) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *VMWareInstanceConfiguration1) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *VMWareInstanceConfiguration1) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *VMWareInstanceConfiguration1) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetHostId

`func (o *VMWareInstanceConfiguration1) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *VMWareInstanceConfiguration1) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *VMWareInstanceConfiguration1) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *VMWareInstanceConfiguration1) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *VMWareInstanceConfiguration1) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *VMWareInstanceConfiguration1) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *VMWareInstanceConfiguration1) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *VMWareInstanceConfiguration1) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *VMWareInstanceConfiguration1) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *VMWareInstanceConfiguration1) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *VMWareInstanceConfiguration1) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *VMWareInstanceConfiguration1) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *VMWareInstanceConfiguration1) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *VMWareInstanceConfiguration1) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *VMWareInstanceConfiguration1) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *VMWareInstanceConfiguration1) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetCreateUser

`func (o *VMWareInstanceConfiguration1) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *VMWareInstanceConfiguration1) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *VMWareInstanceConfiguration1) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *VMWareInstanceConfiguration1) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### SetCreateUserNil

`func (o *VMWareInstanceConfiguration1) SetCreateUserNil(b bool)`

 SetCreateUserNil sets the value for CreateUser to be an explicit nil

### UnsetCreateUser
`func (o *VMWareInstanceConfiguration1) UnsetCreateUser()`

UnsetCreateUser ensures that no value is present for CreateUser, not even an explicit nil
### GetTemplate

`func (o *VMWareInstanceConfiguration1) GetTemplate() int64`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *VMWareInstanceConfiguration1) GetTemplateOk() (*int64, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *VMWareInstanceConfiguration1) SetTemplate(v int64)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *VMWareInstanceConfiguration1) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


