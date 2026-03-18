# UpdateExecuteSchedules200ResponseAllOfSchedule

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

### NewUpdateExecuteSchedules200ResponseAllOfSchedule

`func NewUpdateExecuteSchedules200ResponseAllOfSchedule() *UpdateExecuteSchedules200ResponseAllOfSchedule`

NewUpdateExecuteSchedules200ResponseAllOfSchedule instantiates a new UpdateExecuteSchedules200ResponseAllOfSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExecuteSchedules200ResponseAllOfScheduleWithDefaults

`func NewUpdateExecuteSchedules200ResponseAllOfScheduleWithDefaults() *UpdateExecuteSchedules200ResponseAllOfSchedule`

NewUpdateExecuteSchedules200ResponseAllOfScheduleWithDefaults instantiates a new UpdateExecuteSchedules200ResponseAllOfSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVisibility

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### SetVisibilityNil

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil
### GetEnabled

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetScheduleType

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.

### HasScheduleType

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) HasScheduleType() bool`

HasScheduleType returns a boolean if a field has been set.

### GetScheduleTimezone

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetScheduleTimezone() string`

GetScheduleTimezone returns the ScheduleTimezone field if non-nil, zero value otherwise.

### GetScheduleTimezoneOk

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetScheduleTimezoneOk() (*string, bool)`

GetScheduleTimezoneOk returns a tuple with the ScheduleTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTimezone

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetScheduleTimezone(v string)`

SetScheduleTimezone sets ScheduleTimezone field to given value.

### HasScheduleTimezone

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) HasScheduleTimezone() bool`

HasScheduleTimezone returns a boolean if a field has been set.

### SetScheduleTimezoneNil

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetScheduleTimezoneNil(b bool)`

 SetScheduleTimezoneNil sets the value for ScheduleTimezone to be an explicit nil

### UnsetScheduleTimezone
`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) UnsetScheduleTimezone()`

UnsetScheduleTimezone ensures that no value is present for ScheduleTimezone, not even an explicit nil
### GetCron

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateExecuteSchedules200ResponseAllOfSchedule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


