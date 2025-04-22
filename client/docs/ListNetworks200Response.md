# ListNetworks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to [**[]ListNetworks200ResponseAllOfNetworksInner**](ListNetworks200ResponseAllOfNetworksInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListNetworks200Response

`func NewListNetworks200Response() *ListNetworks200Response`

NewListNetworks200Response instantiates a new ListNetworks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworks200ResponseWithDefaults

`func NewListNetworks200ResponseWithDefaults() *ListNetworks200Response`

NewListNetworks200ResponseWithDefaults instantiates a new ListNetworks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *ListNetworks200Response) GetNetworks() []ListNetworks200ResponseAllOfNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ListNetworks200Response) GetNetworksOk() (*[]ListNetworks200ResponseAllOfNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ListNetworks200Response) SetNetworks(v []ListNetworks200ResponseAllOfNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ListNetworks200Response) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetMeta

`func (o *ListNetworks200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNetworks200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNetworks200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNetworks200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


