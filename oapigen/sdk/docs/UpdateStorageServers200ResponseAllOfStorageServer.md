# UpdateStorageServers200ResponseAllOfStorageServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**UpdateStorageServers200ResponseAllOfStorageServerType**](UpdateStorageServers200ResponseAllOfStorageServerType.md) |  | [optional] 
**Chassis** | Pointer to [**UpdateStorageServers200ResponseAllOfStorageServerChassis**](UpdateStorageServers200ResponseAllOfStorageServerChassis.md) |  | [optional] 
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
**ApiPort** | Pointer to **NullableInt32** |  | [optional] 
**AdminPort** | Pointer to **NullableInt32** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**ServerVendor** | Pointer to **NullableString** |  | [optional] 
**ServerModel** | Pointer to **NullableString** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**MaxStorage** | Pointer to **NullableInt64** |  | [optional] 
**UsedStorage** | Pointer to **NullableInt64** |  | [optional] 
**DiskCount** | Pointer to **NullableInt32** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Groups** | Pointer to [**[]UpdateStorageServers200ResponseAllOfStorageServerGroupsInner**](UpdateStorageServers200ResponseAllOfStorageServerGroupsInner.md) |  | [optional] 
**HostGroups** | Pointer to [**[]UpdateStorageServers200ResponseAllOfStorageServerHostGroupsInner**](UpdateStorageServers200ResponseAllOfStorageServerHostGroupsInner.md) |  | [optional] 
**Hosts** | Pointer to [**[]UpdateStorageServers200ResponseAllOfStorageServerHostsInner**](UpdateStorageServers200ResponseAllOfStorageServerHostsInner.md) |  | [optional] 
**Tenants** | Pointer to [**[]UpdateStorageServers200ResponseAllOfStorageServerTenantsInner**](UpdateStorageServers200ResponseAllOfStorageServerTenantsInner.md) |  | [optional] 
**Owner** | Pointer to [**UpdateStorageServers200ResponseAllOfStorageServerOwner**](UpdateStorageServers200ResponseAllOfStorageServerOwner.md) |  | [optional] 
**Credential** | Pointer to [**UpdateStorageServers200ResponseAllOfStorageServerCredential**](UpdateStorageServers200ResponseAllOfStorageServerCredential.md) |  | [optional] 

## Methods

### NewUpdateStorageServers200ResponseAllOfStorageServer

`func NewUpdateStorageServers200ResponseAllOfStorageServer() *UpdateStorageServers200ResponseAllOfStorageServer`

NewUpdateStorageServers200ResponseAllOfStorageServer instantiates a new UpdateStorageServers200ResponseAllOfStorageServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageServers200ResponseAllOfStorageServerWithDefaults

`func NewUpdateStorageServers200ResponseAllOfStorageServerWithDefaults() *UpdateStorageServers200ResponseAllOfStorageServer`

