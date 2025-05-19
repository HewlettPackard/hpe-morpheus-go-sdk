# ListClusterTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterTypes** | Pointer to [**[]ListClusterTypes200ResponseAllOfClusterTypesInner**](ListClusterTypes200ResponseAllOfClusterTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterTypes200Response

`func NewListClusterTypes200Response() *ListClusterTypes200Response`

NewListClusterTypes200Response instantiates a new ListClusterTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterTypes200ResponseWithDefaults

`func NewListClusterTypes200ResponseWithDefaults() *ListClusterTypes200Response`

NewListClusterTypes200ResponseWithDefaults instantiates a new ListClusterTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterTypes

`func (o *ListClusterTypes200Response) GetClusterTypes() []ListClusterTypes200ResponseAllOfClusterTypesInner`

GetClusterTypes returns the ClusterTypes field if non-nil, zero value otherwise.

### GetClusterTypesOk

`func (o *ListClusterTypes200Response) GetClusterTypesOk() (*[]ListClusterTypes200ResponseAllOfClusterTypesInner, bool)`

GetClusterTypesOk returns a tuple with the ClusterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTypes

`func (o *ListClusterTypes200Response) SetClusterTypes(v []ListClusterTypes200ResponseAllOfClusterTypesInner)`

SetClusterTypes sets ClusterTypes field to given value.

### HasClusterTypes

`func (o *ListClusterTypes200Response) HasClusterTypes() bool`

HasClusterTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


