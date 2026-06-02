# GetNetworkPool200ResponseNetworkPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**GetNetworkPool200ResponseNetworkPoolType**](GetNetworkPool200ResponseNetworkPoolType.md) |  | [optional] 
**Account** | Pointer to [**GetNetworkPool200ResponseNetworkPoolAccount**](GetNetworkPool200ResponseNetworkPoolAccount.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**DnsDomain** | Pointer to **NullableString** |  | [optional] 
**DnsSearchPath** | Pointer to **NullableString** |  | [optional] 
**HostPrefix** | Pointer to **NullableString** |  | [optional] 
**HttpProxy** | Pointer to **NullableString** |  | [optional] 
**DnsServers** | Pointer to **[]string** |  | [optional] 
**DnsSuffixList** | Pointer to **[]string** |  | [optional] 
**DhcpServer** | Pointer to **bool** |  | [optional] 
**DhcpIp** | Pointer to **NullableString** |  | [optional] 
**Gateway** | Pointer to **NullableString** |  | [optional] 
**Netmask** | Pointer to **NullableString** |  | [optional] 
**SubnetAddress** | Pointer to **NullableString** |  | [optional] 
**IpCount** | Pointer to **int64** |  | [optional] 
**FreeCount** | Pointer to **int64** |  | [optional] 
**PoolEnabled** | Pointer to **bool** |  | [optional] 
**TftpServer** | Pointer to **NullableString** |  | [optional] 
**BootFile** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableString** |  | [optional] 
**ParentType** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**PoolGroup** | Pointer to **NullableString** |  | [optional] 
**IpRanges** | Pointer to [**[]GetNetworkPool200ResponseNetworkPoolIpRangesInner**](GetNetworkPool200ResponseNetworkPoolIpRangesInner.md) |  | [optional] 

## Methods

### NewGetNetworkPool200ResponseNetworkPool

`func NewGetNetworkPool200ResponseNetworkPool() *GetNetworkPool200ResponseNetworkPool`

NewGetNetworkPool200ResponseNetworkPool instantiates a new GetNetworkPool200ResponseNetworkPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetNetworkPool200ResponseNetworkPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkPool200ResponseNetworkPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkPool200ResponseNetworkPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkPool200ResponseNetworkPool) GetType() GetNetworkPool200ResponseNetworkPoolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetTypeOk() (*GetNetworkPool200ResponseNetworkPoolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkPool200ResponseNetworkPool) SetType(v GetNetworkPool200ResponseNetworkPoolType)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkPool200ResponseNetworkPool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *GetNetworkPool200ResponseNetworkPool) GetAccount() GetNetworkPool200ResponseNetworkPoolAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetAccountOk() (*GetNetworkPool200ResponseNetworkPoolAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetNetworkPool200ResponseNetworkPool) SetAccount(v GetNetworkPool200ResponseNetworkPoolAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetNetworkPool200ResponseNetworkPool) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCategory

`func (o *GetNetworkPool200ResponseNetworkPool) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetNetworkPool200ResponseNetworkPool) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetNetworkPool200ResponseNetworkPool) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCode

`func (o *GetNetworkPool200ResponseNetworkPool) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkPool200ResponseNetworkPool) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkPool200ResponseNetworkPool) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetName

`func (o *GetNetworkPool200ResponseNetworkPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkPool200ResponseNetworkPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkPool200ResponseNetworkPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetNetworkPool200ResponseNetworkPool) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetNetworkPool200ResponseNetworkPool) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetNetworkPool200ResponseNetworkPool) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetInternalId

`func (o *GetNetworkPool200ResponseNetworkPool) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkPool200ResponseNetworkPool) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkPool200ResponseNetworkPool) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetNetworkPool200ResponseNetworkPool) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkPool200ResponseNetworkPool) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkPool200ResponseNetworkPool) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDnsDomain

`func (o *GetNetworkPool200ResponseNetworkPool) GetDnsDomain() string`

GetDnsDomain returns the DnsDomain field if non-nil, zero value otherwise.

### GetDnsDomainOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetDnsDomainOk() (*string, bool)`

GetDnsDomainOk returns a tuple with the DnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsDomain

`func (o *GetNetworkPool200ResponseNetworkPool) SetDnsDomain(v string)`

SetDnsDomain sets DnsDomain field to given value.

### HasDnsDomain

`func (o *GetNetworkPool200ResponseNetworkPool) HasDnsDomain() bool`

HasDnsDomain returns a boolean if a field has been set.

### SetDnsDomainNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetDnsDomainNil(b bool)`

 SetDnsDomainNil sets the value for DnsDomain to be an explicit nil

### UnsetDnsDomain
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetDnsDomain()`

UnsetDnsDomain ensures that no value is present for DnsDomain, not even an explicit nil
### GetDnsSearchPath

`func (o *GetNetworkPool200ResponseNetworkPool) GetDnsSearchPath() string`

GetDnsSearchPath returns the DnsSearchPath field if non-nil, zero value otherwise.

### GetDnsSearchPathOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetDnsSearchPathOk() (*string, bool)`

GetDnsSearchPathOk returns a tuple with the DnsSearchPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSearchPath

`func (o *GetNetworkPool200ResponseNetworkPool) SetDnsSearchPath(v string)`

SetDnsSearchPath sets DnsSearchPath field to given value.

### HasDnsSearchPath

`func (o *GetNetworkPool200ResponseNetworkPool) HasDnsSearchPath() bool`

HasDnsSearchPath returns a boolean if a field has been set.

### SetDnsSearchPathNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetDnsSearchPathNil(b bool)`

 SetDnsSearchPathNil sets the value for DnsSearchPath to be an explicit nil

### UnsetDnsSearchPath
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetDnsSearchPath()`

UnsetDnsSearchPath ensures that no value is present for DnsSearchPath, not even an explicit nil
### GetHostPrefix

`func (o *GetNetworkPool200ResponseNetworkPool) GetHostPrefix() string`

GetHostPrefix returns the HostPrefix field if non-nil, zero value otherwise.

### GetHostPrefixOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetHostPrefixOk() (*string, bool)`

GetHostPrefixOk returns a tuple with the HostPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPrefix

`func (o *GetNetworkPool200ResponseNetworkPool) SetHostPrefix(v string)`

SetHostPrefix sets HostPrefix field to given value.

### HasHostPrefix

`func (o *GetNetworkPool200ResponseNetworkPool) HasHostPrefix() bool`

HasHostPrefix returns a boolean if a field has been set.

### SetHostPrefixNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetHostPrefixNil(b bool)`

 SetHostPrefixNil sets the value for HostPrefix to be an explicit nil

### UnsetHostPrefix
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetHostPrefix()`

UnsetHostPrefix ensures that no value is present for HostPrefix, not even an explicit nil
### GetHttpProxy

`func (o *GetNetworkPool200ResponseNetworkPool) GetHttpProxy() string`

GetHttpProxy returns the HttpProxy field if non-nil, zero value otherwise.

### GetHttpProxyOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetHttpProxyOk() (*string, bool)`

GetHttpProxyOk returns a tuple with the HttpProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxy

`func (o *GetNetworkPool200ResponseNetworkPool) SetHttpProxy(v string)`

SetHttpProxy sets HttpProxy field to given value.

### HasHttpProxy

`func (o *GetNetworkPool200ResponseNetworkPool) HasHttpProxy() bool`

HasHttpProxy returns a boolean if a field has been set.

### SetHttpProxyNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetHttpProxyNil(b bool)`

 SetHttpProxyNil sets the value for HttpProxy to be an explicit nil

### UnsetHttpProxy
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetHttpProxy()`

UnsetHttpProxy ensures that no value is present for HttpProxy, not even an explicit nil
### GetDnsServers

`func (o *GetNetworkPool200ResponseNetworkPool) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *GetNetworkPool200ResponseNetworkPool) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.

### HasDnsServers

`func (o *GetNetworkPool200ResponseNetworkPool) HasDnsServers() bool`

HasDnsServers returns a boolean if a field has been set.

### GetDnsSuffixList

`func (o *GetNetworkPool200ResponseNetworkPool) GetDnsSuffixList() []string`

GetDnsSuffixList returns the DnsSuffixList field if non-nil, zero value otherwise.

### GetDnsSuffixListOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetDnsSuffixListOk() (*[]string, bool)`

GetDnsSuffixListOk returns a tuple with the DnsSuffixList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSuffixList

`func (o *GetNetworkPool200ResponseNetworkPool) SetDnsSuffixList(v []string)`

SetDnsSuffixList sets DnsSuffixList field to given value.

### HasDnsSuffixList

`func (o *GetNetworkPool200ResponseNetworkPool) HasDnsSuffixList() bool`

HasDnsSuffixList returns a boolean if a field has been set.

### GetDhcpServer

