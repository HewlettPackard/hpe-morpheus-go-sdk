# InstanceCatalogItemType

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
**Enabled** | Pointer to **bool** | Can be used to enable / disable the catalog item type. | [optional] [default to true]
**Featured** | Pointer to **bool** | Can be used to feature the catalog item type. | [optional] [default to false]
**AllowQuantity** | Pointer to **bool** | Can users order more than one of this item at a time. | [optional] [default to false]
**Config** | [**InstanceCatalogItemTypeConfig**](InstanceCatalogItemTypeConfig.md) |  | 
**InstanceSpec** | Pointer to **string** | The instance &#x60;config&#x60; specification as a string in the JSON format. | [optional] 
**FormType** | Pointer to **string** | Form Type determines if the configuration options come from a Form (form) or a list of Inputs (optionTypes). | [optional] [default to "optionTypes"]
**Form** | Pointer to [**InstanceCatalogItemTypeForm**](InstanceCatalogItemTypeForm.md) |  | [optional] 
**OptionTypes** | Pointer to **[]int64** | Array of option type IDs. Only applies to formType &#39;optionTypes&#39;. | [optional] 
**Content** | Pointer to **string** | Documentation content for this Catalog Item. Markdown-formatted text is accepted and displayed appropriately when the item is ordered from the Service Catalog. A new Catalog Item-type Wiki entry will also be added containing this information. | [optional] 

## Methods

### NewInstanceCatalogItemType

`func NewInstanceCatalogItemType(config InstanceCatalogItemTypeConfig, ) *InstanceCatalogItemType`

NewInstanceCatalogItemType instantiates a new InstanceCatalogItemType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCatalogItemTypeWithDefaults

`func NewInstanceCatalogItemTypeWithDefaults() *InstanceCatalogItemType`

NewInstanceCatalogItemTypeWithDefaults instantiates a new InstanceCatalogItemType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceCatalogItemType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceCatalogItemType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceCatalogItemType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceCatalogItemType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *InstanceCatalogItemType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceCatalogItemType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceCatalogItemType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceCatalogItemType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *InstanceCatalogItemType) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *InstanceCatalogItemType) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *InstanceCatalogItemType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InstanceCatalogItemType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InstanceCatalogItemType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InstanceCatalogItemType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *InstanceCatalogItemType) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *InstanceCatalogItemType) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *InstanceCatalogItemType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceCatalogItemType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceCatalogItemType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceCatalogItemType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceCatalogItemType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceCatalogItemType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceCatalogItemType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceCatalogItemType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *InstanceCatalogItemType) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *InstanceCatalogItemType) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *InstanceCatalogItemType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceCatalogItemType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceCatalogItemType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstanceCatalogItemType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *InstanceCatalogItemType) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InstanceCatalogItemType) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InstanceCatalogItemType) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InstanceCatalogItemType) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *InstanceCatalogItemType) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *InstanceCatalogItemType) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *InstanceCatalogItemType) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *InstanceCatalogItemType) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### SetLayoutCodeNil

`func (o *InstanceCatalogItemType) SetLayoutCodeNil(b bool)`

 SetLayoutCodeNil sets the value for LayoutCode to be an explicit nil

### UnsetLayoutCode
`func (o *InstanceCatalogItemType) UnsetLayoutCode()`

UnsetLayoutCode ensures that no value is present for LayoutCode, not even an explicit nil
### GetIconPath

`func (o *InstanceCatalogItemType) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *InstanceCatalogItemType) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *InstanceCatalogItemType) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *InstanceCatalogItemType) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetEnabled

`func (o *InstanceCatalogItemType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InstanceCatalogItemType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InstanceCatalogItemType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InstanceCatalogItemType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatured

`func (o *InstanceCatalogItemType) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *InstanceCatalogItemType) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *InstanceCatalogItemType) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *InstanceCatalogItemType) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *InstanceCatalogItemType) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *InstanceCatalogItemType) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *InstanceCatalogItemType) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *InstanceCatalogItemType) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetConfig

`func (o *InstanceCatalogItemType) GetConfig() InstanceCatalogItemTypeConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InstanceCatalogItemType) GetConfigOk() (*InstanceCatalogItemTypeConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InstanceCatalogItemType) SetConfig(v InstanceCatalogItemTypeConfig)`

SetConfig sets Config field to given value.


### GetInstanceSpec

`func (o *InstanceCatalogItemType) GetInstanceSpec() string`

GetInstanceSpec returns the InstanceSpec field if non-nil, zero value otherwise.

### GetInstanceSpecOk

`func (o *InstanceCatalogItemType) GetInstanceSpecOk() (*string, bool)`

GetInstanceSpecOk returns a tuple with the InstanceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSpec

`func (o *InstanceCatalogItemType) SetInstanceSpec(v string)`

SetInstanceSpec sets InstanceSpec field to given value.

### HasInstanceSpec

`func (o *InstanceCatalogItemType) HasInstanceSpec() bool`

HasInstanceSpec returns a boolean if a field has been set.

### GetFormType

`func (o *InstanceCatalogItemType) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *InstanceCatalogItemType) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *InstanceCatalogItemType) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *InstanceCatalogItemType) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *InstanceCatalogItemType) GetForm() InstanceCatalogItemTypeForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *InstanceCatalogItemType) GetFormOk() (*InstanceCatalogItemTypeForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *InstanceCatalogItemType) SetForm(v InstanceCatalogItemTypeForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *InstanceCatalogItemType) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InstanceCatalogItemType) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InstanceCatalogItemType) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InstanceCatalogItemType) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InstanceCatalogItemType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetContent

`func (o *InstanceCatalogItemType) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *InstanceCatalogItemType) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *InstanceCatalogItemType) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *InstanceCatalogItemType) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


