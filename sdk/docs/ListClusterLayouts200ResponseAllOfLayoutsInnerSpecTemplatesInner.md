# ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to **string** |  | [optional] 
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

### NewListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner

`func NewListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner() *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner`

NewListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner instantiates a new ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerWithDefaults

`func NewListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerWithDefaults() *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner`

NewListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerWithDefaults instantiates a new ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### GetDeploymentId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetStatus

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFile

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetFile() ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetFileOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetFile(v ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile)`

SetFile sets File field to given value.

### HasFile

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


