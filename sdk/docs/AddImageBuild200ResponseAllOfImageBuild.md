# AddImageBuild200ResponseAllOfImageBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Site** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Zone** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BootScript** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript**](ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript.md) |  | [optional] 
**BootCommand** | Pointer to **string** |  | [optional] 
**PreseedScript** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript**](ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript.md) |  | [optional] 
**Scripts** | Pointer to [**[]ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner**](ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner.md) |  | [optional] 
**SshUsername** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to **string** |  | [optional] 
**BuildOutputName** | Pointer to **string** |  | [optional] 
**ConversionFormats** | Pointer to **string** |  | [optional] 
**IsCloudInit** | Pointer to **bool** |  | [optional] 
**VmToolsInstalled** | Pointer to **bool** |  | [optional] 
**KeepResults** | Pointer to **int64** |  | [optional] 
**Config** | Pointer to [**AddImageBuild200ResponseAllOfImageBuildConfig**](AddImageBuild200ResponseAllOfImageBuildConfig.md) |  | [optional] 
**LastResult** | Pointer to [**AddImageBuild200ResponseAllOfImageBuildLastResult**](AddImageBuild200ResponseAllOfImageBuildLastResult.md) |  | [optional] 
**ExecutionCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewAddImageBuild200ResponseAllOfImageBuild

`func NewAddImageBuild200ResponseAllOfImageBuild() *AddImageBuild200ResponseAllOfImageBuild`

NewAddImageBuild200ResponseAllOfImageBuild instantiates a new AddImageBuild200ResponseAllOfImageBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddImageBuild200ResponseAllOfImageBuildWithDefaults

`func NewAddImageBuild200ResponseAllOfImageBuildWithDefaults() *AddImageBuild200ResponseAllOfImageBuild`

NewAddImageBuild200ResponseAllOfImageBuildWithDefaults instantiates a new AddImageBuild200ResponseAllOfImageBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetType

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSite

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetSite() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetSiteOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetSite(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetSite sets Site field to given value.

### HasSite

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetName

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBootScript

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetBootScript() ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetBootScriptOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetBootScript(v ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.

### GetBootCommand

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetBootCommand() string`

GetBootCommand returns the BootCommand field if non-nil, zero value otherwise.

### GetBootCommandOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetBootCommandOk() (*string, bool)`

GetBootCommandOk returns a tuple with the BootCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootCommand

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetBootCommand(v string)`

SetBootCommand sets BootCommand field to given value.

### HasBootCommand

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasBootCommand() bool`

HasBootCommand returns a boolean if a field has been set.

### GetPreseedScript

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetPreseedScript() ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript`

GetPreseedScript returns the PreseedScript field if non-nil, zero value otherwise.

### GetPreseedScriptOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetPreseedScriptOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript, bool)`

GetPreseedScriptOk returns a tuple with the PreseedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScript

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetPreseedScript(v ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript)`

SetPreseedScript sets PreseedScript field to given value.

### HasPreseedScript

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasPreseedScript() bool`

HasPreseedScript returns a boolean if a field has been set.

### GetScripts

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetScripts() []ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetScriptsOk() (*[]ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetScripts(v []ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetSshUsername

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetStorageProvider

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetBuildOutputName

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetBuildOutputName() string`

GetBuildOutputName returns the BuildOutputName field if non-nil, zero value otherwise.

### GetBuildOutputNameOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetBuildOutputNameOk() (*string, bool)`

GetBuildOutputNameOk returns a tuple with the BuildOutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputName

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetBuildOutputName(v string)`

SetBuildOutputName sets BuildOutputName field to given value.

### HasBuildOutputName

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasBuildOutputName() bool`

HasBuildOutputName returns a boolean if a field has been set.

### GetConversionFormats

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetConversionFormats() string`

GetConversionFormats returns the ConversionFormats field if non-nil, zero value otherwise.

### GetConversionFormatsOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetConversionFormatsOk() (*string, bool)`

GetConversionFormatsOk returns a tuple with the ConversionFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFormats

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetConversionFormats(v string)`

SetConversionFormats sets ConversionFormats field to given value.

### HasConversionFormats

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasConversionFormats() bool`

HasConversionFormats returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetKeepResults

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetKeepResults() int64`

GetKeepResults returns the KeepResults field if non-nil, zero value otherwise.

### GetKeepResultsOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetKeepResultsOk() (*int64, bool)`

GetKeepResultsOk returns a tuple with the KeepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepResults

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetKeepResults(v int64)`

SetKeepResults sets KeepResults field to given value.

### HasKeepResults

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasKeepResults() bool`

HasKeepResults returns a boolean if a field has been set.

### GetConfig

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetConfig() AddImageBuild200ResponseAllOfImageBuildConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetConfigOk() (*AddImageBuild200ResponseAllOfImageBuildConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetConfig(v AddImageBuild200ResponseAllOfImageBuildConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLastResult

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetLastResult() AddImageBuild200ResponseAllOfImageBuildLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetLastResultOk() (*AddImageBuild200ResponseAllOfImageBuildLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetLastResult(v AddImageBuild200ResponseAllOfImageBuildLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetExecutionCount

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetExecutionCount() int64`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *AddImageBuild200ResponseAllOfImageBuild) GetExecutionCountOk() (*int64, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *AddImageBuild200ResponseAllOfImageBuild) SetExecutionCount(v int64)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *AddImageBuild200ResponseAllOfImageBuild) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


