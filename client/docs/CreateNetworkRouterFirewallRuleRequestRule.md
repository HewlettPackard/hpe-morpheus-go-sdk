# CreateNetworkRouterFirewallRuleRequestRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Firewall rule name | 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the rule (true, false). Default is on | [optional] [default to true]
**Priority** | Pointer to **int64** | Priority for rule | [optional] 

## Methods

### NewCreateNetworkRouterFirewallRuleRequestRule

`func NewCreateNetworkRouterFirewallRuleRequestRule(name string, ) *CreateNetworkRouterFirewallRuleRequestRule`

NewCreateNetworkRouterFirewallRuleRequestRule instantiates a new CreateNetworkRouterFirewallRuleRequestRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkRouterFirewallRuleRequestRuleWithDefaults

`func NewCreateNetworkRouterFirewallRuleRequestRuleWithDefaults() *CreateNetworkRouterFirewallRuleRequestRule`

NewCreateNetworkRouterFirewallRuleRequestRuleWithDefaults instantiates a new CreateNetworkRouterFirewallRuleRequestRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkRouterFirewallRuleRequestRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkRouterFirewallRuleRequestRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkRouterFirewallRuleRequestRule) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *CreateNetworkRouterFirewallRuleRequestRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateNetworkRouterFirewallRuleRequestRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateNetworkRouterFirewallRuleRequestRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateNetworkRouterFirewallRuleRequestRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPriority

`func (o *CreateNetworkRouterFirewallRuleRequestRule) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateNetworkRouterFirewallRuleRequestRule) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateNetworkRouterFirewallRuleRequestRule) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CreateNetworkRouterFirewallRuleRequestRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


