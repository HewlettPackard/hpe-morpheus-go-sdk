# GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**TaskType** | Pointer to [**GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskType**](GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**TaskOptions** | Pointer to [**GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskOptions**](GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskOptions.md) |  | [optional] 
**File** | Pointer to [**GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskFile**](GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskFile.md) |  | [optional] 
**ResultType** | Pointer to **NullableString** |  | [optional] 
**ExecuteTarget** | Pointer to **string** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 
**RetryCount** | Pointer to **int64** |  | [optional] 
**RetryDelaySeconds** | Pointer to **int64** |  | [optional] 
**AllowCustomConfig** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask

`func NewGetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask() *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask`

NewGetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask instantiates a new GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTaskType

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetTaskType() GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetTaskTypeOk() (*GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetTaskType(v GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTaskOptions

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetTaskOptions() GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetTaskOptionsOk() (*GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetTaskOptions(v GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskTaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetFile() GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetFileOk() (*GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetFile(v GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTaskFile)`

SetFile sets File field to given value.

### HasFile

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetResultType

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetExecuteTarget

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetWorkflows200ResponseAllOfTaskSetTaskSetTasksInnerTask) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


