# ListBillingZone200ResponseAllOfBillingInfoZonesInner

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
**ComputeServers** | Pointer to [**ListBillingZone200ResponseAllOfBillingInfoZonesInnerComputeServers**](ListBillingZone200ResponseAllOfBillingInfoZonesInnerComputeServers.md) |  | [optional] 
**Instances** | Pointer to [**ListBillingZone200ResponseAllOfBillingInfoZonesInnerInstances**](ListBillingZone200ResponseAllOfBillingInfoZonesInnerInstances.md) |  | [optional] 
**DiscoveredServers** | Pointer to [**ListBillingZone200ResponseAllOfBillingInfoZonesInnerDiscoveredServers**](ListBillingZone200ResponseAllOfBillingInfoZonesInnerDiscoveredServers.md) |  | [optional] 
**LoadBalancers** | Pointer to [**ListBillingZone200ResponseAllOfBillingInfoZonesInnerLoadBalancers**](ListBillingZone200ResponseAllOfBillingInfoZonesInnerLoadBalancers.md) |  | [optional] 
**VirtualImages** | Pointer to [**ListBillingZone200ResponseAllOfBillingInfoZonesInnerVirtualImages**](ListBillingZone200ResponseAllOfBillingInfoZonesInnerVirtualImages.md) |  | [optional] 
**Snapshots** | Pointer to [**ListBillingZone200ResponseAllOfBillingInfoZonesInnerSnapshots**](ListBillingZone200ResponseAllOfBillingInfoZonesInnerSnapshots.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 

## Methods

### NewListBillingZone200ResponseAllOfBillingInfoZonesInner

`func NewListBillingZone200ResponseAllOfBillingInfoZonesInner() *ListBillingZone200ResponseAllOfBillingInfoZonesInner`

NewListBillingZone200ResponseAllOfBillingInfoZonesInner instantiates a new ListBillingZone200ResponseAllOfBillingInfoZonesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingZone200ResponseAllOfBillingInfoZonesInnerWithDefaults

`func NewListBillingZone200ResponseAllOfBillingInfoZonesInnerWithDefaults() *ListBillingZone200ResponseAllOfBillingInfoZonesInner`

NewListBillingZone200ResponseAllOfBillingInfoZonesInnerWithDefaults instantiates a new ListBillingZone200ResponseAllOfBillingInfoZonesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneId

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneUUID

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetZoneUUID() string`

GetZoneUUID returns the ZoneUUID field if non-nil, zero value otherwise.

### GetZoneUUIDOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetZoneUUIDOk() (*string, bool)`

GetZoneUUIDOk returns a tuple with the ZoneUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUUID

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetZoneUUID(v string)`

SetZoneUUID sets ZoneUUID field to given value.

### HasZoneUUID

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasZoneUUID() bool`

HasZoneUUID returns a boolean if a field has been set.

### GetZoneCode

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetZoneCode() string`

GetZoneCode returns the ZoneCode field if non-nil, zero value otherwise.

### GetZoneCodeOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetZoneCodeOk() (*string, bool)`

GetZoneCodeOk returns a tuple with the ZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCode

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetZoneCode(v string)`

SetZoneCode sets ZoneCode field to given value.

### HasZoneCode

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasZoneCode() bool`

HasZoneCode returns a boolean if a field has been set.

### SetZoneCodeNil

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetZoneCodeNil(b bool)`

 SetZoneCodeNil sets the value for ZoneCode to be an explicit nil

### UnsetZoneCode
`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) UnsetZoneCode()`

UnsetZoneCode ensures that no value is present for ZoneCode, not even an explicit nil
### GetStartDate

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriceUnit

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetComputeServers

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetComputeServers() ListBillingZone200ResponseAllOfBillingInfoZonesInnerComputeServers`

GetComputeServers returns the ComputeServers field if non-nil, zero value otherwise.

### GetComputeServersOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetComputeServersOk() (*ListBillingZone200ResponseAllOfBillingInfoZonesInnerComputeServers, bool)`

GetComputeServersOk returns a tuple with the ComputeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServers

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetComputeServers(v ListBillingZone200ResponseAllOfBillingInfoZonesInnerComputeServers)`

SetComputeServers sets ComputeServers field to given value.

### HasComputeServers

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasComputeServers() bool`

HasComputeServers returns a boolean if a field has been set.

### GetInstances

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetInstances() ListBillingZone200ResponseAllOfBillingInfoZonesInnerInstances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetInstancesOk() (*ListBillingZone200ResponseAllOfBillingInfoZonesInnerInstances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetInstances(v ListBillingZone200ResponseAllOfBillingInfoZonesInnerInstances)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetDiscoveredServers

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetDiscoveredServers() ListBillingZone200ResponseAllOfBillingInfoZonesInnerDiscoveredServers`

GetDiscoveredServers returns the DiscoveredServers field if non-nil, zero value otherwise.

### GetDiscoveredServersOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetDiscoveredServersOk() (*ListBillingZone200ResponseAllOfBillingInfoZonesInnerDiscoveredServers, bool)`

GetDiscoveredServersOk returns a tuple with the DiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredServers

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetDiscoveredServers(v ListBillingZone200ResponseAllOfBillingInfoZonesInnerDiscoveredServers)`

SetDiscoveredServers sets DiscoveredServers field to given value.

### HasDiscoveredServers

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasDiscoveredServers() bool`

HasDiscoveredServers returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetLoadBalancers() ListBillingZone200ResponseAllOfBillingInfoZonesInnerLoadBalancers`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetLoadBalancersOk() (*ListBillingZone200ResponseAllOfBillingInfoZonesInnerLoadBalancers, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetLoadBalancers(v ListBillingZone200ResponseAllOfBillingInfoZonesInnerLoadBalancers)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetVirtualImages

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetVirtualImages() ListBillingZone200ResponseAllOfBillingInfoZonesInnerVirtualImages`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetVirtualImagesOk() (*ListBillingZone200ResponseAllOfBillingInfoZonesInnerVirtualImages, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetVirtualImages(v ListBillingZone200ResponseAllOfBillingInfoZonesInnerVirtualImages)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetSnapshots

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetSnapshots() ListBillingZone200ResponseAllOfBillingInfoZonesInnerSnapshots`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetSnapshotsOk() (*ListBillingZone200ResponseAllOfBillingInfoZonesInnerSnapshots, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetSnapshots(v ListBillingZone200ResponseAllOfBillingInfoZonesInnerSnapshots)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetPrice

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListBillingZone200ResponseAllOfBillingInfoZonesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


