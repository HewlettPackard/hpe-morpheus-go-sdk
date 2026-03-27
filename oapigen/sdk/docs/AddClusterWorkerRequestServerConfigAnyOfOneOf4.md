# AddClusterWorkerRequestServerConfigAnyOfOneOf4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodCidr** | Pointer to **string** |  | [optional] 
**ServiceCidr** | Pointer to **string** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 
**DefaultRepoAccount** | Pointer to **NullableInt64** | Default Repo Account is the repository to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 
**ImageServer** | Pointer to **string** | Act as Image Server. Set to on to use the Default Repo Account to pull images. | [optional] 

## Methods

### NewAddClusterWorkerRequestServerConfigAnyOfOneOf4

`func NewAddClusterWorkerRequestServerConfigAnyOfOneOf4() *AddClusterWorkerRequestServerConfigAnyOfOneOf4`

NewAddClusterWorkerRequestServerConfigAnyOfOneOf4 instantiates a new AddClusterWorkerRequestServerConfigAnyOfOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterWorkerRequestServerConfigAnyOfOneOf4WithDefaults

`func NewAddClusterWorkerRequestServerConfigAnyOfOneOf4WithDefaults() *AddClusterWorkerRequestServerConfigAnyOfOneOf4`

NewAddClusterWorkerRequestServerConfigAnyOfOneOf4WithDefaults instantiates a new AddClusterWorkerRequestServerConfigAnyOfOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetServiceCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) GetDefaultRepoAccount() int64`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) GetDefaultRepoAccountOk() (*int64, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) SetDefaultRepoAccount(v int64)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.

### SetDefaultRepoAccountNil

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) SetDefaultRepoAccountNil(b bool)`

 SetDefaultRepoAccountNil sets the value for DefaultRepoAccount to be an explicit nil

### UnsetDefaultRepoAccount
`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) UnsetDefaultRepoAccount()`

UnsetDefaultRepoAccount ensures that no value is present for DefaultRepoAccount, not even an explicit nil
### GetImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) GetImageServer() string`

GetImageServer returns the ImageServer field if non-nil, zero value otherwise.

### GetImageServerOk

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) GetImageServerOk() (*string, bool)`

GetImageServerOk returns a tuple with the ImageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) SetImageServer(v string)`

SetImageServer sets ImageServer field to given value.

### HasImageServer

`func (o *AddClusterWorkerRequestServerConfigAnyOfOneOf4) HasImageServer() bool`

HasImageServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


