# ListClusterStatefulSets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statefulsets** | Pointer to [**[]ListClusterPods200ResponseAllOfPodsInner**](ListClusterPods200ResponseAllOfPodsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterStatefulSets200Response

`func NewListClusterStatefulSets200Response() *ListClusterStatefulSets200Response`

NewListClusterStatefulSets200Response instantiates a new ListClusterStatefulSets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterStatefulSets200ResponseWithDefaults

`func NewListClusterStatefulSets200ResponseWithDefaults() *ListClusterStatefulSets200Response`

NewListClusterStatefulSets200ResponseWithDefaults instantiates a new ListClusterStatefulSets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatefulsets

`func (o *ListClusterStatefulSets200Response) GetStatefulsets() []ListClusterPods200ResponseAllOfPodsInner`

GetStatefulsets returns the Statefulsets field if non-nil, zero value otherwise.

### GetStatefulsetsOk

`func (o *ListClusterStatefulSets200Response) GetStatefulsetsOk() (*[]ListClusterPods200ResponseAllOfPodsInner, bool)`

GetStatefulsetsOk returns a tuple with the Statefulsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulsets

`func (o *ListClusterStatefulSets200Response) SetStatefulsets(v []ListClusterPods200ResponseAllOfPodsInner)`

SetStatefulsets sets Statefulsets field to given value.

### HasStatefulsets

`func (o *ListClusterStatefulSets200Response) HasStatefulsets() bool`

HasStatefulsets returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterStatefulSets200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterStatefulSets200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterStatefulSets200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterStatefulSets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


