# GetMigration200ResponseMigration

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
**LinuxKeyPair** | Pointer to [**GetMigration200ResponseMigrationLinuxKeyPair**](GetMigration200ResponseMigrationLinuxKeyPair.md) |  | [optional] 
**WindowsUsername** | Pointer to **NullableString** | Windows Username for migrated servers | [optional] 
**WindowsPassword** | Pointer to **NullableString** | Windows Password for migrated servers | [optional] 
**SourceCloud** | Pointer to [**GetMigration200ResponseMigrationSourceCloud**](GetMigration200ResponseMigrationSourceCloud.md) |  | [optional] 
**TargetCloud** | Pointer to [**GetMigration200ResponseMigrationTargetCloud**](GetMigration200ResponseMigrationTargetCloud.md) |  | [optional] 
**TargetGroup** | Pointer to [**GetMigration200ResponseMigrationTargetGroup**](GetMigration200ResponseMigrationTargetGroup.md) |  | [optional] 
**TargetPool** | Pointer to [**GetMigration200ResponseMigrationTargetPool**](GetMigration200ResponseMigrationTargetPool.md) |  | [optional] 
**Servers** | Pointer to [**[]GetMigration200ResponseMigrationServersInner**](GetMigration200ResponseMigrationServersInner.md) |  | [optional] 

## Methods

### NewGetMigration200ResponseMigration

`func NewGetMigration200ResponseMigration() *GetMigration200ResponseMigration`

NewGetMigration200ResponseMigration instantiates a new GetMigration200ResponseMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMigration200ResponseMigrationWithDefaults

`func NewGetMigration200ResponseMigrationWithDefaults() *GetMigration200ResponseMigration`

NewGetMigration200ResponseMigrationWithDefaults instantiates a new GetMigration200ResponseMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMigration200ResponseMigration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMigration200ResponseMigration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMigration200ResponseMigration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetMigration200ResponseMigration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetMigration200ResponseMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMigration200ResponseMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMigration200ResponseMigration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetMigration200ResponseMigration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *GetMigration200ResponseMigration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMigration200ResponseMigration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMigration200ResponseMigration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMigration200ResponseMigration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetMigration200ResponseMigration) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetMigration200ResponseMigration) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetMigration200ResponseMigration) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetMigration200ResponseMigration) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetMigration200ResponseMigration) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetMigration200ResponseMigration) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetSkippedPrechecks

`func (o *GetMigration200ResponseMigration) GetSkippedPrechecks() bool`

GetSkippedPrechecks returns the SkippedPrechecks field if non-nil, zero value otherwise.

### GetSkippedPrechecksOk

`func (o *GetMigration200ResponseMigration) GetSkippedPrechecksOk() (*bool, bool)`

GetSkippedPrechecksOk returns a tuple with the SkippedPrechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedPrechecks

`func (o *GetMigration200ResponseMigration) SetSkippedPrechecks(v bool)`

SetSkippedPrechecks sets SkippedPrechecks field to given value.

### HasSkippedPrechecks

`func (o *GetMigration200ResponseMigration) HasSkippedPrechecks() bool`

HasSkippedPrechecks returns a boolean if a field has been set.

### GetInstallGuestTools

`func (o *GetMigration200ResponseMigration) GetInstallGuestTools() bool`

GetInstallGuestTools returns the InstallGuestTools field if non-nil, zero value otherwise.

### GetInstallGuestToolsOk

`func (o *GetMigration200ResponseMigration) GetInstallGuestToolsOk() (*bool, bool)`

GetInstallGuestToolsOk returns a tuple with the InstallGuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallGuestTools

`func (o *GetMigration200ResponseMigration) SetInstallGuestTools(v bool)`

SetInstallGuestTools sets InstallGuestTools field to given value.

### HasInstallGuestTools

`func (o *GetMigration200ResponseMigration) HasInstallGuestTools() bool`

HasInstallGuestTools returns a boolean if a field has been set.

### GetReInitializeServerOnMigration

`func (o *GetMigration200ResponseMigration) GetReInitializeServerOnMigration() bool`

GetReInitializeServerOnMigration returns the ReInitializeServerOnMigration field if non-nil, zero value otherwise.

### GetReInitializeServerOnMigrationOk

`func (o *GetMigration200ResponseMigration) GetReInitializeServerOnMigrationOk() (*bool, bool)`

GetReInitializeServerOnMigrationOk returns a tuple with the ReInitializeServerOnMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReInitializeServerOnMigration

`func (o *GetMigration200ResponseMigration) SetReInitializeServerOnMigration(v bool)`

SetReInitializeServerOnMigration sets ReInitializeServerOnMigration field to given value.

### HasReInitializeServerOnMigration

`func (o *GetMigration200ResponseMigration) HasReInitializeServerOnMigration() bool`

