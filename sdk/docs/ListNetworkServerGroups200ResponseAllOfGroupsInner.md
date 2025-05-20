# ListNetworkServerGroups200ResponseAllOfGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
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

### NewListNetworkServerGroups200ResponseAllOfGroupsInner

`func NewListNetworkServerGroups200ResponseAllOfGroupsInner() *ListNetworkServerGroups200ResponseAllOfGroupsInner`

NewListNetworkServerGroups200ResponseAllOfGroupsInner instantiates a new ListNetworkServerGroups200ResponseAllOfGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkServerGroups200ResponseAllOfGroupsInnerWithDefaults

`func NewListNetworkServerGroups200ResponseAllOfGroupsInnerWithDefaults() *ListNetworkServerGroups200ResponseAllOfGroupsInner`

NewListNetworkServerGroups200ResponseAllOfGroupsInnerWithDefaults instantiates a new ListNetworkServerGroups200ResponseAllOfGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetVisibility

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccount

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetAccount() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetAccountOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetAccount(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNetworkServer

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetNetworkServer() GetAlerts200ResponseAllOfChecksInnerAccount`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetNetworkServerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetNetworkServer(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### GetPermissions

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetPermissions() ListNetworkServerGroups200ResponseAllOfGroupsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetPermissionsOk() (*ListNetworkServerGroups200ResponseAllOfGroupsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetPermissions(v ListNetworkServerGroups200ResponseAllOfGroupsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetTags

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetTags() []ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetTagsOk() (*[]ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetTags(v []ListNetworkServerGroups200ResponseAllOfGroupsInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMembers

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetMembers() []ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) GetMembersOk() (*[]ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) SetMembers(v []ListNetworkServerGroups200ResponseAllOfGroupsInnerMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ListNetworkServerGroups200ResponseAllOfGroupsInner) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


