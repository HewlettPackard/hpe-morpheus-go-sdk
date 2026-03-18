# UpdateCatalogItemType200ResponseAllOfCatalogItemType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** | Useful shortcode for provisioning naming schemes and export reference. | [optional] 
**Category** | Pointer to **NullableString** | Catalog Item Type category | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**AllowQuantity** | Pointer to **bool** | Can users order more than one of this item at a time. | [optional] 
**IconPath** | Pointer to **string** |  | [optional] 
**ImagePath** | Pointer to **string** |  | [optional] 
**DarkImagePath** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**LayoutCode** | Pointer to **NullableString** |  | [optional] 
**Blueprint** | Pointer to **map[string]interface{}** |  | [optional] 
**AppSpec** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**InstanceSpec** | Pointer to **NullableString** |  | [optional] 
**Workflow** | Pointer to [**UpdateCatalogItemType200ResponseAllOfCatalogItemTypeWorkflow**](UpdateCatalogItemType200ResponseAllOfCatalogItemTypeWorkflow.md) |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**FormType** | Pointer to **string** |  | [optional] 
**Form** | Pointer to **map[string]interface{}** | Form object that contains input options and/or field groups | [optional] 
**FormConfig** | Pointer to **map[string]interface{}** | Form config object | [optional] 
**OptionTypes** | Pointer to [**[]UpdateCatalogItemType200ResponseAllOfCatalogItemTypeOptionTypesInner**](UpdateCatalogItemType200ResponseAllOfCatalogItemTypeOptionTypesInner.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**UpdateCatalogItemType200ResponseAllOfCatalogItemTypeOwner**](UpdateCatalogItemType200ResponseAllOfCatalogItemTypeOwner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateCatalogItemType200ResponseAllOfCatalogItemType

`func NewUpdateCatalogItemType200ResponseAllOfCatalogItemType() *UpdateCatalogItemType200ResponseAllOfCatalogItemType`

NewUpdateCatalogItemType200ResponseAllOfCatalogItemType instantiates a new UpdateCatalogItemType200ResponseAllOfCatalogItemType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCatalogItemType200ResponseAllOfCatalogItemTypeWithDefaults

`func NewUpdateCatalogItemType200ResponseAllOfCatalogItemTypeWithDefaults() *UpdateCatalogItemType200ResponseAllOfCatalogItemType`

NewUpdateCatalogItemType200ResponseAllOfCatalogItemTypeWithDefaults instantiates a new UpdateCatalogItemType200ResponseAllOfCatalogItemType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatured

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetIconPath

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetImagePath

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetDarkImagePath

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### SetLayoutCodeNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetLayoutCodeNil(b bool)`

 SetLayoutCodeNil sets the value for LayoutCode to be an explicit nil

### UnsetLayoutCode
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetLayoutCode()`

UnsetLayoutCode ensures that no value is present for LayoutCode, not even an explicit nil
### GetBlueprint

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetBlueprint() map[string]interface{}`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetBlueprintOk() (*map[string]interface{}, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetBlueprint(v map[string]interface{})`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### SetBlueprintNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetBlueprintNil(b bool)`

 SetBlueprintNil sets the value for Blueprint to be an explicit nil

### UnsetBlueprint
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetBlueprint()`

UnsetBlueprint ensures that no value is present for Blueprint, not even an explicit nil
### GetAppSpec

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetAppSpec() string`

GetAppSpec returns the AppSpec field if non-nil, zero value otherwise.

### GetAppSpecOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetAppSpecOk() (*string, bool)`

GetAppSpecOk returns a tuple with the AppSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSpec

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetAppSpec(v string)`

SetAppSpec sets AppSpec field to given value.

### HasAppSpec

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasAppSpec() bool`

HasAppSpec returns a boolean if a field has been set.

### SetAppSpecNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetAppSpecNil(b bool)`

 SetAppSpecNil sets the value for AppSpec to be an explicit nil

### UnsetAppSpec
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetAppSpec()`

UnsetAppSpec ensures that no value is present for AppSpec, not even an explicit nil
### GetConfig

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetInstanceSpec

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetInstanceSpec() string`

GetInstanceSpec returns the InstanceSpec field if non-nil, zero value otherwise.

### GetInstanceSpecOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetInstanceSpecOk() (*string, bool)`

GetInstanceSpecOk returns a tuple with the InstanceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSpec

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetInstanceSpec(v string)`

SetInstanceSpec sets InstanceSpec field to given value.

### HasInstanceSpec

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasInstanceSpec() bool`

HasInstanceSpec returns a boolean if a field has been set.

### SetInstanceSpecNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetInstanceSpecNil(b bool)`

 SetInstanceSpecNil sets the value for InstanceSpec to be an explicit nil

### UnsetInstanceSpec
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetInstanceSpec()`

UnsetInstanceSpec ensures that no value is present for InstanceSpec, not even an explicit nil
### GetWorkflow

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetWorkflow() UpdateCatalogItemType200ResponseAllOfCatalogItemTypeWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetWorkflowOk() (*UpdateCatalogItemType200ResponseAllOfCatalogItemTypeWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetWorkflow(v UpdateCatalogItemType200ResponseAllOfCatalogItemTypeWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetContent

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetFormType

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetForm() map[string]interface{}`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetFormOk() (*map[string]interface{}, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetForm(v map[string]interface{})`

SetForm sets Form field to given value.

### HasForm

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasForm() bool`

HasForm returns a boolean if a field has been set.

### SetFormNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetFormNil(b bool)`

 SetFormNil sets the value for Form to be an explicit nil

### UnsetForm
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetForm()`

UnsetForm ensures that no value is present for Form, not even an explicit nil
### GetFormConfig

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetFormConfig() map[string]interface{}`

GetFormConfig returns the FormConfig field if non-nil, zero value otherwise.

### GetFormConfigOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetFormConfigOk() (*map[string]interface{}, bool)`

GetFormConfigOk returns a tuple with the FormConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormConfig

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetFormConfig(v map[string]interface{})`

SetFormConfig sets FormConfig field to given value.

### HasFormConfig

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasFormConfig() bool`

HasFormConfig returns a boolean if a field has been set.

### SetFormConfigNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetFormConfigNil(b bool)`

 SetFormConfigNil sets the value for FormConfig to be an explicit nil

### UnsetFormConfig
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetFormConfig()`

UnsetFormConfig ensures that no value is present for FormConfig, not even an explicit nil
### GetOptionTypes

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetOptionTypes() []UpdateCatalogItemType200ResponseAllOfCatalogItemTypeOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetOptionTypesOk() (*[]UpdateCatalogItemType200ResponseAllOfCatalogItemTypeOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetOptionTypes(v []UpdateCatalogItemType200ResponseAllOfCatalogItemTypeOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetCreatedBy

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetOwner

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetOwner() UpdateCatalogItemType200ResponseAllOfCatalogItemTypeOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetOwnerOk() (*UpdateCatalogItemType200ResponseAllOfCatalogItemTypeOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetOwner(v UpdateCatalogItemType200ResponseAllOfCatalogItemTypeOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateCatalogItemType200ResponseAllOfCatalogItemType) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


