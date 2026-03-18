# BlueprintKubernetesCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Specs** | Pointer to [**[]BlueprintKubernetesCreateConfigSpecsInner**](BlueprintKubernetesCreateConfigSpecsInner.md) | Array of Kubernetes specs in Morpheus | [optional] 

## Methods

### NewBlueprintKubernetesCreateConfig

`func NewBlueprintKubernetesCreateConfig() *BlueprintKubernetesCreateConfig`

NewBlueprintKubernetesCreateConfig instantiates a new BlueprintKubernetesCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintKubernetesCreateConfigWithDefaults

`func NewBlueprintKubernetesCreateConfigWithDefaults() *BlueprintKubernetesCreateConfig`

NewBlueprintKubernetesCreateConfigWithDefaults instantiates a new BlueprintKubernetesCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecs

`func (o *BlueprintKubernetesCreateConfig) GetSpecs() []BlueprintKubernetesCreateConfigSpecsInner`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *BlueprintKubernetesCreateConfig) GetSpecsOk() (*[]BlueprintKubernetesCreateConfigSpecsInner, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *BlueprintKubernetesCreateConfig) SetSpecs(v []BlueprintKubernetesCreateConfigSpecsInner)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *BlueprintKubernetesCreateConfig) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


