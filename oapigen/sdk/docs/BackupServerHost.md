# BackupServerHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationType** | **string** |  | 
**Name** | **string** | A name for the backup | 
**ServerId** | Pointer to **int64** | The ID of the server to backup | [optional] 
**BackupType** | **string** | The backup type code, options vary by the type of cloud and source | 
**JobAction** | **string** | Create a new backup job, clone an existing job or add the new backup to an existing job | 
**JobId** | Pointer to **int64** | The ID of the job to clone or add to. Only applies to jobAction &#x60;clone&#x60; and &#x60;addTo&#x60;. | [optional] 
**JobName** | Pointer to **string** | Name for new job. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**JobSchedule** | Pointer to **int64** | The ID of the execute schedule for new job. See Execute Schedules. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**RetentionCount** | Pointer to **int64** | Retention Count for new job. By default the backup settings value will be used. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 

## Methods

### NewBackupServerHost

`func NewBackupServerHost(locationType string, name string, backupType string, jobAction string, ) *BackupServerHost`

NewBackupServerHost instantiates a new BackupServerHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupServerHostWithDefaults

`func NewBackupServerHostWithDefaults() *BackupServerHost`

NewBackupServerHostWithDefaults instantiates a new BackupServerHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationType

`func (o *BackupServerHost) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *BackupServerHost) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *BackupServerHost) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.


### GetName

`func (o *BackupServerHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupServerHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupServerHost) SetName(v string)`

SetName sets Name field to given value.


### GetServerId

`func (o *BackupServerHost) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *BackupServerHost) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *BackupServerHost) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *BackupServerHost) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetBackupType

`func (o *BackupServerHost) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupServerHost) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupServerHost) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.


### GetJobAction

`func (o *BackupServerHost) GetJobAction() string`

GetJobAction returns the JobAction field if non-nil, zero value otherwise.

### GetJobActionOk

`func (o *BackupServerHost) GetJobActionOk() (*string, bool)`

GetJobActionOk returns a tuple with the JobAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobAction

`func (o *BackupServerHost) SetJobAction(v string)`

SetJobAction sets JobAction field to given value.


### GetJobId

`func (o *BackupServerHost) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BackupServerHost) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BackupServerHost) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BackupServerHost) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobName

`func (o *BackupServerHost) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *BackupServerHost) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *BackupServerHost) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *BackupServerHost) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetJobSchedule

`func (o *BackupServerHost) GetJobSchedule() int64`

GetJobSchedule returns the JobSchedule field if non-nil, zero value otherwise.

### GetJobScheduleOk

`func (o *BackupServerHost) GetJobScheduleOk() (*int64, bool)`

GetJobScheduleOk returns a tuple with the JobSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSchedule

`func (o *BackupServerHost) SetJobSchedule(v int64)`

SetJobSchedule sets JobSchedule field to given value.

### HasJobSchedule

`func (o *BackupServerHost) HasJobSchedule() bool`

HasJobSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *BackupServerHost) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *BackupServerHost) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *BackupServerHost) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *BackupServerHost) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


