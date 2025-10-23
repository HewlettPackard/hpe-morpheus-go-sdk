# AddClouds200ResponseAllOfZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**CostStatus** | Pointer to **NullableString** |  | [optional] 
**CostStatusMessage** | Pointer to **NullableString** |  | [optional] 
**CostStatusDate** | Pointer to **NullableTime** |  | [optional] 
**CostLastSyncDuration** | Pointer to **NullableInt64** |  | [optional] 
**CostLastSync** | Pointer to **NullableTime** |  | [optional] 
**ZoneType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ZoneTypeId** | Pointer to **int64** |  | [optional] 
**GuidanceMode** | Pointer to **NullableString** |  | [optional] 
**StorageMode** | Pointer to **string** |  | [optional] 
**AgentMode** | Pointer to **string** |  | [optional] 
**UserDataLinux** | Pointer to **NullableString** |  | [optional] 
**UserDataWindows** | Pointer to **NullableString** |  | [optional] 
**ConsoleKeymap** | Pointer to **NullableString** |  | [optional] 
**ContainerMode** | Pointer to **string** |  | [optional] 
**CostingMode** | Pointer to **NullableString** |  | [optional] 
**ServiceVersion** | Pointer to **NullableString** |  | [optional] 
**SecurityMode** | Pointer to **string** |  | [optional] 
**InventoryLevel** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**ApiProxy** | Pointer to **NullableString** |  | [optional] 
**ProvisioningProxy** | Pointer to **NullableString** |  | [optional] 
**NetworkDomain** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **NullableString** |  | [optional] 
**AutoRecoverPowerState** | Pointer to **bool** |  | [optional] 
**ScalePriority** | Pointer to **int64** |  | [optional] 
**DefaultDatastoreSyncActive** | Pointer to **bool** |  | [optional] 
**DefaultNetworkSyncActive** | Pointer to **bool** |  | [optional] 
**DefaultFolderSyncActive** | Pointer to **bool** |  | [optional] 
**DefaultSecurityGroupSyncActive** | Pointer to **bool** |  | [optional] 
**DefaultPoolSyncActive** | Pointer to **bool** |  | [optional] 
**DefaultPlanSyncActive** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**ListClouds200ResponseAllOfZonesInnerConfig**](ListClouds200ResponseAllOfZonesInnerConfig.md) |  | [optional] 
**Credential** | Pointer to [**ListClouds200ResponseAllOfZonesInnerCredential**](ListClouds200ResponseAllOfZonesInnerCredential.md) |  | [optional] 
**ImagePath** | Pointer to **NullableString** | Logo image URL | [optional] 
**DarkImagePath** | Pointer to **NullableString** | Dark logo image URL | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastSync** | Pointer to **NullableTime** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableInt64** |  | [optional] 
**NextRunDate** | Pointer to **NullableTime** |  | [optional] 
**Groups** | Pointer to [**[]ListClouds200ResponseAllOfZonesInnerGroupsInner**](ListClouds200ResponseAllOfZonesInnerGroupsInner.md) |  | [optional] 
**SecurityServer** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**NetworkServer** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Stats** | Pointer to [**ListClouds200ResponseAllOfZonesInnerStats**](ListClouds200ResponseAllOfZonesInnerStats.md) |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewAddClouds200ResponseAllOfZone

`func NewAddClouds200ResponseAllOfZone() *AddClouds200ResponseAllOfZone`

NewAddClouds200ResponseAllOfZone instantiates a new AddClouds200ResponseAllOfZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClouds200ResponseAllOfZoneWithDefaults

`func NewAddClouds200ResponseAllOfZoneWithDefaults() *AddClouds200ResponseAllOfZone`

NewAddClouds200ResponseAllOfZoneWithDefaults instantiates a new AddClouds200ResponseAllOfZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddClouds200ResponseAllOfZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddClouds200ResponseAllOfZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddClouds200ResponseAllOfZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddClouds200ResponseAllOfZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *AddClouds200ResponseAllOfZone) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AddClouds200ResponseAllOfZone) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AddClouds200ResponseAllOfZone) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AddClouds200ResponseAllOfZone) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *AddClouds200ResponseAllOfZone) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddClouds200ResponseAllOfZone) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddClouds200ResponseAllOfZone) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddClouds200ResponseAllOfZone) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AddClouds200ResponseAllOfZone) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddClouds200ResponseAllOfZone) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetName

