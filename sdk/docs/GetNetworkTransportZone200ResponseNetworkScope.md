# GetNetworkTransportZone200ResponseNetworkScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**StreamType** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**GetNetworkTransportZone200ResponseNetworkScopeConfig**](GetNetworkTransportZone200ResponseNetworkScopeConfig.md) |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**NetworkServer** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Zone** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewGetNetworkTransportZone200ResponseNetworkScope

`func NewGetNetworkTransportZone200ResponseNetworkScope() *GetNetworkTransportZone200ResponseNetworkScope`

NewGetNetworkTransportZone200ResponseNetworkScope instantiates a new GetNetworkTransportZone200ResponseNetworkScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkTransportZone200ResponseNetworkScopeWithDefaults

`func NewGetNetworkTransportZone200ResponseNetworkScopeWithDefaults() *GetNetworkTransportZone200ResponseNetworkScope`

NewGetNetworkTransportZone200ResponseNetworkScopeWithDefaults instantiates a new GetNetworkTransportZone200ResponseNetworkScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStreamType

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetStreamType() string`

GetStreamType returns the StreamType field if non-nil, zero value otherwise.

### GetStreamTypeOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetStreamTypeOk() (*string, bool)`

GetStreamTypeOk returns a tuple with the StreamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamType

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetStreamType(v string)`

SetStreamType sets StreamType field to given value.

### HasStreamType

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasStreamType() bool`

HasStreamType returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetConfig() GetNetworkTransportZone200ResponseNetworkScopeConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetConfigOk() (*GetNetworkTransportZone200ResponseNetworkScopeConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetConfig(v GetNetworkTransportZone200ResponseNetworkScopeConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOwner

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetNetworkServer() GetAlerts200ResponseAllOfChecksInnerAccount`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetNetworkServerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetNetworkServer(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetZone

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetZone() GetAlerts200ResponseAllOfChecksInnerAccount`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetZoneOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetZone(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetTenants

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetNetworkTransportZone200ResponseNetworkScope) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetNetworkTransportZone200ResponseNetworkScope) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetNetworkTransportZone200ResponseNetworkScope) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


