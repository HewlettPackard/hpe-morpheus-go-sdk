# GetExecuteSchedules200ResponseAllOfSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ScheduleType** | Pointer to **string** |  | [optional] 
**ScheduleTimezone** | Pointer to **NullableString** |  | [optional] 
**Cron** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetExecuteSchedules200ResponseAllOfSchedule

`func NewGetExecuteSchedules200ResponseAllOfSchedule() *GetExecuteSchedules200ResponseAllOfSchedule`

NewGetExecuteSchedules200ResponseAllOfSchedule instantiates a new GetExecuteSchedules200ResponseAllOfSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetExecuteSchedules200ResponseAllOfSchedule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVisibility

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *GetExecuteSchedules200ResponseAllOfSchedule) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetEnabled

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetScheduleType

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### SetScheduleTimezoneNil

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetScheduleTimezoneNil(b bool)`

 SetScheduleTimezoneNil sets the value for ScheduleTimezone to be an explicit nil

### UnsetScheduleTimezone
`func (o *GetExecuteSchedules200ResponseAllOfSchedule) UnsetScheduleTimezone()`

UnsetScheduleTimezone ensures that no value is present for ScheduleTimezone, not even an explicit nil
### GetCron

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetExecuteSchedules200ResponseAllOfSchedule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


