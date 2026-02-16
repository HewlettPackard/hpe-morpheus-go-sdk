# ResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** |  | [optional] 
**Sites** | Pointer to [**[]ResourcePermissionsSitesInner**](ResourcePermissionsSitesInner.md) |  | [optional] 
**AllPlans** | Pointer to **bool** |  | [optional] 
**Plans** | Pointer to [**[]ResourcePermissionsPlansInner**](ResourcePermissionsPlansInner.md) |  | [optional] 

## Methods

### NewResourcePermissions

`func NewResourcePermissions() *ResourcePermissions`

NewResourcePermissions instantiates a new ResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePermissionsWithDefaults

`func NewResourcePermissionsWithDefaults() *ResourcePermissions`

NewResourcePermissionsWithDefaults instantiates a new ResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *ResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *ResourcePermissions) GetSites() []ResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ResourcePermissions) GetSitesOk() (*[]ResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ResourcePermissions) SetSites(v []ResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *ResourcePermissions) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *ResourcePermissions) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetAllPlans

`func (o *ResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *ResourcePermissions) GetPlans() []ResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ResourcePermissions) GetPlansOk() (*[]ResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ResourcePermissions) SetPlans(v []ResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.

### SetPlansNil

`func (o *ResourcePermissions) SetPlansNil(b bool)`

 SetPlansNil sets the value for Plans to be an explicit nil

### UnsetPlans
`func (o *ResourcePermissions) UnsetPlans()`

UnsetPlans ensures that no value is present for Plans, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


