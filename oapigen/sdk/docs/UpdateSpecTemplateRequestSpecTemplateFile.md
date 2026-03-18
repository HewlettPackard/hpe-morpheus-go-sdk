# UpdateSpecTemplateRequestSpecTemplateFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | Pointer to **string** | File Source i.e. local, repository, url. | [optional] [default to "local"]
**Content** | Pointer to **string** | File content, the template text. Only required when sourceType is &#x60;local&#x60;. | [optional] 
**ContentPath** | Pointer to **string** | Content Path, the repo file location or url. Required when sourceType is repository or url. | [optional] 
**ContentRef** | Pointer to **string** | Content Ref, the branch/tag. Only used when sourceType is repo. | [optional] 
**Repository** | Pointer to [**UpdateSpecTemplateRequestSpecTemplateFileRepository**](UpdateSpecTemplateRequestSpecTemplateFileRepository.md) |  | [optional] 

## Methods

### NewUpdateSpecTemplateRequestSpecTemplateFile

`func NewUpdateSpecTemplateRequestSpecTemplateFile() *UpdateSpecTemplateRequestSpecTemplateFile`

NewUpdateSpecTemplateRequestSpecTemplateFile instantiates a new UpdateSpecTemplateRequestSpecTemplateFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSpecTemplateRequestSpecTemplateFileWithDefaults

`func NewUpdateSpecTemplateRequestSpecTemplateFileWithDefaults() *UpdateSpecTemplateRequestSpecTemplateFile`

NewUpdateSpecTemplateRequestSpecTemplateFileWithDefaults instantiates a new UpdateSpecTemplateRequestSpecTemplateFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetContent

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentPath

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### GetContentRef

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) GetContentRef() string`

GetContentRef returns the ContentRef field if non-nil, zero value otherwise.

### GetContentRefOk

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) GetContentRefOk() (*string, bool)`

GetContentRefOk returns a tuple with the ContentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRef

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) SetContentRef(v string)`

SetContentRef sets ContentRef field to given value.

### HasContentRef

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) HasContentRef() bool`

HasContentRef returns a boolean if a field has been set.

### GetRepository

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) GetRepository() UpdateSpecTemplateRequestSpecTemplateFileRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) GetRepositoryOk() (*UpdateSpecTemplateRequestSpecTemplateFileRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) SetRepository(v UpdateSpecTemplateRequestSpecTemplateFileRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *UpdateSpecTemplateRequestSpecTemplateFile) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


