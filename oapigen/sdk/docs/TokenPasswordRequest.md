# TokenPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client ID | [default to "morph-api"]
**GrantType** | **string** | OAuth Grant Type, use &#x60;password&#x60;. | [default to "password"]
**Scope** | **string** | OAuth token scope, use &#x60;write&#x60;. | 
**Username** | **string** | Username Sub-tenant users must format their username as &#x60;subdomain\\username&#x60; with a prefix that is the tenant subdomain or id by default.  | 
**Password** | **string** | Password | 

## Methods

### NewTokenPasswordRequest

`func NewTokenPasswordRequest(clientId string, grantType string, scope string, username string, password string, ) *TokenPasswordRequest`

NewTokenPasswordRequest instantiates a new TokenPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetClientId

`func (o *TokenPasswordRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TokenPasswordRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TokenPasswordRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetGrantType

`func (o *TokenPasswordRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *TokenPasswordRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *TokenPasswordRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetScope

`func (o *TokenPasswordRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TokenPasswordRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TokenPasswordRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetUsername

`func (o *TokenPasswordRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TokenPasswordRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TokenPasswordRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *TokenPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TokenPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TokenPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


