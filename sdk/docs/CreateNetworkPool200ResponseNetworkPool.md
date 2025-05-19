# CreateNetworkPool200ResponseNetworkPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**DnsDomain** | Pointer to **string** |  | [optional] 
**DnsSearchPath** | Pointer to **string** |  | [optional] 
**HostPrefix** | Pointer to **string** |  | [optional] 
**HttpProxy** | Pointer to **string** |  | [optional] 
**DnsServers** | Pointer to **[]string** |  | [optional] 
**DnsSuffixList** | Pointer to **[]string** |  | [optional] 
**DhcpServer** | Pointer to **bool** |  | [optional] 
**DhcpIp** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**SubnetAddress** | Pointer to **string** |  | [optional] 
**IpCount** | Pointer to **int64** |  | [optional] 
**FreeCount** | Pointer to **int64** |  | [optional] 
**PoolEnabled** | Pointer to **bool** |  | [optional] 
**TftpServer** | Pointer to **string** |  | [optional] 
**BootFile** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**ParentType** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**PoolGroup** | Pointer to **string** |  | [optional] 
**IpRanges** | Pointer to [**[]CreateNetworkPool200ResponseNetworkPoolIpRangesInner**](CreateNetworkPool200ResponseNetworkPoolIpRangesInner.md) |  | [optional] 

## Methods

### NewCreateNetworkPool200ResponseNetworkPool

`func NewCreateNetworkPool200ResponseNetworkPool() *CreateNetworkPool200ResponseNetworkPool`

NewCreateNetworkPool200ResponseNetworkPool instantiates a new CreateNetworkPool200ResponseNetworkPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkPool200ResponseNetworkPoolWithDefaults

`func NewCreateNetworkPool200ResponseNetworkPoolWithDefaults() *CreateNetworkPool200ResponseNetworkPool`

NewCreateNetworkPool200ResponseNetworkPoolWithDefaults instantiates a new CreateNetworkPool200ResponseNetworkPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNetworkPool200ResponseNetworkPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkPool200ResponseNetworkPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkPool200ResponseNetworkPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CreateNetworkPool200ResponseNetworkPool) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkPool200ResponseNetworkPool) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *CreateNetworkPool200ResponseNetworkPool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *CreateNetworkPool200ResponseNetworkPool) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateNetworkPool200ResponseNetworkPool) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreateNetworkPool200ResponseNetworkPool) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCategory

`func (o *CreateNetworkPool200ResponseNetworkPool) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateNetworkPool200ResponseNetworkPool) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateNetworkPool200ResponseNetworkPool) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCode

`func (o *CreateNetworkPool200ResponseNetworkPool) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateNetworkPool200ResponseNetworkPool) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateNetworkPool200ResponseNetworkPool) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkPool200ResponseNetworkPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkPool200ResponseNetworkPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkPool200ResponseNetworkPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateNetworkPool200ResponseNetworkPool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateNetworkPool200ResponseNetworkPool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetInternalId

`func (o *CreateNetworkPool200ResponseNetworkPool) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *CreateNetworkPool200ResponseNetworkPool) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *CreateNetworkPool200ResponseNetworkPool) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateNetworkPool200ResponseNetworkPool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateNetworkPool200ResponseNetworkPool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateNetworkPool200ResponseNetworkPool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDnsDomain

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *CreateNetworkPool200ResponseNetworkPool) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *CreateNetworkPool200ResponseNetworkPool) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.

### GetDnsSearchPath

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDnsSearchPath() string`

GetDnsSearchPath returns the DnsSearchPath field if non-nil, zero value otherwise.

### GetDnsSearchPathOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDnsSearchPathOk() (*string, bool)`

GetDnsSearchPathOk returns a tuple with the DnsSearchPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSearchPath

`func (o *CreateNetworkPool200ResponseNetworkPool) SetDnsSearchPath(v string)`

SetDnsSearchPath sets DnsSearchPath field to given value.

### HasDnsSearchPath

`func (o *CreateNetworkPool200ResponseNetworkPool) HasDnsSearchPath() bool`

HasDnsSearchPath returns a boolean if a field has been set.

### GetHostPrefix

`func (o *CreateNetworkPool200ResponseNetworkPool) GetHostPrefix() string`

GetHostPrefix returns the HostPrefix field if non-nil, zero value otherwise.

### GetHostPrefixOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetHostPrefixOk() (*string, bool)`

GetHostPrefixOk returns a tuple with the HostPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPrefix

`func (o *CreateNetworkPool200ResponseNetworkPool) SetHostPrefix(v string)`

SetHostPrefix sets HostPrefix field to given value.

### HasHostPrefix

`func (o *CreateNetworkPool200ResponseNetworkPool) HasHostPrefix() bool`

HasHostPrefix returns a boolean if a field has been set.

### GetHttpProxy

`func (o *CreateNetworkPool200ResponseNetworkPool) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *CreateNetworkPool200ResponseNetworkPool) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *CreateNetworkPool200ResponseNetworkPool) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### GetDnsServers

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *CreateNetworkPool200ResponseNetworkPool) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *CreateNetworkPool200ResponseNetworkPool) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffixList

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDnsSuffixList() []string`

GetDnsSuffixList returns the DnsSuffixList field if non-nil, zero value otherwise.

### GetDnsSuffixListOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDnsSuffixListOk() (*[]string, bool)`

GetDnsSuffixListOk returns a tuple with the DnsSuffixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffixList

`func (o *CreateNetworkPool200ResponseNetworkPool) SetDnsSuffixList(v []string)`

