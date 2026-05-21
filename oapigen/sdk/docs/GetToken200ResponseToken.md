# GetToken200ResponseToken

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

### NewGetToken200ResponseToken

`func NewGetToken200ResponseToken() *GetToken200ResponseToken`

NewGetToken200ResponseToken instantiates a new GetToken200ResponseToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetToken200ResponseTokenWithDefaults

`func NewGetToken200ResponseTokenWithDefaults() *GetToken200ResponseToken`

NewGetToken200ResponseTokenWithDefaults instantiates a new GetToken200ResponseToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetToken200ResponseToken) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetToken200ResponseToken) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetToken200ResponseToken) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetToken200ResponseToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetToken200ResponseToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetToken200ResponseToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetToken200ResponseToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetToken200ResponseToken) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetToken200ResponseToken) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetToken200ResponseToken) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetClientId

`func (o *GetToken200ResponseToken) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetToken200ResponseToken) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetToken200ResponseToken) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetToken200ResponseToken) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetUsername

`func (o *GetToken200ResponseToken) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetToken200ResponseToken) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetToken200ResponseToken) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetToken200ResponseToken) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetExpiration

`func (o *GetToken200ResponseToken) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *GetToken200ResponseToken) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *GetToken200ResponseToken) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *GetToken200ResponseToken) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *GetToken200ResponseToken) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *GetToken200ResponseToken) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetTokenType

`func (o *GetToken200ResponseToken) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *GetToken200ResponseToken) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *GetToken200ResponseToken) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *GetToken200ResponseToken) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetScope

`func (o *GetToken200ResponseToken) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetToken200ResponseToken) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetToken200ResponseToken) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetToken200ResponseToken) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetMaskedAccessToken

`func (o *GetToken200ResponseToken) GetMaskedAccessToken() string`

GetMaskedAccessToken returns the MaskedAccessToken field if non-nil, zero value otherwise.

### GetMaskedAccessTokenOk

`func (o *GetToken200ResponseToken) GetMaskedAccessTokenOk() (*string, bool)`

GetMaskedAccessTokenOk returns a tuple with the MaskedAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedAccessToken

`func (o *GetToken200ResponseToken) SetMaskedAccessToken(v string)`

SetMaskedAccessToken sets MaskedAccessToken field to given value.

### HasMaskedAccessToken

`func (o *GetToken200ResponseToken) HasMaskedAccessToken() bool`

HasMaskedAccessToken returns a boolean if a field has been set.

### GetMaskedRefreshToken

`func (o *GetToken200ResponseToken) GetMaskedRefreshToken() string`

GetMaskedRefreshToken returns the MaskedRefreshToken field if non-nil, zero value otherwise.

### GetMaskedRefreshTokenOk

`func (o *GetToken200ResponseToken) GetMaskedRefreshTokenOk() (*string, bool)`

GetMaskedRefreshTokenOk returns a tuple with the MaskedRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedRefreshToken

`func (o *GetToken200ResponseToken) SetMaskedRefreshToken(v string)`

SetMaskedRefreshToken sets MaskedRefreshToken field to given value.

### HasMaskedRefreshToken

`func (o *GetToken200ResponseToken) HasMaskedRefreshToken() bool`

HasMaskedRefreshToken returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetToken200ResponseToken) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetToken200ResponseToken) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetToken200ResponseToken) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetToken200ResponseToken) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *GetToken200ResponseToken) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *GetToken200ResponseToken) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


