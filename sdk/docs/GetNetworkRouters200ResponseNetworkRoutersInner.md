# GetNetworkRouters200ResponseNetworkRoutersInner

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
**Type** | Pointer to [**GetNetworkRouters200ResponseNetworkRoutersInnerType**](GetNetworkRouters200ResponseNetworkRoutersInnerType.md) |  | [optional] 
**NetworkServer** | Pointer to **string** |  | [optional] 
**Zone** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**ExternalNetwork** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Site** | Pointer to **string** |  | [optional] 
**Interfaces** | Pointer to [**[]GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner**](GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner.md) |  | [optional] 

## Methods

### NewGetNetworkRouters200ResponseNetworkRoutersInner

`func NewGetNetworkRouters200ResponseNetworkRoutersInner() *GetNetworkRouters200ResponseNetworkRoutersInner`

NewGetNetworkRouters200ResponseNetworkRoutersInner instantiates a new GetNetworkRouters200ResponseNetworkRoutersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRouters200ResponseNetworkRoutersInnerWithDefaults

`func NewGetNetworkRouters200ResponseNetworkRoutersInnerWithDefaults() *GetNetworkRouters200ResponseNetworkRoutersInner`

NewGetNetworkRouters200ResponseNetworkRoutersInnerWithDefaults instantiates a new GetNetworkRouters200ResponseNetworkRoutersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRouterType

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetRouterType() string`

GetRouterType returns the RouterType field if non-nil, zero value otherwise.

### GetRouterTypeOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetRouterTypeOk() (*string, bool)`

GetRouterTypeOk returns a tuple with the RouterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterType

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetRouterType(v string)`

SetRouterType sets RouterType field to given value.

### HasRouterType

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasRouterType() bool`

HasRouterType returns a boolean if a field has been set.

### GetStatus

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalIp

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProviderId

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetType() GetNetworkRouters200ResponseNetworkRoutersInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetTypeOk() (*GetNetworkRouters200ResponseNetworkRoutersInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetType(v GetNetworkRouters200ResponseNetworkRoutersInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkServer

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetNetworkServer() string`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetNetworkServerOk() (*string, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetNetworkServer(v string)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetZone

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetZone() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetZoneOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetZone(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetInstance

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetExternalNetwork

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetExternalNetwork() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetExternalNetwork returns the ExternalNetwork field if non-nil, zero value otherwise.

### GetExternalNetworkOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetExternalNetworkOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetExternalNetworkOk returns a tuple with the ExternalNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNetwork

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetExternalNetwork(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetExternalNetwork sets ExternalNetwork field to given value.

### HasExternalNetwork

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasExternalNetwork() bool`

HasExternalNetwork returns a boolean if a field has been set.

### GetSite

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetInterfaces() []GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) GetInterfacesOk() (*[]GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) SetInterfaces(v []GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetNetworkRouters200ResponseNetworkRoutersInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


