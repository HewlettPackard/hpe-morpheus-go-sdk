# PowerSchedule

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
**MondayOn** | Pointer to **int64** |  | [optional] 
**MondayOnTime** | Pointer to **string** |  | [optional] 
**MondayOff** | Pointer to **int64** |  | [optional] 
**MondayOffTime** | Pointer to **string** |  | [optional] 
**TuesdayOn** | Pointer to **int64** |  | [optional] 
**TuesdayOnTime** | Pointer to **string** |  | [optional] 
**TuesdayOff** | Pointer to **int64** |  | [optional] 
**TuesdayOffTime** | Pointer to **string** |  | [optional] 
**WednesdayOn** | Pointer to **int64** |  | [optional] 
**WednesdayOnTime** | Pointer to **string** |  | [optional] 
**WednesdayOff** | Pointer to **int64** |  | [optional] 
**WednesdayOffTime** | Pointer to **string** |  | [optional] 
**ThursdayOn** | Pointer to **int64** |  | [optional] 
**ThursdayOnTime** | Pointer to **string** |  | [optional] 
**ThursdayOff** | Pointer to **int64** |  | [optional] 
**ThursdayOffTime** | Pointer to **string** |  | [optional] 
**FridayOn** | Pointer to **int64** |  | [optional] 
**FridayOnTime** | Pointer to **string** |  | [optional] 
**FridayOff** | Pointer to **int64** |  | [optional] 
**FridayOffTime** | Pointer to **string** |  | [optional] 
**SaturdayOn** | Pointer to **int64** |  | [optional] 
**SaturdayOnTime** | Pointer to **string** |  | [optional] 
**SaturdayOff** | Pointer to **int64** |  | [optional] 
**SaturdayOffTime** | Pointer to **string** |  | [optional] 
**SundayOn** | Pointer to **int64** |  | [optional] 
**SundayOnTime** | Pointer to **string** |  | [optional] 
**SundayOff** | Pointer to **int64** |  | [optional] 
**SundayOffTime** | Pointer to **string** |  | [optional] 
**TotalMonthlyHoursSaved** | Pointer to **float32** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPowerSchedule

`func NewPowerSchedule() *PowerSchedule`

NewPowerSchedule instantiates a new PowerSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerScheduleWithDefaults

`func NewPowerScheduleWithDefaults() *PowerSchedule`

NewPowerScheduleWithDefaults instantiates a new PowerSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerSchedule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerSchedule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerSchedule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PowerSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PowerSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PowerSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PowerSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PowerSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PowerSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PowerSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *PowerSchedule) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PowerSchedule) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PowerSchedule) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PowerSchedule) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnabled

`func (o *PowerSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PowerSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PowerSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PowerSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetScheduleType

`func (o *PowerSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *PowerSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *PowerSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *PowerSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *PowerSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *PowerSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *PowerSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *PowerSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetMondayOn

`func (o *PowerSchedule) GetMondayOn() int64`

GetMondayOn returns the MondayOn field if non-nil, zero value otherwise.

### GetMondayOnOk

`func (o *PowerSchedule) GetMondayOnOk() (*int64, bool)`

GetMondayOnOk returns a tuple with the MondayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOn

`func (o *PowerSchedule) SetMondayOn(v int64)`

SetMondayOn sets MondayOn field to given value.

### HasMondayOn

`func (o *PowerSchedule) HasMondayOn() bool`

HasMondayOn returns a boolean if a field has been set.

### GetMondayOnTime

`func (o *PowerSchedule) GetMondayOnTime() string`

GetMondayOnTime returns the MondayOnTime field if non-nil, zero value otherwise.

### GetMondayOnTimeOk

`func (o *PowerSchedule) GetMondayOnTimeOk() (*string, bool)`

GetMondayOnTimeOk returns a tuple with the MondayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOnTime

`func (o *PowerSchedule) SetMondayOnTime(v string)`

SetMondayOnTime sets MondayOnTime field to given value.

### HasMondayOnTime

`func (o *PowerSchedule) HasMondayOnTime() bool`

HasMondayOnTime returns a boolean if a field has been set.

### GetMondayOff

`func (o *PowerSchedule) GetMondayOff() int64`

GetMondayOff returns the MondayOff field if non-nil, zero value otherwise.

### GetMondayOffOk

`func (o *PowerSchedule) GetMondayOffOk() (*int64, bool)`

GetMondayOffOk returns a tuple with the MondayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOff

`func (o *PowerSchedule) SetMondayOff(v int64)`

SetMondayOff sets MondayOff field to given value.

### HasMondayOff

`func (o *PowerSchedule) HasMondayOff() bool`

HasMondayOff returns a boolean if a field has been set.

### GetMondayOffTime

`func (o *PowerSchedule) GetMondayOffTime() string`

GetMondayOffTime returns the MondayOffTime field if non-nil, zero value otherwise.

### GetMondayOffTimeOk

`func (o *PowerSchedule) GetMondayOffTimeOk() (*string, bool)`

GetMondayOffTimeOk returns a tuple with the MondayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOffTime

`func (o *PowerSchedule) SetMondayOffTime(v string)`

SetMondayOffTime sets MondayOffTime field to given value.

### HasMondayOffTime

`func (o *PowerSchedule) HasMondayOffTime() bool`

HasMondayOffTime returns a boolean if a field has been set.

### GetTuesdayOn

`func (o *PowerSchedule) GetTuesdayOn() int64`

GetTuesdayOn returns the TuesdayOn field if non-nil, zero value otherwise.

### GetTuesdayOnOk

`func (o *PowerSchedule) GetTuesdayOnOk() (*int64, bool)`

GetTuesdayOnOk returns a tuple with the TuesdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOn

`func (o *PowerSchedule) SetTuesdayOn(v int64)`

SetTuesdayOn sets TuesdayOn field to given value.

### HasTuesdayOn

`func (o *PowerSchedule) HasTuesdayOn() bool`

HasTuesdayOn returns a boolean if a field has been set.

### GetTuesdayOnTime

`func (o *PowerSchedule) GetTuesdayOnTime() string`

GetTuesdayOnTime returns the TuesdayOnTime field if non-nil, zero value otherwise.

### GetTuesdayOnTimeOk

`func (o *PowerSchedule) GetTuesdayOnTimeOk() (*string, bool)`

GetTuesdayOnTimeOk returns a tuple with the TuesdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOnTime

`func (o *PowerSchedule) SetTuesdayOnTime(v string)`

SetTuesdayOnTime sets TuesdayOnTime field to given value.

### HasTuesdayOnTime

`func (o *PowerSchedule) HasTuesdayOnTime() bool`

HasTuesdayOnTime returns a boolean if a field has been set.

### GetTuesdayOff

`func (o *PowerSchedule) GetTuesdayOff() int64`

GetTuesdayOff returns the TuesdayOff field if non-nil, zero value otherwise.

### GetTuesdayOffOk

`func (o *PowerSchedule) GetTuesdayOffOk() (*int64, bool)`

GetTuesdayOffOk returns a tuple with the TuesdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOff

`func (o *PowerSchedule) SetTuesdayOff(v int64)`

SetTuesdayOff sets TuesdayOff field to given value.

### HasTuesdayOff

`func (o *PowerSchedule) HasTuesdayOff() bool`

HasTuesdayOff returns a boolean if a field has been set.

### GetTuesdayOffTime

`func (o *PowerSchedule) GetTuesdayOffTime() string`

GetTuesdayOffTime returns the TuesdayOffTime field if non-nil, zero value otherwise.

### GetTuesdayOffTimeOk

`func (o *PowerSchedule) GetTuesdayOffTimeOk() (*string, bool)`

GetTuesdayOffTimeOk returns a tuple with the TuesdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOffTime

`func (o *PowerSchedule) SetTuesdayOffTime(v string)`

SetTuesdayOffTime sets TuesdayOffTime field to given value.

### HasTuesdayOffTime

`func (o *PowerSchedule) HasTuesdayOffTime() bool`

HasTuesdayOffTime returns a boolean if a field has been set.

### GetWednesdayOn

`func (o *PowerSchedule) GetWednesdayOn() int64`

GetWednesdayOn returns the WednesdayOn field if non-nil, zero value otherwise.

### GetWednesdayOnOk

`func (o *PowerSchedule) GetWednesdayOnOk() (*int64, bool)`

GetWednesdayOnOk returns a tuple with the WednesdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOn

`func (o *PowerSchedule) SetWednesdayOn(v int64)`

SetWednesdayOn sets WednesdayOn field to given value.

### HasWednesdayOn

`func (o *PowerSchedule) HasWednesdayOn() bool`

HasWednesdayOn returns a boolean if a field has been set.

### GetWednesdayOnTime

`func (o *PowerSchedule) GetWednesdayOnTime() string`

GetWednesdayOnTime returns the WednesdayOnTime field if non-nil, zero value otherwise.

### GetWednesdayOnTimeOk

`func (o *PowerSchedule) GetWednesdayOnTimeOk() (*string, bool)`

GetWednesdayOnTimeOk returns a tuple with the WednesdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOnTime

`func (o *PowerSchedule) SetWednesdayOnTime(v string)`

SetWednesdayOnTime sets WednesdayOnTime field to given value.

### HasWednesdayOnTime

`func (o *PowerSchedule) HasWednesdayOnTime() bool`

HasWednesdayOnTime returns a boolean if a field has been set.

### GetWednesdayOff

`func (o *PowerSchedule) GetWednesdayOff() int64`

GetWednesdayOff returns the WednesdayOff field if non-nil, zero value otherwise.

### GetWednesdayOffOk

`func (o *PowerSchedule) GetWednesdayOffOk() (*int64, bool)`

GetWednesdayOffOk returns a tuple with the WednesdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOff

`func (o *PowerSchedule) SetWednesdayOff(v int64)`

SetWednesdayOff sets WednesdayOff field to given value.

### HasWednesdayOff

`func (o *PowerSchedule) HasWednesdayOff() bool`

HasWednesdayOff returns a boolean if a field has been set.

### GetWednesdayOffTime

`func (o *PowerSchedule) GetWednesdayOffTime() string`

GetWednesdayOffTime returns the WednesdayOffTime field if non-nil, zero value otherwise.

### GetWednesdayOffTimeOk

`func (o *PowerSchedule) GetWednesdayOffTimeOk() (*string, bool)`

GetWednesdayOffTimeOk returns a tuple with the WednesdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOffTime

`func (o *PowerSchedule) SetWednesdayOffTime(v string)`

SetWednesdayOffTime sets WednesdayOffTime field to given value.

### HasWednesdayOffTime

`func (o *PowerSchedule) HasWednesdayOffTime() bool`

HasWednesdayOffTime returns a boolean if a field has been set.

### GetThursdayOn

`func (o *PowerSchedule) GetThursdayOn() int64`

GetThursdayOn returns the ThursdayOn field if non-nil, zero value otherwise.

### GetThursdayOnOk

`func (o *PowerSchedule) GetThursdayOnOk() (*int64, bool)`

GetThursdayOnOk returns a tuple with the ThursdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOn

`func (o *PowerSchedule) SetThursdayOn(v int64)`

SetThursdayOn sets ThursdayOn field to given value.

### HasThursdayOn

`func (o *PowerSchedule) HasThursdayOn() bool`

HasThursdayOn returns a boolean if a field has been set.

### GetThursdayOnTime

`func (o *PowerSchedule) GetThursdayOnTime() string`

GetThursdayOnTime returns the ThursdayOnTime field if non-nil, zero value otherwise.

### GetThursdayOnTimeOk

`func (o *PowerSchedule) GetThursdayOnTimeOk() (*string, bool)`

GetThursdayOnTimeOk returns a tuple with the ThursdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOnTime

`func (o *PowerSchedule) SetThursdayOnTime(v string)`

SetThursdayOnTime sets ThursdayOnTime field to given value.

### HasThursdayOnTime

`func (o *PowerSchedule) HasThursdayOnTime() bool`

HasThursdayOnTime returns a boolean if a field has been set.

### GetThursdayOff

`func (o *PowerSchedule) GetThursdayOff() int64`

GetThursdayOff returns the ThursdayOff field if non-nil, zero value otherwise.

### GetThursdayOffOk

`func (o *PowerSchedule) GetThursdayOffOk() (*int64, bool)`

GetThursdayOffOk returns a tuple with the ThursdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOff

`func (o *PowerSchedule) SetThursdayOff(v int64)`

SetThursdayOff sets ThursdayOff field to given value.

### HasThursdayOff

`func (o *PowerSchedule) HasThursdayOff() bool`

HasThursdayOff returns a boolean if a field has been set.

### GetThursdayOffTime

`func (o *PowerSchedule) GetThursdayOffTime() string`

GetThursdayOffTime returns the ThursdayOffTime field if non-nil, zero value otherwise.

### GetThursdayOffTimeOk

`func (o *PowerSchedule) GetThursdayOffTimeOk() (*string, bool)`

GetThursdayOffTimeOk returns a tuple with the ThursdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOffTime

`func (o *PowerSchedule) SetThursdayOffTime(v string)`

SetThursdayOffTime sets ThursdayOffTime field to given value.

### HasThursdayOffTime

`func (o *PowerSchedule) HasThursdayOffTime() bool`

HasThursdayOffTime returns a boolean if a field has been set.

### GetFridayOn

`func (o *PowerSchedule) GetFridayOn() int64`

GetFridayOn returns the FridayOn field if non-nil, zero value otherwise.

### GetFridayOnOk

`func (o *PowerSchedule) GetFridayOnOk() (*int64, bool)`

GetFridayOnOk returns a tuple with the FridayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOn

`func (o *PowerSchedule) SetFridayOn(v int64)`

SetFridayOn sets FridayOn field to given value.

### HasFridayOn

`func (o *PowerSchedule) HasFridayOn() bool`

HasFridayOn returns a boolean if a field has been set.

### GetFridayOnTime

`func (o *PowerSchedule) GetFridayOnTime() string`

GetFridayOnTime returns the FridayOnTime field if non-nil, zero value otherwise.

### GetFridayOnTimeOk

`func (o *PowerSchedule) GetFridayOnTimeOk() (*string, bool)`

GetFridayOnTimeOk returns a tuple with the FridayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOnTime

`func (o *PowerSchedule) SetFridayOnTime(v string)`

SetFridayOnTime sets FridayOnTime field to given value.

### HasFridayOnTime

`func (o *PowerSchedule) HasFridayOnTime() bool`

HasFridayOnTime returns a boolean if a field has been set.

### GetFridayOff

`func (o *PowerSchedule) GetFridayOff() int64`

GetFridayOff returns the FridayOff field if non-nil, zero value otherwise.

### GetFridayOffOk

`func (o *PowerSchedule) GetFridayOffOk() (*int64, bool)`

GetFridayOffOk returns a tuple with the FridayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOff

`func (o *PowerSchedule) SetFridayOff(v int64)`

SetFridayOff sets FridayOff field to given value.

### HasFridayOff

`func (o *PowerSchedule) HasFridayOff() bool`

HasFridayOff returns a boolean if a field has been set.

### GetFridayOffTime

`func (o *PowerSchedule) GetFridayOffTime() string`

GetFridayOffTime returns the FridayOffTime field if non-nil, zero value otherwise.

### GetFridayOffTimeOk

`func (o *PowerSchedule) GetFridayOffTimeOk() (*string, bool)`

GetFridayOffTimeOk returns a tuple with the FridayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOffTime

`func (o *PowerSchedule) SetFridayOffTime(v string)`

SetFridayOffTime sets FridayOffTime field to given value.

### HasFridayOffTime

`func (o *PowerSchedule) HasFridayOffTime() bool`

HasFridayOffTime returns a boolean if a field has been set.

### GetSaturdayOn

`func (o *PowerSchedule) GetSaturdayOn() int64`

GetSaturdayOn returns the SaturdayOn field if non-nil, zero value otherwise.

### GetSaturdayOnOk

`func (o *PowerSchedule) GetSaturdayOnOk() (*int64, bool)`

GetSaturdayOnOk returns a tuple with the SaturdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOn

`func (o *PowerSchedule) SetSaturdayOn(v int64)`

SetSaturdayOn sets SaturdayOn field to given value.

### HasSaturdayOn

`func (o *PowerSchedule) HasSaturdayOn() bool`

HasSaturdayOn returns a boolean if a field has been set.

### GetSaturdayOnTime

`func (o *PowerSchedule) GetSaturdayOnTime() string`

GetSaturdayOnTime returns the SaturdayOnTime field if non-nil, zero value otherwise.

### GetSaturdayOnTimeOk

`func (o *PowerSchedule) GetSaturdayOnTimeOk() (*string, bool)`

GetSaturdayOnTimeOk returns a tuple with the SaturdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOnTime

`func (o *PowerSchedule) SetSaturdayOnTime(v string)`

SetSaturdayOnTime sets SaturdayOnTime field to given value.

### HasSaturdayOnTime

`func (o *PowerSchedule) HasSaturdayOnTime() bool`

HasSaturdayOnTime returns a boolean if a field has been set.

### GetSaturdayOff

`func (o *PowerSchedule) GetSaturdayOff() int64`

GetSaturdayOff returns the SaturdayOff field if non-nil, zero value otherwise.

### GetSaturdayOffOk

`func (o *PowerSchedule) GetSaturdayOffOk() (*int64, bool)`

GetSaturdayOffOk returns a tuple with the SaturdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOff

`func (o *PowerSchedule) SetSaturdayOff(v int64)`

SetSaturdayOff sets SaturdayOff field to given value.

### HasSaturdayOff

`func (o *PowerSchedule) HasSaturdayOff() bool`

HasSaturdayOff returns a boolean if a field has been set.

### GetSaturdayOffTime

`func (o *PowerSchedule) GetSaturdayOffTime() string`

GetSaturdayOffTime returns the SaturdayOffTime field if non-nil, zero value otherwise.

### GetSaturdayOffTimeOk

`func (o *PowerSchedule) GetSaturdayOffTimeOk() (*string, bool)`

GetSaturdayOffTimeOk returns a tuple with the SaturdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOffTime

`func (o *PowerSchedule) SetSaturdayOffTime(v string)`

SetSaturdayOffTime sets SaturdayOffTime field to given value.

### HasSaturdayOffTime

`func (o *PowerSchedule) HasSaturdayOffTime() bool`

HasSaturdayOffTime returns a boolean if a field has been set.

### GetSundayOn

`func (o *PowerSchedule) GetSundayOn() int64`

GetSundayOn returns the SundayOn field if non-nil, zero value otherwise.

### GetSundayOnOk

`func (o *PowerSchedule) GetSundayOnOk() (*int64, bool)`

GetSundayOnOk returns a tuple with the SundayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOn

`func (o *PowerSchedule) SetSundayOn(v int64)`

SetSundayOn sets SundayOn field to given value.

### HasSundayOn

`func (o *PowerSchedule) HasSundayOn() bool`

HasSundayOn returns a boolean if a field has been set.

### GetSundayOnTime

`func (o *PowerSchedule) GetSundayOnTime() string`

GetSundayOnTime returns the SundayOnTime field if non-nil, zero value otherwise.

### GetSundayOnTimeOk

`func (o *PowerSchedule) GetSundayOnTimeOk() (*string, bool)`

GetSundayOnTimeOk returns a tuple with the SundayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOnTime

`func (o *PowerSchedule) SetSundayOnTime(v string)`

SetSundayOnTime sets SundayOnTime field to given value.

### HasSundayOnTime

`func (o *PowerSchedule) HasSundayOnTime() bool`

HasSundayOnTime returns a boolean if a field has been set.

### GetSundayOff

`func (o *PowerSchedule) GetSundayOff() int64`

GetSundayOff returns the SundayOff field if non-nil, zero value otherwise.

### GetSundayOffOk

`func (o *PowerSchedule) GetSundayOffOk() (*int64, bool)`

GetSundayOffOk returns a tuple with the SundayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOff

`func (o *PowerSchedule) SetSundayOff(v int64)`

SetSundayOff sets SundayOff field to given value.

### HasSundayOff

`func (o *PowerSchedule) HasSundayOff() bool`

HasSundayOff returns a boolean if a field has been set.

### GetSundayOffTime

`func (o *PowerSchedule) GetSundayOffTime() string`

GetSundayOffTime returns the SundayOffTime field if non-nil, zero value otherwise.

### GetSundayOffTimeOk

`func (o *PowerSchedule) GetSundayOffTimeOk() (*string, bool)`

GetSundayOffTimeOk returns a tuple with the SundayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOffTime

`func (o *PowerSchedule) SetSundayOffTime(v string)`

SetSundayOffTime sets SundayOffTime field to given value.

### HasSundayOffTime

`func (o *PowerSchedule) HasSundayOffTime() bool`

HasSundayOffTime returns a boolean if a field has been set.

### GetTotalMonthlyHoursSaved

`func (o *PowerSchedule) GetTotalMonthlyHoursSaved() float32`

GetTotalMonthlyHoursSaved returns the TotalMonthlyHoursSaved field if non-nil, zero value otherwise.

### GetTotalMonthlyHoursSavedOk

`func (o *PowerSchedule) GetTotalMonthlyHoursSavedOk() (*float32, bool)`

GetTotalMonthlyHoursSavedOk returns a tuple with the TotalMonthlyHoursSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMonthlyHoursSaved

`func (o *PowerSchedule) SetTotalMonthlyHoursSaved(v float32)`

SetTotalMonthlyHoursSaved sets TotalMonthlyHoursSaved field to given value.

### HasTotalMonthlyHoursSaved

`func (o *PowerSchedule) HasTotalMonthlyHoursSaved() bool`

HasTotalMonthlyHoursSaved returns a boolean if a field has been set.

### GetDateCreated

`func (o *PowerSchedule) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *PowerSchedule) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *PowerSchedule) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *PowerSchedule) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PowerSchedule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PowerSchedule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PowerSchedule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PowerSchedule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


