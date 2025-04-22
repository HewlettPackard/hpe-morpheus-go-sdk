# ListInstanceTypesProvisioning200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTypes** | Pointer to [**[]ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner**](ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListInstanceTypesProvisioning200Response

`func NewListInstanceTypesProvisioning200Response() *ListInstanceTypesProvisioning200Response`

NewListInstanceTypesProvisioning200Response instantiates a new ListInstanceTypesProvisioning200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceTypesProvisioning200ResponseWithDefaults

`func NewListInstanceTypesProvisioning200ResponseWithDefaults() *ListInstanceTypesProvisioning200Response`

NewListInstanceTypesProvisioning200ResponseWithDefaults instantiates a new ListInstanceTypesProvisioning200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTypes

`func (o *ListInstanceTypesProvisioning200Response) GetInstanceTypes() []ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *ListInstanceTypesProvisioning200Response) GetInstanceTypesOk() (*[]ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *ListInstanceTypesProvisioning200Response) SetInstanceTypes(v []ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *ListInstanceTypesProvisioning200Response) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListInstanceTypesProvisioning200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInstanceTypesProvisioning200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInstanceTypesProvisioning200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInstanceTypesProvisioning200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


