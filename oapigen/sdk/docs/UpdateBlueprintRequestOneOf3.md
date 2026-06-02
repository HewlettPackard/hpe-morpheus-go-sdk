# UpdateBlueprintRequestOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Kubernetes** | [**UpdateBlueprintRequestOneOf3Kubernetes**](UpdateBlueprintRequestOneOf3Kubernetes.md) |  | 
**Config** | Pointer to [**UpdateBlueprintRequestOneOf3Config**](UpdateBlueprintRequestOneOf3Config.md) |  | [optional] 

## Methods

### NewUpdateBlueprintRequestOneOf3

`func NewUpdateBlueprintRequestOneOf3(name string, type_ string, kubernetes UpdateBlueprintRequestOneOf3Kubernetes, ) *UpdateBlueprintRequestOneOf3`

NewUpdateBlueprintRequestOneOf3 instantiates a new UpdateBlueprintRequestOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateBlueprintRequestOneOf3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBlueprintRequestOneOf3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBlueprintRequestOneOf3) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *UpdateBlueprintRequestOneOf3) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UpdateBlueprintRequestOneOf3) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UpdateBlueprintRequestOneOf3) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *UpdateBlueprintRequestOneOf3) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *UpdateBlueprintRequestOneOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateBlueprintRequestOneOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateBlueprintRequestOneOf3) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *UpdateBlueprintRequestOneOf3) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateBlueprintRequestOneOf3) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateBlueprintRequestOneOf3) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateBlueprintRequestOneOf3) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateBlueprintRequestOneOf3) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateBlueprintRequestOneOf3) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetKubernetes

`func (o *UpdateBlueprintRequestOneOf3) GetKubernetes() UpdateBlueprintRequestOneOf3Kubernetes`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *UpdateBlueprintRequestOneOf3) GetKubernetesOk() (*UpdateBlueprintRequestOneOf3Kubernetes, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *UpdateBlueprintRequestOneOf3) SetKubernetes(v UpdateBlueprintRequestOneOf3Kubernetes)`

SetKubernetes sets Kubernetes field to given value.


### GetConfig

`func (o *UpdateBlueprintRequestOneOf3) GetConfig() UpdateBlueprintRequestOneOf3Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateBlueprintRequestOneOf3) GetConfigOk() (*UpdateBlueprintRequestOneOf3Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateBlueprintRequestOneOf3) SetConfig(v UpdateBlueprintRequestOneOf3Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateBlueprintRequestOneOf3) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


