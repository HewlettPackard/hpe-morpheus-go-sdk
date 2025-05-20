# BlueprintCFTCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**CloudFormation** | [**AddBlueprintRequestOneOf1CloudFormation**](AddBlueprintRequestOneOf1CloudFormation.md) |  | 

## Methods

### NewBlueprintCFTCreate

`func NewBlueprintCFTCreate(name string, type_ string, cloudFormation AddBlueprintRequestOneOf1CloudFormation, ) *BlueprintCFTCreate`

NewBlueprintCFTCreate instantiates a new BlueprintCFTCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCFTCreateWithDefaults

`func NewBlueprintCFTCreateWithDefaults() *BlueprintCFTCreate`

NewBlueprintCFTCreateWithDefaults instantiates a new BlueprintCFTCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintCFTCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintCFTCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintCFTCreate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BlueprintCFTCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintCFTCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintCFTCreate) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *BlueprintCFTCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BlueprintCFTCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BlueprintCFTCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BlueprintCFTCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCloudFormation

`func (o *BlueprintCFTCreate) GetCloudFormation() AddBlueprintRequestOneOf1CloudFormation`

GetCloudFormation returns the CloudFormation field if non-nil, zero value otherwise.

### GetCloudFormationOk

`func (o *BlueprintCFTCreate) GetCloudFormationOk() (*AddBlueprintRequestOneOf1CloudFormation, bool)`

GetCloudFormationOk returns a tuple with the CloudFormation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormation

`func (o *BlueprintCFTCreate) SetCloudFormation(v AddBlueprintRequestOneOf1CloudFormation)`

SetCloudFormation sets CloudFormation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


