# AddIntegrationsRequestOneOf6Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name, a unique identifier for the integration | 
**Type** | **string** | Integration Type Code | 
**ServiceUsername** | **string** | Username | 
**ServicePassword** | Pointer to **string** | Password | [optional] 
**ServiceToken** | Pointer to **string** | Access Token | [optional] 
**ServiceKey** | Pointer to **int64** | Key Pair ID | [optional] 
**Config** | Pointer to [**AddIntegrationsRequestOneOf6IntegrationConfig**](AddIntegrationsRequestOneOf6IntegrationConfig.md) |  | [optional] 

## Methods

### NewAddIntegrationsRequestOneOf6Integration

`func NewAddIntegrationsRequestOneOf6Integration(name string, type_ string, serviceUsername string, ) *AddIntegrationsRequestOneOf6Integration`

NewAddIntegrationsRequestOneOf6Integration instantiates a new AddIntegrationsRequestOneOf6Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrationsRequestOneOf6IntegrationWithDefaults

`func NewAddIntegrationsRequestOneOf6IntegrationWithDefaults() *AddIntegrationsRequestOneOf6Integration`

NewAddIntegrationsRequestOneOf6IntegrationWithDefaults instantiates a new AddIntegrationsRequestOneOf6Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddIntegrationsRequestOneOf6Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIntegrationsRequestOneOf6Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIntegrationsRequestOneOf6Integration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddIntegrationsRequestOneOf6Integration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIntegrationsRequestOneOf6Integration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIntegrationsRequestOneOf6Integration) SetType(v string)`

SetType sets Type field to given value.


### GetServiceUsername

`func (o *AddIntegrationsRequestOneOf6Integration) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *AddIntegrationsRequestOneOf6Integration) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *AddIntegrationsRequestOneOf6Integration) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.


### GetServicePassword

`func (o *AddIntegrationsRequestOneOf6Integration) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *AddIntegrationsRequestOneOf6Integration) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *AddIntegrationsRequestOneOf6Integration) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *AddIntegrationsRequestOneOf6Integration) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceToken

`func (o *AddIntegrationsRequestOneOf6Integration) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *AddIntegrationsRequestOneOf6Integration) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *AddIntegrationsRequestOneOf6Integration) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *AddIntegrationsRequestOneOf6Integration) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceKey

`func (o *AddIntegrationsRequestOneOf6Integration) GetServiceKey() int64`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *AddIntegrationsRequestOneOf6Integration) GetServiceKeyOk() (*int64, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *AddIntegrationsRequestOneOf6Integration) SetServiceKey(v int64)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *AddIntegrationsRequestOneOf6Integration) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetConfig

`func (o *AddIntegrationsRequestOneOf6Integration) GetConfig() AddIntegrationsRequestOneOf6IntegrationConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddIntegrationsRequestOneOf6Integration) GetConfigOk() (*AddIntegrationsRequestOneOf6IntegrationConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddIntegrationsRequestOneOf6Integration) SetConfig(v AddIntegrationsRequestOneOf6IntegrationConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddIntegrationsRequestOneOf6Integration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


