# UpdateHostManagedRequestServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshUsername** | Pointer to **string** | SSH username to use when provisioning | [optional] 
**SshPassword** | Pointer to **string** | SSH password to use, if not specified the account public key can be used | [optional] 
**ServerOs** | Pointer to [**UpdateHostInstallAgentRequestServerServerOs**](UpdateHostInstallAgentRequestServerServerOs.md) |  | [optional] 
**Plan** | Pointer to [**UpdateHostManagedRequestServerPlan**](UpdateHostManagedRequestServerPlan.md) |  | [optional] 
**Account** | Pointer to [**UpdateHostManagedRequestServerAccount**](UpdateHostManagedRequestServerAccount.md) |  | [optional] 
**ProvisionSiteId** | Pointer to **int64** | Specific group to assign the server | [optional] 
**Tags** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. | [optional] 
**Config** | Pointer to [**UpdateHostManagedRequestServerConfig**](UpdateHostManagedRequestServerConfig.md) |  | [optional] 

## Methods

### NewUpdateHostManagedRequestServer

`func NewUpdateHostManagedRequestServer() *UpdateHostManagedRequestServer`

NewUpdateHostManagedRequestServer instantiates a new UpdateHostManagedRequestServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostManagedRequestServerWithDefaults

`func NewUpdateHostManagedRequestServerWithDefaults() *UpdateHostManagedRequestServer`

NewUpdateHostManagedRequestServerWithDefaults instantiates a new UpdateHostManagedRequestServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshUsername

`func (o *UpdateHostManagedRequestServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *UpdateHostManagedRequestServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *UpdateHostManagedRequestServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *UpdateHostManagedRequestServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *UpdateHostManagedRequestServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *UpdateHostManagedRequestServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *UpdateHostManagedRequestServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *UpdateHostManagedRequestServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetServerOs

`func (o *UpdateHostManagedRequestServer) GetServerOs() UpdateHostInstallAgentRequestServerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *UpdateHostManagedRequestServer) GetServerOsOk() (*UpdateHostInstallAgentRequestServerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *UpdateHostManagedRequestServer) SetServerOs(v UpdateHostInstallAgentRequestServerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *UpdateHostManagedRequestServer) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetPlan

`func (o *UpdateHostManagedRequestServer) GetPlan() UpdateHostManagedRequestServerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UpdateHostManagedRequestServer) GetPlanOk() (*UpdateHostManagedRequestServerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UpdateHostManagedRequestServer) SetPlan(v UpdateHostManagedRequestServerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UpdateHostManagedRequestServer) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateHostManagedRequestServer) GetAccount() UpdateHostManagedRequestServerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateHostManagedRequestServer) GetAccountOk() (*UpdateHostManagedRequestServerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateHostManagedRequestServer) SetAccount(v UpdateHostManagedRequestServerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateHostManagedRequestServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetProvisionSiteId

`func (o *UpdateHostManagedRequestServer) GetProvisionSiteId() int64`

GetProvisionSiteId returns the ProvisionSiteId field if non-nil, zero value otherwise.

### GetProvisionSiteIdOk

`func (o *UpdateHostManagedRequestServer) GetProvisionSiteIdOk() (*int64, bool)`

GetProvisionSiteIdOk returns a tuple with the ProvisionSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionSiteId

`func (o *UpdateHostManagedRequestServer) SetProvisionSiteId(v int64)`

SetProvisionSiteId sets ProvisionSiteId field to given value.

### HasProvisionSiteId

`func (o *UpdateHostManagedRequestServer) HasProvisionSiteId() bool`

HasProvisionSiteId returns a boolean if a field has been set.

### GetTags

`func (o *UpdateHostManagedRequestServer) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateHostManagedRequestServer) GetTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateHostManagedRequestServer) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateHostManagedRequestServer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateHostManagedRequestServer) GetConfig() UpdateHostManagedRequestServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateHostManagedRequestServer) GetConfigOk() (*UpdateHostManagedRequestServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateHostManagedRequestServer) SetConfig(v UpdateHostManagedRequestServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateHostManagedRequestServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


