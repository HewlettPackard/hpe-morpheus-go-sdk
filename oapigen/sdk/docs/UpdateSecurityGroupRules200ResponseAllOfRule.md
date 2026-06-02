# UpdateSecurityGroupRules200ResponseAllOfRule

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
**SourceGroup** | Pointer to [**UpdateSecurityGroupRules200ResponseAllOfRuleSourceGroup**](UpdateSecurityGroupRules200ResponseAllOfRuleSourceGroup.md) |  | [optional] 
**SourceTier** | Pointer to [**UpdateSecurityGroupRules200ResponseAllOfRuleSourceTier**](UpdateSecurityGroupRules200ResponseAllOfRuleSourceTier.md) |  | [optional] 
**PortRange** | Pointer to **NullableString** |  | [optional] 
**SourcePortRange** | Pointer to **NullableString** |  | [optional] 
**DestinationPortRange** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **NullableString** |  | [optional] 
**DestinationGroup** | Pointer to [**UpdateSecurityGroupRules200ResponseAllOfRuleDestinationGroup**](UpdateSecurityGroupRules200ResponseAllOfRuleDestinationGroup.md) |  | [optional] 
**DestinationTier** | Pointer to [**UpdateSecurityGroupRules200ResponseAllOfRuleDestinationTier**](UpdateSecurityGroupRules200ResponseAllOfRuleDestinationTier.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateSecurityGroupRules200ResponseAllOfRule

`func NewUpdateSecurityGroupRules200ResponseAllOfRule() *UpdateSecurityGroupRules200ResponseAllOfRule`

NewUpdateSecurityGroupRules200ResponseAllOfRule instantiates a new UpdateSecurityGroupRules200ResponseAllOfRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRuleType

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetCustomRule

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetCustomRule() bool`

GetCustomRule returns the CustomRule field if non-nil, zero value otherwise.

### GetCustomRuleOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetCustomRuleOk() (*bool, bool)`

GetCustomRuleOk returns a tuple with the CustomRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRule

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetCustomRule(v bool)`

SetCustomRule sets CustomRule field to given value.

### HasCustomRule

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasCustomRule() bool`

HasCustomRule returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### SetInstanceTypeIdNil

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetInstanceTypeIdNil(b bool)`

 SetInstanceTypeIdNil sets the value for InstanceTypeId to be an explicit nil

### UnsetInstanceTypeId
`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) UnsetInstanceTypeId()`

UnsetInstanceTypeId ensures that no value is present for InstanceTypeId, not even an explicit nil
### GetDirection

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPolicy

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSourceType

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceGroup

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetSourceGroup() UpdateSecurityGroupRules200ResponseAllOfRuleSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetSourceGroupOk() (*UpdateSecurityGroupRules200ResponseAllOfRuleSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetSourceGroup(v UpdateSecurityGroupRules200ResponseAllOfRuleSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetSourceTier

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetSourceTier() UpdateSecurityGroupRules200ResponseAllOfRuleSourceTier`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetSourceTierOk() (*UpdateSecurityGroupRules200ResponseAllOfRuleSourceTier, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetSourceTier(v UpdateSecurityGroupRules200ResponseAllOfRuleSourceTier)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### GetPortRange

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### SetPortRangeNil

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetPortRangeNil(b bool)`

 SetPortRangeNil sets the value for PortRange to be an explicit nil

### UnsetPortRange
`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) UnsetPortRange()`

UnsetPortRange ensures that no value is present for PortRange, not even an explicit nil
### GetSourcePortRange

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetSourcePortRange() string`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetSourcePortRangeOk() (*string, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetSourcePortRange(v string)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### SetSourcePortRangeNil

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetSourcePortRangeNil(b bool)`

 SetSourcePortRangeNil sets the value for SourcePortRange to be an explicit nil

### UnsetSourcePortRange
`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) UnsetSourcePortRange()`

UnsetSourcePortRange ensures that no value is present for SourcePortRange, not even an explicit nil
### GetDestinationPortRange

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDestinationPortRange() string`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDestinationPortRangeOk() (*string, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetDestinationPortRange(v string)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### SetDestinationPortRangeNil

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetDestinationPortRangeNil(b bool)`

 SetDestinationPortRangeNil sets the value for DestinationPortRange to be an explicit nil

### UnsetDestinationPortRange
`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) UnsetDestinationPortRange()`

UnsetDestinationPortRange ensures that no value is present for DestinationPortRange, not even an explicit nil
### GetProtocol

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestinationType

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetDestinationGroup

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDestinationGroup() UpdateSecurityGroupRules200ResponseAllOfRuleDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDestinationGroupOk() (*UpdateSecurityGroupRules200ResponseAllOfRuleDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetDestinationGroup(v UpdateSecurityGroupRules200ResponseAllOfRuleDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetDestinationTier

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDestinationTier() UpdateSecurityGroupRules200ResponseAllOfRuleDestinationTier`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetDestinationTierOk() (*UpdateSecurityGroupRules200ResponseAllOfRuleDestinationTier, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetDestinationTier(v UpdateSecurityGroupRules200ResponseAllOfRuleDestinationTier)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetEnabled

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *UpdateSecurityGroupRules200ResponseAllOfRule) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


