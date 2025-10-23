# CredentialTenantUsernameKeypairConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**AddCredentialsRequestCredentialOneOfIntegration**](AddCredentialsRequestCredentialOneOfIntegration.md) |  | [optional] 
**AuthPath** | **string** | Tenant name | 
**Username** | **string** | Username | 
**AuthKey** | [**AddCredentialsRequestCredentialOneOf2AuthKey**](AddCredentialsRequestCredentialOneOf2AuthKey.md) |  | 

## Methods

### NewCredentialTenantUsernameKeypairConfig

`func NewCredentialTenantUsernameKeypairConfig(type_ string, name string, authPath string, username string, authKey AddCredentialsRequestCredentialOneOf2AuthKey, ) *CredentialTenantUsernameKeypairConfig`

NewCredentialTenantUsernameKeypairConfig instantiates a new CredentialTenantUsernameKeypairConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTenantUsernameKeypairConfigWithDefaults

`func NewCredentialTenantUsernameKeypairConfigWithDefaults() *CredentialTenantUsernameKeypairConfig`

NewCredentialTenantUsernameKeypairConfigWithDefaults instantiates a new CredentialTenantUsernameKeypairConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CredentialTenantUsernameKeypairConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialTenantUsernameKeypairConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialTenantUsernameKeypairConfig) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CredentialTenantUsernameKeypairConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CredentialTenantUsernameKeypairConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CredentialTenantUsernameKeypairConfig) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CredentialTenantUsernameKeypairConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CredentialTenantUsernameKeypairConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CredentialTenantUsernameKeypairConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CredentialTenantUsernameKeypairConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CredentialTenantUsernameKeypairConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CredentialTenantUsernameKeypairConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CredentialTenantUsernameKeypairConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CredentialTenantUsernameKeypairConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *CredentialTenantUsernameKeypairConfig) GetIntegration() AddCredentialsRequestCredentialOneOfIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *CredentialTenantUsernameKeypairConfig) GetIntegrationOk() (*AddCredentialsRequestCredentialOneOfIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *CredentialTenantUsernameKeypairConfig) SetIntegration(v AddCredentialsRequestCredentialOneOfIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *CredentialTenantUsernameKeypairConfig) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetAuthPath

`func (o *CredentialTenantUsernameKeypairConfig) GetAuthPath() string`

GetAuthPath returns the AuthPath field if non-nil, zero value otherwise.

### GetAuthPathOk

`func (o *CredentialTenantUsernameKeypairConfig) GetAuthPathOk() (*string, bool)`

GetAuthPathOk returns a tuple with the AuthPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPath

`func (o *CredentialTenantUsernameKeypairConfig) SetAuthPath(v string)`

SetAuthPath sets AuthPath field to given value.


### GetUsername

`func (o *CredentialTenantUsernameKeypairConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CredentialTenantUsernameKeypairConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CredentialTenantUsernameKeypairConfig) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAuthKey

`func (o *CredentialTenantUsernameKeypairConfig) GetAuthKey() AddCredentialsRequestCredentialOneOf2AuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *CredentialTenantUsernameKeypairConfig) GetAuthKeyOk() (*AddCredentialsRequestCredentialOneOf2AuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *CredentialTenantUsernameKeypairConfig) SetAuthKey(v AddCredentialsRequestCredentialOneOf2AuthKey)`

SetAuthKey sets AuthKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


