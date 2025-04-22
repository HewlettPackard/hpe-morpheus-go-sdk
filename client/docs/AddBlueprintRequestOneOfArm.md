# AddBlueprintRequestOneOfArm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | ARM Template in JSON | [optional] 
**Yaml** | Pointer to **string** | ARM Template in YAML | [optional] 
**Git** | Pointer to [**AddBlueprintRequestOneOfArmGit**](AddBlueprintRequestOneOfArmGit.md) |  | [optional] 
**OsType** | Pointer to **string** | OS Type | [optional] 
**InstallAgent** | Pointer to [**AddBlueprintRequestOneOfArmInstallAgent**](AddBlueprintRequestOneOfArmInstallAgent.md) |  | [optional] 
**CloudInitEnabled** | Pointer to [**AddBlueprintRequestOneOfArmCloudInitEnabled**](AddBlueprintRequestOneOfArmCloudInitEnabled.md) |  | [optional] 

## Methods

### NewAddBlueprintRequestOneOfArm

`func NewAddBlueprintRequestOneOfArm(configType string, ) *AddBlueprintRequestOneOfArm`

NewAddBlueprintRequestOneOfArm instantiates a new AddBlueprintRequestOneOfArm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprintRequestOneOfArmWithDefaults

`func NewAddBlueprintRequestOneOfArmWithDefaults() *AddBlueprintRequestOneOfArm`

NewAddBlueprintRequestOneOfArmWithDefaults instantiates a new AddBlueprintRequestOneOfArm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *AddBlueprintRequestOneOfArm) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *AddBlueprintRequestOneOfArm) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *AddBlueprintRequestOneOfArm) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *AddBlueprintRequestOneOfArm) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *AddBlueprintRequestOneOfArm) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *AddBlueprintRequestOneOfArm) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *AddBlueprintRequestOneOfArm) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *AddBlueprintRequestOneOfArm) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *AddBlueprintRequestOneOfArm) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *AddBlueprintRequestOneOfArm) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *AddBlueprintRequestOneOfArm) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *AddBlueprintRequestOneOfArm) GetGit() AddBlueprintRequestOneOfArmGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *AddBlueprintRequestOneOfArm) GetGitOk() (*AddBlueprintRequestOneOfArmGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *AddBlueprintRequestOneOfArm) SetGit(v AddBlueprintRequestOneOfArmGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *AddBlueprintRequestOneOfArm) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetOsType

`func (o *AddBlueprintRequestOneOfArm) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *AddBlueprintRequestOneOfArm) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *AddBlueprintRequestOneOfArm) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *AddBlueprintRequestOneOfArm) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetInstallAgent

`func (o *AddBlueprintRequestOneOfArm) GetInstallAgent() AddBlueprintRequestOneOfArmInstallAgent`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *AddBlueprintRequestOneOfArm) GetInstallAgentOk() (*AddBlueprintRequestOneOfArmInstallAgent, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *AddBlueprintRequestOneOfArm) SetInstallAgent(v AddBlueprintRequestOneOfArmInstallAgent)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *AddBlueprintRequestOneOfArm) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *AddBlueprintRequestOneOfArm) GetCloudInitEnabled() AddBlueprintRequestOneOfArmCloudInitEnabled`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *AddBlueprintRequestOneOfArm) GetCloudInitEnabledOk() (*AddBlueprintRequestOneOfArmCloudInitEnabled, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *AddBlueprintRequestOneOfArm) SetCloudInitEnabled(v AddBlueprintRequestOneOfArmCloudInitEnabled)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *AddBlueprintRequestOneOfArm) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


