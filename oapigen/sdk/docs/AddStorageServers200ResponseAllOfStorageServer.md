# AddStorageServers200ResponseAllOfStorageServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
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
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Credential** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAddStorageServers200ResponseAllOfStorageServer

`func NewAddStorageServers200ResponseAllOfStorageServer() *AddStorageServers200ResponseAllOfStorageServer`

NewAddStorageServers200ResponseAllOfStorageServer instantiates a new AddStorageServers200ResponseAllOfStorageServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageServers200ResponseAllOfStorageServerWithDefaults

`func NewAddStorageServers200ResponseAllOfStorageServerWithDefaults() *AddStorageServers200ResponseAllOfStorageServer`

NewAddStorageServers200ResponseAllOfStorageServerWithDefaults instantiates a new AddStorageServers200ResponseAllOfStorageServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChassis

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetChassis() string`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetChassisOk() (*string, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetChassis(v string)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### SetChassisNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetChassisNil(b bool)`

 SetChassisNil sets the value for Chassis to be an explicit nil

### UnsetChassis
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetChassis()`

UnsetChassis ensures that no value is present for Chassis, not even an explicit nil
### GetVisibility

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetServiceUrl

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePath

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### SetServicePathNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServicePathNil(b bool)`

 SetServicePathNil sets the value for ServicePath to be an explicit nil

### UnsetServicePath
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServicePath()`

UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
### GetServiceToken

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetServiceTokenHash

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### SetServiceTokenHashNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceTokenHashNil(b bool)`

 SetServiceTokenHashNil sets the value for ServiceTokenHash to be an explicit nil

### UnsetServiceTokenHash
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServiceTokenHash()`

UnsetServiceTokenHash ensures that no value is present for ServiceTokenHash, not even an explicit nil
### GetServiceVersion

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### SetServiceVersionNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceVersionNil(b bool)`

 SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil

### UnsetServiceVersion
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServiceVersion()`

UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
### GetServiceUsername

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### SetServicePasswordHashNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServicePasswordHashNil(b bool)`

 SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil

### UnsetServicePasswordHash
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServicePasswordHash()`

UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
### GetInternalIp

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetExternalIp

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetApiPort

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetApiPort() string`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetApiPortOk() (*string, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetApiPort(v string)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetAdminPort() string`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetAdminPortOk() (*string, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetAdminPort(v string)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetConfig

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRefType

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetCategory

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetServerVendor

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServerVendor() string`

GetServerVendor returns the ServerVendor field if non-nil, zero value otherwise.

### GetServerVendorOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServerVendorOk() (*string, bool)`

GetServerVendorOk returns a tuple with the ServerVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVendor

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServerVendor(v string)`

SetServerVendor sets ServerVendor field to given value.

### HasServerVendor

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServerVendor() bool`

HasServerVendor returns a boolean if a field has been set.

### SetServerVendorNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServerVendorNil(b bool)`

 SetServerVendorNil sets the value for ServerVendor to be an explicit nil

### UnsetServerVendor
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServerVendor()`

UnsetServerVendor ensures that no value is present for ServerVendor, not even an explicit nil
### GetServerModel

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### SetServerModelNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetServerModelNil(b bool)`

 SetServerModelNil sets the value for ServerModel to be an explicit nil

### UnsetServerModel
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetServerModel()`

UnsetServerModel ensures that no value is present for ServerModel, not even an explicit nil
### GetSerialNumber

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetStatus

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetErrorMessage

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetMaxStorage

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### SetMaxStorageNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetMaxStorageNil(b bool)`

 SetMaxStorageNil sets the value for MaxStorage to be an explicit nil

### UnsetMaxStorage
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetMaxStorage()`

UnsetMaxStorage ensures that no value is present for MaxStorage, not even an explicit nil
### GetUsedStorage

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetUsedStorage() string`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetUsedStorageOk() (*string, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetUsedStorage(v string)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### SetUsedStorageNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetUsedStorageNil(b bool)`

 SetUsedStorageNil sets the value for UsedStorage to be an explicit nil

### UnsetUsedStorage
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetUsedStorage()`

UnsetUsedStorage ensures that no value is present for UsedStorage, not even an explicit nil
### GetDiskCount

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetDiskCount() string`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetDiskCountOk() (*string, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetDiskCount(v string)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### SetDiskCountNil

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetDiskCountNil(b bool)`

 SetDiskCountNil sets the value for DiskCount to be an explicit nil

### UnsetDiskCount
`func (o *AddStorageServers200ResponseAllOfStorageServer) UnsetDiskCount()`

UnsetDiskCount ensures that no value is present for DiskCount, not even an explicit nil
### GetDateCreated

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEnabled

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroups

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetGroups() []map[string]interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetGroupsOk() (*[]map[string]interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetGroups(v []map[string]interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHostGroups

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetHostGroups() []map[string]interface{}`

GetHostGroups returns the HostGroups field if non-nil, zero value otherwise.

### GetHostGroupsOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetHostGroupsOk() (*[]map[string]interface{}, bool)`

GetHostGroupsOk returns a tuple with the HostGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroups

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetHostGroups(v []map[string]interface{})`

SetHostGroups sets HostGroups field to given value.

### HasHostGroups

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasHostGroups() bool`

HasHostGroups returns a boolean if a field has been set.

### GetHosts

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetHosts() []map[string]interface{}`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetHostsOk() (*[]map[string]interface{}, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetHosts(v []map[string]interface{})`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetTenants

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetTenants() []map[string]interface{}`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetTenantsOk() (*[]map[string]interface{}, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetTenants(v []map[string]interface{})`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetOwner

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCredential

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetCredential() map[string]interface{}`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AddStorageServers200ResponseAllOfStorageServer) GetCredentialOk() (*map[string]interface{}, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AddStorageServers200ResponseAllOfStorageServer) SetCredential(v map[string]interface{})`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AddStorageServers200ResponseAllOfStorageServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


