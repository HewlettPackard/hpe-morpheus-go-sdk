# ClusterNamespaceUpdatePermissionsResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass true to allow access to all groups | [optional] 
**Sites** | Pointer to [**[]ClusterNamespaceUpdatePermissionsResourcePermissionsSitesInner**](ClusterNamespaceUpdatePermissionsResourcePermissionsSitesInner.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | Pointer to **bool** | Pass true to allow access to all plans | [optional] 
**Plans** | Pointer to [**[]ClusterNamespaceUpdatePermissionsResourcePermissionsPlansInner**](ClusterNamespaceUpdatePermissionsResourcePermissionsPlansInner.md) | Array of plans that are allowed access | [optional] 

## Methods

### NewClusterNamespaceUpdatePermissionsResourcePermissions

`func NewClusterNamespaceUpdatePermissionsResourcePermissions() *ClusterNamespaceUpdatePermissionsResourcePermissions`

NewClusterNamespaceUpdatePermissionsResourcePermissions instantiates a new ClusterNamespaceUpdatePermissionsResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNamespaceUpdatePermissionsResourcePermissionsWithDefaults

`func NewClusterNamespaceUpdatePermissionsResourcePermissionsWithDefaults() *ClusterNamespaceUpdatePermissionsResourcePermissions`

NewClusterNamespaceUpdatePermissionsResourcePermissionsWithDefaults instantiates a new ClusterNamespaceUpdatePermissionsResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) GetSites() []ClusterNamespaceUpdatePermissionsResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) GetSitesOk() (*[]ClusterNamespaceUpdatePermissionsResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) SetSites(v []ClusterNamespaceUpdatePermissionsResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAllPlans

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) GetPlans() []ClusterNamespaceUpdatePermissionsResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) GetPlansOk() (*[]ClusterNamespaceUpdatePermissionsResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) SetPlans(v []ClusterNamespaceUpdatePermissionsResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ClusterNamespaceUpdatePermissionsResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


