# CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass &#x60;true&#x60; to allow access all groups | [optional] [default to true]
**Sites** | Pointer to [**[]CreateResourcePoolGroupRequestResourcePoolGroupResourcePermissionSitesInner**](CreateResourcePoolGroupRequestResourcePoolGroupResourcePermissionSitesInner.md) | Array of groups that are allowed access | [optional] 

## Methods

### NewCreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission

`func NewCreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission() *CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission`

NewCreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission instantiates a new CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermissionWithDefaults

`func NewCreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermissionWithDefaults() *CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission`

NewCreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermissionWithDefaults instantiates a new CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission) GetSites() []CreateResourcePoolGroupRequestResourcePoolGroupResourcePermissionSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission) GetSitesOk() (*[]CreateResourcePoolGroupRequestResourcePoolGroupResourcePermissionSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission) SetSites(v []CreateResourcePoolGroupRequestResourcePoolGroupResourcePermissionSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *CreateResourcePoolGroup200ResponseResourcePoolGroupResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


