# UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolLoadBalancer**](UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolLoadBalancer.md) |  | [optional] 
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
**Nodes** | Pointer to [**[]UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolNodesInner**](UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolNodesInner.md) |  | [optional] 
**Monitors** | Pointer to [**[]UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolMonitorsInner**](UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolMonitorsInner.md) |  | [optional] 
**Members** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool

`func NewUpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool() *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool`

NewUpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool instantiates a new UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolWithDefaults

`func NewUpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolWithDefaults() *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool`

NewUpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolWithDefaults instantiates a new UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetLoadBalancer() UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetLoadBalancerOk() (*UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetLoadBalancer(v UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVisibility

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVipSticky

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetVipSticky() string`

GetVipSticky returns the VipSticky field if non-nil, zero value otherwise.

### GetVipStickyOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetVipStickyOk() (*string, bool)`

GetVipStickyOk returns a tuple with the VipSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSticky

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetVipSticky(v string)`

SetVipSticky sets VipSticky field to given value.

### HasVipSticky

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasVipSticky() bool`

HasVipSticky returns a boolean if a field has been set.

### SetVipStickyNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetVipStickyNil(b bool)`

 SetVipStickyNil sets the value for VipSticky to be an explicit nil

### UnsetVipSticky
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetVipSticky()`

UnsetVipSticky ensures that no value is present for VipSticky, not even an explicit nil
### GetVipBalance

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetVipBalance() string`

GetVipBalance returns the VipBalance field if non-nil, zero value otherwise.

### GetVipBalanceOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetVipBalanceOk() (*string, bool)`

GetVipBalanceOk returns a tuple with the VipBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipBalance

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetVipBalance(v string)`

SetVipBalance sets VipBalance field to given value.

### HasVipBalance

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasVipBalance() bool`

HasVipBalance returns a boolean if a field has been set.

### GetAllowNat

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetAllowNat() string`

GetAllowNat returns the AllowNat field if non-nil, zero value otherwise.

### GetAllowNatOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetAllowNatOk() (*string, bool)`

GetAllowNatOk returns a tuple with the AllowNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNat

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetAllowNat(v string)`

SetAllowNat sets AllowNat field to given value.

### HasAllowNat

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasAllowNat() bool`

HasAllowNat returns a boolean if a field has been set.

### SetAllowNatNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetAllowNatNil(b bool)`

 SetAllowNatNil sets the value for AllowNat to be an explicit nil

### UnsetAllowNat
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetAllowNat()`

UnsetAllowNat ensures that no value is present for AllowNat, not even an explicit nil
### GetAllowSnat

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetAllowSnat() string`

GetAllowSnat returns the AllowSnat field if non-nil, zero value otherwise.

### GetAllowSnatOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetAllowSnatOk() (*string, bool)`

GetAllowSnatOk returns a tuple with the AllowSnat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSnat

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetAllowSnat(v string)`

SetAllowSnat sets AllowSnat field to given value.

### HasAllowSnat

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasAllowSnat() bool`

HasAllowSnat returns a boolean if a field has been set.

### SetAllowSnatNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetAllowSnatNil(b bool)`

 SetAllowSnatNil sets the value for AllowSnat to be an explicit nil

### UnsetAllowSnat
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetAllowSnat()`

UnsetAllowSnat ensures that no value is present for AllowSnat, not even an explicit nil
### GetVipClientIpMode

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetVipClientIpMode() string`

GetVipClientIpMode returns the VipClientIpMode field if non-nil, zero value otherwise.

### GetVipClientIpModeOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetVipClientIpModeOk() (*string, bool)`

GetVipClientIpModeOk returns a tuple with the VipClientIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipClientIpMode

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetVipClientIpMode(v string)`

SetVipClientIpMode sets VipClientIpMode field to given value.

### HasVipClientIpMode

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasVipClientIpMode() bool`

HasVipClientIpMode returns a boolean if a field has been set.

### SetVipClientIpModeNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetVipClientIpModeNil(b bool)`

 SetVipClientIpModeNil sets the value for VipClientIpMode to be an explicit nil

### UnsetVipClientIpMode
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetVipClientIpMode()`

UnsetVipClientIpMode ensures that no value is present for VipClientIpMode, not even an explicit nil
### GetVipServerIpMode

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetVipServerIpMode() string`

GetVipServerIpMode returns the VipServerIpMode field if non-nil, zero value otherwise.

### GetVipServerIpModeOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetVipServerIpModeOk() (*string, bool)`

GetVipServerIpModeOk returns a tuple with the VipServerIpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipServerIpMode

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetVipServerIpMode(v string)`

SetVipServerIpMode sets VipServerIpMode field to given value.

### HasVipServerIpMode

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasVipServerIpMode() bool`

HasVipServerIpMode returns a boolean if a field has been set.

### SetVipServerIpModeNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetVipServerIpModeNil(b bool)`

 SetVipServerIpModeNil sets the value for VipServerIpMode to be an explicit nil

### UnsetVipServerIpMode
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetVipServerIpMode()`

UnsetVipServerIpMode ensures that no value is present for VipServerIpMode, not even an explicit nil
### GetMinActive

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMinActive() int64`

GetMinActive returns the MinActive field if non-nil, zero value otherwise.

### GetMinActiveOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMinActiveOk() (*int64, bool)`

GetMinActiveOk returns a tuple with the MinActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinActive

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMinActive(v int64)`

SetMinActive sets MinActive field to given value.

### HasMinActive

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasMinActive() bool`

HasMinActive returns a boolean if a field has been set.

### GetMinInService

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMinInService() string`

GetMinInService returns the MinInService field if non-nil, zero value otherwise.

### GetMinInServiceOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMinInServiceOk() (*string, bool)`

GetMinInServiceOk returns a tuple with the MinInService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInService

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMinInService(v string)`

SetMinInService sets MinInService field to given value.

### HasMinInService

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasMinInService() bool`

HasMinInService returns a boolean if a field has been set.

### SetMinInServiceNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMinInServiceNil(b bool)`

 SetMinInServiceNil sets the value for MinInService to be an explicit nil

### UnsetMinInService
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetMinInService()`

UnsetMinInService ensures that no value is present for MinInService, not even an explicit nil
### GetMinUpMonitor

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMinUpMonitor() string`

GetMinUpMonitor returns the MinUpMonitor field if non-nil, zero value otherwise.

### GetMinUpMonitorOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMinUpMonitorOk() (*string, bool)`

GetMinUpMonitorOk returns a tuple with the MinUpMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpMonitor

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMinUpMonitor(v string)`

SetMinUpMonitor sets MinUpMonitor field to given value.

### HasMinUpMonitor

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasMinUpMonitor() bool`

HasMinUpMonitor returns a boolean if a field has been set.

### SetMinUpMonitorNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMinUpMonitorNil(b bool)`

 SetMinUpMonitorNil sets the value for MinUpMonitor to be an explicit nil

### UnsetMinUpMonitor
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetMinUpMonitor()`

UnsetMinUpMonitor ensures that no value is present for MinUpMonitor, not even an explicit nil
### GetMinUpAction

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMinUpAction() string`

GetMinUpAction returns the MinUpAction field if non-nil, zero value otherwise.

### GetMinUpActionOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMinUpActionOk() (*string, bool)`

GetMinUpActionOk returns a tuple with the MinUpAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpAction

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMinUpAction(v string)`

SetMinUpAction sets MinUpAction field to given value.

### HasMinUpAction

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasMinUpAction() bool`

HasMinUpAction returns a boolean if a field has been set.

### SetMinUpActionNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMinUpActionNil(b bool)`

 SetMinUpActionNil sets the value for MinUpAction to be an explicit nil

### UnsetMinUpAction
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetMinUpAction()`

UnsetMinUpAction ensures that no value is present for MinUpAction, not even an explicit nil
### GetMaxQueueDepth

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMaxQueueDepth() string`

GetMaxQueueDepth returns the MaxQueueDepth field if non-nil, zero value otherwise.

### GetMaxQueueDepthOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMaxQueueDepthOk() (*string, bool)`

GetMaxQueueDepthOk returns a tuple with the MaxQueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueueDepth

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMaxQueueDepth(v string)`

SetMaxQueueDepth sets MaxQueueDepth field to given value.

### HasMaxQueueDepth

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasMaxQueueDepth() bool`

HasMaxQueueDepth returns a boolean if a field has been set.

### SetMaxQueueDepthNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMaxQueueDepthNil(b bool)`

 SetMaxQueueDepthNil sets the value for MaxQueueDepth to be an explicit nil

### UnsetMaxQueueDepth
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetMaxQueueDepth()`

UnsetMaxQueueDepth ensures that no value is present for MaxQueueDepth, not even an explicit nil
### GetMaxQueueTime

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMaxQueueTime() string`

GetMaxQueueTime returns the MaxQueueTime field if non-nil, zero value otherwise.

### GetMaxQueueTimeOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMaxQueueTimeOk() (*string, bool)`

GetMaxQueueTimeOk returns a tuple with the MaxQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueueTime

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMaxQueueTime(v string)`

SetMaxQueueTime sets MaxQueueTime field to given value.

### HasMaxQueueTime

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasMaxQueueTime() bool`

HasMaxQueueTime returns a boolean if a field has been set.

### SetMaxQueueTimeNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMaxQueueTimeNil(b bool)`

 SetMaxQueueTimeNil sets the value for MaxQueueTime to be an explicit nil

### UnsetMaxQueueTime
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetMaxQueueTime()`

UnsetMaxQueueTime ensures that no value is present for MaxQueueTime, not even an explicit nil
### GetNumberActive

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetNumberActive() int64`

GetNumberActive returns the NumberActive field if non-nil, zero value otherwise.

### GetNumberActiveOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetNumberActiveOk() (*int64, bool)`

GetNumberActiveOk returns a tuple with the NumberActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberActive

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetNumberActive(v int64)`

SetNumberActive sets NumberActive field to given value.

### HasNumberActive

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasNumberActive() bool`

HasNumberActive returns a boolean if a field has been set.

### GetNumberInService

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetNumberInService() int64`

GetNumberInService returns the NumberInService field if non-nil, zero value otherwise.

### GetNumberInServiceOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetNumberInServiceOk() (*int64, bool)`

GetNumberInServiceOk returns a tuple with the NumberInService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberInService

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetNumberInService(v int64)`

SetNumberInService sets NumberInService field to given value.

### HasNumberInService

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasNumberInService() bool`

HasNumberInService returns a boolean if a field has been set.

### GetHealthScore

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetHealthScore() int64`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetHealthScoreOk() (*int64, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetHealthScore(v int64)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetPerformanceScore

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetPerformanceScore() int64`

GetPerformanceScore returns the PerformanceScore field if non-nil, zero value otherwise.

### GetPerformanceScoreOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetPerformanceScoreOk() (*int64, bool)`

GetPerformanceScoreOk returns a tuple with the PerformanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceScore

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetPerformanceScore(v int64)`

SetPerformanceScore sets PerformanceScore field to given value.

### HasPerformanceScore

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasPerformanceScore() bool`

HasPerformanceScore returns a boolean if a field has been set.

### GetHealthPenalty

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetHealthPenalty() int64`

GetHealthPenalty returns the HealthPenalty field if non-nil, zero value otherwise.

### GetHealthPenaltyOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetHealthPenaltyOk() (*int64, bool)`

GetHealthPenaltyOk returns a tuple with the HealthPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthPenalty

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetHealthPenalty(v int64)`

SetHealthPenalty sets HealthPenalty field to given value.

### HasHealthPenalty

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasHealthPenalty() bool`

HasHealthPenalty returns a boolean if a field has been set.

### GetSecurityPenalty

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetSecurityPenalty() int64`

GetSecurityPenalty returns the SecurityPenalty field if non-nil, zero value otherwise.

### GetSecurityPenaltyOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetSecurityPenaltyOk() (*int64, bool)`

GetSecurityPenaltyOk returns a tuple with the SecurityPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPenalty

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetSecurityPenalty(v int64)`

SetSecurityPenalty sets SecurityPenalty field to given value.

### HasSecurityPenalty

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasSecurityPenalty() bool`

HasSecurityPenalty returns a boolean if a field has been set.

### GetErrorPenalty

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetErrorPenalty() int64`

GetErrorPenalty returns the ErrorPenalty field if non-nil, zero value otherwise.

### GetErrorPenaltyOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetErrorPenaltyOk() (*int64, bool)`

GetErrorPenaltyOk returns a tuple with the ErrorPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPenalty

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetErrorPenalty(v int64)`

SetErrorPenalty sets ErrorPenalty field to given value.

### HasErrorPenalty

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasErrorPenalty() bool`

HasErrorPenalty returns a boolean if a field has been set.

### GetDownAction

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetDownAction() string`

GetDownAction returns the DownAction field if non-nil, zero value otherwise.

### GetDownActionOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetDownActionOk() (*string, bool)`

GetDownActionOk returns a tuple with the DownAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownAction

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetDownAction(v string)`

SetDownAction sets DownAction field to given value.

### HasDownAction

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasDownAction() bool`

HasDownAction returns a boolean if a field has been set.

### SetDownActionNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetDownActionNil(b bool)`

 SetDownActionNil sets the value for DownAction to be an explicit nil

### UnsetDownAction
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetDownAction()`

UnsetDownAction ensures that no value is present for DownAction, not even an explicit nil
### GetRampTime

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetRampTime() string`

GetRampTime returns the RampTime field if non-nil, zero value otherwise.

### GetRampTimeOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetRampTimeOk() (*string, bool)`

GetRampTimeOk returns a tuple with the RampTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRampTime

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetRampTime(v string)`

SetRampTime sets RampTime field to given value.

### HasRampTime

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasRampTime() bool`

HasRampTime returns a boolean if a field has been set.

### SetRampTimeNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetRampTimeNil(b bool)`

 SetRampTimeNil sets the value for RampTime to be an explicit nil

### UnsetRampTime
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetRampTime()`

UnsetRampTime ensures that no value is present for RampTime, not even an explicit nil
### GetPort

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPortType

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### SetPortTypeNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetPortTypeNil(b bool)`

 SetPortTypeNil sets the value for PortType to be an explicit nil

### UnsetPortType
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetPortType()`

UnsetPortType ensures that no value is present for PortType, not even an explicit nil
### GetStatus

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNodes

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetNodes() []UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolNodesInner`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetNodesOk() (*[]UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolNodesInner, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetNodes(v []UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolNodesInner)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetMonitors

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMonitors() []UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolMonitorsInner`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMonitorsOk() (*[]UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolMonitorsInner, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMonitors(v []UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPoolMonitorsInner)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetMembers

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMembers() []map[string]interface{}`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetMembersOk() (*[]map[string]interface{}, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetMembers(v []map[string]interface{})`

SetMembers sets Members field to given value.

### HasMembers

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateLoadBalancerPool200ResponseAllOfLoadBalancerPool) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


