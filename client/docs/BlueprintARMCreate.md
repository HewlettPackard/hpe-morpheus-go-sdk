# BlueprintARMCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Arm** | [**AddBlueprintRequestOneOfArm**](AddBlueprintRequestOneOfArm.md) |  | 

## Methods

### NewBlueprintARMCreate

`func NewBlueprintARMCreate(name string, type_ string, arm AddBlueprintRequestOneOfArm, ) *BlueprintARMCreate`

NewBlueprintARMCreate instantiates a new BlueprintARMCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintARMCreateWithDefaults

`func NewBlueprintARMCreateWithDefaults() *BlueprintARMCreate`

NewBlueprintARMCreateWithDefaults instantiates a new BlueprintARMCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintARMCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintARMCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintARMCreate) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *BlueprintARMCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintARMCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintARMCreate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintARMCreate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *BlueprintARMCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintARMCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintARMCreate) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *BlueprintARMCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BlueprintARMCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BlueprintARMCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BlueprintARMCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetArm

`func (o *BlueprintARMCreate) GetArm() AddBlueprintRequestOneOfArm`

GetArm returns the Arm field if non-nil, zero value otherwise.

### GetArmOk

`func (o *BlueprintARMCreate) GetArmOk() (*AddBlueprintRequestOneOfArm, bool)`

GetArmOk returns a tuple with the Arm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArm

`func (o *BlueprintARMCreate) SetArm(v AddBlueprintRequestOneOfArm)`

SetArm sets Arm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


