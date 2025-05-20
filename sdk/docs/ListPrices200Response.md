# ListPrices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prices** | Pointer to [**[]ListPrices200ResponseAllOfPricesInner**](ListPrices200ResponseAllOfPricesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListPrices200Response

`func NewListPrices200Response() *ListPrices200Response`

NewListPrices200Response instantiates a new ListPrices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPrices200ResponseWithDefaults

`func NewListPrices200ResponseWithDefaults() *ListPrices200Response`

NewListPrices200ResponseWithDefaults instantiates a new ListPrices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrices

`func (o *ListPrices200Response) GetPrices() []ListPrices200ResponseAllOfPricesInner`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ListPrices200Response) GetPricesOk() (*[]ListPrices200ResponseAllOfPricesInner, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ListPrices200Response) SetPrices(v []ListPrices200ResponseAllOfPricesInner)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *ListPrices200Response) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetMeta

`func (o *ListPrices200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPrices200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPrices200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPrices200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


