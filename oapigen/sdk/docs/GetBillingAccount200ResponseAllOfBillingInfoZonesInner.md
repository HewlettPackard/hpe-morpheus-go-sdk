# GetBillingAccount200ResponseAllOfBillingInfoZonesInner

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
**ComputeServers** | Pointer to [**GetBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers**](GetBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers.md) |  | [optional] 
**Instances** | Pointer to [**GetBillingAccount200ResponseAllOfBillingInfoZonesInnerInstances**](GetBillingAccount200ResponseAllOfBillingInfoZonesInnerInstances.md) |  | [optional] 
**DiscoveredServers** | Pointer to [**GetBillingAccount200ResponseAllOfBillingInfoZonesInnerDiscoveredServers**](GetBillingAccount200ResponseAllOfBillingInfoZonesInnerDiscoveredServers.md) |  | [optional] 
**LoadBalancers** | Pointer to [**GetBillingAccount200ResponseAllOfBillingInfoZonesInnerLoadBalancers**](GetBillingAccount200ResponseAllOfBillingInfoZonesInnerLoadBalancers.md) |  | [optional] 
**VirtualImages** | Pointer to [**GetBillingAccount200ResponseAllOfBillingInfoZonesInnerVirtualImages**](GetBillingAccount200ResponseAllOfBillingInfoZonesInnerVirtualImages.md) |  | [optional] 
**Snapshots** | Pointer to [**GetBillingAccount200ResponseAllOfBillingInfoZonesInnerSnapshots**](GetBillingAccount200ResponseAllOfBillingInfoZonesInnerSnapshots.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetBillingAccount200ResponseAllOfBillingInfoZonesInner

`func NewGetBillingAccount200ResponseAllOfBillingInfoZonesInner() *GetBillingAccount200ResponseAllOfBillingInfoZonesInner`

NewGetBillingAccount200ResponseAllOfBillingInfoZonesInner instantiates a new GetBillingAccount200ResponseAllOfBillingInfoZonesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBillingAccount200ResponseAllOfBillingInfoZonesInnerWithDefaults

`func NewGetBillingAccount200ResponseAllOfBillingInfoZonesInnerWithDefaults() *GetBillingAccount200ResponseAllOfBillingInfoZonesInner`

NewGetBillingAccount200ResponseAllOfBillingInfoZonesInnerWithDefaults instantiates a new GetBillingAccount200ResponseAllOfBillingInfoZonesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneId

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneUUID

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneUUID() string`

GetZoneUUID returns the ZoneUUID field if non-nil, zero value otherwise.

### GetZoneUUIDOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneUUIDOk() (*string, bool)`

GetZoneUUIDOk returns a tuple with the ZoneUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUUID

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetZoneUUID(v string)`

SetZoneUUID sets ZoneUUID field to given value.

### HasZoneUUID

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasZoneUUID() bool`

HasZoneUUID returns a boolean if a field has been set.

### GetZoneCode

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneCode() string`

GetZoneCode returns the ZoneCode field if non-nil, zero value otherwise.

### GetZoneCodeOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetZoneCodeOk() (*string, bool)`

GetZoneCodeOk returns a tuple with the ZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCode

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetZoneCode(v string)`

SetZoneCode sets ZoneCode field to given value.

### HasZoneCode

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasZoneCode() bool`

HasZoneCode returns a boolean if a field has been set.

### SetZoneCodeNil

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetZoneCodeNil(b bool)`

 SetZoneCodeNil sets the value for ZoneCode to be an explicit nil

### UnsetZoneCode
`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) UnsetZoneCode()`

UnsetZoneCode ensures that no value is present for ZoneCode, not even an explicit nil
### GetStartDate

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriceUnit

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetComputeServers

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetComputeServers() GetBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers`

GetComputeServers returns the ComputeServers field if non-nil, zero value otherwise.

### GetComputeServersOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetComputeServersOk() (*GetBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers, bool)`

GetComputeServersOk returns a tuple with the ComputeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServers

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetComputeServers(v GetBillingAccount200ResponseAllOfBillingInfoZonesInnerComputeServers)`

SetComputeServers sets ComputeServers field to given value.

### HasComputeServers

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasComputeServers() bool`

HasComputeServers returns a boolean if a field has been set.

### GetInstances

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetInstances() GetBillingAccount200ResponseAllOfBillingInfoZonesInnerInstances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetInstancesOk() (*GetBillingAccount200ResponseAllOfBillingInfoZonesInnerInstances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetInstances(v GetBillingAccount200ResponseAllOfBillingInfoZonesInnerInstances)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetDiscoveredServers

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetDiscoveredServers() GetBillingAccount200ResponseAllOfBillingInfoZonesInnerDiscoveredServers`

GetDiscoveredServers returns the DiscoveredServers field if non-nil, zero value otherwise.

### GetDiscoveredServersOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetDiscoveredServersOk() (*GetBillingAccount200ResponseAllOfBillingInfoZonesInnerDiscoveredServers, bool)`

GetDiscoveredServersOk returns a tuple with the DiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredServers

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetDiscoveredServers(v GetBillingAccount200ResponseAllOfBillingInfoZonesInnerDiscoveredServers)`

SetDiscoveredServers sets DiscoveredServers field to given value.

### HasDiscoveredServers

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasDiscoveredServers() bool`

HasDiscoveredServers returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetLoadBalancers() GetBillingAccount200ResponseAllOfBillingInfoZonesInnerLoadBalancers`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetLoadBalancersOk() (*GetBillingAccount200ResponseAllOfBillingInfoZonesInnerLoadBalancers, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetLoadBalancers(v GetBillingAccount200ResponseAllOfBillingInfoZonesInnerLoadBalancers)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetVirtualImages

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetVirtualImages() GetBillingAccount200ResponseAllOfBillingInfoZonesInnerVirtualImages`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetVirtualImagesOk() (*GetBillingAccount200ResponseAllOfBillingInfoZonesInnerVirtualImages, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetVirtualImages(v GetBillingAccount200ResponseAllOfBillingInfoZonesInnerVirtualImages)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetSnapshots

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetSnapshots() GetBillingAccount200ResponseAllOfBillingInfoZonesInnerSnapshots`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetSnapshotsOk() (*GetBillingAccount200ResponseAllOfBillingInfoZonesInnerSnapshots, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetSnapshots(v GetBillingAccount200ResponseAllOfBillingInfoZonesInnerSnapshots)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetPrice

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetBillingAccount200ResponseAllOfBillingInfoZonesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


