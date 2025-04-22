# AddPrices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to [**ListPrices200ResponseAllOfPricesInner**](ListPrices200ResponseAllOfPricesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddPrices200Response

`func NewAddPrices200Response() *AddPrices200Response`

NewAddPrices200Response instantiates a new AddPrices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPrices200ResponseWithDefaults

`func NewAddPrices200ResponseWithDefaults() *AddPrices200Response`

NewAddPrices200ResponseWithDefaults instantiates a new AddPrices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *AddPrices200Response) GetPrice() ListPrices200ResponseAllOfPricesInner`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AddPrices200Response) GetPriceOk() (*ListPrices200ResponseAllOfPricesInner, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AddPrices200Response) SetPrice(v ListPrices200ResponseAllOfPricesInner)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AddPrices200Response) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSuccess

`func (o *AddPrices200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddPrices200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddPrices200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddPrices200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


