# GetImageBuild200ResponseImageBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Site** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**BootScript** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript**](ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript.md) |  | [optional] 
**BootCommand** | Pointer to **NullableString** |  | [optional] 
**PreseedScript** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript**](ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript.md) |  | [optional] 
**Scripts** | Pointer to [**[]ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner**](ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner.md) |  | [optional] 
**SshUsername** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to **NullableString** |  | [optional] 
**BuildOutputName** | Pointer to **NullableString** |  | [optional] 
**ConversionFormats** | Pointer to **NullableString** |  | [optional] 
**IsCloudInit** | Pointer to **bool** |  | [optional] 
**VmToolsInstalled** | Pointer to **bool** |  | [optional] 
**KeepResults** | Pointer to **NullableInt64** |  | [optional] 
**Config** | Pointer to [**AddImageBuild200ResponseAllOfImageBuildConfig**](AddImageBuild200ResponseAllOfImageBuildConfig.md) |  | [optional] 
**LastResult** | Pointer to [**GetImageBuild200ResponseImageBuildLastResult**](GetImageBuild200ResponseImageBuildLastResult.md) |  | [optional] 
**ExecutionCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetImageBuild200ResponseImageBuild

`func NewGetImageBuild200ResponseImageBuild() *GetImageBuild200ResponseImageBuild`

NewGetImageBuild200ResponseImageBuild instantiates a new GetImageBuild200ResponseImageBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageBuild200ResponseImageBuildWithDefaults

`func NewGetImageBuild200ResponseImageBuildWithDefaults() *GetImageBuild200ResponseImageBuild`

NewGetImageBuild200ResponseImageBuildWithDefaults instantiates a new GetImageBuild200ResponseImageBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetImageBuild200ResponseImageBuild) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetImageBuild200ResponseImageBuild) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetImageBuild200ResponseImageBuild) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetImageBuild200ResponseImageBuild) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetImageBuild200ResponseImageBuild) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetImageBuild200ResponseImageBuild) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetImageBuild200ResponseImageBuild) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetImageBuild200ResponseImageBuild) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetType

`func (o *GetImageBuild200ResponseImageBuild) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetImageBuild200ResponseImageBuild) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetImageBuild200ResponseImageBuild) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *GetImageBuild200ResponseImageBuild) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSite

`func (o *GetImageBuild200ResponseImageBuild) GetSite() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GetImageBuild200ResponseImageBuild) GetSiteOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GetImageBuild200ResponseImageBuild) SetSite(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetSite sets Site field to given value.

### HasSite

`func (o *GetImageBuild200ResponseImageBuild) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *GetImageBuild200ResponseImageBuild) GetZone() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetImageBuild200ResponseImageBuild) GetZoneOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetImageBuild200ResponseImageBuild) SetZone(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetImageBuild200ResponseImageBuild) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetName

`func (o *GetImageBuild200ResponseImageBuild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetImageBuild200ResponseImageBuild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetImageBuild200ResponseImageBuild) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetImageBuild200ResponseImageBuild) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetImageBuild200ResponseImageBuild) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetImageBuild200ResponseImageBuild) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetImageBuild200ResponseImageBuild) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetImageBuild200ResponseImageBuild) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetImageBuild200ResponseImageBuild) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetImageBuild200ResponseImageBuild) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBootScript

`func (o *GetImageBuild200ResponseImageBuild) GetBootScript() ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *GetImageBuild200ResponseImageBuild) GetBootScriptOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *GetImageBuild200ResponseImageBuild) SetBootScript(v ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *GetImageBuild200ResponseImageBuild) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.

### GetBootCommand

`func (o *GetImageBuild200ResponseImageBuild) GetBootCommand() string`

GetBootCommand returns the BootCommand field if non-nil, zero value otherwise.

### GetBootCommandOk

`func (o *GetImageBuild200ResponseImageBuild) GetBootCommandOk() (*string, bool)`

GetBootCommandOk returns a tuple with the BootCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootCommand

`func (o *GetImageBuild200ResponseImageBuild) SetBootCommand(v string)`

SetBootCommand sets BootCommand field to given value.

### HasBootCommand

`func (o *GetImageBuild200ResponseImageBuild) HasBootCommand() bool`

HasBootCommand returns a boolean if a field has been set.

### SetBootCommandNil

`func (o *GetImageBuild200ResponseImageBuild) SetBootCommandNil(b bool)`

 SetBootCommandNil sets the value for BootCommand to be an explicit nil

### UnsetBootCommand
`func (o *GetImageBuild200ResponseImageBuild) UnsetBootCommand()`

UnsetBootCommand ensures that no value is present for BootCommand, not even an explicit nil
### GetPreseedScript

`func (o *GetImageBuild200ResponseImageBuild) GetPreseedScript() ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript`

GetPreseedScript returns the PreseedScript field if non-nil, zero value otherwise.

### GetPreseedScriptOk

`func (o *GetImageBuild200ResponseImageBuild) GetPreseedScriptOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript, bool)`

GetPreseedScriptOk returns a tuple with the PreseedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScript

`func (o *GetImageBuild200ResponseImageBuild) SetPreseedScript(v ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript)`

SetPreseedScript sets PreseedScript field to given value.

### HasPreseedScript

`func (o *GetImageBuild200ResponseImageBuild) HasPreseedScript() bool`

HasPreseedScript returns a boolean if a field has been set.

### GetScripts

`func (o *GetImageBuild200ResponseImageBuild) GetScripts() []ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *GetImageBuild200ResponseImageBuild) GetScriptsOk() (*[]ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *GetImageBuild200ResponseImageBuild) SetScripts(v []ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *GetImageBuild200ResponseImageBuild) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetSshUsername

`func (o *GetImageBuild200ResponseImageBuild) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *GetImageBuild200ResponseImageBuild) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *GetImageBuild200ResponseImageBuild) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *GetImageBuild200ResponseImageBuild) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *GetImageBuild200ResponseImageBuild) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *GetImageBuild200ResponseImageBuild) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *GetImageBuild200ResponseImageBuild) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *GetImageBuild200ResponseImageBuild) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetStorageProvider

`func (o *GetImageBuild200ResponseImageBuild) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *GetImageBuild200ResponseImageBuild) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *GetImageBuild200ResponseImageBuild) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *GetImageBuild200ResponseImageBuild) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### SetStorageProviderNil

`func (o *GetImageBuild200ResponseImageBuild) SetStorageProviderNil(b bool)`

 SetStorageProviderNil sets the value for StorageProvider to be an explicit nil

### UnsetStorageProvider
`func (o *GetImageBuild200ResponseImageBuild) UnsetStorageProvider()`

UnsetStorageProvider ensures that no value is present for StorageProvider, not even an explicit nil
### GetBuildOutputName

`func (o *GetImageBuild200ResponseImageBuild) GetBuildOutputName() string`

GetBuildOutputName returns the BuildOutputName field if non-nil, zero value otherwise.

### GetBuildOutputNameOk

`func (o *GetImageBuild200ResponseImageBuild) GetBuildOutputNameOk() (*string, bool)`

GetBuildOutputNameOk returns a tuple with the BuildOutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputName

`func (o *GetImageBuild200ResponseImageBuild) SetBuildOutputName(v string)`

SetBuildOutputName sets BuildOutputName field to given value.

### HasBuildOutputName

`func (o *GetImageBuild200ResponseImageBuild) HasBuildOutputName() bool`

HasBuildOutputName returns a boolean if a field has been set.

### SetBuildOutputNameNil

`func (o *GetImageBuild200ResponseImageBuild) SetBuildOutputNameNil(b bool)`

 SetBuildOutputNameNil sets the value for BuildOutputName to be an explicit nil

### UnsetBuildOutputName
`func (o *GetImageBuild200ResponseImageBuild) UnsetBuildOutputName()`

