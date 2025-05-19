# CredentialAccessSecretKeyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**AddCredentialsRequestCredentialOneOfIntegration**](AddCredentialsRequestCredentialOneOfIntegration.md) |  | [optional] 
**Username** | **string** | Access Key | 
**Password** | **string** | Secret Key | 

## Methods

### NewCredentialAccessSecretKeyConfig

`func NewCredentialAccessSecretKeyConfig(type_ string, name string, username string, password string, ) *CredentialAccessSecretKeyConfig`

NewCredentialAccessSecretKeyConfig instantiates a new CredentialAccessSecretKeyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialAccessSecretKeyConfigWithDefaults

`func NewCredentialAccessSecretKeyConfigWithDefaults() *CredentialAccessSecretKeyConfig`

NewCredentialAccessSecretKeyConfigWithDefaults instantiates a new CredentialAccessSecretKeyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CredentialAccessSecretKeyConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialAccessSecretKeyConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialAccessSecretKeyConfig) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CredentialAccessSecretKeyConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialAccessSecretKeyConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialAccessSecretKeyConfig) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CredentialAccessSecretKeyConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialAccessSecretKeyConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialAccessSecretKeyConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialAccessSecretKeyConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CredentialAccessSecretKeyConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CredentialAccessSecretKeyConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CredentialAccessSecretKeyConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CredentialAccessSecretKeyConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *CredentialAccessSecretKeyConfig) GetIntegration() AddCredentialsRequestCredentialOneOfIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CredentialAccessSecretKeyConfig) GetIntegrationOk() (*AddCredentialsRequestCredentialOneOfIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CredentialAccessSecretKeyConfig) SetIntegration(v AddCredentialsRequestCredentialOneOfIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *CredentialAccessSecretKeyConfig) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUsername

`func (o *CredentialAccessSecretKeyConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialAccessSecretKeyConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialAccessSecretKeyConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CredentialAccessSecretKeyConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialAccessSecretKeyConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialAccessSecretKeyConfig) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


