# ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**TaskType** | Pointer to [**ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType**](ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**TaskOptions** | Pointer to [**ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskTaskOptions**](ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskTaskOptions.md) |  | [optional] 
**File** | Pointer to [**ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskFile**](ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskFile.md) |  | [optional] 
**ResultType** | Pointer to **string** |  | [optional] 
**ExecuteTarget** | Pointer to **string** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 
**RetryCount** | Pointer to **int64** |  | [optional] 
**RetryDelaySeconds** | Pointer to **int64** |  | [optional] 
**AllowCustomConfig** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask

`func NewListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask() *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask`

NewListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask instantiates a new ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskWithDefaults

`func NewListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskWithDefaults() *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask`

NewListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskWithDefaults instantiates a new ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTaskType

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetTaskType() ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetTaskTypeOk() (*ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetTaskType(v ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTaskOptions

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetTaskOptions() ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskTaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetTaskOptionsOk() (*ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskTaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetTaskOptions(v ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskTaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetFile() ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetFileOk() (*ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetFile(v ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTaskFile)`

SetFile sets File field to given value.

### HasFile

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetResultType

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetExecuteTarget

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListWorkflows200ResponseAllOfTaskSetsInnerTaskSetTasksInnerTask) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


