# UserGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SudoUser** | Pointer to **bool** |  | [optional] 
**ServerGroup** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **[]int64** | A list of IDs of users that are in the user group | [optional] 

## Methods

### NewUserGroupCreate

`func NewUserGroupCreate() *UserGroupCreate`

NewUserGroupCreate instantiates a new UserGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGroupCreateWithDefaults

`func NewUserGroupCreateWithDefaults() *UserGroupCreate`

NewUserGroupCreateWithDefaults instantiates a new UserGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserGroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserGroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserGroupCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserGroupCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UserGroupCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserGroupCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserGroupCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserGroupCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSudoUser

`func (o *UserGroupCreate) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *UserGroupCreate) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *UserGroupCreate) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *UserGroupCreate) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.

### GetServerGroup

`func (o *UserGroupCreate) GetServerGroup() string`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *UserGroupCreate) GetServerGroupOk() (*string, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *UserGroupCreate) SetServerGroup(v string)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *UserGroupCreate) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### GetUsers

`func (o *UserGroupCreate) GetUsers() []int64`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UserGroupCreate) GetUsersOk() (*[]int64, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UserGroupCreate) SetUsers(v []int64)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UserGroupCreate) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


