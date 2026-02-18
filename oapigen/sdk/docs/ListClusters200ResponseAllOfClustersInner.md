# ListClusters200ResponseAllOfClustersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
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
**ServiceToken** | Pointer to **string** |  | [optional] 
**ServiceTokenHash** | Pointer to **string** |  | [optional] 
**ServiceAccess** | Pointer to **string** |  | [optional] 
**ServiceAccessHash** | Pointer to **string** |  | [optional] 
**ServiceCert** | Pointer to **NullableString** |  | [optional] 
**ServiceCertHash** | Pointer to **NullableString** |  | [optional] 
**ServiceVersion** | Pointer to **string** |  | [optional] 
**SearchDomains** | Pointer to **NullableString** |  | [optional] 
**EnableInternalDns** | Pointer to **bool** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**DatacenterId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**InventoryLevel** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **time.Time** |  | [optional] 
**NextRunDate** | Pointer to **time.Time** |  | [optional] 
**LastSyncDuration** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Managed** | Pointer to **bool** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**AutoRecoverPowerState** | Pointer to **bool** |  | [optional] 
**ServiceEntry** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to [**ListClusters200ResponseAllOfClustersInnerCreatedBy**](ListClusters200ResponseAllOfClustersInnerCreatedBy.md) |  | [optional] 
**UserGroup** | Pointer to **NullableString** |  | [optional] 
**Layout** | Pointer to [**ListClusters200ResponseAllOfClustersInnerLayout**](ListClusters200ResponseAllOfClustersInnerLayout.md) |  | [optional] 
**Owner** | Pointer to [**ListClusters200ResponseAllOfClustersInnerOwner**](ListClusters200ResponseAllOfClustersInnerOwner.md) |  | [optional] 
**Servers** | Pointer to [**[]ListClusters200ResponseAllOfClustersInnerServersInner**](ListClusters200ResponseAllOfClustersInnerServersInner.md) |  | [optional] 
**Accounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Integrations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Site** | Pointer to [**ListClusters200ResponseAllOfClustersInnerSite**](ListClusters200ResponseAllOfClustersInnerSite.md) |  | [optional] 
**Type** | Pointer to [**ListClusters200ResponseAllOfClustersInnerType**](ListClusters200ResponseAllOfClustersInnerType.md) |  | [optional] 
**Zone** | Pointer to [**ListClusters200ResponseAllOfClustersInnerZone**](ListClusters200ResponseAllOfClustersInnerZone.md) |  | [optional] 
**WorkerStats** | Pointer to [**ListClusters200ResponseAllOfClustersInnerWorkerStats**](ListClusters200ResponseAllOfClustersInnerWorkerStats.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListClusters200ResponseAllOfClustersInner

`func NewListClusters200ResponseAllOfClustersInner() *ListClusters200ResponseAllOfClustersInner`

NewListClusters200ResponseAllOfClustersInner instantiates a new ListClusters200ResponseAllOfClustersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusters200ResponseAllOfClustersInnerWithDefaults

`func NewListClusters200ResponseAllOfClustersInnerWithDefaults() *ListClusters200ResponseAllOfClustersInner`

NewListClusters200ResponseAllOfClustersInnerWithDefaults instantiates a new ListClusters200ResponseAllOfClustersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusters200ResponseAllOfClustersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusters200ResponseAllOfClustersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusters200ResponseAllOfClustersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusters200ResponseAllOfClustersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusters200ResponseAllOfClustersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusters200ResponseAllOfClustersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListClusters200ResponseAllOfClustersInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusters200ResponseAllOfClustersInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusters200ResponseAllOfClustersInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *ListClusters200ResponseAllOfClustersInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListClusters200ResponseAllOfClustersInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListClusters200ResponseAllOfClustersInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVisibility

`func (o *ListClusters200ResponseAllOfClustersInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListClusters200ResponseAllOfClustersInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListClusters200ResponseAllOfClustersInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListClusters200ResponseAllOfClustersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusters200ResponseAllOfClustersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusters200ResponseAllOfClustersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocation

`func (o *ListClusters200ResponseAllOfClustersInner) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ListClusters200ResponseAllOfClustersInner) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ListClusters200ResponseAllOfClustersInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetEnabled

`func (o *ListClusters200ResponseAllOfClustersInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListClusters200ResponseAllOfClustersInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListClusters200ResponseAllOfClustersInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePath

`func (o *ListClusters200ResponseAllOfClustersInner) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *ListClusters200ResponseAllOfClustersInner) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *ListClusters200ResponseAllOfClustersInner) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### SetServicePathNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetServicePathNil(b bool)`

 SetServicePathNil sets the value for ServicePath to be an explicit nil

