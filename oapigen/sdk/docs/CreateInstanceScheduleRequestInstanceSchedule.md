# CreateInstanceScheduleRequestInstanceSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleType** | Pointer to **string** |  | [optional] [default to "dayOfWeek"]
**ScheduleTimezone** | Pointer to **string** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] [default to "UTC"]
**StartDayOfWeek** | Pointer to **int32** | Start day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**StartTime** | Pointer to **string** | Start time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**EndDayOfWeek** | Pointer to **int32** | End day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**EndTime** | Pointer to **string** | End time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**StartDate** | Pointer to **time.Time** | Start Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional] 
**EndDate** | Pointer to **time.Time** | End Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional] 
**Threshold** | Pointer to [**CreateInstanceScheduleRequestInstanceScheduleThreshold**](CreateInstanceScheduleRequestInstanceScheduleThreshold.md) |  | [optional] 

## Methods

### NewCreateInstanceScheduleRequestInstanceSchedule

`func NewCreateInstanceScheduleRequestInstanceSchedule() *CreateInstanceScheduleRequestInstanceSchedule`

NewCreateInstanceScheduleRequestInstanceSchedule instantiates a new CreateInstanceScheduleRequestInstanceSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceScheduleRequestInstanceScheduleWithDefaults

`func NewCreateInstanceScheduleRequestInstanceScheduleWithDefaults() *CreateInstanceScheduleRequestInstanceSchedule`

NewCreateInstanceScheduleRequestInstanceScheduleWithDefaults instantiates a new CreateInstanceScheduleRequestInstanceSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleType

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *CreateInstanceScheduleRequestInstanceSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *CreateInstanceScheduleRequestInstanceSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *CreateInstanceScheduleRequestInstanceSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *CreateInstanceScheduleRequestInstanceSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetStartDayOfWeek

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetStartDayOfWeek() int32`

GetStartDayOfWeek returns the StartDayOfWeek field if non-nil, zero value otherwise.

### GetStartDayOfWeekOk

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetStartDayOfWeekOk() (*int32, bool)`

GetStartDayOfWeekOk returns a tuple with the StartDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDayOfWeek

`func (o *CreateInstanceScheduleRequestInstanceSchedule) SetStartDayOfWeek(v int32)`

SetStartDayOfWeek sets StartDayOfWeek field to given value.

### HasStartDayOfWeek

`func (o *CreateInstanceScheduleRequestInstanceSchedule) HasStartDayOfWeek() bool`

HasStartDayOfWeek returns a boolean if a field has been set.

### GetStartTime

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateInstanceScheduleRequestInstanceSchedule) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CreateInstanceScheduleRequestInstanceSchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndDayOfWeek

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetEndDayOfWeek() int32`

GetEndDayOfWeek returns the EndDayOfWeek field if non-nil, zero value otherwise.

### GetEndDayOfWeekOk

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetEndDayOfWeekOk() (*int32, bool)`

GetEndDayOfWeekOk returns a tuple with the EndDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDayOfWeek

`func (o *CreateInstanceScheduleRequestInstanceSchedule) SetEndDayOfWeek(v int32)`

SetEndDayOfWeek sets EndDayOfWeek field to given value.

### HasEndDayOfWeek

`func (o *CreateInstanceScheduleRequestInstanceSchedule) HasEndDayOfWeek() bool`

HasEndDayOfWeek returns a boolean if a field has been set.

### GetEndTime

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CreateInstanceScheduleRequestInstanceSchedule) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CreateInstanceScheduleRequestInstanceSchedule) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateInstanceScheduleRequestInstanceSchedule) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateInstanceScheduleRequestInstanceSchedule) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateInstanceScheduleRequestInstanceSchedule) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateInstanceScheduleRequestInstanceSchedule) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetThreshold

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetThreshold() CreateInstanceScheduleRequestInstanceScheduleThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *CreateInstanceScheduleRequestInstanceSchedule) GetThresholdOk() (*CreateInstanceScheduleRequestInstanceScheduleThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *CreateInstanceScheduleRequestInstanceSchedule) SetThreshold(v CreateInstanceScheduleRequestInstanceScheduleThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *CreateInstanceScheduleRequestInstanceSchedule) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


