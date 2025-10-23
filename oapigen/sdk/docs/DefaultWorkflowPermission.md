# DefaultWorkflowPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionCode** | **string** | &#x60;TaskSet&#x60; is the code for Default Workflow Access | 
**Access** | **string** | The new access level. | 

## Methods

### NewDefaultWorkflowPermission

`func NewDefaultWorkflowPermission(permissionCode string, access string, ) *DefaultWorkflowPermission`

NewDefaultWorkflowPermission instantiates a new DefaultWorkflowPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultWorkflowPermissionWithDefaults

`func NewDefaultWorkflowPermissionWithDefaults() *DefaultWorkflowPermission`

NewDefaultWorkflowPermissionWithDefaults instantiates a new DefaultWorkflowPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionCode

`func (o *DefaultWorkflowPermission) GetPermissionCode() string`

GetPermissionCode returns the PermissionCode field if non-nil, zero value otherwise.

### GetPermissionCodeOk

`func (o *DefaultWorkflowPermission) GetPermissionCodeOk() (*string, bool)`

GetPermissionCodeOk returns a tuple with the PermissionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionCode

`func (o *DefaultWorkflowPermission) SetPermissionCode(v string)`

SetPermissionCode sets PermissionCode field to given value.


### GetAccess

`func (o *DefaultWorkflowPermission) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *DefaultWorkflowPermission) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *DefaultWorkflowPermission) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


