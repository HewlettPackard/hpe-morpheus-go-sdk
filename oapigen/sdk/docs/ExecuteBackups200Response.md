# ExecuteBackups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to [**ExecuteBackups200ResponseAllOfBackup**](ExecuteBackups200ResponseAllOfBackup.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewExecuteBackups200Response

`func NewExecuteBackups200Response() *ExecuteBackups200Response`

NewExecuteBackups200Response instantiates a new ExecuteBackups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetBackup

`func (o *ExecuteBackups200Response) GetBackup() ExecuteBackups200ResponseAllOfBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *ExecuteBackups200Response) GetBackupOk() (*ExecuteBackups200ResponseAllOfBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *ExecuteBackups200Response) SetBackup(v ExecuteBackups200ResponseAllOfBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *ExecuteBackups200Response) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetSuccess

`func (o *ExecuteBackups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ExecuteBackups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ExecuteBackups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ExecuteBackups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


