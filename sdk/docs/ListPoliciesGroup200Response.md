# ListPoliciesGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]AddPolicies200ResponseAllOfPolicy**](AddPolicies200ResponseAllOfPolicy.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListPoliciesGroup200Response

`func NewListPoliciesGroup200Response() *ListPoliciesGroup200Response`

NewListPoliciesGroup200Response instantiates a new ListPoliciesGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPoliciesGroup200ResponseWithDefaults

`func NewListPoliciesGroup200ResponseWithDefaults() *ListPoliciesGroup200Response`

NewListPoliciesGroup200ResponseWithDefaults instantiates a new ListPoliciesGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *ListPoliciesGroup200Response) GetPolicies() []AddPolicies200ResponseAllOfPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ListPoliciesGroup200Response) GetPoliciesOk() (*[]AddPolicies200ResponseAllOfPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ListPoliciesGroup200Response) SetPolicies(v []AddPolicies200ResponseAllOfPolicy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ListPoliciesGroup200Response) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetMeta

`func (o *ListPoliciesGroup200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPoliciesGroup200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPoliciesGroup200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPoliciesGroup200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


