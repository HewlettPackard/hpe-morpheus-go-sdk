# ListPolicies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]ListPolicies200ResponseAllOfPoliciesInner**](ListPolicies200ResponseAllOfPoliciesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListPolicies200Response

`func NewListPolicies200Response() *ListPolicies200Response`

NewListPolicies200Response instantiates a new ListPolicies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPolicies200ResponseWithDefaults

`func NewListPolicies200ResponseWithDefaults() *ListPolicies200Response`

NewListPolicies200ResponseWithDefaults instantiates a new ListPolicies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *ListPolicies200Response) GetPolicies() []ListPolicies200ResponseAllOfPoliciesInner`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ListPolicies200Response) GetPoliciesOk() (*[]ListPolicies200ResponseAllOfPoliciesInner, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ListPolicies200Response) SetPolicies(v []ListPolicies200ResponseAllOfPoliciesInner)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ListPolicies200Response) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetMeta

`func (o *ListPolicies200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPolicies200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPolicies200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPolicies200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


