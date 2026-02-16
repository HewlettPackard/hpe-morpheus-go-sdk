# GetClouds200ResponseZone

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
**Owner** | Pointer to [**AddClouds200ResponseAllOfZoneOwner**](AddClouds200ResponseAllOfZoneOwner.md) |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**AddClouds200ResponseAllOfZoneAccount**](AddClouds200ResponseAllOfZoneAccount.md) |  | [optional] 
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
**ZoneType** | Pointer to [**AddClouds200ResponseAllOfZoneZoneType**](AddClouds200ResponseAllOfZoneZoneType.md) |  | [optional] 
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
**NetworkDomain** | Pointer to [**AddClouds200ResponseAllOfZoneNetworkDomain**](AddClouds200ResponseAllOfZoneNetworkDomain.md) |  | [optional] 
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
**Config** | Pointer to [**AddClouds200ResponseAllOfZoneConfig**](AddClouds200ResponseAllOfZoneConfig.md) |  | [optional] 
**Credential** | Pointer to [**AddClouds200ResponseAllOfZoneCredential**](AddClouds200ResponseAllOfZoneCredential.md) |  | [optional] 
**ImagePath** | Pointer to **NullableString** | Logo image URL | [optional] 
**DarkImagePath** | Pointer to **NullableString** | Dark logo image URL | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastSync** | Pointer to **NullableTime** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableInt64** |  | [optional] 
**NextRunDate** | Pointer to **NullableTime** |  | [optional] 
**Groups** | Pointer to [**[]AddClouds200ResponseAllOfZoneGroupsInner**](AddClouds200ResponseAllOfZoneGroupsInner.md) |  | [optional] 
**SecurityServer** | Pointer to [**AddClouds200ResponseAllOfZoneSecurityServer**](AddClouds200ResponseAllOfZoneSecurityServer.md) |  | [optional] 
**NetworkServer** | Pointer to [**AddClouds200ResponseAllOfZoneNetworkServer**](AddClouds200ResponseAllOfZoneNetworkServer.md) |  | [optional] 
**Stats** | Pointer to [**AddClouds200ResponseAllOfZoneStats**](AddClouds200ResponseAllOfZoneStats.md) |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetClouds200ResponseZone

`func NewGetClouds200ResponseZone() *GetClouds200ResponseZone`

NewGetClouds200ResponseZone instantiates a new GetClouds200ResponseZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClouds200ResponseZoneWithDefaults

`func NewGetClouds200ResponseZoneWithDefaults() *GetClouds200ResponseZone`

NewGetClouds200ResponseZoneWithDefaults instantiates a new GetClouds200ResponseZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClouds200ResponseZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClouds200ResponseZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClouds200ResponseZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClouds200ResponseZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetClouds200ResponseZone) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetClouds200ResponseZone) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetClouds200ResponseZone) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetClouds200ResponseZone) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *GetClouds200ResponseZone) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetClouds200ResponseZone) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetClouds200ResponseZone) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetClouds200ResponseZone) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetClouds200ResponseZone) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetClouds200ResponseZone) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetName

`func (o *GetClouds200ResponseZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClouds200ResponseZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClouds200ResponseZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClouds200ResponseZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetClouds200ResponseZone) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetClouds200ResponseZone) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetClouds200ResponseZone) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetClouds200ResponseZone) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabels

`func (o *GetClouds200ResponseZone) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetClouds200ResponseZone) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetClouds200ResponseZone) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetClouds200ResponseZone) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *GetClouds200ResponseZone) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetClouds200ResponseZone) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetClouds200ResponseZone) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetClouds200ResponseZone) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *GetClouds200ResponseZone) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *GetClouds200ResponseZone) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetOwner

`func (o *GetClouds200ResponseZone) GetOwner() AddClouds200ResponseAllOfZoneOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetClouds200ResponseZone) GetOwnerOk() (*AddClouds200ResponseAllOfZoneOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetClouds200ResponseZone) SetOwner(v AddClouds200ResponseAllOfZoneOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetClouds200ResponseZone) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccountId

`func (o *GetClouds200ResponseZone) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetClouds200ResponseZone) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetClouds200ResponseZone) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetClouds200ResponseZone) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *GetClouds200ResponseZone) GetAccount() AddClouds200ResponseAllOfZoneAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetClouds200ResponseZone) GetAccountOk() (*AddClouds200ResponseAllOfZoneAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetClouds200ResponseZone) SetAccount(v AddClouds200ResponseAllOfZoneAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetClouds200ResponseZone) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetVisibility

