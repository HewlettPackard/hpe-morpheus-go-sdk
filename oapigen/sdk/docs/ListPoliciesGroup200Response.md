# ListPoliciesGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]ListPoliciesGroup200ResponseAllOfPoliciesInner**](ListPoliciesGroup200ResponseAllOfPoliciesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListPoliciesGroup200Response

`func NewListPoliciesGroup200Response() *ListPoliciesGroup200Response`

NewListPoliciesGroup200Response instantiates a new ListPoliciesGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetPolicies

`func (o *ListPoliciesGroup200Response) GetPolicies() []ListPoliciesGroup200ResponseAllOfPoliciesInner`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ListPoliciesGroup200Response) GetPoliciesOk() (*[]ListPoliciesGroup200ResponseAllOfPoliciesInner, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ListPoliciesGroup200Response) SetPolicies(v []ListPoliciesGroup200ResponseAllOfPoliciesInner)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ListPoliciesGroup200Response) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetMeta

`func (o *ListPoliciesGroup200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPoliciesGroup200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPoliciesGroup200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPoliciesGroup200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


