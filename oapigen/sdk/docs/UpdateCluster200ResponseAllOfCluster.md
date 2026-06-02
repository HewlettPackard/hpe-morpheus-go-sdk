# UpdateCluster200ResponseAllOfCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ServiceUrl** | Pointer to **NullableString** |  | [optional] 
**ServiceHost** | Pointer to **NullableString** |  | [optional] 
**ServicePath** | Pointer to **NullableString** |  | [optional] 
**ServiceHostname** | Pointer to **NullableString** |  | [optional] 
**ServicePort** | Pointer to **int64** |  | [optional] 
**ServiceUsername** | Pointer to **NullableString** |  | [optional] 
**ServicePassword** | Pointer to **NullableString** |  | [optional] 
**ServicePasswordHash** | Pointer to **NullableString** |  | [optional] 
**ServiceToken** | Pointer to **NullableString** |  | [optional] 
**ServiceTokenHash** | Pointer to **NullableString** |  | [optional] 
**ServiceAccess** | Pointer to **NullableString** |  | [optional] 
**ServiceAccessHash** | Pointer to **NullableString** |  | [optional] 
**ServiceCert** | Pointer to **NullableString** |  | [optional] 
**ServiceCertHash** | Pointer to **NullableString** |  | [optional] 
**ServiceVersion** | Pointer to **NullableString** |  | [optional] 
**SearchDomains** | Pointer to **NullableString** |  | [optional] 
**EnableInternalDns** | Pointer to **bool** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**DatacenterId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**InventoryLevel** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **NullableTime** |  | [optional] 
**NextRunDate** | Pointer to **NullableTime** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableInt64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**AutoRecoverPowerState** | Pointer to **bool** | Automatically Power on VMs | [optional] [default to false]
**CpuPlacementMode** | Pointer to **NullableString** | Cluster CPU placement mode | [optional] 
**UseAgent** | Pointer to **NullableString** | Use the Agent to relay communications for the Kubernetes API instead of direct. | [optional] 
**ProvisionComplete** | Pointer to **bool** | Changes from false to true once provisioning is finished. | [optional] 
**ServiceEntry** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**UpdateCluster200ResponseAllOfClusterCreatedBy**](UpdateCluster200ResponseAllOfClusterCreatedBy.md) |  | [optional] 
**UserGroup** | Pointer to **NullableString** |  | [optional] 
**Layout** | Pointer to [**UpdateCluster200ResponseAllOfClusterLayout**](UpdateCluster200ResponseAllOfClusterLayout.md) |  | [optional] 
**Owner** | Pointer to [**UpdateCluster200ResponseAllOfClusterOwner**](UpdateCluster200ResponseAllOfClusterOwner.md) |  | [optional] 
**Servers** | Pointer to [**[]UpdateCluster200ResponseAllOfClusterServersInner**](UpdateCluster200ResponseAllOfClusterServersInner.md) |  | [optional] 
**Accounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Integrations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Site** | Pointer to [**UpdateCluster200ResponseAllOfClusterSite**](UpdateCluster200ResponseAllOfClusterSite.md) |  | [optional] 
**Type** | Pointer to [**UpdateCluster200ResponseAllOfClusterType**](UpdateCluster200ResponseAllOfClusterType.md) |  | [optional] 
**Zone** | Pointer to [**UpdateCluster200ResponseAllOfClusterZone**](UpdateCluster200ResponseAllOfClusterZone.md) |  | [optional] 
**WorkerStats** | Pointer to [**UpdateCluster200ResponseAllOfClusterWorkerStats**](UpdateCluster200ResponseAllOfClusterWorkerStats.md) |  | [optional] 
**ContainersCount** | Pointer to **int64** |  | [optional] 
**DeploymentsCount** | Pointer to **int64** |  | [optional] 
**PodsCount** | Pointer to **int64** |  | [optional] 
**JobsCount** | Pointer to **int64** |  | [optional] 
**VolumesCount** | Pointer to **int64** |  | [optional] 
**NamespacesCount** | Pointer to **int64** |  | [optional] 
**WorkersCount** | Pointer to **int64** |  | [optional] 
**ServicesCount** | Pointer to **int64** |  | [optional] 
**Permissions** | Pointer to [**UpdateCluster200ResponseAllOfClusterPermissions**](UpdateCluster200ResponseAllOfClusterPermissions.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateCluster200ResponseAllOfCluster

`func NewUpdateCluster200ResponseAllOfCluster() *UpdateCluster200ResponseAllOfCluster`

NewUpdateCluster200ResponseAllOfCluster instantiates a new UpdateCluster200ResponseAllOfCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateCluster200ResponseAllOfCluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCluster200ResponseAllOfCluster) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCluster200ResponseAllOfCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *UpdateCluster200ResponseAllOfCluster) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UpdateCluster200ResponseAllOfCluster) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UpdateCluster200ResponseAllOfCluster) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *UpdateCluster200ResponseAllOfCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCluster200ResponseAllOfCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCluster200ResponseAllOfCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateCluster200ResponseAllOfCluster) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateCluster200ResponseAllOfCluster) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateCluster200ResponseAllOfCluster) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *UpdateCluster200ResponseAllOfCluster) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateCluster200ResponseAllOfCluster) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateCluster200ResponseAllOfCluster) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVisibility

`func (o *UpdateCluster200ResponseAllOfCluster) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCluster200ResponseAllOfCluster) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCluster200ResponseAllOfCluster) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCluster200ResponseAllOfCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCluster200ResponseAllOfCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCluster200ResponseAllOfCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocation

`func (o *UpdateCluster200ResponseAllOfCluster) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateCluster200ResponseAllOfCluster) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateCluster200ResponseAllOfCluster) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetEnabled

`func (o *UpdateCluster200ResponseAllOfCluster) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCluster200ResponseAllOfCluster) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCluster200ResponseAllOfCluster) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePath

`func (o *UpdateCluster200ResponseAllOfCluster) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *UpdateCluster200ResponseAllOfCluster) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *UpdateCluster200ResponseAllOfCluster) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### SetServicePathNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServicePathNil(b bool)`

 SetServicePathNil sets the value for ServicePath to be an explicit nil

### UnsetServicePath
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServicePath()`

UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
### GetServiceHostname

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceHostname() string`

GetServiceHostname returns the ServiceHostname field if non-nil, zero value otherwise.

### GetServiceHostnameOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceHostnameOk() (*string, bool)`

GetServiceHostnameOk returns a tuple with the ServiceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHostname

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceHostname(v string)`

SetServiceHostname sets ServiceHostname field to given value.

### HasServiceHostname

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceHostname() bool`

HasServiceHostname returns a boolean if a field has been set.

### SetServiceHostnameNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceHostnameNil(b bool)`

 SetServiceHostnameNil sets the value for ServiceHostname to be an explicit nil

### UnsetServiceHostname
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceHostname()`

UnsetServiceHostname ensures that no value is present for ServiceHostname, not even an explicit nil
### GetServicePort

`func (o *UpdateCluster200ResponseAllOfCluster) GetServicePort() int64`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServicePortOk() (*int64, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *UpdateCluster200ResponseAllOfCluster) SetServicePort(v int64)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *UpdateCluster200ResponseAllOfCluster) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceUsername

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *UpdateCluster200ResponseAllOfCluster) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *UpdateCluster200ResponseAllOfCluster) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *UpdateCluster200ResponseAllOfCluster) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *UpdateCluster200ResponseAllOfCluster) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *UpdateCluster200ResponseAllOfCluster) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *UpdateCluster200ResponseAllOfCluster) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### SetServicePasswordHashNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServicePasswordHashNil(b bool)`

 SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil

### UnsetServicePasswordHash
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServicePasswordHash()`

UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
### GetServiceToken

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetServiceTokenHash

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### SetServiceTokenHashNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceTokenHashNil(b bool)`

 SetServiceTokenHashNil sets the value for ServiceTokenHash to be an explicit nil

