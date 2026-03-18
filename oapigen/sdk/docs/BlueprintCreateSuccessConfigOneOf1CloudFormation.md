# BlueprintCreateSuccessConfigOneOf1CloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | CloudFormation Template in JSON | [optional] 
**Yaml** | Pointer to **string** | CloudFormation Template in YAML | [optional] 
**Git** | Pointer to [**BlueprintCreateSuccessConfigOneOf1CloudFormationGit**](BlueprintCreateSuccessConfigOneOf1CloudFormationGit.md) |  | [optional] 
**IAM** | Pointer to [**BlueprintCreateSuccessConfigOneOf1CloudFormationIAM**](BlueprintCreateSuccessConfigOneOf1CloudFormationIAM.md) |  | [optional] 
**CAPABILITY_NAMED_IAM** | Pointer to [**BlueprintCreateSuccessConfigOneOf1CloudFormationCAPABILITYNAMEDIAM**](BlueprintCreateSuccessConfigOneOf1CloudFormationCAPABILITYNAMEDIAM.md) |  | [optional] 
**CAPABILITY_AUTO_EXPAND** | Pointer to [**BlueprintCreateSuccessConfigOneOf1CloudFormationCAPABILITYAUTOEXPAND**](BlueprintCreateSuccessConfigOneOf1CloudFormationCAPABILITYAUTOEXPAND.md) |  | [optional] 
**InstallAgent** | Pointer to [**BlueprintCreateSuccessConfigOneOf1CloudFormationInstallAgent**](BlueprintCreateSuccessConfigOneOf1CloudFormationInstallAgent.md) |  | [optional] 
**CloudInitEnabled** | Pointer to [**BlueprintCreateSuccessConfigOneOf1CloudFormationCloudInitEnabled**](BlueprintCreateSuccessConfigOneOf1CloudFormationCloudInitEnabled.md) |  | [optional] 

## Methods

### NewBlueprintCreateSuccessConfigOneOf1CloudFormation

`func NewBlueprintCreateSuccessConfigOneOf1CloudFormation(configType string, ) *BlueprintCreateSuccessConfigOneOf1CloudFormation`

NewBlueprintCreateSuccessConfigOneOf1CloudFormation instantiates a new BlueprintCreateSuccessConfigOneOf1CloudFormation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCreateSuccessConfigOneOf1CloudFormationWithDefaults

`func NewBlueprintCreateSuccessConfigOneOf1CloudFormationWithDefaults() *BlueprintCreateSuccessConfigOneOf1CloudFormation`

NewBlueprintCreateSuccessConfigOneOf1CloudFormationWithDefaults instantiates a new BlueprintCreateSuccessConfigOneOf1CloudFormation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetGit() BlueprintCreateSuccessConfigOneOf1CloudFormationGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetGitOk() (*BlueprintCreateSuccessConfigOneOf1CloudFormationGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) SetGit(v BlueprintCreateSuccessConfigOneOf1CloudFormationGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetIAM

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetIAM() BlueprintCreateSuccessConfigOneOf1CloudFormationIAM`

GetIAM returns the IAM field if non-nil, zero value otherwise.

### GetIAMOk

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetIAMOk() (*BlueprintCreateSuccessConfigOneOf1CloudFormationIAM, bool)`

GetIAMOk returns a tuple with the IAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAM

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) SetIAM(v BlueprintCreateSuccessConfigOneOf1CloudFormationIAM)`

SetIAM sets IAM field to given value.

### HasIAM

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) HasIAM() bool`

HasIAM returns a boolean if a field has been set.

### GetCAPABILITY_NAMED_IAM

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetCAPABILITY_NAMED_IAM() BlueprintCreateSuccessConfigOneOf1CloudFormationCAPABILITYNAMEDIAM`

GetCAPABILITY_NAMED_IAM returns the CAPABILITY_NAMED_IAM field if non-nil, zero value otherwise.

### GetCAPABILITY_NAMED_IAMOk

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetCAPABILITY_NAMED_IAMOk() (*BlueprintCreateSuccessConfigOneOf1CloudFormationCAPABILITYNAMEDIAM, bool)`

GetCAPABILITY_NAMED_IAMOk returns a tuple with the CAPABILITY_NAMED_IAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPABILITY_NAMED_IAM

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) SetCAPABILITY_NAMED_IAM(v BlueprintCreateSuccessConfigOneOf1CloudFormationCAPABILITYNAMEDIAM)`

SetCAPABILITY_NAMED_IAM sets CAPABILITY_NAMED_IAM field to given value.

### HasCAPABILITY_NAMED_IAM

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) HasCAPABILITY_NAMED_IAM() bool`

HasCAPABILITY_NAMED_IAM returns a boolean if a field has been set.

### GetCAPABILITY_AUTO_EXPAND

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetCAPABILITY_AUTO_EXPAND() BlueprintCreateSuccessConfigOneOf1CloudFormationCAPABILITYAUTOEXPAND`

GetCAPABILITY_AUTO_EXPAND returns the CAPABILITY_AUTO_EXPAND field if non-nil, zero value otherwise.

### GetCAPABILITY_AUTO_EXPANDOk

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetCAPABILITY_AUTO_EXPANDOk() (*BlueprintCreateSuccessConfigOneOf1CloudFormationCAPABILITYAUTOEXPAND, bool)`

GetCAPABILITY_AUTO_EXPANDOk returns a tuple with the CAPABILITY_AUTO_EXPAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPABILITY_AUTO_EXPAND

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) SetCAPABILITY_AUTO_EXPAND(v BlueprintCreateSuccessConfigOneOf1CloudFormationCAPABILITYAUTOEXPAND)`

SetCAPABILITY_AUTO_EXPAND sets CAPABILITY_AUTO_EXPAND field to given value.

### HasCAPABILITY_AUTO_EXPAND

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) HasCAPABILITY_AUTO_EXPAND() bool`

HasCAPABILITY_AUTO_EXPAND returns a boolean if a field has been set.

### GetInstallAgent

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetInstallAgent() BlueprintCreateSuccessConfigOneOf1CloudFormationInstallAgent`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetInstallAgentOk() (*BlueprintCreateSuccessConfigOneOf1CloudFormationInstallAgent, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) SetInstallAgent(v BlueprintCreateSuccessConfigOneOf1CloudFormationInstallAgent)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetCloudInitEnabled() BlueprintCreateSuccessConfigOneOf1CloudFormationCloudInitEnabled`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) GetCloudInitEnabledOk() (*BlueprintCreateSuccessConfigOneOf1CloudFormationCloudInitEnabled, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) SetCloudInitEnabled(v BlueprintCreateSuccessConfigOneOf1CloudFormationCloudInitEnabled)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *BlueprintCreateSuccessConfigOneOf1CloudFormation) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


