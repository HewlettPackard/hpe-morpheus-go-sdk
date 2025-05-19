# ExecutionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**ResourceId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**StdOut** | Pointer to **string** |  | [optional] 
**StdErr** | Pointer to **string** |  | [optional] 
**ExitCode** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**CreatedById** | Pointer to **int64** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**RawData** | Pointer to **string** |  | [optional] 

## Methods

### NewExecutionRequest

`func NewExecutionRequest() *ExecutionRequest`

NewExecutionRequest instantiates a new ExecutionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionRequestWithDefaults

`func NewExecutionRequestWithDefaults() *ExecutionRequest`

NewExecutionRequestWithDefaults instantiates a new ExecutionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExecutionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ExecutionRequest) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ExecutionRequest) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ExecutionRequest) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ExecutionRequest) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetContainerId

`func (o *ExecutionRequest) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ExecutionRequest) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ExecutionRequest) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ExecutionRequest) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *ExecutionRequest) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ExecutionRequest) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ExecutionRequest) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ExecutionRequest) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetInstanceId

`func (o *ExecutionRequest) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ExecutionRequest) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ExecutionRequest) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ExecutionRequest) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetResourceId

`func (o *ExecutionRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ExecutionRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ExecutionRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ExecutionRequest) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetAppId

`func (o *ExecutionRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ExecutionRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ExecutionRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ExecutionRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetStdOut

`func (o *ExecutionRequest) GetStdOut() string`

GetStdOut returns the StdOut field if non-nil, zero value otherwise.

### GetStdOutOk

`func (o *ExecutionRequest) GetStdOutOk() (*string, bool)`

GetStdOutOk returns a tuple with the StdOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdOut

`func (o *ExecutionRequest) SetStdOut(v string)`

SetStdOut sets StdOut field to given value.

### HasStdOut

`func (o *ExecutionRequest) HasStdOut() bool`

HasStdOut returns a boolean if a field has been set.

### GetStdErr

`func (o *ExecutionRequest) GetStdErr() string`

GetStdErr returns the StdErr field if non-nil, zero value otherwise.

### GetStdErrOk

`func (o *ExecutionRequest) GetStdErrOk() (*string, bool)`

GetStdErrOk returns a tuple with the StdErr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdErr

`func (o *ExecutionRequest) SetStdErr(v string)`

SetStdErr sets StdErr field to given value.

### HasStdErr

`func (o *ExecutionRequest) HasStdErr() bool`

HasStdErr returns a boolean if a field has been set.

### GetExitCode

`func (o *ExecutionRequest) GetExitCode() int64`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ExecutionRequest) GetExitCodeOk() (*int64, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ExecutionRequest) SetExitCode(v int64)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *ExecutionRequest) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ExecutionRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ExecutionRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ExecutionRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ExecutionRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *ExecutionRequest) GetCreatedById() int64`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ExecutionRequest) GetCreatedByIdOk() (*int64, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ExecutionRequest) SetCreatedById(v int64)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *ExecutionRequest) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ExecutionRequest) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ExecutionRequest) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ExecutionRequest) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ExecutionRequest) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ExecutionRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ExecutionRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ExecutionRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ExecutionRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetConfig

`func (o *ExecutionRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExecutionRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExecutionRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ExecutionRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRawData

`func (o *ExecutionRequest) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *ExecutionRequest) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *ExecutionRequest) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *ExecutionRequest) HasRawData() bool`

HasRawData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


