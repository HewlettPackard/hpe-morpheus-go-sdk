# GetNetworkRoutersNats200ResponseNetworkRouterNATsInner

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


