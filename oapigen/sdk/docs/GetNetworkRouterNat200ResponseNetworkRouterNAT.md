# GetNetworkRouterNat200ResponseNetworkRouterNAT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SourceNetwork** | Pointer to **string** |  | [optional] 
**DestinationNetwork** | Pointer to **NullableString** |  | [optional] 
**TranslatedNetwork** | Pointer to **string** |  | [optional] 
**SourcePorts** | Pointer to **NullableString** |  | [optional] 
**DestinationPorts** | Pointer to **NullableString** |  | [optional] 
**TranslatedPorts** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **NullableString** |  | [optional] 
**MatchIpv6DestinationPrefix** | Pointer to **NullableString** |  | [optional] 
**TranslatedIpv4SourcePrefix** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**SyncSource** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetNetworkRouterNat200ResponseNetworkRouterNAT

`func NewGetNetworkRouterNat200ResponseNetworkRouterNAT() *GetNetworkRouterNat200ResponseNetworkRouterNAT`

NewGetNetworkRouterNat200ResponseNetworkRouterNAT instantiates a new GetNetworkRouterNat200ResponseNetworkRouterNAT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSourceNetwork

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetSourceNetwork() string`

GetSourceNetwork returns the SourceNetwork field if non-nil, zero value otherwise.

### GetSourceNetworkOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetSourceNetworkOk() (*string, bool)`

GetSourceNetworkOk returns a tuple with the SourceNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetwork

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetSourceNetwork(v string)`

SetSourceNetwork sets SourceNetwork field to given value.

### HasSourceNetwork

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasSourceNetwork() bool`

HasSourceNetwork returns a boolean if a field has been set.

### GetDestinationNetwork

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetDestinationNetwork() string`

GetDestinationNetwork returns the DestinationNetwork field if non-nil, zero value otherwise.

### GetDestinationNetworkOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetDestinationNetworkOk() (*string, bool)`

GetDestinationNetworkOk returns a tuple with the DestinationNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNetwork

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetDestinationNetwork(v string)`

SetDestinationNetwork sets DestinationNetwork field to given value.

### HasDestinationNetwork

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasDestinationNetwork() bool`

HasDestinationNetwork returns a boolean if a field has been set.

### SetDestinationNetworkNil

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetDestinationNetworkNil(b bool)`

 SetDestinationNetworkNil sets the value for DestinationNetwork to be an explicit nil

### UnsetDestinationNetwork
`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) UnsetDestinationNetwork()`

UnsetDestinationNetwork ensures that no value is present for DestinationNetwork, not even an explicit nil
### GetTranslatedNetwork

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetTranslatedNetwork() string`

GetTranslatedNetwork returns the TranslatedNetwork field if non-nil, zero value otherwise.

### GetTranslatedNetworkOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetTranslatedNetworkOk() (*string, bool)`

GetTranslatedNetworkOk returns a tuple with the TranslatedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedNetwork

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetTranslatedNetwork(v string)`

SetTranslatedNetwork sets TranslatedNetwork field to given value.

### HasTranslatedNetwork

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasTranslatedNetwork() bool`

HasTranslatedNetwork returns a boolean if a field has been set.

### GetSourcePorts

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetSourcePorts() string`

GetSourcePorts returns the SourcePorts field if non-nil, zero value otherwise.

### GetSourcePortsOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetSourcePortsOk() (*string, bool)`

GetSourcePortsOk returns a tuple with the SourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePorts

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetSourcePorts(v string)`

SetSourcePorts sets SourcePorts field to given value.

### HasSourcePorts

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasSourcePorts() bool`

HasSourcePorts returns a boolean if a field has been set.

### SetSourcePortsNil

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetSourcePortsNil(b bool)`

 SetSourcePortsNil sets the value for SourcePorts to be an explicit nil

### UnsetSourcePorts
`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) UnsetSourcePorts()`

UnsetSourcePorts ensures that no value is present for SourcePorts, not even an explicit nil
### GetDestinationPorts

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetDestinationPorts() string`

GetDestinationPorts returns the DestinationPorts field if non-nil, zero value otherwise.

### GetDestinationPortsOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetDestinationPortsOk() (*string, bool)`

GetDestinationPortsOk returns a tuple with the DestinationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPorts

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetDestinationPorts(v string)`

SetDestinationPorts sets DestinationPorts field to given value.

### HasDestinationPorts

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasDestinationPorts() bool`

HasDestinationPorts returns a boolean if a field has been set.

