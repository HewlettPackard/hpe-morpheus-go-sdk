# UpdateGroups200ResponseGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**ListGroups200ResponseAllOfGroupsInnerConfig**](ListGroups200ResponseAllOfGroupsInnerConfig.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Zones** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Stats** | Pointer to [**ListGroups200ResponseAllOfGroupsInnerStats**](ListGroups200ResponseAllOfGroupsInnerStats.md) |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateGroups200ResponseGroup

`func NewUpdateGroups200ResponseGroup() *UpdateGroups200ResponseGroup`

NewUpdateGroups200ResponseGroup instantiates a new UpdateGroups200ResponseGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroups200ResponseGroupWithDefaults

`func NewUpdateGroups200ResponseGroupWithDefaults() *UpdateGroups200ResponseGroup`

NewUpdateGroups200ResponseGroupWithDefaults instantiates a new UpdateGroups200ResponseGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateGroups200ResponseGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateGroups200ResponseGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateGroups200ResponseGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateGroups200ResponseGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *UpdateGroups200ResponseGroup) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpdateGroups200ResponseGroup) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpdateGroups200ResponseGroup) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpdateGroups200ResponseGroup) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *UpdateGroups200ResponseGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGroups200ResponseGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGroups200ResponseGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateGroups200ResponseGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateGroups200ResponseGroup) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateGroups200ResponseGroup) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateGroups200ResponseGroup) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateGroups200ResponseGroup) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateGroups200ResponseGroup) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateGroups200ResponseGroup) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateGroups200ResponseGroup) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateGroups200ResponseGroup) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateGroups200ResponseGroup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateGroups200ResponseGroup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateGroups200ResponseGroup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateGroups200ResponseGroup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateGroups200ResponseGroup) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateGroups200ResponseGroup) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateGroups200ResponseGroup) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateGroups200ResponseGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActive

`func (o *UpdateGroups200ResponseGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateGroups200ResponseGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateGroups200ResponseGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateGroups200ResponseGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateGroups200ResponseGroup) GetConfig() ListGroups200ResponseAllOfGroupsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateGroups200ResponseGroup) GetConfigOk() (*ListGroups200ResponseAllOfGroupsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateGroups200ResponseGroup) SetConfig(v ListGroups200ResponseAllOfGroupsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateGroups200ResponseGroup) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateGroups200ResponseGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateGroups200ResponseGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateGroups200ResponseGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateGroups200ResponseGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateGroups200ResponseGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateGroups200ResponseGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateGroups200ResponseGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateGroups200ResponseGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetZones

`func (o *UpdateGroups200ResponseGroup) GetZones() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *UpdateGroups200ResponseGroup) GetZonesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *UpdateGroups200ResponseGroup) SetZones(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *UpdateGroups200ResponseGroup) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetStats

`func (o *UpdateGroups200ResponseGroup) GetStats() ListGroups200ResponseAllOfGroupsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *UpdateGroups200ResponseGroup) GetStatsOk() (*ListGroups200ResponseAllOfGroupsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *UpdateGroups200ResponseGroup) SetStats(v ListGroups200ResponseAllOfGroupsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *UpdateGroups200ResponseGroup) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *UpdateGroups200ResponseGroup) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *UpdateGroups200ResponseGroup) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *UpdateGroups200ResponseGroup) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *UpdateGroups200ResponseGroup) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateGroups200ResponseGroup) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateGroups200ResponseGroup) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateGroups200ResponseGroup) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateGroups200ResponseGroup) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


