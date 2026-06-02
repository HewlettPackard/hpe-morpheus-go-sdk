# UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass true to allow access to all groups | [optional] 
**Sites** | Pointer to [**[]UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissionsSitesInner**](UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissionsSitesInner.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | Pointer to **bool** | Pass true to allow access to all plans | [optional] 
**Plans** | Pointer to [**[]UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissionsPlansInner**](UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissionsPlansInner.md) | Array of plans that are allowed access | [optional] 

## Methods

### NewUpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions

`func NewUpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions() *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions`

NewUpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions instantiates a new UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetAll

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) GetSites() []UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) GetSitesOk() (*[]UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) SetSites(v []UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAllPlans

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) GetPlans() []UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) GetPlansOk() (*[]UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) SetPlans(v []UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *UpdateClusterNamespaceRequestNamespacePermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


