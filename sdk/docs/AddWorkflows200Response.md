# AddWorkflows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskSet** | Pointer to [**ListWorkflows200ResponseAllOfTaskSetsInner**](ListWorkflows200ResponseAllOfTaskSetsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddWorkflows200Response

`func NewAddWorkflows200Response() *AddWorkflows200Response`

NewAddWorkflows200Response instantiates a new AddWorkflows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddWorkflows200ResponseWithDefaults

`func NewAddWorkflows200ResponseWithDefaults() *AddWorkflows200Response`

NewAddWorkflows200ResponseWithDefaults instantiates a new AddWorkflows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskSet

`func (o *AddWorkflows200Response) GetTaskSet() ListWorkflows200ResponseAllOfTaskSetsInner`

GetTaskSet returns the TaskSet field if non-nil, zero value otherwise.

### GetTaskSetOk

`func (o *AddWorkflows200Response) GetTaskSetOk() (*ListWorkflows200ResponseAllOfTaskSetsInner, bool)`

GetTaskSetOk returns a tuple with the TaskSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSet

`func (o *AddWorkflows200Response) SetTaskSet(v ListWorkflows200ResponseAllOfTaskSetsInner)`

SetTaskSet sets TaskSet field to given value.

### HasTaskSet

`func (o *AddWorkflows200Response) HasTaskSet() bool`

HasTaskSet returns a boolean if a field has been set.

### GetSuccess

`func (o *AddWorkflows200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddWorkflows200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddWorkflows200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddWorkflows200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


