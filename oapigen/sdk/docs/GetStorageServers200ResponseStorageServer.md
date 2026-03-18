# GetStorageServers200ResponseStorageServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GetStorageServers200ResponseStorageServerType**](GetStorageServers200ResponseStorageServerType.md) |  | [optional] 
**Chassis** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ServiceUrl** | Pointer to **NullableString** |  | [optional] 
**ServiceHost** | Pointer to **NullableString** |  | [optional] 
**ServicePath** | Pointer to **NullableString** |  | [optional] 
**ServiceToken** | Pointer to **NullableString** |  | [optional] 
**ServiceTokenHash** | Pointer to **NullableString** |  | [optional] 
**ServiceVersion** | Pointer to **NullableString** |  | [optional] 
**ServiceUsername** | Pointer to **NullableString** |  | [optional] 
**ServicePassword** | Pointer to **NullableString** |  | [optional] 
**ServicePasswordHash** | Pointer to **NullableString** |  | [optional] 
**InternalIp** | Pointer to **NullableString** |  | [optional] 
**ExternalIp** | Pointer to **NullableString** |  | [optional] 
**ApiPort** | Pointer to **NullableString** |  | [optional] 
**AdminPort** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**ServerVendor** | Pointer to **NullableString** |  | [optional] 
**ServerModel** | Pointer to **NullableString** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**MaxStorage** | Pointer to **NullableString** |  | [optional] 
**UsedStorage** | Pointer to **NullableString** |  | [optional] 
**DiskCount** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Groups** | Pointer to **[]map[string]interface{}** |  | [optional] 
**HostGroups** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Hosts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tenants** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Owner** | Pointer to [**GetStorageServers200ResponseStorageServerOwner**](GetStorageServers200ResponseStorageServerOwner.md) |  | [optional] 
**Credential** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetStorageServers200ResponseStorageServer

`func NewGetStorageServers200ResponseStorageServer() *GetStorageServers200ResponseStorageServer`

NewGetStorageServers200ResponseStorageServer instantiates a new GetStorageServers200ResponseStorageServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorageServers200ResponseStorageServerWithDefaults

`func NewGetStorageServers200ResponseStorageServerWithDefaults() *GetStorageServers200ResponseStorageServer`

NewGetStorageServers200ResponseStorageServerWithDefaults instantiates a new GetStorageServers200ResponseStorageServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetStorageServers200ResponseStorageServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetStorageServers200ResponseStorageServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetStorageServers200ResponseStorageServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetStorageServers200ResponseStorageServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetStorageServers200ResponseStorageServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetStorageServers200ResponseStorageServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetStorageServers200ResponseStorageServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetStorageServers200ResponseStorageServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetStorageServers200ResponseStorageServer) GetType() GetStorageServers200ResponseStorageServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetStorageServers200ResponseStorageServer) GetTypeOk() (*GetStorageServers200ResponseStorageServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetStorageServers200ResponseStorageServer) SetType(v GetStorageServers200ResponseStorageServerType)`

SetType sets Type field to given value.

### HasType

`func (o *GetStorageServers200ResponseStorageServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChassis

`func (o *GetStorageServers200ResponseStorageServer) GetChassis() string`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *GetStorageServers200ResponseStorageServer) GetChassisOk() (*string, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *GetStorageServers200ResponseStorageServer) SetChassis(v string)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *GetStorageServers200ResponseStorageServer) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### SetChassisNil

`func (o *GetStorageServers200ResponseStorageServer) SetChassisNil(b bool)`

 SetChassisNil sets the value for Chassis to be an explicit nil

### UnsetChassis
`func (o *GetStorageServers200ResponseStorageServer) UnsetChassis()`

UnsetChassis ensures that no value is present for Chassis, not even an explicit nil
### GetVisibility

`func (o *GetStorageServers200ResponseStorageServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetStorageServers200ResponseStorageServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetStorageServers200ResponseStorageServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetStorageServers200ResponseStorageServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *GetStorageServers200ResponseStorageServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetStorageServers200ResponseStorageServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetStorageServers200ResponseStorageServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetStorageServers200ResponseStorageServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetStorageServers200ResponseStorageServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetStorageServers200ResponseStorageServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *GetStorageServers200ResponseStorageServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetStorageServers200ResponseStorageServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetStorageServers200ResponseStorageServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetStorageServers200ResponseStorageServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetStorageServers200ResponseStorageServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetStorageServers200ResponseStorageServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetStorageServers200ResponseStorageServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetStorageServers200ResponseStorageServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetStorageServers200ResponseStorageServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetStorageServers200ResponseStorageServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetStorageServers200ResponseStorageServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetStorageServers200ResponseStorageServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetServiceUrl

`func (o *GetStorageServers200ResponseStorageServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *GetStorageServers200ResponseStorageServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *GetStorageServers200ResponseStorageServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *GetStorageServers200ResponseStorageServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *GetStorageServers200ResponseStorageServer) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *GetStorageServers200ResponseStorageServer) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *GetStorageServers200ResponseStorageServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *GetStorageServers200ResponseStorageServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *GetStorageServers200ResponseStorageServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *GetStorageServers200ResponseStorageServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *GetStorageServers200ResponseStorageServer) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *GetStorageServers200ResponseStorageServer) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePath

