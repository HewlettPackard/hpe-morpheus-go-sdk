# GetBackups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | Pointer to [**ListBackups200ResponseAllOfBackupsInner**](ListBackups200ResponseAllOfBackupsInner.md) |  | [optional] 

## Methods

### NewGetBackups200Response

`func NewGetBackups200Response() *GetBackups200Response`

NewGetBackups200Response instantiates a new GetBackups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBackups200ResponseWithDefaults

`func NewGetBackups200ResponseWithDefaults() *GetBackups200Response`

NewGetBackups200ResponseWithDefaults instantiates a new GetBackups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *GetBackups200Response) GetBackup() ListBackups200ResponseAllOfBackupsInner`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *GetBackups200Response) GetBackupOk() (*ListBackups200ResponseAllOfBackupsInner, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *GetBackups200Response) SetBackup(v ListBackups200ResponseAllOfBackupsInner)`

SetBackup sets Backup field to given value.

### HasBackup

`func (o *GetBackups200Response) HasBackup() bool`

HasBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