NewUpdateStorageServers200ResponseAllOfStorageServerWithDefaults instantiates a new UpdateStorageServers200ResponseAllOfStorageServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetType() UpdateStorageServers200ResponseAllOfStorageServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetTypeOk() (*UpdateStorageServers200ResponseAllOfStorageServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetType(v UpdateStorageServers200ResponseAllOfStorageServerType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChassis

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetChassis() UpdateStorageServers200ResponseAllOfStorageServerChassis`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetChassisOk() (*UpdateStorageServers200ResponseAllOfStorageServerChassis, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetChassis(v UpdateStorageServers200ResponseAllOfStorageServerChassis)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetServiceUrl

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePath

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### SetServicePathNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServicePathNil(b bool)`

 SetServicePathNil sets the value for ServicePath to be an explicit nil

### UnsetServicePath
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServicePath()`

UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
### GetServiceToken

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetServiceTokenHash

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### SetServiceTokenHashNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceTokenHashNil(b bool)`

 SetServiceTokenHashNil sets the value for ServiceTokenHash to be an explicit nil

### UnsetServiceTokenHash
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServiceTokenHash()`

UnsetServiceTokenHash ensures that no value is present for ServiceTokenHash, not even an explicit nil
### GetServiceVersion

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### SetServiceVersionNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceVersionNil(b bool)`

 SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil

### UnsetServiceVersion
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServiceVersion()`

UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
### GetServiceUsername

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### SetServicePasswordHashNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServicePasswordHashNil(b bool)`

 SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil

### UnsetServicePasswordHash
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServicePasswordHash()`

UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
### GetInternalIp

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetExternalIp

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetApiPort

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetAdminPort() int32`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetAdminPortOk() (*int32, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetAdminPort(v int32)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetConfig

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRefType

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetServerVendor

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServerVendor() string`

GetServerVendor returns the ServerVendor field if non-nil, zero value otherwise.

### GetServerVendorOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServerVendorOk() (*string, bool)`

GetServerVendorOk returns a tuple with the ServerVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVendor

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServerVendor(v string)`

SetServerVendor sets ServerVendor field to given value.

### HasServerVendor

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServerVendor() bool`

HasServerVendor returns a boolean if a field has been set.

### SetServerVendorNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServerVendorNil(b bool)`

 SetServerVendorNil sets the value for ServerVendor to be an explicit nil

### UnsetServerVendor
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServerVendor()`

UnsetServerVendor ensures that no value is present for ServerVendor, not even an explicit nil
### GetServerModel

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### SetServerModelNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetServerModelNil(b bool)`

 SetServerModelNil sets the value for ServerModel to be an explicit nil

### UnsetServerModel
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetServerModel()`

UnsetServerModel ensures that no value is present for ServerModel, not even an explicit nil
### GetSerialNumber

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetStatus

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusMessage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetErrorMessage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetMaxStorage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### SetMaxStorageNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetMaxStorageNil(b bool)`

 SetMaxStorageNil sets the value for MaxStorage to be an explicit nil

### UnsetMaxStorage
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetMaxStorage()`

UnsetMaxStorage ensures that no value is present for MaxStorage, not even an explicit nil
### GetUsedStorage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### SetUsedStorageNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetUsedStorageNil(b bool)`

 SetUsedStorageNil sets the value for UsedStorage to be an explicit nil

### UnsetUsedStorage
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetUsedStorage()`

UnsetUsedStorage ensures that no value is present for UsedStorage, not even an explicit nil
### GetDiskCount

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetDiskCount() int32`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetDiskCountOk() (*int32, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetDiskCount(v int32)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### SetDiskCountNil

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetDiskCountNil(b bool)`

 SetDiskCountNil sets the value for DiskCount to be an explicit nil

### UnsetDiskCount
`func (o *UpdateStorageServers200ResponseAllOfStorageServer) UnsetDiskCount()`

UnsetDiskCount ensures that no value is present for DiskCount, not even an explicit nil
### GetDateCreated

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroups

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetGroups() []UpdateStorageServers200ResponseAllOfStorageServerGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetGroupsOk() (*[]UpdateStorageServers200ResponseAllOfStorageServerGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetGroups(v []UpdateStorageServers200ResponseAllOfStorageServerGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHostGroups

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetHostGroups() []UpdateStorageServers200ResponseAllOfStorageServerHostGroupsInner`

GetHostGroups returns the HostGroups field if non-nil, zero value otherwise.

### GetHostGroupsOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetHostGroupsOk() (*[]UpdateStorageServers200ResponseAllOfStorageServerHostGroupsInner, bool)`

GetHostGroupsOk returns a tuple with the HostGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroups

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetHostGroups(v []UpdateStorageServers200ResponseAllOfStorageServerHostGroupsInner)`

SetHostGroups sets HostGroups field to given value.

### HasHostGroups

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasHostGroups() bool`

HasHostGroups returns a boolean if a field has been set.

### GetHosts

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetHosts() []UpdateStorageServers200ResponseAllOfStorageServerHostsInner`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetHostsOk() (*[]UpdateStorageServers200ResponseAllOfStorageServerHostsInner, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetHosts(v []UpdateStorageServers200ResponseAllOfStorageServerHostsInner)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetTenants() []UpdateStorageServers200ResponseAllOfStorageServerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetTenantsOk() (*[]UpdateStorageServers200ResponseAllOfStorageServerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetTenants(v []UpdateStorageServers200ResponseAllOfStorageServerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetOwner() UpdateStorageServers200ResponseAllOfStorageServerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetOwnerOk() (*UpdateStorageServers200ResponseAllOfStorageServerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetOwner(v UpdateStorageServers200ResponseAllOfStorageServerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCredential

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetCredential() UpdateStorageServers200ResponseAllOfStorageServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) GetCredentialOk() (*UpdateStorageServers200ResponseAllOfStorageServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) SetCredential(v UpdateStorageServers200ResponseAllOfStorageServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *UpdateStorageServers200ResponseAllOfStorageServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


