# ClusterCreateServerConfigAnyOfOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to [**ClusterCreateServerConfigAnyOfOneOf2Controller**](ClusterCreateServerConfigAnyOfOneOf2Controller.md) |  | [optional] 
**Worker** | Pointer to [**ClusterCreateServerConfigAnyOfOneOf2Worker**](ClusterCreateServerConfigAnyOfOneOf2Worker.md) |  | [optional] 
**PublicIpType** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**DefaultRepoAccount** | Pointer to **NullableInt32** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 

## Methods

### NewClusterCreateServerConfigAnyOfOneOf2

`func NewClusterCreateServerConfigAnyOfOneOf2() *ClusterCreateServerConfigAnyOfOneOf2`

NewClusterCreateServerConfigAnyOfOneOf2 instantiates a new ClusterCreateServerConfigAnyOfOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateServerConfigAnyOfOneOf2WithDefaults

`func NewClusterCreateServerConfigAnyOfOneOf2WithDefaults() *ClusterCreateServerConfigAnyOfOneOf2`

NewClusterCreateServerConfigAnyOfOneOf2WithDefaults instantiates a new ClusterCreateServerConfigAnyOfOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetController

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetController() ClusterCreateServerConfigAnyOfOneOf2Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetControllerOk() (*ClusterCreateServerConfigAnyOfOneOf2Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ClusterCreateServerConfigAnyOfOneOf2) SetController(v ClusterCreateServerConfigAnyOfOneOf2Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *ClusterCreateServerConfigAnyOfOneOf2) HasController() bool`

HasController returns a boolean if a field has been set.

### GetWorker

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetWorker() ClusterCreateServerConfigAnyOfOneOf2Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetWorkerOk() (*ClusterCreateServerConfigAnyOfOneOf2Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ClusterCreateServerConfigAnyOfOneOf2) SetWorker(v ClusterCreateServerConfigAnyOfOneOf2Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *ClusterCreateServerConfigAnyOfOneOf2) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetPublicIpType

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *ClusterCreateServerConfigAnyOfOneOf2) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *ClusterCreateServerConfigAnyOfOneOf2) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterCreateServerConfigAnyOfOneOf2) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterCreateServerConfigAnyOfOneOf2) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterCreateServerConfigAnyOfOneOf2) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterCreateServerConfigAnyOfOneOf2) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultRepoAccount

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetDefaultRepoAccount() int32`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetDefaultRepoAccountOk() (*int32, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *ClusterCreateServerConfigAnyOfOneOf2) SetDefaultRepoAccount(v int32)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *ClusterCreateServerConfigAnyOfOneOf2) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *ClusterCreateServerConfigAnyOfOneOf2) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *ClusterCreateServerConfigAnyOfOneOf2) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *ClusterCreateServerConfigAnyOfOneOf2) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *ClusterCreateServerConfigAnyOfOneOf2) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *ClusterCreateServerConfigAnyOfOneOf2) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


