# UpdateIntegrationsRequestOneOf5Integration

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
**Config** | Pointer to [**UpdateIntegrationsRequestOneOf5IntegrationConfig**](UpdateIntegrationsRequestOneOf5IntegrationConfig.md) |  | [optional] 

## Methods

### NewUpdateIntegrationsRequestOneOf5Integration

`func NewUpdateIntegrationsRequestOneOf5Integration(name string, type_ string, serviceUrl string, serviceUsername string, ) *UpdateIntegrationsRequestOneOf5Integration`

NewUpdateIntegrationsRequestOneOf5Integration instantiates a new UpdateIntegrationsRequestOneOf5Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationsRequestOneOf5IntegrationWithDefaults

`func NewUpdateIntegrationsRequestOneOf5IntegrationWithDefaults() *UpdateIntegrationsRequestOneOf5Integration`

NewUpdateIntegrationsRequestOneOf5IntegrationWithDefaults instantiates a new UpdateIntegrationsRequestOneOf5Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIntegrationsRequestOneOf5Integration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIntegrationsRequestOneOf5Integration) SetType(v string)`

SetType sets Type field to given value.


### GetServiceUrl

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *UpdateIntegrationsRequestOneOf5Integration) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetServiceUsername

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *UpdateIntegrationsRequestOneOf5Integration) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.


### GetServicePassword

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *UpdateIntegrationsRequestOneOf5Integration) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *UpdateIntegrationsRequestOneOf5Integration) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceToken

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *UpdateIntegrationsRequestOneOf5Integration) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *UpdateIntegrationsRequestOneOf5Integration) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceKey

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetServiceKey() int64`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetServiceKeyOk() (*int64, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *UpdateIntegrationsRequestOneOf5Integration) SetServiceKey(v int64)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *UpdateIntegrationsRequestOneOf5Integration) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetConfig() UpdateIntegrationsRequestOneOf5IntegrationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIntegrationsRequestOneOf5Integration) GetConfigOk() (*UpdateIntegrationsRequestOneOf5IntegrationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIntegrationsRequestOneOf5Integration) SetConfig(v UpdateIntegrationsRequestOneOf5IntegrationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIntegrationsRequestOneOf5Integration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


