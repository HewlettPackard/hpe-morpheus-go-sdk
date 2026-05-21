# ListTokens200ResponseAllOfTokensInner

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

### NewListTokens200ResponseAllOfTokensInner

`func NewListTokens200ResponseAllOfTokensInner() *ListTokens200ResponseAllOfTokensInner`

NewListTokens200ResponseAllOfTokensInner instantiates a new ListTokens200ResponseAllOfTokensInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTokens200ResponseAllOfTokensInnerWithDefaults

`func NewListTokens200ResponseAllOfTokensInnerWithDefaults() *ListTokens200ResponseAllOfTokensInner`

NewListTokens200ResponseAllOfTokensInnerWithDefaults instantiates a new ListTokens200ResponseAllOfTokensInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListTokens200ResponseAllOfTokensInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListTokens200ResponseAllOfTokensInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListTokens200ResponseAllOfTokensInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListTokens200ResponseAllOfTokensInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListTokens200ResponseAllOfTokensInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListTokens200ResponseAllOfTokensInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListTokens200ResponseAllOfTokensInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListTokens200ResponseAllOfTokensInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ListTokens200ResponseAllOfTokensInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ListTokens200ResponseAllOfTokensInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetClientId

`func (o *ListTokens200ResponseAllOfTokensInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ListTokens200ResponseAllOfTokensInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ListTokens200ResponseAllOfTokensInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ListTokens200ResponseAllOfTokensInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetUsername

`func (o *ListTokens200ResponseAllOfTokensInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListTokens200ResponseAllOfTokensInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListTokens200ResponseAllOfTokensInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListTokens200ResponseAllOfTokensInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetExpiration

`func (o *ListTokens200ResponseAllOfTokensInner) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ListTokens200ResponseAllOfTokensInner) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ListTokens200ResponseAllOfTokensInner) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ListTokens200ResponseAllOfTokensInner) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *ListTokens200ResponseAllOfTokensInner) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *ListTokens200ResponseAllOfTokensInner) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetTokenType

`func (o *ListTokens200ResponseAllOfTokensInner) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ListTokens200ResponseAllOfTokensInner) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ListTokens200ResponseAllOfTokensInner) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *ListTokens200ResponseAllOfTokensInner) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetScope

`func (o *ListTokens200ResponseAllOfTokensInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ListTokens200ResponseAllOfTokensInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ListTokens200ResponseAllOfTokensInner) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ListTokens200ResponseAllOfTokensInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetMaskedAccessToken

`func (o *ListTokens200ResponseAllOfTokensInner) GetMaskedAccessToken() string`

GetMaskedAccessToken returns the MaskedAccessToken field if non-nil, zero value otherwise.

### GetMaskedAccessTokenOk

`func (o *ListTokens200ResponseAllOfTokensInner) GetMaskedAccessTokenOk() (*string, bool)`

GetMaskedAccessTokenOk returns a tuple with the MaskedAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedAccessToken

`func (o *ListTokens200ResponseAllOfTokensInner) SetMaskedAccessToken(v string)`

SetMaskedAccessToken sets MaskedAccessToken field to given value.

### HasMaskedAccessToken

`func (o *ListTokens200ResponseAllOfTokensInner) HasMaskedAccessToken() bool`

HasMaskedAccessToken returns a boolean if a field has been set.

### GetMaskedRefreshToken

`func (o *ListTokens200ResponseAllOfTokensInner) GetMaskedRefreshToken() string`

GetMaskedRefreshToken returns the MaskedRefreshToken field if non-nil, zero value otherwise.

### GetMaskedRefreshTokenOk

`func (o *ListTokens200ResponseAllOfTokensInner) GetMaskedRefreshTokenOk() (*string, bool)`

GetMaskedRefreshTokenOk returns a tuple with the MaskedRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedRefreshToken

`func (o *ListTokens200ResponseAllOfTokensInner) SetMaskedRefreshToken(v string)`

SetMaskedRefreshToken sets MaskedRefreshToken field to given value.

### HasMaskedRefreshToken

`func (o *ListTokens200ResponseAllOfTokensInner) HasMaskedRefreshToken() bool`

HasMaskedRefreshToken returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListTokens200ResponseAllOfTokensInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListTokens200ResponseAllOfTokensInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListTokens200ResponseAllOfTokensInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListTokens200ResponseAllOfTokensInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ListTokens200ResponseAllOfTokensInner) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ListTokens200ResponseAllOfTokensInner) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


