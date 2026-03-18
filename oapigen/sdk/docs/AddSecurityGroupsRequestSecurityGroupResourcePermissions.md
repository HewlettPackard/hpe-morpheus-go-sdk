# AddSecurityGroupsRequestSecurityGroupResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass &#x60;true&#x60; to allow access all groups | [optional] [default to true]
**Sites** | Pointer to [**[]UpdateCloudFoldersRequestFolderResourcePermissionsSitesInner**](UpdateCloudFoldersRequestFolderResourcePermissionsSitesInner.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | Pointer to **bool** | Pass true to allow access all plans | [optional] [default to true]
**Plans** | Pointer to [**[]UpdateCloudFoldersRequestFolderResourcePermissionsPlansInner**](UpdateCloudFoldersRequestFolderResourcePermissionsPlansInner.md) | Array of plans that are allowed access | [optional] 

## Methods

### NewAddSecurityGroupsRequestSecurityGroupResourcePermissions

`func NewAddSecurityGroupsRequestSecurityGroupResourcePermissions() *AddSecurityGroupsRequestSecurityGroupResourcePermissions`

NewAddSecurityGroupsRequestSecurityGroupResourcePermissions instantiates a new AddSecurityGroupsRequestSecurityGroupResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroupsRequestSecurityGroupResourcePermissionsWithDefaults

`func NewAddSecurityGroupsRequestSecurityGroupResourcePermissionsWithDefaults() *AddSecurityGroupsRequestSecurityGroupResourcePermissions`

NewAddSecurityGroupsRequestSecurityGroupResourcePermissionsWithDefaults instantiates a new AddSecurityGroupsRequestSecurityGroupResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) GetSites() []UpdateCloudFoldersRequestFolderResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) GetSitesOk() (*[]UpdateCloudFoldersRequestFolderResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) SetSites(v []UpdateCloudFoldersRequestFolderResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAllPlans

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) GetPlans() []UpdateCloudFoldersRequestFolderResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) GetPlansOk() (*[]UpdateCloudFoldersRequestFolderResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) SetPlans(v []UpdateCloudFoldersRequestFolderResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *AddSecurityGroupsRequestSecurityGroupResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


