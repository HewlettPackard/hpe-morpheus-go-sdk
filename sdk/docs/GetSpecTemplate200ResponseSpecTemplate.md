# GetSpecTemplate200ResponseSpecTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalType** | Pointer to **NullableString** |  | [optional] 
**DeploymentId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**File** | Pointer to [**ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile**](ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetSpecTemplate200ResponseSpecTemplate

`func NewGetSpecTemplate200ResponseSpecTemplate() *GetSpecTemplate200ResponseSpecTemplate`

NewGetSpecTemplate200ResponseSpecTemplate instantiates a new GetSpecTemplate200ResponseSpecTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpecTemplate200ResponseSpecTemplateWithDefaults

`func NewGetSpecTemplate200ResponseSpecTemplateWithDefaults() *GetSpecTemplate200ResponseSpecTemplate`

NewGetSpecTemplate200ResponseSpecTemplateWithDefaults instantiates a new GetSpecTemplate200ResponseSpecTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCode

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetSpecTemplate200ResponseSpecTemplate) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetType

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetType sets Type field to given value.

### HasType

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetSpecTemplate200ResponseSpecTemplate) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalType

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *GetSpecTemplate200ResponseSpecTemplate) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetDeploymentId

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *GetSpecTemplate200ResponseSpecTemplate) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetStatus

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetSpecTemplate200ResponseSpecTemplate) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFile

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetFile() ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetFileOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetFile(v ListClusterLayouts200ResponseAllOfLayoutsInnerSpecTemplatesInnerFile)`

SetFile sets File field to given value.

### HasFile

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetConfig

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *GetSpecTemplate200ResponseSpecTemplate) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetDateCreated

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetSpecTemplate200ResponseSpecTemplate) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetSpecTemplate200ResponseSpecTemplate) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetSpecTemplate200ResponseSpecTemplate) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


