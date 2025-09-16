# ListSubnets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | Pointer to [**[]ListSubnets200ResponseAllOfSubnetsInner**](ListSubnets200ResponseAllOfSubnetsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListSubnets200Response

`func NewListSubnets200Response() *ListSubnets200Response`

NewListSubnets200Response instantiates a new ListSubnets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubnets200ResponseWithDefaults

`func NewListSubnets200ResponseWithDefaults() *ListSubnets200Response`

NewListSubnets200ResponseWithDefaults instantiates a new ListSubnets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *ListSubnets200Response) GetSubnets() []ListSubnets200ResponseAllOfSubnetsInner`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *ListSubnets200Response) GetSubnetsOk() (*[]ListSubnets200ResponseAllOfSubnetsInner, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *ListSubnets200Response) SetSubnets(v []ListSubnets200ResponseAllOfSubnetsInner)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *ListSubnets200Response) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetMeta

`func (o *ListSubnets200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSubnets200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSubnets200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSubnets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


