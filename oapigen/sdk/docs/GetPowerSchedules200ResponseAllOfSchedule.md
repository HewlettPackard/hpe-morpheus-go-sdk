# GetPowerSchedules200ResponseAllOfSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
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

### NewGetPowerSchedules200ResponseAllOfSchedule

`func NewGetPowerSchedules200ResponseAllOfSchedule() *GetPowerSchedules200ResponseAllOfSchedule`

NewGetPowerSchedules200ResponseAllOfSchedule instantiates a new GetPowerSchedules200ResponseAllOfSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPowerSchedules200ResponseAllOfScheduleWithDefaults

`func NewGetPowerSchedules200ResponseAllOfScheduleWithDefaults() *GetPowerSchedules200ResponseAllOfSchedule`

NewGetPowerSchedules200ResponseAllOfScheduleWithDefaults instantiates a new GetPowerSchedules200ResponseAllOfSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetPowerSchedules200ResponseAllOfSchedule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVisibility

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnabled

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetScheduleType

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetMondayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetMondayOn() int64`

GetMondayOn returns the MondayOn field if non-nil, zero value otherwise.

### GetMondayOnOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetMondayOnOk() (*int64, bool)`

GetMondayOnOk returns a tuple with the MondayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetMondayOn(v int64)`

SetMondayOn sets MondayOn field to given value.

### HasMondayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasMondayOn() bool`

HasMondayOn returns a boolean if a field has been set.

### GetMondayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetMondayOnTime() string`

GetMondayOnTime returns the MondayOnTime field if non-nil, zero value otherwise.

### GetMondayOnTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetMondayOnTimeOk() (*string, bool)`

GetMondayOnTimeOk returns a tuple with the MondayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetMondayOnTime(v string)`

SetMondayOnTime sets MondayOnTime field to given value.

### HasMondayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasMondayOnTime() bool`

HasMondayOnTime returns a boolean if a field has been set.

### GetMondayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetMondayOff() int64`

GetMondayOff returns the MondayOff field if non-nil, zero value otherwise.

### GetMondayOffOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetMondayOffOk() (*int64, bool)`

GetMondayOffOk returns a tuple with the MondayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetMondayOff(v int64)`

SetMondayOff sets MondayOff field to given value.

### HasMondayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasMondayOff() bool`

HasMondayOff returns a boolean if a field has been set.

### GetMondayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetMondayOffTime() string`

GetMondayOffTime returns the MondayOffTime field if non-nil, zero value otherwise.

### GetMondayOffTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetMondayOffTimeOk() (*string, bool)`

GetMondayOffTimeOk returns a tuple with the MondayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMondayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetMondayOffTime(v string)`

SetMondayOffTime sets MondayOffTime field to given value.

### HasMondayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasMondayOffTime() bool`

HasMondayOffTime returns a boolean if a field has been set.

### GetTuesdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetTuesdayOn() int64`

GetTuesdayOn returns the TuesdayOn field if non-nil, zero value otherwise.

### GetTuesdayOnOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetTuesdayOnOk() (*int64, bool)`

GetTuesdayOnOk returns a tuple with the TuesdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetTuesdayOn(v int64)`

SetTuesdayOn sets TuesdayOn field to given value.

### HasTuesdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasTuesdayOn() bool`

HasTuesdayOn returns a boolean if a field has been set.

### GetTuesdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetTuesdayOnTime() string`

GetTuesdayOnTime returns the TuesdayOnTime field if non-nil, zero value otherwise.

### GetTuesdayOnTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetTuesdayOnTimeOk() (*string, bool)`

GetTuesdayOnTimeOk returns a tuple with the TuesdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetTuesdayOnTime(v string)`

SetTuesdayOnTime sets TuesdayOnTime field to given value.

### HasTuesdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasTuesdayOnTime() bool`

HasTuesdayOnTime returns a boolean if a field has been set.

### GetTuesdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetTuesdayOff() int64`

