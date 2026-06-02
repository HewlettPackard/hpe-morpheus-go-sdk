# GetJobExecutions200ResponseAllOfJobExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Process** | Pointer to **NullableString** |  | [optional] 
**Job** | Pointer to [**GetJobExecutions200ResponseAllOfJobExecutionJob**](GetJobExecutions200ResponseAllOfJobExecutionJob.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**ResultData** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetJobExecutions200ResponseAllOfJobExecution

`func NewGetJobExecutions200ResponseAllOfJobExecution() *GetJobExecutions200ResponseAllOfJobExecution`

NewGetJobExecutions200ResponseAllOfJobExecution instantiates a new GetJobExecutions200ResponseAllOfJobExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcess

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetProcess() string`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetProcessOk() (*string, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetProcess(v string)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### SetProcessNil

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetProcessNil(b bool)`

 SetProcessNil sets the value for Process to be an explicit nil

### UnsetProcess
`func (o *GetJobExecutions200ResponseAllOfJobExecution) UnsetProcess()`

UnsetProcess ensures that no value is present for Process, not even an explicit nil
### GetJob

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetJob() GetJobExecutions200ResponseAllOfJobExecutionJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetJobOk() (*GetJobExecutions200ResponseAllOfJobExecutionJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetJob(v GetJobExecutions200ResponseAllOfJobExecutionJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetDescription

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetJobExecutions200ResponseAllOfJobExecution) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDateCreated

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetStartDate

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetResultData

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetResultData() string`

GetResultData returns the ResultData field if non-nil, zero value otherwise.

### GetResultDataOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetResultDataOk() (*string, bool)`

GetResultDataOk returns a tuple with the ResultData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultData

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetResultData(v string)`

SetResultData sets ResultData field to given value.

### HasResultData

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasResultData() bool`

HasResultData returns a boolean if a field has been set.

### SetResultDataNil

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetResultDataNil(b bool)`

 SetResultDataNil sets the value for ResultData to be an explicit nil

### UnsetResultData
`func (o *GetJobExecutions200ResponseAllOfJobExecution) UnsetResultData()`

UnsetResultData ensures that no value is present for ResultData, not even an explicit nil
### GetStatus

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetJobExecutions200ResponseAllOfJobExecution) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetCreatedBy

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetJobExecutions200ResponseAllOfJobExecution) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetJobExecutions200ResponseAllOfJobExecution) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GetJobExecutions200ResponseAllOfJobExecution) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GetJobExecutions200ResponseAllOfJobExecution) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


