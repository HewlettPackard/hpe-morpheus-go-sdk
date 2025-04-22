# ListCloudTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneTypes** | Pointer to [**[]ListCloudTypes200ResponseAllOfZoneTypesInner**](ListCloudTypes200ResponseAllOfZoneTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCloudTypes200Response

`func NewListCloudTypes200Response() *ListCloudTypes200Response`

NewListCloudTypes200Response instantiates a new ListCloudTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudTypes200ResponseWithDefaults

`func NewListCloudTypes200ResponseWithDefaults() *ListCloudTypes200Response`

NewListCloudTypes200ResponseWithDefaults instantiates a new ListCloudTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneTypes

`func (o *ListCloudTypes200Response) GetZoneTypes() []ListCloudTypes200ResponseAllOfZoneTypesInner`

GetZoneTypes returns the ZoneTypes field if non-nil, zero value otherwise.

### GetZoneTypesOk

`func (o *ListCloudTypes200Response) GetZoneTypesOk() (*[]ListCloudTypes200ResponseAllOfZoneTypesInner, bool)`

GetZoneTypesOk returns a tuple with the ZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTypes

`func (o *ListCloudTypes200Response) SetZoneTypes(v []ListCloudTypes200ResponseAllOfZoneTypesInner)`

SetZoneTypes sets ZoneTypes field to given value.

### HasZoneTypes

`func (o *ListCloudTypes200Response) HasZoneTypes() bool`

HasZoneTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListCloudTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCloudTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCloudTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCloudTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


