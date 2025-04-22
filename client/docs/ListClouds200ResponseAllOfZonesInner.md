# ListClouds200ResponseAllOfZonesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**CostStatus** | Pointer to **string** |  | [optional] 
**CostStatusMessage** | Pointer to **string** |  | [optional] 
**CostStatusDate** | Pointer to **time.Time** |  | [optional] 
**CostLastSyncDuration** | Pointer to **int64** |  | [optional] 
**CostLastSync** | Pointer to **time.Time** |  | [optional] 
**ZoneType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ZoneTypeId** | Pointer to **int64** |  | [optional] 
**GuidanceMode** | Pointer to **string** |  | [optional] 
**StorageMode** | Pointer to **string** |  | [optional] 
**AgentMode** | Pointer to **string** |  | [optional] 
**UserDataLinux** | Pointer to **string** |  | [optional] 
**UserDataWindows** | Pointer to **string** |  | [optional] 
**ConsoleKeymap** | Pointer to **string** |  | [optional] 
**ContainerMode** | Pointer to **string** |  | [optional] 
**CostingMode** | Pointer to **string** |  | [optional] 
**ServiceVersion** | Pointer to **string** |  | [optional] 
**SecurityMode** | Pointer to **string** |  | [optional] 
**InventoryLevel** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**ApiProxy** | Pointer to **string** |  | [optional] 
**ProvisioningProxy** | Pointer to **string** |  | [optional] 
**NetworkDomain** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
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
**ImagePath** | Pointer to **string** | Logo image URL | [optional] 
**DarkImagePath** | Pointer to **string** | Dark logo image URL | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastSync** | Pointer to **time.Time** |  | [optional] 
**LastSyncDuration** | Pointer to **int64** |  | [optional] 
**NextRunDate** | Pointer to **time.Time** |  | [optional] 
**Groups** | Pointer to [**[]ListClouds200ResponseAllOfZonesInnerGroupsInner**](ListClouds200ResponseAllOfZonesInnerGroupsInner.md) |  | [optional] 
**SecurityServer** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**NetworkServer** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Stats** | Pointer to [**ListClouds200ResponseAllOfZonesInnerStats**](ListClouds200ResponseAllOfZonesInnerStats.md) |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewListClouds200ResponseAllOfZonesInner

`func NewListClouds200ResponseAllOfZonesInner() *ListClouds200ResponseAllOfZonesInner`

NewListClouds200ResponseAllOfZonesInner instantiates a new ListClouds200ResponseAllOfZonesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClouds200ResponseAllOfZonesInnerWithDefaults

`func NewListClouds200ResponseAllOfZonesInnerWithDefaults() *ListClouds200ResponseAllOfZonesInner`

NewListClouds200ResponseAllOfZonesInnerWithDefaults instantiates a new ListClouds200ResponseAllOfZonesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClouds200ResponseAllOfZonesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClouds200ResponseAllOfZonesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClouds200ResponseAllOfZonesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListClouds200ResponseAllOfZonesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListClouds200ResponseAllOfZonesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListClouds200ResponseAllOfZonesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *ListClouds200ResponseAllOfZonesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListClouds200ResponseAllOfZonesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListClouds200ResponseAllOfZonesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetName

`func (o *ListClouds200ResponseAllOfZonesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClouds200ResponseAllOfZonesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClouds200ResponseAllOfZonesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListClouds200ResponseAllOfZonesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClouds200ResponseAllOfZonesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClouds200ResponseAllOfZonesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabels

`func (o *ListClouds200ResponseAllOfZonesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListClouds200ResponseAllOfZonesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListClouds200ResponseAllOfZonesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *ListClouds200ResponseAllOfZonesInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ListClouds200ResponseAllOfZonesInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ListClouds200ResponseAllOfZonesInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetOwner

`func (o *ListClouds200ResponseAllOfZonesInner) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListClouds200ResponseAllOfZonesInner) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListClouds200ResponseAllOfZonesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccountId

`func (o *ListClouds200ResponseAllOfZonesInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListClouds200ResponseAllOfZonesInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListClouds200ResponseAllOfZonesInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *ListClouds200ResponseAllOfZonesInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListClouds200ResponseAllOfZonesInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListClouds200ResponseAllOfZonesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetVisibility

`func (o *ListClouds200ResponseAllOfZonesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListClouds200ResponseAllOfZonesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListClouds200ResponseAllOfZonesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnabled

`func (o *ListClouds200ResponseAllOfZonesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListClouds200ResponseAllOfZonesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListClouds200ResponseAllOfZonesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *ListClouds200ResponseAllOfZonesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClouds200ResponseAllOfZonesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClouds200ResponseAllOfZonesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListClouds200ResponseAllOfZonesInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListClouds200ResponseAllOfZonesInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListClouds200ResponseAllOfZonesInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListClouds200ResponseAllOfZonesInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListClouds200ResponseAllOfZonesInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListClouds200ResponseAllOfZonesInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetCostStatus

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostStatus() string`

GetCostStatus returns the CostStatus field if non-nil, zero value otherwise.

### GetCostStatusOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostStatusOk() (*string, bool)`

GetCostStatusOk returns a tuple with the CostStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatus

`func (o *ListClouds200ResponseAllOfZonesInner) SetCostStatus(v string)`

SetCostStatus sets CostStatus field to given value.

### HasCostStatus

`func (o *ListClouds200ResponseAllOfZonesInner) HasCostStatus() bool`

HasCostStatus returns a boolean if a field has been set.

### GetCostStatusMessage

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostStatusMessage() string`

GetCostStatusMessage returns the CostStatusMessage field if non-nil, zero value otherwise.

### GetCostStatusMessageOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostStatusMessageOk() (*string, bool)`

GetCostStatusMessageOk returns a tuple with the CostStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatusMessage

`func (o *ListClouds200ResponseAllOfZonesInner) SetCostStatusMessage(v string)`

SetCostStatusMessage sets CostStatusMessage field to given value.

### HasCostStatusMessage

`func (o *ListClouds200ResponseAllOfZonesInner) HasCostStatusMessage() bool`

HasCostStatusMessage returns a boolean if a field has been set.

### GetCostStatusDate

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostStatusDate() time.Time`

GetCostStatusDate returns the CostStatusDate field if non-nil, zero value otherwise.

### GetCostStatusDateOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostStatusDateOk() (*time.Time, bool)`

GetCostStatusDateOk returns a tuple with the CostStatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatusDate

`func (o *ListClouds200ResponseAllOfZonesInner) SetCostStatusDate(v time.Time)`

SetCostStatusDate sets CostStatusDate field to given value.

### HasCostStatusDate

`func (o *ListClouds200ResponseAllOfZonesInner) HasCostStatusDate() bool`

HasCostStatusDate returns a boolean if a field has been set.

### GetCostLastSyncDuration

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostLastSyncDuration() int64`

GetCostLastSyncDuration returns the CostLastSyncDuration field if non-nil, zero value otherwise.

### GetCostLastSyncDurationOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostLastSyncDurationOk() (*int64, bool)`

GetCostLastSyncDurationOk returns a tuple with the CostLastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLastSyncDuration

`func (o *ListClouds200ResponseAllOfZonesInner) SetCostLastSyncDuration(v int64)`

SetCostLastSyncDuration sets CostLastSyncDuration field to given value.

### HasCostLastSyncDuration

`func (o *ListClouds200ResponseAllOfZonesInner) HasCostLastSyncDuration() bool`

HasCostLastSyncDuration returns a boolean if a field has been set.

### GetCostLastSync

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostLastSync() time.Time`

GetCostLastSync returns the CostLastSync field if non-nil, zero value otherwise.

### GetCostLastSyncOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostLastSyncOk() (*time.Time, bool)`

GetCostLastSyncOk returns a tuple with the CostLastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLastSync

`func (o *ListClouds200ResponseAllOfZonesInner) SetCostLastSync(v time.Time)`

SetCostLastSync sets CostLastSync field to given value.

### HasCostLastSync

`func (o *ListClouds200ResponseAllOfZonesInner) HasCostLastSync() bool`

HasCostLastSync returns a boolean if a field has been set.

### GetZoneType

`func (o *ListClouds200ResponseAllOfZonesInner) GetZoneType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetZoneTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *ListClouds200ResponseAllOfZonesInner) SetZoneType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *ListClouds200ResponseAllOfZonesInner) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetZoneTypeId

`func (o *ListClouds200ResponseAllOfZonesInner) GetZoneTypeId() int64`

GetZoneTypeId returns the ZoneTypeId field if non-nil, zero value otherwise.

### GetZoneTypeIdOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetZoneTypeIdOk() (*int64, bool)`

GetZoneTypeIdOk returns a tuple with the ZoneTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTypeId

`func (o *ListClouds200ResponseAllOfZonesInner) SetZoneTypeId(v int64)`

SetZoneTypeId sets ZoneTypeId field to given value.

### HasZoneTypeId

`func (o *ListClouds200ResponseAllOfZonesInner) HasZoneTypeId() bool`

HasZoneTypeId returns a boolean if a field has been set.

### GetGuidanceMode

`func (o *ListClouds200ResponseAllOfZonesInner) GetGuidanceMode() string`

GetGuidanceMode returns the GuidanceMode field if non-nil, zero value otherwise.

### GetGuidanceModeOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetGuidanceModeOk() (*string, bool)`

GetGuidanceModeOk returns a tuple with the GuidanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidanceMode

`func (o *ListClouds200ResponseAllOfZonesInner) SetGuidanceMode(v string)`

SetGuidanceMode sets GuidanceMode field to given value.

### HasGuidanceMode

`func (o *ListClouds200ResponseAllOfZonesInner) HasGuidanceMode() bool`

HasGuidanceMode returns a boolean if a field has been set.

### GetStorageMode

`func (o *ListClouds200ResponseAllOfZonesInner) GetStorageMode() string`

GetStorageMode returns the StorageMode field if non-nil, zero value otherwise.

### GetStorageModeOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetStorageModeOk() (*string, bool)`

GetStorageModeOk returns a tuple with the StorageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageMode

`func (o *ListClouds200ResponseAllOfZonesInner) SetStorageMode(v string)`

SetStorageMode sets StorageMode field to given value.

### HasStorageMode

`func (o *ListClouds200ResponseAllOfZonesInner) HasStorageMode() bool`

HasStorageMode returns a boolean if a field has been set.

### GetAgentMode

`func (o *ListClouds200ResponseAllOfZonesInner) GetAgentMode() string`

GetAgentMode returns the AgentMode field if non-nil, zero value otherwise.

### GetAgentModeOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetAgentModeOk() (*string, bool)`

GetAgentModeOk returns a tuple with the AgentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMode

`func (o *ListClouds200ResponseAllOfZonesInner) SetAgentMode(v string)`

SetAgentMode sets AgentMode field to given value.

### HasAgentMode

`func (o *ListClouds200ResponseAllOfZonesInner) HasAgentMode() bool`

HasAgentMode returns a boolean if a field has been set.

### GetUserDataLinux

`func (o *ListClouds200ResponseAllOfZonesInner) GetUserDataLinux() string`

GetUserDataLinux returns the UserDataLinux field if non-nil, zero value otherwise.

### GetUserDataLinuxOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetUserDataLinuxOk() (*string, bool)`

GetUserDataLinuxOk returns a tuple with the UserDataLinux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataLinux

`func (o *ListClouds200ResponseAllOfZonesInner) SetUserDataLinux(v string)`

SetUserDataLinux sets UserDataLinux field to given value.

### HasUserDataLinux

`func (o *ListClouds200ResponseAllOfZonesInner) HasUserDataLinux() bool`

HasUserDataLinux returns a boolean if a field has been set.

### GetUserDataWindows

`func (o *ListClouds200ResponseAllOfZonesInner) GetUserDataWindows() string`

GetUserDataWindows returns the UserDataWindows field if non-nil, zero value otherwise.

### GetUserDataWindowsOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetUserDataWindowsOk() (*string, bool)`

GetUserDataWindowsOk returns a tuple with the UserDataWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataWindows

`func (o *ListClouds200ResponseAllOfZonesInner) SetUserDataWindows(v string)`

SetUserDataWindows sets UserDataWindows field to given value.

### HasUserDataWindows

`func (o *ListClouds200ResponseAllOfZonesInner) HasUserDataWindows() bool`

HasUserDataWindows returns a boolean if a field has been set.

### GetConsoleKeymap

`func (o *ListClouds200ResponseAllOfZonesInner) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *ListClouds200ResponseAllOfZonesInner) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *ListClouds200ResponseAllOfZonesInner) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### GetContainerMode

`func (o *ListClouds200ResponseAllOfZonesInner) GetContainerMode() string`

GetContainerMode returns the ContainerMode field if non-nil, zero value otherwise.

### GetContainerModeOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetContainerModeOk() (*string, bool)`

GetContainerModeOk returns a tuple with the ContainerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMode

`func (o *ListClouds200ResponseAllOfZonesInner) SetContainerMode(v string)`

SetContainerMode sets ContainerMode field to given value.

### HasContainerMode

`func (o *ListClouds200ResponseAllOfZonesInner) HasContainerMode() bool`

HasContainerMode returns a boolean if a field has been set.

### GetCostingMode

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostingMode() string`

GetCostingMode returns the CostingMode field if non-nil, zero value otherwise.

### GetCostingModeOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetCostingModeOk() (*string, bool)`

GetCostingModeOk returns a tuple with the CostingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingMode

`func (o *ListClouds200ResponseAllOfZonesInner) SetCostingMode(v string)`

SetCostingMode sets CostingMode field to given value.

### HasCostingMode

`func (o *ListClouds200ResponseAllOfZonesInner) HasCostingMode() bool`

HasCostingMode returns a boolean if a field has been set.

### GetServiceVersion

`func (o *ListClouds200ResponseAllOfZonesInner) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ListClouds200ResponseAllOfZonesInner) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ListClouds200ResponseAllOfZonesInner) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSecurityMode

`func (o *ListClouds200ResponseAllOfZonesInner) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *ListClouds200ResponseAllOfZonesInner) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *ListClouds200ResponseAllOfZonesInner) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *ListClouds200ResponseAllOfZonesInner) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *ListClouds200ResponseAllOfZonesInner) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *ListClouds200ResponseAllOfZonesInner) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetTimezone

`func (o *ListClouds200ResponseAllOfZonesInner) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ListClouds200ResponseAllOfZonesInner) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ListClouds200ResponseAllOfZonesInner) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetApiProxy

`func (o *ListClouds200ResponseAllOfZonesInner) GetApiProxy() string`

GetApiProxy returns the ApiProxy field if non-nil, zero value otherwise.

### GetApiProxyOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetApiProxyOk() (*string, bool)`

GetApiProxyOk returns a tuple with the ApiProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProxy

`func (o *ListClouds200ResponseAllOfZonesInner) SetApiProxy(v string)`

SetApiProxy sets ApiProxy field to given value.

### HasApiProxy

`func (o *ListClouds200ResponseAllOfZonesInner) HasApiProxy() bool`

HasApiProxy returns a boolean if a field has been set.

### GetProvisioningProxy

`func (o *ListClouds200ResponseAllOfZonesInner) GetProvisioningProxy() string`

GetProvisioningProxy returns the ProvisioningProxy field if non-nil, zero value otherwise.

### GetProvisioningProxyOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetProvisioningProxyOk() (*string, bool)`

GetProvisioningProxyOk returns a tuple with the ProvisioningProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProxy

`func (o *ListClouds200ResponseAllOfZonesInner) SetProvisioningProxy(v string)`

