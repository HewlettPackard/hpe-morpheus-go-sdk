# AddInstanceDeploy200ResponseAppDeploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**Instance** | Pointer to [**AddInstanceDeploy200ResponseAppDeployInstance**](AddInstanceDeploy200ResponseAppDeployInstance.md) |  | [optional] 
**Deployment** | Pointer to [**AddInstanceDeploy200ResponseAppDeployDeployment**](AddInstanceDeploy200ResponseAppDeployDeployment.md) |  | [optional] 
**DeploymentVersionId** | Pointer to **int64** |  | [optional] 
**DeploymentVersion** | Pointer to [**AddInstanceDeploy200ResponseAppDeployDeploymentVersion**](AddInstanceDeploy200ResponseAppDeployDeploymentVersion.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DeployDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddInstanceDeploy200ResponseAppDeploy

`func NewAddInstanceDeploy200ResponseAppDeploy() *AddInstanceDeploy200ResponseAppDeploy`

NewAddInstanceDeploy200ResponseAppDeploy instantiates a new AddInstanceDeploy200ResponseAppDeploy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceId

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstance

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetInstance() AddInstanceDeploy200ResponseAppDeployInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetInstanceOk() (*AddInstanceDeploy200ResponseAppDeployInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetInstance(v AddInstanceDeploy200ResponseAppDeployInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetDeployment

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetDeployment() AddInstanceDeploy200ResponseAppDeployDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetDeploymentOk() (*AddInstanceDeploy200ResponseAppDeployDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetDeployment(v AddInstanceDeploy200ResponseAppDeployDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetDeploymentVersionId

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetDeploymentVersionId() int64`

GetDeploymentVersionId returns the DeploymentVersionId field if non-nil, zero value otherwise.

### GetDeploymentVersionIdOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetDeploymentVersionIdOk() (*int64, bool)`

GetDeploymentVersionIdOk returns a tuple with the DeploymentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersionId

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetDeploymentVersionId(v int64)`

SetDeploymentVersionId sets DeploymentVersionId field to given value.

### HasDeploymentVersionId

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasDeploymentVersionId() bool`

HasDeploymentVersionId returns a boolean if a field has been set.

### GetDeploymentVersion

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetDeploymentVersion() AddInstanceDeploy200ResponseAppDeployDeploymentVersion`

GetDeploymentVersion returns the DeploymentVersion field if non-nil, zero value otherwise.

### GetDeploymentVersionOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetDeploymentVersionOk() (*AddInstanceDeploy200ResponseAppDeployDeploymentVersion, bool)`

GetDeploymentVersionOk returns a tuple with the DeploymentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersion

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetDeploymentVersion(v AddInstanceDeploy200ResponseAppDeployDeploymentVersion)`

SetDeploymentVersion sets DeploymentVersion field to given value.

### HasDeploymentVersion

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasDeploymentVersion() bool`

HasDeploymentVersion returns a boolean if a field has been set.

### GetConfig

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeployDate

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetDeployDate() time.Time`

GetDeployDate returns the DeployDate field if non-nil, zero value otherwise.

### GetDeployDateOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetDeployDateOk() (*time.Time, bool)`

GetDeployDateOk returns a tuple with the DeployDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployDate

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetDeployDate(v time.Time)`

SetDeployDate sets DeployDate field to given value.

### HasDeployDate

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasDeployDate() bool`

HasDeployDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddInstanceDeploy200ResponseAppDeploy) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddInstanceDeploy200ResponseAppDeploy) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddInstanceDeploy200ResponseAppDeploy) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


