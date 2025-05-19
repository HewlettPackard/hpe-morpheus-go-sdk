# ListExecuteSchedules200ResponseAllOfSchedulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ScheduleType** | Pointer to **string** |  | [optional] 
**ScheduleTimezone** | Pointer to **string** |  | [optional] 
**Cron** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListExecuteSchedules200ResponseAllOfSchedulesInner

`func NewListExecuteSchedules200ResponseAllOfSchedulesInner() *ListExecuteSchedules200ResponseAllOfSchedulesInner`

NewListExecuteSchedules200ResponseAllOfSchedulesInner instantiates a new ListExecuteSchedules200ResponseAllOfSchedulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListExecuteSchedules200ResponseAllOfSchedulesInnerWithDefaults

`func NewListExecuteSchedules200ResponseAllOfSchedulesInnerWithDefaults() *ListExecuteSchedules200ResponseAllOfSchedulesInner`

NewListExecuteSchedules200ResponseAllOfSchedulesInnerWithDefaults instantiates a new ListExecuteSchedules200ResponseAllOfSchedulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnabled

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetScheduleType

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetCron

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListExecuteSchedules200ResponseAllOfSchedulesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