`func (o *AddClouds200ResponseAllOfZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddClouds200ResponseAllOfZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddClouds200ResponseAllOfZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddClouds200ResponseAllOfZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddClouds200ResponseAllOfZone) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddClouds200ResponseAllOfZone) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddClouds200ResponseAllOfZone) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddClouds200ResponseAllOfZone) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabels

`func (o *AddClouds200ResponseAllOfZone) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddClouds200ResponseAllOfZone) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddClouds200ResponseAllOfZone) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddClouds200ResponseAllOfZone) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *AddClouds200ResponseAllOfZone) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AddClouds200ResponseAllOfZone) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AddClouds200ResponseAllOfZone) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AddClouds200ResponseAllOfZone) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AddClouds200ResponseAllOfZone) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AddClouds200ResponseAllOfZone) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetOwner

`func (o *AddClouds200ResponseAllOfZone) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddClouds200ResponseAllOfZone) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddClouds200ResponseAllOfZone) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddClouds200ResponseAllOfZone) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccountId

`func (o *AddClouds200ResponseAllOfZone) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddClouds200ResponseAllOfZone) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddClouds200ResponseAllOfZone) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddClouds200ResponseAllOfZone) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *AddClouds200ResponseAllOfZone) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddClouds200ResponseAllOfZone) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddClouds200ResponseAllOfZone) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddClouds200ResponseAllOfZone) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetVisibility

`func (o *AddClouds200ResponseAllOfZone) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddClouds200ResponseAllOfZone) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddClouds200ResponseAllOfZone) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddClouds200ResponseAllOfZone) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnabled

`func (o *AddClouds200ResponseAllOfZone) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddClouds200ResponseAllOfZone) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddClouds200ResponseAllOfZone) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddClouds200ResponseAllOfZone) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *AddClouds200ResponseAllOfZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddClouds200ResponseAllOfZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddClouds200ResponseAllOfZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddClouds200ResponseAllOfZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddClouds200ResponseAllOfZone) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddClouds200ResponseAllOfZone) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddClouds200ResponseAllOfZone) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddClouds200ResponseAllOfZone) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AddClouds200ResponseAllOfZone) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AddClouds200ResponseAllOfZone) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *AddClouds200ResponseAllOfZone) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *AddClouds200ResponseAllOfZone) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *AddClouds200ResponseAllOfZone) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *AddClouds200ResponseAllOfZone) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *AddClouds200ResponseAllOfZone) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *AddClouds200ResponseAllOfZone) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetCostStatus

`func (o *AddClouds200ResponseAllOfZone) GetCostStatus() string`

GetCostStatus returns the CostStatus field if non-nil, zero value otherwise.

### GetCostStatusOk

`func (o *AddClouds200ResponseAllOfZone) GetCostStatusOk() (*string, bool)`

GetCostStatusOk returns a tuple with the CostStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatus

`func (o *AddClouds200ResponseAllOfZone) SetCostStatus(v string)`

SetCostStatus sets CostStatus field to given value.

### HasCostStatus

`func (o *AddClouds200ResponseAllOfZone) HasCostStatus() bool`

HasCostStatus returns a boolean if a field has been set.

### SetCostStatusNil

`func (o *AddClouds200ResponseAllOfZone) SetCostStatusNil(b bool)`

 SetCostStatusNil sets the value for CostStatus to be an explicit nil

### UnsetCostStatus
`func (o *AddClouds200ResponseAllOfZone) UnsetCostStatus()`

UnsetCostStatus ensures that no value is present for CostStatus, not even an explicit nil
### GetCostStatusMessage

`func (o *AddClouds200ResponseAllOfZone) GetCostStatusMessage() string`

GetCostStatusMessage returns the CostStatusMessage field if non-nil, zero value otherwise.

### GetCostStatusMessageOk

`func (o *AddClouds200ResponseAllOfZone) GetCostStatusMessageOk() (*string, bool)`

GetCostStatusMessageOk returns a tuple with the CostStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatusMessage

`func (o *AddClouds200ResponseAllOfZone) SetCostStatusMessage(v string)`

SetCostStatusMessage sets CostStatusMessage field to given value.

### HasCostStatusMessage

`func (o *AddClouds200ResponseAllOfZone) HasCostStatusMessage() bool`

HasCostStatusMessage returns a boolean if a field has been set.

### SetCostStatusMessageNil

`func (o *AddClouds200ResponseAllOfZone) SetCostStatusMessageNil(b bool)`

 SetCostStatusMessageNil sets the value for CostStatusMessage to be an explicit nil

### UnsetCostStatusMessage
`func (o *AddClouds200ResponseAllOfZone) UnsetCostStatusMessage()`

UnsetCostStatusMessage ensures that no value is present for CostStatusMessage, not even an explicit nil
### GetCostStatusDate

