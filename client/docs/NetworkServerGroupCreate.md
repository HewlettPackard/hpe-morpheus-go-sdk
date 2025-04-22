# NetworkServerGroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**NetworkServer** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 
**Permissions** | Pointer to [**ListNetworkServerGroups200ResponseAllOfGroupsInnerPermissions**](ListNetworkServerGroups200ResponseAllOfGroupsInnerPermissions.md) |  | [optional] 
**Tags** | Pointer to [**[]ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner**](ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner.md) |  | [optional] 
**Members** | Pointer to [**[]ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner**](ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner.md) |  | [optional] 

## Methods

### NewNetworkServerGroupCreate

`func NewNetworkServerGroupCreate(name string, ) *NetworkServerGroupCreate`

NewNetworkServerGroupCreate instantiates a new NetworkServerGroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServerGroupCreateWithDefaults

`func NewNetworkServerGroupCreateWithDefaults() *NetworkServerGroupCreate`

NewNetworkServerGroupCreateWithDefaults instantiates a new NetworkServerGroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkServerGroupCreate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkServerGroupCreate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkServerGroupCreate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkServerGroupCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkServerGroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkServerGroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkServerGroupCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NetworkServerGroupCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkServerGroupCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkServerGroupCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkServerGroupCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *NetworkServerGroupCreate) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *NetworkServerGroupCreate) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *NetworkServerGroupCreate) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *NetworkServerGroupCreate) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *NetworkServerGroupCreate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetworkServerGroupCreate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetworkServerGroupCreate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NetworkServerGroupCreate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *NetworkServerGroupCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *NetworkServerGroupCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *NetworkServerGroupCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *NetworkServerGroupCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *NetworkServerGroupCreate) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkServerGroupCreate) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkServerGroupCreate) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NetworkServerGroupCreate) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *NetworkServerGroupCreate) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NetworkServerGroupCreate) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NetworkServerGroupCreate) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NetworkServerGroupCreate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *NetworkServerGroupCreate) GetNetworkServer() GetAlerts200ResponseAllOfChecksInnerAccount`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *NetworkServerGroupCreate) GetNetworkServerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *NetworkServerGroupCreate) SetNetworkServer(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *NetworkServerGroupCreate) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetPermissions

`func (o *NetworkServerGroupCreate) GetPermissions() ListNetworkServerGroups200ResponseAllOfGroupsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *NetworkServerGroupCreate) GetPermissionsOk() (*ListNetworkServerGroups200ResponseAllOfGroupsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *NetworkServerGroupCreate) SetPermissions(v ListNetworkServerGroups200ResponseAllOfGroupsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *NetworkServerGroupCreate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTags

`func (o *NetworkServerGroupCreate) GetTags() []ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetworkServerGroupCreate) GetTagsOk() (*[]ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetworkServerGroupCreate) SetTags(v []ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NetworkServerGroupCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMembers

`func (o *NetworkServerGroupCreate) GetMembers() []ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *NetworkServerGroupCreate) GetMembersOk() (*[]ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *NetworkServerGroupCreate) SetMembers(v []ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *NetworkServerGroupCreate) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


