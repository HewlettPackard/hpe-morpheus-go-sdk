# AddClusterRequestClusterServerNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The primary management interface name to establish a management bridge (i.e. eth0,ens192,bond0,etc) | 

## Methods

### NewAddClusterRequestClusterServerNetwork

`func NewAddClusterRequestClusterServerNetwork(name string, ) *AddClusterRequestClusterServerNetwork`

NewAddClusterRequestClusterServerNetwork instantiates a new AddClusterRequestClusterServerNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *AddClusterRequestClusterServerNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddClusterRequestClusterServerNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddClusterRequestClusterServerNetwork) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


