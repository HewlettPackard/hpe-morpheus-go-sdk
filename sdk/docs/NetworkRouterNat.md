# NetworkRouterNat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**SourceNetwork** | Pointer to **string** |  | [optional] 
**DestinationNetwork** | Pointer to **string** |  | [optional] 
**TranslatedNetwork** | Pointer to **string** |  | [optional] 
**SourcePorts** | Pointer to **string** |  | [optional] 
**DestinationPorts** | Pointer to **string** |  | [optional] 
**TranslatedPorts** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**MatchIpv6DestinationPrefix** | Pointer to **string** |  | [optional] 
**TranslatedIpv4SourcePrefix** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**SyncSource** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewNetworkRouterNat

`func NewNetworkRouterNat() *NetworkRouterNat`

NewNetworkRouterNat instantiates a new NetworkRouterNat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkRouterNatWithDefaults

`func NewNetworkRouterNatWithDefaults() *NetworkRouterNat`

NewNetworkRouterNatWithDefaults instantiates a new NetworkRouterNat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkRouterNat) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkRouterNat) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkRouterNat) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkRouterNat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkRouterNat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkRouterNat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkRouterNat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkRouterNat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkRouterNat) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkRouterNat) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkRouterNat) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkRouterNat) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkRouterNat) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkRouterNat) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkRouterNat) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkRouterNat) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSourceNetwork

`func (o *NetworkRouterNat) GetSourceNetwork() string`

GetSourceNetwork returns the SourceNetwork field if non-nil, zero value otherwise.

### GetSourceNetworkOk

`func (o *NetworkRouterNat) GetSourceNetworkOk() (*string, bool)`

GetSourceNetworkOk returns a tuple with the SourceNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetwork

`func (o *NetworkRouterNat) SetSourceNetwork(v string)`

SetSourceNetwork sets SourceNetwork field to given value.

### HasSourceNetwork

`func (o *NetworkRouterNat) HasSourceNetwork() bool`

HasSourceNetwork returns a boolean if a field has been set.

### GetDestinationNetwork

`func (o *NetworkRouterNat) GetDestinationNetwork() string`

GetDestinationNetwork returns the DestinationNetwork field if non-nil, zero value otherwise.

### GetDestinationNetworkOk

`func (o *NetworkRouterNat) GetDestinationNetworkOk() (*string, bool)`

GetDestinationNetworkOk returns a tuple with the DestinationNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNetwork

`func (o *NetworkRouterNat) SetDestinationNetwork(v string)`

SetDestinationNetwork sets DestinationNetwork field to given value.

### HasDestinationNetwork

`func (o *NetworkRouterNat) HasDestinationNetwork() bool`

HasDestinationNetwork returns a boolean if a field has been set.

### GetTranslatedNetwork

`func (o *NetworkRouterNat) GetTranslatedNetwork() string`

GetTranslatedNetwork returns the TranslatedNetwork field if non-nil, zero value otherwise.

### GetTranslatedNetworkOk

`func (o *NetworkRouterNat) GetTranslatedNetworkOk() (*string, bool)`

GetTranslatedNetworkOk returns a tuple with the TranslatedNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedNetwork

`func (o *NetworkRouterNat) SetTranslatedNetwork(v string)`

SetTranslatedNetwork sets TranslatedNetwork field to given value.

### HasTranslatedNetwork

`func (o *NetworkRouterNat) HasTranslatedNetwork() bool`

HasTranslatedNetwork returns a boolean if a field has been set.

### GetSourcePorts

`func (o *NetworkRouterNat) GetSourcePorts() string`

GetSourcePorts returns the SourcePorts field if non-nil, zero value otherwise.

### GetSourcePortsOk

`func (o *NetworkRouterNat) GetSourcePortsOk() (*string, bool)`

GetSourcePortsOk returns a tuple with the SourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePorts

`func (o *NetworkRouterNat) SetSourcePorts(v string)`

SetSourcePorts sets SourcePorts field to given value.

### HasSourcePorts

`func (o *NetworkRouterNat) HasSourcePorts() bool`

HasSourcePorts returns a boolean if a field has been set.

### GetDestinationPorts

`func (o *NetworkRouterNat) GetDestinationPorts() string`

GetDestinationPorts returns the DestinationPorts field if non-nil, zero value otherwise.

### GetDestinationPortsOk

`func (o *NetworkRouterNat) GetDestinationPortsOk() (*string, bool)`

GetDestinationPortsOk returns a tuple with the DestinationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPorts

`func (o *NetworkRouterNat) SetDestinationPorts(v string)`

SetDestinationPorts sets DestinationPorts field to given value.

### HasDestinationPorts

