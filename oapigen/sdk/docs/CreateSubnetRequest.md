# CreateSubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to [**CreateSubnetRequestSubnet**](CreateSubnetRequestSubnet.md) |  | [optional] 
**ResourcePermission** | Pointer to [**CreateSubnetRequestResourcePermission**](CreateSubnetRequestResourcePermission.md) |  | [optional] 

## Methods

### NewCreateSubnetRequest

`func NewCreateSubnetRequest() *CreateSubnetRequest`

NewCreateSubnetRequest instantiates a new CreateSubnetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubnetRequestWithDefaults

`func NewCreateSubnetRequestWithDefaults() *CreateSubnetRequest`

NewCreateSubnetRequestWithDefaults instantiates a new CreateSubnetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *CreateSubnetRequest) GetSubnet() CreateSubnetRequestSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateSubnetRequest) GetSubnetOk() (*CreateSubnetRequestSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateSubnetRequest) SetSubnet(v CreateSubnetRequestSubnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *CreateSubnetRequest) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetResourcePermission

`func (o *CreateSubnetRequest) GetResourcePermission() CreateSubnetRequestResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *CreateSubnetRequest) GetResourcePermissionOk() (*CreateSubnetRequestResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *CreateSubnetRequest) SetResourcePermission(v CreateSubnetRequestResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *CreateSubnetRequest) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


