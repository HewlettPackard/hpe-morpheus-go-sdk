# ListClusterPolicies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner**](ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterPolicies200Response

`func NewListClusterPolicies200Response() *ListClusterPolicies200Response`

NewListClusterPolicies200Response instantiates a new ListClusterPolicies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterPolicies200ResponseWithDefaults

`func NewListClusterPolicies200ResponseWithDefaults() *ListClusterPolicies200Response`

NewListClusterPolicies200ResponseWithDefaults instantiates a new ListClusterPolicies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *ListClusterPolicies200Response) GetPolicies() []ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ListClusterPolicies200Response) GetPoliciesOk() (*[]ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ListClusterPolicies200Response) SetPolicies(v []ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ListClusterPolicies200Response) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterPolicies200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterPolicies200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterPolicies200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterPolicies200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


