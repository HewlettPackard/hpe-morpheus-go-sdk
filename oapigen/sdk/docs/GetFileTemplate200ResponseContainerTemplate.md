# GetFileTemplate200ResponseContainerTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**GetFileTemplate200ResponseContainerTemplateAccount**](GetFileTemplate200ResponseContainerTemplateAccount.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**TemplateType** | Pointer to **NullableString** |  | [optional] 
**TemplatePhase** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**SettingCategory** | Pointer to **NullableString** |  | [optional] 
**SettingName** | Pointer to **NullableString** |  | [optional] 
**AutoRun** | Pointer to **bool** |  | [optional] 
**RunOnScale** | Pointer to **NullableBool** |  | [optional] 
**RunOnDeploy** | Pointer to **NullableBool** |  | [optional] 
**FileOwner** | Pointer to **NullableString** |  | [optional] 
**FileGroup** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetFileTemplate200ResponseContainerTemplate

`func NewGetFileTemplate200ResponseContainerTemplate() *GetFileTemplate200ResponseContainerTemplate`

NewGetFileTemplate200ResponseContainerTemplate instantiates a new GetFileTemplate200ResponseContainerTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetFileTemplate200ResponseContainerTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetFileTemplate200ResponseContainerTemplate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetFileTemplate200ResponseContainerTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetFileTemplate200ResponseContainerTemplate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetFileTemplate200ResponseContainerTemplate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetFileTemplate200ResponseContainerTemplate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAccount

`func (o *GetFileTemplate200ResponseContainerTemplate) GetAccount() GetFileTemplate200ResponseContainerTemplateAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetAccountOk() (*GetFileTemplate200ResponseContainerTemplateAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetFileTemplate200ResponseContainerTemplate) SetAccount(v GetFileTemplate200ResponseContainerTemplateAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetFileTemplate200ResponseContainerTemplate) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *GetFileTemplate200ResponseContainerTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetFileTemplate200ResponseContainerTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetFileTemplate200ResponseContainerTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetFileTemplate200ResponseContainerTemplate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetFileTemplate200ResponseContainerTemplate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetFileTemplate200ResponseContainerTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetFileTemplate200ResponseContainerTemplate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetFileTemplate200ResponseContainerTemplate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetFileName

`func (o *GetFileTemplate200ResponseContainerTemplate) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *GetFileTemplate200ResponseContainerTemplate) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *GetFileTemplate200ResponseContainerTemplate) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilePath

`func (o *GetFileTemplate200ResponseContainerTemplate) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *GetFileTemplate200ResponseContainerTemplate) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *GetFileTemplate200ResponseContainerTemplate) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetTemplateType

`func (o *GetFileTemplate200ResponseContainerTemplate) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *GetFileTemplate200ResponseContainerTemplate) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *GetFileTemplate200ResponseContainerTemplate) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.

### SetTemplateTypeNil

`func (o *GetFileTemplate200ResponseContainerTemplate) SetTemplateTypeNil(b bool)`

 SetTemplateTypeNil sets the value for TemplateType to be an explicit nil

### UnsetTemplateType
`func (o *GetFileTemplate200ResponseContainerTemplate) UnsetTemplateType()`

UnsetTemplateType ensures that no value is present for TemplateType, not even an explicit nil
### GetTemplatePhase

`func (o *GetFileTemplate200ResponseContainerTemplate) GetTemplatePhase() string`

GetTemplatePhase returns the TemplatePhase field if non-nil, zero value otherwise.

### GetTemplatePhaseOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetTemplatePhaseOk() (*string, bool)`

GetTemplatePhaseOk returns a tuple with the TemplatePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePhase

`func (o *GetFileTemplate200ResponseContainerTemplate) SetTemplatePhase(v string)`

SetTemplatePhase sets TemplatePhase field to given value.

### HasTemplatePhase

`func (o *GetFileTemplate200ResponseContainerTemplate) HasTemplatePhase() bool`

HasTemplatePhase returns a boolean if a field has been set.

### GetTemplate

