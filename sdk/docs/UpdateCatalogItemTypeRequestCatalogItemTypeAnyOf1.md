# UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Catalog Item Type name | [optional] 
**Code** | Pointer to **string** | Useful shortcode for provisioning naming schemes and export reference. | [optional] 
**Category** | Pointer to **string** | Catalog Item Type category | [optional] 
**Description** | Pointer to **string** | Catalog Item Type description | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | Pointer to **string** | Type, &#x60;instance&#x60;, &#x60;blueprint&#x60; or &#x60;workflow&#x60;. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context. | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**LayoutCode** | Pointer to **string** | Identifier primarily used for Plugin Catalog Item Types | [optional] 
**IconPath** | Pointer to **string** | Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png. | [optional] 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the catalog item type. | [optional] [default to true]
**Featured** | Pointer to **bool** | Can be used to feature the catalog item type. | [optional] [default to false]
**AllowQuantity** | Pointer to **bool** | Can users order more than one of this item at a time. | [optional] [default to false]
**Blueprint** | Pointer to [**AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint**](AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint.md) |  | [optional] 
**AppSpec** | Pointer to **string** | The appSpec for blueprint type catalog items is a string in the Scribe YAML format with fields | [optional] 
**FormType** | Pointer to **string** | Form Type determines if the configuration options come from a Form (form) or a list of Inputs (optionTypes). | [optional] [default to "optionTypes"]
**Form** | Pointer to [**AddCatalogItemTypeRequestCatalogItemTypeOneOfForm**](AddCatalogItemTypeRequestCatalogItemTypeOneOfForm.md) |  | [optional] 
**OptionTypes** | Pointer to **[]int64** | Array of option type IDs, see Inputs. Only applies to formType &#39;optionTypes&#39;. | [optional] 

## Methods

### NewUpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1

`func NewUpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1() *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1`

NewUpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1 instantiates a new UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1WithDefaults

`func NewUpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1WithDefaults() *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1`

NewUpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1WithDefaults instantiates a new UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### GetIconPath

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatured

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetBlueprint

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetBlueprint() AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetBlueprintOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetBlueprint(v AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetAppSpec

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetAppSpec() string`

GetAppSpec returns the AppSpec field if non-nil, zero value otherwise.

### GetAppSpecOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetAppSpecOk() (*string, bool)`

GetAppSpecOk returns a tuple with the AppSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSpec

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetAppSpec(v string)`

SetAppSpec sets AppSpec field to given value.

### HasAppSpec

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasAppSpec() bool`

HasAppSpec returns a boolean if a field has been set.

### GetFormType

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetForm() AddCatalogItemTypeRequestCatalogItemTypeOneOfForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetFormOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetForm(v AddCatalogItemTypeRequestCatalogItemTypeOneOfForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetOptionTypes

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *UpdateCatalogItemTypeRequestCatalogItemTypeAnyOf1) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


