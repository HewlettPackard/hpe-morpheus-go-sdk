# AddBaremetalHostRequestServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IloIpAddress** | Pointer to **string** |  | [optional] 
**IloUsername** | Pointer to **string** |  | [optional] 
**IloPassword** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**ResourcePoolId** | Pointer to **int32** |  | [optional] 
**PreProvisioned** | Pointer to **bool** | Set to &#39;true&#39; if the server is pre-provisioned (brownfield) | [optional] 
**WindowsOS** | Pointer to **bool** | Set to &#39;true&#39; if preProvisioned is &#39;true&#39; and the server has a Windows OS installed | [optional] 
**OsIpAddress** | Pointer to **string** | The IP address of the OS installed on the server. Required if preProvisioned is &#39;true&#39;. | [optional] 
**OsUsername** | Pointer to **string** | The username to access the OS installed on the server. Required if preProvisioned is &#39;true&#39;. | [optional] 
**OsPassword** | Pointer to **string** | The password to access the OS installed on the server. Required if preProvisioned is &#39;true&#39;. | [optional] 
**OsSSHKeyId** | Pointer to **int32** | ID of the SSH Key Pair needed to access the OS installed on the server. Required if preProvisioned is &#39;true&#39; and the OS is Linux and password authentication is disabled.  | [optional] 

## Methods

### NewAddBaremetalHostRequestServerConfig

`func NewAddBaremetalHostRequestServerConfig() *AddBaremetalHostRequestServerConfig`

NewAddBaremetalHostRequestServerConfig instantiates a new AddBaremetalHostRequestServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBaremetalHostRequestServerConfigWithDefaults

`func NewAddBaremetalHostRequestServerConfigWithDefaults() *AddBaremetalHostRequestServerConfig`

NewAddBaremetalHostRequestServerConfigWithDefaults instantiates a new AddBaremetalHostRequestServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIloIpAddress

`func (o *AddBaremetalHostRequestServerConfig) GetIloIpAddress() string`

GetIloIpAddress returns the IloIpAddress field if non-nil, zero value otherwise.

### GetIloIpAddressOk

`func (o *AddBaremetalHostRequestServerConfig) GetIloIpAddressOk() (*string, bool)`

GetIloIpAddressOk returns a tuple with the IloIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIloIpAddress

`func (o *AddBaremetalHostRequestServerConfig) SetIloIpAddress(v string)`

SetIloIpAddress sets IloIpAddress field to given value.

### HasIloIpAddress

`func (o *AddBaremetalHostRequestServerConfig) HasIloIpAddress() bool`

HasIloIpAddress returns a boolean if a field has been set.

### GetIloUsername

`func (o *AddBaremetalHostRequestServerConfig) GetIloUsername() string`

GetIloUsername returns the IloUsername field if non-nil, zero value otherwise.

### GetIloUsernameOk

`func (o *AddBaremetalHostRequestServerConfig) GetIloUsernameOk() (*string, bool)`

GetIloUsernameOk returns a tuple with the IloUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIloUsername

`func (o *AddBaremetalHostRequestServerConfig) SetIloUsername(v string)`

SetIloUsername sets IloUsername field to given value.

### HasIloUsername

`func (o *AddBaremetalHostRequestServerConfig) HasIloUsername() bool`

HasIloUsername returns a boolean if a field has been set.

### GetIloPassword

`func (o *AddBaremetalHostRequestServerConfig) GetIloPassword() string`

GetIloPassword returns the IloPassword field if non-nil, zero value otherwise.

### GetIloPasswordOk

`func (o *AddBaremetalHostRequestServerConfig) GetIloPasswordOk() (*string, bool)`

GetIloPasswordOk returns a tuple with the IloPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIloPassword

`func (o *AddBaremetalHostRequestServerConfig) SetIloPassword(v string)`

SetIloPassword sets IloPassword field to given value.

### HasIloPassword

`func (o *AddBaremetalHostRequestServerConfig) HasIloPassword() bool`

HasIloPassword returns a boolean if a field has been set.

### GetMacAddress

`func (o *AddBaremetalHostRequestServerConfig) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AddBaremetalHostRequestServerConfig) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AddBaremetalHostRequestServerConfig) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AddBaremetalHostRequestServerConfig) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *AddBaremetalHostRequestServerConfig) GetResourcePoolId() int32`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AddBaremetalHostRequestServerConfig) GetResourcePoolIdOk() (*int32, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AddBaremetalHostRequestServerConfig) SetResourcePoolId(v int32)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AddBaremetalHostRequestServerConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetPreProvisioned

`func (o *AddBaremetalHostRequestServerConfig) GetPreProvisioned() bool`

GetPreProvisioned returns the PreProvisioned field if non-nil, zero value otherwise.

### GetPreProvisionedOk

`func (o *AddBaremetalHostRequestServerConfig) GetPreProvisionedOk() (*bool, bool)`

GetPreProvisionedOk returns a tuple with the PreProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProvisioned

`func (o *AddBaremetalHostRequestServerConfig) SetPreProvisioned(v bool)`

SetPreProvisioned sets PreProvisioned field to given value.

### HasPreProvisioned

`func (o *AddBaremetalHostRequestServerConfig) HasPreProvisioned() bool`

HasPreProvisioned returns a boolean if a field has been set.

### GetWindowsOS

`func (o *AddBaremetalHostRequestServerConfig) GetWindowsOS() bool`

GetWindowsOS returns the WindowsOS field if non-nil, zero value otherwise.

### GetWindowsOSOk

`func (o *AddBaremetalHostRequestServerConfig) GetWindowsOSOk() (*bool, bool)`

GetWindowsOSOk returns a tuple with the WindowsOS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsOS

`func (o *AddBaremetalHostRequestServerConfig) SetWindowsOS(v bool)`

SetWindowsOS sets WindowsOS field to given value.

### HasWindowsOS

`func (o *AddBaremetalHostRequestServerConfig) HasWindowsOS() bool`

HasWindowsOS returns a boolean if a field has been set.

### GetOsIpAddress

`func (o *AddBaremetalHostRequestServerConfig) GetOsIpAddress() string`

GetOsIpAddress returns the OsIpAddress field if non-nil, zero value otherwise.

### GetOsIpAddressOk

`func (o *AddBaremetalHostRequestServerConfig) GetOsIpAddressOk() (*string, bool)`

GetOsIpAddressOk returns a tuple with the OsIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsIpAddress

`func (o *AddBaremetalHostRequestServerConfig) SetOsIpAddress(v string)`

SetOsIpAddress sets OsIpAddress field to given value.

### HasOsIpAddress

`func (o *AddBaremetalHostRequestServerConfig) HasOsIpAddress() bool`

HasOsIpAddress returns a boolean if a field has been set.

### GetOsUsername

`func (o *AddBaremetalHostRequestServerConfig) GetOsUsername() string`

GetOsUsername returns the OsUsername field if non-nil, zero value otherwise.

### GetOsUsernameOk

`func (o *AddBaremetalHostRequestServerConfig) GetOsUsernameOk() (*string, bool)`

GetOsUsernameOk returns a tuple with the OsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUsername

`func (o *AddBaremetalHostRequestServerConfig) SetOsUsername(v string)`

SetOsUsername sets OsUsername field to given value.

### HasOsUsername

`func (o *AddBaremetalHostRequestServerConfig) HasOsUsername() bool`

HasOsUsername returns a boolean if a field has been set.

### GetOsPassword

`func (o *AddBaremetalHostRequestServerConfig) GetOsPassword() string`

GetOsPassword returns the OsPassword field if non-nil, zero value otherwise.

### GetOsPasswordOk

`func (o *AddBaremetalHostRequestServerConfig) GetOsPasswordOk() (*string, bool)`

GetOsPasswordOk returns a tuple with the OsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPassword

`func (o *AddBaremetalHostRequestServerConfig) SetOsPassword(v string)`

SetOsPassword sets OsPassword field to given value.

### HasOsPassword

`func (o *AddBaremetalHostRequestServerConfig) HasOsPassword() bool`

HasOsPassword returns a boolean if a field has been set.

### GetOsSSHKeyId

`func (o *AddBaremetalHostRequestServerConfig) GetOsSSHKeyId() int32`

GetOsSSHKeyId returns the OsSSHKeyId field if non-nil, zero value otherwise.

### GetOsSSHKeyIdOk

`func (o *AddBaremetalHostRequestServerConfig) GetOsSSHKeyIdOk() (*int32, bool)`

GetOsSSHKeyIdOk returns a tuple with the OsSSHKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSSHKeyId

`func (o *AddBaremetalHostRequestServerConfig) SetOsSSHKeyId(v int32)`

SetOsSSHKeyId sets OsSSHKeyId field to given value.

### HasOsSSHKeyId

`func (o *AddBaremetalHostRequestServerConfig) HasOsSSHKeyId() bool`

HasOsSSHKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


