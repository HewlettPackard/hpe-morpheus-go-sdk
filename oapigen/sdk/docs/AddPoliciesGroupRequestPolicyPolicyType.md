# AddPoliciesGroupRequestPolicyPolicyType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The policy type | [optional] 
**Config** | Pointer to [**AddPoliciesGroupRequestPolicyPolicyTypeConfig**](AddPoliciesGroupRequestPolicyPolicyTypeConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Set to false to disable | [optional] [default to true]
**RefType** | Pointer to **string** | Scope object type | [optional] 
**RefId** | Pointer to **int64** | Scope object ID (&#x60;group&#x60;) | [optional] 
**Accounts** | Pointer to **[]int64** | Array of tenants to scope the policy to | [optional] 
**EachUser** | Pointer to **bool** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; | [optional] 

## Methods

### NewAddPoliciesGroupRequestPolicyPolicyType

`func NewAddPoliciesGroupRequestPolicyPolicyType() *AddPoliciesGroupRequestPolicyPolicyType`

NewAddPoliciesGroupRequestPolicyPolicyType instantiates a new AddPoliciesGroupRequestPolicyPolicyType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPoliciesGroupRequestPolicyPolicyTypeWithDefaults

`func NewAddPoliciesGroupRequestPolicyPolicyTypeWithDefaults() *AddPoliciesGroupRequestPolicyPolicyType`

NewAddPoliciesGroupRequestPolicyPolicyTypeWithDefaults instantiates a new AddPoliciesGroupRequestPolicyPolicyType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddPoliciesGroupRequestPolicyPolicyType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddPoliciesGroupRequestPolicyPolicyType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConfig

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetConfig() AddPoliciesGroupRequestPolicyPolicyTypeConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetConfigOk() (*AddPoliciesGroupRequestPolicyPolicyTypeConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddPoliciesGroupRequestPolicyPolicyType) SetConfig(v AddPoliciesGroupRequestPolicyPolicyTypeConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddPoliciesGroupRequestPolicyPolicyType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPoliciesGroupRequestPolicyPolicyType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddPoliciesGroupRequestPolicyPolicyType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefType

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *AddPoliciesGroupRequestPolicyPolicyType) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *AddPoliciesGroupRequestPolicyPolicyType) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AddPoliciesGroupRequestPolicyPolicyType) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AddPoliciesGroupRequestPolicyPolicyType) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetAccounts

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AddPoliciesGroupRequestPolicyPolicyType) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AddPoliciesGroupRequestPolicyPolicyType) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetEachUser

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *AddPoliciesGroupRequestPolicyPolicyType) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *AddPoliciesGroupRequestPolicyPolicyType) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *AddPoliciesGroupRequestPolicyPolicyType) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


