# ImageBuildCreate

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

### NewImageBuildCreate

`func NewImageBuildCreate() *ImageBuildCreate`

NewImageBuildCreate instantiates a new ImageBuildCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuildCreateWithDefaults

`func NewImageBuildCreateWithDefaults() *ImageBuildCreate`

NewImageBuildCreateWithDefaults instantiates a new ImageBuildCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ImageBuildCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageBuildCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageBuildCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageBuildCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ImageBuildCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageBuildCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageBuildCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageBuildCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ImageBuildCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageBuildCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageBuildCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ImageBuildCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSite

`func (o *ImageBuildCreate) GetSite() AddImageBuildRequestImageBuildSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ImageBuildCreate) GetSiteOk() (*AddImageBuildRequestImageBuildSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ImageBuildCreate) SetSite(v AddImageBuildRequestImageBuildSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *ImageBuildCreate) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *ImageBuildCreate) GetZone() AddImageBuildRequestImageBuildZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ImageBuildCreate) GetZoneOk() (*AddImageBuildRequestImageBuildZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ImageBuildCreate) SetZone(v AddImageBuildRequestImageBuildZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ImageBuildCreate) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetConfig

`func (o *ImageBuildCreate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ImageBuildCreate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ImageBuildCreate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ImageBuildCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetBootScript

`func (o *ImageBuildCreate) GetBootScript() AddImageBuildRequestImageBuildBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *ImageBuildCreate) GetBootScriptOk() (*AddImageBuildRequestImageBuildBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *ImageBuildCreate) SetBootScript(v AddImageBuildRequestImageBuildBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *ImageBuildCreate) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.

### GetPreseedScript

`func (o *ImageBuildCreate) GetPreseedScript() AddImageBuildRequestImageBuildPreseedScript`

GetPreseedScript returns the PreseedScript field if non-nil, zero value otherwise.

### GetPreseedScriptOk

`func (o *ImageBuildCreate) GetPreseedScriptOk() (*AddImageBuildRequestImageBuildPreseedScript, bool)`

GetPreseedScriptOk returns a tuple with the PreseedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScript

`func (o *ImageBuildCreate) SetPreseedScript(v AddImageBuildRequestImageBuildPreseedScript)`

SetPreseedScript sets PreseedScript field to given value.

### HasPreseedScript

`func (o *ImageBuildCreate) HasPreseedScript() bool`

HasPreseedScript returns a boolean if a field has been set.

### GetSshUsername

`func (o *ImageBuildCreate) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ImageBuildCreate) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ImageBuildCreate) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ImageBuildCreate) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ImageBuildCreate) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ImageBuildCreate) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ImageBuildCreate) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ImageBuildCreate) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetStorageProvider

`func (o *ImageBuildCreate) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ImageBuildCreate) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ImageBuildCreate) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ImageBuildCreate) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *ImageBuildCreate) GetIsCloudInit() string`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *ImageBuildCreate) GetIsCloudInitOk() (*string, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *ImageBuildCreate) SetIsCloudInit(v string)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *ImageBuildCreate) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetBuildOutputName

`func (o *ImageBuildCreate) GetBuildOutputName() string`

GetBuildOutputName returns the BuildOutputName field if non-nil, zero value otherwise.

### GetBuildOutputNameOk

`func (o *ImageBuildCreate) GetBuildOutputNameOk() (*string, bool)`

GetBuildOutputNameOk returns a tuple with the BuildOutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputName

`func (o *ImageBuildCreate) SetBuildOutputName(v string)`

SetBuildOutputName sets BuildOutputName field to given value.

### HasBuildOutputName

`func (o *ImageBuildCreate) HasBuildOutputName() bool`

HasBuildOutputName returns a boolean if a field has been set.

### GetConversionFormats

`func (o *ImageBuildCreate) GetConversionFormats() string`

GetConversionFormats returns the ConversionFormats field if non-nil, zero value otherwise.

### GetConversionFormatsOk

`func (o *ImageBuildCreate) GetConversionFormatsOk() (*string, bool)`

GetConversionFormatsOk returns a tuple with the ConversionFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFormats

`func (o *ImageBuildCreate) SetConversionFormats(v string)`

SetConversionFormats sets ConversionFormats field to given value.

### HasConversionFormats

`func (o *ImageBuildCreate) HasConversionFormats() bool`

HasConversionFormats returns a boolean if a field has been set.

### SetConversionFormatsNil

`func (o *ImageBuildCreate) SetConversionFormatsNil(b bool)`

 SetConversionFormatsNil sets the value for ConversionFormats to be an explicit nil

### UnsetConversionFormats
`func (o *ImageBuildCreate) UnsetConversionFormats()`

UnsetConversionFormats ensures that no value is present for ConversionFormats, not even an explicit nil
### GetKeepResults

`func (o *ImageBuildCreate) GetKeepResults() int64`

GetKeepResults returns the KeepResults field if non-nil, zero value otherwise.

### GetKeepResultsOk

`func (o *ImageBuildCreate) GetKeepResultsOk() (*int64, bool)`

GetKeepResultsOk returns a tuple with the KeepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepResults

`func (o *ImageBuildCreate) SetKeepResults(v int64)`

SetKeepResults sets KeepResults field to given value.

### HasKeepResults

`func (o *ImageBuildCreate) HasKeepResults() bool`

HasKeepResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


