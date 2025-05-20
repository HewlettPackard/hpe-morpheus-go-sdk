# UpdateTasksRequestTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name for the task | [optional] 
**Code** | Pointer to **string** | A unique code for the task | [optional] 
**Visibility** | Pointer to **string** | Visibility | [optional] [default to "private"]
**TaskType** | Pointer to [**AddTasksRequestTaskTaskType**](AddTasksRequestTaskTaskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** | An array of Category labels for filtering | [optional] 
**TaskOptions** | Pointer to **map[string]interface{}** | Map of options specific to each &#x60;task type&#x60;. eg. script | [optional] 
**ResultType** | Pointer to **string** |  | [optional] 
**ExecuteTarget** | Pointer to **string** | The execution target. eg. local,remote,resource. The default value varies by task type.  | [optional] 
**Retryable** | Pointer to **bool** | If the task should be retried or not. | [optional] [default to false]
**RetryCount** | Pointer to **int32** | The number of times to retry. | [optional] 
**RetryDelaySeconds** | Pointer to **float32** | The delay, between retries. | [optional] 
**File** | Pointer to [**AddTasksRequestTaskFile**](AddTasksRequestTaskFile.md) |  | [optional] 
**Credential** | Pointer to [**AddIntegrationsRequestOneOfIntegrationCredential**](AddIntegrationsRequestOneOfIntegrationCredential.md) |  | [optional] 

## Methods

### NewUpdateTasksRequestTask

`func NewUpdateTasksRequestTask() *UpdateTasksRequestTask`

NewUpdateTasksRequestTask instantiates a new UpdateTasksRequestTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTasksRequestTaskWithDefaults

`func NewUpdateTasksRequestTaskWithDefaults() *UpdateTasksRequestTask`

NewUpdateTasksRequestTaskWithDefaults instantiates a new UpdateTasksRequestTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTasksRequestTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTasksRequestTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTasksRequestTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTasksRequestTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateTasksRequestTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateTasksRequestTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateTasksRequestTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateTasksRequestTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateTasksRequestTask) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateTasksRequestTask) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateTasksRequestTask) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateTasksRequestTask) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskType

`func (o *UpdateTasksRequestTask) GetTaskType() AddTasksRequestTaskTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *UpdateTasksRequestTask) GetTaskTypeOk() (*AddTasksRequestTaskTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *UpdateTasksRequestTask) SetTaskType(v AddTasksRequestTaskTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *UpdateTasksRequestTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateTasksRequestTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateTasksRequestTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateTasksRequestTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateTasksRequestTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTaskOptions

`func (o *UpdateTasksRequestTask) GetTaskOptions() map[string]interface{}`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *UpdateTasksRequestTask) GetTaskOptionsOk() (*map[string]interface{}, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *UpdateTasksRequestTask) SetTaskOptions(v map[string]interface{})`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *UpdateTasksRequestTask) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetResultType

`func (o *UpdateTasksRequestTask) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *UpdateTasksRequestTask) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *UpdateTasksRequestTask) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *UpdateTasksRequestTask) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetExecuteTarget

`func (o *UpdateTasksRequestTask) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *UpdateTasksRequestTask) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *UpdateTasksRequestTask) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *UpdateTasksRequestTask) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *UpdateTasksRequestTask) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *UpdateTasksRequestTask) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *UpdateTasksRequestTask) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *UpdateTasksRequestTask) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *UpdateTasksRequestTask) GetRetryCount() int32`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *UpdateTasksRequestTask) GetRetryCountOk() (*int32, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *UpdateTasksRequestTask) SetRetryCount(v int32)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *UpdateTasksRequestTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *UpdateTasksRequestTask) GetRetryDelaySeconds() float32`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *UpdateTasksRequestTask) GetRetryDelaySecondsOk() (*float32, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *UpdateTasksRequestTask) SetRetryDelaySeconds(v float32)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *UpdateTasksRequestTask) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetFile

`func (o *UpdateTasksRequestTask) GetFile() AddTasksRequestTaskFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UpdateTasksRequestTask) GetFileOk() (*AddTasksRequestTaskFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UpdateTasksRequestTask) SetFile(v AddTasksRequestTaskFile)`

SetFile sets File field to given value.

### HasFile

`func (o *UpdateTasksRequestTask) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetCredential

`func (o *UpdateTasksRequestTask) GetCredential() AddIntegrationsRequestOneOfIntegrationCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *UpdateTasksRequestTask) GetCredentialOk() (*AddIntegrationsRequestOneOfIntegrationCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *UpdateTasksRequestTask) SetCredential(v AddIntegrationsRequestOneOfIntegrationCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *UpdateTasksRequestTask) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


