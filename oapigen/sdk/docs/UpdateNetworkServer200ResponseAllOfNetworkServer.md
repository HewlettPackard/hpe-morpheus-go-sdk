# UpdateNetworkServer200ResponseAllOfNetworkServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network Server ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **NullableString** | Description | [optional] 
**Type** | Pointer to [**UpdateNetworkServer200ResponseAllOfNetworkServerType**](UpdateNetworkServer200ResponseAllOfNetworkServerType.md) |  | [optional] 
**Integration** | Pointer to [**UpdateNetworkServer200ResponseAllOfNetworkServerIntegration**](UpdateNetworkServer200ResponseAllOfNetworkServerIntegration.md) |  | [optional] 
**Account** | Pointer to [**UpdateNetworkServer200ResponseAllOfNetworkServerAccount**](UpdateNetworkServer200ResponseAllOfNetworkServerAccount.md) |  | [optional] 
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
**Credential** | Pointer to [**UpdateNetworkServer200ResponseAllOfNetworkServerCredential**](UpdateNetworkServer200ResponseAllOfNetworkServerCredential.md) |  | [optional] 
**Tenants** | Pointer to [**[]UpdateNetworkServer200ResponseAllOfNetworkServerTenantsInner**](UpdateNetworkServer200ResponseAllOfNetworkServerTenantsInner.md) |  | [optional] 

## Methods

### NewUpdateNetworkServer200ResponseAllOfNetworkServer

`func NewUpdateNetworkServer200ResponseAllOfNetworkServer() *UpdateNetworkServer200ResponseAllOfNetworkServer`

NewUpdateNetworkServer200ResponseAllOfNetworkServer instantiates a new UpdateNetworkServer200ResponseAllOfNetworkServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkServer200ResponseAllOfNetworkServerWithDefaults

`func NewUpdateNetworkServer200ResponseAllOfNetworkServerWithDefaults() *UpdateNetworkServer200ResponseAllOfNetworkServer`

NewUpdateNetworkServer200ResponseAllOfNetworkServerWithDefaults instantiates a new UpdateNetworkServer200ResponseAllOfNetworkServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetType() UpdateNetworkServer200ResponseAllOfNetworkServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetTypeOk() (*UpdateNetworkServer200ResponseAllOfNetworkServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetType(v UpdateNetworkServer200ResponseAllOfNetworkServerType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegration

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetIntegration() UpdateNetworkServer200ResponseAllOfNetworkServerIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetIntegrationOk() (*UpdateNetworkServer200ResponseAllOfNetworkServerIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetIntegration(v UpdateNetworkServer200ResponseAllOfNetworkServerIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetAccount() UpdateNetworkServer200ResponseAllOfNetworkServerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetAccountOk() (*UpdateNetworkServer200ResponseAllOfNetworkServerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetAccount(v UpdateNetworkServer200ResponseAllOfNetworkServerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetInternalId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetServiceUrl

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePort

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### SetServicePortNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServicePortNil(b bool)`

 SetServicePortNil sets the value for ServicePort to be an explicit nil

### UnsetServicePort
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetServicePort()`

UnsetServicePort ensures that no value is present for ServicePort, not even an explicit nil
### GetServiceMode

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### SetServiceModeNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceModeNil(b bool)`

 SetServiceModeNil sets the value for ServiceMode to be an explicit nil

### UnsetServiceMode
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceMode()`

UnsetServiceMode ensures that no value is present for ServiceMode, not even an explicit nil
### GetServicePath

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### SetServicePathNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServicePathNil(b bool)`

 SetServicePathNil sets the value for ServicePath to be an explicit nil

### UnsetServicePath
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetServicePath()`

UnsetServicePath ensures that no value is present for ServicePath, not even an explicit nil
### GetServiceUsername

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### SetServicePasswordHashNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServicePasswordHashNil(b bool)`

 SetServicePasswordHashNil sets the value for ServicePasswordHash to be an explicit nil

### UnsetServicePasswordHash
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetServicePasswordHash()`

UnsetServicePasswordHash ensures that no value is present for ServicePasswordHash, not even an explicit nil
### GetServiceToken

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### SetServiceTokenNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceTokenNil(b bool)`

 SetServiceTokenNil sets the value for ServiceToken to be an explicit nil

### UnsetServiceToken
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceToken()`

UnsetServiceToken ensures that no value is present for ServiceToken, not even an explicit nil
### GetServiceTokenHash

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### SetServiceTokenHashNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetServiceTokenHashNil(b bool)`

 SetServiceTokenHashNil sets the value for ServiceTokenHash to be an explicit nil

### UnsetServiceTokenHash
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetServiceTokenHash()`

UnsetServiceTokenHash ensures that no value is present for ServiceTokenHash, not even an explicit nil
### GetApiPort

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### SetApiPortNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetApiPortNil(b bool)`

 SetApiPortNil sets the value for ApiPort to be an explicit nil

### UnsetApiPort
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetApiPort()`

UnsetApiPort ensures that no value is present for ApiPort, not even an explicit nil
### GetAdminPort

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetAdminPort() int32`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetAdminPortOk() (*int32, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetAdminPort(v int32)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### SetAdminPortNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetAdminPortNil(b bool)`

 SetAdminPortNil sets the value for AdminPort to be an explicit nil

### UnsetAdminPort
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetAdminPort()`

UnsetAdminPort ensures that no value is present for AdminPort, not even an explicit nil
### GetStatus

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetLastSync

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetNextRunDate

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetLastSyncDuration

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetConfig

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### SetNetworkFilterNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetNetworkFilterNil(b bool)`

 SetNetworkFilterNil sets the value for NetworkFilter to be an explicit nil

### UnsetNetworkFilter
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetNetworkFilter()`

UnsetNetworkFilter ensures that no value is present for NetworkFilter, not even an explicit nil
### GetTenantMatch

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### SetTenantMatchNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetTenantMatchNil(b bool)`

 SetTenantMatchNil sets the value for TenantMatch to be an explicit nil

### UnsetTenantMatch
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetTenantMatch()`

UnsetTenantMatch ensures that no value is present for TenantMatch, not even an explicit nil
### GetZoneId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetDateCreated

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVisible

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetCredential

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetCredential() UpdateNetworkServer200ResponseAllOfNetworkServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetCredentialOk() (*UpdateNetworkServer200ResponseAllOfNetworkServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetCredential(v UpdateNetworkServer200ResponseAllOfNetworkServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetTenants() []UpdateNetworkServer200ResponseAllOfNetworkServerTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) GetTenantsOk() (*[]UpdateNetworkServer200ResponseAllOfNetworkServerTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) SetTenants(v []UpdateNetworkServer200ResponseAllOfNetworkServerTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateNetworkServer200ResponseAllOfNetworkServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


