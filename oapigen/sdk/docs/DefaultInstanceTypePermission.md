# DefaultInstanceTypePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionCode** | **string** | &#x60;InstanceType&#x60; is the code for Default Instance Type Access | 
**Access** | **string** | The new access level. | 

## Methods

### NewDefaultInstanceTypePermission

`func NewDefaultInstanceTypePermission(permissionCode string, access string, ) *DefaultInstanceTypePermission`

NewDefaultInstanceTypePermission instantiates a new DefaultInstanceTypePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultInstanceTypePermissionWithDefaults

`func NewDefaultInstanceTypePermissionWithDefaults() *DefaultInstanceTypePermission`

NewDefaultInstanceTypePermissionWithDefaults instantiates a new DefaultInstanceTypePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionCode

`func (o *DefaultInstanceTypePermission) GetPermissionCode() string`

GetPermissionCode returns the PermissionCode field if non-nil, zero value otherwise.

### GetPermissionCodeOk

`func (o *DefaultInstanceTypePermission) GetPermissionCodeOk() (*string, bool)`

GetPermissionCodeOk returns a tuple with the PermissionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionCode

`func (o *DefaultInstanceTypePermission) SetPermissionCode(v string)`

SetPermissionCode sets PermissionCode field to given value.


### GetAccess

`func (o *DefaultInstanceTypePermission) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DefaultInstanceTypePermission) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DefaultInstanceTypePermission) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


