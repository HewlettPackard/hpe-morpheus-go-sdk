# BillingZonesInner

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
**ComputeServers** | Pointer to [**BillingZonesInnerComputeServers**](BillingZonesInnerComputeServers.md) |  | [optional] 
**Instances** | Pointer to [**BillingZonesInnerInstances**](BillingZonesInnerInstances.md) |  | [optional] 
**DiscoveredServers** | Pointer to [**BillingZonesInnerDiscoveredServers**](BillingZonesInnerDiscoveredServers.md) |  | [optional] 
**LoadBalancers** | Pointer to [**BillingZonesInnerLoadBalancers**](BillingZonesInnerLoadBalancers.md) |  | [optional] 
**VirtualImages** | Pointer to [**BillingZonesInnerVirtualImages**](BillingZonesInnerVirtualImages.md) |  | [optional] 
**Snapshots** | Pointer to [**BillingZonesInnerSnapshots**](BillingZonesInnerSnapshots.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 

## Methods

### NewBillingZonesInner

`func NewBillingZonesInner() *BillingZonesInner`

NewBillingZonesInner instantiates a new BillingZonesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetZoneName

`func (o *BillingZonesInner) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *BillingZonesInner) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *BillingZonesInner) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *BillingZonesInner) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneId

`func (o *BillingZonesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *BillingZonesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *BillingZonesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *BillingZonesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneUUID

`func (o *BillingZonesInner) GetZoneUUID() string`

GetZoneUUID returns the ZoneUUID field if non-nil, zero value otherwise.

### GetZoneUUIDOk

`func (o *BillingZonesInner) GetZoneUUIDOk() (*string, bool)`

GetZoneUUIDOk returns a tuple with the ZoneUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUUID

`func (o *BillingZonesInner) SetZoneUUID(v string)`

SetZoneUUID sets ZoneUUID field to given value.

### HasZoneUUID

`func (o *BillingZonesInner) HasZoneUUID() bool`

HasZoneUUID returns a boolean if a field has been set.

### GetZoneCode

`func (o *BillingZonesInner) GetZoneCode() string`

GetZoneCode returns the ZoneCode field if non-nil, zero value otherwise.

### GetZoneCodeOk

`func (o *BillingZonesInner) GetZoneCodeOk() (*string, bool)`

GetZoneCodeOk returns a tuple with the ZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCode

`func (o *BillingZonesInner) SetZoneCode(v string)`

SetZoneCode sets ZoneCode field to given value.

### HasZoneCode

`func (o *BillingZonesInner) HasZoneCode() bool`

HasZoneCode returns a boolean if a field has been set.

### SetZoneCodeNil

`func (o *BillingZonesInner) SetZoneCodeNil(b bool)`

 SetZoneCodeNil sets the value for ZoneCode to be an explicit nil

### UnsetZoneCode
`func (o *BillingZonesInner) UnsetZoneCode()`

UnsetZoneCode ensures that no value is present for ZoneCode, not even an explicit nil
### GetStartDate

`func (o *BillingZonesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingZonesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingZonesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingZonesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingZonesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingZonesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingZonesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingZonesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriceUnit

`func (o *BillingZonesInner) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *BillingZonesInner) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *BillingZonesInner) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *BillingZonesInner) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetComputeServers

`func (o *BillingZonesInner) GetComputeServers() BillingZonesInnerComputeServers`

GetComputeServers returns the ComputeServers field if non-nil, zero value otherwise.

### GetComputeServersOk

`func (o *BillingZonesInner) GetComputeServersOk() (*BillingZonesInnerComputeServers, bool)`

GetComputeServersOk returns a tuple with the ComputeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServers

`func (o *BillingZonesInner) SetComputeServers(v BillingZonesInnerComputeServers)`

SetComputeServers sets ComputeServers field to given value.

### HasComputeServers

`func (o *BillingZonesInner) HasComputeServers() bool`

HasComputeServers returns a boolean if a field has been set.

### GetInstances

`func (o *BillingZonesInner) GetInstances() BillingZonesInnerInstances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *BillingZonesInner) GetInstancesOk() (*BillingZonesInnerInstances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *BillingZonesInner) SetInstances(v BillingZonesInnerInstances)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *BillingZonesInner) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetDiscoveredServers

`func (o *BillingZonesInner) GetDiscoveredServers() BillingZonesInnerDiscoveredServers`

GetDiscoveredServers returns the DiscoveredServers field if non-nil, zero value otherwise.

### GetDiscoveredServersOk

`func (o *BillingZonesInner) GetDiscoveredServersOk() (*BillingZonesInnerDiscoveredServers, bool)`

GetDiscoveredServersOk returns a tuple with the DiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredServers

`func (o *BillingZonesInner) SetDiscoveredServers(v BillingZonesInnerDiscoveredServers)`

SetDiscoveredServers sets DiscoveredServers field to given value.

### HasDiscoveredServers

`func (o *BillingZonesInner) HasDiscoveredServers() bool`

HasDiscoveredServers returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *BillingZonesInner) GetLoadBalancers() BillingZonesInnerLoadBalancers`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BillingZonesInner) GetLoadBalancersOk() (*BillingZonesInnerLoadBalancers, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BillingZonesInner) SetLoadBalancers(v BillingZonesInnerLoadBalancers)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *BillingZonesInner) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetVirtualImages

`func (o *BillingZonesInner) GetVirtualImages() BillingZonesInnerVirtualImages`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *BillingZonesInner) GetVirtualImagesOk() (*BillingZonesInnerVirtualImages, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *BillingZonesInner) SetVirtualImages(v BillingZonesInnerVirtualImages)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *BillingZonesInner) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetSnapshots

`func (o *BillingZonesInner) GetSnapshots() BillingZonesInnerSnapshots`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *BillingZonesInner) GetSnapshotsOk() (*BillingZonesInnerSnapshots, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *BillingZonesInner) SetSnapshots(v BillingZonesInnerSnapshots)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *BillingZonesInner) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetPrice

`func (o *BillingZonesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingZonesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingZonesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingZonesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingZonesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingZonesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingZonesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingZonesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


