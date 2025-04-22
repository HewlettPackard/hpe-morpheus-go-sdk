# ListStorageServers200ResponseAllOfStorageServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Chassis** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ServiceUrl** | Pointer to **string** |  | [optional] 
**ServiceHost** | Pointer to **string** |  | [optional] 
**ServicePath** | Pointer to **string** |  | [optional] 
**ServiceToken** | Pointer to **string** |  | [optional] 
**ServiceTokenHash** | Pointer to **string** |  | [optional] 
**ServiceVersion** | Pointer to **string** |  | [optional] 
**ServiceUsername** | Pointer to **string** |  | [optional] 
**ServicePassword** | Pointer to **string** |  | [optional] 
**ServicePasswordHash** | Pointer to **string** |  | [optional] 
**InternalIp** | Pointer to **string** |  | [optional] 
**ExternalIp** | Pointer to **string** |  | [optional] 
**ApiPort** | Pointer to **string** |  | [optional] 
**AdminPort** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ServerVendor** | Pointer to **string** |  | [optional] 
**ServerModel** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **string** |  | [optional] 
**UsedStorage** | Pointer to **string** |  | [optional] 
**DiskCount** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Groups** | Pointer to **[]map[string]interface{}** |  | [optional] 
**HostGroups** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Hosts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tenants** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Owner** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Credential** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewListStorageServers200ResponseAllOfStorageServersInner

`func NewListStorageServers200ResponseAllOfStorageServersInner() *ListStorageServers200ResponseAllOfStorageServersInner`

NewListStorageServers200ResponseAllOfStorageServersInner instantiates a new ListStorageServers200ResponseAllOfStorageServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageServers200ResponseAllOfStorageServersInnerWithDefaults

`func NewListStorageServers200ResponseAllOfStorageServersInnerWithDefaults() *ListStorageServers200ResponseAllOfStorageServersInner`

NewListStorageServers200ResponseAllOfStorageServersInnerWithDefaults instantiates a new ListStorageServers200ResponseAllOfStorageServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChassis

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetChassis() string`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetChassisOk() (*string, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetChassis(v string)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetVisibility

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetServiceUrl

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceHost

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### GetServicePath

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### GetServiceToken

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceTokenHash

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### GetServiceVersion

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetServiceUsername

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServicePasswordHash

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### GetInternalIp

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### GetExternalIp

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### GetApiPort

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetApiPort() string`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetApiPortOk() (*string, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetApiPort(v string)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetAdminPort

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetAdminPort() string`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetAdminPortOk() (*string, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetAdminPort(v string)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### GetConfig

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetRefType

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetCategory

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetServerVendor

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServerVendor() string`

GetServerVendor returns the ServerVendor field if non-nil, zero value otherwise.

### GetServerVendorOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServerVendorOk() (*string, bool)`

GetServerVendorOk returns a tuple with the ServerVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVendor

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServerVendor(v string)`

SetServerVendor sets ServerVendor field to given value.

### HasServerVendor

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServerVendor() bool`

HasServerVendor returns a boolean if a field has been set.

### GetServerModel

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetStatus

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetUsedStorage() string`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetUsedStorageOk() (*string, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetUsedStorage(v string)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetDiskCount

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetDiskCount() string`

GetDiskCount returns the DiskCount field if non-nil, zero value otherwise.

### GetDiskCountOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetDiskCountOk() (*string, bool)`

GetDiskCountOk returns a tuple with the DiskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCount

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetDiskCount(v string)`

SetDiskCount sets DiskCount field to given value.

### HasDiskCount

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasDiskCount() bool`

HasDiskCount returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEnabled

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroups

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetGroups() []map[string]interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetGroupsOk() (*[]map[string]interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetGroups(v []map[string]interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHostGroups

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetHostGroups() []map[string]interface{}`

GetHostGroups returns the HostGroups field if non-nil, zero value otherwise.

### GetHostGroupsOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetHostGroupsOk() (*[]map[string]interface{}, bool)`

GetHostGroupsOk returns a tuple with the HostGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostGroups

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetHostGroups(v []map[string]interface{})`

SetHostGroups sets HostGroups field to given value.

### HasHostGroups

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasHostGroups() bool`

HasHostGroups returns a boolean if a field has been set.

### GetHosts

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetHosts() []map[string]interface{}`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetHostsOk() (*[]map[string]interface{}, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetHosts(v []map[string]interface{})`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetTenants

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetTenants() []map[string]interface{}`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetTenantsOk() (*[]map[string]interface{}, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetTenants(v []map[string]interface{})`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetOwner

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCredential

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetCredential() map[string]interface{}`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) GetCredentialOk() (*map[string]interface{}, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) SetCredential(v map[string]interface{})`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListStorageServers200ResponseAllOfStorageServersInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