`func (o *NetworkRouterNat) HasDestinationPorts() bool`

HasDestinationPorts returns a boolean if a field has been set.

### GetTranslatedPorts

`func (o *NetworkRouterNat) GetTranslatedPorts() string`

GetTranslatedPorts returns the TranslatedPorts field if non-nil, zero value otherwise.

### GetTranslatedPortsOk

`func (o *NetworkRouterNat) GetTranslatedPortsOk() (*string, bool)`

GetTranslatedPortsOk returns a tuple with the TranslatedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedPorts

`func (o *NetworkRouterNat) SetTranslatedPorts(v string)`

SetTranslatedPorts sets TranslatedPorts field to given value.

### HasTranslatedPorts

`func (o *NetworkRouterNat) HasTranslatedPorts() bool`

HasTranslatedPorts returns a boolean if a field has been set.

### GetPriority

`func (o *NetworkRouterNat) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NetworkRouterNat) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NetworkRouterNat) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NetworkRouterNat) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProtocol

`func (o *NetworkRouterNat) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NetworkRouterNat) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NetworkRouterNat) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NetworkRouterNat) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetMatchIpv6DestinationPrefix

`func (o *NetworkRouterNat) GetMatchIpv6DestinationPrefix() string`

GetMatchIpv6DestinationPrefix returns the MatchIpv6DestinationPrefix field if non-nil, zero value otherwise.

### GetMatchIpv6DestinationPrefixOk

`func (o *NetworkRouterNat) GetMatchIpv6DestinationPrefixOk() (*string, bool)`

GetMatchIpv6DestinationPrefixOk returns a tuple with the MatchIpv6DestinationPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchIpv6DestinationPrefix

`func (o *NetworkRouterNat) SetMatchIpv6DestinationPrefix(v string)`

SetMatchIpv6DestinationPrefix sets MatchIpv6DestinationPrefix field to given value.

### HasMatchIpv6DestinationPrefix

`func (o *NetworkRouterNat) HasMatchIpv6DestinationPrefix() bool`

HasMatchIpv6DestinationPrefix returns a boolean if a field has been set.

### GetTranslatedIpv4SourcePrefix

`func (o *NetworkRouterNat) GetTranslatedIpv4SourcePrefix() string`

GetTranslatedIpv4SourcePrefix returns the TranslatedIpv4SourcePrefix field if non-nil, zero value otherwise.

### GetTranslatedIpv4SourcePrefixOk

`func (o *NetworkRouterNat) GetTranslatedIpv4SourcePrefixOk() (*string, bool)`

GetTranslatedIpv4SourcePrefixOk returns a tuple with the TranslatedIpv4SourcePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedIpv4SourcePrefix

`func (o *NetworkRouterNat) SetTranslatedIpv4SourcePrefix(v string)`

SetTranslatedIpv4SourcePrefix sets TranslatedIpv4SourcePrefix field to given value.

### HasTranslatedIpv4SourcePrefix

`func (o *NetworkRouterNat) HasTranslatedIpv4SourcePrefix() bool`

HasTranslatedIpv4SourcePrefix returns a boolean if a field has been set.

### GetRefType

`func (o *NetworkRouterNat) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *NetworkRouterNat) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *NetworkRouterNat) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *NetworkRouterNat) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *NetworkRouterNat) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *NetworkRouterNat) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *NetworkRouterNat) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *NetworkRouterNat) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSyncSource

`func (o *NetworkRouterNat) GetSyncSource() string`

GetSyncSource returns the SyncSource field if non-nil, zero value otherwise.

### GetSyncSourceOk

`func (o *NetworkRouterNat) GetSyncSourceOk() (*string, bool)`

GetSyncSourceOk returns a tuple with the SyncSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncSource

`func (o *NetworkRouterNat) SetSyncSource(v string)`

SetSyncSource sets SyncSource field to given value.

### HasSyncSource

`func (o *NetworkRouterNat) HasSyncSource() bool`

HasSyncSource returns a boolean if a field has been set.

### GetInternalId

`func (o *NetworkRouterNat) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *NetworkRouterNat) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *NetworkRouterNat) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *NetworkRouterNat) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *NetworkRouterNat) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetworkRouterNat) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetworkRouterNat) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NetworkRouterNat) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProviderId

`func (o *NetworkRouterNat) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *NetworkRouterNat) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *NetworkRouterNat) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *NetworkRouterNat) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetDateCreated

`func (o *NetworkRouterNat) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NetworkRouterNat) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NetworkRouterNat) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NetworkRouterNat) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NetworkRouterNat) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NetworkRouterNat) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NetworkRouterNat) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NetworkRouterNat) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


