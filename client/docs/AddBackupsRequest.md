# AddBackupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | [**AddBackupsRequestBackup**](AddBackupsRequestBackup.md) |  | 

## Methods

### NewAddBackupsRequest

`func NewAddBackupsRequest(backup AddBackupsRequestBackup, ) *AddBackupsRequest`

NewAddBackupsRequest instantiates a new AddBackupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBackupsRequestWithDefaults

`func NewAddBackupsRequestWithDefaults() *AddBackupsRequest`

NewAddBackupsRequestWithDefaults instantiates a new AddBackupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *AddBackupsRequest) GetBackup() AddBackupsRequestBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *AddBackupsRequest) GetBackupOk() (*AddBackupsRequestBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *AddBackupsRequest) SetBackup(v AddBackupsRequestBackup)`

SetBackup sets Backup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


