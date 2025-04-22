# ListHealth200ResponseAllOfHealthDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ConnectionList** | Pointer to **[]map[string]interface{}** |  | [optional] 
**BusyConnections** | Pointer to **[]string** |  | [optional] 
**MaxConnections** | Pointer to **int64** |  | [optional] 
**MaxUsedConnections** | Pointer to **int64** |  | [optional] 
**UsedConnections** | Pointer to **int64** |  | [optional] 
**AbortedConnections** | Pointer to **int64** |  | [optional] 
**InnodbStatus** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**ListHealth200ResponseAllOfHealthDatabaseStats**](ListHealth200ResponseAllOfHealthDatabaseStats.md) |  | [optional] 
**Scans** | Pointer to [**ListHealth200ResponseAllOfHealthDatabaseScans**](ListHealth200ResponseAllOfHealthDatabaseScans.md) |  | [optional] 
**SlowQueries** | Pointer to [**[]ListHealth200ResponseAllOfHealthDatabaseSlowQueriesInner**](ListHealth200ResponseAllOfHealthDatabaseSlowQueriesInner.md) |  | [optional] 
**InnodbStats** | Pointer to [**ListHealth200ResponseAllOfHealthDatabaseInnodbStats**](ListHealth200ResponseAllOfHealthDatabaseInnodbStats.md) |  | [optional] 
**ScanPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewListHealth200ResponseAllOfHealthDatabase

`func NewListHealth200ResponseAllOfHealthDatabase() *ListHealth200ResponseAllOfHealthDatabase`

NewListHealth200ResponseAllOfHealthDatabase instantiates a new ListHealth200ResponseAllOfHealthDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHealth200ResponseAllOfHealthDatabaseWithDefaults

`func NewListHealth200ResponseAllOfHealthDatabaseWithDefaults() *ListHealth200ResponseAllOfHealthDatabase`

NewListHealth200ResponseAllOfHealthDatabaseWithDefaults instantiates a new ListHealth200ResponseAllOfHealthDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetConnectionList

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetConnectionList() []map[string]interface{}`

GetConnectionList returns the ConnectionList field if non-nil, zero value otherwise.

### GetConnectionListOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetConnectionListOk() (*[]map[string]interface{}, bool)`

GetConnectionListOk returns a tuple with the ConnectionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionList

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetConnectionList(v []map[string]interface{})`

SetConnectionList sets ConnectionList field to given value.

### HasConnectionList

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasConnectionList() bool`

HasConnectionList returns a boolean if a field has been set.

### GetBusyConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetBusyConnections() []string`

GetBusyConnections returns the BusyConnections field if non-nil, zero value otherwise.

### GetBusyConnectionsOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetBusyConnectionsOk() (*[]string, bool)`

GetBusyConnectionsOk returns a tuple with the BusyConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetBusyConnections(v []string)`

SetBusyConnections sets BusyConnections field to given value.

### HasBusyConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasBusyConnections() bool`

HasBusyConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetMaxUsedConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetMaxUsedConnections() int64`

GetMaxUsedConnections returns the MaxUsedConnections field if non-nil, zero value otherwise.

### GetMaxUsedConnectionsOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetMaxUsedConnectionsOk() (*int64, bool)`

GetMaxUsedConnectionsOk returns a tuple with the MaxUsedConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsedConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetMaxUsedConnections(v int64)`

SetMaxUsedConnections sets MaxUsedConnections field to given value.

### HasMaxUsedConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasMaxUsedConnections() bool`

HasMaxUsedConnections returns a boolean if a field has been set.

### GetUsedConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetUsedConnections() int64`

GetUsedConnections returns the UsedConnections field if non-nil, zero value otherwise.

### GetUsedConnectionsOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetUsedConnectionsOk() (*int64, bool)`

GetUsedConnectionsOk returns a tuple with the UsedConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetUsedConnections(v int64)`

SetUsedConnections sets UsedConnections field to given value.

### HasUsedConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasUsedConnections() bool`

