# ListClusters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]ListClusters200ResponseAllOfClustersInner**](ListClusters200ResponseAllOfClustersInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusters200Response

`func NewListClusters200Response() *ListClusters200Response`

NewListClusters200Response instantiates a new ListClusters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusters200ResponseWithDefaults

`func NewListClusters200ResponseWithDefaults() *ListClusters200Response`

NewListClusters200ResponseWithDefaults instantiates a new ListClusters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *ListClusters200Response) GetClusters() []ListClusters200ResponseAllOfClustersInner`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *ListClusters200Response) GetClustersOk() (*[]ListClusters200ResponseAllOfClustersInner, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *ListClusters200Response) SetClusters(v []ListClusters200ResponseAllOfClustersInner)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *ListClusters200Response) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusters200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusters200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusters200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusters200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