`func (o *GetClouds200ResponseZone) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetClouds200ResponseZone) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetClouds200ResponseZone) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetClouds200ResponseZone) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnabled

`func (o *GetClouds200ResponseZone) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetClouds200ResponseZone) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetClouds200ResponseZone) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetClouds200ResponseZone) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *GetClouds200ResponseZone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClouds200ResponseZone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClouds200ResponseZone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetClouds200ResponseZone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetClouds200ResponseZone) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetClouds200ResponseZone) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetClouds200ResponseZone) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetClouds200ResponseZone) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetClouds200ResponseZone) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetClouds200ResponseZone) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *GetClouds200ResponseZone) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetClouds200ResponseZone) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetClouds200ResponseZone) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetClouds200ResponseZone) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetClouds200ResponseZone) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetClouds200ResponseZone) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetCostStatus

`func (o *GetClouds200ResponseZone) GetCostStatus() string`

GetCostStatus returns the CostStatus field if non-nil, zero value otherwise.

### GetCostStatusOk

`func (o *GetClouds200ResponseZone) GetCostStatusOk() (*string, bool)`

GetCostStatusOk returns a tuple with the CostStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatus

`func (o *GetClouds200ResponseZone) SetCostStatus(v string)`

SetCostStatus sets CostStatus field to given value.

### HasCostStatus

`func (o *GetClouds200ResponseZone) HasCostStatus() bool`

HasCostStatus returns a boolean if a field has been set.

### SetCostStatusNil

`func (o *GetClouds200ResponseZone) SetCostStatusNil(b bool)`

 SetCostStatusNil sets the value for CostStatus to be an explicit nil

### UnsetCostStatus
`func (o *GetClouds200ResponseZone) UnsetCostStatus()`

UnsetCostStatus ensures that no value is present for CostStatus, not even an explicit nil
### GetCostStatusMessage

`func (o *GetClouds200ResponseZone) GetCostStatusMessage() string`

GetCostStatusMessage returns the CostStatusMessage field if non-nil, zero value otherwise.

### GetCostStatusMessageOk

`func (o *GetClouds200ResponseZone) GetCostStatusMessageOk() (*string, bool)`

GetCostStatusMessageOk returns a tuple with the CostStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatusMessage

`func (o *GetClouds200ResponseZone) SetCostStatusMessage(v string)`

SetCostStatusMessage sets CostStatusMessage field to given value.

### HasCostStatusMessage

`func (o *GetClouds200ResponseZone) HasCostStatusMessage() bool`

HasCostStatusMessage returns a boolean if a field has been set.

### SetCostStatusMessageNil

`func (o *GetClouds200ResponseZone) SetCostStatusMessageNil(b bool)`

 SetCostStatusMessageNil sets the value for CostStatusMessage to be an explicit nil

### UnsetCostStatusMessage
`func (o *GetClouds200ResponseZone) UnsetCostStatusMessage()`

UnsetCostStatusMessage ensures that no value is present for CostStatusMessage, not even an explicit nil
### GetCostStatusDate

`func (o *GetClouds200ResponseZone) GetCostStatusDate() time.Time`

GetCostStatusDate returns the CostStatusDate field if non-nil, zero value otherwise.

### GetCostStatusDateOk

`func (o *GetClouds200ResponseZone) GetCostStatusDateOk() (*time.Time, bool)`

GetCostStatusDateOk returns a tuple with the CostStatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatusDate

`func (o *GetClouds200ResponseZone) SetCostStatusDate(v time.Time)`

SetCostStatusDate sets CostStatusDate field to given value.

### HasCostStatusDate

`func (o *GetClouds200ResponseZone) HasCostStatusDate() bool`

HasCostStatusDate returns a boolean if a field has been set.

### SetCostStatusDateNil

`func (o *GetClouds200ResponseZone) SetCostStatusDateNil(b bool)`

 SetCostStatusDateNil sets the value for CostStatusDate to be an explicit nil

### UnsetCostStatusDate
`func (o *GetClouds200ResponseZone) UnsetCostStatusDate()`

