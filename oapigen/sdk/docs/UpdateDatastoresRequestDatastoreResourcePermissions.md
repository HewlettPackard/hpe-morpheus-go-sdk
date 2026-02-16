# UpdateDatastoresRequestDatastoreResourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass &#x60;true&#x60; to allow access all groups | [optional] [default to true]
**Sites** | Pointer to [**[]UpdateDatastoresRequestDatastoreResourcePermissionsSitesInner**](UpdateDatastoresRequestDatastoreResourcePermissionsSitesInner.md) | Array of groups that are allowed access | [optional] 
**AllPlans** | Pointer to **bool** | Pass true to allow access all plans | [optional] [default to true]
**Plans** | Pointer to [**[]UpdateDatastoresRequestDatastoreResourcePermissionsPlansInner**](UpdateDatastoresRequestDatastoreResourcePermissionsPlansInner.md) | Array of plans that are allowed access | [optional] 

## Methods

### NewUpdateDatastoresRequestDatastoreResourcePermissions

`func NewUpdateDatastoresRequestDatastoreResourcePermissions() *UpdateDatastoresRequestDatastoreResourcePermissions`

NewUpdateDatastoresRequestDatastoreResourcePermissions instantiates a new UpdateDatastoresRequestDatastoreResourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatastoresRequestDatastoreResourcePermissionsWithDefaults

`func NewUpdateDatastoresRequestDatastoreResourcePermissionsWithDefaults() *UpdateDatastoresRequestDatastoreResourcePermissions`

NewUpdateDatastoresRequestDatastoreResourcePermissionsWithDefaults instantiates a new UpdateDatastoresRequestDatastoreResourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) GetSites() []UpdateDatastoresRequestDatastoreResourcePermissionsSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) GetSitesOk() (*[]UpdateDatastoresRequestDatastoreResourcePermissionsSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) SetSites(v []UpdateDatastoresRequestDatastoreResourcePermissionsSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetAllPlans

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) GetAllPlans() bool`

GetAllPlans returns the AllPlans field if non-nil, zero value otherwise.

### GetAllPlansOk

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) GetAllPlansOk() (*bool, bool)`

GetAllPlansOk returns a tuple with the AllPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlans

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) SetAllPlans(v bool)`

SetAllPlans sets AllPlans field to given value.

### HasAllPlans

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) HasAllPlans() bool`

HasAllPlans returns a boolean if a field has been set.

### GetPlans

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) GetPlans() []UpdateDatastoresRequestDatastoreResourcePermissionsPlansInner`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) GetPlansOk() (*[]UpdateDatastoresRequestDatastoreResourcePermissionsPlansInner, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) SetPlans(v []UpdateDatastoresRequestDatastoreResourcePermissionsPlansInner)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *UpdateDatastoresRequestDatastoreResourcePermissions) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


