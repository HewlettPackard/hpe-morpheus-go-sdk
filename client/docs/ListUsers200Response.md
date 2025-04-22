# ListUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]AddUserTenant200ResponseAllOfUser**](AddUserTenant200ResponseAllOfUser.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 
**Global** | Pointer to **bool** | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. | [optional] [default to false]

## Methods

### NewListUsers200Response

`func NewListUsers200Response() *ListUsers200Response`

NewListUsers200Response instantiates a new ListUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsers200ResponseWithDefaults

`func NewListUsers200ResponseWithDefaults() *ListUsers200Response`

NewListUsers200ResponseWithDefaults instantiates a new ListUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ListUsers200Response) GetUsers() []AddUserTenant200ResponseAllOfUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListUsers200Response) GetUsersOk() (*[]AddUserTenant200ResponseAllOfUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListUsers200Response) SetUsers(v []AddUserTenant200ResponseAllOfUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ListUsers200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetMeta

`func (o *ListUsers200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListUsers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListUsers200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListUsers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetGlobal

`func (o *ListUsers200Response) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *ListUsers200Response) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *ListUsers200Response) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *ListUsers200Response) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