`func (o *AddClouds200ResponseAllOfZone) GetCostStatusDate() time.Time`

GetCostStatusDate returns the CostStatusDate field if non-nil, zero value otherwise.

### GetCostStatusDateOk

`func (o *AddClouds200ResponseAllOfZone) GetCostStatusDateOk() (*time.Time, bool)`

GetCostStatusDateOk returns a tuple with the CostStatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatusDate

`func (o *AddClouds200ResponseAllOfZone) SetCostStatusDate(v time.Time)`

SetCostStatusDate sets CostStatusDate field to given value.

### HasCostStatusDate

`func (o *AddClouds200ResponseAllOfZone) HasCostStatusDate() bool`

HasCostStatusDate returns a boolean if a field has been set.

### SetCostStatusDateNil

`func (o *AddClouds200ResponseAllOfZone) SetCostStatusDateNil(b bool)`

 SetCostStatusDateNil sets the value for CostStatusDate to be an explicit nil

### UnsetCostStatusDate
`func (o *AddClouds200ResponseAllOfZone) UnsetCostStatusDate()`

UnsetCostStatusDate ensures that no value is present for CostStatusDate, not even an explicit nil
### GetCostLastSyncDuration

`func (o *AddClouds200ResponseAllOfZone) GetCostLastSyncDuration() int64`

GetCostLastSyncDuration returns the CostLastSyncDuration field if non-nil, zero value otherwise.

### GetCostLastSyncDurationOk

`func (o *AddClouds200ResponseAllOfZone) GetCostLastSyncDurationOk() (*int64, bool)`

GetCostLastSyncDurationOk returns a tuple with the CostLastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLastSyncDuration

`func (o *AddClouds200ResponseAllOfZone) SetCostLastSyncDuration(v int64)`

SetCostLastSyncDuration sets CostLastSyncDuration field to given value.

### HasCostLastSyncDuration

`func (o *AddClouds200ResponseAllOfZone) HasCostLastSyncDuration() bool`

HasCostLastSyncDuration returns a boolean if a field has been set.

### SetCostLastSyncDurationNil

`func (o *AddClouds200ResponseAllOfZone) SetCostLastSyncDurationNil(b bool)`

 SetCostLastSyncDurationNil sets the value for CostLastSyncDuration to be an explicit nil

### UnsetCostLastSyncDuration
`func (o *AddClouds200ResponseAllOfZone) UnsetCostLastSyncDuration()`

UnsetCostLastSyncDuration ensures that no value is present for CostLastSyncDuration, not even an explicit nil
### GetCostLastSync

`func (o *AddClouds200ResponseAllOfZone) GetCostLastSync() time.Time`

GetCostLastSync returns the CostLastSync field if non-nil, zero value otherwise.

### GetCostLastSyncOk

`func (o *AddClouds200ResponseAllOfZone) GetCostLastSyncOk() (*time.Time, bool)`

GetCostLastSyncOk returns a tuple with the CostLastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLastSync

`func (o *AddClouds200ResponseAllOfZone) SetCostLastSync(v time.Time)`

SetCostLastSync sets CostLastSync field to given value.

### HasCostLastSync

`func (o *AddClouds200ResponseAllOfZone) HasCostLastSync() bool`

HasCostLastSync returns a boolean if a field has been set.

### SetCostLastSyncNil

`func (o *AddClouds200ResponseAllOfZone) SetCostLastSyncNil(b bool)`

 SetCostLastSyncNil sets the value for CostLastSync to be an explicit nil

### UnsetCostLastSync
`func (o *AddClouds200ResponseAllOfZone) UnsetCostLastSync()`

UnsetCostLastSync ensures that no value is present for CostLastSync, not even an explicit nil
### GetZoneType

