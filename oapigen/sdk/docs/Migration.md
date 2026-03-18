# Migration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Migration ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Status** | Pointer to **string** | Migration Status. The possible status values are: &#39;pending&#39;, &#39;running&#39;, &#39;failed&#39;, &#39;completed&#39; | [optional] 
**StatusMessage** | Pointer to **NullableString** | Status Message | [optional] 
**SkippedPrechecks** | Pointer to **bool** | Indicates if the precheck was skipped | [optional] 
**InstallGuestTools** | Pointer to **bool** | Indicates if guest tools should be installed | [optional] 
**ReInitializeServerOnMigration** | Pointer to **bool** | Indicates if the server should be re-initialized on migration | [optional] 
**LinuxUsername** | Pointer to **NullableString** | Linux Username for migrated servers | [optional] 
**LinuxPassword** | Pointer to **NullableString** | Linux Password for migrated servers | [optional] 
**LinuxKeyPair** | Pointer to [**AddMigration200ResponseAnyOfMigrationLinuxKeyPair**](AddMigration200ResponseAnyOfMigrationLinuxKeyPair.md) |  | [optional] 
**WindowsUsername** | Pointer to **NullableString** | Windows Username for migrated servers | [optional] 
**WindowsPassword** | Pointer to **NullableString** | Windows Password for migrated servers | [optional] 
**SourceCloud** | Pointer to [**AddMigration200ResponseAnyOfMigrationSourceCloud**](AddMigration200ResponseAnyOfMigrationSourceCloud.md) |  | [optional] 
**TargetCloud** | Pointer to [**AddMigration200ResponseAnyOfMigrationTargetCloud**](AddMigration200ResponseAnyOfMigrationTargetCloud.md) |  | [optional] 
**TargetGroup** | Pointer to [**AddMigration200ResponseAnyOfMigrationTargetGroup**](AddMigration200ResponseAnyOfMigrationTargetGroup.md) |  | [optional] 
**TargetPool** | Pointer to [**AddMigration200ResponseAnyOfMigrationTargetPool**](AddMigration200ResponseAnyOfMigrationTargetPool.md) |  | [optional] 
**Servers** | Pointer to [**[]AddMigration200ResponseAnyOfMigrationServersInner**](AddMigration200ResponseAnyOfMigrationServersInner.md) |  | [optional] 

## Methods

### NewMigration

`func NewMigration() *Migration`

NewMigration instantiates a new Migration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationWithDefaults

`func NewMigrationWithDefaults() *Migration`

NewMigrationWithDefaults instantiates a new Migration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Migration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Migration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Migration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Migration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Migration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Migration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Migration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Migration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *Migration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Migration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Migration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Migration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Migration) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Migration) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Migration) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Migration) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *Migration) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *Migration) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetSkippedPrechecks

`func (o *Migration) GetSkippedPrechecks() bool`

GetSkippedPrechecks returns the SkippedPrechecks field if non-nil, zero value otherwise.

### GetSkippedPrechecksOk

`func (o *Migration) GetSkippedPrechecksOk() (*bool, bool)`

GetSkippedPrechecksOk returns a tuple with the SkippedPrechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedPrechecks

`func (o *Migration) SetSkippedPrechecks(v bool)`

SetSkippedPrechecks sets SkippedPrechecks field to given value.

### HasSkippedPrechecks

`func (o *Migration) HasSkippedPrechecks() bool`

HasSkippedPrechecks returns a boolean if a field has been set.

### GetInstallGuestTools

`func (o *Migration) GetInstallGuestTools() bool`

GetInstallGuestTools returns the InstallGuestTools field if non-nil, zero value otherwise.

### GetInstallGuestToolsOk

`func (o *Migration) GetInstallGuestToolsOk() (*bool, bool)`

GetInstallGuestToolsOk returns a tuple with the InstallGuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallGuestTools

`func (o *Migration) SetInstallGuestTools(v bool)`

SetInstallGuestTools sets InstallGuestTools field to given value.

### HasInstallGuestTools

`func (o *Migration) HasInstallGuestTools() bool`

HasInstallGuestTools returns a boolean if a field has been set.

### GetReInitializeServerOnMigration

`func (o *Migration) GetReInitializeServerOnMigration() bool`

GetReInitializeServerOnMigration returns the ReInitializeServerOnMigration field if non-nil, zero value otherwise.

### GetReInitializeServerOnMigrationOk

`func (o *Migration) GetReInitializeServerOnMigrationOk() (*bool, bool)`

GetReInitializeServerOnMigrationOk returns a tuple with the ReInitializeServerOnMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReInitializeServerOnMigration

