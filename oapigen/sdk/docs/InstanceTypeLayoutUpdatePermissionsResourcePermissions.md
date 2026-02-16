# InstanceTypeLayoutUpdatePermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Set to true to grant access to all groups | [optional] 
**Sites** | Pointer to [**[]InstanceTypeLayoutUpdatePermissionsResourcePermissionsSitesInner**](InstanceTypeLayoutUpdatePermissionsResourcePermissionsSitesInner.md) | Array of objects identifying groups with access | [optional] 

## Methods

### NewInstanceTypeLayoutUpdatePermissionsResourcePermissions

`func NewInstanceTypeLayoutUpdatePermissionsResourcePermissions() *InstanceTypeLayoutUpdatePermissionsResourcePermissions`

NewInstanceTypeLayoutUpdatePermissionsResourcePermissions instantiates a new InstanceTypeLayoutUpdatePermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeLayoutUpdatePermissionsResourcePermissionsWithDefaults

`func NewInstanceTypeLayoutUpdatePermissionsResourcePermissionsWithDefaults() *InstanceTypeLayoutUpdatePermissionsResourcePermissions`

NewInstanceTypeLayoutUpdatePermissionsResourcePermissionsWithDefaults instantiates a new InstanceTypeLayoutUpdatePermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *InstanceTypeLayoutUpdatePermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *InstanceTypeLayoutUpdatePermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *InstanceTypeLayoutUpdatePermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *InstanceTypeLayoutUpdatePermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *InstanceTypeLayoutUpdatePermissionsResourcePermissions) GetSites() []InstanceTypeLayoutUpdatePermissionsResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *InstanceTypeLayoutUpdatePermissionsResourcePermissions) GetSitesOk() (*[]InstanceTypeLayoutUpdatePermissionsResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *InstanceTypeLayoutUpdatePermissionsResourcePermissions) SetSites(v []InstanceTypeLayoutUpdatePermissionsResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *InstanceTypeLayoutUpdatePermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


