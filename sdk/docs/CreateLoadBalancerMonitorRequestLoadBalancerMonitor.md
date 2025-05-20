# CreateLoadBalancerMonitorRequestLoadBalancerMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**MonitorType** | Pointer to **string** |  | [optional] 
**MonitorTimeout** | Pointer to **int64** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by type. | [optional] 

## Methods

### NewCreateLoadBalancerMonitorRequestLoadBalancerMonitor

`func NewCreateLoadBalancerMonitorRequestLoadBalancerMonitor() *CreateLoadBalancerMonitorRequestLoadBalancerMonitor`

NewCreateLoadBalancerMonitorRequestLoadBalancerMonitor instantiates a new CreateLoadBalancerMonitorRequestLoadBalancerMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerMonitorRequestLoadBalancerMonitorWithDefaults

`func NewCreateLoadBalancerMonitorRequestLoadBalancerMonitorWithDefaults() *CreateLoadBalancerMonitorRequestLoadBalancerMonitor`

NewCreateLoadBalancerMonitorRequestLoadBalancerMonitorWithDefaults instantiates a new CreateLoadBalancerMonitorRequestLoadBalancerMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMonitorType

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetMonitorTimeout

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTimeout() int64`

GetMonitorTimeout returns the MonitorTimeout field if non-nil, zero value otherwise.

### GetMonitorTimeoutOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTimeoutOk() (*int64, bool)`

GetMonitorTimeoutOk returns a tuple with the MonitorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeout

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorTimeout(v int64)`

SetMonitorTimeout sets MonitorTimeout field to given value.

### HasMonitorTimeout

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorTimeout() bool`

HasMonitorTimeout returns a boolean if a field has been set.

### GetConfig

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


