# UpdatePowerSchedulesRequestSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the power schedule | [optional] 
**Description** | Pointer to **string** | A description for the power schedule | [optional] 
**ScheduleType** | Pointer to **string** | Type of schedule &#x60;power&#x60; on or &#x60;power off&#x60; | [optional] 
**ScheduleTimezone** | Pointer to **string** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional] [default to "UTC"]
**Enabled** | Pointer to **bool** | Is the power schedule enabled | [optional] [default to true]
**MondayOnTime** | Pointer to **string** | Monday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**MondayOffTime** | Pointer to **string** | Monday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**TuesdayOnTime** | Pointer to **string** | Tuesday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**TuesdayOffTime** | Pointer to **string** | Tuesday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**WednesdayOnTime** | Pointer to **string** | Wednesday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**WednesdayOffTime** | Pointer to **string** | Wednesday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**ThursdayOnTime** | Pointer to **string** | Thursday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**ThursdayOffTime** | Pointer to **string** | Thursday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**FridayOnTime** | Pointer to **string** | Friday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**FridayOffTime** | Pointer to **string** | Friday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**SaturdayOnTime** | Pointer to **string** | Saturday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**SaturdayOffTime** | Pointer to **string** | Saturday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**SundayOnTime** | Pointer to **string** | Sunday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**SundayOffTime** | Pointer to **string** | Sunday Off time of the day in 24-hour format | [optional] [default to "24:00"]

## Methods

### NewUpdatePowerSchedulesRequestSchedule

`func NewUpdatePowerSchedulesRequestSchedule() *UpdatePowerSchedulesRequestSchedule`

NewUpdatePowerSchedulesRequestSchedule instantiates a new UpdatePowerSchedulesRequestSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePowerSchedulesRequestScheduleWithDefaults

`func NewUpdatePowerSchedulesRequestScheduleWithDefaults() *UpdatePowerSchedulesRequestSchedule`

NewUpdatePowerSchedulesRequestScheduleWithDefaults instantiates a new UpdatePowerSchedulesRequestSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePowerSchedulesRequestSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePowerSchedulesRequestSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePowerSchedulesRequestSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePowerSchedulesRequestSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePowerSchedulesRequestSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePowerSchedulesRequestSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScheduleType

`func (o *UpdatePowerSchedulesRequestSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *UpdatePowerSchedulesRequestSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *UpdatePowerSchedulesRequestSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *UpdatePowerSchedulesRequestSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *UpdatePowerSchedulesRequestSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *UpdatePowerSchedulesRequestSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdatePowerSchedulesRequestSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdatePowerSchedulesRequestSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdatePowerSchedulesRequestSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMondayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetMondayOnTime() string`

GetMondayOnTime returns the MondayOnTime field if non-nil, zero value otherwise.

### GetMondayOnTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetMondayOnTimeOk() (*string, bool)`

GetMondayOnTimeOk returns a tuple with the MondayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetMondayOnTime(v string)`

SetMondayOnTime sets MondayOnTime field to given value.

### HasMondayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasMondayOnTime() bool`

HasMondayOnTime returns a boolean if a field has been set.

### GetMondayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetMondayOffTime() string`

GetMondayOffTime returns the MondayOffTime field if non-nil, zero value otherwise.

### GetMondayOffTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetMondayOffTimeOk() (*string, bool)`

GetMondayOffTimeOk returns a tuple with the MondayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetMondayOffTime(v string)`

SetMondayOffTime sets MondayOffTime field to given value.

### HasMondayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasMondayOffTime() bool`

HasMondayOffTime returns a boolean if a field has been set.

### GetTuesdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetTuesdayOnTime() string`

GetTuesdayOnTime returns the TuesdayOnTime field if non-nil, zero value otherwise.

### GetTuesdayOnTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetTuesdayOnTimeOk() (*string, bool)`

GetTuesdayOnTimeOk returns a tuple with the TuesdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetTuesdayOnTime(v string)`

SetTuesdayOnTime sets TuesdayOnTime field to given value.

### HasTuesdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasTuesdayOnTime() bool`

HasTuesdayOnTime returns a boolean if a field has been set.

### GetTuesdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetTuesdayOffTime() string`

GetTuesdayOffTime returns the TuesdayOffTime field if non-nil, zero value otherwise.

### GetTuesdayOffTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetTuesdayOffTimeOk() (*string, bool)`

GetTuesdayOffTimeOk returns a tuple with the TuesdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetTuesdayOffTime(v string)`

SetTuesdayOffTime sets TuesdayOffTime field to given value.

### HasTuesdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasTuesdayOffTime() bool`

HasTuesdayOffTime returns a boolean if a field has been set.

### GetWednesdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetWednesdayOnTime() string`

GetWednesdayOnTime returns the WednesdayOnTime field if non-nil, zero value otherwise.

### GetWednesdayOnTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetWednesdayOnTimeOk() (*string, bool)`