SetProvisioningProxy sets ProvisioningProxy field to given value.

### HasProvisioningProxy

`func (o *ListClouds200ResponseAllOfZonesInner) HasProvisioningProxy() bool`

HasProvisioningProxy returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *ListClouds200ResponseAllOfZonesInner) GetNetworkDomain() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetNetworkDomainOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *ListClouds200ResponseAllOfZonesInner) SetNetworkDomain(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *ListClouds200ResponseAllOfZonesInner) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetDomainName

`func (o *ListClouds200ResponseAllOfZonesInner) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ListClouds200ResponseAllOfZonesInner) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ListClouds200ResponseAllOfZonesInner) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetRegionCode

`func (o *ListClouds200ResponseAllOfZonesInner) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ListClouds200ResponseAllOfZonesInner) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ListClouds200ResponseAllOfZonesInner) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetAutoRecoverPowerState

`func (o *ListClouds200ResponseAllOfZonesInner) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *ListClouds200ResponseAllOfZonesInner) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *ListClouds200ResponseAllOfZonesInner) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetScalePriority

`func (o *ListClouds200ResponseAllOfZonesInner) GetScalePriority() int64`

GetScalePriority returns the ScalePriority field if non-nil, zero value otherwise.

### GetScalePriorityOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetScalePriorityOk() (*int64, bool)`

