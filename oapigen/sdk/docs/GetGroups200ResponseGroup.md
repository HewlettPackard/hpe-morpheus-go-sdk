# GetGroups200ResponseGroup

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
**Config** | Pointer to [**GetGroups200ResponseGroupConfig**](GetGroups200ResponseGroupConfig.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Zones** | Pointer to [**[]GetGroups200ResponseGroupZonesInner**](GetGroups200ResponseGroupZonesInner.md) |  | [optional] 
**Stats** | Pointer to [**GetGroups200ResponseGroupStats**](GetGroups200ResponseGroupStats.md) |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetGroups200ResponseGroup

`func NewGetGroups200ResponseGroup() *GetGroups200ResponseGroup`

NewGetGroups200ResponseGroup instantiates a new GetGroups200ResponseGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroups200ResponseGroupWithDefaults

`func NewGetGroups200ResponseGroupWithDefaults() *GetGroups200ResponseGroup`

NewGetGroups200ResponseGroupWithDefaults instantiates a new GetGroups200ResponseGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGroups200ResponseGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGroups200ResponseGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGroups200ResponseGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetGroups200ResponseGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetGroups200ResponseGroup) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGroups200ResponseGroup) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGroups200ResponseGroup) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGroups200ResponseGroup) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *GetGroups200ResponseGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGroups200ResponseGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGroups200ResponseGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGroups200ResponseGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetGroups200ResponseGroup) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetGroups200ResponseGroup) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetGroups200ResponseGroup) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetGroups200ResponseGroup) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetGroups200ResponseGroup) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetGroups200ResponseGroup) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLabels

`func (o *GetGroups200ResponseGroup) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetGroups200ResponseGroup) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetGroups200ResponseGroup) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetGroups200ResponseGroup) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *GetGroups200ResponseGroup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetGroups200ResponseGroup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetGroups200ResponseGroup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetGroups200ResponseGroup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *GetGroups200ResponseGroup) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *GetGroups200ResponseGroup) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetAccountId

`func (o *GetGroups200ResponseGroup) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGroups200ResponseGroup) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGroups200ResponseGroup) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGroups200ResponseGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActive

`func (o *GetGroups200ResponseGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetGroups200ResponseGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetGroups200ResponseGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetGroups200ResponseGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetConfig

`func (o *GetGroups200ResponseGroup) GetConfig() GetGroups200ResponseGroupConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetGroups200ResponseGroup) GetConfigOk() (*GetGroups200ResponseGroupConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetGroups200ResponseGroup) SetConfig(v GetGroups200ResponseGroupConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetGroups200ResponseGroup) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetGroups200ResponseGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetGroups200ResponseGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetGroups200ResponseGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetGroups200ResponseGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetGroups200ResponseGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetGroups200ResponseGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetGroups200ResponseGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetGroups200ResponseGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetZones

`func (o *GetGroups200ResponseGroup) GetZones() []GetGroups200ResponseGroupZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *GetGroups200ResponseGroup) GetZonesOk() (*[]GetGroups200ResponseGroupZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *GetGroups200ResponseGroup) SetZones(v []GetGroups200ResponseGroupZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *GetGroups200ResponseGroup) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetStats

`func (o *GetGroups200ResponseGroup) GetStats() GetGroups200ResponseGroupStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetGroups200ResponseGroup) GetStatsOk() (*GetGroups200ResponseGroupStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetGroups200ResponseGroup) SetStats(v GetGroups200ResponseGroupStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetGroups200ResponseGroup) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *GetGroups200ResponseGroup) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *GetGroups200ResponseGroup) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *GetGroups200ResponseGroup) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *GetGroups200ResponseGroup) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


