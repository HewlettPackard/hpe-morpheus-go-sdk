# RunReports200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportResult** | Pointer to [**ListReports200ResponseAllOfReportResultsInner**](ListReports200ResponseAllOfReportResultsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewRunReports200Response

`func NewRunReports200Response() *RunReports200Response`

NewRunReports200Response instantiates a new RunReports200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunReports200ResponseWithDefaults

`func NewRunReports200ResponseWithDefaults() *RunReports200Response`

NewRunReports200ResponseWithDefaults instantiates a new RunReports200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportResult

`func (o *RunReports200Response) GetReportResult() ListReports200ResponseAllOfReportResultsInner`

GetReportResult returns the ReportResult field if non-nil, zero value otherwise.

### GetReportResultOk

`func (o *RunReports200Response) GetReportResultOk() (*ListReports200ResponseAllOfReportResultsInner, bool)`

GetReportResultOk returns a tuple with the ReportResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportResult

`func (o *RunReports200Response) SetReportResult(v ListReports200ResponseAllOfReportResultsInner)`

SetReportResult sets ReportResult field to given value.

### HasReportResult

`func (o *RunReports200Response) HasReportResult() bool`

HasReportResult returns a boolean if a field has been set.

### GetSuccess

`func (o *RunReports200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RunReports200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RunReports200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *RunReports200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


