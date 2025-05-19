# ListUserGroups200ResponseAllOfUserGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SudoUser** | Pointer to **bool** |  | [optional] 
**ServerGroup** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]ListCredentials200ResponseAllOfCredentialsInnerUser**](ListCredentials200ResponseAllOfCredentialsInnerUser.md) |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListUserGroups200ResponseAllOfUserGroupsInner

`func NewListUserGroups200ResponseAllOfUserGroupsInner() *ListUserGroups200ResponseAllOfUserGroupsInner`

NewListUserGroups200ResponseAllOfUserGroupsInner instantiates a new ListUserGroups200ResponseAllOfUserGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserGroups200ResponseAllOfUserGroupsInnerWithDefaults

`func NewListUserGroups200ResponseAllOfUserGroupsInnerWithDefaults() *ListUserGroups200ResponseAllOfUserGroupsInner`

NewListUserGroups200ResponseAllOfUserGroupsInnerWithDefaults instantiates a new ListUserGroups200ResponseAllOfUserGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSudoUser

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetSudoUser() bool`

GetSudoUser returns the SudoUser field if non-nil, zero value otherwise.

### GetSudoUserOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetSudoUserOk() (*bool, bool)`

GetSudoUserOk returns a tuple with the SudoUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoUser

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetSudoUser(v bool)`

SetSudoUser sets SudoUser field to given value.

### HasSudoUser

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasSudoUser() bool`

HasSudoUser returns a boolean if a field has been set.

### GetServerGroup

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetServerGroup() string`

GetServerGroup returns the ServerGroup field if non-nil, zero value otherwise.

### GetServerGroupOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetServerGroupOk() (*string, bool)`

GetServerGroupOk returns a tuple with the ServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroup

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetServerGroup(v string)`

SetServerGroup sets ServerGroup field to given value.

### HasServerGroup

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasServerGroup() bool`

HasServerGroup returns a boolean if a field has been set.

### GetUsers

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetUsers() []ListCredentials200ResponseAllOfCredentialsInnerUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetUsersOk() (*[]ListCredentials200ResponseAllOfCredentialsInnerUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetUsers(v []ListCredentials200ResponseAllOfCredentialsInnerUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAccount

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListUserGroups200ResponseAllOfUserGroupsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


