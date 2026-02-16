# TaskResponseFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**SourceType** | Pointer to **NullableString** |  | [optional] 
**ContentRef** | Pointer to **NullableString** |  | [optional] 
**ContentPath** | Pointer to **NullableString** |  | [optional] 
**Repository** | Pointer to [**GetTasks200ResponseAllOfTaskFileRepository**](GetTasks200ResponseAllOfTaskFileRepository.md) |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTaskResponseFile

`func NewTaskResponseFile() *TaskResponseFile`

NewTaskResponseFile instantiates a new TaskResponseFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResponseFileWithDefaults

`func NewTaskResponseFileWithDefaults() *TaskResponseFile`

NewTaskResponseFileWithDefaults instantiates a new TaskResponseFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskResponseFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskResponseFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskResponseFile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TaskResponseFile) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaskResponseFile) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaskResponseFile) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSourceType

`func (o *TaskResponseFile) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TaskResponseFile) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TaskResponseFile) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *TaskResponseFile) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *TaskResponseFile) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *TaskResponseFile) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetContentRef

`func (o *TaskResponseFile) GetContentRef() string`

GetContentRef returns the ContentRef field if non-nil, zero value otherwise.

### GetContentRefOk

`func (o *TaskResponseFile) GetContentRefOk() (*string, bool)`

GetContentRefOk returns a tuple with the ContentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRef

`func (o *TaskResponseFile) SetContentRef(v string)`

SetContentRef sets ContentRef field to given value.

### HasContentRef

`func (o *TaskResponseFile) HasContentRef() bool`

HasContentRef returns a boolean if a field has been set.

### SetContentRefNil

`func (o *TaskResponseFile) SetContentRefNil(b bool)`

 SetContentRefNil sets the value for ContentRef to be an explicit nil

### UnsetContentRef
`func (o *TaskResponseFile) UnsetContentRef()`

UnsetContentRef ensures that no value is present for ContentRef, not even an explicit nil
### GetContentPath

`func (o *TaskResponseFile) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *TaskResponseFile) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *TaskResponseFile) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *TaskResponseFile) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### SetContentPathNil

`func (o *TaskResponseFile) SetContentPathNil(b bool)`

 SetContentPathNil sets the value for ContentPath to be an explicit nil

### UnsetContentPath
`func (o *TaskResponseFile) UnsetContentPath()`

UnsetContentPath ensures that no value is present for ContentPath, not even an explicit nil
### GetRepository

`func (o *TaskResponseFile) GetRepository() GetTasks200ResponseAllOfTaskFileRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *TaskResponseFile) GetRepositoryOk() (*GetTasks200ResponseAllOfTaskFileRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *TaskResponseFile) SetRepository(v GetTasks200ResponseAllOfTaskFileRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *TaskResponseFile) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetContent

`func (o *TaskResponseFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *TaskResponseFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *TaskResponseFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *TaskResponseFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *TaskResponseFile) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *TaskResponseFile) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


