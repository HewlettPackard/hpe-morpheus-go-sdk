# GetNetworkFirewallRuleGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleGroups** | Pointer to [**[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner**](GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetNetworkFirewallRuleGroups200Response

`func NewGetNetworkFirewallRuleGroups200Response() *GetNetworkFirewallRuleGroups200Response`

NewGetNetworkFirewallRuleGroups200Response instantiates a new GetNetworkFirewallRuleGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirewallRuleGroups200ResponseWithDefaults

`func NewGetNetworkFirewallRuleGroups200ResponseWithDefaults() *GetNetworkFirewallRuleGroups200Response`

NewGetNetworkFirewallRuleGroups200ResponseWithDefaults instantiates a new GetNetworkFirewallRuleGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleGroups

`func (o *GetNetworkFirewallRuleGroups200Response) GetRuleGroups() []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner`

GetRuleGroups returns the RuleGroups field if non-nil, zero value otherwise.

### GetRuleGroupsOk

`func (o *GetNetworkFirewallRuleGroups200Response) GetRuleGroupsOk() (*[]GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner, bool)`

GetRuleGroupsOk returns a tuple with the RuleGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroups

`func (o *GetNetworkFirewallRuleGroups200Response) SetRuleGroups(v []GetNetworkFirewallRuleGroups200ResponseAllOfRuleGroupsInner)`

SetRuleGroups sets RuleGroups field to given value.

### HasRuleGroups

`func (o *GetNetworkFirewallRuleGroups200Response) HasRuleGroups() bool`

HasRuleGroups returns a boolean if a field has been set.

### GetMeta

`func (o *GetNetworkFirewallRuleGroups200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetNetworkFirewallRuleGroups200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetNetworkFirewallRuleGroups200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetNetworkFirewallRuleGroups200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


