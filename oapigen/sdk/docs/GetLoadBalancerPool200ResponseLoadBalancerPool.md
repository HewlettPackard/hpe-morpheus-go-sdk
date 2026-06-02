# GetLoadBalancerPool200ResponseLoadBalancerPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**GetLoadBalancerPool200ResponseLoadBalancerPoolLoadBalancer**](GetLoadBalancerPool200ResponseLoadBalancerPoolLoadBalancer.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**VipSticky** | Pointer to **NullableString** |  | [optional] 
**VipBalance** | Pointer to **string** |  | [optional] 
**AllowNat** | Pointer to **NullableString** |  | [optional] 
**AllowSnat** | Pointer to **NullableString** |  | [optional] 
**VipClientIpMode** | Pointer to **NullableString** |  | [optional] 
**VipServerIpMode** | Pointer to **NullableString** |  | [optional] 
**MinActive** | Pointer to **int64** |  | [optional] 
**MinInService** | Pointer to **NullableString** |  | [optional] 
**MinUpMonitor** | Pointer to **NullableString** |  | [optional] 
**MinUpAction** | Pointer to **NullableString** |  | [optional] 
**MaxQueueDepth** | Pointer to **NullableString** |  | [optional] 
**MaxQueueTime** | Pointer to **NullableString** |  | [optional] 
**NumberActive** | Pointer to **int64** |  | [optional] 
**NumberInService** | Pointer to **int64** |  | [optional] 
**HealthScore** | Pointer to **int64** |  | [optional] 
**PerformanceScore** | Pointer to **int64** |  | [optional] 
**HealthPenalty** | Pointer to **int64** |  | [optional] 
**SecurityPenalty** | Pointer to **int64** |  | [optional] 
**ErrorPenalty** | Pointer to **int64** |  | [optional] 
**DownAction** | Pointer to **NullableString** |  | [optional] 
**RampTime** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableString** |  | [optional] 
**PortType** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]GetLoadBalancerPool200ResponseLoadBalancerPoolNodesInner**](GetLoadBalancerPool200ResponseLoadBalancerPoolNodesInner.md) |  | [optional] 
**Monitors** | Pointer to [**[]GetLoadBalancerPool200ResponseLoadBalancerPoolMonitorsInner**](GetLoadBalancerPool200ResponseLoadBalancerPoolMonitorsInner.md) |  | [optional] 
**Members** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetLoadBalancerPool200ResponseLoadBalancerPool

`func NewGetLoadBalancerPool200ResponseLoadBalancerPool() *GetLoadBalancerPool200ResponseLoadBalancerPool`

