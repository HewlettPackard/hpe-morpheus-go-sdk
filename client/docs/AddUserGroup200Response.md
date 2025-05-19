# AddUserGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroup** | Pointer to [**ListUserGroups200ResponseAllOfUserGroupsInner**](ListUserGroups200ResponseAllOfUserGroupsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddUserGroup200Response

`func NewAddUserGroup200Response() *AddUserGroup200Response`

NewAddUserGroup200Response instantiates a new AddUserGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserGroup200ResponseWithDefaults

`func NewAddUserGroup200ResponseWithDefaults() *AddUserGroup200Response`

NewAddUserGroup200ResponseWithDefaults instantiates a new AddUserGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroup

`func (o *AddUserGroup200Response) GetUserGroup() ListUserGroups200ResponseAllOfUserGroupsInner`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *AddUserGroup200Response) GetUserGroupOk() (*ListUserGroups200ResponseAllOfUserGroupsInner, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *AddUserGroup200Response) SetUserGroup(v ListUserGroups200ResponseAllOfUserGroupsInner)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *AddUserGroup200Response) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetSuccess

`func (o *AddUserGroup200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddUserGroup200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddUserGroup200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddUserGroup200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


