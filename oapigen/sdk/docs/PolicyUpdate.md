# PolicyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the policy | [optional] 
**Description** | Pointer to **string** | A description for the policy | [optional] 
**Config** | Pointer to [**UpdatePoliciesRequestPolicyConfig**](UpdatePoliciesRequestPolicyConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Set to false to disable | [optional] [default to true]
**RefType** | Pointer to **string** | Scope object type | [optional] 
**RefId** | Pointer to **int64** | Scope object ID (&#x60;group&#x60;,&#x60;cloud&#x60;,&#x60;user&#x60;, etc) | [optional] 
**Accounts** | Pointer to **[]int64** | Array of tenants to scope the policy to | [optional] 
**EachUser** | Pointer to **bool** | Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60; | [optional] 

## Methods

### NewPolicyUpdate

`func NewPolicyUpdate() *PolicyUpdate`

NewPolicyUpdate instantiates a new PolicyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyUpdateWithDefaults

`func NewPolicyUpdateWithDefaults() *PolicyUpdate`

NewPolicyUpdateWithDefaults instantiates a new PolicyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PolicyUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfig

`func (o *PolicyUpdate) GetConfig() UpdatePoliciesRequestPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PolicyUpdate) GetConfigOk() (*UpdatePoliciesRequestPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PolicyUpdate) SetConfig(v UpdatePoliciesRequestPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PolicyUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *PolicyUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PolicyUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PolicyUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PolicyUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefType

`func (o *PolicyUpdate) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *PolicyUpdate) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *PolicyUpdate) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *PolicyUpdate) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *PolicyUpdate) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *PolicyUpdate) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *PolicyUpdate) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *PolicyUpdate) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetAccounts

`func (o *PolicyUpdate) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *PolicyUpdate) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *PolicyUpdate) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *PolicyUpdate) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetEachUser

`func (o *PolicyUpdate) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *PolicyUpdate) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *PolicyUpdate) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *PolicyUpdate) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


