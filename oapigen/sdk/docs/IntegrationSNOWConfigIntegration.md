# IntegrationSNOWConfigIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name, a unique identifier for the integration | 
**Type** | **string** | Integration Type Code | 
**Enabled** | Pointer to **bool** | Set &#x60;true&#x60; to enable integration | [optional] 
**Refresh** | Pointer to **bool** | Pass &#x60;false&#x60; to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  | [optional] [default to true]
**ServiceUrl** | **string** | ServiceNow Host | 
**ServiceUsername** | **string** | ServiceNow Username | 
**ServicePassword** | **string** | ServiceNow Password | 
**Config** | Pointer to [**IntegrationSNOWConfigIntegrationConfig**](IntegrationSNOWConfigIntegrationConfig.md) |  | [optional] 

## Methods

### NewIntegrationSNOWConfigIntegration

`func NewIntegrationSNOWConfigIntegration(name string, type_ string, serviceUrl string, serviceUsername string, servicePassword string, ) *IntegrationSNOWConfigIntegration`

NewIntegrationSNOWConfigIntegration instantiates a new IntegrationSNOWConfigIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationSNOWConfigIntegrationWithDefaults

`func NewIntegrationSNOWConfigIntegrationWithDefaults() *IntegrationSNOWConfigIntegration`

NewIntegrationSNOWConfigIntegrationWithDefaults instantiates a new IntegrationSNOWConfigIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IntegrationSNOWConfigIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationSNOWConfigIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationSNOWConfigIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *IntegrationSNOWConfigIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationSNOWConfigIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationSNOWConfigIntegration) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *IntegrationSNOWConfigIntegration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegrationSNOWConfigIntegration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegrationSNOWConfigIntegration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IntegrationSNOWConfigIntegration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefresh

`func (o *IntegrationSNOWConfigIntegration) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *IntegrationSNOWConfigIntegration) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *IntegrationSNOWConfigIntegration) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *IntegrationSNOWConfigIntegration) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetServiceUrl

`func (o *IntegrationSNOWConfigIntegration) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *IntegrationSNOWConfigIntegration) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *IntegrationSNOWConfigIntegration) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetServiceUsername

`func (o *IntegrationSNOWConfigIntegration) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *IntegrationSNOWConfigIntegration) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *IntegrationSNOWConfigIntegration) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.


### GetServicePassword

`func (o *IntegrationSNOWConfigIntegration) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *IntegrationSNOWConfigIntegration) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *IntegrationSNOWConfigIntegration) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.


### GetConfig

`func (o *IntegrationSNOWConfigIntegration) GetConfig() IntegrationSNOWConfigIntegrationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationSNOWConfigIntegration) GetConfigOk() (*IntegrationSNOWConfigIntegrationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationSNOWConfigIntegration) SetConfig(v IntegrationSNOWConfigIntegrationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IntegrationSNOWConfigIntegration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