UnsetCostStatusDate ensures that no value is present for CostStatusDate, not even an explicit nil
### GetCostLastSyncDuration

`func (o *GetClouds200ResponseZone) GetCostLastSyncDuration() int64`

GetCostLastSyncDuration returns the CostLastSyncDuration field if non-nil, zero value otherwise.

### GetCostLastSyncDurationOk

`func (o *GetClouds200ResponseZone) GetCostLastSyncDurationOk() (*int64, bool)`

GetCostLastSyncDurationOk returns a tuple with the CostLastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLastSyncDuration

`func (o *GetClouds200ResponseZone) SetCostLastSyncDuration(v int64)`

SetCostLastSyncDuration sets CostLastSyncDuration field to given value.

### HasCostLastSyncDuration

`func (o *GetClouds200ResponseZone) HasCostLastSyncDuration() bool`

HasCostLastSyncDuration returns a boolean if a field has been set.

### SetCostLastSyncDurationNil

`func (o *GetClouds200ResponseZone) SetCostLastSyncDurationNil(b bool)`

 SetCostLastSyncDurationNil sets the value for CostLastSyncDuration to be an explicit nil

### UnsetCostLastSyncDuration
`func (o *GetClouds200ResponseZone) UnsetCostLastSyncDuration()`

UnsetCostLastSyncDuration ensures that no value is present for CostLastSyncDuration, not even an explicit nil
### GetCostLastSync

`func (o *GetClouds200ResponseZone) GetCostLastSync() time.Time`

GetCostLastSync returns the CostLastSync field if non-nil, zero value otherwise.

### GetCostLastSyncOk

`func (o *GetClouds200ResponseZone) GetCostLastSyncOk() (*time.Time, bool)`

GetCostLastSyncOk returns a tuple with the CostLastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLastSync

`func (o *GetClouds200ResponseZone) SetCostLastSync(v time.Time)`

SetCostLastSync sets CostLastSync field to given value.

### HasCostLastSync

`func (o *GetClouds200ResponseZone) HasCostLastSync() bool`

HasCostLastSync returns a boolean if a field has been set.

### SetCostLastSyncNil

`func (o *GetClouds200ResponseZone) SetCostLastSyncNil(b bool)`

 SetCostLastSyncNil sets the value for CostLastSync to be an explicit nil

### UnsetCostLastSync
`func (o *GetClouds200ResponseZone) UnsetCostLastSync()`

UnsetCostLastSync ensures that no value is present for CostLastSync, not even an explicit nil
### GetZoneType

`func (o *GetClouds200ResponseZone) GetZoneType() AddClouds200ResponseAllOfZoneZoneType`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *GetClouds200ResponseZone) GetZoneTypeOk() (*AddClouds200ResponseAllOfZoneZoneType, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *GetClouds200ResponseZone) SetZoneType(v AddClouds200ResponseAllOfZoneZoneType)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *GetClouds200ResponseZone) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetZoneTypeId

`func (o *GetClouds200ResponseZone) GetZoneTypeId() int64`

GetZoneTypeId returns the ZoneTypeId field if non-nil, zero value otherwise.

### GetZoneTypeIdOk

`func (o *GetClouds200ResponseZone) GetZoneTypeIdOk() (*int64, bool)`

GetZoneTypeIdOk returns a tuple with the ZoneTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTypeId

`func (o *GetClouds200ResponseZone) SetZoneTypeId(v int64)`

SetZoneTypeId sets ZoneTypeId field to given value.

### HasZoneTypeId

`func (o *GetClouds200ResponseZone) HasZoneTypeId() bool`

HasZoneTypeId returns a boolean if a field has been set.

### GetGuidanceMode

`func (o *GetClouds200ResponseZone) GetGuidanceMode() string`

GetGuidanceMode returns the GuidanceMode field if non-nil, zero value otherwise.

### GetGuidanceModeOk

`func (o *GetClouds200ResponseZone) GetGuidanceModeOk() (*string, bool)`

GetGuidanceModeOk returns a tuple with the GuidanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidanceMode

`func (o *GetClouds200ResponseZone) SetGuidanceMode(v string)`

SetGuidanceMode sets GuidanceMode field to given value.

### HasGuidanceMode

