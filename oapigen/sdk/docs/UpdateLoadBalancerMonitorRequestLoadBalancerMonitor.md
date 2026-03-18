# UpdateLoadBalancerMonitorRequestLoadBalancerMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**MonitorType** | Pointer to **string** |  | [optional] 
**MonitorTimeout** | Pointer to **int64** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by type. | [optional] 

## Methods

### NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitor

`func NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitor() *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor`

NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitor instantiates a new UpdateLoadBalancerMonitorRequestLoadBalancerMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorWithDefaults

`func NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorWithDefaults() *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor`

NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorWithDefaults instantiates a new UpdateLoadBalancerMonitorRequestLoadBalancerMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMonitorType

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetMonitorTimeout

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTimeout() int64`

GetMonitorTimeout returns the MonitorTimeout field if non-nil, zero value otherwise.

### GetMonitorTimeoutOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTimeoutOk() (*int64, bool)`

GetMonitorTimeoutOk returns a tuple with the MonitorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeout

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorTimeout(v int64)`

SetMonitorTimeout sets MonitorTimeout field to given value.

### HasMonitorTimeout

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorTimeout() bool`

HasMonitorTimeout returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


