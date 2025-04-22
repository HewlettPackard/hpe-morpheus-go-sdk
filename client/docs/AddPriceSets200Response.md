# AddPriceSets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**ListPriceSets200ResponseAllOfPriceSetsInner**](ListPriceSets200ResponseAllOfPriceSetsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddPriceSets200Response

`func NewAddPriceSets200Response() *AddPriceSets200Response`

NewAddPriceSets200Response instantiates a new AddPriceSets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPriceSets200ResponseWithDefaults

`func NewAddPriceSets200ResponseWithDefaults() *AddPriceSets200Response`

NewAddPriceSets200ResponseWithDefaults instantiates a new AddPriceSets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *AddPriceSets200Response) GetBudget() ListPriceSets200ResponseAllOfPriceSetsInner`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AddPriceSets200Response) GetBudgetOk() (*ListPriceSets200ResponseAllOfPriceSetsInner, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AddPriceSets200Response) SetBudget(v ListPriceSets200ResponseAllOfPriceSetsInner)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AddPriceSets200Response) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetSuccess

`func (o *AddPriceSets200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddPriceSets200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddPriceSets200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddPriceSets200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


