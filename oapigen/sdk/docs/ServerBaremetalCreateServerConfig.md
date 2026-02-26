# ServerBaremetalCreateServerConfig

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

### NewServerBaremetalCreateServerConfig

`func NewServerBaremetalCreateServerConfig() *ServerBaremetalCreateServerConfig`

NewServerBaremetalCreateServerConfig instantiates a new ServerBaremetalCreateServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerBaremetalCreateServerConfigWithDefaults

`func NewServerBaremetalCreateServerConfigWithDefaults() *ServerBaremetalCreateServerConfig`

NewServerBaremetalCreateServerConfigWithDefaults instantiates a new ServerBaremetalCreateServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIloIpAddress

`func (o *ServerBaremetalCreateServerConfig) GetIloIpAddress() string`

GetIloIpAddress returns the IloIpAddress field if non-nil, zero value otherwise.

### GetIloIpAddressOk

`func (o *ServerBaremetalCreateServerConfig) GetIloIpAddressOk() (*string, bool)`

GetIloIpAddressOk returns a tuple with the IloIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIloIpAddress

`func (o *ServerBaremetalCreateServerConfig) SetIloIpAddress(v string)`

SetIloIpAddress sets IloIpAddress field to given value.

### HasIloIpAddress

`func (o *ServerBaremetalCreateServerConfig) HasIloIpAddress() bool`

HasIloIpAddress returns a boolean if a field has been set.

### GetIloUsername

`func (o *ServerBaremetalCreateServerConfig) GetIloUsername() string`

GetIloUsername returns the IloUsername field if non-nil, zero value otherwise.

### GetIloUsernameOk

`func (o *ServerBaremetalCreateServerConfig) GetIloUsernameOk() (*string, bool)`

GetIloUsernameOk returns a tuple with the IloUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIloUsername

`func (o *ServerBaremetalCreateServerConfig) SetIloUsername(v string)`

SetIloUsername sets IloUsername field to given value.

### HasIloUsername

`func (o *ServerBaremetalCreateServerConfig) HasIloUsername() bool`

HasIloUsername returns a boolean if a field has been set.

### GetIloPassword

`func (o *ServerBaremetalCreateServerConfig) GetIloPassword() string`

GetIloPassword returns the IloPassword field if non-nil, zero value otherwise.

### GetIloPasswordOk

`func (o *ServerBaremetalCreateServerConfig) GetIloPasswordOk() (*string, bool)`

GetIloPasswordOk returns a tuple with the IloPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIloPassword

`func (o *ServerBaremetalCreateServerConfig) SetIloPassword(v string)`

SetIloPassword sets IloPassword field to given value.

### HasIloPassword

`func (o *ServerBaremetalCreateServerConfig) HasIloPassword() bool`

HasIloPassword returns a boolean if a field has been set.

### GetMacAddress

`func (o *ServerBaremetalCreateServerConfig) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ServerBaremetalCreateServerConfig) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ServerBaremetalCreateServerConfig) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ServerBaremetalCreateServerConfig) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ServerBaremetalCreateServerConfig) GetResourcePoolId() int32`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ServerBaremetalCreateServerConfig) GetResourcePoolIdOk() (*int32, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ServerBaremetalCreateServerConfig) SetResourcePoolId(v int32)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ServerBaremetalCreateServerConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetPreProvisioned

`func (o *ServerBaremetalCreateServerConfig) GetPreProvisioned() bool`

GetPreProvisioned returns the PreProvisioned field if non-nil, zero value otherwise.

### GetPreProvisionedOk

`func (o *ServerBaremetalCreateServerConfig) GetPreProvisionedOk() (*bool, bool)`

GetPreProvisionedOk returns a tuple with the PreProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProvisioned

`func (o *ServerBaremetalCreateServerConfig) SetPreProvisioned(v bool)`

SetPreProvisioned sets PreProvisioned field to given value.

### HasPreProvisioned

`func (o *ServerBaremetalCreateServerConfig) HasPreProvisioned() bool`

HasPreProvisioned returns a boolean if a field has been set.

### GetWindowsOS

`func (o *ServerBaremetalCreateServerConfig) GetWindowsOS() bool`

GetWindowsOS returns the WindowsOS field if non-nil, zero value otherwise.

### GetWindowsOSOk

`func (o *ServerBaremetalCreateServerConfig) GetWindowsOSOk() (*bool, bool)`

GetWindowsOSOk returns a tuple with the WindowsOS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsOS

`func (o *ServerBaremetalCreateServerConfig) SetWindowsOS(v bool)`

SetWindowsOS sets WindowsOS field to given value.

### HasWindowsOS

`func (o *ServerBaremetalCreateServerConfig) HasWindowsOS() bool`

HasWindowsOS returns a boolean if a field has been set.

### GetOsIpAddress

`func (o *ServerBaremetalCreateServerConfig) GetOsIpAddress() string`

GetOsIpAddress returns the OsIpAddress field if non-nil, zero value otherwise.

### GetOsIpAddressOk

`func (o *ServerBaremetalCreateServerConfig) GetOsIpAddressOk() (*string, bool)`

GetOsIpAddressOk returns a tuple with the OsIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsIpAddress

`func (o *ServerBaremetalCreateServerConfig) SetOsIpAddress(v string)`

SetOsIpAddress sets OsIpAddress field to given value.

### HasOsIpAddress

`func (o *ServerBaremetalCreateServerConfig) HasOsIpAddress() bool`

HasOsIpAddress returns a boolean if a field has been set.

### GetOsUsername

`func (o *ServerBaremetalCreateServerConfig) GetOsUsername() string`

GetOsUsername returns the OsUsername field if non-nil, zero value otherwise.

### GetOsUsernameOk

`func (o *ServerBaremetalCreateServerConfig) GetOsUsernameOk() (*string, bool)`

GetOsUsernameOk returns a tuple with the OsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsUsername

`func (o *ServerBaremetalCreateServerConfig) SetOsUsername(v string)`

SetOsUsername sets OsUsername field to given value.

### HasOsUsername

`func (o *ServerBaremetalCreateServerConfig) HasOsUsername() bool`

HasOsUsername returns a boolean if a field has been set.

### GetOsPassword

`func (o *ServerBaremetalCreateServerConfig) GetOsPassword() string`

GetOsPassword returns the OsPassword field if non-nil, zero value otherwise.

### GetOsPasswordOk

`func (o *ServerBaremetalCreateServerConfig) GetOsPasswordOk() (*string, bool)`

GetOsPasswordOk returns a tuple with the OsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsPassword

`func (o *ServerBaremetalCreateServerConfig) SetOsPassword(v string)`

SetOsPassword sets OsPassword field to given value.

### HasOsPassword

`func (o *ServerBaremetalCreateServerConfig) HasOsPassword() bool`

HasOsPassword returns a boolean if a field has been set.

### GetOsSSHKeyId

`func (o *ServerBaremetalCreateServerConfig) GetOsSSHKeyId() int32`

GetOsSSHKeyId returns the OsSSHKeyId field if non-nil, zero value otherwise.

### GetOsSSHKeyIdOk

`func (o *ServerBaremetalCreateServerConfig) GetOsSSHKeyIdOk() (*int32, bool)`

GetOsSSHKeyIdOk returns a tuple with the OsSSHKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsSSHKeyId

`func (o *ServerBaremetalCreateServerConfig) SetOsSSHKeyId(v int32)`

SetOsSSHKeyId sets OsSSHKeyId field to given value.

### HasOsSSHKeyId

`func (o *ServerBaremetalCreateServerConfig) HasOsSSHKeyId() bool`

HasOsSSHKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