`func (o *GetStorageServers200ResponseStorageServer) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *GetStorageServers200ResponseStorageServer) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *GetStorageServers200ResponseStorageServer) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *GetStorageServers200ResponseStorageServer) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### SetServicePathNil

`func (o *GetStorageServers200ResponseStorageServer) SetServicePathNil(b bool)`

 SetServicePathNil sets the value for ServicePath to be an explicit nil

### UnsetServicePath
`func (o *GetStorageServers200ResponseStorageServer) UnsetServicePath()`

UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
### GetServiceToken

`func (o *GetStorageServers200ResponseStorageServer) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *GetStorageServers200ResponseStorageServer) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *GetStorageServers200ResponseStorageServer) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *GetStorageServers200ResponseStorageServer) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *GetStorageServers200ResponseStorageServer) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *GetStorageServers200ResponseStorageServer) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetServiceTokenHash

`func (o *GetStorageServers200ResponseStorageServer) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *GetStorageServers200ResponseStorageServer) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *GetStorageServers200ResponseStorageServer) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *GetStorageServers200ResponseStorageServer) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### SetServiceTokenHashNil

`func (o *GetStorageServers200ResponseStorageServer) SetServiceTokenHashNil(b bool)`

 SetServiceTokenHashNil sets the value for ServiceTokenHash to be an explicit nil

### UnsetServiceTokenHash
`func (o *GetStorageServers200ResponseStorageServer) UnsetServiceTokenHash()`

UnsetServiceTokenHash ensures that no value is present for ServiceTokenHash, not even an explicit nil
### GetServiceVersion

`func (o *GetStorageServers200ResponseStorageServer) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *GetStorageServers200ResponseStorageServer) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *GetStorageServers200ResponseStorageServer) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *GetStorageServers200ResponseStorageServer) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### SetServiceVersionNil

`func (o *GetStorageServers200ResponseStorageServer) SetServiceVersionNil(b bool)`

 SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil

### UnsetServiceVersion
`func (o *GetStorageServers200ResponseStorageServer) UnsetServiceVersion()`

UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
### GetServiceUsername

`func (o *GetStorageServers200ResponseStorageServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *GetStorageServers200ResponseStorageServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *GetStorageServers200ResponseStorageServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *GetStorageServers200ResponseStorageServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *GetStorageServers200ResponseStorageServer) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *GetStorageServers200ResponseStorageServer) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *GetStorageServers200ResponseStorageServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *GetStorageServers200ResponseStorageServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *GetStorageServers200ResponseStorageServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *GetStorageServers200ResponseStorageServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *GetStorageServers200ResponseStorageServer) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *GetStorageServers200ResponseStorageServer) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *GetStorageServers200ResponseStorageServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *GetStorageServers200ResponseStorageServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *GetStorageServers200ResponseStorageServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *GetStorageServers200ResponseStorageServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### SetServicePasswordHashNil

`func (o *GetStorageServers200ResponseStorageServer) SetServicePasswordHashNil(b bool)`

 SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil

### UnsetServicePasswordHash
`func (o *GetStorageServers200ResponseStorageServer) UnsetServicePasswordHash()`

UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
### GetInternalIp

`func (o *GetStorageServers200ResponseStorageServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *GetStorageServers200ResponseStorageServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *GetStorageServers200ResponseStorageServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *GetStorageServers200ResponseStorageServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *GetStorageServers200ResponseStorageServer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *GetStorageServers200ResponseStorageServer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetExternalIp

`func (o *GetStorageServers200ResponseStorageServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *GetStorageServers200ResponseStorageServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *GetStorageServers200ResponseStorageServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *GetStorageServers200ResponseStorageServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *GetStorageServers200ResponseStorageServer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *GetStorageServers200ResponseStorageServer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetApiPort

`func (o *GetStorageServers200ResponseStorageServer) GetApiPort() string`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *GetStorageServers200ResponseStorageServer) GetApiPortOk() (*string, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *GetStorageServers200ResponseStorageServer) SetApiPort(v string)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *GetStorageServers200ResponseStorageServer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *GetStorageServers200ResponseStorageServer) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *GetStorageServers200ResponseStorageServer) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *GetStorageServers200ResponseStorageServer) GetAdminPort() string`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *GetStorageServers200ResponseStorageServer) GetAdminPortOk() (*string, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *GetStorageServers200ResponseStorageServer) SetAdminPort(v string)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *GetStorageServers200ResponseStorageServer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *GetStorageServers200ResponseStorageServer) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *GetStorageServers200ResponseStorageServer) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetConfig

