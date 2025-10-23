# GetCluster200ResponseCluster

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
**UseAgent** | Pointer to **NullableString** | Use the Agent to relay communications for the Kubernetes API instead of direct. | [optional] 
**ProvisionComplete** | Pointer to **bool** | Changes from false to true once provisioning is finished. | [optional] 
**ServiceEntry** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**UserGroup** | Pointer to **NullableString** |  | [optional] 
**Layout** | Pointer to [**ListClusters200ResponseAllOfClustersInnerLayout**](ListClusters200ResponseAllOfClustersInnerLayout.md) |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Servers** | Pointer to [**[]ListClusters200ResponseAllOfClustersInnerServersInner**](ListClusters200ResponseAllOfClustersInnerServersInner.md) |  | [optional] 
**Accounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Integrations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Site** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Type** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Zone** | Pointer to [**ListClusters200ResponseAllOfClustersInnerZone**](ListClusters200ResponseAllOfClustersInnerZone.md) |  | [optional] 
**WorkerStats** | Pointer to [**ListClusters200ResponseAllOfClustersInnerWorkerStats**](ListClusters200ResponseAllOfClustersInnerWorkerStats.md) |  | [optional] 
**ContainersCount** | Pointer to **int64** |  | [optional] 
**DeploymentsCount** | Pointer to **int64** |  | [optional] 
**PodsCount** | Pointer to **int64** |  | [optional] 
**JobsCount** | Pointer to **int64** |  | [optional] 
**VolumesCount** | Pointer to **int64** |  | [optional] 
**NamespacesCount** | Pointer to **int64** |  | [optional] 
**WorkersCount** | Pointer to **int64** |  | [optional] 
**ServicesCount** | Pointer to **int64** |  | [optional] 
**Permissions** | Pointer to [**AddCluster200ResponseAllOfClusterPermissions**](AddCluster200ResponseAllOfClusterPermissions.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetCluster200ResponseCluster

`func NewGetCluster200ResponseCluster() *GetCluster200ResponseCluster`

NewGetCluster200ResponseCluster instantiates a new GetCluster200ResponseCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCluster200ResponseClusterWithDefaults

`func NewGetCluster200ResponseClusterWithDefaults() *GetCluster200ResponseCluster`

NewGetCluster200ResponseClusterWithDefaults instantiates a new GetCluster200ResponseCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCluster200ResponseCluster) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCluster200ResponseCluster) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCluster200ResponseCluster) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCluster200ResponseCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *GetCluster200ResponseCluster) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetCluster200ResponseCluster) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetCluster200ResponseCluster) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetCluster200ResponseCluster) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *GetCluster200ResponseCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCluster200ResponseCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCluster200ResponseCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCluster200ResponseCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetCluster200ResponseCluster) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetCluster200ResponseCluster) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetCluster200ResponseCluster) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetCluster200ResponseCluster) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetCluster200ResponseCluster) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetCluster200ResponseCluster) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *GetCluster200ResponseCluster) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetCluster200ResponseCluster) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetCluster200ResponseCluster) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetCluster200ResponseCluster) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetCluster200ResponseCluster) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetCluster200ResponseCluster) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVisibility

`func (o *GetCluster200ResponseCluster) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetCluster200ResponseCluster) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetCluster200ResponseCluster) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetCluster200ResponseCluster) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *GetCluster200ResponseCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCluster200ResponseCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCluster200ResponseCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCluster200ResponseCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetCluster200ResponseCluster) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetCluster200ResponseCluster) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocation

`func (o *GetCluster200ResponseCluster) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetCluster200ResponseCluster) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetCluster200ResponseCluster) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetCluster200ResponseCluster) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *GetCluster200ResponseCluster) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *GetCluster200ResponseCluster) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetEnabled

`func (o *GetCluster200ResponseCluster) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetCluster200ResponseCluster) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetCluster200ResponseCluster) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetCluster200ResponseCluster) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *GetCluster200ResponseCluster) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *GetCluster200ResponseCluster) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *GetCluster200ResponseCluster) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *GetCluster200ResponseCluster) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *GetCluster200ResponseCluster) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *GetCluster200ResponseCluster) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *GetCluster200ResponseCluster) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *GetCluster200ResponseCluster) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *GetCluster200ResponseCluster) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *GetCluster200ResponseCluster) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *GetCluster200ResponseCluster) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *GetCluster200ResponseCluster) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePath

`func (o *GetCluster200ResponseCluster) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *GetCluster200ResponseCluster) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *GetCluster200ResponseCluster) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *GetCluster200ResponseCluster) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### SetServicePathNil

