# UpdatePricesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | [**UpdatePricesRequestPrice**](UpdatePricesRequestPrice.md) |  | 

## Methods

### NewUpdatePricesRequest

`func NewUpdatePricesRequest(price UpdatePricesRequestPrice, ) *UpdatePricesRequest`

NewUpdatePricesRequest instantiates a new UpdatePricesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePricesRequestWithDefaults

`func NewUpdatePricesRequestWithDefaults() *UpdatePricesRequest`

NewUpdatePricesRequestWithDefaults instantiates a new UpdatePricesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *UpdatePricesRequest) GetPrice() UpdatePricesRequestPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *UpdatePricesRequest) GetPriceOk() (*UpdatePricesRequestPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *UpdatePricesRequest) SetPrice(v UpdatePricesRequestPrice)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


