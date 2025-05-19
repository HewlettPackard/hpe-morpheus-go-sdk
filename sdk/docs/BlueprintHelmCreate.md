# BlueprintHelmCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Helm** | [**AddBlueprintRequestOneOf2Helm**](AddBlueprintRequestOneOf2Helm.md) |  | 

## Methods

### NewBlueprintHelmCreate

`func NewBlueprintHelmCreate(name string, type_ string, helm AddBlueprintRequestOneOf2Helm, ) *BlueprintHelmCreate`

NewBlueprintHelmCreate instantiates a new BlueprintHelmCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintHelmCreateWithDefaults

`func NewBlueprintHelmCreateWithDefaults() *BlueprintHelmCreate`

NewBlueprintHelmCreateWithDefaults instantiates a new BlueprintHelmCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintHelmCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintHelmCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintHelmCreate) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *BlueprintHelmCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintHelmCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintHelmCreate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintHelmCreate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *BlueprintHelmCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintHelmCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintHelmCreate) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *BlueprintHelmCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BlueprintHelmCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BlueprintHelmCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BlueprintHelmCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetHelm

`func (o *BlueprintHelmCreate) GetHelm() AddBlueprintRequestOneOf2Helm`

GetHelm returns the Helm field if non-nil, zero value otherwise.

### GetHelmOk

`func (o *BlueprintHelmCreate) GetHelmOk() (*AddBlueprintRequestOneOf2Helm, bool)`

GetHelmOk returns a tuple with the Helm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelm

`func (o *BlueprintHelmCreate) SetHelm(v AddBlueprintRequestOneOf2Helm)`

SetHelm sets Helm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


