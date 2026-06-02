# ContainerTypeCreateContainerPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Port** | **int64** |  | 
**LoadBalanceProtocol** | Pointer to **string** |  | [optional] 

## Methods

### NewContainerTypeCreateContainerPortsInner

`func NewContainerTypeCreateContainerPortsInner(name string, port int64, ) *ContainerTypeCreateContainerPortsInner`

NewContainerTypeCreateContainerPortsInner instantiates a new ContainerTypeCreateContainerPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *ContainerTypeCreateContainerPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerTypeCreateContainerPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerTypeCreateContainerPortsInner) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *ContainerTypeCreateContainerPortsInner) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ContainerTypeCreateContainerPortsInner) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ContainerTypeCreateContainerPortsInner) SetPort(v int64)`

SetPort sets Port field to given value.


### GetLoadBalanceProtocol

`func (o *ContainerTypeCreateContainerPortsInner) GetLoadBalanceProtocol() string`

GetLoadBalanceProtocol returns the LoadBalanceProtocol field if non-nil, zero value otherwise.

### GetLoadBalanceProtocolOk

`func (o *ContainerTypeCreateContainerPortsInner) GetLoadBalanceProtocolOk() (*string, bool)`

GetLoadBalanceProtocolOk returns a tuple with the LoadBalanceProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceProtocol

`func (o *ContainerTypeCreateContainerPortsInner) SetLoadBalanceProtocol(v string)`

SetLoadBalanceProtocol sets LoadBalanceProtocol field to given value.

### HasLoadBalanceProtocol

`func (o *ContainerTypeCreateContainerPortsInner) HasLoadBalanceProtocol() bool`

HasLoadBalanceProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