GetTuesdayOff returns the TuesdayOff field if non-nil, zero value otherwise.

### GetTuesdayOffOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetTuesdayOffOk() (*int64, bool)`

GetTuesdayOffOk returns a tuple with the TuesdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetTuesdayOff(v int64)`

SetTuesdayOff sets TuesdayOff field to given value.

### HasTuesdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasTuesdayOff() bool`

HasTuesdayOff returns a boolean if a field has been set.

### GetTuesdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetTuesdayOffTime() string`

GetTuesdayOffTime returns the TuesdayOffTime field if non-nil, zero value otherwise.

### GetTuesdayOffTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetTuesdayOffTimeOk() (*string, bool)`

GetTuesdayOffTimeOk returns a tuple with the TuesdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetTuesdayOffTime(v string)`

SetTuesdayOffTime sets TuesdayOffTime field to given value.

### HasTuesdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasTuesdayOffTime() bool`

HasTuesdayOffTime returns a boolean if a field has been set.

### GetWednesdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetWednesdayOn() int64`

GetWednesdayOn returns the WednesdayOn field if non-nil, zero value otherwise.

### GetWednesdayOnOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetWednesdayOnOk() (*int64, bool)`

GetWednesdayOnOk returns a tuple with the WednesdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetWednesdayOn(v int64)`

SetWednesdayOn sets WednesdayOn field to given value.

### HasWednesdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasWednesdayOn() bool`

HasWednesdayOn returns a boolean if a field has been set.

### GetWednesdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetWednesdayOnTime() string`

GetWednesdayOnTime returns the WednesdayOnTime field if non-nil, zero value otherwise.

### GetWednesdayOnTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetWednesdayOnTimeOk() (*string, bool)`

GetWednesdayOnTimeOk returns a tuple with the WednesdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetWednesdayOnTime(v string)`

SetWednesdayOnTime sets WednesdayOnTime field to given value.

### HasWednesdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasWednesdayOnTime() bool`

HasWednesdayOnTime returns a boolean if a field has been set.

### GetWednesdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetWednesdayOff() int64`

GetWednesdayOff returns the WednesdayOff field if non-nil, zero value otherwise.

### GetWednesdayOffOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetWednesdayOffOk() (*int64, bool)`

GetWednesdayOffOk returns a tuple with the WednesdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetWednesdayOff(v int64)`

SetWednesdayOff sets WednesdayOff field to given value.

### HasWednesdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasWednesdayOff() bool`

HasWednesdayOff returns a boolean if a field has been set.

### GetWednesdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetWednesdayOffTime() string`

GetWednesdayOffTime returns the WednesdayOffTime field if non-nil, zero value otherwise.

### GetWednesdayOffTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetWednesdayOffTimeOk() (*string, bool)`

GetWednesdayOffTimeOk returns a tuple with the WednesdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetWednesdayOffTime(v string)`

SetWednesdayOffTime sets WednesdayOffTime field to given value.

### HasWednesdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasWednesdayOffTime() bool`

HasWednesdayOffTime returns a boolean if a field has been set.

### GetThursdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetThursdayOn() int64`

GetThursdayOn returns the ThursdayOn field if non-nil, zero value otherwise.

### GetThursdayOnOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetThursdayOnOk() (*int64, bool)`

GetThursdayOnOk returns a tuple with the ThursdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetThursdayOn(v int64)`

SetThursdayOn sets ThursdayOn field to given value.

### HasThursdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasThursdayOn() bool`

HasThursdayOn returns a boolean if a field has been set.

### GetThursdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetThursdayOnTime() string`

GetThursdayOnTime returns the ThursdayOnTime field if non-nil, zero value otherwise.

### GetThursdayOnTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetThursdayOnTimeOk() (*string, bool)`

GetThursdayOnTimeOk returns a tuple with the ThursdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetThursdayOnTime(v string)`

SetThursdayOnTime sets ThursdayOnTime field to given value.

### HasThursdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasThursdayOnTime() bool`

