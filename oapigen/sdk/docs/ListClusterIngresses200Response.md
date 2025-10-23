# ListClusterIngresses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingresses** | Pointer to [**[]ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner**](ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterIngresses200Response

`func NewListClusterIngresses200Response() *ListClusterIngresses200Response`

NewListClusterIngresses200Response instantiates a new ListClusterIngresses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterIngresses200ResponseWithDefaults

`func NewListClusterIngresses200ResponseWithDefaults() *ListClusterIngresses200Response`

NewListClusterIngresses200ResponseWithDefaults instantiates a new ListClusterIngresses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngresses

`func (o *ListClusterIngresses200Response) GetIngresses() []ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner`

GetIngresses returns the Ingresses field if non-nil, zero value otherwise.

### GetIngressesOk

`func (o *ListClusterIngresses200Response) GetIngressesOk() (*[]ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner, bool)`

GetIngressesOk returns a tuple with the Ingresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngresses

`func (o *ListClusterIngresses200Response) SetIngresses(v []ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner)`

SetIngresses sets Ingresses field to given value.

### HasIngresses

`func (o *ListClusterIngresses200Response) HasIngresses() bool`

HasIngresses returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterIngresses200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterIngresses200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterIngresses200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterIngresses200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


