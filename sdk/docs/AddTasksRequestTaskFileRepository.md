# AddTasksRequestTaskFileRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Code Repository ID, required for type &#x60;repository&#x60;. Use &#x60;/api/options/codeRepositories&#x60; to see available repositories. | [optional] 

## Methods

### NewAddTasksRequestTaskFileRepository

`func NewAddTasksRequestTaskFileRepository() *AddTasksRequestTaskFileRepository`

NewAddTasksRequestTaskFileRepository instantiates a new AddTasksRequestTaskFileRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTasksRequestTaskFileRepositoryWithDefaults

`func NewAddTasksRequestTaskFileRepositoryWithDefaults() *AddTasksRequestTaskFileRepository`

NewAddTasksRequestTaskFileRepositoryWithDefaults instantiates a new AddTasksRequestTaskFileRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddTasksRequestTaskFileRepository) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddTasksRequestTaskFileRepository) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddTasksRequestTaskFileRepository) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddTasksRequestTaskFileRepository) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


