# DefaultPersonaPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionCode** | **string** | &#x60;Persona&#x60; is the code for Default Persona Access | 
**Access** | **string** | The new access level. | 

## Methods

### NewDefaultPersonaPermission

`func NewDefaultPersonaPermission(permissionCode string, access string, ) *DefaultPersonaPermission`

NewDefaultPersonaPermission instantiates a new DefaultPersonaPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultPersonaPermissionWithDefaults

`func NewDefaultPersonaPermissionWithDefaults() *DefaultPersonaPermission`

NewDefaultPersonaPermissionWithDefaults instantiates a new DefaultPersonaPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionCode

`func (o *DefaultPersonaPermission) GetPermissionCode() string`

GetPermissionCode returns the PermissionCode field if non-nil, zero value otherwise.

### GetPermissionCodeOk

`func (o *DefaultPersonaPermission) GetPermissionCodeOk() (*string, bool)`

GetPermissionCodeOk returns a tuple with the PermissionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionCode

`func (o *DefaultPersonaPermission) SetPermissionCode(v string)`

SetPermissionCode sets PermissionCode field to given value.


### GetAccess

`func (o *DefaultPersonaPermission) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DefaultPersonaPermission) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DefaultPersonaPermission) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


