# ListOsTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsTypes** | Pointer to [**[]ListOsTypes200ResponseAllOfOsTypesInner**](ListOsTypes200ResponseAllOfOsTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListOsTypes200Response

`func NewListOsTypes200Response() *ListOsTypes200Response`

NewListOsTypes200Response instantiates a new ListOsTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOsTypes200ResponseWithDefaults

`func NewListOsTypes200ResponseWithDefaults() *ListOsTypes200Response`

NewListOsTypes200ResponseWithDefaults instantiates a new ListOsTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsTypes

`func (o *ListOsTypes200Response) GetOsTypes() []ListOsTypes200ResponseAllOfOsTypesInner`

GetOsTypes returns the OsTypes field if non-nil, zero value otherwise.

### GetOsTypesOk

`func (o *ListOsTypes200Response) GetOsTypesOk() (*[]ListOsTypes200ResponseAllOfOsTypesInner, bool)`

GetOsTypesOk returns a tuple with the OsTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTypes

`func (o *ListOsTypes200Response) SetOsTypes(v []ListOsTypes200ResponseAllOfOsTypesInner)`

SetOsTypes sets OsTypes field to given value.

### HasOsTypes

`func (o *ListOsTypes200Response) HasOsTypes() bool`

HasOsTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListOsTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListOsTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListOsTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListOsTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


