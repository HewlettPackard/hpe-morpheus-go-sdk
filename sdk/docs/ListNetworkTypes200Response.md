# ListNetworkTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkTypes** | Pointer to [**[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner**](GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListNetworkTypes200Response

`func NewListNetworkTypes200Response() *ListNetworkTypes200Response`

NewListNetworkTypes200Response instantiates a new ListNetworkTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkTypes200ResponseWithDefaults

`func NewListNetworkTypes200ResponseWithDefaults() *ListNetworkTypes200Response`

NewListNetworkTypes200ResponseWithDefaults instantiates a new ListNetworkTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkTypes

`func (o *ListNetworkTypes200Response) GetNetworkTypes() []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner`

GetNetworkTypes returns the NetworkTypes field if non-nil, zero value otherwise.

### GetNetworkTypesOk

`func (o *ListNetworkTypes200Response) GetNetworkTypesOk() (*[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner, bool)`

GetNetworkTypesOk returns a tuple with the NetworkTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTypes

`func (o *ListNetworkTypes200Response) SetNetworkTypes(v []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionTypeNetworkTypesInner)`

SetNetworkTypes sets NetworkTypes field to given value.

### HasNetworkTypes

`func (o *ListNetworkTypes200Response) HasNetworkTypes() bool`

HasNetworkTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListNetworkTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNetworkTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNetworkTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNetworkTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