`func (o *AddClouds200ResponseAllOfZone) GetZoneType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *AddClouds200ResponseAllOfZone) GetZoneTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *AddClouds200ResponseAllOfZone) SetZoneType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *AddClouds200ResponseAllOfZone) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetZoneTypeId

`func (o *AddClouds200ResponseAllOfZone) GetZoneTypeId() int64`

GetZoneTypeId returns the ZoneTypeId field if non-nil, zero value otherwise.

### GetZoneTypeIdOk

`func (o *AddClouds200ResponseAllOfZone) GetZoneTypeIdOk() (*int64, bool)`

GetZoneTypeIdOk returns a tuple with the ZoneTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTypeId

`func (o *AddClouds200ResponseAllOfZone) SetZoneTypeId(v int64)`

SetZoneTypeId sets ZoneTypeId field to given value.

### HasZoneTypeId

`func (o *AddClouds200ResponseAllOfZone) HasZoneTypeId() bool`

HasZoneTypeId returns a boolean if a field has been set.

### GetGuidanceMode

`func (o *AddClouds200ResponseAllOfZone) GetGuidanceMode() string`

GetGuidanceMode returns the GuidanceMode field if non-nil, zero value otherwise.

### GetGuidanceModeOk

`func (o *AddClouds200ResponseAllOfZone) GetGuidanceModeOk() (*string, bool)`

GetGuidanceModeOk returns a tuple with the GuidanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidanceMode

`func (o *AddClouds200ResponseAllOfZone) SetGuidanceMode(v string)`

SetGuidanceMode sets GuidanceMode field to given value.

### HasGuidanceMode

`func (o *AddClouds200ResponseAllOfZone) HasGuidanceMode() bool`

HasGuidanceMode returns a boolean if a field has been set.

### SetGuidanceModeNil

`func (o *AddClouds200ResponseAllOfZone) SetGuidanceModeNil(b bool)`

 SetGuidanceModeNil sets the value for GuidanceMode to be an explicit nil

### UnsetGuidanceMode
`func (o *AddClouds200ResponseAllOfZone) UnsetGuidanceMode()`

UnsetGuidanceMode ensures that no value is present for GuidanceMode, not even an explicit nil
### GetStorageMode

`func (o *AddClouds200ResponseAllOfZone) GetStorageMode() string`

GetStorageMode returns the StorageMode field if non-nil, zero value otherwise.

### GetStorageModeOk

`func (o *AddClouds200ResponseAllOfZone) GetStorageModeOk() (*string, bool)`

GetStorageModeOk returns a tuple with the StorageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageMode

`func (o *AddClouds200ResponseAllOfZone) SetStorageMode(v string)`

SetStorageMode sets StorageMode field to given value.

### HasStorageMode

`func (o *AddClouds200ResponseAllOfZone) HasStorageMode() bool`

HasStorageMode returns a boolean if a field has been set.

### GetAgentMode

`func (o *AddClouds200ResponseAllOfZone) GetAgentMode() string`

GetAgentMode returns the AgentMode field if non-nil, zero value otherwise.

### GetAgentModeOk

`func (o *AddClouds200ResponseAllOfZone) GetAgentModeOk() (*string, bool)`

GetAgentModeOk returns a tuple with the AgentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMode

`func (o *AddClouds200ResponseAllOfZone) SetAgentMode(v string)`

SetAgentMode sets AgentMode field to given value.

### HasAgentMode

`func (o *AddClouds200ResponseAllOfZone) HasAgentMode() bool`

HasAgentMode returns a boolean if a field has been set.

### GetUserDataLinux

`func (o *AddClouds200ResponseAllOfZone) GetUserDataLinux() string`

GetUserDataLinux returns the UserDataLinux field if non-nil, zero value otherwise.

### GetUserDataLinuxOk

`func (o *AddClouds200ResponseAllOfZone) GetUserDataLinuxOk() (*string, bool)`

GetUserDataLinuxOk returns a tuple with the UserDataLinux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataLinux

`func (o *AddClouds200ResponseAllOfZone) SetUserDataLinux(v string)`

SetUserDataLinux sets UserDataLinux field to given value.

### HasUserDataLinux

`func (o *AddClouds200ResponseAllOfZone) HasUserDataLinux() bool`

HasUserDataLinux returns a boolean if a field has been set.

### SetUserDataLinuxNil

`func (o *AddClouds200ResponseAllOfZone) SetUserDataLinuxNil(b bool)`

 SetUserDataLinuxNil sets the value for UserDataLinux to be an explicit nil

### UnsetUserDataLinux
`func (o *AddClouds200ResponseAllOfZone) UnsetUserDataLinux()`

UnsetUserDataLinux ensures that no value is present for UserDataLinux, not even an explicit nil
### GetUserDataWindows

`func (o *AddClouds200ResponseAllOfZone) GetUserDataWindows() string`

GetUserDataWindows returns the UserDataWindows field if non-nil, zero value otherwise.

### GetUserDataWindowsOk

`func (o *AddClouds200ResponseAllOfZone) GetUserDataWindowsOk() (*string, bool)`

GetUserDataWindowsOk returns a tuple with the UserDataWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataWindows

`func (o *AddClouds200ResponseAllOfZone) SetUserDataWindows(v string)`

SetUserDataWindows sets UserDataWindows field to given value.

### HasUserDataWindows

`func (o *AddClouds200ResponseAllOfZone) HasUserDataWindows() bool`

HasUserDataWindows returns a boolean if a field has been set.

### SetUserDataWindowsNil

`func (o *AddClouds200ResponseAllOfZone) SetUserDataWindowsNil(b bool)`

 SetUserDataWindowsNil sets the value for UserDataWindows to be an explicit nil

### UnsetUserDataWindows
`func (o *AddClouds200ResponseAllOfZone) UnsetUserDataWindows()`

UnsetUserDataWindows ensures that no value is present for UserDataWindows, not even an explicit nil
### GetConsoleKeymap

`func (o *AddClouds200ResponseAllOfZone) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *AddClouds200ResponseAllOfZone) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *AddClouds200ResponseAllOfZone) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *AddClouds200ResponseAllOfZone) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### SetConsoleKeymapNil