`func (o *GetFileTemplate200ResponseContainerTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetFileTemplate200ResponseContainerTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetFileTemplate200ResponseContainerTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetCategory

`func (o *GetFileTemplate200ResponseContainerTemplate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetFileTemplate200ResponseContainerTemplate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetFileTemplate200ResponseContainerTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetFileTemplate200ResponseContainerTemplate) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetFileTemplate200ResponseContainerTemplate) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetSettingCategory

`func (o *GetFileTemplate200ResponseContainerTemplate) GetSettingCategory() string`

GetSettingCategory returns the SettingCategory field if non-nil, zero value otherwise.

### GetSettingCategoryOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetSettingCategoryOk() (*string, bool)`

GetSettingCategoryOk returns a tuple with the SettingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingCategory

`func (o *GetFileTemplate200ResponseContainerTemplate) SetSettingCategory(v string)`

SetSettingCategory sets SettingCategory field to given value.

### HasSettingCategory

`func (o *GetFileTemplate200ResponseContainerTemplate) HasSettingCategory() bool`

HasSettingCategory returns a boolean if a field has been set.

### SetSettingCategoryNil

`func (o *GetFileTemplate200ResponseContainerTemplate) SetSettingCategoryNil(b bool)`

 SetSettingCategoryNil sets the value for SettingCategory to be an explicit nil

### UnsetSettingCategory
`func (o *GetFileTemplate200ResponseContainerTemplate) UnsetSettingCategory()`

UnsetSettingCategory ensures that no value is present for SettingCategory, not even an explicit nil
### GetSettingName

`func (o *GetFileTemplate200ResponseContainerTemplate) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *GetFileTemplate200ResponseContainerTemplate) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *GetFileTemplate200ResponseContainerTemplate) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### SetSettingNameNil

`func (o *GetFileTemplate200ResponseContainerTemplate) SetSettingNameNil(b bool)`

 SetSettingNameNil sets the value for SettingName to be an explicit nil

### UnsetSettingName
`func (o *GetFileTemplate200ResponseContainerTemplate) UnsetSettingName()`

UnsetSettingName ensures that no value is present for SettingName, not even an explicit nil
### GetAutoRun

`func (o *GetFileTemplate200ResponseContainerTemplate) GetAutoRun() bool`

GetAutoRun returns the AutoRun field if non-nil, zero value otherwise.

### GetAutoRunOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetAutoRunOk() (*bool, bool)`

GetAutoRunOk returns a tuple with the AutoRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRun

`func (o *GetFileTemplate200ResponseContainerTemplate) SetAutoRun(v bool)`

SetAutoRun sets AutoRun field to given value.

### HasAutoRun

`func (o *GetFileTemplate200ResponseContainerTemplate) HasAutoRun() bool`

HasAutoRun returns a boolean if a field has been set.

### GetRunOnScale

`func (o *GetFileTemplate200ResponseContainerTemplate) GetRunOnScale() bool`

GetRunOnScale returns the RunOnScale field if non-nil, zero value otherwise.

### GetRunOnScaleOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetRunOnScaleOk() (*bool, bool)`

GetRunOnScaleOk returns a tuple with the RunOnScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnScale

`func (o *GetFileTemplate200ResponseContainerTemplate) SetRunOnScale(v bool)`

SetRunOnScale sets RunOnScale field to given value.

### HasRunOnScale

`func (o *GetFileTemplate200ResponseContainerTemplate) HasRunOnScale() bool`

HasRunOnScale returns a boolean if a field has been set.

### SetRunOnScaleNil

`func (o *GetFileTemplate200ResponseContainerTemplate) SetRunOnScaleNil(b bool)`

 SetRunOnScaleNil sets the value for RunOnScale to be an explicit nil

### UnsetRunOnScale
`func (o *GetFileTemplate200ResponseContainerTemplate) UnsetRunOnScale()`

UnsetRunOnScale ensures that no value is present for RunOnScale, not even an explicit nil
### GetRunOnDeploy

