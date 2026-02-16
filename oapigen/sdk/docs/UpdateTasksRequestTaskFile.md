# UpdateTasksRequestTaskFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | File Source i.e. &#x60;local&#x60;, &#x60;repository&#x60;, &#x60;url&#x60;. Default is &#x60;local&#x60;. | [default to "local"]
**Content** | Pointer to **string** | File content, the script text. Only required when sourceType is &#x60;local&#x60;. | [optional] 
**ContentPath** | Pointer to **string** | Content Path, the repo file location or url. Required when sourceType is &#x60;repository&#x60; or &#x60;url&#x60;. | [optional] 
**ContentRef** | Pointer to **string** | Content Ref, the branch/tag. Only used when sourceType is &#x60;repository&#x60;. | [optional] 
**Repository** | Pointer to [**UpdateTasksRequestTaskFileRepository**](UpdateTasksRequestTaskFileRepository.md) |  | [optional] 

## Methods

### NewUpdateTasksRequestTaskFile

`func NewUpdateTasksRequestTaskFile(sourceType string, ) *UpdateTasksRequestTaskFile`

NewUpdateTasksRequestTaskFile instantiates a new UpdateTasksRequestTaskFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTasksRequestTaskFileWithDefaults

`func NewUpdateTasksRequestTaskFileWithDefaults() *UpdateTasksRequestTaskFile`

NewUpdateTasksRequestTaskFileWithDefaults instantiates a new UpdateTasksRequestTaskFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *UpdateTasksRequestTaskFile) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *UpdateTasksRequestTaskFile) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *UpdateTasksRequestTaskFile) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetContent

`func (o *UpdateTasksRequestTaskFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateTasksRequestTaskFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateTasksRequestTaskFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdateTasksRequestTaskFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentPath

`func (o *UpdateTasksRequestTaskFile) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *UpdateTasksRequestTaskFile) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *UpdateTasksRequestTaskFile) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *UpdateTasksRequestTaskFile) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### GetContentRef

`func (o *UpdateTasksRequestTaskFile) GetContentRef() string`

GetContentRef returns the ContentRef field if non-nil, zero value otherwise.

### GetContentRefOk

`func (o *UpdateTasksRequestTaskFile) GetContentRefOk() (*string, bool)`

GetContentRefOk returns a tuple with the ContentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRef

`func (o *UpdateTasksRequestTaskFile) SetContentRef(v string)`

SetContentRef sets ContentRef field to given value.

### HasContentRef

`func (o *UpdateTasksRequestTaskFile) HasContentRef() bool`

HasContentRef returns a boolean if a field has been set.

### GetRepository

`func (o *UpdateTasksRequestTaskFile) GetRepository() UpdateTasksRequestTaskFileRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *UpdateTasksRequestTaskFile) GetRepositoryOk() (*UpdateTasksRequestTaskFileRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *UpdateTasksRequestTaskFile) SetRepository(v UpdateTasksRequestTaskFileRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *UpdateTasksRequestTaskFile) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


