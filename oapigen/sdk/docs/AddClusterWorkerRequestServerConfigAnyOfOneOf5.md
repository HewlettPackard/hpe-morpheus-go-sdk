# AddClusterWorkerRequestServerConfigAnyOfOneOf5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodCidr** | Pointer to **string** |  | [optional] 
**ServiceCidr** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**DefaultRepoAccount** | Pointer to **NullableInt32** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 

## Methods

### NewAddClusterWorkerRequestServerConfigAnyOfOneOf5

`func NewAddClusterWorkerRequestServerConfigAnyOfOneOf5() *AddClusterWorkerRequestServerConfigAnyOfOneOf5`

NewAddClusterWorkerRequestServerConfigAnyOfOneOf5 instantiates a new AddClusterWorkerRequestServerConfigAnyOfOneOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterWorkerRequestServerConfigAnyOfOneOf5WithDefaults

`func NewAddClusterWorkerRequestServerConfigAnyOfOneOf5WithDefaults() *AddClusterWorkerRequestServerConfigAnyOfOneOf5`

NewAddClusterWorkerRequestServerConfigAnyOfOneOf5WithDefaults instantiates a new AddClusterWorkerRequestServerConfigAnyOfOneOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetServiceCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetDefaultRepoAccount() int32`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetDefaultRepoAccountOk() (*int32, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) SetDefaultRepoAccount(v int32)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf5) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


