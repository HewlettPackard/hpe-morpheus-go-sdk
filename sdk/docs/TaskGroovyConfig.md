# TaskGroovyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**TaskType** | Pointer to [**ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType**](ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**TaskOptions** | Pointer to [**ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions**](ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions.md) |  | [optional] 
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

### NewTaskGroovyConfig

`func NewTaskGroovyConfig() *TaskGroovyConfig`

NewTaskGroovyConfig instantiates a new TaskGroovyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskGroovyConfigWithDefaults

`func NewTaskGroovyConfigWithDefaults() *TaskGroovyConfig`

NewTaskGroovyConfigWithDefaults instantiates a new TaskGroovyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskGroovyConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskGroovyConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskGroovyConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TaskGroovyConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *TaskGroovyConfig) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TaskGroovyConfig) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TaskGroovyConfig) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TaskGroovyConfig) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *TaskGroovyConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskGroovyConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskGroovyConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskGroovyConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *TaskGroovyConfig) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TaskGroovyConfig) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TaskGroovyConfig) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TaskGroovyConfig) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTaskType

`func (o *TaskGroovyConfig) GetTaskType() ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *TaskGroovyConfig) GetTaskTypeOk() (*ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *TaskGroovyConfig) SetTaskType(v ListTasks200ResponseAllOfTasksInnerAnyOf4TaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *TaskGroovyConfig) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *TaskGroovyConfig) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *TaskGroovyConfig) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *TaskGroovyConfig) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *TaskGroovyConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVisibility

`func (o *TaskGroovyConfig) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *TaskGroovyConfig) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *TaskGroovyConfig) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *TaskGroovyConfig) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskOptions

`func (o *TaskGroovyConfig) GetTaskOptions() ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *TaskGroovyConfig) GetTaskOptionsOk() (*ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *TaskGroovyConfig) SetTaskOptions(v ListTasks200ResponseAllOfTasksInnerAnyOf4TaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *TaskGroovyConfig) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *TaskGroovyConfig) GetFile() ListTasks200ResponseAllOfTasksInnerAnyOfFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *TaskGroovyConfig) GetFileOk() (*ListTasks200ResponseAllOfTasksInnerAnyOfFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *TaskGroovyConfig) SetFile(v ListTasks200ResponseAllOfTasksInnerAnyOfFile)`

SetFile sets File field to given value.

### HasFile

`func (o *TaskGroovyConfig) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetResultType

`func (o *TaskGroovyConfig) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *TaskGroovyConfig) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *TaskGroovyConfig) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *TaskGroovyConfig) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetExecuteTarget

`func (o *TaskGroovyConfig) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *TaskGroovyConfig) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *TaskGroovyConfig) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *TaskGroovyConfig) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *TaskGroovyConfig) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *TaskGroovyConfig) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *TaskGroovyConfig) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *TaskGroovyConfig) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *TaskGroovyConfig) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *TaskGroovyConfig) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *TaskGroovyConfig) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *TaskGroovyConfig) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *TaskGroovyConfig) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *TaskGroovyConfig) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *TaskGroovyConfig) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *TaskGroovyConfig) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *TaskGroovyConfig) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *TaskGroovyConfig) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *TaskGroovyConfig) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *TaskGroovyConfig) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetCredential

`func (o *TaskGroovyConfig) GetCredential() ListClouds200ResponseAllOfZonesInnerCredentialAnyOf`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *TaskGroovyConfig) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredentialAnyOf, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *TaskGroovyConfig) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredentialAnyOf)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *TaskGroovyConfig) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDateCreated

`func (o *TaskGroovyConfig) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TaskGroovyConfig) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TaskGroovyConfig) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TaskGroovyConfig) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *TaskGroovyConfig) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *TaskGroovyConfig) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *TaskGroovyConfig) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *TaskGroovyConfig) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