SetDnsSuffixList sets DnsSuffixList field to given value.

### HasDnsSuffixList

`func (o *CreateNetworkPool200ResponseNetworkPool) HasDnsSuffixList() bool`

HasDnsSuffixList returns a boolean if a field has been set.

### GetDhcpServer

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *CreateNetworkPool200ResponseNetworkPool) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *CreateNetworkPool200ResponseNetworkPool) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpIp

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *CreateNetworkPool200ResponseNetworkPool) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *CreateNetworkPool200ResponseNetworkPool) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### GetGateway

`func (o *CreateNetworkPool200ResponseNetworkPool) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreateNetworkPool200ResponseNetworkPool) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CreateNetworkPool200ResponseNetworkPool) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *CreateNetworkPool200ResponseNetworkPool) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *CreateNetworkPool200ResponseNetworkPool) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *CreateNetworkPool200ResponseNetworkPool) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetSubnetAddress

`func (o *CreateNetworkPool200ResponseNetworkPool) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *CreateNetworkPool200ResponseNetworkPool) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *CreateNetworkPool200ResponseNetworkPool) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### GetIpCount

`func (o *CreateNetworkPool200ResponseNetworkPool) GetIpCount() int64`

GetIpCount returns the IpCount field if non-nil, zero value otherwise.

### GetIpCountOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetIpCountOk() (*int64, bool)`

GetIpCountOk returns a tuple with the IpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpCount

`func (o *CreateNetworkPool200ResponseNetworkPool) SetIpCount(v int64)`

SetIpCount sets IpCount field to given value.

### HasIpCount

`func (o *CreateNetworkPool200ResponseNetworkPool) HasIpCount() bool`

HasIpCount returns a boolean if a field has been set.

### GetFreeCount

`func (o *CreateNetworkPool200ResponseNetworkPool) GetFreeCount() int64`

GetFreeCount returns the FreeCount field if non-nil, zero value otherwise.

### GetFreeCountOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetFreeCountOk() (*int64, bool)`

GetFreeCountOk returns a tuple with the FreeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCount

`func (o *CreateNetworkPool200ResponseNetworkPool) SetFreeCount(v int64)`

SetFreeCount sets FreeCount field to given value.

### HasFreeCount

`func (o *CreateNetworkPool200ResponseNetworkPool) HasFreeCount() bool`

HasFreeCount returns a boolean if a field has been set.

### GetPoolEnabled

`func (o *CreateNetworkPool200ResponseNetworkPool) GetPoolEnabled() bool`

GetPoolEnabled returns the PoolEnabled field if non-nil, zero value otherwise.

### GetPoolEnabledOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetPoolEnabledOk() (*bool, bool)`

GetPoolEnabledOk returns a tuple with the PoolEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEnabled

`func (o *CreateNetworkPool200ResponseNetworkPool) SetPoolEnabled(v bool)`

SetPoolEnabled sets PoolEnabled field to given value.

### HasPoolEnabled

`func (o *CreateNetworkPool200ResponseNetworkPool) HasPoolEnabled() bool`

HasPoolEnabled returns a boolean if a field has been set.

### GetTftpServer

`func (o *CreateNetworkPool200ResponseNetworkPool) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *CreateNetworkPool200ResponseNetworkPool) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *CreateNetworkPool200ResponseNetworkPool) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### GetBootFile

`func (o *CreateNetworkPool200ResponseNetworkPool) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *CreateNetworkPool200ResponseNetworkPool) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *CreateNetworkPool200ResponseNetworkPool) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### GetRefType

`func (o *CreateNetworkPool200ResponseNetworkPool) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *CreateNetworkPool200ResponseNetworkPool) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *CreateNetworkPool200ResponseNetworkPool) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *CreateNetworkPool200ResponseNetworkPool) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *CreateNetworkPool200ResponseNetworkPool) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *CreateNetworkPool200ResponseNetworkPool) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetParentType

`func (o *CreateNetworkPool200ResponseNetworkPool) GetParentType() string`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetParentTypeOk() (*string, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *CreateNetworkPool200ResponseNetworkPool) SetParentType(v string)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *CreateNetworkPool200ResponseNetworkPool) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetParentId

`func (o *CreateNetworkPool200ResponseNetworkPool) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateNetworkPool200ResponseNetworkPool) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateNetworkPool200ResponseNetworkPool) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPoolGroup

`func (o *CreateNetworkPool200ResponseNetworkPool) GetPoolGroup() string`

GetPoolGroup returns the PoolGroup field if non-nil, zero value otherwise.

### GetPoolGroupOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetPoolGroupOk() (*string, bool)`

GetPoolGroupOk returns a tuple with the PoolGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolGroup

`func (o *CreateNetworkPool200ResponseNetworkPool) SetPoolGroup(v string)`

SetPoolGroup sets PoolGroup field to given value.

### HasPoolGroup

`func (o *CreateNetworkPool200ResponseNetworkPool) HasPoolGroup() bool`

HasPoolGroup returns a boolean if a field has been set.

### GetIpRanges

`func (o *CreateNetworkPool200ResponseNetworkPool) GetIpRanges() []CreateNetworkPool200ResponseNetworkPoolIpRangesInner`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *CreateNetworkPool200ResponseNetworkPool) GetIpRangesOk() (*[]CreateNetworkPool200ResponseNetworkPoolIpRangesInner, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *CreateNetworkPool200ResponseNetworkPool) SetIpRanges(v []CreateNetworkPool200ResponseNetworkPoolIpRangesInner)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *CreateNetworkPool200ResponseNetworkPool) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


