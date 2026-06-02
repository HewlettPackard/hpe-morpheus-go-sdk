# GetUserGroup200ResponseUserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SudoUser** | Pointer to **bool** |  | [optional] 
**ServerGroup** | Pointer to **NullableString** |  | [optional] 
**Users** | Pointer to [**[]GetUserGroup200ResponseUserGroupUsersInner**](GetUserGroup200ResponseUserGroupUsersInner.md) |  | [optional] 
**Account** | Pointer to [**GetUserGroup200ResponseUserGroupAccount**](GetUserGroup200ResponseUserGroupAccount.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetUserGroup200ResponseUserGroup

`func NewGetUserGroup200ResponseUserGroup() *GetUserGroup200ResponseUserGroup`

NewGetUserGroup200ResponseUserGroup instantiates a new GetUserGroup200ResponseUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetUserGroup200ResponseUserGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUserGroup200ResponseUserGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUserGroup200ResponseUserGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetUserGroup200ResponseUserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetUserGroup200ResponseUserGroup) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetUserGroup200ResponseUserGroup) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetUserGroup200ResponseUserGroup) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetUserGroup200ResponseUserGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *GetUserGroup200ResponseUserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUserGroup200ResponseUserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUserGroup200ResponseUserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUserGroup200ResponseUserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetUserGroup200ResponseUserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetUserGroup200ResponseUserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetUserGroup200ResponseUserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetUserGroup200ResponseUserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetUserGroup200ResponseUserGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetUserGroup200ResponseUserGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *GetUserGroup200ResponseUserGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetUserGroup200ResponseUserGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetUserGroup200ResponseUserGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetUserGroup200ResponseUserGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSudoUser

`func (o *GetUserGroup200ResponseUserGroup) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *GetUserGroup200ResponseUserGroup) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *GetUserGroup200ResponseUserGroup) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *GetUserGroup200ResponseUserGroup) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.

### GetServerGroup

`func (o *GetUserGroup200ResponseUserGroup) GetServerGroup() string`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *GetUserGroup200ResponseUserGroup) GetServerGroupOk() (*string, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *GetUserGroup200ResponseUserGroup) SetServerGroup(v string)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *GetUserGroup200ResponseUserGroup) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### SetServerGroupNil

`func (o *GetUserGroup200ResponseUserGroup) SetServerGroupNil(b bool)`

 SetServerGroupNil sets the value for ServerGroup to be an explicit nil

### UnsetServerGroup
`func (o *GetUserGroup200ResponseUserGroup) UnsetServerGroup()`

UnsetServerGroup ensures that no value is present for ServerGroup, not even an explicit nil
### GetUsers

`func (o *GetUserGroup200ResponseUserGroup) GetUsers() []GetUserGroup200ResponseUserGroupUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetUserGroup200ResponseUserGroup) GetUsersOk() (*[]GetUserGroup200ResponseUserGroupUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetUserGroup200ResponseUserGroup) SetUsers(v []GetUserGroup200ResponseUserGroupUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetUserGroup200ResponseUserGroup) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAccount

`func (o *GetUserGroup200ResponseUserGroup) GetAccount() GetUserGroup200ResponseUserGroupAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetUserGroup200ResponseUserGroup) GetAccountOk() (*GetUserGroup200ResponseUserGroupAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetUserGroup200ResponseUserGroup) SetAccount(v GetUserGroup200ResponseUserGroupAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetUserGroup200ResponseUserGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetUserGroup200ResponseUserGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetUserGroup200ResponseUserGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetUserGroup200ResponseUserGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetUserGroup200ResponseUserGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetUserGroup200ResponseUserGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetUserGroup200ResponseUserGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetUserGroup200ResponseUserGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetUserGroup200ResponseUserGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


