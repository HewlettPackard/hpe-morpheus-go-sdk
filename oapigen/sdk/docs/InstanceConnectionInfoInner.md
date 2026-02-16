# InstanceConnectionInfoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **NullableInt64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewInstanceConnectionInfoInner

`func NewInstanceConnectionInfoInner() *InstanceConnectionInfoInner`

NewInstanceConnectionInfoInner instantiates a new InstanceConnectionInfoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConnectionInfoInnerWithDefaults

`func NewInstanceConnectionInfoInnerWithDefaults() *InstanceConnectionInfoInner`

NewInstanceConnectionInfoInnerWithDefaults instantiates a new InstanceConnectionInfoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *InstanceConnectionInfoInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InstanceConnectionInfoInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InstanceConnectionInfoInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InstanceConnectionInfoInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPort

`func (o *InstanceConnectionInfoInner) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InstanceConnectionInfoInner) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InstanceConnectionInfoInner) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *InstanceConnectionInfoInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *InstanceConnectionInfoInner) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *InstanceConnectionInfoInner) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetName

`func (o *InstanceConnectionInfoInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceConnectionInfoInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceConnectionInfoInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceConnectionInfoInner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


