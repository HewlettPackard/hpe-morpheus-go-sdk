# AddTasksRequestTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for the task | 
**Code** | Pointer to **string** | A unique code for the task | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**TaskType** | [**AddTasksRequestTaskTaskType**](AddTasksRequestTaskTaskType.md) |  | 
**Labels** | Pointer to **[]string** | An array of Category labels for filtering | [optional] 
**TaskOptions** | Pointer to **map[string]interface{}** | Map of options specific to each &#x60;task type&#x60;. eg. script | [optional] 
**ResultType** | Pointer to **string** |  | [optional] 
**ExecuteTarget** | **string** | The execution target. eg. local,remote,resource. The default value varies by task type.  | 
**Retryable** | Pointer to **bool** | If the task should be retried or not. | [optional] [default to false]
**RetryCount** | Pointer to **int32** | The number of times to retry. | [optional] 
**RetryDelaySeconds** | Pointer to **float32** | The delay, between retries. | [optional] 
**File** | Pointer to [**AddTasksRequestTaskFile**](AddTasksRequestTaskFile.md) |  | [optional] 
**Credential** | Pointer to [**AddTasksRequestTaskCredential**](AddTasksRequestTaskCredential.md) |  | [optional] 

## Methods

### NewAddTasksRequestTask

`func NewAddTasksRequestTask(name string, taskType AddTasksRequestTaskTaskType, executeTarget string, ) *AddTasksRequestTask`

NewAddTasksRequestTask instantiates a new AddTasksRequestTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTasksRequestTaskWithDefaults

`func NewAddTasksRequestTaskWithDefaults() *AddTasksRequestTask`

NewAddTasksRequestTaskWithDefaults instantiates a new AddTasksRequestTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddTasksRequestTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddTasksRequestTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddTasksRequestTask) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *AddTasksRequestTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddTasksRequestTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddTasksRequestTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddTasksRequestTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetVisibility

`func (o *AddTasksRequestTask) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddTasksRequestTask) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddTasksRequestTask) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddTasksRequestTask) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskType

`func (o *AddTasksRequestTask) GetTaskType() AddTasksRequestTaskTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *AddTasksRequestTask) GetTaskTypeOk() (*AddTasksRequestTaskTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *AddTasksRequestTask) SetTaskType(v AddTasksRequestTaskTaskType)`

SetTaskType sets TaskType field to given value.


### GetLabels

`func (o *AddTasksRequestTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddTasksRequestTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddTasksRequestTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddTasksRequestTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTaskOptions

`func (o *AddTasksRequestTask) GetTaskOptions() map[string]interface{}`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *AddTasksRequestTask) GetTaskOptionsOk() (*map[string]interface{}, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *AddTasksRequestTask) SetTaskOptions(v map[string]interface{})`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *AddTasksRequestTask) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetResultType

`func (o *AddTasksRequestTask) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *AddTasksRequestTask) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *AddTasksRequestTask) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *AddTasksRequestTask) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetExecuteTarget

`func (o *AddTasksRequestTask) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *AddTasksRequestTask) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *AddTasksRequestTask) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.


### GetRetryable

`func (o *AddTasksRequestTask) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *AddTasksRequestTask) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *AddTasksRequestTask) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *AddTasksRequestTask) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *AddTasksRequestTask) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *AddTasksRequestTask) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *AddTasksRequestTask) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *AddTasksRequestTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *AddTasksRequestTask) GetRetryDelaySeconds() float32`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *AddTasksRequestTask) GetRetryDelaySecondsOk() (*float32, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *AddTasksRequestTask) SetRetryDelaySeconds(v float32)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *AddTasksRequestTask) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetFile

`func (o *AddTasksRequestTask) GetFile() AddTasksRequestTaskFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AddTasksRequestTask) GetFileOk() (*AddTasksRequestTaskFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AddTasksRequestTask) SetFile(v AddTasksRequestTaskFile)`

SetFile sets File field to given value.

### HasFile

`func (o *AddTasksRequestTask) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetCredential

`func (o *AddTasksRequestTask) GetCredential() AddTasksRequestTaskCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AddTasksRequestTask) GetCredentialOk() (*AddTasksRequestTaskCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AddTasksRequestTask) SetCredential(v AddTasksRequestTaskCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AddTasksRequestTask) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


