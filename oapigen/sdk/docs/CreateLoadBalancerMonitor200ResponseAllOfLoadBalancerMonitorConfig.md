# CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | Pointer to [**CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigMonitor**](CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigMonitor.md) |  | [optional] 
**MonitorConfig** | Pointer to **NullableString** | Free-form monitor configuration content. | [optional] 

## Methods

### NewCreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig

`func NewCreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig() *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig`

NewCreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig instantiates a new CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetMonitor

`func (o *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) GetMonitor() CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigMonitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) GetMonitorOk() (*CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigMonitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) SetMonitor(v CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigMonitor)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetMonitorConfig

`func (o *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) GetMonitorConfig() string`

GetMonitorConfig returns the MonitorConfig field if non-nil, zero value otherwise.

### GetMonitorConfigOk

`func (o *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) GetMonitorConfigOk() (*string, bool)`

GetMonitorConfigOk returns a tuple with the MonitorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorConfig

`func (o *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) SetMonitorConfig(v string)`

SetMonitorConfig sets MonitorConfig field to given value.

### HasMonitorConfig

`func (o *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) HasMonitorConfig() bool`

HasMonitorConfig returns a boolean if a field has been set.

### SetMonitorConfigNil

`func (o *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) SetMonitorConfigNil(b bool)`

 SetMonitorConfigNil sets the value for MonitorConfig to be an explicit nil

### UnsetMonitorConfig
`func (o *CreateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) UnsetMonitorConfig()`

UnsetMonitorConfig ensures that no value is present for MonitorConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