GetScalePriorityOk returns a tuple with the ScalePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalePriority

`func (o *ListClouds200ResponseAllOfZonesInner) SetScalePriority(v int64)`

SetScalePriority sets ScalePriority field to given value.

### HasScalePriority

`func (o *ListClouds200ResponseAllOfZonesInner) HasScalePriority() bool`

HasScalePriority returns a boolean if a field has been set.

### GetDefaultDatastoreSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultDatastoreSyncActive() bool`

GetDefaultDatastoreSyncActive returns the DefaultDatastoreSyncActive field if non-nil, zero value otherwise.

### GetDefaultDatastoreSyncActiveOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultDatastoreSyncActiveOk() (*bool, bool)`

GetDefaultDatastoreSyncActiveOk returns a tuple with the DefaultDatastoreSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDatastoreSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) SetDefaultDatastoreSyncActive(v bool)`

SetDefaultDatastoreSyncActive sets DefaultDatastoreSyncActive field to given value.

### HasDefaultDatastoreSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) HasDefaultDatastoreSyncActive() bool`

HasDefaultDatastoreSyncActive returns a boolean if a field has been set.

### GetDefaultNetworkSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultNetworkSyncActive() bool`

GetDefaultNetworkSyncActive returns the DefaultNetworkSyncActive field if non-nil, zero value otherwise.

### GetDefaultNetworkSyncActiveOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultNetworkSyncActiveOk() (*bool, bool)`

GetDefaultNetworkSyncActiveOk returns a tuple with the DefaultNetworkSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) SetDefaultNetworkSyncActive(v bool)`

SetDefaultNetworkSyncActive sets DefaultNetworkSyncActive field to given value.

### HasDefaultNetworkSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) HasDefaultNetworkSyncActive() bool`

HasDefaultNetworkSyncActive returns a boolean if a field has been set.

### GetDefaultFolderSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultFolderSyncActive() bool`

GetDefaultFolderSyncActive returns the DefaultFolderSyncActive field if non-nil, zero value otherwise.

### GetDefaultFolderSyncActiveOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultFolderSyncActiveOk() (*bool, bool)`

GetDefaultFolderSyncActiveOk returns a tuple with the DefaultFolderSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolderSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) SetDefaultFolderSyncActive(v bool)`

SetDefaultFolderSyncActive sets DefaultFolderSyncActive field to given value.

### HasDefaultFolderSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) HasDefaultFolderSyncActive() bool`

HasDefaultFolderSyncActive returns a boolean if a field has been set.

### GetDefaultSecurityGroupSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultSecurityGroupSyncActive() bool`

GetDefaultSecurityGroupSyncActive returns the DefaultSecurityGroupSyncActive field if non-nil, zero value otherwise.

### GetDefaultSecurityGroupSyncActiveOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultSecurityGroupSyncActiveOk() (*bool, bool)`

GetDefaultSecurityGroupSyncActiveOk returns a tuple with the DefaultSecurityGroupSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurityGroupSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) SetDefaultSecurityGroupSyncActive(v bool)`

SetDefaultSecurityGroupSyncActive sets DefaultSecurityGroupSyncActive field to given value.

### HasDefaultSecurityGroupSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) HasDefaultSecurityGroupSyncActive() bool`

HasDefaultSecurityGroupSyncActive returns a boolean if a field has been set.

### GetDefaultPoolSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultPoolSyncActive() bool`

GetDefaultPoolSyncActive returns the DefaultPoolSyncActive field if non-nil, zero value otherwise.

### GetDefaultPoolSyncActiveOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultPoolSyncActiveOk() (*bool, bool)`

GetDefaultPoolSyncActiveOk returns a tuple with the DefaultPoolSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPoolSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) SetDefaultPoolSyncActive(v bool)`

SetDefaultPoolSyncActive sets DefaultPoolSyncActive field to given value.

### HasDefaultPoolSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) HasDefaultPoolSyncActive() bool`

