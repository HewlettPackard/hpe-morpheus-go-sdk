# IntegrationAnsibleConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name, a unique identifier for the integration | 
**Type** | **string** | Integration Type Code | 
**Enabled** | Pointer to **bool** | Set &#x60;true&#x60; to enable integration | [optional] 
**Refresh** | Pointer to **bool** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  | [optional] [default to true]
**ServiceUrl** | **string** | Ansible Git URL | 
**ServiceUsername** | Pointer to **string** | Git Username | [optional] 
**ServicePassword** | Pointer to **string** | Git Password or Token depending on the Git host | [optional] 
**ServiceToken** | Pointer to **string** | Git Token | [optional] 
**ServiceKey** | Pointer to **int64** | Keypair ID | [optional] 
**Config** | Pointer to [**IntegrationAnsibleConfigIntegrationConfig**](IntegrationAnsibleConfigIntegrationConfig.md) |  | [optional] 

## Methods

### NewIntegrationAnsibleConfigIntegration

`func NewIntegrationAnsibleConfigIntegration(name string, type_ string, serviceUrl string, ) *IntegrationAnsibleConfigIntegration`

NewIntegrationAnsibleConfigIntegration instantiates a new IntegrationAnsibleConfigIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationAnsibleConfigIntegrationWithDefaults

`func NewIntegrationAnsibleConfigIntegrationWithDefaults() *IntegrationAnsibleConfigIntegration`

NewIntegrationAnsibleConfigIntegrationWithDefaults instantiates a new IntegrationAnsibleConfigIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IntegrationAnsibleConfigIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationAnsibleConfigIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationAnsibleConfigIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IntegrationAnsibleConfigIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationAnsibleConfigIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationAnsibleConfigIntegration) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *IntegrationAnsibleConfigIntegration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegrationAnsibleConfigIntegration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegrationAnsibleConfigIntegration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IntegrationAnsibleConfigIntegration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefresh

`func (o *IntegrationAnsibleConfigIntegration) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *IntegrationAnsibleConfigIntegration) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *IntegrationAnsibleConfigIntegration) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *IntegrationAnsibleConfigIntegration) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetServiceUrl

`func (o *IntegrationAnsibleConfigIntegration) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *IntegrationAnsibleConfigIntegration) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *IntegrationAnsibleConfigIntegration) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetServiceUsername

`func (o *IntegrationAnsibleConfigIntegration) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *IntegrationAnsibleConfigIntegration) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *IntegrationAnsibleConfigIntegration) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *IntegrationAnsibleConfigIntegration) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *IntegrationAnsibleConfigIntegration) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *IntegrationAnsibleConfigIntegration) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *IntegrationAnsibleConfigIntegration) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *IntegrationAnsibleConfigIntegration) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceToken

`func (o *IntegrationAnsibleConfigIntegration) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *IntegrationAnsibleConfigIntegration) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *IntegrationAnsibleConfigIntegration) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *IntegrationAnsibleConfigIntegration) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceKey

`func (o *IntegrationAnsibleConfigIntegration) GetServiceKey() int64`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *IntegrationAnsibleConfigIntegration) GetServiceKeyOk() (*int64, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *IntegrationAnsibleConfigIntegration) SetServiceKey(v int64)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *IntegrationAnsibleConfigIntegration) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetConfig

`func (o *IntegrationAnsibleConfigIntegration) GetConfig() IntegrationAnsibleConfigIntegrationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationAnsibleConfigIntegration) GetConfigOk() (*IntegrationAnsibleConfigIntegrationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationAnsibleConfigIntegration) SetConfig(v IntegrationAnsibleConfigIntegrationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IntegrationAnsibleConfigIntegration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


