# UpdateUserGroup200ResponseAllOfUserGroup

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
**Users** | Pointer to [**[]UpdateUserGroup200ResponseAllOfUserGroupUsersInner**](UpdateUserGroup200ResponseAllOfUserGroupUsersInner.md) |  | [optional] 
**Account** | Pointer to [**UpdateUserGroup200ResponseAllOfUserGroupAccount**](UpdateUserGroup200ResponseAllOfUserGroupAccount.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateUserGroup200ResponseAllOfUserGroup

`func NewUpdateUserGroup200ResponseAllOfUserGroup() *UpdateUserGroup200ResponseAllOfUserGroup`

NewUpdateUserGroup200ResponseAllOfUserGroup instantiates a new UpdateUserGroup200ResponseAllOfUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateUserGroup200ResponseAllOfUserGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSudoUser

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.

### GetServerGroup

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetServerGroup() string`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetServerGroupOk() (*string, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetServerGroup(v string)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### SetServerGroupNil

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetServerGroupNil(b bool)`

 SetServerGroupNil sets the value for ServerGroup to be an explicit nil

### UnsetServerGroup
`func (o *UpdateUserGroup200ResponseAllOfUserGroup) UnsetServerGroup()`

UnsetServerGroup ensures that no value is present for ServerGroup, not even an explicit nil
### GetUsers

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetUsers() []UpdateUserGroup200ResponseAllOfUserGroupUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetUsersOk() (*[]UpdateUserGroup200ResponseAllOfUserGroupUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetUsers(v []UpdateUserGroup200ResponseAllOfUserGroupUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetAccount() UpdateUserGroup200ResponseAllOfUserGroupAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetAccountOk() (*UpdateUserGroup200ResponseAllOfUserGroupAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetAccount(v UpdateUserGroup200ResponseAllOfUserGroupAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateUserGroup200ResponseAllOfUserGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


