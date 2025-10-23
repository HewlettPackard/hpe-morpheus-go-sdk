# GetCatalogItemType200ResponseCatalogItemType

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
**Workflow** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 
**FormType** | Pointer to **string** |  | [optional] 
**Form** | Pointer to **map[string]interface{}** | Form object that contains input options and/or field groups | [optional] 
**FormConfig** | Pointer to **map[string]interface{}** | Form config object | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetCatalogItemType200ResponseCatalogItemType

`func NewGetCatalogItemType200ResponseCatalogItemType() *GetCatalogItemType200ResponseCatalogItemType`

NewGetCatalogItemType200ResponseCatalogItemType instantiates a new GetCatalogItemType200ResponseCatalogItemType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogItemType200ResponseCatalogItemTypeWithDefaults

`func NewGetCatalogItemType200ResponseCatalogItemTypeWithDefaults() *GetCatalogItemType200ResponseCatalogItemType`

NewGetCatalogItemType200ResponseCatalogItemTypeWithDefaults instantiates a new GetCatalogItemType200ResponseCatalogItemType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatured

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetIconPath

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetImagePath

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetDarkImagePath

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### GetVisibility

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### SetLayoutCodeNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetLayoutCodeNil(b bool)`

 SetLayoutCodeNil sets the value for LayoutCode to be an explicit nil

### UnsetLayoutCode
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetLayoutCode()`

UnsetLayoutCode ensures that no value is present for LayoutCode, not even an explicit nil
### GetBlueprint

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetBlueprint() map[string]interface{}`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetBlueprintOk() (*map[string]interface{}, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetBlueprint(v map[string]interface{})`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### SetBlueprintNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetBlueprintNil(b bool)`

 SetBlueprintNil sets the value for Blueprint to be an explicit nil

### UnsetBlueprint
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetBlueprint()`

UnsetBlueprint ensures that no value is present for Blueprint, not even an explicit nil
### GetAppSpec

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetAppSpec() string`

GetAppSpec returns the AppSpec field if non-nil, zero value otherwise.

### GetAppSpecOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetAppSpecOk() (*string, bool)`

GetAppSpecOk returns a tuple with the AppSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSpec

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetAppSpec(v string)`

SetAppSpec sets AppSpec field to given value.

### HasAppSpec

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasAppSpec() bool`

HasAppSpec returns a boolean if a field has been set.

### SetAppSpecNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetAppSpecNil(b bool)`

 SetAppSpecNil sets the value for AppSpec to be an explicit nil

### UnsetAppSpec
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetAppSpec()`

UnsetAppSpec ensures that no value is present for AppSpec, not even an explicit nil
### GetConfig

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetInstanceSpec

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetInstanceSpec() string`

GetInstanceSpec returns the InstanceSpec field if non-nil, zero value otherwise.

### GetInstanceSpecOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetInstanceSpecOk() (*string, bool)`

GetInstanceSpecOk returns a tuple with the InstanceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSpec

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetInstanceSpec(v string)`

SetInstanceSpec sets InstanceSpec field to given value.

### HasInstanceSpec

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasInstanceSpec() bool`

HasInstanceSpec returns a boolean if a field has been set.

### SetInstanceSpecNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetInstanceSpecNil(b bool)`

 SetInstanceSpecNil sets the value for InstanceSpec to be an explicit nil

### UnsetInstanceSpec
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetInstanceSpec()`

UnsetInstanceSpec ensures that no value is present for InstanceSpec, not even an explicit nil
### GetWorkflow

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetWorkflow() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetWorkflowOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetWorkflow(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetContent

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetFormType

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetForm() map[string]interface{}`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetFormOk() (*map[string]interface{}, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetForm(v map[string]interface{})`

SetForm sets Form field to given value.

### HasForm

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasForm() bool`

HasForm returns a boolean if a field has been set.

### SetFormNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetFormNil(b bool)`

 SetFormNil sets the value for Form to be an explicit nil

### UnsetForm
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetForm()`

UnsetForm ensures that no value is present for Form, not even an explicit nil
### GetFormConfig

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetFormConfig() map[string]interface{}`

GetFormConfig returns the FormConfig field if non-nil, zero value otherwise.

### GetFormConfigOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetFormConfigOk() (*map[string]interface{}, bool)`

GetFormConfigOk returns a tuple with the FormConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormConfig

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetFormConfig(v map[string]interface{})`

SetFormConfig sets FormConfig field to given value.

### HasFormConfig

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasFormConfig() bool`

HasFormConfig returns a boolean if a field has been set.

### SetFormConfigNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetFormConfigNil(b bool)`

 SetFormConfigNil sets the value for FormConfig to be an explicit nil

### UnsetFormConfig
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetFormConfig()`

UnsetFormConfig ensures that no value is present for FormConfig, not even an explicit nil
### GetOptionTypes

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetCreatedBy

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GetCatalogItemType200ResponseCatalogItemType) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetOwner

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetCatalogItemType200ResponseCatalogItemType) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetCatalogItemType200ResponseCatalogItemType) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetCatalogItemType200ResponseCatalogItemType) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


