# BlueprintKubernetesCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Kubernetes** | [**AddBlueprintRequestOneOf3Kubernetes**](AddBlueprintRequestOneOf3Kubernetes.md) |  | 
**Config** | Pointer to [**AddBlueprintRequestOneOf3Config**](AddBlueprintRequestOneOf3Config.md) |  | [optional] 

## Methods

### NewBlueprintKubernetesCreate

`func NewBlueprintKubernetesCreate(name string, type_ string, kubernetes AddBlueprintRequestOneOf3Kubernetes, ) *BlueprintKubernetesCreate`

NewBlueprintKubernetesCreate instantiates a new BlueprintKubernetesCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintKubernetesCreateWithDefaults

`func NewBlueprintKubernetesCreateWithDefaults() *BlueprintKubernetesCreate`

NewBlueprintKubernetesCreateWithDefaults instantiates a new BlueprintKubernetesCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintKubernetesCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintKubernetesCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintKubernetesCreate) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *BlueprintKubernetesCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintKubernetesCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintKubernetesCreate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintKubernetesCreate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *BlueprintKubernetesCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintKubernetesCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintKubernetesCreate) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *BlueprintKubernetesCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BlueprintKubernetesCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BlueprintKubernetesCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BlueprintKubernetesCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetKubernetes

`func (o *BlueprintKubernetesCreate) GetKubernetes() AddBlueprintRequestOneOf3Kubernetes`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *BlueprintKubernetesCreate) GetKubernetesOk() (*AddBlueprintRequestOneOf3Kubernetes, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *BlueprintKubernetesCreate) SetKubernetes(v AddBlueprintRequestOneOf3Kubernetes)`

SetKubernetes sets Kubernetes field to given value.


### GetConfig

`func (o *BlueprintKubernetesCreate) GetConfig() AddBlueprintRequestOneOf3Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BlueprintKubernetesCreate) GetConfigOk() (*AddBlueprintRequestOneOf3Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BlueprintKubernetesCreate) SetConfig(v AddBlueprintRequestOneOf3Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BlueprintKubernetesCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


