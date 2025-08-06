# AddRolesRequestRoleTaskSetPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | &#x60;id&#x60; of the workflow (taskSet) | 
**Access** | **string** | The new access level. | 

## Methods

### NewAddRolesRequestRoleTaskSetPermissionsInner

`func NewAddRolesRequestRoleTaskSetPermissionsInner(id int64, access string, ) *AddRolesRequestRoleTaskSetPermissionsInner`

NewAddRolesRequestRoleTaskSetPermissionsInner instantiates a new AddRolesRequestRoleTaskSetPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRolesRequestRoleTaskSetPermissionsInnerWithDefaults

`func NewAddRolesRequestRoleTaskSetPermissionsInnerWithDefaults() *AddRolesRequestRoleTaskSetPermissionsInner`

NewAddRolesRequestRoleTaskSetPermissionsInnerWithDefaults instantiates a new AddRolesRequestRoleTaskSetPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddRolesRequestRoleTaskSetPermissionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddRolesRequestRoleTaskSetPermissionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddRolesRequestRoleTaskSetPermissionsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetAccess

`func (o *AddRolesRequestRoleTaskSetPermissionsInner) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AddRolesRequestRoleTaskSetPermissionsInner) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AddRolesRequestRoleTaskSetPermissionsInner) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


