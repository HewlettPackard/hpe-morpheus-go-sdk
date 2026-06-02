# PolicyGroupCreatePolicyType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The policy type | [optional] 
**Config** | Pointer to [**PolicyGroupCreatePolicyTypeConfig**](PolicyGroupCreatePolicyTypeConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Set to false to disable | [optional] [default to true]
**RefType** | Pointer to **string** | Scope object type | [optional] 
**RefId** | Pointer to **int64** | Scope object ID (&#x60;group&#x60;) | [optional] 
**Accounts** | Pointer to **[]int64** | Array of tenants to scope the policy to | [optional] 
**EachUser** | Pointer to **bool** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; | [optional] 

## Methods

### NewPolicyGroupCreatePolicyType

`func NewPolicyGroupCreatePolicyType() *PolicyGroupCreatePolicyType`

NewPolicyGroupCreatePolicyType instantiates a new PolicyGroupCreatePolicyType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetCode

`func (o *PolicyGroupCreatePolicyType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PolicyGroupCreatePolicyType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PolicyGroupCreatePolicyType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PolicyGroupCreatePolicyType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConfig

`func (o *PolicyGroupCreatePolicyType) GetConfig() PolicyGroupCreatePolicyTypeConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PolicyGroupCreatePolicyType) GetConfigOk() (*PolicyGroupCreatePolicyTypeConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PolicyGroupCreatePolicyType) SetConfig(v PolicyGroupCreatePolicyTypeConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PolicyGroupCreatePolicyType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *PolicyGroupCreatePolicyType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PolicyGroupCreatePolicyType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PolicyGroupCreatePolicyType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PolicyGroupCreatePolicyType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefType

`func (o *PolicyGroupCreatePolicyType) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *PolicyGroupCreatePolicyType) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *PolicyGroupCreatePolicyType) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *PolicyGroupCreatePolicyType) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *PolicyGroupCreatePolicyType) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *PolicyGroupCreatePolicyType) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *PolicyGroupCreatePolicyType) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *PolicyGroupCreatePolicyType) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetAccounts

`func (o *PolicyGroupCreatePolicyType) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *PolicyGroupCreatePolicyType) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *PolicyGroupCreatePolicyType) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *PolicyGroupCreatePolicyType) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetEachUser

`func (o *PolicyGroupCreatePolicyType) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *PolicyGroupCreatePolicyType) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *PolicyGroupCreatePolicyType) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *PolicyGroupCreatePolicyType) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


