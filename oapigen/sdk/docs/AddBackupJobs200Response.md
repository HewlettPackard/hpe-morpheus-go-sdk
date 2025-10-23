# AddBackupJobs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**ListBackupJobs200ResponseAllOfJobsInner**](ListBackupJobs200ResponseAllOfJobsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddBackupJobs200Response

`func NewAddBackupJobs200Response() *AddBackupJobs200Response`

NewAddBackupJobs200Response instantiates a new AddBackupJobs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBackupJobs200ResponseWithDefaults

`func NewAddBackupJobs200ResponseWithDefaults() *AddBackupJobs200Response`

NewAddBackupJobs200ResponseWithDefaults instantiates a new AddBackupJobs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *AddBackupJobs200Response) GetJob() ListBackupJobs200ResponseAllOfJobsInner`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *AddBackupJobs200Response) GetJobOk() (*ListBackupJobs200ResponseAllOfJobsInner, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *AddBackupJobs200Response) SetJob(v ListBackupJobs200ResponseAllOfJobsInner)`

SetJob sets Job field to given value.

### HasJob

`func (o *AddBackupJobs200Response) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetSuccess

`func (o *AddBackupJobs200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddBackupJobs200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddBackupJobs200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddBackupJobs200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


