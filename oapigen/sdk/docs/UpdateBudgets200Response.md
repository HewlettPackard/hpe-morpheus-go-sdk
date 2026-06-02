# UpdateBudgets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**UpdateBudgets200ResponseAllOfBudget**](UpdateBudgets200ResponseAllOfBudget.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateBudgets200Response

`func NewUpdateBudgets200Response() *UpdateBudgets200Response`

NewUpdateBudgets200Response instantiates a new UpdateBudgets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetBudget

`func (o *UpdateBudgets200Response) GetBudget() UpdateBudgets200ResponseAllOfBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *UpdateBudgets200Response) GetBudgetOk() (*UpdateBudgets200ResponseAllOfBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *UpdateBudgets200Response) SetBudget(v UpdateBudgets200ResponseAllOfBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *UpdateBudgets200Response) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateBudgets200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateBudgets200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateBudgets200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateBudgets200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


