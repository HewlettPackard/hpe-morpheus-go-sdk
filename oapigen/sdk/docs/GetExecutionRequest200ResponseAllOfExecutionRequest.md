# GetExecutionRequest200ResponseAllOfExecutionRequest

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

### NewGetExecutionRequest200ResponseAllOfExecutionRequest

`func NewGetExecutionRequest200ResponseAllOfExecutionRequest() *GetExecutionRequest200ResponseAllOfExecutionRequest`

NewGetExecutionRequest200ResponseAllOfExecutionRequest instantiates a new GetExecutionRequest200ResponseAllOfExecutionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExecutionRequest200ResponseAllOfExecutionRequestWithDefaults

`func NewGetExecutionRequest200ResponseAllOfExecutionRequestWithDefaults() *GetExecutionRequest200ResponseAllOfExecutionRequest`

NewGetExecutionRequest200ResponseAllOfExecutionRequestWithDefaults instantiates a new GetExecutionRequest200ResponseAllOfExecutionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUniqueId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetContainerId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetServerId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetInstanceId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetResourceId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetAppId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetStdOut

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetStdOut() string`

GetStdOut returns the StdOut field if non-nil, zero value otherwise.

### GetStdOutOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetStdOutOk() (*string, bool)`

GetStdOutOk returns a tuple with the StdOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdOut

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetStdOut(v string)`

SetStdOut sets StdOut field to given value.

### HasStdOut

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasStdOut() bool`

HasStdOut returns a boolean if a field has been set.

### GetStdErr

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetStdErr() string`

GetStdErr returns the StdErr field if non-nil, zero value otherwise.

### GetStdErrOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetStdErrOk() (*string, bool)`

GetStdErrOk returns a tuple with the StdErr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdErr

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetStdErr(v string)`

SetStdErr sets StdErr field to given value.

### HasStdErr

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasStdErr() bool`

HasStdErr returns a boolean if a field has been set.

### GetExitCode

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetExitCode() int64`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetExitCodeOk() (*int64, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetExitCode(v int64)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetStatus

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedById

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetCreatedById() int64`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetCreatedByIdOk() (*int64, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetCreatedById(v int64)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetConfig

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRawData

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### SetRawDataNil

`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) SetRawDataNil(b bool)`

 SetRawDataNil sets the value for RawData to be an explicit nil

### UnsetRawData
`func (o *GetExecutionRequest200ResponseAllOfExecutionRequest) UnsetRawData()`

UnsetRawData ensures that no value is present for RawData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


