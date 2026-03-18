# UpdateIntegrationsRequestOneOf3Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name, a unique identifier for the integration | 
**Type** | **string** | Integration Type Code | 
**ServiceMode** | Pointer to **string** | Topology | [optional] [default to "single"]
**ServiceUrl** | **string** | Salt Master | 
**Secondary** | Pointer to **string** | Salt Syndic | [optional] 
**ServicePort** | Pointer to **int32** | SSH Port | [optional] [default to 22]
**ServiceUsername** | **string** | Username | 
**ServicePassword** | Pointer to **string** | Password | [optional] 
**ServiceKey** | Pointer to **string** | Master Key Pair | [optional] 
**AuthKey** | Pointer to **string** | Signing Key | [optional] 
**ServicePath** | Pointer to **string** | Working Directory | [optional] 
**ServiceVersion** | Pointer to **string** | Salt Version | [optional] 
**ServiceWindowsVersion** | Pointer to **string** | Salt Version (Windows) | [optional] 
**RepoUrl** | Pointer to **string** | Repo URL | [optional] 
**ServiceConfig** | Pointer to **string** | Minion Config | [optional] 
**ServiceCommand** | Pointer to **string** | Post Provision Commands | [optional] 
**Config** | Pointer to [**UpdateIntegrationsRequestOneOf3IntegrationConfig**](UpdateIntegrationsRequestOneOf3IntegrationConfig.md) |  | [optional] 

## Methods

### NewUpdateIntegrationsRequestOneOf3Integration

`func NewUpdateIntegrationsRequestOneOf3Integration(name string, type_ string, serviceUrl string, serviceUsername string, ) *UpdateIntegrationsRequestOneOf3Integration`

NewUpdateIntegrationsRequestOneOf3Integration instantiates a new UpdateIntegrationsRequestOneOf3Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationsRequestOneOf3IntegrationWithDefaults

`func NewUpdateIntegrationsRequestOneOf3IntegrationWithDefaults() *UpdateIntegrationsRequestOneOf3Integration`

NewUpdateIntegrationsRequestOneOf3IntegrationWithDefaults instantiates a new UpdateIntegrationsRequestOneOf3Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetType(v string)`

SetType sets Type field to given value.


### GetServiceMode

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetServiceUrl

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetSecondary

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetSecondary() string`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetSecondaryOk() (*string, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetSecondary(v string)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetServicePort

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceUsername

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.


### GetServicePassword

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceKey

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetAuthKey

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetServicePath

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### GetServiceVersion

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetServiceWindowsVersion

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceWindowsVersion() string`

GetServiceWindowsVersion returns the ServiceWindowsVersion field if non-nil, zero value otherwise.

### GetServiceWindowsVersionOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceWindowsVersionOk() (*string, bool)`

GetServiceWindowsVersionOk returns a tuple with the ServiceWindowsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWindowsVersion

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServiceWindowsVersion(v string)`

SetServiceWindowsVersion sets ServiceWindowsVersion field to given value.

### HasServiceWindowsVersion

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasServiceWindowsVersion() bool`

HasServiceWindowsVersion returns a boolean if a field has been set.

### GetRepoUrl

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetServiceConfig

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceConfig() string`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceConfigOk() (*string, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServiceConfig(v string)`

SetServiceConfig sets ServiceConfig field to given value.

### HasServiceConfig

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasServiceConfig() bool`

HasServiceConfig returns a boolean if a field has been set.

### GetServiceCommand

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceCommand() string`

GetServiceCommand returns the ServiceCommand field if non-nil, zero value otherwise.

### GetServiceCommandOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetServiceCommandOk() (*string, bool)`

GetServiceCommandOk returns a tuple with the ServiceCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCommand

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetServiceCommand(v string)`

SetServiceCommand sets ServiceCommand field to given value.

### HasServiceCommand

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasServiceCommand() bool`

HasServiceCommand returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetConfig() UpdateIntegrationsRequestOneOf3IntegrationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIntegrationsRequestOneOf3Integration) GetConfigOk() (*UpdateIntegrationsRequestOneOf3IntegrationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIntegrationsRequestOneOf3Integration) SetConfig(v UpdateIntegrationsRequestOneOf3IntegrationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIntegrationsRequestOneOf3Integration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