### UnsetServicePath
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetServicePath()`

UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
### GetServiceHostname

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceHostname() string`

GetServiceHostname returns the ServiceHostname field if non-nil, zero value otherwise.

### GetServiceHostnameOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceHostnameOk() (*string, bool)`

GetServiceHostnameOk returns a tuple with the ServiceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHostname

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceHostname(v string)`

SetServiceHostname sets ServiceHostname field to given value.

### HasServiceHostname

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceHostname() bool`

HasServiceHostname returns a boolean if a field has been set.

### SetServiceHostnameNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceHostnameNil(b bool)`

 SetServiceHostnameNil sets the value for ServiceHostname to be an explicit nil

### UnsetServiceHostname
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetServiceHostname()`

UnsetServiceHostname ensures that no value is present for ServiceHostname, not even an explicit nil
### GetServicePort

`func (o *ListClusters200ResponseAllOfClustersInner) GetServicePort() int64`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServicePortOk() (*int64, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *ListClusters200ResponseAllOfClustersInner) SetServicePort(v int64)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *ListClusters200ResponseAllOfClustersInner) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceUsername

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *ListClusters200ResponseAllOfClustersInner) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *ListClusters200ResponseAllOfClustersInner) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *ListClusters200ResponseAllOfClustersInner) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *ListClusters200ResponseAllOfClustersInner) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *ListClusters200ResponseAllOfClustersInner) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *ListClusters200ResponseAllOfClustersInner) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### SetServicePasswordHashNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetServicePasswordHashNil(b bool)`

 SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil

### UnsetServicePasswordHash
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetServicePasswordHash()`

UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
### GetServiceToken

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceTokenHash

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### GetServiceAccess

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceAccess() string`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceAccessOk() (*string, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceAccess(v string)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.

### GetServiceAccessHash

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceAccessHash() string`

GetServiceAccessHash returns the ServiceAccessHash field if non-nil, zero value otherwise.

### GetServiceAccessHashOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceAccessHashOk() (*string, bool)`

GetServiceAccessHashOk returns a tuple with the ServiceAccessHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccessHash

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceAccessHash(v string)`

SetServiceAccessHash sets ServiceAccessHash field to given value.

### HasServiceAccessHash

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceAccessHash() bool`

HasServiceAccessHash returns a boolean if a field has been set.

### GetServiceCert

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceCert() string`

GetServiceCert returns the ServiceCert field if non-nil, zero value otherwise.

### GetServiceCertOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceCertOk() (*string, bool)`

GetServiceCertOk returns a tuple with the ServiceCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCert

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceCert(v string)`

SetServiceCert sets ServiceCert field to given value.

### HasServiceCert

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceCert() bool`

HasServiceCert returns a boolean if a field has been set.

### SetServiceCertNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceCertNil(b bool)`

 SetServiceCertNil sets the value for ServiceCert to be an explicit nil

### UnsetServiceCert
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetServiceCert()`

UnsetServiceCert ensures that no value is present for ServiceCert, not even an explicit nil
### GetServiceCertHash

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceCertHash() string`

GetServiceCertHash returns the ServiceCertHash field if non-nil, zero value otherwise.

### GetServiceCertHashOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceCertHashOk() (*string, bool)`

GetServiceCertHashOk returns a tuple with the ServiceCertHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCertHash

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceCertHash(v string)`

SetServiceCertHash sets ServiceCertHash field to given value.

### HasServiceCertHash

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceCertHash() bool`

HasServiceCertHash returns a boolean if a field has been set.

### SetServiceCertHashNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceCertHashNil(b bool)`

 SetServiceCertHashNil sets the value for ServiceCertHash to be an explicit nil

### UnsetServiceCertHash
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetServiceCertHash()`

UnsetServiceCertHash ensures that no value is present for ServiceCertHash, not even an explicit nil
### GetServiceVersion

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetSearchDomains

`func (o *ListClusters200ResponseAllOfClustersInner) GetSearchDomains() string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetSearchDomainsOk() (*string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *ListClusters200ResponseAllOfClustersInner) SetSearchDomains(v string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *ListClusters200ResponseAllOfClustersInner) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.

### SetSearchDomainsNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetSearchDomainsNil(b bool)`

 SetSearchDomainsNil sets the value for SearchDomains to be an explicit nil

### UnsetSearchDomains
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetSearchDomains()`

UnsetSearchDomains ensures that no value is present for SearchDomains, not even an explicit nil
### GetEnableInternalDns

`func (o *ListClusters200ResponseAllOfClustersInner) GetEnableInternalDns() bool`

GetEnableInternalDns returns the EnableInternalDns field if non-nil, zero value otherwise.

### GetEnableInternalDnsOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetEnableInternalDnsOk() (*bool, bool)`

GetEnableInternalDnsOk returns a tuple with the EnableInternalDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternalDns

`func (o *ListClusters200ResponseAllOfClustersInner) SetEnableInternalDns(v bool)`

SetEnableInternalDns sets EnableInternalDns field to given value.

### HasEnableInternalDns

`func (o *ListClusters200ResponseAllOfClustersInner) HasEnableInternalDns() bool`

HasEnableInternalDns returns a boolean if a field has been set.

### GetInternalId

`func (o *ListClusters200ResponseAllOfClustersInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListClusters200ResponseAllOfClustersInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListClusters200ResponseAllOfClustersInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *ListClusters200ResponseAllOfClustersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListClusters200ResponseAllOfClustersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListClusters200ResponseAllOfClustersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDatacenterId

`func (o *ListClusters200ResponseAllOfClustersInner) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *ListClusters200ResponseAllOfClustersInner) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *ListClusters200ResponseAllOfClustersInner) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### SetDatacenterIdNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetDatacenterIdNil(b bool)`

 SetDatacenterIdNil sets the value for DatacenterId to be an explicit nil

