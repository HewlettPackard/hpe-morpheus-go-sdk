# ListClusterJobs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]ListClusterJobs200ResponseAllOfJobsInner**](ListClusterJobs200ResponseAllOfJobsInner.md) |  | [optional] 
**Stats** | Pointer to **map[string]interface{}** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterJobs200Response

`func NewListClusterJobs200Response() *ListClusterJobs200Response`

NewListClusterJobs200Response instantiates a new ListClusterJobs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterJobs200ResponseWithDefaults

`func NewListClusterJobs200ResponseWithDefaults() *ListClusterJobs200Response`

NewListClusterJobs200ResponseWithDefaults instantiates a new ListClusterJobs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *ListClusterJobs200Response) GetJobs() []ListClusterJobs200ResponseAllOfJobsInner`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ListClusterJobs200Response) GetJobsOk() (*[]ListClusterJobs200ResponseAllOfJobsInner, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ListClusterJobs200Response) SetJobs(v []ListClusterJobs200ResponseAllOfJobsInner)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ListClusterJobs200Response) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetStats

`func (o *ListClusterJobs200Response) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListClusterJobs200Response) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListClusterJobs200Response) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListClusterJobs200Response) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterJobs200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterJobs200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterJobs200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterJobs200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


