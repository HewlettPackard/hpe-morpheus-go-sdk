# UpdateBudgetsRequestBudget

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
**ForecastType** | Pointer to [**UpdateBudgetsRequestBudgetForecastType**](UpdateBudgetsRequestBudgetForecastType.md) |  | [optional] 

## Methods

### NewUpdateBudgetsRequestBudget

`func NewUpdateBudgetsRequestBudget(name string, ) *UpdateBudgetsRequestBudget`

NewUpdateBudgetsRequestBudget instantiates a new UpdateBudgetsRequestBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateBudgetsRequestBudget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBudgetsRequestBudget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBudgetsRequestBudget) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateBudgetsRequestBudget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateBudgetsRequestBudget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateBudgetsRequestBudget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateBudgetsRequestBudget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScope

`func (o *UpdateBudgetsRequestBudget) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateBudgetsRequestBudget) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateBudgetsRequestBudget) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UpdateBudgetsRequestBudget) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetPeriod

`func (o *UpdateBudgetsRequestBudget) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *UpdateBudgetsRequestBudget) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *UpdateBudgetsRequestBudget) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *UpdateBudgetsRequestBudget) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetYear

`func (o *UpdateBudgetsRequestBudget) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *UpdateBudgetsRequestBudget) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *UpdateBudgetsRequestBudget) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *UpdateBudgetsRequestBudget) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateBudgetsRequestBudget) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateBudgetsRequestBudget) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateBudgetsRequestBudget) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateBudgetsRequestBudget) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateBudgetsRequestBudget) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateBudgetsRequestBudget) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateBudgetsRequestBudget) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateBudgetsRequestBudget) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetInterval

`func (o *UpdateBudgetsRequestBudget) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UpdateBudgetsRequestBudget) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UpdateBudgetsRequestBudget) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UpdateBudgetsRequestBudget) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetScopeTenantId

`func (o *UpdateBudgetsRequestBudget) GetScopeTenantId() int64`

GetScopeTenantId returns the ScopeTenantId field if non-nil, zero value otherwise.

### GetScopeTenantIdOk

`func (o *UpdateBudgetsRequestBudget) GetScopeTenantIdOk() (*int64, bool)`

GetScopeTenantIdOk returns a tuple with the ScopeTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeTenantId

`func (o *UpdateBudgetsRequestBudget) SetScopeTenantId(v int64)`

SetScopeTenantId sets ScopeTenantId field to given value.

### HasScopeTenantId

`func (o *UpdateBudgetsRequestBudget) HasScopeTenantId() bool`

HasScopeTenantId returns a boolean if a field has been set.

### GetScopeGroupId

`func (o *UpdateBudgetsRequestBudget) GetScopeGroupId() int64`

GetScopeGroupId returns the ScopeGroupId field if non-nil, zero value otherwise.

### GetScopeGroupIdOk

`func (o *UpdateBudgetsRequestBudget) GetScopeGroupIdOk() (*int64, bool)`

GetScopeGroupIdOk returns a tuple with the ScopeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeGroupId

`func (o *UpdateBudgetsRequestBudget) SetScopeGroupId(v int64)`

SetScopeGroupId sets ScopeGroupId field to given value.

### HasScopeGroupId

`func (o *UpdateBudgetsRequestBudget) HasScopeGroupId() bool`

HasScopeGroupId returns a boolean if a field has been set.

### GetScopeCloudId

`func (o *UpdateBudgetsRequestBudget) GetScopeCloudId() int64`

GetScopeCloudId returns the ScopeCloudId field if non-nil, zero value otherwise.

### GetScopeCloudIdOk

`func (o *UpdateBudgetsRequestBudget) GetScopeCloudIdOk() (*int64, bool)`

GetScopeCloudIdOk returns a tuple with the ScopeCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeCloudId

`func (o *UpdateBudgetsRequestBudget) SetScopeCloudId(v int64)`

SetScopeCloudId sets ScopeCloudId field to given value.

### HasScopeCloudId

`func (o *UpdateBudgetsRequestBudget) HasScopeCloudId() bool`

HasScopeCloudId returns a boolean if a field has been set.

### GetScopeUserId

`func (o *UpdateBudgetsRequestBudget) GetScopeUserId() int64`

GetScopeUserId returns the ScopeUserId field if non-nil, zero value otherwise.

### GetScopeUserIdOk

`func (o *UpdateBudgetsRequestBudget) GetScopeUserIdOk() (*int64, bool)`

GetScopeUserIdOk returns a tuple with the ScopeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeUserId

`func (o *UpdateBudgetsRequestBudget) SetScopeUserId(v int64)`

SetScopeUserId sets ScopeUserId field to given value.

### HasScopeUserId

`func (o *UpdateBudgetsRequestBudget) HasScopeUserId() bool`

HasScopeUserId returns a boolean if a field has been set.

### GetCosts

`func (o *UpdateBudgetsRequestBudget) GetCosts() []int64`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *UpdateBudgetsRequestBudget) GetCostsOk() (*[]int64, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *UpdateBudgetsRequestBudget) SetCosts(v []int64)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *UpdateBudgetsRequestBudget) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateBudgetsRequestBudget) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateBudgetsRequestBudget) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateBudgetsRequestBudget) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateBudgetsRequestBudget) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetForecastType

`func (o *UpdateBudgetsRequestBudget) GetForecastType() UpdateBudgetsRequestBudgetForecastType`

GetForecastType returns the ForecastType field if non-nil, zero value otherwise.

### GetForecastTypeOk

`func (o *UpdateBudgetsRequestBudget) GetForecastTypeOk() (*UpdateBudgetsRequestBudgetForecastType, bool)`

GetForecastTypeOk returns a tuple with the ForecastType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastType

`func (o *UpdateBudgetsRequestBudget) SetForecastType(v UpdateBudgetsRequestBudgetForecastType)`

SetForecastType sets ForecastType field to given value.

### HasForecastType

`func (o *UpdateBudgetsRequestBudget) HasForecastType() bool`

HasForecastType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


