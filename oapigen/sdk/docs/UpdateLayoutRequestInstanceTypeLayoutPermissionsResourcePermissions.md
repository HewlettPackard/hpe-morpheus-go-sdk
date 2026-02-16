# UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Set to true to grant access to all groups | [optional] 
**Sites** | Pointer to [**[]UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissionsSitesInner**](UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissionsSitesInner.md) | Array of objects identifying groups with access | [optional] 

## Methods

### NewUpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions

`func NewUpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions() *UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions`

NewUpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions instantiates a new UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissionsWithDefaults

`func NewUpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissionsWithDefaults() *UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions`

NewUpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissionsWithDefaults instantiates a new UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions) GetSites() []UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions) GetSitesOk() (*[]UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions) SetSites(v []UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateLayoutRequestInstanceTypeLayoutPermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


