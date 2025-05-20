# ListSecurityGroupRules200ResponseAllOfRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RuleType** | Pointer to **string** |  | [optional] 
**CustomRule** | Pointer to **bool** |  | [optional] 
**InstanceTypeId** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SourceGroup** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**SourceTier** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**PortRange** | Pointer to **string** |  | [optional] 
**SourcePortRange** | Pointer to **string** |  | [optional] 
**DestinationPortRange** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **string** |  | [optional] 
**DestinationGroup** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**DestinationTier** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **string** |  | [optional] 

## Methods

### NewListSecurityGroupRules200ResponseAllOfRulesInner

`func NewListSecurityGroupRules200ResponseAllOfRulesInner() *ListSecurityGroupRules200ResponseAllOfRulesInner`

NewListSecurityGroupRules200ResponseAllOfRulesInner instantiates a new ListSecurityGroupRules200ResponseAllOfRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityGroupRules200ResponseAllOfRulesInnerWithDefaults

`func NewListSecurityGroupRules200ResponseAllOfRulesInnerWithDefaults() *ListSecurityGroupRules200ResponseAllOfRulesInner`

NewListSecurityGroupRules200ResponseAllOfRulesInnerWithDefaults instantiates a new ListSecurityGroupRules200ResponseAllOfRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleType

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetCustomRule

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetCustomRule() bool`

GetCustomRule returns the CustomRule field if non-nil, zero value otherwise.

### GetCustomRuleOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetCustomRuleOk() (*bool, bool)`

GetCustomRuleOk returns a tuple with the CustomRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRule

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetCustomRule(v bool)`

SetCustomRule sets CustomRule field to given value.

### HasCustomRule

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasCustomRule() bool`

HasCustomRule returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### GetDirection

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPolicy

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSourceType

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceGroup

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetSourceGroup() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetSourceGroupOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetSourceGroup(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetSourceTier

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetSourceTier() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetSourceTierOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetSourceTier(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### GetPortRange

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### GetSourcePortRange

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetSourcePortRange() string`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetSourcePortRangeOk() (*string, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetSourcePortRange(v string)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### GetDestinationPortRange

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDestinationPortRange() string`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDestinationPortRangeOk() (*string, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetDestinationPortRange(v string)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### GetProtocol

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestinationType

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationGroup

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDestinationGroup() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDestinationGroupOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetDestinationGroup(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetDestinationTier

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDestinationTier() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetDestinationTierOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetDestinationTier(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### GetExternalId

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEnabled

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListSecurityGroupRules200ResponseAllOfRulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


