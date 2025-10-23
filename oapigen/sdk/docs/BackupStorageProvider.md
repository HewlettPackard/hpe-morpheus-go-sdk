# BackupStorageProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationType** | **string** |  | 
**Name** | **string** | A name for the backup | 
**SourceProviderId** | Pointer to **int64** | Source Object Store. The ID of the storage provider (bucket) to be backed up. | [optional] 
**StorageProviderId** | Pointer to **int64** | Target Object Store. The ID of the storage provider (bucket) the backup will be copied to. | [optional] 
**BackupType** | **string** | The backup type code, options vary by the type of cloud and source | 
**JobAction** | **string** | Create a new backup job, clone an existing job or add the new backup to an existing job | 
**JobId** | Pointer to **int64** | The ID of the job to clone or add to. Only applies to jobAction &#x60;clone&#x60; and &#x60;addTo&#x60;. | [optional] 
**JobName** | Pointer to **string** | Name for new job. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**JobSchedule** | Pointer to **int64** | The ID of the execute schedule for new job. See Execute Schedules. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**RetentionCount** | Pointer to **int64** | Retention Count for new job. By default the backup settings value will be used. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 

## Methods

### NewBackupStorageProvider

`func NewBackupStorageProvider(locationType string, name string, backupType string, jobAction string, ) *BackupStorageProvider`

NewBackupStorageProvider instantiates a new BackupStorageProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupStorageProviderWithDefaults

`func NewBackupStorageProviderWithDefaults() *BackupStorageProvider`

NewBackupStorageProviderWithDefaults instantiates a new BackupStorageProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationType

`func (o *BackupStorageProvider) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *BackupStorageProvider) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *BackupStorageProvider) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.


### GetName

`func (o *BackupStorageProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupStorageProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupStorageProvider) SetName(v string)`

SetName sets Name field to given value.


### GetSourceProviderId

`func (o *BackupStorageProvider) GetSourceProviderId() int64`

GetSourceProviderId returns the SourceProviderId field if non-nil, zero value otherwise.

### GetSourceProviderIdOk

`func (o *BackupStorageProvider) GetSourceProviderIdOk() (*int64, bool)`

GetSourceProviderIdOk returns a tuple with the SourceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProviderId

`func (o *BackupStorageProvider) SetSourceProviderId(v int64)`

SetSourceProviderId sets SourceProviderId field to given value.

### HasSourceProviderId

`func (o *BackupStorageProvider) HasSourceProviderId() bool`

HasSourceProviderId returns a boolean if a field has been set.

### GetStorageProviderId

`func (o *BackupStorageProvider) GetStorageProviderId() int64`

GetStorageProviderId returns the StorageProviderId field if non-nil, zero value otherwise.

### GetStorageProviderIdOk

`func (o *BackupStorageProvider) GetStorageProviderIdOk() (*int64, bool)`

GetStorageProviderIdOk returns a tuple with the StorageProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProviderId

`func (o *BackupStorageProvider) SetStorageProviderId(v int64)`

SetStorageProviderId sets StorageProviderId field to given value.

### HasStorageProviderId

`func (o *BackupStorageProvider) HasStorageProviderId() bool`

HasStorageProviderId returns a boolean if a field has been set.

### GetBackupType

`func (o *BackupStorageProvider) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *BackupStorageProvider) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *BackupStorageProvider) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.


### GetJobAction

`func (o *BackupStorageProvider) GetJobAction() string`

GetJobAction returns the JobAction field if non-nil, zero value otherwise.

### GetJobActionOk

`func (o *BackupStorageProvider) GetJobActionOk() (*string, bool)`

GetJobActionOk returns a tuple with the JobAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobAction

`func (o *BackupStorageProvider) SetJobAction(v string)`

SetJobAction sets JobAction field to given value.


### GetJobId

`func (o *BackupStorageProvider) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BackupStorageProvider) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BackupStorageProvider) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *BackupStorageProvider) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobName

`func (o *BackupStorageProvider) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *BackupStorageProvider) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *BackupStorageProvider) SetJobName(v string)`

SetJobName sets JobName field to given value.

### HasJobName

`func (o *BackupStorageProvider) HasJobName() bool`

HasJobName returns a boolean if a field has been set.

### GetJobSchedule

`func (o *BackupStorageProvider) GetJobSchedule() int64`

GetJobSchedule returns the JobSchedule field if non-nil, zero value otherwise.

### GetJobScheduleOk

`func (o *BackupStorageProvider) GetJobScheduleOk() (*int64, bool)`

GetJobScheduleOk returns a tuple with the JobSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSchedule

`func (o *BackupStorageProvider) SetJobSchedule(v int64)`

SetJobSchedule sets JobSchedule field to given value.

### HasJobSchedule

`func (o *BackupStorageProvider) HasJobSchedule() bool`

HasJobSchedule returns a boolean if a field has been set.

### GetRetentionCount

`func (o *BackupStorageProvider) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *BackupStorageProvider) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *BackupStorageProvider) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *BackupStorageProvider) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


