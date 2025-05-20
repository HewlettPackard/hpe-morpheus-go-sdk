# AddBlueprintRequestOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Helm** | [**AddBlueprintRequestOneOf2Helm**](AddBlueprintRequestOneOf2Helm.md) |  | 

## Methods

### NewAddBlueprintRequestOneOf2

`func NewAddBlueprintRequestOneOf2(name string, type_ string, helm AddBlueprintRequestOneOf2Helm, ) *AddBlueprintRequestOneOf2`

NewAddBlueprintRequestOneOf2 instantiates a new AddBlueprintRequestOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprintRequestOneOf2WithDefaults

`func NewAddBlueprintRequestOneOf2WithDefaults() *AddBlueprintRequestOneOf2`

NewAddBlueprintRequestOneOf2WithDefaults instantiates a new AddBlueprintRequestOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddBlueprintRequestOneOf2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBlueprintRequestOneOf2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBlueprintRequestOneOf2) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *AddBlueprintRequestOneOf2) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AddBlueprintRequestOneOf2) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AddBlueprintRequestOneOf2) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AddBlueprintRequestOneOf2) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *AddBlueprintRequestOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddBlueprintRequestOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddBlueprintRequestOneOf2) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *AddBlueprintRequestOneOf2) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddBlueprintRequestOneOf2) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddBlueprintRequestOneOf2) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddBlueprintRequestOneOf2) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetHelm

`func (o *AddBlueprintRequestOneOf2) GetHelm() AddBlueprintRequestOneOf2Helm`

GetHelm returns the Helm field if non-nil, zero value otherwise.

### GetHelmOk

`func (o *AddBlueprintRequestOneOf2) GetHelmOk() (*AddBlueprintRequestOneOf2Helm, bool)`

GetHelmOk returns a tuple with the Helm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelm

`func (o *AddBlueprintRequestOneOf2) SetHelm(v AddBlueprintRequestOneOf2Helm)`

SetHelm sets Helm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


