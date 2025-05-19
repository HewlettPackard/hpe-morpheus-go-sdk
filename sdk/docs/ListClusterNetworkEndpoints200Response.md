# ListClusterNetworkEndpoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner**](ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterNetworkEndpoints200Response

`func NewListClusterNetworkEndpoints200Response() *ListClusterNetworkEndpoints200Response`

NewListClusterNetworkEndpoints200Response instantiates a new ListClusterNetworkEndpoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterNetworkEndpoints200ResponseWithDefaults

`func NewListClusterNetworkEndpoints200ResponseWithDefaults() *ListClusterNetworkEndpoints200Response`

NewListClusterNetworkEndpoints200ResponseWithDefaults instantiates a new ListClusterNetworkEndpoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *ListClusterNetworkEndpoints200Response) GetEndpoints() []ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ListClusterNetworkEndpoints200Response) GetEndpointsOk() (*[]ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ListClusterNetworkEndpoints200Response) SetEndpoints(v []ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ListClusterNetworkEndpoints200Response) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterNetworkEndpoints200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterNetworkEndpoints200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterNetworkEndpoints200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterNetworkEndpoints200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


