# AddBlueprintRequestOneOf3Kubernetes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Yaml** | Pointer to **string** | Kubernetes Spec in YAML | [optional] 
**Git** | Pointer to [**AddBlueprintRequestOneOf3KubernetesGit**](AddBlueprintRequestOneOf3KubernetesGit.md) |  | [optional] 

## Methods

### NewAddBlueprintRequestOneOf3Kubernetes

`func NewAddBlueprintRequestOneOf3Kubernetes(configType string, ) *AddBlueprintRequestOneOf3Kubernetes`

NewAddBlueprintRequestOneOf3Kubernetes instantiates a new AddBlueprintRequestOneOf3Kubernetes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprintRequestOneOf3KubernetesWithDefaults

`func NewAddBlueprintRequestOneOf3KubernetesWithDefaults() *AddBlueprintRequestOneOf3Kubernetes`

NewAddBlueprintRequestOneOf3KubernetesWithDefaults instantiates a new AddBlueprintRequestOneOf3Kubernetes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *AddBlueprintRequestOneOf3Kubernetes) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *AddBlueprintRequestOneOf3Kubernetes) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *AddBlueprintRequestOneOf3Kubernetes) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetYaml

`func (o *AddBlueprintRequestOneOf3Kubernetes) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *AddBlueprintRequestOneOf3Kubernetes) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *AddBlueprintRequestOneOf3Kubernetes) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *AddBlueprintRequestOneOf3Kubernetes) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *AddBlueprintRequestOneOf3Kubernetes) GetGit() AddBlueprintRequestOneOf3KubernetesGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *AddBlueprintRequestOneOf3Kubernetes) GetGitOk() (*AddBlueprintRequestOneOf3KubernetesGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *AddBlueprintRequestOneOf3Kubernetes) SetGit(v AddBlueprintRequestOneOf3KubernetesGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *AddBlueprintRequestOneOf3Kubernetes) HasGit() bool`

HasGit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


