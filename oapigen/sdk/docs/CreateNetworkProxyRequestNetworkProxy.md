# CreateNetworkProxyRequestNetworkProxy

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
**Account** | Pointer to [**CreateNetworkProxyRequestNetworkProxyAccount**](CreateNetworkProxyRequestNetworkProxyAccount.md) |  | [optional] 

## Methods

### NewCreateNetworkProxyRequestNetworkProxy

`func NewCreateNetworkProxyRequestNetworkProxy() *CreateNetworkProxyRequestNetworkProxy`

NewCreateNetworkProxyRequestNetworkProxy instantiates a new CreateNetworkProxyRequestNetworkProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkProxyRequestNetworkProxyWithDefaults

`func NewCreateNetworkProxyRequestNetworkProxyWithDefaults() *CreateNetworkProxyRequestNetworkProxy`

NewCreateNetworkProxyRequestNetworkProxyWithDefaults instantiates a new CreateNetworkProxyRequestNetworkProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkProxyRequestNetworkProxy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkProxyRequestNetworkProxy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkProxyRequestNetworkProxy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateNetworkProxyRequestNetworkProxy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProxyHost

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyHost() string`

GetProxyHost returns the ProxyHost field if non-nil, zero value otherwise.

### GetProxyHostOk

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyHostOk() (*string, bool)`

GetProxyHostOk returns a tuple with the ProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHost

`func (o *CreateNetworkProxyRequestNetworkProxy) SetProxyHost(v string)`

SetProxyHost sets ProxyHost field to given value.

### HasProxyHost

`func (o *CreateNetworkProxyRequestNetworkProxy) HasProxyHost() bool`

HasProxyHost returns a boolean if a field has been set.

### GetProxyPort

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyPort() string`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyPortOk() (*string, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *CreateNetworkProxyRequestNetworkProxy) SetProxyPort(v string)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *CreateNetworkProxyRequestNetworkProxy) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyUser

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *CreateNetworkProxyRequestNetworkProxy) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *CreateNetworkProxyRequestNetworkProxy) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### GetProxyPassword

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *CreateNetworkProxyRequestNetworkProxy) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *CreateNetworkProxyRequestNetworkProxy) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetProxyDomain

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyDomain() string`

GetProxyDomain returns the ProxyDomain field if non-nil, zero value otherwise.

### GetProxyDomainOk

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyDomainOk() (*string, bool)`

GetProxyDomainOk returns a tuple with the ProxyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyDomain

`func (o *CreateNetworkProxyRequestNetworkProxy) SetProxyDomain(v string)`

SetProxyDomain sets ProxyDomain field to given value.

### HasProxyDomain

`func (o *CreateNetworkProxyRequestNetworkProxy) HasProxyDomain() bool`

HasProxyDomain returns a boolean if a field has been set.

### GetProxyWorkstation

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyWorkstation() string`

GetProxyWorkstation returns the ProxyWorkstation field if non-nil, zero value otherwise.

### GetProxyWorkstationOk

`func (o *CreateNetworkProxyRequestNetworkProxy) GetProxyWorkstationOk() (*string, bool)`

GetProxyWorkstationOk returns a tuple with the ProxyWorkstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyWorkstation

`func (o *CreateNetworkProxyRequestNetworkProxy) SetProxyWorkstation(v string)`

SetProxyWorkstation sets ProxyWorkstation field to given value.

### HasProxyWorkstation

`func (o *CreateNetworkProxyRequestNetworkProxy) HasProxyWorkstation() bool`

HasProxyWorkstation returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateNetworkProxyRequestNetworkProxy) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateNetworkProxyRequestNetworkProxy) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateNetworkProxyRequestNetworkProxy) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateNetworkProxyRequestNetworkProxy) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *CreateNetworkProxyRequestNetworkProxy) GetAccount() CreateNetworkProxyRequestNetworkProxyAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateNetworkProxyRequestNetworkProxy) GetAccountOk() (*CreateNetworkProxyRequestNetworkProxyAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateNetworkProxyRequestNetworkProxy) SetAccount(v CreateNetworkProxyRequestNetworkProxyAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreateNetworkProxyRequestNetworkProxy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


