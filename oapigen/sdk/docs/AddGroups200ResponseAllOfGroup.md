# AddGroups200ResponseAllOfGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**AddGroups200ResponseAllOfGroupConfig**](AddGroups200ResponseAllOfGroupConfig.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Zones** | Pointer to [**[]AddGroups200ResponseAllOfGroupZonesInner**](AddGroups200ResponseAllOfGroupZonesInner.md) |  | [optional] 
**Stats** | Pointer to [**AddGroups200ResponseAllOfGroupStats**](AddGroups200ResponseAllOfGroupStats.md) |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewAddGroups200ResponseAllOfGroup

`func NewAddGroups200ResponseAllOfGroup() *AddGroups200ResponseAllOfGroup`

NewAddGroups200ResponseAllOfGroup instantiates a new AddGroups200ResponseAllOfGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroups200ResponseAllOfGroupWithDefaults

`func NewAddGroups200ResponseAllOfGroupWithDefaults() *AddGroups200ResponseAllOfGroup`

NewAddGroups200ResponseAllOfGroupWithDefaults instantiates a new AddGroups200ResponseAllOfGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddGroups200ResponseAllOfGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddGroups200ResponseAllOfGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddGroups200ResponseAllOfGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddGroups200ResponseAllOfGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *AddGroups200ResponseAllOfGroup) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AddGroups200ResponseAllOfGroup) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AddGroups200ResponseAllOfGroup) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AddGroups200ResponseAllOfGroup) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *AddGroups200ResponseAllOfGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddGroups200ResponseAllOfGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddGroups200ResponseAllOfGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddGroups200ResponseAllOfGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddGroups200ResponseAllOfGroup) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddGroups200ResponseAllOfGroup) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddGroups200ResponseAllOfGroup) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddGroups200ResponseAllOfGroup) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AddGroups200ResponseAllOfGroup) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AddGroups200ResponseAllOfGroup) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLabels

`func (o *AddGroups200ResponseAllOfGroup) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddGroups200ResponseAllOfGroup) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddGroups200ResponseAllOfGroup) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddGroups200ResponseAllOfGroup) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *AddGroups200ResponseAllOfGroup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AddGroups200ResponseAllOfGroup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AddGroups200ResponseAllOfGroup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AddGroups200ResponseAllOfGroup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AddGroups200ResponseAllOfGroup) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AddGroups200ResponseAllOfGroup) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetAccountId

`func (o *AddGroups200ResponseAllOfGroup) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddGroups200ResponseAllOfGroup) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddGroups200ResponseAllOfGroup) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddGroups200ResponseAllOfGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActive

`func (o *AddGroups200ResponseAllOfGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddGroups200ResponseAllOfGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddGroups200ResponseAllOfGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddGroups200ResponseAllOfGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *AddGroups200ResponseAllOfGroup) GetConfig() AddGroups200ResponseAllOfGroupConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddGroups200ResponseAllOfGroup) GetConfigOk() (*AddGroups200ResponseAllOfGroupConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddGroups200ResponseAllOfGroup) SetConfig(v AddGroups200ResponseAllOfGroupConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddGroups200ResponseAllOfGroup) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddGroups200ResponseAllOfGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddGroups200ResponseAllOfGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddGroups200ResponseAllOfGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddGroups200ResponseAllOfGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddGroups200ResponseAllOfGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddGroups200ResponseAllOfGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddGroups200ResponseAllOfGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddGroups200ResponseAllOfGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetZones

`func (o *AddGroups200ResponseAllOfGroup) GetZones() []AddGroups200ResponseAllOfGroupZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AddGroups200ResponseAllOfGroup) GetZonesOk() (*[]AddGroups200ResponseAllOfGroupZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AddGroups200ResponseAllOfGroup) SetZones(v []AddGroups200ResponseAllOfGroupZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AddGroups200ResponseAllOfGroup) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetStats

`func (o *AddGroups200ResponseAllOfGroup) GetStats() AddGroups200ResponseAllOfGroupStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AddGroups200ResponseAllOfGroup) GetStatsOk() (*AddGroups200ResponseAllOfGroupStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AddGroups200ResponseAllOfGroup) SetStats(v AddGroups200ResponseAllOfGroupStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AddGroups200ResponseAllOfGroup) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *AddGroups200ResponseAllOfGroup) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *AddGroups200ResponseAllOfGroup) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *AddGroups200ResponseAllOfGroup) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *AddGroups200ResponseAllOfGroup) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


