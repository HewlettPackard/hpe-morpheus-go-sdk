# AddRolesRequestRoleZonesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | &#x60;id&#x60; of the cloud (zone) | 
**Access** | **string** | The new access level. | 

## Methods

### NewAddRolesRequestRoleZonesInner

`func NewAddRolesRequestRoleZonesInner(id int32, access string, ) *AddRolesRequestRoleZonesInner`

NewAddRolesRequestRoleZonesInner instantiates a new AddRolesRequestRoleZonesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRolesRequestRoleZonesInnerWithDefaults

`func NewAddRolesRequestRoleZonesInnerWithDefaults() *AddRolesRequestRoleZonesInner`

NewAddRolesRequestRoleZonesInnerWithDefaults instantiates a new AddRolesRequestRoleZonesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddRolesRequestRoleZonesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddRolesRequestRoleZonesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddRolesRequestRoleZonesInner) SetId(v int32)`

SetId sets Id field to given value.


### GetAccess

`func (o *AddRolesRequestRoleZonesInner) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AddRolesRequestRoleZonesInner) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AddRolesRequestRoleZonesInner) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


