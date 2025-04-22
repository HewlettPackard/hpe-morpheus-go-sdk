# ListGroups200ResponseAllOfGroupsInner

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

## Methods

### NewListGroups200ResponseAllOfGroupsInner

`func NewListGroups200ResponseAllOfGroupsInner() *ListGroups200ResponseAllOfGroupsInner`

NewListGroups200ResponseAllOfGroupsInner instantiates a new ListGroups200ResponseAllOfGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGroups200ResponseAllOfGroupsInnerWithDefaults

`func NewListGroups200ResponseAllOfGroupsInnerWithDefaults() *ListGroups200ResponseAllOfGroupsInner`

NewListGroups200ResponseAllOfGroupsInnerWithDefaults instantiates a new ListGroups200ResponseAllOfGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListGroups200ResponseAllOfGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListGroups200ResponseAllOfGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListGroups200ResponseAllOfGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListGroups200ResponseAllOfGroupsInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListGroups200ResponseAllOfGroupsInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListGroups200ResponseAllOfGroupsInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ListGroups200ResponseAllOfGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListGroups200ResponseAllOfGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListGroups200ResponseAllOfGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListGroups200ResponseAllOfGroupsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListGroups200ResponseAllOfGroupsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListGroups200ResponseAllOfGroupsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabels

`func (o *ListGroups200ResponseAllOfGroupsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListGroups200ResponseAllOfGroupsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListGroups200ResponseAllOfGroupsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *ListGroups200ResponseAllOfGroupsInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ListGroups200ResponseAllOfGroupsInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ListGroups200ResponseAllOfGroupsInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAccountId

`func (o *ListGroups200ResponseAllOfGroupsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListGroups200ResponseAllOfGroupsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListGroups200ResponseAllOfGroupsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActive

`func (o *ListGroups200ResponseAllOfGroupsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListGroups200ResponseAllOfGroupsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListGroups200ResponseAllOfGroupsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *ListGroups200ResponseAllOfGroupsInner) GetConfig() ListGroups200ResponseAllOfGroupsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetConfigOk() (*ListGroups200ResponseAllOfGroupsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListGroups200ResponseAllOfGroupsInner) SetConfig(v ListGroups200ResponseAllOfGroupsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListGroups200ResponseAllOfGroupsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListGroups200ResponseAllOfGroupsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListGroups200ResponseAllOfGroupsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListGroups200ResponseAllOfGroupsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListGroups200ResponseAllOfGroupsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListGroups200ResponseAllOfGroupsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListGroups200ResponseAllOfGroupsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetZones

`func (o *ListGroups200ResponseAllOfGroupsInner) GetZones() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetZonesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ListGroups200ResponseAllOfGroupsInner) SetZones(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ListGroups200ResponseAllOfGroupsInner) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetStats

`func (o *ListGroups200ResponseAllOfGroupsInner) GetStats() ListGroups200ResponseAllOfGroupsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetStatsOk() (*ListGroups200ResponseAllOfGroupsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListGroups200ResponseAllOfGroupsInner) SetStats(v ListGroups200ResponseAllOfGroupsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListGroups200ResponseAllOfGroupsInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *ListGroups200ResponseAllOfGroupsInner) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *ListGroups200ResponseAllOfGroupsInner) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *ListGroups200ResponseAllOfGroupsInner) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *ListGroups200ResponseAllOfGroupsInner) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


