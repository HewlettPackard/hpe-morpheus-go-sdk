# UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass &#x60;true&#x60; to allow access all groups | [optional] [default to true]
**Sites** | Pointer to [**[]UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermissionSitesInner**](UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermissionSitesInner.md) | Array of groups that are allowed access | [optional] 

## Methods

### NewUpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission

`func NewUpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission() *UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission`

NewUpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission instantiates a new UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetAll

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission) GetSites() []UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermissionSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission) GetSitesOk() (*[]UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermissionSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission) SetSites(v []UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermissionSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateResourcePoolGroupRequestResourcePoolGroupResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


