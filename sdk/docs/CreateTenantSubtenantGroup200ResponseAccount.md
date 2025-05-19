# CreateTenantSubtenantGroup200ResponseAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Zones** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Stats** | Pointer to [**ListGroups200ResponseAllOfGroupsInnerStats**](ListGroups200ResponseAllOfGroupsInnerStats.md) |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateTenantSubtenantGroup200ResponseAccount

`func NewCreateTenantSubtenantGroup200ResponseAccount() *CreateTenantSubtenantGroup200ResponseAccount`

NewCreateTenantSubtenantGroup200ResponseAccount instantiates a new CreateTenantSubtenantGroup200ResponseAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantSubtenantGroup200ResponseAccountWithDefaults

`func NewCreateTenantSubtenantGroup200ResponseAccountWithDefaults() *CreateTenantSubtenantGroup200ResponseAccount`

NewCreateTenantSubtenantGroup200ResponseAccountWithDefaults instantiates a new CreateTenantSubtenantGroup200ResponseAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLocation

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetZones

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetZones() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetZonesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetZones(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetStats

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetStats() ListGroups200ResponseAllOfGroupsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetStatsOk() (*ListGroups200ResponseAllOfGroupsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetStats(v ListGroups200ResponseAllOfGroupsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.

### GetSuccess

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateTenantSubtenantGroup200ResponseAccount) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateTenantSubtenantGroup200ResponseAccount) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateTenantSubtenantGroup200ResponseAccount) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


