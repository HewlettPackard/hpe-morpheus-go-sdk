# ListSpecTemplates200ResponseAllOfSpecTemplatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalType** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**File** | Pointer to [**ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile**](ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListSpecTemplates200ResponseAllOfSpecTemplatesInner

`func NewListSpecTemplates200ResponseAllOfSpecTemplatesInner() *ListSpecTemplates200ResponseAllOfSpecTemplatesInner`

NewListSpecTemplates200ResponseAllOfSpecTemplatesInner instantiates a new ListSpecTemplates200ResponseAllOfSpecTemplatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSpecTemplates200ResponseAllOfSpecTemplatesInnerWithDefaults

`func NewListSpecTemplates200ResponseAllOfSpecTemplatesInnerWithDefaults() *ListSpecTemplates200ResponseAllOfSpecTemplatesInner`

NewListSpecTemplates200ResponseAllOfSpecTemplatesInnerWithDefaults instantiates a new ListSpecTemplates200ResponseAllOfSpecTemplatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCode

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalType

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetDeploymentId

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetStatus

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFile

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetFile() ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetFileOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetFile(v ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile)`

SetFile sets File field to given value.

### HasFile

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetConfig

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListSpecTemplates200ResponseAllOfSpecTemplatesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


