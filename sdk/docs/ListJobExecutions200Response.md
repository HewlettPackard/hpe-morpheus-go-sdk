# ListJobExecutions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobExecutions** | Pointer to [**[]ListJobExecutions200ResponseAllOfJobExecutionsInner**](ListJobExecutions200ResponseAllOfJobExecutionsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListJobExecutions200Response

`func NewListJobExecutions200Response() *ListJobExecutions200Response`

NewListJobExecutions200Response instantiates a new ListJobExecutions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJobExecutions200ResponseWithDefaults

`func NewListJobExecutions200ResponseWithDefaults() *ListJobExecutions200Response`

NewListJobExecutions200ResponseWithDefaults instantiates a new ListJobExecutions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobExecutions

`func (o *ListJobExecutions200Response) GetJobExecutions() []ListJobExecutions200ResponseAllOfJobExecutionsInner`

GetJobExecutions returns the JobExecutions field if non-nil, zero value otherwise.

### GetJobExecutionsOk

`func (o *ListJobExecutions200Response) GetJobExecutionsOk() (*[]ListJobExecutions200ResponseAllOfJobExecutionsInner, bool)`

GetJobExecutionsOk returns a tuple with the JobExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExecutions

`func (o *ListJobExecutions200Response) SetJobExecutions(v []ListJobExecutions200ResponseAllOfJobExecutionsInner)`

SetJobExecutions sets JobExecutions field to given value.

### HasJobExecutions

`func (o *ListJobExecutions200Response) HasJobExecutions() bool`

HasJobExecutions returns a boolean if a field has been set.

### GetMeta

`func (o *ListJobExecutions200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListJobExecutions200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListJobExecutions200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListJobExecutions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


