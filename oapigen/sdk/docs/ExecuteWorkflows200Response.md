# ExecuteWorkflows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**ExecuteWorkflows200ResponseAllOfJob**](ExecuteWorkflows200ResponseAllOfJob.md) |  | [optional] 
**JobExecution** | Pointer to [**ExecuteWorkflows200ResponseAllOfJobExecution**](ExecuteWorkflows200ResponseAllOfJobExecution.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewExecuteWorkflows200Response

`func NewExecuteWorkflows200Response() *ExecuteWorkflows200Response`

NewExecuteWorkflows200Response instantiates a new ExecuteWorkflows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteWorkflows200ResponseWithDefaults

`func NewExecuteWorkflows200ResponseWithDefaults() *ExecuteWorkflows200Response`

NewExecuteWorkflows200ResponseWithDefaults instantiates a new ExecuteWorkflows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *ExecuteWorkflows200Response) GetJob() ExecuteWorkflows200ResponseAllOfJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ExecuteWorkflows200Response) GetJobOk() (*ExecuteWorkflows200ResponseAllOfJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ExecuteWorkflows200Response) SetJob(v ExecuteWorkflows200ResponseAllOfJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *ExecuteWorkflows200Response) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobExecution

`func (o *ExecuteWorkflows200Response) GetJobExecution() ExecuteWorkflows200ResponseAllOfJobExecution`

GetJobExecution returns the JobExecution field if non-nil, zero value otherwise.

### GetJobExecutionOk

`func (o *ExecuteWorkflows200Response) GetJobExecutionOk() (*ExecuteWorkflows200ResponseAllOfJobExecution, bool)`

GetJobExecutionOk returns a tuple with the JobExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExecution

`func (o *ExecuteWorkflows200Response) SetJobExecution(v ExecuteWorkflows200ResponseAllOfJobExecution)`

SetJobExecution sets JobExecution field to given value.

### HasJobExecution

`func (o *ExecuteWorkflows200Response) HasJobExecution() bool`

HasJobExecution returns a boolean if a field has been set.

### GetSuccess

`func (o *ExecuteWorkflows200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ExecuteWorkflows200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ExecuteWorkflows200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ExecuteWorkflows200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


