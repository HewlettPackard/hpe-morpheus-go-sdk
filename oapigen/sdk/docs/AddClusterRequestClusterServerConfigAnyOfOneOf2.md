# AddClusterRequestClusterServerConfigAnyOfOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to [**AddClusterRequestClusterServerConfigAnyOfOneOf2Controller**](AddClusterRequestClusterServerConfigAnyOfOneOf2Controller.md) |  | [optional] 
**Worker** | Pointer to [**AddClusterRequestClusterServerConfigAnyOfOneOf2Worker**](AddClusterRequestClusterServerConfigAnyOfOneOf2Worker.md) |  | [optional] 
**PublicIpType** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**DefaultRepoAccount** | Pointer to **NullableInt64** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 

## Methods

### NewAddClusterRequestClusterServerConfigAnyOfOneOf2

`func NewAddClusterRequestClusterServerConfigAnyOfOneOf2() *AddClusterRequestClusterServerConfigAnyOfOneOf2`

NewAddClusterRequestClusterServerConfigAnyOfOneOf2 instantiates a new AddClusterRequestClusterServerConfigAnyOfOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterRequestClusterServerConfigAnyOfOneOf2WithDefaults

`func NewAddClusterRequestClusterServerConfigAnyOfOneOf2WithDefaults() *AddClusterRequestClusterServerConfigAnyOfOneOf2`

NewAddClusterRequestClusterServerConfigAnyOfOneOf2WithDefaults instantiates a new AddClusterRequestClusterServerConfigAnyOfOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetController

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetController() AddClusterRequestClusterServerConfigAnyOfOneOf2Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetControllerOk() (*AddClusterRequestClusterServerConfigAnyOfOneOf2Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) SetController(v AddClusterRequestClusterServerConfigAnyOfOneOf2Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) HasController() bool`

HasController returns a boolean if a field has been set.

### GetWorker

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetWorker() AddClusterRequestClusterServerConfigAnyOfOneOf2Worker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetWorkerOk() (*AddClusterRequestClusterServerConfigAnyOfOneOf2Worker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) SetWorker(v AddClusterRequestClusterServerConfigAnyOfOneOf2Worker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetPublicIpType

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetNodeCount

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultRepoAccount

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetDefaultRepoAccount() int64`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetDefaultRepoAccountOk() (*int64, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) SetDefaultRepoAccount(v int64)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *AddClusterRequestClusterServerConfigAnyOfOneOf2) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


