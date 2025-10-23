# ListReports200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportResults** | Pointer to [**[]ListReports200ResponseAllOfReportResultsInner**](ListReports200ResponseAllOfReportResultsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListReports200Response

`func NewListReports200Response() *ListReports200Response`

NewListReports200Response instantiates a new ListReports200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReports200ResponseWithDefaults

`func NewListReports200ResponseWithDefaults() *ListReports200Response`

NewListReports200ResponseWithDefaults instantiates a new ListReports200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportResults

`func (o *ListReports200Response) GetReportResults() []ListReports200ResponseAllOfReportResultsInner`

GetReportResults returns the ReportResults field if non-nil, zero value otherwise.

### GetReportResultsOk

`func (o *ListReports200Response) GetReportResultsOk() (*[]ListReports200ResponseAllOfReportResultsInner, bool)`

GetReportResultsOk returns a tuple with the ReportResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportResults

`func (o *ListReports200Response) SetReportResults(v []ListReports200ResponseAllOfReportResultsInner)`

SetReportResults sets ReportResults field to given value.

### HasReportResults

`func (o *ListReports200Response) HasReportResults() bool`

HasReportResults returns a boolean if a field has been set.

### GetMeta

`func (o *ListReports200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListReports200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListReports200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListReports200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


