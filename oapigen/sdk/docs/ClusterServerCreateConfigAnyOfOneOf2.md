# ClusterServerCreateConfigAnyOfOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to [**ClusterServerCreateConfigAnyOfOneOf2Controller**](ClusterServerCreateConfigAnyOfOneOf2Controller.md) |  | [optional] 
**Worker** | Pointer to [**ClusterServerCreateConfigAnyOfOneOf2Worker**](ClusterServerCreateConfigAnyOfOneOf2Worker.md) |  | [optional] 
**PublicIpType** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**DefaultRepoAccount** | Pointer to **NullableInt32** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 

## Methods

### NewClusterServerCreateConfigAnyOfOneOf2

`func NewClusterServerCreateConfigAnyOfOneOf2() *ClusterServerCreateConfigAnyOfOneOf2`

NewClusterServerCreateConfigAnyOfOneOf2 instantiates a new ClusterServerCreateConfigAnyOfOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigAnyOfOneOf2WithDefaults

`func NewClusterServerCreateConfigAnyOfOneOf2WithDefaults() *ClusterServerCreateConfigAnyOfOneOf2`

NewClusterServerCreateConfigAnyOfOneOf2WithDefaults instantiates a new ClusterServerCreateConfigAnyOfOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetController

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetController() ClusterServerCreateConfigAnyOfOneOf2Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetControllerOk() (*ClusterServerCreateConfigAnyOfOneOf2Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ClusterServerCreateConfigAnyOfOneOf2) SetController(v ClusterServerCreateConfigAnyOfOneOf2Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *ClusterServerCreateConfigAnyOfOneOf2) HasController() bool`

HasController returns a boolean if a field has been set.

### GetWorker

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetWorker() ClusterServerCreateConfigAnyOfOneOf2Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetWorkerOk() (*ClusterServerCreateConfigAnyOfOneOf2Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ClusterServerCreateConfigAnyOfOneOf2) SetWorker(v ClusterServerCreateConfigAnyOfOneOf2Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *ClusterServerCreateConfigAnyOfOneOf2) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetPublicIpType

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *ClusterServerCreateConfigAnyOfOneOf2) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *ClusterServerCreateConfigAnyOfOneOf2) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreateConfigAnyOfOneOf2) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreateConfigAnyOfOneOf2) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterServerCreateConfigAnyOfOneOf2) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterServerCreateConfigAnyOfOneOf2) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultRepoAccount

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetDefaultRepoAccount() int32`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetDefaultRepoAccountOk() (*int32, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *ClusterServerCreateConfigAnyOfOneOf2) SetDefaultRepoAccount(v int32)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *ClusterServerCreateConfigAnyOfOneOf2) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *ClusterServerCreateConfigAnyOfOneOf2) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *ClusterServerCreateConfigAnyOfOneOf2) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *ClusterServerCreateConfigAnyOfOneOf2) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *ClusterServerCreateConfigAnyOfOneOf2) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *ClusterServerCreateConfigAnyOfOneOf2) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