`func (o *GetFileTemplate200ResponseContainerTemplate) GetRunOnDeploy() bool`

GetRunOnDeploy returns the RunOnDeploy field if non-nil, zero value otherwise.

### GetRunOnDeployOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetRunOnDeployOk() (*bool, bool)`

GetRunOnDeployOk returns a tuple with the RunOnDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunOnDeploy

`func (o *GetFileTemplate200ResponseContainerTemplate) SetRunOnDeploy(v bool)`

SetRunOnDeploy sets RunOnDeploy field to given value.

### HasRunOnDeploy

`func (o *GetFileTemplate200ResponseContainerTemplate) HasRunOnDeploy() bool`

HasRunOnDeploy returns a boolean if a field has been set.

### SetRunOnDeployNil

`func (o *GetFileTemplate200ResponseContainerTemplate) SetRunOnDeployNil(b bool)`

 SetRunOnDeployNil sets the value for RunOnDeploy to be an explicit nil

### UnsetRunOnDeploy
`func (o *GetFileTemplate200ResponseContainerTemplate) UnsetRunOnDeploy()`

UnsetRunOnDeploy ensures that no value is present for RunOnDeploy, not even an explicit nil
### GetFileOwner

`func (o *GetFileTemplate200ResponseContainerTemplate) GetFileOwner() string`

GetFileOwner returns the FileOwner field if non-nil, zero value otherwise.

### GetFileOwnerOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetFileOwnerOk() (*string, bool)`

GetFileOwnerOk returns a tuple with the FileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOwner

`func (o *GetFileTemplate200ResponseContainerTemplate) SetFileOwner(v string)`

SetFileOwner sets FileOwner field to given value.

### HasFileOwner

`func (o *GetFileTemplate200ResponseContainerTemplate) HasFileOwner() bool`

HasFileOwner returns a boolean if a field has been set.

### SetFileOwnerNil

`func (o *GetFileTemplate200ResponseContainerTemplate) SetFileOwnerNil(b bool)`

 SetFileOwnerNil sets the value for FileOwner to be an explicit nil

### UnsetFileOwner
`func (o *GetFileTemplate200ResponseContainerTemplate) UnsetFileOwner()`

UnsetFileOwner ensures that no value is present for FileOwner, not even an explicit nil
### GetFileGroup

`func (o *GetFileTemplate200ResponseContainerTemplate) GetFileGroup() string`

GetFileGroup returns the FileGroup field if non-nil, zero value otherwise.

### GetFileGroupOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetFileGroupOk() (*string, bool)`

GetFileGroupOk returns a tuple with the FileGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileGroup

`func (o *GetFileTemplate200ResponseContainerTemplate) SetFileGroup(v string)`

SetFileGroup sets FileGroup field to given value.

### HasFileGroup

`func (o *GetFileTemplate200ResponseContainerTemplate) HasFileGroup() bool`

HasFileGroup returns a boolean if a field has been set.

### SetFileGroupNil

`func (o *GetFileTemplate200ResponseContainerTemplate) SetFileGroupNil(b bool)`

 SetFileGroupNil sets the value for FileGroup to be an explicit nil

### UnsetFileGroup
`func (o *GetFileTemplate200ResponseContainerTemplate) UnsetFileGroup()`

UnsetFileGroup ensures that no value is present for FileGroup, not even an explicit nil
### GetPermissions

`func (o *GetFileTemplate200ResponseContainerTemplate) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetFileTemplate200ResponseContainerTemplate) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetFileTemplate200ResponseContainerTemplate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *GetFileTemplate200ResponseContainerTemplate) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *GetFileTemplate200ResponseContainerTemplate) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetDateCreated

`func (o *GetFileTemplate200ResponseContainerTemplate) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetFileTemplate200ResponseContainerTemplate) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetFileTemplate200ResponseContainerTemplate) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetFileTemplate200ResponseContainerTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetFileTemplate200ResponseContainerTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetFileTemplate200ResponseContainerTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetFileTemplate200ResponseContainerTemplate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


