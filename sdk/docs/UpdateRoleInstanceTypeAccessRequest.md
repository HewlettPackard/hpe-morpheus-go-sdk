# UpdateRoleInstanceTypeAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTypeId** | **int32** | &#x60;id&#x60; of the instance type | 
**Access** | **string** | The new access level. | 
**AllInstanceTypes** | Pointer to **bool** | Apply to all instance types | [optional] 

## Methods

### NewUpdateRoleInstanceTypeAccessRequest

`func NewUpdateRoleInstanceTypeAccessRequest(instanceTypeId int32, access string, ) *UpdateRoleInstanceTypeAccessRequest`

NewUpdateRoleInstanceTypeAccessRequest instantiates a new UpdateRoleInstanceTypeAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleInstanceTypeAccessRequestWithDefaults

`func NewUpdateRoleInstanceTypeAccessRequestWithDefaults() *UpdateRoleInstanceTypeAccessRequest`

NewUpdateRoleInstanceTypeAccessRequestWithDefaults instantiates a new UpdateRoleInstanceTypeAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTypeId

`func (o *UpdateRoleInstanceTypeAccessRequest) GetInstanceTypeId() int32`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *UpdateRoleInstanceTypeAccessRequest) GetInstanceTypeIdOk() (*int32, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *UpdateRoleInstanceTypeAccessRequest) SetInstanceTypeId(v int32)`

SetInstanceTypeId sets InstanceTypeId field to given value.


### GetAccess

`func (o *UpdateRoleInstanceTypeAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleInstanceTypeAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleInstanceTypeAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllInstanceTypes

`func (o *UpdateRoleInstanceTypeAccessRequest) GetAllInstanceTypes() bool`

GetAllInstanceTypes returns the AllInstanceTypes field if non-nil, zero value otherwise.

### GetAllInstanceTypesOk

`func (o *UpdateRoleInstanceTypeAccessRequest) GetAllInstanceTypesOk() (*bool, bool)`

GetAllInstanceTypesOk returns a tuple with the AllInstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllInstanceTypes

`func (o *UpdateRoleInstanceTypeAccessRequest) SetAllInstanceTypes(v bool)`

SetAllInstanceTypes sets AllInstanceTypes field to given value.

### HasAllInstanceTypes

`func (o *UpdateRoleInstanceTypeAccessRequest) HasAllInstanceTypes() bool`

HasAllInstanceTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


