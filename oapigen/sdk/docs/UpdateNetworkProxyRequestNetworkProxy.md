# UpdateNetworkProxyRequestNetworkProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**ProxyHost** | Pointer to **string** | Proxy Host | [optional] 
**ProxyPort** | Pointer to **string** | Proxy Port | [optional] 
**ProxyUser** | Pointer to **string** | Proxy Username | [optional] 
**ProxyPassword** | Pointer to **string** | Proxy Password | [optional] 
**ProxyDomain** | Pointer to **string** | Proxy Domain | [optional] 
**ProxyWorkstation** | Pointer to **string** | Proxy Workstation | [optional] 
**Visibility** | Pointer to **string** | Visibility | [optional] [default to "private"]
**Account** | Pointer to [**UpdateNetworkProxyRequestNetworkProxyAccount**](UpdateNetworkProxyRequestNetworkProxyAccount.md) |  | [optional] 

## Methods

### NewUpdateNetworkProxyRequestNetworkProxy

`func NewUpdateNetworkProxyRequestNetworkProxy() *UpdateNetworkProxyRequestNetworkProxy`

NewUpdateNetworkProxyRequestNetworkProxy instantiates a new UpdateNetworkProxyRequestNetworkProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkProxyRequestNetworkProxyWithDefaults

`func NewUpdateNetworkProxyRequestNetworkProxyWithDefaults() *UpdateNetworkProxyRequestNetworkProxy`

NewUpdateNetworkProxyRequestNetworkProxyWithDefaults instantiates a new UpdateNetworkProxyRequestNetworkProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkProxyRequestNetworkProxy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkProxyRequestNetworkProxy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProxyHost

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyHost() string`

GetProxyHost returns the ProxyHost field if non-nil, zero value otherwise.

### GetProxyHostOk

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyHostOk() (*string, bool)`

GetProxyHostOk returns a tuple with the ProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHost

`func (o *UpdateNetworkProxyRequestNetworkProxy) SetProxyHost(v string)`

SetProxyHost sets ProxyHost field to given value.

### HasProxyHost

`func (o *UpdateNetworkProxyRequestNetworkProxy) HasProxyHost() bool`

HasProxyHost returns a boolean if a field has been set.

### GetProxyPort

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyPort() string`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyPortOk() (*string, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *UpdateNetworkProxyRequestNetworkProxy) SetProxyPort(v string)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *UpdateNetworkProxyRequestNetworkProxy) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyUser

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *UpdateNetworkProxyRequestNetworkProxy) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *UpdateNetworkProxyRequestNetworkProxy) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### GetProxyPassword

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *UpdateNetworkProxyRequestNetworkProxy) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *UpdateNetworkProxyRequestNetworkProxy) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetProxyDomain

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyDomain() string`

GetProxyDomain returns the ProxyDomain field if non-nil, zero value otherwise.

### GetProxyDomainOk

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyDomainOk() (*string, bool)`

GetProxyDomainOk returns a tuple with the ProxyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyDomain

`func (o *UpdateNetworkProxyRequestNetworkProxy) SetProxyDomain(v string)`

SetProxyDomain sets ProxyDomain field to given value.

### HasProxyDomain

`func (o *UpdateNetworkProxyRequestNetworkProxy) HasProxyDomain() bool`

HasProxyDomain returns a boolean if a field has been set.

### GetProxyWorkstation

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyWorkstation() string`

GetProxyWorkstation returns the ProxyWorkstation field if non-nil, zero value otherwise.

### GetProxyWorkstationOk

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetProxyWorkstationOk() (*string, bool)`

GetProxyWorkstationOk returns a tuple with the ProxyWorkstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyWorkstation

`func (o *UpdateNetworkProxyRequestNetworkProxy) SetProxyWorkstation(v string)`

SetProxyWorkstation sets ProxyWorkstation field to given value.

### HasProxyWorkstation

`func (o *UpdateNetworkProxyRequestNetworkProxy) HasProxyWorkstation() bool`

HasProxyWorkstation returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateNetworkProxyRequestNetworkProxy) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateNetworkProxyRequestNetworkProxy) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetAccount() UpdateNetworkProxyRequestNetworkProxyAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateNetworkProxyRequestNetworkProxy) GetAccountOk() (*UpdateNetworkProxyRequestNetworkProxyAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateNetworkProxyRequestNetworkProxy) SetAccount(v UpdateNetworkProxyRequestNetworkProxyAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateNetworkProxyRequestNetworkProxy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


