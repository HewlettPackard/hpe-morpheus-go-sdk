# ListBackups200ResponseAllOfBackupsInnerStats

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

### NewListBackups200ResponseAllOfBackupsInnerStats

`func NewListBackups200ResponseAllOfBackupsInnerStats() *ListBackups200ResponseAllOfBackupsInnerStats`

NewListBackups200ResponseAllOfBackupsInnerStats instantiates a new ListBackups200ResponseAllOfBackupsInnerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackups200ResponseAllOfBackupsInnerStatsWithDefaults

`func NewListBackups200ResponseAllOfBackupsInnerStatsWithDefaults() *ListBackups200ResponseAllOfBackupsInnerStats`

NewListBackups200ResponseAllOfBackupsInnerStatsWithDefaults instantiates a new ListBackups200ResponseAllOfBackupsInnerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalSize

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetAvgSize

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetAvgSize() int64`

GetAvgSize returns the AvgSize field if non-nil, zero value otherwise.

### GetAvgSizeOk

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetAvgSizeOk() (*int64, bool)`

GetAvgSizeOk returns a tuple with the AvgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSize

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) SetAvgSize(v int64)`

SetAvgSize sets AvgSize field to given value.

### HasAvgSize

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) HasAvgSize() bool`

HasAvgSize returns a boolean if a field has been set.

### GetTotalCompleted

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetTotalCompleted() int64`

GetTotalCompleted returns the TotalCompleted field if non-nil, zero value otherwise.

### GetTotalCompletedOk

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetTotalCompletedOk() (*int64, bool)`

GetTotalCompletedOk returns a tuple with the TotalCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCompleted

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) SetTotalCompleted(v int64)`

SetTotalCompleted sets TotalCompleted field to given value.

### HasTotalCompleted

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) HasTotalCompleted() bool`

HasTotalCompleted returns a boolean if a field has been set.

### GetSuccess

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetSuccess() int64`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetSuccessOk() (*int64, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) SetSuccess(v int64)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetSuccessRate

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetSuccessRate() float64`

GetSuccessRate returns the SuccessRate field if non-nil, zero value otherwise.

### GetSuccessRateOk

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetSuccessRateOk() (*float64, bool)`

GetSuccessRateOk returns a tuple with the SuccessRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRate

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) SetSuccessRate(v float64)`

SetSuccessRate sets SuccessRate field to given value.

### HasSuccessRate

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) HasSuccessRate() bool`

HasSuccessRate returns a boolean if a field has been set.

### GetFailRate

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetFailRate() float64`

GetFailRate returns the FailRate field if non-nil, zero value otherwise.

### GetFailRateOk

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetFailRateOk() (*float64, bool)`

GetFailRateOk returns a tuple with the FailRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailRate

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) SetFailRate(v float64)`

SetFailRate sets FailRate field to given value.

### HasFailRate

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) HasFailRate() bool`

HasFailRate returns a boolean if a field has been set.

### GetLastFiveResults

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetLastFiveResults() []string`

GetLastFiveResults returns the LastFiveResults field if non-nil, zero value otherwise.

### GetLastFiveResultsOk

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) GetLastFiveResultsOk() (*[]string, bool)`

GetLastFiveResultsOk returns a tuple with the LastFiveResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFiveResults

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) SetLastFiveResults(v []string)`

SetLastFiveResults sets LastFiveResults field to given value.

### HasLastFiveResults

`func (o *ListBackups200ResponseAllOfBackupsInnerStats) HasLastFiveResults() bool`

HasLastFiveResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


