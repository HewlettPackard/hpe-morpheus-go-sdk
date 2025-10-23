# ListWorkflows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskSets** | Pointer to [**[]ListWorkflows200ResponseAllOfTaskSetsInner**](ListWorkflows200ResponseAllOfTaskSetsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListWorkflows200Response

`func NewListWorkflows200Response() *ListWorkflows200Response`

NewListWorkflows200Response instantiates a new ListWorkflows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflows200ResponseWithDefaults

`func NewListWorkflows200ResponseWithDefaults() *ListWorkflows200Response`

NewListWorkflows200ResponseWithDefaults instantiates a new ListWorkflows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskSets

`func (o *ListWorkflows200Response) GetTaskSets() []ListWorkflows200ResponseAllOfTaskSetsInner`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *ListWorkflows200Response) GetTaskSetsOk() (*[]ListWorkflows200ResponseAllOfTaskSetsInner, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *ListWorkflows200Response) SetTaskSets(v []ListWorkflows200ResponseAllOfTaskSetsInner)`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *ListWorkflows200Response) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### GetMeta

`func (o *ListWorkflows200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListWorkflows200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListWorkflows200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListWorkflows200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


