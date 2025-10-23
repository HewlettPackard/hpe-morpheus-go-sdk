# AddBudgetsRequestBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] [default to "account"]
**Period** | Pointer to **string** |  | [optional] [default to "year"]
**Year** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Interval** | Pointer to **string** |  | [optional] [default to "year"]
**ScopeTenantId** | Pointer to **int64** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;tenant  | [optional] 
**ScopeGroupId** | Pointer to **int64** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;group   | [optional] 
**ScopeCloudId** | Pointer to **int64** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;cloud  | [optional] 
**ScopeUserId** | Pointer to **int64** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;user  | [optional] 
**Costs** | Pointer to **[]int64** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**ForecastType** | Pointer to [**AddBudgetsRequestBudgetForecastType**](AddBudgetsRequestBudgetForecastType.md) |  | [optional] 

## Methods

### NewAddBudgetsRequestBudget

`func NewAddBudgetsRequestBudget(name string, ) *AddBudgetsRequestBudget`

NewAddBudgetsRequestBudget instantiates a new AddBudgetsRequestBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBudgetsRequestBudgetWithDefaults

`func NewAddBudgetsRequestBudgetWithDefaults() *AddBudgetsRequestBudget`

NewAddBudgetsRequestBudgetWithDefaults instantiates a new AddBudgetsRequestBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddBudgetsRequestBudget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBudgetsRequestBudget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBudgetsRequestBudget) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddBudgetsRequestBudget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddBudgetsRequestBudget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddBudgetsRequestBudget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddBudgetsRequestBudget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScope

`func (o *AddBudgetsRequestBudget) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AddBudgetsRequestBudget) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AddBudgetsRequestBudget) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AddBudgetsRequestBudget) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetPeriod

`func (o *AddBudgetsRequestBudget) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AddBudgetsRequestBudget) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AddBudgetsRequestBudget) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AddBudgetsRequestBudget) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetYear

`func (o *AddBudgetsRequestBudget) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AddBudgetsRequestBudget) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AddBudgetsRequestBudget) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *AddBudgetsRequestBudget) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetStartDate

`func (o *AddBudgetsRequestBudget) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AddBudgetsRequestBudget) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AddBudgetsRequestBudget) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AddBudgetsRequestBudget) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *AddBudgetsRequestBudget) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AddBudgetsRequestBudget) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AddBudgetsRequestBudget) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AddBudgetsRequestBudget) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInterval

`func (o *AddBudgetsRequestBudget) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *AddBudgetsRequestBudget) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *AddBudgetsRequestBudget) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *AddBudgetsRequestBudget) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetScopeTenantId

`func (o *AddBudgetsRequestBudget) GetScopeTenantId() int64`

GetScopeTenantId returns the ScopeTenantId field if non-nil, zero value otherwise.

### GetScopeTenantIdOk

`func (o *AddBudgetsRequestBudget) GetScopeTenantIdOk() (*int64, bool)`

GetScopeTenantIdOk returns a tuple with the ScopeTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeTenantId

`func (o *AddBudgetsRequestBudget) SetScopeTenantId(v int64)`

SetScopeTenantId sets ScopeTenantId field to given value.

### HasScopeTenantId

`func (o *AddBudgetsRequestBudget) HasScopeTenantId() bool`

HasScopeTenantId returns a boolean if a field has been set.

### GetScopeGroupId

`func (o *AddBudgetsRequestBudget) GetScopeGroupId() int64`

GetScopeGroupId returns the ScopeGroupId field if non-nil, zero value otherwise.

### GetScopeGroupIdOk

`func (o *AddBudgetsRequestBudget) GetScopeGroupIdOk() (*int64, bool)`

GetScopeGroupIdOk returns a tuple with the ScopeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeGroupId

`func (o *AddBudgetsRequestBudget) SetScopeGroupId(v int64)`

SetScopeGroupId sets ScopeGroupId field to given value.

### HasScopeGroupId

`func (o *AddBudgetsRequestBudget) HasScopeGroupId() bool`

HasScopeGroupId returns a boolean if a field has been set.

### GetScopeCloudId

`func (o *AddBudgetsRequestBudget) GetScopeCloudId() int64`

GetScopeCloudId returns the ScopeCloudId field if non-nil, zero value otherwise.

### GetScopeCloudIdOk

`func (o *AddBudgetsRequestBudget) GetScopeCloudIdOk() (*int64, bool)`

GetScopeCloudIdOk returns a tuple with the ScopeCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeCloudId

`func (o *AddBudgetsRequestBudget) SetScopeCloudId(v int64)`

SetScopeCloudId sets ScopeCloudId field to given value.

### HasScopeCloudId

`func (o *AddBudgetsRequestBudget) HasScopeCloudId() bool`

HasScopeCloudId returns a boolean if a field has been set.

### GetScopeUserId

`func (o *AddBudgetsRequestBudget) GetScopeUserId() int64`

GetScopeUserId returns the ScopeUserId field if non-nil, zero value otherwise.

### GetScopeUserIdOk

`func (o *AddBudgetsRequestBudget) GetScopeUserIdOk() (*int64, bool)`

GetScopeUserIdOk returns a tuple with the ScopeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeUserId

`func (o *AddBudgetsRequestBudget) SetScopeUserId(v int64)`

SetScopeUserId sets ScopeUserId field to given value.

### HasScopeUserId

`func (o *AddBudgetsRequestBudget) HasScopeUserId() bool`

HasScopeUserId returns a boolean if a field has been set.

### GetCosts

`func (o *AddBudgetsRequestBudget) GetCosts() []int64`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *AddBudgetsRequestBudget) GetCostsOk() (*[]int64, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *AddBudgetsRequestBudget) SetCosts(v []int64)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *AddBudgetsRequestBudget) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetEnabled

`func (o *AddBudgetsRequestBudget) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddBudgetsRequestBudget) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddBudgetsRequestBudget) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddBudgetsRequestBudget) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForecastType

`func (o *AddBudgetsRequestBudget) GetForecastType() AddBudgetsRequestBudgetForecastType`

GetForecastType returns the ForecastType field if non-nil, zero value otherwise.

### GetForecastTypeOk

`func (o *AddBudgetsRequestBudget) GetForecastTypeOk() (*AddBudgetsRequestBudgetForecastType, bool)`

GetForecastTypeOk returns a tuple with the ForecastType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastType

`func (o *AddBudgetsRequestBudget) SetForecastType(v AddBudgetsRequestBudgetForecastType)`

SetForecastType sets ForecastType field to given value.

### HasForecastType

`func (o *AddBudgetsRequestBudget) HasForecastType() bool`

HasForecastType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


