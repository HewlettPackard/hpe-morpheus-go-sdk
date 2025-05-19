# ListDeploys200ResponseAllOfAppDeploysInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**Instance** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Deployment** | Pointer to [**ListDeploys200ResponseAllOfAppDeploysInnerDeployment**](ListDeploys200ResponseAllOfAppDeploysInnerDeployment.md) |  | [optional] 
**DeploymentVersionId** | Pointer to **int64** |  | [optional] 
**DeploymentVersion** | Pointer to [**ListDeploys200ResponseAllOfAppDeploysInnerDeploymentVersion**](ListDeploys200ResponseAllOfAppDeploysInnerDeploymentVersion.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DeployDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListDeploys200ResponseAllOfAppDeploysInner

`func NewListDeploys200ResponseAllOfAppDeploysInner() *ListDeploys200ResponseAllOfAppDeploysInner`

NewListDeploys200ResponseAllOfAppDeploysInner instantiates a new ListDeploys200ResponseAllOfAppDeploysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeploys200ResponseAllOfAppDeploysInnerWithDefaults

`func NewListDeploys200ResponseAllOfAppDeploysInnerWithDefaults() *ListDeploys200ResponseAllOfAppDeploysInner`

NewListDeploys200ResponseAllOfAppDeploysInnerWithDefaults instantiates a new ListDeploys200ResponseAllOfAppDeploysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceId

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstance

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetInstance() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetInstanceOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetInstance(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetDeployment

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetDeployment() ListDeploys200ResponseAllOfAppDeploysInnerDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetDeploymentOk() (*ListDeploys200ResponseAllOfAppDeploysInnerDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetDeployment(v ListDeploys200ResponseAllOfAppDeploysInnerDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetDeploymentVersionId

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetDeploymentVersionId() int64`

GetDeploymentVersionId returns the DeploymentVersionId field if non-nil, zero value otherwise.

### GetDeploymentVersionIdOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetDeploymentVersionIdOk() (*int64, bool)`

GetDeploymentVersionIdOk returns a tuple with the DeploymentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersionId

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetDeploymentVersionId(v int64)`

SetDeploymentVersionId sets DeploymentVersionId field to given value.

### HasDeploymentVersionId

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasDeploymentVersionId() bool`

HasDeploymentVersionId returns a boolean if a field has been set.

### GetDeploymentVersion

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetDeploymentVersion() ListDeploys200ResponseAllOfAppDeploysInnerDeploymentVersion`

GetDeploymentVersion returns the DeploymentVersion field if non-nil, zero value otherwise.

### GetDeploymentVersionOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetDeploymentVersionOk() (*ListDeploys200ResponseAllOfAppDeploysInnerDeploymentVersion, bool)`

GetDeploymentVersionOk returns a tuple with the DeploymentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentVersion

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetDeploymentVersion(v ListDeploys200ResponseAllOfAppDeploysInnerDeploymentVersion)`

SetDeploymentVersion sets DeploymentVersion field to given value.

### HasDeploymentVersion

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasDeploymentVersion() bool`

HasDeploymentVersion returns a boolean if a field has been set.

### GetConfig

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeployDate

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetDeployDate() time.Time`

GetDeployDate returns the DeployDate field if non-nil, zero value otherwise.

### GetDeployDateOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetDeployDateOk() (*time.Time, bool)`

GetDeployDateOk returns a tuple with the DeployDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployDate

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetDeployDate(v time.Time)`

SetDeployDate sets DeployDate field to given value.

### HasDeployDate

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasDeployDate() bool`

HasDeployDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListDeploys200ResponseAllOfAppDeploysInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