HasUsedConnections returns a boolean if a field has been set.

### GetAbortedConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetAbortedConnections() int64`

GetAbortedConnections returns the AbortedConnections field if non-nil, zero value otherwise.

### GetAbortedConnectionsOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetAbortedConnectionsOk() (*int64, bool)`

GetAbortedConnectionsOk returns a tuple with the AbortedConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortedConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetAbortedConnections(v int64)`

SetAbortedConnections sets AbortedConnections field to given value.

### HasAbortedConnections

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasAbortedConnections() bool`

HasAbortedConnections returns a boolean if a field has been set.

### GetInnodbStatus

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetInnodbStatus() string`

GetInnodbStatus returns the InnodbStatus field if non-nil, zero value otherwise.

### GetInnodbStatusOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetInnodbStatusOk() (*string, bool)`

GetInnodbStatusOk returns a tuple with the InnodbStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbStatus

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetInnodbStatus(v string)`

SetInnodbStatus sets InnodbStatus field to given value.

### HasInnodbStatus

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasInnodbStatus() bool`

HasInnodbStatus returns a boolean if a field has been set.

### GetStats

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetStats() ListHealth200ResponseAllOfHealthDatabaseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetStatsOk() (*ListHealth200ResponseAllOfHealthDatabaseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetStats(v ListHealth200ResponseAllOfHealthDatabaseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetScans

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetScans() ListHealth200ResponseAllOfHealthDatabaseScans`

GetScans returns the Scans field if non-nil, zero value otherwise.

### GetScansOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetScansOk() (*ListHealth200ResponseAllOfHealthDatabaseScans, bool)`

GetScansOk returns a tuple with the Scans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScans

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetScans(v ListHealth200ResponseAllOfHealthDatabaseScans)`

SetScans sets Scans field to given value.

### HasScans

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasScans() bool`

HasScans returns a boolean if a field has been set.

### GetSlowQueries

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetSlowQueries() []ListHealth200ResponseAllOfHealthDatabaseSlowQueriesInner`

GetSlowQueries returns the SlowQueries field if non-nil, zero value otherwise.

### GetSlowQueriesOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetSlowQueriesOk() (*[]ListHealth200ResponseAllOfHealthDatabaseSlowQueriesInner, bool)`

GetSlowQueriesOk returns a tuple with the SlowQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowQueries

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetSlowQueries(v []ListHealth200ResponseAllOfHealthDatabaseSlowQueriesInner)`

SetSlowQueries sets SlowQueries field to given value.

### HasSlowQueries

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasSlowQueries() bool`

HasSlowQueries returns a boolean if a field has been set.

### GetInnodbStats

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetInnodbStats() ListHealth200ResponseAllOfHealthDatabaseInnodbStats`

GetInnodbStats returns the InnodbStats field if non-nil, zero value otherwise.

### GetInnodbStatsOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetInnodbStatsOk() (*ListHealth200ResponseAllOfHealthDatabaseInnodbStats, bool)`

GetInnodbStatsOk returns a tuple with the InnodbStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbStats

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetInnodbStats(v ListHealth200ResponseAllOfHealthDatabaseInnodbStats)`

SetInnodbStats sets InnodbStats field to given value.

### HasInnodbStats

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasInnodbStats() bool`

HasInnodbStats returns a boolean if a field has been set.

### GetScanPercent

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetScanPercent() float32`

GetScanPercent returns the ScanPercent field if non-nil, zero value otherwise.

### GetScanPercentOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetScanPercentOk() (*float32, bool)`

GetScanPercentOk returns a tuple with the ScanPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPercent

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetScanPercent(v float32)`

SetScanPercent sets ScanPercent field to given value.

### HasScanPercent

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasScanPercent() bool`

HasScanPercent returns a boolean if a field has been set.

### GetStatus

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHealth200ResponseAllOfHealthDatabase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHealth200ResponseAllOfHealthDatabase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListHealth200ResponseAllOfHealthDatabase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


