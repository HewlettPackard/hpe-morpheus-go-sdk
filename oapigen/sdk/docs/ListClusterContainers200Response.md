# ListClusterContainers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**[]ListClusterContainers200ResponseAllOfContainersInner**](ListClusterContainers200ResponseAllOfContainersInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterContainers200Response

`func NewListClusterContainers200Response() *ListClusterContainers200Response`

NewListClusterContainers200Response instantiates a new ListClusterContainers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterContainers200ResponseWithDefaults

`func NewListClusterContainers200ResponseWithDefaults() *ListClusterContainers200Response`

NewListClusterContainers200ResponseWithDefaults instantiates a new ListClusterContainers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *ListClusterContainers200Response) GetContainers() []ListClusterContainers200ResponseAllOfContainersInner`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *ListClusterContainers200Response) GetContainersOk() (*[]ListClusterContainers200ResponseAllOfContainersInner, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *ListClusterContainers200Response) SetContainers(v []ListClusterContainers200ResponseAllOfContainersInner)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *ListClusterContainers200Response) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterContainers200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterContainers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterContainers200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterContainers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


