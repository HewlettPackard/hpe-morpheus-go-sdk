# UpdateMigrationRequestMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**SkippedPrechecks** | Pointer to **bool** | Indicates if the precheck should be skipped | [optional] 
**InstallGuestTools** | Pointer to **bool** | Indicates if guest tools should be installed | [optional] 
**ReInitializeServerOnMigration** | Pointer to **bool** | Indicates if the server should be re-initialized on migration | [optional] 
**LinuxUsername** | Pointer to **NullableString** | Linux Username for migrated servers | [optional] 
**LinuxPassword** | Pointer to **NullableString** | Linux Password for migrated servers | [optional] 
**LinuxKeyPair** | Pointer to [**UpdateMigrationRequestMigrationLinuxKeyPair**](UpdateMigrationRequestMigrationLinuxKeyPair.md) |  | [optional] 
**WindowsUsername** | Pointer to **NullableString** | Windows Username for migrated servers | [optional] 
**WindowsPassword** | Pointer to **NullableString** | Windows Password for migrated servers | [optional] 
**SourceCloudId** | Pointer to **int64** | Source Cloud ID. The API &#x60;/api/migrations/source-clouds&#x60; can be used to find available options.  | [optional] 
**TargetCloudId** | Pointer to **int64** | Target Cloud ID. The API &#x60;/api/migrations/target-clouds?sourceCloudId&#x3D;34&#x60; can be used to find available options.  | [optional] 
**TargetGroupId** | Pointer to **int64** | Target Group ID.  The Options API &#x60;/api/options/targetGroups?sourceCloudId&#x3D;34&amp;targetCloudId&#x3D;129&#x60; can be used to find available options.  | [optional] 
**TargetPoolId** | Pointer to [**UpdateMigrationRequestMigrationTargetPoolId**](UpdateMigrationRequestMigrationTargetPoolId.md) |  | [optional] 
**SourceServerIds** | Pointer to **[]int64** | Array of Server IDs to be migrated. The API &#x60;/api/migrations/source-servers?sourceCloudId&#x3D;34&#x60; can be used to find available options.  | [optional] 
**Datastores** | Pointer to [**[]UpdateMigrationRequestMigrationDatastoresInner**](UpdateMigrationRequestMigrationDatastoresInner.md) | Array of datastore mappings. | [optional] 
**Networks** | Pointer to [**[]UpdateMigrationRequestMigrationNetworksInner**](UpdateMigrationRequestMigrationNetworksInner.md) | Array of network mappings. | [optional] 

## Methods

### NewUpdateMigrationRequestMigration

`func NewUpdateMigrationRequestMigration() *UpdateMigrationRequestMigration`

NewUpdateMigrationRequestMigration instantiates a new UpdateMigrationRequestMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateMigrationRequestMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMigrationRequestMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMigrationRequestMigration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMigrationRequestMigration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSkippedPrechecks

`func (o *UpdateMigrationRequestMigration) GetSkippedPrechecks() bool`

GetSkippedPrechecks returns the SkippedPrechecks field if non-nil, zero value otherwise.

### GetSkippedPrechecksOk

`func (o *UpdateMigrationRequestMigration) GetSkippedPrechecksOk() (*bool, bool)`

GetSkippedPrechecksOk returns a tuple with the SkippedPrechecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedPrechecks

`func (o *UpdateMigrationRequestMigration) SetSkippedPrechecks(v bool)`

SetSkippedPrechecks sets SkippedPrechecks field to given value.

### HasSkippedPrechecks

`func (o *UpdateMigrationRequestMigration) HasSkippedPrechecks() bool`

HasSkippedPrechecks returns a boolean if a field has been set.

### GetInstallGuestTools

`func (o *UpdateMigrationRequestMigration) GetInstallGuestTools() bool`

GetInstallGuestTools returns the InstallGuestTools field if non-nil, zero value otherwise.

### GetInstallGuestToolsOk

`func (o *UpdateMigrationRequestMigration) GetInstallGuestToolsOk() (*bool, bool)`

GetInstallGuestToolsOk returns a tuple with the InstallGuestTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallGuestTools

`func (o *UpdateMigrationRequestMigration) SetInstallGuestTools(v bool)`

SetInstallGuestTools sets InstallGuestTools field to given value.

### HasInstallGuestTools

`func (o *UpdateMigrationRequestMigration) HasInstallGuestTools() bool`

HasInstallGuestTools returns a boolean if a field has been set.

### GetReInitializeServerOnMigration

`func (o *UpdateMigrationRequestMigration) GetReInitializeServerOnMigration() bool`

GetReInitializeServerOnMigration returns the ReInitializeServerOnMigration field if non-nil, zero value otherwise.

### GetReInitializeServerOnMigrationOk

`func (o *UpdateMigrationRequestMigration) GetReInitializeServerOnMigrationOk() (*bool, bool)`

GetReInitializeServerOnMigrationOk returns a tuple with the ReInitializeServerOnMigration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReInitializeServerOnMigration

`func (o *UpdateMigrationRequestMigration) SetReInitializeServerOnMigration(v bool)`

SetReInitializeServerOnMigration sets ReInitializeServerOnMigration field to given value.

### HasReInitializeServerOnMigration

`func (o *UpdateMigrationRequestMigration) HasReInitializeServerOnMigration() bool`

HasReInitializeServerOnMigration returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *UpdateMigrationRequestMigration) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *UpdateMigrationRequestMigration) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *UpdateMigrationRequestMigration) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *UpdateMigrationRequestMigration) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *UpdateMigrationRequestMigration) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *UpdateMigrationRequestMigration) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *UpdateMigrationRequestMigration) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *UpdateMigrationRequestMigration) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *UpdateMigrationRequestMigration) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *UpdateMigrationRequestMigration) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *UpdateMigrationRequestMigration) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *UpdateMigrationRequestMigration) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPair

