# GetAccessToken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | Token that grants API Access | [optional] 
**RefreshToken** | Pointer to **string** | Token that can request an new API access_token | [optional] 
**ExpiresIn** | Pointer to **float32** | Seconds until token expires | [optional] 
**TokenType** | Pointer to **string** | Token type granted | [optional] 
**Scope** | Pointer to **string** | Scope granted | [optional] 

## Methods

### NewGetAccessToken200Response

`func NewGetAccessToken200Response() *GetAccessToken200Response`

NewGetAccessToken200Response instantiates a new GetAccessToken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessToken200ResponseWithDefaults

`func NewGetAccessToken200ResponseWithDefaults() *GetAccessToken200Response`

NewGetAccessToken200ResponseWithDefaults instantiates a new GetAccessToken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *GetAccessToken200Response) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GetAccessToken200Response) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GetAccessToken200Response) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *GetAccessToken200Response) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *GetAccessToken200Response) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *GetAccessToken200Response) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *GetAccessToken200Response) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *GetAccessToken200Response) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *GetAccessToken200Response) GetExpiresIn() float32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GetAccessToken200Response) GetExpiresInOk() (*float32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GetAccessToken200Response) SetExpiresIn(v float32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GetAccessToken200Response) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetTokenType

`func (o *GetAccessToken200Response) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *GetAccessToken200Response) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *GetAccessToken200Response) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *GetAccessToken200Response) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetScope

`func (o *GetAccessToken200Response) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetAccessToken200Response) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetAccessToken200Response) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetAccessToken200Response) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


