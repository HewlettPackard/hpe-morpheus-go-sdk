# ListClusterConfigmaps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configmaps** | Pointer to [**[]ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner**](ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterConfigmaps200Response

`func NewListClusterConfigmaps200Response() *ListClusterConfigmaps200Response`

NewListClusterConfigmaps200Response instantiates a new ListClusterConfigmaps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterConfigmaps200ResponseWithDefaults

`func NewListClusterConfigmaps200ResponseWithDefaults() *ListClusterConfigmaps200Response`

NewListClusterConfigmaps200ResponseWithDefaults instantiates a new ListClusterConfigmaps200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigmaps

`func (o *ListClusterConfigmaps200Response) GetConfigmaps() []ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner`

GetConfigmaps returns the Configmaps field if non-nil, zero value otherwise.

### GetConfigmapsOk

`func (o *ListClusterConfigmaps200Response) GetConfigmapsOk() (*[]ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner, bool)`

GetConfigmapsOk returns a tuple with the Configmaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigmaps

`func (o *ListClusterConfigmaps200Response) SetConfigmaps(v []ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner)`

SetConfigmaps sets Configmaps field to given value.

### HasConfigmaps

`func (o *ListClusterConfigmaps200Response) HasConfigmaps() bool`

HasConfigmaps returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterConfigmaps200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterConfigmaps200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterConfigmaps200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterConfigmaps200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


