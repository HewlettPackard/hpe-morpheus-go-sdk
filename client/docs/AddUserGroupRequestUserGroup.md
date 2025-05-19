# AddUserGroupRequestUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SudoUser** | Pointer to **bool** |  | [optional] 
**ServerGroup** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **[]int64** | A list of IDs of users that are in the user group | [optional] 

## Methods

### NewAddUserGroupRequestUserGroup

`func NewAddUserGroupRequestUserGroup() *AddUserGroupRequestUserGroup`

NewAddUserGroupRequestUserGroup instantiates a new AddUserGroupRequestUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserGroupRequestUserGroupWithDefaults

`func NewAddUserGroupRequestUserGroupWithDefaults() *AddUserGroupRequestUserGroup`

NewAddUserGroupRequestUserGroupWithDefaults instantiates a new AddUserGroupRequestUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddUserGroupRequestUserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddUserGroupRequestUserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddUserGroupRequestUserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddUserGroupRequestUserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddUserGroupRequestUserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUserGroupRequestUserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUserGroupRequestUserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUserGroupRequestUserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSudoUser

`func (o *AddUserGroupRequestUserGroup) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *AddUserGroupRequestUserGroup) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *AddUserGroupRequestUserGroup) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *AddUserGroupRequestUserGroup) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.

### GetServerGroup

`func (o *AddUserGroupRequestUserGroup) GetServerGroup() string`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *AddUserGroupRequestUserGroup) GetServerGroupOk() (*string, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *AddUserGroupRequestUserGroup) SetServerGroup(v string)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *AddUserGroupRequestUserGroup) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### GetUsers

`func (o *AddUserGroupRequestUserGroup) GetUsers() []int64`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AddUserGroupRequestUserGroup) GetUsersOk() (*[]int64, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AddUserGroupRequestUserGroup) SetUsers(v []int64)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AddUserGroupRequestUserGroup) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


