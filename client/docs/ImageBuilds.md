# ImageBuilds

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
**Config** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerConfig**](ListImageBuilds200ResponseAllOfImageBuildsInnerConfig.md) |  | [optional] 
**LastResult** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult**](ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult.md) |  | [optional] 
**ExecutionCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewImageBuilds

`func NewImageBuilds() *ImageBuilds`

NewImageBuilds instantiates a new ImageBuilds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuildsWithDefaults

`func NewImageBuildsWithDefaults() *ImageBuilds`

NewImageBuildsWithDefaults instantiates a new ImageBuilds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageBuilds) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageBuilds) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageBuilds) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ImageBuilds) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ImageBuilds) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ImageBuilds) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ImageBuilds) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ImageBuilds) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetType

`func (o *ImageBuilds) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageBuilds) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageBuilds) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ImageBuilds) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSite

`func (o *ImageBuilds) GetSite() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ImageBuilds) GetSiteOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ImageBuilds) SetSite(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetSite sets Site field to given value.

### HasSite

`func (o *ImageBuilds) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *ImageBuilds) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ImageBuilds) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ImageBuilds) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ImageBuilds) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetName

`func (o *ImageBuilds) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageBuilds) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageBuilds) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageBuilds) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ImageBuilds) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageBuilds) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageBuilds) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageBuilds) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBootScript

`func (o *ImageBuilds) GetBootScript() ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *ImageBuilds) GetBootScriptOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *ImageBuilds) SetBootScript(v ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *ImageBuilds) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.

### GetBootCommand

`func (o *ImageBuilds) GetBootCommand() string`

GetBootCommand returns the BootCommand field if non-nil, zero value otherwise.

### GetBootCommandOk

`func (o *ImageBuilds) GetBootCommandOk() (*string, bool)`

GetBootCommandOk returns a tuple with the BootCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootCommand

`func (o *ImageBuilds) SetBootCommand(v string)`

SetBootCommand sets BootCommand field to given value.

### HasBootCommand

`func (o *ImageBuilds) HasBootCommand() bool`

HasBootCommand returns a boolean if a field has been set.

### GetPreseedScript

`func (o *ImageBuilds) GetPreseedScript() ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript`

GetPreseedScript returns the PreseedScript field if non-nil, zero value otherwise.

### GetPreseedScriptOk

`func (o *ImageBuilds) GetPreseedScriptOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript, bool)`

GetPreseedScriptOk returns a tuple with the PreseedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScript

`func (o *ImageBuilds) SetPreseedScript(v ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript)`

SetPreseedScript sets PreseedScript field to given value.

### HasPreseedScript

`func (o *ImageBuilds) HasPreseedScript() bool`

HasPreseedScript returns a boolean if a field has been set.

### GetScripts

`func (o *ImageBuilds) GetScripts() []ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *ImageBuilds) GetScriptsOk() (*[]ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *ImageBuilds) SetScripts(v []ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *ImageBuilds) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetSshUsername

`func (o *ImageBuilds) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ImageBuilds) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ImageBuilds) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ImageBuilds) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ImageBuilds) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ImageBuilds) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ImageBuilds) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ImageBuilds) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetStorageProvider

`func (o *ImageBuilds) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ImageBuilds) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ImageBuilds) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ImageBuilds) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetBuildOutputName

`func (o *ImageBuilds) GetBuildOutputName() string`

GetBuildOutputName returns the BuildOutputName field if non-nil, zero value otherwise.

### GetBuildOutputNameOk

`func (o *ImageBuilds) GetBuildOutputNameOk() (*string, bool)`

GetBuildOutputNameOk returns a tuple with the BuildOutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputName

`func (o *ImageBuilds) SetBuildOutputName(v string)`

SetBuildOutputName sets BuildOutputName field to given value.

### HasBuildOutputName

`func (o *ImageBuilds) HasBuildOutputName() bool`

HasBuildOutputName returns a boolean if a field has been set.

### GetConversionFormats

`func (o *ImageBuilds) GetConversionFormats() string`

GetConversionFormats returns the ConversionFormats field if non-nil, zero value otherwise.

### GetConversionFormatsOk

`func (o *ImageBuilds) GetConversionFormatsOk() (*string, bool)`

GetConversionFormatsOk returns a tuple with the ConversionFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFormats

`func (o *ImageBuilds) SetConversionFormats(v string)`

SetConversionFormats sets ConversionFormats field to given value.

### HasConversionFormats

`func (o *ImageBuilds) HasConversionFormats() bool`

HasConversionFormats returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *ImageBuilds) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *ImageBuilds) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *ImageBuilds) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *ImageBuilds) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *ImageBuilds) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *ImageBuilds) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *ImageBuilds) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *ImageBuilds) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetKeepResults

`func (o *ImageBuilds) GetKeepResults() int64`

GetKeepResults returns the KeepResults field if non-nil, zero value otherwise.

### GetKeepResultsOk

`func (o *ImageBuilds) GetKeepResultsOk() (*int64, bool)`

GetKeepResultsOk returns a tuple with the KeepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepResults

`func (o *ImageBuilds) SetKeepResults(v int64)`

SetKeepResults sets KeepResults field to given value.

### HasKeepResults

`func (o *ImageBuilds) HasKeepResults() bool`

HasKeepResults returns a boolean if a field has been set.

### GetConfig

`func (o *ImageBuilds) GetConfig() ListImageBuilds200ResponseAllOfImageBuildsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ImageBuilds) GetConfigOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ImageBuilds) SetConfig(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ImageBuilds) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLastResult

`func (o *ImageBuilds) GetLastResult() ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *ImageBuilds) GetLastResultOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *ImageBuilds) SetLastResult(v ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *ImageBuilds) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetExecutionCount

`func (o *ImageBuilds) GetExecutionCount() int64`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *ImageBuilds) GetExecutionCountOk() (*int64, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *ImageBuilds) SetExecutionCount(v int64)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *ImageBuilds) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


