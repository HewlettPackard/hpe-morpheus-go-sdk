# ClusterLayoutSpecTemplatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ClusterLayoutSpecTemplatesInnerType**](ClusterLayoutSpecTemplatesInnerType.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalType** | Pointer to **NullableString** |  | [optional] 
**DeploymentId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**File** | Pointer to [**ClusterLayoutSpecTemplatesInnerFile**](ClusterLayoutSpecTemplatesInnerFile.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewClusterLayoutSpecTemplatesInner

`func NewClusterLayoutSpecTemplatesInner() *ClusterLayoutSpecTemplatesInner`

NewClusterLayoutSpecTemplatesInner instantiates a new ClusterLayoutSpecTemplatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutSpecTemplatesInnerWithDefaults

`func NewClusterLayoutSpecTemplatesInnerWithDefaults() *ClusterLayoutSpecTemplatesInner`

NewClusterLayoutSpecTemplatesInnerWithDefaults instantiates a new ClusterLayoutSpecTemplatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterLayoutSpecTemplatesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterLayoutSpecTemplatesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterLayoutSpecTemplatesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterLayoutSpecTemplatesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ClusterLayoutSpecTemplatesInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterLayoutSpecTemplatesInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterLayoutSpecTemplatesInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ClusterLayoutSpecTemplatesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ClusterLayoutSpecTemplatesInner) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ClusterLayoutSpecTemplatesInner) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetName

`func (o *ClusterLayoutSpecTemplatesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLayoutSpecTemplatesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLayoutSpecTemplatesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterLayoutSpecTemplatesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ClusterLayoutSpecTemplatesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterLayoutSpecTemplatesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterLayoutSpecTemplatesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterLayoutSpecTemplatesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ClusterLayoutSpecTemplatesInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ClusterLayoutSpecTemplatesInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetCode

`func (o *ClusterLayoutSpecTemplatesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClusterLayoutSpecTemplatesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClusterLayoutSpecTemplatesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClusterLayoutSpecTemplatesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *ClusterLayoutSpecTemplatesInner) GetType() ClusterLayoutSpecTemplatesInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterLayoutSpecTemplatesInner) GetTypeOk() (*ClusterLayoutSpecTemplatesInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterLayoutSpecTemplatesInner) SetType(v ClusterLayoutSpecTemplatesInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterLayoutSpecTemplatesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *ClusterLayoutSpecTemplatesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ClusterLayoutSpecTemplatesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ClusterLayoutSpecTemplatesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ClusterLayoutSpecTemplatesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ClusterLayoutSpecTemplatesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ClusterLayoutSpecTemplatesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalType

`func (o *ClusterLayoutSpecTemplatesInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *ClusterLayoutSpecTemplatesInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *ClusterLayoutSpecTemplatesInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *ClusterLayoutSpecTemplatesInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *ClusterLayoutSpecTemplatesInner) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *ClusterLayoutSpecTemplatesInner) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetDeploymentId

`func (o *ClusterLayoutSpecTemplatesInner) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *ClusterLayoutSpecTemplatesInner) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *ClusterLayoutSpecTemplatesInner) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *ClusterLayoutSpecTemplatesInner) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *ClusterLayoutSpecTemplatesInner) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *ClusterLayoutSpecTemplatesInner) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetStatus

`func (o *ClusterLayoutSpecTemplatesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterLayoutSpecTemplatesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterLayoutSpecTemplatesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterLayoutSpecTemplatesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ClusterLayoutSpecTemplatesInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ClusterLayoutSpecTemplatesInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFile

`func (o *ClusterLayoutSpecTemplatesInner) GetFile() ClusterLayoutSpecTemplatesInnerFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ClusterLayoutSpecTemplatesInner) GetFileOk() (*ClusterLayoutSpecTemplatesInnerFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ClusterLayoutSpecTemplatesInner) SetFile(v ClusterLayoutSpecTemplatesInnerFile)`

SetFile sets File field to given value.

### HasFile

`func (o *ClusterLayoutSpecTemplatesInner) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetConfig

`func (o *ClusterLayoutSpecTemplatesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterLayoutSpecTemplatesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterLayoutSpecTemplatesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ClusterLayoutSpecTemplatesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ClusterLayoutSpecTemplatesInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ClusterLayoutSpecTemplatesInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ClusterLayoutSpecTemplatesInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ClusterLayoutSpecTemplatesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ClusterLayoutSpecTemplatesInner) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ClusterLayoutSpecTemplatesInner) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ClusterLayoutSpecTemplatesInner) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ClusterLayoutSpecTemplatesInner) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ClusterLayoutSpecTemplatesInner) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ClusterLayoutSpecTemplatesInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ClusterLayoutSpecTemplatesInner) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ClusterLayoutSpecTemplatesInner) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetDateCreated

`func (o *ClusterLayoutSpecTemplatesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ClusterLayoutSpecTemplatesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ClusterLayoutSpecTemplatesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ClusterLayoutSpecTemplatesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ClusterLayoutSpecTemplatesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ClusterLayoutSpecTemplatesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ClusterLayoutSpecTemplatesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ClusterLayoutSpecTemplatesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


