# PolicyGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the policy | [optional] 
**Description** | Pointer to **string** | A description for the policy | [optional] 
**PolicyType** | Pointer to [**UpdatePoliciesGroupRequestPolicyPolicyType**](UpdatePoliciesGroupRequestPolicyPolicyType.md) |  | [optional] 

## Methods

### NewPolicyGroupUpdate

`func NewPolicyGroupUpdate() *PolicyGroupUpdate`

NewPolicyGroupUpdate instantiates a new PolicyGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyGroupUpdateWithDefaults

`func NewPolicyGroupUpdateWithDefaults() *PolicyGroupUpdate`

NewPolicyGroupUpdateWithDefaults instantiates a new PolicyGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyGroupUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyGroupUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyGroupUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyGroupUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyGroupUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyGroupUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyGroupUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyGroupUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyType

`func (o *PolicyGroupUpdate) GetPolicyType() UpdatePoliciesGroupRequestPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *PolicyGroupUpdate) GetPolicyTypeOk() (*UpdatePoliciesGroupRequestPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *PolicyGroupUpdate) SetPolicyType(v UpdatePoliciesGroupRequestPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *PolicyGroupUpdate) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


