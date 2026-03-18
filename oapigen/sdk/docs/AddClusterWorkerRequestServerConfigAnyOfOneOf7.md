# AddClusterWorkerRequestServerConfigAnyOfOneOf7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**PodCidr** | Pointer to **string** |  | [optional] 
**VmwareFolderId** | Pointer to **string** |  | [optional] 
**ServiceCidr** | Pointer to **string** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**DefaultRepoAccount** | Pointer to **NullableInt32** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 

## Methods

### NewAddClusterWorkerRequestServerConfigAnyOfOneOf7

`func NewAddClusterWorkerRequestServerConfigAnyOfOneOf7() *AddClusterWorkerRequestServerConfigAnyOfOneOf7`

NewAddClusterWorkerRequestServerConfigAnyOfOneOf7 instantiates a new AddClusterWorkerRequestServerConfigAnyOfOneOf7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterWorkerRequestServerConfigAnyOfOneOf7WithDefaults

`func NewAddClusterWorkerRequestServerConfigAnyOfOneOf7WithDefaults() *AddClusterWorkerRequestServerConfigAnyOfOneOf7`

NewAddClusterWorkerRequestServerConfigAnyOfOneOf7WithDefaults instantiates a new AddClusterWorkerRequestServerConfigAnyOfOneOf7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetPodCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetServiceCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetDefaultRepoAccount() int32`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetDefaultRepoAccountOk() (*int32, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) SetDefaultRepoAccount(v int32)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf7) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


