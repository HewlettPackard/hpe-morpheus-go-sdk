# UpdateCredentialsRequestCredentialOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Credential Type Code | 
**Name** | **string** | A unique name scoped to your account for the credential | 
**Description** | Pointer to **string** | Optional Description | [optional] 
**Enabled** | Pointer to **bool** | Credential enabled | [optional] [default to true]
**Integration** | Pointer to [**UpdateCredentialsRequestCredentialOneOf3Integration**](UpdateCredentialsRequestCredentialOneOf3Integration.md) |  | [optional] 
**AuthPath** | **string** | Tenant name | 
**Username** | **string** | Username | 
**AuthKey** | [**UpdateCredentialsRequestCredentialOneOf3AuthKey**](UpdateCredentialsRequestCredentialOneOf3AuthKey.md) |  | 

## Methods

### NewUpdateCredentialsRequestCredentialOneOf3

`func NewUpdateCredentialsRequestCredentialOneOf3(type_ string, name string, authPath string, username string, authKey UpdateCredentialsRequestCredentialOneOf3AuthKey, ) *UpdateCredentialsRequestCredentialOneOf3`

NewUpdateCredentialsRequestCredentialOneOf3 instantiates a new UpdateCredentialsRequestCredentialOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCredentialsRequestCredentialOneOf3WithDefaults

`func NewUpdateCredentialsRequestCredentialOneOf3WithDefaults() *UpdateCredentialsRequestCredentialOneOf3`

NewUpdateCredentialsRequestCredentialOneOf3WithDefaults instantiates a new UpdateCredentialsRequestCredentialOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCredentialsRequestCredentialOneOf3) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCredentialsRequestCredentialOneOf3) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCredentialsRequestCredentialOneOf3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCredentialsRequestCredentialOneOf3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCredentialsRequestCredentialOneOf3) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCredentialsRequestCredentialOneOf3) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIntegration

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetIntegration() UpdateCredentialsRequestCredentialOneOf3Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetIntegrationOk() (*UpdateCredentialsRequestCredentialOneOf3Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UpdateCredentialsRequestCredentialOneOf3) SetIntegration(v UpdateCredentialsRequestCredentialOneOf3Integration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *UpdateCredentialsRequestCredentialOneOf3) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetAuthPath

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetAuthPath() string`

GetAuthPath returns the AuthPath field if non-nil, zero value otherwise.

### GetAuthPathOk

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetAuthPathOk() (*string, bool)`

GetAuthPathOk returns a tuple with the AuthPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPath

`func (o *UpdateCredentialsRequestCredentialOneOf3) SetAuthPath(v string)`

SetAuthPath sets AuthPath field to given value.


### GetUsername

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateCredentialsRequestCredentialOneOf3) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAuthKey

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetAuthKey() UpdateCredentialsRequestCredentialOneOf3AuthKey`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *UpdateCredentialsRequestCredentialOneOf3) GetAuthKeyOk() (*UpdateCredentialsRequestCredentialOneOf3AuthKey, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *UpdateCredentialsRequestCredentialOneOf3) SetAuthKey(v UpdateCredentialsRequestCredentialOneOf3AuthKey)`

SetAuthKey sets AuthKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


