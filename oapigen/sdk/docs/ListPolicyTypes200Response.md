# ListPolicyTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyTypes** | Pointer to [**[]ListPolicyTypes200ResponseAllOfPolicyTypesInner**](ListPolicyTypes200ResponseAllOfPolicyTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListPolicyTypes200Response

`func NewListPolicyTypes200Response() *ListPolicyTypes200Response`

NewListPolicyTypes200Response instantiates a new ListPolicyTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPolicyTypes200ResponseWithDefaults

`func NewListPolicyTypes200ResponseWithDefaults() *ListPolicyTypes200Response`

NewListPolicyTypes200ResponseWithDefaults instantiates a new ListPolicyTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyTypes

`func (o *ListPolicyTypes200Response) GetPolicyTypes() []ListPolicyTypes200ResponseAllOfPolicyTypesInner`

GetPolicyTypes returns the PolicyTypes field if non-nil, zero value otherwise.

### GetPolicyTypesOk

`func (o *ListPolicyTypes200Response) GetPolicyTypesOk() (*[]ListPolicyTypes200ResponseAllOfPolicyTypesInner, bool)`

GetPolicyTypesOk returns a tuple with the PolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTypes

`func (o *ListPolicyTypes200Response) SetPolicyTypes(v []ListPolicyTypes200ResponseAllOfPolicyTypesInner)`

SetPolicyTypes sets PolicyTypes field to given value.

### HasPolicyTypes

`func (o *ListPolicyTypes200Response) HasPolicyTypes() bool`

HasPolicyTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListPolicyTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPolicyTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPolicyTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPolicyTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


