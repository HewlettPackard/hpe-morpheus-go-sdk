# NetworkFirewallRuleCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **[]string** |  | [optional] 
**Profile** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkFirewallRuleCreateConfig

`func NewNetworkFirewallRuleCreateConfig() *NetworkFirewallRuleCreateConfig`

NewNetworkFirewallRuleCreateConfig instantiates a new NetworkFirewallRuleCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkFirewallRuleCreateConfigWithDefaults

`func NewNetworkFirewallRuleCreateConfigWithDefaults() *NetworkFirewallRuleCreateConfig`

NewNetworkFirewallRuleCreateConfigWithDefaults instantiates a new NetworkFirewallRuleCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *NetworkFirewallRuleCreateConfig) GetApplication() []string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *NetworkFirewallRuleCreateConfig) GetApplicationOk() (*[]string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *NetworkFirewallRuleCreateConfig) SetApplication(v []string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *NetworkFirewallRuleCreateConfig) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetProfile

`func (o *NetworkFirewallRuleCreateConfig) GetProfile() []string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *NetworkFirewallRuleCreateConfig) GetProfileOk() (*[]string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *NetworkFirewallRuleCreateConfig) SetProfile(v []string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *NetworkFirewallRuleCreateConfig) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