### UnsetDatacenterId
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetDatacenterId()`

UnsetDatacenterId ensures that no value is present for DatacenterId, not even an explicit nil
### GetStatus

`func (o *ListClusters200ResponseAllOfClustersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusters200ResponseAllOfClustersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusters200ResponseAllOfClustersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListClusters200ResponseAllOfClustersInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListClusters200ResponseAllOfClustersInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListClusters200ResponseAllOfClustersInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListClusters200ResponseAllOfClustersInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListClusters200ResponseAllOfClustersInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListClusters200ResponseAllOfClustersInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *ListClusters200ResponseAllOfClustersInner) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *ListClusters200ResponseAllOfClustersInner) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *ListClusters200ResponseAllOfClustersInner) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetLastSync

`func (o *ListClusters200ResponseAllOfClustersInner) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *ListClusters200ResponseAllOfClustersInner) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *ListClusters200ResponseAllOfClustersInner) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetNextRunDate

`func (o *ListClusters200ResponseAllOfClustersInner) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *ListClusters200ResponseAllOfClustersInner) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *ListClusters200ResponseAllOfClustersInner) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *ListClusters200ResponseAllOfClustersInner) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *ListClusters200ResponseAllOfClustersInner) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *ListClusters200ResponseAllOfClustersInner) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListClusters200ResponseAllOfClustersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListClusters200ResponseAllOfClustersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListClusters200ResponseAllOfClustersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListClusters200ResponseAllOfClustersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClusters200ResponseAllOfClustersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClusters200ResponseAllOfClustersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetManaged

