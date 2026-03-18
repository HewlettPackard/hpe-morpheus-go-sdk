# ListCloudAffinityGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffinityGroups** | Pointer to [**[]ListCloudAffinityGroups200ResponseAllOfAffinityGroupsInner**](ListCloudAffinityGroups200ResponseAllOfAffinityGroupsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListAlerts200ResponseAllOfMeta**](ListAlerts200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCloudAffinityGroups200Response

`func NewListCloudAffinityGroups200Response() *ListCloudAffinityGroups200Response`

NewListCloudAffinityGroups200Response instantiates a new ListCloudAffinityGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudAffinityGroups200ResponseWithDefaults

`func NewListCloudAffinityGroups200ResponseWithDefaults() *ListCloudAffinityGroups200Response`

NewListCloudAffinityGroups200ResponseWithDefaults instantiates a new ListCloudAffinityGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffinityGroups

`func (o *ListCloudAffinityGroups200Response) GetAffinityGroups() []ListCloudAffinityGroups200ResponseAllOfAffinityGroupsInner`

GetAffinityGroups returns the AffinityGroups field if non-nil, zero value otherwise.

### GetAffinityGroupsOk

`func (o *ListCloudAffinityGroups200Response) GetAffinityGroupsOk() (*[]ListCloudAffinityGroups200ResponseAllOfAffinityGroupsInner, bool)`

GetAffinityGroupsOk returns a tuple with the AffinityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinityGroups

`func (o *ListCloudAffinityGroups200Response) SetAffinityGroups(v []ListCloudAffinityGroups200ResponseAllOfAffinityGroupsInner)`

SetAffinityGroups sets AffinityGroups field to given value.

### HasAffinityGroups

`func (o *ListCloudAffinityGroups200Response) HasAffinityGroups() bool`

HasAffinityGroups returns a boolean if a field has been set.

### GetMeta

`func (o *ListCloudAffinityGroups200Response) GetMeta() ListAlerts200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCloudAffinityGroups200Response) GetMetaOk() (*ListAlerts200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCloudAffinityGroups200Response) SetMeta(v ListAlerts200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCloudAffinityGroups200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


