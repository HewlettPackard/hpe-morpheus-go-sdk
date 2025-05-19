# ListLoadBalancers200ResponseAllOfLoadBalancersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Cloud** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**ApiPort** | Pointer to **string** |  | [optional] 
**AdminPort** | Pointer to **string** |  | [optional] 
**SslEnabled** | Pointer to **bool** |  | [optional] 
**SslCert** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Credential** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential.md) |  | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**ResourcePermission** | Pointer to [**ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission**](ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission.md) |  | [optional] 

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

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetCloud() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetCloudOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetCloud(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetType

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

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

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetCredential() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetCredentialOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetCredential(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetResourcePermission

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetResourcePermission() ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) GetResourcePermissionOk() (*ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) SetResourcePermission(v ListCloudDatastores200ResponseAllOfDatastoresInnerResourcePermission)`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *ListLoadBalancers200ResponseAllOfLoadBalancersInner) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