### UnsetServiceTokenHash
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceTokenHash()`

UnsetServiceTokenHash ensures that no value is present for ServiceTokenHash, not even an explicit nil
### GetServiceAccess

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceAccess() string`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceAccessOk() (*string, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceAccess(v string)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.

### SetServiceAccessNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceAccessNil(b bool)`

 SetServiceAccessNil sets the value for ServiceAccess to be an explicit nil

### UnsetServiceAccess
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceAccess()`

UnsetServiceAccess ensures that no value is present for ServiceAccess, not even an explicit nil
### GetServiceAccessHash

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceAccessHash() string`

GetServiceAccessHash returns the ServiceAccessHash field if non-nil, zero value otherwise.

### GetServiceAccessHashOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceAccessHashOk() (*string, bool)`

GetServiceAccessHashOk returns a tuple with the ServiceAccessHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccessHash

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceAccessHash(v string)`

SetServiceAccessHash sets ServiceAccessHash field to given value.

### HasServiceAccessHash

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceAccessHash() bool`

HasServiceAccessHash returns a boolean if a field has been set.

### SetServiceAccessHashNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceAccessHashNil(b bool)`

 SetServiceAccessHashNil sets the value for ServiceAccessHash to be an explicit nil

### UnsetServiceAccessHash
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceAccessHash()`

UnsetServiceAccessHash ensures that no value is present for ServiceAccessHash, not even an explicit nil
### GetServiceCert

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceCert() string`

GetServiceCert returns the ServiceCert field if non-nil, zero value otherwise.

### GetServiceCertOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceCertOk() (*string, bool)`

GetServiceCertOk returns a tuple with the ServiceCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCert

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceCert(v string)`

SetServiceCert sets ServiceCert field to given value.

### HasServiceCert

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceCert() bool`

HasServiceCert returns a boolean if a field has been set.

### SetServiceCertNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceCertNil(b bool)`

 SetServiceCertNil sets the value for ServiceCert to be an explicit nil

### UnsetServiceCert
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceCert()`

UnsetServiceCert ensures that no value is present for ServiceCert, not even an explicit nil
### GetServiceCertHash

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceCertHash() string`

GetServiceCertHash returns the ServiceCertHash field if non-nil, zero value otherwise.

### GetServiceCertHashOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceCertHashOk() (*string, bool)`

GetServiceCertHashOk returns a tuple with the ServiceCertHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCertHash

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceCertHash(v string)`

SetServiceCertHash sets ServiceCertHash field to given value.

### HasServiceCertHash

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceCertHash() bool`

HasServiceCertHash returns a boolean if a field has been set.

### SetServiceCertHashNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceCertHashNil(b bool)`

 SetServiceCertHashNil sets the value for ServiceCertHash to be an explicit nil

### UnsetServiceCertHash
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceCertHash()`

UnsetServiceCertHash ensures that no value is present for ServiceCertHash, not even an explicit nil
### GetServiceVersion

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### SetServiceVersionNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceVersionNil(b bool)`

 SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil

### UnsetServiceVersion
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceVersion()`

UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
### GetSearchDomains

`func (o *UpdateCluster200ResponseAllOfCluster) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *UpdateCluster200ResponseAllOfCluster) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *UpdateCluster200ResponseAllOfCluster) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### SetSearchDomainsNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetSearchDomainsNil(b bool)`

 SetSearchDomainsNil sets the value for SearchDomains to be an explicit nil

### UnsetSearchDomains
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetSearchDomains()`

UnsetSearchDomains ensures that no value is present for SearchDomains, not even an explicit nil
### GetEnableInternalDns

`func (o *UpdateCluster200ResponseAllOfCluster) GetEnableInternalDns() bool`

GetEnableInternalDns returns the EnableInternalDns field if non-nil, zero value otherwise.

### GetEnableInternalDnsOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetEnableInternalDnsOk() (*bool, bool)`

GetEnableInternalDnsOk returns a tuple with the EnableInternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternalDns

`func (o *UpdateCluster200ResponseAllOfCluster) SetEnableInternalDns(v bool)`

SetEnableInternalDns sets EnableInternalDns field to given value.

### HasEnableInternalDns

`func (o *UpdateCluster200ResponseAllOfCluster) HasEnableInternalDns() bool`

HasEnableInternalDns returns a boolean if a field has been set.

### GetInternalId

`func (o *UpdateCluster200ResponseAllOfCluster) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateCluster200ResponseAllOfCluster) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateCluster200ResponseAllOfCluster) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *UpdateCluster200ResponseAllOfCluster) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateCluster200ResponseAllOfCluster) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateCluster200ResponseAllOfCluster) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDatacenterId

`func (o *UpdateCluster200ResponseAllOfCluster) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *UpdateCluster200ResponseAllOfCluster) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *UpdateCluster200ResponseAllOfCluster) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### SetDatacenterIdNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetDatacenterIdNil(b bool)`

 SetDatacenterIdNil sets the value for DatacenterId to be an explicit nil

