# DefaultVDIPoolPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionCode** | **string** | &#x60;VdiPools&#x60; is the code for Default VDI Pool Access | 
**Access** | **string** | The new access level. | 

## Methods

### NewDefaultVDIPoolPermission

`func NewDefaultVDIPoolPermission(permissionCode string, access string, ) *DefaultVDIPoolPermission`

NewDefaultVDIPoolPermission instantiates a new DefaultVDIPoolPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultVDIPoolPermissionWithDefaults

`func NewDefaultVDIPoolPermissionWithDefaults() *DefaultVDIPoolPermission`

NewDefaultVDIPoolPermissionWithDefaults instantiates a new DefaultVDIPoolPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionCode

`func (o *DefaultVDIPoolPermission) GetPermissionCode() string`

GetPermissionCode returns the PermissionCode field if non-nil, zero value otherwise.

### GetPermissionCodeOk

`func (o *DefaultVDIPoolPermission) GetPermissionCodeOk() (*string, bool)`

GetPermissionCodeOk returns a tuple with the PermissionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionCode

`func (o *DefaultVDIPoolPermission) SetPermissionCode(v string)`

SetPermissionCode sets PermissionCode field to given value.


### GetAccess

`func (o *DefaultVDIPoolPermission) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DefaultVDIPoolPermission) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DefaultVDIPoolPermission) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


