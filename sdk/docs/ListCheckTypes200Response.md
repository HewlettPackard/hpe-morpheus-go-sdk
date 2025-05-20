# ListCheckTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckTypes** | Pointer to [**[]ListCheckTypes200ResponseAllOfCheckTypesInner**](ListCheckTypes200ResponseAllOfCheckTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCheckTypes200Response

`func NewListCheckTypes200Response() *ListCheckTypes200Response`

NewListCheckTypes200Response instantiates a new ListCheckTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCheckTypes200ResponseWithDefaults

`func NewListCheckTypes200ResponseWithDefaults() *ListCheckTypes200Response`

NewListCheckTypes200ResponseWithDefaults instantiates a new ListCheckTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckTypes

`func (o *ListCheckTypes200Response) GetCheckTypes() []ListCheckTypes200ResponseAllOfCheckTypesInner`

GetCheckTypes returns the CheckTypes field if non-nil, zero value otherwise.

### GetCheckTypesOk

`func (o *ListCheckTypes200Response) GetCheckTypesOk() (*[]ListCheckTypes200ResponseAllOfCheckTypesInner, bool)`

GetCheckTypesOk returns a tuple with the CheckTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckTypes

`func (o *ListCheckTypes200Response) SetCheckTypes(v []ListCheckTypes200ResponseAllOfCheckTypesInner)`

SetCheckTypes sets CheckTypes field to given value.

### HasCheckTypes

`func (o *ListCheckTypes200Response) HasCheckTypes() bool`

HasCheckTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListCheckTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCheckTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCheckTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCheckTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


