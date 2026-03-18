# UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**TaskType** | Pointer to [**UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskType**](UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**TaskOptions** | Pointer to [**UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskOptions**](UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskOptions.md) |  | [optional] 
**File** | Pointer to [**UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskFile**](UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskFile.md) |  | [optional] 
**ResultType** | Pointer to **NullableString** |  | [optional] 
**ExecuteTarget** | Pointer to **string** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 
**RetryCount** | Pointer to **int64** |  | [optional] 
**RetryDelaySeconds** | Pointer to **int64** |  | [optional] 
**AllowCustomConfig** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask

`func NewUpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask() *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask`

NewUpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask instantiates a new UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskWithDefaults

`func NewUpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskWithDefaults() *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask`

NewUpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskWithDefaults instantiates a new UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTaskType

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetTaskType() UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetTaskTypeOk() (*UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetTaskType(v UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTaskOptions

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetTaskOptions() UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetTaskOptionsOk() (*UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetTaskOptions(v UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetFile() UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetFileOk() (*UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetFile(v UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskFile)`

SetFile sets File field to given value.

### HasFile

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetResultType

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetExecuteTarget

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


