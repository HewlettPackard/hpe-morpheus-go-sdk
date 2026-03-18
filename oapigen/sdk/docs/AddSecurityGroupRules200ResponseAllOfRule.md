# AddSecurityGroupRules200ResponseAllOfRule

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
**SourceGroup** | Pointer to [**AddSecurityGroupRules200ResponseAllOfRuleSourceGroup**](AddSecurityGroupRules200ResponseAllOfRuleSourceGroup.md) |  | [optional] 
**SourceTier** | Pointer to [**AddSecurityGroupRules200ResponseAllOfRuleSourceTier**](AddSecurityGroupRules200ResponseAllOfRuleSourceTier.md) |  | [optional] 
**PortRange** | Pointer to **NullableString** |  | [optional] 
**SourcePortRange** | Pointer to **NullableString** |  | [optional] 
**DestinationPortRange** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **NullableString** |  | [optional] 
**DestinationGroup** | Pointer to [**AddSecurityGroupRules200ResponseAllOfRuleDestinationGroup**](AddSecurityGroupRules200ResponseAllOfRuleDestinationGroup.md) |  | [optional] 
**DestinationTier** | Pointer to [**AddSecurityGroupRules200ResponseAllOfRuleDestinationTier**](AddSecurityGroupRules200ResponseAllOfRuleDestinationTier.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAddSecurityGroupRules200ResponseAllOfRule

`func NewAddSecurityGroupRules200ResponseAllOfRule() *AddSecurityGroupRules200ResponseAllOfRule`

NewAddSecurityGroupRules200ResponseAllOfRule instantiates a new AddSecurityGroupRules200ResponseAllOfRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroupRules200ResponseAllOfRuleWithDefaults

`func NewAddSecurityGroupRules200ResponseAllOfRuleWithDefaults() *AddSecurityGroupRules200ResponseAllOfRule`

NewAddSecurityGroupRules200ResponseAllOfRuleWithDefaults instantiates a new AddSecurityGroupRules200ResponseAllOfRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddSecurityGroupRules200ResponseAllOfRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRuleType

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetCustomRule

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetCustomRule() bool`

GetCustomRule returns the CustomRule field if non-nil, zero value otherwise.

### GetCustomRuleOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetCustomRuleOk() (*bool, bool)`

GetCustomRuleOk returns a tuple with the CustomRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRule

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetCustomRule(v bool)`

SetCustomRule sets CustomRule field to given value.

### HasCustomRule

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasCustomRule() bool`

HasCustomRule returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### SetInstanceTypeIdNil

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetInstanceTypeIdNil(b bool)`

 SetInstanceTypeIdNil sets the value for InstanceTypeId to be an explicit nil

### UnsetInstanceTypeId
`func (o *AddSecurityGroupRules200ResponseAllOfRule) UnsetInstanceTypeId()`

UnsetInstanceTypeId ensures that no value is present for InstanceTypeId, not even an explicit nil
### GetDirection

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPolicy

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSourceType

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *AddSecurityGroupRules200ResponseAllOfRule) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceGroup

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetSourceGroup() AddSecurityGroupRules200ResponseAllOfRuleSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetSourceGroupOk() (*AddSecurityGroupRules200ResponseAllOfRuleSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetSourceGroup(v AddSecurityGroupRules200ResponseAllOfRuleSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetSourceTier

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetSourceTier() AddSecurityGroupRules200ResponseAllOfRuleSourceTier`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetSourceTierOk() (*AddSecurityGroupRules200ResponseAllOfRuleSourceTier, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetSourceTier(v AddSecurityGroupRules200ResponseAllOfRuleSourceTier)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### GetPortRange

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### SetPortRangeNil

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetPortRangeNil(b bool)`

 SetPortRangeNil sets the value for PortRange to be an explicit nil

### UnsetPortRange
`func (o *AddSecurityGroupRules200ResponseAllOfRule) UnsetPortRange()`

UnsetPortRange ensures that no value is present for PortRange, not even an explicit nil
### GetSourcePortRange

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetSourcePortRange() string`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetSourcePortRangeOk() (*string, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetSourcePortRange(v string)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### SetSourcePortRangeNil

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetSourcePortRangeNil(b bool)`

 SetSourcePortRangeNil sets the value for SourcePortRange to be an explicit nil

### UnsetSourcePortRange
`func (o *AddSecurityGroupRules200ResponseAllOfRule) UnsetSourcePortRange()`

UnsetSourcePortRange ensures that no value is present for SourcePortRange, not even an explicit nil
### GetDestinationPortRange

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDestinationPortRange() string`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDestinationPortRangeOk() (*string, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetDestinationPortRange(v string)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### SetDestinationPortRangeNil

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetDestinationPortRangeNil(b bool)`

 SetDestinationPortRangeNil sets the value for DestinationPortRange to be an explicit nil

### UnsetDestinationPortRange
`func (o *AddSecurityGroupRules200ResponseAllOfRule) UnsetDestinationPortRange()`

UnsetDestinationPortRange ensures that no value is present for DestinationPortRange, not even an explicit nil
### GetProtocol

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestinationType

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *AddSecurityGroupRules200ResponseAllOfRule) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetDestinationGroup

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDestinationGroup() AddSecurityGroupRules200ResponseAllOfRuleDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDestinationGroupOk() (*AddSecurityGroupRules200ResponseAllOfRuleDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetDestinationGroup(v AddSecurityGroupRules200ResponseAllOfRuleDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetDestinationTier

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDestinationTier() AddSecurityGroupRules200ResponseAllOfRuleDestinationTier`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetDestinationTierOk() (*AddSecurityGroupRules200ResponseAllOfRuleDestinationTier, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetDestinationTier(v AddSecurityGroupRules200ResponseAllOfRuleDestinationTier)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### GetExternalId

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddSecurityGroupRules200ResponseAllOfRule) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetEnabled

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSecurityGroupRules200ResponseAllOfRule) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddSecurityGroupRules200ResponseAllOfRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *AddSecurityGroupRules200ResponseAllOfRule) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *AddSecurityGroupRules200ResponseAllOfRule) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


