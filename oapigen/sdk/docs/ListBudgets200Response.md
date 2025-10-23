# ListBudgets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budgets** | Pointer to [**[]ListBudgets200ResponseAllOfBudgetsInner**](ListBudgets200ResponseAllOfBudgetsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListBudgets200Response

`func NewListBudgets200Response() *ListBudgets200Response`

NewListBudgets200Response instantiates a new ListBudgets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBudgets200ResponseWithDefaults

`func NewListBudgets200ResponseWithDefaults() *ListBudgets200Response`

NewListBudgets200ResponseWithDefaults instantiates a new ListBudgets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgets

`func (o *ListBudgets200Response) GetBudgets() []ListBudgets200ResponseAllOfBudgetsInner`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *ListBudgets200Response) GetBudgetsOk() (*[]ListBudgets200ResponseAllOfBudgetsInner, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgets

`func (o *ListBudgets200Response) SetBudgets(v []ListBudgets200ResponseAllOfBudgetsInner)`

SetBudgets sets Budgets field to given value.

### HasBudgets

`func (o *ListBudgets200Response) HasBudgets() bool`

HasBudgets returns a boolean if a field has been set.

### GetMeta

`func (o *ListBudgets200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBudgets200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBudgets200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBudgets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