HasThursdayOnTime returns a boolean if a field has been set.

### GetThursdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetThursdayOff() int64`

GetThursdayOff returns the ThursdayOff field if non-nil, zero value otherwise.

### GetThursdayOffOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetThursdayOffOk() (*int64, bool)`

GetThursdayOffOk returns a tuple with the ThursdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetThursdayOff(v int64)`

SetThursdayOff sets ThursdayOff field to given value.

### HasThursdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasThursdayOff() bool`

HasThursdayOff returns a boolean if a field has been set.

### GetThursdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetThursdayOffTime() string`

GetThursdayOffTime returns the ThursdayOffTime field if non-nil, zero value otherwise.

### GetThursdayOffTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetThursdayOffTimeOk() (*string, bool)`

GetThursdayOffTimeOk returns a tuple with the ThursdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetThursdayOffTime(v string)`

SetThursdayOffTime sets ThursdayOffTime field to given value.

### HasThursdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasThursdayOffTime() bool`

HasThursdayOffTime returns a boolean if a field has been set.

### GetFridayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetFridayOn() int64`

GetFridayOn returns the FridayOn field if non-nil, zero value otherwise.

### GetFridayOnOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetFridayOnOk() (*int64, bool)`

GetFridayOnOk returns a tuple with the FridayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetFridayOn(v int64)`

SetFridayOn sets FridayOn field to given value.

### HasFridayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasFridayOn() bool`

HasFridayOn returns a boolean if a field has been set.

### GetFridayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetFridayOnTime() string`

GetFridayOnTime returns the FridayOnTime field if non-nil, zero value otherwise.

### GetFridayOnTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetFridayOnTimeOk() (*string, bool)`

GetFridayOnTimeOk returns a tuple with the FridayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetFridayOnTime(v string)`

SetFridayOnTime sets FridayOnTime field to given value.

### HasFridayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasFridayOnTime() bool`

HasFridayOnTime returns a boolean if a field has been set.

### GetFridayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetFridayOff() int64`

GetFridayOff returns the FridayOff field if non-nil, zero value otherwise.

### GetFridayOffOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetFridayOffOk() (*int64, bool)`

GetFridayOffOk returns a tuple with the FridayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetFridayOff(v int64)`

SetFridayOff sets FridayOff field to given value.

### HasFridayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasFridayOff() bool`

HasFridayOff returns a boolean if a field has been set.

### GetFridayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetFridayOffTime() string`

GetFridayOffTime returns the FridayOffTime field if non-nil, zero value otherwise.

### GetFridayOffTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetFridayOffTimeOk() (*string, bool)`

GetFridayOffTimeOk returns a tuple with the FridayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFridayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetFridayOffTime(v string)`

SetFridayOffTime sets FridayOffTime field to given value.

### HasFridayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasFridayOffTime() bool`

HasFridayOffTime returns a boolean if a field has been set.

### GetSaturdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSaturdayOn() int64`

GetSaturdayOn returns the SaturdayOn field if non-nil, zero value otherwise.

### GetSaturdayOnOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSaturdayOnOk() (*int64, bool)`

GetSaturdayOnOk returns a tuple with the SaturdayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetSaturdayOn(v int64)`

SetSaturdayOn sets SaturdayOn field to given value.

### HasSaturdayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasSaturdayOn() bool`

HasSaturdayOn returns a boolean if a field has been set.

### GetSaturdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSaturdayOnTime() string`

GetSaturdayOnTime returns the SaturdayOnTime field if non-nil, zero value otherwise.

### GetSaturdayOnTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSaturdayOnTimeOk() (*string, bool)`

GetSaturdayOnTimeOk returns a tuple with the SaturdayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetSaturdayOnTime(v string)`

SetSaturdayOnTime sets SaturdayOnTime field to given value.

### HasSaturdayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasSaturdayOnTime() bool`

