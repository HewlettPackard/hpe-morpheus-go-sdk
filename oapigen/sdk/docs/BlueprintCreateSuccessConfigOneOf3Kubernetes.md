# BlueprintCreateSuccessConfigOneOf3Kubernetes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Yaml** | Pointer to **string** | Kubernetes Spec in YAML | [optional] 
**Git** | Pointer to [**BlueprintCreateSuccessConfigOneOf3KubernetesGit**](BlueprintCreateSuccessConfigOneOf3KubernetesGit.md) |  | [optional] 

## Methods

### NewBlueprintCreateSuccessConfigOneOf3Kubernetes

`func NewBlueprintCreateSuccessConfigOneOf3Kubernetes(configType string, ) *BlueprintCreateSuccessConfigOneOf3Kubernetes`

NewBlueprintCreateSuccessConfigOneOf3Kubernetes instantiates a new BlueprintCreateSuccessConfigOneOf3Kubernetes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetConfigType

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetYaml

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) GetGit() BlueprintCreateSuccessConfigOneOf3KubernetesGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) GetGitOk() (*BlueprintCreateSuccessConfigOneOf3KubernetesGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) SetGit(v BlueprintCreateSuccessConfigOneOf3KubernetesGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintCreateSuccessConfigOneOf3Kubernetes) HasGit() bool`

HasGit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


