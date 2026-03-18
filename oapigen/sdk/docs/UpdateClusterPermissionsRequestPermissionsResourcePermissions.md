# UpdateClusterPermissionsRequestPermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass true to allow access to all groups | [optional] 
**Sites** | Pointer to [**[]UpdateClusterPermissionsRequestPermissionsResourcePermissionsSitesInner**](UpdateClusterPermissionsRequestPermissionsResourcePermissionsSitesInner.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | Pointer to **bool** | Pass true to allow access to all plans | [optional] 
**Plans** | Pointer to [**[]UpdateClusterPermissionsRequestPermissionsResourcePermissionsPlansInner**](UpdateClusterPermissionsRequestPermissionsResourcePermissionsPlansInner.md) | Array of plans that are allowed access | [optional] 

## Methods

### NewUpdateClusterPermissionsRequestPermissionsResourcePermissions

`func NewUpdateClusterPermissionsRequestPermissionsResourcePermissions() *UpdateClusterPermissionsRequestPermissionsResourcePermissions`

NewUpdateClusterPermissionsRequestPermissionsResourcePermissions instantiates a new UpdateClusterPermissionsRequestPermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterPermissionsRequestPermissionsResourcePermissionsWithDefaults

`func NewUpdateClusterPermissionsRequestPermissionsResourcePermissionsWithDefaults() *UpdateClusterPermissionsRequestPermissionsResourcePermissions`

NewUpdateClusterPermissionsRequestPermissionsResourcePermissionsWithDefaults instantiates a new UpdateClusterPermissionsRequestPermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) GetSites() []UpdateClusterPermissionsRequestPermissionsResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) GetSitesOk() (*[]UpdateClusterPermissionsRequestPermissionsResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) SetSites(v []UpdateClusterPermissionsRequestPermissionsResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAllPlans

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) GetPlans() []UpdateClusterPermissionsRequestPermissionsResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) GetPlansOk() (*[]UpdateClusterPermissionsRequestPermissionsResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) SetPlans(v []UpdateClusterPermissionsRequestPermissionsResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *UpdateClusterPermissionsRequestPermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


