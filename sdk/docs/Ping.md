# Ping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**BuildVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewPing

`func NewPing() *Ping`

NewPing instantiates a new Ping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingWithDefaults

`func NewPingWithDefaults() *Ping`

NewPingWithDefaults instantiates a new Ping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *Ping) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Ping) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Ping) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Ping) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBuildVersion

`func (o *Ping) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *Ping) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *Ping) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *Ping) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


