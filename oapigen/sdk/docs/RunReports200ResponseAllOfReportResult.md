# RunReports200ResponseAllOfReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**RunReports200ResponseAllOfReportResultType**](RunReports200ResponseAllOfReportResultType.md) |  | [optional] 
**ReportTitle** | Pointer to **NullableString** |  | [optional] 
**FilterTitle** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Config** | Pointer to [**RunReports200ResponseAllOfReportResultConfig**](RunReports200ResponseAllOfReportResultConfig.md) |  | [optional] 
**CreatedBy** | Pointer to [**RunReports200ResponseAllOfReportResultCreatedBy**](RunReports200ResponseAllOfReportResultCreatedBy.md) |  | [optional] 
**Rows** | Pointer to [**[]RunReports200ResponseAllOfReportResultRowsInner**](RunReports200ResponseAllOfReportResultRowsInner.md) |  | [optional] 

## Methods

### NewRunReports200ResponseAllOfReportResult

`func NewRunReports200ResponseAllOfReportResult() *RunReports200ResponseAllOfReportResult`

NewRunReports200ResponseAllOfReportResult instantiates a new RunReports200ResponseAllOfReportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *RunReports200ResponseAllOfReportResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunReports200ResponseAllOfReportResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunReports200ResponseAllOfReportResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RunReports200ResponseAllOfReportResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RunReports200ResponseAllOfReportResult) GetType() RunReports200ResponseAllOfReportResultType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunReports200ResponseAllOfReportResult) GetTypeOk() (*RunReports200ResponseAllOfReportResultType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunReports200ResponseAllOfReportResult) SetType(v RunReports200ResponseAllOfReportResultType)`

SetType sets Type field to given value.

### HasType

`func (o *RunReports200ResponseAllOfReportResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReportTitle

`func (o *RunReports200ResponseAllOfReportResult) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *RunReports200ResponseAllOfReportResult) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *RunReports200ResponseAllOfReportResult) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.

### HasReportTitle

`func (o *RunReports200ResponseAllOfReportResult) HasReportTitle() bool`

HasReportTitle returns a boolean if a field has been set.

### SetReportTitleNil

`func (o *RunReports200ResponseAllOfReportResult) SetReportTitleNil(b bool)`

 SetReportTitleNil sets the value for ReportTitle to be an explicit nil

### UnsetReportTitle
`func (o *RunReports200ResponseAllOfReportResult) UnsetReportTitle()`

UnsetReportTitle ensures that no value is present for ReportTitle, not even an explicit nil
### GetFilterTitle

`func (o *RunReports200ResponseAllOfReportResult) GetFilterTitle() string`

GetFilterTitle returns the FilterTitle field if non-nil, zero value otherwise.

### GetFilterTitleOk

`func (o *RunReports200ResponseAllOfReportResult) GetFilterTitleOk() (*string, bool)`

GetFilterTitleOk returns a tuple with the FilterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTitle

`func (o *RunReports200ResponseAllOfReportResult) SetFilterTitle(v string)`

SetFilterTitle sets FilterTitle field to given value.

### HasFilterTitle

`func (o *RunReports200ResponseAllOfReportResult) HasFilterTitle() bool`

HasFilterTitle returns a boolean if a field has been set.

### SetFilterTitleNil

`func (o *RunReports200ResponseAllOfReportResult) SetFilterTitleNil(b bool)`

 SetFilterTitleNil sets the value for FilterTitle to be an explicit nil

### UnsetFilterTitle
`func (o *RunReports200ResponseAllOfReportResult) UnsetFilterTitle()`

UnsetFilterTitle ensures that no value is present for FilterTitle, not even an explicit nil
### GetStatus

`func (o *RunReports200ResponseAllOfReportResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunReports200ResponseAllOfReportResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunReports200ResponseAllOfReportResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RunReports200ResponseAllOfReportResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *RunReports200ResponseAllOfReportResult) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *RunReports200ResponseAllOfReportResult) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *RunReports200ResponseAllOfReportResult) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *RunReports200ResponseAllOfReportResult) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RunReports200ResponseAllOfReportResult) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RunReports200ResponseAllOfReportResult) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RunReports200ResponseAllOfReportResult) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RunReports200ResponseAllOfReportResult) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStartDate

`func (o *RunReports200ResponseAllOfReportResult) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RunReports200ResponseAllOfReportResult) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RunReports200ResponseAllOfReportResult) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *RunReports200ResponseAllOfReportResult) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *RunReports200ResponseAllOfReportResult) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *RunReports200ResponseAllOfReportResult) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *RunReports200ResponseAllOfReportResult) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *RunReports200ResponseAllOfReportResult) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *RunReports200ResponseAllOfReportResult) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *RunReports200ResponseAllOfReportResult) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *RunReports200ResponseAllOfReportResult) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *RunReports200ResponseAllOfReportResult) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetConfig

`func (o *RunReports200ResponseAllOfReportResult) GetConfig() RunReports200ResponseAllOfReportResultConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RunReports200ResponseAllOfReportResult) GetConfigOk() (*RunReports200ResponseAllOfReportResultConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RunReports200ResponseAllOfReportResult) SetConfig(v RunReports200ResponseAllOfReportResultConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *RunReports200ResponseAllOfReportResult) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RunReports200ResponseAllOfReportResult) GetCreatedBy() RunReports200ResponseAllOfReportResultCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RunReports200ResponseAllOfReportResult) GetCreatedByOk() (*RunReports200ResponseAllOfReportResultCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RunReports200ResponseAllOfReportResult) SetCreatedBy(v RunReports200ResponseAllOfReportResultCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RunReports200ResponseAllOfReportResult) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetRows

`func (o *RunReports200ResponseAllOfReportResult) GetRows() []RunReports200ResponseAllOfReportResultRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *RunReports200ResponseAllOfReportResult) GetRowsOk() (*[]RunReports200ResponseAllOfReportResultRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *RunReports200ResponseAllOfReportResult) SetRows(v []RunReports200ResponseAllOfReportResultRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *RunReports200ResponseAllOfReportResult) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


