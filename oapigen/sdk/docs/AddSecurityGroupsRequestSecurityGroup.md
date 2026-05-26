# AddSecurityGroupsRequestSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name for your security group | 
**Description** | Pointer to **string** | Optional description field | [optional] 
**ZoneId** | **int64** | Scoped Cloud ID | 
**Active** | Pointer to **bool** | Set to &#x60;false&#x60; to disable a security group. | [optional] 
**Visibility** | Pointer to **string** | Visibility for the security group. | [optional] [default to "private"]
**NetworkServerId** | Pointer to **int64** | Network Server ID to scope the security group to a network integration (e.g. NSX-T). Use as an alternative to zoneId for network-server-scoped groups. | [optional] 
**CustomOptions** | Pointer to [**AddSecurityGroupsRequestSecurityGroupCustomOptions**](AddSecurityGroupsRequestSecurityGroupCustomOptions.md) |  | [optional] 
**TenantPermissions** | Pointer to [**AddSecurityGroupsRequestSecurityGroupTenantPermissions**](AddSecurityGroupsRequestSecurityGroupTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**AddSecurityGroupsRequestSecurityGroupResourcePermissions**](AddSecurityGroupsRequestSecurityGroupResourcePermissions.md) |  | [optional] 

## Methods

### NewAddSecurityGroupsRequestSecurityGroup

`func NewAddSecurityGroupsRequestSecurityGroup(name string, zoneId int64, ) *AddSecurityGroupsRequestSecurityGroup`

NewAddSecurityGroupsRequestSecurityGroup instantiates a new AddSecurityGroupsRequestSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroupsRequestSecurityGroupWithDefaults

`func NewAddSecurityGroupsRequestSecurityGroupWithDefaults() *AddSecurityGroupsRequestSecurityGroup`

NewAddSecurityGroupsRequestSecurityGroupWithDefaults instantiates a new AddSecurityGroupsRequestSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddSecurityGroupsRequestSecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSecurityGroupsRequestSecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSecurityGroupsRequestSecurityGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddSecurityGroupsRequestSecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSecurityGroupsRequestSecurityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSecurityGroupsRequestSecurityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSecurityGroupsRequestSecurityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetZoneId

`func (o *AddSecurityGroupsRequestSecurityGroup) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddSecurityGroupsRequestSecurityGroup) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddSecurityGroupsRequestSecurityGroup) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.


### GetActive

`func (o *AddSecurityGroupsRequestSecurityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddSecurityGroupsRequestSecurityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddSecurityGroupsRequestSecurityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddSecurityGroupsRequestSecurityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetVisibility

`func (o *AddSecurityGroupsRequestSecurityGroup) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddSecurityGroupsRequestSecurityGroup) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddSecurityGroupsRequestSecurityGroup) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddSecurityGroupsRequestSecurityGroup) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetNetworkServerId

`func (o *AddSecurityGroupsRequestSecurityGroup) GetNetworkServerId() int64`

GetNetworkServerId returns the NetworkServerId field if non-nil, zero value otherwise.

### GetNetworkServerIdOk

`func (o *AddSecurityGroupsRequestSecurityGroup) GetNetworkServerIdOk() (*int64, bool)`

GetNetworkServerIdOk returns a tuple with the NetworkServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServerId

`func (o *AddSecurityGroupsRequestSecurityGroup) SetNetworkServerId(v int64)`

SetNetworkServerId sets NetworkServerId field to given value.

### HasNetworkServerId

`func (o *AddSecurityGroupsRequestSecurityGroup) HasNetworkServerId() bool`

HasNetworkServerId returns a boolean if a field has been set.

### GetCustomOptions

`func (o *AddSecurityGroupsRequestSecurityGroup) GetCustomOptions() AddSecurityGroupsRequestSecurityGroupCustomOptions`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *AddSecurityGroupsRequestSecurityGroup) GetCustomOptionsOk() (*AddSecurityGroupsRequestSecurityGroupCustomOptions, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *AddSecurityGroupsRequestSecurityGroup) SetCustomOptions(v AddSecurityGroupsRequestSecurityGroupCustomOptions)`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *AddSecurityGroupsRequestSecurityGroup) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *AddSecurityGroupsRequestSecurityGroup) GetTenantPermissions() AddSecurityGroupsRequestSecurityGroupTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *AddSecurityGroupsRequestSecurityGroup) GetTenantPermissionsOk() (*AddSecurityGroupsRequestSecurityGroupTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *AddSecurityGroupsRequestSecurityGroup) SetTenantPermissions(v AddSecurityGroupsRequestSecurityGroupTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *AddSecurityGroupsRequestSecurityGroup) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *AddSecurityGroupsRequestSecurityGroup) GetResourcePermissions() AddSecurityGroupsRequestSecurityGroupResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *AddSecurityGroupsRequestSecurityGroup) GetResourcePermissionsOk() (*AddSecurityGroupsRequestSecurityGroupResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *AddSecurityGroupsRequestSecurityGroup) SetResourcePermissions(v AddSecurityGroupsRequestSecurityGroupResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *AddSecurityGroupsRequestSecurityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


