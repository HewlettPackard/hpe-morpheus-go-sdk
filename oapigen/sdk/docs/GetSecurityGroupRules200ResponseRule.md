# GetSecurityGroupRules200ResponseRule

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
**SourceGroup** | Pointer to [**GetSecurityGroupRules200ResponseRuleSourceGroup**](GetSecurityGroupRules200ResponseRuleSourceGroup.md) |  | [optional] 
**SourceTier** | Pointer to [**GetSecurityGroupRules200ResponseRuleSourceTier**](GetSecurityGroupRules200ResponseRuleSourceTier.md) |  | [optional] 
**PortRange** | Pointer to **NullableString** |  | [optional] 
**SourcePortRange** | Pointer to **NullableString** |  | [optional] 
**DestinationPortRange** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to **NullableString** |  | [optional] 
**DestinationGroup** | Pointer to [**GetSecurityGroupRules200ResponseRuleDestinationGroup**](GetSecurityGroupRules200ResponseRuleDestinationGroup.md) |  | [optional] 
**DestinationTier** | Pointer to [**GetSecurityGroupRules200ResponseRuleDestinationTier**](GetSecurityGroupRules200ResponseRuleDestinationTier.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetSecurityGroupRules200ResponseRule

`func NewGetSecurityGroupRules200ResponseRule() *GetSecurityGroupRules200ResponseRule`

NewGetSecurityGroupRules200ResponseRule instantiates a new GetSecurityGroupRules200ResponseRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecurityGroupRules200ResponseRuleWithDefaults

`func NewGetSecurityGroupRules200ResponseRuleWithDefaults() *GetSecurityGroupRules200ResponseRule`

NewGetSecurityGroupRules200ResponseRuleWithDefaults instantiates a new GetSecurityGroupRules200ResponseRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSecurityGroupRules200ResponseRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSecurityGroupRules200ResponseRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSecurityGroupRules200ResponseRule) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSecurityGroupRules200ResponseRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetSecurityGroupRules200ResponseRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSecurityGroupRules200ResponseRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSecurityGroupRules200ResponseRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSecurityGroupRules200ResponseRule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetSecurityGroupRules200ResponseRule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetSecurityGroupRules200ResponseRule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRuleType

`func (o *GetSecurityGroupRules200ResponseRule) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *GetSecurityGroupRules200ResponseRule) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *GetSecurityGroupRules200ResponseRule) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *GetSecurityGroupRules200ResponseRule) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetCustomRule

`func (o *GetSecurityGroupRules200ResponseRule) GetCustomRule() bool`

GetCustomRule returns the CustomRule field if non-nil, zero value otherwise.

### GetCustomRuleOk

`func (o *GetSecurityGroupRules200ResponseRule) GetCustomRuleOk() (*bool, bool)`

GetCustomRuleOk returns a tuple with the CustomRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRule

`func (o *GetSecurityGroupRules200ResponseRule) SetCustomRule(v bool)`

SetCustomRule sets CustomRule field to given value.

### HasCustomRule

`func (o *GetSecurityGroupRules200ResponseRule) HasCustomRule() bool`

HasCustomRule returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *GetSecurityGroupRules200ResponseRule) GetInstanceTypeId() string`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *GetSecurityGroupRules200ResponseRule) GetInstanceTypeIdOk() (*string, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *GetSecurityGroupRules200ResponseRule) SetInstanceTypeId(v string)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *GetSecurityGroupRules200ResponseRule) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### SetInstanceTypeIdNil

`func (o *GetSecurityGroupRules200ResponseRule) SetInstanceTypeIdNil(b bool)`

 SetInstanceTypeIdNil sets the value for InstanceTypeId to be an explicit nil

### UnsetInstanceTypeId
`func (o *GetSecurityGroupRules200ResponseRule) UnsetInstanceTypeId()`

UnsetInstanceTypeId ensures that no value is present for InstanceTypeId, not even an explicit nil
### GetDirection

`func (o *GetSecurityGroupRules200ResponseRule) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetSecurityGroupRules200ResponseRule) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetSecurityGroupRules200ResponseRule) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetSecurityGroupRules200ResponseRule) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetPolicy

`func (o *GetSecurityGroupRules200ResponseRule) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetSecurityGroupRules200ResponseRule) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetSecurityGroupRules200ResponseRule) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetSecurityGroupRules200ResponseRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSourceType

