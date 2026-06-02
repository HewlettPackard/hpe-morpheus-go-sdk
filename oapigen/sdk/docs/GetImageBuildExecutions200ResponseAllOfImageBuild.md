# GetImageBuildExecutions200ResponseAllOfImageBuild

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

### NewGetImageBuildExecutions200ResponseAllOfImageBuild

`func NewGetImageBuildExecutions200ResponseAllOfImageBuild() *GetImageBuildExecutions200ResponseAllOfImageBuild`

NewGetImageBuildExecutions200ResponseAllOfImageBuild instantiates a new GetImageBuildExecutions200ResponseAllOfImageBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetAccount() GetImageBuild200ResponseImageBuildAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetAccountOk() (*GetImageBuild200ResponseImageBuildAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetAccount(v GetImageBuild200ResponseImageBuildAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetType

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetType() GetImageBuild200ResponseImageBuildType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetTypeOk() (*GetImageBuild200ResponseImageBuildType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetType(v GetImageBuild200ResponseImageBuildType)`

SetType sets Type field to given value.

### HasType

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSite

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetSite() GetImageBuild200ResponseImageBuildSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetSiteOk() (*GetImageBuild200ResponseImageBuildSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetSite(v GetImageBuild200ResponseImageBuildSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetZone() GetImageBuild200ResponseImageBuildZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetZoneOk() (*GetImageBuild200ResponseImageBuildZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetZone(v GetImageBuild200ResponseImageBuildZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetName

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBootScript

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetBootScript() GetImageBuild200ResponseImageBuildBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetBootScriptOk() (*GetImageBuild200ResponseImageBuildBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetBootScript(v GetImageBuild200ResponseImageBuildBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.

### GetBootCommand

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetBootCommand() string`

GetBootCommand returns the BootCommand field if non-nil, zero value otherwise.

### GetBootCommandOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetBootCommandOk() (*string, bool)`

GetBootCommandOk returns a tuple with the BootCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootCommand

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetBootCommand(v string)`

SetBootCommand sets BootCommand field to given value.

### HasBootCommand

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasBootCommand() bool`

HasBootCommand returns a boolean if a field has been set.

### SetBootCommandNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetBootCommandNil(b bool)`

 SetBootCommandNil sets the value for BootCommand to be an explicit nil

### UnsetBootCommand
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) UnsetBootCommand()`

UnsetBootCommand ensures that no value is present for BootCommand, not even an explicit nil
### GetPreseedScript

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetPreseedScript() GetImageBuild200ResponseImageBuildPreseedScript`

GetPreseedScript returns the PreseedScript field if non-nil, zero value otherwise.

### GetPreseedScriptOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetPreseedScriptOk() (*GetImageBuild200ResponseImageBuildPreseedScript, bool)`

GetPreseedScriptOk returns a tuple with the PreseedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScript

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetPreseedScript(v GetImageBuild200ResponseImageBuildPreseedScript)`

SetPreseedScript sets PreseedScript field to given value.

### HasPreseedScript

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasPreseedScript() bool`

HasPreseedScript returns a boolean if a field has been set.

### GetScripts

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetScripts() []GetImageBuild200ResponseImageBuildScriptsInner`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetScriptsOk() (*[]GetImageBuild200ResponseImageBuildScriptsInner, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetScripts(v []GetImageBuild200ResponseImageBuildScriptsInner)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetSshUsername

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetStorageProvider

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### SetStorageProviderNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetStorageProviderNil(b bool)`

 SetStorageProviderNil sets the value for StorageProvider to be an explicit nil

### UnsetStorageProvider
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) UnsetStorageProvider()`

UnsetStorageProvider ensures that no value is present for StorageProvider, not even an explicit nil
### GetBuildOutputName

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetBuildOutputName() string`

GetBuildOutputName returns the BuildOutputName field if non-nil, zero value otherwise.

### GetBuildOutputNameOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetBuildOutputNameOk() (*string, bool)`

GetBuildOutputNameOk returns a tuple with the BuildOutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputName

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetBuildOutputName(v string)`

SetBuildOutputName sets BuildOutputName field to given value.

### HasBuildOutputName

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasBuildOutputName() bool`

HasBuildOutputName returns a boolean if a field has been set.

### SetBuildOutputNameNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetBuildOutputNameNil(b bool)`

 SetBuildOutputNameNil sets the value for BuildOutputName to be an explicit nil

### UnsetBuildOutputName
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) UnsetBuildOutputName()`

UnsetBuildOutputName ensures that no value is present for BuildOutputName, not even an explicit nil
### GetConversionFormats

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetConversionFormats() string`

GetConversionFormats returns the ConversionFormats field if non-nil, zero value otherwise.

### GetConversionFormatsOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetConversionFormatsOk() (*string, bool)`

GetConversionFormatsOk returns a tuple with the ConversionFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFormats

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetConversionFormats(v string)`

SetConversionFormats sets ConversionFormats field to given value.

### HasConversionFormats

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasConversionFormats() bool`

HasConversionFormats returns a boolean if a field has been set.

### SetConversionFormatsNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetConversionFormatsNil(b bool)`

 SetConversionFormatsNil sets the value for ConversionFormats to be an explicit nil

### UnsetConversionFormats
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) UnsetConversionFormats()`

UnsetConversionFormats ensures that no value is present for ConversionFormats, not even an explicit nil
### GetIsCloudInit

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetKeepResults

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetKeepResults() int64`

GetKeepResults returns the KeepResults field if non-nil, zero value otherwise.

### GetKeepResultsOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetKeepResultsOk() (*int64, bool)`

GetKeepResultsOk returns a tuple with the KeepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepResults

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetKeepResults(v int64)`

SetKeepResults sets KeepResults field to given value.

### HasKeepResults

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasKeepResults() bool`

HasKeepResults returns a boolean if a field has been set.

### SetKeepResultsNil

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetKeepResultsNil(b bool)`

 SetKeepResultsNil sets the value for KeepResults to be an explicit nil

### UnsetKeepResults
`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) UnsetKeepResults()`

UnsetKeepResults ensures that no value is present for KeepResults, not even an explicit nil
### GetConfig

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetConfig() GetImageBuild200ResponseImageBuildConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetConfigOk() (*GetImageBuild200ResponseImageBuildConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetConfig(v GetImageBuild200ResponseImageBuildConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLastResult

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetLastResult() GetImageBuild200ResponseImageBuildLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetLastResultOk() (*GetImageBuild200ResponseImageBuildLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetLastResult(v GetImageBuild200ResponseImageBuildLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetExecutionCount

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetExecutionCount() int64`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) GetExecutionCountOk() (*int64, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) SetExecutionCount(v int64)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *GetImageBuildExecutions200ResponseAllOfImageBuild) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