HasReInitializeServerOnMigration returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *GetMigration200ResponseMigration) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *GetMigration200ResponseMigration) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *GetMigration200ResponseMigration) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *GetMigration200ResponseMigration) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *GetMigration200ResponseMigration) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *GetMigration200ResponseMigration) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *GetMigration200ResponseMigration) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *GetMigration200ResponseMigration) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *GetMigration200ResponseMigration) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *GetMigration200ResponseMigration) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *GetMigration200ResponseMigration) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *GetMigration200ResponseMigration) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPair

`func (o *GetMigration200ResponseMigration) GetLinuxKeyPair() GetMigration200ResponseMigrationLinuxKeyPair`

GetLinuxKeyPair returns the LinuxKeyPair field if non-nil, zero value otherwise.

### GetLinuxKeyPairOk

`func (o *GetMigration200ResponseMigration) GetLinuxKeyPairOk() (*GetMigration200ResponseMigrationLinuxKeyPair, bool)`

GetLinuxKeyPairOk returns a tuple with the LinuxKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPair

`func (o *GetMigration200ResponseMigration) SetLinuxKeyPair(v GetMigration200ResponseMigrationLinuxKeyPair)`

SetLinuxKeyPair sets LinuxKeyPair field to given value.

### HasLinuxKeyPair

`func (o *GetMigration200ResponseMigration) HasLinuxKeyPair() bool`

HasLinuxKeyPair returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *GetMigration200ResponseMigration) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *GetMigration200ResponseMigration) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *GetMigration200ResponseMigration) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *GetMigration200ResponseMigration) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *GetMigration200ResponseMigration) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *GetMigration200ResponseMigration) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *GetMigration200ResponseMigration) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *GetMigration200ResponseMigration) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *GetMigration200ResponseMigration) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *GetMigration200ResponseMigration) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *GetMigration200ResponseMigration) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *GetMigration200ResponseMigration) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetSourceCloud

`func (o *GetMigration200ResponseMigration) GetSourceCloud() GetMigration200ResponseMigrationSourceCloud`

GetSourceCloud returns the SourceCloud field if non-nil, zero value otherwise.

### GetSourceCloudOk

`func (o *GetMigration200ResponseMigration) GetSourceCloudOk() (*GetMigration200ResponseMigrationSourceCloud, bool)`

GetSourceCloudOk returns a tuple with the SourceCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCloud

`func (o *GetMigration200ResponseMigration) SetSourceCloud(v GetMigration200ResponseMigrationSourceCloud)`

SetSourceCloud sets SourceCloud field to given value.

### HasSourceCloud

`func (o *GetMigration200ResponseMigration) HasSourceCloud() bool`

HasSourceCloud returns a boolean if a field has been set.

### GetTargetCloud

`func (o *GetMigration200ResponseMigration) GetTargetCloud() GetMigration200ResponseMigrationTargetCloud`

GetTargetCloud returns the TargetCloud field if non-nil, zero value otherwise.

### GetTargetCloudOk

`func (o *GetMigration200ResponseMigration) GetTargetCloudOk() (*GetMigration200ResponseMigrationTargetCloud, bool)`

GetTargetCloudOk returns a tuple with the TargetCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCloud

`func (o *GetMigration200ResponseMigration) SetTargetCloud(v GetMigration200ResponseMigrationTargetCloud)`

SetTargetCloud sets TargetCloud field to given value.

### HasTargetCloud

`func (o *GetMigration200ResponseMigration) HasTargetCloud() bool`

HasTargetCloud returns a boolean if a field has been set.

### GetTargetGroup

`func (o *GetMigration200ResponseMigration) GetTargetGroup() GetMigration200ResponseMigrationTargetGroup`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *GetMigration200ResponseMigration) GetTargetGroupOk() (*GetMigration200ResponseMigrationTargetGroup, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *GetMigration200ResponseMigration) SetTargetGroup(v GetMigration200ResponseMigrationTargetGroup)`

SetTargetGroup sets TargetGroup field to given value.

### HasTargetGroup

`func (o *GetMigration200ResponseMigration) HasTargetGroup() bool`

HasTargetGroup returns a boolean if a field has been set.

### GetTargetPool

`func (o *GetMigration200ResponseMigration) GetTargetPool() GetMigration200ResponseMigrationTargetPool`

GetTargetPool returns the TargetPool field if non-nil, zero value otherwise.

### GetTargetPoolOk

`func (o *GetMigration200ResponseMigration) GetTargetPoolOk() (*GetMigration200ResponseMigrationTargetPool, bool)`

GetTargetPoolOk returns a tuple with the TargetPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPool

`func (o *GetMigration200ResponseMigration) SetTargetPool(v GetMigration200ResponseMigrationTargetPool)`

SetTargetPool sets TargetPool field to given value.

### HasTargetPool

`func (o *GetMigration200ResponseMigration) HasTargetPool() bool`

HasTargetPool returns a boolean if a field has been set.

### GetServers

`func (o *GetMigration200ResponseMigration) GetServers() []GetMigration200ResponseMigrationServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetMigration200ResponseMigration) GetServersOk() (*[]GetMigration200ResponseMigrationServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetMigration200ResponseMigration) SetServers(v []GetMigration200ResponseMigrationServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetMigration200ResponseMigration) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


