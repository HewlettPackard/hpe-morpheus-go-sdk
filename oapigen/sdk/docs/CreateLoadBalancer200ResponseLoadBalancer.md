# CreateLoadBalancer200ResponseLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Cloud** | Pointer to [**CreateLoadBalancer200ResponseLoadBalancerCloud**](CreateLoadBalancer200ResponseLoadBalancerCloud.md) |  | [optional] 
**Type** | Pointer to [**CreateLoadBalancer200ResponseLoadBalancerType**](CreateLoadBalancer200ResponseLoadBalancerType.md) |  | [optional] 
**Owner** | Pointer to [**CreateLoadBalancer200ResponseLoadBalancerOwner**](CreateLoadBalancer200ResponseLoadBalancerOwner.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**PasswordHash** | Pointer to **NullableString** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **NullableString** |  | [optional] 
**ExternalIp** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ApiPort** | Pointer to **NullableString** |  | [optional] 
**AdminPort** | Pointer to **NullableString** |  | [optional] 
**SslEnabled** | Pointer to **NullableBool** |  | [optional] 
**SslCert** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AllowVipEntry** | Pointer to **bool** |  | [optional] 
**VipPools** | Pointer to **[]map[string]interface{}** |  | [optional] 
**VirtualServiceName** | Pointer to **NullableString** |  | [optional] 
**PoolName** | Pointer to **NullableString** |  | [optional] 
**ServerName** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Credential** | Pointer to [**CreateLoadBalancer200ResponseLoadBalancerCredential**](CreateLoadBalancer200ResponseLoadBalancerCredential.md) |  | [optional] 
**Tenants** | Pointer to [**[]CreateLoadBalancer200ResponseLoadBalancerTenantsInner**](CreateLoadBalancer200ResponseLoadBalancerTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**CreateLoadBalancer200ResponseLoadBalancerResourcePermission**](CreateLoadBalancer200ResponseLoadBalancerResourcePermission.md) |  | [optional] 
**InstancePrice** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCreateLoadBalancer200ResponseLoadBalancer

`func NewCreateLoadBalancer200ResponseLoadBalancer() *CreateLoadBalancer200ResponseLoadBalancer`

NewCreateLoadBalancer200ResponseLoadBalancer instantiates a new CreateLoadBalancer200ResponseLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancer200ResponseLoadBalancerWithDefaults

`func NewCreateLoadBalancer200ResponseLoadBalancerWithDefaults() *CreateLoadBalancer200ResponseLoadBalancer`

NewCreateLoadBalancer200ResponseLoadBalancerWithDefaults instantiates a new CreateLoadBalancer200ResponseLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCloud

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetCloud() CreateLoadBalancer200ResponseLoadBalancerCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetCloudOk() (*CreateLoadBalancer200ResponseLoadBalancerCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetCloud(v CreateLoadBalancer200ResponseLoadBalancerCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetType

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetType() CreateLoadBalancer200ResponseLoadBalancerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetTypeOk() (*CreateLoadBalancer200ResponseLoadBalancerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetType(v CreateLoadBalancer200ResponseLoadBalancerType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetOwner() CreateLoadBalancer200ResponseLoadBalancerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetOwnerOk() (*CreateLoadBalancer200ResponseLoadBalancerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetOwner(v CreateLoadBalancer200ResponseLoadBalancerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetIp

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetExternalIp

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetExternalId

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetApiPort

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetApiPort() string`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetApiPortOk() (*string, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetApiPort(v string)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetAdminPort() string`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetAdminPortOk() (*string, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetAdminPort(v string)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetSslEnabled

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### SetSslEnabledNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetSslEnabledNil(b bool)`

 SetSslEnabledNil sets the value for SslEnabled to be an explicit nil

### UnsetSslEnabled
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetSslEnabled()`

UnsetSslEnabled ensures that no value is present for SslEnabled, not even an explicit nil
### GetSslCert

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetEnabled

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAllowVipEntry

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetAllowVipEntry() bool`

GetAllowVipEntry returns the AllowVipEntry field if non-nil, zero value otherwise.

### GetAllowVipEntryOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetAllowVipEntryOk() (*bool, bool)`

GetAllowVipEntryOk returns a tuple with the AllowVipEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVipEntry

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetAllowVipEntry(v bool)`

SetAllowVipEntry sets AllowVipEntry field to given value.

### HasAllowVipEntry

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasAllowVipEntry() bool`

HasAllowVipEntry returns a boolean if a field has been set.

### GetVipPools

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetVipPools() []map[string]interface{}`

GetVipPools returns the VipPools field if non-nil, zero value otherwise.

### GetVipPoolsOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetVipPoolsOk() (*[]map[string]interface{}, bool)`

GetVipPoolsOk returns a tuple with the VipPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPools

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetVipPools(v []map[string]interface{})`

SetVipPools sets VipPools field to given value.

### HasVipPools

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasVipPools() bool`

HasVipPools returns a boolean if a field has been set.

### GetVirtualServiceName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetVirtualServiceName() string`

GetVirtualServiceName returns the VirtualServiceName field if non-nil, zero value otherwise.

### GetVirtualServiceNameOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetVirtualServiceNameOk() (*string, bool)`

GetVirtualServiceNameOk returns a tuple with the VirtualServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServiceName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetVirtualServiceName(v string)`

SetVirtualServiceName sets VirtualServiceName field to given value.

### HasVirtualServiceName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasVirtualServiceName() bool`

HasVirtualServiceName returns a boolean if a field has been set.

### SetVirtualServiceNameNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetVirtualServiceNameNil(b bool)`

 SetVirtualServiceNameNil sets the value for VirtualServiceName to be an explicit nil

### UnsetVirtualServiceName
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetVirtualServiceName()`

UnsetVirtualServiceName ensures that no value is present for VirtualServiceName, not even an explicit nil
### GetPoolName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### SetPoolNameNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetPoolNameNil(b bool)`

 SetPoolNameNil sets the value for PoolName to be an explicit nil

### UnsetPoolName
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetPoolName()`

UnsetPoolName ensures that no value is present for PoolName, not even an explicit nil
### GetServerName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *CreateLoadBalancer200ResponseLoadBalancer) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetConfig

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCredential

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetCredential() CreateLoadBalancer200ResponseLoadBalancerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetCredentialOk() (*CreateLoadBalancer200ResponseLoadBalancerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetCredential(v CreateLoadBalancer200ResponseLoadBalancerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetTenants() []CreateLoadBalancer200ResponseLoadBalancerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetTenantsOk() (*[]CreateLoadBalancer200ResponseLoadBalancerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetTenants(v []CreateLoadBalancer200ResponseLoadBalancerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetResourcePermission() CreateLoadBalancer200ResponseLoadBalancerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetResourcePermissionOk() (*CreateLoadBalancer200ResponseLoadBalancerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetResourcePermission(v CreateLoadBalancer200ResponseLoadBalancerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetInstancePrice

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetInstancePrice() map[string]interface{}`

GetInstancePrice returns the InstancePrice field if non-nil, zero value otherwise.

### GetInstancePriceOk

`func (o *CreateLoadBalancer200ResponseLoadBalancer) GetInstancePriceOk() (*map[string]interface{}, bool)`

GetInstancePriceOk returns a tuple with the InstancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePrice

`func (o *CreateLoadBalancer200ResponseLoadBalancer) SetInstancePrice(v map[string]interface{})`

SetInstancePrice sets InstancePrice field to given value.

### HasInstancePrice

`func (o *CreateLoadBalancer200ResponseLoadBalancer) HasInstancePrice() bool`

HasInstancePrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


