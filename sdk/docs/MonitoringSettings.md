# MonitoringSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoManageChecks** | Pointer to **bool** | Auto Create Checks | [optional] 
**AvailabilityTimeFrame** | Pointer to **int32** | Availability Time Frame. The number of days availability should be calculated for. Changes will not take effect until your checks have passed their check interval. | [optional] 
**AvailabilityPrecision** | Pointer to **int32** | Availability Precision. The number of decimal places availability should be displayed in. Can be anywhere between 0 and 5. | [optional] 
**DefaultCheckInterval** | Pointer to **int32** | Default Check Interval. The number of minutes to use as the default interval to use when creating new checks. | [optional] 
**ServiceNow** | Pointer to [**GetMonitoringSettings200ResponseMonitoringSettingsServiceNow**](GetMonitoringSettings200ResponseMonitoringSettingsServiceNow.md) |  | [optional] 

## Methods

### NewMonitoringSettings

`func NewMonitoringSettings() *MonitoringSettings`

NewMonitoringSettings instantiates a new MonitoringSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringSettingsWithDefaults

`func NewMonitoringSettingsWithDefaults() *MonitoringSettings`

NewMonitoringSettingsWithDefaults instantiates a new MonitoringSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoManageChecks

`func (o *MonitoringSettings) GetAutoManageChecks() bool`

GetAutoManageChecks returns the AutoManageChecks field if non-nil, zero value otherwise.

### GetAutoManageChecksOk

`func (o *MonitoringSettings) GetAutoManageChecksOk() (*bool, bool)`

GetAutoManageChecksOk returns a tuple with the AutoManageChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoManageChecks

`func (o *MonitoringSettings) SetAutoManageChecks(v bool)`

SetAutoManageChecks sets AutoManageChecks field to given value.

### HasAutoManageChecks

`func (o *MonitoringSettings) HasAutoManageChecks() bool`

HasAutoManageChecks returns a boolean if a field has been set.

### GetAvailabilityTimeFrame

`func (o *MonitoringSettings) GetAvailabilityTimeFrame() int32`

GetAvailabilityTimeFrame returns the AvailabilityTimeFrame field if non-nil, zero value otherwise.

### GetAvailabilityTimeFrameOk

`func (o *MonitoringSettings) GetAvailabilityTimeFrameOk() (*int32, bool)`

GetAvailabilityTimeFrameOk returns a tuple with the AvailabilityTimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityTimeFrame

`func (o *MonitoringSettings) SetAvailabilityTimeFrame(v int32)`

SetAvailabilityTimeFrame sets AvailabilityTimeFrame field to given value.

### HasAvailabilityTimeFrame

`func (o *MonitoringSettings) HasAvailabilityTimeFrame() bool`

HasAvailabilityTimeFrame returns a boolean if a field has been set.

### GetAvailabilityPrecision

`func (o *MonitoringSettings) GetAvailabilityPrecision() int32`

GetAvailabilityPrecision returns the AvailabilityPrecision field if non-nil, zero value otherwise.

### GetAvailabilityPrecisionOk

`func (o *MonitoringSettings) GetAvailabilityPrecisionOk() (*int32, bool)`

GetAvailabilityPrecisionOk returns a tuple with the AvailabilityPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityPrecision

`func (o *MonitoringSettings) SetAvailabilityPrecision(v int32)`

SetAvailabilityPrecision sets AvailabilityPrecision field to given value.

### HasAvailabilityPrecision

`func (o *MonitoringSettings) HasAvailabilityPrecision() bool`

HasAvailabilityPrecision returns a boolean if a field has been set.

### GetDefaultCheckInterval

`func (o *MonitoringSettings) GetDefaultCheckInterval() int32`

GetDefaultCheckInterval returns the DefaultCheckInterval field if non-nil, zero value otherwise.

### GetDefaultCheckIntervalOk

`func (o *MonitoringSettings) GetDefaultCheckIntervalOk() (*int32, bool)`

GetDefaultCheckIntervalOk returns a tuple with the DefaultCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCheckInterval

`func (o *MonitoringSettings) SetDefaultCheckInterval(v int32)`

SetDefaultCheckInterval sets DefaultCheckInterval field to given value.

### HasDefaultCheckInterval

`func (o *MonitoringSettings) HasDefaultCheckInterval() bool`

HasDefaultCheckInterval returns a boolean if a field has been set.

### GetServiceNow

`func (o *MonitoringSettings) GetServiceNow() GetMonitoringSettings200ResponseMonitoringSettingsServiceNow`

GetServiceNow returns the ServiceNow field if non-nil, zero value otherwise.

### GetServiceNowOk

`func (o *MonitoringSettings) GetServiceNowOk() (*GetMonitoringSettings200ResponseMonitoringSettingsServiceNow, bool)`

GetServiceNowOk returns a tuple with the ServiceNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNow

`func (o *MonitoringSettings) SetServiceNow(v GetMonitoringSettings200ResponseMonitoringSettingsServiceNow)`

SetServiceNow sets ServiceNow field to given value.

### HasServiceNow

`func (o *MonitoringSettings) HasServiceNow() bool`

HasServiceNow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


