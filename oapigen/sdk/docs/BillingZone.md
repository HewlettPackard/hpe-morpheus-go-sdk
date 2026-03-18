# BillingZone

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
**ComputeServers** | Pointer to [**GetBillingZoneIdentifier200ResponseAllOfBillingInfoComputeServers**](GetBillingZoneIdentifier200ResponseAllOfBillingInfoComputeServers.md) |  | [optional] 
**Instances** | Pointer to [**GetBillingZoneIdentifier200ResponseAllOfBillingInfoInstances**](GetBillingZoneIdentifier200ResponseAllOfBillingInfoInstances.md) |  | [optional] 
**DiscoveredServers** | Pointer to [**GetBillingZoneIdentifier200ResponseAllOfBillingInfoDiscoveredServers**](GetBillingZoneIdentifier200ResponseAllOfBillingInfoDiscoveredServers.md) |  | [optional] 
**LoadBalancers** | Pointer to [**GetBillingZoneIdentifier200ResponseAllOfBillingInfoLoadBalancers**](GetBillingZoneIdentifier200ResponseAllOfBillingInfoLoadBalancers.md) |  | [optional] 
**VirtualImages** | Pointer to [**GetBillingZoneIdentifier200ResponseAllOfBillingInfoVirtualImages**](GetBillingZoneIdentifier200ResponseAllOfBillingInfoVirtualImages.md) |  | [optional] 
**Snapshots** | Pointer to [**GetBillingZoneIdentifier200ResponseAllOfBillingInfoSnapshots**](GetBillingZoneIdentifier200ResponseAllOfBillingInfoSnapshots.md) |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 

## Methods

### NewBillingZone

`func NewBillingZone() *BillingZone`

NewBillingZone instantiates a new BillingZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingZoneWithDefaults

`func NewBillingZoneWithDefaults() *BillingZone`

NewBillingZoneWithDefaults instantiates a new BillingZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *BillingZone) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *BillingZone) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *BillingZone) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *BillingZone) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneId

`func (o *BillingZone) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *BillingZone) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *BillingZone) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *BillingZone) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneUUID

`func (o *BillingZone) GetZoneUUID() string`

GetZoneUUID returns the ZoneUUID field if non-nil, zero value otherwise.

### GetZoneUUIDOk

`func (o *BillingZone) GetZoneUUIDOk() (*string, bool)`

GetZoneUUIDOk returns a tuple with the ZoneUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUUID

`func (o *BillingZone) SetZoneUUID(v string)`

SetZoneUUID sets ZoneUUID field to given value.

### HasZoneUUID

`func (o *BillingZone) HasZoneUUID() bool`

HasZoneUUID returns a boolean if a field has been set.

### GetZoneCode

`func (o *BillingZone) GetZoneCode() string`

GetZoneCode returns the ZoneCode field if non-nil, zero value otherwise.

### GetZoneCodeOk

`func (o *BillingZone) GetZoneCodeOk() (*string, bool)`

GetZoneCodeOk returns a tuple with the ZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCode

`func (o *BillingZone) SetZoneCode(v string)`

SetZoneCode sets ZoneCode field to given value.

### HasZoneCode

`func (o *BillingZone) HasZoneCode() bool`

HasZoneCode returns a boolean if a field has been set.

### SetZoneCodeNil

`func (o *BillingZone) SetZoneCodeNil(b bool)`

 SetZoneCodeNil sets the value for ZoneCode to be an explicit nil

### UnsetZoneCode
`func (o *BillingZone) UnsetZoneCode()`

UnsetZoneCode ensures that no value is present for ZoneCode, not even an explicit nil
### GetStartDate

`func (o *BillingZone) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BillingZone) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BillingZone) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BillingZone) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *BillingZone) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BillingZone) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BillingZone) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BillingZone) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriceUnit

`func (o *BillingZone) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *BillingZone) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *BillingZone) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *BillingZone) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetComputeServers

`func (o *BillingZone) GetComputeServers() GetBillingZoneIdentifier200ResponseAllOfBillingInfoComputeServers`

GetComputeServers returns the ComputeServers field if non-nil, zero value otherwise.

### GetComputeServersOk

`func (o *BillingZone) GetComputeServersOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoComputeServers, bool)`

GetComputeServersOk returns a tuple with the ComputeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServers

`func (o *BillingZone) SetComputeServers(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoComputeServers)`

SetComputeServers sets ComputeServers field to given value.

### HasComputeServers

`func (o *BillingZone) HasComputeServers() bool`

HasComputeServers returns a boolean if a field has been set.

### GetInstances

`func (o *BillingZone) GetInstances() GetBillingZoneIdentifier200ResponseAllOfBillingInfoInstances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *BillingZone) GetInstancesOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoInstances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *BillingZone) SetInstances(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoInstances)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *BillingZone) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetDiscoveredServers

`func (o *BillingZone) GetDiscoveredServers() GetBillingZoneIdentifier200ResponseAllOfBillingInfoDiscoveredServers`

GetDiscoveredServers returns the DiscoveredServers field if non-nil, zero value otherwise.

### GetDiscoveredServersOk

`func (o *BillingZone) GetDiscoveredServersOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoDiscoveredServers, bool)`

GetDiscoveredServersOk returns a tuple with the DiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredServers

`func (o *BillingZone) SetDiscoveredServers(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoDiscoveredServers)`

SetDiscoveredServers sets DiscoveredServers field to given value.

### HasDiscoveredServers

`func (o *BillingZone) HasDiscoveredServers() bool`

HasDiscoveredServers returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *BillingZone) GetLoadBalancers() GetBillingZoneIdentifier200ResponseAllOfBillingInfoLoadBalancers`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BillingZone) GetLoadBalancersOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoLoadBalancers, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BillingZone) SetLoadBalancers(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoLoadBalancers)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *BillingZone) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetVirtualImages

`func (o *BillingZone) GetVirtualImages() GetBillingZoneIdentifier200ResponseAllOfBillingInfoVirtualImages`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *BillingZone) GetVirtualImagesOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoVirtualImages, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *BillingZone) SetVirtualImages(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoVirtualImages)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *BillingZone) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetSnapshots

`func (o *BillingZone) GetSnapshots() GetBillingZoneIdentifier200ResponseAllOfBillingInfoSnapshots`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *BillingZone) GetSnapshotsOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoSnapshots, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *BillingZone) SetSnapshots(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoSnapshots)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *BillingZone) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetPrice

`func (o *BillingZone) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillingZone) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillingZone) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *BillingZone) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *BillingZone) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *BillingZone) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *BillingZone) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *BillingZone) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


