# ListLoadBalancers200ResponseAllOfLoadBalancersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Cloud** | Pointer to [**NullableListLoadBalancers200ResponseAllOfLoadBalancersInnerCloud**](ListLoadBalancers200ResponseAllOfLoadBalancersInnerCloud.md) |  | [optional] 
**Type** | Pointer to [**ListLoadBalancers200ResponseAllOfLoadBalancersInnerType**](ListLoadBalancers200ResponseAllOfLoadBalancersInnerType.md) |  | [optional] 
**Owner** | Pointer to [**ListLoadBalancers200ResponseAllOfLoadBalancersInnerOwner**](ListLoadBalancers200ResponseAllOfLoadBalancersInnerOwner.md) |  | [optional] 
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
**Credential** | Pointer to [**ListLoadBalancers200ResponseAllOfLoadBalancersInnerCredential**](ListLoadBalancers200ResponseAllOfLoadBalancersInnerCredential.md) |  | [optional] 
**Tenants** | Pointer to [**[]ListLoadBalancers200ResponseAllOfLoadBalancersInnerTenantsInner**](ListLoadBalancers200ResponseAllOfLoadBalancersInnerTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission**](ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission.md) |  | [optional] 
**InstancePrice** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListLoadBalancers200ResponseAllOfLoadBalancersInner

`func NewListLoadBalancers200ResponseAllOfLoadBalancersInner() *ListLoadBalancers200ResponseAllOfLoadBalancersInner`

NewListLoadBalancers200ResponseAllOfLoadBalancersInner instantiates a new ListLoadBalancers200ResponseAllOfLoadBalancersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancers200ResponseAllOfLoadBalancersInnerWithDefaults

`func NewListLoadBalancers200ResponseAllOfLoadBalancersInnerWithDefaults() *ListLoadBalancers200ResponseAllOfLoadBalancersInner`

NewListLoadBalancers200ResponseAllOfLoadBalancersInnerWithDefaults instantiates a new ListLoadBalancers200ResponseAllOfLoadBalancersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCloud

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetCloud() ListLoadBalancers200ResponseAllOfLoadBalancersInnerCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetCloudOk() (*ListLoadBalancers200ResponseAllOfLoadBalancersInnerCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetCloud(v ListLoadBalancers200ResponseAllOfLoadBalancersInnerCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### SetCloudNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetCloudNil(b bool)`

 SetCloudNil sets the value for Cloud to be an explicit nil

### UnsetCloud
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetCloud()`

UnsetCloud ensures that no value is present for Cloud, not even an explicit nil
### GetType

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetType() ListLoadBalancers200ResponseAllOfLoadBalancersInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetTypeOk() (*ListLoadBalancers200ResponseAllOfLoadBalancersInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetType(v ListLoadBalancers200ResponseAllOfLoadBalancersInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetOwner() ListLoadBalancers200ResponseAllOfLoadBalancersInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetOwnerOk() (*ListLoadBalancers200ResponseAllOfLoadBalancersInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetOwner(v ListLoadBalancers200ResponseAllOfLoadBalancersInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetVisibility

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetIp

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetExternalIp

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetExternalId

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetApiPort

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetApiPort() string`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetApiPortOk() (*string, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetApiPort(v string)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetAdminPort() string`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetAdminPortOk() (*string, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetAdminPort(v string)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetSslEnabled

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### SetSslEnabledNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetSslEnabledNil(b bool)`

 SetSslEnabledNil sets the value for SslEnabled to be an explicit nil

### UnsetSslEnabled
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetSslEnabled()`

UnsetSslEnabled ensures that no value is present for SslEnabled, not even an explicit nil
### GetSslCert

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetEnabled

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAllowVipEntry

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetAllowVipEntry() bool`

GetAllowVipEntry returns the AllowVipEntry field if non-nil, zero value otherwise.

### GetAllowVipEntryOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetAllowVipEntryOk() (*bool, bool)`

GetAllowVipEntryOk returns a tuple with the AllowVipEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVipEntry

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetAllowVipEntry(v bool)`

SetAllowVipEntry sets AllowVipEntry field to given value.

### HasAllowVipEntry

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasAllowVipEntry() bool`

HasAllowVipEntry returns a boolean if a field has been set.

### GetVipPools

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetVipPools() []map[string]interface{}`

GetVipPools returns the VipPools field if non-nil, zero value otherwise.

### GetVipPoolsOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetVipPoolsOk() (*[]map[string]interface{}, bool)`

GetVipPoolsOk returns a tuple with the VipPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPools

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetVipPools(v []map[string]interface{})`

SetVipPools sets VipPools field to given value.

### HasVipPools

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasVipPools() bool`

HasVipPools returns a boolean if a field has been set.

### GetVirtualServiceName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetVirtualServiceName() string`

GetVirtualServiceName returns the VirtualServiceName field if non-nil, zero value otherwise.

### GetVirtualServiceNameOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetVirtualServiceNameOk() (*string, bool)`

GetVirtualServiceNameOk returns a tuple with the VirtualServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServiceName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetVirtualServiceName(v string)`

SetVirtualServiceName sets VirtualServiceName field to given value.

### HasVirtualServiceName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasVirtualServiceName() bool`

HasVirtualServiceName returns a boolean if a field has been set.

### SetVirtualServiceNameNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetVirtualServiceNameNil(b bool)`

 SetVirtualServiceNameNil sets the value for VirtualServiceName to be an explicit nil

### UnsetVirtualServiceName
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetVirtualServiceName()`

UnsetVirtualServiceName ensures that no value is present for VirtualServiceName, not even an explicit nil
### GetPoolName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### SetPoolNameNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetPoolNameNil(b bool)`

 SetPoolNameNil sets the value for PoolName to be an explicit nil

### UnsetPoolName
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetPoolName()`

UnsetPoolName ensures that no value is present for PoolName, not even an explicit nil
### GetServerName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetConfig

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCredential

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetCredential() ListLoadBalancers200ResponseAllOfLoadBalancersInnerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetCredentialOk() (*ListLoadBalancers200ResponseAllOfLoadBalancersInnerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetCredential(v ListLoadBalancers200ResponseAllOfLoadBalancersInnerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetTenants() []ListLoadBalancers200ResponseAllOfLoadBalancersInnerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetTenantsOk() (*[]ListLoadBalancers200ResponseAllOfLoadBalancersInnerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetTenants(v []ListLoadBalancers200ResponseAllOfLoadBalancersInnerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetResourcePermission() ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetResourcePermissionOk() (*ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetResourcePermission(v ListLoadBalancers200ResponseAllOfLoadBalancersInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetInstancePrice

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetInstancePrice() map[string]interface{}`

GetInstancePrice returns the InstancePrice field if non-nil, zero value otherwise.

### GetInstancePriceOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetInstancePriceOk() (*map[string]interface{}, bool)`

GetInstancePriceOk returns a tuple with the InstancePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancePrice

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetInstancePrice(v map[string]interface{})`

SetInstancePrice sets InstancePrice field to given value.

### HasInstancePrice

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasInstancePrice() bool`

HasInstancePrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


