# AddIntegrationsRequestOneOf1Integration

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
**Config** | Pointer to [**AddIntegrationsRequestOneOf1IntegrationConfig**](AddIntegrationsRequestOneOf1IntegrationConfig.md) |  | [optional] 

## Methods

### NewAddIntegrationsRequestOneOf1Integration

`func NewAddIntegrationsRequestOneOf1Integration(name string, type_ string, serviceUrl string, ) *AddIntegrationsRequestOneOf1Integration`

NewAddIntegrationsRequestOneOf1Integration instantiates a new AddIntegrationsRequestOneOf1Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrationsRequestOneOf1IntegrationWithDefaults

`func NewAddIntegrationsRequestOneOf1IntegrationWithDefaults() *AddIntegrationsRequestOneOf1Integration`

NewAddIntegrationsRequestOneOf1IntegrationWithDefaults instantiates a new AddIntegrationsRequestOneOf1Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddIntegrationsRequestOneOf1Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIntegrationsRequestOneOf1Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIntegrationsRequestOneOf1Integration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddIntegrationsRequestOneOf1Integration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIntegrationsRequestOneOf1Integration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIntegrationsRequestOneOf1Integration) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *AddIntegrationsRequestOneOf1Integration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIntegrationsRequestOneOf1Integration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIntegrationsRequestOneOf1Integration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddIntegrationsRequestOneOf1Integration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRefresh

`func (o *AddIntegrationsRequestOneOf1Integration) GetRefresh() bool`

GetRefresh returns the Refresh field if non-nil, zero value otherwise.

### GetRefreshOk

`func (o *AddIntegrationsRequestOneOf1Integration) GetRefreshOk() (*bool, bool)`

GetRefreshOk returns a tuple with the Refresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefresh

`func (o *AddIntegrationsRequestOneOf1Integration) SetRefresh(v bool)`

SetRefresh sets Refresh field to given value.

### HasRefresh

`func (o *AddIntegrationsRequestOneOf1Integration) HasRefresh() bool`

HasRefresh returns a boolean if a field has been set.

### GetServiceUrl

`func (o *AddIntegrationsRequestOneOf1Integration) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *AddIntegrationsRequestOneOf1Integration) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *AddIntegrationsRequestOneOf1Integration) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetServiceUsername

`func (o *AddIntegrationsRequestOneOf1Integration) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *AddIntegrationsRequestOneOf1Integration) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *AddIntegrationsRequestOneOf1Integration) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *AddIntegrationsRequestOneOf1Integration) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *AddIntegrationsRequestOneOf1Integration) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *AddIntegrationsRequestOneOf1Integration) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *AddIntegrationsRequestOneOf1Integration) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *AddIntegrationsRequestOneOf1Integration) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceToken

`func (o *AddIntegrationsRequestOneOf1Integration) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *AddIntegrationsRequestOneOf1Integration) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *AddIntegrationsRequestOneOf1Integration) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *AddIntegrationsRequestOneOf1Integration) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceKey

`func (o *AddIntegrationsRequestOneOf1Integration) GetServiceKey() int64`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *AddIntegrationsRequestOneOf1Integration) GetServiceKeyOk() (*int64, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *AddIntegrationsRequestOneOf1Integration) SetServiceKey(v int64)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *AddIntegrationsRequestOneOf1Integration) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetConfig

`func (o *AddIntegrationsRequestOneOf1Integration) GetConfig() AddIntegrationsRequestOneOf1IntegrationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddIntegrationsRequestOneOf1Integration) GetConfigOk() (*AddIntegrationsRequestOneOf1IntegrationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddIntegrationsRequestOneOf1Integration) SetConfig(v AddIntegrationsRequestOneOf1IntegrationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddIntegrationsRequestOneOf1Integration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