GetWednesdayOnTimeOk returns a tuple with the WednesdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetWednesdayOnTime(v string)`

SetWednesdayOnTime sets WednesdayOnTime field to given value.

### HasWednesdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasWednesdayOnTime() bool`

HasWednesdayOnTime returns a boolean if a field has been set.

### GetWednesdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetWednesdayOffTime() string`

GetWednesdayOffTime returns the WednesdayOffTime field if non-nil, zero value otherwise.

### GetWednesdayOffTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetWednesdayOffTimeOk() (*string, bool)`

GetWednesdayOffTimeOk returns a tuple with the WednesdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetWednesdayOffTime(v string)`

SetWednesdayOffTime sets WednesdayOffTime field to given value.

### HasWednesdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasWednesdayOffTime() bool`

HasWednesdayOffTime returns a boolean if a field has been set.

### GetThursdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetThursdayOnTime() string`

GetThursdayOnTime returns the ThursdayOnTime field if non-nil, zero value otherwise.

### GetThursdayOnTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetThursdayOnTimeOk() (*string, bool)`

GetThursdayOnTimeOk returns a tuple with the ThursdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetThursdayOnTime(v string)`

SetThursdayOnTime sets ThursdayOnTime field to given value.

### HasThursdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasThursdayOnTime() bool`

HasThursdayOnTime returns a boolean if a field has been set.

### GetThursdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetThursdayOffTime() string`

GetThursdayOffTime returns the ThursdayOffTime field if non-nil, zero value otherwise.

### GetThursdayOffTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetThursdayOffTimeOk() (*string, bool)`

GetThursdayOffTimeOk returns a tuple with the ThursdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetThursdayOffTime(v string)`

SetThursdayOffTime sets ThursdayOffTime field to given value.

### HasThursdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasThursdayOffTime() bool`

HasThursdayOffTime returns a boolean if a field has been set.

### GetFridayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetFridayOnTime() string`

GetFridayOnTime returns the FridayOnTime field if non-nil, zero value otherwise.

### GetFridayOnTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetFridayOnTimeOk() (*string, bool)`

GetFridayOnTimeOk returns a tuple with the FridayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetFridayOnTime(v string)`

SetFridayOnTime sets FridayOnTime field to given value.

### HasFridayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasFridayOnTime() bool`

HasFridayOnTime returns a boolean if a field has been set.

### GetFridayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetFridayOffTime() string`

GetFridayOffTime returns the FridayOffTime field if non-nil, zero value otherwise.

### GetFridayOffTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetFridayOffTimeOk() (*string, bool)`

GetFridayOffTimeOk returns a tuple with the FridayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetFridayOffTime(v string)`

SetFridayOffTime sets FridayOffTime field to given value.

### HasFridayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasFridayOffTime() bool`

HasFridayOffTime returns a boolean if a field has been set.

### GetSaturdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetSaturdayOnTime() string`

GetSaturdayOnTime returns the SaturdayOnTime field if non-nil, zero value otherwise.

### GetSaturdayOnTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetSaturdayOnTimeOk() (*string, bool)`

GetSaturdayOnTimeOk returns a tuple with the SaturdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetSaturdayOnTime(v string)`

SetSaturdayOnTime sets SaturdayOnTime field to given value.

### HasSaturdayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasSaturdayOnTime() bool`

HasSaturdayOnTime returns a boolean if a field has been set.

### GetSaturdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetSaturdayOffTime() string`

GetSaturdayOffTime returns the SaturdayOffTime field if non-nil, zero value otherwise.

### GetSaturdayOffTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetSaturdayOffTimeOk() (*string, bool)`

GetSaturdayOffTimeOk returns a tuple with the SaturdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetSaturdayOffTime(v string)`

SetSaturdayOffTime sets SaturdayOffTime field to given value.

### HasSaturdayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasSaturdayOffTime() bool`

HasSaturdayOffTime returns a boolean if a field has been set.

### GetSundayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetSundayOnTime() string`

GetSundayOnTime returns the SundayOnTime field if non-nil, zero value otherwise.

### GetSundayOnTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetSundayOnTimeOk() (*string, bool)`

GetSundayOnTimeOk returns a tuple with the SundayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetSundayOnTime(v string)`

SetSundayOnTime sets SundayOnTime field to given value.

### HasSundayOnTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasSundayOnTime() bool`

HasSundayOnTime returns a boolean if a field has been set.

### GetSundayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) GetSundayOffTime() string`

GetSundayOffTime returns the SundayOffTime field if non-nil, zero value otherwise.

### GetSundayOffTimeOk

`func (o *UpdatePowerSchedulesRequestSchedule) GetSundayOffTimeOk() (*string, bool)`

GetSundayOffTimeOk returns a tuple with the SundayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) SetSundayOffTime(v string)`

SetSundayOffTime sets SundayOffTime field to given value.

### HasSundayOffTime

`func (o *UpdatePowerSchedulesRequestSchedule) HasSundayOffTime() bool`

HasSundayOffTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


