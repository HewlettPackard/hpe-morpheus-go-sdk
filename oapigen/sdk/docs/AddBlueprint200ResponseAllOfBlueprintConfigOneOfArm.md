# AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | ARM Template in JSON | [optional] 
**Yaml** | Pointer to **string** | ARM Template in YAML | [optional] 
**Git** | Pointer to [**AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmGit**](AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmGit.md) |  | [optional] 
**OsType** | Pointer to **string** | OS Type | [optional] 
**InstallAgent** | Pointer to [**AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmInstallAgent**](AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmInstallAgent.md) |  | [optional] 
**CloudInitEnabled** | Pointer to [**AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmCloudInitEnabled**](AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmCloudInitEnabled.md) |  | [optional] 

## Methods

### NewAddBlueprint200ResponseAllOfBlueprintConfigOneOfArm

`func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOfArm(configType string, ) *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm`

NewAddBlueprint200ResponseAllOfBlueprintConfigOneOfArm instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetConfigType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetGit() AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetGitOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) SetGit(v AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetOsType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetInstallAgent

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetInstallAgent() AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmInstallAgent`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetInstallAgentOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmInstallAgent, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) SetInstallAgent(v AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmInstallAgent)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetCloudInitEnabled() AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmCloudInitEnabled`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) GetCloudInitEnabledOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmCloudInitEnabled, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) SetCloudInitEnabled(v AddBlueprint200ResponseAllOfBlueprintConfigOneOfArmCloudInitEnabled)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOfArm) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