`func (o *GetClouds200ResponseZone) HasGuidanceMode() bool`

HasGuidanceMode returns a boolean if a field has been set.

### SetGuidanceModeNil

`func (o *GetClouds200ResponseZone) SetGuidanceModeNil(b bool)`

 SetGuidanceModeNil sets the value for GuidanceMode to be an explicit nil

### UnsetGuidanceMode
`func (o *GetClouds200ResponseZone) UnsetGuidanceMode()`

UnsetGuidanceMode ensures that no value is present for GuidanceMode, not even an explicit nil
### GetStorageMode

`func (o *GetClouds200ResponseZone) GetStorageMode() string`

GetStorageMode returns the StorageMode field if non-nil, zero value otherwise.

### GetStorageModeOk

`func (o *GetClouds200ResponseZone) GetStorageModeOk() (*string, bool)`

GetStorageModeOk returns a tuple with the StorageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageMode

`func (o *GetClouds200ResponseZone) SetStorageMode(v string)`

SetStorageMode sets StorageMode field to given value.

### HasStorageMode

`func (o *GetClouds200ResponseZone) HasStorageMode() bool`

HasStorageMode returns a boolean if a field has been set.

### GetAgentMode

`func (o *GetClouds200ResponseZone) GetAgentMode() string`

GetAgentMode returns the AgentMode field if non-nil, zero value otherwise.

### GetAgentModeOk

`func (o *GetClouds200ResponseZone) GetAgentModeOk() (*string, bool)`

GetAgentModeOk returns a tuple with the AgentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMode

`func (o *GetClouds200ResponseZone) SetAgentMode(v string)`

SetAgentMode sets AgentMode field to given value.

### HasAgentMode

`func (o *GetClouds200ResponseZone) HasAgentMode() bool`

HasAgentMode returns a boolean if a field has been set.

### GetUserDataLinux

`func (o *GetClouds200ResponseZone) GetUserDataLinux() string`

GetUserDataLinux returns the UserDataLinux field if non-nil, zero value otherwise.

### GetUserDataLinuxOk

`func (o *GetClouds200ResponseZone) GetUserDataLinuxOk() (*string, bool)`

GetUserDataLinuxOk returns a tuple with the UserDataLinux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataLinux

`func (o *GetClouds200ResponseZone) SetUserDataLinux(v string)`

SetUserDataLinux sets UserDataLinux field to given value.

### HasUserDataLinux

`func (o *GetClouds200ResponseZone) HasUserDataLinux() bool`

HasUserDataLinux returns a boolean if a field has been set.

### SetUserDataLinuxNil

`func (o *GetClouds200ResponseZone) SetUserDataLinuxNil(b bool)`

 SetUserDataLinuxNil sets the value for UserDataLinux to be an explicit nil

### UnsetUserDataLinux
`func (o *GetClouds200ResponseZone) UnsetUserDataLinux()`

UnsetUserDataLinux ensures that no value is present for UserDataLinux, not even an explicit nil
### GetUserDataWindows

`func (o *GetClouds200ResponseZone) GetUserDataWindows() string`

GetUserDataWindows returns the UserDataWindows field if non-nil, zero value otherwise.

### GetUserDataWindowsOk

`func (o *GetClouds200ResponseZone) GetUserDataWindowsOk() (*string, bool)`

GetUserDataWindowsOk returns a tuple with the UserDataWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataWindows

`func (o *GetClouds200ResponseZone) SetUserDataWindows(v string)`

SetUserDataWindows sets UserDataWindows field to given value.

### HasUserDataWindows

`func (o *GetClouds200ResponseZone) HasUserDataWindows() bool`

HasUserDataWindows returns a boolean if a field has been set.

### SetUserDataWindowsNil

`func (o *GetClouds200ResponseZone) SetUserDataWindowsNil(b bool)`

 SetUserDataWindowsNil sets the value for UserDataWindows to be an explicit nil

### UnsetUserDataWindows
`func (o *GetClouds200ResponseZone) UnsetUserDataWindows()`

UnsetUserDataWindows ensures that no value is present for UserDataWindows, not even an explicit nil
### GetConsoleKeymap

