# UpdatePoliciesGroupRequestPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the policy | [optional] 
**Description** | Pointer to **string** | A description for the policy | [optional] 
**PolicyType** | Pointer to [**UpdatePoliciesGroupRequestPolicyPolicyType**](UpdatePoliciesGroupRequestPolicyPolicyType.md) |  | [optional] 

## Methods

### NewUpdatePoliciesGroupRequestPolicy

`func NewUpdatePoliciesGroupRequestPolicy() *UpdatePoliciesGroupRequestPolicy`

NewUpdatePoliciesGroupRequestPolicy instantiates a new UpdatePoliciesGroupRequestPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePoliciesGroupRequestPolicyWithDefaults

`func NewUpdatePoliciesGroupRequestPolicyWithDefaults() *UpdatePoliciesGroupRequestPolicy`

NewUpdatePoliciesGroupRequestPolicyWithDefaults instantiates a new UpdatePoliciesGroupRequestPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePoliciesGroupRequestPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePoliciesGroupRequestPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePoliciesGroupRequestPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePoliciesGroupRequestPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePoliciesGroupRequestPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePoliciesGroupRequestPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePoliciesGroupRequestPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePoliciesGroupRequestPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyType

`func (o *UpdatePoliciesGroupRequestPolicy) GetPolicyType() UpdatePoliciesGroupRequestPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *UpdatePoliciesGroupRequestPolicy) GetPolicyTypeOk() (*UpdatePoliciesGroupRequestPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *UpdatePoliciesGroupRequestPolicy) SetPolicyType(v UpdatePoliciesGroupRequestPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *UpdatePoliciesGroupRequestPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


