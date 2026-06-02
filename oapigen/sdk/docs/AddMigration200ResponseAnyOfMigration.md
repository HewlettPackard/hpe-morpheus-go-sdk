# AddMigration200ResponseAnyOfMigration

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

### NewAddMigration200ResponseAnyOfMigration

`func NewAddMigration200ResponseAnyOfMigration() *AddMigration200ResponseAnyOfMigration`

NewAddMigration200ResponseAnyOfMigration instantiates a new AddMigration200ResponseAnyOfMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddMigration200ResponseAnyOfMigration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddMigration200ResponseAnyOfMigration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddMigration200ResponseAnyOfMigration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddMigration200ResponseAnyOfMigration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddMigration200ResponseAnyOfMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddMigration200ResponseAnyOfMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddMigration200ResponseAnyOfMigration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddMigration200ResponseAnyOfMigration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AddMigration200ResponseAnyOfMigration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddMigration200ResponseAnyOfMigration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddMigration200ResponseAnyOfMigration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddMigration200ResponseAnyOfMigration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddMigration200ResponseAnyOfMigration) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddMigration200ResponseAnyOfMigration) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddMigration200ResponseAnyOfMigration) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddMigration200ResponseAnyOfMigration) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AddMigration200ResponseAnyOfMigration) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AddMigration200ResponseAnyOfMigration) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetSkippedPrechecks

`func (o *AddMigration200ResponseAnyOfMigration) GetSkippedPrechecks() bool`

GetSkippedPrechecks returns the SkippedPrechecks field if non-nil, zero value otherwise.

### GetSkippedPrechecksOk

`func (o *AddMigration200ResponseAnyOfMigration) GetSkippedPrechecksOk() (*bool, bool)`

GetSkippedPrechecksOk returns a tuple with the SkippedPrechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedPrechecks

`func (o *AddMigration200ResponseAnyOfMigration) SetSkippedPrechecks(v bool)`

SetSkippedPrechecks sets SkippedPrechecks field to given value.

### HasSkippedPrechecks

`func (o *AddMigration200ResponseAnyOfMigration) HasSkippedPrechecks() bool`

HasSkippedPrechecks returns a boolean if a field has been set.

### GetInstallGuestTools

`func (o *AddMigration200ResponseAnyOfMigration) GetInstallGuestTools() bool`

GetInstallGuestTools returns the InstallGuestTools field if non-nil, zero value otherwise.

### GetInstallGuestToolsOk

`func (o *AddMigration200ResponseAnyOfMigration) GetInstallGuestToolsOk() (*bool, bool)`

GetInstallGuestToolsOk returns a tuple with the InstallGuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallGuestTools

`func (o *AddMigration200ResponseAnyOfMigration) SetInstallGuestTools(v bool)`

SetInstallGuestTools sets InstallGuestTools field to given value.

### HasInstallGuestTools

`func (o *AddMigration200ResponseAnyOfMigration) HasInstallGuestTools() bool`

HasInstallGuestTools returns a boolean if a field has been set.

### GetReInitializeServerOnMigration

`func (o *AddMigration200ResponseAnyOfMigration) GetReInitializeServerOnMigration() bool`

GetReInitializeServerOnMigration returns the ReInitializeServerOnMigration field if non-nil, zero value otherwise.

### GetReInitializeServerOnMigrationOk

`func (o *AddMigration200ResponseAnyOfMigration) GetReInitializeServerOnMigrationOk() (*bool, bool)`

GetReInitializeServerOnMigrationOk returns a tuple with the ReInitializeServerOnMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReInitializeServerOnMigration

`func (o *AddMigration200ResponseAnyOfMigration) SetReInitializeServerOnMigration(v bool)`

SetReInitializeServerOnMigration sets ReInitializeServerOnMigration field to given value.

### HasReInitializeServerOnMigration

`func (o *AddMigration200ResponseAnyOfMigration) HasReInitializeServerOnMigration() bool`

HasReInitializeServerOnMigration returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *AddMigration200ResponseAnyOfMigration) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *AddMigration200ResponseAnyOfMigration) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *AddMigration200ResponseAnyOfMigration) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *AddMigration200ResponseAnyOfMigration) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *AddMigration200ResponseAnyOfMigration) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *AddMigration200ResponseAnyOfMigration) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *AddMigration200ResponseAnyOfMigration) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *AddMigration200ResponseAnyOfMigration) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *AddMigration200ResponseAnyOfMigration) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *AddMigration200ResponseAnyOfMigration) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *AddMigration200ResponseAnyOfMigration) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *AddMigration200ResponseAnyOfMigration) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPair

`func (o *AddMigration200ResponseAnyOfMigration) GetLinuxKeyPair() AddMigration200ResponseAnyOfMigrationLinuxKeyPair`

GetLinuxKeyPair returns the LinuxKeyPair field if non-nil, zero value otherwise.

### GetLinuxKeyPairOk

`func (o *AddMigration200ResponseAnyOfMigration) GetLinuxKeyPairOk() (*AddMigration200ResponseAnyOfMigrationLinuxKeyPair, bool)`

