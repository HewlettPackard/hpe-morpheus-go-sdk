# Policy

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

### NewPolicy

`func NewPolicy() *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Policy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Policy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Policy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Policy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Policy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Policy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Policy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Policy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyType

`func (o *Policy) GetPolicyType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *Policy) GetPolicyTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *Policy) SetPolicyType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *Policy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *Policy) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Policy) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Policy) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Policy) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSite

`func (o *Policy) GetSite() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Policy) GetSiteOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Policy) SetSite(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetSite sets Site field to given value.

### HasSite

`func (o *Policy) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUser

`func (o *Policy) GetUser() GetAlerts200ResponseAllOfChecksInnerCreatedBy`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Policy) GetUserOk() (*GetAlerts200ResponseAllOfChecksInnerCreatedBy, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Policy) SetUser(v GetAlerts200ResponseAllOfChecksInnerCreatedBy)`

SetUser sets User field to given value.

### HasUser

`func (o *Policy) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *Policy) GetRole() AddPolicies200ResponseAllOfPolicyRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Policy) GetRoleOk() (*AddPolicies200ResponseAllOfPolicyRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Policy) SetRole(v AddPolicies200ResponseAllOfPolicyRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *Policy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRefType

`func (o *Policy) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *Policy) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *Policy) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *Policy) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *Policy) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Policy) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Policy) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Policy) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetEachUser

`func (o *Policy) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *Policy) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *Policy) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *Policy) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### GetConfig

`func (o *Policy) GetConfig() AddPolicies200ResponseAllOfPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Policy) GetConfigOk() (*AddPolicies200ResponseAllOfPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Policy) SetConfig(v AddPolicies200ResponseAllOfPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Policy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *Policy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Policy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Policy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Policy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *Policy) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Policy) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Policy) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Policy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccounts

`func (o *Policy) GetAccounts() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Policy) GetAccountsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Policy) SetAccounts(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Policy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


