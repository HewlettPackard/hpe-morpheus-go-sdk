# AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner

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

### NewAddSecurityGroups200ResponseSecurityGroupAllOfRulesInner

`func NewAddSecurityGroups200ResponseSecurityGroupAllOfRulesInner() *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner`

NewAddSecurityGroups200ResponseSecurityGroupAllOfRulesInner instantiates a new AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroups200ResponseSecurityGroupAllOfRulesInnerWithDefaults

`func NewAddSecurityGroups200ResponseSecurityGroupAllOfRulesInnerWithDefaults() *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner`

NewAddSecurityGroups200ResponseSecurityGroupAllOfRulesInnerWithDefaults instantiates a new AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRuleType

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetCustomRule

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetCustomRule() bool`

GetCustomRule returns the CustomRule field if non-nil, zero value otherwise.

### GetCustomRuleOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetCustomRuleOk() (*bool, bool)`

GetCustomRuleOk returns a tuple with the CustomRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRule

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetCustomRule(v bool)`

SetCustomRule sets CustomRule field to given value.

### HasCustomRule

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasCustomRule() bool`

HasCustomRule returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### SetInstanceTypeIdNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetInstanceTypeIdNil(b bool)`

 SetInstanceTypeIdNil sets the value for InstanceTypeId to be an explicit nil

### UnsetInstanceTypeId
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) UnsetInstanceTypeId()`

UnsetInstanceTypeId ensures that no value is present for InstanceTypeId, not even an explicit nil
### GetDirection

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPolicy

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSourceType

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceGroup

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetSourceGroup() string`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetSourceGroupOk() (*string, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetSourceGroup(v string)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### SetSourceGroupNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetSourceGroupNil(b bool)`

 SetSourceGroupNil sets the value for SourceGroup to be an explicit nil

### UnsetSourceGroup
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) UnsetSourceGroup()`

UnsetSourceGroup ensures that no value is present for SourceGroup, not even an explicit nil
### GetSourceTier

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetSourceTier() string`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetSourceTierOk() (*string, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetSourceTier(v string)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### SetSourceTierNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetSourceTierNil(b bool)`

 SetSourceTierNil sets the value for SourceTier to be an explicit nil

### UnsetSourceTier
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) UnsetSourceTier()`

UnsetSourceTier ensures that no value is present for SourceTier, not even an explicit nil
### GetPortRange

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### SetPortRangeNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetPortRangeNil(b bool)`

 SetPortRangeNil sets the value for PortRange to be an explicit nil

### UnsetPortRange
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) UnsetPortRange()`

UnsetPortRange ensures that no value is present for PortRange, not even an explicit nil
### GetProtocol

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestinationType

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetDestinationGroup

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetDestinationGroup() string`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetDestinationGroupOk() (*string, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetDestinationGroup(v string)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### SetDestinationGroupNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetDestinationGroupNil(b bool)`

 SetDestinationGroupNil sets the value for DestinationGroup to be an explicit nil

### UnsetDestinationGroup
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) UnsetDestinationGroup()`

UnsetDestinationGroup ensures that no value is present for DestinationGroup, not even an explicit nil
### GetDestinationTier

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetDestinationTier() string`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetDestinationTierOk() (*string, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetDestinationTier(v string)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### SetDestinationTierNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetDestinationTierNil(b bool)`

 SetDestinationTierNil sets the value for DestinationTier to be an explicit nil

### UnsetDestinationTier
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) UnsetDestinationTier()`

UnsetDestinationTier ensures that no value is present for DestinationTier, not even an explicit nil
### GetExternalId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *AddSecurityGroups200ResponseSecurityGroupAllOfRulesInner) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


