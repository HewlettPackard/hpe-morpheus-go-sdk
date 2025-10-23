# AddCatalogItemTypeRequestCatalogItemType

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
**Config** | [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig.md) |  | 
**InstanceSpec** | Pointer to **string** | The instance &#x60;config&#x60; specification as a string in the JSON format. | [optional] 
**FormType** | Pointer to **string** | Form Type determines if the configuration options come from a Form (form) or a list of Inputs (optionTypes). | [optional] [default to "optionTypes"]
**Form** | Pointer to [**AddCatalogItemTypeRequestCatalogItemTypeOneOfForm**](AddCatalogItemTypeRequestCatalogItemTypeOneOfForm.md) |  | [optional] 
**OptionTypes** | Pointer to **[]int64** | Array of option type IDs. Only applies to formType &#39;optionTypes&#39;. | [optional] 
**Content** | Pointer to **string** | Documentation content for this Catalog Item. Markdown-formatted text is accepted and displayed appropriately when the item is ordered from the Service Catalog. A new Catalog Item-type Wiki entry will also be added containing this information. | [optional] 
**Blueprint** | [**AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint**](AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint.md) |  | 
**AppSpec** | Pointer to **string** | The appSpec for blueprint type catalog items is a string in the Scribe YAML format with fields | [optional] 
**Workflow** | [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | 
**Context** | Pointer to **string** | Context for running the workflow, determines if a target resource must be selected. | [optional] 
**WorkflowConfig** | Pointer to **string** | Configuration object that contains settings for the workflow. | [optional] 

## Methods

### NewAddCatalogItemTypeRequestCatalogItemType

`func NewAddCatalogItemTypeRequestCatalogItemType(config AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig, blueprint AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint, workflow GetAlerts200ResponseAllOfChecksInnerAccount, ) *AddCatalogItemTypeRequestCatalogItemType`

NewAddCatalogItemTypeRequestCatalogItemType instantiates a new AddCatalogItemTypeRequestCatalogItemType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogItemTypeRequestCatalogItemTypeWithDefaults

`func NewAddCatalogItemTypeRequestCatalogItemTypeWithDefaults() *AddCatalogItemTypeRequestCatalogItemType`

NewAddCatalogItemTypeRequestCatalogItemTypeWithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVisibility

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### GetIconPath

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatured

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetConfig

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetConfig() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetConfigOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetConfig(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig)`

SetConfig sets Config field to given value.


### GetInstanceSpec

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetInstanceSpec() string`

GetInstanceSpec returns the InstanceSpec field if non-nil, zero value otherwise.

### GetInstanceSpecOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetInstanceSpecOk() (*string, bool)`

GetInstanceSpecOk returns a tuple with the InstanceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSpec

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetInstanceSpec(v string)`

SetInstanceSpec sets InstanceSpec field to given value.

### HasInstanceSpec

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasInstanceSpec() bool`

HasInstanceSpec returns a boolean if a field has been set.

### GetFormType

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetForm() AddCatalogItemTypeRequestCatalogItemTypeOneOfForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetFormOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetForm(v AddCatalogItemTypeRequestCatalogItemTypeOneOfForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetOptionTypes

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetContent

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetBlueprint

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetBlueprint() AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetBlueprintOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetBlueprint(v AddCatalogItemTypeRequestCatalogItemTypeOneOf1Blueprint)`

SetBlueprint sets Blueprint field to given value.


### GetAppSpec

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetAppSpec() string`

GetAppSpec returns the AppSpec field if non-nil, zero value otherwise.

### GetAppSpecOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetAppSpecOk() (*string, bool)`

GetAppSpecOk returns a tuple with the AppSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSpec

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetAppSpec(v string)`

SetAppSpec sets AppSpec field to given value.

### HasAppSpec

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasAppSpec() bool`

HasAppSpec returns a boolean if a field has been set.

### GetWorkflow

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetWorkflow() GetAlerts200ResponseAllOfChecksInnerAccount`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetWorkflowOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetWorkflow(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetWorkflow sets Workflow field to given value.


### GetContext

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetWorkflowConfig

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetWorkflowConfig() string`

GetWorkflowConfig returns the WorkflowConfig field if non-nil, zero value otherwise.

### GetWorkflowConfigOk

`func (o *AddCatalogItemTypeRequestCatalogItemType) GetWorkflowConfigOk() (*string, bool)`

GetWorkflowConfigOk returns a tuple with the WorkflowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowConfig

`func (o *AddCatalogItemTypeRequestCatalogItemType) SetWorkflowConfig(v string)`

SetWorkflowConfig sets WorkflowConfig field to given value.

### HasWorkflowConfig

`func (o *AddCatalogItemTypeRequestCatalogItemType) HasWorkflowConfig() bool`

HasWorkflowConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


