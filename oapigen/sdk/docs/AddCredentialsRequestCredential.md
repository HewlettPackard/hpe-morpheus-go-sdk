# AddCredentialsRequestCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**AddCredentialsRequestCredentialOneOfIntegration**](AddCredentialsRequestCredentialOneOfIntegration.md) |  | [optional] 
**Username** | **string** | Username | 
**Password** | **string** | Password | 
**AuthKey** | [**AddCredentialsRequestCredentialOneOf2AuthKey**](AddCredentialsRequestCredentialOneOf2AuthKey.md) |  | 
**AuthPath** | **string** | Tenant name | 
**Config** | [**AddCredentialsRequestCredentialOneOf8Config**](AddCredentialsRequestCredentialOneOf8Config.md) |  | 

## Methods

### NewAddCredentialsRequestCredential

`func NewAddCredentialsRequestCredential(type_ string, name string, username string, password string, authKey AddCredentialsRequestCredentialOneOf2AuthKey, authPath string, config AddCredentialsRequestCredentialOneOf8Config, ) *AddCredentialsRequestCredential`

NewAddCredentialsRequestCredential instantiates a new AddCredentialsRequestCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCredentialsRequestCredentialWithDefaults

`func NewAddCredentialsRequestCredentialWithDefaults() *AddCredentialsRequestCredential`

NewAddCredentialsRequestCredentialWithDefaults instantiates a new AddCredentialsRequestCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddCredentialsRequestCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCredentialsRequestCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCredentialsRequestCredential) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AddCredentialsRequestCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCredentialsRequestCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCredentialsRequestCredential) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddCredentialsRequestCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCredentialsRequestCredential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCredentialsRequestCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCredentialsRequestCredential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCredentialsRequestCredential) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCredentialsRequestCredential) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCredentialsRequestCredential) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddCredentialsRequestCredential) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *AddCredentialsRequestCredential) GetIntegration() AddCredentialsRequestCredentialOneOfIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AddCredentialsRequestCredential) GetIntegrationOk() (*AddCredentialsRequestCredentialOneOfIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AddCredentialsRequestCredential) SetIntegration(v AddCredentialsRequestCredentialOneOfIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *AddCredentialsRequestCredential) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUsername

`func (o *AddCredentialsRequestCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddCredentialsRequestCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddCredentialsRequestCredential) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AddCredentialsRequestCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddCredentialsRequestCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddCredentialsRequestCredential) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAuthKey

`func (o *AddCredentialsRequestCredential) GetAuthKey() AddCredentialsRequestCredentialOneOf2AuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *AddCredentialsRequestCredential) GetAuthKeyOk() (*AddCredentialsRequestCredentialOneOf2AuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *AddCredentialsRequestCredential) SetAuthKey(v AddCredentialsRequestCredentialOneOf2AuthKey)`

SetAuthKey sets AuthKey field to given value.


### GetAuthPath

`func (o *AddCredentialsRequestCredential) GetAuthPath() string`

GetAuthPath returns the AuthPath field if non-nil, zero value otherwise.

### GetAuthPathOk

`func (o *AddCredentialsRequestCredential) GetAuthPathOk() (*string, bool)`

GetAuthPathOk returns a tuple with the AuthPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPath

`func (o *AddCredentialsRequestCredential) SetAuthPath(v string)`

SetAuthPath sets AuthPath field to given value.


### GetConfig

`func (o *AddCredentialsRequestCredential) GetConfig() AddCredentialsRequestCredentialOneOf8Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddCredentialsRequestCredential) GetConfigOk() (*AddCredentialsRequestCredentialOneOf8Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddCredentialsRequestCredential) SetConfig(v AddCredentialsRequestCredentialOneOf8Config)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


