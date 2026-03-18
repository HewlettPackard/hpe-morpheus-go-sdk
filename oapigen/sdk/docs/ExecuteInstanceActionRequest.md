# ExecuteInstanceActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The instance type code for the node being added during scale operation | [optional] 
**SelectedResourcePoolId** | Pointer to **int64** | The ID of the resource pool where the node will be created | [optional] 
**PreProvisioned** | Pointer to **string** | Set to &#39;on&#39; to use a pre-provisioned server | [optional] 
**SelectedServerId** | Pointer to **int64** | The ID of the pre-provisioned server | [optional] 
**SshUsername** | Pointer to **string** | SSH username for connecting to the pre-provisioned server | [optional] 
**SshPassword** | Pointer to **string** | SSH password for connecting to the pre-provisioned server | [optional] 
**SshHost** | Pointer to **string** | The SSH host IP address for the pre-provisioned server | [optional] 
**SshKeyPair** | Pointer to [**ExecuteInstanceActionRequestSshKeyPair**](ExecuteInstanceActionRequestSshKeyPair.md) |  | [optional] 

## Methods

### NewExecuteInstanceActionRequest

`func NewExecuteInstanceActionRequest() *ExecuteInstanceActionRequest`

NewExecuteInstanceActionRequest instantiates a new ExecuteInstanceActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteInstanceActionRequestWithDefaults

`func NewExecuteInstanceActionRequestWithDefaults() *ExecuteInstanceActionRequest`

NewExecuteInstanceActionRequestWithDefaults instantiates a new ExecuteInstanceActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ExecuteInstanceActionRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExecuteInstanceActionRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExecuteInstanceActionRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ExecuteInstanceActionRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetSelectedResourcePoolId

`func (o *ExecuteInstanceActionRequest) GetSelectedResourcePoolId() int64`

GetSelectedResourcePoolId returns the SelectedResourcePoolId field if non-nil, zero value otherwise.

### GetSelectedResourcePoolIdOk

`func (o *ExecuteInstanceActionRequest) GetSelectedResourcePoolIdOk() (*int64, bool)`

GetSelectedResourcePoolIdOk returns a tuple with the SelectedResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedResourcePoolId

`func (o *ExecuteInstanceActionRequest) SetSelectedResourcePoolId(v int64)`

SetSelectedResourcePoolId sets SelectedResourcePoolId field to given value.

### HasSelectedResourcePoolId

`func (o *ExecuteInstanceActionRequest) HasSelectedResourcePoolId() bool`

HasSelectedResourcePoolId returns a boolean if a field has been set.

### GetPreProvisioned

`func (o *ExecuteInstanceActionRequest) GetPreProvisioned() string`

GetPreProvisioned returns the PreProvisioned field if non-nil, zero value otherwise.

### GetPreProvisionedOk

`func (o *ExecuteInstanceActionRequest) GetPreProvisionedOk() (*string, bool)`

GetPreProvisionedOk returns a tuple with the PreProvisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProvisioned

`func (o *ExecuteInstanceActionRequest) SetPreProvisioned(v string)`

SetPreProvisioned sets PreProvisioned field to given value.

### HasPreProvisioned

`func (o *ExecuteInstanceActionRequest) HasPreProvisioned() bool`

HasPreProvisioned returns a boolean if a field has been set.

### GetSelectedServerId

`func (o *ExecuteInstanceActionRequest) GetSelectedServerId() int64`

GetSelectedServerId returns the SelectedServerId field if non-nil, zero value otherwise.

### GetSelectedServerIdOk

`func (o *ExecuteInstanceActionRequest) GetSelectedServerIdOk() (*int64, bool)`

GetSelectedServerIdOk returns a tuple with the SelectedServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedServerId

`func (o *ExecuteInstanceActionRequest) SetSelectedServerId(v int64)`

SetSelectedServerId sets SelectedServerId field to given value.

### HasSelectedServerId

`func (o *ExecuteInstanceActionRequest) HasSelectedServerId() bool`

HasSelectedServerId returns a boolean if a field has been set.

### GetSshUsername

`func (o *ExecuteInstanceActionRequest) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ExecuteInstanceActionRequest) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ExecuteInstanceActionRequest) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ExecuteInstanceActionRequest) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ExecuteInstanceActionRequest) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ExecuteInstanceActionRequest) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ExecuteInstanceActionRequest) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ExecuteInstanceActionRequest) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshHost

`func (o *ExecuteInstanceActionRequest) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *ExecuteInstanceActionRequest) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *ExecuteInstanceActionRequest) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *ExecuteInstanceActionRequest) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### GetSshKeyPair

`func (o *ExecuteInstanceActionRequest) GetSshKeyPair() ExecuteInstanceActionRequestSshKeyPair`

GetSshKeyPair returns the SshKeyPair field if non-nil, zero value otherwise.

### GetSshKeyPairOk

`func (o *ExecuteInstanceActionRequest) GetSshKeyPairOk() (*ExecuteInstanceActionRequestSshKeyPair, bool)`

GetSshKeyPairOk returns a tuple with the SshKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPair

`func (o *ExecuteInstanceActionRequest) SetSshKeyPair(v ExecuteInstanceActionRequestSshKeyPair)`

SetSshKeyPair sets SshKeyPair field to given value.

### HasSshKeyPair

`func (o *ExecuteInstanceActionRequest) HasSshKeyPair() bool`

HasSshKeyPair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