`func (o *GetStorageServers200ResponseStorageServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetStorageServers200ResponseStorageServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetStorageServers200ResponseStorageServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetStorageServers200ResponseStorageServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRefType

`func (o *GetStorageServers200ResponseStorageServer) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetStorageServers200ResponseStorageServer) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetStorageServers200ResponseStorageServer) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetStorageServers200ResponseStorageServer) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetStorageServers200ResponseStorageServer) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetStorageServers200ResponseStorageServer) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetStorageServers200ResponseStorageServer) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetStorageServers200ResponseStorageServer) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetCategory

`func (o *GetStorageServers200ResponseStorageServer) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetStorageServers200ResponseStorageServer) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetStorageServers200ResponseStorageServer) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetStorageServers200ResponseStorageServer) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetStorageServers200ResponseStorageServer) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetStorageServers200ResponseStorageServer) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetServerVendor

`func (o *GetStorageServers200ResponseStorageServer) GetServerVendor() string`

GetServerVendor returns the ServerVendor field if non-nil, zero value otherwise.

### GetServerVendorOk

`func (o *GetStorageServers200ResponseStorageServer) GetServerVendorOk() (*string, bool)`

GetServerVendorOk returns a tuple with the ServerVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVendor

`func (o *GetStorageServers200ResponseStorageServer) SetServerVendor(v string)`

SetServerVendor sets ServerVendor field to given value.

### HasServerVendor

`func (o *GetStorageServers200ResponseStorageServer) HasServerVendor() bool`

HasServerVendor returns a boolean if a field has been set.

### SetServerVendorNil

`func (o *GetStorageServers200ResponseStorageServer) SetServerVendorNil(b bool)`

 SetServerVendorNil sets the value for ServerVendor to be an explicit nil

### UnsetServerVendor
`func (o *GetStorageServers200ResponseStorageServer) UnsetServerVendor()`

UnsetServerVendor ensures that no value is present for ServerVendor, not even an explicit nil
### GetServerModel

`func (o *GetStorageServers200ResponseStorageServer) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *GetStorageServers200ResponseStorageServer) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *GetStorageServers200ResponseStorageServer) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *GetStorageServers200ResponseStorageServer) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### SetServerModelNil

`func (o *GetStorageServers200ResponseStorageServer) SetServerModelNil(b bool)`

 SetServerModelNil sets the value for ServerModel to be an explicit nil

### UnsetServerModel
`func (o *GetStorageServers200ResponseStorageServer) UnsetServerModel()`

UnsetServerModel ensures that no value is present for ServerModel, not even an explicit nil
### GetSerialNumber

`func (o *GetStorageServers200ResponseStorageServer) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *GetStorageServers200ResponseStorageServer) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *GetStorageServers200ResponseStorageServer) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *GetStorageServers200ResponseStorageServer) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *GetStorageServers200ResponseStorageServer) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *GetStorageServers200ResponseStorageServer) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetStatus

`func (o *GetStorageServers200ResponseStorageServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStorageServers200ResponseStorageServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStorageServers200ResponseStorageServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetStorageServers200ResponseStorageServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetStorageServers200ResponseStorageServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetStorageServers200ResponseStorageServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetStorageServers200ResponseStorageServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetStorageServers200ResponseStorageServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetStorageServers200ResponseStorageServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetStorageServers200ResponseStorageServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *GetStorageServers200ResponseStorageServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetStorageServers200ResponseStorageServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetStorageServers200ResponseStorageServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetStorageServers200ResponseStorageServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetErrorMessage

`func (o *GetStorageServers200ResponseStorageServer) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetStorageServers200ResponseStorageServer) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetStorageServers200ResponseStorageServer) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetStorageServers200ResponseStorageServer) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetStorageServers200ResponseStorageServer) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetStorageServers200ResponseStorageServer) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetMaxStorage

`func (o *GetStorageServers200ResponseStorageServer) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetStorageServers200ResponseStorageServer) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetStorageServers200ResponseStorageServer) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetStorageServers200ResponseStorageServer) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### SetMaxStorageNil

`func (o *GetStorageServers200ResponseStorageServer) SetMaxStorageNil(b bool)`

 SetMaxStorageNil sets the value for MaxStorage to be an explicit nil

