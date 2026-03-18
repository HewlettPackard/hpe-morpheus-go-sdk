# ListClusterAffinityGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffinityGroups** | Pointer to [**[]ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner**](ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListAlerts200ResponseAllOfMeta**](ListAlerts200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterAffinityGroups200Response

`func NewListClusterAffinityGroups200Response() *ListClusterAffinityGroups200Response`

NewListClusterAffinityGroups200Response instantiates a new ListClusterAffinityGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterAffinityGroups200ResponseWithDefaults

`func NewListClusterAffinityGroups200ResponseWithDefaults() *ListClusterAffinityGroups200Response`

NewListClusterAffinityGroups200ResponseWithDefaults instantiates a new ListClusterAffinityGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffinityGroups

`func (o *ListClusterAffinityGroups200Response) GetAffinityGroups() []ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner`

GetAffinityGroups returns the AffinityGroups field if non-nil, zero value otherwise.

### GetAffinityGroupsOk

`func (o *ListClusterAffinityGroups200Response) GetAffinityGroupsOk() (*[]ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner, bool)`

GetAffinityGroupsOk returns a tuple with the AffinityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityGroups

`func (o *ListClusterAffinityGroups200Response) SetAffinityGroups(v []ListClusterAffinityGroups200ResponseAllOfAffinityGroupsInner)`

SetAffinityGroups sets AffinityGroups field to given value.

### HasAffinityGroups

`func (o *ListClusterAffinityGroups200Response) HasAffinityGroups() bool`

HasAffinityGroups returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterAffinityGroups200Response) GetMeta() ListAlerts200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterAffinityGroups200Response) GetMetaOk() (*ListAlerts200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterAffinityGroups200Response) SetMeta(v ListAlerts200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterAffinityGroups200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


