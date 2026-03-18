# BlueprintTerraformCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Terraform** | [**BlueprintTerraformCreateTerraform**](BlueprintTerraformCreateTerraform.md) |  | 
**Config** | Pointer to [**BlueprintTerraformCreateConfig**](BlueprintTerraformCreateConfig.md) |  | [optional] 

## Methods

### NewBlueprintTerraformCreate

`func NewBlueprintTerraformCreate(name string, type_ string, terraform BlueprintTerraformCreateTerraform, ) *BlueprintTerraformCreate`

NewBlueprintTerraformCreate instantiates a new BlueprintTerraformCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintTerraformCreateWithDefaults

`func NewBlueprintTerraformCreateWithDefaults() *BlueprintTerraformCreate`

NewBlueprintTerraformCreateWithDefaults instantiates a new BlueprintTerraformCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintTerraformCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintTerraformCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintTerraformCreate) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *BlueprintTerraformCreate) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintTerraformCreate) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintTerraformCreate) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintTerraformCreate) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *BlueprintTerraformCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintTerraformCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintTerraformCreate) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *BlueprintTerraformCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BlueprintTerraformCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BlueprintTerraformCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BlueprintTerraformCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *BlueprintTerraformCreate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *BlueprintTerraformCreate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTerraform

`func (o *BlueprintTerraformCreate) GetTerraform() BlueprintTerraformCreateTerraform`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *BlueprintTerraformCreate) GetTerraformOk() (*BlueprintTerraformCreateTerraform, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *BlueprintTerraformCreate) SetTerraform(v BlueprintTerraformCreateTerraform)`

SetTerraform sets Terraform field to given value.


### GetConfig

`func (o *BlueprintTerraformCreate) GetConfig() BlueprintTerraformCreateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BlueprintTerraformCreate) GetConfigOk() (*BlueprintTerraformCreateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BlueprintTerraformCreate) SetConfig(v BlueprintTerraformCreateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BlueprintTerraformCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


