# ListVirtualImageLocations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | Pointer to [**[]ListVirtualImageLocations200ResponseAllOfLocationsInner**](ListVirtualImageLocations200ResponseAllOfLocationsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListVirtualImageLocations200Response

`func NewListVirtualImageLocations200Response() *ListVirtualImageLocations200Response`

NewListVirtualImageLocations200Response instantiates a new ListVirtualImageLocations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVirtualImageLocations200ResponseWithDefaults

`func NewListVirtualImageLocations200ResponseWithDefaults() *ListVirtualImageLocations200Response`

NewListVirtualImageLocations200ResponseWithDefaults instantiates a new ListVirtualImageLocations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocations

`func (o *ListVirtualImageLocations200Response) GetLocations() []ListVirtualImageLocations200ResponseAllOfLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ListVirtualImageLocations200Response) GetLocationsOk() (*[]ListVirtualImageLocations200ResponseAllOfLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ListVirtualImageLocations200Response) SetLocations(v []ListVirtualImageLocations200ResponseAllOfLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ListVirtualImageLocations200Response) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetMeta

`func (o *ListVirtualImageLocations200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVirtualImageLocations200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVirtualImageLocations200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVirtualImageLocations200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


