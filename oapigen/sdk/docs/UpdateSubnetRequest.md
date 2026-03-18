# UpdateSubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to [**UpdateSubnetRequestSubnet**](UpdateSubnetRequestSubnet.md) |  | [optional] 
**ResourcePermission** | Pointer to [**UpdateSubnetRequestResourcePermission**](UpdateSubnetRequestResourcePermission.md) |  | [optional] 

## Methods

### NewUpdateSubnetRequest

`func NewUpdateSubnetRequest() *UpdateSubnetRequest`

NewUpdateSubnetRequest instantiates a new UpdateSubnetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubnetRequestWithDefaults

`func NewUpdateSubnetRequestWithDefaults() *UpdateSubnetRequest`

NewUpdateSubnetRequestWithDefaults instantiates a new UpdateSubnetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *UpdateSubnetRequest) GetSubnet() UpdateSubnetRequestSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdateSubnetRequest) GetSubnetOk() (*UpdateSubnetRequestSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdateSubnetRequest) SetSubnet(v UpdateSubnetRequestSubnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *UpdateSubnetRequest) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetResourcePermission

`func (o *UpdateSubnetRequest) GetResourcePermission() UpdateSubnetRequestResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *UpdateSubnetRequest) GetResourcePermissionOk() (*UpdateSubnetRequestResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *UpdateSubnetRequest) SetResourcePermission(v UpdateSubnetRequestResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *UpdateSubnetRequest) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


