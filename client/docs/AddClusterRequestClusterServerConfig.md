# AddClusterRequestClusterServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRepoAccount** | Pointer to **int32** | Default Git Account to be used when pulling images.  Default behavior is to be anonymous, which does have limits on allowed image pulls from public Docker Repos. | [optional] 

## Methods

### NewAddClusterRequestClusterServerConfig

`func NewAddClusterRequestClusterServerConfig() *AddClusterRequestClusterServerConfig`

NewAddClusterRequestClusterServerConfig instantiates a new AddClusterRequestClusterServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterRequestClusterServerConfigWithDefaults

`func NewAddClusterRequestClusterServerConfigWithDefaults() *AddClusterRequestClusterServerConfig`

NewAddClusterRequestClusterServerConfigWithDefaults instantiates a new AddClusterRequestClusterServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRepoAccount

`func (o *AddClusterRequestClusterServerConfig) GetDefaultRepoAccount() int32`

GetDefaultRepoAccount returns the DefaultRepoAccount field if non-nil, zero value otherwise.

### GetDefaultRepoAccountOk

`func (o *AddClusterRequestClusterServerConfig) GetDefaultRepoAccountOk() (*int32, bool)`

GetDefaultRepoAccountOk returns a tuple with the DefaultRepoAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepoAccount

`func (o *AddClusterRequestClusterServerConfig) SetDefaultRepoAccount(v int32)`

SetDefaultRepoAccount sets DefaultRepoAccount field to given value.

### HasDefaultRepoAccount

`func (o *AddClusterRequestClusterServerConfig) HasDefaultRepoAccount() bool`

HasDefaultRepoAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


