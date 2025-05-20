# ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**VipSticky** | Pointer to **string** |  | [optional] 
**VipBalance** | Pointer to **string** |  | [optional] 
**AllowNat** | Pointer to **string** |  | [optional] 
**AllowSnat** | Pointer to **string** |  | [optional] 
**VipClientIpMode** | Pointer to **string** |  | [optional] 
**VipServerIpMode** | Pointer to **string** |  | [optional] 
**MinActive** | Pointer to **int64** |  | [optional] 
**MinInService** | Pointer to **string** |  | [optional] 
**MinUpMonitor** | Pointer to **string** |  | [optional] 
**MinUpAction** | Pointer to **string** |  | [optional] 
**MaxQueueDepth** | Pointer to **string** |  | [optional] 
**MaxQueueTime** | Pointer to **string** |  | [optional] 
**NumberActive** | Pointer to **int64** |  | [optional] 
**NumberInService** | Pointer to **int64** |  | [optional] 
**HealthScore** | Pointer to **int64** |  | [optional] 
**PerformanceScore** | Pointer to **int64** |  | [optional] 
**HealthPenalty** | Pointer to **int64** |  | [optional] 
**SecurityPenalty** | Pointer to **int64** |  | [optional] 
**ErrorPenalty** | Pointer to **int64** |  | [optional] 
**DownAction** | Pointer to **string** |  | [optional] 
**RampTime** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**PortType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Monitors** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Members** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner

`func NewListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner() *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner`

NewListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner instantiates a new ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInnerWithDefaults

`func NewListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInnerWithDefaults() *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner`

NewListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInnerWithDefaults instantiates a new ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetLoadBalancer() ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetLoadBalancerOk() (*ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetLoadBalancer(v ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVisibility

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEnabled

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVipSticky

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetVipSticky() string`

GetVipSticky returns the VipSticky field if non-nil, zero value otherwise.

### GetVipStickyOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetVipStickyOk() (*string, bool)`

GetVipStickyOk returns a tuple with the VipSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSticky

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetVipSticky(v string)`

SetVipSticky sets VipSticky field to given value.

### HasVipSticky

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasVipSticky() bool`

HasVipSticky returns a boolean if a field has been set.

### GetVipBalance

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetVipBalance() string`

GetVipBalance returns the VipBalance field if non-nil, zero value otherwise.

### GetVipBalanceOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetVipBalanceOk() (*string, bool)`

GetVipBalanceOk returns a tuple with the VipBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipBalance

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetVipBalance(v string)`

SetVipBalance sets VipBalance field to given value.

### HasVipBalance

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasVipBalance() bool`

HasVipBalance returns a boolean if a field has been set.

### GetAllowNat

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetAllowNat() string`

GetAllowNat returns the AllowNat field if non-nil, zero value otherwise.

### GetAllowNatOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetAllowNatOk() (*string, bool)`

GetAllowNatOk returns a tuple with the AllowNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNat

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetAllowNat(v string)`

SetAllowNat sets AllowNat field to given value.

### HasAllowNat

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasAllowNat() bool`

HasAllowNat returns a boolean if a field has been set.

### GetAllowSnat

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetAllowSnat() string`

GetAllowSnat returns the AllowSnat field if non-nil, zero value otherwise.

### GetAllowSnatOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetAllowSnatOk() (*string, bool)`

GetAllowSnatOk returns a tuple with the AllowSnat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSnat

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetAllowSnat(v string)`

SetAllowSnat sets AllowSnat field to given value.

### HasAllowSnat

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasAllowSnat() bool`

HasAllowSnat returns a boolean if a field has been set.

### GetVipClientIpMode

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetVipClientIpMode() string`

GetVipClientIpMode returns the VipClientIpMode field if non-nil, zero value otherwise.

### GetVipClientIpModeOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetVipClientIpModeOk() (*string, bool)`

GetVipClientIpModeOk returns a tuple with the VipClientIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipClientIpMode

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetVipClientIpMode(v string)`

SetVipClientIpMode sets VipClientIpMode field to given value.

### HasVipClientIpMode

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasVipClientIpMode() bool`

HasVipClientIpMode returns a boolean if a field has been set.

### GetVipServerIpMode

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetVipServerIpMode() string`

GetVipServerIpMode returns the VipServerIpMode field if non-nil, zero value otherwise.

### GetVipServerIpModeOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetVipServerIpModeOk() (*string, bool)`

GetVipServerIpModeOk returns a tuple with the VipServerIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipServerIpMode

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetVipServerIpMode(v string)`

SetVipServerIpMode sets VipServerIpMode field to given value.

### HasVipServerIpMode

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasVipServerIpMode() bool`

HasVipServerIpMode returns a boolean if a field has been set.

### GetMinActive

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMinActive() int64`

GetMinActive returns the MinActive field if non-nil, zero value otherwise.

### GetMinActiveOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMinActiveOk() (*int64, bool)`

GetMinActiveOk returns a tuple with the MinActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinActive

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetMinActive(v int64)`

SetMinActive sets MinActive field to given value.

### HasMinActive

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasMinActive() bool`

HasMinActive returns a boolean if a field has been set.

### GetMinInService

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMinInService() string`

GetMinInService returns the MinInService field if non-nil, zero value otherwise.

### GetMinInServiceOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMinInServiceOk() (*string, bool)`

GetMinInServiceOk returns a tuple with the MinInService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInService

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetMinInService(v string)`

SetMinInService sets MinInService field to given value.

### HasMinInService

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasMinInService() bool`

HasMinInService returns a boolean if a field has been set.

### GetMinUpMonitor

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMinUpMonitor() string`

GetMinUpMonitor returns the MinUpMonitor field if non-nil, zero value otherwise.

### GetMinUpMonitorOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMinUpMonitorOk() (*string, bool)`

GetMinUpMonitorOk returns a tuple with the MinUpMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpMonitor

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetMinUpMonitor(v string)`

SetMinUpMonitor sets MinUpMonitor field to given value.

### HasMinUpMonitor

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasMinUpMonitor() bool`

HasMinUpMonitor returns a boolean if a field has been set.

### GetMinUpAction

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMinUpAction() string`

GetMinUpAction returns the MinUpAction field if non-nil, zero value otherwise.

### GetMinUpActionOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMinUpActionOk() (*string, bool)`

GetMinUpActionOk returns a tuple with the MinUpAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpAction

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetMinUpAction(v string)`

SetMinUpAction sets MinUpAction field to given value.

### HasMinUpAction

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasMinUpAction() bool`

HasMinUpAction returns a boolean if a field has been set.

### GetMaxQueueDepth

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMaxQueueDepth() string`

GetMaxQueueDepth returns the MaxQueueDepth field if non-nil, zero value otherwise.

### GetMaxQueueDepthOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMaxQueueDepthOk() (*string, bool)`

GetMaxQueueDepthOk returns a tuple with the MaxQueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueueDepth

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetMaxQueueDepth(v string)`

SetMaxQueueDepth sets MaxQueueDepth field to given value.

### HasMaxQueueDepth

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasMaxQueueDepth() bool`

HasMaxQueueDepth returns a boolean if a field has been set.

### GetMaxQueueTime

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMaxQueueTime() string`

GetMaxQueueTime returns the MaxQueueTime field if non-nil, zero value otherwise.

### GetMaxQueueTimeOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMaxQueueTimeOk() (*string, bool)`

GetMaxQueueTimeOk returns a tuple with the MaxQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueueTime

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetMaxQueueTime(v string)`

SetMaxQueueTime sets MaxQueueTime field to given value.

### HasMaxQueueTime

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasMaxQueueTime() bool`

HasMaxQueueTime returns a boolean if a field has been set.

### GetNumberActive

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetNumberActive() int64`

GetNumberActive returns the NumberActive field if non-nil, zero value otherwise.

### GetNumberActiveOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetNumberActiveOk() (*int64, bool)`

GetNumberActiveOk returns a tuple with the NumberActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberActive

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetNumberActive(v int64)`

SetNumberActive sets NumberActive field to given value.

### HasNumberActive

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasNumberActive() bool`

HasNumberActive returns a boolean if a field has been set.

### GetNumberInService

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetNumberInService() int64`

GetNumberInService returns the NumberInService field if non-nil, zero value otherwise.

### GetNumberInServiceOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetNumberInServiceOk() (*int64, bool)`

GetNumberInServiceOk returns a tuple with the NumberInService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberInService

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetNumberInService(v int64)`

SetNumberInService sets NumberInService field to given value.

### HasNumberInService

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasNumberInService() bool`

HasNumberInService returns a boolean if a field has been set.

### GetHealthScore

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetHealthScore() int64`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetHealthScoreOk() (*int64, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetHealthScore(v int64)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetPerformanceScore

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetPerformanceScore() int64`

GetPerformanceScore returns the PerformanceScore field if non-nil, zero value otherwise.

### GetPerformanceScoreOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetPerformanceScoreOk() (*int64, bool)`

GetPerformanceScoreOk returns a tuple with the PerformanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceScore

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetPerformanceScore(v int64)`

SetPerformanceScore sets PerformanceScore field to given value.

### HasPerformanceScore

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasPerformanceScore() bool`

HasPerformanceScore returns a boolean if a field has been set.

### GetHealthPenalty

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetHealthPenalty() int64`

GetHealthPenalty returns the HealthPenalty field if non-nil, zero value otherwise.

### GetHealthPenaltyOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetHealthPenaltyOk() (*int64, bool)`

GetHealthPenaltyOk returns a tuple with the HealthPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthPenalty

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetHealthPenalty(v int64)`

SetHealthPenalty sets HealthPenalty field to given value.

### HasHealthPenalty

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasHealthPenalty() bool`

HasHealthPenalty returns a boolean if a field has been set.

### GetSecurityPenalty

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetSecurityPenalty() int64`

GetSecurityPenalty returns the SecurityPenalty field if non-nil, zero value otherwise.

### GetSecurityPenaltyOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetSecurityPenaltyOk() (*int64, bool)`

GetSecurityPenaltyOk returns a tuple with the SecurityPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPenalty

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetSecurityPenalty(v int64)`

SetSecurityPenalty sets SecurityPenalty field to given value.

### HasSecurityPenalty

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasSecurityPenalty() bool`

HasSecurityPenalty returns a boolean if a field has been set.

### GetErrorPenalty

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetErrorPenalty() int64`

GetErrorPenalty returns the ErrorPenalty field if non-nil, zero value otherwise.

### GetErrorPenaltyOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetErrorPenaltyOk() (*int64, bool)`

GetErrorPenaltyOk returns a tuple with the ErrorPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPenalty

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetErrorPenalty(v int64)`

SetErrorPenalty sets ErrorPenalty field to given value.

### HasErrorPenalty

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasErrorPenalty() bool`

HasErrorPenalty returns a boolean if a field has been set.

### GetDownAction

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetDownAction() string`

GetDownAction returns the DownAction field if non-nil, zero value otherwise.

### GetDownActionOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetDownActionOk() (*string, bool)`

GetDownActionOk returns a tuple with the DownAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownAction

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetDownAction(v string)`

SetDownAction sets DownAction field to given value.

### HasDownAction

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasDownAction() bool`

HasDownAction returns a boolean if a field has been set.

### GetRampTime

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetRampTime() string`

GetRampTime returns the RampTime field if non-nil, zero value otherwise.

### GetRampTimeOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetRampTimeOk() (*string, bool)`

GetRampTimeOk returns a tuple with the RampTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRampTime

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetRampTime(v string)`

SetRampTime sets RampTime field to given value.

### HasRampTime

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasRampTime() bool`

HasRampTime returns a boolean if a field has been set.

### GetPort

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortType

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetStatus

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNodes

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetNodes() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetNodesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetNodes(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetMonitors

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMonitors() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMonitorsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetMonitors(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetMembers

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMembers() []map[string]interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetMembersOk() (*[]map[string]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetMembers(v []map[string]interface{})`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetConfig

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


