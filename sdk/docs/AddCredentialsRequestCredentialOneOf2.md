# AddCredentialsRequestCredentialOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**AddCredentialsRequestCredentialOneOfIntegration**](AddCredentialsRequestCredentialOneOfIntegration.md) |  | [optional] 
**Username** | **string** | Email | 
**AuthKey** | [**AddCredentialsRequestCredentialOneOf2AuthKey**](AddCredentialsRequestCredentialOneOf2AuthKey.md) |  | 

## Methods

### NewAddCredentialsRequestCredentialOneOf2

`func NewAddCredentialsRequestCredentialOneOf2(type_ string, name string, username string, authKey AddCredentialsRequestCredentialOneOf2AuthKey, ) *AddCredentialsRequestCredentialOneOf2`

NewAddCredentialsRequestCredentialOneOf2 instantiates a new AddCredentialsRequestCredentialOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCredentialsRequestCredentialOneOf2WithDefaults

`func NewAddCredentialsRequestCredentialOneOf2WithDefaults() *AddCredentialsRequestCredentialOneOf2`

NewAddCredentialsRequestCredentialOneOf2WithDefaults instantiates a new AddCredentialsRequestCredentialOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddCredentialsRequestCredentialOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCredentialsRequestCredentialOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCredentialsRequestCredentialOneOf2) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AddCredentialsRequestCredentialOneOf2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCredentialsRequestCredentialOneOf2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCredentialsRequestCredentialOneOf2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddCredentialsRequestCredentialOneOf2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCredentialsRequestCredentialOneOf2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCredentialsRequestCredentialOneOf2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCredentialsRequestCredentialOneOf2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCredentialsRequestCredentialOneOf2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCredentialsRequestCredentialOneOf2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCredentialsRequestCredentialOneOf2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddCredentialsRequestCredentialOneOf2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *AddCredentialsRequestCredentialOneOf2) GetIntegration() AddCredentialsRequestCredentialOneOfIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AddCredentialsRequestCredentialOneOf2) GetIntegrationOk() (*AddCredentialsRequestCredentialOneOfIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AddCredentialsRequestCredentialOneOf2) SetIntegration(v AddCredentialsRequestCredentialOneOfIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *AddCredentialsRequestCredentialOneOf2) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUsername

`func (o *AddCredentialsRequestCredentialOneOf2) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddCredentialsRequestCredentialOneOf2) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddCredentialsRequestCredentialOneOf2) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAuthKey

`func (o *AddCredentialsRequestCredentialOneOf2) GetAuthKey() AddCredentialsRequestCredentialOneOf2AuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *AddCredentialsRequestCredentialOneOf2) GetAuthKeyOk() (*AddCredentialsRequestCredentialOneOf2AuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *AddCredentialsRequestCredentialOneOf2) SetAuthKey(v AddCredentialsRequestCredentialOneOf2AuthKey)`

SetAuthKey sets AuthKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


