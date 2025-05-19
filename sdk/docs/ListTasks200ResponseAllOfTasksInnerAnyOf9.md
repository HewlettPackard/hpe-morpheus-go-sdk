# ListTasks200ResponseAllOfTasksInnerAnyOf9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**TaskType** | Pointer to [**ListTasks200ResponseAllOfTasksInnerAnyOf9TaskType**](ListTasks200ResponseAllOfTasksInnerAnyOf9TaskType.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**TaskOptions** | Pointer to [**ListTasks200ResponseAllOfTasksInnerAnyOf9TaskOptions**](ListTasks200ResponseAllOfTasksInnerAnyOf9TaskOptions.md) |  | [optional] 
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

### NewListTasks200ResponseAllOfTasksInnerAnyOf9

`func NewListTasks200ResponseAllOfTasksInnerAnyOf9() *ListTasks200ResponseAllOfTasksInnerAnyOf9`

NewListTasks200ResponseAllOfTasksInnerAnyOf9 instantiates a new ListTasks200ResponseAllOfTasksInnerAnyOf9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTasks200ResponseAllOfTasksInnerAnyOf9WithDefaults

`func NewListTasks200ResponseAllOfTasksInnerAnyOf9WithDefaults() *ListTasks200ResponseAllOfTasksInnerAnyOf9`

NewListTasks200ResponseAllOfTasksInnerAnyOf9WithDefaults instantiates a new ListTasks200ResponseAllOfTasksInnerAnyOf9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTaskType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetTaskType() ListTasks200ResponseAllOfTasksInnerAnyOf9TaskType`

GetTaskType returns the TaskType field if non-nil, zero value otherwise.

### GetTaskTypeOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetTaskTypeOk() (*ListTasks200ResponseAllOfTasksInnerAnyOf9TaskType, bool)`

GetTaskTypeOk returns a tuple with the TaskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetTaskType(v ListTasks200ResponseAllOfTasksInnerAnyOf9TaskType)`

SetTaskType sets TaskType field to given value.

### HasTaskType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasTaskType() bool`

HasTaskType returns a boolean if a field has been set.

### GetLabels

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetVisibility

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTaskOptions

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetTaskOptions() ListTasks200ResponseAllOfTasksInnerAnyOf9TaskOptions`

GetTaskOptions returns the TaskOptions field if non-nil, zero value otherwise.

### GetTaskOptionsOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetTaskOptionsOk() (*ListTasks200ResponseAllOfTasksInnerAnyOf9TaskOptions, bool)`

GetTaskOptionsOk returns a tuple with the TaskOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskOptions

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetTaskOptions(v ListTasks200ResponseAllOfTasksInnerAnyOf9TaskOptions)`

SetTaskOptions sets TaskOptions field to given value.

### HasTaskOptions

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasTaskOptions() bool`

HasTaskOptions returns a boolean if a field has been set.

### GetResultType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetExecuteTarget

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetExecuteTarget() string`

GetExecuteTarget returns the ExecuteTarget field if non-nil, zero value otherwise.

### GetExecuteTargetOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetExecuteTargetOk() (*string, bool)`

GetExecuteTargetOk returns a tuple with the ExecuteTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteTarget

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetExecuteTarget(v string)`

SetExecuteTarget sets ExecuteTarget field to given value.

### HasExecuteTarget

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasExecuteTarget() bool`

HasExecuteTarget returns a boolean if a field has been set.

### GetRetryable

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetRetryCount

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetRetryCount() int64`

GetRetryCount returns the RetryCount field if non-nil, zero value otherwise.

### GetRetryCountOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetRetryCountOk() (*int64, bool)`

GetRetryCountOk returns a tuple with the RetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryCount

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetRetryCount(v int64)`

SetRetryCount sets RetryCount field to given value.

### HasRetryCount

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasRetryCount() bool`

HasRetryCount returns a boolean if a field has been set.

### GetRetryDelaySeconds

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetRetryDelaySeconds() int64`

GetRetryDelaySeconds returns the RetryDelaySeconds field if non-nil, zero value otherwise.

### GetRetryDelaySecondsOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetRetryDelaySecondsOk() (*int64, bool)`

GetRetryDelaySecondsOk returns a tuple with the RetryDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelaySeconds

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetRetryDelaySeconds(v int64)`

SetRetryDelaySeconds sets RetryDelaySeconds field to given value.

### HasRetryDelaySeconds

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasRetryDelaySeconds() bool`

HasRetryDelaySeconds returns a boolean if a field has been set.

### GetAllowCustomConfig

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetAllowCustomConfig() bool`

GetAllowCustomConfig returns the AllowCustomConfig field if non-nil, zero value otherwise.

### GetAllowCustomConfigOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetAllowCustomConfigOk() (*bool, bool)`

GetAllowCustomConfigOk returns a tuple with the AllowCustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomConfig

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetAllowCustomConfig(v bool)`

SetAllowCustomConfig sets AllowCustomConfig field to given value.

### HasAllowCustomConfig

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasAllowCustomConfig() bool`

HasAllowCustomConfig returns a boolean if a field has been set.

### GetCredential

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetCredential() ListClouds200ResponseAllOfZonesInnerCredentialAnyOf`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredentialAnyOf, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredentialAnyOf)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOf9) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


