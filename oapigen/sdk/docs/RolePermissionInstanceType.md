# RolePermissionInstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTypeId** | **int32** | &#x60;id&#x60; of the instance type | 
**Access** | **string** | The new access level. | 

## Methods

### NewRolePermissionInstanceType

`func NewRolePermissionInstanceType(instanceTypeId int32, access string, ) *RolePermissionInstanceType`

NewRolePermissionInstanceType instantiates a new RolePermissionInstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionInstanceTypeWithDefaults

`func NewRolePermissionInstanceTypeWithDefaults() *RolePermissionInstanceType`

NewRolePermissionInstanceTypeWithDefaults instantiates a new RolePermissionInstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTypeId

`func (o *RolePermissionInstanceType) GetInstanceTypeId() int32`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *RolePermissionInstanceType) GetInstanceTypeIdOk() (*int32, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *RolePermissionInstanceType) SetInstanceTypeId(v int32)`

SetInstanceTypeId sets InstanceTypeId field to given value.


### GetAccess

`func (o *RolePermissionInstanceType) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RolePermissionInstanceType) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RolePermissionInstanceType) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


