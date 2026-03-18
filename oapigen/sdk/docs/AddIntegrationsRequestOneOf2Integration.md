# AddIntegrationsRequestOneOf2Integration

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
**Config** | Pointer to [**AddIntegrationsRequestOneOf2IntegrationConfig**](AddIntegrationsRequestOneOf2IntegrationConfig.md) |  | [optional] 

## Methods

### NewAddIntegrationsRequestOneOf2Integration

`func NewAddIntegrationsRequestOneOf2Integration(name string, type_ string, serviceUrl string, serviceUsername string, servicePassword string, ) *AddIntegrationsRequestOneOf2Integration`

NewAddIntegrationsRequestOneOf2Integration instantiates a new AddIntegrationsRequestOneOf2Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrationsRequestOneOf2IntegrationWithDefaults

`func NewAddIntegrationsRequestOneOf2IntegrationWithDefaults() *AddIntegrationsRequestOneOf2Integration`

NewAddIntegrationsRequestOneOf2IntegrationWithDefaults instantiates a new AddIntegrationsRequestOneOf2Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddIntegrationsRequestOneOf2Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIntegrationsRequestOneOf2Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIntegrationsRequestOneOf2Integration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddIntegrationsRequestOneOf2Integration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIntegrationsRequestOneOf2Integration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIntegrationsRequestOneOf2Integration) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *AddIntegrationsRequestOneOf2Integration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIntegrationsRequestOneOf2Integration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIntegrationsRequestOneOf2Integration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddIntegrationsRequestOneOf2Integration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefresh

`func (o *AddIntegrationsRequestOneOf2Integration) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *AddIntegrationsRequestOneOf2Integration) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *AddIntegrationsRequestOneOf2Integration) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *AddIntegrationsRequestOneOf2Integration) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetServiceUrl

`func (o *AddIntegrationsRequestOneOf2Integration) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *AddIntegrationsRequestOneOf2Integration) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *AddIntegrationsRequestOneOf2Integration) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetServiceUsername

`func (o *AddIntegrationsRequestOneOf2Integration) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *AddIntegrationsRequestOneOf2Integration) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *AddIntegrationsRequestOneOf2Integration) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.


### GetServicePassword

`func (o *AddIntegrationsRequestOneOf2Integration) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *AddIntegrationsRequestOneOf2Integration) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *AddIntegrationsRequestOneOf2Integration) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.


### GetConfig

`func (o *AddIntegrationsRequestOneOf2Integration) GetConfig() AddIntegrationsRequestOneOf2IntegrationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddIntegrationsRequestOneOf2Integration) GetConfigOk() (*AddIntegrationsRequestOneOf2IntegrationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddIntegrationsRequestOneOf2Integration) SetConfig(v AddIntegrationsRequestOneOf2IntegrationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddIntegrationsRequestOneOf2Integration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


