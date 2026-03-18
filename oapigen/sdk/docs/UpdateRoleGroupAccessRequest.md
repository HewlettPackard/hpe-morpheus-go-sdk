# UpdateRoleGroupAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** | &#x60;id&#x60; of the group (site) | 
**Access** | **string** | The new access level. | 
**AllGroups** | **bool** | Apply to all groups (site) | 

## Methods

### NewUpdateRoleGroupAccessRequest

`func NewUpdateRoleGroupAccessRequest(groupId int32, access string, allGroups bool, ) *UpdateRoleGroupAccessRequest`

NewUpdateRoleGroupAccessRequest instantiates a new UpdateRoleGroupAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleGroupAccessRequestWithDefaults

`func NewUpdateRoleGroupAccessRequestWithDefaults() *UpdateRoleGroupAccessRequest`

NewUpdateRoleGroupAccessRequestWithDefaults instantiates a new UpdateRoleGroupAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UpdateRoleGroupAccessRequest) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateRoleGroupAccessRequest) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateRoleGroupAccessRequest) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetAccess

`func (o *UpdateRoleGroupAccessRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleGroupAccessRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleGroupAccessRequest) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAllGroups

`func (o *UpdateRoleGroupAccessRequest) GetAllGroups() bool`

GetAllGroups returns the AllGroups field if non-nil, zero value otherwise.

### GetAllGroupsOk

`func (o *UpdateRoleGroupAccessRequest) GetAllGroupsOk() (*bool, bool)`

GetAllGroupsOk returns a tuple with the AllGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllGroups

`func (o *UpdateRoleGroupAccessRequest) SetAllGroups(v bool)`

SetAllGroups sets AllGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


