# ListHostTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerTypes** | Pointer to [**[]ListHostTypes200ResponseAllOfServerTypesInner**](ListHostTypes200ResponseAllOfServerTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListHostTypes200Response

`func NewListHostTypes200Response() *ListHostTypes200Response`

NewListHostTypes200Response instantiates a new ListHostTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHostTypes200ResponseWithDefaults

`func NewListHostTypes200ResponseWithDefaults() *ListHostTypes200Response`

NewListHostTypes200ResponseWithDefaults instantiates a new ListHostTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerTypes

`func (o *ListHostTypes200Response) GetServerTypes() []ListHostTypes200ResponseAllOfServerTypesInner`

GetServerTypes returns the ServerTypes field if non-nil, zero value otherwise.

### GetServerTypesOk

`func (o *ListHostTypes200Response) GetServerTypesOk() (*[]ListHostTypes200ResponseAllOfServerTypesInner, bool)`

GetServerTypesOk returns a tuple with the ServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypes

`func (o *ListHostTypes200Response) SetServerTypes(v []ListHostTypes200ResponseAllOfServerTypesInner)`

SetServerTypes sets ServerTypes field to given value.

### HasServerTypes

`func (o *ListHostTypes200Response) HasServerTypes() bool`

HasServerTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListHostTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListHostTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListHostTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListHostTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