`func (o *AddClouds200ResponseAllOfZone) SetConsoleKeymapNil(b bool)`

 SetConsoleKeymapNil sets the value for ConsoleKeymap to be an explicit nil

### UnsetConsoleKeymap
`func (o *AddClouds200ResponseAllOfZone) UnsetConsoleKeymap()`

UnsetConsoleKeymap ensures that no value is present for ConsoleKeymap, not even an explicit nil
### GetContainerMode

`func (o *AddClouds200ResponseAllOfZone) GetContainerMode() string`

GetContainerMode returns the ContainerMode field if non-nil, zero value otherwise.

### GetContainerModeOk

`func (o *AddClouds200ResponseAllOfZone) GetContainerModeOk() (*string, bool)`

GetContainerModeOk returns a tuple with the ContainerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMode

`func (o *AddClouds200ResponseAllOfZone) SetContainerMode(v string)`

SetContainerMode sets ContainerMode field to given value.

### HasContainerMode

`func (o *AddClouds200ResponseAllOfZone) HasContainerMode() bool`

HasContainerMode returns a boolean if a field has been set.

### GetCostingMode

`func (o *AddClouds200ResponseAllOfZone) GetCostingMode() string`

GetCostingMode returns the CostingMode field if non-nil, zero value otherwise.

### GetCostingModeOk

`func (o *AddClouds200ResponseAllOfZone) GetCostingModeOk() (*string, bool)`

GetCostingModeOk returns a tuple with the CostingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingMode

`func (o *AddClouds200ResponseAllOfZone) SetCostingMode(v string)`

SetCostingMode sets CostingMode field to given value.

### HasCostingMode

`func (o *AddClouds200ResponseAllOfZone) HasCostingMode() bool`

HasCostingMode returns a boolean if a field has been set.

### SetCostingModeNil

`func (o *AddClouds200ResponseAllOfZone) SetCostingModeNil(b bool)`

 SetCostingModeNil sets the value for CostingMode to be an explicit nil

### UnsetCostingMode
`func (o *AddClouds200ResponseAllOfZone) UnsetCostingMode()`

UnsetCostingMode ensures that no value is present for CostingMode, not even an explicit nil
### GetServiceVersion

`func (o *AddClouds200ResponseAllOfZone) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *AddClouds200ResponseAllOfZone) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *AddClouds200ResponseAllOfZone) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *AddClouds200ResponseAllOfZone) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### SetServiceVersionNil

`func (o *AddClouds200ResponseAllOfZone) SetServiceVersionNil(b bool)`

 SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil

### UnsetServiceVersion
`func (o *AddClouds200ResponseAllOfZone) UnsetServiceVersion()`

UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
### GetSecurityMode

`func (o *AddClouds200ResponseAllOfZone) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *AddClouds200ResponseAllOfZone) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *AddClouds200ResponseAllOfZone) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *AddClouds200ResponseAllOfZone) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *AddClouds200ResponseAllOfZone) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *AddClouds200ResponseAllOfZone) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *AddClouds200ResponseAllOfZone) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *AddClouds200ResponseAllOfZone) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetTimezone

`func (o *AddClouds200ResponseAllOfZone) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AddClouds200ResponseAllOfZone) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AddClouds200ResponseAllOfZone) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AddClouds200ResponseAllOfZone) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *AddClouds200ResponseAllOfZone) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *AddClouds200ResponseAllOfZone) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetApiProxy

`func (o *AddClouds200ResponseAllOfZone) GetApiProxy() string`

GetApiProxy returns the ApiProxy field if non-nil, zero value otherwise.

### GetApiProxyOk

`func (o *AddClouds200ResponseAllOfZone) GetApiProxyOk() (*string, bool)`

GetApiProxyOk returns a tuple with the ApiProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProxy

`func (o *AddClouds200ResponseAllOfZone) SetApiProxy(v string)`

SetApiProxy sets ApiProxy field to given value.

### HasApiProxy

`func (o *AddClouds200ResponseAllOfZone) HasApiProxy() bool`

HasApiProxy returns a boolean if a field has been set.

### SetApiProxyNil

`func (o *AddClouds200ResponseAllOfZone) SetApiProxyNil(b bool)`

 SetApiProxyNil sets the value for ApiProxy to be an explicit nil

### UnsetApiProxy
`func (o *AddClouds200ResponseAllOfZone) UnsetApiProxy()`

UnsetApiProxy ensures that no value is present for ApiProxy, not even an explicit nil
### GetProvisioningProxy

`func (o *AddClouds200ResponseAllOfZone) GetProvisioningProxy() string`

GetProvisioningProxy returns the ProvisioningProxy field if non-nil, zero value otherwise.

### GetProvisioningProxyOk

`func (o *AddClouds200ResponseAllOfZone) GetProvisioningProxyOk() (*string, bool)`

GetProvisioningProxyOk returns a tuple with the ProvisioningProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProxy

`func (o *AddClouds200ResponseAllOfZone) SetProvisioningProxy(v string)`

SetProvisioningProxy sets ProvisioningProxy field to given value.

### HasProvisioningProxy

`func (o *AddClouds200ResponseAllOfZone) HasProvisioningProxy() bool`

HasProvisioningProxy returns a boolean if a field has been set.

### SetProvisioningProxyNil

`func (o *AddClouds200ResponseAllOfZone) SetProvisioningProxyNil(b bool)`

 SetProvisioningProxyNil sets the value for ProvisioningProxy to be an explicit nil

### UnsetProvisioningProxy
`func (o *AddClouds200ResponseAllOfZone) UnsetProvisioningProxy()`

UnsetProvisioningProxy ensures that no value is present for ProvisioningProxy, not even an explicit nil
### GetNetworkDomain

`func (o *AddClouds200ResponseAllOfZone) GetNetworkDomain() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *AddClouds200ResponseAllOfZone) GetNetworkDomainOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *AddClouds200ResponseAllOfZone) SetNetworkDomain(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *AddClouds200ResponseAllOfZone) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetDomainName

`func (o *AddClouds200ResponseAllOfZone) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AddClouds200ResponseAllOfZone) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *AddClouds200ResponseAllOfZone) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *AddClouds200ResponseAllOfZone) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetRegionCode

`func (o *AddClouds200ResponseAllOfZone) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *AddClouds200ResponseAllOfZone) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *AddClouds200ResponseAllOfZone) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *AddClouds200ResponseAllOfZone) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *AddClouds200ResponseAllOfZone) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *AddClouds200ResponseAllOfZone) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetAutoRecoverPowerState

