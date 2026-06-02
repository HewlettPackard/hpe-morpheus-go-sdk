# UpdateBlueprintRequestOneOf1CloudFormation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | CloudFormation Template in JSON | [optional] 
**Yaml** | Pointer to **string** | CloudFormation Template in YAML | [optional] 
**Git** | Pointer to [**UpdateBlueprintRequestOneOf1CloudFormationGit**](UpdateBlueprintRequestOneOf1CloudFormationGit.md) |  | [optional] 
**IAM** | Pointer to **bool** | CloudFormation Attribute CAPABILITY_IAM | [optional] [default to false]
**CAPABILITY_NAMED_IAM** | Pointer to **bool** | CloudFormation Attribute CAPABILITY_NAMED_IAM | [optional] [default to false]
**CAPABILITY_AUTO_EXPAND** | Pointer to **bool** | CloudFormation Attribute CAPABILITY_AUTO_EXPAND | [optional] [default to false]
**InstallAgent** | Pointer to **bool** | Install Morpheus Agent | [optional] [default to false]
**CloudInitEnabled** | Pointer to **bool** | Cloud Init Enabled | [optional] [default to false]

## Methods

### NewUpdateBlueprintRequestOneOf1CloudFormation

`func NewUpdateBlueprintRequestOneOf1CloudFormation(configType string, ) *UpdateBlueprintRequestOneOf1CloudFormation`

NewUpdateBlueprintRequestOneOf1CloudFormation instantiates a new UpdateBlueprintRequestOneOf1CloudFormation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetConfigType

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetGit() UpdateBlueprintRequestOneOf1CloudFormationGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetGitOk() (*UpdateBlueprintRequestOneOf1CloudFormationGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) SetGit(v UpdateBlueprintRequestOneOf1CloudFormationGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetIAM

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetIAM() bool`

GetIAM returns the IAM field if non-nil, zero value otherwise.

### GetIAMOk

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetIAMOk() (*bool, bool)`

GetIAMOk returns a tuple with the IAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIAM

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) SetIAM(v bool)`

SetIAM sets IAM field to given value.

### HasIAM

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) HasIAM() bool`

HasIAM returns a boolean if a field has been set.

### GetCAPABILITY_NAMED_IAM

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetCAPABILITY_NAMED_IAM() bool`

GetCAPABILITY_NAMED_IAM returns the CAPABILITY_NAMED_IAM field if non-nil, zero value otherwise.

### GetCAPABILITY_NAMED_IAMOk

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetCAPABILITY_NAMED_IAMOk() (*bool, bool)`

GetCAPABILITY_NAMED_IAMOk returns a tuple with the CAPABILITY_NAMED_IAM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPABILITY_NAMED_IAM

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) SetCAPABILITY_NAMED_IAM(v bool)`

SetCAPABILITY_NAMED_IAM sets CAPABILITY_NAMED_IAM field to given value.

### HasCAPABILITY_NAMED_IAM

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) HasCAPABILITY_NAMED_IAM() bool`

HasCAPABILITY_NAMED_IAM returns a boolean if a field has been set.

### GetCAPABILITY_AUTO_EXPAND

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetCAPABILITY_AUTO_EXPAND() bool`

GetCAPABILITY_AUTO_EXPAND returns the CAPABILITY_AUTO_EXPAND field if non-nil, zero value otherwise.

### GetCAPABILITY_AUTO_EXPANDOk

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetCAPABILITY_AUTO_EXPANDOk() (*bool, bool)`

GetCAPABILITY_AUTO_EXPANDOk returns a tuple with the CAPABILITY_AUTO_EXPAND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAPABILITY_AUTO_EXPAND

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) SetCAPABILITY_AUTO_EXPAND(v bool)`

SetCAPABILITY_AUTO_EXPAND sets CAPABILITY_AUTO_EXPAND field to given value.

### HasCAPABILITY_AUTO_EXPAND

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) HasCAPABILITY_AUTO_EXPAND() bool`

HasCAPABILITY_AUTO_EXPAND returns a boolean if a field has been set.

### GetInstallAgent

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetCloudInitEnabled() bool`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) GetCloudInitEnabledOk() (*bool, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) SetCloudInitEnabled(v bool)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *UpdateBlueprintRequestOneOf1CloudFormation) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