HasSaturdayOnTime returns a boolean if a field has been set.

### GetSaturdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSaturdayOff() int64`

GetSaturdayOff returns the SaturdayOff field if non-nil, zero value otherwise.

### GetSaturdayOffOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSaturdayOffOk() (*int64, bool)`

GetSaturdayOffOk returns a tuple with the SaturdayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetSaturdayOff(v int64)`

SetSaturdayOff sets SaturdayOff field to given value.

### HasSaturdayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasSaturdayOff() bool`

HasSaturdayOff returns a boolean if a field has been set.

### GetSaturdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSaturdayOffTime() string`

GetSaturdayOffTime returns the SaturdayOffTime field if non-nil, zero value otherwise.

### GetSaturdayOffTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSaturdayOffTimeOk() (*string, bool)`

GetSaturdayOffTimeOk returns a tuple with the SaturdayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetSaturdayOffTime(v string)`

SetSaturdayOffTime sets SaturdayOffTime field to given value.

### HasSaturdayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasSaturdayOffTime() bool`

HasSaturdayOffTime returns a boolean if a field has been set.

### GetSundayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSundayOn() int64`

GetSundayOn returns the SundayOn field if non-nil, zero value otherwise.

### GetSundayOnOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSundayOnOk() (*int64, bool)`

GetSundayOnOk returns a tuple with the SundayOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetSundayOn(v int64)`

SetSundayOn sets SundayOn field to given value.

### HasSundayOn

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasSundayOn() bool`

HasSundayOn returns a boolean if a field has been set.

### GetSundayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSundayOnTime() string`

GetSundayOnTime returns the SundayOnTime field if non-nil, zero value otherwise.

### GetSundayOnTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSundayOnTimeOk() (*string, bool)`

GetSundayOnTimeOk returns a tuple with the SundayOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetSundayOnTime(v string)`

SetSundayOnTime sets SundayOnTime field to given value.

### HasSundayOnTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasSundayOnTime() bool`

HasSundayOnTime returns a boolean if a field has been set.

### GetSundayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSundayOff() int64`

GetSundayOff returns the SundayOff field if non-nil, zero value otherwise.

### GetSundayOffOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSundayOffOk() (*int64, bool)`

GetSundayOffOk returns a tuple with the SundayOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetSundayOff(v int64)`

SetSundayOff sets SundayOff field to given value.

### HasSundayOff

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasSundayOff() bool`

HasSundayOff returns a boolean if a field has been set.

### GetSundayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSundayOffTime() string`

GetSundayOffTime returns the SundayOffTime field if non-nil, zero value otherwise.

### GetSundayOffTimeOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetSundayOffTimeOk() (*string, bool)`

GetSundayOffTimeOk returns a tuple with the SundayOffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSundayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetSundayOffTime(v string)`

SetSundayOffTime sets SundayOffTime field to given value.

### HasSundayOffTime

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasSundayOffTime() bool`

HasSundayOffTime returns a boolean if a field has been set.

### GetTotalMonthlyHoursSaved

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetTotalMonthlyHoursSaved() float32`

GetTotalMonthlyHoursSaved returns the TotalMonthlyHoursSaved field if non-nil, zero value otherwise.

### GetTotalMonthlyHoursSavedOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetTotalMonthlyHoursSavedOk() (*float32, bool)`

GetTotalMonthlyHoursSavedOk returns a tuple with the TotalMonthlyHoursSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMonthlyHoursSaved

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetTotalMonthlyHoursSaved(v float32)`

SetTotalMonthlyHoursSaved sets TotalMonthlyHoursSaved field to given value.

### HasTotalMonthlyHoursSaved

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasTotalMonthlyHoursSaved() bool`

HasTotalMonthlyHoursSaved returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetPowerSchedules200ResponseAllOfSchedule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetPowerSchedules200ResponseAllOfSchedule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetPowerSchedules200ResponseAllOfSchedule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