`func (o *GetClouds200ResponseZone) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *GetClouds200ResponseZone) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *GetClouds200ResponseZone) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *GetClouds200ResponseZone) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### SetConsoleKeymapNil

`func (o *GetClouds200ResponseZone) SetConsoleKeymapNil(b bool)`

 SetConsoleKeymapNil sets the value for ConsoleKeymap to be an explicit nil

### UnsetConsoleKeymap
`func (o *GetClouds200ResponseZone) UnsetConsoleKeymap()`

UnsetConsoleKeymap ensures that no value is present for ConsoleKeymap, not even an explicit nil
### GetContainerMode

`func (o *GetClouds200ResponseZone) GetContainerMode() string`

GetContainerMode returns the ContainerMode field if non-nil, zero value otherwise.

### GetContainerModeOk

`func (o *GetClouds200ResponseZone) GetContainerModeOk() (*string, bool)`

GetContainerModeOk returns a tuple with the ContainerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMode

`func (o *GetClouds200ResponseZone) SetContainerMode(v string)`

SetContainerMode sets ContainerMode field to given value.

### HasContainerMode

`func (o *GetClouds200ResponseZone) HasContainerMode() bool`

HasContainerMode returns a boolean if a field has been set.

### GetCostingMode

`func (o *GetClouds200ResponseZone) GetCostingMode() string`

GetCostingMode returns the CostingMode field if non-nil, zero value otherwise.

### GetCostingModeOk

`func (o *GetClouds200ResponseZone) GetCostingModeOk() (*string, bool)`

GetCostingModeOk returns a tuple with the CostingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingMode

`func (o *GetClouds200ResponseZone) SetCostingMode(v string)`

SetCostingMode sets CostingMode field to given value.

### HasCostingMode

`func (o *GetClouds200ResponseZone) HasCostingMode() bool`

HasCostingMode returns a boolean if a field has been set.

### SetCostingModeNil

`func (o *GetClouds200ResponseZone) SetCostingModeNil(b bool)`

 SetCostingModeNil sets the value for CostingMode to be an explicit nil

### UnsetCostingMode
`func (o *GetClouds200ResponseZone) UnsetCostingMode()`

UnsetCostingMode ensures that no value is present for CostingMode, not even an explicit nil
### GetServiceVersion

`func (o *GetClouds200ResponseZone) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *GetClouds200ResponseZone) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *GetClouds200ResponseZone) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *GetClouds200ResponseZone) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### SetServiceVersionNil

`func (o *GetClouds200ResponseZone) SetServiceVersionNil(b bool)`

 SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil

### UnsetServiceVersion
`func (o *GetClouds200ResponseZone) UnsetServiceVersion()`

UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
### GetSecurityMode

`func (o *GetClouds200ResponseZone) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *GetClouds200ResponseZone) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *GetClouds200ResponseZone) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *GetClouds200ResponseZone) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *GetClouds200ResponseZone) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *GetClouds200ResponseZone) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *GetClouds200ResponseZone) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *GetClouds200ResponseZone) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetTimezone

`func (o *GetClouds200ResponseZone) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetClouds200ResponseZone) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetClouds200ResponseZone) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetClouds200ResponseZone) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *GetClouds200ResponseZone) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *GetClouds200ResponseZone) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetApiProxy

`func (o *GetClouds200ResponseZone) GetApiProxy() string`

GetApiProxy returns the ApiProxy field if non-nil, zero value otherwise.

### GetApiProxyOk

`func (o *GetClouds200ResponseZone) GetApiProxyOk() (*string, bool)`

GetApiProxyOk returns a tuple with the ApiProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProxy

`func (o *GetClouds200ResponseZone) SetApiProxy(v string)`

SetApiProxy sets ApiProxy field to given value.

### HasApiProxy

`func (o *GetClouds200ResponseZone) HasApiProxy() bool`

HasApiProxy returns a boolean if a field has been set.

### SetApiProxyNil

`func (o *GetClouds200ResponseZone) SetApiProxyNil(b bool)`

 SetApiProxyNil sets the value for ApiProxy to be an explicit nil

### UnsetApiProxy
`func (o *GetClouds200ResponseZone) UnsetApiProxy()`

UnsetApiProxy ensures that no value is present for ApiProxy, not even an explicit nil
### GetProvisioningProxy

`func (o *GetClouds200ResponseZone) GetProvisioningProxy() string`

GetProvisioningProxy returns the ProvisioningProxy field if non-nil, zero value otherwise.

