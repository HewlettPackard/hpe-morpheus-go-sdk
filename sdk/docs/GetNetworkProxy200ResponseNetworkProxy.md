# GetNetworkProxy200ResponseNetworkProxy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProxyHost** | Pointer to **string** |  | [optional] 
**ProxyPort** | Pointer to **int64** |  | [optional] 
**ProxyUser** | Pointer to **NullableString** |  | [optional] 
**ProxyPassword** | Pointer to **NullableString** |  | [optional] 
**ProxyDomain** | Pointer to **string** |  | [optional] 
**ProxyWorkstation** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**NullableListApprovals200ResponseAllOfApprovalsInnerAccount**](ListApprovals200ResponseAllOfApprovalsInnerAccount.md) |  | [optional] 
**Owner** | Pointer to [**NullableListApprovals200ResponseAllOfApprovalsInnerAccount**](ListApprovals200ResponseAllOfApprovalsInnerAccount.md) |  | [optional] 

## Methods

### NewGetNetworkProxy200ResponseNetworkProxy

`func NewGetNetworkProxy200ResponseNetworkProxy() *GetNetworkProxy200ResponseNetworkProxy`

NewGetNetworkProxy200ResponseNetworkProxy instantiates a new GetNetworkProxy200ResponseNetworkProxy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkProxy200ResponseNetworkProxyWithDefaults

`func NewGetNetworkProxy200ResponseNetworkProxyWithDefaults() *GetNetworkProxy200ResponseNetworkProxy`

NewGetNetworkProxy200ResponseNetworkProxyWithDefaults instantiates a new GetNetworkProxy200ResponseNetworkProxy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProxyHost

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyHost() string`

GetProxyHost returns the ProxyHost field if non-nil, zero value otherwise.

### GetProxyHostOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyHostOk() (*string, bool)`

GetProxyHostOk returns a tuple with the ProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHost

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetProxyHost(v string)`

SetProxyHost sets ProxyHost field to given value.

### HasProxyHost

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasProxyHost() bool`

HasProxyHost returns a boolean if a field has been set.

### GetProxyPort

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyPort() int64`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyPortOk() (*int64, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetProxyPort(v int64)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyUser

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### SetProxyUserNil

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetProxyUserNil(b bool)`

 SetProxyUserNil sets the value for ProxyUser to be an explicit nil

### UnsetProxyUser
`func (o *GetNetworkProxy200ResponseNetworkProxy) UnsetProxyUser()`

UnsetProxyUser ensures that no value is present for ProxyUser, not even an explicit nil
### GetProxyPassword

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### SetProxyPasswordNil

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetProxyPasswordNil(b bool)`

 SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil

### UnsetProxyPassword
`func (o *GetNetworkProxy200ResponseNetworkProxy) UnsetProxyPassword()`

UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
### GetProxyDomain

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyDomain() string`

GetProxyDomain returns the ProxyDomain field if non-nil, zero value otherwise.

### GetProxyDomainOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyDomainOk() (*string, bool)`

GetProxyDomainOk returns a tuple with the ProxyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyDomain

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetProxyDomain(v string)`

SetProxyDomain sets ProxyDomain field to given value.

### HasProxyDomain

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasProxyDomain() bool`

HasProxyDomain returns a boolean if a field has been set.

### GetProxyWorkstation

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyWorkstation() string`

GetProxyWorkstation returns the ProxyWorkstation field if non-nil, zero value otherwise.

### GetProxyWorkstationOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetProxyWorkstationOk() (*string, bool)`

GetProxyWorkstationOk returns a tuple with the ProxyWorkstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyWorkstation

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetProxyWorkstation(v string)`

SetProxyWorkstation sets ProxyWorkstation field to given value.

### HasProxyWorkstation

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasProxyWorkstation() bool`

HasProxyWorkstation returns a boolean if a field has been set.

### SetProxyWorkstationNil

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetProxyWorkstationNil(b bool)`

 SetProxyWorkstationNil sets the value for ProxyWorkstation to be an explicit nil

### UnsetProxyWorkstation
`func (o *GetNetworkProxy200ResponseNetworkProxy) UnsetProxyWorkstation()`

UnsetProxyWorkstation ensures that no value is present for ProxyWorkstation, not even an explicit nil
### GetVisibility

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetAccount() ListApprovals200ResponseAllOfApprovalsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetAccountOk() (*ListApprovals200ResponseAllOfApprovalsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetAccount(v ListApprovals200ResponseAllOfApprovalsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *GetNetworkProxy200ResponseNetworkProxy) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetOwner

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetOwner() ListApprovals200ResponseAllOfApprovalsInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetNetworkProxy200ResponseNetworkProxy) GetOwnerOk() (*ListApprovals200ResponseAllOfApprovalsInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetOwner(v ListApprovals200ResponseAllOfApprovalsInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetNetworkProxy200ResponseNetworkProxy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *GetNetworkProxy200ResponseNetworkProxy) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *GetNetworkProxy200ResponseNetworkProxy) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


