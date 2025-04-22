# UpdateFileTemplateRequestContainerTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | File template name | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**FileName** | Pointer to **string** | Filename for the file template | [optional] 
**FilePath** | Pointer to **string** | Path for the file template | [optional] 
**Category** | Pointer to **string** | Category | [optional] 
**TemplatePhase** | Pointer to **string** | Template Phase, provision, start, etc. | [optional] 
**Template** | Pointer to **string** | Template content, that is, the file template content itself. | [optional] 
**FileOwner** | Pointer to **int64** | File Owner | [optional] 
**SettingName** | Pointer to **string** | Setting Name | [optional] 
**SettingCategory** | Pointer to **string** | Setting Category | [optional] 

## Methods

### NewUpdateFileTemplateRequestContainerTemplate

`func NewUpdateFileTemplateRequestContainerTemplate() *UpdateFileTemplateRequestContainerTemplate`

NewUpdateFileTemplateRequestContainerTemplate instantiates a new UpdateFileTemplateRequestContainerTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFileTemplateRequestContainerTemplateWithDefaults

`func NewUpdateFileTemplateRequestContainerTemplateWithDefaults() *UpdateFileTemplateRequestContainerTemplate`

NewUpdateFileTemplateRequestContainerTemplateWithDefaults instantiates a new UpdateFileTemplateRequestContainerTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateFileTemplateRequestContainerTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFileTemplateRequestContainerTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFileTemplateRequestContainerTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFileTemplateRequestContainerTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateFileTemplateRequestContainerTemplate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateFileTemplateRequestContainerTemplate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateFileTemplateRequestContainerTemplate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateFileTemplateRequestContainerTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFileName

`func (o *UpdateFileTemplateRequestContainerTemplate) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *UpdateFileTemplateRequestContainerTemplate) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *UpdateFileTemplateRequestContainerTemplate) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *UpdateFileTemplateRequestContainerTemplate) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilePath

`func (o *UpdateFileTemplateRequestContainerTemplate) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *UpdateFileTemplateRequestContainerTemplate) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *UpdateFileTemplateRequestContainerTemplate) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *UpdateFileTemplateRequestContainerTemplate) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateFileTemplateRequestContainerTemplate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateFileTemplateRequestContainerTemplate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateFileTemplateRequestContainerTemplate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateFileTemplateRequestContainerTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTemplatePhase

`func (o *UpdateFileTemplateRequestContainerTemplate) GetTemplatePhase() string`

GetTemplatePhase returns the TemplatePhase field if non-nil, zero value otherwise.

### GetTemplatePhaseOk

`func (o *UpdateFileTemplateRequestContainerTemplate) GetTemplatePhaseOk() (*string, bool)`

GetTemplatePhaseOk returns a tuple with the TemplatePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePhase

`func (o *UpdateFileTemplateRequestContainerTemplate) SetTemplatePhase(v string)`

SetTemplatePhase sets TemplatePhase field to given value.

### HasTemplatePhase

`func (o *UpdateFileTemplateRequestContainerTemplate) HasTemplatePhase() bool`

HasTemplatePhase returns a boolean if a field has been set.

### GetTemplate

`func (o *UpdateFileTemplateRequestContainerTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UpdateFileTemplateRequestContainerTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UpdateFileTemplateRequestContainerTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *UpdateFileTemplateRequestContainerTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetFileOwner

`func (o *UpdateFileTemplateRequestContainerTemplate) GetFileOwner() int64`

GetFileOwner returns the FileOwner field if non-nil, zero value otherwise.

### GetFileOwnerOk

`func (o *UpdateFileTemplateRequestContainerTemplate) GetFileOwnerOk() (*int64, bool)`

GetFileOwnerOk returns a tuple with the FileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOwner

`func (o *UpdateFileTemplateRequestContainerTemplate) SetFileOwner(v int64)`

SetFileOwner sets FileOwner field to given value.

### HasFileOwner

`func (o *UpdateFileTemplateRequestContainerTemplate) HasFileOwner() bool`

HasFileOwner returns a boolean if a field has been set.

### GetSettingName

`func (o *UpdateFileTemplateRequestContainerTemplate) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *UpdateFileTemplateRequestContainerTemplate) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *UpdateFileTemplateRequestContainerTemplate) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *UpdateFileTemplateRequestContainerTemplate) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetSettingCategory

`func (o *UpdateFileTemplateRequestContainerTemplate) GetSettingCategory() string`

GetSettingCategory returns the SettingCategory field if non-nil, zero value otherwise.

### GetSettingCategoryOk

`func (o *UpdateFileTemplateRequestContainerTemplate) GetSettingCategoryOk() (*string, bool)`

GetSettingCategoryOk returns a tuple with the SettingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingCategory

`func (o *UpdateFileTemplateRequestContainerTemplate) SetSettingCategory(v string)`

SetSettingCategory sets SettingCategory field to given value.

### HasSettingCategory

`func (o *UpdateFileTemplateRequestContainerTemplate) HasSettingCategory() bool`

HasSettingCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


