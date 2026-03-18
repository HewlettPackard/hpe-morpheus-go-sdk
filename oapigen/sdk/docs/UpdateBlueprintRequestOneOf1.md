# UpdateBlueprintRequestOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**CloudFormation** | [**UpdateBlueprintRequestOneOf1CloudFormation**](UpdateBlueprintRequestOneOf1CloudFormation.md) |  | 

## Methods

### NewUpdateBlueprintRequestOneOf1

`func NewUpdateBlueprintRequestOneOf1(name string, type_ string, cloudFormation UpdateBlueprintRequestOneOf1CloudFormation, ) *UpdateBlueprintRequestOneOf1`

NewUpdateBlueprintRequestOneOf1 instantiates a new UpdateBlueprintRequestOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBlueprintRequestOneOf1WithDefaults

`func NewUpdateBlueprintRequestOneOf1WithDefaults() *UpdateBlueprintRequestOneOf1`

NewUpdateBlueprintRequestOneOf1WithDefaults instantiates a new UpdateBlueprintRequestOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBlueprintRequestOneOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBlueprintRequestOneOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBlueprintRequestOneOf1) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UpdateBlueprintRequestOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateBlueprintRequestOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateBlueprintRequestOneOf1) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *UpdateBlueprintRequestOneOf1) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateBlueprintRequestOneOf1) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateBlueprintRequestOneOf1) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateBlueprintRequestOneOf1) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateBlueprintRequestOneOf1) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateBlueprintRequestOneOf1) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetCloudFormation

`func (o *UpdateBlueprintRequestOneOf1) GetCloudFormation() UpdateBlueprintRequestOneOf1CloudFormation`

GetCloudFormation returns the CloudFormation field if non-nil, zero value otherwise.

### GetCloudFormationOk

`func (o *UpdateBlueprintRequestOneOf1) GetCloudFormationOk() (*UpdateBlueprintRequestOneOf1CloudFormation, bool)`

GetCloudFormationOk returns a tuple with the CloudFormation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormation

`func (o *UpdateBlueprintRequestOneOf1) SetCloudFormation(v UpdateBlueprintRequestOneOf1CloudFormation)`

SetCloudFormation sets CloudFormation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


