# GetLoadBalancer200ResponseLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Cloud** | Pointer to [**GetLoadBalancer200ResponseLoadBalancerCloud**](GetLoadBalancer200ResponseLoadBalancerCloud.md) |  | [optional] 
**Type** | Pointer to [**GetLoadBalancer200ResponseLoadBalancerType**](GetLoadBalancer200ResponseLoadBalancerType.md) |  | [optional] 
**Owner** | Pointer to [**GetLoadBalancer200ResponseLoadBalancerOwner**](GetLoadBalancer200ResponseLoadBalancerOwner.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **NullableString** |  | [optional] 
**ExternalIp** | Pointer to **NullableString** |  | [optional] 
**ApiPort** | Pointer to **NullableString** |  | [optional] 
**AdminPort** | Pointer to **NullableString** |  | [optional] 
**SslEnabled** | Pointer to **NullableBool** |  | [optional] 
**SslCert** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Credential** | Pointer to [**GetLoadBalancer200ResponseLoadBalancerCredential**](GetLoadBalancer200ResponseLoadBalancerCredential.md) |  | [optional] 
**Tenants** | Pointer to [**[]GetLoadBalancer200ResponseLoadBalancerTenantsInner**](GetLoadBalancer200ResponseLoadBalancerTenantsInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**GetLoadBalancer200ResponseLoadBalancerResourcePermission**](GetLoadBalancer200ResponseLoadBalancerResourcePermission.md) |  | [optional] 

## Methods

### NewGetLoadBalancer200ResponseLoadBalancer

`func NewGetLoadBalancer200ResponseLoadBalancer() *GetLoadBalancer200ResponseLoadBalancer`

NewGetLoadBalancer200ResponseLoadBalancer instantiates a new GetLoadBalancer200ResponseLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoadBalancer200ResponseLoadBalancerWithDefaults

`func NewGetLoadBalancer200ResponseLoadBalancerWithDefaults() *GetLoadBalancer200ResponseLoadBalancer`

NewGetLoadBalancer200ResponseLoadBalancerWithDefaults instantiates a new GetLoadBalancer200ResponseLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCloud

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetCloud() GetLoadBalancer200ResponseLoadBalancerCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetCloudOk() (*GetLoadBalancer200ResponseLoadBalancerCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetCloud(v GetLoadBalancer200ResponseLoadBalancerCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetType

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetType() GetLoadBalancer200ResponseLoadBalancerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetTypeOk() (*GetLoadBalancer200ResponseLoadBalancerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetType(v GetLoadBalancer200ResponseLoadBalancerType)`

SetType sets Type field to given value.

### HasType

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetOwner() GetLoadBalancer200ResponseLoadBalancerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetOwnerOk() (*GetLoadBalancer200ResponseLoadBalancerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetOwner(v GetLoadBalancer200ResponseLoadBalancerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetVisibility

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GetLoadBalancer200ResponseLoadBalancer) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetIp

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetInternalIp

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *GetLoadBalancer200ResponseLoadBalancer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetExternalIp

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *GetLoadBalancer200ResponseLoadBalancer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetApiPort

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetApiPort() string`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetApiPortOk() (*string, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetApiPort(v string)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *GetLoadBalancer200ResponseLoadBalancer) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetAdminPort() string`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetAdminPortOk() (*string, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetAdminPort(v string)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *GetLoadBalancer200ResponseLoadBalancer) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetSslEnabled

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### SetSslEnabledNil

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetSslEnabledNil(b bool)`

 SetSslEnabledNil sets the value for SslEnabled to be an explicit nil

### UnsetSslEnabled
`func (o *GetLoadBalancer200ResponseLoadBalancer) UnsetSslEnabled()`

UnsetSslEnabled ensures that no value is present for SslEnabled, not even an explicit nil
### GetSslCert

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *GetLoadBalancer200ResponseLoadBalancer) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetConfig

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCredential

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetCredential() GetLoadBalancer200ResponseLoadBalancerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetCredentialOk() (*GetLoadBalancer200ResponseLoadBalancerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetCredential(v GetLoadBalancer200ResponseLoadBalancerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetTenants() []GetLoadBalancer200ResponseLoadBalancerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetTenantsOk() (*[]GetLoadBalancer200ResponseLoadBalancerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetTenants(v []GetLoadBalancer200ResponseLoadBalancerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetResourcePermission() GetLoadBalancer200ResponseLoadBalancerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *GetLoadBalancer200ResponseLoadBalancer) GetResourcePermissionOk() (*GetLoadBalancer200ResponseLoadBalancerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *GetLoadBalancer200ResponseLoadBalancer) SetResourcePermission(v GetLoadBalancer200ResponseLoadBalancerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *GetLoadBalancer200ResponseLoadBalancer) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


