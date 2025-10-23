# GetBudgets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**GetBudgets200ResponseAllOfBudget**](GetBudgets200ResponseAllOfBudget.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetBudgets200Response

`func NewGetBudgets200Response() *GetBudgets200Response`

NewGetBudgets200Response instantiates a new GetBudgets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBudgets200ResponseWithDefaults

`func NewGetBudgets200ResponseWithDefaults() *GetBudgets200Response`

NewGetBudgets200ResponseWithDefaults instantiates a new GetBudgets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *GetBudgets200Response) GetBudget() GetBudgets200ResponseAllOfBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *GetBudgets200Response) GetBudgetOk() (*GetBudgets200ResponseAllOfBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *GetBudgets200Response) SetBudget(v GetBudgets200ResponseAllOfBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *GetBudgets200Response) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetSuccess

`func (o *GetBudgets200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetBudgets200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetBudgets200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetBudgets200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


