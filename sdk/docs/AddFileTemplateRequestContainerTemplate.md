# AddFileTemplateRequestContainerTemplate

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

### NewAddFileTemplateRequestContainerTemplate

`func NewAddFileTemplateRequestContainerTemplate(name string, fileName string, ) *AddFileTemplateRequestContainerTemplate`

NewAddFileTemplateRequestContainerTemplate instantiates a new AddFileTemplateRequestContainerTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddFileTemplateRequestContainerTemplateWithDefaults

`func NewAddFileTemplateRequestContainerTemplateWithDefaults() *AddFileTemplateRequestContainerTemplate`

NewAddFileTemplateRequestContainerTemplateWithDefaults instantiates a new AddFileTemplateRequestContainerTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddFileTemplateRequestContainerTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddFileTemplateRequestContainerTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddFileTemplateRequestContainerTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *AddFileTemplateRequestContainerTemplate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddFileTemplateRequestContainerTemplate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddFileTemplateRequestContainerTemplate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddFileTemplateRequestContainerTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFileName

`func (o *AddFileTemplateRequestContainerTemplate) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AddFileTemplateRequestContainerTemplate) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AddFileTemplateRequestContainerTemplate) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFilePath

`func (o *AddFileTemplateRequestContainerTemplate) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *AddFileTemplateRequestContainerTemplate) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *AddFileTemplateRequestContainerTemplate) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *AddFileTemplateRequestContainerTemplate) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetCategory

`func (o *AddFileTemplateRequestContainerTemplate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddFileTemplateRequestContainerTemplate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddFileTemplateRequestContainerTemplate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddFileTemplateRequestContainerTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTemplatePhase

`func (o *AddFileTemplateRequestContainerTemplate) GetTemplatePhase() string`

GetTemplatePhase returns the TemplatePhase field if non-nil, zero value otherwise.

### GetTemplatePhaseOk

`func (o *AddFileTemplateRequestContainerTemplate) GetTemplatePhaseOk() (*string, bool)`

GetTemplatePhaseOk returns a tuple with the TemplatePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePhase

`func (o *AddFileTemplateRequestContainerTemplate) SetTemplatePhase(v string)`

SetTemplatePhase sets TemplatePhase field to given value.

### HasTemplatePhase

`func (o *AddFileTemplateRequestContainerTemplate) HasTemplatePhase() bool`

HasTemplatePhase returns a boolean if a field has been set.

### GetTemplate

`func (o *AddFileTemplateRequestContainerTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AddFileTemplateRequestContainerTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AddFileTemplateRequestContainerTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *AddFileTemplateRequestContainerTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetFileOwner

`func (o *AddFileTemplateRequestContainerTemplate) GetFileOwner() int64`

GetFileOwner returns the FileOwner field if non-nil, zero value otherwise.

### GetFileOwnerOk

`func (o *AddFileTemplateRequestContainerTemplate) GetFileOwnerOk() (*int64, bool)`

GetFileOwnerOk returns a tuple with the FileOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileOwner

`func (o *AddFileTemplateRequestContainerTemplate) SetFileOwner(v int64)`

SetFileOwner sets FileOwner field to given value.

### HasFileOwner

`func (o *AddFileTemplateRequestContainerTemplate) HasFileOwner() bool`

HasFileOwner returns a boolean if a field has been set.

### GetSettingName

`func (o *AddFileTemplateRequestContainerTemplate) GetSettingName() string`

GetSettingName returns the SettingName field if non-nil, zero value otherwise.

### GetSettingNameOk

`func (o *AddFileTemplateRequestContainerTemplate) GetSettingNameOk() (*string, bool)`

GetSettingNameOk returns a tuple with the SettingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingName

`func (o *AddFileTemplateRequestContainerTemplate) SetSettingName(v string)`

SetSettingName sets SettingName field to given value.

### HasSettingName

`func (o *AddFileTemplateRequestContainerTemplate) HasSettingName() bool`

HasSettingName returns a boolean if a field has been set.

### GetSettingCategory

`func (o *AddFileTemplateRequestContainerTemplate) GetSettingCategory() string`

GetSettingCategory returns the SettingCategory field if non-nil, zero value otherwise.

### GetSettingCategoryOk

`func (o *AddFileTemplateRequestContainerTemplate) GetSettingCategoryOk() (*string, bool)`

GetSettingCategoryOk returns a tuple with the SettingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingCategory

`func (o *AddFileTemplateRequestContainerTemplate) SetSettingCategory(v string)`

SetSettingCategory sets SettingCategory field to given value.

### HasSettingCategory

`func (o *AddFileTemplateRequestContainerTemplate) HasSettingCategory() bool`

HasSettingCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


