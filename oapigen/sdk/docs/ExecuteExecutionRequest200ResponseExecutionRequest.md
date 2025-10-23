# ExecuteExecutionRequest200ResponseExecutionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **NullableString** |  | [optional] 
**ServerId** | Pointer to **NullableString** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**AppId** | Pointer to **NullableString** |  | [optional] 
**StdOut** | Pointer to **string** |  | [optional] 
**StdErr** | Pointer to **string** |  | [optional] 
**ExitCode** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**CreatedById** | Pointer to **int64** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**RawData** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExecuteExecutionRequest200ResponseExecutionRequest

`func NewExecuteExecutionRequest200ResponseExecutionRequest() *ExecuteExecutionRequest200ResponseExecutionRequest`

NewExecuteExecutionRequest200ResponseExecutionRequest instantiates a new ExecuteExecutionRequest200ResponseExecutionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteExecutionRequest200ResponseExecutionRequestWithDefaults

`func NewExecuteExecutionRequest200ResponseExecutionRequestWithDefaults() *ExecuteExecutionRequest200ResponseExecutionRequest`

NewExecuteExecutionRequest200ResponseExecutionRequestWithDefaults instantiates a new ExecuteExecutionRequest200ResponseExecutionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetContainerId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetServerId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetInstanceId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetResourceId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetAppId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetStdOut

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetStdOut() string`

GetStdOut returns the StdOut field if non-nil, zero value otherwise.

### GetStdOutOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetStdOutOk() (*string, bool)`

GetStdOutOk returns a tuple with the StdOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdOut

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetStdOut(v string)`

SetStdOut sets StdOut field to given value.

### HasStdOut

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasStdOut() bool`

HasStdOut returns a boolean if a field has been set.

### GetStdErr

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetStdErr() string`

GetStdErr returns the StdErr field if non-nil, zero value otherwise.

### GetStdErrOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetStdErrOk() (*string, bool)`

GetStdErrOk returns a tuple with the StdErr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdErr

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetStdErr(v string)`

SetStdErr sets StdErr field to given value.

### HasStdErr

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasStdErr() bool`

HasStdErr returns a boolean if a field has been set.

### GetExitCode

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetExitCode() int64`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetExitCodeOk() (*int64, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetExitCode(v int64)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetStatus

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetCreatedById() int64`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetCreatedByIdOk() (*int64, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetCreatedById(v int64)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetConfig

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRawData

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *ExecuteExecutionRequest200ResponseExecutionRequest) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


