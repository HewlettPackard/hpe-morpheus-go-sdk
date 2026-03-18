# GetNetworkServer200ResponseNetworkServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network Server ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **NullableString** | Description | [optional] 
**Type** | Pointer to [**GetNetworkServer200ResponseNetworkServerType**](GetNetworkServer200ResponseNetworkServerType.md) |  | [optional] 
**Integration** | Pointer to [**GetNetworkServer200ResponseNetworkServerIntegration**](GetNetworkServer200ResponseNetworkServerIntegration.md) |  | [optional] 
**Account** | Pointer to [**GetNetworkServer200ResponseNetworkServerAccount**](GetNetworkServer200ResponseNetworkServerAccount.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **NullableString** | Internal ID | [optional] 
**ExternalId** | Pointer to **NullableString** | External ID | [optional] 
**ServiceUrl** | Pointer to **NullableString** | Service URL | [optional] 
**ServiceHost** | Pointer to **NullableString** | Service Host | [optional] 
**ServicePort** | Pointer to **NullableInt32** | Service Port | [optional] 
**ServiceMode** | Pointer to **NullableString** | Service Mode | [optional] 
**ServicePath** | Pointer to **NullableString** | Service Path | [optional] 
**ServiceUsername** | Pointer to **NullableString** | Service Username | [optional] 
**ServicePassword** | Pointer to **NullableString** | Service Password | [optional] 
**ServicePasswordHash** | Pointer to **NullableString** |  | [optional] 
**ServiceToken** | Pointer to **NullableString** | Service Token | [optional] 
**ServiceTokenHash** | Pointer to **NullableString** |  | [optional] 
**ApiPort** | Pointer to **NullableInt32** |  | [optional] 
**AdminPort** | Pointer to **NullableInt32** |  | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**StatusMessage** | Pointer to **NullableString** | Status Message | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**LastSync** | Pointer to **NullableTime** | Last Sync Date | [optional] 
**NextRunDate** | Pointer to **NullableTime** | Next Run Date | [optional] 
**LastSyncDuration** | Pointer to **NullableInt64** | Last Sync Duration in milliseconds | [optional] 
**Config** | Pointer to **map[string]interface{}** | Config object varies with network server type. | [optional] 
**NetworkFilter** | Pointer to **NullableString** | Network Filter | [optional] 
**TenantMatch** | Pointer to **NullableString** | Tenant Match | [optional] 
**ZoneId** | Pointer to **NullableInt64** | Cloud ID | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 
**Credential** | Pointer to [**GetNetworkServer200ResponseNetworkServerCredential**](GetNetworkServer200ResponseNetworkServerCredential.md) |  | [optional] 
**Tenants** | Pointer to [**[]GetNetworkServer200ResponseNetworkServerTenantsInner**](GetNetworkServer200ResponseNetworkServerTenantsInner.md) |  | [optional] 

## Methods

### NewGetNetworkServer200ResponseNetworkServer

`func NewGetNetworkServer200ResponseNetworkServer() *GetNetworkServer200ResponseNetworkServer`

NewGetNetworkServer200ResponseNetworkServer instantiates a new GetNetworkServer200ResponseNetworkServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkServer200ResponseNetworkServerWithDefaults

`func NewGetNetworkServer200ResponseNetworkServerWithDefaults() *GetNetworkServer200ResponseNetworkServer`

NewGetNetworkServer200ResponseNetworkServerWithDefaults instantiates a new GetNetworkServer200ResponseNetworkServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkServer200ResponseNetworkServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkServer200ResponseNetworkServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkServer200ResponseNetworkServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkServer200ResponseNetworkServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkServer200ResponseNetworkServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkServer200ResponseNetworkServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkServer200ResponseNetworkServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkServer200ResponseNetworkServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkServer200ResponseNetworkServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *GetNetworkServer200ResponseNetworkServer) GetType() GetNetworkServer200ResponseNetworkServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetTypeOk() (*GetNetworkServer200ResponseNetworkServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkServer200ResponseNetworkServer) SetType(v GetNetworkServer200ResponseNetworkServerType)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkServer200ResponseNetworkServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegration

`func (o *GetNetworkServer200ResponseNetworkServer) GetIntegration() GetNetworkServer200ResponseNetworkServerIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetIntegrationOk() (*GetNetworkServer200ResponseNetworkServerIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *GetNetworkServer200ResponseNetworkServer) SetIntegration(v GetNetworkServer200ResponseNetworkServerIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *GetNetworkServer200ResponseNetworkServer) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetAccount

`func (o *GetNetworkServer200ResponseNetworkServer) GetAccount() GetNetworkServer200ResponseNetworkServerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetAccountOk() (*GetNetworkServer200ResponseNetworkServerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetNetworkServer200ResponseNetworkServer) SetAccount(v GetNetworkServer200ResponseNetworkServerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetNetworkServer200ResponseNetworkServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetVisibility

`func (o *GetNetworkServer200ResponseNetworkServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetNetworkServer200ResponseNetworkServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetNetworkServer200ResponseNetworkServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetInternalId

`func (o *GetNetworkServer200ResponseNetworkServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkServer200ResponseNetworkServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkServer200ResponseNetworkServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *GetNetworkServer200ResponseNetworkServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkServer200ResponseNetworkServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkServer200ResponseNetworkServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetServiceUrl

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *GetNetworkServer200ResponseNetworkServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *GetNetworkServer200ResponseNetworkServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePort

`func (o *GetNetworkServer200ResponseNetworkServer) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *GetNetworkServer200ResponseNetworkServer) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *GetNetworkServer200ResponseNetworkServer) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### SetServicePortNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetServicePortNil(b bool)`

 SetServicePortNil sets the value for ServicePort to be an explicit nil

### UnsetServicePort
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetServicePort()`

UnsetServicePort ensures that no value is present for ServicePort, not even an explicit nil
### GetServiceMode

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *GetNetworkServer200ResponseNetworkServer) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### SetServiceModeNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceModeNil(b bool)`

 SetServiceModeNil sets the value for ServiceMode to be an explicit nil

### UnsetServiceMode
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetServiceMode()`

UnsetServiceMode ensures that no value is present for ServiceMode, not even an explicit nil
### GetServicePath

`func (o *GetNetworkServer200ResponseNetworkServer) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *GetNetworkServer200ResponseNetworkServer) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *GetNetworkServer200ResponseNetworkServer) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### SetServicePathNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetServicePathNil(b bool)`

 SetServicePathNil sets the value for ServicePath to be an explicit nil

### UnsetServicePath
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetServicePath()`

UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
### GetServiceUsername

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *GetNetworkServer200ResponseNetworkServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *GetNetworkServer200ResponseNetworkServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *GetNetworkServer200ResponseNetworkServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *GetNetworkServer200ResponseNetworkServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *GetNetworkServer200ResponseNetworkServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *GetNetworkServer200ResponseNetworkServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *GetNetworkServer200ResponseNetworkServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### SetServicePasswordHashNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetServicePasswordHashNil(b bool)`

 SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil

### UnsetServicePasswordHash
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetServicePasswordHash()`

UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
### GetServiceToken

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *GetNetworkServer200ResponseNetworkServer) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetServiceTokenHash

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *GetNetworkServer200ResponseNetworkServer) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### SetServiceTokenHashNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetServiceTokenHashNil(b bool)`

 SetServiceTokenHashNil sets the value for ServiceTokenHash to be an explicit nil

