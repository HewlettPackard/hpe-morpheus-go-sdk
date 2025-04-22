# ListBackupJobs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | Pointer to [**[]ListBackupJobs200ResponseAllOfJobsInner**](ListBackupJobs200ResponseAllOfJobsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListBackupJobs200Response

`func NewListBackupJobs200Response() *ListBackupJobs200Response`

NewListBackupJobs200Response instantiates a new ListBackupJobs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBackupJobs200ResponseWithDefaults

`func NewListBackupJobs200ResponseWithDefaults() *ListBackupJobs200Response`

NewListBackupJobs200ResponseWithDefaults instantiates a new ListBackupJobs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *ListBackupJobs200Response) GetJobs() []ListBackupJobs200ResponseAllOfJobsInner`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *ListBackupJobs200Response) GetJobsOk() (*[]ListBackupJobs200ResponseAllOfJobsInner, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *ListBackupJobs200Response) SetJobs(v []ListBackupJobs200ResponseAllOfJobsInner)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *ListBackupJobs200Response) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetMeta

`func (o *ListBackupJobs200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBackupJobs200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBackupJobs200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBackupJobs200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


