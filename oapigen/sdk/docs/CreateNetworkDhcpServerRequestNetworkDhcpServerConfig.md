# CreateNetworkDhcpServerRequestNetworkDhcpServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeCluster** | Pointer to **string** | Edge Cluster | [optional] 
**PreferredEdgeNode1** | Pointer to **string** | Active Edge Node Options obtained by calling option source with :optionSource &#x3D; nsxtEdgeNodes and networkServerId param | [optional] 
**PreferredEdgeNode2** | Pointer to **string** | Standby Edge Node Options obtained by calling option source with optionSource &#x3D; nsxtEdgeNodes and networkServerId param | [optional] 

## Methods

### NewCreateNetworkDhcpServerRequestNetworkDhcpServerConfig

`func NewCreateNetworkDhcpServerRequestNetworkDhcpServerConfig() *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig`

NewCreateNetworkDhcpServerRequestNetworkDhcpServerConfig instantiates a new CreateNetworkDhcpServerRequestNetworkDhcpServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkDhcpServerRequestNetworkDhcpServerConfigWithDefaults

`func NewCreateNetworkDhcpServerRequestNetworkDhcpServerConfigWithDefaults() *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig`

NewCreateNetworkDhcpServerRequestNetworkDhcpServerConfigWithDefaults instantiates a new CreateNetworkDhcpServerRequestNetworkDhcpServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeCluster

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) GetEdgeCluster() string`

GetEdgeCluster returns the EdgeCluster field if non-nil, zero value otherwise.

### GetEdgeClusterOk

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) GetEdgeClusterOk() (*string, bool)`

GetEdgeClusterOk returns a tuple with the EdgeCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCluster

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) SetEdgeCluster(v string)`

SetEdgeCluster sets EdgeCluster field to given value.

### HasEdgeCluster

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) HasEdgeCluster() bool`

HasEdgeCluster returns a boolean if a field has been set.

### GetPreferredEdgeNode1

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) GetPreferredEdgeNode1() string`

GetPreferredEdgeNode1 returns the PreferredEdgeNode1 field if non-nil, zero value otherwise.

### GetPreferredEdgeNode1Ok

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) GetPreferredEdgeNode1Ok() (*string, bool)`

GetPreferredEdgeNode1Ok returns a tuple with the PreferredEdgeNode1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredEdgeNode1

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) SetPreferredEdgeNode1(v string)`

SetPreferredEdgeNode1 sets PreferredEdgeNode1 field to given value.

### HasPreferredEdgeNode1

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) HasPreferredEdgeNode1() bool`

HasPreferredEdgeNode1 returns a boolean if a field has been set.

### GetPreferredEdgeNode2

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) GetPreferredEdgeNode2() string`

GetPreferredEdgeNode2 returns the PreferredEdgeNode2 field if non-nil, zero value otherwise.

### GetPreferredEdgeNode2Ok

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) GetPreferredEdgeNode2Ok() (*string, bool)`

GetPreferredEdgeNode2Ok returns a tuple with the PreferredEdgeNode2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredEdgeNode2

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) SetPreferredEdgeNode2(v string)`

SetPreferredEdgeNode2 sets PreferredEdgeNode2 field to given value.

### HasPreferredEdgeNode2

`func (o *CreateNetworkDhcpServerRequestNetworkDhcpServerConfig) HasPreferredEdgeNode2() bool`

HasPreferredEdgeNode2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


