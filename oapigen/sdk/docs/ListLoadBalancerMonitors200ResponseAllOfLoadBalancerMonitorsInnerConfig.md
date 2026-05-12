# ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | Pointer to [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfigMonitor**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfigMonitor.md) |  | [optional] 
**MonitorConfig** | Pointer to **NullableString** | Free-form monitor configuration content. | [optional] 

## Methods

### NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig

`func NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig() *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig`

NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig instantiates a new ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfigWithDefaults

`func NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfigWithDefaults() *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig`

NewListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfigWithDefaults instantiates a new ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig) GetMonitor() ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfigMonitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig) GetMonitorOk() (*ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfigMonitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig) SetMonitor(v ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfigMonitor)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetMonitorConfig

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig) GetMonitorConfig() string`

GetMonitorConfig returns the MonitorConfig field if non-nil, zero value otherwise.

### GetMonitorConfigOk

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig) GetMonitorConfigOk() (*string, bool)`

GetMonitorConfigOk returns a tuple with the MonitorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorConfig

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig) SetMonitorConfig(v string)`

SetMonitorConfig sets MonitorConfig field to given value.

### HasMonitorConfig

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig) HasMonitorConfig() bool`

HasMonitorConfig returns a boolean if a field has been set.

### SetMonitorConfigNil

`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig) SetMonitorConfigNil(b bool)`

 SetMonitorConfigNil sets the value for MonitorConfig to be an explicit nil

### UnsetMonitorConfig
`func (o *ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerConfig) UnsetMonitorConfig()`

UnsetMonitorConfig ensures that no value is present for MonitorConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


