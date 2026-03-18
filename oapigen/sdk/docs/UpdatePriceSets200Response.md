# UpdatePriceSets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**UpdatePriceSets200ResponseAllOfBudget**](UpdatePriceSets200ResponseAllOfBudget.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdatePriceSets200Response

`func NewUpdatePriceSets200Response() *UpdatePriceSets200Response`

NewUpdatePriceSets200Response instantiates a new UpdatePriceSets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePriceSets200ResponseWithDefaults

`func NewUpdatePriceSets200ResponseWithDefaults() *UpdatePriceSets200Response`

NewUpdatePriceSets200ResponseWithDefaults instantiates a new UpdatePriceSets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *UpdatePriceSets200Response) GetBudget() UpdatePriceSets200ResponseAllOfBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *UpdatePriceSets200Response) GetBudgetOk() (*UpdatePriceSets200ResponseAllOfBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *UpdatePriceSets200Response) SetBudget(v UpdatePriceSets200ResponseAllOfBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *UpdatePriceSets200Response) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdatePriceSets200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdatePriceSets200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdatePriceSets200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdatePriceSets200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


