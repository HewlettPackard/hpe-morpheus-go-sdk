# GetNetworkRouter200ResponseNetworkRouter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**RouterType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ProviderId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GetNetworkRouter200ResponseNetworkRouterType**](GetNetworkRouter200ResponseNetworkRouterType.md) |  | [optional] 
**NetworkServer** | Pointer to [**GetNetworkRouter200ResponseNetworkRouterNetworkServer**](GetNetworkRouter200ResponseNetworkRouterNetworkServer.md) |  | [optional] 
**Zone** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**ExternalNetwork** | Pointer to **string** |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 
**Interfaces** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Firewall** | Pointer to [**GetNetworkRouter200ResponseNetworkRouterFirewall**](GetNetworkRouter200ResponseNetworkRouterFirewall.md) |  | [optional] 
**Routes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Nats** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Permissions** | Pointer to [**GetNetworkRouter200ResponseNetworkRouterPermissions**](GetNetworkRouter200ResponseNetworkRouterPermissions.md) |  | [optional] 

## Methods

### NewGetNetworkRouter200ResponseNetworkRouter

`func NewGetNetworkRouter200ResponseNetworkRouter() *GetNetworkRouter200ResponseNetworkRouter`

NewGetNetworkRouter200ResponseNetworkRouter instantiates a new GetNetworkRouter200ResponseNetworkRouter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRouter200ResponseNetworkRouterWithDefaults

`func NewGetNetworkRouter200ResponseNetworkRouterWithDefaults() *GetNetworkRouter200ResponseNetworkRouter`

NewGetNetworkRouter200ResponseNetworkRouterWithDefaults instantiates a new GetNetworkRouter200ResponseNetworkRouter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRouterType

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetRouterType() string`

GetRouterType returns the RouterType field if non-nil, zero value otherwise.

### GetRouterTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetRouterTypeOk() (*string, bool)`

GetRouterTypeOk returns a tuple with the RouterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterType

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetRouterType(v string)`

SetRouterType sets RouterType field to given value.

### HasRouterType

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasRouterType() bool`

HasRouterType returns a boolean if a field has been set.

### GetStatus

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalIp

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetType() GetNetworkRouter200ResponseNetworkRouterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetTypeOk() (*GetNetworkRouter200ResponseNetworkRouterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetType(v GetNetworkRouter200ResponseNetworkRouterType)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkServer

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetNetworkServer() GetNetworkRouter200ResponseNetworkRouterNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetNetworkServerOk() (*GetNetworkRouter200ResponseNetworkRouterNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetNetworkServer(v GetNetworkRouter200ResponseNetworkRouterNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetZone

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetZone() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetZoneOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetZone(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetInstance

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetExternalNetwork

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetExternalNetwork() string`

GetExternalNetwork returns the ExternalNetwork field if non-nil, zero value otherwise.

### GetExternalNetworkOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetExternalNetworkOk() (*string, bool)`

GetExternalNetworkOk returns a tuple with the ExternalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNetwork

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetExternalNetwork(v string)`

SetExternalNetwork sets ExternalNetwork field to given value.

### HasExternalNetwork

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasExternalNetwork() bool`

HasExternalNetwork returns a boolean if a field has been set.

### GetSite

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetInterfaces() []map[string]interface{}`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetInterfacesOk() (*[]map[string]interface{}, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetInterfaces(v []map[string]interface{})`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetFirewall

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetFirewall() GetNetworkRouter200ResponseNetworkRouterFirewall`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetFirewallOk() (*GetNetworkRouter200ResponseNetworkRouterFirewall, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetFirewall(v GetNetworkRouter200ResponseNetworkRouterFirewall)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetRoutes

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetRoutes() []map[string]interface{}`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetRoutesOk() (*[]map[string]interface{}, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetRoutes(v []map[string]interface{})`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetNats

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetNats() []map[string]interface{}`

GetNats returns the Nats field if non-nil, zero value otherwise.

### GetNatsOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetNatsOk() (*[]map[string]interface{}, bool)`

GetNatsOk returns a tuple with the Nats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNats

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetNats(v []map[string]interface{})`

SetNats sets Nats field to given value.

### HasNats

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasNats() bool`

HasNats returns a boolean if a field has been set.

### GetPermissions

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetPermissions() GetNetworkRouter200ResponseNetworkRouterPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetNetworkRouter200ResponseNetworkRouter) GetPermissionsOk() (*GetNetworkRouter200ResponseNetworkRouterPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetNetworkRouter200ResponseNetworkRouter) SetPermissions(v GetNetworkRouter200ResponseNetworkRouterPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetNetworkRouter200ResponseNetworkRouter) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


