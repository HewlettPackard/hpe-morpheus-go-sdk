# BillingZonesZonesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneName** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**ZoneUUID** | Pointer to **string** |  | [optional] 
**ZoneCode** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**ComputeServers** | Pointer to [**BillingZonesZonesInnerComputeServers**](BillingZonesZonesInnerComputeServers.md) |  | [optional] 
**Instances** | Pointer to [**BillingZonesZonesInnerInstances**](BillingZonesZonesInnerInstances.md) |  | [optional] 
**DiscoveredServers** | Pointer to [**BillingZonesZonesInnerDiscoveredServers**](BillingZonesZonesInnerDiscoveredServers.md) |  | [optional] 
**LoadBalancers** | Pointer to [**BillingZonesZonesInnerLoadBalancers**](BillingZonesZonesInnerLoadBalancers.md) |  | [optional] 
**VirtualImages** | Pointer to [**BillingZonesZonesInnerVirtualImages**](BillingZonesZonesInnerVirtualImages.md) |  | [optional] 
**Snapshots** | Pointer to [**BillingZonesZonesInnerSnapshots**](BillingZonesZonesInnerSnapshots.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 

## Methods

### NewBillingZonesZonesInner

`func NewBillingZonesZonesInner() *BillingZonesZonesInner`

NewBillingZonesZonesInner instantiates a new BillingZonesZonesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingZonesZonesInnerWithDefaults

`func NewBillingZonesZonesInnerWithDefaults() *BillingZonesZonesInner`

NewBillingZonesZonesInnerWithDefaults instantiates a new BillingZonesZonesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *BillingZonesZonesInner) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *BillingZonesZonesInner) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *BillingZonesZonesInner) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *BillingZonesZonesInner) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneId

`func (o *BillingZonesZonesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *BillingZonesZonesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *BillingZonesZonesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *BillingZonesZonesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneUUID

`func (o *BillingZonesZonesInner) GetZoneUUID() string`

GetZoneUUID returns the ZoneUUID field if non-nil, zero value otherwise.

### GetZoneUUIDOk

`func (o *BillingZonesZonesInner) GetZoneUUIDOk() (*string, bool)`

GetZoneUUIDOk returns a tuple with the ZoneUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUUID

`func (o *BillingZonesZonesInner) SetZoneUUID(v string)`

SetZoneUUID sets ZoneUUID field to given value.

### HasZoneUUID

`func (o *BillingZonesZonesInner) HasZoneUUID() bool`

HasZoneUUID returns a boolean if a field has been set.

### GetZoneCode

`func (o *BillingZonesZonesInner) GetZoneCode() string`

GetZoneCode returns the ZoneCode field if non-nil, zero value otherwise.

### GetZoneCodeOk

`func (o *BillingZonesZonesInner) GetZoneCodeOk() (*string, bool)`

GetZoneCodeOk returns a tuple with the ZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCode

`func (o *BillingZonesZonesInner) SetZoneCode(v string)`

SetZoneCode sets ZoneCode field to given value.

### HasZoneCode

`func (o *BillingZonesZonesInner) HasZoneCode() bool`

HasZoneCode returns a boolean if a field has been set.

### SetZoneCodeNil

`func (o *BillingZonesZonesInner) SetZoneCodeNil(b bool)`

 SetZoneCodeNil sets the value for ZoneCode to be an explicit nil

### UnsetZoneCode
`func (o *BillingZonesZonesInner) UnsetZoneCode()`

UnsetZoneCode ensures that no value is present for ZoneCode, not even an explicit nil
### GetStartDate

`func (o *BillingZonesZonesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingZonesZonesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingZonesZonesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingZonesZonesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingZonesZonesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingZonesZonesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingZonesZonesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingZonesZonesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriceUnit

`func (o *BillingZonesZonesInner) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *BillingZonesZonesInner) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *BillingZonesZonesInner) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *BillingZonesZonesInner) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetComputeServers

`func (o *BillingZonesZonesInner) GetComputeServers() BillingZonesZonesInnerComputeServers`

GetComputeServers returns the ComputeServers field if non-nil, zero value otherwise.

### GetComputeServersOk

`func (o *BillingZonesZonesInner) GetComputeServersOk() (*BillingZonesZonesInnerComputeServers, bool)`

GetComputeServersOk returns a tuple with the ComputeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServers

`func (o *BillingZonesZonesInner) SetComputeServers(v BillingZonesZonesInnerComputeServers)`

SetComputeServers sets ComputeServers field to given value.

### HasComputeServers

`func (o *BillingZonesZonesInner) HasComputeServers() bool`

HasComputeServers returns a boolean if a field has been set.

### GetInstances

`func (o *BillingZonesZonesInner) GetInstances() BillingZonesZonesInnerInstances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *BillingZonesZonesInner) GetInstancesOk() (*BillingZonesZonesInnerInstances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *BillingZonesZonesInner) SetInstances(v BillingZonesZonesInnerInstances)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *BillingZonesZonesInner) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetDiscoveredServers

`func (o *BillingZonesZonesInner) GetDiscoveredServers() BillingZonesZonesInnerDiscoveredServers`

GetDiscoveredServers returns the DiscoveredServers field if non-nil, zero value otherwise.

### GetDiscoveredServersOk

`func (o *BillingZonesZonesInner) GetDiscoveredServersOk() (*BillingZonesZonesInnerDiscoveredServers, bool)`

GetDiscoveredServersOk returns a tuple with the DiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredServers

`func (o *BillingZonesZonesInner) SetDiscoveredServers(v BillingZonesZonesInnerDiscoveredServers)`

SetDiscoveredServers sets DiscoveredServers field to given value.

### HasDiscoveredServers

`func (o *BillingZonesZonesInner) HasDiscoveredServers() bool`

HasDiscoveredServers returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *BillingZonesZonesInner) GetLoadBalancers() BillingZonesZonesInnerLoadBalancers`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BillingZonesZonesInner) GetLoadBalancersOk() (*BillingZonesZonesInnerLoadBalancers, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BillingZonesZonesInner) SetLoadBalancers(v BillingZonesZonesInnerLoadBalancers)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *BillingZonesZonesInner) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetVirtualImages

`func (o *BillingZonesZonesInner) GetVirtualImages() BillingZonesZonesInnerVirtualImages`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *BillingZonesZonesInner) GetVirtualImagesOk() (*BillingZonesZonesInnerVirtualImages, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *BillingZonesZonesInner) SetVirtualImages(v BillingZonesZonesInnerVirtualImages)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *BillingZonesZonesInner) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetSnapshots

`func (o *BillingZonesZonesInner) GetSnapshots() BillingZonesZonesInnerSnapshots`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *BillingZonesZonesInner) GetSnapshotsOk() (*BillingZonesZonesInnerSnapshots, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *BillingZonesZonesInner) SetSnapshots(v BillingZonesZonesInnerSnapshots)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *BillingZonesZonesInner) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetPrice

`func (o *BillingZonesZonesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingZonesZonesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingZonesZonesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingZonesZonesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingZonesZonesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingZonesZonesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingZonesZonesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingZonesZonesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


