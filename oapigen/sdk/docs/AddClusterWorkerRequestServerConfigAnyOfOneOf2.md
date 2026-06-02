# AddClusterWorkerRequestServerConfigAnyOfOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to [**AddClusterWorkerRequestServerConfigAnyOfOneOf2Controller**](AddClusterWorkerRequestServerConfigAnyOfOneOf2Controller.md) |  | [optional] 
**Worker** | Pointer to [**AddClusterWorkerRequestServerConfigAnyOfOneOf2Worker**](AddClusterWorkerRequestServerConfigAnyOfOneOf2Worker.md) |  | [optional] 
**PublicIpType** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**DefaultRepoAccount** | Pointer to **NullableInt64** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 

## Methods

### NewAddClusterWorkerRequestServerConfigAnyOfOneOf2

`func NewAddClusterWorkerRequestServerConfigAnyOfOneOf2() *AddClusterWorkerRequestServerConfigAnyOfOneOf2`

NewAddClusterWorkerRequestServerConfigAnyOfOneOf2 instantiates a new AddClusterWorkerRequestServerConfigAnyOfOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetController

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetController() AddClusterWorkerRequestServerConfigAnyOfOneOf2Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetControllerOk() (*AddClusterWorkerRequestServerConfigAnyOfOneOf2Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) SetController(v AddClusterWorkerRequestServerConfigAnyOfOneOf2Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) HasController() bool`

HasController returns a boolean if a field has been set.

### GetWorker

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetWorker() AddClusterWorkerRequestServerConfigAnyOfOneOf2Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetWorkerOk() (*AddClusterWorkerRequestServerConfigAnyOfOneOf2Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) SetWorker(v AddClusterWorkerRequestServerConfigAnyOfOneOf2Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetPublicIpType

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetDefaultRepoAccount() int64`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetDefaultRepoAccountOk() (*int64, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) SetDefaultRepoAccount(v int64)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf2) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


