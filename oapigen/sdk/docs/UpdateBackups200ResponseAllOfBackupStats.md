# UpdateBackups200ResponseAllOfBackupStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalSize** | Pointer to **int64** | Total size of all backups in bytes | [optional] 
**AvgSize** | Pointer to **int64** | Average size of each backup in bytes | [optional] 
**TotalCompleted** | Pointer to **int64** | Total completed count | [optional] 
**Success** | Pointer to **int64** | Successful backup count | [optional] 
**Failed** | Pointer to **int64** | Failed backup count | [optional] 
**SuccessRate** | Pointer to **float64** | Success rate 0-1 | [optional] 
**FailRate** | Pointer to **float64** | Failure rate 0-1 | [optional] 
**LastFiveResults** | Pointer to **[]string** | List of the last 5 backup result statuses | [optional] 

## Methods

### NewUpdateBackups200ResponseAllOfBackupStats

`func NewUpdateBackups200ResponseAllOfBackupStats() *UpdateBackups200ResponseAllOfBackupStats`

NewUpdateBackups200ResponseAllOfBackupStats instantiates a new UpdateBackups200ResponseAllOfBackupStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetTotalSize

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *UpdateBackups200ResponseAllOfBackupStats) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *UpdateBackups200ResponseAllOfBackupStats) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetAvgSize

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetAvgSize() int64`

GetAvgSize returns the AvgSize field if non-nil, zero value otherwise.

### GetAvgSizeOk

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetAvgSizeOk() (*int64, bool)`

GetAvgSizeOk returns a tuple with the AvgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSize

`func (o *UpdateBackups200ResponseAllOfBackupStats) SetAvgSize(v int64)`

SetAvgSize sets AvgSize field to given value.

### HasAvgSize

`func (o *UpdateBackups200ResponseAllOfBackupStats) HasAvgSize() bool`

HasAvgSize returns a boolean if a field has been set.

### GetTotalCompleted

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetTotalCompleted() int64`

GetTotalCompleted returns the TotalCompleted field if non-nil, zero value otherwise.

### GetTotalCompletedOk

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetTotalCompletedOk() (*int64, bool)`

GetTotalCompletedOk returns a tuple with the TotalCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCompleted

`func (o *UpdateBackups200ResponseAllOfBackupStats) SetTotalCompleted(v int64)`

SetTotalCompleted sets TotalCompleted field to given value.

### HasTotalCompleted

`func (o *UpdateBackups200ResponseAllOfBackupStats) HasTotalCompleted() bool`

HasTotalCompleted returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetSuccess() int64`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetSuccessOk() (*int64, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateBackups200ResponseAllOfBackupStats) SetSuccess(v int64)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateBackups200ResponseAllOfBackupStats) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *UpdateBackups200ResponseAllOfBackupStats) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *UpdateBackups200ResponseAllOfBackupStats) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetSuccessRate

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetSuccessRate() float64`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetSuccessRateOk() (*float64, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *UpdateBackups200ResponseAllOfBackupStats) SetSuccessRate(v float64)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *UpdateBackups200ResponseAllOfBackupStats) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetFailRate

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetFailRate() float64`

GetFailRate returns the FailRate field if non-nil, zero value otherwise.

### GetFailRateOk

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetFailRateOk() (*float64, bool)`

GetFailRateOk returns a tuple with the FailRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailRate

`func (o *UpdateBackups200ResponseAllOfBackupStats) SetFailRate(v float64)`

SetFailRate sets FailRate field to given value.

### HasFailRate

`func (o *UpdateBackups200ResponseAllOfBackupStats) HasFailRate() bool`

HasFailRate returns a boolean if a field has been set.

### GetLastFiveResults

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetLastFiveResults() []string`

GetLastFiveResults returns the LastFiveResults field if non-nil, zero value otherwise.

### GetLastFiveResultsOk

`func (o *UpdateBackups200ResponseAllOfBackupStats) GetLastFiveResultsOk() (*[]string, bool)`

GetLastFiveResultsOk returns a tuple with the LastFiveResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFiveResults

`func (o *UpdateBackups200ResponseAllOfBackupStats) SetLastFiveResults(v []string)`

SetLastFiveResults sets LastFiveResults field to given value.

### HasLastFiveResults

`func (o *UpdateBackups200ResponseAllOfBackupStats) HasLastFiveResults() bool`

HasLastFiveResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


