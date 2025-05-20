# UpdateSecurityGroupsRequestSecurityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for your security group | [optional] 
**Description** | Pointer to **string** | Optional description field | [optional] 
**Active** | Pointer to **bool** | Set to &#x60;false&#x60; to disable a security group. | [optional] 
**TenantPermissions** | Pointer to [**AddSecurityGroupsRequestSecurityGroupTenantPermissions**](AddSecurityGroupsRequestSecurityGroupTenantPermissions.md) |  | [optional] 
**ResourcePermissions** | Pointer to [**UpdateCloudDatastoresRequestDatastoreResourcePermissions**](UpdateCloudDatastoresRequestDatastoreResourcePermissions.md) |  | [optional] 

## Methods

### NewUpdateSecurityGroupsRequestSecurityGroup

`func NewUpdateSecurityGroupsRequestSecurityGroup() *UpdateSecurityGroupsRequestSecurityGroup`

NewUpdateSecurityGroupsRequestSecurityGroup instantiates a new UpdateSecurityGroupsRequestSecurityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecurityGroupsRequestSecurityGroupWithDefaults

`func NewUpdateSecurityGroupsRequestSecurityGroupWithDefaults() *UpdateSecurityGroupsRequestSecurityGroup`

NewUpdateSecurityGroupsRequestSecurityGroupWithDefaults instantiates a new UpdateSecurityGroupsRequestSecurityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSecurityGroupsRequestSecurityGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSecurityGroupsRequestSecurityGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSecurityGroupsRequestSecurityGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSecurityGroupsRequestSecurityGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSecurityGroupsRequestSecurityGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSecurityGroupsRequestSecurityGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSecurityGroupsRequestSecurityGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSecurityGroupsRequestSecurityGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetActive

`func (o *UpdateSecurityGroupsRequestSecurityGroup) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateSecurityGroupsRequestSecurityGroup) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateSecurityGroupsRequestSecurityGroup) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateSecurityGroupsRequestSecurityGroup) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *UpdateSecurityGroupsRequestSecurityGroup) GetTenantPermissions() AddSecurityGroupsRequestSecurityGroupTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *UpdateSecurityGroupsRequestSecurityGroup) GetTenantPermissionsOk() (*AddSecurityGroupsRequestSecurityGroupTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *UpdateSecurityGroupsRequestSecurityGroup) SetTenantPermissions(v AddSecurityGroupsRequestSecurityGroupTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *UpdateSecurityGroupsRequestSecurityGroup) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.

### GetResourcePermissions

`func (o *UpdateSecurityGroupsRequestSecurityGroup) GetResourcePermissions() UpdateCloudDatastoresRequestDatastoreResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *UpdateSecurityGroupsRequestSecurityGroup) GetResourcePermissionsOk() (*UpdateCloudDatastoresRequestDatastoreResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *UpdateSecurityGroupsRequestSecurityGroup) SetResourcePermissions(v UpdateCloudDatastoresRequestDatastoreResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *UpdateSecurityGroupsRequestSecurityGroup) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