NewGetLoadBalancerPool200ResponseLoadBalancerPool instantiates a new GetLoadBalancerPool200ResponseLoadBalancerPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetLoadBalancer() GetLoadBalancerPool200ResponseLoadBalancerPoolLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetLoadBalancerOk() (*GetLoadBalancerPool200ResponseLoadBalancerPoolLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetLoadBalancer(v GetLoadBalancerPool200ResponseLoadBalancerPoolLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVisibility

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEnabled

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVipSticky

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetVipSticky() string`

GetVipSticky returns the VipSticky field if non-nil, zero value otherwise.

### GetVipStickyOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetVipStickyOk() (*string, bool)`

GetVipStickyOk returns a tuple with the VipSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSticky

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetVipSticky(v string)`

SetVipSticky sets VipSticky field to given value.

### HasVipSticky

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasVipSticky() bool`

HasVipSticky returns a boolean if a field has been set.

### SetVipStickyNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetVipStickyNil(b bool)`

 SetVipStickyNil sets the value for VipSticky to be an explicit nil

### UnsetVipSticky
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetVipSticky()`

UnsetVipSticky ensures that no value is present for VipSticky, not even an explicit nil
### GetVipBalance

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetVipBalance() string`

GetVipBalance returns the VipBalance field if non-nil, zero value otherwise.

### GetVipBalanceOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetVipBalanceOk() (*string, bool)`

GetVipBalanceOk returns a tuple with the VipBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipBalance

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetVipBalance(v string)`

SetVipBalance sets VipBalance field to given value.

### HasVipBalance

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasVipBalance() bool`

HasVipBalance returns a boolean if a field has been set.

### GetAllowNat

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetAllowNat() string`

GetAllowNat returns the AllowNat field if non-nil, zero value otherwise.

### GetAllowNatOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetAllowNatOk() (*string, bool)`

GetAllowNatOk returns a tuple with the AllowNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNat

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetAllowNat(v string)`

SetAllowNat sets AllowNat field to given value.

### HasAllowNat

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasAllowNat() bool`

HasAllowNat returns a boolean if a field has been set.

### SetAllowNatNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetAllowNatNil(b bool)`

 SetAllowNatNil sets the value for AllowNat to be an explicit nil

### UnsetAllowNat
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetAllowNat()`

UnsetAllowNat ensures that no value is present for AllowNat, not even an explicit nil
### GetAllowSnat

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetAllowSnat() string`

GetAllowSnat returns the AllowSnat field if non-nil, zero value otherwise.

### GetAllowSnatOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetAllowSnatOk() (*string, bool)`

GetAllowSnatOk returns a tuple with the AllowSnat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSnat

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetAllowSnat(v string)`

SetAllowSnat sets AllowSnat field to given value.

### HasAllowSnat

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasAllowSnat() bool`

HasAllowSnat returns a boolean if a field has been set.

### SetAllowSnatNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetAllowSnatNil(b bool)`

 SetAllowSnatNil sets the value for AllowSnat to be an explicit nil

### UnsetAllowSnat
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetAllowSnat()`

UnsetAllowSnat ensures that no value is present for AllowSnat, not even an explicit nil
### GetVipClientIpMode

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetVipClientIpMode() string`

GetVipClientIpMode returns the VipClientIpMode field if non-nil, zero value otherwise.

### GetVipClientIpModeOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetVipClientIpModeOk() (*string, bool)`

GetVipClientIpModeOk returns a tuple with the VipClientIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipClientIpMode

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetVipClientIpMode(v string)`

SetVipClientIpMode sets VipClientIpMode field to given value.

### HasVipClientIpMode

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasVipClientIpMode() bool`

HasVipClientIpMode returns a boolean if a field has been set.

### SetVipClientIpModeNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetVipClientIpModeNil(b bool)`

 SetVipClientIpModeNil sets the value for VipClientIpMode to be an explicit nil

### UnsetVipClientIpMode
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetVipClientIpMode()`

UnsetVipClientIpMode ensures that no value is present for VipClientIpMode, not even an explicit nil
### GetVipServerIpMode

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetVipServerIpMode() string`

GetVipServerIpMode returns the VipServerIpMode field if non-nil, zero value otherwise.

### GetVipServerIpModeOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetVipServerIpModeOk() (*string, bool)`

GetVipServerIpModeOk returns a tuple with the VipServerIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipServerIpMode

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetVipServerIpMode(v string)`

SetVipServerIpMode sets VipServerIpMode field to given value.

### HasVipServerIpMode

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasVipServerIpMode() bool`

HasVipServerIpMode returns a boolean if a field has been set.

### SetVipServerIpModeNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetVipServerIpModeNil(b bool)`

 SetVipServerIpModeNil sets the value for VipServerIpMode to be an explicit nil

### UnsetVipServerIpMode
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetVipServerIpMode()`

UnsetVipServerIpMode ensures that no value is present for VipServerIpMode, not even an explicit nil
### GetMinActive

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMinActive() int64`

GetMinActive returns the MinActive field if non-nil, zero value otherwise.

### GetMinActiveOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMinActiveOk() (*int64, bool)`

GetMinActiveOk returns a tuple with the MinActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinActive

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMinActive(v int64)`

SetMinActive sets MinActive field to given value.

### HasMinActive

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasMinActive() bool`

HasMinActive returns a boolean if a field has been set.

### GetMinInService

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMinInService() string`

GetMinInService returns the MinInService field if non-nil, zero value otherwise.

### GetMinInServiceOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMinInServiceOk() (*string, bool)`

GetMinInServiceOk returns a tuple with the MinInService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInService

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMinInService(v string)`

SetMinInService sets MinInService field to given value.

### HasMinInService

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasMinInService() bool`

HasMinInService returns a boolean if a field has been set.

### SetMinInServiceNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMinInServiceNil(b bool)`

 SetMinInServiceNil sets the value for MinInService to be an explicit nil

### UnsetMinInService
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetMinInService()`

UnsetMinInService ensures that no value is present for MinInService, not even an explicit nil
### GetMinUpMonitor

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMinUpMonitor() string`

GetMinUpMonitor returns the MinUpMonitor field if non-nil, zero value otherwise.

### GetMinUpMonitorOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMinUpMonitorOk() (*string, bool)`

GetMinUpMonitorOk returns a tuple with the MinUpMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpMonitor

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMinUpMonitor(v string)`

SetMinUpMonitor sets MinUpMonitor field to given value.

### HasMinUpMonitor

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasMinUpMonitor() bool`

HasMinUpMonitor returns a boolean if a field has been set.

### SetMinUpMonitorNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMinUpMonitorNil(b bool)`

 SetMinUpMonitorNil sets the value for MinUpMonitor to be an explicit nil

### UnsetMinUpMonitor
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetMinUpMonitor()`

UnsetMinUpMonitor ensures that no value is present for MinUpMonitor, not even an explicit nil
### GetMinUpAction

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMinUpAction() string`

GetMinUpAction returns the MinUpAction field if non-nil, zero value otherwise.

### GetMinUpActionOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMinUpActionOk() (*string, bool)`

GetMinUpActionOk returns a tuple with the MinUpAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpAction

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMinUpAction(v string)`

SetMinUpAction sets MinUpAction field to given value.

### HasMinUpAction

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasMinUpAction() bool`

HasMinUpAction returns a boolean if a field has been set.

### SetMinUpActionNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMinUpActionNil(b bool)`

 SetMinUpActionNil sets the value for MinUpAction to be an explicit nil

### UnsetMinUpAction
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetMinUpAction()`

UnsetMinUpAction ensures that no value is present for MinUpAction, not even an explicit nil
### GetMaxQueueDepth

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMaxQueueDepth() string`

GetMaxQueueDepth returns the MaxQueueDepth field if non-nil, zero value otherwise.

### GetMaxQueueDepthOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMaxQueueDepthOk() (*string, bool)`

GetMaxQueueDepthOk returns a tuple with the MaxQueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueueDepth

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMaxQueueDepth(v string)`

SetMaxQueueDepth sets MaxQueueDepth field to given value.

### HasMaxQueueDepth

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasMaxQueueDepth() bool`

HasMaxQueueDepth returns a boolean if a field has been set.

### SetMaxQueueDepthNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMaxQueueDepthNil(b bool)`

 SetMaxQueueDepthNil sets the value for MaxQueueDepth to be an explicit nil

### UnsetMaxQueueDepth
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetMaxQueueDepth()`

UnsetMaxQueueDepth ensures that no value is present for MaxQueueDepth, not even an explicit nil
### GetMaxQueueTime

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMaxQueueTime() string`

GetMaxQueueTime returns the MaxQueueTime field if non-nil, zero value otherwise.

### GetMaxQueueTimeOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMaxQueueTimeOk() (*string, bool)`

GetMaxQueueTimeOk returns a tuple with the MaxQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueueTime

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMaxQueueTime(v string)`

SetMaxQueueTime sets MaxQueueTime field to given value.

### HasMaxQueueTime

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasMaxQueueTime() bool`

HasMaxQueueTime returns a boolean if a field has been set.

### SetMaxQueueTimeNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMaxQueueTimeNil(b bool)`

 SetMaxQueueTimeNil sets the value for MaxQueueTime to be an explicit nil

### UnsetMaxQueueTime
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetMaxQueueTime()`

UnsetMaxQueueTime ensures that no value is present for MaxQueueTime, not even an explicit nil
### GetNumberActive

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetNumberActive() int64`

GetNumberActive returns the NumberActive field if non-nil, zero value otherwise.

### GetNumberActiveOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetNumberActiveOk() (*int64, bool)`

GetNumberActiveOk returns a tuple with the NumberActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberActive

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetNumberActive(v int64)`

SetNumberActive sets NumberActive field to given value.

### HasNumberActive

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasNumberActive() bool`

HasNumberActive returns a boolean if a field has been set.

### GetNumberInService

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetNumberInService() int64`

GetNumberInService returns the NumberInService field if non-nil, zero value otherwise.

### GetNumberInServiceOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetNumberInServiceOk() (*int64, bool)`

GetNumberInServiceOk returns a tuple with the NumberInService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberInService

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetNumberInService(v int64)`

SetNumberInService sets NumberInService field to given value.

### HasNumberInService

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasNumberInService() bool`

HasNumberInService returns a boolean if a field has been set.

### GetHealthScore

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetHealthScore() int64`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetHealthScoreOk() (*int64, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetHealthScore(v int64)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetPerformanceScore

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetPerformanceScore() int64`

GetPerformanceScore returns the PerformanceScore field if non-nil, zero value otherwise.

### GetPerformanceScoreOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetPerformanceScoreOk() (*int64, bool)`

GetPerformanceScoreOk returns a tuple with the PerformanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceScore

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetPerformanceScore(v int64)`

SetPerformanceScore sets PerformanceScore field to given value.

### HasPerformanceScore

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasPerformanceScore() bool`

HasPerformanceScore returns a boolean if a field has been set.

### GetHealthPenalty

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetHealthPenalty() int64`

GetHealthPenalty returns the HealthPenalty field if non-nil, zero value otherwise.

### GetHealthPenaltyOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetHealthPenaltyOk() (*int64, bool)`

GetHealthPenaltyOk returns a tuple with the HealthPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthPenalty

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetHealthPenalty(v int64)`

SetHealthPenalty sets HealthPenalty field to given value.

### HasHealthPenalty

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasHealthPenalty() bool`

HasHealthPenalty returns a boolean if a field has been set.

### GetSecurityPenalty

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetSecurityPenalty() int64`

GetSecurityPenalty returns the SecurityPenalty field if non-nil, zero value otherwise.

### GetSecurityPenaltyOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetSecurityPenaltyOk() (*int64, bool)`

GetSecurityPenaltyOk returns a tuple with the SecurityPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPenalty

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetSecurityPenalty(v int64)`

SetSecurityPenalty sets SecurityPenalty field to given value.

### HasSecurityPenalty

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasSecurityPenalty() bool`

HasSecurityPenalty returns a boolean if a field has been set.

### GetErrorPenalty

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetErrorPenalty() int64`

GetErrorPenalty returns the ErrorPenalty field if non-nil, zero value otherwise.

### GetErrorPenaltyOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetErrorPenaltyOk() (*int64, bool)`

GetErrorPenaltyOk returns a tuple with the ErrorPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPenalty

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetErrorPenalty(v int64)`

SetErrorPenalty sets ErrorPenalty field to given value.

### HasErrorPenalty

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasErrorPenalty() bool`

HasErrorPenalty returns a boolean if a field has been set.

### GetDownAction

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetDownAction() string`

GetDownAction returns the DownAction field if non-nil, zero value otherwise.

### GetDownActionOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetDownActionOk() (*string, bool)`

GetDownActionOk returns a tuple with the DownAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownAction

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetDownAction(v string)`

SetDownAction sets DownAction field to given value.

### HasDownAction

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasDownAction() bool`

HasDownAction returns a boolean if a field has been set.

### SetDownActionNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetDownActionNil(b bool)`

 SetDownActionNil sets the value for DownAction to be an explicit nil

### UnsetDownAction
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetDownAction()`

UnsetDownAction ensures that no value is present for DownAction, not even an explicit nil
### GetRampTime

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetRampTime() string`

GetRampTime returns the RampTime field if non-nil, zero value otherwise.

### GetRampTimeOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetRampTimeOk() (*string, bool)`

GetRampTimeOk returns a tuple with the RampTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRampTime

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetRampTime(v string)`

SetRampTime sets RampTime field to given value.

### HasRampTime

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasRampTime() bool`

HasRampTime returns a boolean if a field has been set.

### SetRampTimeNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetRampTimeNil(b bool)`

 SetRampTimeNil sets the value for RampTime to be an explicit nil

### UnsetRampTime
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetRampTime()`

UnsetRampTime ensures that no value is present for RampTime, not even an explicit nil
### GetPort

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPortType

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### SetPortTypeNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetPortTypeNil(b bool)`

 SetPortTypeNil sets the value for PortType to be an explicit nil

### UnsetPortType
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetPortType()`

UnsetPortType ensures that no value is present for PortType, not even an explicit nil
### GetStatus

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNodes

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetNodes() []GetLoadBalancerPool200ResponseLoadBalancerPoolNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetNodesOk() (*[]GetLoadBalancerPool200ResponseLoadBalancerPoolNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetNodes(v []GetLoadBalancerPool200ResponseLoadBalancerPoolNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetMonitors

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMonitors() []GetLoadBalancerPool200ResponseLoadBalancerPoolMonitorsInner`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMonitorsOk() (*[]GetLoadBalancerPool200ResponseLoadBalancerPoolMonitorsInner, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMonitors(v []GetLoadBalancerPool200ResponseLoadBalancerPoolMonitorsInner)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetMembers

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMembers() []map[string]interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetMembersOk() (*[]map[string]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetMembers(v []map[string]interface{})`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetConfig

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetLoadBalancerPool200ResponseLoadBalancerPool) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


