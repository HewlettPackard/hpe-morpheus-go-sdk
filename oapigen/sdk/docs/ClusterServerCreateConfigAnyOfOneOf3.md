# ClusterServerCreateConfigAnyOfOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GoogleZoneId** | Pointer to **int64** |  | [optional] 
**Worker** | Pointer to [**ClusterServerCreateConfigAnyOfOneOf3Worker**](ClusterServerCreateConfigAnyOfOneOf3Worker.md) |  | [optional] 
**Channel** | Pointer to **string** |  | [optional] 
**ControlPlaneVersion** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**DefaultRepoAccount** | Pointer to **NullableInt32** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 

## Methods

### NewClusterServerCreateConfigAnyOfOneOf3

`func NewClusterServerCreateConfigAnyOfOneOf3() *ClusterServerCreateConfigAnyOfOneOf3`

NewClusterServerCreateConfigAnyOfOneOf3 instantiates a new ClusterServerCreateConfigAnyOfOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigAnyOfOneOf3WithDefaults

`func NewClusterServerCreateConfigAnyOfOneOf3WithDefaults() *ClusterServerCreateConfigAnyOfOneOf3`

NewClusterServerCreateConfigAnyOfOneOf3WithDefaults instantiates a new ClusterServerCreateConfigAnyOfOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGoogleZoneId

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *ClusterServerCreateConfigAnyOfOneOf3) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *ClusterServerCreateConfigAnyOfOneOf3) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetWorker

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetWorker() ClusterServerCreateConfigAnyOfOneOf3Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetWorkerOk() (*ClusterServerCreateConfigAnyOfOneOf3Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ClusterServerCreateConfigAnyOfOneOf3) SetWorker(v ClusterServerCreateConfigAnyOfOneOf3Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *ClusterServerCreateConfigAnyOfOneOf3) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetChannel

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ClusterServerCreateConfigAnyOfOneOf3) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ClusterServerCreateConfigAnyOfOneOf3) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetControlPlaneVersion

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetControlPlaneVersion() string`

GetControlPlaneVersion returns the ControlPlaneVersion field if non-nil, zero value otherwise.

### GetControlPlaneVersionOk

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetControlPlaneVersionOk() (*string, bool)`

GetControlPlaneVersionOk returns a tuple with the ControlPlaneVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlaneVersion

`func (o *ClusterServerCreateConfigAnyOfOneOf3) SetControlPlaneVersion(v string)`

SetControlPlaneVersion sets ControlPlaneVersion field to given value.

### HasControlPlaneVersion

`func (o *ClusterServerCreateConfigAnyOfOneOf3) HasControlPlaneVersion() bool`

HasControlPlaneVersion returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreateConfigAnyOfOneOf3) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreateConfigAnyOfOneOf3) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterServerCreateConfigAnyOfOneOf3) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterServerCreateConfigAnyOfOneOf3) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultRepoAccount

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetDefaultRepoAccount() int32`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetDefaultRepoAccountOk() (*int32, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *ClusterServerCreateConfigAnyOfOneOf3) SetDefaultRepoAccount(v int32)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *ClusterServerCreateConfigAnyOfOneOf3) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *ClusterServerCreateConfigAnyOfOneOf3) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *ClusterServerCreateConfigAnyOfOneOf3) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *ClusterServerCreateConfigAnyOfOneOf3) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *ClusterServerCreateConfigAnyOfOneOf3) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *ClusterServerCreateConfigAnyOfOneOf3) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


