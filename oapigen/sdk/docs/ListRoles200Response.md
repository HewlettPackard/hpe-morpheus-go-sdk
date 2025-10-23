# ListRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]ListRoles200ResponseAllOfRolesInner**](ListRoles200ResponseAllOfRolesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListRoles200Response

`func NewListRoles200Response() *ListRoles200Response`

NewListRoles200Response instantiates a new ListRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoles200ResponseWithDefaults

`func NewListRoles200ResponseWithDefaults() *ListRoles200Response`

NewListRoles200ResponseWithDefaults instantiates a new ListRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *ListRoles200Response) GetRoles() []ListRoles200ResponseAllOfRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ListRoles200Response) GetRolesOk() (*[]ListRoles200ResponseAllOfRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ListRoles200Response) SetRoles(v []ListRoles200ResponseAllOfRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ListRoles200Response) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetMeta

`func (o *ListRoles200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListRoles200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListRoles200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListRoles200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


