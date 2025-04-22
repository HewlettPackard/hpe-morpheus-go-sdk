# ListBillingAccount200ResponseAllOfBillingInfoZonesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneName** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**ZoneUUID** | Pointer to **string** |  | [optional] 
**ZoneCode** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**ComputeServers** | Pointer to [**ListBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers**](ListBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers.md) |  | [optional] 
**Instances** | Pointer to [**ListBillingAccount200ResponseAllOfBillingInfoZonesInnerInstances**](ListBillingAccount200ResponseAllOfBillingInfoZonesInnerInstances.md) |  | [optional] 
**DiscoveredServers** | Pointer to [**ListBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers**](ListBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers.md) |  | [optional] 
**LoadBalancers** | Pointer to [**ListBillingAccount200ResponseAllOfBillingInfoZonesInnerLoadBalancers**](ListBillingAccount200ResponseAllOfBillingInfoZonesInnerLoadBalancers.md) |  | [optional] 
**VirtualImages** | Pointer to [**ListBillingAccount200ResponseAllOfBillingInfoZonesInnerVirtualImages**](ListBillingAccount200ResponseAllOfBillingInfoZonesInnerVirtualImages.md) |  | [optional] 
**Snapshots** | Pointer to [**ListBillingAccount200ResponseAllOfBillingInfoZonesInnerSnapshots**](ListBillingAccount200ResponseAllOfBillingInfoZonesInnerSnapshots.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 

## Methods

### NewListBillingAccount200ResponseAllOfBillingInfoZonesInner

`func NewListBillingAccount200ResponseAllOfBillingInfoZonesInner() *ListBillingAccount200ResponseAllOfBillingInfoZonesInner`

NewListBillingAccount200ResponseAllOfBillingInfoZonesInner instantiates a new ListBillingAccount200ResponseAllOfBillingInfoZonesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingAccount200ResponseAllOfBillingInfoZonesInnerWithDefaults

`func NewListBillingAccount200ResponseAllOfBillingInfoZonesInnerWithDefaults() *ListBillingAccount200ResponseAllOfBillingInfoZonesInner`

NewListBillingAccount200ResponseAllOfBillingInfoZonesInnerWithDefaults instantiates a new ListBillingAccount200ResponseAllOfBillingInfoZonesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneId

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneUUID

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneUUID() string`

GetZoneUUID returns the ZoneUUID field if non-nil, zero value otherwise.

### GetZoneUUIDOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneUUIDOk() (*string, bool)`

GetZoneUUIDOk returns a tuple with the ZoneUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUUID

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetZoneUUID(v string)`

SetZoneUUID sets ZoneUUID field to given value.

### HasZoneUUID

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasZoneUUID() bool`

HasZoneUUID returns a boolean if a field has been set.

### GetZoneCode

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneCode() string`

GetZoneCode returns the ZoneCode field if non-nil, zero value otherwise.

### GetZoneCodeOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneCodeOk() (*string, bool)`

GetZoneCodeOk returns a tuple with the ZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCode

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetZoneCode(v string)`

SetZoneCode sets ZoneCode field to given value.

### HasZoneCode

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasZoneCode() bool`

HasZoneCode returns a boolean if a field has been set.

### GetStartDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriceUnit

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetComputeServers

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetComputeServers() ListBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers`

GetComputeServers returns the ComputeServers field if non-nil, zero value otherwise.

### GetComputeServersOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetComputeServersOk() (*ListBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers, bool)`

GetComputeServersOk returns a tuple with the ComputeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServers

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetComputeServers(v ListBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers)`

SetComputeServers sets ComputeServers field to given value.

### HasComputeServers

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasComputeServers() bool`

HasComputeServers returns a boolean if a field has been set.

### GetInstances

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetInstances() ListBillingAccount200ResponseAllOfBillingInfoZonesInnerInstances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetInstancesOk() (*ListBillingAccount200ResponseAllOfBillingInfoZonesInnerInstances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetInstances(v ListBillingAccount200ResponseAllOfBillingInfoZonesInnerInstances)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetDiscoveredServers

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetDiscoveredServers() ListBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers`

GetDiscoveredServers returns the DiscoveredServers field if non-nil, zero value otherwise.

### GetDiscoveredServersOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetDiscoveredServersOk() (*ListBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers, bool)`

GetDiscoveredServersOk returns a tuple with the DiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredServers

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetDiscoveredServers(v ListBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers)`

SetDiscoveredServers sets DiscoveredServers field to given value.

### HasDiscoveredServers

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasDiscoveredServers() bool`

HasDiscoveredServers returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetLoadBalancers() ListBillingAccount200ResponseAllOfBillingInfoZonesInnerLoadBalancers`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetLoadBalancersOk() (*ListBillingAccount200ResponseAllOfBillingInfoZonesInnerLoadBalancers, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetLoadBalancers(v ListBillingAccount200ResponseAllOfBillingInfoZonesInnerLoadBalancers)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetVirtualImages

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetVirtualImages() ListBillingAccount200ResponseAllOfBillingInfoZonesInnerVirtualImages`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetVirtualImagesOk() (*ListBillingAccount200ResponseAllOfBillingInfoZonesInnerVirtualImages, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetVirtualImages(v ListBillingAccount200ResponseAllOfBillingInfoZonesInnerVirtualImages)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetSnapshots

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetSnapshots() ListBillingAccount200ResponseAllOfBillingInfoZonesInnerSnapshots`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetSnapshotsOk() (*ListBillingAccount200ResponseAllOfBillingInfoZonesInnerSnapshots, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetSnapshots(v ListBillingAccount200ResponseAllOfBillingInfoZonesInnerSnapshots)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetPrice

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingAccount200ResponseAllOfBillingInfoZonesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


