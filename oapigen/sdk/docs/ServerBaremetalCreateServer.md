# ServerBaremetalCreateServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | Pointer to [**ServerBaremetalCreateServerCloud**](ServerBaremetalCreateServerCloud.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ServerBaremetalCreateServerComputeServerType**](ServerBaremetalCreateServerComputeServerType.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**ServerBaremetalCreateServerGroup**](ServerBaremetalCreateServerGroup.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ServerBaremetalCreateServerConfig**](ServerBaremetalCreateServerConfig.md) |  | [optional] 

## Methods

### NewServerBaremetalCreateServer

`func NewServerBaremetalCreateServer() *ServerBaremetalCreateServer`

NewServerBaremetalCreateServer instantiates a new ServerBaremetalCreateServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerBaremetalCreateServerWithDefaults

`func NewServerBaremetalCreateServerWithDefaults() *ServerBaremetalCreateServer`

NewServerBaremetalCreateServerWithDefaults instantiates a new ServerBaremetalCreateServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *ServerBaremetalCreateServer) GetCloud() ServerBaremetalCreateServerCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ServerBaremetalCreateServer) GetCloudOk() (*ServerBaremetalCreateServerCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ServerBaremetalCreateServer) SetCloud(v ServerBaremetalCreateServerCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ServerBaremetalCreateServer) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ServerBaremetalCreateServer) GetComputeServerType() ServerBaremetalCreateServerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ServerBaremetalCreateServer) GetComputeServerTypeOk() (*ServerBaremetalCreateServerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ServerBaremetalCreateServer) SetComputeServerType(v ServerBaremetalCreateServerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ServerBaremetalCreateServer) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetDescription

`func (o *ServerBaremetalCreateServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerBaremetalCreateServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerBaremetalCreateServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerBaremetalCreateServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *ServerBaremetalCreateServer) GetGroup() ServerBaremetalCreateServerGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ServerBaremetalCreateServer) GetGroupOk() (*ServerBaremetalCreateServerGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ServerBaremetalCreateServer) SetGroup(v ServerBaremetalCreateServerGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ServerBaremetalCreateServer) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLabels

`func (o *ServerBaremetalCreateServer) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ServerBaremetalCreateServer) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ServerBaremetalCreateServer) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ServerBaremetalCreateServer) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *ServerBaremetalCreateServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerBaremetalCreateServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerBaremetalCreateServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerBaremetalCreateServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisibility

`func (o *ServerBaremetalCreateServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ServerBaremetalCreateServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ServerBaremetalCreateServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ServerBaremetalCreateServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetConfig

`func (o *ServerBaremetalCreateServer) GetConfig() ServerBaremetalCreateServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServerBaremetalCreateServer) GetConfigOk() (*ServerBaremetalCreateServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServerBaremetalCreateServer) SetConfig(v ServerBaremetalCreateServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServerBaremetalCreateServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


