# UpdateBackupJobsRequestJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the backup job | [optional] 
**Code** | Pointer to **string** | A code for the backup job | [optional] 
**ScheduleId** | Pointer to **int64** | Execute Schedule ID to use for the backup job | [optional] 
**RetentionCount** | Pointer to **int64** | Retention Count | [optional] 

## Methods

### NewUpdateBackupJobsRequestJob

`func NewUpdateBackupJobsRequestJob() *UpdateBackupJobsRequestJob`

NewUpdateBackupJobsRequestJob instantiates a new UpdateBackupJobsRequestJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBackupJobsRequestJobWithDefaults

`func NewUpdateBackupJobsRequestJobWithDefaults() *UpdateBackupJobsRequestJob`

NewUpdateBackupJobsRequestJobWithDefaults instantiates a new UpdateBackupJobsRequestJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBackupJobsRequestJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBackupJobsRequestJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBackupJobsRequestJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBackupJobsRequestJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateBackupJobsRequestJob) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateBackupJobsRequestJob) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateBackupJobsRequestJob) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateBackupJobsRequestJob) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetScheduleId

`func (o *UpdateBackupJobsRequestJob) GetScheduleId() int64`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *UpdateBackupJobsRequestJob) GetScheduleIdOk() (*int64, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *UpdateBackupJobsRequestJob) SetScheduleId(v int64)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *UpdateBackupJobsRequestJob) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetRetentionCount

`func (o *UpdateBackupJobsRequestJob) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *UpdateBackupJobsRequestJob) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *UpdateBackupJobsRequestJob) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *UpdateBackupJobsRequestJob) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


