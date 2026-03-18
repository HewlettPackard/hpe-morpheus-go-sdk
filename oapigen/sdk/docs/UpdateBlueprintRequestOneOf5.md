# UpdateBlueprintRequestOneOf5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Terraform** | [**UpdateBlueprintRequestOneOf5Terraform**](UpdateBlueprintRequestOneOf5Terraform.md) |  | 
**Config** | Pointer to [**UpdateBlueprintRequestOneOf5Config**](UpdateBlueprintRequestOneOf5Config.md) |  | [optional] 

## Methods

### NewUpdateBlueprintRequestOneOf5

`func NewUpdateBlueprintRequestOneOf5(name string, type_ string, terraform UpdateBlueprintRequestOneOf5Terraform, ) *UpdateBlueprintRequestOneOf5`

NewUpdateBlueprintRequestOneOf5 instantiates a new UpdateBlueprintRequestOneOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBlueprintRequestOneOf5WithDefaults

`func NewUpdateBlueprintRequestOneOf5WithDefaults() *UpdateBlueprintRequestOneOf5`

NewUpdateBlueprintRequestOneOf5WithDefaults instantiates a new UpdateBlueprintRequestOneOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBlueprintRequestOneOf5) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBlueprintRequestOneOf5) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBlueprintRequestOneOf5) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *UpdateBlueprintRequestOneOf5) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UpdateBlueprintRequestOneOf5) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UpdateBlueprintRequestOneOf5) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *UpdateBlueprintRequestOneOf5) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *UpdateBlueprintRequestOneOf5) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateBlueprintRequestOneOf5) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateBlueprintRequestOneOf5) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *UpdateBlueprintRequestOneOf5) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateBlueprintRequestOneOf5) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateBlueprintRequestOneOf5) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateBlueprintRequestOneOf5) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateBlueprintRequestOneOf5) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateBlueprintRequestOneOf5) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetTerraform

`func (o *UpdateBlueprintRequestOneOf5) GetTerraform() UpdateBlueprintRequestOneOf5Terraform`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *UpdateBlueprintRequestOneOf5) GetTerraformOk() (*UpdateBlueprintRequestOneOf5Terraform, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *UpdateBlueprintRequestOneOf5) SetTerraform(v UpdateBlueprintRequestOneOf5Terraform)`

SetTerraform sets Terraform field to given value.


### GetConfig

`func (o *UpdateBlueprintRequestOneOf5) GetConfig() UpdateBlueprintRequestOneOf5Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateBlueprintRequestOneOf5) GetConfigOk() (*UpdateBlueprintRequestOneOf5Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateBlueprintRequestOneOf5) SetConfig(v UpdateBlueprintRequestOneOf5Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateBlueprintRequestOneOf5) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


