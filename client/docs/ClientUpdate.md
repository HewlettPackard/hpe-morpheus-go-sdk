# ClientUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** |  | [optional] 
**AccessTokenValiditySeconds** | Pointer to **int64** |  | [optional] 
**RefreshTokenValiditySeconds** | Pointer to **int64** |  | [optional] 
**RedirectUris** | Pointer to **[]string** | List of Redirect URIs for use with the OpenID Authorization Code Flow | [optional] 

## Methods

### NewClientUpdate

`func NewClientUpdate() *ClientUpdate`

NewClientUpdate instantiates a new ClientUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientUpdateWithDefaults

`func NewClientUpdateWithDefaults() *ClientUpdate`

NewClientUpdateWithDefaults instantiates a new ClientUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ClientUpdate) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientUpdate) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientUpdate) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ClientUpdate) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAccessTokenValiditySeconds

`func (o *ClientUpdate) GetAccessTokenValiditySeconds() int64`

GetAccessTokenValiditySeconds returns the AccessTokenValiditySeconds field if non-nil, zero value otherwise.

### GetAccessTokenValiditySecondsOk

`func (o *ClientUpdate) GetAccessTokenValiditySecondsOk() (*int64, bool)`

GetAccessTokenValiditySecondsOk returns a tuple with the AccessTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValiditySeconds

`func (o *ClientUpdate) SetAccessTokenValiditySeconds(v int64)`

SetAccessTokenValiditySeconds sets AccessTokenValiditySeconds field to given value.

### HasAccessTokenValiditySeconds

`func (o *ClientUpdate) HasAccessTokenValiditySeconds() bool`

HasAccessTokenValiditySeconds returns a boolean if a field has been set.

### GetRefreshTokenValiditySeconds

`func (o *ClientUpdate) GetRefreshTokenValiditySeconds() int64`

GetRefreshTokenValiditySeconds returns the RefreshTokenValiditySeconds field if non-nil, zero value otherwise.

### GetRefreshTokenValiditySecondsOk

`func (o *ClientUpdate) GetRefreshTokenValiditySecondsOk() (*int64, bool)`

GetRefreshTokenValiditySecondsOk returns a tuple with the RefreshTokenValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenValiditySeconds

`func (o *ClientUpdate) SetRefreshTokenValiditySeconds(v int64)`

SetRefreshTokenValiditySeconds sets RefreshTokenValiditySeconds field to given value.

### HasRefreshTokenValiditySeconds

`func (o *ClientUpdate) HasRefreshTokenValiditySeconds() bool`

HasRefreshTokenValiditySeconds returns a boolean if a field has been set.

### GetRedirectUris

`func (o *ClientUpdate) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ClientUpdate) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ClientUpdate) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ClientUpdate) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


