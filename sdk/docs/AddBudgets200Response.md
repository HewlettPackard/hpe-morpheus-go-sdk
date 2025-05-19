# AddBudgets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**ListBudgets200ResponseAllOfBudgetsInner**](ListBudgets200ResponseAllOfBudgetsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddBudgets200Response

`func NewAddBudgets200Response() *AddBudgets200Response`

NewAddBudgets200Response instantiates a new AddBudgets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBudgets200ResponseWithDefaults

`func NewAddBudgets200ResponseWithDefaults() *AddBudgets200Response`

NewAddBudgets200ResponseWithDefaults instantiates a new AddBudgets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *AddBudgets200Response) GetBudget() ListBudgets200ResponseAllOfBudgetsInner`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AddBudgets200Response) GetBudgetOk() (*ListBudgets200ResponseAllOfBudgetsInner, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AddBudgets200Response) SetBudget(v ListBudgets200ResponseAllOfBudgetsInner)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AddBudgets200Response) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetSuccess

`func (o *AddBudgets200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddBudgets200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddBudgets200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddBudgets200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


