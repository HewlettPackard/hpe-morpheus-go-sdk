# GetTenantSubtenantGroup200ResponseGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Zones** | Pointer to [**[]CreateTenantSubtenantGroup200ResponseAccountAllOfZonesInner**](CreateTenantSubtenantGroup200ResponseAccountAllOfZonesInner.md) |  | [optional] 
**Stats** | Pointer to [**CreateTenantSubtenantGroup200ResponseAccountAllOfStats**](CreateTenantSubtenantGroup200ResponseAccountAllOfStats.md) |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetTenantSubtenantGroup200ResponseGroup

`func NewGetTenantSubtenantGroup200ResponseGroup() *GetTenantSubtenantGroup200ResponseGroup`

NewGetTenantSubtenantGroup200ResponseGroup instantiates a new GetTenantSubtenantGroup200ResponseGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetTenantSubtenantGroup200ResponseGroup) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLocation

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *GetTenantSubtenantGroup200ResponseGroup) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetAccountId

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetVisibility

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetZones

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetZones() []CreateTenantSubtenantGroup200ResponseAccountAllOfZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetZonesOk() (*[]CreateTenantSubtenantGroup200ResponseAccountAllOfZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetZones(v []CreateTenantSubtenantGroup200ResponseAccountAllOfZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetStats

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetStats() CreateTenantSubtenantGroup200ResponseAccountAllOfStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetStatsOk() (*CreateTenantSubtenantGroup200ResponseAccountAllOfStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetStats(v CreateTenantSubtenantGroup200ResponseAccountAllOfStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *GetTenantSubtenantGroup200ResponseGroup) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *GetTenantSubtenantGroup200ResponseGroup) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *GetTenantSubtenantGroup200ResponseGroup) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


