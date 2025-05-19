# AddCredentialsRequestCredentialOneOf5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**AddCredentialsRequestCredentialOneOfIntegration**](AddCredentialsRequestCredentialOneOfIntegration.md) |  | [optional] 
**Username** | **string** | Username | 
**AuthKey** | [**AddCredentialsRequestCredentialOneOf2AuthKey**](AddCredentialsRequestCredentialOneOf2AuthKey.md) |  | 

## Methods

### NewAddCredentialsRequestCredentialOneOf5

`func NewAddCredentialsRequestCredentialOneOf5(type_ string, name string, username string, authKey AddCredentialsRequestCredentialOneOf2AuthKey, ) *AddCredentialsRequestCredentialOneOf5`

NewAddCredentialsRequestCredentialOneOf5 instantiates a new AddCredentialsRequestCredentialOneOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCredentialsRequestCredentialOneOf5WithDefaults

`func NewAddCredentialsRequestCredentialOneOf5WithDefaults() *AddCredentialsRequestCredentialOneOf5`

NewAddCredentialsRequestCredentialOneOf5WithDefaults instantiates a new AddCredentialsRequestCredentialOneOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddCredentialsRequestCredentialOneOf5) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCredentialsRequestCredentialOneOf5) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCredentialsRequestCredentialOneOf5) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AddCredentialsRequestCredentialOneOf5) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCredentialsRequestCredentialOneOf5) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCredentialsRequestCredentialOneOf5) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddCredentialsRequestCredentialOneOf5) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCredentialsRequestCredentialOneOf5) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCredentialsRequestCredentialOneOf5) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCredentialsRequestCredentialOneOf5) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCredentialsRequestCredentialOneOf5) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCredentialsRequestCredentialOneOf5) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCredentialsRequestCredentialOneOf5) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddCredentialsRequestCredentialOneOf5) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *AddCredentialsRequestCredentialOneOf5) GetIntegration() AddCredentialsRequestCredentialOneOfIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *AddCredentialsRequestCredentialOneOf5) GetIntegrationOk() (*AddCredentialsRequestCredentialOneOfIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *AddCredentialsRequestCredentialOneOf5) SetIntegration(v AddCredentialsRequestCredentialOneOfIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *AddCredentialsRequestCredentialOneOf5) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetUsername

`func (o *AddCredentialsRequestCredentialOneOf5) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddCredentialsRequestCredentialOneOf5) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddCredentialsRequestCredentialOneOf5) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAuthKey

`func (o *AddCredentialsRequestCredentialOneOf5) GetAuthKey() AddCredentialsRequestCredentialOneOf2AuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *AddCredentialsRequestCredentialOneOf5) GetAuthKeyOk() (*AddCredentialsRequestCredentialOneOf2AuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *AddCredentialsRequestCredentialOneOf5) SetAuthKey(v AddCredentialsRequestCredentialOneOf2AuthKey)`

SetAuthKey sets AuthKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


