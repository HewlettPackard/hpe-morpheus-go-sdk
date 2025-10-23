# ListSecurityGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | Pointer to [**[]ListSecurityGroups200ResponseAllOfSecurityGroupsInner**](ListSecurityGroups200ResponseAllOfSecurityGroupsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListSecurityGroups200Response

`func NewListSecurityGroups200Response() *ListSecurityGroups200Response`

NewListSecurityGroups200Response instantiates a new ListSecurityGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityGroups200ResponseWithDefaults

`func NewListSecurityGroups200ResponseWithDefaults() *ListSecurityGroups200Response`

NewListSecurityGroups200ResponseWithDefaults instantiates a new ListSecurityGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroups

`func (o *ListSecurityGroups200Response) GetSecurityGroups() []ListSecurityGroups200ResponseAllOfSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ListSecurityGroups200Response) GetSecurityGroupsOk() (*[]ListSecurityGroups200ResponseAllOfSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ListSecurityGroups200Response) SetSecurityGroups(v []ListSecurityGroups200ResponseAllOfSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ListSecurityGroups200Response) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetMeta

`func (o *ListSecurityGroups200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSecurityGroups200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSecurityGroups200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSecurityGroups200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


