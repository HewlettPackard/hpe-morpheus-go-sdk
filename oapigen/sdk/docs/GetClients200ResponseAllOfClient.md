# GetClients200ResponseAllOfClient

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

### NewGetClients200ResponseAllOfClient

`func NewGetClients200ResponseAllOfClient() *GetClients200ResponseAllOfClient`

NewGetClients200ResponseAllOfClient instantiates a new GetClients200ResponseAllOfClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClients200ResponseAllOfClientWithDefaults

`func NewGetClients200ResponseAllOfClientWithDefaults() *GetClients200ResponseAllOfClient`

NewGetClients200ResponseAllOfClientWithDefaults instantiates a new GetClients200ResponseAllOfClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClients200ResponseAllOfClient) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClients200ResponseAllOfClient) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClients200ResponseAllOfClient) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClients200ResponseAllOfClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *GetClients200ResponseAllOfClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetClients200ResponseAllOfClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetClients200ResponseAllOfClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetClients200ResponseAllOfClient) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAccessTokenValiditySeconds

`func (o *GetClients200ResponseAllOfClient) GetAccessTokenValiditySeconds() int64`

GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field if non-nil, zero value otherwise.

### GetAccessTokenValiditySecondsOk

`func (o *GetClients200ResponseAllOfClient) GetAccessTokenValiditySecondsOk() (*int64, bool)`

GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValiditySeconds

`func (o *GetClients200ResponseAllOfClient) SetAccessTokenValiditySeconds(v int64)`

SetAccessTokenValiditySeconds sets AccessTokenValiditySeconds field to given value.

### HasAccessTokenValiditySeconds

`func (o *GetClients200ResponseAllOfClient) HasAccessTokenValiditySeconds() bool`

HasAccessTokenValiditySeconds returns a boolean if a field has been set.

### GetRefreshTokenValiditySeconds

`func (o *GetClients200ResponseAllOfClient) GetRefreshTokenValiditySeconds() int64`

GetRefreshTokenValiditySeconds returns the RefreshTokenValiditySeconds field if non-nil, zero value otherwise.

### GetRefreshTokenValiditySecondsOk

`func (o *GetClients200ResponseAllOfClient) GetRefreshTokenValiditySecondsOk() (*int64, bool)`

GetRefreshTokenValiditySecondsOk returns a tuple with the RefreshTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValiditySeconds

`func (o *GetClients200ResponseAllOfClient) SetRefreshTokenValiditySeconds(v int64)`

SetRefreshTokenValiditySeconds sets RefreshTokenValiditySeconds field to given value.

### HasRefreshTokenValiditySeconds

`func (o *GetClients200ResponseAllOfClient) HasRefreshTokenValiditySeconds() bool`

HasRefreshTokenValiditySeconds returns a boolean if a field has been set.

### GetAuthorities

`func (o *GetClients200ResponseAllOfClient) GetAuthorities() []string`

GetAuthorities returns the Authorities field if non-nil, zero value otherwise.

### GetAuthoritiesOk

`func (o *GetClients200ResponseAllOfClient) GetAuthoritiesOk() (*[]string, bool)`

GetAuthoritiesOk returns a tuple with the Authorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorities

`func (o *GetClients200ResponseAllOfClient) SetAuthorities(v []string)`

SetAuthorities sets Authorities field to given value.

### HasAuthorities

`func (o *GetClients200ResponseAllOfClient) HasAuthorities() bool`

HasAuthorities returns a boolean if a field has been set.

### GetAuthorizedGrantTypes

`func (o *GetClients200ResponseAllOfClient) GetAuthorizedGrantTypes() []string`

GetAuthorizedGrantTypes returns the AuthorizedGrantTypes field if non-nil, zero value otherwise.

### GetAuthorizedGrantTypesOk

`func (o *GetClients200ResponseAllOfClient) GetAuthorizedGrantTypesOk() (*[]string, bool)`

GetAuthorizedGrantTypesOk returns a tuple with the AuthorizedGrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedGrantTypes

`func (o *GetClients200ResponseAllOfClient) SetAuthorizedGrantTypes(v []string)`

SetAuthorizedGrantTypes sets AuthorizedGrantTypes field to given value.

### HasAuthorizedGrantTypes

`func (o *GetClients200ResponseAllOfClient) HasAuthorizedGrantTypes() bool`

HasAuthorizedGrantTypes returns a boolean if a field has been set.

### GetScopes

`func (o *GetClients200ResponseAllOfClient) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetClients200ResponseAllOfClient) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetClients200ResponseAllOfClient) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GetClients200ResponseAllOfClient) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetRedirectUris

`func (o *GetClients200ResponseAllOfClient) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *GetClients200ResponseAllOfClient) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *GetClients200ResponseAllOfClient) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *GetClients200ResponseAllOfClient) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


