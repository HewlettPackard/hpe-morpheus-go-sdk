# AddSecurityGroupsRequestSecurityGroupTenantPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]int64** | Array of tenant account ids that are allowed access | [optional] 
**CanManageAccounts** | Pointer to **[]int64** | Array of tenant account ids that can manage | [optional] 

## Methods

### NewAddSecurityGroupsRequestSecurityGroupTenantPermissions

`func NewAddSecurityGroupsRequestSecurityGroupTenantPermissions() *AddSecurityGroupsRequestSecurityGroupTenantPermissions`

NewAddSecurityGroupsRequestSecurityGroupTenantPermissions instantiates a new AddSecurityGroupsRequestSecurityGroupTenantPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroupsRequestSecurityGroupTenantPermissionsWithDefaults

`func NewAddSecurityGroupsRequestSecurityGroupTenantPermissionsWithDefaults() *AddSecurityGroupsRequestSecurityGroupTenantPermissions`

NewAddSecurityGroupsRequestSecurityGroupTenantPermissionsWithDefaults instantiates a new AddSecurityGroupsRequestSecurityGroupTenantPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AddSecurityGroupsRequestSecurityGroupTenantPermissions) GetAccounts() []int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AddSecurityGroupsRequestSecurityGroupTenantPermissions) GetAccountsOk() (*[]int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AddSecurityGroupsRequestSecurityGroupTenantPermissions) SetAccounts(v []int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AddSecurityGroupsRequestSecurityGroupTenantPermissions) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetCanManageAccounts

`func (o *AddSecurityGroupsRequestSecurityGroupTenantPermissions) GetCanManageAccounts() []int64`

GetCanManageAccounts returns the CanManageAccounts field if non-nil, zero value otherwise.

### GetCanManageAccountsOk

`func (o *AddSecurityGroupsRequestSecurityGroupTenantPermissions) GetCanManageAccountsOk() (*[]int64, bool)`

GetCanManageAccountsOk returns a tuple with the CanManageAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageAccounts

`func (o *AddSecurityGroupsRequestSecurityGroupTenantPermissions) SetCanManageAccounts(v []int64)`

SetCanManageAccounts sets CanManageAccounts field to given value.

### HasCanManageAccounts

`func (o *AddSecurityGroupsRequestSecurityGroupTenantPermissions) HasCanManageAccounts() bool`

HasCanManageAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


