# GetNetworkServerGroup200ResponseGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**GetNetworkServerGroup200ResponseGroupAccount**](GetNetworkServerGroup200ResponseGroupAccount.md) |  | [optional] 
**Owner** | Pointer to [**GetNetworkServerGroup200ResponseGroupOwner**](GetNetworkServerGroup200ResponseGroupOwner.md) |  | [optional] 
**NetworkServer** | Pointer to [**GetNetworkServerGroup200ResponseGroupNetworkServer**](GetNetworkServerGroup200ResponseGroupNetworkServer.md) |  | [optional] 
**Permissions** | Pointer to [**GetNetworkServerGroup200ResponseGroupPermissions**](GetNetworkServerGroup200ResponseGroupPermissions.md) |  | [optional] 
**Tags** | Pointer to [**[]GetNetworkServerGroup200ResponseGroupTagsInner**](GetNetworkServerGroup200ResponseGroupTagsInner.md) |  | [optional] 
**Members** | Pointer to [**[]GetNetworkServerGroup200ResponseGroupMembersInner**](GetNetworkServerGroup200ResponseGroupMembersInner.md) |  | [optional] 

## Methods

### NewGetNetworkServerGroup200ResponseGroup

`func NewGetNetworkServerGroup200ResponseGroup() *GetNetworkServerGroup200ResponseGroup`

NewGetNetworkServerGroup200ResponseGroup instantiates a new GetNetworkServerGroup200ResponseGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkServerGroup200ResponseGroupWithDefaults

`func NewGetNetworkServerGroup200ResponseGroupWithDefaults() *GetNetworkServerGroup200ResponseGroup`

NewGetNetworkServerGroup200ResponseGroupWithDefaults instantiates a new GetNetworkServerGroup200ResponseGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkServerGroup200ResponseGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkServerGroup200ResponseGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkServerGroup200ResponseGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkServerGroup200ResponseGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkServerGroup200ResponseGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkServerGroup200ResponseGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkServerGroup200ResponseGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkServerGroup200ResponseGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkServerGroup200ResponseGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetNetworkServerGroup200ResponseGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetNetworkServerGroup200ResponseGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *GetNetworkServerGroup200ResponseGroup) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetNetworkServerGroup200ResponseGroup) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetNetworkServerGroup200ResponseGroup) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetNetworkServerGroup200ResponseGroup) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetNetworkServerGroup200ResponseGroup) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetNetworkServerGroup200ResponseGroup) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *GetNetworkServerGroup200ResponseGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetNetworkServerGroup200ResponseGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetNetworkServerGroup200ResponseGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *GetNetworkServerGroup200ResponseGroup) GetAccount() GetNetworkServerGroup200ResponseGroupAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetAccountOk() (*GetNetworkServerGroup200ResponseGroupAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetNetworkServerGroup200ResponseGroup) SetAccount(v GetNetworkServerGroup200ResponseGroupAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetNetworkServerGroup200ResponseGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *GetNetworkServerGroup200ResponseGroup) GetOwner() GetNetworkServerGroup200ResponseGroupOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetOwnerOk() (*GetNetworkServerGroup200ResponseGroupOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetNetworkServerGroup200ResponseGroup) SetOwner(v GetNetworkServerGroup200ResponseGroupOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetNetworkServerGroup200ResponseGroup) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *GetNetworkServerGroup200ResponseGroup) GetNetworkServer() GetNetworkServerGroup200ResponseGroupNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetNetworkServerOk() (*GetNetworkServerGroup200ResponseGroupNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *GetNetworkServerGroup200ResponseGroup) SetNetworkServer(v GetNetworkServerGroup200ResponseGroupNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *GetNetworkServerGroup200ResponseGroup) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetPermissions

`func (o *GetNetworkServerGroup200ResponseGroup) GetPermissions() GetNetworkServerGroup200ResponseGroupPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetPermissionsOk() (*GetNetworkServerGroup200ResponseGroupPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetNetworkServerGroup200ResponseGroup) SetPermissions(v GetNetworkServerGroup200ResponseGroupPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetNetworkServerGroup200ResponseGroup) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTags

`func (o *GetNetworkServerGroup200ResponseGroup) GetTags() []GetNetworkServerGroup200ResponseGroupTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetTagsOk() (*[]GetNetworkServerGroup200ResponseGroupTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetNetworkServerGroup200ResponseGroup) SetTags(v []GetNetworkServerGroup200ResponseGroupTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetNetworkServerGroup200ResponseGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMembers

`func (o *GetNetworkServerGroup200ResponseGroup) GetMembers() []GetNetworkServerGroup200ResponseGroupMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetNetworkServerGroup200ResponseGroup) GetMembersOk() (*[]GetNetworkServerGroup200ResponseGroupMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetNetworkServerGroup200ResponseGroup) SetMembers(v []GetNetworkServerGroup200ResponseGroupMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetNetworkServerGroup200ResponseGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


