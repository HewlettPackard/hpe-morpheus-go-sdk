# AddTasks200ResponseAllOfTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**TaskType** | Pointer to [**AddTasks200ResponseAllOfTaskTaskType**](AddTasks200ResponseAllOfTaskTaskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**TaskOptions** | Pointer to [**AddTasks200ResponseAllOfTaskTaskOptions**](AddTasks200ResponseAllOfTaskTaskOptions.md) |  | [optional] 
**File** | Pointer to [**AddTasks200ResponseAllOfTaskFile**](AddTasks200ResponseAllOfTaskFile.md) |  | [optional] 
**ResultType** | Pointer to **NullableString** |  | [optional] 
**ExecuteTarget** | Pointer to **string** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 
**RetryCount** | Pointer to **int64** |  | [optional] 
**RetryDelaySeconds** | Pointer to **int64** |  | [optional] 
**AllowCustomConfig** | Pointer to **bool** |  | [optional] 
**Credential** | Pointer to [**AddTasks200ResponseAllOfTaskCredential**](AddTasks200ResponseAllOfTaskCredential.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddTasks200ResponseAllOfTask

`func NewAddTasks200ResponseAllOfTask() *AddTasks200ResponseAllOfTask`

NewAddTasks200ResponseAllOfTask instantiates a new AddTasks200ResponseAllOfTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTasks200ResponseAllOfTaskWithDefaults

`func NewAddTasks200ResponseAllOfTaskWithDefaults() *AddTasks200ResponseAllOfTask`

NewAddTasks200ResponseAllOfTaskWithDefaults instantiates a new AddTasks200ResponseAllOfTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddTasks200ResponseAllOfTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddTasks200ResponseAllOfTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddTasks200ResponseAllOfTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddTasks200ResponseAllOfTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *AddTasks200ResponseAllOfTask) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddTasks200ResponseAllOfTask) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddTasks200ResponseAllOfTask) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddTasks200ResponseAllOfTask) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *AddTasks200ResponseAllOfTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddTasks200ResponseAllOfTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddTasks200ResponseAllOfTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddTasks200ResponseAllOfTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddTasks200ResponseAllOfTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddTasks200ResponseAllOfTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddTasks200ResponseAllOfTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddTasks200ResponseAllOfTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AddTasks200ResponseAllOfTask) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AddTasks200ResponseAllOfTask) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTaskType

`func (o *AddTasks200ResponseAllOfTask) GetTaskType() AddTasks200ResponseAllOfTaskTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *AddTasks200ResponseAllOfTask) GetTaskTypeOk() (*AddTasks200ResponseAllOfTaskTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *AddTasks200ResponseAllOfTask) SetTaskType(v AddTasks200ResponseAllOfTaskTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *AddTasks200ResponseAllOfTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *AddTasks200ResponseAllOfTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddTasks200ResponseAllOfTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddTasks200ResponseAllOfTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddTasks200ResponseAllOfTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVisibility

`func (o *AddTasks200ResponseAllOfTask) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddTasks200ResponseAllOfTask) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddTasks200ResponseAllOfTask) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddTasks200ResponseAllOfTask) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskOptions

`func (o *AddTasks200ResponseAllOfTask) GetTaskOptions() AddTasks200ResponseAllOfTaskTaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *AddTasks200ResponseAllOfTask) GetTaskOptionsOk() (*AddTasks200ResponseAllOfTaskTaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *AddTasks200ResponseAllOfTask) SetTaskOptions(v AddTasks200ResponseAllOfTaskTaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *AddTasks200ResponseAllOfTask) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *AddTasks200ResponseAllOfTask) GetFile() AddTasks200ResponseAllOfTaskFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AddTasks200ResponseAllOfTask) GetFileOk() (*AddTasks200ResponseAllOfTaskFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AddTasks200ResponseAllOfTask) SetFile(v AddTasks200ResponseAllOfTaskFile)`

SetFile sets File field to given value.

### HasFile

`func (o *AddTasks200ResponseAllOfTask) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetResultType

`func (o *AddTasks200ResponseAllOfTask) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *AddTasks200ResponseAllOfTask) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *AddTasks200ResponseAllOfTask) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *AddTasks200ResponseAllOfTask) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *AddTasks200ResponseAllOfTask) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *AddTasks200ResponseAllOfTask) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetExecuteTarget

`func (o *AddTasks200ResponseAllOfTask) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *AddTasks200ResponseAllOfTask) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *AddTasks200ResponseAllOfTask) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *AddTasks200ResponseAllOfTask) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *AddTasks200ResponseAllOfTask) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *AddTasks200ResponseAllOfTask) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *AddTasks200ResponseAllOfTask) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *AddTasks200ResponseAllOfTask) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *AddTasks200ResponseAllOfTask) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *AddTasks200ResponseAllOfTask) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *AddTasks200ResponseAllOfTask) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *AddTasks200ResponseAllOfTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *AddTasks200ResponseAllOfTask) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *AddTasks200ResponseAllOfTask) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *AddTasks200ResponseAllOfTask) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *AddTasks200ResponseAllOfTask) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *AddTasks200ResponseAllOfTask) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *AddTasks200ResponseAllOfTask) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *AddTasks200ResponseAllOfTask) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *AddTasks200ResponseAllOfTask) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetCredential

`func (o *AddTasks200ResponseAllOfTask) GetCredential() AddTasks200ResponseAllOfTaskCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AddTasks200ResponseAllOfTask) GetCredentialOk() (*AddTasks200ResponseAllOfTaskCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AddTasks200ResponseAllOfTask) SetCredential(v AddTasks200ResponseAllOfTaskCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AddTasks200ResponseAllOfTask) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddTasks200ResponseAllOfTask) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddTasks200ResponseAllOfTask) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddTasks200ResponseAllOfTask) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddTasks200ResponseAllOfTask) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddTasks200ResponseAllOfTask) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddTasks200ResponseAllOfTask) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddTasks200ResponseAllOfTask) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddTasks200ResponseAllOfTask) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


