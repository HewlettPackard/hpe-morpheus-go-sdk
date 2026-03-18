# UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner

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
**Threshold** | Pointer to [**UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInnerThreshold**](UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInnerThreshold.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner

`func NewUpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner() *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner`

NewUpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner instantiates a new UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInnerWithDefaults

`func NewUpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInnerWithDefaults() *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner`

NewUpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInnerWithDefaults instantiates a new UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduleType

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetStartDayOfWeek

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetStartDayOfWeek() int32`

GetStartDayOfWeek returns the StartDayOfWeek field if non-nil, zero value otherwise.

### GetStartDayOfWeekOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetStartDayOfWeekOk() (*int32, bool)`

GetStartDayOfWeekOk returns a tuple with the StartDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDayOfWeek

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetStartDayOfWeek(v int32)`

SetStartDayOfWeek sets StartDayOfWeek field to given value.

### HasStartDayOfWeek

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasStartDayOfWeek() bool`

HasStartDayOfWeek returns a boolean if a field has been set.

### GetStartTime

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndDayOfWeek

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetEndDayOfWeek() int32`

GetEndDayOfWeek returns the EndDayOfWeek field if non-nil, zero value otherwise.

### GetEndDayOfWeekOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetEndDayOfWeekOk() (*int32, bool)`

GetEndDayOfWeekOk returns a tuple with the EndDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDayOfWeek

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetEndDayOfWeek(v int32)`

SetEndDayOfWeek sets EndDayOfWeek field to given value.

### HasEndDayOfWeek

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasEndDayOfWeek() bool`

HasEndDayOfWeek returns a boolean if a field has been set.

### GetEndTime

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDisplay

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetStartDisplay() string`

GetStartDisplay returns the StartDisplay field if non-nil, zero value otherwise.

### GetStartDisplayOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetStartDisplayOk() (*string, bool)`

GetStartDisplayOk returns a tuple with the StartDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDisplay

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetStartDisplay(v string)`

SetStartDisplay sets StartDisplay field to given value.

### HasStartDisplay

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasStartDisplay() bool`

HasStartDisplay returns a boolean if a field has been set.

### GetEndDisplay

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetEndDisplay() string`

GetEndDisplay returns the EndDisplay field if non-nil, zero value otherwise.

### GetEndDisplayOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetEndDisplayOk() (*string, bool)`

GetEndDisplayOk returns a tuple with the EndDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDisplay

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetEndDisplay(v string)`

SetEndDisplay sets EndDisplay field to given value.

### HasEndDisplay

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasEndDisplay() bool`

HasEndDisplay returns a boolean if a field has been set.

### GetThreshold

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetThreshold() UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInnerThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetThresholdOk() (*UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInnerThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetThreshold(v UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInnerThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


