# UpdateUserGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroup** | Pointer to [**UpdateUserGroup200ResponseAllOfUserGroup**](UpdateUserGroup200ResponseAllOfUserGroup.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateUserGroup200Response

`func NewUpdateUserGroup200Response() *UpdateUserGroup200Response`

NewUpdateUserGroup200Response instantiates a new UpdateUserGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetUserGroup

`func (o *UpdateUserGroup200Response) GetUserGroup() UpdateUserGroup200ResponseAllOfUserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *UpdateUserGroup200Response) GetUserGroupOk() (*UpdateUserGroup200ResponseAllOfUserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *UpdateUserGroup200Response) SetUserGroup(v UpdateUserGroup200ResponseAllOfUserGroup)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *UpdateUserGroup200Response) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateUserGroup200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateUserGroup200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateUserGroup200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateUserGroup200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


