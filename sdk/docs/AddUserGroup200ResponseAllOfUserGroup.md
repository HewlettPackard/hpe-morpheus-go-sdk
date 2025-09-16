# AddUserGroup200ResponseAllOfUserGroup

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
**Users** | Pointer to [**[]ListCredentials200ResponseAllOfCredentialsInnerUser**](ListCredentials200ResponseAllOfCredentialsInnerUser.md) |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddUserGroup200ResponseAllOfUserGroup

`func NewAddUserGroup200ResponseAllOfUserGroup() *AddUserGroup200ResponseAllOfUserGroup`

NewAddUserGroup200ResponseAllOfUserGroup instantiates a new AddUserGroup200ResponseAllOfUserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserGroup200ResponseAllOfUserGroupWithDefaults

`func NewAddUserGroup200ResponseAllOfUserGroupWithDefaults() *AddUserGroup200ResponseAllOfUserGroup`

NewAddUserGroup200ResponseAllOfUserGroupWithDefaults instantiates a new AddUserGroup200ResponseAllOfUserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddUserGroup200ResponseAllOfUserGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSudoUser

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.

### GetServerGroup

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetServerGroup() string`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetServerGroupOk() (*string, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetServerGroup(v string)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### SetServerGroupNil

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetServerGroupNil(b bool)`

 SetServerGroupNil sets the value for ServerGroup to be an explicit nil

### UnsetServerGroup
`func (o *AddUserGroup200ResponseAllOfUserGroup) UnsetServerGroup()`

UnsetServerGroup ensures that no value is present for ServerGroup, not even an explicit nil
### GetUsers

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetUsers() []ListCredentials200ResponseAllOfCredentialsInnerUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetUsersOk() (*[]ListCredentials200ResponseAllOfCredentialsInnerUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetUsers(v []ListCredentials200ResponseAllOfCredentialsInnerUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAccount

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddUserGroup200ResponseAllOfUserGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddUserGroup200ResponseAllOfUserGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddUserGroup200ResponseAllOfUserGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


