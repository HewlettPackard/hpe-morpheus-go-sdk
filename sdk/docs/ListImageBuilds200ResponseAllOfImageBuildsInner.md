# ListImageBuilds200ResponseAllOfImageBuildsInner

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

### NewListImageBuilds200ResponseAllOfImageBuildsInner

`func NewListImageBuilds200ResponseAllOfImageBuildsInner() *ListImageBuilds200ResponseAllOfImageBuildsInner`

NewListImageBuilds200ResponseAllOfImageBuildsInner instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageBuilds200ResponseAllOfImageBuildsInnerWithDefaults

`func NewListImageBuilds200ResponseAllOfImageBuildsInnerWithDefaults() *ListImageBuilds200ResponseAllOfImageBuildsInner`

NewListImageBuilds200ResponseAllOfImageBuildsInnerWithDefaults instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetType

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSite

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetSite() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetSiteOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetSite(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetSite sets Site field to given value.

### HasSite

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetName

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBootScript

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetBootScript() ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetBootScriptOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetBootScript(v ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.

### GetBootCommand

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetBootCommand() string`

GetBootCommand returns the BootCommand field if non-nil, zero value otherwise.

### GetBootCommandOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetBootCommandOk() (*string, bool)`

GetBootCommandOk returns a tuple with the BootCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootCommand

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetBootCommand(v string)`

SetBootCommand sets BootCommand field to given value.

### HasBootCommand

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasBootCommand() bool`

HasBootCommand returns a boolean if a field has been set.

### GetPreseedScript

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetPreseedScript() ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript`

GetPreseedScript returns the PreseedScript field if non-nil, zero value otherwise.

### GetPreseedScriptOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetPreseedScriptOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript, bool)`

GetPreseedScriptOk returns a tuple with the PreseedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScript

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetPreseedScript(v ListImageBuilds200ResponseAllOfImageBuildsInnerBootScript)`

SetPreseedScript sets PreseedScript field to given value.

### HasPreseedScript

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasPreseedScript() bool`

HasPreseedScript returns a boolean if a field has been set.

### GetScripts

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetScripts() []ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetScriptsOk() (*[]ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetScripts(v []ListImageBuilds200ResponseAllOfImageBuildsInnerScriptsInner)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetSshUsername

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetStorageProvider

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetBuildOutputName

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetBuildOutputName() string`

GetBuildOutputName returns the BuildOutputName field if non-nil, zero value otherwise.

### GetBuildOutputNameOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetBuildOutputNameOk() (*string, bool)`

GetBuildOutputNameOk returns a tuple with the BuildOutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputName

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetBuildOutputName(v string)`

SetBuildOutputName sets BuildOutputName field to given value.

### HasBuildOutputName

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasBuildOutputName() bool`

HasBuildOutputName returns a boolean if a field has been set.

### GetConversionFormats

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetConversionFormats() string`

GetConversionFormats returns the ConversionFormats field if non-nil, zero value otherwise.

### GetConversionFormatsOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetConversionFormatsOk() (*string, bool)`

GetConversionFormatsOk returns a tuple with the ConversionFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFormats

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetConversionFormats(v string)`

SetConversionFormats sets ConversionFormats field to given value.

### HasConversionFormats

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasConversionFormats() bool`

HasConversionFormats returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetKeepResults

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetKeepResults() int64`

GetKeepResults returns the KeepResults field if non-nil, zero value otherwise.

### GetKeepResultsOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetKeepResultsOk() (*int64, bool)`

GetKeepResultsOk returns a tuple with the KeepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepResults

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetKeepResults(v int64)`

SetKeepResults sets KeepResults field to given value.

### HasKeepResults

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasKeepResults() bool`

HasKeepResults returns a boolean if a field has been set.

### GetConfig

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetConfig() ListImageBuilds200ResponseAllOfImageBuildsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetConfigOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetConfig(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLastResult

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetLastResult() ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetLastResultOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetLastResult(v ListImageBuilds200ResponseAllOfImageBuildsInnerLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetExecutionCount

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetExecutionCount() int64`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) GetExecutionCountOk() (*int64, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) SetExecutionCount(v int64)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInner) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


