# UpdatePrices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to [**UpdatePrices200ResponseAllOfPrice**](UpdatePrices200ResponseAllOfPrice.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdatePrices200Response

`func NewUpdatePrices200Response() *UpdatePrices200Response`

NewUpdatePrices200Response instantiates a new UpdatePrices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePrices200ResponseWithDefaults

`func NewUpdatePrices200ResponseWithDefaults() *UpdatePrices200Response`

NewUpdatePrices200ResponseWithDefaults instantiates a new UpdatePrices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *UpdatePrices200Response) GetPrice() UpdatePrices200ResponseAllOfPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdatePrices200Response) GetPriceOk() (*UpdatePrices200ResponseAllOfPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdatePrices200Response) SetPrice(v UpdatePrices200ResponseAllOfPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *UpdatePrices200Response) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdatePrices200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdatePrices200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdatePrices200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdatePrices200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


