# AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Cloud** | Pointer to [**AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpCloud**](AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpCloud.md) |  | [optional] 
**Server** | Pointer to [**AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpServer**](AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpServer.md) |  | [optional] 
**IpStatus** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** | IP Address | [optional] 
**IpRange** | Pointer to **string** |  | [optional] 
**PtrId** | Pointer to **string** |  | [optional] 
**NetworkDomain** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerNetworkDomain**](ListNetworks200ResponseAllOfNetworksInnerNetworkDomain.md) |  | [optional] 
**CreatedBy** | Pointer to [**ListBackupResults200ResponseAllOfResultsInnerCreatedBy**](ListBackupResults200ResponseAllOfResultsInnerCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp

`func NewAllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp() *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp`

NewAllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp instantiates a new AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpWithDefaults

`func NewAllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpWithDefaults() *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp`

NewAllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpWithDefaults instantiates a new AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCloud

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetCloud() AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetCloudOk() (*AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetCloud(v AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetServer

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetServer() AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetServerOk() (*AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetServer(v AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIpServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetIpStatus

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetIpStatus() string`

GetIpStatus returns the IpStatus field if non-nil, zero value otherwise.

### GetIpStatusOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetIpStatusOk() (*string, bool)`

GetIpStatusOk returns a tuple with the IpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStatus

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetIpStatus(v string)`

SetIpStatus sets IpStatus field to given value.

### HasIpStatus

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasIpStatus() bool`

HasIpStatus returns a boolean if a field has been set.

### GetIpAddress

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpRange

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetPtrId

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetPtrId() string`

GetPtrId returns the PtrId field if non-nil, zero value otherwise.

### GetPtrIdOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetPtrIdOk() (*string, bool)`

GetPtrIdOk returns a tuple with the PtrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrId

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetPtrId(v string)`

SetPtrId sets PtrId field to given value.

### HasPtrId

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasPtrId() bool`

HasPtrId returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetNetworkDomain() ListNetworks200ResponseAllOfNetworksInnerNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetNetworkDomainOk() (*ListNetworks200ResponseAllOfNetworksInnerNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetNetworkDomain(v ListNetworks200ResponseAllOfNetworksInnerNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetCreatedBy() ListBackupResults200ResponseAllOfResultsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetCreatedByOk() (*ListBackupResults200ResponseAllOfResultsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetCreatedBy(v ListBackupResults200ResponseAllOfResultsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AllocateNetworkFloatingIp200ResponseAllOfNetworkFloatingIp) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


