# TokenAuthorizationCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client ID | 
**ClientSecret** | **string** | Client Secret | 
**GrantType** | **string** | OAuth Grant Type, use authorization_code. | [default to "authorization_code"]
**AuthorizationCode** | **string** | Authorization code must be obtained with a valid request to &#x60;/oauth/authorize&#x60;. This code is used to request an access token in the OAuth 2.0 Authorization Code Flow. | 

## Methods

### NewTokenAuthorizationCodeRequest

`func NewTokenAuthorizationCodeRequest(clientId string, clientSecret string, grantType string, authorizationCode string, ) *TokenAuthorizationCodeRequest`

NewTokenAuthorizationCodeRequest instantiates a new TokenAuthorizationCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenAuthorizationCodeRequestWithDefaults

`func NewTokenAuthorizationCodeRequestWithDefaults() *TokenAuthorizationCodeRequest`

NewTokenAuthorizationCodeRequestWithDefaults instantiates a new TokenAuthorizationCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TokenAuthorizationCodeRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenAuthorizationCodeRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenAuthorizationCodeRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *TokenAuthorizationCodeRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *TokenAuthorizationCodeRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *TokenAuthorizationCodeRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetGrantType

`func (o *TokenAuthorizationCodeRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *TokenAuthorizationCodeRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *TokenAuthorizationCodeRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetAuthorizationCode

`func (o *TokenAuthorizationCodeRequest) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *TokenAuthorizationCodeRequest) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *TokenAuthorizationCodeRequest) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


