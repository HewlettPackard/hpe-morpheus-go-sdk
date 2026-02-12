# ClusterServerCreateConfigVsphere

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**PodCidr** | Pointer to **string** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**ServiceCidr** | Pointer to **string** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterServerCreateConfigVsphere

`func NewClusterServerCreateConfigVsphere() *ClusterServerCreateConfigVsphere`

NewClusterServerCreateConfigVsphere instantiates a new ClusterServerCreateConfigVsphere object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigVsphereWithDefaults

`func NewClusterServerCreateConfigVsphereWithDefaults() *ClusterServerCreateConfigVsphere`

NewClusterServerCreateConfigVsphereWithDefaults instantiates a new ClusterServerCreateConfigVsphere object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *ClusterServerCreateConfigVsphere) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreateConfigVsphere) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreateConfigVsphere) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreateConfigVsphere) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetPodCidr

`func (o *ClusterServerCreateConfigVsphere) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *ClusterServerCreateConfigVsphere) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *ClusterServerCreateConfigVsphere) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *ClusterServerCreateConfigVsphere) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *ClusterServerCreateConfigVsphere) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *ClusterServerCreateConfigVsphere) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *ClusterServerCreateConfigVsphere) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *ClusterServerCreateConfigVsphere) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetServiceCidr

`func (o *ClusterServerCreateConfigVsphere) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *ClusterServerCreateConfigVsphere) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *ClusterServerCreateConfigVsphere) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *ClusterServerCreateConfigVsphere) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterServerCreateConfigVsphere) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterServerCreateConfigVsphere) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterServerCreateConfigVsphere) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterServerCreateConfigVsphere) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


