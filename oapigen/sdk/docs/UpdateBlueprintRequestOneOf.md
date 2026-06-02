# UpdateBlueprintRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Arm** | [**UpdateBlueprintRequestOneOfArm**](UpdateBlueprintRequestOneOfArm.md) |  | 

## Methods

### NewUpdateBlueprintRequestOneOf

`func NewUpdateBlueprintRequestOneOf(name string, type_ string, arm UpdateBlueprintRequestOneOfArm, ) *UpdateBlueprintRequestOneOf`

NewUpdateBlueprintRequestOneOf instantiates a new UpdateBlueprintRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateBlueprintRequestOneOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBlueprintRequestOneOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBlueprintRequestOneOf) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *UpdateBlueprintRequestOneOf) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UpdateBlueprintRequestOneOf) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UpdateBlueprintRequestOneOf) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *UpdateBlueprintRequestOneOf) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *UpdateBlueprintRequestOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateBlueprintRequestOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateBlueprintRequestOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *UpdateBlueprintRequestOneOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateBlueprintRequestOneOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateBlueprintRequestOneOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateBlueprintRequestOneOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateBlueprintRequestOneOf) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateBlueprintRequestOneOf) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetArm

`func (o *UpdateBlueprintRequestOneOf) GetArm() UpdateBlueprintRequestOneOfArm`

GetArm returns the Arm field if non-nil, zero value otherwise.

### GetArmOk

`func (o *UpdateBlueprintRequestOneOf) GetArmOk() (*UpdateBlueprintRequestOneOfArm, bool)`

GetArmOk returns a tuple with the Arm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArm

`func (o *UpdateBlueprintRequestOneOf) SetArm(v UpdateBlueprintRequestOneOfArm)`

SetArm sets Arm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


