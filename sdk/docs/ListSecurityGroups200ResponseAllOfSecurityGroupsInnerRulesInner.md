# ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RuleType** | Pointer to **string** |  | [optional] 
**CustomRule** | Pointer to **bool** |  | [optional] 
**InstanceTypeId** | Pointer to **NullableString** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**SourceGroup** | Pointer to **NullableString** |  | [optional] 
**SourceTier** | Pointer to **NullableString** |  | [optional] 
**PortRange** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **NullableString** |  | [optional] 
**DestinationGroup** | Pointer to **NullableString** |  | [optional] 
**DestinationTier** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner

`func NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner() *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner`

NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner instantiates a new ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInnerWithDefaults

`func NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInnerWithDefaults() *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner`

NewListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInnerWithDefaults instantiates a new ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRuleType

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetCustomRule

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetCustomRule() bool`

GetCustomRule returns the CustomRule field if non-nil, zero value otherwise.

### GetCustomRuleOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetCustomRuleOk() (*bool, bool)`

GetCustomRuleOk returns a tuple with the CustomRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRule

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetCustomRule(v bool)`

SetCustomRule sets CustomRule field to given value.

### HasCustomRule

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasCustomRule() bool`

HasCustomRule returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### SetInstanceTypeIdNil

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetInstanceTypeIdNil(b bool)`

 SetInstanceTypeIdNil sets the value for InstanceTypeId to be an explicit nil

### UnsetInstanceTypeId
`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) UnsetInstanceTypeId()`

UnsetInstanceTypeId ensures that no value is present for InstanceTypeId, not even an explicit nil
### GetDirection

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPolicy

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSourceType

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceGroup

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetSourceGroup() string`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetSourceGroupOk() (*string, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetSourceGroup(v string)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### SetSourceGroupNil

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetSourceGroupNil(b bool)`

 SetSourceGroupNil sets the value for SourceGroup to be an explicit nil

### UnsetSourceGroup
`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) UnsetSourceGroup()`

UnsetSourceGroup ensures that no value is present for SourceGroup, not even an explicit nil
### GetSourceTier

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetSourceTier() string`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetSourceTierOk() (*string, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetSourceTier(v string)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### SetSourceTierNil

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetSourceTierNil(b bool)`

 SetSourceTierNil sets the value for SourceTier to be an explicit nil

### UnsetSourceTier
`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) UnsetSourceTier()`

UnsetSourceTier ensures that no value is present for SourceTier, not even an explicit nil
### GetPortRange

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### SetPortRangeNil

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetPortRangeNil(b bool)`

 SetPortRangeNil sets the value for PortRange to be an explicit nil

### UnsetPortRange
`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) UnsetPortRange()`

UnsetPortRange ensures that no value is present for PortRange, not even an explicit nil
### GetProtocol

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestinationType

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetDestinationGroup

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetDestinationGroup() string`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetDestinationGroupOk() (*string, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetDestinationGroup(v string)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### SetDestinationGroupNil

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetDestinationGroupNil(b bool)`

 SetDestinationGroupNil sets the value for DestinationGroup to be an explicit nil

### UnsetDestinationGroup
`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) UnsetDestinationGroup()`

UnsetDestinationGroup ensures that no value is present for DestinationGroup, not even an explicit nil
### GetDestinationTier

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetDestinationTier() string`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetDestinationTierOk() (*string, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetDestinationTier(v string)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### SetDestinationTierNil

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetDestinationTierNil(b bool)`

 SetDestinationTierNil sets the value for DestinationTier to be an explicit nil

### UnsetDestinationTier
`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) UnsetDestinationTier()`

UnsetDestinationTier ensures that no value is present for DestinationTier, not even an explicit nil
### GetExternalId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEnabled

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ListSecurityGroups200ResponseAllOfSecurityGroupsInnerRulesInner) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


