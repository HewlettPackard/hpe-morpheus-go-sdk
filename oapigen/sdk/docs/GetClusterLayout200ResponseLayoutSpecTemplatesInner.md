# GetClusterLayout200ResponseLayoutSpecTemplatesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GetClusterLayout200ResponseLayoutSpecTemplatesInnerType**](GetClusterLayout200ResponseLayoutSpecTemplatesInnerType.md) |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalType** | Pointer to **NullableString** |  | [optional] 
**DeploymentId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**File** | Pointer to [**GetClusterLayout200ResponseLayoutSpecTemplatesInnerFile**](GetClusterLayout200ResponseLayoutSpecTemplatesInnerFile.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetClusterLayout200ResponseLayoutSpecTemplatesInner

`func NewGetClusterLayout200ResponseLayoutSpecTemplatesInner() *GetClusterLayout200ResponseLayoutSpecTemplatesInner`

NewGetClusterLayout200ResponseLayoutSpecTemplatesInner instantiates a new GetClusterLayout200ResponseLayoutSpecTemplatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetName

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetCode

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetType

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetType() GetClusterLayout200ResponseLayoutSpecTemplatesInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetTypeOk() (*GetClusterLayout200ResponseLayoutSpecTemplatesInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetType(v GetClusterLayout200ResponseLayoutSpecTemplatesInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalType

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetExternalType() string`

GetExternalType returns the ExternalType field if non-nil, zero value otherwise.

### GetExternalTypeOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetExternalTypeOk() (*string, bool)`

GetExternalTypeOk returns a tuple with the ExternalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalType

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetExternalType(v string)`

SetExternalType sets ExternalType field to given value.

### HasExternalType

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasExternalType() bool`

HasExternalType returns a boolean if a field has been set.

### SetExternalTypeNil

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetExternalTypeNil(b bool)`

 SetExternalTypeNil sets the value for ExternalType to be an explicit nil

### UnsetExternalType
`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) UnsetExternalType()`

UnsetExternalType ensures that no value is present for ExternalType, not even an explicit nil
### GetDeploymentId

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### SetDeploymentIdNil

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetDeploymentIdNil(b bool)`

 SetDeploymentIdNil sets the value for DeploymentId to be an explicit nil

### UnsetDeploymentId
`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) UnsetDeploymentId()`

UnsetDeploymentId ensures that no value is present for DeploymentId, not even an explicit nil
### GetStatus

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFile

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetFile() GetClusterLayout200ResponseLayoutSpecTemplatesInnerFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetFileOk() (*GetClusterLayout200ResponseLayoutSpecTemplatesInnerFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetFile(v GetClusterLayout200ResponseLayoutSpecTemplatesInnerFile)`

SetFile sets File field to given value.

### HasFile

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetConfig

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetDateCreated

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetClusterLayout200ResponseLayoutSpecTemplatesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


