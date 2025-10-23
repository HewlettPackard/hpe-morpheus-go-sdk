# ListSubnetTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubnetTypes** | Pointer to [**[]ListSubnetTypes200ResponseAllOfSubnetTypesInner**](ListSubnetTypes200ResponseAllOfSubnetTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListSubnetTypes200Response

`func NewListSubnetTypes200Response() *ListSubnetTypes200Response`

NewListSubnetTypes200Response instantiates a new ListSubnetTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubnetTypes200ResponseWithDefaults

`func NewListSubnetTypes200ResponseWithDefaults() *ListSubnetTypes200Response`

NewListSubnetTypes200ResponseWithDefaults instantiates a new ListSubnetTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnetTypes

`func (o *ListSubnetTypes200Response) GetSubnetTypes() []ListSubnetTypes200ResponseAllOfSubnetTypesInner`

GetSubnetTypes returns the SubnetTypes field if non-nil, zero value otherwise.

### GetSubnetTypesOk

`func (o *ListSubnetTypes200Response) GetSubnetTypesOk() (*[]ListSubnetTypes200ResponseAllOfSubnetTypesInner, bool)`

GetSubnetTypesOk returns a tuple with the SubnetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetTypes

`func (o *ListSubnetTypes200Response) SetSubnetTypes(v []ListSubnetTypes200ResponseAllOfSubnetTypesInner)`

SetSubnetTypes sets SubnetTypes field to given value.

### HasSubnetTypes

`func (o *ListSubnetTypes200Response) HasSubnetTypes() bool`

HasSubnetTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListSubnetTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSubnetTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSubnetTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSubnetTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


