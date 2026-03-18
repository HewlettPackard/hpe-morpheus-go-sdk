# UpdateImageBuild200ResponseImageBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetImageBuild200ResponseImageBuildAccount**](GetImageBuild200ResponseImageBuildAccount.md) |  | [optional] 
**Type** | Pointer to [**GetImageBuild200ResponseImageBuildType**](GetImageBuild200ResponseImageBuildType.md) |  | [optional] 
**Site** | Pointer to [**GetImageBuild200ResponseImageBuildSite**](GetImageBuild200ResponseImageBuildSite.md) |  | [optional] 
**Zone** | Pointer to [**GetImageBuild200ResponseImageBuildZone**](GetImageBuild200ResponseImageBuildZone.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**BootScript** | Pointer to [**GetImageBuild200ResponseImageBuildBootScript**](GetImageBuild200ResponseImageBuildBootScript.md) |  | [optional] 
**BootCommand** | Pointer to **NullableString** |  | [optional] 
**PreseedScript** | Pointer to [**GetImageBuild200ResponseImageBuildPreseedScript**](GetImageBuild200ResponseImageBuildPreseedScript.md) |  | [optional] 
**Scripts** | Pointer to [**[]GetImageBuild200ResponseImageBuildScriptsInner**](GetImageBuild200ResponseImageBuildScriptsInner.md) |  | [optional] 
**SshUsername** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to **NullableString** |  | [optional] 
**BuildOutputName** | Pointer to **NullableString** |  | [optional] 
**ConversionFormats** | Pointer to **NullableString** |  | [optional] 
**IsCloudInit** | Pointer to **bool** |  | [optional] 
**VmToolsInstalled** | Pointer to **bool** |  | [optional] 
**KeepResults** | Pointer to **NullableInt64** |  | [optional] 
**Config** | Pointer to [**GetImageBuild200ResponseImageBuildConfig**](GetImageBuild200ResponseImageBuildConfig.md) |  | [optional] 
**LastResult** | Pointer to [**GetImageBuild200ResponseImageBuildLastResult**](GetImageBuild200ResponseImageBuildLastResult.md) |  | [optional] 
**ExecutionCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewUpdateImageBuild200ResponseImageBuild

`func NewUpdateImageBuild200ResponseImageBuild() *UpdateImageBuild200ResponseImageBuild`

NewUpdateImageBuild200ResponseImageBuild instantiates a new UpdateImageBuild200ResponseImageBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateImageBuild200ResponseImageBuildWithDefaults

`func NewUpdateImageBuild200ResponseImageBuildWithDefaults() *UpdateImageBuild200ResponseImageBuild`

NewUpdateImageBuild200ResponseImageBuildWithDefaults instantiates a new UpdateImageBuild200ResponseImageBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateImageBuild200ResponseImageBuild) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateImageBuild200ResponseImageBuild) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateImageBuild200ResponseImageBuild) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateImageBuild200ResponseImageBuild) GetAccount() GetImageBuild200ResponseImageBuildAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetAccountOk() (*GetImageBuild200ResponseImageBuildAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateImageBuild200ResponseImageBuild) SetAccount(v GetImageBuild200ResponseImageBuildAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateImageBuild200ResponseImageBuild) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetType

`func (o *UpdateImageBuild200ResponseImageBuild) GetType() GetImageBuild200ResponseImageBuildType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetTypeOk() (*GetImageBuild200ResponseImageBuildType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateImageBuild200ResponseImageBuild) SetType(v GetImageBuild200ResponseImageBuildType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateImageBuild200ResponseImageBuild) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSite

`func (o *UpdateImageBuild200ResponseImageBuild) GetSite() GetImageBuild200ResponseImageBuildSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetSiteOk() (*GetImageBuild200ResponseImageBuildSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *UpdateImageBuild200ResponseImageBuild) SetSite(v GetImageBuild200ResponseImageBuildSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *UpdateImageBuild200ResponseImageBuild) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *UpdateImageBuild200ResponseImageBuild) GetZone() GetImageBuild200ResponseImageBuildZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetZoneOk() (*GetImageBuild200ResponseImageBuildZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateImageBuild200ResponseImageBuild) SetZone(v GetImageBuild200ResponseImageBuildZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateImageBuild200ResponseImageBuild) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetName

`func (o *UpdateImageBuild200ResponseImageBuild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateImageBuild200ResponseImageBuild) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateImageBuild200ResponseImageBuild) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateImageBuild200ResponseImageBuild) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateImageBuild200ResponseImageBuild) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateImageBuild200ResponseImageBuild) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateImageBuild200ResponseImageBuild) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateImageBuild200ResponseImageBuild) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBootScript

`func (o *UpdateImageBuild200ResponseImageBuild) GetBootScript() GetImageBuild200ResponseImageBuildBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetBootScriptOk() (*GetImageBuild200ResponseImageBuildBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *UpdateImageBuild200ResponseImageBuild) SetBootScript(v GetImageBuild200ResponseImageBuildBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *UpdateImageBuild200ResponseImageBuild) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.

### GetBootCommand

`func (o *UpdateImageBuild200ResponseImageBuild) GetBootCommand() string`

GetBootCommand returns the BootCommand field if non-nil, zero value otherwise.

### GetBootCommandOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetBootCommandOk() (*string, bool)`

GetBootCommandOk returns a tuple with the BootCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootCommand

`func (o *UpdateImageBuild200ResponseImageBuild) SetBootCommand(v string)`

SetBootCommand sets BootCommand field to given value.

### HasBootCommand

`func (o *UpdateImageBuild200ResponseImageBuild) HasBootCommand() bool`

HasBootCommand returns a boolean if a field has been set.

### SetBootCommandNil

`func (o *UpdateImageBuild200ResponseImageBuild) SetBootCommandNil(b bool)`

 SetBootCommandNil sets the value for BootCommand to be an explicit nil

### UnsetBootCommand
`func (o *UpdateImageBuild200ResponseImageBuild) UnsetBootCommand()`

UnsetBootCommand ensures that no value is present for BootCommand, not even an explicit nil
### GetPreseedScript

`func (o *UpdateImageBuild200ResponseImageBuild) GetPreseedScript() GetImageBuild200ResponseImageBuildPreseedScript`

GetPreseedScript returns the PreseedScript field if non-nil, zero value otherwise.

### GetPreseedScriptOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetPreseedScriptOk() (*GetImageBuild200ResponseImageBuildPreseedScript, bool)`

GetPreseedScriptOk returns a tuple with the PreseedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScript

`func (o *UpdateImageBuild200ResponseImageBuild) SetPreseedScript(v GetImageBuild200ResponseImageBuildPreseedScript)`

SetPreseedScript sets PreseedScript field to given value.

### HasPreseedScript

`func (o *UpdateImageBuild200ResponseImageBuild) HasPreseedScript() bool`

HasPreseedScript returns a boolean if a field has been set.

### GetScripts

`func (o *UpdateImageBuild200ResponseImageBuild) GetScripts() []GetImageBuild200ResponseImageBuildScriptsInner`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetScriptsOk() (*[]GetImageBuild200ResponseImageBuildScriptsInner, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *UpdateImageBuild200ResponseImageBuild) SetScripts(v []GetImageBuild200ResponseImageBuildScriptsInner)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *UpdateImageBuild200ResponseImageBuild) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetSshUsername

`func (o *UpdateImageBuild200ResponseImageBuild) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *UpdateImageBuild200ResponseImageBuild) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *UpdateImageBuild200ResponseImageBuild) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *UpdateImageBuild200ResponseImageBuild) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *UpdateImageBuild200ResponseImageBuild) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *UpdateImageBuild200ResponseImageBuild) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetStorageProvider

`func (o *UpdateImageBuild200ResponseImageBuild) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *UpdateImageBuild200ResponseImageBuild) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *UpdateImageBuild200ResponseImageBuild) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### SetStorageProviderNil

`func (o *UpdateImageBuild200ResponseImageBuild) SetStorageProviderNil(b bool)`

 SetStorageProviderNil sets the value for StorageProvider to be an explicit nil

### UnsetStorageProvider
`func (o *UpdateImageBuild200ResponseImageBuild) UnsetStorageProvider()`

UnsetStorageProvider ensures that no value is present for StorageProvider, not even an explicit nil
### GetBuildOutputName

`func (o *UpdateImageBuild200ResponseImageBuild) GetBuildOutputName() string`

GetBuildOutputName returns the BuildOutputName field if non-nil, zero value otherwise.

### GetBuildOutputNameOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetBuildOutputNameOk() (*string, bool)`

GetBuildOutputNameOk returns a tuple with the BuildOutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputName

`func (o *UpdateImageBuild200ResponseImageBuild) SetBuildOutputName(v string)`

