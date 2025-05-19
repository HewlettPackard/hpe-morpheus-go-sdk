# ListFileTemplates200ResponseAllOfContainerTemplatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**TemplateType** | Pointer to **string** |  | [optional] 
**TemplatePhase** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**SettingCategory** | Pointer to **string** |  | [optional] 
**SettingName** | Pointer to **string** |  | [optional] 
**AutoRun** | Pointer to **bool** |  | [optional] 
**RunOnScale** | Pointer to **bool** |  | [optional] 
**RunOnDeploy** | Pointer to **bool** |  | [optional] 
**FileOwner** | Pointer to **string** |  | [optional] 
**FileGroup** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListFileTemplates200ResponseAllOfContainerTemplatesInner

`func NewListFileTemplates200ResponseAllOfContainerTemplatesInner() *ListFileTemplates200ResponseAllOfContainerTemplatesInner`

NewListFileTemplates200ResponseAllOfContainerTemplatesInner instantiates a new ListFileTemplates200ResponseAllOfContainerTemplatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFileTemplates200ResponseAllOfContainerTemplatesInnerWithDefaults

`func NewListFileTemplates200ResponseAllOfContainerTemplatesInnerWithDefaults() *ListFileTemplates200ResponseAllOfContainerTemplatesInner`

NewListFileTemplates200ResponseAllOfContainerTemplatesInnerWithDefaults instantiates a new ListFileTemplates200ResponseAllOfContainerTemplatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAccount

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFileName

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilePath

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetTemplateType

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.

### GetTemplatePhase

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplatePhase() string`

GetTemplatePhase returns the TemplatePhase field if non-nil, zero value otherwise.

### GetTemplatePhaseOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplatePhaseOk() (*string, bool)`

GetTemplatePhaseOk returns a tuple with the TemplatePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePhase

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetTemplatePhase(v string)`

SetTemplatePhase sets TemplatePhase field to given value.

### HasTemplatePhase

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasTemplatePhase() bool`

HasTemplatePhase returns a boolean if a field has been set.

### GetTemplate

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetCategory

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSettingCategory

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetSettingCategory() string`

GetSettingCategory returns the SettingCategory field if non-nil, zero value otherwise.

### GetSettingCategoryOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetSettingCategoryOk() (*string, bool)`

GetSettingCategoryOk returns a tuple with the SettingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingCategory

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetSettingCategory(v string)`

SetSettingCategory sets SettingCategory field to given value.

### HasSettingCategory

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasSettingCategory() bool`

HasSettingCategory returns a boolean if a field has been set.

### GetSettingName

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetAutoRun

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetAutoRun() bool`

GetAutoRun returns the AutoRun field if non-nil, zero value otherwise.

### GetAutoRunOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetAutoRunOk() (*bool, bool)`

GetAutoRunOk returns a tuple with the AutoRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRun

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetAutoRun(v bool)`

SetAutoRun sets AutoRun field to given value.

### HasAutoRun

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasAutoRun() bool`

HasAutoRun returns a boolean if a field has been set.

### GetRunOnScale

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetRunOnScale() bool`

GetRunOnScale returns the RunOnScale field if non-nil, zero value otherwise.

### GetRunOnScaleOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetRunOnScaleOk() (*bool, bool)`

GetRunOnScaleOk returns a tuple with the RunOnScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnScale

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetRunOnScale(v bool)`

SetRunOnScale sets RunOnScale field to given value.

### HasRunOnScale

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasRunOnScale() bool`

HasRunOnScale returns a boolean if a field has been set.

### GetRunOnDeploy

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetRunOnDeploy() bool`

GetRunOnDeploy returns the RunOnDeploy field if non-nil, zero value otherwise.

### GetRunOnDeployOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetRunOnDeployOk() (*bool, bool)`

GetRunOnDeployOk returns a tuple with the RunOnDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnDeploy

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetRunOnDeploy(v bool)`

SetRunOnDeploy sets RunOnDeploy field to given value.

### HasRunOnDeploy

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasRunOnDeploy() bool`

HasRunOnDeploy returns a boolean if a field has been set.

### GetFileOwner

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileOwner() string`

GetFileOwner returns the FileOwner field if non-nil, zero value otherwise.

### GetFileOwnerOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileOwnerOk() (*string, bool)`

GetFileOwnerOk returns a tuple with the FileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOwner

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetFileOwner(v string)`

SetFileOwner sets FileOwner field to given value.

### HasFileOwner

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasFileOwner() bool`

HasFileOwner returns a boolean if a field has been set.

### GetFileGroup

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileGroup() string`

GetFileGroup returns the FileGroup field if non-nil, zero value otherwise.

### GetFileGroupOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetFileGroupOk() (*string, bool)`

GetFileGroupOk returns a tuple with the FileGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileGroup

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetFileGroup(v string)`

SetFileGroup sets FileGroup field to given value.

### HasFileGroup

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasFileGroup() bool`

HasFileGroup returns a boolean if a field has been set.

### GetPermissions

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListFileTemplates200ResponseAllOfContainerTemplatesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


