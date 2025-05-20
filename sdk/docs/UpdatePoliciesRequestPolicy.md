# UpdatePoliciesRequestPolicy

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

### NewUpdatePoliciesRequestPolicy

`func NewUpdatePoliciesRequestPolicy() *UpdatePoliciesRequestPolicy`

NewUpdatePoliciesRequestPolicy instantiates a new UpdatePoliciesRequestPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePoliciesRequestPolicyWithDefaults

`func NewUpdatePoliciesRequestPolicyWithDefaults() *UpdatePoliciesRequestPolicy`

NewUpdatePoliciesRequestPolicyWithDefaults instantiates a new UpdatePoliciesRequestPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePoliciesRequestPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePoliciesRequestPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePoliciesRequestPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePoliciesRequestPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePoliciesRequestPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePoliciesRequestPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePoliciesRequestPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePoliciesRequestPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfig

`func (o *UpdatePoliciesRequestPolicy) GetConfig() UpdatePoliciesRequestPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdatePoliciesRequestPolicy) GetConfigOk() (*UpdatePoliciesRequestPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdatePoliciesRequestPolicy) SetConfig(v UpdatePoliciesRequestPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdatePoliciesRequestPolicy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdatePoliciesRequestPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdatePoliciesRequestPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdatePoliciesRequestPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdatePoliciesRequestPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefType

`func (o *UpdatePoliciesRequestPolicy) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdatePoliciesRequestPolicy) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdatePoliciesRequestPolicy) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdatePoliciesRequestPolicy) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *UpdatePoliciesRequestPolicy) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdatePoliciesRequestPolicy) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdatePoliciesRequestPolicy) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdatePoliciesRequestPolicy) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetAccounts

`func (o *UpdatePoliciesRequestPolicy) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdatePoliciesRequestPolicy) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdatePoliciesRequestPolicy) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UpdatePoliciesRequestPolicy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetEachUser

`func (o *UpdatePoliciesRequestPolicy) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *UpdatePoliciesRequestPolicy) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *UpdatePoliciesRequestPolicy) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *UpdatePoliciesRequestPolicy) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


