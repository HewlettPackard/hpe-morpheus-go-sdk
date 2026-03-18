# MigrationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**SkippedPrechecks** | Pointer to **bool** | Indicates if the precheck should be skipped | [optional] 
**InstallGuestTools** | Pointer to **bool** | Indicates if guest tools should be installed | [optional] 
**ReInitializeServerOnMigration** | Pointer to **bool** | Indicates if the server should be re-initialized on migration | [optional] 
**LinuxUsername** | Pointer to **NullableString** | Linux Username for migrated servers | [optional] 
**LinuxPassword** | Pointer to **NullableString** | Linux Password for migrated servers | [optional] 
**LinuxKeyPair** | Pointer to [**MigrationUpdateLinuxKeyPair**](MigrationUpdateLinuxKeyPair.md) |  | [optional] 
**WindowsUsername** | Pointer to **NullableString** | Windows Username for migrated servers | [optional] 
**WindowsPassword** | Pointer to **NullableString** | Windows Password for migrated servers | [optional] 
**SourceCloudId** | Pointer to **int64** | Source Cloud ID. The API &#x60;/api/migrations/source-clouds&#x60; can be used to find available options.  | [optional] 
**TargetCloudId** | Pointer to **int64** | Target Cloud ID. The API &#x60;/api/migrations/target-clouds?sourceCloudId&#x3D;34&#x60; can be used to find available options.  | [optional] 
**TargetGroupId** | Pointer to **int64** | Target Group ID.  The Options API &#x60;/api/options/targetGroups?sourceCloudId&#x3D;34&amp;targetCloudId&#x3D;129&#x60; can be used to find available options.  | [optional] 
**TargetPoolId** | Pointer to [**MigrationUpdateTargetPoolId**](MigrationUpdateTargetPoolId.md) |  | [optional] 
**SourceServerIds** | Pointer to **[]int64** | Array of Server IDs to be migrated. The API &#x60;/api/migrations/source-servers?sourceCloudId&#x3D;34&#x60; can be used to find available options.  | [optional] 
**Datastores** | Pointer to [**[]MigrationUpdateDatastoresInner**](MigrationUpdateDatastoresInner.md) | Array of datastore mappings. | [optional] 
**Networks** | Pointer to [**[]MigrationUpdateNetworksInner**](MigrationUpdateNetworksInner.md) | Array of network mappings. | [optional] 

## Methods

### NewMigrationUpdate

`func NewMigrationUpdate() *MigrationUpdate`

NewMigrationUpdate instantiates a new MigrationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationUpdateWithDefaults

`func NewMigrationUpdateWithDefaults() *MigrationUpdate`

NewMigrationUpdateWithDefaults instantiates a new MigrationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MigrationUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MigrationUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MigrationUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MigrationUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSkippedPrechecks

`func (o *MigrationUpdate) GetSkippedPrechecks() bool`

GetSkippedPrechecks returns the SkippedPrechecks field if non-nil, zero value otherwise.

### GetSkippedPrechecksOk

`func (o *MigrationUpdate) GetSkippedPrechecksOk() (*bool, bool)`

GetSkippedPrechecksOk returns a tuple with the SkippedPrechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedPrechecks

`func (o *MigrationUpdate) SetSkippedPrechecks(v bool)`

SetSkippedPrechecks sets SkippedPrechecks field to given value.

### HasSkippedPrechecks

`func (o *MigrationUpdate) HasSkippedPrechecks() bool`

HasSkippedPrechecks returns a boolean if a field has been set.

### GetInstallGuestTools

`func (o *MigrationUpdate) GetInstallGuestTools() bool`

GetInstallGuestTools returns the InstallGuestTools field if non-nil, zero value otherwise.

### GetInstallGuestToolsOk

`func (o *MigrationUpdate) GetInstallGuestToolsOk() (*bool, bool)`

GetInstallGuestToolsOk returns a tuple with the InstallGuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallGuestTools

`func (o *MigrationUpdate) SetInstallGuestTools(v bool)`

SetInstallGuestTools sets InstallGuestTools field to given value.

### HasInstallGuestTools

`func (o *MigrationUpdate) HasInstallGuestTools() bool`

HasInstallGuestTools returns a boolean if a field has been set.

### GetReInitializeServerOnMigration

`func (o *MigrationUpdate) GetReInitializeServerOnMigration() bool`

GetReInitializeServerOnMigration returns the ReInitializeServerOnMigration field if non-nil, zero value otherwise.

### GetReInitializeServerOnMigrationOk

`func (o *MigrationUpdate) GetReInitializeServerOnMigrationOk() (*bool, bool)`

GetReInitializeServerOnMigrationOk returns a tuple with the ReInitializeServerOnMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReInitializeServerOnMigration

`func (o *MigrationUpdate) SetReInitializeServerOnMigration(v bool)`

SetReInitializeServerOnMigration sets ReInitializeServerOnMigration field to given value.

### HasReInitializeServerOnMigration

`func (o *MigrationUpdate) HasReInitializeServerOnMigration() bool`

