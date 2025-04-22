# ListClusterWorkers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workers** | Pointer to [**[]GetClusterMasters200ResponseMastersInner**](GetClusterMasters200ResponseMastersInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterWorkers200Response

`func NewListClusterWorkers200Response() *ListClusterWorkers200Response`

NewListClusterWorkers200Response instantiates a new ListClusterWorkers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterWorkers200ResponseWithDefaults

`func NewListClusterWorkers200ResponseWithDefaults() *ListClusterWorkers200Response`

NewListClusterWorkers200ResponseWithDefaults instantiates a new ListClusterWorkers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkers

`func (o *ListClusterWorkers200Response) GetWorkers() []GetClusterMasters200ResponseMastersInner`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *ListClusterWorkers200Response) GetWorkersOk() (*[]GetClusterMasters200ResponseMastersInner, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *ListClusterWorkers200Response) SetWorkers(v []GetClusterMasters200ResponseMastersInner)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *ListClusterWorkers200Response) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterWorkers200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterWorkers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterWorkers200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterWorkers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


