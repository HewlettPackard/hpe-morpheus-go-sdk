# AddBlueprintRequestOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**CloudFormation** | [**AddBlueprintRequestOneOf1CloudFormation**](AddBlueprintRequestOneOf1CloudFormation.md) |  | 

## Methods

### NewAddBlueprintRequestOneOf1

`func NewAddBlueprintRequestOneOf1(name string, type_ string, cloudFormation AddBlueprintRequestOneOf1CloudFormation, ) *AddBlueprintRequestOneOf1`

NewAddBlueprintRequestOneOf1 instantiates a new AddBlueprintRequestOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprintRequestOneOf1WithDefaults

`func NewAddBlueprintRequestOneOf1WithDefaults() *AddBlueprintRequestOneOf1`

NewAddBlueprintRequestOneOf1WithDefaults instantiates a new AddBlueprintRequestOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddBlueprintRequestOneOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBlueprintRequestOneOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBlueprintRequestOneOf1) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddBlueprintRequestOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddBlueprintRequestOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddBlueprintRequestOneOf1) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *AddBlueprintRequestOneOf1) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddBlueprintRequestOneOf1) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddBlueprintRequestOneOf1) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddBlueprintRequestOneOf1) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCloudFormation

`func (o *AddBlueprintRequestOneOf1) GetCloudFormation() AddBlueprintRequestOneOf1CloudFormation`

GetCloudFormation returns the CloudFormation field if non-nil, zero value otherwise.

### GetCloudFormationOk

`func (o *AddBlueprintRequestOneOf1) GetCloudFormationOk() (*AddBlueprintRequestOneOf1CloudFormation, bool)`

GetCloudFormationOk returns a tuple with the CloudFormation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormation

`func (o *AddBlueprintRequestOneOf1) SetCloudFormation(v AddBlueprintRequestOneOf1CloudFormation)`

SetCloudFormation sets CloudFormation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


