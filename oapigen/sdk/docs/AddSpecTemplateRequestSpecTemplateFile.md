# AddSpecTemplateRequestSpecTemplateFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | File Source i.e. local, repository, url. | [default to "local"]
**Content** | Pointer to **string** | File content, the template text. Only required when sourceType is &#x60;local&#x60;. | [optional] 
**ContentPath** | Pointer to **string** | Content Path, the repo file location or url. Required when sourceType is repository or url. | [optional] 
**ContentRef** | Pointer to **string** | Content Ref, the branch/tag. Only used when sourceType is repo. | [optional] 
**Repository** | Pointer to [**AddSpecTemplateRequestSpecTemplateFileRepository**](AddSpecTemplateRequestSpecTemplateFileRepository.md) |  | [optional] 

## Methods

### NewAddSpecTemplateRequestSpecTemplateFile

`func NewAddSpecTemplateRequestSpecTemplateFile(sourceType string, ) *AddSpecTemplateRequestSpecTemplateFile`

NewAddSpecTemplateRequestSpecTemplateFile instantiates a new AddSpecTemplateRequestSpecTemplateFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSpecTemplateRequestSpecTemplateFileWithDefaults

`func NewAddSpecTemplateRequestSpecTemplateFileWithDefaults() *AddSpecTemplateRequestSpecTemplateFile`

NewAddSpecTemplateRequestSpecTemplateFileWithDefaults instantiates a new AddSpecTemplateRequestSpecTemplateFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *AddSpecTemplateRequestSpecTemplateFile) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AddSpecTemplateRequestSpecTemplateFile) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AddSpecTemplateRequestSpecTemplateFile) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetContent

`func (o *AddSpecTemplateRequestSpecTemplateFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AddSpecTemplateRequestSpecTemplateFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AddSpecTemplateRequestSpecTemplateFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AddSpecTemplateRequestSpecTemplateFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentPath

`func (o *AddSpecTemplateRequestSpecTemplateFile) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *AddSpecTemplateRequestSpecTemplateFile) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *AddSpecTemplateRequestSpecTemplateFile) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *AddSpecTemplateRequestSpecTemplateFile) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### GetContentRef

`func (o *AddSpecTemplateRequestSpecTemplateFile) GetContentRef() string`

GetContentRef returns the ContentRef field if non-nil, zero value otherwise.

### GetContentRefOk

`func (o *AddSpecTemplateRequestSpecTemplateFile) GetContentRefOk() (*string, bool)`

GetContentRefOk returns a tuple with the ContentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRef

`func (o *AddSpecTemplateRequestSpecTemplateFile) SetContentRef(v string)`

SetContentRef sets ContentRef field to given value.

### HasContentRef

`func (o *AddSpecTemplateRequestSpecTemplateFile) HasContentRef() bool`

HasContentRef returns a boolean if a field has been set.

### GetRepository

`func (o *AddSpecTemplateRequestSpecTemplateFile) GetRepository() AddSpecTemplateRequestSpecTemplateFileRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *AddSpecTemplateRequestSpecTemplateFile) GetRepositoryOk() (*AddSpecTemplateRequestSpecTemplateFileRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *AddSpecTemplateRequestSpecTemplateFile) SetRepository(v AddSpecTemplateRequestSpecTemplateFileRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *AddSpecTemplateRequestSpecTemplateFile) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


