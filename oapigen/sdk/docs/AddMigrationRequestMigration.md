# AddMigrationRequestMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**SkippedPrechecks** | Pointer to **bool** | Indicates if the precheck should be skipped | [optional] [default to false]
**InstallGuestTools** | Pointer to **bool** | Indicates if guest tools should be installed | [optional] [default to true]
**ReInitializeServerOnMigration** | Pointer to **bool** | Indicates if the server should be re-initialized on migration | [optional] [default to false]
**LinuxUsername** | Pointer to **NullableString** | Linux Username for migrated servers | [optional] 
**LinuxPassword** | Pointer to **NullableString** | Linux Password for migrated servers | [optional] 
**LinuxKeyPair** | Pointer to [**AddMigrationRequestMigrationLinuxKeyPair**](AddMigrationRequestMigrationLinuxKeyPair.md) |  | [optional] 
**WindowsUsername** | Pointer to **NullableString** | Windows Username for migrated servers | [optional] 
**WindowsPassword** | Pointer to **NullableString** | Windows Password for migrated servers | [optional] 
**SourceCloudId** | Pointer to **int64** | Source Cloud ID. The API &#x60;/api/migrations/source-clouds&#x60; can be used to find available options.  | [optional] 
**TargetCloudId** | Pointer to **int64** | Target Cloud ID. The API &#x60;/api/migrations/target-clouds?sourceCloudId&#x3D;34&#x60; can be used to find available options.  | [optional] 
**TargetGroupId** | Pointer to **int64** | Target Group ID.  The Options API &#x60;/api/options/targetGroups?sourceCloudId&#x3D;34&amp;targetCloudId&#x3D;129&#x60; can be used to find available options.  | [optional] 
**TargetPoolId** | Pointer to [**AddMigrationRequestMigrationTargetPoolId**](AddMigrationRequestMigrationTargetPoolId.md) |  | [optional] 
**SourceServerIds** | Pointer to **[]int64** | Array of Server IDs to be migrated. The API &#x60;/api/migrations/source-servers?sourceCloudId&#x3D;34&#x60; can be used to find available options.  | [optional] 
**Datastores** | Pointer to [**[]AddMigrationRequestMigrationDatastoresInner**](AddMigrationRequestMigrationDatastoresInner.md) | Array of datastore mappings. | [optional] 
**Networks** | Pointer to [**[]AddMigrationRequestMigrationNetworksInner**](AddMigrationRequestMigrationNetworksInner.md) | Array of network mappings. | [optional] 

## Methods

### NewAddMigrationRequestMigration

`func NewAddMigrationRequestMigration() *AddMigrationRequestMigration`

NewAddMigrationRequestMigration instantiates a new AddMigrationRequestMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *AddMigrationRequestMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddMigrationRequestMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddMigrationRequestMigration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddMigrationRequestMigration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSkippedPrechecks

`func (o *AddMigrationRequestMigration) GetSkippedPrechecks() bool`

GetSkippedPrechecks returns the SkippedPrechecks field if non-nil, zero value otherwise.

### GetSkippedPrechecksOk

`func (o *AddMigrationRequestMigration) GetSkippedPrechecksOk() (*bool, bool)`

GetSkippedPrechecksOk returns a tuple with the SkippedPrechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedPrechecks

`func (o *AddMigrationRequestMigration) SetSkippedPrechecks(v bool)`

SetSkippedPrechecks sets SkippedPrechecks field to given value.

### HasSkippedPrechecks

`func (o *AddMigrationRequestMigration) HasSkippedPrechecks() bool`

HasSkippedPrechecks returns a boolean if a field has been set.

### GetInstallGuestTools

`func (o *AddMigrationRequestMigration) GetInstallGuestTools() bool`

GetInstallGuestTools returns the InstallGuestTools field if non-nil, zero value otherwise.

### GetInstallGuestToolsOk

`func (o *AddMigrationRequestMigration) GetInstallGuestToolsOk() (*bool, bool)`

GetInstallGuestToolsOk returns a tuple with the InstallGuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallGuestTools

`func (o *AddMigrationRequestMigration) SetInstallGuestTools(v bool)`

SetInstallGuestTools sets InstallGuestTools field to given value.

### HasInstallGuestTools

`func (o *AddMigrationRequestMigration) HasInstallGuestTools() bool`

HasInstallGuestTools returns a boolean if a field has been set.

### GetReInitializeServerOnMigration

`func (o *AddMigrationRequestMigration) GetReInitializeServerOnMigration() bool`

GetReInitializeServerOnMigration returns the ReInitializeServerOnMigration field if non-nil, zero value otherwise.

### GetReInitializeServerOnMigrationOk

`func (o *AddMigrationRequestMigration) GetReInitializeServerOnMigrationOk() (*bool, bool)`

GetReInitializeServerOnMigrationOk returns a tuple with the ReInitializeServerOnMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReInitializeServerOnMigration

`func (o *AddMigrationRequestMigration) SetReInitializeServerOnMigration(v bool)`

SetReInitializeServerOnMigration sets ReInitializeServerOnMigration field to given value.

### HasReInitializeServerOnMigration

`func (o *AddMigrationRequestMigration) HasReInitializeServerOnMigration() bool`

HasReInitializeServerOnMigration returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *AddMigrationRequestMigration) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *AddMigrationRequestMigration) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *AddMigrationRequestMigration) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *AddMigrationRequestMigration) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *AddMigrationRequestMigration) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *AddMigrationRequestMigration) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *AddMigrationRequestMigration) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *AddMigrationRequestMigration) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *AddMigrationRequestMigration) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *AddMigrationRequestMigration) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *AddMigrationRequestMigration) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *AddMigrationRequestMigration) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPair

`func (o *AddMigrationRequestMigration) GetLinuxKeyPair() AddMigrationRequestMigrationLinuxKeyPair`

GetLinuxKeyPair returns the LinuxKeyPair field if non-nil, zero value otherwise.

### GetLinuxKeyPairOk

`func (o *AddMigrationRequestMigration) GetLinuxKeyPairOk() (*AddMigrationRequestMigrationLinuxKeyPair, bool)`

GetLinuxKeyPairOk returns a tuple with the LinuxKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPair

`func (o *AddMigrationRequestMigration) SetLinuxKeyPair(v AddMigrationRequestMigrationLinuxKeyPair)`

SetLinuxKeyPair sets LinuxKeyPair field to given value.

### HasLinuxKeyPair

`func (o *AddMigrationRequestMigration) HasLinuxKeyPair() bool`

HasLinuxKeyPair returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *AddMigrationRequestMigration) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *AddMigrationRequestMigration) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *AddMigrationRequestMigration) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *AddMigrationRequestMigration) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *AddMigrationRequestMigration) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *AddMigrationRequestMigration) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *AddMigrationRequestMigration) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *AddMigrationRequestMigration) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *AddMigrationRequestMigration) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *AddMigrationRequestMigration) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *AddMigrationRequestMigration) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *AddMigrationRequestMigration) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetSourceCloudId

`func (o *AddMigrationRequestMigration) GetSourceCloudId() int64`

GetSourceCloudId returns the SourceCloudId field if non-nil, zero value otherwise.

### GetSourceCloudIdOk

`func (o *AddMigrationRequestMigration) GetSourceCloudIdOk() (*int64, bool)`

GetSourceCloudIdOk returns a tuple with the SourceCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCloudId

`func (o *AddMigrationRequestMigration) SetSourceCloudId(v int64)`

SetSourceCloudId sets SourceCloudId field to given value.

### HasSourceCloudId

`func (o *AddMigrationRequestMigration) HasSourceCloudId() bool`

HasSourceCloudId returns a boolean if a field has been set.

### GetTargetCloudId

`func (o *AddMigrationRequestMigration) GetTargetCloudId() int64`

GetTargetCloudId returns the TargetCloudId field if non-nil, zero value otherwise.

### GetTargetCloudIdOk

`func (o *AddMigrationRequestMigration) GetTargetCloudIdOk() (*int64, bool)`

GetTargetCloudIdOk returns a tuple with the TargetCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCloudId

`func (o *AddMigrationRequestMigration) SetTargetCloudId(v int64)`

SetTargetCloudId sets TargetCloudId field to given value.

### HasTargetCloudId

`func (o *AddMigrationRequestMigration) HasTargetCloudId() bool`

HasTargetCloudId returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *AddMigrationRequestMigration) GetTargetGroupId() int64`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *AddMigrationRequestMigration) GetTargetGroupIdOk() (*int64, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *AddMigrationRequestMigration) SetTargetGroupId(v int64)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *AddMigrationRequestMigration) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetTargetPoolId

`func (o *AddMigrationRequestMigration) GetTargetPoolId() AddMigrationRequestMigrationTargetPoolId`

GetTargetPoolId returns the TargetPoolId field if non-nil, zero value otherwise.

### GetTargetPoolIdOk

`func (o *AddMigrationRequestMigration) GetTargetPoolIdOk() (*AddMigrationRequestMigrationTargetPoolId, bool)`

GetTargetPoolIdOk returns a tuple with the TargetPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPoolId

`func (o *AddMigrationRequestMigration) SetTargetPoolId(v AddMigrationRequestMigrationTargetPoolId)`

SetTargetPoolId sets TargetPoolId field to given value.

### HasTargetPoolId

`func (o *AddMigrationRequestMigration) HasTargetPoolId() bool`

HasTargetPoolId returns a boolean if a field has been set.

### GetSourceServerIds

`func (o *AddMigrationRequestMigration) GetSourceServerIds() []int64`

GetSourceServerIds returns the SourceServerIds field if non-nil, zero value otherwise.

### GetSourceServerIdsOk

`func (o *AddMigrationRequestMigration) GetSourceServerIdsOk() (*[]int64, bool)`

GetSourceServerIdsOk returns a tuple with the SourceServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceServerIds

`func (o *AddMigrationRequestMigration) SetSourceServerIds(v []int64)`

SetSourceServerIds sets SourceServerIds field to given value.

### HasSourceServerIds

`func (o *AddMigrationRequestMigration) HasSourceServerIds() bool`

HasSourceServerIds returns a boolean if a field has been set.

### GetDatastores

`func (o *AddMigrationRequestMigration) GetDatastores() []AddMigrationRequestMigrationDatastoresInner`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *AddMigrationRequestMigration) GetDatastoresOk() (*[]AddMigrationRequestMigrationDatastoresInner, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *AddMigrationRequestMigration) SetDatastores(v []AddMigrationRequestMigrationDatastoresInner)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *AddMigrationRequestMigration) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetNetworks

`func (o *AddMigrationRequestMigration) GetNetworks() []AddMigrationRequestMigrationNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *AddMigrationRequestMigration) GetNetworksOk() (*[]AddMigrationRequestMigrationNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *AddMigrationRequestMigration) SetNetworks(v []AddMigrationRequestMigrationNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *AddMigrationRequestMigration) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


