# AddPrices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to [**AddPrices200ResponseAllOfPrice**](AddPrices200ResponseAllOfPrice.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddPrices200Response

`func NewAddPrices200Response() *AddPrices200Response`

NewAddPrices200Response instantiates a new AddPrices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetPrice

`func (o *AddPrices200Response) GetPrice() AddPrices200ResponseAllOfPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AddPrices200Response) GetPriceOk() (*AddPrices200ResponseAllOfPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AddPrices200Response) SetPrice(v AddPrices200ResponseAllOfPrice)`

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


