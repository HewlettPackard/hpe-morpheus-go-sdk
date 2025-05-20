# FileTemplateUpdate

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

### NewFileTemplateUpdate

`func NewFileTemplateUpdate() *FileTemplateUpdate`

NewFileTemplateUpdate instantiates a new FileTemplateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileTemplateUpdateWithDefaults

`func NewFileTemplateUpdateWithDefaults() *FileTemplateUpdate`

NewFileTemplateUpdateWithDefaults instantiates a new FileTemplateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileTemplateUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileTemplateUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileTemplateUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileTemplateUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *FileTemplateUpdate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FileTemplateUpdate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FileTemplateUpdate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *FileTemplateUpdate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFileName

`func (o *FileTemplateUpdate) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FileTemplateUpdate) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FileTemplateUpdate) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *FileTemplateUpdate) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilePath

`func (o *FileTemplateUpdate) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FileTemplateUpdate) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FileTemplateUpdate) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FileTemplateUpdate) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetCategory

`func (o *FileTemplateUpdate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FileTemplateUpdate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FileTemplateUpdate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FileTemplateUpdate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTemplatePhase

`func (o *FileTemplateUpdate) GetTemplatePhase() string`

GetTemplatePhase returns the TemplatePhase field if non-nil, zero value otherwise.

### GetTemplatePhaseOk

`func (o *FileTemplateUpdate) GetTemplatePhaseOk() (*string, bool)`

GetTemplatePhaseOk returns a tuple with the TemplatePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePhase

`func (o *FileTemplateUpdate) SetTemplatePhase(v string)`

SetTemplatePhase sets TemplatePhase field to given value.

### HasTemplatePhase

`func (o *FileTemplateUpdate) HasTemplatePhase() bool`

HasTemplatePhase returns a boolean if a field has been set.

### GetTemplate

`func (o *FileTemplateUpdate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *FileTemplateUpdate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *FileTemplateUpdate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *FileTemplateUpdate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetFileOwner

`func (o *FileTemplateUpdate) GetFileOwner() int64`

GetFileOwner returns the FileOwner field if non-nil, zero value otherwise.

### GetFileOwnerOk

`func (o *FileTemplateUpdate) GetFileOwnerOk() (*int64, bool)`

GetFileOwnerOk returns a tuple with the FileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOwner

`func (o *FileTemplateUpdate) SetFileOwner(v int64)`

SetFileOwner sets FileOwner field to given value.

### HasFileOwner

`func (o *FileTemplateUpdate) HasFileOwner() bool`

HasFileOwner returns a boolean if a field has been set.

### GetSettingName

`func (o *FileTemplateUpdate) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *FileTemplateUpdate) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *FileTemplateUpdate) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *FileTemplateUpdate) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetSettingCategory

`func (o *FileTemplateUpdate) GetSettingCategory() string`

GetSettingCategory returns the SettingCategory field if non-nil, zero value otherwise.

### GetSettingCategoryOk

`func (o *FileTemplateUpdate) GetSettingCategoryOk() (*string, bool)`

GetSettingCategoryOk returns a tuple with the SettingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingCategory

`func (o *FileTemplateUpdate) SetSettingCategory(v string)`

SetSettingCategory sets SettingCategory field to given value.

### HasSettingCategory

`func (o *FileTemplateUpdate) HasSettingCategory() bool`

HasSettingCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


