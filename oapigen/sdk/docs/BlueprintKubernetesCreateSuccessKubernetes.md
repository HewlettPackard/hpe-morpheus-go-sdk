# BlueprintKubernetesCreateSuccessKubernetes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Yaml** | Pointer to **string** | Kubernetes Spec in YAML | [optional] 
**Git** | Pointer to [**BlueprintKubernetesCreateSuccessKubernetesGit**](BlueprintKubernetesCreateSuccessKubernetesGit.md) |  | [optional] 

## Methods

### NewBlueprintKubernetesCreateSuccessKubernetes

`func NewBlueprintKubernetesCreateSuccessKubernetes(configType string, ) *BlueprintKubernetesCreateSuccessKubernetes`

NewBlueprintKubernetesCreateSuccessKubernetes instantiates a new BlueprintKubernetesCreateSuccessKubernetes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintKubernetesCreateSuccessKubernetesWithDefaults

`func NewBlueprintKubernetesCreateSuccessKubernetesWithDefaults() *BlueprintKubernetesCreateSuccessKubernetes`

NewBlueprintKubernetesCreateSuccessKubernetesWithDefaults instantiates a new BlueprintKubernetesCreateSuccessKubernetes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *BlueprintKubernetesCreateSuccessKubernetes) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintKubernetesCreateSuccessKubernetes) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintKubernetesCreateSuccessKubernetes) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetYaml

`func (o *BlueprintKubernetesCreateSuccessKubernetes) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *BlueprintKubernetesCreateSuccessKubernetes) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *BlueprintKubernetesCreateSuccessKubernetes) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *BlueprintKubernetesCreateSuccessKubernetes) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *BlueprintKubernetesCreateSuccessKubernetes) GetGit() BlueprintKubernetesCreateSuccessKubernetesGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintKubernetesCreateSuccessKubernetes) GetGitOk() (*BlueprintKubernetesCreateSuccessKubernetesGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintKubernetesCreateSuccessKubernetes) SetGit(v BlueprintKubernetesCreateSuccessKubernetesGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintKubernetesCreateSuccessKubernetes) HasGit() bool`

HasGit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


