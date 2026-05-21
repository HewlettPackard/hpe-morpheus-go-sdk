# TokenRefreshTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client ID | [default to "morph-api"]
**GrantType** | **string** | OAuth Grant Type, use &#x60;refresh_token&#x60;. | [default to "refresh_token"]
**RefreshToken** | **string** | Refresh Token | 

## Methods

### NewTokenRefreshTokenRequest

`func NewTokenRefreshTokenRequest(clientId string, grantType string, refreshToken string, ) *TokenRefreshTokenRequest`

NewTokenRefreshTokenRequest instantiates a new TokenRefreshTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRefreshTokenRequestWithDefaults

`func NewTokenRefreshTokenRequestWithDefaults() *TokenRefreshTokenRequest`

NewTokenRefreshTokenRequestWithDefaults instantiates a new TokenRefreshTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TokenRefreshTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenRefreshTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenRefreshTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetGrantType

`func (o *TokenRefreshTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *TokenRefreshTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *TokenRefreshTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetRefreshToken

`func (o *TokenRefreshTokenRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenRefreshTokenRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenRefreshTokenRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


