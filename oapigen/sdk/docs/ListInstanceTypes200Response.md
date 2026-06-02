# ListInstanceTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTypes** | Pointer to [**[]ListInstanceTypes200ResponseAllOfInstanceTypesInner**](ListInstanceTypes200ResponseAllOfInstanceTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListInstanceTypes200ResponseAllOfMeta**](ListInstanceTypes200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListInstanceTypes200Response

`func NewListInstanceTypes200Response() *ListInstanceTypes200Response`

NewListInstanceTypes200Response instantiates a new ListInstanceTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetInstanceTypes

`func (o *ListInstanceTypes200Response) GetInstanceTypes() []ListInstanceTypes200ResponseAllOfInstanceTypesInner`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *ListInstanceTypes200Response) GetInstanceTypesOk() (*[]ListInstanceTypes200ResponseAllOfInstanceTypesInner, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *ListInstanceTypes200Response) SetInstanceTypes(v []ListInstanceTypes200ResponseAllOfInstanceTypesInner)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *ListInstanceTypes200Response) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListInstanceTypes200Response) GetMeta() ListInstanceTypes200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInstanceTypes200Response) GetMetaOk() (*ListInstanceTypes200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInstanceTypes200Response) SetMeta(v ListInstanceTypes200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInstanceTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


