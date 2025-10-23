# AddInstanceDeployRequestAppDeploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **int64** | Deployment ID. | [optional] 
**Version** | Pointer to **string** | Deployment Version number identifier (userVersion). Can be passed along with deploymentId to identify the version | [optional] 
**VersionId** | Pointer to **int64** | Deployment Version ID. This can be passed instead of deploymentId and version. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Map of configuration properties that vary by instance type. | [optional] 
**StageOnly** | Pointer to **bool** | Stage Only, do not run the deploy right away and instead set status to staged so it can be deployed later on. | [optional] [default to false]

## Methods

### NewAddInstanceDeployRequestAppDeploy

`func NewAddInstanceDeployRequestAppDeploy() *AddInstanceDeployRequestAppDeploy`

NewAddInstanceDeployRequestAppDeploy instantiates a new AddInstanceDeployRequestAppDeploy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstanceDeployRequestAppDeployWithDefaults

`func NewAddInstanceDeployRequestAppDeployWithDefaults() *AddInstanceDeployRequestAppDeploy`

NewAddInstanceDeployRequestAppDeployWithDefaults instantiates a new AddInstanceDeployRequestAppDeploy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *AddInstanceDeployRequestAppDeploy) GetDeploymentId() int64`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *AddInstanceDeployRequestAppDeploy) GetDeploymentIdOk() (*int64, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *AddInstanceDeployRequestAppDeploy) SetDeploymentId(v int64)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *AddInstanceDeployRequestAppDeploy) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetVersion

`func (o *AddInstanceDeployRequestAppDeploy) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddInstanceDeployRequestAppDeploy) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddInstanceDeployRequestAppDeploy) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AddInstanceDeployRequestAppDeploy) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionId

`func (o *AddInstanceDeployRequestAppDeploy) GetVersionId() int64`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *AddInstanceDeployRequestAppDeploy) GetVersionIdOk() (*int64, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *AddInstanceDeployRequestAppDeploy) SetVersionId(v int64)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *AddInstanceDeployRequestAppDeploy) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetConfig

`func (o *AddInstanceDeployRequestAppDeploy) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddInstanceDeployRequestAppDeploy) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddInstanceDeployRequestAppDeploy) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddInstanceDeployRequestAppDeploy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStageOnly

`func (o *AddInstanceDeployRequestAppDeploy) GetStageOnly() bool`

GetStageOnly returns the StageOnly field if non-nil, zero value otherwise.

### GetStageOnlyOk

`func (o *AddInstanceDeployRequestAppDeploy) GetStageOnlyOk() (*bool, bool)`

GetStageOnlyOk returns a tuple with the StageOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageOnly

`func (o *AddInstanceDeployRequestAppDeploy) SetStageOnly(v bool)`

SetStageOnly sets StageOnly field to given value.

### HasStageOnly

`func (o *AddInstanceDeployRequestAppDeploy) HasStageOnly() bool`

HasStageOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


