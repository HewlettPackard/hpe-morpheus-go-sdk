# ListProvisionTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvisionTypes** | Pointer to [**[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType**](GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListProvisionTypes200Response

`func NewListProvisionTypes200Response() *ListProvisionTypes200Response`

NewListProvisionTypes200Response instantiates a new ListProvisionTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProvisionTypes200ResponseWithDefaults

`func NewListProvisionTypes200ResponseWithDefaults() *ListProvisionTypes200Response`

NewListProvisionTypes200ResponseWithDefaults instantiates a new ListProvisionTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvisionTypes

`func (o *ListProvisionTypes200Response) GetProvisionTypes() []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType`

GetProvisionTypes returns the ProvisionTypes field if non-nil, zero value otherwise.

### GetProvisionTypesOk

`func (o *ListProvisionTypes200Response) GetProvisionTypesOk() (*[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType, bool)`

GetProvisionTypesOk returns a tuple with the ProvisionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypes

`func (o *ListProvisionTypes200Response) SetProvisionTypes(v []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType)`

SetProvisionTypes sets ProvisionTypes field to given value.

### HasProvisionTypes

`func (o *ListProvisionTypes200Response) HasProvisionTypes() bool`

HasProvisionTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListProvisionTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListProvisionTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListProvisionTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListProvisionTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


