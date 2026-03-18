# UpdateBackups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to [**UpdateBackups200ResponseAllOfBackup**](UpdateBackups200ResponseAllOfBackup.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateBackups200Response

`func NewUpdateBackups200Response() *UpdateBackups200Response`

NewUpdateBackups200Response instantiates a new UpdateBackups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBackups200ResponseWithDefaults

`func NewUpdateBackups200ResponseWithDefaults() *UpdateBackups200Response`

NewUpdateBackups200ResponseWithDefaults instantiates a new UpdateBackups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *UpdateBackups200Response) GetBackup() UpdateBackups200ResponseAllOfBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *UpdateBackups200Response) GetBackupOk() (*UpdateBackups200ResponseAllOfBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *UpdateBackups200Response) SetBackup(v UpdateBackups200ResponseAllOfBackup)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *UpdateBackups200Response) HasBackup() bool`

HasBackup returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateBackups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateBackups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateBackups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateBackups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


