# ListSecurityScans200ResponseAllOfSecurityScansInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**SecurityPackage** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerCluster**](ListInstances200ResponseAllOfInstancesInnerCluster.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ScanDate** | Pointer to **time.Time** |  | [optional] 
**ScanDuration** | Pointer to **int64** |  | [optional] 
**TestCount** | Pointer to **int64** |  | [optional] 
**RunCount** | Pointer to **int64** |  | [optional] 
**PassCount** | Pointer to **int64** |  | [optional] 
**FailCount** | Pointer to **int64** |  | [optional] 
**OtherCount** | Pointer to **int64** |  | [optional] 
**ScanScore** | Pointer to **float32** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Results** | Pointer to **map[string]interface{}** | Results Summary (only returned when using query parameter results&#x3D;true) | [optional] 

## Methods

### NewListSecurityScans200ResponseAllOfSecurityScansInner

`func NewListSecurityScans200ResponseAllOfSecurityScansInner() *ListSecurityScans200ResponseAllOfSecurityScansInner`

NewListSecurityScans200ResponseAllOfSecurityScansInner instantiates a new ListSecurityScans200ResponseAllOfSecurityScansInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityScans200ResponseAllOfSecurityScansInnerWithDefaults

`func NewListSecurityScans200ResponseAllOfSecurityScansInnerWithDefaults() *ListSecurityScans200ResponseAllOfSecurityScansInner`

NewListSecurityScans200ResponseAllOfSecurityScansInnerWithDefaults instantiates a new ListSecurityScans200ResponseAllOfSecurityScansInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSecurityPackage

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetSecurityPackage() ListInstances200ResponseAllOfInstancesInnerCluster`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetSecurityPackageOk() (*ListInstances200ResponseAllOfInstancesInnerCluster, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetSecurityPackage(v ListInstances200ResponseAllOfInstancesInnerCluster)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### GetStatus

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScanDate

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetScanDate() time.Time`

GetScanDate returns the ScanDate field if non-nil, zero value otherwise.

### GetScanDateOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetScanDateOk() (*time.Time, bool)`

GetScanDateOk returns a tuple with the ScanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDate

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetScanDate(v time.Time)`

SetScanDate sets ScanDate field to given value.

### HasScanDate

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasScanDate() bool`

HasScanDate returns a boolean if a field has been set.

### GetScanDuration

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetScanDuration() int64`

GetScanDuration returns the ScanDuration field if non-nil, zero value otherwise.

### GetScanDurationOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetScanDurationOk() (*int64, bool)`

GetScanDurationOk returns a tuple with the ScanDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDuration

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetScanDuration(v int64)`

SetScanDuration sets ScanDuration field to given value.

### HasScanDuration

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasScanDuration() bool`

HasScanDuration returns a boolean if a field has been set.

### GetTestCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetTestCount() int64`

GetTestCount returns the TestCount field if non-nil, zero value otherwise.

### GetTestCountOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetTestCountOk() (*int64, bool)`

GetTestCountOk returns a tuple with the TestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetTestCount(v int64)`

SetTestCount sets TestCount field to given value.

### HasTestCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasTestCount() bool`

HasTestCount returns a boolean if a field has been set.

### GetRunCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetRunCount() int64`

GetRunCount returns the RunCount field if non-nil, zero value otherwise.

### GetRunCountOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetRunCountOk() (*int64, bool)`

GetRunCountOk returns a tuple with the RunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetRunCount(v int64)`

SetRunCount sets RunCount field to given value.

### HasRunCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasRunCount() bool`

HasRunCount returns a boolean if a field has been set.

### GetPassCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetPassCount() int64`

GetPassCount returns the PassCount field if non-nil, zero value otherwise.

### GetPassCountOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetPassCountOk() (*int64, bool)`

GetPassCountOk returns a tuple with the PassCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetPassCount(v int64)`

SetPassCount sets PassCount field to given value.

### HasPassCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasPassCount() bool`

HasPassCount returns a boolean if a field has been set.

### GetFailCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetFailCount() int64`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetFailCountOk() (*int64, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetFailCount(v int64)`

SetFailCount sets FailCount field to given value.

### HasFailCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasFailCount() bool`

HasFailCount returns a boolean if a field has been set.

### GetOtherCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetOtherCount() int64`

GetOtherCount returns the OtherCount field if non-nil, zero value otherwise.

### GetOtherCountOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetOtherCountOk() (*int64, bool)`

GetOtherCountOk returns a tuple with the OtherCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetOtherCount(v int64)`

SetOtherCount sets OtherCount field to given value.

### HasOtherCount

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasOtherCount() bool`

HasOtherCount returns a boolean if a field has been set.

### GetScanScore

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetScanScore() float32`

GetScanScore returns the ScanScore field if non-nil, zero value otherwise.

### GetScanScoreOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetScanScoreOk() (*float32, bool)`

GetScanScoreOk returns a tuple with the ScanScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanScore

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetScanScore(v float32)`

SetScanScore sets ScanScore field to given value.

### HasScanScore

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasScanScore() bool`

HasScanScore returns a boolean if a field has been set.

### GetExternalId

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetResults

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *ListSecurityScans200ResponseAllOfSecurityScansInner) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


