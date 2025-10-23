# ListTenantSubtenantIdentitySources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserSources** | Pointer to [**[]GetIdentitySources200ResponseUserSource**](GetIdentitySources200ResponseUserSource.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListTenantSubtenantIdentitySources200Response

`func NewListTenantSubtenantIdentitySources200Response() *ListTenantSubtenantIdentitySources200Response`

NewListTenantSubtenantIdentitySources200Response instantiates a new ListTenantSubtenantIdentitySources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTenantSubtenantIdentitySources200ResponseWithDefaults

`func NewListTenantSubtenantIdentitySources200ResponseWithDefaults() *ListTenantSubtenantIdentitySources200Response`

NewListTenantSubtenantIdentitySources200ResponseWithDefaults instantiates a new ListTenantSubtenantIdentitySources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserSources

`func (o *ListTenantSubtenantIdentitySources200Response) GetUserSources() []GetIdentitySources200ResponseUserSource`

GetUserSources returns the UserSources field if non-nil, zero value otherwise.

### GetUserSourcesOk

`func (o *ListTenantSubtenantIdentitySources200Response) GetUserSourcesOk() (*[]GetIdentitySources200ResponseUserSource, bool)`

GetUserSourcesOk returns a tuple with the UserSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSources

`func (o *ListTenantSubtenantIdentitySources200Response) SetUserSources(v []GetIdentitySources200ResponseUserSource)`

SetUserSources sets UserSources field to given value.

### HasUserSources

`func (o *ListTenantSubtenantIdentitySources200Response) HasUserSources() bool`

HasUserSources returns a boolean if a field has been set.

### GetMeta

`func (o *ListTenantSubtenantIdentitySources200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTenantSubtenantIdentitySources200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTenantSubtenantIdentitySources200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTenantSubtenantIdentitySources200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


