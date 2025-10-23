# ListCheckGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckGroups** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInner**](GetAlerts200ResponseAllOfCheckGroupsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCheckGroups200Response

`func NewListCheckGroups200Response() *ListCheckGroups200Response`

NewListCheckGroups200Response instantiates a new ListCheckGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCheckGroups200ResponseWithDefaults

`func NewListCheckGroups200ResponseWithDefaults() *ListCheckGroups200Response`

NewListCheckGroups200ResponseWithDefaults instantiates a new ListCheckGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckGroups

`func (o *ListCheckGroups200Response) GetCheckGroups() []GetAlerts200ResponseAllOfCheckGroupsInner`

GetCheckGroups returns the CheckGroups field if non-nil, zero value otherwise.

### GetCheckGroupsOk

`func (o *ListCheckGroups200Response) GetCheckGroupsOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInner, bool)`

GetCheckGroupsOk returns a tuple with the CheckGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroups

`func (o *ListCheckGroups200Response) SetCheckGroups(v []GetAlerts200ResponseAllOfCheckGroupsInner)`

SetCheckGroups sets CheckGroups field to given value.

### HasCheckGroups

`func (o *ListCheckGroups200Response) HasCheckGroups() bool`

HasCheckGroups returns a boolean if a field has been set.

### GetMeta

`func (o *ListCheckGroups200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCheckGroups200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCheckGroups200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCheckGroups200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


