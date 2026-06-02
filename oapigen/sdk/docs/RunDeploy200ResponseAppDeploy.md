# RunDeploy200ResponseAppDeploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**Instance** | Pointer to [**RunDeploy200ResponseAppDeployInstance**](RunDeploy200ResponseAppDeployInstance.md) |  | [optional] 
**Deployment** | Pointer to [**RunDeploy200ResponseAppDeployDeployment**](RunDeploy200ResponseAppDeployDeployment.md) |  | [optional] 
**DeploymentVersionId** | Pointer to **int64** |  | [optional] 
**DeploymentVersion** | Pointer to [**RunDeploy200ResponseAppDeployDeploymentVersion**](RunDeploy200ResponseAppDeployDeploymentVersion.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DeployDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRunDeploy200ResponseAppDeploy

`func NewRunDeploy200ResponseAppDeploy() *RunDeploy200ResponseAppDeploy`

NewRunDeploy200ResponseAppDeploy instantiates a new RunDeploy200ResponseAppDeploy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *RunDeploy200ResponseAppDeploy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunDeploy200ResponseAppDeploy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunDeploy200ResponseAppDeploy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RunDeploy200ResponseAppDeploy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceId

`func (o *RunDeploy200ResponseAppDeploy) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *RunDeploy200ResponseAppDeploy) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *RunDeploy200ResponseAppDeploy) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *RunDeploy200ResponseAppDeploy) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstance

`func (o *RunDeploy200ResponseAppDeploy) GetInstance() RunDeploy200ResponseAppDeployInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *RunDeploy200ResponseAppDeploy) GetInstanceOk() (*RunDeploy200ResponseAppDeployInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *RunDeploy200ResponseAppDeploy) SetInstance(v RunDeploy200ResponseAppDeployInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *RunDeploy200ResponseAppDeploy) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetDeployment

`func (o *RunDeploy200ResponseAppDeploy) GetDeployment() RunDeploy200ResponseAppDeployDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *RunDeploy200ResponseAppDeploy) GetDeploymentOk() (*RunDeploy200ResponseAppDeployDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *RunDeploy200ResponseAppDeploy) SetDeployment(v RunDeploy200ResponseAppDeployDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *RunDeploy200ResponseAppDeploy) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetDeploymentVersionId

`func (o *RunDeploy200ResponseAppDeploy) GetDeploymentVersionId() int64`

GetDeploymentVersionId returns the DeploymentVersionId field if non-nil, zero value otherwise.

### GetDeploymentVersionIdOk

`func (o *RunDeploy200ResponseAppDeploy) GetDeploymentVersionIdOk() (*int64, bool)`

GetDeploymentVersionIdOk returns a tuple with the DeploymentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersionId

`func (o *RunDeploy200ResponseAppDeploy) SetDeploymentVersionId(v int64)`

SetDeploymentVersionId sets DeploymentVersionId field to given value.

### HasDeploymentVersionId

`func (o *RunDeploy200ResponseAppDeploy) HasDeploymentVersionId() bool`

HasDeploymentVersionId returns a boolean if a field has been set.

### GetDeploymentVersion

`func (o *RunDeploy200ResponseAppDeploy) GetDeploymentVersion() RunDeploy200ResponseAppDeployDeploymentVersion`

GetDeploymentVersion returns the DeploymentVersion field if non-nil, zero value otherwise.

### GetDeploymentVersionOk

`func (o *RunDeploy200ResponseAppDeploy) GetDeploymentVersionOk() (*RunDeploy200ResponseAppDeployDeploymentVersion, bool)`

GetDeploymentVersionOk returns a tuple with the DeploymentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersion

`func (o *RunDeploy200ResponseAppDeploy) SetDeploymentVersion(v RunDeploy200ResponseAppDeployDeploymentVersion)`

SetDeploymentVersion sets DeploymentVersion field to given value.

### HasDeploymentVersion

`func (o *RunDeploy200ResponseAppDeploy) HasDeploymentVersion() bool`

HasDeploymentVersion returns a boolean if a field has been set.

### GetConfig

`func (o *RunDeploy200ResponseAppDeploy) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *RunDeploy200ResponseAppDeploy) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *RunDeploy200ResponseAppDeploy) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *RunDeploy200ResponseAppDeploy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *RunDeploy200ResponseAppDeploy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunDeploy200ResponseAppDeploy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunDeploy200ResponseAppDeploy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RunDeploy200ResponseAppDeploy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeployDate

`func (o *RunDeploy200ResponseAppDeploy) GetDeployDate() time.Time`

GetDeployDate returns the DeployDate field if non-nil, zero value otherwise.

### GetDeployDateOk

`func (o *RunDeploy200ResponseAppDeploy) GetDeployDateOk() (*time.Time, bool)`

GetDeployDateOk returns a tuple with the DeployDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployDate

`func (o *RunDeploy200ResponseAppDeploy) SetDeployDate(v time.Time)`

SetDeployDate sets DeployDate field to given value.

### HasDeployDate

`func (o *RunDeploy200ResponseAppDeploy) HasDeployDate() bool`

HasDeployDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *RunDeploy200ResponseAppDeploy) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *RunDeploy200ResponseAppDeploy) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *RunDeploy200ResponseAppDeploy) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *RunDeploy200ResponseAppDeploy) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RunDeploy200ResponseAppDeploy) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RunDeploy200ResponseAppDeploy) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RunDeploy200ResponseAppDeploy) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RunDeploy200ResponseAppDeploy) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


