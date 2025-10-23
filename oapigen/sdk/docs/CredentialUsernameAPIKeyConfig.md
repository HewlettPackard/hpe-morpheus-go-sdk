# CredentialUsernameAPIKeyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**AddCredentialsRequestCredentialOneOfIntegration**](AddCredentialsRequestCredentialOneOfIntegration.md) |  | [optional] 
**Username** | **string** | Username | 
**Password** | **string** | API Key | 

## Methods

### NewCredentialUsernameAPIKeyConfig

`func NewCredentialUsernameAPIKeyConfig(type_ string, name string, username string, password string, ) *CredentialUsernameAPIKeyConfig`

NewCredentialUsernameAPIKeyConfig instantiates a new CredentialUsernameAPIKeyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialUsernameAPIKeyConfigWithDefaults

`func NewCredentialUsernameAPIKeyConfigWithDefaults() *CredentialUsernameAPIKeyConfig`

NewCredentialUsernameAPIKeyConfigWithDefaults instantiates a new CredentialUsernameAPIKeyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CredentialUsernameAPIKeyConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialUsernameAPIKeyConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialUsernameAPIKeyConfig) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CredentialUsernameAPIKeyConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialUsernameAPIKeyConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialUsernameAPIKeyConfig) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CredentialUsernameAPIKeyConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialUsernameAPIKeyConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialUsernameAPIKeyConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialUsernameAPIKeyConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CredentialUsernameAPIKeyConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CredentialUsernameAPIKeyConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CredentialUsernameAPIKeyConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CredentialUsernameAPIKeyConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *CredentialUsernameAPIKeyConfig) GetIntegration() AddCredentialsRequestCredentialOneOfIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CredentialUsernameAPIKeyConfig) GetIntegrationOk() (*AddCredentialsRequestCredentialOneOfIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CredentialUsernameAPIKeyConfig) SetIntegration(v AddCredentialsRequestCredentialOneOfIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *CredentialUsernameAPIKeyConfig) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUsername

`func (o *CredentialUsernameAPIKeyConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialUsernameAPIKeyConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialUsernameAPIKeyConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CredentialUsernameAPIKeyConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CredentialUsernameAPIKeyConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CredentialUsernameAPIKeyConfig) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


