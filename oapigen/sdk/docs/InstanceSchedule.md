# InstanceSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ScheduleType** | Pointer to **string** |  | [optional] [default to "dayOfWeek"]
**ScheduleTimezone** | Pointer to **string** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] [default to "UTC"]
**StartDayOfWeek** | Pointer to **int32** | Start day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**StartTime** | Pointer to **string** | Start time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**EndDayOfWeek** | Pointer to **int32** | End day of the week 1-7 (Sun-Sat). Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**EndTime** | Pointer to **string** | End time of the day in 24-hour format. Only used and required for scheduleType &#x60;dayOfWeek&#x60; | [optional] 
**StartDate** | Pointer to **time.Time** | Start Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional] 
**EndDate** | Pointer to **time.Time** | End Date. Only used and required for scheduleType &#x60;exact&#x60; | [optional] 
**StartDisplay** | Pointer to **string** | Start day and time or start date formatted for display | [optional] 
**EndDisplay** | Pointer to **string** | End day and time or end date formatted for display | [optional] 
**Threshold** | Pointer to [**GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold**](GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInstanceSchedule

`func NewInstanceSchedule() *InstanceSchedule`

NewInstanceSchedule instantiates a new InstanceSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceScheduleWithDefaults

`func NewInstanceScheduleWithDefaults() *InstanceSchedule`

NewInstanceScheduleWithDefaults instantiates a new InstanceSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceSchedule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceSchedule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceSchedule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduleType

`func (o *InstanceSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *InstanceSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *InstanceSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *InstanceSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *InstanceSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *InstanceSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *InstanceSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *InstanceSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetStartDayOfWeek

`func (o *InstanceSchedule) GetStartDayOfWeek() int32`

GetStartDayOfWeek returns the StartDayOfWeek field if non-nil, zero value otherwise.

### GetStartDayOfWeekOk

`func (o *InstanceSchedule) GetStartDayOfWeekOk() (*int32, bool)`

GetStartDayOfWeekOk returns a tuple with the StartDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDayOfWeek

`func (o *InstanceSchedule) SetStartDayOfWeek(v int32)`

SetStartDayOfWeek sets StartDayOfWeek field to given value.

### HasStartDayOfWeek

`func (o *InstanceSchedule) HasStartDayOfWeek() bool`

HasStartDayOfWeek returns a boolean if a field has been set.

### GetStartTime

`func (o *InstanceSchedule) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InstanceSchedule) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InstanceSchedule) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InstanceSchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndDayOfWeek

`func (o *InstanceSchedule) GetEndDayOfWeek() int32`

GetEndDayOfWeek returns the EndDayOfWeek field if non-nil, zero value otherwise.

### GetEndDayOfWeekOk

`func (o *InstanceSchedule) GetEndDayOfWeekOk() (*int32, bool)`

GetEndDayOfWeekOk returns a tuple with the EndDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDayOfWeek

`func (o *InstanceSchedule) SetEndDayOfWeek(v int32)`

SetEndDayOfWeek sets EndDayOfWeek field to given value.

### HasEndDayOfWeek

`func (o *InstanceSchedule) HasEndDayOfWeek() bool`

HasEndDayOfWeek returns a boolean if a field has been set.

### GetEndTime

`func (o *InstanceSchedule) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InstanceSchedule) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InstanceSchedule) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *InstanceSchedule) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartDate

`func (o *InstanceSchedule) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InstanceSchedule) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InstanceSchedule) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *InstanceSchedule) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *InstanceSchedule) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InstanceSchedule) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InstanceSchedule) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *InstanceSchedule) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDisplay

`func (o *InstanceSchedule) GetStartDisplay() string`

GetStartDisplay returns the StartDisplay field if non-nil, zero value otherwise.

### GetStartDisplayOk

`func (o *InstanceSchedule) GetStartDisplayOk() (*string, bool)`

GetStartDisplayOk returns a tuple with the StartDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDisplay

`func (o *InstanceSchedule) SetStartDisplay(v string)`

SetStartDisplay sets StartDisplay field to given value.

### HasStartDisplay

`func (o *InstanceSchedule) HasStartDisplay() bool`

HasStartDisplay returns a boolean if a field has been set.

### GetEndDisplay

`func (o *InstanceSchedule) GetEndDisplay() string`

GetEndDisplay returns the EndDisplay field if non-nil, zero value otherwise.

### GetEndDisplayOk

`func (o *InstanceSchedule) GetEndDisplayOk() (*string, bool)`

GetEndDisplayOk returns a tuple with the EndDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDisplay

`func (o *InstanceSchedule) SetEndDisplay(v string)`

SetEndDisplay sets EndDisplay field to given value.

### HasEndDisplay

`func (o *InstanceSchedule) HasEndDisplay() bool`

HasEndDisplay returns a boolean if a field has been set.

### GetThreshold

`func (o *InstanceSchedule) GetThreshold() GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *InstanceSchedule) GetThresholdOk() (*GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *InstanceSchedule) SetThreshold(v GetInstanceThreshold200ResponseInstanceSchedulesInnerThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *InstanceSchedule) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetDateCreated

`func (o *InstanceSchedule) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InstanceSchedule) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InstanceSchedule) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InstanceSchedule) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InstanceSchedule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InstanceSchedule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InstanceSchedule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InstanceSchedule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


