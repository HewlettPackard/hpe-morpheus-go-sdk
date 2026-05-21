# AddToken200ResponseAllOfToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** | Username associated with the token. Sub-tenant users will have their username formatted as &#x60;subdomain\\username&#x60; with a prefix that is the tenant subdomain or id by default. | [optional] 
**Expiration** | Pointer to **NullableTime** | Expiration date of the token | [optional] 
**TokenType** | Pointer to **string** | Type of the token | [optional] 
**Scope** | Pointer to **string** | Authorized scope(s), separated by spaces. Either &#x60;write&#x60; or &#x60;write openid&#x60;. | [optional] 
**MaskedAccessToken** | Pointer to **string** | Masked Access Token, with all but the first 8 characters replaced by asterisks for security. | [optional] 
**MaskedRefreshToken** | Pointer to **string** | Masked Refresh Token, with all but the first 8 characters replaced by asterisks for security. | [optional] 
**DateCreated** | Pointer to **NullableTime** | Date the token was created | [optional] 

## Methods

### NewAddToken200ResponseAllOfToken

`func NewAddToken200ResponseAllOfToken() *AddToken200ResponseAllOfToken`

NewAddToken200ResponseAllOfToken instantiates a new AddToken200ResponseAllOfToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddToken200ResponseAllOfTokenWithDefaults

`func NewAddToken200ResponseAllOfTokenWithDefaults() *AddToken200ResponseAllOfToken`

NewAddToken200ResponseAllOfTokenWithDefaults instantiates a new AddToken200ResponseAllOfToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddToken200ResponseAllOfToken) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddToken200ResponseAllOfToken) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddToken200ResponseAllOfToken) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddToken200ResponseAllOfToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddToken200ResponseAllOfToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddToken200ResponseAllOfToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddToken200ResponseAllOfToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddToken200ResponseAllOfToken) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AddToken200ResponseAllOfToken) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AddToken200ResponseAllOfToken) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetClientId

`func (o *AddToken200ResponseAllOfToken) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AddToken200ResponseAllOfToken) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AddToken200ResponseAllOfToken) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AddToken200ResponseAllOfToken) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetUsername

`func (o *AddToken200ResponseAllOfToken) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddToken200ResponseAllOfToken) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddToken200ResponseAllOfToken) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddToken200ResponseAllOfToken) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetExpiration

`func (o *AddToken200ResponseAllOfToken) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *AddToken200ResponseAllOfToken) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *AddToken200ResponseAllOfToken) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *AddToken200ResponseAllOfToken) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *AddToken200ResponseAllOfToken) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *AddToken200ResponseAllOfToken) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetTokenType

`func (o *AddToken200ResponseAllOfToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AddToken200ResponseAllOfToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AddToken200ResponseAllOfToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AddToken200ResponseAllOfToken) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetScope

`func (o *AddToken200ResponseAllOfToken) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AddToken200ResponseAllOfToken) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AddToken200ResponseAllOfToken) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AddToken200ResponseAllOfToken) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetMaskedAccessToken

`func (o *AddToken200ResponseAllOfToken) GetMaskedAccessToken() string`

GetMaskedAccessToken returns the MaskedAccessToken field if non-nil, zero value otherwise.

### GetMaskedAccessTokenOk

`func (o *AddToken200ResponseAllOfToken) GetMaskedAccessTokenOk() (*string, bool)`

GetMaskedAccessTokenOk returns a tuple with the MaskedAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedAccessToken

`func (o *AddToken200ResponseAllOfToken) SetMaskedAccessToken(v string)`

SetMaskedAccessToken sets MaskedAccessToken field to given value.

### HasMaskedAccessToken

`func (o *AddToken200ResponseAllOfToken) HasMaskedAccessToken() bool`

HasMaskedAccessToken returns a boolean if a field has been set.

### GetMaskedRefreshToken

`func (o *AddToken200ResponseAllOfToken) GetMaskedRefreshToken() string`

GetMaskedRefreshToken returns the MaskedRefreshToken field if non-nil, zero value otherwise.

### GetMaskedRefreshTokenOk

`func (o *AddToken200ResponseAllOfToken) GetMaskedRefreshTokenOk() (*string, bool)`

GetMaskedRefreshTokenOk returns a tuple with the MaskedRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedRefreshToken

`func (o *AddToken200ResponseAllOfToken) SetMaskedRefreshToken(v string)`

SetMaskedRefreshToken sets MaskedRefreshToken field to given value.

### HasMaskedRefreshToken

`func (o *AddToken200ResponseAllOfToken) HasMaskedRefreshToken() bool`

HasMaskedRefreshToken returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddToken200ResponseAllOfToken) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddToken200ResponseAllOfToken) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddToken200ResponseAllOfToken) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddToken200ResponseAllOfToken) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AddToken200ResponseAllOfToken) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AddToken200ResponseAllOfToken) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


