# ListPowerSchedules200ResponseAllOfSchedulesInner

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

### NewListPowerSchedules200ResponseAllOfSchedulesInner

`func NewListPowerSchedules200ResponseAllOfSchedulesInner() *ListPowerSchedules200ResponseAllOfSchedulesInner`

NewListPowerSchedules200ResponseAllOfSchedulesInner instantiates a new ListPowerSchedules200ResponseAllOfSchedulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPowerSchedules200ResponseAllOfSchedulesInnerWithDefaults

`func NewListPowerSchedules200ResponseAllOfSchedulesInnerWithDefaults() *ListPowerSchedules200ResponseAllOfSchedulesInner`

NewListPowerSchedules200ResponseAllOfSchedulesInnerWithDefaults instantiates a new ListPowerSchedules200ResponseAllOfSchedulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnabled

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetScheduleType

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetMondayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetMondayOn() int64`

GetMondayOn returns the MondayOn field if non-nil, zero value otherwise.

### GetMondayOnOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetMondayOnOk() (*int64, bool)`

GetMondayOnOk returns a tuple with the MondayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetMondayOn(v int64)`

SetMondayOn sets MondayOn field to given value.

### HasMondayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasMondayOn() bool`

HasMondayOn returns a boolean if a field has been set.

### GetMondayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetMondayOnTime() string`

GetMondayOnTime returns the MondayOnTime field if non-nil, zero value otherwise.

### GetMondayOnTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetMondayOnTimeOk() (*string, bool)`

GetMondayOnTimeOk returns a tuple with the MondayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetMondayOnTime(v string)`

SetMondayOnTime sets MondayOnTime field to given value.

### HasMondayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasMondayOnTime() bool`

HasMondayOnTime returns a boolean if a field has been set.

### GetMondayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetMondayOff() int64`

GetMondayOff returns the MondayOff field if non-nil, zero value otherwise.

### GetMondayOffOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetMondayOffOk() (*int64, bool)`

GetMondayOffOk returns a tuple with the MondayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetMondayOff(v int64)`

SetMondayOff sets MondayOff field to given value.

### HasMondayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasMondayOff() bool`

HasMondayOff returns a boolean if a field has been set.

### GetMondayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetMondayOffTime() string`

GetMondayOffTime returns the MondayOffTime field if non-nil, zero value otherwise.

### GetMondayOffTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetMondayOffTimeOk() (*string, bool)`

GetMondayOffTimeOk returns a tuple with the MondayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetMondayOffTime(v string)`

SetMondayOffTime sets MondayOffTime field to given value.

### HasMondayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasMondayOffTime() bool`

HasMondayOffTime returns a boolean if a field has been set.

### GetTuesdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetTuesdayOn() int64`

GetTuesdayOn returns the TuesdayOn field if non-nil, zero value otherwise.

### GetTuesdayOnOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetTuesdayOnOk() (*int64, bool)`

GetTuesdayOnOk returns a tuple with the TuesdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetTuesdayOn(v int64)`

SetTuesdayOn sets TuesdayOn field to given value.

### HasTuesdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasTuesdayOn() bool`

HasTuesdayOn returns a boolean if a field has been set.

### GetTuesdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetTuesdayOnTime() string`

GetTuesdayOnTime returns the TuesdayOnTime field if non-nil, zero value otherwise.

### GetTuesdayOnTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetTuesdayOnTimeOk() (*string, bool)`

GetTuesdayOnTimeOk returns a tuple with the TuesdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetTuesdayOnTime(v string)`

SetTuesdayOnTime sets TuesdayOnTime field to given value.

### HasTuesdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasTuesdayOnTime() bool`

HasTuesdayOnTime returns a boolean if a field has been set.

### GetTuesdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetTuesdayOff() int64`

GetTuesdayOff returns the TuesdayOff field if non-nil, zero value otherwise.

### GetTuesdayOffOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetTuesdayOffOk() (*int64, bool)`

GetTuesdayOffOk returns a tuple with the TuesdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetTuesdayOff(v int64)`

SetTuesdayOff sets TuesdayOff field to given value.

### HasTuesdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasTuesdayOff() bool`

HasTuesdayOff returns a boolean if a field has been set.

### GetTuesdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetTuesdayOffTime() string`

GetTuesdayOffTime returns the TuesdayOffTime field if non-nil, zero value otherwise.

### GetTuesdayOffTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetTuesdayOffTimeOk() (*string, bool)`

GetTuesdayOffTimeOk returns a tuple with the TuesdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetTuesdayOffTime(v string)`

SetTuesdayOffTime sets TuesdayOffTime field to given value.

### HasTuesdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasTuesdayOffTime() bool`

HasTuesdayOffTime returns a boolean if a field has been set.

### GetWednesdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetWednesdayOn() int64`

GetWednesdayOn returns the WednesdayOn field if non-nil, zero value otherwise.

### GetWednesdayOnOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetWednesdayOnOk() (*int64, bool)`

GetWednesdayOnOk returns a tuple with the WednesdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetWednesdayOn(v int64)`

SetWednesdayOn sets WednesdayOn field to given value.

### HasWednesdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasWednesdayOn() bool`