### UnsetDatacenterId
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetDatacenterId()`

UnsetDatacenterId ensures that no value is present for DatacenterId, not even an explicit nil
### GetStatus

`func (o *UpdateCluster200ResponseAllOfCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateCluster200ResponseAllOfCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateCluster200ResponseAllOfCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *UpdateCluster200ResponseAllOfCluster) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *UpdateCluster200ResponseAllOfCluster) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *UpdateCluster200ResponseAllOfCluster) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusMessage

`func (o *UpdateCluster200ResponseAllOfCluster) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateCluster200ResponseAllOfCluster) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateCluster200ResponseAllOfCluster) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetInventoryLevel

`func (o *UpdateCluster200ResponseAllOfCluster) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *UpdateCluster200ResponseAllOfCluster) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *UpdateCluster200ResponseAllOfCluster) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetLastSync

`func (o *UpdateCluster200ResponseAllOfCluster) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *UpdateCluster200ResponseAllOfCluster) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *UpdateCluster200ResponseAllOfCluster) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetNextRunDate

`func (o *UpdateCluster200ResponseAllOfCluster) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *UpdateCluster200ResponseAllOfCluster) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *UpdateCluster200ResponseAllOfCluster) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetLastSyncDuration

`func (o *UpdateCluster200ResponseAllOfCluster) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *UpdateCluster200ResponseAllOfCluster) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *UpdateCluster200ResponseAllOfCluster) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetDateCreated

`func (o *UpdateCluster200ResponseAllOfCluster) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateCluster200ResponseAllOfCluster) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateCluster200ResponseAllOfCluster) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateCluster200ResponseAllOfCluster) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateCluster200ResponseAllOfCluster) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateCluster200ResponseAllOfCluster) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetManaged

`func (o *UpdateCluster200ResponseAllOfCluster) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *UpdateCluster200ResponseAllOfCluster) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *UpdateCluster200ResponseAllOfCluster) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateCluster200ResponseAllOfCluster) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateCluster200ResponseAllOfCluster) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateCluster200ResponseAllOfCluster) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetAutoRecoverPowerState

`func (o *UpdateCluster200ResponseAllOfCluster) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *UpdateCluster200ResponseAllOfCluster) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *UpdateCluster200ResponseAllOfCluster) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetCpuPlacementMode

`func (o *UpdateCluster200ResponseAllOfCluster) GetCpuPlacementMode() string`

GetCpuPlacementMode returns the CpuPlacementMode field if non-nil, zero value otherwise.

### GetCpuPlacementModeOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetCpuPlacementModeOk() (*string, bool)`

GetCpuPlacementModeOk returns a tuple with the CpuPlacementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuPlacementMode

`func (o *UpdateCluster200ResponseAllOfCluster) SetCpuPlacementMode(v string)`

SetCpuPlacementMode sets CpuPlacementMode field to given value.

### HasCpuPlacementMode

`func (o *UpdateCluster200ResponseAllOfCluster) HasCpuPlacementMode() bool`

HasCpuPlacementMode returns a boolean if a field has been set.

### SetCpuPlacementModeNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetCpuPlacementModeNil(b bool)`

 SetCpuPlacementModeNil sets the value for CpuPlacementMode to be an explicit nil

### UnsetCpuPlacementMode
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetCpuPlacementMode()`

UnsetCpuPlacementMode ensures that no value is present for CpuPlacementMode, not even an explicit nil
### GetUseAgent

`func (o *UpdateCluster200ResponseAllOfCluster) GetUseAgent() string`

GetUseAgent returns the UseAgent field if non-nil, zero value otherwise.

### GetUseAgentOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetUseAgentOk() (*string, bool)`

GetUseAgentOk returns a tuple with the UseAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgent

`func (o *UpdateCluster200ResponseAllOfCluster) SetUseAgent(v string)`

SetUseAgent sets UseAgent field to given value.

### HasUseAgent

`func (o *UpdateCluster200ResponseAllOfCluster) HasUseAgent() bool`