UnsetBuildOutputName ensures that no value is present for BuildOutputName, not even an explicit nil
### GetConversionFormats

`func (o *GetImageBuild200ResponseImageBuild) GetConversionFormats() string`

GetConversionFormats returns the ConversionFormats field if non-nil, zero value otherwise.

### GetConversionFormatsOk

`func (o *GetImageBuild200ResponseImageBuild) GetConversionFormatsOk() (*string, bool)`

GetConversionFormatsOk returns a tuple with the ConversionFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFormats

`func (o *GetImageBuild200ResponseImageBuild) SetConversionFormats(v string)`

SetConversionFormats sets ConversionFormats field to given value.

### HasConversionFormats

`func (o *GetImageBuild200ResponseImageBuild) HasConversionFormats() bool`

HasConversionFormats returns a boolean if a field has been set.

### SetConversionFormatsNil

`func (o *GetImageBuild200ResponseImageBuild) SetConversionFormatsNil(b bool)`

 SetConversionFormatsNil sets the value for ConversionFormats to be an explicit nil

### UnsetConversionFormats
`func (o *GetImageBuild200ResponseImageBuild) UnsetConversionFormats()`

UnsetConversionFormats ensures that no value is present for ConversionFormats, not even an explicit nil
### GetIsCloudInit

`func (o *GetImageBuild200ResponseImageBuild) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *GetImageBuild200ResponseImageBuild) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *GetImageBuild200ResponseImageBuild) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *GetImageBuild200ResponseImageBuild) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *GetImageBuild200ResponseImageBuild) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *GetImageBuild200ResponseImageBuild) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *GetImageBuild200ResponseImageBuild) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *GetImageBuild200ResponseImageBuild) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetKeepResults

`func (o *GetImageBuild200ResponseImageBuild) GetKeepResults() int64`

GetKeepResults returns the KeepResults field if non-nil, zero value otherwise.

### GetKeepResultsOk

`func (o *GetImageBuild200ResponseImageBuild) GetKeepResultsOk() (*int64, bool)`

GetKeepResultsOk returns a tuple with the KeepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepResults

`func (o *GetImageBuild200ResponseImageBuild) SetKeepResults(v int64)`

SetKeepResults sets KeepResults field to given value.

### HasKeepResults

`func (o *GetImageBuild200ResponseImageBuild) HasKeepResults() bool`

HasKeepResults returns a boolean if a field has been set.

### SetKeepResultsNil

`func (o *GetImageBuild200ResponseImageBuild) SetKeepResultsNil(b bool)`

 SetKeepResultsNil sets the value for KeepResults to be an explicit nil

### UnsetKeepResults
`func (o *GetImageBuild200ResponseImageBuild) UnsetKeepResults()`

UnsetKeepResults ensures that no value is present for KeepResults, not even an explicit nil
### GetConfig

`func (o *GetImageBuild200ResponseImageBuild) GetConfig() AddImageBuild200ResponseAllOfImageBuildConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetImageBuild200ResponseImageBuild) GetConfigOk() (*AddImageBuild200ResponseAllOfImageBuildConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetImageBuild200ResponseImageBuild) SetConfig(v AddImageBuild200ResponseAllOfImageBuildConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetImageBuild200ResponseImageBuild) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLastResult

`func (o *GetImageBuild200ResponseImageBuild) GetLastResult() GetImageBuild200ResponseImageBuildLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *GetImageBuild200ResponseImageBuild) GetLastResultOk() (*GetImageBuild200ResponseImageBuildLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *GetImageBuild200ResponseImageBuild) SetLastResult(v GetImageBuild200ResponseImageBuildLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *GetImageBuild200ResponseImageBuild) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetExecutionCount

`func (o *GetImageBuild200ResponseImageBuild) GetExecutionCount() int64`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *GetImageBuild200ResponseImageBuild) GetExecutionCountOk() (*int64, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *GetImageBuild200ResponseImageBuild) SetExecutionCount(v int64)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *GetImageBuild200ResponseImageBuild) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


