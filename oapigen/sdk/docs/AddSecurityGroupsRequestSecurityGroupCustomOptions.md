# AddSecurityGroupsRequestSecurityGroupCustomOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroup** | Pointer to **string** | External ID of Azure Resource Group | [optional] 
**Vpc** | Pointer to **string** | External ID of Amazon VPC | [optional] 
**ResourcePoolId** | Pointer to **int64** | Resource Pool ID (applicable to cloud types Openstack/OpenTelekom/Huawei) | [optional] 

## Methods

### NewAddSecurityGroupsRequestSecurityGroupCustomOptions

`func NewAddSecurityGroupsRequestSecurityGroupCustomOptions() *AddSecurityGroupsRequestSecurityGroupCustomOptions`

NewAddSecurityGroupsRequestSecurityGroupCustomOptions instantiates a new AddSecurityGroupsRequestSecurityGroupCustomOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroupsRequestSecurityGroupCustomOptionsWithDefaults

`func NewAddSecurityGroupsRequestSecurityGroupCustomOptionsWithDefaults() *AddSecurityGroupsRequestSecurityGroupCustomOptions`

NewAddSecurityGroupsRequestSecurityGroupCustomOptionsWithDefaults instantiates a new AddSecurityGroupsRequestSecurityGroupCustomOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceGroup

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) GetResourceGroup() string`

GetResourceGroup returns the ResourceGroup field if non-nil, zero value otherwise.

### GetResourceGroupOk

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) GetResourceGroupOk() (*string, bool)`

GetResourceGroupOk returns a tuple with the ResourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroup

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) SetResourceGroup(v string)`

SetResourceGroup sets ResourceGroup field to given value.

### HasResourceGroup

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) HasResourceGroup() bool`

HasResourceGroup returns a boolean if a field has been set.

### GetVpc

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) HasVpc() bool`

HasVpc returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AddSecurityGroupsRequestSecurityGroupCustomOptions) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


