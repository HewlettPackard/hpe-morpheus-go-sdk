# AddCatalogItemTypeRequestCatalogItemTypeOneOf2

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
**Workflow** | [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | 
**Context** | Pointer to **string** | Context for running the workflow, determines if a target resource must be selected. | [optional] 
**WorkflowConfig** | Pointer to **string** | Configuration object that contains settings for the workflow. | [optional] 
**FormType** | Pointer to **string** | Form Type determines if the configuration options come from a Form (form) or a list of Inputs (optionTypes). | [optional] [default to "optionTypes"]
**Form** | Pointer to [**AddCatalogItemTypeRequestCatalogItemTypeOneOfForm**](AddCatalogItemTypeRequestCatalogItemTypeOneOfForm.md) |  | [optional] 
**OptionTypes** | Pointer to **[]int64** | Array of option type IDs. Only applies to formType &#39;optionTypes&#39;. | [optional] 
**Content** | Pointer to **string** | Documentation content for this Catalog Item. Markdown-formatted text is accepted and displayed appropriately when the item is ordered from the Service Catalog. A new Catalog Item-type Wiki entry will also be added containing this information. | [optional] 

## Methods

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOf2

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOf2(workflow GetAlerts200ResponseAllOfChecksInnerAccount, ) *AddCatalogItemTypeRequestCatalogItemTypeOneOf2`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOf2 instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOf2WithDefaults

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOf2WithDefaults() *AddCatalogItemTypeRequestCatalogItemTypeOneOf2`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOf2WithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### GetIconPath

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatured

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetWorkflow

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetWorkflow() GetAlerts200ResponseAllOfChecksInnerAccount`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetWorkflowOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetWorkflow(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetWorkflow sets Workflow field to given value.


### GetContext

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetWorkflowConfig

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetWorkflowConfig() string`

GetWorkflowConfig returns the WorkflowConfig field if non-nil, zero value otherwise.

### GetWorkflowConfigOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetWorkflowConfigOk() (*string, bool)`

GetWorkflowConfigOk returns a tuple with the WorkflowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowConfig

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetWorkflowConfig(v string)`

SetWorkflowConfig sets WorkflowConfig field to given value.

### HasWorkflowConfig

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasWorkflowConfig() bool`

HasWorkflowConfig returns a boolean if a field has been set.

### GetFormType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetForm() AddCatalogItemTypeRequestCatalogItemTypeOneOfForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetFormOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetForm(v AddCatalogItemTypeRequestCatalogItemTypeOneOfForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetOptionTypes

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetContent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOf2) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


