# ListTenantSubtenantGroups200ResponseAllOfGroupsInner

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

## Methods

### NewListTenantSubtenantGroups200ResponseAllOfGroupsInner

`func NewListTenantSubtenantGroups200ResponseAllOfGroupsInner() *ListTenantSubtenantGroups200ResponseAllOfGroupsInner`

NewListTenantSubtenantGroups200ResponseAllOfGroupsInner instantiates a new ListTenantSubtenantGroups200ResponseAllOfGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTenantSubtenantGroups200ResponseAllOfGroupsInnerWithDefaults

`func NewListTenantSubtenantGroups200ResponseAllOfGroupsInnerWithDefaults() *ListTenantSubtenantGroups200ResponseAllOfGroupsInner`

NewListTenantSubtenantGroups200ResponseAllOfGroupsInnerWithDefaults instantiates a new ListTenantSubtenantGroups200ResponseAllOfGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLocation

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAccountId

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetVisibility

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetZones

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetZones() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetZonesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetZones(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetStats

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetStats() ListGroups200ResponseAllOfGroupsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetStatsOk() (*ListGroups200ResponseAllOfGroupsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetStats(v ListGroups200ResponseAllOfGroupsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *ListTenantSubtenantGroups200ResponseAllOfGroupsInner) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


