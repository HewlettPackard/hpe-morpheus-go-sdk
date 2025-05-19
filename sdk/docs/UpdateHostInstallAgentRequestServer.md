# UpdateHostInstallAgentRequestServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshUsername** | Pointer to **string** | SSH username to use when provisioning | [optional] 
**SshPassword** | Pointer to **string** | SSH password to use, if not specified the account public key can be used | [optional] 
**ServerOs** | Pointer to [**UpdateHostInstallAgentRequestServerServerOs**](UpdateHostInstallAgentRequestServerServerOs.md) |  | [optional] 

## Methods

### NewUpdateHostInstallAgentRequestServer

`func NewUpdateHostInstallAgentRequestServer() *UpdateHostInstallAgentRequestServer`

NewUpdateHostInstallAgentRequestServer instantiates a new UpdateHostInstallAgentRequestServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostInstallAgentRequestServerWithDefaults

`func NewUpdateHostInstallAgentRequestServerWithDefaults() *UpdateHostInstallAgentRequestServer`

NewUpdateHostInstallAgentRequestServerWithDefaults instantiates a new UpdateHostInstallAgentRequestServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshUsername

`func (o *UpdateHostInstallAgentRequestServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *UpdateHostInstallAgentRequestServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *UpdateHostInstallAgentRequestServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *UpdateHostInstallAgentRequestServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *UpdateHostInstallAgentRequestServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *UpdateHostInstallAgentRequestServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *UpdateHostInstallAgentRequestServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *UpdateHostInstallAgentRequestServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetServerOs

`func (o *UpdateHostInstallAgentRequestServer) GetServerOs() UpdateHostInstallAgentRequestServerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *UpdateHostInstallAgentRequestServer) GetServerOsOk() (*UpdateHostInstallAgentRequestServerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *UpdateHostInstallAgentRequestServer) SetServerOs(v UpdateHostInstallAgentRequestServerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *UpdateHostInstallAgentRequestServer) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


