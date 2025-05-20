# SecurityScan

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

### NewSecurityScan

`func NewSecurityScan() *SecurityScan`

NewSecurityScan instantiates a new SecurityScan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanWithDefaults

`func NewSecurityScanWithDefaults() *SecurityScan`

NewSecurityScanWithDefaults instantiates a new SecurityScan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityScan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityScan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityScan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityScan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSecurityPackage

`func (o *SecurityScan) GetSecurityPackage() ListInstances200ResponseAllOfInstancesInnerCluster`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *SecurityScan) GetSecurityPackageOk() (*ListInstances200ResponseAllOfInstancesInnerCluster, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *SecurityScan) SetSecurityPackage(v ListInstances200ResponseAllOfInstancesInnerCluster)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *SecurityScan) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityScan) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityScan) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityScan) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityScan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScanDate

`func (o *SecurityScan) GetScanDate() time.Time`

GetScanDate returns the ScanDate field if non-nil, zero value otherwise.

### GetScanDateOk

`func (o *SecurityScan) GetScanDateOk() (*time.Time, bool)`

GetScanDateOk returns a tuple with the ScanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDate

`func (o *SecurityScan) SetScanDate(v time.Time)`

SetScanDate sets ScanDate field to given value.

### HasScanDate

`func (o *SecurityScan) HasScanDate() bool`

HasScanDate returns a boolean if a field has been set.

### GetScanDuration

`func (o *SecurityScan) GetScanDuration() int64`

GetScanDuration returns the ScanDuration field if non-nil, zero value otherwise.

### GetScanDurationOk

`func (o *SecurityScan) GetScanDurationOk() (*int64, bool)`

GetScanDurationOk returns a tuple with the ScanDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDuration

`func (o *SecurityScan) SetScanDuration(v int64)`

SetScanDuration sets ScanDuration field to given value.

### HasScanDuration

`func (o *SecurityScan) HasScanDuration() bool`

HasScanDuration returns a boolean if a field has been set.

### GetTestCount

`func (o *SecurityScan) GetTestCount() int64`

GetTestCount returns the TestCount field if non-nil, zero value otherwise.

### GetTestCountOk

`func (o *SecurityScan) GetTestCountOk() (*int64, bool)`

GetTestCountOk returns a tuple with the TestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCount

`func (o *SecurityScan) SetTestCount(v int64)`

SetTestCount sets TestCount field to given value.

### HasTestCount

`func (o *SecurityScan) HasTestCount() bool`

HasTestCount returns a boolean if a field has been set.

### GetRunCount

`func (o *SecurityScan) GetRunCount() int64`

GetRunCount returns the RunCount field if non-nil, zero value otherwise.

### GetRunCountOk

`func (o *SecurityScan) GetRunCountOk() (*int64, bool)`

GetRunCountOk returns a tuple with the RunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCount

`func (o *SecurityScan) SetRunCount(v int64)`

SetRunCount sets RunCount field to given value.

### HasRunCount

`func (o *SecurityScan) HasRunCount() bool`

HasRunCount returns a boolean if a field has been set.

### GetPassCount

`func (o *SecurityScan) GetPassCount() int64`

GetPassCount returns the PassCount field if non-nil, zero value otherwise.

### GetPassCountOk

`func (o *SecurityScan) GetPassCountOk() (*int64, bool)`

GetPassCountOk returns a tuple with the PassCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCount

`func (o *SecurityScan) SetPassCount(v int64)`

SetPassCount sets PassCount field to given value.

### HasPassCount

`func (o *SecurityScan) HasPassCount() bool`

HasPassCount returns a boolean if a field has been set.

### GetFailCount

`func (o *SecurityScan) GetFailCount() int64`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *SecurityScan) GetFailCountOk() (*int64, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCount

`func (o *SecurityScan) SetFailCount(v int64)`

SetFailCount sets FailCount field to given value.

### HasFailCount

`func (o *SecurityScan) HasFailCount() bool`

HasFailCount returns a boolean if a field has been set.

### GetOtherCount

`func (o *SecurityScan) GetOtherCount() int64`

GetOtherCount returns the OtherCount field if non-nil, zero value otherwise.

### GetOtherCountOk

`func (o *SecurityScan) GetOtherCountOk() (*int64, bool)`

GetOtherCountOk returns a tuple with the OtherCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCount

`func (o *SecurityScan) SetOtherCount(v int64)`

SetOtherCount sets OtherCount field to given value.

### HasOtherCount

`func (o *SecurityScan) HasOtherCount() bool`

HasOtherCount returns a boolean if a field has been set.

### GetScanScore

`func (o *SecurityScan) GetScanScore() float32`

GetScanScore returns the ScanScore field if non-nil, zero value otherwise.

### GetScanScoreOk

`func (o *SecurityScan) GetScanScoreOk() (*float32, bool)`

GetScanScoreOk returns a tuple with the ScanScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanScore

`func (o *SecurityScan) SetScanScore(v float32)`

SetScanScore sets ScanScore field to given value.

### HasScanScore

`func (o *SecurityScan) HasScanScore() bool`

HasScanScore returns a boolean if a field has been set.

### GetExternalId

`func (o *SecurityScan) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SecurityScan) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SecurityScan) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SecurityScan) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SecurityScan) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SecurityScan) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SecurityScan) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SecurityScan) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *SecurityScan) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SecurityScan) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SecurityScan) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SecurityScan) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *SecurityScan) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *SecurityScan) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *SecurityScan) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *SecurityScan) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *SecurityScan) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SecurityScan) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SecurityScan) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SecurityScan) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetResults

`func (o *SecurityScan) GetResults() map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SecurityScan) GetResultsOk() (*map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SecurityScan) SetResults(v map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *SecurityScan) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


