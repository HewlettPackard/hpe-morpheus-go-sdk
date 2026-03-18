# UpdateTenantSubtenantGroup200ResponseAccount

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
**Zones** | Pointer to [**[]UpdateTenantSubtenantGroup200ResponseAccountAllOfZonesInner**](UpdateTenantSubtenantGroup200ResponseAccountAllOfZonesInner.md) |  | [optional] 
**Stats** | Pointer to [**UpdateTenantSubtenantGroup200ResponseAccountAllOfStats**](UpdateTenantSubtenantGroup200ResponseAccountAllOfStats.md) |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateTenantSubtenantGroup200ResponseAccount

`func NewUpdateTenantSubtenantGroup200ResponseAccount() *UpdateTenantSubtenantGroup200ResponseAccount`

NewUpdateTenantSubtenantGroup200ResponseAccount instantiates a new UpdateTenantSubtenantGroup200ResponseAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantSubtenantGroup200ResponseAccountWithDefaults

`func NewUpdateTenantSubtenantGroup200ResponseAccountWithDefaults() *UpdateTenantSubtenantGroup200ResponseAccount`

NewUpdateTenantSubtenantGroup200ResponseAccountWithDefaults instantiates a new UpdateTenantSubtenantGroup200ResponseAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateTenantSubtenantGroup200ResponseAccount) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLocation

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *UpdateTenantSubtenantGroup200ResponseAccount) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetAccountId

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetZones

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetZones() []UpdateTenantSubtenantGroup200ResponseAccountAllOfZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetZonesOk() (*[]UpdateTenantSubtenantGroup200ResponseAccountAllOfZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetZones(v []UpdateTenantSubtenantGroup200ResponseAccountAllOfZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetStats

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetStats() UpdateTenantSubtenantGroup200ResponseAccountAllOfStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetStatsOk() (*UpdateTenantSubtenantGroup200ResponseAccountAllOfStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetStats(v UpdateTenantSubtenantGroup200ResponseAccountAllOfStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateTenantSubtenantGroup200ResponseAccount) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


