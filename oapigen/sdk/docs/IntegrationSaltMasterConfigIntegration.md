# IntegrationSaltMasterConfigIntegration

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
**Config** | Pointer to [**IntegrationSaltMasterConfigIntegrationConfig**](IntegrationSaltMasterConfigIntegrationConfig.md) |  | [optional] 

## Methods

### NewIntegrationSaltMasterConfigIntegration

`func NewIntegrationSaltMasterConfigIntegration(name string, type_ string, serviceUrl string, serviceUsername string, ) *IntegrationSaltMasterConfigIntegration`

NewIntegrationSaltMasterConfigIntegration instantiates a new IntegrationSaltMasterConfigIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *IntegrationSaltMasterConfigIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationSaltMasterConfigIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationSaltMasterConfigIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IntegrationSaltMasterConfigIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationSaltMasterConfigIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationSaltMasterConfigIntegration) SetType(v string)`

SetType sets Type field to given value.


### GetServiceMode

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *IntegrationSaltMasterConfigIntegration) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *IntegrationSaltMasterConfigIntegration) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetServiceUrl

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *IntegrationSaltMasterConfigIntegration) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetSecondary

`func (o *IntegrationSaltMasterConfigIntegration) GetSecondary() string`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *IntegrationSaltMasterConfigIntegration) GetSecondaryOk() (*string, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *IntegrationSaltMasterConfigIntegration) SetSecondary(v string)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *IntegrationSaltMasterConfigIntegration) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetServicePort

`func (o *IntegrationSaltMasterConfigIntegration) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *IntegrationSaltMasterConfigIntegration) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *IntegrationSaltMasterConfigIntegration) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceUsername

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *IntegrationSaltMasterConfigIntegration) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.


### GetServicePassword

`func (o *IntegrationSaltMasterConfigIntegration) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *IntegrationSaltMasterConfigIntegration) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *IntegrationSaltMasterConfigIntegration) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceKey

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *IntegrationSaltMasterConfigIntegration) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *IntegrationSaltMasterConfigIntegration) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetAuthKey

`func (o *IntegrationSaltMasterConfigIntegration) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *IntegrationSaltMasterConfigIntegration) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *IntegrationSaltMasterConfigIntegration) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *IntegrationSaltMasterConfigIntegration) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetServicePath

`func (o *IntegrationSaltMasterConfigIntegration) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *IntegrationSaltMasterConfigIntegration) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *IntegrationSaltMasterConfigIntegration) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### GetServiceVersion

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *IntegrationSaltMasterConfigIntegration) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *IntegrationSaltMasterConfigIntegration) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### GetServiceWindowsVersion

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceWindowsVersion() string`

GetServiceWindowsVersion returns the ServiceWindowsVersion field if non-nil, zero value otherwise.

### GetServiceWindowsVersionOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceWindowsVersionOk() (*string, bool)`

GetServiceWindowsVersionOk returns a tuple with the ServiceWindowsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWindowsVersion

`func (o *IntegrationSaltMasterConfigIntegration) SetServiceWindowsVersion(v string)`

SetServiceWindowsVersion sets ServiceWindowsVersion field to given value.

### HasServiceWindowsVersion

`func (o *IntegrationSaltMasterConfigIntegration) HasServiceWindowsVersion() bool`

HasServiceWindowsVersion returns a boolean if a field has been set.

### GetRepoUrl

`func (o *IntegrationSaltMasterConfigIntegration) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *IntegrationSaltMasterConfigIntegration) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *IntegrationSaltMasterConfigIntegration) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *IntegrationSaltMasterConfigIntegration) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetServiceConfig

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceConfig() string`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceConfigOk() (*string, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *IntegrationSaltMasterConfigIntegration) SetServiceConfig(v string)`

SetServiceConfig sets ServiceConfig field to given value.

### HasServiceConfig

`func (o *IntegrationSaltMasterConfigIntegration) HasServiceConfig() bool`

HasServiceConfig returns a boolean if a field has been set.

### GetServiceCommand

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceCommand() string`

GetServiceCommand returns the ServiceCommand field if non-nil, zero value otherwise.

### GetServiceCommandOk

`func (o *IntegrationSaltMasterConfigIntegration) GetServiceCommandOk() (*string, bool)`

GetServiceCommandOk returns a tuple with the ServiceCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCommand

`func (o *IntegrationSaltMasterConfigIntegration) SetServiceCommand(v string)`

SetServiceCommand sets ServiceCommand field to given value.

### HasServiceCommand

`func (o *IntegrationSaltMasterConfigIntegration) HasServiceCommand() bool`

HasServiceCommand returns a boolean if a field has been set.

### GetConfig

`func (o *IntegrationSaltMasterConfigIntegration) GetConfig() IntegrationSaltMasterConfigIntegrationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationSaltMasterConfigIntegration) GetConfigOk() (*IntegrationSaltMasterConfigIntegrationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationSaltMasterConfigIntegration) SetConfig(v IntegrationSaltMasterConfigIntegrationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IntegrationSaltMasterConfigIntegration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


