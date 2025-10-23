# ListClusterPods200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pods** | Pointer to [**[]ListClusterPods200ResponseAllOfPodsInner**](ListClusterPods200ResponseAllOfPodsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterPods200Response

`func NewListClusterPods200Response() *ListClusterPods200Response`

NewListClusterPods200Response instantiates a new ListClusterPods200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterPods200ResponseWithDefaults

`func NewListClusterPods200ResponseWithDefaults() *ListClusterPods200Response`

NewListClusterPods200ResponseWithDefaults instantiates a new ListClusterPods200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPods

`func (o *ListClusterPods200Response) GetPods() []ListClusterPods200ResponseAllOfPodsInner`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *ListClusterPods200Response) GetPodsOk() (*[]ListClusterPods200ResponseAllOfPodsInner, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *ListClusterPods200Response) SetPods(v []ListClusterPods200ResponseAllOfPodsInner)`

SetPods sets Pods field to given value.

### HasPods

`func (o *ListClusterPods200Response) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterPods200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterPods200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterPods200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterPods200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


