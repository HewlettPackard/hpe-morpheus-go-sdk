# UpdateBlueprintRequestOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Helm** | [**UpdateBlueprintRequestOneOf2Helm**](UpdateBlueprintRequestOneOf2Helm.md) |  | 

## Methods

### NewUpdateBlueprintRequestOneOf2

`func NewUpdateBlueprintRequestOneOf2(name string, type_ string, helm UpdateBlueprintRequestOneOf2Helm, ) *UpdateBlueprintRequestOneOf2`

NewUpdateBlueprintRequestOneOf2 instantiates a new UpdateBlueprintRequestOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateBlueprintRequestOneOf2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBlueprintRequestOneOf2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBlueprintRequestOneOf2) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *UpdateBlueprintRequestOneOf2) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UpdateBlueprintRequestOneOf2) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UpdateBlueprintRequestOneOf2) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *UpdateBlueprintRequestOneOf2) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *UpdateBlueprintRequestOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateBlueprintRequestOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateBlueprintRequestOneOf2) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *UpdateBlueprintRequestOneOf2) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateBlueprintRequestOneOf2) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateBlueprintRequestOneOf2) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateBlueprintRequestOneOf2) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateBlueprintRequestOneOf2) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateBlueprintRequestOneOf2) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetHelm

`func (o *UpdateBlueprintRequestOneOf2) GetHelm() UpdateBlueprintRequestOneOf2Helm`

GetHelm returns the Helm field if non-nil, zero value otherwise.

### GetHelmOk

`func (o *UpdateBlueprintRequestOneOf2) GetHelmOk() (*UpdateBlueprintRequestOneOf2Helm, bool)`

GetHelmOk returns a tuple with the Helm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelm

`func (o *UpdateBlueprintRequestOneOf2) SetHelm(v UpdateBlueprintRequestOneOf2Helm)`

SetHelm sets Helm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


