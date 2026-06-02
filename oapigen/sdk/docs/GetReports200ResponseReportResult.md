# GetReports200ResponseReportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**GetReports200ResponseReportResultType**](GetReports200ResponseReportResultType.md) |  | [optional] 
**ReportTitle** | Pointer to **NullableString** |  | [optional] 
**FilterTitle** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**Config** | Pointer to [**GetReports200ResponseReportResultConfig**](GetReports200ResponseReportResultConfig.md) |  | [optional] 
**CreatedBy** | Pointer to [**GetReports200ResponseReportResultCreatedBy**](GetReports200ResponseReportResultCreatedBy.md) |  | [optional] 
**Rows** | Pointer to [**[]GetReports200ResponseReportResultRowsInner**](GetReports200ResponseReportResultRowsInner.md) |  | [optional] 

## Methods

### NewGetReports200ResponseReportResult

`func NewGetReports200ResponseReportResult() *GetReports200ResponseReportResult`

NewGetReports200ResponseReportResult instantiates a new GetReports200ResponseReportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetReports200ResponseReportResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReports200ResponseReportResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReports200ResponseReportResult) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetReports200ResponseReportResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetReports200ResponseReportResult) GetType() GetReports200ResponseReportResultType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetReports200ResponseReportResult) GetTypeOk() (*GetReports200ResponseReportResultType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetReports200ResponseReportResult) SetType(v GetReports200ResponseReportResultType)`

SetType sets Type field to given value.

### HasType

`func (o *GetReports200ResponseReportResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReportTitle

`func (o *GetReports200ResponseReportResult) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *GetReports200ResponseReportResult) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *GetReports200ResponseReportResult) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.

### HasReportTitle

`func (o *GetReports200ResponseReportResult) HasReportTitle() bool`

HasReportTitle returns a boolean if a field has been set.

### SetReportTitleNil

`func (o *GetReports200ResponseReportResult) SetReportTitleNil(b bool)`

 SetReportTitleNil sets the value for ReportTitle to be an explicit nil

### UnsetReportTitle
`func (o *GetReports200ResponseReportResult) UnsetReportTitle()`

UnsetReportTitle ensures that no value is present for ReportTitle, not even an explicit nil
### GetFilterTitle

`func (o *GetReports200ResponseReportResult) GetFilterTitle() string`

GetFilterTitle returns the FilterTitle field if non-nil, zero value otherwise.

### GetFilterTitleOk

`func (o *GetReports200ResponseReportResult) GetFilterTitleOk() (*string, bool)`

GetFilterTitleOk returns a tuple with the FilterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTitle

`func (o *GetReports200ResponseReportResult) SetFilterTitle(v string)`

SetFilterTitle sets FilterTitle field to given value.

### HasFilterTitle

`func (o *GetReports200ResponseReportResult) HasFilterTitle() bool`

HasFilterTitle returns a boolean if a field has been set.

### SetFilterTitleNil

`func (o *GetReports200ResponseReportResult) SetFilterTitleNil(b bool)`

 SetFilterTitleNil sets the value for FilterTitle to be an explicit nil

### UnsetFilterTitle
`func (o *GetReports200ResponseReportResult) UnsetFilterTitle()`

UnsetFilterTitle ensures that no value is present for FilterTitle, not even an explicit nil
### GetStatus

`func (o *GetReports200ResponseReportResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetReports200ResponseReportResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetReports200ResponseReportResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetReports200ResponseReportResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetReports200ResponseReportResult) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetReports200ResponseReportResult) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetReports200ResponseReportResult) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetReports200ResponseReportResult) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetReports200ResponseReportResult) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetReports200ResponseReportResult) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetReports200ResponseReportResult) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetReports200ResponseReportResult) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStartDate

`func (o *GetReports200ResponseReportResult) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetReports200ResponseReportResult) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetReports200ResponseReportResult) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetReports200ResponseReportResult) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *GetReports200ResponseReportResult) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *GetReports200ResponseReportResult) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *GetReports200ResponseReportResult) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetReports200ResponseReportResult) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetReports200ResponseReportResult) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetReports200ResponseReportResult) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *GetReports200ResponseReportResult) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *GetReports200ResponseReportResult) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetConfig

`func (o *GetReports200ResponseReportResult) GetConfig() GetReports200ResponseReportResultConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetReports200ResponseReportResult) GetConfigOk() (*GetReports200ResponseReportResultConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetReports200ResponseReportResult) SetConfig(v GetReports200ResponseReportResultConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetReports200ResponseReportResult) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetReports200ResponseReportResult) GetCreatedBy() GetReports200ResponseReportResultCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetReports200ResponseReportResult) GetCreatedByOk() (*GetReports200ResponseReportResultCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetReports200ResponseReportResult) SetCreatedBy(v GetReports200ResponseReportResultCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetReports200ResponseReportResult) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetRows

`func (o *GetReports200ResponseReportResult) GetRows() []GetReports200ResponseReportResultRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetReports200ResponseReportResult) GetRowsOk() (*[]GetReports200ResponseReportResultRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetReports200ResponseReportResult) SetRows(v []GetReports200ResponseReportResultRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetReports200ResponseReportResult) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