`func (o *AddClouds200ResponseAllOfZone) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *AddClouds200ResponseAllOfZone) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *AddClouds200ResponseAllOfZone) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *AddClouds200ResponseAllOfZone) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetScalePriority

`func (o *AddClouds200ResponseAllOfZone) GetScalePriority() int64`

GetScalePriority returns the ScalePriority field if non-nil, zero value otherwise.

### GetScalePriorityOk

`func (o *AddClouds200ResponseAllOfZone) GetScalePriorityOk() (*int64, bool)`

GetScalePriorityOk returns a tuple with the ScalePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalePriority

`func (o *AddClouds200ResponseAllOfZone) SetScalePriority(v int64)`

SetScalePriority sets ScalePriority field to given value.

### HasScalePriority

`func (o *AddClouds200ResponseAllOfZone) HasScalePriority() bool`

HasScalePriority returns a boolean if a field has been set.

### GetDefaultDatastoreSyncActive

`func (o *AddClouds200ResponseAllOfZone) GetDefaultDatastoreSyncActive() bool`

GetDefaultDatastoreSyncActive returns the DefaultDatastoreSyncActive field if non-nil, zero value otherwise.

### GetDefaultDatastoreSyncActiveOk

`func (o *AddClouds200ResponseAllOfZone) GetDefaultDatastoreSyncActiveOk() (*bool, bool)`

GetDefaultDatastoreSyncActiveOk returns a tuple with the DefaultDatastoreSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDatastoreSyncActive

`func (o *AddClouds200ResponseAllOfZone) SetDefaultDatastoreSyncActive(v bool)`

SetDefaultDatastoreSyncActive sets DefaultDatastoreSyncActive field to given value.

### HasDefaultDatastoreSyncActive

`func (o *AddClouds200ResponseAllOfZone) HasDefaultDatastoreSyncActive() bool`

HasDefaultDatastoreSyncActive returns a boolean if a field has been set.

### GetDefaultNetworkSyncActive

`func (o *AddClouds200ResponseAllOfZone) GetDefaultNetworkSyncActive() bool`

GetDefaultNetworkSyncActive returns the DefaultNetworkSyncActive field if non-nil, zero value otherwise.

### GetDefaultNetworkSyncActiveOk

`func (o *AddClouds200ResponseAllOfZone) GetDefaultNetworkSyncActiveOk() (*bool, bool)`

GetDefaultNetworkSyncActiveOk returns a tuple with the DefaultNetworkSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkSyncActive

`func (o *AddClouds200ResponseAllOfZone) SetDefaultNetworkSyncActive(v bool)`

SetDefaultNetworkSyncActive sets DefaultNetworkSyncActive field to given value.

### HasDefaultNetworkSyncActive

`func (o *AddClouds200ResponseAllOfZone) HasDefaultNetworkSyncActive() bool`

HasDefaultNetworkSyncActive returns a boolean if a field has been set.

### GetDefaultFolderSyncActive

`func (o *AddClouds200ResponseAllOfZone) GetDefaultFolderSyncActive() bool`

GetDefaultFolderSyncActive returns the DefaultFolderSyncActive field if non-nil, zero value otherwise.

### GetDefaultFolderSyncActiveOk

`func (o *AddClouds200ResponseAllOfZone) GetDefaultFolderSyncActiveOk() (*bool, bool)`

GetDefaultFolderSyncActiveOk returns a tuple with the DefaultFolderSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolderSyncActive

`func (o *AddClouds200ResponseAllOfZone) SetDefaultFolderSyncActive(v bool)`

SetDefaultFolderSyncActive sets DefaultFolderSyncActive field to given value.

### HasDefaultFolderSyncActive

`func (o *AddClouds200ResponseAllOfZone) HasDefaultFolderSyncActive() bool`

HasDefaultFolderSyncActive returns a boolean if a field has been set.

### GetDefaultSecurityGroupSyncActive

`func (o *AddClouds200ResponseAllOfZone) GetDefaultSecurityGroupSyncActive() bool`

GetDefaultSecurityGroupSyncActive returns the DefaultSecurityGroupSyncActive field if non-nil, zero value otherwise.

### GetDefaultSecurityGroupSyncActiveOk

`func (o *AddClouds200ResponseAllOfZone) GetDefaultSecurityGroupSyncActiveOk() (*bool, bool)`

GetDefaultSecurityGroupSyncActiveOk returns a tuple with the DefaultSecurityGroupSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurityGroupSyncActive

`func (o *AddClouds200ResponseAllOfZone) SetDefaultSecurityGroupSyncActive(v bool)`

SetDefaultSecurityGroupSyncActive sets DefaultSecurityGroupSyncActive field to given value.

### HasDefaultSecurityGroupSyncActive

`func (o *AddClouds200ResponseAllOfZone) HasDefaultSecurityGroupSyncActive() bool`

HasDefaultSecurityGroupSyncActive returns a boolean if a field has been set.

### GetDefaultPoolSyncActive

`func (o *AddClouds200ResponseAllOfZone) GetDefaultPoolSyncActive() bool`

GetDefaultPoolSyncActive returns the DefaultPoolSyncActive field if non-nil, zero value otherwise.

### GetDefaultPoolSyncActiveOk

`func (o *AddClouds200ResponseAllOfZone) GetDefaultPoolSyncActiveOk() (*bool, bool)`

GetDefaultPoolSyncActiveOk returns a tuple with the DefaultPoolSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPoolSyncActive

`func (o *AddClouds200ResponseAllOfZone) SetDefaultPoolSyncActive(v bool)`

SetDefaultPoolSyncActive sets DefaultPoolSyncActive field to given value.

### HasDefaultPoolSyncActive

`func (o *AddClouds200ResponseAllOfZone) HasDefaultPoolSyncActive() bool`

HasDefaultPoolSyncActive returns a boolean if a field has been set.

### GetDefaultPlanSyncActive

`func (o *AddClouds200ResponseAllOfZone) GetDefaultPlanSyncActive() bool`

GetDefaultPlanSyncActive returns the DefaultPlanSyncActive field if non-nil, zero value otherwise.

### GetDefaultPlanSyncActiveOk

`func (o *AddClouds200ResponseAllOfZone) GetDefaultPlanSyncActiveOk() (*bool, bool)`

GetDefaultPlanSyncActiveOk returns a tuple with the DefaultPlanSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlanSyncActive

`func (o *AddClouds200ResponseAllOfZone) SetDefaultPlanSyncActive(v bool)`

SetDefaultPlanSyncActive sets DefaultPlanSyncActive field to given value.

### HasDefaultPlanSyncActive

`func (o *AddClouds200ResponseAllOfZone) HasDefaultPlanSyncActive() bool`

HasDefaultPlanSyncActive returns a boolean if a field has been set.

### GetConfig

`func (o *AddClouds200ResponseAllOfZone) GetConfig() ListClouds200ResponseAllOfZonesInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddClouds200ResponseAllOfZone) GetConfigOk() (*ListClouds200ResponseAllOfZonesInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddClouds200ResponseAllOfZone) SetConfig(v ListClouds200ResponseAllOfZonesInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddClouds200ResponseAllOfZone) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *AddClouds200ResponseAllOfZone) GetCredential() ListClouds200ResponseAllOfZonesInnerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AddClouds200ResponseAllOfZone) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AddClouds200ResponseAllOfZone) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AddClouds200ResponseAllOfZone) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetImagePath

