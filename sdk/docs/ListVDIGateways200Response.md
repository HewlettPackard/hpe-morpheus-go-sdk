# ListVDIGateways200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdiGateways** | Pointer to [**[]ListVDIGateways200ResponseAllOfVdiGatewaysInner**](ListVDIGateways200ResponseAllOfVdiGatewaysInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListVDIGateways200Response

`func NewListVDIGateways200Response() *ListVDIGateways200Response`

NewListVDIGateways200Response instantiates a new ListVDIGateways200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVDIGateways200ResponseWithDefaults

`func NewListVDIGateways200ResponseWithDefaults() *ListVDIGateways200Response`

NewListVDIGateways200ResponseWithDefaults instantiates a new ListVDIGateways200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdiGateways

`func (o *ListVDIGateways200Response) GetVdiGateways() []ListVDIGateways200ResponseAllOfVdiGatewaysInner`

GetVdiGateways returns the VdiGateways field if non-nil, zero value otherwise.

### GetVdiGatewaysOk

`func (o *ListVDIGateways200Response) GetVdiGatewaysOk() (*[]ListVDIGateways200ResponseAllOfVdiGatewaysInner, bool)`

GetVdiGatewaysOk returns a tuple with the VdiGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiGateways

`func (o *ListVDIGateways200Response) SetVdiGateways(v []ListVDIGateways200ResponseAllOfVdiGatewaysInner)`

SetVdiGateways sets VdiGateways field to given value.

### HasVdiGateways

`func (o *ListVDIGateways200Response) HasVdiGateways() bool`

HasVdiGateways returns a boolean if a field has been set.

### GetMeta

`func (o *ListVDIGateways200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVDIGateways200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVDIGateways200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVDIGateways200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


