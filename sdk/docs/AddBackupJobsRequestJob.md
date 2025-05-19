# AddBackupJobsRequestJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the backup job | 
**Code** | Pointer to **string** | A code for the backup job | [optional] 
**ScheduleId** | Pointer to **int64** | Execute Schedule ID to use for the backup job | [optional] 
**RetentionCount** | Pointer to **int64** | Retention Count. By default the backup settings value will be used. | [optional] 

## Methods

### NewAddBackupJobsRequestJob

`func NewAddBackupJobsRequestJob(name string, ) *AddBackupJobsRequestJob`

NewAddBackupJobsRequestJob instantiates a new AddBackupJobsRequestJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBackupJobsRequestJobWithDefaults

`func NewAddBackupJobsRequestJobWithDefaults() *AddBackupJobsRequestJob`

NewAddBackupJobsRequestJobWithDefaults instantiates a new AddBackupJobsRequestJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddBackupJobsRequestJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBackupJobsRequestJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBackupJobsRequestJob) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *AddBackupJobsRequestJob) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddBackupJobsRequestJob) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddBackupJobsRequestJob) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddBackupJobsRequestJob) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetScheduleId

`func (o *AddBackupJobsRequestJob) GetScheduleId() int64`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *AddBackupJobsRequestJob) GetScheduleIdOk() (*int64, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *AddBackupJobsRequestJob) SetScheduleId(v int64)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *AddBackupJobsRequestJob) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetRetentionCount

`func (o *AddBackupJobsRequestJob) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *AddBackupJobsRequestJob) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *AddBackupJobsRequestJob) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *AddBackupJobsRequestJob) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


