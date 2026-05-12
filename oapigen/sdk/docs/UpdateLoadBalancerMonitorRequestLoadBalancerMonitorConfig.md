# UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | Pointer to [**UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfigMonitor**](UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfigMonitor.md) |  | [optional] 
**MonitorConfig** | Pointer to **NullableString** | Free-form monitor configuration content. | [optional] 

## Methods

### NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig

`func NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig() *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig`

NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig instantiates a new UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfigWithDefaults

`func NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfigWithDefaults() *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig`

NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfigWithDefaults instantiates a new UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig) GetMonitor() UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfigMonitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig) GetMonitorOk() (*UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfigMonitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig) SetMonitor(v UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfigMonitor)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetMonitorConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig) GetMonitorConfig() string`

GetMonitorConfig returns the MonitorConfig field if non-nil, zero value otherwise.

### GetMonitorConfigOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig) GetMonitorConfigOk() (*string, bool)`

GetMonitorConfigOk returns a tuple with the MonitorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig) SetMonitorConfig(v string)`

SetMonitorConfig sets MonitorConfig field to given value.

### HasMonitorConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig) HasMonitorConfig() bool`

HasMonitorConfig returns a boolean if a field has been set.

### SetMonitorConfigNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig) SetMonitorConfigNil(b bool)`

 SetMonitorConfigNil sets the value for MonitorConfig to be an explicit nil

### UnsetMonitorConfig
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig) UnsetMonitorConfig()`

UnsetMonitorConfig ensures that no value is present for MonitorConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


