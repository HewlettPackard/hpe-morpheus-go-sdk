# UpdateBlueprintRequestOneOfArm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | ARM Template in JSON | [optional] 
**Yaml** | Pointer to **string** | ARM Template in YAML | [optional] 
**Git** | Pointer to [**UpdateBlueprintRequestOneOfArmGit**](UpdateBlueprintRequestOneOfArmGit.md) |  | [optional] 
**OsType** | Pointer to **string** | OS Type | [optional] 
**InstallAgent** | Pointer to [**UpdateBlueprintRequestOneOfArmInstallAgent**](UpdateBlueprintRequestOneOfArmInstallAgent.md) |  | [optional] 
**CloudInitEnabled** | Pointer to [**UpdateBlueprintRequestOneOfArmCloudInitEnabled**](UpdateBlueprintRequestOneOfArmCloudInitEnabled.md) |  | [optional] 

## Methods

### NewUpdateBlueprintRequestOneOfArm

`func NewUpdateBlueprintRequestOneOfArm(configType string, ) *UpdateBlueprintRequestOneOfArm`

NewUpdateBlueprintRequestOneOfArm instantiates a new UpdateBlueprintRequestOneOfArm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBlueprintRequestOneOfArmWithDefaults

`func NewUpdateBlueprintRequestOneOfArmWithDefaults() *UpdateBlueprintRequestOneOfArm`

NewUpdateBlueprintRequestOneOfArmWithDefaults instantiates a new UpdateBlueprintRequestOneOfArm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *UpdateBlueprintRequestOneOfArm) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *UpdateBlueprintRequestOneOfArm) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *UpdateBlueprintRequestOneOfArm) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *UpdateBlueprintRequestOneOfArm) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateBlueprintRequestOneOfArm) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateBlueprintRequestOneOfArm) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateBlueprintRequestOneOfArm) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetYaml

`func (o *UpdateBlueprintRequestOneOfArm) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *UpdateBlueprintRequestOneOfArm) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *UpdateBlueprintRequestOneOfArm) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *UpdateBlueprintRequestOneOfArm) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *UpdateBlueprintRequestOneOfArm) GetGit() UpdateBlueprintRequestOneOfArmGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *UpdateBlueprintRequestOneOfArm) GetGitOk() (*UpdateBlueprintRequestOneOfArmGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *UpdateBlueprintRequestOneOfArm) SetGit(v UpdateBlueprintRequestOneOfArmGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *UpdateBlueprintRequestOneOfArm) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetOsType

`func (o *UpdateBlueprintRequestOneOfArm) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *UpdateBlueprintRequestOneOfArm) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *UpdateBlueprintRequestOneOfArm) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *UpdateBlueprintRequestOneOfArm) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetInstallAgent

`func (o *UpdateBlueprintRequestOneOfArm) GetInstallAgent() UpdateBlueprintRequestOneOfArmInstallAgent`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *UpdateBlueprintRequestOneOfArm) GetInstallAgentOk() (*UpdateBlueprintRequestOneOfArmInstallAgent, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *UpdateBlueprintRequestOneOfArm) SetInstallAgent(v UpdateBlueprintRequestOneOfArmInstallAgent)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *UpdateBlueprintRequestOneOfArm) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetCloudInitEnabled

`func (o *UpdateBlueprintRequestOneOfArm) GetCloudInitEnabled() UpdateBlueprintRequestOneOfArmCloudInitEnabled`

GetCloudInitEnabled returns the CloudInitEnabled field if non-nil, zero value otherwise.

### GetCloudInitEnabledOk

`func (o *UpdateBlueprintRequestOneOfArm) GetCloudInitEnabledOk() (*UpdateBlueprintRequestOneOfArmCloudInitEnabled, bool)`

GetCloudInitEnabledOk returns a tuple with the CloudInitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitEnabled

`func (o *UpdateBlueprintRequestOneOfArm) SetCloudInitEnabled(v UpdateBlueprintRequestOneOfArmCloudInitEnabled)`

SetCloudInitEnabled sets CloudInitEnabled field to given value.

### HasCloudInitEnabled

`func (o *UpdateBlueprintRequestOneOfArm) HasCloudInitEnabled() bool`

HasCloudInitEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


