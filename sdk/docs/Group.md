# Group

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

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Group) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *Group) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Group) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Group) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Group) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *Group) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Group) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Group) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Group) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabels

`func (o *Group) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Group) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Group) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Group) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *Group) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Group) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Group) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Group) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAccountId

`func (o *Group) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Group) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Group) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Group) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActive

`func (o *Group) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Group) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Group) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Group) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *Group) GetConfig() ListGroups200ResponseAllOfGroupsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Group) GetConfigOk() (*ListGroups200ResponseAllOfGroupsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Group) SetConfig(v ListGroups200ResponseAllOfGroupsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Group) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *Group) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Group) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Group) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Group) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Group) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Group) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Group) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Group) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetZones

`func (o *Group) GetZones() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *Group) GetZonesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *Group) SetZones(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *Group) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetStats

`func (o *Group) GetStats() ListGroups200ResponseAllOfGroupsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Group) GetStatsOk() (*ListGroups200ResponseAllOfGroupsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Group) SetStats(v ListGroups200ResponseAllOfGroupsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Group) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *Group) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *Group) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *Group) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *Group) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