`func (o *ListClusters200ResponseAllOfClustersInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListClusters200ResponseAllOfClustersInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListClusters200ResponseAllOfClustersInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetLabels

`func (o *ListClusters200ResponseAllOfClustersInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListClusters200ResponseAllOfClustersInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListClusters200ResponseAllOfClustersInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAutoRecoverPowerState

`func (o *ListClusters200ResponseAllOfClustersInner) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *ListClusters200ResponseAllOfClustersInner) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *ListClusters200ResponseAllOfClustersInner) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetServiceEntry

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceEntry() string`

GetServiceEntry returns the ServiceEntry field if non-nil, zero value otherwise.

### GetServiceEntryOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServiceEntryOk() (*string, bool)`

GetServiceEntryOk returns a tuple with the ServiceEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEntry

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceEntry(v string)`

SetServiceEntry sets ServiceEntry field to given value.

### HasServiceEntry

`func (o *ListClusters200ResponseAllOfClustersInner) HasServiceEntry() bool`

HasServiceEntry returns a boolean if a field has been set.

### SetServiceEntryNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetServiceEntryNil(b bool)`

 SetServiceEntryNil sets the value for ServiceEntry to be an explicit nil

### UnsetServiceEntry
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetServiceEntry()`

UnsetServiceEntry ensures that no value is present for ServiceEntry, not even an explicit nil
### GetCreatedBy

`func (o *ListClusters200ResponseAllOfClustersInner) GetCreatedBy() ListClusters200ResponseAllOfClustersInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetCreatedByOk() (*ListClusters200ResponseAllOfClustersInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListClusters200ResponseAllOfClustersInner) SetCreatedBy(v ListClusters200ResponseAllOfClustersInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListClusters200ResponseAllOfClustersInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUserGroup

`func (o *ListClusters200ResponseAllOfClustersInner) GetUserGroup() string`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetUserGroupOk() (*string, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *ListClusters200ResponseAllOfClustersInner) SetUserGroup(v string)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *ListClusters200ResponseAllOfClustersInner) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### SetUserGroupNil

`func (o *ListClusters200ResponseAllOfClustersInner) SetUserGroupNil(b bool)`

 SetUserGroupNil sets the value for UserGroup to be an explicit nil

### UnsetUserGroup
`func (o *ListClusters200ResponseAllOfClustersInner) UnsetUserGroup()`

UnsetUserGroup ensures that no value is present for UserGroup, not even an explicit nil
### GetLayout

`func (o *ListClusters200ResponseAllOfClustersInner) GetLayout() ListClusters200ResponseAllOfClustersInnerLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetLayoutOk() (*ListClusters200ResponseAllOfClustersInnerLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ListClusters200ResponseAllOfClustersInner) SetLayout(v ListClusters200ResponseAllOfClustersInnerLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *ListClusters200ResponseAllOfClustersInner) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetOwner

`func (o *ListClusters200ResponseAllOfClustersInner) GetOwner() ListClusters200ResponseAllOfClustersInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetOwnerOk() (*ListClusters200ResponseAllOfClustersInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListClusters200ResponseAllOfClustersInner) SetOwner(v ListClusters200ResponseAllOfClustersInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListClusters200ResponseAllOfClustersInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetServers

`func (o *ListClusters200ResponseAllOfClustersInner) GetServers() []ListClusters200ResponseAllOfClustersInnerServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetServersOk() (*[]ListClusters200ResponseAllOfClustersInnerServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ListClusters200ResponseAllOfClustersInner) SetServers(v []ListClusters200ResponseAllOfClustersInnerServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ListClusters200ResponseAllOfClustersInner) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetAccounts

`func (o *ListClusters200ResponseAllOfClustersInner) GetAccounts() []map[string]interface{}`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetAccountsOk() (*[]map[string]interface{}, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ListClusters200ResponseAllOfClustersInner) SetAccounts(v []map[string]interface{})`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ListClusters200ResponseAllOfClustersInner) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetIntegrations

`func (o *ListClusters200ResponseAllOfClustersInner) GetIntegrations() []map[string]interface{}`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetIntegrationsOk() (*[]map[string]interface{}, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *ListClusters200ResponseAllOfClustersInner) SetIntegrations(v []map[string]interface{})`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *ListClusters200ResponseAllOfClustersInner) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetSite

`func (o *ListClusters200ResponseAllOfClustersInner) GetSite() ListClusters200ResponseAllOfClustersInnerSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetSiteOk() (*ListClusters200ResponseAllOfClustersInnerSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ListClusters200ResponseAllOfClustersInner) SetSite(v ListClusters200ResponseAllOfClustersInnerSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *ListClusters200ResponseAllOfClustersInner) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetType

`func (o *ListClusters200ResponseAllOfClustersInner) GetType() ListClusters200ResponseAllOfClustersInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetTypeOk() (*ListClusters200ResponseAllOfClustersInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusters200ResponseAllOfClustersInner) SetType(v ListClusters200ResponseAllOfClustersInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusters200ResponseAllOfClustersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetZone

`func (o *ListClusters200ResponseAllOfClustersInner) GetZone() ListClusters200ResponseAllOfClustersInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetZoneOk() (*ListClusters200ResponseAllOfClustersInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListClusters200ResponseAllOfClustersInner) SetZone(v ListClusters200ResponseAllOfClustersInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListClusters200ResponseAllOfClustersInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetWorkerStats

`func (o *ListClusters200ResponseAllOfClustersInner) GetWorkerStats() ListClusters200ResponseAllOfClustersInnerWorkerStats`

GetWorkerStats returns the WorkerStats field if non-nil, zero value otherwise.

### GetWorkerStatsOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetWorkerStatsOk() (*ListClusters200ResponseAllOfClustersInnerWorkerStats, bool)`

GetWorkerStatsOk returns a tuple with the WorkerStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerStats

`func (o *ListClusters200ResponseAllOfClustersInner) SetWorkerStats(v ListClusters200ResponseAllOfClustersInnerWorkerStats)`

SetWorkerStats sets WorkerStats field to given value.

### HasWorkerStats

`func (o *ListClusters200ResponseAllOfClustersInner) HasWorkerStats() bool`

HasWorkerStats returns a boolean if a field has been set.

### GetConfig

`func (o *ListClusters200ResponseAllOfClustersInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListClusters200ResponseAllOfClustersInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListClusters200ResponseAllOfClustersInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListClusters200ResponseAllOfClustersInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


