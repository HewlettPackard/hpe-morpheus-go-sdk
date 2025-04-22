# ListGuidances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discoveries** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInner**](ListGuidances200ResponseAllOfDiscoveriesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListGuidances200Response

`func NewListGuidances200Response() *ListGuidances200Response`

NewListGuidances200Response instantiates a new ListGuidances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGuidances200ResponseWithDefaults

`func NewListGuidances200ResponseWithDefaults() *ListGuidances200Response`

NewListGuidances200ResponseWithDefaults instantiates a new ListGuidances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveries

`func (o *ListGuidances200Response) GetDiscoveries() []ListGuidances200ResponseAllOfDiscoveriesInner`

GetDiscoveries returns the Discoveries field if non-nil, zero value otherwise.

### GetDiscoveriesOk

`func (o *ListGuidances200Response) GetDiscoveriesOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInner, bool)`

GetDiscoveriesOk returns a tuple with the Discoveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveries

`func (o *ListGuidances200Response) SetDiscoveries(v []ListGuidances200ResponseAllOfDiscoveriesInner)`

SetDiscoveries sets Discoveries field to given value.

### HasDiscoveries

`func (o *ListGuidances200Response) HasDiscoveries() bool`

HasDiscoveries returns a boolean if a field has been set.

### GetMeta

`func (o *ListGuidances200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListGuidances200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListGuidances200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListGuidances200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