HasDefaultPoolSyncActive returns a boolean if a field has been set.

### GetDefaultPlanSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultPlanSyncActive() bool`

GetDefaultPlanSyncActive returns the DefaultPlanSyncActive field if non-nil, zero value otherwise.

### GetDefaultPlanSyncActiveOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetDefaultPlanSyncActiveOk() (*bool, bool)`

GetDefaultPlanSyncActiveOk returns a tuple with the DefaultPlanSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlanSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) SetDefaultPlanSyncActive(v bool)`

SetDefaultPlanSyncActive sets DefaultPlanSyncActive field to given value.

### HasDefaultPlanSyncActive

`func (o *ListClouds200ResponseAllOfZonesInner) HasDefaultPlanSyncActive() bool`

HasDefaultPlanSyncActive returns a boolean if a field has been set.

### GetConfig

`func (o *ListClouds200ResponseAllOfZonesInner) GetConfig() ListClouds200ResponseAllOfZonesInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetConfigOk() (*ListClouds200ResponseAllOfZonesInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListClouds200ResponseAllOfZonesInner) SetConfig(v ListClouds200ResponseAllOfZonesInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListClouds200ResponseAllOfZonesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *ListClouds200ResponseAllOfZonesInner) GetCredential() ListClouds200ResponseAllOfZonesInnerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListClouds200ResponseAllOfZonesInner) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListClouds200ResponseAllOfZonesInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetImagePath

`func (o *ListClouds200ResponseAllOfZonesInner) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *ListClouds200ResponseAllOfZonesInner) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *ListClouds200ResponseAllOfZonesInner) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetDarkImagePath

`func (o *ListClouds200ResponseAllOfZonesInner) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *ListClouds200ResponseAllOfZonesInner) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *ListClouds200ResponseAllOfZonesInner) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListClouds200ResponseAllOfZonesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListClouds200ResponseAllOfZonesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListClouds200ResponseAllOfZonesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListClouds200ResponseAllOfZonesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClouds200ResponseAllOfZonesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClouds200ResponseAllOfZonesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastSync

`func (o *ListClouds200ResponseAllOfZonesInner) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *ListClouds200ResponseAllOfZonesInner) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *ListClouds200ResponseAllOfZonesInner) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *ListClouds200ResponseAllOfZonesInner) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *ListClouds200ResponseAllOfZonesInner) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *ListClouds200ResponseAllOfZonesInner) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### GetNextRunDate

`func (o *ListClouds200ResponseAllOfZonesInner) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *ListClouds200ResponseAllOfZonesInner) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *ListClouds200ResponseAllOfZonesInner) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### GetGroups

`func (o *ListClouds200ResponseAllOfZonesInner) GetGroups() []ListClouds200ResponseAllOfZonesInnerGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetGroupsOk() (*[]ListClouds200ResponseAllOfZonesInnerGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ListClouds200ResponseAllOfZonesInner) SetGroups(v []ListClouds200ResponseAllOfZonesInnerGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ListClouds200ResponseAllOfZonesInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetSecurityServer

`func (o *ListClouds200ResponseAllOfZonesInner) GetSecurityServer() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetSecurityServerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *ListClouds200ResponseAllOfZonesInner) SetSecurityServer(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *ListClouds200ResponseAllOfZonesInner) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetNetworkServer

`func (o *ListClouds200ResponseAllOfZonesInner) GetNetworkServer() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetNetworkServerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *ListClouds200ResponseAllOfZonesInner) SetNetworkServer(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *ListClouds200ResponseAllOfZonesInner) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetStats

`func (o *ListClouds200ResponseAllOfZonesInner) GetStats() ListClouds200ResponseAllOfZonesInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetStatsOk() (*ListClouds200ResponseAllOfZonesInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListClouds200ResponseAllOfZonesInner) SetStats(v ListClouds200ResponseAllOfZonesInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListClouds200ResponseAllOfZonesInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *ListClouds200ResponseAllOfZonesInner) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *ListClouds200ResponseAllOfZonesInner) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *ListClouds200ResponseAllOfZonesInner) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *ListClouds200ResponseAllOfZonesInner) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


