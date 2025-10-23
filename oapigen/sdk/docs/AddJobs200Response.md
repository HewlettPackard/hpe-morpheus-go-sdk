# AddJobs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**ListJobs200ResponseAllOfJobsInnerAnyOf**](ListJobs200ResponseAllOfJobsInnerAnyOf.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddJobs200Response

`func NewAddJobs200Response() *AddJobs200Response`

NewAddJobs200Response instantiates a new AddJobs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJobs200ResponseWithDefaults

`func NewAddJobs200ResponseWithDefaults() *AddJobs200Response`

NewAddJobs200ResponseWithDefaults instantiates a new AddJobs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *AddJobs200Response) GetJob() ListJobs200ResponseAllOfJobsInnerAnyOf`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *AddJobs200Response) GetJobOk() (*ListJobs200ResponseAllOfJobsInnerAnyOf, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *AddJobs200Response) SetJob(v ListJobs200ResponseAllOfJobsInnerAnyOf)`

SetJob sets Job field to given value.

### HasJob

`func (o *AddJobs200Response) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetSuccess

`func (o *AddJobs200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddJobs200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddJobs200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddJobs200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


