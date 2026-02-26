# BlueprintCatalogItemType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Catalog Item Type name | [optional] 
**Code** | Pointer to **NullableString** | Useful shortcode for provisioning naming schemes and export reference. | [optional] 
**Category** | Pointer to **NullableString** | Catalog Item Type category | [optional] 
**Description** | Pointer to **string** | Catalog Item Type description | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | Pointer to **string** | Type, &#x60;instance&#x60;, &#x60;blueprint&#x60; or &#x60;workflow&#x60;. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context. | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**LayoutCode** | Pointer to **NullableString** | Identifier primarily used for Plugin Catalog Item Types | [optional] 
**IconPath** | Pointer to **string** | Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png. | [optional] 
**AllowQuantity** | Pointer to **bool** | Can users order more than one of this item at a time. | [optional] 
**Blueprint** | [**BlueprintCatalogItemTypeBlueprint**](BlueprintCatalogItemTypeBlueprint.md) |  | 
**AppSpec** | Pointer to **string** | The appSpec for blueprint type catalog items is a string in the Scribe YAML format with fields | [optional] 
**FormType** | Pointer to **string** | Form Type determines if the configuration options come from a Form (form) or a list of Inputs (optionTypes). | [optional] [default to "optionTypes"]
**Form** | Pointer to [**BlueprintCatalogItemTypeForm**](BlueprintCatalogItemTypeForm.md) |  | [optional] 
**OptionTypes** | Pointer to **[]int64** | Array of option type IDs, see Inputs. Only applies to formType &#39;optionTypes&#39;. | [optional] 

## Methods

### NewBlueprintCatalogItemType

`func NewBlueprintCatalogItemType(blueprint BlueprintCatalogItemTypeBlueprint, ) *BlueprintCatalogItemType`

NewBlueprintCatalogItemType instantiates a new BlueprintCatalogItemType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCatalogItemTypeWithDefaults

`func NewBlueprintCatalogItemTypeWithDefaults() *BlueprintCatalogItemType`

NewBlueprintCatalogItemTypeWithDefaults instantiates a new BlueprintCatalogItemType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintCatalogItemType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintCatalogItemType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintCatalogItemType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintCatalogItemType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *BlueprintCatalogItemType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BlueprintCatalogItemType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BlueprintCatalogItemType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BlueprintCatalogItemType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *BlueprintCatalogItemType) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *BlueprintCatalogItemType) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *BlueprintCatalogItemType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BlueprintCatalogItemType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BlueprintCatalogItemType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BlueprintCatalogItemType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *BlueprintCatalogItemType) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *BlueprintCatalogItemType) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *BlueprintCatalogItemType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintCatalogItemType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintCatalogItemType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintCatalogItemType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *BlueprintCatalogItemType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BlueprintCatalogItemType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BlueprintCatalogItemType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BlueprintCatalogItemType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *BlueprintCatalogItemType) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *BlueprintCatalogItemType) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *BlueprintCatalogItemType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintCatalogItemType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintCatalogItemType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintCatalogItemType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *BlueprintCatalogItemType) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BlueprintCatalogItemType) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BlueprintCatalogItemType) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BlueprintCatalogItemType) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *BlueprintCatalogItemType) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *BlueprintCatalogItemType) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *BlueprintCatalogItemType) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *BlueprintCatalogItemType) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### SetLayoutCodeNil

`func (o *BlueprintCatalogItemType) SetLayoutCodeNil(b bool)`

 SetLayoutCodeNil sets the value for LayoutCode to be an explicit nil

### UnsetLayoutCode
`func (o *BlueprintCatalogItemType) UnsetLayoutCode()`

UnsetLayoutCode ensures that no value is present for LayoutCode, not even an explicit nil
### GetIconPath

`func (o *BlueprintCatalogItemType) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *BlueprintCatalogItemType) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *BlueprintCatalogItemType) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *BlueprintCatalogItemType) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *BlueprintCatalogItemType) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *BlueprintCatalogItemType) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *BlueprintCatalogItemType) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *BlueprintCatalogItemType) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetBlueprint

`func (o *BlueprintCatalogItemType) GetBlueprint() BlueprintCatalogItemTypeBlueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *BlueprintCatalogItemType) GetBlueprintOk() (*BlueprintCatalogItemTypeBlueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *BlueprintCatalogItemType) SetBlueprint(v BlueprintCatalogItemTypeBlueprint)`

SetBlueprint sets Blueprint field to given value.


### GetAppSpec

`func (o *BlueprintCatalogItemType) GetAppSpec() string`

GetAppSpec returns the AppSpec field if non-nil, zero value otherwise.

### GetAppSpecOk

`func (o *BlueprintCatalogItemType) GetAppSpecOk() (*string, bool)`

GetAppSpecOk returns a tuple with the AppSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSpec

`func (o *BlueprintCatalogItemType) SetAppSpec(v string)`

SetAppSpec sets AppSpec field to given value.

### HasAppSpec

`func (o *BlueprintCatalogItemType) HasAppSpec() bool`

HasAppSpec returns a boolean if a field has been set.

### GetFormType

`func (o *BlueprintCatalogItemType) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *BlueprintCatalogItemType) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *BlueprintCatalogItemType) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *BlueprintCatalogItemType) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *BlueprintCatalogItemType) GetForm() BlueprintCatalogItemTypeForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *BlueprintCatalogItemType) GetFormOk() (*BlueprintCatalogItemTypeForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *BlueprintCatalogItemType) SetForm(v BlueprintCatalogItemTypeForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *BlueprintCatalogItemType) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetOptionTypes

`func (o *BlueprintCatalogItemType) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *BlueprintCatalogItemType) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *BlueprintCatalogItemType) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *BlueprintCatalogItemType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


