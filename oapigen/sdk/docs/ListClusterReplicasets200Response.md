# ListClusterReplicasets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Replicasets** | Pointer to [**[]ListClusterReplicasets200ResponseAllOfReplicasetsInner**](ListClusterReplicasets200ResponseAllOfReplicasetsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterReplicasets200Response

`func NewListClusterReplicasets200Response() *ListClusterReplicasets200Response`

NewListClusterReplicasets200Response instantiates a new ListClusterReplicasets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterReplicasets200ResponseWithDefaults

`func NewListClusterReplicasets200ResponseWithDefaults() *ListClusterReplicasets200Response`

NewListClusterReplicasets200ResponseWithDefaults instantiates a new ListClusterReplicasets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicasets

`func (o *ListClusterReplicasets200Response) GetReplicasets() []ListClusterReplicasets200ResponseAllOfReplicasetsInner`

GetReplicasets returns the Replicasets field if non-nil, zero value otherwise.

### GetReplicasetsOk

`func (o *ListClusterReplicasets200Response) GetReplicasetsOk() (*[]ListClusterReplicasets200ResponseAllOfReplicasetsInner, bool)`

GetReplicasetsOk returns a tuple with the Replicasets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicasets

`func (o *ListClusterReplicasets200Response) SetReplicasets(v []ListClusterReplicasets200ResponseAllOfReplicasetsInner)`

SetReplicasets sets Replicasets field to given value.

### HasReplicasets

`func (o *ListClusterReplicasets200Response) HasReplicasets() bool`

HasReplicasets returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterReplicasets200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterReplicasets200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterReplicasets200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterReplicasets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