`func (o *AddClouds200ResponseAllOfZone) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *AddClouds200ResponseAllOfZone) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *AddClouds200ResponseAllOfZone) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *AddClouds200ResponseAllOfZone) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### SetImagePathNil

`func (o *AddClouds200ResponseAllOfZone) SetImagePathNil(b bool)`

 SetImagePathNil sets the value for ImagePath to be an explicit nil

### UnsetImagePath
`func (o *AddClouds200ResponseAllOfZone) UnsetImagePath()`

UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
### GetDarkImagePath

`func (o *AddClouds200ResponseAllOfZone) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *AddClouds200ResponseAllOfZone) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *AddClouds200ResponseAllOfZone) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *AddClouds200ResponseAllOfZone) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### SetDarkImagePathNil

`func (o *AddClouds200ResponseAllOfZone) SetDarkImagePathNil(b bool)`

 SetDarkImagePathNil sets the value for DarkImagePath to be an explicit nil

### UnsetDarkImagePath
`func (o *AddClouds200ResponseAllOfZone) UnsetDarkImagePath()`

UnsetDarkImagePath ensures that no value is present for DarkImagePath, not even an explicit nil
### GetDateCreated

`func (o *AddClouds200ResponseAllOfZone) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddClouds200ResponseAllOfZone) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddClouds200ResponseAllOfZone) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddClouds200ResponseAllOfZone) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddClouds200ResponseAllOfZone) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddClouds200ResponseAllOfZone) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddClouds200ResponseAllOfZone) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddClouds200ResponseAllOfZone) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastSync

`func (o *AddClouds200ResponseAllOfZone) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AddClouds200ResponseAllOfZone) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AddClouds200ResponseAllOfZone) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AddClouds200ResponseAllOfZone) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *AddClouds200ResponseAllOfZone) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *AddClouds200ResponseAllOfZone) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *AddClouds200ResponseAllOfZone) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *AddClouds200ResponseAllOfZone) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *AddClouds200ResponseAllOfZone) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *AddClouds200ResponseAllOfZone) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *AddClouds200ResponseAllOfZone) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *AddClouds200ResponseAllOfZone) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetNextRunDate

`func (o *AddClouds200ResponseAllOfZone) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *AddClouds200ResponseAllOfZone) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *AddClouds200ResponseAllOfZone) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *AddClouds200ResponseAllOfZone) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *AddClouds200ResponseAllOfZone) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *AddClouds200ResponseAllOfZone) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetGroups

`func (o *AddClouds200ResponseAllOfZone) GetGroups() []ListClouds200ResponseAllOfZonesInnerGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AddClouds200ResponseAllOfZone) GetGroupsOk() (*[]ListClouds200ResponseAllOfZonesInnerGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AddClouds200ResponseAllOfZone) SetGroups(v []ListClouds200ResponseAllOfZonesInnerGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AddClouds200ResponseAllOfZone) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetSecurityServer

`func (o *AddClouds200ResponseAllOfZone) GetSecurityServer() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *AddClouds200ResponseAllOfZone) GetSecurityServerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *AddClouds200ResponseAllOfZone) SetSecurityServer(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *AddClouds200ResponseAllOfZone) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetNetworkServer

`func (o *AddClouds200ResponseAllOfZone) GetNetworkServer() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *AddClouds200ResponseAllOfZone) GetNetworkServerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *AddClouds200ResponseAllOfZone) SetNetworkServer(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *AddClouds200ResponseAllOfZone) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetStats

`func (o *AddClouds200ResponseAllOfZone) GetStats() ListClouds200ResponseAllOfZonesInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AddClouds200ResponseAllOfZone) GetStatsOk() (*ListClouds200ResponseAllOfZonesInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AddClouds200ResponseAllOfZone) SetStats(v ListClouds200ResponseAllOfZonesInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AddClouds200ResponseAllOfZone) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *AddClouds200ResponseAllOfZone) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *AddClouds200ResponseAllOfZone) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *AddClouds200ResponseAllOfZone) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *AddClouds200ResponseAllOfZone) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


