# ListClouds200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zones** | Pointer to [**[]ListClouds200ResponseAllOfZonesInner**](ListClouds200ResponseAllOfZonesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClouds200Response

`func NewListClouds200Response() *ListClouds200Response`

NewListClouds200Response instantiates a new ListClouds200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClouds200ResponseWithDefaults

`func NewListClouds200ResponseWithDefaults() *ListClouds200Response`

NewListClouds200ResponseWithDefaults instantiates a new ListClouds200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZones

`func (o *ListClouds200Response) GetZones() []ListClouds200ResponseAllOfZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ListClouds200Response) GetZonesOk() (*[]ListClouds200ResponseAllOfZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ListClouds200Response) SetZones(v []ListClouds200ResponseAllOfZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ListClouds200Response) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetMeta

`func (o *ListClouds200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClouds200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClouds200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClouds200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


