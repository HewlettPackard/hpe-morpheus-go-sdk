# WorkflowCatalogItemType

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
**Workflow** | [**WorkflowCatalogItemTypeWorkflow**](WorkflowCatalogItemTypeWorkflow.md) |  | 
**Context** | Pointer to **string** | Context for running the workflow, determines if a target resource must be selected. | [optional] 
**WorkflowConfig** | Pointer to **string** | Configuration object that contains settings for the workflow. | [optional] 
**FormType** | Pointer to **string** | Form Type determines if the configuration options come from a Form (form) or a list of Inputs (optionTypes). | [optional] [default to "optionTypes"]
**Form** | Pointer to [**WorkflowCatalogItemTypeForm**](WorkflowCatalogItemTypeForm.md) |  | [optional] 
**OptionTypes** | Pointer to **[]int64** | Array of option type IDs. Only applies to formType &#39;optionTypes&#39;. | [optional] 
**Content** | Pointer to **string** | Documentation content for this Catalog Item. Markdown-formatted text is accepted and displayed appropriately when the item is ordered from the Service Catalog. A new Catalog Item-type Wiki entry will also be added containing this information. | [optional] 

## Methods

### NewWorkflowCatalogItemType

`func NewWorkflowCatalogItemType(workflow WorkflowCatalogItemTypeWorkflow, ) *WorkflowCatalogItemType`

NewWorkflowCatalogItemType instantiates a new WorkflowCatalogItemType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCatalogItemTypeWithDefaults

`func NewWorkflowCatalogItemTypeWithDefaults() *WorkflowCatalogItemType`

NewWorkflowCatalogItemTypeWithDefaults instantiates a new WorkflowCatalogItemType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WorkflowCatalogItemType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowCatalogItemType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowCatalogItemType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowCatalogItemType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *WorkflowCatalogItemType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WorkflowCatalogItemType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WorkflowCatalogItemType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *WorkflowCatalogItemType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *WorkflowCatalogItemType) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *WorkflowCatalogItemType) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *WorkflowCatalogItemType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *WorkflowCatalogItemType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *WorkflowCatalogItemType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *WorkflowCatalogItemType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *WorkflowCatalogItemType) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *WorkflowCatalogItemType) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDescription

`func (o *WorkflowCatalogItemType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowCatalogItemType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowCatalogItemType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowCatalogItemType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *WorkflowCatalogItemType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WorkflowCatalogItemType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WorkflowCatalogItemType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WorkflowCatalogItemType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WorkflowCatalogItemType) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WorkflowCatalogItemType) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetType

`func (o *WorkflowCatalogItemType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowCatalogItemType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowCatalogItemType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowCatalogItemType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *WorkflowCatalogItemType) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *WorkflowCatalogItemType) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *WorkflowCatalogItemType) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *WorkflowCatalogItemType) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *WorkflowCatalogItemType) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *WorkflowCatalogItemType) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *WorkflowCatalogItemType) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *WorkflowCatalogItemType) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### SetLayoutCodeNil

`func (o *WorkflowCatalogItemType) SetLayoutCodeNil(b bool)`

 SetLayoutCodeNil sets the value for LayoutCode to be an explicit nil

### UnsetLayoutCode
`func (o *WorkflowCatalogItemType) UnsetLayoutCode()`

UnsetLayoutCode ensures that no value is present for LayoutCode, not even an explicit nil
### GetIconPath

`func (o *WorkflowCatalogItemType) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *WorkflowCatalogItemType) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *WorkflowCatalogItemType) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *WorkflowCatalogItemType) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetEnabled

`func (o *WorkflowCatalogItemType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowCatalogItemType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowCatalogItemType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowCatalogItemType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatured

`func (o *WorkflowCatalogItemType) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *WorkflowCatalogItemType) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *WorkflowCatalogItemType) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *WorkflowCatalogItemType) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *WorkflowCatalogItemType) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *WorkflowCatalogItemType) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *WorkflowCatalogItemType) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *WorkflowCatalogItemType) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetWorkflow

`func (o *WorkflowCatalogItemType) GetWorkflow() WorkflowCatalogItemTypeWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WorkflowCatalogItemType) GetWorkflowOk() (*WorkflowCatalogItemTypeWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WorkflowCatalogItemType) SetWorkflow(v WorkflowCatalogItemTypeWorkflow)`

SetWorkflow sets Workflow field to given value.


### GetContext

`func (o *WorkflowCatalogItemType) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *WorkflowCatalogItemType) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *WorkflowCatalogItemType) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *WorkflowCatalogItemType) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetWorkflowConfig

`func (o *WorkflowCatalogItemType) GetWorkflowConfig() string`

GetWorkflowConfig returns the WorkflowConfig field if non-nil, zero value otherwise.

### GetWorkflowConfigOk

`func (o *WorkflowCatalogItemType) GetWorkflowConfigOk() (*string, bool)`

GetWorkflowConfigOk returns a tuple with the WorkflowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowConfig

`func (o *WorkflowCatalogItemType) SetWorkflowConfig(v string)`

SetWorkflowConfig sets WorkflowConfig field to given value.

### HasWorkflowConfig

`func (o *WorkflowCatalogItemType) HasWorkflowConfig() bool`

HasWorkflowConfig returns a boolean if a field has been set.

### GetFormType

`func (o *WorkflowCatalogItemType) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *WorkflowCatalogItemType) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *WorkflowCatalogItemType) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *WorkflowCatalogItemType) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *WorkflowCatalogItemType) GetForm() WorkflowCatalogItemTypeForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *WorkflowCatalogItemType) GetFormOk() (*WorkflowCatalogItemTypeForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *WorkflowCatalogItemType) SetForm(v WorkflowCatalogItemTypeForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *WorkflowCatalogItemType) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetOptionTypes

`func (o *WorkflowCatalogItemType) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *WorkflowCatalogItemType) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *WorkflowCatalogItemType) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *WorkflowCatalogItemType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetContent

`func (o *WorkflowCatalogItemType) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WorkflowCatalogItemType) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WorkflowCatalogItemType) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WorkflowCatalogItemType) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