GetLinuxKeyPairOk returns a tuple with the LinuxKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPair

`func (o *AddMigration200ResponseAnyOfMigration) SetLinuxKeyPair(v AddMigration200ResponseAnyOfMigrationLinuxKeyPair)`

SetLinuxKeyPair sets LinuxKeyPair field to given value.

### HasLinuxKeyPair

`func (o *AddMigration200ResponseAnyOfMigration) HasLinuxKeyPair() bool`

HasLinuxKeyPair returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *AddMigration200ResponseAnyOfMigration) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *AddMigration200ResponseAnyOfMigration) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *AddMigration200ResponseAnyOfMigration) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *AddMigration200ResponseAnyOfMigration) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *AddMigration200ResponseAnyOfMigration) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *AddMigration200ResponseAnyOfMigration) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *AddMigration200ResponseAnyOfMigration) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *AddMigration200ResponseAnyOfMigration) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *AddMigration200ResponseAnyOfMigration) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *AddMigration200ResponseAnyOfMigration) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *AddMigration200ResponseAnyOfMigration) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *AddMigration200ResponseAnyOfMigration) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetSourceCloud

`func (o *AddMigration200ResponseAnyOfMigration) GetSourceCloud() AddMigration200ResponseAnyOfMigrationSourceCloud`

GetSourceCloud returns the SourceCloud field if non-nil, zero value otherwise.

### GetSourceCloudOk

`func (o *AddMigration200ResponseAnyOfMigration) GetSourceCloudOk() (*AddMigration200ResponseAnyOfMigrationSourceCloud, bool)`

GetSourceCloudOk returns a tuple with the SourceCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCloud

`func (o *AddMigration200ResponseAnyOfMigration) SetSourceCloud(v AddMigration200ResponseAnyOfMigrationSourceCloud)`

SetSourceCloud sets SourceCloud field to given value.

### HasSourceCloud

`func (o *AddMigration200ResponseAnyOfMigration) HasSourceCloud() bool`

HasSourceCloud returns a boolean if a field has been set.

### GetTargetCloud

`func (o *AddMigration200ResponseAnyOfMigration) GetTargetCloud() AddMigration200ResponseAnyOfMigrationTargetCloud`

GetTargetCloud returns the TargetCloud field if non-nil, zero value otherwise.

### GetTargetCloudOk

`func (o *AddMigration200ResponseAnyOfMigration) GetTargetCloudOk() (*AddMigration200ResponseAnyOfMigrationTargetCloud, bool)`

GetTargetCloudOk returns a tuple with the TargetCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCloud

`func (o *AddMigration200ResponseAnyOfMigration) SetTargetCloud(v AddMigration200ResponseAnyOfMigrationTargetCloud)`

SetTargetCloud sets TargetCloud field to given value.

### HasTargetCloud

`func (o *AddMigration200ResponseAnyOfMigration) HasTargetCloud() bool`

HasTargetCloud returns a boolean if a field has been set.

### GetTargetGroup

`func (o *AddMigration200ResponseAnyOfMigration) GetTargetGroup() AddMigration200ResponseAnyOfMigrationTargetGroup`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *AddMigration200ResponseAnyOfMigration) GetTargetGroupOk() (*AddMigration200ResponseAnyOfMigrationTargetGroup, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *AddMigration200ResponseAnyOfMigration) SetTargetGroup(v AddMigration200ResponseAnyOfMigrationTargetGroup)`

SetTargetGroup sets TargetGroup field to given value.

### HasTargetGroup

`func (o *AddMigration200ResponseAnyOfMigration) HasTargetGroup() bool`

HasTargetGroup returns a boolean if a field has been set.

### GetTargetPool

`func (o *AddMigration200ResponseAnyOfMigration) GetTargetPool() AddMigration200ResponseAnyOfMigrationTargetPool`

GetTargetPool returns the TargetPool field if non-nil, zero value otherwise.

### GetTargetPoolOk

`func (o *AddMigration200ResponseAnyOfMigration) GetTargetPoolOk() (*AddMigration200ResponseAnyOfMigrationTargetPool, bool)`

GetTargetPoolOk returns a tuple with the TargetPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPool

`func (o *AddMigration200ResponseAnyOfMigration) SetTargetPool(v AddMigration200ResponseAnyOfMigrationTargetPool)`

SetTargetPool sets TargetPool field to given value.

### HasTargetPool

`func (o *AddMigration200ResponseAnyOfMigration) HasTargetPool() bool`

HasTargetPool returns a boolean if a field has been set.

### GetServers

`func (o *AddMigration200ResponseAnyOfMigration) GetServers() []AddMigration200ResponseAnyOfMigrationServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *AddMigration200ResponseAnyOfMigration) GetServersOk() (*[]AddMigration200ResponseAnyOfMigrationServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *AddMigration200ResponseAnyOfMigration) SetServers(v []AddMigration200ResponseAnyOfMigrationServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *AddMigration200ResponseAnyOfMigration) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


