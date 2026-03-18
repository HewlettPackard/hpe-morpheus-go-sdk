# UpdateCredentialsRequestCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**UpdateCredentialsRequestCredentialOneOf8Integration**](UpdateCredentialsRequestCredentialOneOf8Integration.md) |  | [optional] 
**Username** | **string** | Username | 
**Password** | **string** | Password | 
**AuthKey** | [**UpdateCredentialsRequestCredentialOneOf7AuthKey**](UpdateCredentialsRequestCredentialOneOf7AuthKey.md) |  | 
**AuthPath** | **string** | Tenant name | 
**Config** | [**UpdateCredentialsRequestCredentialOneOf8Config**](UpdateCredentialsRequestCredentialOneOf8Config.md) |  | 

## Methods

### NewUpdateCredentialsRequestCredential

`func NewUpdateCredentialsRequestCredential(type_ string, name string, username string, password string, authKey UpdateCredentialsRequestCredentialOneOf7AuthKey, authPath string, config UpdateCredentialsRequestCredentialOneOf8Config, ) *UpdateCredentialsRequestCredential`

NewUpdateCredentialsRequestCredential instantiates a new UpdateCredentialsRequestCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCredentialsRequestCredentialWithDefaults

`func NewUpdateCredentialsRequestCredentialWithDefaults() *UpdateCredentialsRequestCredential`

NewUpdateCredentialsRequestCredentialWithDefaults instantiates a new UpdateCredentialsRequestCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateCredentialsRequestCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCredentialsRequestCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCredentialsRequestCredential) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *UpdateCredentialsRequestCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCredentialsRequestCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCredentialsRequestCredential) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateCredentialsRequestCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCredentialsRequestCredential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCredentialsRequestCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCredentialsRequestCredential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateCredentialsRequestCredential) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCredentialsRequestCredential) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCredentialsRequestCredential) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCredentialsRequestCredential) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *UpdateCredentialsRequestCredential) GetIntegration() UpdateCredentialsRequestCredentialOneOf8Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UpdateCredentialsRequestCredential) GetIntegrationOk() (*UpdateCredentialsRequestCredentialOneOf8Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UpdateCredentialsRequestCredential) SetIntegration(v UpdateCredentialsRequestCredentialOneOf8Integration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *UpdateCredentialsRequestCredential) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateCredentialsRequestCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateCredentialsRequestCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateCredentialsRequestCredential) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UpdateCredentialsRequestCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateCredentialsRequestCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateCredentialsRequestCredential) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAuthKey

`func (o *UpdateCredentialsRequestCredential) GetAuthKey() UpdateCredentialsRequestCredentialOneOf7AuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *UpdateCredentialsRequestCredential) GetAuthKeyOk() (*UpdateCredentialsRequestCredentialOneOf7AuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *UpdateCredentialsRequestCredential) SetAuthKey(v UpdateCredentialsRequestCredentialOneOf7AuthKey)`

SetAuthKey sets AuthKey field to given value.


### GetAuthPath

`func (o *UpdateCredentialsRequestCredential) GetAuthPath() string`

GetAuthPath returns the AuthPath field if non-nil, zero value otherwise.

### GetAuthPathOk

`func (o *UpdateCredentialsRequestCredential) GetAuthPathOk() (*string, bool)`

GetAuthPathOk returns a tuple with the AuthPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPath

`func (o *UpdateCredentialsRequestCredential) SetAuthPath(v string)`

SetAuthPath sets AuthPath field to given value.


### GetConfig

`func (o *UpdateCredentialsRequestCredential) GetConfig() UpdateCredentialsRequestCredentialOneOf8Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateCredentialsRequestCredential) GetConfigOk() (*UpdateCredentialsRequestCredentialOneOf8Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateCredentialsRequestCredential) SetConfig(v UpdateCredentialsRequestCredentialOneOf8Config)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