`func (o *GetCluster200ResponseCluster) SetServicePathNil(b bool)`

 SetServicePathNil sets the value for ServicePath to be an explicit nil

### UnsetServicePath
`func (o *GetCluster200ResponseCluster) UnsetServicePath()`

UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
### GetServiceHostname

`func (o *GetCluster200ResponseCluster) GetServiceHostname() string`

GetServiceHostname returns the ServiceHostname field if non-nil, zero value otherwise.

### GetServiceHostnameOk

`func (o *GetCluster200ResponseCluster) GetServiceHostnameOk() (*string, bool)`

GetServiceHostnameOk returns a tuple with the ServiceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHostname

`func (o *GetCluster200ResponseCluster) SetServiceHostname(v string)`

SetServiceHostname sets ServiceHostname field to given value.

### HasServiceHostname

`func (o *GetCluster200ResponseCluster) HasServiceHostname() bool`

HasServiceHostname returns a boolean if a field has been set.

### SetServiceHostnameNil

`func (o *GetCluster200ResponseCluster) SetServiceHostnameNil(b bool)`

 SetServiceHostnameNil sets the value for ServiceHostname to be an explicit nil

### UnsetServiceHostname
`func (o *GetCluster200ResponseCluster) UnsetServiceHostname()`

UnsetServiceHostname ensures that no value is present for ServiceHostname, not even an explicit nil
### GetServicePort

`func (o *GetCluster200ResponseCluster) GetServicePort() int64`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *GetCluster200ResponseCluster) GetServicePortOk() (*int64, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *GetCluster200ResponseCluster) SetServicePort(v int64)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *GetCluster200ResponseCluster) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceUsername

`func (o *GetCluster200ResponseCluster) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *GetCluster200ResponseCluster) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *GetCluster200ResponseCluster) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *GetCluster200ResponseCluster) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *GetCluster200ResponseCluster) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *GetCluster200ResponseCluster) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *GetCluster200ResponseCluster) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *GetCluster200ResponseCluster) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *GetCluster200ResponseCluster) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *GetCluster200ResponseCluster) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *GetCluster200ResponseCluster) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *GetCluster200ResponseCluster) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *GetCluster200ResponseCluster) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *GetCluster200ResponseCluster) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *GetCluster200ResponseCluster) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *GetCluster200ResponseCluster) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### SetServicePasswordHashNil

`func (o *GetCluster200ResponseCluster) SetServicePasswordHashNil(b bool)`

 SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil

### UnsetServicePasswordHash
`func (o *GetCluster200ResponseCluster) UnsetServicePasswordHash()`

UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
### GetServiceToken

`func (o *GetCluster200ResponseCluster) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *GetCluster200ResponseCluster) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *GetCluster200ResponseCluster) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *GetCluster200ResponseCluster) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *GetCluster200ResponseCluster) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *GetCluster200ResponseCluster) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetServiceTokenHash

`func (o *GetCluster200ResponseCluster) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *GetCluster200ResponseCluster) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *GetCluster200ResponseCluster) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *GetCluster200ResponseCluster) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### SetServiceTokenHashNil

`func (o *GetCluster200ResponseCluster) SetServiceTokenHashNil(b bool)`

 SetServiceTokenHashNil sets the value for ServiceTokenHash to be an explicit nil

### UnsetServiceTokenHash
`func (o *GetCluster200ResponseCluster) UnsetServiceTokenHash()`

UnsetServiceTokenHash ensures that no value is present for ServiceTokenHash, not even an explicit nil
### GetServiceAccess

`func (o *GetCluster200ResponseCluster) GetServiceAccess() string`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *GetCluster200ResponseCluster) GetServiceAccessOk() (*string, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *GetCluster200ResponseCluster) SetServiceAccess(v string)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *GetCluster200ResponseCluster) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.

### SetServiceAccessNil

`func (o *GetCluster200ResponseCluster) SetServiceAccessNil(b bool)`

 SetServiceAccessNil sets the value for ServiceAccess to be an explicit nil

### UnsetServiceAccess
`func (o *GetCluster200ResponseCluster) UnsetServiceAccess()`

UnsetServiceAccess ensures that no value is present for ServiceAccess, not even an explicit nil
### GetServiceAccessHash

`func (o *GetCluster200ResponseCluster) GetServiceAccessHash() string`

GetServiceAccessHash returns the ServiceAccessHash field if non-nil, zero value otherwise.

### GetServiceAccessHashOk

`func (o *GetCluster200ResponseCluster) GetServiceAccessHashOk() (*string, bool)`

GetServiceAccessHashOk returns a tuple with the ServiceAccessHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccessHash

`func (o *GetCluster200ResponseCluster) SetServiceAccessHash(v string)`

SetServiceAccessHash sets ServiceAccessHash field to given value.

### HasServiceAccessHash

`func (o *GetCluster200ResponseCluster) HasServiceAccessHash() bool`

HasServiceAccessHash returns a boolean if a field has been set.

### SetServiceAccessHashNil

`func (o *GetCluster200ResponseCluster) SetServiceAccessHashNil(b bool)`

 SetServiceAccessHashNil sets the value for ServiceAccessHash to be an explicit nil

### UnsetServiceAccessHash
`func (o *GetCluster200ResponseCluster) UnsetServiceAccessHash()`

UnsetServiceAccessHash ensures that no value is present for ServiceAccessHash, not even an explicit nil
### GetServiceCert

`func (o *GetCluster200ResponseCluster) GetServiceCert() string`

GetServiceCert returns the ServiceCert field if non-nil, zero value otherwise.

### GetServiceCertOk

`func (o *GetCluster200ResponseCluster) GetServiceCertOk() (*string, bool)`

GetServiceCertOk returns a tuple with the ServiceCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCert

`func (o *GetCluster200ResponseCluster) SetServiceCert(v string)`

SetServiceCert sets ServiceCert field to given value.

### HasServiceCert

`func (o *GetCluster200ResponseCluster) HasServiceCert() bool`

HasServiceCert returns a boolean if a field has been set.

### SetServiceCertNil

`func (o *GetCluster200ResponseCluster) SetServiceCertNil(b bool)`

 SetServiceCertNil sets the value for ServiceCert to be an explicit nil

### UnsetServiceCert
`func (o *GetCluster200ResponseCluster) UnsetServiceCert()`

UnsetServiceCert ensures that no value is present for ServiceCert, not even an explicit nil
### GetServiceCertHash

`func (o *GetCluster200ResponseCluster) GetServiceCertHash() string`

GetServiceCertHash returns the ServiceCertHash field if non-nil, zero value otherwise.

### GetServiceCertHashOk

`func (o *GetCluster200ResponseCluster) GetServiceCertHashOk() (*string, bool)`

GetServiceCertHashOk returns a tuple with the ServiceCertHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCertHash

`func (o *GetCluster200ResponseCluster) SetServiceCertHash(v string)`

SetServiceCertHash sets ServiceCertHash field to given value.

### HasServiceCertHash

`func (o *GetCluster200ResponseCluster) HasServiceCertHash() bool`

HasServiceCertHash returns a boolean if a field has been set.

### SetServiceCertHashNil

`func (o *GetCluster200ResponseCluster) SetServiceCertHashNil(b bool)`

 SetServiceCertHashNil sets the value for ServiceCertHash to be an explicit nil

### UnsetServiceCertHash
`func (o *GetCluster200ResponseCluster) UnsetServiceCertHash()`

UnsetServiceCertHash ensures that no value is present for ServiceCertHash, not even an explicit nil
### GetServiceVersion

`func (o *GetCluster200ResponseCluster) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *GetCluster200ResponseCluster) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *GetCluster200ResponseCluster) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *GetCluster200ResponseCluster) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### SetServiceVersionNil

`func (o *GetCluster200ResponseCluster) SetServiceVersionNil(b bool)`

 SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil

### UnsetServiceVersion
`func (o *GetCluster200ResponseCluster) UnsetServiceVersion()`

UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
### GetSearchDomains

