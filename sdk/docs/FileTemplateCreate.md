# FileTemplateCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | File template name | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**FileName** | **string** | Filename for the file template | 
**FilePath** | Pointer to **string** | Path for the file template | [optional] 
**Category** | Pointer to **string** | Category | [optional] 
**TemplatePhase** | Pointer to **string** | Template Phase, provision, start, etc. | [optional] 
**Template** | Pointer to **string** | Template content, that is, the file template content itself. | [optional] 
**FileOwner** | Pointer to **int64** | File Owner | [optional] 
**SettingName** | Pointer to **string** | Setting Name | [optional] 
**SettingCategory** | Pointer to **string** | Setting Category | [optional] 

## Methods

### NewFileTemplateCreate

`func NewFileTemplateCreate(name string, fileName string, ) *FileTemplateCreate`

NewFileTemplateCreate instantiates a new FileTemplateCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileTemplateCreateWithDefaults

`func NewFileTemplateCreateWithDefaults() *FileTemplateCreate`

NewFileTemplateCreateWithDefaults instantiates a new FileTemplateCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileTemplateCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileTemplateCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileTemplateCreate) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *FileTemplateCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FileTemplateCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FileTemplateCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *FileTemplateCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFileName

`func (o *FileTemplateCreate) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *FileTemplateCreate) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *FileTemplateCreate) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFilePath

`func (o *FileTemplateCreate) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *FileTemplateCreate) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *FileTemplateCreate) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *FileTemplateCreate) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetCategory

`func (o *FileTemplateCreate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *FileTemplateCreate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *FileTemplateCreate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *FileTemplateCreate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTemplatePhase

`func (o *FileTemplateCreate) GetTemplatePhase() string`

GetTemplatePhase returns the TemplatePhase field if non-nil, zero value otherwise.

### GetTemplatePhaseOk

`func (o *FileTemplateCreate) GetTemplatePhaseOk() (*string, bool)`

GetTemplatePhaseOk returns a tuple with the TemplatePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePhase

`func (o *FileTemplateCreate) SetTemplatePhase(v string)`

SetTemplatePhase sets TemplatePhase field to given value.

### HasTemplatePhase

`func (o *FileTemplateCreate) HasTemplatePhase() bool`

HasTemplatePhase returns a boolean if a field has been set.

### GetTemplate

`func (o *FileTemplateCreate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *FileTemplateCreate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *FileTemplateCreate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *FileTemplateCreate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetFileOwner

`func (o *FileTemplateCreate) GetFileOwner() int64`

GetFileOwner returns the FileOwner field if non-nil, zero value otherwise.

### GetFileOwnerOk

`func (o *FileTemplateCreate) GetFileOwnerOk() (*int64, bool)`

GetFileOwnerOk returns a tuple with the FileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOwner

`func (o *FileTemplateCreate) SetFileOwner(v int64)`

SetFileOwner sets FileOwner field to given value.

### HasFileOwner

`func (o *FileTemplateCreate) HasFileOwner() bool`

HasFileOwner returns a boolean if a field has been set.

### GetSettingName

`func (o *FileTemplateCreate) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *FileTemplateCreate) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *FileTemplateCreate) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *FileTemplateCreate) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetSettingCategory

`func (o *FileTemplateCreate) GetSettingCategory() string`

GetSettingCategory returns the SettingCategory field if non-nil, zero value otherwise.

### GetSettingCategoryOk

`func (o *FileTemplateCreate) GetSettingCategoryOk() (*string, bool)`

GetSettingCategoryOk returns a tuple with the SettingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingCategory

`func (o *FileTemplateCreate) SetSettingCategory(v string)`

SetSettingCategory sets SettingCategory field to given value.

### HasSettingCategory

`func (o *FileTemplateCreate) HasSettingCategory() bool`

HasSettingCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


