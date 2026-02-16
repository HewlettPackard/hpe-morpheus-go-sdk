# NetworkCreateResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Pass true to allow access all groups | [optional] 
**Sites** | Pointer to **[]int64** | Array of groups that are allowed access | [optional] 

## Methods

### NewNetworkCreateResourcePermission

`func NewNetworkCreateResourcePermission() *NetworkCreateResourcePermission`

NewNetworkCreateResourcePermission instantiates a new NetworkCreateResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkCreateResourcePermissionWithDefaults

`func NewNetworkCreateResourcePermissionWithDefaults() *NetworkCreateResourcePermission`

NewNetworkCreateResourcePermissionWithDefaults instantiates a new NetworkCreateResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *NetworkCreateResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *NetworkCreateResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *NetworkCreateResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *NetworkCreateResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *NetworkCreateResourcePermission) GetSites() []int64`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *NetworkCreateResourcePermission) GetSitesOk() (*[]int64, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *NetworkCreateResourcePermission) SetSites(v []int64)`

SetSites sets Sites field to given value.

### HasSites

`func (o *NetworkCreateResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