`func (o *UpdateMigrationRequestMigration) GetLinuxKeyPair() UpdateMigrationRequestMigrationLinuxKeyPair`

GetLinuxKeyPair returns the LinuxKeyPair field if non-nil, zero value otherwise.

### GetLinuxKeyPairOk

`func (o *UpdateMigrationRequestMigration) GetLinuxKeyPairOk() (*UpdateMigrationRequestMigrationLinuxKeyPair, bool)`

GetLinuxKeyPairOk returns a tuple with the LinuxKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPair

`func (o *UpdateMigrationRequestMigration) SetLinuxKeyPair(v UpdateMigrationRequestMigrationLinuxKeyPair)`

SetLinuxKeyPair sets LinuxKeyPair field to given value.

### HasLinuxKeyPair

`func (o *UpdateMigrationRequestMigration) HasLinuxKeyPair() bool`

HasLinuxKeyPair returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *UpdateMigrationRequestMigration) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *UpdateMigrationRequestMigration) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *UpdateMigrationRequestMigration) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *UpdateMigrationRequestMigration) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *UpdateMigrationRequestMigration) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *UpdateMigrationRequestMigration) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *UpdateMigrationRequestMigration) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *UpdateMigrationRequestMigration) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *UpdateMigrationRequestMigration) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *UpdateMigrationRequestMigration) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *UpdateMigrationRequestMigration) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *UpdateMigrationRequestMigration) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetSourceCloudId

`func (o *UpdateMigrationRequestMigration) GetSourceCloudId() int64`

GetSourceCloudId returns the SourceCloudId field if non-nil, zero value otherwise.

### GetSourceCloudIdOk

`func (o *UpdateMigrationRequestMigration) GetSourceCloudIdOk() (*int64, bool)`

GetSourceCloudIdOk returns a tuple with the SourceCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCloudId

`func (o *UpdateMigrationRequestMigration) SetSourceCloudId(v int64)`

SetSourceCloudId sets SourceCloudId field to given value.

### HasSourceCloudId

`func (o *UpdateMigrationRequestMigration) HasSourceCloudId() bool`

HasSourceCloudId returns a boolean if a field has been set.

### GetTargetCloudId

`func (o *UpdateMigrationRequestMigration) GetTargetCloudId() int64`

GetTargetCloudId returns the TargetCloudId field if non-nil, zero value otherwise.

### GetTargetCloudIdOk

`func (o *UpdateMigrationRequestMigration) GetTargetCloudIdOk() (*int64, bool)`

GetTargetCloudIdOk returns a tuple with the TargetCloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCloudId

`func (o *UpdateMigrationRequestMigration) SetTargetCloudId(v int64)`

SetTargetCloudId sets TargetCloudId field to given value.

### HasTargetCloudId

`func (o *UpdateMigrationRequestMigration) HasTargetCloudId() bool`

HasTargetCloudId returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *UpdateMigrationRequestMigration) GetTargetGroupId() int64`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *UpdateMigrationRequestMigration) GetTargetGroupIdOk() (*int64, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *UpdateMigrationRequestMigration) SetTargetGroupId(v int64)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *UpdateMigrationRequestMigration) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetTargetPoolId

`func (o *UpdateMigrationRequestMigration) GetTargetPoolId() UpdateMigrationRequestMigrationTargetPoolId`

GetTargetPoolId returns the TargetPoolId field if non-nil, zero value otherwise.

### GetTargetPoolIdOk

`func (o *UpdateMigrationRequestMigration) GetTargetPoolIdOk() (*UpdateMigrationRequestMigrationTargetPoolId, bool)`

GetTargetPoolIdOk returns a tuple with the TargetPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPoolId

`func (o *UpdateMigrationRequestMigration) SetTargetPoolId(v UpdateMigrationRequestMigrationTargetPoolId)`

SetTargetPoolId sets TargetPoolId field to given value.

### HasTargetPoolId

`func (o *UpdateMigrationRequestMigration) HasTargetPoolId() bool`

HasTargetPoolId returns a boolean if a field has been set.

### GetSourceServerIds

`func (o *UpdateMigrationRequestMigration) GetSourceServerIds() []int64`

GetSourceServerIds returns the SourceServerIds field if non-nil, zero value otherwise.

### GetSourceServerIdsOk

`func (o *UpdateMigrationRequestMigration) GetSourceServerIdsOk() (*[]int64, bool)`

GetSourceServerIdsOk returns a tuple with the SourceServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceServerIds

`func (o *UpdateMigrationRequestMigration) SetSourceServerIds(v []int64)`

SetSourceServerIds sets SourceServerIds field to given value.

### HasSourceServerIds

`func (o *UpdateMigrationRequestMigration) HasSourceServerIds() bool`

HasSourceServerIds returns a boolean if a field has been set.

### GetDatastores

`func (o *UpdateMigrationRequestMigration) GetDatastores() []UpdateMigrationRequestMigrationDatastoresInner`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *UpdateMigrationRequestMigration) GetDatastoresOk() (*[]UpdateMigrationRequestMigrationDatastoresInner, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *UpdateMigrationRequestMigration) SetDatastores(v []UpdateMigrationRequestMigrationDatastoresInner)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *UpdateMigrationRequestMigration) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetNetworks

`func (o *UpdateMigrationRequestMigration) GetNetworks() []UpdateMigrationRequestMigrationNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *UpdateMigrationRequestMigration) GetNetworksOk() (*[]UpdateMigrationRequestMigrationNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *UpdateMigrationRequestMigration) SetNetworks(v []UpdateMigrationRequestMigrationNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *UpdateMigrationRequestMigration) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


