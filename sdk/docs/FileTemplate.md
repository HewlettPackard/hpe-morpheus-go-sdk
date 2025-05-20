# FileTemplate

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

### NewFileTemplate

`func NewFileTemplate() *FileTemplate`

NewFileTemplate instantiates a new FileTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileTemplateWithDefaults

`func NewFileTemplateWithDefaults() *FileTemplate`

NewFileTemplateWithDefaults instantiates a new FileTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileTemplate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FileTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *FileTemplate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FileTemplate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FileTemplate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FileTemplate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAccount

`func (o *FileTemplate) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *FileTemplate) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *FileTemplate) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *FileTemplate) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *FileTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *FileTemplate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FileTemplate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FileTemplate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *FileTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFileName

`func (o *FileTemplate) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FileTemplate) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FileTemplate) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FileTemplate) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilePath

`func (o *FileTemplate) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FileTemplate) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FileTemplate) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FileTemplate) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetTemplateType

`func (o *FileTemplate) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *FileTemplate) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *FileTemplate) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *FileTemplate) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.

### GetTemplatePhase

`func (o *FileTemplate) GetTemplatePhase() string`

GetTemplatePhase returns the TemplatePhase field if non-nil, zero value otherwise.

### GetTemplatePhaseOk

`func (o *FileTemplate) GetTemplatePhaseOk() (*string, bool)`

GetTemplatePhaseOk returns a tuple with the TemplatePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePhase

`func (o *FileTemplate) SetTemplatePhase(v string)`

SetTemplatePhase sets TemplatePhase field to given value.

### HasTemplatePhase

`func (o *FileTemplate) HasTemplatePhase() bool`

HasTemplatePhase returns a boolean if a field has been set.

### GetTemplate

`func (o *FileTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *FileTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *FileTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *FileTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetCategory

`func (o *FileTemplate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FileTemplate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FileTemplate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FileTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSettingCategory

`func (o *FileTemplate) GetSettingCategory() string`

GetSettingCategory returns the SettingCategory field if non-nil, zero value otherwise.

### GetSettingCategoryOk

`func (o *FileTemplate) GetSettingCategoryOk() (*string, bool)`

GetSettingCategoryOk returns a tuple with the SettingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingCategory

`func (o *FileTemplate) SetSettingCategory(v string)`

SetSettingCategory sets SettingCategory field to given value.

### HasSettingCategory

`func (o *FileTemplate) HasSettingCategory() bool`

HasSettingCategory returns a boolean if a field has been set.

### GetSettingName

`func (o *FileTemplate) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *FileTemplate) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *FileTemplate) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *FileTemplate) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetAutoRun

`func (o *FileTemplate) GetAutoRun() bool`

GetAutoRun returns the AutoRun field if non-nil, zero value otherwise.

### GetAutoRunOk

`func (o *FileTemplate) GetAutoRunOk() (*bool, bool)`

GetAutoRunOk returns a tuple with the AutoRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRun

`func (o *FileTemplate) SetAutoRun(v bool)`

SetAutoRun sets AutoRun field to given value.

### HasAutoRun

`func (o *FileTemplate) HasAutoRun() bool`

HasAutoRun returns a boolean if a field has been set.

### GetRunOnScale

`func (o *FileTemplate) GetRunOnScale() bool`

GetRunOnScale returns the RunOnScale field if non-nil, zero value otherwise.

### GetRunOnScaleOk

`func (o *FileTemplate) GetRunOnScaleOk() (*bool, bool)`

GetRunOnScaleOk returns a tuple with the RunOnScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnScale

`func (o *FileTemplate) SetRunOnScale(v bool)`

SetRunOnScale sets RunOnScale field to given value.

### HasRunOnScale

`func (o *FileTemplate) HasRunOnScale() bool`

HasRunOnScale returns a boolean if a field has been set.

### GetRunOnDeploy

`func (o *FileTemplate) GetRunOnDeploy() bool`

GetRunOnDeploy returns the RunOnDeploy field if non-nil, zero value otherwise.

### GetRunOnDeployOk

`func (o *FileTemplate) GetRunOnDeployOk() (*bool, bool)`

GetRunOnDeployOk returns a tuple with the RunOnDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnDeploy

`func (o *FileTemplate) SetRunOnDeploy(v bool)`

SetRunOnDeploy sets RunOnDeploy field to given value.

### HasRunOnDeploy

`func (o *FileTemplate) HasRunOnDeploy() bool`

HasRunOnDeploy returns a boolean if a field has been set.

### GetFileOwner

`func (o *FileTemplate) GetFileOwner() string`

GetFileOwner returns the FileOwner field if non-nil, zero value otherwise.

### GetFileOwnerOk

`func (o *FileTemplate) GetFileOwnerOk() (*string, bool)`

GetFileOwnerOk returns a tuple with the FileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOwner

`func (o *FileTemplate) SetFileOwner(v string)`

SetFileOwner sets FileOwner field to given value.

### HasFileOwner

`func (o *FileTemplate) HasFileOwner() bool`

HasFileOwner returns a boolean if a field has been set.

### GetFileGroup

`func (o *FileTemplate) GetFileGroup() string`

GetFileGroup returns the FileGroup field if non-nil, zero value otherwise.

### GetFileGroupOk

`func (o *FileTemplate) GetFileGroupOk() (*string, bool)`

GetFileGroupOk returns a tuple with the FileGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileGroup

`func (o *FileTemplate) SetFileGroup(v string)`

SetFileGroup sets FileGroup field to given value.

### HasFileGroup

`func (o *FileTemplate) HasFileGroup() bool`

HasFileGroup returns a boolean if a field has been set.

### GetPermissions

`func (o *FileTemplate) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FileTemplate) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FileTemplate) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *FileTemplate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetDateCreated

`func (o *FileTemplate) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *FileTemplate) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *FileTemplate) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *FileTemplate) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *FileTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *FileTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *FileTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *FileTemplate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


