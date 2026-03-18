# GetBillingZoneIdentifier200ResponseAllOfBillingInfo

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

### NewGetBillingZoneIdentifier200ResponseAllOfBillingInfo

`func NewGetBillingZoneIdentifier200ResponseAllOfBillingInfo() *GetBillingZoneIdentifier200ResponseAllOfBillingInfo`

NewGetBillingZoneIdentifier200ResponseAllOfBillingInfo instantiates a new GetBillingZoneIdentifier200ResponseAllOfBillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBillingZoneIdentifier200ResponseAllOfBillingInfoWithDefaults

`func NewGetBillingZoneIdentifier200ResponseAllOfBillingInfoWithDefaults() *GetBillingZoneIdentifier200ResponseAllOfBillingInfo`

NewGetBillingZoneIdentifier200ResponseAllOfBillingInfoWithDefaults instantiates a new GetBillingZoneIdentifier200ResponseAllOfBillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneId

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneUUID

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetZoneUUID() string`

GetZoneUUID returns the ZoneUUID field if non-nil, zero value otherwise.

### GetZoneUUIDOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetZoneUUIDOk() (*string, bool)`

GetZoneUUIDOk returns a tuple with the ZoneUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUUID

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetZoneUUID(v string)`

SetZoneUUID sets ZoneUUID field to given value.

### HasZoneUUID

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasZoneUUID() bool`

HasZoneUUID returns a boolean if a field has been set.

### GetZoneCode

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetZoneCode() string`

GetZoneCode returns the ZoneCode field if non-nil, zero value otherwise.

### GetZoneCodeOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetZoneCodeOk() (*string, bool)`

GetZoneCodeOk returns a tuple with the ZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCode

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetZoneCode(v string)`

SetZoneCode sets ZoneCode field to given value.

### HasZoneCode

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasZoneCode() bool`

HasZoneCode returns a boolean if a field has been set.

### SetZoneCodeNil

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetZoneCodeNil(b bool)`

 SetZoneCodeNil sets the value for ZoneCode to be an explicit nil

### UnsetZoneCode
`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) UnsetZoneCode()`

UnsetZoneCode ensures that no value is present for ZoneCode, not even an explicit nil
### GetStartDate

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriceUnit

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetComputeServers

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetComputeServers() GetBillingZoneIdentifier200ResponseAllOfBillingInfoComputeServers`

GetComputeServers returns the ComputeServers field if non-nil, zero value otherwise.

### GetComputeServersOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetComputeServersOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoComputeServers, bool)`

GetComputeServersOk returns a tuple with the ComputeServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServers

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetComputeServers(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoComputeServers)`

SetComputeServers sets ComputeServers field to given value.

### HasComputeServers

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasComputeServers() bool`

HasComputeServers returns a boolean if a field has been set.

### GetInstances

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetInstances() GetBillingZoneIdentifier200ResponseAllOfBillingInfoInstances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetInstancesOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoInstances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetInstances(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoInstances)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetDiscoveredServers

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetDiscoveredServers() GetBillingZoneIdentifier200ResponseAllOfBillingInfoDiscoveredServers`

GetDiscoveredServers returns the DiscoveredServers field if non-nil, zero value otherwise.

### GetDiscoveredServersOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetDiscoveredServersOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoDiscoveredServers, bool)`

GetDiscoveredServersOk returns a tuple with the DiscoveredServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredServers

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetDiscoveredServers(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoDiscoveredServers)`

SetDiscoveredServers sets DiscoveredServers field to given value.

### HasDiscoveredServers

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasDiscoveredServers() bool`

HasDiscoveredServers returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetLoadBalancers() GetBillingZoneIdentifier200ResponseAllOfBillingInfoLoadBalancers`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetLoadBalancersOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoLoadBalancers, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetLoadBalancers(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoLoadBalancers)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetVirtualImages

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetVirtualImages() GetBillingZoneIdentifier200ResponseAllOfBillingInfoVirtualImages`

GetVirtualImages returns the VirtualImages field if non-nil, zero value otherwise.

### GetVirtualImagesOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetVirtualImagesOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoVirtualImages, bool)`

GetVirtualImagesOk returns a tuple with the VirtualImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImages

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetVirtualImages(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoVirtualImages)`

SetVirtualImages sets VirtualImages field to given value.

### HasVirtualImages

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasVirtualImages() bool`

HasVirtualImages returns a boolean if a field has been set.

### GetSnapshots

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetSnapshots() GetBillingZoneIdentifier200ResponseAllOfBillingInfoSnapshots`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetSnapshotsOk() (*GetBillingZoneIdentifier200ResponseAllOfBillingInfoSnapshots, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetSnapshots(v GetBillingZoneIdentifier200ResponseAllOfBillingInfoSnapshots)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetPrice

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCost

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetBillingZoneIdentifier200ResponseAllOfBillingInfo) HasCost() bool`

HasCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


