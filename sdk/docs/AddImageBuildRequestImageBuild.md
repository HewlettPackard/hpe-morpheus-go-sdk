# AddImageBuildRequestImageBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the image build | [optional] 
**Description** | Pointer to **string** | A description for the image build | [optional] 
**Type** | Pointer to **string** | The image builder type. | [optional] 
**Site** | Pointer to [**AddImageBuildRequestImageBuildSite**](AddImageBuildRequestImageBuildSite.md) |  | [optional] 
**Zone** | Pointer to [**AddImageBuildRequestImageBuildZone**](AddImageBuildRequestImageBuildZone.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | A map of config values. This is the instance config that is used for provisioning. See Provisioning Types. | [optional] 
**BootScript** | Pointer to [**AddImageBuildRequestImageBuildBootScript**](AddImageBuildRequestImageBuildBootScript.md) |  | [optional] 
**PreseedScript** | Pointer to [**AddImageBuildRequestImageBuildPreseedScript**](AddImageBuildRequestImageBuildPreseedScript.md) |  | [optional] 
**SshUsername** | Pointer to **string** | SSH Username | [optional] 
**SshPassword** | Pointer to **string** | SSH Password | [optional] 
**StorageProvider** | Pointer to **string** | Storage Provider | [optional] 
**IsCloudInit** | Pointer to **string** | Cloud Init | [optional] 
**BuildOutputName** | Pointer to **string** | Build Output Name | [optional] 
**ConversionFormats** | Pointer to **NullableString** |  | [optional] 
**KeepResults** | Pointer to **int64** | Keep Results - Keep only the most recent builds. Older executions will be deleted along with their associated Virtual Images. The value 0 disables this functionality. | [optional] [default to 0]

## Methods

### NewAddImageBuildRequestImageBuild

`func NewAddImageBuildRequestImageBuild() *AddImageBuildRequestImageBuild`

NewAddImageBuildRequestImageBuild instantiates a new AddImageBuildRequestImageBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddImageBuildRequestImageBuildWithDefaults

`func NewAddImageBuildRequestImageBuildWithDefaults() *AddImageBuildRequestImageBuild`

NewAddImageBuildRequestImageBuildWithDefaults instantiates a new AddImageBuildRequestImageBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddImageBuildRequestImageBuild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddImageBuildRequestImageBuild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddImageBuildRequestImageBuild) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddImageBuildRequestImageBuild) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddImageBuildRequestImageBuild) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddImageBuildRequestImageBuild) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddImageBuildRequestImageBuild) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddImageBuildRequestImageBuild) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *AddImageBuildRequestImageBuild) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddImageBuildRequestImageBuild) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddImageBuildRequestImageBuild) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddImageBuildRequestImageBuild) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSite

`func (o *AddImageBuildRequestImageBuild) GetSite() AddImageBuildRequestImageBuildSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *AddImageBuildRequestImageBuild) GetSiteOk() (*AddImageBuildRequestImageBuildSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *AddImageBuildRequestImageBuild) SetSite(v AddImageBuildRequestImageBuildSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *AddImageBuildRequestImageBuild) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *AddImageBuildRequestImageBuild) GetZone() AddImageBuildRequestImageBuildZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddImageBuildRequestImageBuild) GetZoneOk() (*AddImageBuildRequestImageBuildZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddImageBuildRequestImageBuild) SetZone(v AddImageBuildRequestImageBuildZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddImageBuildRequestImageBuild) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetConfig

`func (o *AddImageBuildRequestImageBuild) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddImageBuildRequestImageBuild) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddImageBuildRequestImageBuild) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddImageBuildRequestImageBuild) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetBootScript

`func (o *AddImageBuildRequestImageBuild) GetBootScript() AddImageBuildRequestImageBuildBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *AddImageBuildRequestImageBuild) GetBootScriptOk() (*AddImageBuildRequestImageBuildBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *AddImageBuildRequestImageBuild) SetBootScript(v AddImageBuildRequestImageBuildBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *AddImageBuildRequestImageBuild) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.

### GetPreseedScript

`func (o *AddImageBuildRequestImageBuild) GetPreseedScript() AddImageBuildRequestImageBuildPreseedScript`

GetPreseedScript returns the PreseedScript field if non-nil, zero value otherwise.

### GetPreseedScriptOk

`func (o *AddImageBuildRequestImageBuild) GetPreseedScriptOk() (*AddImageBuildRequestImageBuildPreseedScript, bool)`

GetPreseedScriptOk returns a tuple with the PreseedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScript

`func (o *AddImageBuildRequestImageBuild) SetPreseedScript(v AddImageBuildRequestImageBuildPreseedScript)`

SetPreseedScript sets PreseedScript field to given value.

### HasPreseedScript

`func (o *AddImageBuildRequestImageBuild) HasPreseedScript() bool`

HasPreseedScript returns a boolean if a field has been set.

### GetSshUsername

`func (o *AddImageBuildRequestImageBuild) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *AddImageBuildRequestImageBuild) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *AddImageBuildRequestImageBuild) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *AddImageBuildRequestImageBuild) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *AddImageBuildRequestImageBuild) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *AddImageBuildRequestImageBuild) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *AddImageBuildRequestImageBuild) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *AddImageBuildRequestImageBuild) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetStorageProvider

`func (o *AddImageBuildRequestImageBuild) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *AddImageBuildRequestImageBuild) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *AddImageBuildRequestImageBuild) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *AddImageBuildRequestImageBuild) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *AddImageBuildRequestImageBuild) GetIsCloudInit() string`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *AddImageBuildRequestImageBuild) GetIsCloudInitOk() (*string, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *AddImageBuildRequestImageBuild) SetIsCloudInit(v string)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *AddImageBuildRequestImageBuild) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetBuildOutputName

`func (o *AddImageBuildRequestImageBuild) GetBuildOutputName() string`

GetBuildOutputName returns the BuildOutputName field if non-nil, zero value otherwise.

### GetBuildOutputNameOk

`func (o *AddImageBuildRequestImageBuild) GetBuildOutputNameOk() (*string, bool)`

GetBuildOutputNameOk returns a tuple with the BuildOutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputName

`func (o *AddImageBuildRequestImageBuild) SetBuildOutputName(v string)`

SetBuildOutputName sets BuildOutputName field to given value.

### HasBuildOutputName

`func (o *AddImageBuildRequestImageBuild) HasBuildOutputName() bool`

HasBuildOutputName returns a boolean if a field has been set.

### GetConversionFormats

`func (o *AddImageBuildRequestImageBuild) GetConversionFormats() string`

GetConversionFormats returns the ConversionFormats field if non-nil, zero value otherwise.

### GetConversionFormatsOk

`func (o *AddImageBuildRequestImageBuild) GetConversionFormatsOk() (*string, bool)`

GetConversionFormatsOk returns a tuple with the ConversionFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFormats

`func (o *AddImageBuildRequestImageBuild) SetConversionFormats(v string)`

SetConversionFormats sets ConversionFormats field to given value.

### HasConversionFormats

`func (o *AddImageBuildRequestImageBuild) HasConversionFormats() bool`

HasConversionFormats returns a boolean if a field has been set.

### SetConversionFormatsNil

`func (o *AddImageBuildRequestImageBuild) SetConversionFormatsNil(b bool)`

 SetConversionFormatsNil sets the value for ConversionFormats to be an explicit nil

### UnsetConversionFormats
`func (o *AddImageBuildRequestImageBuild) UnsetConversionFormats()`

UnsetConversionFormats ensures that no value is present for ConversionFormats, not even an explicit nil
### GetKeepResults

`func (o *AddImageBuildRequestImageBuild) GetKeepResults() int64`

GetKeepResults returns the KeepResults field if non-nil, zero value otherwise.

### GetKeepResultsOk

`func (o *AddImageBuildRequestImageBuild) GetKeepResultsOk() (*int64, bool)`

GetKeepResultsOk returns a tuple with the KeepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepResults

`func (o *AddImageBuildRequestImageBuild) SetKeepResults(v int64)`

SetKeepResults sets KeepResults field to given value.

### HasKeepResults

`func (o *AddImageBuildRequestImageBuild) HasKeepResults() bool`

HasKeepResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


