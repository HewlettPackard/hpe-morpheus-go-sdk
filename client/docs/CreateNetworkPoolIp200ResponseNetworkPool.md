# CreateNetworkPoolIp200ResponseNetworkPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**NetworkPoolId** | Pointer to **int64** |  | [optional] 
**IpType** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**GatewayAddress** | Pointer to **string** |  | [optional] 
**SubnetMask** | Pointer to **string** |  | [optional] 
**DnsServer** | Pointer to **string** |  | [optional] 
**InterfaceName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**StaticIp** | Pointer to **bool** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**PtrId** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**SubRefId** | Pointer to **int64** |  | [optional] 
**NetworkDomain** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 

## Methods

### NewCreateNetworkPoolIp200ResponseNetworkPool

`func NewCreateNetworkPoolIp200ResponseNetworkPool() *CreateNetworkPoolIp200ResponseNetworkPool`

NewCreateNetworkPoolIp200ResponseNetworkPool instantiates a new CreateNetworkPoolIp200ResponseNetworkPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkPoolIp200ResponseNetworkPoolWithDefaults

`func NewCreateNetworkPoolIp200ResponseNetworkPoolWithDefaults() *CreateNetworkPoolIp200ResponseNetworkPool`

NewCreateNetworkPoolIp200ResponseNetworkPoolWithDefaults instantiates a new CreateNetworkPoolIp200ResponseNetworkPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkPoolId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetNetworkPoolId() int64`

GetNetworkPoolId returns the NetworkPoolId field if non-nil, zero value otherwise.

### GetNetworkPoolIdOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetNetworkPoolIdOk() (*int64, bool)`

GetNetworkPoolIdOk returns a tuple with the NetworkPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetNetworkPoolId(v int64)`

SetNetworkPoolId sets NetworkPoolId field to given value.

### HasNetworkPoolId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasNetworkPoolId() bool`

HasNetworkPoolId returns a boolean if a field has been set.

### GetIpType

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetIpAddress

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetGatewayAddress

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetGatewayAddress() string`

GetGatewayAddress returns the GatewayAddress field if non-nil, zero value otherwise.

### GetGatewayAddressOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetGatewayAddressOk() (*string, bool)`

GetGatewayAddressOk returns a tuple with the GatewayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayAddress

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetGatewayAddress(v string)`

SetGatewayAddress sets GatewayAddress field to given value.

### HasGatewayAddress

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasGatewayAddress() bool`

HasGatewayAddress returns a boolean if a field has been set.

### GetSubnetMask

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetDnsServer

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetDnsServer() string`

GetDnsServer returns the DnsServer field if non-nil, zero value otherwise.

### GetDnsServerOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetDnsServerOk() (*string, bool)`

GetDnsServerOk returns a tuple with the DnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServer

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetDnsServer(v string)`

SetDnsServer sets DnsServer field to given value.

### HasDnsServer

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasDnsServer() bool`

HasDnsServer returns a boolean if a field has been set.

### GetInterfaceName

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActive

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStaticIp

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetStaticIp() bool`

GetStaticIp returns the StaticIp field if non-nil, zero value otherwise.

### GetStaticIpOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetStaticIpOk() (*bool, bool)`

GetStaticIpOk returns a tuple with the StaticIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIp

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetStaticIp(v bool)`

SetStaticIp sets StaticIp field to given value.

### HasStaticIp

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasStaticIp() bool`

HasStaticIp returns a boolean if a field has been set.

### GetFqdn

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetDomainName

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetHostname

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInternalId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetPtrId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetPtrId() string`

GetPtrId returns the PtrId field if non-nil, zero value otherwise.

### GetPtrIdOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetPtrIdOk() (*string, bool)`

GetPtrIdOk returns a tuple with the PtrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtrId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetPtrId(v string)`

SetPtrId sets PtrId field to given value.

### HasPtrId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasPtrId() bool`

HasPtrId returns a boolean if a field has been set.

### GetDateCreated

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRefType

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSubRefId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetSubRefId() int64`

GetSubRefId returns the SubRefId field if non-nil, zero value otherwise.

### GetSubRefIdOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetSubRefIdOk() (*int64, bool)`

GetSubRefIdOk returns a tuple with the SubRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetSubRefId(v int64)`

SetSubRefId sets SubRefId field to given value.

### HasSubRefId

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasSubRefId() bool`

HasSubRefId returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CreateNetworkPoolIp200ResponseNetworkPool) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


