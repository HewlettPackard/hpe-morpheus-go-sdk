# UpdateLogSettingsRequestLogSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**RetentionDays** | Pointer to **string** |  | [optional] 
**SyslogRules** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Integrations** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateLogSettingsRequestLogSettings

`func NewUpdateLogSettingsRequestLogSettings() *UpdateLogSettingsRequestLogSettings`

NewUpdateLogSettingsRequestLogSettings instantiates a new UpdateLogSettingsRequestLogSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetEnabled

`func (o *UpdateLogSettingsRequestLogSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateLogSettingsRequestLogSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateLogSettingsRequestLogSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateLogSettingsRequestLogSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRetentionDays

`func (o *UpdateLogSettingsRequestLogSettings) GetRetentionDays() string`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *UpdateLogSettingsRequestLogSettings) GetRetentionDaysOk() (*string, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *UpdateLogSettingsRequestLogSettings) SetRetentionDays(v string)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *UpdateLogSettingsRequestLogSettings) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetSyslogRules

`func (o *UpdateLogSettingsRequestLogSettings) GetSyslogRules() []map[string]interface{}`

GetSyslogRules returns the SyslogRules field if non-nil, zero value otherwise.

### GetSyslogRulesOk

`func (o *UpdateLogSettingsRequestLogSettings) GetSyslogRulesOk() (*[]map[string]interface{}, bool)`

GetSyslogRulesOk returns a tuple with the SyslogRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogRules

`func (o *UpdateLogSettingsRequestLogSettings) SetSyslogRules(v []map[string]interface{})`

SetSyslogRules sets SyslogRules field to given value.

### HasSyslogRules

`func (o *UpdateLogSettingsRequestLogSettings) HasSyslogRules() bool`

HasSyslogRules returns a boolean if a field has been set.

### SetSyslogRulesNil

`func (o *UpdateLogSettingsRequestLogSettings) SetSyslogRulesNil(b bool)`

 SetSyslogRulesNil sets the value for SyslogRules to be an explicit nil

### UnsetSyslogRules
`func (o *UpdateLogSettingsRequestLogSettings) UnsetSyslogRules()`

UnsetSyslogRules ensures that no value is present for SyslogRules, not even an explicit nil
### GetIntegrations

`func (o *UpdateLogSettingsRequestLogSettings) GetIntegrations() []map[string]interface{}`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *UpdateLogSettingsRequestLogSettings) GetIntegrationsOk() (*[]map[string]interface{}, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *UpdateLogSettingsRequestLogSettings) SetIntegrations(v []map[string]interface{})`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *UpdateLogSettingsRequestLogSettings) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### SetIntegrationsNil

`func (o *UpdateLogSettingsRequestLogSettings) SetIntegrationsNil(b bool)`

 SetIntegrationsNil sets the value for Integrations to be an explicit nil

### UnsetIntegrations
`func (o *UpdateLogSettingsRequestLogSettings) UnsetIntegrations()`

UnsetIntegrations ensures that no value is present for Integrations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