HasUseAgent returns a boolean if a field has been set.

### SetUseAgentNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetUseAgentNil(b bool)`

 SetUseAgentNil sets the value for UseAgent to be an explicit nil

### UnsetUseAgent
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetUseAgent()`

UnsetUseAgent ensures that no value is present for UseAgent, not even an explicit nil
### GetProvisionComplete

`func (o *UpdateCluster200ResponseAllOfCluster) GetProvisionComplete() bool`

GetProvisionComplete returns the ProvisionComplete field if non-nil, zero value otherwise.

### GetProvisionCompleteOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetProvisionCompleteOk() (*bool, bool)`

GetProvisionCompleteOk returns a tuple with the ProvisionComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionComplete

`func (o *UpdateCluster200ResponseAllOfCluster) SetProvisionComplete(v bool)`

SetProvisionComplete sets ProvisionComplete field to given value.

### HasProvisionComplete

`func (o *UpdateCluster200ResponseAllOfCluster) HasProvisionComplete() bool`

HasProvisionComplete returns a boolean if a field has been set.

### GetServiceEntry

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceEntry() string`

GetServiceEntry returns the ServiceEntry field if non-nil, zero value otherwise.

### GetServiceEntryOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServiceEntryOk() (*string, bool)`

GetServiceEntryOk returns a tuple with the ServiceEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEntry

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceEntry(v string)`

SetServiceEntry sets ServiceEntry field to given value.

### HasServiceEntry

`func (o *UpdateCluster200ResponseAllOfCluster) HasServiceEntry() bool`

HasServiceEntry returns a boolean if a field has been set.

### SetServiceEntryNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetServiceEntryNil(b bool)`

 SetServiceEntryNil sets the value for ServiceEntry to be an explicit nil

### UnsetServiceEntry
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetServiceEntry()`

UnsetServiceEntry ensures that no value is present for ServiceEntry, not even an explicit nil
### GetCreatedBy

`func (o *UpdateCluster200ResponseAllOfCluster) GetCreatedBy() UpdateCluster200ResponseAllOfClusterCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetCreatedByOk() (*UpdateCluster200ResponseAllOfClusterCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateCluster200ResponseAllOfCluster) SetCreatedBy(v UpdateCluster200ResponseAllOfClusterCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateCluster200ResponseAllOfCluster) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUserGroup

`func (o *UpdateCluster200ResponseAllOfCluster) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *UpdateCluster200ResponseAllOfCluster) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *UpdateCluster200ResponseAllOfCluster) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### SetUserGroupNil

`func (o *UpdateCluster200ResponseAllOfCluster) SetUserGroupNil(b bool)`

 SetUserGroupNil sets the value for UserGroup to be an explicit nil

### UnsetUserGroup
`func (o *UpdateCluster200ResponseAllOfCluster) UnsetUserGroup()`

UnsetUserGroup ensures that no value is present for UserGroup, not even an explicit nil
### GetLayout

`func (o *UpdateCluster200ResponseAllOfCluster) GetLayout() UpdateCluster200ResponseAllOfClusterLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetLayoutOk() (*UpdateCluster200ResponseAllOfClusterLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *UpdateCluster200ResponseAllOfCluster) SetLayout(v UpdateCluster200ResponseAllOfClusterLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *UpdateCluster200ResponseAllOfCluster) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateCluster200ResponseAllOfCluster) GetOwner() UpdateCluster200ResponseAllOfClusterOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetOwnerOk() (*UpdateCluster200ResponseAllOfClusterOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateCluster200ResponseAllOfCluster) SetOwner(v UpdateCluster200ResponseAllOfClusterOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateCluster200ResponseAllOfCluster) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetServers

`func (o *UpdateCluster200ResponseAllOfCluster) GetServers() []UpdateCluster200ResponseAllOfClusterServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServersOk() (*[]UpdateCluster200ResponseAllOfClusterServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *UpdateCluster200ResponseAllOfCluster) SetServers(v []UpdateCluster200ResponseAllOfClusterServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *UpdateCluster200ResponseAllOfCluster) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetAccounts

`func (o *UpdateCluster200ResponseAllOfCluster) GetAccounts() []map[string]interface{}`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetAccountsOk() (*[]map[string]interface{}, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdateCluster200ResponseAllOfCluster) SetAccounts(v []map[string]interface{})`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UpdateCluster200ResponseAllOfCluster) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetIntegrations

`func (o *UpdateCluster200ResponseAllOfCluster) GetIntegrations() []map[string]interface{}`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetIntegrationsOk() (*[]map[string]interface{}, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *UpdateCluster200ResponseAllOfCluster) SetIntegrations(v []map[string]interface{})`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *UpdateCluster200ResponseAllOfCluster) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetSite

`func (o *UpdateCluster200ResponseAllOfCluster) GetSite() UpdateCluster200ResponseAllOfClusterSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetSiteOk() (*UpdateCluster200ResponseAllOfClusterSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *UpdateCluster200ResponseAllOfCluster) SetSite(v UpdateCluster200ResponseAllOfClusterSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *UpdateCluster200ResponseAllOfCluster) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetType

`func (o *UpdateCluster200ResponseAllOfCluster) GetType() UpdateCluster200ResponseAllOfClusterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetTypeOk() (*UpdateCluster200ResponseAllOfClusterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCluster200ResponseAllOfCluster) SetType(v UpdateCluster200ResponseAllOfClusterType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCluster200ResponseAllOfCluster) HasType() bool`

HasType returns a boolean if a field has been set.

### GetZone

`func (o *UpdateCluster200ResponseAllOfCluster) GetZone() UpdateCluster200ResponseAllOfClusterZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetZoneOk() (*UpdateCluster200ResponseAllOfClusterZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdateCluster200ResponseAllOfCluster) SetZone(v UpdateCluster200ResponseAllOfClusterZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdateCluster200ResponseAllOfCluster) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetWorkerStats

`func (o *UpdateCluster200ResponseAllOfCluster) GetWorkerStats() UpdateCluster200ResponseAllOfClusterWorkerStats`

GetWorkerStats returns the WorkerStats field if non-nil, zero value otherwise.

### GetWorkerStatsOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetWorkerStatsOk() (*UpdateCluster200ResponseAllOfClusterWorkerStats, bool)`

GetWorkerStatsOk returns a tuple with the WorkerStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerStats

`func (o *UpdateCluster200ResponseAllOfCluster) SetWorkerStats(v UpdateCluster200ResponseAllOfClusterWorkerStats)`

SetWorkerStats sets WorkerStats field to given value.

### HasWorkerStats

`func (o *UpdateCluster200ResponseAllOfCluster) HasWorkerStats() bool`

HasWorkerStats returns a boolean if a field has been set.

### GetContainersCount

`func (o *UpdateCluster200ResponseAllOfCluster) GetContainersCount() int64`

GetContainersCount returns the ContainersCount field if non-nil, zero value otherwise.

### GetContainersCountOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetContainersCountOk() (*int64, bool)`

GetContainersCountOk returns a tuple with the ContainersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainersCount

`func (o *UpdateCluster200ResponseAllOfCluster) SetContainersCount(v int64)`

SetContainersCount sets ContainersCount field to given value.

### HasContainersCount

`func (o *UpdateCluster200ResponseAllOfCluster) HasContainersCount() bool`

HasContainersCount returns a boolean if a field has been set.

### GetDeploymentsCount

`func (o *UpdateCluster200ResponseAllOfCluster) GetDeploymentsCount() int64`

GetDeploymentsCount returns the DeploymentsCount field if non-nil, zero value otherwise.

### GetDeploymentsCountOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetDeploymentsCountOk() (*int64, bool)`

GetDeploymentsCountOk returns a tuple with the DeploymentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsCount

`func (o *UpdateCluster200ResponseAllOfCluster) SetDeploymentsCount(v int64)`

SetDeploymentsCount sets DeploymentsCount field to given value.

### HasDeploymentsCount

`func (o *UpdateCluster200ResponseAllOfCluster) HasDeploymentsCount() bool`

HasDeploymentsCount returns a boolean if a field has been set.

### GetPodsCount

`func (o *UpdateCluster200ResponseAllOfCluster) GetPodsCount() int64`

GetPodsCount returns the PodsCount field if non-nil, zero value otherwise.

### GetPodsCountOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetPodsCountOk() (*int64, bool)`

GetPodsCountOk returns a tuple with the PodsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsCount

`func (o *UpdateCluster200ResponseAllOfCluster) SetPodsCount(v int64)`

SetPodsCount sets PodsCount field to given value.

### HasPodsCount

`func (o *UpdateCluster200ResponseAllOfCluster) HasPodsCount() bool`

HasPodsCount returns a boolean if a field has been set.

### GetJobsCount

`func (o *UpdateCluster200ResponseAllOfCluster) GetJobsCount() int64`

GetJobsCount returns the JobsCount field if non-nil, zero value otherwise.

### GetJobsCountOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetJobsCountOk() (*int64, bool)`

GetJobsCountOk returns a tuple with the JobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCount

`func (o *UpdateCluster200ResponseAllOfCluster) SetJobsCount(v int64)`

SetJobsCount sets JobsCount field to given value.

### HasJobsCount

`func (o *UpdateCluster200ResponseAllOfCluster) HasJobsCount() bool`

HasJobsCount returns a boolean if a field has been set.

### GetVolumesCount

`func (o *UpdateCluster200ResponseAllOfCluster) GetVolumesCount() int64`

GetVolumesCount returns the VolumesCount field if non-nil, zero value otherwise.

### GetVolumesCountOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetVolumesCountOk() (*int64, bool)`

GetVolumesCountOk returns a tuple with the VolumesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesCount

`func (o *UpdateCluster200ResponseAllOfCluster) SetVolumesCount(v int64)`

SetVolumesCount sets VolumesCount field to given value.

### HasVolumesCount

`func (o *UpdateCluster200ResponseAllOfCluster) HasVolumesCount() bool`

HasVolumesCount returns a boolean if a field has been set.

### GetNamespacesCount

`func (o *UpdateCluster200ResponseAllOfCluster) GetNamespacesCount() int64`

GetNamespacesCount returns the NamespacesCount field if non-nil, zero value otherwise.

### GetNamespacesCountOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetNamespacesCountOk() (*int64, bool)`

GetNamespacesCountOk returns a tuple with the NamespacesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespacesCount

`func (o *UpdateCluster200ResponseAllOfCluster) SetNamespacesCount(v int64)`

SetNamespacesCount sets NamespacesCount field to given value.

### HasNamespacesCount

`func (o *UpdateCluster200ResponseAllOfCluster) HasNamespacesCount() bool`

HasNamespacesCount returns a boolean if a field has been set.

### GetWorkersCount

`func (o *UpdateCluster200ResponseAllOfCluster) GetWorkersCount() int64`

GetWorkersCount returns the WorkersCount field if non-nil, zero value otherwise.

### GetWorkersCountOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetWorkersCountOk() (*int64, bool)`

GetWorkersCountOk returns a tuple with the WorkersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkersCount

`func (o *UpdateCluster200ResponseAllOfCluster) SetWorkersCount(v int64)`

SetWorkersCount sets WorkersCount field to given value.

### HasWorkersCount

`func (o *UpdateCluster200ResponseAllOfCluster) HasWorkersCount() bool`

HasWorkersCount returns a boolean if a field has been set.

### GetServicesCount

`func (o *UpdateCluster200ResponseAllOfCluster) GetServicesCount() int64`

GetServicesCount returns the ServicesCount field if non-nil, zero value otherwise.

### GetServicesCountOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetServicesCountOk() (*int64, bool)`

GetServicesCountOk returns a tuple with the ServicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesCount

`func (o *UpdateCluster200ResponseAllOfCluster) SetServicesCount(v int64)`

SetServicesCount sets ServicesCount field to given value.

### HasServicesCount

`func (o *UpdateCluster200ResponseAllOfCluster) HasServicesCount() bool`

HasServicesCount returns a boolean if a field has been set.

### GetPermissions

`func (o *UpdateCluster200ResponseAllOfCluster) GetPermissions() UpdateCluster200ResponseAllOfClusterPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetPermissionsOk() (*UpdateCluster200ResponseAllOfClusterPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateCluster200ResponseAllOfCluster) SetPermissions(v UpdateCluster200ResponseAllOfClusterPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateCluster200ResponseAllOfCluster) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateCluster200ResponseAllOfCluster) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateCluster200ResponseAllOfCluster) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateCluster200ResponseAllOfCluster) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateCluster200ResponseAllOfCluster) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


