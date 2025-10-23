# AddBackupsRequestBackup

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
**ServerId** | Pointer to **int64** | The ID of the server to backup | [optional] 
**SourceProviderId** | Pointer to **int64** | Source Object Store. The ID of the storage provider (bucket) to be backed up. | [optional] 
**StorageProviderId** | Pointer to **int64** | Target Object Store. The ID of the storage provider (bucket) the backup will be copied to. | [optional] 

## Methods

### NewAddBackupsRequestBackup

`func NewAddBackupsRequestBackup(locationType string, name string, instanceId int64, containerId int64, backupType string, jobAction string, ) *AddBackupsRequestBackup`

NewAddBackupsRequestBackup instantiates a new AddBackupsRequestBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBackupsRequestBackupWithDefaults

`func NewAddBackupsRequestBackupWithDefaults() *AddBackupsRequestBackup`

NewAddBackupsRequestBackupWithDefaults instantiates a new AddBackupsRequestBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationType

`func (o *AddBackupsRequestBackup) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *AddBackupsRequestBackup) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *AddBackupsRequestBackup) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.


### GetName

`func (o *AddBackupsRequestBackup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBackupsRequestBackup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBackupsRequestBackup) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceId

`func (o *AddBackupsRequestBackup) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AddBackupsRequestBackup) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AddBackupsRequestBackup) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetContainerId

`func (o *AddBackupsRequestBackup) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *AddBackupsRequestBackup) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *AddBackupsRequestBackup) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.


### GetBackupType

`func (o *AddBackupsRequestBackup) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *AddBackupsRequestBackup) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *AddBackupsRequestBackup) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.


### GetJobAction

`func (o *AddBackupsRequestBackup) GetJobAction() string`

GetJobAction returns the JobAction field if non-nil, zero value otherwise.

### GetJobActionOk

`func (o *AddBackupsRequestBackup) GetJobActionOk() (*string, bool)`

GetJobActionOk returns a tuple with the JobAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobAction

`func (o *AddBackupsRequestBackup) SetJobAction(v string)`

SetJobAction sets JobAction field to given value.


### GetJobId

`func (o *AddBackupsRequestBackup) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *AddBackupsRequestBackup) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *AddBackupsRequestBackup) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *AddBackupsRequestBackup) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobName

`func (o *AddBackupsRequestBackup) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *AddBackupsRequestBackup) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *AddBackupsRequestBackup) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *AddBackupsRequestBackup) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetJobSchedule

`func (o *AddBackupsRequestBackup) GetJobSchedule() int64`

GetJobSchedule returns the JobSchedule field if non-nil, zero value otherwise.

### GetJobScheduleOk

`func (o *AddBackupsRequestBackup) GetJobScheduleOk() (*int64, bool)`

GetJobScheduleOk returns a tuple with the JobSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSchedule

`func (o *AddBackupsRequestBackup) SetJobSchedule(v int64)`

SetJobSchedule sets JobSchedule field to given value.

### HasJobSchedule

`func (o *AddBackupsRequestBackup) HasJobSchedule() bool`

HasJobSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *AddBackupsRequestBackup) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *AddBackupsRequestBackup) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *AddBackupsRequestBackup) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *AddBackupsRequestBackup) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetBackupJob

`func (o *AddBackupsRequestBackup) GetBackupJob() BackupInstanceBackupJob`

GetBackupJob returns the BackupJob field if non-nil, zero value otherwise.

### GetBackupJobOk

`func (o *AddBackupsRequestBackup) GetBackupJobOk() (*BackupInstanceBackupJob, bool)`

GetBackupJobOk returns a tuple with the BackupJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupJob

`func (o *AddBackupsRequestBackup) SetBackupJob(v BackupInstanceBackupJob)`

SetBackupJob sets BackupJob field to given value.

### HasBackupJob

`func (o *AddBackupsRequestBackup) HasBackupJob() bool`

HasBackupJob returns a boolean if a field has been set.

### GetServerId

`func (o *AddBackupsRequestBackup) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *AddBackupsRequestBackup) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *AddBackupsRequestBackup) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *AddBackupsRequestBackup) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetSourceProviderId

`func (o *AddBackupsRequestBackup) GetSourceProviderId() int64`

GetSourceProviderId returns the SourceProviderId field if non-nil, zero value otherwise.

### GetSourceProviderIdOk

`func (o *AddBackupsRequestBackup) GetSourceProviderIdOk() (*int64, bool)`

GetSourceProviderIdOk returns a tuple with the SourceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProviderId

`func (o *AddBackupsRequestBackup) SetSourceProviderId(v int64)`

SetSourceProviderId sets SourceProviderId field to given value.

### HasSourceProviderId

`func (o *AddBackupsRequestBackup) HasSourceProviderId() bool`

HasSourceProviderId returns a boolean if a field has been set.

### GetStorageProviderId

`func (o *AddBackupsRequestBackup) GetStorageProviderId() int64`

GetStorageProviderId returns the StorageProviderId field if non-nil, zero value otherwise.

### GetStorageProviderIdOk

`func (o *AddBackupsRequestBackup) GetStorageProviderIdOk() (*int64, bool)`

GetStorageProviderIdOk returns a tuple with the StorageProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProviderId

`func (o *AddBackupsRequestBackup) SetStorageProviderId(v int64)`

SetStorageProviderId sets StorageProviderId field to given value.

### HasStorageProviderId

`func (o *AddBackupsRequestBackup) HasStorageProviderId() bool`

HasStorageProviderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