HasReInitializeServerOnMigration returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *MigrationUpdate) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *MigrationUpdate) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *MigrationUpdate) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *MigrationUpdate) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *MigrationUpdate) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *MigrationUpdate) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *MigrationUpdate) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *MigrationUpdate) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *MigrationUpdate) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *MigrationUpdate) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *MigrationUpdate) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *MigrationUpdate) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPair

`func (o *MigrationUpdate) GetLinuxKeyPair() MigrationUpdateLinuxKeyPair`

GetLinuxKeyPair returns the LinuxKeyPair field if non-nil, zero value otherwise.

### GetLinuxKeyPairOk

`func (o *MigrationUpdate) GetLinuxKeyPairOk() (*MigrationUpdateLinuxKeyPair, bool)`

GetLinuxKeyPairOk returns a tuple with the LinuxKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPair

`func (o *MigrationUpdate) SetLinuxKeyPair(v MigrationUpdateLinuxKeyPair)`

SetLinuxKeyPair sets LinuxKeyPair field to given value.

### HasLinuxKeyPair

`func (o *MigrationUpdate) HasLinuxKeyPair() bool`

HasLinuxKeyPair returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *MigrationUpdate) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *MigrationUpdate) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *MigrationUpdate) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *MigrationUpdate) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *MigrationUpdate) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *MigrationUpdate) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *MigrationUpdate) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *MigrationUpdate) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *MigrationUpdate) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *MigrationUpdate) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *MigrationUpdate) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *MigrationUpdate) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetSourceCloudId

`func (o *MigrationUpdate) GetSourceCloudId() int64`

GetSourceCloudId returns the SourceCloudId field if non-nil, zero value otherwise.

### GetSourceCloudIdOk

`func (o *MigrationUpdate) GetSourceCloudIdOk() (*int64, bool)`

GetSourceCloudIdOk returns a tuple with the SourceCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCloudId

`func (o *MigrationUpdate) SetSourceCloudId(v int64)`

SetSourceCloudId sets SourceCloudId field to given value.

### HasSourceCloudId

`func (o *MigrationUpdate) HasSourceCloudId() bool`

HasSourceCloudId returns a boolean if a field has been set.

### GetTargetCloudId

`func (o *MigrationUpdate) GetTargetCloudId() int64`

GetTargetCloudId returns the TargetCloudId field if non-nil, zero value otherwise.

### GetTargetCloudIdOk

`func (o *MigrationUpdate) GetTargetCloudIdOk() (*int64, bool)`

GetTargetCloudIdOk returns a tuple with the TargetCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCloudId

`func (o *MigrationUpdate) SetTargetCloudId(v int64)`

SetTargetCloudId sets TargetCloudId field to given value.

### HasTargetCloudId

`func (o *MigrationUpdate) HasTargetCloudId() bool`

HasTargetCloudId returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *MigrationUpdate) GetTargetGroupId() int64`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *MigrationUpdate) GetTargetGroupIdOk() (*int64, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *MigrationUpdate) SetTargetGroupId(v int64)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *MigrationUpdate) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetTargetPoolId

`func (o *MigrationUpdate) GetTargetPoolId() MigrationUpdateTargetPoolId`

GetTargetPoolId returns the TargetPoolId field if non-nil, zero value otherwise.

### GetTargetPoolIdOk

`func (o *MigrationUpdate) GetTargetPoolIdOk() (*MigrationUpdateTargetPoolId, bool)`

GetTargetPoolIdOk returns a tuple with the TargetPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPoolId

`func (o *MigrationUpdate) SetTargetPoolId(v MigrationUpdateTargetPoolId)`

SetTargetPoolId sets TargetPoolId field to given value.

### HasTargetPoolId

`func (o *MigrationUpdate) HasTargetPoolId() bool`

HasTargetPoolId returns a boolean if a field has been set.

### GetSourceServerIds

`func (o *MigrationUpdate) GetSourceServerIds() []int64`

GetSourceServerIds returns the SourceServerIds field if non-nil, zero value otherwise.

### GetSourceServerIdsOk

`func (o *MigrationUpdate) GetSourceServerIdsOk() (*[]int64, bool)`

GetSourceServerIdsOk returns a tuple with the SourceServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceServerIds

`func (o *MigrationUpdate) SetSourceServerIds(v []int64)`

SetSourceServerIds sets SourceServerIds field to given value.

### HasSourceServerIds

`func (o *MigrationUpdate) HasSourceServerIds() bool`

HasSourceServerIds returns a boolean if a field has been set.

### GetDatastores

`func (o *MigrationUpdate) GetDatastores() []MigrationUpdateDatastoresInner`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *MigrationUpdate) GetDatastoresOk() (*[]MigrationUpdateDatastoresInner, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *MigrationUpdate) SetDatastores(v []MigrationUpdateDatastoresInner)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *MigrationUpdate) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetNetworks

`func (o *MigrationUpdate) GetNetworks() []MigrationUpdateNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *MigrationUpdate) GetNetworksOk() (*[]MigrationUpdateNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *MigrationUpdate) SetNetworks(v []MigrationUpdateNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *MigrationUpdate) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


