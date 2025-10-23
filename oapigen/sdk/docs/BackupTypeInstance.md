# BackupTypeInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationType** | **string** |  | 
**Name** | **string** | A name for the backup | 
**InstanceId** | **int64** | The ID of the instance to backup | 
**ContainerId** | **int64** | The ID of the container to backup | 
**BackupType** | **string** | The backup type code, options vary by the type of cloud and source | 
**JobAction** | **string** | Create a new backup job, clone an existing job or add the new backup to an existing job | 
**JobId** | Pointer to **int64** | The ID of the job to clone or add to. Only applies to jobAction &#x60;clone&#x60; and &#x60;addTo&#x60;. | [optional] 
**JobName** | Pointer to **string** | Name for new job. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**JobSchedule** | Pointer to **int64** | The ID of the execute schedule for new job. See Execute Schedules. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**RetentionCount** | Pointer to **int64** | Retention Count for new job. By default the backup settings value will be used. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**BackupJob** | Pointer to [**BackupInstanceBackupJob**](BackupInstanceBackupJob.md) |  | [optional] 

## Methods

### NewBackupTypeInstance

`func NewBackupTypeInstance(locationType string, name string, instanceId int64, containerId int64, backupType string, jobAction string, ) *BackupTypeInstance`

NewBackupTypeInstance instantiates a new BackupTypeInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupTypeInstanceWithDefaults

`func NewBackupTypeInstanceWithDefaults() *BackupTypeInstance`

NewBackupTypeInstanceWithDefaults instantiates a new BackupTypeInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationType

`func (o *BackupTypeInstance) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *BackupTypeInstance) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *BackupTypeInstance) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.


### GetName

`func (o *BackupTypeInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupTypeInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupTypeInstance) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceId

`func (o *BackupTypeInstance) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *BackupTypeInstance) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *BackupTypeInstance) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetContainerId

`func (o *BackupTypeInstance) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *BackupTypeInstance) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *BackupTypeInstance) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.


### GetBackupType

`func (o *BackupTypeInstance) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupTypeInstance) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupTypeInstance) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.


### GetJobAction

`func (o *BackupTypeInstance) GetJobAction() string`

GetJobAction returns the JobAction field if non-nil, zero value otherwise.

### GetJobActionOk

`func (o *BackupTypeInstance) GetJobActionOk() (*string, bool)`

GetJobActionOk returns a tuple with the JobAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobAction

`func (o *BackupTypeInstance) SetJobAction(v string)`

SetJobAction sets JobAction field to given value.


### GetJobId

`func (o *BackupTypeInstance) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BackupTypeInstance) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BackupTypeInstance) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BackupTypeInstance) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobName

`func (o *BackupTypeInstance) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *BackupTypeInstance) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *BackupTypeInstance) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *BackupTypeInstance) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetJobSchedule

`func (o *BackupTypeInstance) GetJobSchedule() int64`

GetJobSchedule returns the JobSchedule field if non-nil, zero value otherwise.

### GetJobScheduleOk

`func (o *BackupTypeInstance) GetJobScheduleOk() (*int64, bool)`

GetJobScheduleOk returns a tuple with the JobSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSchedule

`func (o *BackupTypeInstance) SetJobSchedule(v int64)`

SetJobSchedule sets JobSchedule field to given value.

### HasJobSchedule

`func (o *BackupTypeInstance) HasJobSchedule() bool`

HasJobSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *BackupTypeInstance) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *BackupTypeInstance) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *BackupTypeInstance) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *BackupTypeInstance) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetBackupJob

`func (o *BackupTypeInstance) GetBackupJob() BackupInstanceBackupJob`

GetBackupJob returns the BackupJob field if non-nil, zero value otherwise.

### GetBackupJobOk

`func (o *BackupTypeInstance) GetBackupJobOk() (*BackupInstanceBackupJob, bool)`

GetBackupJobOk returns a tuple with the BackupJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJob

`func (o *BackupTypeInstance) SetBackupJob(v BackupInstanceBackupJob)`

SetBackupJob sets BackupJob field to given value.

### HasBackupJob

`func (o *BackupTypeInstance) HasBackupJob() bool`

HasBackupJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


