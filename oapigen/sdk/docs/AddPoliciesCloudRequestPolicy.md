# AddPoliciesCloudRequestPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the policy | 
**Description** | Pointer to **string** | A description for the policy | [optional] 
**PolicyType** | [**AddPoliciesCloudRequestPolicyPolicyType**](AddPoliciesCloudRequestPolicyPolicyType.md) |  | 

## Methods

### NewAddPoliciesCloudRequestPolicy

`func NewAddPoliciesCloudRequestPolicy(name string, policyType AddPoliciesCloudRequestPolicyPolicyType, ) *AddPoliciesCloudRequestPolicy`

NewAddPoliciesCloudRequestPolicy instantiates a new AddPoliciesCloudRequestPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPoliciesCloudRequestPolicyWithDefaults

`func NewAddPoliciesCloudRequestPolicyWithDefaults() *AddPoliciesCloudRequestPolicy`

NewAddPoliciesCloudRequestPolicyWithDefaults instantiates a new AddPoliciesCloudRequestPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddPoliciesCloudRequestPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddPoliciesCloudRequestPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddPoliciesCloudRequestPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddPoliciesCloudRequestPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPoliciesCloudRequestPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPoliciesCloudRequestPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPoliciesCloudRequestPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyType

`func (o *AddPoliciesCloudRequestPolicy) GetPolicyType() AddPoliciesCloudRequestPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *AddPoliciesCloudRequestPolicy) GetPolicyTypeOk() (*AddPoliciesCloudRequestPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *AddPoliciesCloudRequestPolicy) SetPolicyType(v AddPoliciesCloudRequestPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