`func (o *GetCluster200ResponseCluster) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *GetCluster200ResponseCluster) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *GetCluster200ResponseCluster) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *GetCluster200ResponseCluster) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### SetSearchDomainsNil

`func (o *GetCluster200ResponseCluster) SetSearchDomainsNil(b bool)`

 SetSearchDomainsNil sets the value for SearchDomains to be an explicit nil

### UnsetSearchDomains
`func (o *GetCluster200ResponseCluster) UnsetSearchDomains()`

UnsetSearchDomains ensures that no value is present for SearchDomains, not even an explicit nil
### GetEnableInternalDns

`func (o *GetCluster200ResponseCluster) GetEnableInternalDns() bool`

GetEnableInternalDns returns the EnableInternalDns field if non-nil, zero value otherwise.

### GetEnableInternalDnsOk

`func (o *GetCluster200ResponseCluster) GetEnableInternalDnsOk() (*bool, bool)`

GetEnableInternalDnsOk returns a tuple with the EnableInternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternalDns

`func (o *GetCluster200ResponseCluster) SetEnableInternalDns(v bool)`

SetEnableInternalDns sets EnableInternalDns field to given value.

### HasEnableInternalDns

`func (o *GetCluster200ResponseCluster) HasEnableInternalDns() bool`

HasEnableInternalDns returns a boolean if a field has been set.

### GetInternalId

`func (o *GetCluster200ResponseCluster) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetCluster200ResponseCluster) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetCluster200ResponseCluster) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetCluster200ResponseCluster) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetCluster200ResponseCluster) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetCluster200ResponseCluster) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetCluster200ResponseCluster) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetCluster200ResponseCluster) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetCluster200ResponseCluster) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetCluster200ResponseCluster) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetCluster200ResponseCluster) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetCluster200ResponseCluster) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDatacenterId

`func (o *GetCluster200ResponseCluster) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *GetCluster200ResponseCluster) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *GetCluster200ResponseCluster) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *GetCluster200ResponseCluster) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### SetDatacenterIdNil

`func (o *GetCluster200ResponseCluster) SetDatacenterIdNil(b bool)`

 SetDatacenterIdNil sets the value for DatacenterId to be an explicit nil

### UnsetDatacenterId
`func (o *GetCluster200ResponseCluster) UnsetDatacenterId()`

UnsetDatacenterId ensures that no value is present for DatacenterId, not even an explicit nil
### GetStatus

`func (o *GetCluster200ResponseCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCluster200ResponseCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCluster200ResponseCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCluster200ResponseCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *GetCluster200ResponseCluster) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetCluster200ResponseCluster) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetCluster200ResponseCluster) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetCluster200ResponseCluster) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetCluster200ResponseCluster) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetCluster200ResponseCluster) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusMessage

`func (o *GetCluster200ResponseCluster) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetCluster200ResponseCluster) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetCluster200ResponseCluster) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetCluster200ResponseCluster) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetCluster200ResponseCluster) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetCluster200ResponseCluster) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetInventoryLevel

`func (o *GetCluster200ResponseCluster) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *GetCluster200ResponseCluster) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *GetCluster200ResponseCluster) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *GetCluster200ResponseCluster) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetLastSync

`func (o *GetCluster200ResponseCluster) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *GetCluster200ResponseCluster) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *GetCluster200ResponseCluster) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *GetCluster200ResponseCluster) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *GetCluster200ResponseCluster) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *GetCluster200ResponseCluster) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetNextRunDate

`func (o *GetCluster200ResponseCluster) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *GetCluster200ResponseCluster) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *GetCluster200ResponseCluster) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *GetCluster200ResponseCluster) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *GetCluster200ResponseCluster) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *GetCluster200ResponseCluster) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetLastSyncDuration

`func (o *GetCluster200ResponseCluster) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *GetCluster200ResponseCluster) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *GetCluster200ResponseCluster) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *GetCluster200ResponseCluster) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *GetCluster200ResponseCluster) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *GetCluster200ResponseCluster) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetDateCreated

`func (o *GetCluster200ResponseCluster) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetCluster200ResponseCluster) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetCluster200ResponseCluster) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetCluster200ResponseCluster) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetCluster200ResponseCluster) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetCluster200ResponseCluster) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetCluster200ResponseCluster) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetCluster200ResponseCluster) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetManaged

`func (o *GetCluster200ResponseCluster) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *GetCluster200ResponseCluster) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *GetCluster200ResponseCluster) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *GetCluster200ResponseCluster) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetLabels

`func (o *GetCluster200ResponseCluster) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetCluster200ResponseCluster) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetCluster200ResponseCluster) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetCluster200ResponseCluster) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetCluster200ResponseCluster) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetCluster200ResponseCluster) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetAutoRecoverPowerState

`func (o *GetCluster200ResponseCluster) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *GetCluster200ResponseCluster) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *GetCluster200ResponseCluster) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *GetCluster200ResponseCluster) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetUseAgent

`func (o *GetCluster200ResponseCluster) GetUseAgent() string`

GetUseAgent returns the UseAgent field if non-nil, zero value otherwise.

### GetUseAgentOk

`func (o *GetCluster200ResponseCluster) GetUseAgentOk() (*string, bool)`

GetUseAgentOk returns a tuple with the UseAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAgent

`func (o *GetCluster200ResponseCluster) SetUseAgent(v string)`

SetUseAgent sets UseAgent field to given value.

### HasUseAgent

`func (o *GetCluster200ResponseCluster) HasUseAgent() bool`

HasUseAgent returns a boolean if a field has been set.

### SetUseAgentNil

`func (o *GetCluster200ResponseCluster) SetUseAgentNil(b bool)`

 SetUseAgentNil sets the value for UseAgent to be an explicit nil

### UnsetUseAgent
`func (o *GetCluster200ResponseCluster) UnsetUseAgent()`

UnsetUseAgent ensures that no value is present for UseAgent, not even an explicit nil
### GetProvisionComplete

`func (o *GetCluster200ResponseCluster) GetProvisionComplete() bool`

GetProvisionComplete returns the ProvisionComplete field if non-nil, zero value otherwise.

### GetProvisionCompleteOk

`func (o *GetCluster200ResponseCluster) GetProvisionCompleteOk() (*bool, bool)`

GetProvisionCompleteOk returns a tuple with the ProvisionComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionComplete

`func (o *GetCluster200ResponseCluster) SetProvisionComplete(v bool)`

SetProvisionComplete sets ProvisionComplete field to given value.

### HasProvisionComplete

`func (o *GetCluster200ResponseCluster) HasProvisionComplete() bool`

HasProvisionComplete returns a boolean if a field has been set.

### GetServiceEntry

`func (o *GetCluster200ResponseCluster) GetServiceEntry() string`

GetServiceEntry returns the ServiceEntry field if non-nil, zero value otherwise.

### GetServiceEntryOk

`func (o *GetCluster200ResponseCluster) GetServiceEntryOk() (*string, bool)`

GetServiceEntryOk returns a tuple with the ServiceEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEntry

`func (o *GetCluster200ResponseCluster) SetServiceEntry(v string)`

SetServiceEntry sets ServiceEntry field to given value.

### HasServiceEntry

`func (o *GetCluster200ResponseCluster) HasServiceEntry() bool`

HasServiceEntry returns a boolean if a field has been set.

### SetServiceEntryNil

`func (o *GetCluster200ResponseCluster) SetServiceEntryNil(b bool)`

 SetServiceEntryNil sets the value for ServiceEntry to be an explicit nil

### UnsetServiceEntry
`func (o *GetCluster200ResponseCluster) UnsetServiceEntry()`

UnsetServiceEntry ensures that no value is present for ServiceEntry, not even an explicit nil
### GetCreatedBy

`func (o *GetCluster200ResponseCluster) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetCluster200ResponseCluster) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetCluster200ResponseCluster) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetCluster200ResponseCluster) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUserGroup

`func (o *GetCluster200ResponseCluster) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *GetCluster200ResponseCluster) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *GetCluster200ResponseCluster) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *GetCluster200ResponseCluster) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### SetUserGroupNil

`func (o *GetCluster200ResponseCluster) SetUserGroupNil(b bool)`

 SetUserGroupNil sets the value for UserGroup to be an explicit nil

### UnsetUserGroup
`func (o *GetCluster200ResponseCluster) UnsetUserGroup()`

UnsetUserGroup ensures that no value is present for UserGroup, not even an explicit nil
### GetLayout

`func (o *GetCluster200ResponseCluster) GetLayout() ListClusters200ResponseAllOfClustersInnerLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *GetCluster200ResponseCluster) GetLayoutOk() (*ListClusters200ResponseAllOfClustersInnerLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *GetCluster200ResponseCluster) SetLayout(v ListClusters200ResponseAllOfClustersInnerLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *GetCluster200ResponseCluster) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetOwner

`func (o *GetCluster200ResponseCluster) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetCluster200ResponseCluster) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetCluster200ResponseCluster) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetCluster200ResponseCluster) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetServers

`func (o *GetCluster200ResponseCluster) GetServers() []ListClusters200ResponseAllOfClustersInnerServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetCluster200ResponseCluster) GetServersOk() (*[]ListClusters200ResponseAllOfClustersInnerServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetCluster200ResponseCluster) SetServers(v []ListClusters200ResponseAllOfClustersInnerServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetCluster200ResponseCluster) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetAccounts

`func (o *GetCluster200ResponseCluster) GetAccounts() []map[string]interface{}`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetCluster200ResponseCluster) GetAccountsOk() (*[]map[string]interface{}, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetCluster200ResponseCluster) SetAccounts(v []map[string]interface{})`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetCluster200ResponseCluster) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetIntegrations

`func (o *GetCluster200ResponseCluster) GetIntegrations() []map[string]interface{}`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *GetCluster200ResponseCluster) GetIntegrationsOk() (*[]map[string]interface{}, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *GetCluster200ResponseCluster) SetIntegrations(v []map[string]interface{})`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *GetCluster200ResponseCluster) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetSite

`func (o *GetCluster200ResponseCluster) GetSite() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GetCluster200ResponseCluster) GetSiteOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GetCluster200ResponseCluster) SetSite(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetSite sets Site field to given value.

### HasSite

`func (o *GetCluster200ResponseCluster) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetType

`func (o *GetCluster200ResponseCluster) GetType() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCluster200ResponseCluster) GetTypeOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCluster200ResponseCluster) SetType(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetType sets Type field to given value.

### HasType

`func (o *GetCluster200ResponseCluster) HasType() bool`

HasType returns a boolean if a field has been set.

### GetZone

`func (o *GetCluster200ResponseCluster) GetZone() ListClusters200ResponseAllOfClustersInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetCluster200ResponseCluster) GetZoneOk() (*ListClusters200ResponseAllOfClustersInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetCluster200ResponseCluster) SetZone(v ListClusters200ResponseAllOfClustersInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetCluster200ResponseCluster) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetWorkerStats

`func (o *GetCluster200ResponseCluster) GetWorkerStats() ListClusters200ResponseAllOfClustersInnerWorkerStats`

GetWorkerStats returns the WorkerStats field if non-nil, zero value otherwise.

### GetWorkerStatsOk

`func (o *GetCluster200ResponseCluster) GetWorkerStatsOk() (*ListClusters200ResponseAllOfClustersInnerWorkerStats, bool)`

GetWorkerStatsOk returns a tuple with the WorkerStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerStats

`func (o *GetCluster200ResponseCluster) SetWorkerStats(v ListClusters200ResponseAllOfClustersInnerWorkerStats)`

SetWorkerStats sets WorkerStats field to given value.

### HasWorkerStats

`func (o *GetCluster200ResponseCluster) HasWorkerStats() bool`

HasWorkerStats returns a boolean if a field has been set.

### GetContainersCount

`func (o *GetCluster200ResponseCluster) GetContainersCount() int64`

GetContainersCount returns the ContainersCount field if non-nil, zero value otherwise.

### GetContainersCountOk

`func (o *GetCluster200ResponseCluster) GetContainersCountOk() (*int64, bool)`

GetContainersCountOk returns a tuple with the ContainersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainersCount

`func (o *GetCluster200ResponseCluster) SetContainersCount(v int64)`

SetContainersCount sets ContainersCount field to given value.

### HasContainersCount

`func (o *GetCluster200ResponseCluster) HasContainersCount() bool`

HasContainersCount returns a boolean if a field has been set.

### GetDeploymentsCount

`func (o *GetCluster200ResponseCluster) GetDeploymentsCount() int64`

GetDeploymentsCount returns the DeploymentsCount field if non-nil, zero value otherwise.

### GetDeploymentsCountOk

`func (o *GetCluster200ResponseCluster) GetDeploymentsCountOk() (*int64, bool)`

GetDeploymentsCountOk returns a tuple with the DeploymentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsCount

`func (o *GetCluster200ResponseCluster) SetDeploymentsCount(v int64)`

SetDeploymentsCount sets DeploymentsCount field to given value.

### HasDeploymentsCount

`func (o *GetCluster200ResponseCluster) HasDeploymentsCount() bool`

HasDeploymentsCount returns a boolean if a field has been set.

### GetPodsCount

`func (o *GetCluster200ResponseCluster) GetPodsCount() int64`

GetPodsCount returns the PodsCount field if non-nil, zero value otherwise.

### GetPodsCountOk