### SetDestinationPortsNil

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetDestinationPortsNil(b bool)`

 SetDestinationPortsNil sets the value for DestinationPorts to be an explicit nil

### UnsetDestinationPorts
`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) UnsetDestinationPorts()`

UnsetDestinationPorts ensures that no value is present for DestinationPorts, not even an explicit nil
### GetTranslatedPorts

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetTranslatedPorts() string`

GetTranslatedPorts returns the TranslatedPorts field if non-nil, zero value otherwise.

### GetTranslatedPortsOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetTranslatedPortsOk() (*string, bool)`

GetTranslatedPortsOk returns a tuple with the TranslatedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedPorts

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetTranslatedPorts(v string)`

SetTranslatedPorts sets TranslatedPorts field to given value.

### HasTranslatedPorts

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasTranslatedPorts() bool`

HasTranslatedPorts returns a boolean if a field has been set.

### SetTranslatedPortsNil

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetTranslatedPortsNil(b bool)`

 SetTranslatedPortsNil sets the value for TranslatedPorts to be an explicit nil

### UnsetTranslatedPorts
`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) UnsetTranslatedPorts()`

UnsetTranslatedPorts ensures that no value is present for TranslatedPorts, not even an explicit nil
### GetPriority

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProtocol

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetMatchIpv6DestinationPrefix

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetMatchIpv6DestinationPrefix() string`

GetMatchIpv6DestinationPrefix returns the MatchIpv6DestinationPrefix field if non-nil, zero value otherwise.

### GetMatchIpv6DestinationPrefixOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetMatchIpv6DestinationPrefixOk() (*string, bool)`

GetMatchIpv6DestinationPrefixOk returns a tuple with the MatchIpv6DestinationPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpv6DestinationPrefix

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetMatchIpv6DestinationPrefix(v string)`

SetMatchIpv6DestinationPrefix sets MatchIpv6DestinationPrefix field to given value.

### HasMatchIpv6DestinationPrefix

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasMatchIpv6DestinationPrefix() bool`

HasMatchIpv6DestinationPrefix returns a boolean if a field has been set.

### SetMatchIpv6DestinationPrefixNil

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetMatchIpv6DestinationPrefixNil(b bool)`

 SetMatchIpv6DestinationPrefixNil sets the value for MatchIpv6DestinationPrefix to be an explicit nil

### UnsetMatchIpv6DestinationPrefix
`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) UnsetMatchIpv6DestinationPrefix()`

UnsetMatchIpv6DestinationPrefix ensures that no value is present for MatchIpv6DestinationPrefix, not even an explicit nil
### GetTranslatedIpv4SourcePrefix

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetTranslatedIpv4SourcePrefix() string`

GetTranslatedIpv4SourcePrefix returns the TranslatedIpv4SourcePrefix field if non-nil, zero value otherwise.

### GetTranslatedIpv4SourcePrefixOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetTranslatedIpv4SourcePrefixOk() (*string, bool)`

GetTranslatedIpv4SourcePrefixOk returns a tuple with the TranslatedIpv4SourcePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedIpv4SourcePrefix

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetTranslatedIpv4SourcePrefix(v string)`

SetTranslatedIpv4SourcePrefix sets TranslatedIpv4SourcePrefix field to given value.

### HasTranslatedIpv4SourcePrefix

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasTranslatedIpv4SourcePrefix() bool`

HasTranslatedIpv4SourcePrefix returns a boolean if a field has been set.

### SetTranslatedIpv4SourcePrefixNil

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetTranslatedIpv4SourcePrefixNil(b bool)`

 SetTranslatedIpv4SourcePrefixNil sets the value for TranslatedIpv4SourcePrefix to be an explicit nil

### UnsetTranslatedIpv4SourcePrefix
`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) UnsetTranslatedIpv4SourcePrefix()`

UnsetTranslatedIpv4SourcePrefix ensures that no value is present for TranslatedIpv4SourcePrefix, not even an explicit nil
### GetRefType

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetSyncSource

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetSyncSource() string`

GetSyncSource returns the SyncSource field if non-nil, zero value otherwise.

### GetSyncSourceOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetSyncSourceOk() (*string, bool)`

GetSyncSourceOk returns a tuple with the SyncSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSource

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetSyncSource(v string)`

SetSyncSource sets SyncSource field to given value.

### HasSyncSource

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasSyncSource() bool`

HasSyncSource returns a boolean if a field has been set.

### GetInternalId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkRouterNat200ResponseNetworkRouterNAT) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


