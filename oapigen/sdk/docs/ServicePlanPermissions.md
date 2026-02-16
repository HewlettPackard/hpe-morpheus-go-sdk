# ServicePlanPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePermissions** | Pointer to [**ServicePlanPermissionsResourcePermissions**](ServicePlanPermissionsResourcePermissions.md) |  | [optional] 
**TenantPermissions** | Pointer to [**ServicePlanPermissionsTenantPermissions**](ServicePlanPermissionsTenantPermissions.md) |  | [optional] 

## Methods

### NewServicePlanPermissions

`func NewServicePlanPermissions() *ServicePlanPermissions`

NewServicePlanPermissions instantiates a new ServicePlanPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanPermissionsWithDefaults

`func NewServicePlanPermissionsWithDefaults() *ServicePlanPermissions`

NewServicePlanPermissionsWithDefaults instantiates a new ServicePlanPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePermissions

`func (o *ServicePlanPermissions) GetResourcePermissions() ServicePlanPermissionsResourcePermissions`

GetResourcePermissions returns the ResourcePermissions field if non-nil, zero value otherwise.

### GetResourcePermissionsOk

`func (o *ServicePlanPermissions) GetResourcePermissionsOk() (*ServicePlanPermissionsResourcePermissions, bool)`

GetResourcePermissionsOk returns a tuple with the ResourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermissions

`func (o *ServicePlanPermissions) SetResourcePermissions(v ServicePlanPermissionsResourcePermissions)`

SetResourcePermissions sets ResourcePermissions field to given value.

### HasResourcePermissions

`func (o *ServicePlanPermissions) HasResourcePermissions() bool`

HasResourcePermissions returns a boolean if a field has been set.

### GetTenantPermissions

`func (o *ServicePlanPermissions) GetTenantPermissions() ServicePlanPermissionsTenantPermissions`

GetTenantPermissions returns the TenantPermissions field if non-nil, zero value otherwise.

### GetTenantPermissionsOk

`func (o *ServicePlanPermissions) GetTenantPermissionsOk() (*ServicePlanPermissionsTenantPermissions, bool)`

GetTenantPermissionsOk returns a tuple with the TenantPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantPermissions

`func (o *ServicePlanPermissions) SetTenantPermissions(v ServicePlanPermissionsTenantPermissions)`

SetTenantPermissions sets TenantPermissions field to given value.

### HasTenantPermissions

`func (o *ServicePlanPermissions) HasTenantPermissions() bool`

HasTenantPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


