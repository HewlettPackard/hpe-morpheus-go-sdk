# UpdateInstanceSchedule200ResponseAllOfInstanceSchedule

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
**Threshold** | Pointer to [**UpdateInstanceSchedule200ResponseAllOfInstanceScheduleThreshold**](UpdateInstanceSchedule200ResponseAllOfInstanceScheduleThreshold.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateInstanceSchedule200ResponseAllOfInstanceSchedule

`func NewUpdateInstanceSchedule200ResponseAllOfInstanceSchedule() *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule`

NewUpdateInstanceSchedule200ResponseAllOfInstanceSchedule instantiates a new UpdateInstanceSchedule200ResponseAllOfInstanceSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduleType

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetStartDayOfWeek

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetStartDayOfWeek() int32`

GetStartDayOfWeek returns the StartDayOfWeek field if non-nil, zero value otherwise.

### GetStartDayOfWeekOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetStartDayOfWeekOk() (*int32, bool)`

GetStartDayOfWeekOk returns a tuple with the StartDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDayOfWeek

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetStartDayOfWeek(v int32)`

SetStartDayOfWeek sets StartDayOfWeek field to given value.

### HasStartDayOfWeek

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasStartDayOfWeek() bool`

HasStartDayOfWeek returns a boolean if a field has been set.

### GetStartTime

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndDayOfWeek

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetEndDayOfWeek() int32`

GetEndDayOfWeek returns the EndDayOfWeek field if non-nil, zero value otherwise.

### GetEndDayOfWeekOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetEndDayOfWeekOk() (*int32, bool)`

GetEndDayOfWeekOk returns a tuple with the EndDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDayOfWeek

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetEndDayOfWeek(v int32)`

SetEndDayOfWeek sets EndDayOfWeek field to given value.

### HasEndDayOfWeek

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasEndDayOfWeek() bool`

HasEndDayOfWeek returns a boolean if a field has been set.

### GetEndTime

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDisplay

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetStartDisplay() string`

GetStartDisplay returns the StartDisplay field if non-nil, zero value otherwise.

### GetStartDisplayOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetStartDisplayOk() (*string, bool)`

GetStartDisplayOk returns a tuple with the StartDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDisplay

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetStartDisplay(v string)`

SetStartDisplay sets StartDisplay field to given value.

### HasStartDisplay

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasStartDisplay() bool`

HasStartDisplay returns a boolean if a field has been set.

### GetEndDisplay

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetEndDisplay() string`

GetEndDisplay returns the EndDisplay field if non-nil, zero value otherwise.

### GetEndDisplayOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetEndDisplayOk() (*string, bool)`

GetEndDisplayOk returns a tuple with the EndDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDisplay

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetEndDisplay(v string)`

SetEndDisplay sets EndDisplay field to given value.

### HasEndDisplay

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasEndDisplay() bool`

HasEndDisplay returns a boolean if a field has been set.

### GetThreshold

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetThreshold() UpdateInstanceSchedule200ResponseAllOfInstanceScheduleThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetThresholdOk() (*UpdateInstanceSchedule200ResponseAllOfInstanceScheduleThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetThreshold(v UpdateInstanceSchedule200ResponseAllOfInstanceScheduleThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateInstanceSchedule200ResponseAllOfInstanceSchedule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