`func (o *GetNetworkPool200ResponseNetworkPool) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *GetNetworkPool200ResponseNetworkPool) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *GetNetworkPool200ResponseNetworkPool) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetDhcpIp

`func (o *GetNetworkPool200ResponseNetworkPool) GetDhcpIp() string`

GetDhcpIp returns the DhcpIp field if non-nil, zero value otherwise.

### GetDhcpIpOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetDhcpIpOk() (*string, bool)`

GetDhcpIpOk returns a tuple with the DhcpIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIp

`func (o *GetNetworkPool200ResponseNetworkPool) SetDhcpIp(v string)`

SetDhcpIp sets DhcpIp field to given value.

### HasDhcpIp

`func (o *GetNetworkPool200ResponseNetworkPool) HasDhcpIp() bool`

HasDhcpIp returns a boolean if a field has been set.

### SetDhcpIpNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetDhcpIpNil(b bool)`

 SetDhcpIpNil sets the value for DhcpIp to be an explicit nil

### UnsetDhcpIp
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetDhcpIp()`

UnsetDhcpIp ensures that no value is present for DhcpIp, not even an explicit nil
### GetGateway

`func (o *GetNetworkPool200ResponseNetworkPool) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetNetworkPool200ResponseNetworkPool) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetNetworkPool200ResponseNetworkPool) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### SetGatewayNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetGatewayNil(b bool)`

 SetGatewayNil sets the value for Gateway to be an explicit nil

### UnsetGateway
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetGateway()`

UnsetGateway ensures that no value is present for Gateway, not even an explicit nil
### GetNetmask

`func (o *GetNetworkPool200ResponseNetworkPool) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *GetNetworkPool200ResponseNetworkPool) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *GetNetworkPool200ResponseNetworkPool) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### SetNetmaskNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetNetmaskNil(b bool)`

 SetNetmaskNil sets the value for Netmask to be an explicit nil

### UnsetNetmask
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetNetmask()`

UnsetNetmask ensures that no value is present for Netmask, not even an explicit nil
### GetSubnetAddress

`func (o *GetNetworkPool200ResponseNetworkPool) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *GetNetworkPool200ResponseNetworkPool) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *GetNetworkPool200ResponseNetworkPool) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### SetSubnetAddressNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetSubnetAddressNil(b bool)`

 SetSubnetAddressNil sets the value for SubnetAddress to be an explicit nil

### UnsetSubnetAddress
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetSubnetAddress()`

UnsetSubnetAddress ensures that no value is present for SubnetAddress, not even an explicit nil
### GetIpCount

`func (o *GetNetworkPool200ResponseNetworkPool) GetIpCount() int64`

GetIpCount returns the IpCount field if non-nil, zero value otherwise.

### GetIpCountOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetIpCountOk() (*int64, bool)`

GetIpCountOk returns a tuple with the IpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpCount

`func (o *GetNetworkPool200ResponseNetworkPool) SetIpCount(v int64)`

SetIpCount sets IpCount field to given value.

### HasIpCount

`func (o *GetNetworkPool200ResponseNetworkPool) HasIpCount() bool`

HasIpCount returns a boolean if a field has been set.

### GetFreeCount

`func (o *GetNetworkPool200ResponseNetworkPool) GetFreeCount() int64`

GetFreeCount returns the FreeCount field if non-nil, zero value otherwise.

### GetFreeCountOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetFreeCountOk() (*int64, bool)`

GetFreeCountOk returns a tuple with the FreeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCount

`func (o *GetNetworkPool200ResponseNetworkPool) SetFreeCount(v int64)`

SetFreeCount sets FreeCount field to given value.

### HasFreeCount

`func (o *GetNetworkPool200ResponseNetworkPool) HasFreeCount() bool`

HasFreeCount returns a boolean if a field has been set.

### GetPoolEnabled

`func (o *GetNetworkPool200ResponseNetworkPool) GetPoolEnabled() bool`

GetPoolEnabled returns the PoolEnabled field if non-nil, zero value otherwise.

### GetPoolEnabledOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetPoolEnabledOk() (*bool, bool)`

GetPoolEnabledOk returns a tuple with the PoolEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEnabled

`func (o *GetNetworkPool200ResponseNetworkPool) SetPoolEnabled(v bool)`

SetPoolEnabled sets PoolEnabled field to given value.

### HasPoolEnabled

`func (o *GetNetworkPool200ResponseNetworkPool) HasPoolEnabled() bool`

HasPoolEnabled returns a boolean if a field has been set.

