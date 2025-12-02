# BackupCreationPolicyTypeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateBackupType** | **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict pattern) | 
**CreateBackup** | Pointer to **bool** | Enforce backup creation | [optional] 

## Methods

### NewBackupCreationPolicyTypeConfiguration

`func NewBackupCreationPolicyTypeConfiguration(createBackupType string, ) *BackupCreationPolicyTypeConfiguration`

NewBackupCreationPolicyTypeConfiguration instantiates a new BackupCreationPolicyTypeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupCreationPolicyTypeConfigurationWithDefaults

`func NewBackupCreationPolicyTypeConfigurationWithDefaults() *BackupCreationPolicyTypeConfiguration`

NewBackupCreationPolicyTypeConfigurationWithDefaults instantiates a new BackupCreationPolicyTypeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateBackupType

`func (o *BackupCreationPolicyTypeConfiguration) GetCreateBackupType() string`

GetCreateBackupType returns the CreateBackupType field if non-nil, zero value otherwise.

### GetCreateBackupTypeOk

`func (o *BackupCreationPolicyTypeConfiguration) GetCreateBackupTypeOk() (*string, bool)`

GetCreateBackupTypeOk returns a tuple with the CreateBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackupType

`func (o *BackupCreationPolicyTypeConfiguration) SetCreateBackupType(v string)`

SetCreateBackupType sets CreateBackupType field to given value.


### GetCreateBackup

`func (o *BackupCreationPolicyTypeConfiguration) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *BackupCreationPolicyTypeConfiguration) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *BackupCreationPolicyTypeConfiguration) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *BackupCreationPolicyTypeConfiguration) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


