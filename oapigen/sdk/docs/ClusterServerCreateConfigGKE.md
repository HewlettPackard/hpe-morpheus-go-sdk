# ClusterServerCreateConfigGKE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoogleZoneId** | Pointer to **int64** |  | [optional] 
**Worker** | Pointer to [**AddClusterRequestClusterServerConfigAnyOfOneOf3Worker**](AddClusterRequestClusterServerConfigAnyOfOneOf3Worker.md) |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**ControlPlaneVersion** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterServerCreateConfigGKE

`func NewClusterServerCreateConfigGKE() *ClusterServerCreateConfigGKE`

NewClusterServerCreateConfigGKE instantiates a new ClusterServerCreateConfigGKE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigGKEWithDefaults

`func NewClusterServerCreateConfigGKEWithDefaults() *ClusterServerCreateConfigGKE`

NewClusterServerCreateConfigGKEWithDefaults instantiates a new ClusterServerCreateConfigGKE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoogleZoneId

`func (o *ClusterServerCreateConfigGKE) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *ClusterServerCreateConfigGKE) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *ClusterServerCreateConfigGKE) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *ClusterServerCreateConfigGKE) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetWorker

`func (o *ClusterServerCreateConfigGKE) GetWorker() AddClusterRequestClusterServerConfigAnyOfOneOf3Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ClusterServerCreateConfigGKE) GetWorkerOk() (*AddClusterRequestClusterServerConfigAnyOfOneOf3Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ClusterServerCreateConfigGKE) SetWorker(v AddClusterRequestClusterServerConfigAnyOfOneOf3Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *ClusterServerCreateConfigGKE) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetChannel

`func (o *ClusterServerCreateConfigGKE) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ClusterServerCreateConfigGKE) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ClusterServerCreateConfigGKE) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ClusterServerCreateConfigGKE) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetControlPlaneVersion

`func (o *ClusterServerCreateConfigGKE) GetControlPlaneVersion() string`

GetControlPlaneVersion returns the ControlPlaneVersion field if non-nil, zero value otherwise.

### GetControlPlaneVersionOk

`func (o *ClusterServerCreateConfigGKE) GetControlPlaneVersionOk() (*string, bool)`

GetControlPlaneVersionOk returns a tuple with the ControlPlaneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneVersion

`func (o *ClusterServerCreateConfigGKE) SetControlPlaneVersion(v string)`

SetControlPlaneVersion sets ControlPlaneVersion field to given value.

### HasControlPlaneVersion

`func (o *ClusterServerCreateConfigGKE) HasControlPlaneVersion() bool`

HasControlPlaneVersion returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterServerCreateConfigGKE) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreateConfigGKE) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreateConfigGKE) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreateConfigGKE) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterServerCreateConfigGKE) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterServerCreateConfigGKE) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterServerCreateConfigGKE) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterServerCreateConfigGKE) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


