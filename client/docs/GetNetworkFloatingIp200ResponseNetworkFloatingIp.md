# GetNetworkFloatingIp200ResponseNetworkFloatingIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Cloud** | Pointer to [**GetNetworkFloatingIp200ResponseNetworkFloatingIpCloud**](GetNetworkFloatingIp200ResponseNetworkFloatingIpCloud.md) |  | [optional] 
**Server** | Pointer to [**GetNetworkFloatingIp200ResponseNetworkFloatingIpServer**](GetNetworkFloatingIp200ResponseNetworkFloatingIpServer.md) |  | [optional] 
**IpStatus** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** | IP Address | [optional] 
**IpRange** | Pointer to **string** |  | [optional] 
**PtrId** | Pointer to **string** |  | [optional] 
**NetworkDomain** | Pointer to [**ListNetworks200ResponseAllOfNetworksInnerNetworkDomain**](ListNetworks200ResponseAllOfNetworksInnerNetworkDomain.md) |  | [optional] 
**CreatedBy** | Pointer to [**ListBackupResults200ResponseAllOfResultsInnerCreatedBy**](ListBackupResults200ResponseAllOfResultsInnerCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetNetworkFloatingIp200ResponseNetworkFloatingIp

`func NewGetNetworkFloatingIp200ResponseNetworkFloatingIp() *GetNetworkFloatingIp200ResponseNetworkFloatingIp`

NewGetNetworkFloatingIp200ResponseNetworkFloatingIp instantiates a new GetNetworkFloatingIp200ResponseNetworkFloatingIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFloatingIp200ResponseNetworkFloatingIpWithDefaults

`func NewGetNetworkFloatingIp200ResponseNetworkFloatingIpWithDefaults() *GetNetworkFloatingIp200ResponseNetworkFloatingIp`

NewGetNetworkFloatingIp200ResponseNetworkFloatingIpWithDefaults instantiates a new GetNetworkFloatingIp200ResponseNetworkFloatingIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCloud

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetCloud() GetNetworkFloatingIp200ResponseNetworkFloatingIpCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetCloudOk() (*GetNetworkFloatingIp200ResponseNetworkFloatingIpCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetCloud(v GetNetworkFloatingIp200ResponseNetworkFloatingIpCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetServer

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetServer() GetNetworkFloatingIp200ResponseNetworkFloatingIpServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetServerOk() (*GetNetworkFloatingIp200ResponseNetworkFloatingIpServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetServer(v GetNetworkFloatingIp200ResponseNetworkFloatingIpServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetIpStatus

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetIpStatus() string`

GetIpStatus returns the IpStatus field if non-nil, zero value otherwise.

### GetIpStatusOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetIpStatusOk() (*string, bool)`

GetIpStatusOk returns a tuple with the IpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpStatus

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetIpStatus(v string)`

SetIpStatus sets IpStatus field to given value.

### HasIpStatus

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasIpStatus() bool`

HasIpStatus returns a boolean if a field has been set.

### GetIpAddress

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpRange

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetPtrId

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetPtrId() string`

GetPtrId returns the PtrId field if non-nil, zero value otherwise.

### GetPtrIdOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetPtrIdOk() (*string, bool)`

GetPtrIdOk returns a tuple with the PtrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrId

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetPtrId(v string)`

SetPtrId sets PtrId field to given value.

### HasPtrId

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasPtrId() bool`

HasPtrId returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetNetworkDomain() ListNetworks200ResponseAllOfNetworksInnerNetworkDomain`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetNetworkDomainOk() (*ListNetworks200ResponseAllOfNetworksInnerNetworkDomain, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetNetworkDomain(v ListNetworks200ResponseAllOfNetworksInnerNetworkDomain)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetCreatedBy() ListBackupResults200ResponseAllOfResultsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetCreatedByOk() (*ListBackupResults200ResponseAllOfResultsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetCreatedBy(v ListBackupResults200ResponseAllOfResultsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkFloatingIp200ResponseNetworkFloatingIp) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


