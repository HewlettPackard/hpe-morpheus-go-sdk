# AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | CloudFormation Template in JSON | [optional] 
**Yaml** | Pointer to **string** | CloudFormation Template in YAML | [optional] 
**Git** | Pointer to [**AddBlueprintRequestOneOf1CloudFormationGit**](AddBlueprintRequestOneOf1CloudFormationGit.md) |  | [optional] 
**IAM** | Pointer to [**AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationIAM**](AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationIAM.md) |  | [optional] 
**CAPABILITY_NAMED_IAM** | Pointer to [**AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationCAPABILITYNAMEDIAM**](AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationCAPABILITYNAMEDIAM.md) |  | [optional] 
**CAPABILITY_AUTO_EXPAND** | Pointer to [**AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationCAPABILITYAUTOEXPAND**](AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationCAPABILITYAUTOEXPAND.md) |  | [optional] 
**InstallAgent** | Pointer to [**AddBlueprintRequestOneOfArmInstallAgent**](AddBlueprintRequestOneOfArmInstallAgent.md) |  | [optional] 
**CloudInitEnabled** | Pointer to [**AddBlueprintRequestOneOfArmCloudInitEnabled**](AddBlueprintRequestOneOfArmCloudInitEnabled.md) |  | [optional] 

## Methods

### NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation

`func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation(configType string, ) *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation`

NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationWithDefaults

`func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationWithDefaults() *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation`

NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationWithDefaults instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetGit() AddBlueprintRequestOneOf1CloudFormationGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetGitOk() (*AddBlueprintRequestOneOf1CloudFormationGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) SetGit(v AddBlueprintRequestOneOf1CloudFormationGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetIAM

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetIAM() AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationIAM`

GetIAM returns the IAM field if non-nil, zero value otherwise.

### GetIAMOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetIAMOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationIAM, bool)`

GetIAMOk returns a tuple with the IAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAM

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) SetIAM(v AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationIAM)`

SetIAM sets IAM field to given value.

### HasIAM

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) HasIAM() bool`

HasIAM returns a boolean if a field has been set.

### GetCAPABILITY_NAMED_IAM

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetCAPABILITY_NAMED_IAM() AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationCAPABILITYNAMEDIAM`

GetCAPABILITY_NAMED_IAM returns the CAPABILITY_NAMED_IAM field if non-nil, zero value otherwise.

### GetCAPABILITY_NAMED_IAMOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetCAPABILITY_NAMED_IAMOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationCAPABILITYNAMEDIAM, bool)`

GetCAPABILITY_NAMED_IAMOk returns a tuple with the CAPABILITY_NAMED_IAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPABILITY_NAMED_IAM

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) SetCAPABILITY_NAMED_IAM(v AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationCAPABILITYNAMEDIAM)`

SetCAPABILITY_NAMED_IAM sets CAPABILITY_NAMED_IAM field to given value.

### HasCAPABILITY_NAMED_IAM

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) HasCAPABILITY_NAMED_IAM() bool`

HasCAPABILITY_NAMED_IAM returns a boolean if a field has been set.

### GetCAPABILITY_AUTO_EXPAND

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetCAPABILITY_AUTO_EXPAND() AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationCAPABILITYAUTOEXPAND`

GetCAPABILITY_AUTO_EXPAND returns the CAPABILITY_AUTO_EXPAND field if non-nil, zero value otherwise.

### GetCAPABILITY_AUTO_EXPANDOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetCAPABILITY_AUTO_EXPANDOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationCAPABILITYAUTOEXPAND, bool)`

GetCAPABILITY_AUTO_EXPANDOk returns a tuple with the CAPABILITY_AUTO_EXPAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPABILITY_AUTO_EXPAND

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) SetCAPABILITY_AUTO_EXPAND(v AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormationCAPABILITYAUTOEXPAND)`

SetCAPABILITY_AUTO_EXPAND sets CAPABILITY_AUTO_EXPAND field to given value.

### HasCAPABILITY_AUTO_EXPAND

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) HasCAPABILITY_AUTO_EXPAND() bool`

HasCAPABILITY_AUTO_EXPAND returns a boolean if a field has been set.

### GetInstallAgent

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetInstallAgent() AddBlueprintRequestOneOfArmInstallAgent`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetInstallAgentOk() (*AddBlueprintRequestOneOfArmInstallAgent, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) SetInstallAgent(v AddBlueprintRequestOneOfArmInstallAgent)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetCloudInitEnabled() AddBlueprintRequestOneOfArmCloudInitEnabled`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) GetCloudInitEnabledOk() (*AddBlueprintRequestOneOfArmCloudInitEnabled, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) SetCloudInitEnabled(v AddBlueprintRequestOneOfArmCloudInitEnabled)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


