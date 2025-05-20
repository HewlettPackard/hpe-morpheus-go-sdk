# ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** | Useful shortcode for provisioning naming schemes and export reference. | [optional] 
**Category** | Pointer to **string** | Catalog Item Type category | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**AllowQuantity** | Pointer to **bool** | Can users order more than one of this item at a time. | [optional] 
**IconPath** | Pointer to **string** |  | [optional] 
**ImagePath** | Pointer to **string** |  | [optional] 
**DarkImagePath** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**LayoutCode** | Pointer to **string** |  | [optional] 
**Blueprint** | Pointer to **map[string]interface{}** |  | [optional] 
**AppSpec** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**InstanceSpec** | Pointer to **string** |  | [optional] 
**Workflow** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**FormType** | Pointer to **string** |  | [optional] 
**Form** | Pointer to **map[string]interface{}** | Form object that contains input options and/or field groups | [optional] 
**FormConfig** | Pointer to **map[string]interface{}** | Form config object | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner

`func NewListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner() *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner`

NewListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner instantiates a new ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerWithDefaults

`func NewListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerWithDefaults() *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner`

NewListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerWithDefaults instantiates a new ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFeatured

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetIconPath

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetImagePath

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetDarkImagePath

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### GetVisibility

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLayoutCode

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetLayoutCode() string`

GetLayoutCode returns the LayoutCode field if non-nil, zero value otherwise.

### GetLayoutCodeOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetLayoutCodeOk() (*string, bool)`

GetLayoutCodeOk returns a tuple with the LayoutCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutCode

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetLayoutCode(v string)`

SetLayoutCode sets LayoutCode field to given value.

### HasLayoutCode

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasLayoutCode() bool`

HasLayoutCode returns a boolean if a field has been set.

### GetBlueprint

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetBlueprint() map[string]interface{}`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetBlueprintOk() (*map[string]interface{}, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetBlueprint(v map[string]interface{})`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetAppSpec

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetAppSpec() string`

GetAppSpec returns the AppSpec field if non-nil, zero value otherwise.

### GetAppSpecOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetAppSpecOk() (*string, bool)`

GetAppSpecOk returns a tuple with the AppSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSpec

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetAppSpec(v string)`

SetAppSpec sets AppSpec field to given value.

### HasAppSpec

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasAppSpec() bool`

HasAppSpec returns a boolean if a field has been set.

### GetConfig

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetInstanceSpec

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetInstanceSpec() string`

GetInstanceSpec returns the InstanceSpec field if non-nil, zero value otherwise.

### GetInstanceSpecOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetInstanceSpecOk() (*string, bool)`

GetInstanceSpecOk returns a tuple with the InstanceSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSpec

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetInstanceSpec(v string)`

SetInstanceSpec sets InstanceSpec field to given value.

### HasInstanceSpec

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasInstanceSpec() bool`

HasInstanceSpec returns a boolean if a field has been set.

### GetWorkflow

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetWorkflow() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetWorkflowOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetWorkflow(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetContent

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFormType

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetForm() map[string]interface{}`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetFormOk() (*map[string]interface{}, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetForm(v map[string]interface{})`

SetForm sets Form field to given value.

### HasForm

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetFormConfig

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetFormConfig() map[string]interface{}`

GetFormConfig returns the FormConfig field if non-nil, zero value otherwise.

### GetFormConfigOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetFormConfigOk() (*map[string]interface{}, bool)`

GetFormConfigOk returns a tuple with the FormConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormConfig

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetFormConfig(v map[string]interface{})`

SetFormConfig sets FormConfig field to given value.

### HasFormConfig

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasFormConfig() bool`

HasFormConfig returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetOwner

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


