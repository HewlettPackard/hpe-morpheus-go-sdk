# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**AccessTokenValiditySeconds** | Pointer to **int64** |  | [optional] 
**RefreshTokenValiditySeconds** | Pointer to **int64** |  | [optional] 
**Authorities** | Pointer to **[]string** |  | [optional] 
**AuthorizedGrantTypes** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 

## Methods

### NewClient

`func NewClient() *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Client) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Client) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Client) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Client) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *Client) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Client) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Client) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Client) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAccessTokenValiditySeconds

`func (o *Client) GetAccessTokenValiditySeconds() int64`

GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field if non-nil, zero value otherwise.

### GetAccessTokenValiditySecondsOk

`func (o *Client) GetAccessTokenValiditySecondsOk() (*int64, bool)`

GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValiditySeconds

`func (o *Client) SetAccessTokenValiditySeconds(v int64)`

SetAccessTokenValiditySeconds sets AccessTokenValiditySeconds field to given value.

### HasAccessTokenValiditySeconds

`func (o *Client) HasAccessTokenValiditySeconds() bool`

HasAccessTokenValiditySeconds returns a boolean if a field has been set.

### GetRefreshTokenValiditySeconds

`func (o *Client) GetRefreshTokenValiditySeconds() int64`

GetRefreshTokenValiditySeconds returns the RefreshTokenValiditySeconds field if non-nil, zero value otherwise.

### GetRefreshTokenValiditySecondsOk

`func (o *Client) GetRefreshTokenValiditySecondsOk() (*int64, bool)`

GetRefreshTokenValiditySecondsOk returns a tuple with the RefreshTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValiditySeconds

`func (o *Client) SetRefreshTokenValiditySeconds(v int64)`

SetRefreshTokenValiditySeconds sets RefreshTokenValiditySeconds field to given value.

### HasRefreshTokenValiditySeconds

`func (o *Client) HasRefreshTokenValiditySeconds() bool`

HasRefreshTokenValiditySeconds returns a boolean if a field has been set.

### GetAuthorities

`func (o *Client) GetAuthorities() []string`

GetAuthorities returns the Authorities field if non-nil, zero value otherwise.

### GetAuthoritiesOk

`func (o *Client) GetAuthoritiesOk() (*[]string, bool)`

GetAuthoritiesOk returns a tuple with the Authorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorities

`func (o *Client) SetAuthorities(v []string)`

SetAuthorities sets Authorities field to given value.

### HasAuthorities

`func (o *Client) HasAuthorities() bool`

HasAuthorities returns a boolean if a field has been set.

### GetAuthorizedGrantTypes

`func (o *Client) GetAuthorizedGrantTypes() []string`

GetAuthorizedGrantTypes returns the AuthorizedGrantTypes field if non-nil, zero value otherwise.

### GetAuthorizedGrantTypesOk

`func (o *Client) GetAuthorizedGrantTypesOk() (*[]string, bool)`

GetAuthorizedGrantTypesOk returns a tuple with the AuthorizedGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedGrantTypes

`func (o *Client) SetAuthorizedGrantTypes(v []string)`

SetAuthorizedGrantTypes sets AuthorizedGrantTypes field to given value.

### HasAuthorizedGrantTypes

`func (o *Client) HasAuthorizedGrantTypes() bool`

HasAuthorizedGrantTypes returns a boolean if a field has been set.

### GetScopes

`func (o *Client) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Client) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Client) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *Client) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetRedirectUris

`func (o *Client) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *Client) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *Client) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *Client) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


