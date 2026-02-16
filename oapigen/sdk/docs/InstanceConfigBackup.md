# InstanceConfigBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateBackup** | Pointer to **bool** |  | [optional] 
**JobAction** | Pointer to **string** |  | [optional] 
**JobRetentionCount** | Pointer to **string** |  | [optional] 
**ProviderBackupType** | Pointer to **int64** |  | [optional] 

## Methods

### NewInstanceConfigBackup

`func NewInstanceConfigBackup() *InstanceConfigBackup`

NewInstanceConfigBackup instantiates a new InstanceConfigBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConfigBackupWithDefaults

`func NewInstanceConfigBackupWithDefaults() *InstanceConfigBackup`

NewInstanceConfigBackupWithDefaults instantiates a new InstanceConfigBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateBackup

`func (o *InstanceConfigBackup) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *InstanceConfigBackup) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *InstanceConfigBackup) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *InstanceConfigBackup) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.

### GetJobAction

`func (o *InstanceConfigBackup) GetJobAction() string`

GetJobAction returns the JobAction field if non-nil, zero value otherwise.

### GetJobActionOk

`func (o *InstanceConfigBackup) GetJobActionOk() (*string, bool)`

GetJobActionOk returns a tuple with the JobAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobAction

`func (o *InstanceConfigBackup) SetJobAction(v string)`

SetJobAction sets JobAction field to given value.

### HasJobAction

`func (o *InstanceConfigBackup) HasJobAction() bool`

HasJobAction returns a boolean if a field has been set.

### GetJobRetentionCount

`func (o *InstanceConfigBackup) GetJobRetentionCount() string`

GetJobRetentionCount returns the JobRetentionCount field if non-nil, zero value otherwise.

### GetJobRetentionCountOk

`func (o *InstanceConfigBackup) GetJobRetentionCountOk() (*string, bool)`

GetJobRetentionCountOk returns a tuple with the JobRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRetentionCount

`func (o *InstanceConfigBackup) SetJobRetentionCount(v string)`

SetJobRetentionCount sets JobRetentionCount field to given value.

### HasJobRetentionCount

`func (o *InstanceConfigBackup) HasJobRetentionCount() bool`

HasJobRetentionCount returns a boolean if a field has been set.

### GetProviderBackupType

`func (o *InstanceConfigBackup) GetProviderBackupType() int64`

GetProviderBackupType returns the ProviderBackupType field if non-nil, zero value otherwise.

### GetProviderBackupTypeOk

`func (o *InstanceConfigBackup) GetProviderBackupTypeOk() (*int64, bool)`

GetProviderBackupTypeOk returns a tuple with the ProviderBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderBackupType

`func (o *InstanceConfigBackup) SetProviderBackupType(v int64)`

SetProviderBackupType sets ProviderBackupType field to given value.

### HasProviderBackupType

`func (o *InstanceConfigBackup) HasProviderBackupType() bool`

HasProviderBackupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


