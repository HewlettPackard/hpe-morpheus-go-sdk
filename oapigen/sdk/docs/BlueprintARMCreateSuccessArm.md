# BlueprintARMCreateSuccessArm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | ARM Template in JSON | [optional] 
**Yaml** | Pointer to **string** | ARM Template in YAML | [optional] 
**Git** | Pointer to [**BlueprintARMCreateSuccessArmGit**](BlueprintARMCreateSuccessArmGit.md) |  | [optional] 
**OsType** | Pointer to **string** | OS Type | [optional] 
**InstallAgent** | Pointer to [**BlueprintARMCreateSuccessArmInstallAgent**](BlueprintARMCreateSuccessArmInstallAgent.md) |  | [optional] 
**CloudInitEnabled** | Pointer to [**BlueprintARMCreateSuccessArmCloudInitEnabled**](BlueprintARMCreateSuccessArmCloudInitEnabled.md) |  | [optional] 

## Methods

### NewBlueprintARMCreateSuccessArm

`func NewBlueprintARMCreateSuccessArm(configType string, ) *BlueprintARMCreateSuccessArm`

NewBlueprintARMCreateSuccessArm instantiates a new BlueprintARMCreateSuccessArm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetConfigType

`func (o *BlueprintARMCreateSuccessArm) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintARMCreateSuccessArm) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintARMCreateSuccessArm) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *BlueprintARMCreateSuccessArm) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *BlueprintARMCreateSuccessArm) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *BlueprintARMCreateSuccessArm) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *BlueprintARMCreateSuccessArm) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *BlueprintARMCreateSuccessArm) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *BlueprintARMCreateSuccessArm) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *BlueprintARMCreateSuccessArm) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *BlueprintARMCreateSuccessArm) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *BlueprintARMCreateSuccessArm) GetGit() BlueprintARMCreateSuccessArmGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintARMCreateSuccessArm) GetGitOk() (*BlueprintARMCreateSuccessArmGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintARMCreateSuccessArm) SetGit(v BlueprintARMCreateSuccessArmGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintARMCreateSuccessArm) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetOsType

`func (o *BlueprintARMCreateSuccessArm) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *BlueprintARMCreateSuccessArm) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *BlueprintARMCreateSuccessArm) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *BlueprintARMCreateSuccessArm) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetInstallAgent

`func (o *BlueprintARMCreateSuccessArm) GetInstallAgent() BlueprintARMCreateSuccessArmInstallAgent`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *BlueprintARMCreateSuccessArm) GetInstallAgentOk() (*BlueprintARMCreateSuccessArmInstallAgent, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *BlueprintARMCreateSuccessArm) SetInstallAgent(v BlueprintARMCreateSuccessArmInstallAgent)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *BlueprintARMCreateSuccessArm) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *BlueprintARMCreateSuccessArm) GetCloudInitEnabled() BlueprintARMCreateSuccessArmCloudInitEnabled`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *BlueprintARMCreateSuccessArm) GetCloudInitEnabledOk() (*BlueprintARMCreateSuccessArmCloudInitEnabled, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *BlueprintARMCreateSuccessArm) SetCloudInitEnabled(v BlueprintARMCreateSuccessArmCloudInitEnabled)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *BlueprintARMCreateSuccessArm) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


