# ListJobs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]ListJobs200ResponseAllOfJobsInner**](ListJobs200ResponseAllOfJobsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListJobs200Response

`func NewListJobs200Response() *ListJobs200Response`

NewListJobs200Response instantiates a new ListJobs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJobs200ResponseWithDefaults

`func NewListJobs200ResponseWithDefaults() *ListJobs200Response`

NewListJobs200ResponseWithDefaults instantiates a new ListJobs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *ListJobs200Response) GetJobs() []ListJobs200ResponseAllOfJobsInner`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ListJobs200Response) GetJobsOk() (*[]ListJobs200ResponseAllOfJobsInner, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ListJobs200Response) SetJobs(v []ListJobs200ResponseAllOfJobsInner)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ListJobs200Response) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetMeta

`func (o *ListJobs200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListJobs200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListJobs200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListJobs200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