### UnsetServiceTokenHash
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetServiceTokenHash()`

UnsetServiceTokenHash ensures that no value is present for ServiceTokenHash, not even an explicit nil
### GetApiPort

`func (o *GetNetworkServer200ResponseNetworkServer) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *GetNetworkServer200ResponseNetworkServer) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *GetNetworkServer200ResponseNetworkServer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *GetNetworkServer200ResponseNetworkServer) GetAdminPort() int32`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetAdminPortOk() (*int32, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *GetNetworkServer200ResponseNetworkServer) SetAdminPort(v int32)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *GetNetworkServer200ResponseNetworkServer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetStatus

`func (o *GetNetworkServer200ResponseNetworkServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkServer200ResponseNetworkServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkServer200ResponseNetworkServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetNetworkServer200ResponseNetworkServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetNetworkServer200ResponseNetworkServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetNetworkServer200ResponseNetworkServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *GetNetworkServer200ResponseNetworkServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetNetworkServer200ResponseNetworkServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetNetworkServer200ResponseNetworkServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetLastSync

`func (o *GetNetworkServer200ResponseNetworkServer) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *GetNetworkServer200ResponseNetworkServer) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *GetNetworkServer200ResponseNetworkServer) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetNextRunDate

`func (o *GetNetworkServer200ResponseNetworkServer) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *GetNetworkServer200ResponseNetworkServer) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *GetNetworkServer200ResponseNetworkServer) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetLastSyncDuration

`func (o *GetNetworkServer200ResponseNetworkServer) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *GetNetworkServer200ResponseNetworkServer) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *GetNetworkServer200ResponseNetworkServer) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetConfig

`func (o *GetNetworkServer200ResponseNetworkServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkServer200ResponseNetworkServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkServer200ResponseNetworkServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *GetNetworkServer200ResponseNetworkServer) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *GetNetworkServer200ResponseNetworkServer) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *GetNetworkServer200ResponseNetworkServer) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### SetNetworkFilterNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetNetworkFilterNil(b bool)`

 SetNetworkFilterNil sets the value for NetworkFilter to be an explicit nil

### UnsetNetworkFilter
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetNetworkFilter()`

UnsetNetworkFilter ensures that no value is present for NetworkFilter, not even an explicit nil
### GetTenantMatch

`func (o *GetNetworkServer200ResponseNetworkServer) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *GetNetworkServer200ResponseNetworkServer) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *GetNetworkServer200ResponseNetworkServer) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### SetTenantMatchNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetTenantMatchNil(b bool)`

 SetTenantMatchNil sets the value for TenantMatch to be an explicit nil

### UnsetTenantMatch
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetTenantMatch()`

UnsetTenantMatch ensures that no value is present for TenantMatch, not even an explicit nil
### GetZoneId

`func (o *GetNetworkServer200ResponseNetworkServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetNetworkServer200ResponseNetworkServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetNetworkServer200ResponseNetworkServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *GetNetworkServer200ResponseNetworkServer) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *GetNetworkServer200ResponseNetworkServer) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetDateCreated

`func (o *GetNetworkServer200ResponseNetworkServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkServer200ResponseNetworkServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkServer200ResponseNetworkServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkServer200ResponseNetworkServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkServer200ResponseNetworkServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkServer200ResponseNetworkServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkServer200ResponseNetworkServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkServer200ResponseNetworkServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkServer200ResponseNetworkServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVisible

`func (o *GetNetworkServer200ResponseNetworkServer) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *GetNetworkServer200ResponseNetworkServer) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *GetNetworkServer200ResponseNetworkServer) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetCredential

`func (o *GetNetworkServer200ResponseNetworkServer) GetCredential() GetNetworkServer200ResponseNetworkServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetCredentialOk() (*GetNetworkServer200ResponseNetworkServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetNetworkServer200ResponseNetworkServer) SetCredential(v GetNetworkServer200ResponseNetworkServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetNetworkServer200ResponseNetworkServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *GetNetworkServer200ResponseNetworkServer) GetTenants() []GetNetworkServer200ResponseNetworkServerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetNetworkServer200ResponseNetworkServer) GetTenantsOk() (*[]GetNetworkServer200ResponseNetworkServerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetNetworkServer200ResponseNetworkServer) SetTenants(v []GetNetworkServer200ResponseNetworkServerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetNetworkServer200ResponseNetworkServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


