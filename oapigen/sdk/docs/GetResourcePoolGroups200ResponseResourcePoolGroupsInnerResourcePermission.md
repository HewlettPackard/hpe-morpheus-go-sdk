# GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass &#x60;true&#x60; to allow access all groups | [optional] [default to true]
**Sites** | Pointer to [**[]GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionSitesInner**](GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionSitesInner.md) | Array of groups that are allowed access | [optional] 

## Methods

### NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission

`func NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission() *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission`

NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission instantiates a new GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionWithDefaults

`func NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionWithDefaults() *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission`

NewGetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionWithDefaults instantiates a new GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission) GetSites() []GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission) GetSitesOk() (*[]GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission) SetSites(v []GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermissionSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


