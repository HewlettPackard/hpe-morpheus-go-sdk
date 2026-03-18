# ExecuteBackupRestore200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to [**ExecuteBackups200ResponseAllOfBackup**](ExecuteBackups200ResponseAllOfBackup.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewExecuteBackupRestore200Response

`func NewExecuteBackupRestore200Response() *ExecuteBackupRestore200Response`

NewExecuteBackupRestore200Response instantiates a new ExecuteBackupRestore200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteBackupRestore200ResponseWithDefaults

`func NewExecuteBackupRestore200ResponseWithDefaults() *ExecuteBackupRestore200Response`

NewExecuteBackupRestore200ResponseWithDefaults instantiates a new ExecuteBackupRestore200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *ExecuteBackupRestore200Response) GetBackup() ExecuteBackups200ResponseAllOfBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *ExecuteBackupRestore200Response) GetBackupOk() (*ExecuteBackups200ResponseAllOfBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *ExecuteBackupRestore200Response) SetBackup(v ExecuteBackups200ResponseAllOfBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *ExecuteBackupRestore200Response) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetSuccess

`func (o *ExecuteBackupRestore200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ExecuteBackupRestore200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ExecuteBackupRestore200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ExecuteBackupRestore200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


