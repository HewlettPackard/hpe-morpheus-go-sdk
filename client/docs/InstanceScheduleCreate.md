# InstanceScheduleCreate

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

### NewInstanceScheduleCreate

`func NewInstanceScheduleCreate() *InstanceScheduleCreate`

NewInstanceScheduleCreate instantiates a new InstanceScheduleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceScheduleCreateWithDefaults

`func NewInstanceScheduleCreateWithDefaults() *InstanceScheduleCreate`

NewInstanceScheduleCreateWithDefaults instantiates a new InstanceScheduleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleType

`func (o *InstanceScheduleCreate) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *InstanceScheduleCreate) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *InstanceScheduleCreate) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *InstanceScheduleCreate) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *InstanceScheduleCreate) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *InstanceScheduleCreate) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *InstanceScheduleCreate) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *InstanceScheduleCreate) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetStartDayOfWeek

`func (o *InstanceScheduleCreate) GetStartDayOfWeek() int32`

GetStartDayOfWeek returns the StartDayOfWeek field if non-nil, zero value otherwise.

### GetStartDayOfWeekOk

`func (o *InstanceScheduleCreate) GetStartDayOfWeekOk() (*int32, bool)`

GetStartDayOfWeekOk returns a tuple with the StartDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDayOfWeek

`func (o *InstanceScheduleCreate) SetStartDayOfWeek(v int32)`

SetStartDayOfWeek sets StartDayOfWeek field to given value.

### HasStartDayOfWeek

`func (o *InstanceScheduleCreate) HasStartDayOfWeek() bool`

HasStartDayOfWeek returns a boolean if a field has been set.

### GetStartTime

`func (o *InstanceScheduleCreate) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InstanceScheduleCreate) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InstanceScheduleCreate) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InstanceScheduleCreate) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndDayOfWeek

`func (o *InstanceScheduleCreate) GetEndDayOfWeek() int32`

GetEndDayOfWeek returns the EndDayOfWeek field if non-nil, zero value otherwise.

### GetEndDayOfWeekOk

`func (o *InstanceScheduleCreate) GetEndDayOfWeekOk() (*int32, bool)`

GetEndDayOfWeekOk returns a tuple with the EndDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDayOfWeek

`func (o *InstanceScheduleCreate) SetEndDayOfWeek(v int32)`

SetEndDayOfWeek sets EndDayOfWeek field to given value.

### HasEndDayOfWeek

`func (o *InstanceScheduleCreate) HasEndDayOfWeek() bool`

HasEndDayOfWeek returns a boolean if a field has been set.

### GetEndTime

`func (o *InstanceScheduleCreate) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InstanceScheduleCreate) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InstanceScheduleCreate) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *InstanceScheduleCreate) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartDate

`func (o *InstanceScheduleCreate) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InstanceScheduleCreate) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InstanceScheduleCreate) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InstanceScheduleCreate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InstanceScheduleCreate) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InstanceScheduleCreate) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InstanceScheduleCreate) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InstanceScheduleCreate) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetThreshold

`func (o *InstanceScheduleCreate) GetThreshold() CreateInstanceScheduleRequestInstanceScheduleThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *InstanceScheduleCreate) GetThresholdOk() (*CreateInstanceScheduleRequestInstanceScheduleThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *InstanceScheduleCreate) SetThreshold(v CreateInstanceScheduleRequestInstanceScheduleThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *InstanceScheduleCreate) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