### GetProvisioningProxyOk

`func (o *GetClouds200ResponseZone) GetProvisioningProxyOk() (*string, bool)`

GetProvisioningProxyOk returns a tuple with the ProvisioningProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProxy

`func (o *GetClouds200ResponseZone) SetProvisioningProxy(v string)`

SetProvisioningProxy sets ProvisioningProxy field to given value.

### HasProvisioningProxy

`func (o *GetClouds200ResponseZone) HasProvisioningProxy() bool`

HasProvisioningProxy returns a boolean if a field has been set.

### SetProvisioningProxyNil

`func (o *GetClouds200ResponseZone) SetProvisioningProxyNil(b bool)`

 SetProvisioningProxyNil sets the value for ProvisioningProxy to be an explicit nil

### UnsetProvisioningProxy
`func (o *GetClouds200ResponseZone) UnsetProvisioningProxy()`

UnsetProvisioningProxy ensures that no value is present for ProvisioningProxy, not even an explicit nil
### GetNetworkDomain

`func (o *GetClouds200ResponseZone) GetNetworkDomain() AddClouds200ResponseAllOfZoneNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *GetClouds200ResponseZone) GetNetworkDomainOk() (*AddClouds200ResponseAllOfZoneNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *GetClouds200ResponseZone) SetNetworkDomain(v AddClouds200ResponseAllOfZoneNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *GetClouds200ResponseZone) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetDomainName

`func (o *GetClouds200ResponseZone) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GetClouds200ResponseZone) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GetClouds200ResponseZone) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GetClouds200ResponseZone) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetRegionCode

`func (o *GetClouds200ResponseZone) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *GetClouds200ResponseZone) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *GetClouds200ResponseZone) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *GetClouds200ResponseZone) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *GetClouds200ResponseZone) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *GetClouds200ResponseZone) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetAutoRecoverPowerState

`func (o *GetClouds200ResponseZone) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *GetClouds200ResponseZone) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *GetClouds200ResponseZone) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *GetClouds200ResponseZone) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetScalePriority

`func (o *GetClouds200ResponseZone) GetScalePriority() int64`

GetScalePriority returns the ScalePriority field if non-nil, zero value otherwise.

### GetScalePriorityOk

`func (o *GetClouds200ResponseZone) GetScalePriorityOk() (*int64, bool)`

GetScalePriorityOk returns a tuple with the ScalePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalePriority

`func (o *GetClouds200ResponseZone) SetScalePriority(v int64)`

SetScalePriority sets ScalePriority field to given value.

### HasScalePriority

`func (o *GetClouds200ResponseZone) HasScalePriority() bool`

HasScalePriority returns a boolean if a field has been set.

### GetDefaultDatastoreSyncActive

`func (o *GetClouds200ResponseZone) GetDefaultDatastoreSyncActive() bool`

GetDefaultDatastoreSyncActive returns the DefaultDatastoreSyncActive field if non-nil, zero value otherwise.

### GetDefaultDatastoreSyncActiveOk

`func (o *GetClouds200ResponseZone) GetDefaultDatastoreSyncActiveOk() (*bool, bool)`

GetDefaultDatastoreSyncActiveOk returns a tuple with the DefaultDatastoreSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDatastoreSyncActive

`func (o *GetClouds200ResponseZone) SetDefaultDatastoreSyncActive(v bool)`

SetDefaultDatastoreSyncActive sets DefaultDatastoreSyncActive field to given value.

### HasDefaultDatastoreSyncActive

`func (o *GetClouds200ResponseZone) HasDefaultDatastoreSyncActive() bool`

HasDefaultDatastoreSyncActive returns a boolean if a field has been set.

### GetDefaultNetworkSyncActive

`func (o *GetClouds200ResponseZone) GetDefaultNetworkSyncActive() bool`

GetDefaultNetworkSyncActive returns the DefaultNetworkSyncActive field if non-nil, zero value otherwise.

### GetDefaultNetworkSyncActiveOk

`func (o *GetClouds200ResponseZone) GetDefaultNetworkSyncActiveOk() (*bool, bool)`

GetDefaultNetworkSyncActiveOk returns a tuple with the DefaultNetworkSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkSyncActive

`func (o *GetClouds200ResponseZone) SetDefaultNetworkSyncActive(v bool)`

