# AddSecurityGroupRulesRequestRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the rule | [optional] 
**Direction** | Pointer to **string** | Either &#x60;ingress&#x60; or &#x60;egress&#x60; | [optional] [default to "ingress"]
**SourceType** | Pointer to **string** | Either &#x60;cidr&#x60;, &#x60;group&#x60;, &#x60;tier&#x60;, &#x60;all&#x60; | [optional] [default to "cidr"]
**Source** | Pointer to **string** | CIDR representing the source IP(s) which should receive access. Required for &#x60;sourceType&#x60;&#x3D;cidr | [optional] 
**SourceGroup** | Pointer to [**AddSecurityGroupRulesRequestRuleSourceGroup**](AddSecurityGroupRulesRequestRuleSourceGroup.md) |  | [optional] 
**SourceTier** | Pointer to [**AddSecurityGroupRulesRequestRuleSourceTier**](AddSecurityGroupRulesRequestRuleSourceTier.md) |  | [optional] 
**PortRange** | Pointer to **string** | Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the source. Required if customRule is true, otherwise, ignored.  | [optional] 
**DestinationPortRange** | Pointer to **string** | Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the destination.  | [optional] 
**Protocol** | **string** | Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored. | 
**DestinationType** | Pointer to **string** | Either cidr, group, tier, instance. | [optional] [default to "cidr"]
**Destination** | Pointer to **string** | CIDR representing the destination IP(s) which should receive access. Required for &#x60;destinationType&#x60;&#x3D;cidr. | [optional] 
**DestinationGroup** | Pointer to [**AddSecurityGroupRulesRequestRuleDestinationGroup**](AddSecurityGroupRulesRequestRuleDestinationGroup.md) |  | [optional] 
**DestinationTier** | Pointer to [**AddSecurityGroupRulesRequestRuleDestinationTier**](AddSecurityGroupRulesRequestRuleDestinationTier.md) |  | [optional] 
**RuleType** | **string** | Either &#x60;customRule&#x60; or an &#x60;instance type&#x60; code. | [default to "customRule"]
**Policy** | Pointer to **string** | Either &#x60;accept&#x60; or &#x60;deny&#x60;. | [optional] 
**InstanceTypeId** | Pointer to **int64** | The id of an Instance Type. If specified, the source CIDR will have access to all ports exposed by the particular instance in the cloud, app, or instance. Required if customRule is false, otherwise ignored.  | [optional] 

## Methods

### NewAddSecurityGroupRulesRequestRule

`func NewAddSecurityGroupRulesRequestRule(protocol string, ruleType string, ) *AddSecurityGroupRulesRequestRule`

NewAddSecurityGroupRulesRequestRule instantiates a new AddSecurityGroupRulesRequestRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroupRulesRequestRuleWithDefaults

`func NewAddSecurityGroupRulesRequestRuleWithDefaults() *AddSecurityGroupRulesRequestRule`

NewAddSecurityGroupRulesRequestRuleWithDefaults instantiates a new AddSecurityGroupRulesRequestRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddSecurityGroupRulesRequestRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSecurityGroupRulesRequestRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSecurityGroupRulesRequestRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddSecurityGroupRulesRequestRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirection

`func (o *AddSecurityGroupRulesRequestRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddSecurityGroupRulesRequestRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddSecurityGroupRulesRequestRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AddSecurityGroupRulesRequestRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetSourceType

`func (o *AddSecurityGroupRulesRequestRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AddSecurityGroupRulesRequestRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AddSecurityGroupRulesRequestRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *AddSecurityGroupRulesRequestRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *AddSecurityGroupRulesRequestRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AddSecurityGroupRulesRequestRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AddSecurityGroupRulesRequestRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AddSecurityGroupRulesRequestRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceGroup

`func (o *AddSecurityGroupRulesRequestRule) GetSourceGroup() AddSecurityGroupRulesRequestRuleSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *AddSecurityGroupRulesRequestRule) GetSourceGroupOk() (*AddSecurityGroupRulesRequestRuleSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *AddSecurityGroupRulesRequestRule) SetSourceGroup(v AddSecurityGroupRulesRequestRuleSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *AddSecurityGroupRulesRequestRule) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetSourceTier

`func (o *AddSecurityGroupRulesRequestRule) GetSourceTier() AddSecurityGroupRulesRequestRuleSourceTier`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *AddSecurityGroupRulesRequestRule) GetSourceTierOk() (*AddSecurityGroupRulesRequestRuleSourceTier, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *AddSecurityGroupRulesRequestRule) SetSourceTier(v AddSecurityGroupRulesRequestRuleSourceTier)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *AddSecurityGroupRulesRequestRule) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### GetPortRange

`func (o *AddSecurityGroupRulesRequestRule) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *AddSecurityGroupRulesRequestRule) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *AddSecurityGroupRulesRequestRule) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *AddSecurityGroupRulesRequestRule) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetDestinationPortRange

`func (o *AddSecurityGroupRulesRequestRule) GetDestinationPortRange() string`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *AddSecurityGroupRulesRequestRule) GetDestinationPortRangeOk() (*string, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *AddSecurityGroupRulesRequestRule) SetDestinationPortRange(v string)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *AddSecurityGroupRulesRequestRule) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### GetProtocol

`func (o *AddSecurityGroupRulesRequestRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AddSecurityGroupRulesRequestRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AddSecurityGroupRulesRequestRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestinationType

`func (o *AddSecurityGroupRulesRequestRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *AddSecurityGroupRulesRequestRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *AddSecurityGroupRulesRequestRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *AddSecurityGroupRulesRequestRule) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *AddSecurityGroupRulesRequestRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AddSecurityGroupRulesRequestRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AddSecurityGroupRulesRequestRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *AddSecurityGroupRulesRequestRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationGroup

`func (o *AddSecurityGroupRulesRequestRule) GetDestinationGroup() AddSecurityGroupRulesRequestRuleDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *AddSecurityGroupRulesRequestRule) GetDestinationGroupOk() (*AddSecurityGroupRulesRequestRuleDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *AddSecurityGroupRulesRequestRule) SetDestinationGroup(v AddSecurityGroupRulesRequestRuleDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *AddSecurityGroupRulesRequestRule) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetDestinationTier

`func (o *AddSecurityGroupRulesRequestRule) GetDestinationTier() AddSecurityGroupRulesRequestRuleDestinationTier`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *AddSecurityGroupRulesRequestRule) GetDestinationTierOk() (*AddSecurityGroupRulesRequestRuleDestinationTier, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *AddSecurityGroupRulesRequestRule) SetDestinationTier(v AddSecurityGroupRulesRequestRuleDestinationTier)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *AddSecurityGroupRulesRequestRule) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### GetRuleType

`func (o *AddSecurityGroupRulesRequestRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AddSecurityGroupRulesRequestRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AddSecurityGroupRulesRequestRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.


### GetPolicy

`func (o *AddSecurityGroupRulesRequestRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AddSecurityGroupRulesRequestRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AddSecurityGroupRulesRequestRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *AddSecurityGroupRulesRequestRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *AddSecurityGroupRulesRequestRule) GetInstanceTypeId() int64`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *AddSecurityGroupRulesRequestRule) GetInstanceTypeIdOk() (*int64, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *AddSecurityGroupRulesRequestRule) SetInstanceTypeId(v int64)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *AddSecurityGroupRulesRequestRule) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


