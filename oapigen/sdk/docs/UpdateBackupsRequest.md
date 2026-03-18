# UpdateBackupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | [**UpdateBackupsRequestBackup**](UpdateBackupsRequestBackup.md) |  | 

## Methods

### NewUpdateBackupsRequest

`func NewUpdateBackupsRequest(backup UpdateBackupsRequestBackup, ) *UpdateBackupsRequest`

NewUpdateBackupsRequest instantiates a new UpdateBackupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBackupsRequestWithDefaults

`func NewUpdateBackupsRequestWithDefaults() *UpdateBackupsRequest`

NewUpdateBackupsRequestWithDefaults instantiates a new UpdateBackupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *UpdateBackupsRequest) GetBackup() UpdateBackupsRequestBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *UpdateBackupsRequest) GetBackupOk() (*UpdateBackupsRequestBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *UpdateBackupsRequest) SetBackup(v UpdateBackupsRequestBackup)`

SetBackup sets Backup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


