# AddInstanceRequestPortsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int64** | Port number. | 
**Name** | Pointer to **string** | A name for the port. | [optional] 
**Lb** | Pointer to **string** | The load balancer protocol. HTTP, HTTPS, or TCP. | [optional] 

## Methods

### NewAddInstanceRequestPortsInner

`func NewAddInstanceRequestPortsInner(port int64, ) *AddInstanceRequestPortsInner`

NewAddInstanceRequestPortsInner instantiates a new AddInstanceRequestPortsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstanceRequestPortsInnerWithDefaults

`func NewAddInstanceRequestPortsInnerWithDefaults() *AddInstanceRequestPortsInner`

NewAddInstanceRequestPortsInnerWithDefaults instantiates a new AddInstanceRequestPortsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *AddInstanceRequestPortsInner) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AddInstanceRequestPortsInner) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AddInstanceRequestPortsInner) SetPort(v int64)`

SetPort sets Port field to given value.


### GetName

`func (o *AddInstanceRequestPortsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddInstanceRequestPortsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddInstanceRequestPortsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddInstanceRequestPortsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLb

`func (o *AddInstanceRequestPortsInner) GetLb() string`

GetLb returns the Lb field if non-nil, zero value otherwise.

### GetLbOk

`func (o *AddInstanceRequestPortsInner) GetLbOk() (*string, bool)`

GetLbOk returns a tuple with the Lb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLb

`func (o *AddInstanceRequestPortsInner) SetLb(v string)`

SetLb sets Lb field to given value.

### HasLb

`func (o *AddInstanceRequestPortsInner) HasLb() bool`

HasLb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


