# CreateNetworkServerGroupRequestGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**CreateNetworkServerGroupRequestGroupAccount**](CreateNetworkServerGroupRequestGroupAccount.md) |  | [optional] 
**Owner** | Pointer to [**CreateNetworkServerGroupRequestGroupOwner**](CreateNetworkServerGroupRequestGroupOwner.md) |  | [optional] 
**NetworkServer** | Pointer to [**CreateNetworkServerGroupRequestGroupNetworkServer**](CreateNetworkServerGroupRequestGroupNetworkServer.md) |  | [optional] 
**Permissions** | Pointer to [**CreateNetworkServerGroupRequestGroupPermissions**](CreateNetworkServerGroupRequestGroupPermissions.md) |  | [optional] 
**Tags** | Pointer to [**[]CreateNetworkServerGroupRequestGroupTagsInner**](CreateNetworkServerGroupRequestGroupTagsInner.md) |  | [optional] 
**Members** | Pointer to [**[]CreateNetworkServerGroupRequestGroupMembersInner**](CreateNetworkServerGroupRequestGroupMembersInner.md) |  | [optional] 

## Methods

### NewCreateNetworkServerGroupRequestGroup

`func NewCreateNetworkServerGroupRequestGroup(name string, ) *CreateNetworkServerGroupRequestGroup`

NewCreateNetworkServerGroupRequestGroup instantiates a new CreateNetworkServerGroupRequestGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkServerGroupRequestGroupWithDefaults

`func NewCreateNetworkServerGroupRequestGroupWithDefaults() *CreateNetworkServerGroupRequestGroup`

NewCreateNetworkServerGroupRequestGroupWithDefaults instantiates a new CreateNetworkServerGroupRequestGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateNetworkServerGroupRequestGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateNetworkServerGroupRequestGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateNetworkServerGroupRequestGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateNetworkServerGroupRequestGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateNetworkServerGroupRequestGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkServerGroupRequestGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkServerGroupRequestGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateNetworkServerGroupRequestGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateNetworkServerGroupRequestGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateNetworkServerGroupRequestGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateNetworkServerGroupRequestGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *CreateNetworkServerGroupRequestGroup) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *CreateNetworkServerGroupRequestGroup) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *CreateNetworkServerGroupRequestGroup) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *CreateNetworkServerGroupRequestGroup) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateNetworkServerGroupRequestGroup) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateNetworkServerGroupRequestGroup) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateNetworkServerGroupRequestGroup) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateNetworkServerGroupRequestGroup) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateNetworkServerGroupRequestGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateNetworkServerGroupRequestGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateNetworkServerGroupRequestGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateNetworkServerGroupRequestGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *CreateNetworkServerGroupRequestGroup) GetAccount() CreateNetworkServerGroupRequestGroupAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateNetworkServerGroupRequestGroup) GetAccountOk() (*CreateNetworkServerGroupRequestGroupAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateNetworkServerGroupRequestGroup) SetAccount(v CreateNetworkServerGroupRequestGroupAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreateNetworkServerGroupRequestGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *CreateNetworkServerGroupRequestGroup) GetOwner() CreateNetworkServerGroupRequestGroupOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CreateNetworkServerGroupRequestGroup) GetOwnerOk() (*CreateNetworkServerGroupRequestGroupOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CreateNetworkServerGroupRequestGroup) SetOwner(v CreateNetworkServerGroupRequestGroupOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CreateNetworkServerGroupRequestGroup) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *CreateNetworkServerGroupRequestGroup) GetNetworkServer() CreateNetworkServerGroupRequestGroupNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *CreateNetworkServerGroupRequestGroup) GetNetworkServerOk() (*CreateNetworkServerGroupRequestGroupNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *CreateNetworkServerGroupRequestGroup) SetNetworkServer(v CreateNetworkServerGroupRequestGroupNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *CreateNetworkServerGroupRequestGroup) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateNetworkServerGroupRequestGroup) GetPermissions() CreateNetworkServerGroupRequestGroupPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateNetworkServerGroupRequestGroup) GetPermissionsOk() (*CreateNetworkServerGroupRequestGroupPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateNetworkServerGroupRequestGroup) SetPermissions(v CreateNetworkServerGroupRequestGroupPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateNetworkServerGroupRequestGroup) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTags

`func (o *CreateNetworkServerGroupRequestGroup) GetTags() []CreateNetworkServerGroupRequestGroupTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateNetworkServerGroupRequestGroup) GetTagsOk() (*[]CreateNetworkServerGroupRequestGroupTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateNetworkServerGroupRequestGroup) SetTags(v []CreateNetworkServerGroupRequestGroupTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateNetworkServerGroupRequestGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMembers

`func (o *CreateNetworkServerGroupRequestGroup) GetMembers() []CreateNetworkServerGroupRequestGroupMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CreateNetworkServerGroupRequestGroup) GetMembersOk() (*[]CreateNetworkServerGroupRequestGroupMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CreateNetworkServerGroupRequestGroup) SetMembers(v []CreateNetworkServerGroupRequestGroupMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *CreateNetworkServerGroupRequestGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


