# GetNetworkSubnets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | Pointer to [**[]GetNetworkSubnets200ResponseAllOfSubnetsInner**](GetNetworkSubnets200ResponseAllOfSubnetsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetNetworkSubnets200Response

`func NewGetNetworkSubnets200Response() *GetNetworkSubnets200Response`

NewGetNetworkSubnets200Response instantiates a new GetNetworkSubnets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSubnets200ResponseWithDefaults

`func NewGetNetworkSubnets200ResponseWithDefaults() *GetNetworkSubnets200Response`

NewGetNetworkSubnets200ResponseWithDefaults instantiates a new GetNetworkSubnets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *GetNetworkSubnets200Response) GetSubnets() []GetNetworkSubnets200ResponseAllOfSubnetsInner`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *GetNetworkSubnets200Response) GetSubnetsOk() (*[]GetNetworkSubnets200ResponseAllOfSubnetsInner, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *GetNetworkSubnets200Response) SetSubnets(v []GetNetworkSubnets200ResponseAllOfSubnetsInner)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *GetNetworkSubnets200Response) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetMeta

`func (o *GetNetworkSubnets200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetNetworkSubnets200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetNetworkSubnets200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetNetworkSubnets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


