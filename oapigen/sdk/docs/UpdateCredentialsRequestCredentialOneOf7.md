# UpdateCredentialsRequestCredentialOneOf7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**UpdateCredentialsRequestCredentialOneOf7Integration**](UpdateCredentialsRequestCredentialOneOf7Integration.md) |  | [optional] 
**Username** | **string** | Username | 
**Password** | **string** | User password, API Key, or applicable secret | 
**AuthKey** | [**UpdateCredentialsRequestCredentialOneOf7AuthKey**](UpdateCredentialsRequestCredentialOneOf7AuthKey.md) |  | 

## Methods

### NewUpdateCredentialsRequestCredentialOneOf7

`func NewUpdateCredentialsRequestCredentialOneOf7(type_ string, name string, username string, password string, authKey UpdateCredentialsRequestCredentialOneOf7AuthKey, ) *UpdateCredentialsRequestCredentialOneOf7`

NewUpdateCredentialsRequestCredentialOneOf7 instantiates a new UpdateCredentialsRequestCredentialOneOf7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCredentialsRequestCredentialOneOf7WithDefaults

`func NewUpdateCredentialsRequestCredentialOneOf7WithDefaults() *UpdateCredentialsRequestCredentialOneOf7`

NewUpdateCredentialsRequestCredentialOneOf7WithDefaults instantiates a new UpdateCredentialsRequestCredentialOneOf7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCredentialsRequestCredentialOneOf7) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCredentialsRequestCredentialOneOf7) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCredentialsRequestCredentialOneOf7) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCredentialsRequestCredentialOneOf7) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCredentialsRequestCredentialOneOf7) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCredentialsRequestCredentialOneOf7) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetIntegration() UpdateCredentialsRequestCredentialOneOf7Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetIntegrationOk() (*UpdateCredentialsRequestCredentialOneOf7Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UpdateCredentialsRequestCredentialOneOf7) SetIntegration(v UpdateCredentialsRequestCredentialOneOf7Integration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *UpdateCredentialsRequestCredentialOneOf7) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateCredentialsRequestCredentialOneOf7) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateCredentialsRequestCredentialOneOf7) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAuthKey

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetAuthKey() UpdateCredentialsRequestCredentialOneOf7AuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *UpdateCredentialsRequestCredentialOneOf7) GetAuthKeyOk() (*UpdateCredentialsRequestCredentialOneOf7AuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *UpdateCredentialsRequestCredentialOneOf7) SetAuthKey(v UpdateCredentialsRequestCredentialOneOf7AuthKey)`

SetAuthKey sets AuthKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


