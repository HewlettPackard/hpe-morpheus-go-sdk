# ListCloudResourcePools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePools** | Pointer to [**[]ListCloudResourcePools200ResponseAllOfResourcePoolsInner**](ListCloudResourcePools200ResponseAllOfResourcePoolsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCloudResourcePools200Response

`func NewListCloudResourcePools200Response() *ListCloudResourcePools200Response`

NewListCloudResourcePools200Response instantiates a new ListCloudResourcePools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudResourcePools200ResponseWithDefaults

`func NewListCloudResourcePools200ResponseWithDefaults() *ListCloudResourcePools200Response`

NewListCloudResourcePools200ResponseWithDefaults instantiates a new ListCloudResourcePools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePools

`func (o *ListCloudResourcePools200Response) GetResourcePools() []ListCloudResourcePools200ResponseAllOfResourcePoolsInner`

GetResourcePools returns the ResourcePools field if non-nil, zero value otherwise.

### GetResourcePoolsOk

`func (o *ListCloudResourcePools200Response) GetResourcePoolsOk() (*[]ListCloudResourcePools200ResponseAllOfResourcePoolsInner, bool)`

GetResourcePoolsOk returns a tuple with the ResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePools

`func (o *ListCloudResourcePools200Response) SetResourcePools(v []ListCloudResourcePools200ResponseAllOfResourcePoolsInner)`

SetResourcePools sets ResourcePools field to given value.

### HasResourcePools

`func (o *ListCloudResourcePools200Response) HasResourcePools() bool`

HasResourcePools returns a boolean if a field has been set.

### GetMeta

`func (o *ListCloudResourcePools200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCloudResourcePools200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCloudResourcePools200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCloudResourcePools200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


