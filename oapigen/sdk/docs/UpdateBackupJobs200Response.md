# UpdateBackupJobs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**UpdateBackupJobs200ResponseAllOfJob**](UpdateBackupJobs200ResponseAllOfJob.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateBackupJobs200Response

`func NewUpdateBackupJobs200Response() *UpdateBackupJobs200Response`

NewUpdateBackupJobs200Response instantiates a new UpdateBackupJobs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetJob

`func (o *UpdateBackupJobs200Response) GetJob() UpdateBackupJobs200ResponseAllOfJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *UpdateBackupJobs200Response) GetJobOk() (*UpdateBackupJobs200ResponseAllOfJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *UpdateBackupJobs200Response) SetJob(v UpdateBackupJobs200ResponseAllOfJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *UpdateBackupJobs200Response) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateBackupJobs200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateBackupJobs200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateBackupJobs200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateBackupJobs200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


