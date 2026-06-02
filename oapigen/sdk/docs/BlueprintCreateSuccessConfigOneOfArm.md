# BlueprintCreateSuccessConfigOneOfArm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | ARM Template in JSON | [optional] 
**Yaml** | Pointer to **string** | ARM Template in YAML | [optional] 
**Git** | Pointer to [**BlueprintCreateSuccessConfigOneOfArmGit**](BlueprintCreateSuccessConfigOneOfArmGit.md) |  | [optional] 
**OsType** | Pointer to **string** | OS Type | [optional] 
**InstallAgent** | Pointer to [**BlueprintCreateSuccessConfigOneOfArmInstallAgent**](BlueprintCreateSuccessConfigOneOfArmInstallAgent.md) |  | [optional] 
**CloudInitEnabled** | Pointer to [**BlueprintCreateSuccessConfigOneOfArmCloudInitEnabled**](BlueprintCreateSuccessConfigOneOfArmCloudInitEnabled.md) |  | [optional] 

## Methods

### NewBlueprintCreateSuccessConfigOneOfArm

`func NewBlueprintCreateSuccessConfigOneOfArm(configType string, ) *BlueprintCreateSuccessConfigOneOfArm`

NewBlueprintCreateSuccessConfigOneOfArm instantiates a new BlueprintCreateSuccessConfigOneOfArm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetConfigType

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintCreateSuccessConfigOneOfArm) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *BlueprintCreateSuccessConfigOneOfArm) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *BlueprintCreateSuccessConfigOneOfArm) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *BlueprintCreateSuccessConfigOneOfArm) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *BlueprintCreateSuccessConfigOneOfArm) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetGit() BlueprintCreateSuccessConfigOneOfArmGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetGitOk() (*BlueprintCreateSuccessConfigOneOfArmGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintCreateSuccessConfigOneOfArm) SetGit(v BlueprintCreateSuccessConfigOneOfArmGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintCreateSuccessConfigOneOfArm) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetOsType

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *BlueprintCreateSuccessConfigOneOfArm) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *BlueprintCreateSuccessConfigOneOfArm) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetInstallAgent

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetInstallAgent() BlueprintCreateSuccessConfigOneOfArmInstallAgent`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetInstallAgentOk() (*BlueprintCreateSuccessConfigOneOfArmInstallAgent, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *BlueprintCreateSuccessConfigOneOfArm) SetInstallAgent(v BlueprintCreateSuccessConfigOneOfArmInstallAgent)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *BlueprintCreateSuccessConfigOneOfArm) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetCloudInitEnabled() BlueprintCreateSuccessConfigOneOfArmCloudInitEnabled`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *BlueprintCreateSuccessConfigOneOfArm) GetCloudInitEnabledOk() (*BlueprintCreateSuccessConfigOneOfArmCloudInitEnabled, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *BlueprintCreateSuccessConfigOneOfArm) SetCloudInitEnabled(v BlueprintCreateSuccessConfigOneOfArmCloudInitEnabled)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *BlueprintCreateSuccessConfigOneOfArm) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


