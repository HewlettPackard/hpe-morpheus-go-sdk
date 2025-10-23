# PolicyCloudUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the policy | [optional] 
**Description** | Pointer to **string** | A description for the policy | [optional] 
**PolicyType** | Pointer to [**UpdatePoliciesCloudRequestPolicyPolicyType**](UpdatePoliciesCloudRequestPolicyPolicyType.md) |  | [optional] 

## Methods

### NewPolicyCloudUpdate

`func NewPolicyCloudUpdate() *PolicyCloudUpdate`

NewPolicyCloudUpdate instantiates a new PolicyCloudUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCloudUpdateWithDefaults

`func NewPolicyCloudUpdateWithDefaults() *PolicyCloudUpdate`

NewPolicyCloudUpdateWithDefaults instantiates a new PolicyCloudUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyCloudUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyCloudUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyCloudUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyCloudUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyCloudUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyCloudUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyCloudUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyCloudUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyType

`func (o *PolicyCloudUpdate) GetPolicyType() UpdatePoliciesCloudRequestPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *PolicyCloudUpdate) GetPolicyTypeOk() (*UpdatePoliciesCloudRequestPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *PolicyCloudUpdate) SetPolicyType(v UpdatePoliciesCloudRequestPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *PolicyCloudUpdate) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


