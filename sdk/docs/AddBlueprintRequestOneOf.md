# AddBlueprintRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Arm** | [**AddBlueprintRequestOneOfArm**](AddBlueprintRequestOneOfArm.md) |  | 

## Methods

### NewAddBlueprintRequestOneOf

`func NewAddBlueprintRequestOneOf(name string, type_ string, arm AddBlueprintRequestOneOfArm, ) *AddBlueprintRequestOneOf`

NewAddBlueprintRequestOneOf instantiates a new AddBlueprintRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprintRequestOneOfWithDefaults

`func NewAddBlueprintRequestOneOfWithDefaults() *AddBlueprintRequestOneOf`

NewAddBlueprintRequestOneOfWithDefaults instantiates a new AddBlueprintRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddBlueprintRequestOneOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBlueprintRequestOneOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBlueprintRequestOneOf) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *AddBlueprintRequestOneOf) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AddBlueprintRequestOneOf) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AddBlueprintRequestOneOf) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AddBlueprintRequestOneOf) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *AddBlueprintRequestOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddBlueprintRequestOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddBlueprintRequestOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *AddBlueprintRequestOneOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddBlueprintRequestOneOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddBlueprintRequestOneOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddBlueprintRequestOneOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetArm

`func (o *AddBlueprintRequestOneOf) GetArm() AddBlueprintRequestOneOfArm`

GetArm returns the Arm field if non-nil, zero value otherwise.

### GetArmOk

`func (o *AddBlueprintRequestOneOf) GetArmOk() (*AddBlueprintRequestOneOfArm, bool)`

GetArmOk returns a tuple with the Arm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArm

`func (o *AddBlueprintRequestOneOf) SetArm(v AddBlueprintRequestOneOfArm)`

SetArm sets Arm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