`func (o *GetSecurityGroupRules200ResponseRule) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetSecurityGroupRules200ResponseRule) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetSecurityGroupRules200ResponseRule) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetSecurityGroupRules200ResponseRule) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSource

`func (o *GetSecurityGroupRules200ResponseRule) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetSecurityGroupRules200ResponseRule) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetSecurityGroupRules200ResponseRule) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetSecurityGroupRules200ResponseRule) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *GetSecurityGroupRules200ResponseRule) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *GetSecurityGroupRules200ResponseRule) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceGroup

`func (o *GetSecurityGroupRules200ResponseRule) GetSourceGroup() GetSecurityGroupRules200ResponseRuleSourceGroup`

GetSourceGroup returns the SourceGroup field if non-nil, zero value otherwise.

### GetSourceGroupOk

`func (o *GetSecurityGroupRules200ResponseRule) GetSourceGroupOk() (*GetSecurityGroupRules200ResponseRuleSourceGroup, bool)`

GetSourceGroupOk returns a tuple with the SourceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroup

`func (o *GetSecurityGroupRules200ResponseRule) SetSourceGroup(v GetSecurityGroupRules200ResponseRuleSourceGroup)`

SetSourceGroup sets SourceGroup field to given value.

### HasSourceGroup

`func (o *GetSecurityGroupRules200ResponseRule) HasSourceGroup() bool`

HasSourceGroup returns a boolean if a field has been set.

### GetSourceTier

`func (o *GetSecurityGroupRules200ResponseRule) GetSourceTier() GetSecurityGroupRules200ResponseRuleSourceTier`

GetSourceTier returns the SourceTier field if non-nil, zero value otherwise.

### GetSourceTierOk

`func (o *GetSecurityGroupRules200ResponseRule) GetSourceTierOk() (*GetSecurityGroupRules200ResponseRuleSourceTier, bool)`

GetSourceTierOk returns a tuple with the SourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTier

`func (o *GetSecurityGroupRules200ResponseRule) SetSourceTier(v GetSecurityGroupRules200ResponseRuleSourceTier)`

SetSourceTier sets SourceTier field to given value.

### HasSourceTier

`func (o *GetSecurityGroupRules200ResponseRule) HasSourceTier() bool`

HasSourceTier returns a boolean if a field has been set.

### GetPortRange

`func (o *GetSecurityGroupRules200ResponseRule) GetPortRange() string`

GetPortRange returns the PortRange field if non-nil, zero value otherwise.

### GetPortRangeOk

`func (o *GetSecurityGroupRules200ResponseRule) GetPortRangeOk() (*string, bool)`

GetPortRangeOk returns a tuple with the PortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRange

`func (o *GetSecurityGroupRules200ResponseRule) SetPortRange(v string)`

SetPortRange sets PortRange field to given value.

### HasPortRange

`func (o *GetSecurityGroupRules200ResponseRule) HasPortRange() bool`

HasPortRange returns a boolean if a field has been set.

### SetPortRangeNil

`func (o *GetSecurityGroupRules200ResponseRule) SetPortRangeNil(b bool)`

 SetPortRangeNil sets the value for PortRange to be an explicit nil

### UnsetPortRange
`func (o *GetSecurityGroupRules200ResponseRule) UnsetPortRange()`

UnsetPortRange ensures that no value is present for PortRange, not even an explicit nil
### GetSourcePortRange

`func (o *GetSecurityGroupRules200ResponseRule) GetSourcePortRange() string`

GetSourcePortRange returns the SourcePortRange field if non-nil, zero value otherwise.

### GetSourcePortRangeOk

`func (o *GetSecurityGroupRules200ResponseRule) GetSourcePortRangeOk() (*string, bool)`

GetSourcePortRangeOk returns a tuple with the SourcePortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortRange

`func (o *GetSecurityGroupRules200ResponseRule) SetSourcePortRange(v string)`

SetSourcePortRange sets SourcePortRange field to given value.

### HasSourcePortRange

`func (o *GetSecurityGroupRules200ResponseRule) HasSourcePortRange() bool`

HasSourcePortRange returns a boolean if a field has been set.

### SetSourcePortRangeNil

`func (o *GetSecurityGroupRules200ResponseRule) SetSourcePortRangeNil(b bool)`

 SetSourcePortRangeNil sets the value for SourcePortRange to be an explicit nil

### UnsetSourcePortRange
`func (o *GetSecurityGroupRules200ResponseRule) UnsetSourcePortRange()`

UnsetSourcePortRange ensures that no value is present for SourcePortRange, not even an explicit nil
### GetDestinationPortRange

`func (o *GetSecurityGroupRules200ResponseRule) GetDestinationPortRange() string`

