# GetBackups200ResponseBackupStats

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

### NewGetBackups200ResponseBackupStats

`func NewGetBackups200ResponseBackupStats() *GetBackups200ResponseBackupStats`

NewGetBackups200ResponseBackupStats instantiates a new GetBackups200ResponseBackupStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBackups200ResponseBackupStatsWithDefaults

`func NewGetBackups200ResponseBackupStatsWithDefaults() *GetBackups200ResponseBackupStats`

NewGetBackups200ResponseBackupStatsWithDefaults instantiates a new GetBackups200ResponseBackupStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSize

`func (o *GetBackups200ResponseBackupStats) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *GetBackups200ResponseBackupStats) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *GetBackups200ResponseBackupStats) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *GetBackups200ResponseBackupStats) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetAvgSize

`func (o *GetBackups200ResponseBackupStats) GetAvgSize() int64`

GetAvgSize returns the AvgSize field if non-nil, zero value otherwise.

### GetAvgSizeOk

`func (o *GetBackups200ResponseBackupStats) GetAvgSizeOk() (*int64, bool)`

GetAvgSizeOk returns a tuple with the AvgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSize

`func (o *GetBackups200ResponseBackupStats) SetAvgSize(v int64)`

SetAvgSize sets AvgSize field to given value.

### HasAvgSize

`func (o *GetBackups200ResponseBackupStats) HasAvgSize() bool`

HasAvgSize returns a boolean if a field has been set.

### GetTotalCompleted

`func (o *GetBackups200ResponseBackupStats) GetTotalCompleted() int64`

GetTotalCompleted returns the TotalCompleted field if non-nil, zero value otherwise.

### GetTotalCompletedOk

`func (o *GetBackups200ResponseBackupStats) GetTotalCompletedOk() (*int64, bool)`

GetTotalCompletedOk returns a tuple with the TotalCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCompleted

`func (o *GetBackups200ResponseBackupStats) SetTotalCompleted(v int64)`

SetTotalCompleted sets TotalCompleted field to given value.

### HasTotalCompleted

`func (o *GetBackups200ResponseBackupStats) HasTotalCompleted() bool`

HasTotalCompleted returns a boolean if a field has been set.

### GetSuccess

`func (o *GetBackups200ResponseBackupStats) GetSuccess() int64`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetBackups200ResponseBackupStats) GetSuccessOk() (*int64, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetBackups200ResponseBackupStats) SetSuccess(v int64)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetBackups200ResponseBackupStats) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *GetBackups200ResponseBackupStats) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *GetBackups200ResponseBackupStats) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *GetBackups200ResponseBackupStats) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *GetBackups200ResponseBackupStats) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetSuccessRate

`func (o *GetBackups200ResponseBackupStats) GetSuccessRate() float64`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *GetBackups200ResponseBackupStats) GetSuccessRateOk() (*float64, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *GetBackups200ResponseBackupStats) SetSuccessRate(v float64)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *GetBackups200ResponseBackupStats) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetFailRate

`func (o *GetBackups200ResponseBackupStats) GetFailRate() float64`

GetFailRate returns the FailRate field if non-nil, zero value otherwise.

### GetFailRateOk

`func (o *GetBackups200ResponseBackupStats) GetFailRateOk() (*float64, bool)`

GetFailRateOk returns a tuple with the FailRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailRate

`func (o *GetBackups200ResponseBackupStats) SetFailRate(v float64)`

SetFailRate sets FailRate field to given value.

### HasFailRate

`func (o *GetBackups200ResponseBackupStats) HasFailRate() bool`

HasFailRate returns a boolean if a field has been set.

### GetLastFiveResults

`func (o *GetBackups200ResponseBackupStats) GetLastFiveResults() []string`

GetLastFiveResults returns the LastFiveResults field if non-nil, zero value otherwise.

### GetLastFiveResultsOk

`func (o *GetBackups200ResponseBackupStats) GetLastFiveResultsOk() (*[]string, bool)`

GetLastFiveResultsOk returns a tuple with the LastFiveResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFiveResults

`func (o *GetBackups200ResponseBackupStats) SetLastFiveResults(v []string)`

SetLastFiveResults sets LastFiveResults field to given value.

### HasLastFiveResults

`func (o *GetBackups200ResponseBackupStats) HasLastFiveResults() bool`

HasLastFiveResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


