# UpdateBlueprintPermissionsRequestResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Set to true to grant access to all groups | [optional] 
**Sites** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of objects identifying groups with access | [optional] 
**OwnerId** | Pointer to **int64** | User ID, can be used to change blueprint owner. | [optional] 

## Methods

### NewUpdateBlueprintPermissionsRequestResourcePermission

`func NewUpdateBlueprintPermissionsRequestResourcePermission() *UpdateBlueprintPermissionsRequestResourcePermission`

NewUpdateBlueprintPermissionsRequestResourcePermission instantiates a new UpdateBlueprintPermissionsRequestResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBlueprintPermissionsRequestResourcePermissionWithDefaults

`func NewUpdateBlueprintPermissionsRequestResourcePermissionWithDefaults() *UpdateBlueprintPermissionsRequestResourcePermission`

NewUpdateBlueprintPermissionsRequestResourcePermissionWithDefaults instantiates a new UpdateBlueprintPermissionsRequestResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetSites() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetSitesOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) SetSites(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetOwnerId

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UpdateBlueprintPermissionsRequestResourcePermission) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