SetDefaultNetworkSyncActive sets DefaultNetworkSyncActive field to given value.

### HasDefaultNetworkSyncActive

`func (o *GetClouds200ResponseZone) HasDefaultNetworkSyncActive() bool`

HasDefaultNetworkSyncActive returns a boolean if a field has been set.

### GetDefaultFolderSyncActive

`func (o *GetClouds200ResponseZone) GetDefaultFolderSyncActive() bool`

GetDefaultFolderSyncActive returns the DefaultFolderSyncActive field if non-nil, zero value otherwise.

### GetDefaultFolderSyncActiveOk

`func (o *GetClouds200ResponseZone) GetDefaultFolderSyncActiveOk() (*bool, bool)`

GetDefaultFolderSyncActiveOk returns a tuple with the DefaultFolderSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolderSyncActive

`func (o *GetClouds200ResponseZone) SetDefaultFolderSyncActive(v bool)`

SetDefaultFolderSyncActive sets DefaultFolderSyncActive field to given value.

### HasDefaultFolderSyncActive

`func (o *GetClouds200ResponseZone) HasDefaultFolderSyncActive() bool`

HasDefaultFolderSyncActive returns a boolean if a field has been set.

### GetDefaultSecurityGroupSyncActive

`func (o *GetClouds200ResponseZone) GetDefaultSecurityGroupSyncActive() bool`

GetDefaultSecurityGroupSyncActive returns the DefaultSecurityGroupSyncActive field if non-nil, zero value otherwise.

### GetDefaultSecurityGroupSyncActiveOk

`func (o *GetClouds200ResponseZone) GetDefaultSecurityGroupSyncActiveOk() (*bool, bool)`

GetDefaultSecurityGroupSyncActiveOk returns a tuple with the DefaultSecurityGroupSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurityGroupSyncActive

`func (o *GetClouds200ResponseZone) SetDefaultSecurityGroupSyncActive(v bool)`

SetDefaultSecurityGroupSyncActive sets DefaultSecurityGroupSyncActive field to given value.

### HasDefaultSecurityGroupSyncActive

`func (o *GetClouds200ResponseZone) HasDefaultSecurityGroupSyncActive() bool`

HasDefaultSecurityGroupSyncActive returns a boolean if a field has been set.

### GetDefaultPoolSyncActive

`func (o *GetClouds200ResponseZone) GetDefaultPoolSyncActive() bool`

GetDefaultPoolSyncActive returns the DefaultPoolSyncActive field if non-nil, zero value otherwise.

### GetDefaultPoolSyncActiveOk

`func (o *GetClouds200ResponseZone) GetDefaultPoolSyncActiveOk() (*bool, bool)`

GetDefaultPoolSyncActiveOk returns a tuple with the DefaultPoolSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPoolSyncActive

`func (o *GetClouds200ResponseZone) SetDefaultPoolSyncActive(v bool)`

SetDefaultPoolSyncActive sets DefaultPoolSyncActive field to given value.

### HasDefaultPoolSyncActive

`func (o *GetClouds200ResponseZone) HasDefaultPoolSyncActive() bool`

HasDefaultPoolSyncActive returns a boolean if a field has been set.

### GetDefaultPlanSyncActive

`func (o *GetClouds200ResponseZone) GetDefaultPlanSyncActive() bool`

GetDefaultPlanSyncActive returns the DefaultPlanSyncActive field if non-nil, zero value otherwise.

### GetDefaultPlanSyncActiveOk

`func (o *GetClouds200ResponseZone) GetDefaultPlanSyncActiveOk() (*bool, bool)`

GetDefaultPlanSyncActiveOk returns a tuple with the DefaultPlanSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlanSyncActive

`func (o *GetClouds200ResponseZone) SetDefaultPlanSyncActive(v bool)`

SetDefaultPlanSyncActive sets DefaultPlanSyncActive field to given value.

### HasDefaultPlanSyncActive

`func (o *GetClouds200ResponseZone) HasDefaultPlanSyncActive() bool`

HasDefaultPlanSyncActive returns a boolean if a field has been set.

### GetConfig

