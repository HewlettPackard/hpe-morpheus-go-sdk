# ExecuteBackupRestoreRequestRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupResultId** | **int64** | Id of backup result | 
**RestoreInstance** | **string** | Type of restore | 
**InstanceId** | Pointer to **int64** | Id of instance | [optional] 
**Config** | Pointer to **map[string]interface{}** | Additional config | [optional] 
**InstanceConfig** | Pointer to **map[string]interface{}** | Instance config for restore to new. | [optional] 

## Methods

### NewExecuteBackupRestoreRequestRestore

`func NewExecuteBackupRestoreRequestRestore(backupResultId int64, restoreInstance string, ) *ExecuteBackupRestoreRequestRestore`

NewExecuteBackupRestoreRequestRestore instantiates a new ExecuteBackupRestoreRequestRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteBackupRestoreRequestRestoreWithDefaults

`func NewExecuteBackupRestoreRequestRestoreWithDefaults() *ExecuteBackupRestoreRequestRestore`

NewExecuteBackupRestoreRequestRestoreWithDefaults instantiates a new ExecuteBackupRestoreRequestRestore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupResultId

`func (o *ExecuteBackupRestoreRequestRestore) GetBackupResultId() int64`

GetBackupResultId returns the BackupResultId field if non-nil, zero value otherwise.

### GetBackupResultIdOk

`func (o *ExecuteBackupRestoreRequestRestore) GetBackupResultIdOk() (*int64, bool)`

GetBackupResultIdOk returns a tuple with the BackupResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupResultId

`func (o *ExecuteBackupRestoreRequestRestore) SetBackupResultId(v int64)`

SetBackupResultId sets BackupResultId field to given value.


### GetRestoreInstance

`func (o *ExecuteBackupRestoreRequestRestore) GetRestoreInstance() string`

GetRestoreInstance returns the RestoreInstance field if non-nil, zero value otherwise.

### GetRestoreInstanceOk

`func (o *ExecuteBackupRestoreRequestRestore) GetRestoreInstanceOk() (*string, bool)`

GetRestoreInstanceOk returns a tuple with the RestoreInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreInstance

`func (o *ExecuteBackupRestoreRequestRestore) SetRestoreInstance(v string)`

SetRestoreInstance sets RestoreInstance field to given value.


### GetInstanceId

`func (o *ExecuteBackupRestoreRequestRestore) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ExecuteBackupRestoreRequestRestore) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ExecuteBackupRestoreRequestRestore) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ExecuteBackupRestoreRequestRestore) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetConfig

`func (o *ExecuteBackupRestoreRequestRestore) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ExecuteBackupRestoreRequestRestore) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ExecuteBackupRestoreRequestRestore) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ExecuteBackupRestoreRequestRestore) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetInstanceConfig

`func (o *ExecuteBackupRestoreRequestRestore) GetInstanceConfig() map[string]interface{}`

GetInstanceConfig returns the InstanceConfig field if non-nil, zero value otherwise.

### GetInstanceConfigOk

`func (o *ExecuteBackupRestoreRequestRestore) GetInstanceConfigOk() (*map[string]interface{}, bool)`

GetInstanceConfigOk returns a tuple with the InstanceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceConfig

`func (o *ExecuteBackupRestoreRequestRestore) SetInstanceConfig(v map[string]interface{})`

SetInstanceConfig sets InstanceConfig field to given value.

### HasInstanceConfig

`func (o *ExecuteBackupRestoreRequestRestore) HasInstanceConfig() bool`

HasInstanceConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


