# GetNetworkRoutersNats200ResponseNetworkRouterNATsInner

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

### NewGetNetworkRoutersNats200ResponseNetworkRouterNATsInner

`func NewGetNetworkRoutersNats200ResponseNetworkRouterNATsInner() *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner`

NewGetNetworkRoutersNats200ResponseNetworkRouterNATsInner instantiates a new GetNetworkRoutersNats200ResponseNetworkRouterNATsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRoutersNats200ResponseNetworkRouterNATsInnerWithDefaults

`func NewGetNetworkRoutersNats200ResponseNetworkRouterNATsInnerWithDefaults() *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner`

NewGetNetworkRoutersNats200ResponseNetworkRouterNATsInnerWithDefaults instantiates a new GetNetworkRoutersNats200ResponseNetworkRouterNATsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSourceNetwork

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetSourceNetwork() string`

GetSourceNetwork returns the SourceNetwork field if non-nil, zero value otherwise.

### GetSourceNetworkOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetSourceNetworkOk() (*string, bool)`

GetSourceNetworkOk returns a tuple with the SourceNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetwork

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetSourceNetwork(v string)`

SetSourceNetwork sets SourceNetwork field to given value.

### HasSourceNetwork

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasSourceNetwork() bool`

HasSourceNetwork returns a boolean if a field has been set.

### GetDestinationNetwork

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetDestinationNetwork() string`

GetDestinationNetwork returns the DestinationNetwork field if non-nil, zero value otherwise.

### GetDestinationNetworkOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetDestinationNetworkOk() (*string, bool)`

GetDestinationNetworkOk returns a tuple with the DestinationNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNetwork

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetDestinationNetwork(v string)`

SetDestinationNetwork sets DestinationNetwork field to given value.

### HasDestinationNetwork

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasDestinationNetwork() bool`

HasDestinationNetwork returns a boolean if a field has been set.

### SetDestinationNetworkNil

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetDestinationNetworkNil(b bool)`

 SetDestinationNetworkNil sets the value for DestinationNetwork to be an explicit nil

### UnsetDestinationNetwork
`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) UnsetDestinationNetwork()`

UnsetDestinationNetwork ensures that no value is present for DestinationNetwork, not even an explicit nil
### GetTranslatedNetwork

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetTranslatedNetwork() string`

GetTranslatedNetwork returns the TranslatedNetwork field if non-nil, zero value otherwise.

### GetTranslatedNetworkOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetTranslatedNetworkOk() (*string, bool)`

GetTranslatedNetworkOk returns a tuple with the TranslatedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedNetwork

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetTranslatedNetwork(v string)`

SetTranslatedNetwork sets TranslatedNetwork field to given value.

### HasTranslatedNetwork

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasTranslatedNetwork() bool`

HasTranslatedNetwork returns a boolean if a field has been set.

### GetSourcePorts

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetSourcePorts() string`

GetSourcePorts returns the SourcePorts field if non-nil, zero value otherwise.

### GetSourcePortsOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetSourcePortsOk() (*string, bool)`

GetSourcePortsOk returns a tuple with the SourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePorts

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetSourcePorts(v string)`

SetSourcePorts sets SourcePorts field to given value.

### HasSourcePorts

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasSourcePorts() bool`

HasSourcePorts returns a boolean if a field has been set.

### SetSourcePortsNil

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetSourcePortsNil(b bool)`

 SetSourcePortsNil sets the value for SourcePorts to be an explicit nil

### UnsetSourcePorts
`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) UnsetSourcePorts()`

UnsetSourcePorts ensures that no value is present for SourcePorts, not even an explicit nil
### GetDestinationPorts

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetDestinationPorts() string`

GetDestinationPorts returns the DestinationPorts field if non-nil, zero value otherwise.

### GetDestinationPortsOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetDestinationPortsOk() (*string, bool)`

GetDestinationPortsOk returns a tuple with the DestinationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPorts

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetDestinationPorts(v string)`

SetDestinationPorts sets DestinationPorts field to given value.

### HasDestinationPorts

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasDestinationPorts() bool`

HasDestinationPorts returns a boolean if a field has been set.