GetDestinationPortRange returns the DestinationPortRange field if non-nil, zero value otherwise.

### GetDestinationPortRangeOk

`func (o *GetSecurityGroupRules200ResponseRule) GetDestinationPortRangeOk() (*string, bool)`

GetDestinationPortRangeOk returns a tuple with the DestinationPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPortRange

`func (o *GetSecurityGroupRules200ResponseRule) SetDestinationPortRange(v string)`

SetDestinationPortRange sets DestinationPortRange field to given value.

### HasDestinationPortRange

`func (o *GetSecurityGroupRules200ResponseRule) HasDestinationPortRange() bool`

HasDestinationPortRange returns a boolean if a field has been set.

### SetDestinationPortRangeNil

`func (o *GetSecurityGroupRules200ResponseRule) SetDestinationPortRangeNil(b bool)`

 SetDestinationPortRangeNil sets the value for DestinationPortRange to be an explicit nil

### UnsetDestinationPortRange
`func (o *GetSecurityGroupRules200ResponseRule) UnsetDestinationPortRange()`

UnsetDestinationPortRange ensures that no value is present for DestinationPortRange, not even an explicit nil
### GetProtocol

`func (o *GetSecurityGroupRules200ResponseRule) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetSecurityGroupRules200ResponseRule) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetSecurityGroupRules200ResponseRule) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetSecurityGroupRules200ResponseRule) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestinationType

`func (o *GetSecurityGroupRules200ResponseRule) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *GetSecurityGroupRules200ResponseRule) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *GetSecurityGroupRules200ResponseRule) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *GetSecurityGroupRules200ResponseRule) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetDestination

`func (o *GetSecurityGroupRules200ResponseRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetSecurityGroupRules200ResponseRule) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetSecurityGroupRules200ResponseRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetSecurityGroupRules200ResponseRule) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### SetDestinationNil

`func (o *GetSecurityGroupRules200ResponseRule) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *GetSecurityGroupRules200ResponseRule) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetDestinationGroup

`func (o *GetSecurityGroupRules200ResponseRule) GetDestinationGroup() GetSecurityGroupRules200ResponseRuleDestinationGroup`

GetDestinationGroup returns the DestinationGroup field if non-nil, zero value otherwise.

### GetDestinationGroupOk

`func (o *GetSecurityGroupRules200ResponseRule) GetDestinationGroupOk() (*GetSecurityGroupRules200ResponseRuleDestinationGroup, bool)`

GetDestinationGroupOk returns a tuple with the DestinationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroup

`func (o *GetSecurityGroupRules200ResponseRule) SetDestinationGroup(v GetSecurityGroupRules200ResponseRuleDestinationGroup)`

SetDestinationGroup sets DestinationGroup field to given value.

### HasDestinationGroup

`func (o *GetSecurityGroupRules200ResponseRule) HasDestinationGroup() bool`

HasDestinationGroup returns a boolean if a field has been set.

### GetDestinationTier

`func (o *GetSecurityGroupRules200ResponseRule) GetDestinationTier() GetSecurityGroupRules200ResponseRuleDestinationTier`

GetDestinationTier returns the DestinationTier field if non-nil, zero value otherwise.

### GetDestinationTierOk

`func (o *GetSecurityGroupRules200ResponseRule) GetDestinationTierOk() (*GetSecurityGroupRules200ResponseRuleDestinationTier, bool)`

GetDestinationTierOk returns a tuple with the DestinationTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTier

`func (o *GetSecurityGroupRules200ResponseRule) SetDestinationTier(v GetSecurityGroupRules200ResponseRuleDestinationTier)`

SetDestinationTier sets DestinationTier field to given value.

### HasDestinationTier

`func (o *GetSecurityGroupRules200ResponseRule) HasDestinationTier() bool`

HasDestinationTier returns a boolean if a field has been set.

### GetExternalId

`func (o *GetSecurityGroupRules200ResponseRule) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetSecurityGroupRules200ResponseRule) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetSecurityGroupRules200ResponseRule) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetSecurityGroupRules200ResponseRule) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetSecurityGroupRules200ResponseRule) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetSecurityGroupRules200ResponseRule) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetEnabled

`func (o *GetSecurityGroupRules200ResponseRule) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetSecurityGroupRules200ResponseRule) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetSecurityGroupRules200ResponseRule) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetSecurityGroupRules200ResponseRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *GetSecurityGroupRules200ResponseRule) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *GetSecurityGroupRules200ResponseRule) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