SetBuildOutputName sets BuildOutputName field to given value.

### HasBuildOutputName

`func (o *UpdateImageBuild200ResponseImageBuild) HasBuildOutputName() bool`

HasBuildOutputName returns a boolean if a field has been set.

### SetBuildOutputNameNil

`func (o *UpdateImageBuild200ResponseImageBuild) SetBuildOutputNameNil(b bool)`

 SetBuildOutputNameNil sets the value for BuildOutputName to be an explicit nil

### UnsetBuildOutputName
`func (o *UpdateImageBuild200ResponseImageBuild) UnsetBuildOutputName()`

UnsetBuildOutputName ensures that no value is present for BuildOutputName, not even an explicit nil
### GetConversionFormats

`func (o *UpdateImageBuild200ResponseImageBuild) GetConversionFormats() string`

GetConversionFormats returns the ConversionFormats field if non-nil, zero value otherwise.

### GetConversionFormatsOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetConversionFormatsOk() (*string, bool)`

GetConversionFormatsOk returns a tuple with the ConversionFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFormats

`func (o *UpdateImageBuild200ResponseImageBuild) SetConversionFormats(v string)`

SetConversionFormats sets ConversionFormats field to given value.

### HasConversionFormats

`func (o *UpdateImageBuild200ResponseImageBuild) HasConversionFormats() bool`

HasConversionFormats returns a boolean if a field has been set.

### SetConversionFormatsNil

`func (o *UpdateImageBuild200ResponseImageBuild) SetConversionFormatsNil(b bool)`

 SetConversionFormatsNil sets the value for ConversionFormats to be an explicit nil

### UnsetConversionFormats
`func (o *UpdateImageBuild200ResponseImageBuild) UnsetConversionFormats()`

UnsetConversionFormats ensures that no value is present for ConversionFormats, not even an explicit nil
### GetIsCloudInit

`func (o *UpdateImageBuild200ResponseImageBuild) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *UpdateImageBuild200ResponseImageBuild) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *UpdateImageBuild200ResponseImageBuild) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *UpdateImageBuild200ResponseImageBuild) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *UpdateImageBuild200ResponseImageBuild) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *UpdateImageBuild200ResponseImageBuild) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetKeepResults

`func (o *UpdateImageBuild200ResponseImageBuild) GetKeepResults() int64`

GetKeepResults returns the KeepResults field if non-nil, zero value otherwise.

### GetKeepResultsOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetKeepResultsOk() (*int64, bool)`

GetKeepResultsOk returns a tuple with the KeepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepResults

`func (o *UpdateImageBuild200ResponseImageBuild) SetKeepResults(v int64)`

SetKeepResults sets KeepResults field to given value.

### HasKeepResults

`func (o *UpdateImageBuild200ResponseImageBuild) HasKeepResults() bool`

HasKeepResults returns a boolean if a field has been set.

### SetKeepResultsNil

`func (o *UpdateImageBuild200ResponseImageBuild) SetKeepResultsNil(b bool)`

 SetKeepResultsNil sets the value for KeepResults to be an explicit nil

### UnsetKeepResults
`func (o *UpdateImageBuild200ResponseImageBuild) UnsetKeepResults()`

UnsetKeepResults ensures that no value is present for KeepResults, not even an explicit nil
### GetConfig

`func (o *UpdateImageBuild200ResponseImageBuild) GetConfig() GetImageBuild200ResponseImageBuildConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetConfigOk() (*GetImageBuild200ResponseImageBuildConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateImageBuild200ResponseImageBuild) SetConfig(v GetImageBuild200ResponseImageBuildConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateImageBuild200ResponseImageBuild) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLastResult

`func (o *UpdateImageBuild200ResponseImageBuild) GetLastResult() GetImageBuild200ResponseImageBuildLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetLastResultOk() (*GetImageBuild200ResponseImageBuildLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *UpdateImageBuild200ResponseImageBuild) SetLastResult(v GetImageBuild200ResponseImageBuildLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *UpdateImageBuild200ResponseImageBuild) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetExecutionCount

`func (o *UpdateImageBuild200ResponseImageBuild) GetExecutionCount() int64`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *UpdateImageBuild200ResponseImageBuild) GetExecutionCountOk() (*int64, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *UpdateImageBuild200ResponseImageBuild) SetExecutionCount(v int64)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *UpdateImageBuild200ResponseImageBuild) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