### SetDestinationPortsNil

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetDestinationPortsNil(b bool)`

 SetDestinationPortsNil sets the value for DestinationPorts to be an explicit nil

### UnsetDestinationPorts
`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) UnsetDestinationPorts()`

UnsetDestinationPorts ensures that no value is present for DestinationPorts, not even an explicit nil
### GetTranslatedPorts

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetTranslatedPorts() string`

GetTranslatedPorts returns the TranslatedPorts field if non-nil, zero value otherwise.

### GetTranslatedPortsOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetTranslatedPortsOk() (*string, bool)`

GetTranslatedPortsOk returns a tuple with the TranslatedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedPorts

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetTranslatedPorts(v string)`

SetTranslatedPorts sets TranslatedPorts field to given value.

### HasTranslatedPorts

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasTranslatedPorts() bool`

HasTranslatedPorts returns a boolean if a field has been set.

### SetTranslatedPortsNil

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetTranslatedPortsNil(b bool)`

 SetTranslatedPortsNil sets the value for TranslatedPorts to be an explicit nil

### UnsetTranslatedPorts
`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) UnsetTranslatedPorts()`

UnsetTranslatedPorts ensures that no value is present for TranslatedPorts, not even an explicit nil
### GetPriority

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProtocol

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetMatchIpv6DestinationPrefix

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetMatchIpv6DestinationPrefix() string`

GetMatchIpv6DestinationPrefix returns the MatchIpv6DestinationPrefix field if non-nil, zero value otherwise.

### GetMatchIpv6DestinationPrefixOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetMatchIpv6DestinationPrefixOk() (*string, bool)`

GetMatchIpv6DestinationPrefixOk returns a tuple with the MatchIpv6DestinationPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpv6DestinationPrefix

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetMatchIpv6DestinationPrefix(v string)`

SetMatchIpv6DestinationPrefix sets MatchIpv6DestinationPrefix field to given value.

### HasMatchIpv6DestinationPrefix

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasMatchIpv6DestinationPrefix() bool`

HasMatchIpv6DestinationPrefix returns a boolean if a field has been set.

### SetMatchIpv6DestinationPrefixNil

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetMatchIpv6DestinationPrefixNil(b bool)`

 SetMatchIpv6DestinationPrefixNil sets the value for MatchIpv6DestinationPrefix to be an explicit nil

### UnsetMatchIpv6DestinationPrefix
`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) UnsetMatchIpv6DestinationPrefix()`

UnsetMatchIpv6DestinationPrefix ensures that no value is present for MatchIpv6DestinationPrefix, not even an explicit nil
### GetTranslatedIpv4SourcePrefix

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetTranslatedIpv4SourcePrefix() string`

GetTranslatedIpv4SourcePrefix returns the TranslatedIpv4SourcePrefix field if non-nil, zero value otherwise.

### GetTranslatedIpv4SourcePrefixOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetTranslatedIpv4SourcePrefixOk() (*string, bool)`

GetTranslatedIpv4SourcePrefixOk returns a tuple with the TranslatedIpv4SourcePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedIpv4SourcePrefix

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetTranslatedIpv4SourcePrefix(v string)`

SetTranslatedIpv4SourcePrefix sets TranslatedIpv4SourcePrefix field to given value.

### HasTranslatedIpv4SourcePrefix

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasTranslatedIpv4SourcePrefix() bool`

HasTranslatedIpv4SourcePrefix returns a boolean if a field has been set.

### SetTranslatedIpv4SourcePrefixNil

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetTranslatedIpv4SourcePrefixNil(b bool)`

 SetTranslatedIpv4SourcePrefixNil sets the value for TranslatedIpv4SourcePrefix to be an explicit nil

### UnsetTranslatedIpv4SourcePrefix
`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) UnsetTranslatedIpv4SourcePrefix()`

UnsetTranslatedIpv4SourcePrefix ensures that no value is present for TranslatedIpv4SourcePrefix, not even an explicit nil
### GetRefType

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetSyncSource

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetSyncSource() string`

GetSyncSource returns the SyncSource field if non-nil, zero value otherwise.

### GetSyncSourceOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetSyncSourceOk() (*string, bool)`

GetSyncSourceOk returns a tuple with the SyncSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSource

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetSyncSource(v string)`

SetSyncSource sets SyncSource field to given value.

### HasSyncSource

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasSyncSource() bool`

HasSyncSource returns a boolean if a field has been set.

### GetInternalId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkRoutersNats200ResponseNetworkRouterNATsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


