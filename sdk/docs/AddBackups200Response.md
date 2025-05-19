# AddBackups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to [**ListBackups200ResponseAllOfBackupsInner**](ListBackups200ResponseAllOfBackupsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddBackups200Response

`func NewAddBackups200Response() *AddBackups200Response`

NewAddBackups200Response instantiates a new AddBackups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBackups200ResponseWithDefaults

`func NewAddBackups200ResponseWithDefaults() *AddBackups200Response`

NewAddBackups200ResponseWithDefaults instantiates a new AddBackups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *AddBackups200Response) GetBackup() ListBackups200ResponseAllOfBackupsInner`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *AddBackups200Response) GetBackupOk() (*ListBackups200ResponseAllOfBackupsInner, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *AddBackups200Response) SetBackup(v ListBackups200ResponseAllOfBackupsInner)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *AddBackups200Response) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetSuccess

`func (o *AddBackups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddBackups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddBackups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddBackups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


