# ListNetworkServerGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]ListNetworkServerGroups200ResponseAllOfGroupsInner**](ListNetworkServerGroups200ResponseAllOfGroupsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListNetworkServerGroups200Response

`func NewListNetworkServerGroups200Response() *ListNetworkServerGroups200Response`

NewListNetworkServerGroups200Response instantiates a new ListNetworkServerGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkServerGroups200ResponseWithDefaults

`func NewListNetworkServerGroups200ResponseWithDefaults() *ListNetworkServerGroups200Response`

NewListNetworkServerGroups200ResponseWithDefaults instantiates a new ListNetworkServerGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ListNetworkServerGroups200Response) GetGroups() []ListNetworkServerGroups200ResponseAllOfGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ListNetworkServerGroups200Response) GetGroupsOk() (*[]ListNetworkServerGroups200ResponseAllOfGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ListNetworkServerGroups200Response) SetGroups(v []ListNetworkServerGroups200ResponseAllOfGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ListNetworkServerGroups200Response) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetMeta

`func (o *ListNetworkServerGroups200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNetworkServerGroups200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNetworkServerGroups200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNetworkServerGroups200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


