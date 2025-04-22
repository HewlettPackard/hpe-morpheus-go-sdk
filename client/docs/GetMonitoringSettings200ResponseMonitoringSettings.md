# GetMonitoringSettings200ResponseMonitoringSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoManageChecks** | Pointer to **bool** | Auto Create Checks | [optional] 
**AvailabilityTimeFrame** | Pointer to **int32** | Availability Time Frame. The number of days availability should be calculated for. Changes will not take effect until your checks have passed their check interval. | [optional] 
**AvailabilityPrecision** | Pointer to **int32** | Availability Precision. The number of decimal places availability should be displayed in. Can be anywhere between 0 and 5. | [optional] 
**DefaultCheckInterval** | Pointer to **int32** | Default Check Interval. The number of minutes to use as the default interval to use when creating new checks. | [optional] 
**ServiceNow** | Pointer to [**GetMonitoringSettings200ResponseMonitoringSettingsServiceNow**](GetMonitoringSettings200ResponseMonitoringSettingsServiceNow.md) |  | [optional] 

## Methods

### NewGetMonitoringSettings200ResponseMonitoringSettings

`func NewGetMonitoringSettings200ResponseMonitoringSettings() *GetMonitoringSettings200ResponseMonitoringSettings`

NewGetMonitoringSettings200ResponseMonitoringSettings instantiates a new GetMonitoringSettings200ResponseMonitoringSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMonitoringSettings200ResponseMonitoringSettingsWithDefaults

`func NewGetMonitoringSettings200ResponseMonitoringSettingsWithDefaults() *GetMonitoringSettings200ResponseMonitoringSettings`

NewGetMonitoringSettings200ResponseMonitoringSettingsWithDefaults instantiates a new GetMonitoringSettings200ResponseMonitoringSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoManageChecks

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) GetAutoManageChecks() bool`

GetAutoManageChecks returns the AutoManageChecks field if non-nil, zero value otherwise.

### GetAutoManageChecksOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) GetAutoManageChecksOk() (*bool, bool)`

GetAutoManageChecksOk returns a tuple with the AutoManageChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoManageChecks

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) SetAutoManageChecks(v bool)`

SetAutoManageChecks sets AutoManageChecks field to given value.

### HasAutoManageChecks

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) HasAutoManageChecks() bool`

HasAutoManageChecks returns a boolean if a field has been set.

### GetAvailabilityTimeFrame

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) GetAvailabilityTimeFrame() int32`

GetAvailabilityTimeFrame returns the AvailabilityTimeFrame field if non-nil, zero value otherwise.

### GetAvailabilityTimeFrameOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) GetAvailabilityTimeFrameOk() (*int32, bool)`

GetAvailabilityTimeFrameOk returns a tuple with the AvailabilityTimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityTimeFrame

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) SetAvailabilityTimeFrame(v int32)`

SetAvailabilityTimeFrame sets AvailabilityTimeFrame field to given value.

### HasAvailabilityTimeFrame

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) HasAvailabilityTimeFrame() bool`

HasAvailabilityTimeFrame returns a boolean if a field has been set.

### GetAvailabilityPrecision

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) GetAvailabilityPrecision() int32`

GetAvailabilityPrecision returns the AvailabilityPrecision field if non-nil, zero value otherwise.

### GetAvailabilityPrecisionOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) GetAvailabilityPrecisionOk() (*int32, bool)`

GetAvailabilityPrecisionOk returns a tuple with the AvailabilityPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityPrecision

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) SetAvailabilityPrecision(v int32)`

SetAvailabilityPrecision sets AvailabilityPrecision field to given value.

### HasAvailabilityPrecision

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) HasAvailabilityPrecision() bool`

HasAvailabilityPrecision returns a boolean if a field has been set.

### GetDefaultCheckInterval

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) GetDefaultCheckInterval() int32`

GetDefaultCheckInterval returns the DefaultCheckInterval field if non-nil, zero value otherwise.

### GetDefaultCheckIntervalOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) GetDefaultCheckIntervalOk() (*int32, bool)`

GetDefaultCheckIntervalOk returns a tuple with the DefaultCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCheckInterval

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) SetDefaultCheckInterval(v int32)`

SetDefaultCheckInterval sets DefaultCheckInterval field to given value.

### HasDefaultCheckInterval

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) HasDefaultCheckInterval() bool`

HasDefaultCheckInterval returns a boolean if a field has been set.

### GetServiceNow

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) GetServiceNow() GetMonitoringSettings200ResponseMonitoringSettingsServiceNow`

GetServiceNow returns the ServiceNow field if non-nil, zero value otherwise.

### GetServiceNowOk

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) GetServiceNowOk() (*GetMonitoringSettings200ResponseMonitoringSettingsServiceNow, bool)`

GetServiceNowOk returns a tuple with the ServiceNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNow

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) SetServiceNow(v GetMonitoringSettings200ResponseMonitoringSettingsServiceNow)`

SetServiceNow sets ServiceNow field to given value.

### HasServiceNow

`func (o *GetMonitoringSettings200ResponseMonitoringSettings) HasServiceNow() bool`

HasServiceNow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


