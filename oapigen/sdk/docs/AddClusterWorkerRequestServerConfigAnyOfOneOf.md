# AddClusterWorkerRequestServerConfigAnyOfOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**DefaultRepoAccount** | Pointer to **NullableInt64** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 

## Methods

### NewAddClusterWorkerRequestServerConfigAnyOfOneOf

`func NewAddClusterWorkerRequestServerConfigAnyOfOneOf() *AddClusterWorkerRequestServerConfigAnyOfOneOf`

NewAddClusterWorkerRequestServerConfigAnyOfOneOf instantiates a new AddClusterWorkerRequestServerConfigAnyOfOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterWorkerRequestServerConfigAnyOfOneOfWithDefaults

`func NewAddClusterWorkerRequestServerConfigAnyOfOneOfWithDefaults() *AddClusterWorkerRequestServerConfigAnyOfOneOf`

NewAddClusterWorkerRequestServerConfigAnyOfOneOfWithDefaults instantiates a new AddClusterWorkerRequestServerConfigAnyOfOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) GetDefaultRepoAccount() int64`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) GetDefaultRepoAccountOk() (*int64, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) SetDefaultRepoAccount(v int64)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