### UnsetMaxStorage
`func (o *GetStorageServers200ResponseStorageServer) UnsetMaxStorage()`

UnsetMaxStorage ensures that no value is present for MaxStorage, not even an explicit nil
### GetUsedStorage

`func (o *GetStorageServers200ResponseStorageServer) GetUsedStorage() string`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *GetStorageServers200ResponseStorageServer) GetUsedStorageOk() (*string, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *GetStorageServers200ResponseStorageServer) SetUsedStorage(v string)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *GetStorageServers200ResponseStorageServer) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### SetUsedStorageNil

`func (o *GetStorageServers200ResponseStorageServer) SetUsedStorageNil(b bool)`

 SetUsedStorageNil sets the value for UsedStorage to be an explicit nil

### UnsetUsedStorage
`func (o *GetStorageServers200ResponseStorageServer) UnsetUsedStorage()`

UnsetUsedStorage ensures that no value is present for UsedStorage, not even an explicit nil
### GetDiskCount

`func (o *GetStorageServers200ResponseStorageServer) GetDiskCount() string`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *GetStorageServers200ResponseStorageServer) GetDiskCountOk() (*string, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *GetStorageServers200ResponseStorageServer) SetDiskCount(v string)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *GetStorageServers200ResponseStorageServer) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### SetDiskCountNil

`func (o *GetStorageServers200ResponseStorageServer) SetDiskCountNil(b bool)`

 SetDiskCountNil sets the value for DiskCount to be an explicit nil

### UnsetDiskCount
`func (o *GetStorageServers200ResponseStorageServer) UnsetDiskCount()`

UnsetDiskCount ensures that no value is present for DiskCount, not even an explicit nil
### GetDateCreated

`func (o *GetStorageServers200ResponseStorageServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetStorageServers200ResponseStorageServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetStorageServers200ResponseStorageServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetStorageServers200ResponseStorageServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetStorageServers200ResponseStorageServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetStorageServers200ResponseStorageServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetStorageServers200ResponseStorageServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetStorageServers200ResponseStorageServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEnabled

`func (o *GetStorageServers200ResponseStorageServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetStorageServers200ResponseStorageServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetStorageServers200ResponseStorageServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetStorageServers200ResponseStorageServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroups

`func (o *GetStorageServers200ResponseStorageServer) GetGroups() []map[string]interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetStorageServers200ResponseStorageServer) GetGroupsOk() (*[]map[string]interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetStorageServers200ResponseStorageServer) SetGroups(v []map[string]interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GetStorageServers200ResponseStorageServer) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHostGroups

`func (o *GetStorageServers200ResponseStorageServer) GetHostGroups() []map[string]interface{}`

GetHostGroups returns the HostGroups field if non-nil, zero value otherwise.

### GetHostGroupsOk

`func (o *GetStorageServers200ResponseStorageServer) GetHostGroupsOk() (*[]map[string]interface{}, bool)`

GetHostGroupsOk returns a tuple with the HostGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroups

`func (o *GetStorageServers200ResponseStorageServer) SetHostGroups(v []map[string]interface{})`

SetHostGroups sets HostGroups field to given value.

### HasHostGroups

`func (o *GetStorageServers200ResponseStorageServer) HasHostGroups() bool`

HasHostGroups returns a boolean if a field has been set.

### GetHosts

`func (o *GetStorageServers200ResponseStorageServer) GetHosts() []map[string]interface{}`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *GetStorageServers200ResponseStorageServer) GetHostsOk() (*[]map[string]interface{}, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *GetStorageServers200ResponseStorageServer) SetHosts(v []map[string]interface{})`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *GetStorageServers200ResponseStorageServer) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetTenants

`func (o *GetStorageServers200ResponseStorageServer) GetTenants() []map[string]interface{}`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetStorageServers200ResponseStorageServer) GetTenantsOk() (*[]map[string]interface{}, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetStorageServers200ResponseStorageServer) SetTenants(v []map[string]interface{})`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetStorageServers200ResponseStorageServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetOwner

`func (o *GetStorageServers200ResponseStorageServer) GetOwner() GetStorageServers200ResponseStorageServerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetStorageServers200ResponseStorageServer) GetOwnerOk() (*GetStorageServers200ResponseStorageServerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetStorageServers200ResponseStorageServer) SetOwner(v GetStorageServers200ResponseStorageServerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetStorageServers200ResponseStorageServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCredential

`func (o *GetStorageServers200ResponseStorageServer) GetCredential() map[string]interface{}`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetStorageServers200ResponseStorageServer) GetCredentialOk() (*map[string]interface{}, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetStorageServers200ResponseStorageServer) SetCredential(v map[string]interface{})`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetStorageServers200ResponseStorageServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


