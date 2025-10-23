# GetWorkflows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskSet** | Pointer to [**ListWorkflows200ResponseAllOfTaskSetsInner**](ListWorkflows200ResponseAllOfTaskSetsInner.md) |  | [optional] 

## Methods

### NewGetWorkflows200Response

`func NewGetWorkflows200Response() *GetWorkflows200Response`

NewGetWorkflows200Response instantiates a new GetWorkflows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflows200ResponseWithDefaults

`func NewGetWorkflows200ResponseWithDefaults() *GetWorkflows200Response`

NewGetWorkflows200ResponseWithDefaults instantiates a new GetWorkflows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskSet

`func (o *GetWorkflows200Response) GetTaskSet() ListWorkflows200ResponseAllOfTaskSetsInner`

GetTaskSet returns the TaskSet field if non-nil, zero value otherwise.

### GetTaskSetOk

`func (o *GetWorkflows200Response) GetTaskSetOk() (*ListWorkflows200ResponseAllOfTaskSetsInner, bool)`

GetTaskSetOk returns a tuple with the TaskSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSet

`func (o *GetWorkflows200Response) SetTaskSet(v ListWorkflows200ResponseAllOfTaskSetsInner)`

SetTaskSet sets TaskSet field to given value.

### HasTaskSet

`func (o *GetWorkflows200Response) HasTaskSet() bool`

HasTaskSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