`func (o *Migration) SetReInitializeServerOnMigration(v bool)`

SetReInitializeServerOnMigration sets ReInitializeServerOnMigration field to given value.

### HasReInitializeServerOnMigration

`func (o *Migration) HasReInitializeServerOnMigration() bool`

HasReInitializeServerOnMigration returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *Migration) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *Migration) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *Migration) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *Migration) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *Migration) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *Migration) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *Migration) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *Migration) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *Migration) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *Migration) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *Migration) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *Migration) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPair

`func (o *Migration) GetLinuxKeyPair() AddMigration200ResponseAnyOfMigrationLinuxKeyPair`

GetLinuxKeyPair returns the LinuxKeyPair field if non-nil, zero value otherwise.

### GetLinuxKeyPairOk

`func (o *Migration) GetLinuxKeyPairOk() (*AddMigration200ResponseAnyOfMigrationLinuxKeyPair, bool)`

GetLinuxKeyPairOk returns a tuple with the LinuxKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPair

`func (o *Migration) SetLinuxKeyPair(v AddMigration200ResponseAnyOfMigrationLinuxKeyPair)`

SetLinuxKeyPair sets LinuxKeyPair field to given value.

### HasLinuxKeyPair

`func (o *Migration) HasLinuxKeyPair() bool`

HasLinuxKeyPair returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *Migration) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *Migration) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *Migration) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *Migration) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *Migration) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *Migration) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *Migration) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *Migration) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *Migration) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *Migration) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *Migration) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *Migration) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetSourceCloud

`func (o *Migration) GetSourceCloud() AddMigration200ResponseAnyOfMigrationSourceCloud`

GetSourceCloud returns the SourceCloud field if non-nil, zero value otherwise.

### GetSourceCloudOk

`func (o *Migration) GetSourceCloudOk() (*AddMigration200ResponseAnyOfMigrationSourceCloud, bool)`

GetSourceCloudOk returns a tuple with the SourceCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCloud

`func (o *Migration) SetSourceCloud(v AddMigration200ResponseAnyOfMigrationSourceCloud)`

SetSourceCloud sets SourceCloud field to given value.

### HasSourceCloud

`func (o *Migration) HasSourceCloud() bool`

HasSourceCloud returns a boolean if a field has been set.

### GetTargetCloud

`func (o *Migration) GetTargetCloud() AddMigration200ResponseAnyOfMigrationTargetCloud`

GetTargetCloud returns the TargetCloud field if non-nil, zero value otherwise.

### GetTargetCloudOk

`func (o *Migration) GetTargetCloudOk() (*AddMigration200ResponseAnyOfMigrationTargetCloud, bool)`

GetTargetCloudOk returns a tuple with the TargetCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCloud

`func (o *Migration) SetTargetCloud(v AddMigration200ResponseAnyOfMigrationTargetCloud)`

SetTargetCloud sets TargetCloud field to given value.

### HasTargetCloud

`func (o *Migration) HasTargetCloud() bool`

HasTargetCloud returns a boolean if a field has been set.

### GetTargetGroup

`func (o *Migration) GetTargetGroup() AddMigration200ResponseAnyOfMigrationTargetGroup`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *Migration) GetTargetGroupOk() (*AddMigration200ResponseAnyOfMigrationTargetGroup, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *Migration) SetTargetGroup(v AddMigration200ResponseAnyOfMigrationTargetGroup)`

SetTargetGroup sets TargetGroup field to given value.

### HasTargetGroup

`func (o *Migration) HasTargetGroup() bool`

HasTargetGroup returns a boolean if a field has been set.

### GetTargetPool

`func (o *Migration) GetTargetPool() AddMigration200ResponseAnyOfMigrationTargetPool`

GetTargetPool returns the TargetPool field if non-nil, zero value otherwise.

### GetTargetPoolOk

`func (o *Migration) GetTargetPoolOk() (*AddMigration200ResponseAnyOfMigrationTargetPool, bool)`

GetTargetPoolOk returns a tuple with the TargetPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPool

`func (o *Migration) SetTargetPool(v AddMigration200ResponseAnyOfMigrationTargetPool)`

SetTargetPool sets TargetPool field to given value.

### HasTargetPool

`func (o *Migration) HasTargetPool() bool`

HasTargetPool returns a boolean if a field has been set.

### GetServers

`func (o *Migration) GetServers() []AddMigration200ResponseAnyOfMigrationServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *Migration) GetServersOk() (*[]AddMigration200ResponseAnyOfMigrationServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *Migration) SetServers(v []AddMigration200ResponseAnyOfMigrationServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *Migration) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


