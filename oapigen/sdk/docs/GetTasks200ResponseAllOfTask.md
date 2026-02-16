# GetTasks200ResponseAllOfTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**TaskType** | Pointer to [**GetTasks200ResponseAllOfTaskTaskType**](GetTasks200ResponseAllOfTaskTaskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**TaskOptions** | Pointer to [**GetTasks200ResponseAllOfTaskTaskOptions**](GetTasks200ResponseAllOfTaskTaskOptions.md) |  | [optional] 
**File** | Pointer to [**GetTasks200ResponseAllOfTaskFile**](GetTasks200ResponseAllOfTaskFile.md) |  | [optional] 
**ResultType** | Pointer to **NullableString** |  | [optional] 
**ExecuteTarget** | Pointer to **string** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 
**RetryCount** | Pointer to **int64** |  | [optional] 
**RetryDelaySeconds** | Pointer to **int64** |  | [optional] 
**AllowCustomConfig** | Pointer to **bool** |  | [optional] 
**Credential** | Pointer to [**GetTasks200ResponseAllOfTaskCredential**](GetTasks200ResponseAllOfTaskCredential.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetTasks200ResponseAllOfTask

`func NewGetTasks200ResponseAllOfTask() *GetTasks200ResponseAllOfTask`

NewGetTasks200ResponseAllOfTask instantiates a new GetTasks200ResponseAllOfTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTasks200ResponseAllOfTaskWithDefaults

`func NewGetTasks200ResponseAllOfTaskWithDefaults() *GetTasks200ResponseAllOfTask`

NewGetTasks200ResponseAllOfTaskWithDefaults instantiates a new GetTasks200ResponseAllOfTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTasks200ResponseAllOfTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTasks200ResponseAllOfTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTasks200ResponseAllOfTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetTasks200ResponseAllOfTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetTasks200ResponseAllOfTask) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetTasks200ResponseAllOfTask) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetTasks200ResponseAllOfTask) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetTasks200ResponseAllOfTask) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *GetTasks200ResponseAllOfTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTasks200ResponseAllOfTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTasks200ResponseAllOfTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTasks200ResponseAllOfTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetTasks200ResponseAllOfTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetTasks200ResponseAllOfTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetTasks200ResponseAllOfTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetTasks200ResponseAllOfTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetTasks200ResponseAllOfTask) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetTasks200ResponseAllOfTask) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTaskType

`func (o *GetTasks200ResponseAllOfTask) GetTaskType() GetTasks200ResponseAllOfTaskTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *GetTasks200ResponseAllOfTask) GetTaskTypeOk() (*GetTasks200ResponseAllOfTaskTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *GetTasks200ResponseAllOfTask) SetTaskType(v GetTasks200ResponseAllOfTaskTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *GetTasks200ResponseAllOfTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *GetTasks200ResponseAllOfTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetTasks200ResponseAllOfTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetTasks200ResponseAllOfTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetTasks200ResponseAllOfTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVisibility

`func (o *GetTasks200ResponseAllOfTask) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetTasks200ResponseAllOfTask) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetTasks200ResponseAllOfTask) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetTasks200ResponseAllOfTask) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskOptions

`func (o *GetTasks200ResponseAllOfTask) GetTaskOptions() GetTasks200ResponseAllOfTaskTaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *GetTasks200ResponseAllOfTask) GetTaskOptionsOk() (*GetTasks200ResponseAllOfTaskTaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *GetTasks200ResponseAllOfTask) SetTaskOptions(v GetTasks200ResponseAllOfTaskTaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *GetTasks200ResponseAllOfTask) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *GetTasks200ResponseAllOfTask) GetFile() GetTasks200ResponseAllOfTaskFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *GetTasks200ResponseAllOfTask) GetFileOk() (*GetTasks200ResponseAllOfTaskFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *GetTasks200ResponseAllOfTask) SetFile(v GetTasks200ResponseAllOfTaskFile)`

SetFile sets File field to given value.

### HasFile

`func (o *GetTasks200ResponseAllOfTask) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetResultType

`func (o *GetTasks200ResponseAllOfTask) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *GetTasks200ResponseAllOfTask) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *GetTasks200ResponseAllOfTask) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *GetTasks200ResponseAllOfTask) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *GetTasks200ResponseAllOfTask) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *GetTasks200ResponseAllOfTask) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetExecuteTarget

`func (o *GetTasks200ResponseAllOfTask) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *GetTasks200ResponseAllOfTask) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *GetTasks200ResponseAllOfTask) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *GetTasks200ResponseAllOfTask) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *GetTasks200ResponseAllOfTask) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *GetTasks200ResponseAllOfTask) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *GetTasks200ResponseAllOfTask) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *GetTasks200ResponseAllOfTask) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *GetTasks200ResponseAllOfTask) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *GetTasks200ResponseAllOfTask) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *GetTasks200ResponseAllOfTask) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *GetTasks200ResponseAllOfTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *GetTasks200ResponseAllOfTask) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *GetTasks200ResponseAllOfTask) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *GetTasks200ResponseAllOfTask) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *GetTasks200ResponseAllOfTask) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *GetTasks200ResponseAllOfTask) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *GetTasks200ResponseAllOfTask) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *GetTasks200ResponseAllOfTask) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *GetTasks200ResponseAllOfTask) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetCredential

`func (o *GetTasks200ResponseAllOfTask) GetCredential() GetTasks200ResponseAllOfTaskCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetTasks200ResponseAllOfTask) GetCredentialOk() (*GetTasks200ResponseAllOfTaskCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetTasks200ResponseAllOfTask) SetCredential(v GetTasks200ResponseAllOfTaskCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetTasks200ResponseAllOfTask) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetTasks200ResponseAllOfTask) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetTasks200ResponseAllOfTask) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetTasks200ResponseAllOfTask) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetTasks200ResponseAllOfTask) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetTasks200ResponseAllOfTask) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetTasks200ResponseAllOfTask) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetTasks200ResponseAllOfTask) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetTasks200ResponseAllOfTask) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


