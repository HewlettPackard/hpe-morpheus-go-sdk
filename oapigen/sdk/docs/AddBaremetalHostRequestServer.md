# AddBaremetalHostRequestServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | Pointer to [**AddBaremetalHostRequestServerCloud**](AddBaremetalHostRequestServerCloud.md) |  | [optional] 
**ComputeServerType** | Pointer to [**AddBaremetalHostRequestServerComputeServerType**](AddBaremetalHostRequestServerComputeServerType.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**AddBaremetalHostRequestServerGroup**](AddBaremetalHostRequestServerGroup.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**AddBaremetalHostRequestServerConfig**](AddBaremetalHostRequestServerConfig.md) |  | [optional] 

## Methods

### NewAddBaremetalHostRequestServer

`func NewAddBaremetalHostRequestServer() *AddBaremetalHostRequestServer`

NewAddBaremetalHostRequestServer instantiates a new AddBaremetalHostRequestServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBaremetalHostRequestServerWithDefaults

`func NewAddBaremetalHostRequestServerWithDefaults() *AddBaremetalHostRequestServer`

NewAddBaremetalHostRequestServerWithDefaults instantiates a new AddBaremetalHostRequestServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *AddBaremetalHostRequestServer) GetCloud() AddBaremetalHostRequestServerCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AddBaremetalHostRequestServer) GetCloudOk() (*AddBaremetalHostRequestServerCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AddBaremetalHostRequestServer) SetCloud(v AddBaremetalHostRequestServerCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *AddBaremetalHostRequestServer) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetComputeServerType

`func (o *AddBaremetalHostRequestServer) GetComputeServerType() AddBaremetalHostRequestServerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *AddBaremetalHostRequestServer) GetComputeServerTypeOk() (*AddBaremetalHostRequestServerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *AddBaremetalHostRequestServer) SetComputeServerType(v AddBaremetalHostRequestServerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *AddBaremetalHostRequestServer) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetDescription

`func (o *AddBaremetalHostRequestServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddBaremetalHostRequestServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddBaremetalHostRequestServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddBaremetalHostRequestServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *AddBaremetalHostRequestServer) GetGroup() AddBaremetalHostRequestServerGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AddBaremetalHostRequestServer) GetGroupOk() (*AddBaremetalHostRequestServerGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AddBaremetalHostRequestServer) SetGroup(v AddBaremetalHostRequestServerGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AddBaremetalHostRequestServer) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLabels

`func (o *AddBaremetalHostRequestServer) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddBaremetalHostRequestServer) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddBaremetalHostRequestServer) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddBaremetalHostRequestServer) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *AddBaremetalHostRequestServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBaremetalHostRequestServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBaremetalHostRequestServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddBaremetalHostRequestServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisibility

`func (o *AddBaremetalHostRequestServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddBaremetalHostRequestServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddBaremetalHostRequestServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddBaremetalHostRequestServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetConfig

`func (o *AddBaremetalHostRequestServer) GetConfig() AddBaremetalHostRequestServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddBaremetalHostRequestServer) GetConfigOk() (*AddBaremetalHostRequestServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddBaremetalHostRequestServer) SetConfig(v AddBaremetalHostRequestServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddBaremetalHostRequestServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


