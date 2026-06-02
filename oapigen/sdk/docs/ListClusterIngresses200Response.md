# ListClusterIngresses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ingresses** | Pointer to [**[]ListClusterIngresses200ResponseAllOfIngressesInner**](ListClusterIngresses200ResponseAllOfIngressesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterIngresses200Response

`func NewListClusterIngresses200Response() *ListClusterIngresses200Response`

NewListClusterIngresses200Response instantiates a new ListClusterIngresses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetIngresses

`func (o *ListClusterIngresses200Response) GetIngresses() []ListClusterIngresses200ResponseAllOfIngressesInner`

GetIngresses returns the Ingresses field if non-nil, zero value otherwise.

### GetIngressesOk

`func (o *ListClusterIngresses200Response) GetIngressesOk() (*[]ListClusterIngresses200ResponseAllOfIngressesInner, bool)`

GetIngressesOk returns a tuple with the Ingresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngresses

`func (o *ListClusterIngresses200Response) SetIngresses(v []ListClusterIngresses200ResponseAllOfIngressesInner)`

SetIngresses sets Ingresses field to given value.

### HasIngresses

`func (o *ListClusterIngresses200Response) HasIngresses() bool`

HasIngresses returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterIngresses200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterIngresses200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterIngresses200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterIngresses200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


