# AddExecuteSchedulesRequestSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the execute schedule | 
**Description** | Pointer to **string** | A description for the execute schedule | [optional] 
**ScheduleType** | Pointer to **string** | Type of schedule | [optional] 
**ScheduleTimezone** | Pointer to **string** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional] [default to "UTC"]
**Cron** | Pointer to **string** | Cron Expression. The default is daily at midnight | [optional] [default to "0 0 * * *"]
**Enabled** | Pointer to **bool** | Is enabled | [optional] [default to true]

## Methods

### NewAddExecuteSchedulesRequestSchedule

`func NewAddExecuteSchedulesRequestSchedule(name string, ) *AddExecuteSchedulesRequestSchedule`

NewAddExecuteSchedulesRequestSchedule instantiates a new AddExecuteSchedulesRequestSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddExecuteSchedulesRequestScheduleWithDefaults

`func NewAddExecuteSchedulesRequestScheduleWithDefaults() *AddExecuteSchedulesRequestSchedule`

NewAddExecuteSchedulesRequestScheduleWithDefaults instantiates a new AddExecuteSchedulesRequestSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddExecuteSchedulesRequestSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddExecuteSchedulesRequestSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddExecuteSchedulesRequestSchedule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddExecuteSchedulesRequestSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddExecuteSchedulesRequestSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddExecuteSchedulesRequestSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddExecuteSchedulesRequestSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetScheduleType

`func (o *AddExecuteSchedulesRequestSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *AddExecuteSchedulesRequestSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *AddExecuteSchedulesRequestSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *AddExecuteSchedulesRequestSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *AddExecuteSchedulesRequestSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *AddExecuteSchedulesRequestSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *AddExecuteSchedulesRequestSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *AddExecuteSchedulesRequestSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### GetCron

`func (o *AddExecuteSchedulesRequestSchedule) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *AddExecuteSchedulesRequestSchedule) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *AddExecuteSchedulesRequestSchedule) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *AddExecuteSchedulesRequestSchedule) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetEnabled

`func (o *AddExecuteSchedulesRequestSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddExecuteSchedulesRequestSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddExecuteSchedulesRequestSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddExecuteSchedulesRequestSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