`func (o *GetCluster200ResponseCluster) GetPodsCountOk() (*int64, bool)`

GetPodsCountOk returns a tuple with the PodsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodsCount

`func (o *GetCluster200ResponseCluster) SetPodsCount(v int64)`

SetPodsCount sets PodsCount field to given value.

### HasPodsCount

`func (o *GetCluster200ResponseCluster) HasPodsCount() bool`

HasPodsCount returns a boolean if a field has been set.

### GetJobsCount

`func (o *GetCluster200ResponseCluster) GetJobsCount() int64`

GetJobsCount returns the JobsCount field if non-nil, zero value otherwise.

### GetJobsCountOk

`func (o *GetCluster200ResponseCluster) GetJobsCountOk() (*int64, bool)`

GetJobsCountOk returns a tuple with the JobsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsCount

`func (o *GetCluster200ResponseCluster) SetJobsCount(v int64)`

SetJobsCount sets JobsCount field to given value.

### HasJobsCount

`func (o *GetCluster200ResponseCluster) HasJobsCount() bool`

HasJobsCount returns a boolean if a field has been set.

### GetVolumesCount

`func (o *GetCluster200ResponseCluster) GetVolumesCount() int64`

GetVolumesCount returns the VolumesCount field if non-nil, zero value otherwise.

### GetVolumesCountOk

`func (o *GetCluster200ResponseCluster) GetVolumesCountOk() (*int64, bool)`

GetVolumesCountOk returns a tuple with the VolumesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesCount

`func (o *GetCluster200ResponseCluster) SetVolumesCount(v int64)`

SetVolumesCount sets VolumesCount field to given value.

### HasVolumesCount

`func (o *GetCluster200ResponseCluster) HasVolumesCount() bool`

HasVolumesCount returns a boolean if a field has been set.

### GetNamespacesCount

`func (o *GetCluster200ResponseCluster) GetNamespacesCount() int64`

GetNamespacesCount returns the NamespacesCount field if non-nil, zero value otherwise.

### GetNamespacesCountOk

`func (o *GetCluster200ResponseCluster) GetNamespacesCountOk() (*int64, bool)`

GetNamespacesCountOk returns a tuple with the NamespacesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespacesCount

`func (o *GetCluster200ResponseCluster) SetNamespacesCount(v int64)`

SetNamespacesCount sets NamespacesCount field to given value.

### HasNamespacesCount

`func (o *GetCluster200ResponseCluster) HasNamespacesCount() bool`

HasNamespacesCount returns a boolean if a field has been set.

### GetWorkersCount

`func (o *GetCluster200ResponseCluster) GetWorkersCount() int64`

GetWorkersCount returns the WorkersCount field if non-nil, zero value otherwise.

### GetWorkersCountOk

`func (o *GetCluster200ResponseCluster) GetWorkersCountOk() (*int64, bool)`

GetWorkersCountOk returns a tuple with the WorkersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkersCount

`func (o *GetCluster200ResponseCluster) SetWorkersCount(v int64)`

SetWorkersCount sets WorkersCount field to given value.

### HasWorkersCount

`func (o *GetCluster200ResponseCluster) HasWorkersCount() bool`

HasWorkersCount returns a boolean if a field has been set.

### GetServicesCount

`func (o *GetCluster200ResponseCluster) GetServicesCount() int64`

GetServicesCount returns the ServicesCount field if non-nil, zero value otherwise.

### GetServicesCountOk

`func (o *GetCluster200ResponseCluster) GetServicesCountOk() (*int64, bool)`

GetServicesCountOk returns a tuple with the ServicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesCount

`func (o *GetCluster200ResponseCluster) SetServicesCount(v int64)`

SetServicesCount sets ServicesCount field to given value.

### HasServicesCount

`func (o *GetCluster200ResponseCluster) HasServicesCount() bool`

HasServicesCount returns a boolean if a field has been set.

### GetPermissions

`func (o *GetCluster200ResponseCluster) GetPermissions() AddCluster200ResponseAllOfClusterPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetCluster200ResponseCluster) GetPermissionsOk() (*AddCluster200ResponseAllOfClusterPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetCluster200ResponseCluster) SetPermissions(v AddCluster200ResponseAllOfClusterPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetCluster200ResponseCluster) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetConfig

`func (o *GetCluster200ResponseCluster) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetCluster200ResponseCluster) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetCluster200ResponseCluster) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetCluster200ResponseCluster) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


