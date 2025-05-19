# InstanceTypeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Instance type name | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Description** | Pointer to **string** | Instance type description | [optional] 
**Code** | Pointer to **string** | Instance type code | [optional] 
**Category** | Pointer to **string** | Category | [optional] 
**Visibility** | Pointer to **string** | Visibility | [optional] [default to "private"]
**Featured** | Pointer to **bool** | Featured | [optional] 
**HasSettings** | Pointer to **bool** | Enable Settings | [optional] 
**HasAutoScale** | Pointer to **bool** | Enable Scaling (Horizontal) | [optional] 
**HasDeployment** | Pointer to **bool** | Supports Deployments | [optional] 
**EnvironmentPrefix** | Pointer to **string** | Environment Prefix, can be used to make exported evars unique. | [optional] 
**EnvironmentVariables** | Pointer to [**[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner**](AddClusterLayoutsRequestLayoutEnvironmentVariablesInner.md) | Array of instance type env variables. | [optional] 
**PriceSets** | Pointer to [**[]AddInstanceTypeRequestInstanceTypePriceSetsInner**](AddInstanceTypeRequestInstanceTypePriceSetsInner.md) | Array of price set objects | [optional] 
**OptionTypes** | Pointer to **[]int64** | Array of instance type option type IDs | [optional] 

## Methods

### NewInstanceTypeCreate

`func NewInstanceTypeCreate(name string, ) *InstanceTypeCreate`

NewInstanceTypeCreate instantiates a new InstanceTypeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeCreateWithDefaults

`func NewInstanceTypeCreateWithDefaults() *InstanceTypeCreate`

NewInstanceTypeCreateWithDefaults instantiates a new InstanceTypeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceTypeCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeCreate) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *InstanceTypeCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceTypeCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceTypeCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceTypeCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceTypeCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceTypeCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceTypeCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceTypeCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *InstanceTypeCreate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceTypeCreate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceTypeCreate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceTypeCreate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *InstanceTypeCreate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InstanceTypeCreate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InstanceTypeCreate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InstanceTypeCreate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVisibility

`func (o *InstanceTypeCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InstanceTypeCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InstanceTypeCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InstanceTypeCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetFeatured

`func (o *InstanceTypeCreate) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *InstanceTypeCreate) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *InstanceTypeCreate) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *InstanceTypeCreate) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetHasSettings

`func (o *InstanceTypeCreate) GetHasSettings() bool`

GetHasSettings returns the HasSettings field if non-nil, zero value otherwise.

### GetHasSettingsOk

`func (o *InstanceTypeCreate) GetHasSettingsOk() (*bool, bool)`

GetHasSettingsOk returns a tuple with the HasSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSettings

`func (o *InstanceTypeCreate) SetHasSettings(v bool)`

SetHasSettings sets HasSettings field to given value.

### HasHasSettings

`func (o *InstanceTypeCreate) HasHasSettings() bool`

HasHasSettings returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *InstanceTypeCreate) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *InstanceTypeCreate) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *InstanceTypeCreate) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *InstanceTypeCreate) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetHasDeployment

`func (o *InstanceTypeCreate) GetHasDeployment() bool`

GetHasDeployment returns the HasDeployment field if non-nil, zero value otherwise.

### GetHasDeploymentOk

`func (o *InstanceTypeCreate) GetHasDeploymentOk() (*bool, bool)`

GetHasDeploymentOk returns a tuple with the HasDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeployment

`func (o *InstanceTypeCreate) SetHasDeployment(v bool)`

SetHasDeployment sets HasDeployment field to given value.

### HasHasDeployment

`func (o *InstanceTypeCreate) HasHasDeployment() bool`

HasHasDeployment returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *InstanceTypeCreate) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *InstanceTypeCreate) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *InstanceTypeCreate) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *InstanceTypeCreate) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *InstanceTypeCreate) GetEnvironmentVariables() []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *InstanceTypeCreate) GetEnvironmentVariablesOk() (*[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *InstanceTypeCreate) SetEnvironmentVariables(v []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *InstanceTypeCreate) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetPriceSets

`func (o *InstanceTypeCreate) GetPriceSets() []AddInstanceTypeRequestInstanceTypePriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *InstanceTypeCreate) GetPriceSetsOk() (*[]AddInstanceTypeRequestInstanceTypePriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *InstanceTypeCreate) SetPriceSets(v []AddInstanceTypeRequestInstanceTypePriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *InstanceTypeCreate) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InstanceTypeCreate) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InstanceTypeCreate) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InstanceTypeCreate) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InstanceTypeCreate) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


