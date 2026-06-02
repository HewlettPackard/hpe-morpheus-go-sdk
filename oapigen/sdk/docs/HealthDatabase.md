# HealthDatabase

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
**InnodbStatus** | Pointer to **NullableString** |  | [optional] 
**Stats** | Pointer to [**HealthDatabaseStats**](HealthDatabaseStats.md) |  | [optional] 
**Scans** | Pointer to [**HealthDatabaseScans**](HealthDatabaseScans.md) |  | [optional] 
**SlowQueries** | Pointer to [**[]HealthDatabaseSlowQueriesInner**](HealthDatabaseSlowQueriesInner.md) |  | [optional] 
**InnodbStats** | Pointer to [**HealthDatabaseInnodbStats**](HealthDatabaseInnodbStats.md) |  | [optional] 
**ScanPercent** | Pointer to **float32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthDatabase

`func NewHealthDatabase() *HealthDatabase`

NewHealthDatabase instantiates a new HealthDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetSuccess

`func (o *HealthDatabase) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *HealthDatabase) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *HealthDatabase) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *HealthDatabase) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetConnectionList

`func (o *HealthDatabase) GetConnectionList() []map[string]interface{}`

GetConnectionList returns the ConnectionList field if non-nil, zero value otherwise.

### GetConnectionListOk

`func (o *HealthDatabase) GetConnectionListOk() (*[]map[string]interface{}, bool)`

GetConnectionListOk returns a tuple with the ConnectionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionList

`func (o *HealthDatabase) SetConnectionList(v []map[string]interface{})`

SetConnectionList sets ConnectionList field to given value.

### HasConnectionList

`func (o *HealthDatabase) HasConnectionList() bool`

HasConnectionList returns a boolean if a field has been set.

### SetConnectionListNil

`func (o *HealthDatabase) SetConnectionListNil(b bool)`

 SetConnectionListNil sets the value for ConnectionList to be an explicit nil

### UnsetConnectionList
`func (o *HealthDatabase) UnsetConnectionList()`

UnsetConnectionList ensures that no value is present for ConnectionList, not even an explicit nil
### GetBusyConnections

`func (o *HealthDatabase) GetBusyConnections() []*string`

GetBusyConnections returns the BusyConnections field if non-nil, zero value otherwise.

### GetBusyConnectionsOk

`func (o *HealthDatabase) GetBusyConnectionsOk() (*[]*string, bool)`

GetBusyConnectionsOk returns a tuple with the BusyConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyConnections

`func (o *HealthDatabase) SetBusyConnections(v []*string)`

SetBusyConnections sets BusyConnections field to given value.

### HasBusyConnections

`func (o *HealthDatabase) HasBusyConnections() bool`

HasBusyConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *HealthDatabase) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *HealthDatabase) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *HealthDatabase) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *HealthDatabase) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetMaxUsedConnections

`func (o *HealthDatabase) GetMaxUsedConnections() int64`

GetMaxUsedConnections returns the MaxUsedConnections field if non-nil, zero value otherwise.

### GetMaxUsedConnectionsOk

`func (o *HealthDatabase) GetMaxUsedConnectionsOk() (*int64, bool)`

GetMaxUsedConnectionsOk returns a tuple with the MaxUsedConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsedConnections

`func (o *HealthDatabase) SetMaxUsedConnections(v int64)`

SetMaxUsedConnections sets MaxUsedConnections field to given value.

### HasMaxUsedConnections

`func (o *HealthDatabase) HasMaxUsedConnections() bool`

HasMaxUsedConnections returns a boolean if a field has been set.

### GetUsedConnections

`func (o *HealthDatabase) GetUsedConnections() int64`

GetUsedConnections returns the UsedConnections field if non-nil, zero value otherwise.

### GetUsedConnectionsOk

`func (o *HealthDatabase) GetUsedConnectionsOk() (*int64, bool)`

GetUsedConnectionsOk returns a tuple with the UsedConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedConnections

`func (o *HealthDatabase) SetUsedConnections(v int64)`

SetUsedConnections sets UsedConnections field to given value.

### HasUsedConnections

`func (o *HealthDatabase) HasUsedConnections() bool`

HasUsedConnections returns a boolean if a field has been set.

### GetAbortedConnections

`func (o *HealthDatabase) GetAbortedConnections() int64`

GetAbortedConnections returns the AbortedConnections field if non-nil, zero value otherwise.

### GetAbortedConnectionsOk

`func (o *HealthDatabase) GetAbortedConnectionsOk() (*int64, bool)`

GetAbortedConnectionsOk returns a tuple with the AbortedConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortedConnections

`func (o *HealthDatabase) SetAbortedConnections(v int64)`

SetAbortedConnections sets AbortedConnections field to given value.

### HasAbortedConnections

`func (o *HealthDatabase) HasAbortedConnections() bool`

HasAbortedConnections returns a boolean if a field has been set.

### GetInnodbStatus

`func (o *HealthDatabase) GetInnodbStatus() string`

GetInnodbStatus returns the InnodbStatus field if non-nil, zero value otherwise.

### GetInnodbStatusOk

`func (o *HealthDatabase) GetInnodbStatusOk() (*string, bool)`

GetInnodbStatusOk returns a tuple with the InnodbStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbStatus

`func (o *HealthDatabase) SetInnodbStatus(v string)`

SetInnodbStatus sets InnodbStatus field to given value.

### HasInnodbStatus

`func (o *HealthDatabase) HasInnodbStatus() bool`

HasInnodbStatus returns a boolean if a field has been set.

### SetInnodbStatusNil

`func (o *HealthDatabase) SetInnodbStatusNil(b bool)`

 SetInnodbStatusNil sets the value for InnodbStatus to be an explicit nil

### UnsetInnodbStatus
`func (o *HealthDatabase) UnsetInnodbStatus()`

UnsetInnodbStatus ensures that no value is present for InnodbStatus, not even an explicit nil
### GetStats

`func (o *HealthDatabase) GetStats() HealthDatabaseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *HealthDatabase) GetStatsOk() (*HealthDatabaseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *HealthDatabase) SetStats(v HealthDatabaseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *HealthDatabase) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetScans

`func (o *HealthDatabase) GetScans() HealthDatabaseScans`

GetScans returns the Scans field if non-nil, zero value otherwise.

### GetScansOk

`func (o *HealthDatabase) GetScansOk() (*HealthDatabaseScans, bool)`

GetScansOk returns a tuple with the Scans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScans

`func (o *HealthDatabase) SetScans(v HealthDatabaseScans)`

SetScans sets Scans field to given value.

### HasScans

`func (o *HealthDatabase) HasScans() bool`

HasScans returns a boolean if a field has been set.

### GetSlowQueries

`func (o *HealthDatabase) GetSlowQueries() []HealthDatabaseSlowQueriesInner`

GetSlowQueries returns the SlowQueries field if non-nil, zero value otherwise.

### GetSlowQueriesOk

`func (o *HealthDatabase) GetSlowQueriesOk() (*[]HealthDatabaseSlowQueriesInner, bool)`

GetSlowQueriesOk returns a tuple with the SlowQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowQueries

`func (o *HealthDatabase) SetSlowQueries(v []HealthDatabaseSlowQueriesInner)`

SetSlowQueries sets SlowQueries field to given value.

### HasSlowQueries

`func (o *HealthDatabase) HasSlowQueries() bool`

HasSlowQueries returns a boolean if a field has been set.

### GetInnodbStats

`func (o *HealthDatabase) GetInnodbStats() HealthDatabaseInnodbStats`

GetInnodbStats returns the InnodbStats field if non-nil, zero value otherwise.

### GetInnodbStatsOk

`func (o *HealthDatabase) GetInnodbStatsOk() (*HealthDatabaseInnodbStats, bool)`

GetInnodbStatsOk returns a tuple with the InnodbStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnodbStats

`func (o *HealthDatabase) SetInnodbStats(v HealthDatabaseInnodbStats)`

SetInnodbStats sets InnodbStats field to given value.

### HasInnodbStats

`func (o *HealthDatabase) HasInnodbStats() bool`

HasInnodbStats returns a boolean if a field has been set.

### GetScanPercent

`func (o *HealthDatabase) GetScanPercent() float32`

GetScanPercent returns the ScanPercent field if non-nil, zero value otherwise.

### GetScanPercentOk

`func (o *HealthDatabase) GetScanPercentOk() (*float32, bool)`

GetScanPercentOk returns a tuple with the ScanPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPercent

`func (o *HealthDatabase) SetScanPercent(v float32)`

SetScanPercent sets ScanPercent field to given value.

### HasScanPercent

`func (o *HealthDatabase) HasScanPercent() bool`

HasScanPercent returns a boolean if a field has been set.

### GetStatus

`func (o *HealthDatabase) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthDatabase) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthDatabase) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthDatabase) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