`func (o *GetClouds200ResponseZone) GetConfig() AddClouds200ResponseAllOfZoneConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetClouds200ResponseZone) GetConfigOk() (*AddClouds200ResponseAllOfZoneConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetClouds200ResponseZone) SetConfig(v AddClouds200ResponseAllOfZoneConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetClouds200ResponseZone) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *GetClouds200ResponseZone) GetCredential() AddClouds200ResponseAllOfZoneCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetClouds200ResponseZone) GetCredentialOk() (*AddClouds200ResponseAllOfZoneCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetClouds200ResponseZone) SetCredential(v AddClouds200ResponseAllOfZoneCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetClouds200ResponseZone) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetImagePath

`func (o *GetClouds200ResponseZone) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *GetClouds200ResponseZone) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *GetClouds200ResponseZone) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *GetClouds200ResponseZone) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### SetImagePathNil

`func (o *GetClouds200ResponseZone) SetImagePathNil(b bool)`

 SetImagePathNil sets the value for ImagePath to be an explicit nil

### UnsetImagePath
`func (o *GetClouds200ResponseZone) UnsetImagePath()`

UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
### GetDarkImagePath

`func (o *GetClouds200ResponseZone) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *GetClouds200ResponseZone) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *GetClouds200ResponseZone) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *GetClouds200ResponseZone) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### SetDarkImagePathNil

`func (o *GetClouds200ResponseZone) SetDarkImagePathNil(b bool)`

 SetDarkImagePathNil sets the value for DarkImagePath to be an explicit nil

### UnsetDarkImagePath
`func (o *GetClouds200ResponseZone) UnsetDarkImagePath()`

UnsetDarkImagePath ensures that no value is present for DarkImagePath, not even an explicit nil
### GetDateCreated

`func (o *GetClouds200ResponseZone) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetClouds200ResponseZone) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetClouds200ResponseZone) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetClouds200ResponseZone) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetClouds200ResponseZone) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetClouds200ResponseZone) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetClouds200ResponseZone) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetClouds200ResponseZone) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastSync

`func (o *GetClouds200ResponseZone) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *GetClouds200ResponseZone) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *GetClouds200ResponseZone) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *GetClouds200ResponseZone) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *GetClouds200ResponseZone) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *GetClouds200ResponseZone) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *GetClouds200ResponseZone) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *GetClouds200ResponseZone) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *GetClouds200ResponseZone) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *GetClouds200ResponseZone) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *GetClouds200ResponseZone) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *GetClouds200ResponseZone) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetNextRunDate

`func (o *GetClouds200ResponseZone) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *GetClouds200ResponseZone) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *GetClouds200ResponseZone) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *GetClouds200ResponseZone) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *GetClouds200ResponseZone) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *GetClouds200ResponseZone) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetGroups

`func (o *GetClouds200ResponseZone) GetGroups() []AddClouds200ResponseAllOfZoneGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetClouds200ResponseZone) GetGroupsOk() (*[]AddClouds200ResponseAllOfZoneGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetClouds200ResponseZone) SetGroups(v []AddClouds200ResponseAllOfZoneGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GetClouds200ResponseZone) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetSecurityServer

`func (o *GetClouds200ResponseZone) GetSecurityServer() AddClouds200ResponseAllOfZoneSecurityServer`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *GetClouds200ResponseZone) GetSecurityServerOk() (*AddClouds200ResponseAllOfZoneSecurityServer, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *GetClouds200ResponseZone) SetSecurityServer(v AddClouds200ResponseAllOfZoneSecurityServer)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *GetClouds200ResponseZone) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### GetNetworkServer

`func (o *GetClouds200ResponseZone) GetNetworkServer() AddClouds200ResponseAllOfZoneNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetClouds200ResponseZone) GetNetworkServerOk() (*AddClouds200ResponseAllOfZoneNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetClouds200ResponseZone) SetNetworkServer(v AddClouds200ResponseAllOfZoneNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetClouds200ResponseZone) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetStats

`func (o *GetClouds200ResponseZone) GetStats() AddClouds200ResponseAllOfZoneStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetClouds200ResponseZone) GetStatsOk() (*AddClouds200ResponseAllOfZoneStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetClouds200ResponseZone) SetStats(v AddClouds200ResponseAllOfZoneStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetClouds200ResponseZone) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *GetClouds200ResponseZone) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *GetClouds200ResponseZone) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *GetClouds200ResponseZone) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *GetClouds200ResponseZone) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


