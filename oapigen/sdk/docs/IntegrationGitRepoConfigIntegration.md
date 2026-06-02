# IntegrationGitRepoConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name, a unique identifier for the integration | 
**Type** | **string** | Integration Type Code | 
**ServiceUrl** | **string** | Git URL | 
**ServiceUsername** | **string** | Username | 
**ServicePassword** | Pointer to **string** | Password | [optional] 
**ServiceToken** | Pointer to **string** | Access Token | [optional] 
**ServiceKey** | Pointer to **int64** | Key Pair ID | [optional] 
**Config** | Pointer to [**IntegrationGitRepoConfigIntegrationConfig**](IntegrationGitRepoConfigIntegrationConfig.md) |  | [optional] 

## Methods

### NewIntegrationGitRepoConfigIntegration

`func NewIntegrationGitRepoConfigIntegration(name string, type_ string, serviceUrl string, serviceUsername string, ) *IntegrationGitRepoConfigIntegration`

NewIntegrationGitRepoConfigIntegration instantiates a new IntegrationGitRepoConfigIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *IntegrationGitRepoConfigIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationGitRepoConfigIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationGitRepoConfigIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IntegrationGitRepoConfigIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationGitRepoConfigIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationGitRepoConfigIntegration) SetType(v string)`

SetType sets Type field to given value.


### GetServiceUrl

`func (o *IntegrationGitRepoConfigIntegration) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *IntegrationGitRepoConfigIntegration) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *IntegrationGitRepoConfigIntegration) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetServiceUsername

`func (o *IntegrationGitRepoConfigIntegration) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *IntegrationGitRepoConfigIntegration) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *IntegrationGitRepoConfigIntegration) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.


### GetServicePassword

`func (o *IntegrationGitRepoConfigIntegration) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *IntegrationGitRepoConfigIntegration) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *IntegrationGitRepoConfigIntegration) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *IntegrationGitRepoConfigIntegration) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceToken

`func (o *IntegrationGitRepoConfigIntegration) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *IntegrationGitRepoConfigIntegration) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *IntegrationGitRepoConfigIntegration) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *IntegrationGitRepoConfigIntegration) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceKey

`func (o *IntegrationGitRepoConfigIntegration) GetServiceKey() int64`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *IntegrationGitRepoConfigIntegration) GetServiceKeyOk() (*int64, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *IntegrationGitRepoConfigIntegration) SetServiceKey(v int64)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *IntegrationGitRepoConfigIntegration) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetConfig

`func (o *IntegrationGitRepoConfigIntegration) GetConfig() IntegrationGitRepoConfigIntegrationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationGitRepoConfigIntegration) GetConfigOk() (*IntegrationGitRepoConfigIntegrationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationGitRepoConfigIntegration) SetConfig(v IntegrationGitRepoConfigIntegrationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IntegrationGitRepoConfigIntegration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