### GetTftpServer

`func (o *GetNetworkPool200ResponseNetworkPool) GetTftpServer() string`

GetTftpServer returns the TftpServer field if non-nil, zero value otherwise.

### GetTftpServerOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetTftpServerOk() (*string, bool)`

GetTftpServerOk returns a tuple with the TftpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServer

`func (o *GetNetworkPool200ResponseNetworkPool) SetTftpServer(v string)`

SetTftpServer sets TftpServer field to given value.

### HasTftpServer

`func (o *GetNetworkPool200ResponseNetworkPool) HasTftpServer() bool`

HasTftpServer returns a boolean if a field has been set.

### SetTftpServerNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetTftpServerNil(b bool)`

 SetTftpServerNil sets the value for TftpServer to be an explicit nil

### UnsetTftpServer
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetTftpServer()`

UnsetTftpServer ensures that no value is present for TftpServer, not even an explicit nil
### GetBootFile

`func (o *GetNetworkPool200ResponseNetworkPool) GetBootFile() string`

GetBootFile returns the BootFile field if non-nil, zero value otherwise.

### GetBootFileOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetBootFileOk() (*string, bool)`

GetBootFileOk returns a tuple with the BootFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFile

`func (o *GetNetworkPool200ResponseNetworkPool) SetBootFile(v string)`

SetBootFile sets BootFile field to given value.

### HasBootFile

`func (o *GetNetworkPool200ResponseNetworkPool) HasBootFile() bool`

HasBootFile returns a boolean if a field has been set.

### SetBootFileNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetBootFileNil(b bool)`

 SetBootFileNil sets the value for BootFile to be an explicit nil

### UnsetBootFile
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetBootFile()`

UnsetBootFile ensures that no value is present for BootFile, not even an explicit nil
### GetRefType

`func (o *GetNetworkPool200ResponseNetworkPool) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetNetworkPool200ResponseNetworkPool) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetNetworkPool200ResponseNetworkPool) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *GetNetworkPool200ResponseNetworkPool) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetNetworkPool200ResponseNetworkPool) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetNetworkPool200ResponseNetworkPool) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetParentType

`func (o *GetNetworkPool200ResponseNetworkPool) GetParentType() string`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetParentTypeOk() (*string, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *GetNetworkPool200ResponseNetworkPool) SetParentType(v string)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *GetNetworkPool200ResponseNetworkPool) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### SetParentTypeNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetParentTypeNil(b bool)`

 SetParentTypeNil sets the value for ParentType to be an explicit nil

### UnsetParentType
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetParentType()`

UnsetParentType ensures that no value is present for ParentType, not even an explicit nil
### GetParentId

`func (o *GetNetworkPool200ResponseNetworkPool) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GetNetworkPool200ResponseNetworkPool) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GetNetworkPool200ResponseNetworkPool) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetPoolGroup

`func (o *GetNetworkPool200ResponseNetworkPool) GetPoolGroup() string`

GetPoolGroup returns the PoolGroup field if non-nil, zero value otherwise.

### GetPoolGroupOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetPoolGroupOk() (*string, bool)`

GetPoolGroupOk returns a tuple with the PoolGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolGroup

`func (o *GetNetworkPool200ResponseNetworkPool) SetPoolGroup(v string)`

SetPoolGroup sets PoolGroup field to given value.

### HasPoolGroup

`func (o *GetNetworkPool200ResponseNetworkPool) HasPoolGroup() bool`

HasPoolGroup returns a boolean if a field has been set.

### SetPoolGroupNil

`func (o *GetNetworkPool200ResponseNetworkPool) SetPoolGroupNil(b bool)`

 SetPoolGroupNil sets the value for PoolGroup to be an explicit nil

### UnsetPoolGroup
`func (o *GetNetworkPool200ResponseNetworkPool) UnsetPoolGroup()`

UnsetPoolGroup ensures that no value is present for PoolGroup, not even an explicit nil
### GetIpRanges

`func (o *GetNetworkPool200ResponseNetworkPool) GetIpRanges() []GetNetworkPool200ResponseNetworkPoolIpRangesInner`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *GetNetworkPool200ResponseNetworkPool) GetIpRangesOk() (*[]GetNetworkPool200ResponseNetworkPoolIpRangesInner, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *GetNetworkPool200ResponseNetworkPool) SetIpRanges(v []GetNetworkPool200ResponseNetworkPoolIpRangesInner)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *GetNetworkPool200ResponseNetworkPool) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


