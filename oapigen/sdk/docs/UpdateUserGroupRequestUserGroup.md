# UpdateUserGroupRequestUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SudoUser** | Pointer to **bool** |  | [optional] 
**ServerGroup** | Pointer to **NullableString** |  | [optional] 
**Users** | Pointer to **[]int64** | A list of IDs of users that are in the user group | [optional] 

## Methods

### NewUpdateUserGroupRequestUserGroup

`func NewUpdateUserGroupRequestUserGroup() *UpdateUserGroupRequestUserGroup`

NewUpdateUserGroupRequestUserGroup instantiates a new UpdateUserGroupRequestUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateUserGroupRequestUserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserGroupRequestUserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserGroupRequestUserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserGroupRequestUserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateUserGroupRequestUserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateUserGroupRequestUserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateUserGroupRequestUserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateUserGroupRequestUserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateUserGroupRequestUserGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateUserGroupRequestUserGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSudoUser

`func (o *UpdateUserGroupRequestUserGroup) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *UpdateUserGroupRequestUserGroup) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *UpdateUserGroupRequestUserGroup) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *UpdateUserGroupRequestUserGroup) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.

### GetServerGroup

`func (o *UpdateUserGroupRequestUserGroup) GetServerGroup() string`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *UpdateUserGroupRequestUserGroup) GetServerGroupOk() (*string, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *UpdateUserGroupRequestUserGroup) SetServerGroup(v string)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *UpdateUserGroupRequestUserGroup) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### SetServerGroupNil

`func (o *UpdateUserGroupRequestUserGroup) SetServerGroupNil(b bool)`

 SetServerGroupNil sets the value for ServerGroup to be an explicit nil

### UnsetServerGroup
`func (o *UpdateUserGroupRequestUserGroup) UnsetServerGroup()`

UnsetServerGroup ensures that no value is present for ServerGroup, not even an explicit nil
### GetUsers

`func (o *UpdateUserGroupRequestUserGroup) GetUsers() []int64`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateUserGroupRequestUserGroup) GetUsersOk() (*[]int64, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateUserGroupRequestUserGroup) SetUsers(v []int64)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdateUserGroupRequestUserGroup) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


