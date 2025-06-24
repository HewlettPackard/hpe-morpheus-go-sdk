# ListTasks200ResponseAllOfTasksInnerAnyOfFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**SourceType** | Pointer to **NullableString** |  | [optional] 
**ContentRef** | Pointer to **NullableString** |  | [optional] 
**ContentPath** | Pointer to **NullableString** |  | [optional] 
**Repository** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Content** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListTasks200ResponseAllOfTasksInnerAnyOfFile

`func NewListTasks200ResponseAllOfTasksInnerAnyOfFile() *ListTasks200ResponseAllOfTasksInnerAnyOfFile`

NewListTasks200ResponseAllOfTasksInnerAnyOfFile instantiates a new ListTasks200ResponseAllOfTasksInnerAnyOfFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTasks200ResponseAllOfTasksInnerAnyOfFileWithDefaults

`func NewListTasks200ResponseAllOfTasksInnerAnyOfFileWithDefaults() *ListTasks200ResponseAllOfTasksInnerAnyOfFile`

NewListTasks200ResponseAllOfTasksInnerAnyOfFileWithDefaults instantiates a new ListTasks200ResponseAllOfTasksInnerAnyOfFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSourceType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetContentRef

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetContentRef() string`

GetContentRef returns the ContentRef field if non-nil, zero value otherwise.

### GetContentRefOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetContentRefOk() (*string, bool)`

GetContentRefOk returns a tuple with the ContentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentRef

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetContentRef(v string)`

SetContentRef sets ContentRef field to given value.

### HasContentRef

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) HasContentRef() bool`

HasContentRef returns a boolean if a field has been set.

### SetContentRefNil

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetContentRefNil(b bool)`

 SetContentRefNil sets the value for ContentRef to be an explicit nil

### UnsetContentRef
`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) UnsetContentRef()`

UnsetContentRef ensures that no value is present for ContentRef, not even an explicit nil
### GetContentPath

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### SetContentPathNil

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetContentPathNil(b bool)`

 SetContentPathNil sets the value for ContentPath to be an explicit nil

### UnsetContentPath
`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) UnsetContentPath()`

UnsetContentPath ensures that no value is present for ContentPath, not even an explicit nil
### GetRepository

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetRepository() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetRepositoryOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetRepository(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetContent

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ListTasks200ResponseAllOfTasksInnerAnyOfFile) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


