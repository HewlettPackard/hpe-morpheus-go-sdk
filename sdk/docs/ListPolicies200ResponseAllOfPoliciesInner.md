# ListPolicies200ResponseAllOfPoliciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Site** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**User** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**Role** | Pointer to [**ListPolicies200ResponseAllOfPoliciesInnerRole**](ListPolicies200ResponseAllOfPoliciesInnerRole.md) |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**EachUser** | Pointer to **NullableBool** |  | [optional] 
**Config** | Pointer to [**ListPolicies200ResponseAllOfPoliciesInnerConfig**](ListPolicies200ResponseAllOfPoliciesInnerConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**NullableListApprovals200ResponseAllOfApprovalsInnerAccount**](ListApprovals200ResponseAllOfApprovalsInnerAccount.md) |  | [optional] 
**Accounts** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewListPolicies200ResponseAllOfPoliciesInner

`func NewListPolicies200ResponseAllOfPoliciesInner() *ListPolicies200ResponseAllOfPoliciesInner`

NewListPolicies200ResponseAllOfPoliciesInner instantiates a new ListPolicies200ResponseAllOfPoliciesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPolicies200ResponseAllOfPoliciesInnerWithDefaults

`func NewListPolicies200ResponseAllOfPoliciesInnerWithDefaults() *ListPolicies200ResponseAllOfPoliciesInner`

NewListPolicies200ResponseAllOfPoliciesInnerWithDefaults instantiates a new ListPolicies200ResponseAllOfPoliciesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListPolicies200ResponseAllOfPoliciesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyType

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetPolicyType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetPolicyTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetPolicyType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSite

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetSite() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetSiteOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetSite(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetSite sets Site field to given value.

### HasSite

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUser

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetUser() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetUserOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetUser(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetRole() ListPolicies200ResponseAllOfPoliciesInnerRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetRoleOk() (*ListPolicies200ResponseAllOfPoliciesInnerRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetRole(v ListPolicies200ResponseAllOfPoliciesInnerRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRefType

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *ListPolicies200ResponseAllOfPoliciesInner) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *ListPolicies200ResponseAllOfPoliciesInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetEachUser

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### SetEachUserNil

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetEachUserNil(b bool)`

 SetEachUserNil sets the value for EachUser to be an explicit nil

### UnsetEachUser
`func (o *ListPolicies200ResponseAllOfPoliciesInner) UnsetEachUser()`

UnsetEachUser ensures that no value is present for EachUser, not even an explicit nil
### GetConfig

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetConfig() ListPolicies200ResponseAllOfPoliciesInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetConfigOk() (*ListPolicies200ResponseAllOfPoliciesInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetConfig(v ListPolicies200ResponseAllOfPoliciesInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetOwner() ListApprovals200ResponseAllOfApprovalsInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetOwnerOk() (*ListApprovals200ResponseAllOfApprovalsInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetOwner(v ListApprovals200ResponseAllOfApprovalsInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ListPolicies200ResponseAllOfPoliciesInner) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetAccounts

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetAccounts() []GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ListPolicies200ResponseAllOfPoliciesInner) GetAccountsOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetAccounts(v []GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ListPolicies200ResponseAllOfPoliciesInner) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *ListPolicies200ResponseAllOfPoliciesInner) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *ListPolicies200ResponseAllOfPoliciesInner) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


