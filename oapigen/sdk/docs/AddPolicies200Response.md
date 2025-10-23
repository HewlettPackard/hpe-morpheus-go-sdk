# AddPolicies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to [**AddPolicies200ResponseAllOfPolicy**](AddPolicies200ResponseAllOfPolicy.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddPolicies200Response

`func NewAddPolicies200Response() *AddPolicies200Response`

NewAddPolicies200Response instantiates a new AddPolicies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPolicies200ResponseWithDefaults

`func NewAddPolicies200ResponseWithDefaults() *AddPolicies200Response`

NewAddPolicies200ResponseWithDefaults instantiates a new AddPolicies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *AddPolicies200Response) GetPolicy() AddPolicies200ResponseAllOfPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AddPolicies200Response) GetPolicyOk() (*AddPolicies200ResponseAllOfPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AddPolicies200Response) SetPolicy(v AddPolicies200ResponseAllOfPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *AddPolicies200Response) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSuccess

`func (o *AddPolicies200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddPolicies200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddPolicies200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddPolicies200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


