# UpdateIntegrationsRequestOneOf6Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name, a unique identifier for the integration | 
**Type** | **string** | Integration Type Code | 
**ServiceUsername** | **string** | Username | 
**ServicePassword** | Pointer to **string** | Password | [optional] 
**ServiceToken** | Pointer to **string** | Access Token | [optional] 
**ServiceKey** | Pointer to **int64** | Key Pair ID | [optional] 
**Config** | Pointer to [**UpdateIntegrationsRequestOneOf6IntegrationConfig**](UpdateIntegrationsRequestOneOf6IntegrationConfig.md) |  | [optional] 

## Methods

### NewUpdateIntegrationsRequestOneOf6Integration

`func NewUpdateIntegrationsRequestOneOf6Integration(name string, type_ string, serviceUsername string, ) *UpdateIntegrationsRequestOneOf6Integration`

NewUpdateIntegrationsRequestOneOf6Integration instantiates a new UpdateIntegrationsRequestOneOf6Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationsRequestOneOf6IntegrationWithDefaults

`func NewUpdateIntegrationsRequestOneOf6IntegrationWithDefaults() *UpdateIntegrationsRequestOneOf6Integration`

NewUpdateIntegrationsRequestOneOf6IntegrationWithDefaults instantiates a new UpdateIntegrationsRequestOneOf6Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateIntegrationsRequestOneOf6Integration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIntegrationsRequestOneOf6Integration) SetType(v string)`

SetType sets Type field to given value.


### GetServiceUsername

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *UpdateIntegrationsRequestOneOf6Integration) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.


### GetServicePassword

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *UpdateIntegrationsRequestOneOf6Integration) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *UpdateIntegrationsRequestOneOf6Integration) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceToken

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *UpdateIntegrationsRequestOneOf6Integration) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *UpdateIntegrationsRequestOneOf6Integration) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceKey

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetServiceKey() int64`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetServiceKeyOk() (*int64, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *UpdateIntegrationsRequestOneOf6Integration) SetServiceKey(v int64)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *UpdateIntegrationsRequestOneOf6Integration) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetConfig() UpdateIntegrationsRequestOneOf6IntegrationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIntegrationsRequestOneOf6Integration) GetConfigOk() (*UpdateIntegrationsRequestOneOf6IntegrationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIntegrationsRequestOneOf6Integration) SetConfig(v UpdateIntegrationsRequestOneOf6IntegrationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIntegrationsRequestOneOf6Integration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


