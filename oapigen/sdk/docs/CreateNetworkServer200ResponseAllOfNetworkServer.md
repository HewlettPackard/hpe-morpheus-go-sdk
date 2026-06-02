# CreateNetworkServer200ResponseAllOfNetworkServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network Server ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **NullableString** | Description | [optional] 
**Type** | Pointer to [**CreateNetworkServer200ResponseAllOfNetworkServerType**](CreateNetworkServer200ResponseAllOfNetworkServerType.md) |  | [optional] 
**Integration** | Pointer to [**CreateNetworkServer200ResponseAllOfNetworkServerIntegration**](CreateNetworkServer200ResponseAllOfNetworkServerIntegration.md) |  | [optional] 
**Account** | Pointer to [**CreateNetworkServer200ResponseAllOfNetworkServerAccount**](CreateNetworkServer200ResponseAllOfNetworkServerAccount.md) |  | [optional] 
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
**Credential** | Pointer to [**CreateNetworkServer200ResponseAllOfNetworkServerCredential**](CreateNetworkServer200ResponseAllOfNetworkServerCredential.md) |  | [optional] 
**Tenants** | Pointer to [**[]CreateNetworkServer200ResponseAllOfNetworkServerTenantsInner**](CreateNetworkServer200ResponseAllOfNetworkServerTenantsInner.md) |  | [optional] 

## Methods

### NewCreateNetworkServer200ResponseAllOfNetworkServer

`func NewCreateNetworkServer200ResponseAllOfNetworkServer() *CreateNetworkServer200ResponseAllOfNetworkServer`

NewCreateNetworkServer200ResponseAllOfNetworkServer instantiates a new CreateNetworkServer200ResponseAllOfNetworkServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetType() CreateNetworkServer200ResponseAllOfNetworkServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetTypeOk() (*CreateNetworkServer200ResponseAllOfNetworkServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetType(v CreateNetworkServer200ResponseAllOfNetworkServerType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegration

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetIntegration() CreateNetworkServer200ResponseAllOfNetworkServerIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetIntegrationOk() (*CreateNetworkServer200ResponseAllOfNetworkServerIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetIntegration(v CreateNetworkServer200ResponseAllOfNetworkServerIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetAccount

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetAccount() CreateNetworkServer200ResponseAllOfNetworkServerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetAccountOk() (*CreateNetworkServer200ResponseAllOfNetworkServerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetAccount(v CreateNetworkServer200ResponseAllOfNetworkServerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetInternalId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetServiceUrl

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePort

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### SetServicePortNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServicePortNil(b bool)`

 SetServicePortNil sets the value for ServicePort to be an explicit nil

### UnsetServicePort
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetServicePort()`

UnsetServicePort ensures that no value is present for ServicePort, not even an explicit nil
### GetServiceMode

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### SetServiceModeNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceModeNil(b bool)`

 SetServiceModeNil sets the value for ServiceMode to be an explicit nil

### UnsetServiceMode
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceMode()`

UnsetServiceMode ensures that no value is present for ServiceMode, not even an explicit nil
### GetServicePath

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### SetServicePathNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServicePathNil(b bool)`

 SetServicePathNil sets the value for ServicePath to be an explicit nil

### UnsetServicePath
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetServicePath()`

UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
### GetServiceUsername

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### SetServicePasswordHashNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServicePasswordHashNil(b bool)`

 SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil

### UnsetServicePasswordHash
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetServicePasswordHash()`

UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
### GetServiceToken

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetServiceTokenHash

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### SetServiceTokenHashNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetServiceTokenHashNil(b bool)`

 SetServiceTokenHashNil sets the value for ServiceTokenHash to be an explicit nil

### UnsetServiceTokenHash
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceTokenHash()`

UnsetServiceTokenHash ensures that no value is present for ServiceTokenHash, not even an explicit nil
### GetApiPort

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetAdminPort() int32`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetAdminPortOk() (*int32, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetAdminPort(v int32)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetStatus

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetLastSync

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetNextRunDate

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetLastSyncDuration

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetConfig

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### SetNetworkFilterNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetNetworkFilterNil(b bool)`

 SetNetworkFilterNil sets the value for NetworkFilter to be an explicit nil

### UnsetNetworkFilter
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetNetworkFilter()`

UnsetNetworkFilter ensures that no value is present for NetworkFilter, not even an explicit nil
### GetTenantMatch

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### SetTenantMatchNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetTenantMatchNil(b bool)`

 SetTenantMatchNil sets the value for TenantMatch to be an explicit nil

### UnsetTenantMatch
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetTenantMatch()`

UnsetTenantMatch ensures that no value is present for TenantMatch, not even an explicit nil
### GetZoneId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetDateCreated

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVisible

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetCredential

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetCredential() CreateNetworkServer200ResponseAllOfNetworkServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetCredentialOk() (*CreateNetworkServer200ResponseAllOfNetworkServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetCredential(v CreateNetworkServer200ResponseAllOfNetworkServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetTenants() []CreateNetworkServer200ResponseAllOfNetworkServerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) GetTenantsOk() (*[]CreateNetworkServer200ResponseAllOfNetworkServerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) SetTenants(v []CreateNetworkServer200ResponseAllOfNetworkServerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateNetworkServer200ResponseAllOfNetworkServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


