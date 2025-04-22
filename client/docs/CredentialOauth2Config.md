# CredentialOauth2Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**AddCredentialsRequestCredentialOneOfIntegration**](AddCredentialsRequestCredentialOneOfIntegration.md) |  | [optional] 
**Username** | Pointer to **string** | Username | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**Config** | [**AddCredentialsRequestCredentialOneOf8Config**](AddCredentialsRequestCredentialOneOf8Config.md) |  | 

## Methods

### NewCredentialOauth2Config

`func NewCredentialOauth2Config(type_ string, name string, config AddCredentialsRequestCredentialOneOf8Config, ) *CredentialOauth2Config`

NewCredentialOauth2Config instantiates a new CredentialOauth2Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialOauth2ConfigWithDefaults

`func NewCredentialOauth2ConfigWithDefaults() *CredentialOauth2Config`

NewCredentialOauth2ConfigWithDefaults instantiates a new CredentialOauth2Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CredentialOauth2Config) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialOauth2Config) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialOauth2Config) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CredentialOauth2Config) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialOauth2Config) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialOauth2Config) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CredentialOauth2Config) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialOauth2Config) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialOauth2Config) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialOauth2Config) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CredentialOauth2Config) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CredentialOauth2Config) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CredentialOauth2Config) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CredentialOauth2Config) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *CredentialOauth2Config) GetIntegration() AddCredentialsRequestCredentialOneOfIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CredentialOauth2Config) GetIntegrationOk() (*AddCredentialsRequestCredentialOneOfIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CredentialOauth2Config) SetIntegration(v AddCredentialsRequestCredentialOneOfIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *CredentialOauth2Config) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUsername

`func (o *CredentialOauth2Config) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialOauth2Config) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialOauth2Config) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CredentialOauth2Config) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *CredentialOauth2Config) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialOauth2Config) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialOauth2Config) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CredentialOauth2Config) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetConfig

`func (o *CredentialOauth2Config) GetConfig() AddCredentialsRequestCredentialOneOf8Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CredentialOauth2Config) GetConfigOk() (*AddCredentialsRequestCredentialOneOf8Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CredentialOauth2Config) SetConfig(v AddCredentialsRequestCredentialOneOf8Config)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


