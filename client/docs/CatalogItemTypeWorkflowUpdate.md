# CatalogItemTypeWorkflowUpdate

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
**Workflow** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Context** | Pointer to **string** | Context for running the workflow, determines if a target resource must be selected. | [optional] 
**WorkflowConfig** | Pointer to **string** | Configuration object that contains settings for the workflow. | [optional] 
**FormType** | Pointer to **string** | Form Type determines if the configuration options come from a Form (form) or a list of Inputs (optionTypes). | [optional] [default to "optionTypes"]
**Form** | Pointer to [**AddCatalogItemTypeRequestCatalogItemTypeOneOfForm**](AddCatalogItemTypeRequestCatalogItemTypeOneOfForm.md) |  | [optional] 
**OptionTypes** | Pointer to **[]int64** | Array of option type IDs. Only applies to formType &#39;optionTypes&#39;. | [optional] 
**Content** | Pointer to **string** | Documentation content for this Catalog Item. Markdown-formatted text is accepted and displayed appropriately when the item is ordered from the Service Catalog. A new Catalog Item-type Wiki entry will also be added containing this information. | [optional] 

## Methods

### NewCatalogItemTypeWorkflowUpdate

`func NewCatalogItemTypeWorkflowUpdate() *CatalogItemTypeWorkflowUpdate`

NewCatalogItemTypeWorkflowUpdate instantiates a new CatalogItemTypeWorkflowUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemTypeWorkflowUpdateWithDefaults

`func NewCatalogItemTypeWorkflowUpdateWithDefaults() *CatalogItemTypeWorkflowUpdate`

NewCatalogItemTypeWorkflowUpdateWithDefaults instantiates a new CatalogItemTypeWorkflowUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CatalogItemTypeWorkflowUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogItemTypeWorkflowUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogItemTypeWorkflowUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogItemTypeWorkflowUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *CatalogItemTypeWorkflowUpdate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CatalogItemTypeWorkflowUpdate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CatalogItemTypeWorkflowUpdate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CatalogItemTypeWorkflowUpdate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *CatalogItemTypeWorkflowUpdate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CatalogItemTypeWorkflowUpdate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CatalogItemTypeWorkflowUpdate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CatalogItemTypeWorkflowUpdate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogItemTypeWorkflowUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogItemTypeWorkflowUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogItemTypeWorkflowUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogItemTypeWorkflowUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *CatalogItemTypeWorkflowUpdate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CatalogItemTypeWorkflowUpdate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CatalogItemTypeWorkflowUpdate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CatalogItemTypeWorkflowUpdate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *CatalogItemTypeWorkflowUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogItemTypeWorkflowUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogItemTypeWorkflowUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogItemTypeWorkflowUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *CatalogItemTypeWorkflowUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CatalogItemTypeWorkflowUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CatalogItemTypeWorkflowUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CatalogItemTypeWorkflowUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *CatalogItemTypeWorkflowUpdate) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *CatalogItemTypeWorkflowUpdate) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *CatalogItemTypeWorkflowUpdate) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *CatalogItemTypeWorkflowUpdate) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### GetIconPath

`func (o *CatalogItemTypeWorkflowUpdate) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *CatalogItemTypeWorkflowUpdate) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *CatalogItemTypeWorkflowUpdate) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *CatalogItemTypeWorkflowUpdate) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetEnabled

`func (o *CatalogItemTypeWorkflowUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CatalogItemTypeWorkflowUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CatalogItemTypeWorkflowUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CatalogItemTypeWorkflowUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatured

`func (o *CatalogItemTypeWorkflowUpdate) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *CatalogItemTypeWorkflowUpdate) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *CatalogItemTypeWorkflowUpdate) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *CatalogItemTypeWorkflowUpdate) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *CatalogItemTypeWorkflowUpdate) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *CatalogItemTypeWorkflowUpdate) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *CatalogItemTypeWorkflowUpdate) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *CatalogItemTypeWorkflowUpdate) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetWorkflow

`func (o *CatalogItemTypeWorkflowUpdate) GetWorkflow() GetAlerts200ResponseAllOfChecksInnerAccount`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *CatalogItemTypeWorkflowUpdate) GetWorkflowOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *CatalogItemTypeWorkflowUpdate) SetWorkflow(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *CatalogItemTypeWorkflowUpdate) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetContext

`func (o *CatalogItemTypeWorkflowUpdate) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CatalogItemTypeWorkflowUpdate) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CatalogItemTypeWorkflowUpdate) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *CatalogItemTypeWorkflowUpdate) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetWorkflowConfig

`func (o *CatalogItemTypeWorkflowUpdate) GetWorkflowConfig() string`

GetWorkflowConfig returns the WorkflowConfig field if non-nil, zero value otherwise.

### GetWorkflowConfigOk

`func (o *CatalogItemTypeWorkflowUpdate) GetWorkflowConfigOk() (*string, bool)`

GetWorkflowConfigOk returns a tuple with the WorkflowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowConfig

`func (o *CatalogItemTypeWorkflowUpdate) SetWorkflowConfig(v string)`

SetWorkflowConfig sets WorkflowConfig field to given value.

### HasWorkflowConfig

`func (o *CatalogItemTypeWorkflowUpdate) HasWorkflowConfig() bool`

HasWorkflowConfig returns a boolean if a field has been set.

### GetFormType

`func (o *CatalogItemTypeWorkflowUpdate) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *CatalogItemTypeWorkflowUpdate) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *CatalogItemTypeWorkflowUpdate) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *CatalogItemTypeWorkflowUpdate) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *CatalogItemTypeWorkflowUpdate) GetForm() AddCatalogItemTypeRequestCatalogItemTypeOneOfForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CatalogItemTypeWorkflowUpdate) GetFormOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CatalogItemTypeWorkflowUpdate) SetForm(v AddCatalogItemTypeRequestCatalogItemTypeOneOfForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *CatalogItemTypeWorkflowUpdate) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetOptionTypes

`func (o *CatalogItemTypeWorkflowUpdate) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *CatalogItemTypeWorkflowUpdate) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *CatalogItemTypeWorkflowUpdate) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *CatalogItemTypeWorkflowUpdate) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetContent

`func (o *CatalogItemTypeWorkflowUpdate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CatalogItemTypeWorkflowUpdate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CatalogItemTypeWorkflowUpdate) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CatalogItemTypeWorkflowUpdate) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


