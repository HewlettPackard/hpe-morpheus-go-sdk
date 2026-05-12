# UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | Pointer to [**UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigMonitor**](UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigMonitor.md) |  | [optional] 
**MonitorConfig** | Pointer to **NullableString** | Free-form monitor configuration content. | [optional] 

## Methods

### NewUpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig

`func NewUpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig() *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig`

NewUpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig instantiates a new UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigWithDefaults

`func NewUpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigWithDefaults() *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig`

NewUpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigWithDefaults instantiates a new UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitor

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) GetMonitor() UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigMonitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) GetMonitorOk() (*UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigMonitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) SetMonitor(v UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfigMonitor)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetMonitorConfig

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) GetMonitorConfig() string`

GetMonitorConfig returns the MonitorConfig field if non-nil, zero value otherwise.

### GetMonitorConfigOk

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) GetMonitorConfigOk() (*string, bool)`

GetMonitorConfigOk returns a tuple with the MonitorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorConfig

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) SetMonitorConfig(v string)`

SetMonitorConfig sets MonitorConfig field to given value.

### HasMonitorConfig

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) HasMonitorConfig() bool`

HasMonitorConfig returns a boolean if a field has been set.

### SetMonitorConfigNil

`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) SetMonitorConfigNil(b bool)`

 SetMonitorConfigNil sets the value for MonitorConfig to be an explicit nil

### UnsetMonitorConfig
`func (o *UpdateLoadBalancerMonitor200ResponseAllOfLoadBalancerMonitorConfig) UnsetMonitorConfig()`

UnsetMonitorConfig ensures that no value is present for MonitorConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


