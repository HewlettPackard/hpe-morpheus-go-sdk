# AddPolicies200ResponseAllOfPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**PolicyType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Site** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**User** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerCreatedBy**](GetAlerts200ResponseAllOfChecksInnerCreatedBy.md) |  | [optional] 
**Role** | Pointer to [**AddPolicies200ResponseAllOfPolicyRole**](AddPolicies200ResponseAllOfPolicyRole.md) |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**EachUser** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**AddPolicies200ResponseAllOfPolicyConfig**](AddPolicies200ResponseAllOfPolicyConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Accounts** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewAddPolicies200ResponseAllOfPolicy

`func NewAddPolicies200ResponseAllOfPolicy() *AddPolicies200ResponseAllOfPolicy`

NewAddPolicies200ResponseAllOfPolicy instantiates a new AddPolicies200ResponseAllOfPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPolicies200ResponseAllOfPolicyWithDefaults

`func NewAddPolicies200ResponseAllOfPolicyWithDefaults() *AddPolicies200ResponseAllOfPolicy`

NewAddPolicies200ResponseAllOfPolicyWithDefaults instantiates a new AddPolicies200ResponseAllOfPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddPolicies200ResponseAllOfPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddPolicies200ResponseAllOfPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddPolicies200ResponseAllOfPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddPolicies200ResponseAllOfPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddPolicies200ResponseAllOfPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddPolicies200ResponseAllOfPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddPolicies200ResponseAllOfPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPolicies200ResponseAllOfPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPolicies200ResponseAllOfPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyType

`func (o *AddPolicies200ResponseAllOfPolicy) GetPolicyType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetPolicyTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *AddPolicies200ResponseAllOfPolicy) SetPolicyType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *AddPolicies200ResponseAllOfPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *AddPolicies200ResponseAllOfPolicy) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddPolicies200ResponseAllOfPolicy) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddPolicies200ResponseAllOfPolicy) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSite

`func (o *AddPolicies200ResponseAllOfPolicy) GetSite() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetSiteOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *AddPolicies200ResponseAllOfPolicy) SetSite(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetSite sets Site field to given value.

### HasSite

`func (o *AddPolicies200ResponseAllOfPolicy) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUser

`func (o *AddPolicies200ResponseAllOfPolicy) GetUser() GetAlerts200ResponseAllOfChecksInnerCreatedBy`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetUserOk() (*GetAlerts200ResponseAllOfChecksInnerCreatedBy, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AddPolicies200ResponseAllOfPolicy) SetUser(v GetAlerts200ResponseAllOfChecksInnerCreatedBy)`

SetUser sets User field to given value.

### HasUser

`func (o *AddPolicies200ResponseAllOfPolicy) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *AddPolicies200ResponseAllOfPolicy) GetRole() AddPolicies200ResponseAllOfPolicyRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetRoleOk() (*AddPolicies200ResponseAllOfPolicyRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddPolicies200ResponseAllOfPolicy) SetRole(v AddPolicies200ResponseAllOfPolicyRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *AddPolicies200ResponseAllOfPolicy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRefType

`func (o *AddPolicies200ResponseAllOfPolicy) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *AddPolicies200ResponseAllOfPolicy) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *AddPolicies200ResponseAllOfPolicy) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *AddPolicies200ResponseAllOfPolicy) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AddPolicies200ResponseAllOfPolicy) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AddPolicies200ResponseAllOfPolicy) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetEachUser

`func (o *AddPolicies200ResponseAllOfPolicy) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *AddPolicies200ResponseAllOfPolicy) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *AddPolicies200ResponseAllOfPolicy) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### GetConfig

`func (o *AddPolicies200ResponseAllOfPolicy) GetConfig() AddPolicies200ResponseAllOfPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetConfigOk() (*AddPolicies200ResponseAllOfPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddPolicies200ResponseAllOfPolicy) SetConfig(v AddPolicies200ResponseAllOfPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddPolicies200ResponseAllOfPolicy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPolicies200ResponseAllOfPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPolicies200ResponseAllOfPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddPolicies200ResponseAllOfPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *AddPolicies200ResponseAllOfPolicy) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddPolicies200ResponseAllOfPolicy) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddPolicies200ResponseAllOfPolicy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccounts

`func (o *AddPolicies200ResponseAllOfPolicy) GetAccounts() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AddPolicies200ResponseAllOfPolicy) GetAccountsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AddPolicies200ResponseAllOfPolicy) SetAccounts(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AddPolicies200ResponseAllOfPolicy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


