# UpdateCloudDatastoresRequestDatastoreResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass &#x60;true&#x60; to allow access all groups | [optional] [default to true]
**Sites** | Pointer to [**[]UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner**](UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | Pointer to **bool** | Pass true to allow access all plans | [optional] [default to true]
**Plans** | Pointer to [**[]UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner**](UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner.md) | Array of plans that are allowed access | [optional] 

## Methods

### NewUpdateCloudDatastoresRequestDatastoreResourcePermissions

`func NewUpdateCloudDatastoresRequestDatastoreResourcePermissions() *UpdateCloudDatastoresRequestDatastoreResourcePermissions`

NewUpdateCloudDatastoresRequestDatastoreResourcePermissions instantiates a new UpdateCloudDatastoresRequestDatastoreResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudDatastoresRequestDatastoreResourcePermissionsWithDefaults

`func NewUpdateCloudDatastoresRequestDatastoreResourcePermissionsWithDefaults() *UpdateCloudDatastoresRequestDatastoreResourcePermissions`

NewUpdateCloudDatastoresRequestDatastoreResourcePermissionsWithDefaults instantiates a new UpdateCloudDatastoresRequestDatastoreResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) GetSites() []UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) GetSitesOk() (*[]UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) SetSites(v []UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAllPlans

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) GetPlans() []UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) GetPlansOk() (*[]UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) SetPlans(v []UpdateCloudDatastoresRequestDatastoreResourcePermissionsSitesInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *UpdateCloudDatastoresRequestDatastoreResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


