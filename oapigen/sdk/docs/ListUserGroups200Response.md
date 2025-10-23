# ListUserGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroups** | Pointer to [**[]ListUserGroups200ResponseAllOfUserGroupsInner**](ListUserGroups200ResponseAllOfUserGroupsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListUserGroups200Response

`func NewListUserGroups200Response() *ListUserGroups200Response`

NewListUserGroups200Response instantiates a new ListUserGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserGroups200ResponseWithDefaults

`func NewListUserGroups200ResponseWithDefaults() *ListUserGroups200Response`

NewListUserGroups200ResponseWithDefaults instantiates a new ListUserGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroups

`func (o *ListUserGroups200Response) GetUserGroups() []ListUserGroups200ResponseAllOfUserGroupsInner`

GetUserGroups returns the UserGroups field if non-nil, zero value otherwise.

### GetUserGroupsOk

`func (o *ListUserGroups200Response) GetUserGroupsOk() (*[]ListUserGroups200ResponseAllOfUserGroupsInner, bool)`

GetUserGroupsOk returns a tuple with the UserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroups

`func (o *ListUserGroups200Response) SetUserGroups(v []ListUserGroups200ResponseAllOfUserGroupsInner)`

SetUserGroups sets UserGroups field to given value.

### HasUserGroups

`func (o *ListUserGroups200Response) HasUserGroups() bool`

HasUserGroups returns a boolean if a field has been set.

### GetMeta

`func (o *ListUserGroups200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListUserGroups200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListUserGroups200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListUserGroups200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


