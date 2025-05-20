# AddSecurityGroupsRequestSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name for your security group | 
**Description** | Pointer to **string** | Optional description field | [optional] 
**ZoneId** | **int64** | Scoped Cloud ID | 
**Active** | Pointer to **bool** | Set to &#x60;false&#x60; to disable a security group. | [optional] 
**CustomOptions** | Pointer to [**AddSecurityGroupsRequestSecurityGroupCustomOptions**](AddSecurityGroupsRequestSecurityGroupCustomOptions.md) |  | [optional] 
**TenantPermissions** | Pointer to [**AddSecurityGroupsRequestSecurityGroupTenantPermissions**](AddSecurityGroupsRequestSecurityGroupTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**UpdateCloudDatastoresRequestDatastoreResourcePermissions**](UpdateCloudDatastoresRequestDatastoreResourcePermissions.md) |  | [optional] 

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

`func (o *AddSecurityGroupsRequestSecurityGroup) GetResourcePermissions() UpdateCloudDatastoresRequestDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *AddSecurityGroupsRequestSecurityGroup) GetResourcePermissionsOk() (*UpdateCloudDatastoresRequestDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *AddSecurityGroupsRequestSecurityGroup) SetResourcePermissions(v UpdateCloudDatastoresRequestDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *AddSecurityGroupsRequestSecurityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


