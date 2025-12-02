# BackupCreationPolicyTypeConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateBackupType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict pattern) | [optional] 
**CreateBackup** | Pointer to **bool** | Enforce backup creation | [optional] 

## Methods

### NewBackupCreationPolicyTypeConfiguration1

`func NewBackupCreationPolicyTypeConfiguration1() *BackupCreationPolicyTypeConfiguration1`

NewBackupCreationPolicyTypeConfiguration1 instantiates a new BackupCreationPolicyTypeConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupCreationPolicyTypeConfiguration1WithDefaults

`func NewBackupCreationPolicyTypeConfiguration1WithDefaults() *BackupCreationPolicyTypeConfiguration1`

NewBackupCreationPolicyTypeConfiguration1WithDefaults instantiates a new BackupCreationPolicyTypeConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateBackupType

`func (o *BackupCreationPolicyTypeConfiguration1) GetCreateBackupType() string`

GetCreateBackupType returns the CreateBackupType field if non-nil, zero value otherwise.

### GetCreateBackupTypeOk

`func (o *BackupCreationPolicyTypeConfiguration1) GetCreateBackupTypeOk() (*string, bool)`

GetCreateBackupTypeOk returns a tuple with the CreateBackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackupType

`func (o *BackupCreationPolicyTypeConfiguration1) SetCreateBackupType(v string)`

SetCreateBackupType sets CreateBackupType field to given value.

### HasCreateBackupType

`func (o *BackupCreationPolicyTypeConfiguration1) HasCreateBackupType() bool`

HasCreateBackupType returns a boolean if a field has been set.

### GetCreateBackup

`func (o *BackupCreationPolicyTypeConfiguration1) GetCreateBackup() bool`

GetCreateBackup returns the CreateBackup field if non-nil, zero value otherwise.

### GetCreateBackupOk

`func (o *BackupCreationPolicyTypeConfiguration1) GetCreateBackupOk() (*bool, bool)`

GetCreateBackupOk returns a tuple with the CreateBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBackup

`func (o *BackupCreationPolicyTypeConfiguration1) SetCreateBackup(v bool)`

SetCreateBackup sets CreateBackup field to given value.

### HasCreateBackup

`func (o *BackupCreationPolicyTypeConfiguration1) HasCreateBackup() bool`

HasCreateBackup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


