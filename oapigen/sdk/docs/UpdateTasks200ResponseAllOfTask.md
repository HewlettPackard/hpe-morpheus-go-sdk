# UpdateTasks200ResponseAllOfTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**TaskType** | Pointer to [**UpdateTasks200ResponseAllOfTaskTaskType**](UpdateTasks200ResponseAllOfTaskTaskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**TaskOptions** | Pointer to [**UpdateTasks200ResponseAllOfTaskTaskOptions**](UpdateTasks200ResponseAllOfTaskTaskOptions.md) |  | [optional] 
**File** | Pointer to [**UpdateTasks200ResponseAllOfTaskFile**](UpdateTasks200ResponseAllOfTaskFile.md) |  | [optional] 
**ResultType** | Pointer to **NullableString** |  | [optional] 
**ExecuteTarget** | Pointer to **string** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 
**RetryCount** | Pointer to **int64** |  | [optional] 
**RetryDelaySeconds** | Pointer to **int64** |  | [optional] 
**AllowCustomConfig** | Pointer to **bool** |  | [optional] 
**Credential** | Pointer to [**UpdateTasks200ResponseAllOfTaskCredential**](UpdateTasks200ResponseAllOfTaskCredential.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateTasks200ResponseAllOfTask

`func NewUpdateTasks200ResponseAllOfTask() *UpdateTasks200ResponseAllOfTask`

NewUpdateTasks200ResponseAllOfTask instantiates a new UpdateTasks200ResponseAllOfTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTasks200ResponseAllOfTaskWithDefaults

`func NewUpdateTasks200ResponseAllOfTaskWithDefaults() *UpdateTasks200ResponseAllOfTask`

NewUpdateTasks200ResponseAllOfTaskWithDefaults instantiates a new UpdateTasks200ResponseAllOfTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateTasks200ResponseAllOfTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTasks200ResponseAllOfTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTasks200ResponseAllOfTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateTasks200ResponseAllOfTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateTasks200ResponseAllOfTask) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateTasks200ResponseAllOfTask) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateTasks200ResponseAllOfTask) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateTasks200ResponseAllOfTask) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *UpdateTasks200ResponseAllOfTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTasks200ResponseAllOfTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTasks200ResponseAllOfTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTasks200ResponseAllOfTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateTasks200ResponseAllOfTask) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateTasks200ResponseAllOfTask) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateTasks200ResponseAllOfTask) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateTasks200ResponseAllOfTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateTasks200ResponseAllOfTask) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateTasks200ResponseAllOfTask) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetTaskType

`func (o *UpdateTasks200ResponseAllOfTask) GetTaskType() UpdateTasks200ResponseAllOfTaskTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *UpdateTasks200ResponseAllOfTask) GetTaskTypeOk() (*UpdateTasks200ResponseAllOfTaskTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *UpdateTasks200ResponseAllOfTask) SetTaskType(v UpdateTasks200ResponseAllOfTaskTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *UpdateTasks200ResponseAllOfTask) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateTasks200ResponseAllOfTask) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateTasks200ResponseAllOfTask) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateTasks200ResponseAllOfTask) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateTasks200ResponseAllOfTask) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateTasks200ResponseAllOfTask) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateTasks200ResponseAllOfTask) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateTasks200ResponseAllOfTask) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateTasks200ResponseAllOfTask) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskOptions

`func (o *UpdateTasks200ResponseAllOfTask) GetTaskOptions() UpdateTasks200ResponseAllOfTaskTaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *UpdateTasks200ResponseAllOfTask) GetTaskOptionsOk() (*UpdateTasks200ResponseAllOfTaskTaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *UpdateTasks200ResponseAllOfTask) SetTaskOptions(v UpdateTasks200ResponseAllOfTaskTaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *UpdateTasks200ResponseAllOfTask) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *UpdateTasks200ResponseAllOfTask) GetFile() UpdateTasks200ResponseAllOfTaskFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UpdateTasks200ResponseAllOfTask) GetFileOk() (*UpdateTasks200ResponseAllOfTaskFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UpdateTasks200ResponseAllOfTask) SetFile(v UpdateTasks200ResponseAllOfTaskFile)`

SetFile sets File field to given value.

### HasFile

`func (o *UpdateTasks200ResponseAllOfTask) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetResultType

`func (o *UpdateTasks200ResponseAllOfTask) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *UpdateTasks200ResponseAllOfTask) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *UpdateTasks200ResponseAllOfTask) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *UpdateTasks200ResponseAllOfTask) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *UpdateTasks200ResponseAllOfTask) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *UpdateTasks200ResponseAllOfTask) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetExecuteTarget

`func (o *UpdateTasks200ResponseAllOfTask) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *UpdateTasks200ResponseAllOfTask) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *UpdateTasks200ResponseAllOfTask) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *UpdateTasks200ResponseAllOfTask) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *UpdateTasks200ResponseAllOfTask) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *UpdateTasks200ResponseAllOfTask) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *UpdateTasks200ResponseAllOfTask) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *UpdateTasks200ResponseAllOfTask) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *UpdateTasks200ResponseAllOfTask) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *UpdateTasks200ResponseAllOfTask) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *UpdateTasks200ResponseAllOfTask) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *UpdateTasks200ResponseAllOfTask) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *UpdateTasks200ResponseAllOfTask) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *UpdateTasks200ResponseAllOfTask) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *UpdateTasks200ResponseAllOfTask) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *UpdateTasks200ResponseAllOfTask) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *UpdateTasks200ResponseAllOfTask) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *UpdateTasks200ResponseAllOfTask) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *UpdateTasks200ResponseAllOfTask) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *UpdateTasks200ResponseAllOfTask) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetCredential

`func (o *UpdateTasks200ResponseAllOfTask) GetCredential() UpdateTasks200ResponseAllOfTaskCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *UpdateTasks200ResponseAllOfTask) GetCredentialOk() (*UpdateTasks200ResponseAllOfTaskCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *UpdateTasks200ResponseAllOfTask) SetCredential(v UpdateTasks200ResponseAllOfTaskCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *UpdateTasks200ResponseAllOfTask) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateTasks200ResponseAllOfTask) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateTasks200ResponseAllOfTask) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateTasks200ResponseAllOfTask) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateTasks200ResponseAllOfTask) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateTasks200ResponseAllOfTask) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateTasks200ResponseAllOfTask) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateTasks200ResponseAllOfTask) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateTasks200ResponseAllOfTask) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


