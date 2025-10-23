# ListTasks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tasks** | Pointer to [**[]ListTasks200ResponseAllOfTasksInner**](ListTasks200ResponseAllOfTasksInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListTasks200Response

`func NewListTasks200Response() *ListTasks200Response`

NewListTasks200Response instantiates a new ListTasks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTasks200ResponseWithDefaults

`func NewListTasks200ResponseWithDefaults() *ListTasks200Response`

NewListTasks200ResponseWithDefaults instantiates a new ListTasks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTasks

`func (o *ListTasks200Response) GetTasks() []ListTasks200ResponseAllOfTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ListTasks200Response) GetTasksOk() (*[]ListTasks200ResponseAllOfTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ListTasks200Response) SetTasks(v []ListTasks200ResponseAllOfTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ListTasks200Response) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetMeta

`func (o *ListTasks200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTasks200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTasks200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTasks200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


