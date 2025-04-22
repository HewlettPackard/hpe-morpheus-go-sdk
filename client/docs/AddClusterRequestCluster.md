# AddClusterRequestCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AddClusterRequestClusterType**](AddClusterRequestClusterType.md) |  | 
**Name** | **string** | Name of the cluster to be created | 
**Description** | Pointer to **string** | Description of the cluster to be created | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). This will override labels passed under the &#x60;server&#x60; object. | [optional] 
**Group** | [**AddClusterRequestClusterGroup**](AddClusterRequestClusterGroup.md) |  | 
**Cloud** | [**AddClusterRequestClusterCloud**](AddClusterRequestClusterCloud.md) |  | 
**Layout** | [**AddClusterRequestClusterLayout**](AddClusterRequestClusterLayout.md) |  | 
**Server** | [**AddClusterRequestClusterServer**](AddClusterRequestClusterServer.md) |  | 
**AutoRecoverPowerState** | Pointer to **bool** | Automatically Power on VMs | [optional] [default to false]
**TaskSetId** | Pointer to **int64** | Optional Workflow Id desired to be run | [optional] 

## Methods

### NewAddClusterRequestCluster

`func NewAddClusterRequestCluster(type_ AddClusterRequestClusterType, name string, group AddClusterRequestClusterGroup, cloud AddClusterRequestClusterCloud, layout AddClusterRequestClusterLayout, server AddClusterRequestClusterServer, ) *AddClusterRequestCluster`

NewAddClusterRequestCluster instantiates a new AddClusterRequestCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterRequestClusterWithDefaults

`func NewAddClusterRequestClusterWithDefaults() *AddClusterRequestCluster`

NewAddClusterRequestClusterWithDefaults instantiates a new AddClusterRequestCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddClusterRequestCluster) GetType() AddClusterRequestClusterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddClusterRequestCluster) GetTypeOk() (*AddClusterRequestClusterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddClusterRequestCluster) SetType(v AddClusterRequestClusterType)`

SetType sets Type field to given value.


### GetName

`func (o *AddClusterRequestCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddClusterRequestCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddClusterRequestCluster) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddClusterRequestCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddClusterRequestCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddClusterRequestCluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddClusterRequestCluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AddClusterRequestCluster) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddClusterRequestCluster) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddClusterRequestCluster) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddClusterRequestCluster) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetGroup

`func (o *AddClusterRequestCluster) GetGroup() AddClusterRequestClusterGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AddClusterRequestCluster) GetGroupOk() (*AddClusterRequestClusterGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AddClusterRequestCluster) SetGroup(v AddClusterRequestClusterGroup)`

SetGroup sets Group field to given value.


### GetCloud

`func (o *AddClusterRequestCluster) GetCloud() AddClusterRequestClusterCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AddClusterRequestCluster) GetCloudOk() (*AddClusterRequestClusterCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AddClusterRequestCluster) SetCloud(v AddClusterRequestClusterCloud)`

SetCloud sets Cloud field to given value.


### GetLayout

`func (o *AddClusterRequestCluster) GetLayout() AddClusterRequestClusterLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *AddClusterRequestCluster) GetLayoutOk() (*AddClusterRequestClusterLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *AddClusterRequestCluster) SetLayout(v AddClusterRequestClusterLayout)`

SetLayout sets Layout field to given value.


### GetServer

`func (o *AddClusterRequestCluster) GetServer() AddClusterRequestClusterServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddClusterRequestCluster) GetServerOk() (*AddClusterRequestClusterServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddClusterRequestCluster) SetServer(v AddClusterRequestClusterServer)`

SetServer sets Server field to given value.


### GetAutoRecoverPowerState

`func (o *AddClusterRequestCluster) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *AddClusterRequestCluster) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *AddClusterRequestCluster) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *AddClusterRequestCluster) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetTaskSetId

`func (o *AddClusterRequestCluster) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *AddClusterRequestCluster) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *AddClusterRequestCluster) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *AddClusterRequestCluster) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


