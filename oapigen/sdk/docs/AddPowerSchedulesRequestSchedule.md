# AddPowerSchedulesRequestSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the power schedule | 
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

### NewAddPowerSchedulesRequestSchedule

`func NewAddPowerSchedulesRequestSchedule(name string, ) *AddPowerSchedulesRequestSchedule`

NewAddPowerSchedulesRequestSchedule instantiates a new AddPowerSchedulesRequestSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPowerSchedulesRequestScheduleWithDefaults

`func NewAddPowerSchedulesRequestScheduleWithDefaults() *AddPowerSchedulesRequestSchedule`

NewAddPowerSchedulesRequestScheduleWithDefaults instantiates a new AddPowerSchedulesRequestSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddPowerSchedulesRequestSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddPowerSchedulesRequestSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddPowerSchedulesRequestSchedule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddPowerSchedulesRequestSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPowerSchedulesRequestSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPowerSchedulesRequestSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPowerSchedulesRequestSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScheduleType

`func (o *AddPowerSchedulesRequestSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *AddPowerSchedulesRequestSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *AddPowerSchedulesRequestSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *AddPowerSchedulesRequestSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *AddPowerSchedulesRequestSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *AddPowerSchedulesRequestSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *AddPowerSchedulesRequestSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *AddPowerSchedulesRequestSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPowerSchedulesRequestSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPowerSchedulesRequestSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPowerSchedulesRequestSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddPowerSchedulesRequestSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMondayOnTime

`func (o *AddPowerSchedulesRequestSchedule) GetMondayOnTime() string`

GetMondayOnTime returns the MondayOnTime field if non-nil, zero value otherwise.

### GetMondayOnTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetMondayOnTimeOk() (*string, bool)`

GetMondayOnTimeOk returns a tuple with the MondayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOnTime

`func (o *AddPowerSchedulesRequestSchedule) SetMondayOnTime(v string)`

SetMondayOnTime sets MondayOnTime field to given value.

### HasMondayOnTime

`func (o *AddPowerSchedulesRequestSchedule) HasMondayOnTime() bool`

HasMondayOnTime returns a boolean if a field has been set.

### GetMondayOffTime

`func (o *AddPowerSchedulesRequestSchedule) GetMondayOffTime() string`

GetMondayOffTime returns the MondayOffTime field if non-nil, zero value otherwise.

### GetMondayOffTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetMondayOffTimeOk() (*string, bool)`

GetMondayOffTimeOk returns a tuple with the MondayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOffTime

`func (o *AddPowerSchedulesRequestSchedule) SetMondayOffTime(v string)`

SetMondayOffTime sets MondayOffTime field to given value.

### HasMondayOffTime

`func (o *AddPowerSchedulesRequestSchedule) HasMondayOffTime() bool`

HasMondayOffTime returns a boolean if a field has been set.

### GetTuesdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) GetTuesdayOnTime() string`

GetTuesdayOnTime returns the TuesdayOnTime field if non-nil, zero value otherwise.

### GetTuesdayOnTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetTuesdayOnTimeOk() (*string, bool)`

GetTuesdayOnTimeOk returns a tuple with the TuesdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) SetTuesdayOnTime(v string)`

SetTuesdayOnTime sets TuesdayOnTime field to given value.

### HasTuesdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) HasTuesdayOnTime() bool`

HasTuesdayOnTime returns a boolean if a field has been set.

### GetTuesdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) GetTuesdayOffTime() string`

GetTuesdayOffTime returns the TuesdayOffTime field if non-nil, zero value otherwise.

### GetTuesdayOffTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetTuesdayOffTimeOk() (*string, bool)`

GetTuesdayOffTimeOk returns a tuple with the TuesdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) SetTuesdayOffTime(v string)`

SetTuesdayOffTime sets TuesdayOffTime field to given value.

### HasTuesdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) HasTuesdayOffTime() bool`

HasTuesdayOffTime returns a boolean if a field has been set.

### GetWednesdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) GetWednesdayOnTime() string`

GetWednesdayOnTime returns the WednesdayOnTime field if non-nil, zero value otherwise.

### GetWednesdayOnTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetWednesdayOnTimeOk() (*string, bool)`

GetWednesdayOnTimeOk returns a tuple with the WednesdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) SetWednesdayOnTime(v string)`

SetWednesdayOnTime sets WednesdayOnTime field to given value.

### HasWednesdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) HasWednesdayOnTime() bool`

HasWednesdayOnTime returns a boolean if a field has been set.

### GetWednesdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) GetWednesdayOffTime() string`

GetWednesdayOffTime returns the WednesdayOffTime field if non-nil, zero value otherwise.

### GetWednesdayOffTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetWednesdayOffTimeOk() (*string, bool)`

GetWednesdayOffTimeOk returns a tuple with the WednesdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) SetWednesdayOffTime(v string)`

SetWednesdayOffTime sets WednesdayOffTime field to given value.

### HasWednesdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) HasWednesdayOffTime() bool`

HasWednesdayOffTime returns a boolean if a field has been set.

### GetThursdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) GetThursdayOnTime() string`

GetThursdayOnTime returns the ThursdayOnTime field if non-nil, zero value otherwise.

### GetThursdayOnTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetThursdayOnTimeOk() (*string, bool)`

GetThursdayOnTimeOk returns a tuple with the ThursdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) SetThursdayOnTime(v string)`

SetThursdayOnTime sets ThursdayOnTime field to given value.

### HasThursdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) HasThursdayOnTime() bool`

HasThursdayOnTime returns a boolean if a field has been set.

### GetThursdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) GetThursdayOffTime() string`

GetThursdayOffTime returns the ThursdayOffTime field if non-nil, zero value otherwise.

### GetThursdayOffTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetThursdayOffTimeOk() (*string, bool)`

GetThursdayOffTimeOk returns a tuple with the ThursdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) SetThursdayOffTime(v string)`

SetThursdayOffTime sets ThursdayOffTime field to given value.

### HasThursdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) HasThursdayOffTime() bool`

HasThursdayOffTime returns a boolean if a field has been set.

### GetFridayOnTime

`func (o *AddPowerSchedulesRequestSchedule) GetFridayOnTime() string`

GetFridayOnTime returns the FridayOnTime field if non-nil, zero value otherwise.

### GetFridayOnTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetFridayOnTimeOk() (*string, bool)`

GetFridayOnTimeOk returns a tuple with the FridayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOnTime

`func (o *AddPowerSchedulesRequestSchedule) SetFridayOnTime(v string)`

SetFridayOnTime sets FridayOnTime field to given value.

### HasFridayOnTime

`func (o *AddPowerSchedulesRequestSchedule) HasFridayOnTime() bool`

HasFridayOnTime returns a boolean if a field has been set.

### GetFridayOffTime

`func (o *AddPowerSchedulesRequestSchedule) GetFridayOffTime() string`

GetFridayOffTime returns the FridayOffTime field if non-nil, zero value otherwise.

### GetFridayOffTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetFridayOffTimeOk() (*string, bool)`

GetFridayOffTimeOk returns a tuple with the FridayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOffTime

`func (o *AddPowerSchedulesRequestSchedule) SetFridayOffTime(v string)`

SetFridayOffTime sets FridayOffTime field to given value.

### HasFridayOffTime

`func (o *AddPowerSchedulesRequestSchedule) HasFridayOffTime() bool`

HasFridayOffTime returns a boolean if a field has been set.

### GetSaturdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) GetSaturdayOnTime() string`

GetSaturdayOnTime returns the SaturdayOnTime field if non-nil, zero value otherwise.

### GetSaturdayOnTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetSaturdayOnTimeOk() (*string, bool)`

GetSaturdayOnTimeOk returns a tuple with the SaturdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) SetSaturdayOnTime(v string)`

SetSaturdayOnTime sets SaturdayOnTime field to given value.

### HasSaturdayOnTime

`func (o *AddPowerSchedulesRequestSchedule) HasSaturdayOnTime() bool`

HasSaturdayOnTime returns a boolean if a field has been set.

### GetSaturdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) GetSaturdayOffTime() string`

GetSaturdayOffTime returns the SaturdayOffTime field if non-nil, zero value otherwise.

### GetSaturdayOffTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetSaturdayOffTimeOk() (*string, bool)`

GetSaturdayOffTimeOk returns a tuple with the SaturdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) SetSaturdayOffTime(v string)`

SetSaturdayOffTime sets SaturdayOffTime field to given value.

### HasSaturdayOffTime

`func (o *AddPowerSchedulesRequestSchedule) HasSaturdayOffTime() bool`

HasSaturdayOffTime returns a boolean if a field has been set.

### GetSundayOnTime

`func (o *AddPowerSchedulesRequestSchedule) GetSundayOnTime() string`

GetSundayOnTime returns the SundayOnTime field if non-nil, zero value otherwise.

### GetSundayOnTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetSundayOnTimeOk() (*string, bool)`

GetSundayOnTimeOk returns a tuple with the SundayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOnTime

`func (o *AddPowerSchedulesRequestSchedule) SetSundayOnTime(v string)`

SetSundayOnTime sets SundayOnTime field to given value.

### HasSundayOnTime

`func (o *AddPowerSchedulesRequestSchedule) HasSundayOnTime() bool`

HasSundayOnTime returns a boolean if a field has been set.

### GetSundayOffTime

`func (o *AddPowerSchedulesRequestSchedule) GetSundayOffTime() string`

GetSundayOffTime returns the SundayOffTime field if non-nil, zero value otherwise.

### GetSundayOffTimeOk

`func (o *AddPowerSchedulesRequestSchedule) GetSundayOffTimeOk() (*string, bool)`

GetSundayOffTimeOk returns a tuple with the SundayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOffTime

`func (o *AddPowerSchedulesRequestSchedule) SetSundayOffTime(v string)`

SetSundayOffTime sets SundayOffTime field to given value.

### HasSundayOffTime

`func (o *AddPowerSchedulesRequestSchedule) HasSundayOffTime() bool`

HasSundayOffTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


