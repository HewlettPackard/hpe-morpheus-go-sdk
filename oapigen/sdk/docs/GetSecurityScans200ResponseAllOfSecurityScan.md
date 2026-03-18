# GetSecurityScans200ResponseAllOfSecurityScan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**SecurityPackage** | Pointer to [**GetSecurityScans200ResponseAllOfSecurityScanSecurityPackage**](GetSecurityScans200ResponseAllOfSecurityScanSecurityPackage.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ScanDate** | Pointer to **time.Time** |  | [optional] 
**ScanDuration** | Pointer to **int64** |  | [optional] 
**TestCount** | Pointer to **int64** |  | [optional] 
**RunCount** | Pointer to **int64** |  | [optional] 
**PassCount** | Pointer to **int64** |  | [optional] 
**FailCount** | Pointer to **int64** |  | [optional] 
**OtherCount** | Pointer to **int64** |  | [optional] 
**ScanScore** | Pointer to **float32** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Results** | Pointer to **map[string]interface{}** | Results Summary (only returned when using query parameter results&#x3D;true) | [optional] 

## Methods

### NewGetSecurityScans200ResponseAllOfSecurityScan

`func NewGetSecurityScans200ResponseAllOfSecurityScan() *GetSecurityScans200ResponseAllOfSecurityScan`

NewGetSecurityScans200ResponseAllOfSecurityScan instantiates a new GetSecurityScans200ResponseAllOfSecurityScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecurityScans200ResponseAllOfSecurityScanWithDefaults

`func NewGetSecurityScans200ResponseAllOfSecurityScanWithDefaults() *GetSecurityScans200ResponseAllOfSecurityScan`

NewGetSecurityScans200ResponseAllOfSecurityScanWithDefaults instantiates a new GetSecurityScans200ResponseAllOfSecurityScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSecurityPackage

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetSecurityPackage() GetSecurityScans200ResponseAllOfSecurityScanSecurityPackage`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetSecurityPackageOk() (*GetSecurityScans200ResponseAllOfSecurityScanSecurityPackage, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetSecurityPackage(v GetSecurityScans200ResponseAllOfSecurityScanSecurityPackage)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### GetStatus

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScanDate

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetScanDate() time.Time`

GetScanDate returns the ScanDate field if non-nil, zero value otherwise.

### GetScanDateOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetScanDateOk() (*time.Time, bool)`

GetScanDateOk returns a tuple with the ScanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDate

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetScanDate(v time.Time)`

SetScanDate sets ScanDate field to given value.

### HasScanDate

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasScanDate() bool`

HasScanDate returns a boolean if a field has been set.

### GetScanDuration

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetScanDuration() int64`

GetScanDuration returns the ScanDuration field if non-nil, zero value otherwise.

### GetScanDurationOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetScanDurationOk() (*int64, bool)`

GetScanDurationOk returns a tuple with the ScanDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDuration

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetScanDuration(v int64)`

SetScanDuration sets ScanDuration field to given value.

### HasScanDuration

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasScanDuration() bool`

HasScanDuration returns a boolean if a field has been set.

### GetTestCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetTestCount() int64`

GetTestCount returns the TestCount field if non-nil, zero value otherwise.

### GetTestCountOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetTestCountOk() (*int64, bool)`

GetTestCountOk returns a tuple with the TestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetTestCount(v int64)`

SetTestCount sets TestCount field to given value.

### HasTestCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasTestCount() bool`

HasTestCount returns a boolean if a field has been set.

### GetRunCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetRunCount() int64`

GetRunCount returns the RunCount field if non-nil, zero value otherwise.

### GetRunCountOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetRunCountOk() (*int64, bool)`

GetRunCountOk returns a tuple with the RunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetRunCount(v int64)`

SetRunCount sets RunCount field to given value.

### HasRunCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasRunCount() bool`

HasRunCount returns a boolean if a field has been set.

### GetPassCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetPassCount() int64`

GetPassCount returns the PassCount field if non-nil, zero value otherwise.

### GetPassCountOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetPassCountOk() (*int64, bool)`

GetPassCountOk returns a tuple with the PassCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetPassCount(v int64)`

SetPassCount sets PassCount field to given value.

### HasPassCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasPassCount() bool`

HasPassCount returns a boolean if a field has been set.

### GetFailCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetFailCount() int64`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetFailCountOk() (*int64, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetFailCount(v int64)`

SetFailCount sets FailCount field to given value.

### HasFailCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasFailCount() bool`

HasFailCount returns a boolean if a field has been set.

### GetOtherCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetOtherCount() int64`

GetOtherCount returns the OtherCount field if non-nil, zero value otherwise.

### GetOtherCountOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetOtherCountOk() (*int64, bool)`

GetOtherCountOk returns a tuple with the OtherCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetOtherCount(v int64)`

SetOtherCount sets OtherCount field to given value.

### HasOtherCount

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasOtherCount() bool`

HasOtherCount returns a boolean if a field has been set.

### GetScanScore

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetScanScore() float32`

GetScanScore returns the ScanScore field if non-nil, zero value otherwise.

### GetScanScoreOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetScanScoreOk() (*float32, bool)`

GetScanScoreOk returns a tuple with the ScanScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanScore

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetScanScore(v float32)`

SetScanScore sets ScanScore field to given value.

### HasScanScore

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasScanScore() bool`

HasScanScore returns a boolean if a field has been set.

### GetExternalId

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetSecurityScans200ResponseAllOfSecurityScan) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetCreatedBy

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetResults

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *GetSecurityScans200ResponseAllOfSecurityScan) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


