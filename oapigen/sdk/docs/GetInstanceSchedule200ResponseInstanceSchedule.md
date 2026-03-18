# GetInstanceSchedule200ResponseInstanceSchedule

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
**Threshold** | Pointer to [**GetInstanceSchedule200ResponseInstanceScheduleThreshold**](GetInstanceSchedule200ResponseInstanceScheduleThreshold.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetInstanceSchedule200ResponseInstanceSchedule

`func NewGetInstanceSchedule200ResponseInstanceSchedule() *GetInstanceSchedule200ResponseInstanceSchedule`

NewGetInstanceSchedule200ResponseInstanceSchedule instantiates a new GetInstanceSchedule200ResponseInstanceSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceSchedule200ResponseInstanceScheduleWithDefaults

`func NewGetInstanceSchedule200ResponseInstanceScheduleWithDefaults() *GetInstanceSchedule200ResponseInstanceSchedule`

NewGetInstanceSchedule200ResponseInstanceScheduleWithDefaults instantiates a new GetInstanceSchedule200ResponseInstanceSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduleType

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetStartDayOfWeek

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetStartDayOfWeek() int32`

GetStartDayOfWeek returns the StartDayOfWeek field if non-nil, zero value otherwise.

### GetStartDayOfWeekOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetStartDayOfWeekOk() (*int32, bool)`

GetStartDayOfWeekOk returns a tuple with the StartDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDayOfWeek

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetStartDayOfWeek(v int32)`

SetStartDayOfWeek sets StartDayOfWeek field to given value.

### HasStartDayOfWeek

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasStartDayOfWeek() bool`

HasStartDayOfWeek returns a boolean if a field has been set.

### GetStartTime

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndDayOfWeek

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetEndDayOfWeek() int32`

GetEndDayOfWeek returns the EndDayOfWeek field if non-nil, zero value otherwise.

### GetEndDayOfWeekOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetEndDayOfWeekOk() (*int32, bool)`

GetEndDayOfWeekOk returns a tuple with the EndDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDayOfWeek

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetEndDayOfWeek(v int32)`

SetEndDayOfWeek sets EndDayOfWeek field to given value.

### HasEndDayOfWeek

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasEndDayOfWeek() bool`

HasEndDayOfWeek returns a boolean if a field has been set.

### GetEndTime

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartDate

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDisplay

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetStartDisplay() string`

GetStartDisplay returns the StartDisplay field if non-nil, zero value otherwise.

### GetStartDisplayOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetStartDisplayOk() (*string, bool)`

GetStartDisplayOk returns a tuple with the StartDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDisplay

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetStartDisplay(v string)`

SetStartDisplay sets StartDisplay field to given value.

### HasStartDisplay

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasStartDisplay() bool`

HasStartDisplay returns a boolean if a field has been set.

### GetEndDisplay

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetEndDisplay() string`

GetEndDisplay returns the EndDisplay field if non-nil, zero value otherwise.

### GetEndDisplayOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetEndDisplayOk() (*string, bool)`

GetEndDisplayOk returns a tuple with the EndDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDisplay

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetEndDisplay(v string)`

SetEndDisplay sets EndDisplay field to given value.

### HasEndDisplay

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasEndDisplay() bool`

HasEndDisplay returns a boolean if a field has been set.

### GetThreshold

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetThreshold() GetInstanceSchedule200ResponseInstanceScheduleThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetThresholdOk() (*GetInstanceSchedule200ResponseInstanceScheduleThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetThreshold(v GetInstanceSchedule200ResponseInstanceScheduleThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetInstanceSchedule200ResponseInstanceSchedule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