HasWednesdayOn returns a boolean if a field has been set.

### GetWednesdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetWednesdayOnTime() string`

GetWednesdayOnTime returns the WednesdayOnTime field if non-nil, zero value otherwise.

### GetWednesdayOnTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetWednesdayOnTimeOk() (*string, bool)`

GetWednesdayOnTimeOk returns a tuple with the WednesdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetWednesdayOnTime(v string)`

SetWednesdayOnTime sets WednesdayOnTime field to given value.

### HasWednesdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasWednesdayOnTime() bool`

HasWednesdayOnTime returns a boolean if a field has been set.

### GetWednesdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetWednesdayOff() int64`

GetWednesdayOff returns the WednesdayOff field if non-nil, zero value otherwise.

### GetWednesdayOffOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetWednesdayOffOk() (*int64, bool)`

GetWednesdayOffOk returns a tuple with the WednesdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetWednesdayOff(v int64)`

SetWednesdayOff sets WednesdayOff field to given value.

### HasWednesdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasWednesdayOff() bool`

HasWednesdayOff returns a boolean if a field has been set.

### GetWednesdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetWednesdayOffTime() string`

GetWednesdayOffTime returns the WednesdayOffTime field if non-nil, zero value otherwise.

### GetWednesdayOffTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetWednesdayOffTimeOk() (*string, bool)`

GetWednesdayOffTimeOk returns a tuple with the WednesdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetWednesdayOffTime(v string)`

SetWednesdayOffTime sets WednesdayOffTime field to given value.

### HasWednesdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasWednesdayOffTime() bool`

HasWednesdayOffTime returns a boolean if a field has been set.

### GetThursdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetThursdayOn() int64`

GetThursdayOn returns the ThursdayOn field if non-nil, zero value otherwise.

### GetThursdayOnOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetThursdayOnOk() (*int64, bool)`

GetThursdayOnOk returns a tuple with the ThursdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetThursdayOn(v int64)`

SetThursdayOn sets ThursdayOn field to given value.

### HasThursdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasThursdayOn() bool`

HasThursdayOn returns a boolean if a field has been set.

### GetThursdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetThursdayOnTime() string`

GetThursdayOnTime returns the ThursdayOnTime field if non-nil, zero value otherwise.

### GetThursdayOnTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetThursdayOnTimeOk() (*string, bool)`

GetThursdayOnTimeOk returns a tuple with the ThursdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetThursdayOnTime(v string)`

SetThursdayOnTime sets ThursdayOnTime field to given value.

### HasThursdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasThursdayOnTime() bool`

HasThursdayOnTime returns a boolean if a field has been set.

### GetThursdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetThursdayOff() int64`

GetThursdayOff returns the ThursdayOff field if non-nil, zero value otherwise.

### GetThursdayOffOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetThursdayOffOk() (*int64, bool)`

GetThursdayOffOk returns a tuple with the ThursdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetThursdayOff(v int64)`

SetThursdayOff sets ThursdayOff field to given value.

### HasThursdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasThursdayOff() bool`

HasThursdayOff returns a boolean if a field has been set.

### GetThursdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetThursdayOffTime() string`

GetThursdayOffTime returns the ThursdayOffTime field if non-nil, zero value otherwise.

### GetThursdayOffTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetThursdayOffTimeOk() (*string, bool)`

GetThursdayOffTimeOk returns a tuple with the ThursdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetThursdayOffTime(v string)`

SetThursdayOffTime sets ThursdayOffTime field to given value.

### HasThursdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasThursdayOffTime() bool`

HasThursdayOffTime returns a boolean if a field has been set.

### GetFridayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetFridayOn() int64`

GetFridayOn returns the FridayOn field if non-nil, zero value otherwise.

### GetFridayOnOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetFridayOnOk() (*int64, bool)`

GetFridayOnOk returns a tuple with the FridayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetFridayOn(v int64)`

SetFridayOn sets FridayOn field to given value.

### HasFridayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasFridayOn() bool`

HasFridayOn returns a boolean if a field has been set.

### GetFridayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetFridayOnTime() string`

GetFridayOnTime returns the FridayOnTime field if non-nil, zero value otherwise.

### GetFridayOnTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetFridayOnTimeOk() (*string, bool)`

GetFridayOnTimeOk returns a tuple with the FridayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetFridayOnTime(v string)`

SetFridayOnTime sets FridayOnTime field to given value.

### HasFridayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasFridayOnTime() bool`

HasFridayOnTime returns a boolean if a field has been set.

### GetFridayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetFridayOff() int64`

GetFridayOff returns the FridayOff field if non-nil, zero value otherwise.

### GetFridayOffOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetFridayOffOk() (*int64, bool)`

GetFridayOffOk returns a tuple with the FridayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetFridayOff(v int64)`

SetFridayOff sets FridayOff field to given value.

### HasFridayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasFridayOff() bool`

HasFridayOff returns a boolean if a field has been set.

### GetFridayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetFridayOffTime() string`

GetFridayOffTime returns the FridayOffTime field if non-nil, zero value otherwise.

### GetFridayOffTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetFridayOffTimeOk() (*string, bool)`

GetFridayOffTimeOk returns a tuple with the FridayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetFridayOffTime(v string)`

SetFridayOffTime sets FridayOffTime field to given value.

### HasFridayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasFridayOffTime() bool`

HasFridayOffTime returns a boolean if a field has been set.

### GetSaturdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSaturdayOn() int64`

GetSaturdayOn returns the SaturdayOn field if non-nil, zero value otherwise.

### GetSaturdayOnOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSaturdayOnOk() (*int64, bool)`

GetSaturdayOnOk returns a tuple with the SaturdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetSaturdayOn(v int64)`

SetSaturdayOn sets SaturdayOn field to given value.

### HasSaturdayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasSaturdayOn() bool`

HasSaturdayOn returns a boolean if a field has been set.

### GetSaturdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSaturdayOnTime() string`

GetSaturdayOnTime returns the SaturdayOnTime field if non-nil, zero value otherwise.

### GetSaturdayOnTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSaturdayOnTimeOk() (*string, bool)`

GetSaturdayOnTimeOk returns a tuple with the SaturdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetSaturdayOnTime(v string)`

SetSaturdayOnTime sets SaturdayOnTime field to given value.

### HasSaturdayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasSaturdayOnTime() bool`

HasSaturdayOnTime returns a boolean if a field has been set.

### GetSaturdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSaturdayOff() int64`

GetSaturdayOff returns the SaturdayOff field if non-nil, zero value otherwise.

### GetSaturdayOffOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSaturdayOffOk() (*int64, bool)`

GetSaturdayOffOk returns a tuple with the SaturdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetSaturdayOff(v int64)`

SetSaturdayOff sets SaturdayOff field to given value.

### HasSaturdayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasSaturdayOff() bool`

HasSaturdayOff returns a boolean if a field has been set.

### GetSaturdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSaturdayOffTime() string`

GetSaturdayOffTime returns the SaturdayOffTime field if non-nil, zero value otherwise.

### GetSaturdayOffTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSaturdayOffTimeOk() (*string, bool)`

GetSaturdayOffTimeOk returns a tuple with the SaturdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetSaturdayOffTime(v string)`

SetSaturdayOffTime sets SaturdayOffTime field to given value.

### HasSaturdayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasSaturdayOffTime() bool`

HasSaturdayOffTime returns a boolean if a field has been set.

### GetSundayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSundayOn() int64`

GetSundayOn returns the SundayOn field if non-nil, zero value otherwise.

### GetSundayOnOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSundayOnOk() (*int64, bool)`

GetSundayOnOk returns a tuple with the SundayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetSundayOn(v int64)`

SetSundayOn sets SundayOn field to given value.

### HasSundayOn

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasSundayOn() bool`

HasSundayOn returns a boolean if a field has been set.

### GetSundayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSundayOnTime() string`

GetSundayOnTime returns the SundayOnTime field if non-nil, zero value otherwise.

### GetSundayOnTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSundayOnTimeOk() (*string, bool)`

GetSundayOnTimeOk returns a tuple with the SundayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetSundayOnTime(v string)`

SetSundayOnTime sets SundayOnTime field to given value.

### HasSundayOnTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasSundayOnTime() bool`

HasSundayOnTime returns a boolean if a field has been set.

### GetSundayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSundayOff() int64`

GetSundayOff returns the SundayOff field if non-nil, zero value otherwise.

### GetSundayOffOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSundayOffOk() (*int64, bool)`

GetSundayOffOk returns a tuple with the SundayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetSundayOff(v int64)`

SetSundayOff sets SundayOff field to given value.

### HasSundayOff

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasSundayOff() bool`

HasSundayOff returns a boolean if a field has been set.

### GetSundayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSundayOffTime() string`

GetSundayOffTime returns the SundayOffTime field if non-nil, zero value otherwise.

### GetSundayOffTimeOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetSundayOffTimeOk() (*string, bool)`

GetSundayOffTimeOk returns a tuple with the SundayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetSundayOffTime(v string)`

SetSundayOffTime sets SundayOffTime field to given value.

### HasSundayOffTime

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasSundayOffTime() bool`

HasSundayOffTime returns a boolean if a field has been set.

### GetTotalMonthlyHoursSaved

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetTotalMonthlyHoursSaved() float32`

GetTotalMonthlyHoursSaved returns the TotalMonthlyHoursSaved field if non-nil, zero value otherwise.

### GetTotalMonthlyHoursSavedOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetTotalMonthlyHoursSavedOk() (*float32, bool)`

GetTotalMonthlyHoursSavedOk returns a tuple with the TotalMonthlyHoursSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMonthlyHoursSaved

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetTotalMonthlyHoursSaved(v float32)`

SetTotalMonthlyHoursSaved sets TotalMonthlyHoursSaved field to given value.

### HasTotalMonthlyHoursSaved

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasTotalMonthlyHoursSaved() bool`

HasTotalMonthlyHoursSaved returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListPowerSchedules200ResponseAllOfSchedulesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


