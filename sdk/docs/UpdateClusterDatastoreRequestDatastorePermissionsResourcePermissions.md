# UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass true to allow access to all groups | [optional] 
**Sites** | Pointer to [**[]UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsSitesInner**](UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsSitesInner.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | Pointer to **bool** | Pass true to allow access to all plans | [optional] 
**Plans** | Pointer to [**[]UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsSitesInner**](UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsSitesInner.md) | Array of plans that are allowed access | [optional] 

## Methods

### NewUpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions

`func NewUpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions() *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions`

NewUpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions instantiates a new UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsWithDefaults

`func NewUpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsWithDefaults() *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions`

NewUpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsWithDefaults instantiates a new UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) GetSites() []UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) GetSitesOk() (*[]UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) SetSites(v []UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAllPlans

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) GetPlans() []UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsSitesInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) GetPlansOk() (*[]UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsSitesInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) SetPlans(v []UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissionsSitesInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *UpdateClusterDatastoreRequestDatastorePermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


