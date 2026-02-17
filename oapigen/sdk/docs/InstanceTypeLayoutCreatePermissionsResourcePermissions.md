# InstanceTypeLayoutCreatePermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Set to true to grant access to all groups | [optional] 
**Sites** | Pointer to [**[]InstanceTypeLayoutCreatePermissionsResourcePermissionsSitesInner**](InstanceTypeLayoutCreatePermissionsResourcePermissionsSitesInner.md) | Array of objects identifying groups with access | [optional] 

## Methods

### NewInstanceTypeLayoutCreatePermissionsResourcePermissions

`func NewInstanceTypeLayoutCreatePermissionsResourcePermissions() *InstanceTypeLayoutCreatePermissionsResourcePermissions`

NewInstanceTypeLayoutCreatePermissionsResourcePermissions instantiates a new InstanceTypeLayoutCreatePermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeLayoutCreatePermissionsResourcePermissionsWithDefaults

`func NewInstanceTypeLayoutCreatePermissionsResourcePermissionsWithDefaults() *InstanceTypeLayoutCreatePermissionsResourcePermissions`

NewInstanceTypeLayoutCreatePermissionsResourcePermissionsWithDefaults instantiates a new InstanceTypeLayoutCreatePermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *InstanceTypeLayoutCreatePermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *InstanceTypeLayoutCreatePermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *InstanceTypeLayoutCreatePermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *InstanceTypeLayoutCreatePermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *InstanceTypeLayoutCreatePermissionsResourcePermissions) GetSites() []InstanceTypeLayoutCreatePermissionsResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *InstanceTypeLayoutCreatePermissionsResourcePermissions) GetSitesOk() (*[]InstanceTypeLayoutCreatePermissionsResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *InstanceTypeLayoutCreatePermissionsResourcePermissions) SetSites(v []InstanceTypeLayoutCreatePermissionsResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *InstanceTypeLayoutCreatePermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


