# PolicyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the policy | 
**Description** | Pointer to **string** | A description for the policy | [optional] 
**PolicyType** | [**AddPoliciesRequestPolicyPolicyType**](AddPoliciesRequestPolicyPolicyType.md) |  | 
**Config** | [**AddPoliciesRequestPolicyConfig**](AddPoliciesRequestPolicyConfig.md) |  | 
**Enabled** | Pointer to **bool** | Set to false to disable | [optional] [default to true]
**RefType** | Pointer to **string** | Scope object type.  If none specified, will default to Global (null) | [optional] 
**RefId** | Pointer to **int64** | Scope object ID (&#x60;group&#x60;,&#x60;cloud&#x60;,&#x60;user&#x60;, etc) | [optional] 
**Accounts** | Pointer to **[]int64** | Array of tenants to scope the policy to | [optional] 
**EachUser** | Pointer to **bool** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; | [optional] 

## Methods

### NewPolicyCreate

`func NewPolicyCreate(name string, policyType AddPoliciesRequestPolicyPolicyType, config AddPoliciesRequestPolicyConfig, ) *PolicyCreate`

NewPolicyCreate instantiates a new PolicyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCreateWithDefaults

`func NewPolicyCreateWithDefaults() *PolicyCreate`

NewPolicyCreateWithDefaults instantiates a new PolicyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PolicyCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPolicyType

`func (o *PolicyCreate) GetPolicyType() AddPoliciesRequestPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *PolicyCreate) GetPolicyTypeOk() (*AddPoliciesRequestPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *PolicyCreate) SetPolicyType(v AddPoliciesRequestPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.


### GetConfig

`func (o *PolicyCreate) GetConfig() AddPoliciesRequestPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PolicyCreate) GetConfigOk() (*AddPoliciesRequestPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PolicyCreate) SetConfig(v AddPoliciesRequestPolicyConfig)`

SetConfig sets Config field to given value.


### GetEnabled

`func (o *PolicyCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PolicyCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PolicyCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PolicyCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefType

`func (o *PolicyCreate) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *PolicyCreate) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *PolicyCreate) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *PolicyCreate) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *PolicyCreate) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *PolicyCreate) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *PolicyCreate) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *PolicyCreate) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetAccounts

`func (o *PolicyCreate) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *PolicyCreate) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *PolicyCreate) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *PolicyCreate) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetEachUser

`func (o *PolicyCreate) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *PolicyCreate) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *PolicyCreate) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *PolicyCreate) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


