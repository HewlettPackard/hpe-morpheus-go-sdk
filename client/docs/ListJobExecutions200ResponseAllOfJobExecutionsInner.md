# ListJobExecutions200ResponseAllOfJobExecutionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Process** | Pointer to **string** |  | [optional] 
**Job** | Pointer to [**ListJobExecutions200ResponseAllOfJobExecutionsInnerJob**](ListJobExecutions200ResponseAllOfJobExecutionsInnerJob.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**ResultData** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewListJobExecutions200ResponseAllOfJobExecutionsInner

`func NewListJobExecutions200ResponseAllOfJobExecutionsInner() *ListJobExecutions200ResponseAllOfJobExecutionsInner`

NewListJobExecutions200ResponseAllOfJobExecutionsInner instantiates a new ListJobExecutions200ResponseAllOfJobExecutionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJobExecutions200ResponseAllOfJobExecutionsInnerWithDefaults

`func NewListJobExecutions200ResponseAllOfJobExecutionsInnerWithDefaults() *ListJobExecutions200ResponseAllOfJobExecutionsInner`

NewListJobExecutions200ResponseAllOfJobExecutionsInnerWithDefaults instantiates a new ListJobExecutions200ResponseAllOfJobExecutionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcess

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetProcess() string`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetProcessOk() (*string, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetProcess(v string)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetJob

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetJob() ListJobExecutions200ResponseAllOfJobExecutionsInnerJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetJobOk() (*ListJobExecutions200ResponseAllOfJobExecutionsInnerJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetJob(v ListJobExecutions200ResponseAllOfJobExecutionsInnerJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetDescription

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetStartDate

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetResultData

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetResultData() string`

GetResultData returns the ResultData field if non-nil, zero value otherwise.

### GetResultDataOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetResultDataOk() (*string, bool)`

GetResultDataOk returns a tuple with the ResultData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultData

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetResultData(v string)`

SetResultData sets ResultData field to given value.

### HasResultData

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasResultData() bool`

HasResultData returns a boolean if a field has been set.

### GetStatus

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListJobExecutions200ResponseAllOfJobExecutionsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


