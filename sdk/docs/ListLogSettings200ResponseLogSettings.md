# ListLogSettings200ResponseLogSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**RetentionDays** | Pointer to **string** |  | [optional] 
**SyslogRules** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Integrations** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListLogSettings200ResponseLogSettings

`func NewListLogSettings200ResponseLogSettings() *ListLogSettings200ResponseLogSettings`

NewListLogSettings200ResponseLogSettings instantiates a new ListLogSettings200ResponseLogSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogSettings200ResponseLogSettingsWithDefaults

`func NewListLogSettings200ResponseLogSettingsWithDefaults() *ListLogSettings200ResponseLogSettings`

NewListLogSettings200ResponseLogSettingsWithDefaults instantiates a new ListLogSettings200ResponseLogSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ListLogSettings200ResponseLogSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListLogSettings200ResponseLogSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListLogSettings200ResponseLogSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListLogSettings200ResponseLogSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRetentionDays

`func (o *ListLogSettings200ResponseLogSettings) GetRetentionDays() string`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *ListLogSettings200ResponseLogSettings) GetRetentionDaysOk() (*string, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *ListLogSettings200ResponseLogSettings) SetRetentionDays(v string)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *ListLogSettings200ResponseLogSettings) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetSyslogRules

`func (o *ListLogSettings200ResponseLogSettings) GetSyslogRules() []map[string]interface{}`

GetSyslogRules returns the SyslogRules field if non-nil, zero value otherwise.

### GetSyslogRulesOk

`func (o *ListLogSettings200ResponseLogSettings) GetSyslogRulesOk() (*[]map[string]interface{}, bool)`

GetSyslogRulesOk returns a tuple with the SyslogRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogRules

`func (o *ListLogSettings200ResponseLogSettings) SetSyslogRules(v []map[string]interface{})`

SetSyslogRules sets SyslogRules field to given value.

### HasSyslogRules

`func (o *ListLogSettings200ResponseLogSettings) HasSyslogRules() bool`

HasSyslogRules returns a boolean if a field has been set.

### GetIntegrations

`func (o *ListLogSettings200ResponseLogSettings) GetIntegrations() []map[string]interface{}`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *ListLogSettings200ResponseLogSettings) GetIntegrationsOk() (*[]map[string]interface{}, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *ListLogSettings200ResponseLogSettings) SetIntegrations(v []map[string]interface{})`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *ListLogSettings200ResponseLogSettings) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


