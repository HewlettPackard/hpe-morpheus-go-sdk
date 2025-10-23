# UpdateDeploy200ResponseAppDeploy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**Instance** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Deployment** | Pointer to [**ListDeploys200ResponseAllOfAppDeploysInnerDeployment**](ListDeploys200ResponseAllOfAppDeploysInnerDeployment.md) |  | [optional] 
**DeploymentVersionId** | Pointer to **int64** |  | [optional] 
**DeploymentVersion** | Pointer to [**ListDeploys200ResponseAllOfAppDeploysInnerDeploymentVersion**](ListDeploys200ResponseAllOfAppDeploysInnerDeploymentVersion.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DeployDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateDeploy200ResponseAppDeploy

`func NewUpdateDeploy200ResponseAppDeploy() *UpdateDeploy200ResponseAppDeploy`

NewUpdateDeploy200ResponseAppDeploy instantiates a new UpdateDeploy200ResponseAppDeploy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploy200ResponseAppDeployWithDefaults

`func NewUpdateDeploy200ResponseAppDeployWithDefaults() *UpdateDeploy200ResponseAppDeploy`

NewUpdateDeploy200ResponseAppDeployWithDefaults instantiates a new UpdateDeploy200ResponseAppDeploy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateDeploy200ResponseAppDeploy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDeploy200ResponseAppDeploy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateDeploy200ResponseAppDeploy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceId

`func (o *UpdateDeploy200ResponseAppDeploy) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *UpdateDeploy200ResponseAppDeploy) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *UpdateDeploy200ResponseAppDeploy) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstance

`func (o *UpdateDeploy200ResponseAppDeploy) GetInstance() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetInstanceOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *UpdateDeploy200ResponseAppDeploy) SetInstance(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *UpdateDeploy200ResponseAppDeploy) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetDeployment

`func (o *UpdateDeploy200ResponseAppDeploy) GetDeployment() ListDeploys200ResponseAllOfAppDeploysInnerDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetDeploymentOk() (*ListDeploys200ResponseAllOfAppDeploysInnerDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *UpdateDeploy200ResponseAppDeploy) SetDeployment(v ListDeploys200ResponseAllOfAppDeploysInnerDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *UpdateDeploy200ResponseAppDeploy) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetDeploymentVersionId

`func (o *UpdateDeploy200ResponseAppDeploy) GetDeploymentVersionId() int64`

GetDeploymentVersionId returns the DeploymentVersionId field if non-nil, zero value otherwise.

### GetDeploymentVersionIdOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetDeploymentVersionIdOk() (*int64, bool)`

GetDeploymentVersionIdOk returns a tuple with the DeploymentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersionId

`func (o *UpdateDeploy200ResponseAppDeploy) SetDeploymentVersionId(v int64)`

SetDeploymentVersionId sets DeploymentVersionId field to given value.

### HasDeploymentVersionId

`func (o *UpdateDeploy200ResponseAppDeploy) HasDeploymentVersionId() bool`

HasDeploymentVersionId returns a boolean if a field has been set.

### GetDeploymentVersion

`func (o *UpdateDeploy200ResponseAppDeploy) GetDeploymentVersion() ListDeploys200ResponseAllOfAppDeploysInnerDeploymentVersion`

GetDeploymentVersion returns the DeploymentVersion field if non-nil, zero value otherwise.

### GetDeploymentVersionOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetDeploymentVersionOk() (*ListDeploys200ResponseAllOfAppDeploysInnerDeploymentVersion, bool)`

GetDeploymentVersionOk returns a tuple with the DeploymentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersion

`func (o *UpdateDeploy200ResponseAppDeploy) SetDeploymentVersion(v ListDeploys200ResponseAllOfAppDeploysInnerDeploymentVersion)`

SetDeploymentVersion sets DeploymentVersion field to given value.

### HasDeploymentVersion

`func (o *UpdateDeploy200ResponseAppDeploy) HasDeploymentVersion() bool`

HasDeploymentVersion returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateDeploy200ResponseAppDeploy) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateDeploy200ResponseAppDeploy) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateDeploy200ResponseAppDeploy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateDeploy200ResponseAppDeploy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateDeploy200ResponseAppDeploy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateDeploy200ResponseAppDeploy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeployDate

`func (o *UpdateDeploy200ResponseAppDeploy) GetDeployDate() time.Time`

GetDeployDate returns the DeployDate field if non-nil, zero value otherwise.

### GetDeployDateOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetDeployDateOk() (*time.Time, bool)`

GetDeployDateOk returns a tuple with the DeployDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployDate

`func (o *UpdateDeploy200ResponseAppDeploy) SetDeployDate(v time.Time)`

SetDeployDate sets DeployDate field to given value.

### HasDeployDate

`func (o *UpdateDeploy200ResponseAppDeploy) HasDeployDate() bool`

HasDeployDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateDeploy200ResponseAppDeploy) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateDeploy200ResponseAppDeploy) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateDeploy200ResponseAppDeploy) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateDeploy200ResponseAppDeploy) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateDeploy200ResponseAppDeploy) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateDeploy200ResponseAppDeploy) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateDeploy200ResponseAppDeploy) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


