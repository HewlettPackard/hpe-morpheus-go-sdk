# ListPoliciesCloud200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | Pointer to [**[]ListPoliciesCloud200ResponseAllOfPoliciesInner**](ListPoliciesCloud200ResponseAllOfPoliciesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListPoliciesCloud200Response

`func NewListPoliciesCloud200Response() *ListPoliciesCloud200Response`

NewListPoliciesCloud200Response instantiates a new ListPoliciesCloud200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPoliciesCloud200ResponseWithDefaults

`func NewListPoliciesCloud200ResponseWithDefaults() *ListPoliciesCloud200Response`

NewListPoliciesCloud200ResponseWithDefaults instantiates a new ListPoliciesCloud200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *ListPoliciesCloud200Response) GetPolicies() []ListPoliciesCloud200ResponseAllOfPoliciesInner`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ListPoliciesCloud200Response) GetPoliciesOk() (*[]ListPoliciesCloud200ResponseAllOfPoliciesInner, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ListPoliciesCloud200Response) SetPolicies(v []ListPoliciesCloud200ResponseAllOfPoliciesInner)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ListPoliciesCloud200Response) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetMeta

`func (o *ListPoliciesCloud200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPoliciesCloud200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPoliciesCloud200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPoliciesCloud200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


