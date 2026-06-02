# GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Monitor** | Pointer to [**GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfigMonitor**](GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfigMonitor.md) |  | [optional] 
**MonitorConfig** | Pointer to **NullableString** | Free-form monitor configuration content. | [optional] 

## Methods

### NewGetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig

`func NewGetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig() *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig`

NewGetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig instantiates a new GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetMonitor

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig) GetMonitor() GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfigMonitor`

GetMonitor returns the Monitor field if non-nil, zero value otherwise.

### GetMonitorOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig) GetMonitorOk() (*GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfigMonitor, bool)`

GetMonitorOk returns a tuple with the Monitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitor

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig) SetMonitor(v GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfigMonitor)`

SetMonitor sets Monitor field to given value.

### HasMonitor

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig) HasMonitor() bool`

HasMonitor returns a boolean if a field has been set.

### GetMonitorConfig

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig) GetMonitorConfig() string`

GetMonitorConfig returns the MonitorConfig field if non-nil, zero value otherwise.

### GetMonitorConfigOk

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig) GetMonitorConfigOk() (*string, bool)`

GetMonitorConfigOk returns a tuple with the MonitorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorConfig

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig) SetMonitorConfig(v string)`

SetMonitorConfig sets MonitorConfig field to given value.

### HasMonitorConfig

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig) HasMonitorConfig() bool`

HasMonitorConfig returns a boolean if a field has been set.

### SetMonitorConfigNil

`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig) SetMonitorConfigNil(b bool)`

 SetMonitorConfigNil sets the value for MonitorConfig to be an explicit nil

### UnsetMonitorConfig
`func (o *GetLoadBalancerMonitor200ResponseLoadBalancerMonitorConfig) UnsetMonitorConfig()`

UnsetMonitorConfig ensures that no value is present for MonitorConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


