# ExecuteTasks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**JobExecution** | Pointer to [**ExecuteTasks200ResponseAllOfJobExecution**](ExecuteTasks200ResponseAllOfJobExecution.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewExecuteTasks200Response

`func NewExecuteTasks200Response() *ExecuteTasks200Response`

NewExecuteTasks200Response instantiates a new ExecuteTasks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteTasks200ResponseWithDefaults

`func NewExecuteTasks200ResponseWithDefaults() *ExecuteTasks200Response`

NewExecuteTasks200ResponseWithDefaults instantiates a new ExecuteTasks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *ExecuteTasks200Response) GetJob() GetAlerts200ResponseAllOfChecksInnerAccount`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ExecuteTasks200Response) GetJobOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ExecuteTasks200Response) SetJob(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetJob sets Job field to given value.

### HasJob

`func (o *ExecuteTasks200Response) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobExecution

`func (o *ExecuteTasks200Response) GetJobExecution() ExecuteTasks200ResponseAllOfJobExecution`

GetJobExecution returns the JobExecution field if non-nil, zero value otherwise.

### GetJobExecutionOk

`func (o *ExecuteTasks200Response) GetJobExecutionOk() (*ExecuteTasks200ResponseAllOfJobExecution, bool)`

GetJobExecutionOk returns a tuple with the JobExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExecution

`func (o *ExecuteTasks200Response) SetJobExecution(v ExecuteTasks200ResponseAllOfJobExecution)`

SetJobExecution sets JobExecution field to given value.

### HasJobExecution

`func (o *ExecuteTasks200Response) HasJobExecution() bool`

HasJobExecution returns a boolean if a field has been set.

### GetSuccess

`func (o *ExecuteTasks200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ExecuteTasks200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ExecuteTasks200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ExecuteTasks200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


