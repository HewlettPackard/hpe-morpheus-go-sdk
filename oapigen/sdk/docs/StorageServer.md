# StorageServer

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

### NewStorageServer

`func NewStorageServer() *StorageServer`

NewStorageServer instantiates a new StorageServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageServerWithDefaults

`func NewStorageServerWithDefaults() *StorageServer`

NewStorageServerWithDefaults instantiates a new StorageServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StorageServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StorageServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *StorageServer) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageServer) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageServer) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *StorageServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChassis

`func (o *StorageServer) GetChassis() string`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *StorageServer) GetChassisOk() (*string, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *StorageServer) SetChassis(v string)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *StorageServer) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### SetChassisNil

`func (o *StorageServer) SetChassisNil(b bool)`

 SetChassisNil sets the value for Chassis to be an explicit nil

### UnsetChassis
`func (o *StorageServer) UnsetChassis()`

UnsetChassis ensures that no value is present for Chassis, not even an explicit nil
### GetVisibility

`func (o *StorageServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *StorageServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *StorageServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *StorageServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *StorageServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StorageServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StorageServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *StorageServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *StorageServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *StorageServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *StorageServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *StorageServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *StorageServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *StorageServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *StorageServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *StorageServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *StorageServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *StorageServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *StorageServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetServiceUrl

`func (o *StorageServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *StorageServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *StorageServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *StorageServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *StorageServer) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *StorageServer) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *StorageServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *StorageServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *StorageServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *StorageServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *StorageServer) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *StorageServer) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePath

`func (o *StorageServer) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *StorageServer) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *StorageServer) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *StorageServer) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### SetServicePathNil

`func (o *StorageServer) SetServicePathNil(b bool)`

 SetServicePathNil sets the value for ServicePath to be an explicit nil

### UnsetServicePath
`func (o *StorageServer) UnsetServicePath()`

UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
### GetServiceToken

`func (o *StorageServer) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *StorageServer) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *StorageServer) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *StorageServer) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *StorageServer) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *StorageServer) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetServiceTokenHash

`func (o *StorageServer) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *StorageServer) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *StorageServer) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *StorageServer) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### SetServiceTokenHashNil

`func (o *StorageServer) SetServiceTokenHashNil(b bool)`

 SetServiceTokenHashNil sets the value for ServiceTokenHash to be an explicit nil

### UnsetServiceTokenHash
`func (o *StorageServer) UnsetServiceTokenHash()`

UnsetServiceTokenHash ensures that no value is present for ServiceTokenHash, not even an explicit nil
### GetServiceVersion

`func (o *StorageServer) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *StorageServer) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *StorageServer) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *StorageServer) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### SetServiceVersionNil

`func (o *StorageServer) SetServiceVersionNil(b bool)`

 SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil

### UnsetServiceVersion
`func (o *StorageServer) UnsetServiceVersion()`

UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
### GetServiceUsername

`func (o *StorageServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *StorageServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *StorageServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *StorageServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *StorageServer) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *StorageServer) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *StorageServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *StorageServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *StorageServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *StorageServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *StorageServer) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *StorageServer) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *StorageServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *StorageServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *StorageServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *StorageServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### SetServicePasswordHashNil

`func (o *StorageServer) SetServicePasswordHashNil(b bool)`

 SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil

### UnsetServicePasswordHash
`func (o *StorageServer) UnsetServicePasswordHash()`

UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
### GetInternalIp

`func (o *StorageServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *StorageServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *StorageServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *StorageServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *StorageServer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *StorageServer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetExternalIp

`func (o *StorageServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *StorageServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *StorageServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *StorageServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *StorageServer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *StorageServer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetApiPort

`func (o *StorageServer) GetApiPort() string`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *StorageServer) GetApiPortOk() (*string, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *StorageServer) SetApiPort(v string)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *StorageServer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *StorageServer) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *StorageServer) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *StorageServer) GetAdminPort() string`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *StorageServer) GetAdminPortOk() (*string, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *StorageServer) SetAdminPort(v string)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *StorageServer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *StorageServer) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *StorageServer) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetConfig

`func (o *StorageServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StorageServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StorageServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *StorageServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRefType

`func (o *StorageServer) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *StorageServer) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *StorageServer) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *StorageServer) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *StorageServer) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *StorageServer) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *StorageServer) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *StorageServer) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetCategory

`func (o *StorageServer) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StorageServer) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StorageServer) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *StorageServer) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *StorageServer) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *StorageServer) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetServerVendor

`func (o *StorageServer) GetServerVendor() string`

GetServerVendor returns the ServerVendor field if non-nil, zero value otherwise.

### GetServerVendorOk

`func (o *StorageServer) GetServerVendorOk() (*string, bool)`

GetServerVendorOk returns a tuple with the ServerVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVendor

`func (o *StorageServer) SetServerVendor(v string)`

SetServerVendor sets ServerVendor field to given value.

### HasServerVendor

`func (o *StorageServer) HasServerVendor() bool`

HasServerVendor returns a boolean if a field has been set.

### SetServerVendorNil

`func (o *StorageServer) SetServerVendorNil(b bool)`

 SetServerVendorNil sets the value for ServerVendor to be an explicit nil

### UnsetServerVendor
`func (o *StorageServer) UnsetServerVendor()`

UnsetServerVendor ensures that no value is present for ServerVendor, not even an explicit nil
### GetServerModel

`func (o *StorageServer) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *StorageServer) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *StorageServer) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *StorageServer) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### SetServerModelNil

`func (o *StorageServer) SetServerModelNil(b bool)`

 SetServerModelNil sets the value for ServerModel to be an explicit nil

### UnsetServerModel
`func (o *StorageServer) UnsetServerModel()`

UnsetServerModel ensures that no value is present for ServerModel, not even an explicit nil
### GetSerialNumber

`func (o *StorageServer) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *StorageServer) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *StorageServer) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *StorageServer) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *StorageServer) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *StorageServer) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetStatus

`func (o *StorageServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *StorageServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *StorageServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *StorageServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *StorageServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *StorageServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *StorageServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *StorageServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *StorageServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *StorageServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *StorageServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetErrorMessage

`func (o *StorageServer) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *StorageServer) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *StorageServer) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *StorageServer) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *StorageServer) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *StorageServer) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetMaxStorage

`func (o *StorageServer) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *StorageServer) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *StorageServer) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *StorageServer) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### SetMaxStorageNil

`func (o *StorageServer) SetMaxStorageNil(b bool)`

 SetMaxStorageNil sets the value for MaxStorage to be an explicit nil

### UnsetMaxStorage
`func (o *StorageServer) UnsetMaxStorage()`

UnsetMaxStorage ensures that no value is present for MaxStorage, not even an explicit nil
### GetUsedStorage

`func (o *StorageServer) GetUsedStorage() string`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *StorageServer) GetUsedStorageOk() (*string, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *StorageServer) SetUsedStorage(v string)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *StorageServer) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### SetUsedStorageNil

`func (o *StorageServer) SetUsedStorageNil(b bool)`

 SetUsedStorageNil sets the value for UsedStorage to be an explicit nil

### UnsetUsedStorage
`func (o *StorageServer) UnsetUsedStorage()`

UnsetUsedStorage ensures that no value is present for UsedStorage, not even an explicit nil
### GetDiskCount

`func (o *StorageServer) GetDiskCount() string`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *StorageServer) GetDiskCountOk() (*string, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *StorageServer) SetDiskCount(v string)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *StorageServer) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### SetDiskCountNil

`func (o *StorageServer) SetDiskCountNil(b bool)`

 SetDiskCountNil sets the value for DiskCount to be an explicit nil

### UnsetDiskCount
`func (o *StorageServer) UnsetDiskCount()`

UnsetDiskCount ensures that no value is present for DiskCount, not even an explicit nil
### GetDateCreated

`func (o *StorageServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *StorageServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *StorageServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *StorageServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *StorageServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *StorageServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *StorageServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *StorageServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroups

`func (o *StorageServer) GetGroups() []map[string]interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *StorageServer) GetGroupsOk() (*[]map[string]interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *StorageServer) SetGroups(v []map[string]interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *StorageServer) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHostGroups

`func (o *StorageServer) GetHostGroups() []map[string]interface{}`

GetHostGroups returns the HostGroups field if non-nil, zero value otherwise.

### GetHostGroupsOk

`func (o *StorageServer) GetHostGroupsOk() (*[]map[string]interface{}, bool)`

GetHostGroupsOk returns a tuple with the HostGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroups

`func (o *StorageServer) SetHostGroups(v []map[string]interface{})`

SetHostGroups sets HostGroups field to given value.

### HasHostGroups

`func (o *StorageServer) HasHostGroups() bool`

HasHostGroups returns a boolean if a field has been set.

### GetHosts

`func (o *StorageServer) GetHosts() []map[string]interface{}`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *StorageServer) GetHostsOk() (*[]map[string]interface{}, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *StorageServer) SetHosts(v []map[string]interface{})`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *StorageServer) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetTenants

`func (o *StorageServer) GetTenants() []map[string]interface{}`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *StorageServer) GetTenantsOk() (*[]map[string]interface{}, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *StorageServer) SetTenants(v []map[string]interface{})`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *StorageServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetOwner

`func (o *StorageServer) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *StorageServer) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *StorageServer) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *StorageServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCredential

`func (o *StorageServer) GetCredential() map[string]interface{}`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *StorageServer) GetCredentialOk() (*map[string]interface{}, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *StorageServer) SetCredential(v map[string]interface{})`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *StorageServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


