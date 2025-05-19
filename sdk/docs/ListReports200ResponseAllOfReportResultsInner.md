# ListReports200ResponseAllOfReportResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner**](ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner.md) |  | [optional] 
**ReportTitle** | Pointer to **string** |  | [optional] 
**FilterTitle** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Config** | Pointer to [**ListReports200ResponseAllOfReportResultsInnerConfig**](ListReports200ResponseAllOfReportResultsInnerConfig.md) |  | [optional] 
**CreatedBy** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**Rows** | Pointer to [**[]ListReports200ResponseAllOfReportResultsInnerRowsInner**](ListReports200ResponseAllOfReportResultsInnerRowsInner.md) |  | [optional] 

## Methods

### NewListReports200ResponseAllOfReportResultsInner

`func NewListReports200ResponseAllOfReportResultsInner() *ListReports200ResponseAllOfReportResultsInner`

NewListReports200ResponseAllOfReportResultsInner instantiates a new ListReports200ResponseAllOfReportResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReports200ResponseAllOfReportResultsInnerWithDefaults

`func NewListReports200ResponseAllOfReportResultsInnerWithDefaults() *ListReports200ResponseAllOfReportResultsInner`

NewListReports200ResponseAllOfReportResultsInnerWithDefaults instantiates a new ListReports200ResponseAllOfReportResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListReports200ResponseAllOfReportResultsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListReports200ResponseAllOfReportResultsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListReports200ResponseAllOfReportResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ListReports200ResponseAllOfReportResultsInner) GetType() ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetTypeOk() (*ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListReports200ResponseAllOfReportResultsInner) SetType(v ListClusterNetworkEndpoints200ResponseAllOfEndpointsInner)`

SetType sets Type field to given value.

### HasType

`func (o *ListReports200ResponseAllOfReportResultsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReportTitle

`func (o *ListReports200ResponseAllOfReportResultsInner) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *ListReports200ResponseAllOfReportResultsInner) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.

### HasReportTitle

`func (o *ListReports200ResponseAllOfReportResultsInner) HasReportTitle() bool`

HasReportTitle returns a boolean if a field has been set.

### GetFilterTitle

`func (o *ListReports200ResponseAllOfReportResultsInner) GetFilterTitle() string`

GetFilterTitle returns the FilterTitle field if non-nil, zero value otherwise.

### GetFilterTitleOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetFilterTitleOk() (*string, bool)`

GetFilterTitleOk returns a tuple with the FilterTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTitle

`func (o *ListReports200ResponseAllOfReportResultsInner) SetFilterTitle(v string)`

SetFilterTitle sets FilterTitle field to given value.

### HasFilterTitle

`func (o *ListReports200ResponseAllOfReportResultsInner) HasFilterTitle() bool`

HasFilterTitle returns a boolean if a field has been set.

### GetStatus

`func (o *ListReports200ResponseAllOfReportResultsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListReports200ResponseAllOfReportResultsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListReports200ResponseAllOfReportResultsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListReports200ResponseAllOfReportResultsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListReports200ResponseAllOfReportResultsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListReports200ResponseAllOfReportResultsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListReports200ResponseAllOfReportResultsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListReports200ResponseAllOfReportResultsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListReports200ResponseAllOfReportResultsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStartDate

`func (o *ListReports200ResponseAllOfReportResultsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListReports200ResponseAllOfReportResultsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListReports200ResponseAllOfReportResultsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListReports200ResponseAllOfReportResultsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListReports200ResponseAllOfReportResultsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListReports200ResponseAllOfReportResultsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetConfig

`func (o *ListReports200ResponseAllOfReportResultsInner) GetConfig() ListReports200ResponseAllOfReportResultsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetConfigOk() (*ListReports200ResponseAllOfReportResultsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListReports200ResponseAllOfReportResultsInner) SetConfig(v ListReports200ResponseAllOfReportResultsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListReports200ResponseAllOfReportResultsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListReports200ResponseAllOfReportResultsInner) GetCreatedBy() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetCreatedByOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListReports200ResponseAllOfReportResultsInner) SetCreatedBy(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListReports200ResponseAllOfReportResultsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetRows

`func (o *ListReports200ResponseAllOfReportResultsInner) GetRows() []ListReports200ResponseAllOfReportResultsInnerRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ListReports200ResponseAllOfReportResultsInner) GetRowsOk() (*[]ListReports200ResponseAllOfReportResultsInnerRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ListReports200ResponseAllOfReportResultsInner) SetRows(v []ListReports200ResponseAllOfReportResultsInnerRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ListReports200ResponseAllOfReportResultsInner) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


