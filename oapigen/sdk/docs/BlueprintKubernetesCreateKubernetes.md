# BlueprintKubernetesCreateKubernetes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Yaml** | Pointer to **string** | Kubernetes Spec in YAML | [optional] 
**Git** | Pointer to [**BlueprintKubernetesCreateKubernetesGit**](BlueprintKubernetesCreateKubernetesGit.md) |  | [optional] 

## Methods

### NewBlueprintKubernetesCreateKubernetes

`func NewBlueprintKubernetesCreateKubernetes(configType string, ) *BlueprintKubernetesCreateKubernetes`

NewBlueprintKubernetesCreateKubernetes instantiates a new BlueprintKubernetesCreateKubernetes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintKubernetesCreateKubernetesWithDefaults

`func NewBlueprintKubernetesCreateKubernetesWithDefaults() *BlueprintKubernetesCreateKubernetes`

NewBlueprintKubernetesCreateKubernetesWithDefaults instantiates a new BlueprintKubernetesCreateKubernetes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *BlueprintKubernetesCreateKubernetes) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintKubernetesCreateKubernetes) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintKubernetesCreateKubernetes) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetYaml

`func (o *BlueprintKubernetesCreateKubernetes) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *BlueprintKubernetesCreateKubernetes) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *BlueprintKubernetesCreateKubernetes) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *BlueprintKubernetesCreateKubernetes) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *BlueprintKubernetesCreateKubernetes) GetGit() BlueprintKubernetesCreateKubernetesGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintKubernetesCreateKubernetes) GetGitOk() (*BlueprintKubernetesCreateKubernetesGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintKubernetesCreateKubernetes) SetGit(v BlueprintKubernetesCreateKubernetesGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintKubernetesCreateKubernetes) HasGit() bool`

HasGit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


