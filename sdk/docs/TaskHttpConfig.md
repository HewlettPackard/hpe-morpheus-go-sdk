# TaskHttpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**TaskType** | Pointer to [**ListTasks200ResponseAllOfTasksInnerAnyOf5TaskType**](ListTasks200ResponseAllOfTasksInnerAnyOf5TaskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**TaskOptions** | Pointer to [**ListTasks200ResponseAllOfTasksInnerAnyOf5TaskOptions**](ListTasks200ResponseAllOfTasksInnerAnyOf5TaskOptions.md) |  | [optional] 
**File** | Pointer to [**ListTasks200ResponseAllOfTasksInnerAnyOfFile**](ListTasks200ResponseAllOfTasksInnerAnyOfFile.md) |  | [optional] 
**ResultType** | Pointer to **string** |  | [optional] 
**ExecuteTarget** | Pointer to **string** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 
**RetryCount** | Pointer to **int64** |  | [optional] 
**RetryDelaySeconds** | Pointer to **int64** |  | [optional] 
**AllowCustomConfig** | Pointer to **bool** |  | [optional] 
**Credential** | Pointer to [**ListClouds200ResponseAllOfZonesInnerCredentialAnyOf**](ListClouds200ResponseAllOfZonesInnerCredentialAnyOf.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTaskHttpConfig

`func NewTaskHttpConfig() *TaskHttpConfig`

NewTaskHttpConfig instantiates a new TaskHttpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskHttpConfigWithDefaults

`func NewTaskHttpConfigWithDefaults() *TaskHttpConfig`

NewTaskHttpConfigWithDefaults instantiates a new TaskHttpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskHttpConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskHttpConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskHttpConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TaskHttpConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *TaskHttpConfig) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TaskHttpConfig) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TaskHttpConfig) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TaskHttpConfig) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *TaskHttpConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskHttpConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskHttpConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskHttpConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *TaskHttpConfig) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaskHttpConfig) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaskHttpConfig) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaskHttpConfig) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTaskType

`func (o *TaskHttpConfig) GetTaskType() ListTasks200ResponseAllOfTasksInnerAnyOf5TaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *TaskHttpConfig) GetTaskTypeOk() (*ListTasks200ResponseAllOfTasksInnerAnyOf5TaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *TaskHttpConfig) SetTaskType(v ListTasks200ResponseAllOfTasksInnerAnyOf5TaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *TaskHttpConfig) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *TaskHttpConfig) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TaskHttpConfig) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TaskHttpConfig) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *TaskHttpConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVisibility

`func (o *TaskHttpConfig) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *TaskHttpConfig) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *TaskHttpConfig) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *TaskHttpConfig) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskOptions

`func (o *TaskHttpConfig) GetTaskOptions() ListTasks200ResponseAllOfTasksInnerAnyOf5TaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *TaskHttpConfig) GetTaskOptionsOk() (*ListTasks200ResponseAllOfTasksInnerAnyOf5TaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *TaskHttpConfig) SetTaskOptions(v ListTasks200ResponseAllOfTasksInnerAnyOf5TaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *TaskHttpConfig) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *TaskHttpConfig) GetFile() ListTasks200ResponseAllOfTasksInnerAnyOfFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TaskHttpConfig) GetFileOk() (*ListTasks200ResponseAllOfTasksInnerAnyOfFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TaskHttpConfig) SetFile(v ListTasks200ResponseAllOfTasksInnerAnyOfFile)`

SetFile sets File field to given value.

### HasFile

`func (o *TaskHttpConfig) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetResultType

`func (o *TaskHttpConfig) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *TaskHttpConfig) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *TaskHttpConfig) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *TaskHttpConfig) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetExecuteTarget

`func (o *TaskHttpConfig) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *TaskHttpConfig) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *TaskHttpConfig) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *TaskHttpConfig) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *TaskHttpConfig) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *TaskHttpConfig) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *TaskHttpConfig) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *TaskHttpConfig) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *TaskHttpConfig) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *TaskHttpConfig) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *TaskHttpConfig) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *TaskHttpConfig) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *TaskHttpConfig) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *TaskHttpConfig) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *TaskHttpConfig) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *TaskHttpConfig) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *TaskHttpConfig) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *TaskHttpConfig) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *TaskHttpConfig) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *TaskHttpConfig) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetCredential

`func (o *TaskHttpConfig) GetCredential() ListClouds200ResponseAllOfZonesInnerCredentialAnyOf`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *TaskHttpConfig) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredentialAnyOf, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *TaskHttpConfig) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredentialAnyOf)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *TaskHttpConfig) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDateCreated

`func (o *TaskHttpConfig) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskHttpConfig) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskHttpConfig) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskHttpConfig) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *TaskHttpConfig) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TaskHttpConfig) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TaskHttpConfig) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *TaskHttpConfig) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


