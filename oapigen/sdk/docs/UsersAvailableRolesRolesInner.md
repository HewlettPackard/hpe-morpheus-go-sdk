# UsersAvailableRolesRolesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Authority** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**RoleType** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewUsersAvailableRolesRolesInner

`func NewUsersAvailableRolesRolesInner() *UsersAvailableRolesRolesInner`

NewUsersAvailableRolesRolesInner instantiates a new UsersAvailableRolesRolesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersAvailableRolesRolesInnerWithDefaults

`func NewUsersAvailableRolesRolesInnerWithDefaults() *UsersAvailableRolesRolesInner`

NewUsersAvailableRolesRolesInnerWithDefaults instantiates a new UsersAvailableRolesRolesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsersAvailableRolesRolesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersAvailableRolesRolesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersAvailableRolesRolesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UsersAvailableRolesRolesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthority

`func (o *UsersAvailableRolesRolesInner) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *UsersAvailableRolesRolesInner) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *UsersAvailableRolesRolesInner) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *UsersAvailableRolesRolesInner) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetDescription

`func (o *UsersAvailableRolesRolesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UsersAvailableRolesRolesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UsersAvailableRolesRolesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UsersAvailableRolesRolesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UsersAvailableRolesRolesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UsersAvailableRolesRolesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRoleType

`func (o *UsersAvailableRolesRolesInner) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *UsersAvailableRolesRolesInner) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *UsersAvailableRolesRolesInner) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *UsersAvailableRolesRolesInner) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetOwner

`func (o *UsersAvailableRolesRolesInner) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UsersAvailableRolesRolesInner) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UsersAvailableRolesRolesInner) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UsersAvailableRolesRolesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


