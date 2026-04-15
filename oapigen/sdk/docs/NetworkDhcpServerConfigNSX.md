# NetworkDhcpServerConfigNSX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeCluster** | Pointer to **NullableString** | Edge Cluster | [optional] 
**PreferredEdgeNode1** | Pointer to **NullableString** | Active Edge Node Options obtained by calling option source with :optionSource &#x3D; nsxtEdgeNodes and networkServerId param | [optional] 
**PreferredEdgeNode2** | Pointer to **NullableString** | Standby Edge Node Options obtained by calling option source with optionSource &#x3D; nsxtEdgeNodes and networkServerId param | [optional] 

## Methods

### NewNetworkDhcpServerConfigNSX

`func NewNetworkDhcpServerConfigNSX() *NetworkDhcpServerConfigNSX`

NewNetworkDhcpServerConfigNSX instantiates a new NetworkDhcpServerConfigNSX object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDhcpServerConfigNSXWithDefaults

`func NewNetworkDhcpServerConfigNSXWithDefaults() *NetworkDhcpServerConfigNSX`

NewNetworkDhcpServerConfigNSXWithDefaults instantiates a new NetworkDhcpServerConfigNSX object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeCluster

`func (o *NetworkDhcpServerConfigNSX) GetEdgeCluster() string`

GetEdgeCluster returns the EdgeCluster field if non-nil, zero value otherwise.

### GetEdgeClusterOk

`func (o *NetworkDhcpServerConfigNSX) GetEdgeClusterOk() (*string, bool)`

GetEdgeClusterOk returns a tuple with the EdgeCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeCluster

`func (o *NetworkDhcpServerConfigNSX) SetEdgeCluster(v string)`

SetEdgeCluster sets EdgeCluster field to given value.

### HasEdgeCluster

`func (o *NetworkDhcpServerConfigNSX) HasEdgeCluster() bool`

HasEdgeCluster returns a boolean if a field has been set.

### SetEdgeClusterNil

`func (o *NetworkDhcpServerConfigNSX) SetEdgeClusterNil(b bool)`

 SetEdgeClusterNil sets the value for EdgeCluster to be an explicit nil

### UnsetEdgeCluster
`func (o *NetworkDhcpServerConfigNSX) UnsetEdgeCluster()`

UnsetEdgeCluster ensures that no value is present for EdgeCluster, not even an explicit nil
### GetPreferredEdgeNode1

`func (o *NetworkDhcpServerConfigNSX) GetPreferredEdgeNode1() string`

GetPreferredEdgeNode1 returns the PreferredEdgeNode1 field if non-nil, zero value otherwise.

### GetPreferredEdgeNode1Ok

`func (o *NetworkDhcpServerConfigNSX) GetPreferredEdgeNode1Ok() (*string, bool)`

GetPreferredEdgeNode1Ok returns a tuple with the PreferredEdgeNode1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredEdgeNode1

`func (o *NetworkDhcpServerConfigNSX) SetPreferredEdgeNode1(v string)`

SetPreferredEdgeNode1 sets PreferredEdgeNode1 field to given value.

### HasPreferredEdgeNode1

`func (o *NetworkDhcpServerConfigNSX) HasPreferredEdgeNode1() bool`

HasPreferredEdgeNode1 returns a boolean if a field has been set.

### SetPreferredEdgeNode1Nil

`func (o *NetworkDhcpServerConfigNSX) SetPreferredEdgeNode1Nil(b bool)`

 SetPreferredEdgeNode1Nil sets the value for PreferredEdgeNode1 to be an explicit nil

### UnsetPreferredEdgeNode1
`func (o *NetworkDhcpServerConfigNSX) UnsetPreferredEdgeNode1()`

UnsetPreferredEdgeNode1 ensures that no value is present for PreferredEdgeNode1, not even an explicit nil
### GetPreferredEdgeNode2

`func (o *NetworkDhcpServerConfigNSX) GetPreferredEdgeNode2() string`

GetPreferredEdgeNode2 returns the PreferredEdgeNode2 field if non-nil, zero value otherwise.

### GetPreferredEdgeNode2Ok

`func (o *NetworkDhcpServerConfigNSX) GetPreferredEdgeNode2Ok() (*string, bool)`

GetPreferredEdgeNode2Ok returns a tuple with the PreferredEdgeNode2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredEdgeNode2

`func (o *NetworkDhcpServerConfigNSX) SetPreferredEdgeNode2(v string)`

SetPreferredEdgeNode2 sets PreferredEdgeNode2 field to given value.

### HasPreferredEdgeNode2

`func (o *NetworkDhcpServerConfigNSX) HasPreferredEdgeNode2() bool`

HasPreferredEdgeNode2 returns a boolean if a field has been set.

### SetPreferredEdgeNode2Nil

`func (o *NetworkDhcpServerConfigNSX) SetPreferredEdgeNode2Nil(b bool)`

 SetPreferredEdgeNode2Nil sets the value for PreferredEdgeNode2 to be an explicit nil

### UnsetPreferredEdgeNode2
`func (o *NetworkDhcpServerConfigNSX) UnsetPreferredEdgeNode2()`

UnsetPreferredEdgeNode2 ensures that no value is present for PreferredEdgeNode2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


