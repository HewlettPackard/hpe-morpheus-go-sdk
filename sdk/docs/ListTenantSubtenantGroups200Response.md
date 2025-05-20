# ListTenantSubtenantGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]ListTenantSubtenantGroups200ResponseAllOfGroupsInner**](ListTenantSubtenantGroups200ResponseAllOfGroupsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListTenantSubtenantGroups200Response

`func NewListTenantSubtenantGroups200Response() *ListTenantSubtenantGroups200Response`

NewListTenantSubtenantGroups200Response instantiates a new ListTenantSubtenantGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTenantSubtenantGroups200ResponseWithDefaults

`func NewListTenantSubtenantGroups200ResponseWithDefaults() *ListTenantSubtenantGroups200Response`

NewListTenantSubtenantGroups200ResponseWithDefaults instantiates a new ListTenantSubtenantGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ListTenantSubtenantGroups200Response) GetGroups() []ListTenantSubtenantGroups200ResponseAllOfGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ListTenantSubtenantGroups200Response) GetGroupsOk() (*[]ListTenantSubtenantGroups200ResponseAllOfGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ListTenantSubtenantGroups200Response) SetGroups(v []ListTenantSubtenantGroups200ResponseAllOfGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ListTenantSubtenantGroups200Response) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetMeta

`func (o *ListTenantSubtenantGroups200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTenantSubtenantGroups200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTenantSubtenantGroups200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTenantSubtenantGroups200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


