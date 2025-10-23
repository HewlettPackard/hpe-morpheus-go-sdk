# AddPoliciesGroupRequestPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the policy | 
**Description** | Pointer to **string** | A description for the policy | [optional] 
**PolicyType** | [**AddPoliciesGroupRequestPolicyPolicyType**](AddPoliciesGroupRequestPolicyPolicyType.md) |  | 

## Methods

### NewAddPoliciesGroupRequestPolicy

`func NewAddPoliciesGroupRequestPolicy(name string, policyType AddPoliciesGroupRequestPolicyPolicyType, ) *AddPoliciesGroupRequestPolicy`

NewAddPoliciesGroupRequestPolicy instantiates a new AddPoliciesGroupRequestPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPoliciesGroupRequestPolicyWithDefaults

`func NewAddPoliciesGroupRequestPolicyWithDefaults() *AddPoliciesGroupRequestPolicy`

NewAddPoliciesGroupRequestPolicyWithDefaults instantiates a new AddPoliciesGroupRequestPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddPoliciesGroupRequestPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddPoliciesGroupRequestPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddPoliciesGroupRequestPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddPoliciesGroupRequestPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPoliciesGroupRequestPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPoliciesGroupRequestPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPoliciesGroupRequestPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyType

`func (o *AddPoliciesGroupRequestPolicy) GetPolicyType() AddPoliciesGroupRequestPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *AddPoliciesGroupRequestPolicy) GetPolicyTypeOk() (*AddPoliciesGroupRequestPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *AddPoliciesGroupRequestPolicy) SetPolicyType(v AddPoliciesGroupRequestPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


