# ListTasks200ResponseAllOfTasksInnerAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**TaskType** | Pointer to [**ListTasks200ResponseAllOfTasksInnerAnyOfTaskType**](ListTasks200ResponseAllOfTasksInnerAnyOfTaskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**TaskOptions** | Pointer to [**ListTasks200ResponseAllOfTasksInnerAnyOfTaskOptions**](ListTasks200ResponseAllOfTasksInnerAnyOfTaskOptions.md) |  | [optional] 
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

### NewListTasks200ResponseAllOfTasksInnerAnyOf

`func NewListTasks200ResponseAllOfTasksInnerAnyOf() *ListTasks200ResponseAllOfTasksInnerAnyOf`

NewListTasks200ResponseAllOfTasksInnerAnyOf instantiates a new ListTasks200ResponseAllOfTasksInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTasks200ResponseAllOfTasksInnerAnyOfWithDefaults

`func NewListTasks200ResponseAllOfTasksInnerAnyOfWithDefaults() *ListTasks200ResponseAllOfTasksInnerAnyOf`

NewListTasks200ResponseAllOfTasksInnerAnyOfWithDefaults instantiates a new ListTasks200ResponseAllOfTasksInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTaskType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetTaskType() ListTasks200ResponseAllOfTasksInnerAnyOfTaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetTaskTypeOk() (*ListTasks200ResponseAllOfTasksInnerAnyOfTaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetTaskType(v ListTasks200ResponseAllOfTasksInnerAnyOfTaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVisibility

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskOptions

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetTaskOptions() ListTasks200ResponseAllOfTasksInnerAnyOfTaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetTaskOptionsOk() (*ListTasks200ResponseAllOfTasksInnerAnyOfTaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetTaskOptions(v ListTasks200ResponseAllOfTasksInnerAnyOfTaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetFile

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetFile() ListTasks200ResponseAllOfTasksInnerAnyOfFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetFileOk() (*ListTasks200ResponseAllOfTasksInnerAnyOfFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetFile(v ListTasks200ResponseAllOfTasksInnerAnyOfFile)`

SetFile sets File field to given value.

### HasFile

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetResultType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetExecuteTarget

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetCredential

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetCredential() ListClouds200ResponseAllOfZonesInnerCredentialAnyOf`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredentialAnyOf, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredentialAnyOf)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


