# HealthElasticStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**ClusterName** | Pointer to **string** |  | [optional] 
**NodeTotal** | Pointer to **string** |  | [optional] 
**NodeData** | Pointer to **string** |  | [optional] 
**Shards** | Pointer to **string** |  | [optional] 
**Primary** | Pointer to **string** |  | [optional] 
**Relocating** | Pointer to **string** |  | [optional] 
**Initializing** | Pointer to **string** |  | [optional] 
**Unassigned** | Pointer to **string** |  | [optional] 
**PendingTasks** | Pointer to **string** |  | [optional] 
**ActivePercent** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthElasticStats

`func NewHealthElasticStats() *HealthElasticStats`

NewHealthElasticStats instantiates a new HealthElasticStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthElasticStatsWithDefaults

`func NewHealthElasticStatsWithDefaults() *HealthElasticStats`

NewHealthElasticStatsWithDefaults instantiates a new HealthElasticStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthElasticStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthElasticStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthElasticStats) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthElasticStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClusterName

`func (o *HealthElasticStats) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HealthElasticStats) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HealthElasticStats) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HealthElasticStats) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetNodeTotal

`func (o *HealthElasticStats) GetNodeTotal() string`

GetNodeTotal returns the NodeTotal field if non-nil, zero value otherwise.

### GetNodeTotalOk

`func (o *HealthElasticStats) GetNodeTotalOk() (*string, bool)`

GetNodeTotalOk returns a tuple with the NodeTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTotal

`func (o *HealthElasticStats) SetNodeTotal(v string)`

SetNodeTotal sets NodeTotal field to given value.

### HasNodeTotal

`func (o *HealthElasticStats) HasNodeTotal() bool`

HasNodeTotal returns a boolean if a field has been set.

### GetNodeData

`func (o *HealthElasticStats) GetNodeData() string`

GetNodeData returns the NodeData field if non-nil, zero value otherwise.

### GetNodeDataOk

`func (o *HealthElasticStats) GetNodeDataOk() (*string, bool)`

GetNodeDataOk returns a tuple with the NodeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeData

`func (o *HealthElasticStats) SetNodeData(v string)`

SetNodeData sets NodeData field to given value.

### HasNodeData

`func (o *HealthElasticStats) HasNodeData() bool`

HasNodeData returns a boolean if a field has been set.

### GetShards

`func (o *HealthElasticStats) GetShards() string`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *HealthElasticStats) GetShardsOk() (*string, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *HealthElasticStats) SetShards(v string)`

SetShards sets Shards field to given value.

### HasShards

`func (o *HealthElasticStats) HasShards() bool`

HasShards returns a boolean if a field has been set.

### GetPrimary

`func (o *HealthElasticStats) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *HealthElasticStats) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *HealthElasticStats) SetPrimary(v string)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *HealthElasticStats) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetRelocating

`func (o *HealthElasticStats) GetRelocating() string`

GetRelocating returns the Relocating field if non-nil, zero value otherwise.

### GetRelocatingOk

`func (o *HealthElasticStats) GetRelocatingOk() (*string, bool)`

GetRelocatingOk returns a tuple with the Relocating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocating

`func (o *HealthElasticStats) SetRelocating(v string)`

SetRelocating sets Relocating field to given value.

### HasRelocating

`func (o *HealthElasticStats) HasRelocating() bool`

HasRelocating returns a boolean if a field has been set.

### GetInitializing

`func (o *HealthElasticStats) GetInitializing() string`

GetInitializing returns the Initializing field if non-nil, zero value otherwise.

### GetInitializingOk

`func (o *HealthElasticStats) GetInitializingOk() (*string, bool)`

GetInitializingOk returns a tuple with the Initializing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializing

`func (o *HealthElasticStats) SetInitializing(v string)`

SetInitializing sets Initializing field to given value.

### HasInitializing

`func (o *HealthElasticStats) HasInitializing() bool`

HasInitializing returns a boolean if a field has been set.

### GetUnassigned

`func (o *HealthElasticStats) GetUnassigned() string`

GetUnassigned returns the Unassigned field if non-nil, zero value otherwise.

### GetUnassignedOk

`func (o *HealthElasticStats) GetUnassignedOk() (*string, bool)`

GetUnassignedOk returns a tuple with the Unassigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassigned

`func (o *HealthElasticStats) SetUnassigned(v string)`

SetUnassigned sets Unassigned field to given value.

### HasUnassigned

`func (o *HealthElasticStats) HasUnassigned() bool`

HasUnassigned returns a boolean if a field has been set.

### GetPendingTasks

`func (o *HealthElasticStats) GetPendingTasks() string`

GetPendingTasks returns the PendingTasks field if non-nil, zero value otherwise.

### GetPendingTasksOk

`func (o *HealthElasticStats) GetPendingTasksOk() (*string, bool)`

GetPendingTasksOk returns a tuple with the PendingTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTasks

`func (o *HealthElasticStats) SetPendingTasks(v string)`

SetPendingTasks sets PendingTasks field to given value.

### HasPendingTasks

`func (o *HealthElasticStats) HasPendingTasks() bool`

HasPendingTasks returns a boolean if a field has been set.

### GetActivePercent

`func (o *HealthElasticStats) GetActivePercent() string`

GetActivePercent returns the ActivePercent field if non-nil, zero value otherwise.

### GetActivePercentOk

`func (o *HealthElasticStats) GetActivePercentOk() (*string, bool)`

GetActivePercentOk returns a tuple with the ActivePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePercent

`func (o *HealthElasticStats) SetActivePercent(v string)`

SetActivePercent sets ActivePercent field to given value.

### HasActivePercent

`func (o *HealthElasticStats) HasActivePercent() bool`

HasActivePercent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


