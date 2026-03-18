# ClusterNamespaceCreateResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass true to allow access to all groups | [optional] 
**Sites** | Pointer to [**[]ClusterNamespaceCreateResourcePermissionsSitesInner**](ClusterNamespaceCreateResourcePermissionsSitesInner.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | Pointer to **bool** | Pass true to allow access to all plans | [optional] 
**Plans** | Pointer to [**[]ClusterNamespaceCreateResourcePermissionsPlansInner**](ClusterNamespaceCreateResourcePermissionsPlansInner.md) | Array of plans that are allowed access | [optional] 

## Methods

### NewClusterNamespaceCreateResourcePermissions

`func NewClusterNamespaceCreateResourcePermissions() *ClusterNamespaceCreateResourcePermissions`

NewClusterNamespaceCreateResourcePermissions instantiates a new ClusterNamespaceCreateResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNamespaceCreateResourcePermissionsWithDefaults

`func NewClusterNamespaceCreateResourcePermissionsWithDefaults() *ClusterNamespaceCreateResourcePermissions`

NewClusterNamespaceCreateResourcePermissionsWithDefaults instantiates a new ClusterNamespaceCreateResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *ClusterNamespaceCreateResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *ClusterNamespaceCreateResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *ClusterNamespaceCreateResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *ClusterNamespaceCreateResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *ClusterNamespaceCreateResourcePermissions) GetSites() []ClusterNamespaceCreateResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *ClusterNamespaceCreateResourcePermissions) GetSitesOk() (*[]ClusterNamespaceCreateResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *ClusterNamespaceCreateResourcePermissions) SetSites(v []ClusterNamespaceCreateResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *ClusterNamespaceCreateResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAllPlans

`func (o *ClusterNamespaceCreateResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *ClusterNamespaceCreateResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *ClusterNamespaceCreateResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *ClusterNamespaceCreateResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *ClusterNamespaceCreateResourcePermissions) GetPlans() []ClusterNamespaceCreateResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *ClusterNamespaceCreateResourcePermissions) GetPlansOk() (*[]ClusterNamespaceCreateResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *ClusterNamespaceCreateResourcePermissions) SetPlans(v []ClusterNamespaceCreateResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *ClusterNamespaceCreateResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


