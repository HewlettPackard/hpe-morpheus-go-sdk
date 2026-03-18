# UpdateBlueprintRequestOneOf3Kubernetes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Yaml** | Pointer to **string** | Kubernetes Spec in YAML | [optional] 
**Git** | Pointer to [**UpdateBlueprintRequestOneOf3KubernetesGit**](UpdateBlueprintRequestOneOf3KubernetesGit.md) |  | [optional] 

## Methods

### NewUpdateBlueprintRequestOneOf3Kubernetes

`func NewUpdateBlueprintRequestOneOf3Kubernetes(configType string, ) *UpdateBlueprintRequestOneOf3Kubernetes`

NewUpdateBlueprintRequestOneOf3Kubernetes instantiates a new UpdateBlueprintRequestOneOf3Kubernetes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBlueprintRequestOneOf3KubernetesWithDefaults

`func NewUpdateBlueprintRequestOneOf3KubernetesWithDefaults() *UpdateBlueprintRequestOneOf3Kubernetes`

NewUpdateBlueprintRequestOneOf3KubernetesWithDefaults instantiates a new UpdateBlueprintRequestOneOf3Kubernetes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetYaml

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) SetYaml(v string)`

SetYaml sets Yaml field to given value.

### HasYaml

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) HasYaml() bool`

HasYaml returns a boolean if a field has been set.

### GetGit

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) GetGit() UpdateBlueprintRequestOneOf3KubernetesGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) GetGitOk() (*UpdateBlueprintRequestOneOf3KubernetesGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) SetGit(v UpdateBlueprintRequestOneOf3KubernetesGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *UpdateBlueprintRequestOneOf3Kubernetes) HasGit() bool`

HasGit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


