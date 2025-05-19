# UpdateApplianceSettingsRequestApplianceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplianceUrl** | Pointer to **string** | Appliance URL | [optional] 
**InternalApplianceUrl** | Pointer to **string** | Internal Appliance URL (PXE) | [optional] 
**CorsAllowed** | Pointer to **string** | API Allowed Origins | [optional] 
**RegistrationEnabled** | Pointer to **bool** | Registration enabled (true, false) | [optional] 
**DefaultRoleId** | Pointer to **int64** | Default tenant role ID | [optional] 
**DefaultUserRoleId** | Pointer to **int64** | Default user role ID | [optional] 
**DockerPrivilegedMode** | Pointer to **bool** | Docker privileged mode (true, false) | [optional] 
**PasswordMinLength** | Pointer to **string** | Min Password Length | [optional] 
**PasswordMinUpperCase** | Pointer to **string** | Min Password Uppercase | [optional] 
**PasswordMinNumbers** | Pointer to **string** | Min Password Numbers | [optional] 
**PasswordMinSymbols** | Pointer to **string** | Min Password Symbols | [optional] 
**UserBrowserSessionTimeout** | Pointer to **string** | User Browser Session Timeout (Minutes) | [optional] 
**UserBrowserSessionWarning** | Pointer to **string** | User Browser Session Warning (Minutes) | [optional] 
**ExpirePwdDays** | Pointer to **int64** | Expire password after days. Setting to 0 disabled this feature | [optional] 
**DisableAfterAttempts** | Pointer to **int64** | Disable user after number of attempts. Set to 0 to disable this feature | [optional] 
**DisableAfterDaysInactive** | Pointer to **int64** | Disable user if inactive for specified days. Set to 0 to disable this feature | [optional] 
**WarnUserDaysBefore** | Pointer to **int64** | Send warning email number of days in advance before deactivating. Set to 0 to disable this feature | [optional] 
**SmtpMailFrom** | Pointer to **string** | From email address | [optional] 
**SmtpServer** | Pointer to **string** | SMTP server / host | [optional] 
**SmtpPort** | Pointer to **int64** | SMTP port | [optional] 
**SmtpSSL** | Pointer to **bool** | Use SSL for SMTP connection | [optional] 
**SmtpTLS** | Pointer to **bool** | Use TLS for SMTP connections | [optional] 
**SmtpUser** | Pointer to **string** | SMTP username | [optional] 
**SmtpPassword** | Pointer to **string** | SMTP password | [optional] 
**ProxyHost** | Pointer to **string** | Proxy host | [optional] 
**ProxyPort** | Pointer to **string** | Proxy port | [optional] 
**ProxyUser** | Pointer to **string** | Proxy username | [optional] 
**ProxyPassword** | Pointer to **string** | Proxy password | [optional] 
**ProxyDomain** | Pointer to **string** | Proxy domain | [optional] 
**ProxyWorkstation** | Pointer to **string** | Proxy workstation | [optional] 
**CurrencyProvider** | Pointer to **string** | Currency provider | [optional] 
**CurrencyKey** | Pointer to **string** | Currency provider API key | [optional] 
**EnableAllZoneTypes** | Pointer to **bool** | Set all cloud types enabled status on, overrides enableZoneTypes and disableZoneTypes parameters | [optional] 
**EnableZoneTypes** | Pointer to **[]int64** | List of cloud type IDs to set enabled status on | [optional] 
**DisableZoneTypes** | Pointer to **[]int64** | List of cloud type IDs to set enabled status off | [optional] 
**DisableAllZoneTypes** | Pointer to **bool** | Set all cloud types enabled status off, can be used in conjunction with enableZoneTypes | [optional] 
**TwilioAccountSid** | Pointer to **string** | Twilio SMS Account SID | [optional] 
**TwilioSmsFrom** | Pointer to **string** | Twilio SMS From | [optional] 
**TwilioAuthToken** | Pointer to **string** | Twilio SMS Auth Token | [optional] 
**CloudSyncIntervalSeconds** | Pointer to **int64** | Cloud Sync Interval (Seconds) | [optional] 
**ClusterSyncIntervalSeconds** | Pointer to **int64** | Cluster Sync Interval (Seconds) | [optional] 
**UsageRetainmentPeriod** | Pointer to **int64** | Usage Retainment (Days) | [optional] 
**InvoiceRetainmentPeriod** | Pointer to **int64** | Invoice Retainment (Days) | [optional] 
**IncidentRetainmentPeriod** | Pointer to **int64** | Incident Retainment (Days) | [optional] 
**StatsRetainmentPeriod** | Pointer to **int64** | The number of days stats will be retained. (30, 60 or 90) | [optional] 
**ReportsRetainmentPeriod** | Pointer to **int64** | The number of days reports will be retained. | [optional] 
**HttpBlacklistHosts** | Pointer to **string** | Provide a comma delimited list of ips/hostnames to be blocked when using HTTP Task Types or REST Datasource Option Lists | [optional] 
**HttpApprovelistHosts** | Pointer to **string** | Provide a comma delimited list of ips/hostnames to be allowed when using HTTP Task Types or REST Datasource Option Lists. If not specified, only deny list is filtered out. | [optional] 
**NoAgent** | Pointer to **bool** | If true, disables Agent installation globally. | [optional] 
**AgentSSLVerify** | Pointer to **bool** | Enable/Disable SSL Verification of Agent | [optional] 
**DisableSSHPasswordAuth** | Pointer to **bool** | Enable/Disable SSH Password Authentication for the Appliance | [optional] 
**DefaultLocale** | Pointer to **string** | Default appliance Locale. Setting a default locale for the application will override user browser preferences. | [optional] 
**DefaultVdiGateway** | Pointer to **int64** | ID of the VDI gateway. | [optional] 
**MaxOptionListSize** | Pointer to **int64** | Max option list size. Units are x10^3 (thousand). Increasing this value can adversely affect Morpheus performance. Increase with caution. | [optional] 
**ExchangeUrl** | Pointer to **string** | The url used for checking if there is an update for plugins. Default https\\://share.morpheusdata.com | [optional] 

## Methods

### NewUpdateApplianceSettingsRequestApplianceSettings

`func NewUpdateApplianceSettingsRequestApplianceSettings() *UpdateApplianceSettingsRequestApplianceSettings`

NewUpdateApplianceSettingsRequestApplianceSettings instantiates a new UpdateApplianceSettingsRequestApplianceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplianceSettingsRequestApplianceSettingsWithDefaults

`func NewUpdateApplianceSettingsRequestApplianceSettingsWithDefaults() *UpdateApplianceSettingsRequestApplianceSettings`

NewUpdateApplianceSettingsRequestApplianceSettingsWithDefaults instantiates a new UpdateApplianceSettingsRequestApplianceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplianceUrl

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetInternalApplianceUrl

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetInternalApplianceUrl() string`

GetInternalApplianceUrl returns the InternalApplianceUrl field if non-nil, zero value otherwise.

### GetInternalApplianceUrlOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetInternalApplianceUrlOk() (*string, bool)`

GetInternalApplianceUrlOk returns a tuple with the InternalApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalApplianceUrl

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetInternalApplianceUrl(v string)`

SetInternalApplianceUrl sets InternalApplianceUrl field to given value.

### HasInternalApplianceUrl

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasInternalApplianceUrl() bool`

HasInternalApplianceUrl returns a boolean if a field has been set.

### GetCorsAllowed

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetCorsAllowed() string`

GetCorsAllowed returns the CorsAllowed field if non-nil, zero value otherwise.

### GetCorsAllowedOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetCorsAllowedOk() (*string, bool)`

GetCorsAllowedOk returns a tuple with the CorsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowed

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetCorsAllowed(v string)`

SetCorsAllowed sets CorsAllowed field to given value.

### HasCorsAllowed

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasCorsAllowed() bool`

HasCorsAllowed returns a boolean if a field has been set.

### GetRegistrationEnabled

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetRegistrationEnabled() bool`

GetRegistrationEnabled returns the RegistrationEnabled field if non-nil, zero value otherwise.

### GetRegistrationEnabledOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetRegistrationEnabledOk() (*bool, bool)`

GetRegistrationEnabledOk returns a tuple with the RegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEnabled

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetRegistrationEnabled(v bool)`

SetRegistrationEnabled sets RegistrationEnabled field to given value.

### HasRegistrationEnabled

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasRegistrationEnabled() bool`

HasRegistrationEnabled returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDefaultRoleId() int64`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDefaultRoleIdOk() (*int64, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetDefaultRoleId(v int64)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetDefaultUserRoleId

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDefaultUserRoleId() int64`

GetDefaultUserRoleId returns the DefaultUserRoleId field if non-nil, zero value otherwise.

### GetDefaultUserRoleIdOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDefaultUserRoleIdOk() (*int64, bool)`

GetDefaultUserRoleIdOk returns a tuple with the DefaultUserRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserRoleId

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetDefaultUserRoleId(v int64)`

SetDefaultUserRoleId sets DefaultUserRoleId field to given value.

### HasDefaultUserRoleId

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasDefaultUserRoleId() bool`

HasDefaultUserRoleId returns a boolean if a field has been set.

### GetDockerPrivilegedMode

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDockerPrivilegedMode() bool`

GetDockerPrivilegedMode returns the DockerPrivilegedMode field if non-nil, zero value otherwise.

### GetDockerPrivilegedModeOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDockerPrivilegedModeOk() (*bool, bool)`

GetDockerPrivilegedModeOk returns a tuple with the DockerPrivilegedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerPrivilegedMode

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetDockerPrivilegedMode(v bool)`

SetDockerPrivilegedMode sets DockerPrivilegedMode field to given value.

### HasDockerPrivilegedMode

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasDockerPrivilegedMode() bool`

HasDockerPrivilegedMode returns a boolean if a field has been set.

### GetPasswordMinLength

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetPasswordMinLength() string`

GetPasswordMinLength returns the PasswordMinLength field if non-nil, zero value otherwise.

### GetPasswordMinLengthOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetPasswordMinLengthOk() (*string, bool)`

GetPasswordMinLengthOk returns a tuple with the PasswordMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinLength

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetPasswordMinLength(v string)`

SetPasswordMinLength sets PasswordMinLength field to given value.

### HasPasswordMinLength

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasPasswordMinLength() bool`

HasPasswordMinLength returns a boolean if a field has been set.

### GetPasswordMinUpperCase

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetPasswordMinUpperCase() string`

GetPasswordMinUpperCase returns the PasswordMinUpperCase field if non-nil, zero value otherwise.

### GetPasswordMinUpperCaseOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetPasswordMinUpperCaseOk() (*string, bool)`

GetPasswordMinUpperCaseOk returns a tuple with the PasswordMinUpperCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinUpperCase

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetPasswordMinUpperCase(v string)`

SetPasswordMinUpperCase sets PasswordMinUpperCase field to given value.

### HasPasswordMinUpperCase

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasPasswordMinUpperCase() bool`

HasPasswordMinUpperCase returns a boolean if a field has been set.

### GetPasswordMinNumbers

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetPasswordMinNumbers() string`

GetPasswordMinNumbers returns the PasswordMinNumbers field if non-nil, zero value otherwise.

### GetPasswordMinNumbersOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetPasswordMinNumbersOk() (*string, bool)`

GetPasswordMinNumbersOk returns a tuple with the PasswordMinNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinNumbers

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetPasswordMinNumbers(v string)`

SetPasswordMinNumbers sets PasswordMinNumbers field to given value.

### HasPasswordMinNumbers

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasPasswordMinNumbers() bool`

HasPasswordMinNumbers returns a boolean if a field has been set.

### GetPasswordMinSymbols

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetPasswordMinSymbols() string`

GetPasswordMinSymbols returns the PasswordMinSymbols field if non-nil, zero value otherwise.

### GetPasswordMinSymbolsOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetPasswordMinSymbolsOk() (*string, bool)`

GetPasswordMinSymbolsOk returns a tuple with the PasswordMinSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordMinSymbols

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetPasswordMinSymbols(v string)`

SetPasswordMinSymbols sets PasswordMinSymbols field to given value.

### HasPasswordMinSymbols

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasPasswordMinSymbols() bool`

HasPasswordMinSymbols returns a boolean if a field has been set.

### GetUserBrowserSessionTimeout

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetUserBrowserSessionTimeout() string`

GetUserBrowserSessionTimeout returns the UserBrowserSessionTimeout field if non-nil, zero value otherwise.

### GetUserBrowserSessionTimeoutOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetUserBrowserSessionTimeoutOk() (*string, bool)`

GetUserBrowserSessionTimeoutOk returns a tuple with the UserBrowserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBrowserSessionTimeout

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetUserBrowserSessionTimeout(v string)`

SetUserBrowserSessionTimeout sets UserBrowserSessionTimeout field to given value.

### HasUserBrowserSessionTimeout

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasUserBrowserSessionTimeout() bool`

HasUserBrowserSessionTimeout returns a boolean if a field has been set.

### GetUserBrowserSessionWarning

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetUserBrowserSessionWarning() string`

GetUserBrowserSessionWarning returns the UserBrowserSessionWarning field if non-nil, zero value otherwise.

### GetUserBrowserSessionWarningOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetUserBrowserSessionWarningOk() (*string, bool)`

GetUserBrowserSessionWarningOk returns a tuple with the UserBrowserSessionWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBrowserSessionWarning

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetUserBrowserSessionWarning(v string)`

SetUserBrowserSessionWarning sets UserBrowserSessionWarning field to given value.

### HasUserBrowserSessionWarning

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasUserBrowserSessionWarning() bool`

HasUserBrowserSessionWarning returns a boolean if a field has been set.

### GetExpirePwdDays

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetExpirePwdDays() int64`

GetExpirePwdDays returns the ExpirePwdDays field if non-nil, zero value otherwise.

### GetExpirePwdDaysOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetExpirePwdDaysOk() (*int64, bool)`

GetExpirePwdDaysOk returns a tuple with the ExpirePwdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirePwdDays

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetExpirePwdDays(v int64)`

SetExpirePwdDays sets ExpirePwdDays field to given value.

### HasExpirePwdDays

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasExpirePwdDays() bool`

HasExpirePwdDays returns a boolean if a field has been set.

### GetDisableAfterAttempts

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDisableAfterAttempts() int64`

GetDisableAfterAttempts returns the DisableAfterAttempts field if non-nil, zero value otherwise.

### GetDisableAfterAttemptsOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDisableAfterAttemptsOk() (*int64, bool)`

GetDisableAfterAttemptsOk returns a tuple with the DisableAfterAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAfterAttempts

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetDisableAfterAttempts(v int64)`

SetDisableAfterAttempts sets DisableAfterAttempts field to given value.

### HasDisableAfterAttempts

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasDisableAfterAttempts() bool`

HasDisableAfterAttempts returns a boolean if a field has been set.

### GetDisableAfterDaysInactive

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDisableAfterDaysInactive() int64`

GetDisableAfterDaysInactive returns the DisableAfterDaysInactive field if non-nil, zero value otherwise.

### GetDisableAfterDaysInactiveOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDisableAfterDaysInactiveOk() (*int64, bool)`

GetDisableAfterDaysInactiveOk returns a tuple with the DisableAfterDaysInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAfterDaysInactive

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetDisableAfterDaysInactive(v int64)`

SetDisableAfterDaysInactive sets DisableAfterDaysInactive field to given value.

### HasDisableAfterDaysInactive

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasDisableAfterDaysInactive() bool`

HasDisableAfterDaysInactive returns a boolean if a field has been set.

### GetWarnUserDaysBefore

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetWarnUserDaysBefore() int64`

GetWarnUserDaysBefore returns the WarnUserDaysBefore field if non-nil, zero value otherwise.

### GetWarnUserDaysBeforeOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetWarnUserDaysBeforeOk() (*int64, bool)`

GetWarnUserDaysBeforeOk returns a tuple with the WarnUserDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnUserDaysBefore

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetWarnUserDaysBefore(v int64)`

SetWarnUserDaysBefore sets WarnUserDaysBefore field to given value.

### HasWarnUserDaysBefore

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasWarnUserDaysBefore() bool`

HasWarnUserDaysBefore returns a boolean if a field has been set.

### GetSmtpMailFrom

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpMailFrom() string`

GetSmtpMailFrom returns the SmtpMailFrom field if non-nil, zero value otherwise.

### GetSmtpMailFromOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpMailFromOk() (*string, bool)`

GetSmtpMailFromOk returns a tuple with the SmtpMailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpMailFrom

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetSmtpMailFrom(v string)`

SetSmtpMailFrom sets SmtpMailFrom field to given value.

### HasSmtpMailFrom

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasSmtpMailFrom() bool`

HasSmtpMailFrom returns a boolean if a field has been set.

### GetSmtpServer

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetSmtpPort

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpPort() int64`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpPortOk() (*int64, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetSmtpPort(v int64)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpSSL

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpSSL() bool`

GetSmtpSSL returns the SmtpSSL field if non-nil, zero value otherwise.

### GetSmtpSSLOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpSSLOk() (*bool, bool)`

GetSmtpSSLOk returns a tuple with the SmtpSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSSL

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetSmtpSSL(v bool)`

SetSmtpSSL sets SmtpSSL field to given value.

### HasSmtpSSL

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasSmtpSSL() bool`

HasSmtpSSL returns a boolean if a field has been set.

### GetSmtpTLS

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpTLS() bool`

GetSmtpTLS returns the SmtpTLS field if non-nil, zero value otherwise.

### GetSmtpTLSOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpTLSOk() (*bool, bool)`

GetSmtpTLSOk returns a tuple with the SmtpTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpTLS

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetSmtpTLS(v bool)`

SetSmtpTLS sets SmtpTLS field to given value.

### HasSmtpTLS

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasSmtpTLS() bool`

HasSmtpTLS returns a boolean if a field has been set.

### GetSmtpUser

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpUser() string`

GetSmtpUser returns the SmtpUser field if non-nil, zero value otherwise.

### GetSmtpUserOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpUserOk() (*string, bool)`

GetSmtpUserOk returns a tuple with the SmtpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpUser

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetSmtpUser(v string)`

SetSmtpUser sets SmtpUser field to given value.

### HasSmtpUser

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasSmtpUser() bool`

HasSmtpUser returns a boolean if a field has been set.

### GetSmtpPassword

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpPassword() string`

GetSmtpPassword returns the SmtpPassword field if non-nil, zero value otherwise.

### GetSmtpPasswordOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetSmtpPasswordOk() (*string, bool)`

GetSmtpPasswordOk returns a tuple with the SmtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPassword

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetSmtpPassword(v string)`

SetSmtpPassword sets SmtpPassword field to given value.

### HasSmtpPassword

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasSmtpPassword() bool`

HasSmtpPassword returns a boolean if a field has been set.

### GetProxyHost

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyHost() string`

GetProxyHost returns the ProxyHost field if non-nil, zero value otherwise.

### GetProxyHostOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyHostOk() (*string, bool)`

GetProxyHostOk returns a tuple with the ProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHost

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetProxyHost(v string)`

SetProxyHost sets ProxyHost field to given value.

### HasProxyHost

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasProxyHost() bool`

HasProxyHost returns a boolean if a field has been set.

### GetProxyPort

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyPort() string`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyPortOk() (*string, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetProxyPort(v string)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyUser

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### GetProxyPassword

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetProxyDomain

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyDomain() string`

GetProxyDomain returns the ProxyDomain field if non-nil, zero value otherwise.

### GetProxyDomainOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyDomainOk() (*string, bool)`

GetProxyDomainOk returns a tuple with the ProxyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyDomain

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetProxyDomain(v string)`

SetProxyDomain sets ProxyDomain field to given value.

### HasProxyDomain

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasProxyDomain() bool`

HasProxyDomain returns a boolean if a field has been set.

### GetProxyWorkstation

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyWorkstation() string`

GetProxyWorkstation returns the ProxyWorkstation field if non-nil, zero value otherwise.

### GetProxyWorkstationOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetProxyWorkstationOk() (*string, bool)`

GetProxyWorkstationOk returns a tuple with the ProxyWorkstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyWorkstation

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetProxyWorkstation(v string)`

SetProxyWorkstation sets ProxyWorkstation field to given value.

### HasProxyWorkstation

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasProxyWorkstation() bool`

HasProxyWorkstation returns a boolean if a field has been set.

### GetCurrencyProvider

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetCurrencyProvider() string`

GetCurrencyProvider returns the CurrencyProvider field if non-nil, zero value otherwise.

### GetCurrencyProviderOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetCurrencyProviderOk() (*string, bool)`

GetCurrencyProviderOk returns a tuple with the CurrencyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyProvider

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetCurrencyProvider(v string)`

SetCurrencyProvider sets CurrencyProvider field to given value.

### HasCurrencyProvider

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasCurrencyProvider() bool`

HasCurrencyProvider returns a boolean if a field has been set.

### GetCurrencyKey

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetCurrencyKey() string`

GetCurrencyKey returns the CurrencyKey field if non-nil, zero value otherwise.

### GetCurrencyKeyOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetCurrencyKeyOk() (*string, bool)`

GetCurrencyKeyOk returns a tuple with the CurrencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyKey

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetCurrencyKey(v string)`

SetCurrencyKey sets CurrencyKey field to given value.

### HasCurrencyKey

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasCurrencyKey() bool`

HasCurrencyKey returns a boolean if a field has been set.

### GetEnableAllZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetEnableAllZoneTypes() bool`

GetEnableAllZoneTypes returns the EnableAllZoneTypes field if non-nil, zero value otherwise.

### GetEnableAllZoneTypesOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetEnableAllZoneTypesOk() (*bool, bool)`

GetEnableAllZoneTypesOk returns a tuple with the EnableAllZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetEnableAllZoneTypes(v bool)`

SetEnableAllZoneTypes sets EnableAllZoneTypes field to given value.

### HasEnableAllZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasEnableAllZoneTypes() bool`

HasEnableAllZoneTypes returns a boolean if a field has been set.

### GetEnableZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetEnableZoneTypes() []int64`

GetEnableZoneTypes returns the EnableZoneTypes field if non-nil, zero value otherwise.

### GetEnableZoneTypesOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetEnableZoneTypesOk() (*[]int64, bool)`

GetEnableZoneTypesOk returns a tuple with the EnableZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetEnableZoneTypes(v []int64)`

SetEnableZoneTypes sets EnableZoneTypes field to given value.

### HasEnableZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasEnableZoneTypes() bool`

HasEnableZoneTypes returns a boolean if a field has been set.

### GetDisableZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDisableZoneTypes() []int64`

GetDisableZoneTypes returns the DisableZoneTypes field if non-nil, zero value otherwise.

### GetDisableZoneTypesOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDisableZoneTypesOk() (*[]int64, bool)`

GetDisableZoneTypesOk returns a tuple with the DisableZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetDisableZoneTypes(v []int64)`

SetDisableZoneTypes sets DisableZoneTypes field to given value.

### HasDisableZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasDisableZoneTypes() bool`

HasDisableZoneTypes returns a boolean if a field has been set.

### GetDisableAllZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDisableAllZoneTypes() bool`

GetDisableAllZoneTypes returns the DisableAllZoneTypes field if non-nil, zero value otherwise.

### GetDisableAllZoneTypesOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDisableAllZoneTypesOk() (*bool, bool)`

GetDisableAllZoneTypesOk returns a tuple with the DisableAllZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAllZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetDisableAllZoneTypes(v bool)`

SetDisableAllZoneTypes sets DisableAllZoneTypes field to given value.

### HasDisableAllZoneTypes

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasDisableAllZoneTypes() bool`

HasDisableAllZoneTypes returns a boolean if a field has been set.

### GetTwilioAccountSid

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetTwilioAccountSid() string`

GetTwilioAccountSid returns the TwilioAccountSid field if non-nil, zero value otherwise.

### GetTwilioAccountSidOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetTwilioAccountSidOk() (*string, bool)`

GetTwilioAccountSidOk returns a tuple with the TwilioAccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAccountSid

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetTwilioAccountSid(v string)`

SetTwilioAccountSid sets TwilioAccountSid field to given value.

### HasTwilioAccountSid

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasTwilioAccountSid() bool`

HasTwilioAccountSid returns a boolean if a field has been set.

### GetTwilioSmsFrom

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetTwilioSmsFrom() string`

GetTwilioSmsFrom returns the TwilioSmsFrom field if non-nil, zero value otherwise.

### GetTwilioSmsFromOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetTwilioSmsFromOk() (*string, bool)`

GetTwilioSmsFromOk returns a tuple with the TwilioSmsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioSmsFrom

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetTwilioSmsFrom(v string)`

SetTwilioSmsFrom sets TwilioSmsFrom field to given value.

### HasTwilioSmsFrom

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasTwilioSmsFrom() bool`

HasTwilioSmsFrom returns a boolean if a field has been set.

### GetTwilioAuthToken

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetTwilioAuthToken() string`

GetTwilioAuthToken returns the TwilioAuthToken field if non-nil, zero value otherwise.

### GetTwilioAuthTokenOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetTwilioAuthTokenOk() (*string, bool)`

GetTwilioAuthTokenOk returns a tuple with the TwilioAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioAuthToken

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetTwilioAuthToken(v string)`

SetTwilioAuthToken sets TwilioAuthToken field to given value.

### HasTwilioAuthToken

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasTwilioAuthToken() bool`

HasTwilioAuthToken returns a boolean if a field has been set.

### GetCloudSyncIntervalSeconds

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetCloudSyncIntervalSeconds() int64`

GetCloudSyncIntervalSeconds returns the CloudSyncIntervalSeconds field if non-nil, zero value otherwise.

### GetCloudSyncIntervalSecondsOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetCloudSyncIntervalSecondsOk() (*int64, bool)`

GetCloudSyncIntervalSecondsOk returns a tuple with the CloudSyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSyncIntervalSeconds

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetCloudSyncIntervalSeconds(v int64)`

SetCloudSyncIntervalSeconds sets CloudSyncIntervalSeconds field to given value.

### HasCloudSyncIntervalSeconds

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasCloudSyncIntervalSeconds() bool`

HasCloudSyncIntervalSeconds returns a boolean if a field has been set.

### GetClusterSyncIntervalSeconds

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetClusterSyncIntervalSeconds() int64`

GetClusterSyncIntervalSeconds returns the ClusterSyncIntervalSeconds field if non-nil, zero value otherwise.

### GetClusterSyncIntervalSecondsOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetClusterSyncIntervalSecondsOk() (*int64, bool)`

GetClusterSyncIntervalSecondsOk returns a tuple with the ClusterSyncIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSyncIntervalSeconds

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetClusterSyncIntervalSeconds(v int64)`

SetClusterSyncIntervalSeconds sets ClusterSyncIntervalSeconds field to given value.

### HasClusterSyncIntervalSeconds

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasClusterSyncIntervalSeconds() bool`

HasClusterSyncIntervalSeconds returns a boolean if a field has been set.

### GetUsageRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetUsageRetainmentPeriod() int64`

GetUsageRetainmentPeriod returns the UsageRetainmentPeriod field if non-nil, zero value otherwise.

### GetUsageRetainmentPeriodOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetUsageRetainmentPeriodOk() (*int64, bool)`

GetUsageRetainmentPeriodOk returns a tuple with the UsageRetainmentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetUsageRetainmentPeriod(v int64)`

SetUsageRetainmentPeriod sets UsageRetainmentPeriod field to given value.

### HasUsageRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasUsageRetainmentPeriod() bool`

HasUsageRetainmentPeriod returns a boolean if a field has been set.

### GetInvoiceRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetInvoiceRetainmentPeriod() int64`

GetInvoiceRetainmentPeriod returns the InvoiceRetainmentPeriod field if non-nil, zero value otherwise.

### GetInvoiceRetainmentPeriodOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetInvoiceRetainmentPeriodOk() (*int64, bool)`

GetInvoiceRetainmentPeriodOk returns a tuple with the InvoiceRetainmentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetInvoiceRetainmentPeriod(v int64)`

SetInvoiceRetainmentPeriod sets InvoiceRetainmentPeriod field to given value.

### HasInvoiceRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasInvoiceRetainmentPeriod() bool`

HasInvoiceRetainmentPeriod returns a boolean if a field has been set.

### GetIncidentRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetIncidentRetainmentPeriod() int64`

GetIncidentRetainmentPeriod returns the IncidentRetainmentPeriod field if non-nil, zero value otherwise.

### GetIncidentRetainmentPeriodOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetIncidentRetainmentPeriodOk() (*int64, bool)`

GetIncidentRetainmentPeriodOk returns a tuple with the IncidentRetainmentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetIncidentRetainmentPeriod(v int64)`

SetIncidentRetainmentPeriod sets IncidentRetainmentPeriod field to given value.

### HasIncidentRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasIncidentRetainmentPeriod() bool`

HasIncidentRetainmentPeriod returns a boolean if a field has been set.

### GetStatsRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetStatsRetainmentPeriod() int64`

GetStatsRetainmentPeriod returns the StatsRetainmentPeriod field if non-nil, zero value otherwise.

### GetStatsRetainmentPeriodOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetStatsRetainmentPeriodOk() (*int64, bool)`

GetStatsRetainmentPeriodOk returns a tuple with the StatsRetainmentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetStatsRetainmentPeriod(v int64)`

SetStatsRetainmentPeriod sets StatsRetainmentPeriod field to given value.

### HasStatsRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasStatsRetainmentPeriod() bool`

HasStatsRetainmentPeriod returns a boolean if a field has been set.

### GetReportsRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetReportsRetainmentPeriod() int64`

GetReportsRetainmentPeriod returns the ReportsRetainmentPeriod field if non-nil, zero value otherwise.

### GetReportsRetainmentPeriodOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetReportsRetainmentPeriodOk() (*int64, bool)`

GetReportsRetainmentPeriodOk returns a tuple with the ReportsRetainmentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetReportsRetainmentPeriod(v int64)`

SetReportsRetainmentPeriod sets ReportsRetainmentPeriod field to given value.

### HasReportsRetainmentPeriod

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasReportsRetainmentPeriod() bool`

HasReportsRetainmentPeriod returns a boolean if a field has been set.

### GetHttpBlacklistHosts

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetHttpBlacklistHosts() string`

GetHttpBlacklistHosts returns the HttpBlacklistHosts field if non-nil, zero value otherwise.

### GetHttpBlacklistHostsOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetHttpBlacklistHostsOk() (*string, bool)`

GetHttpBlacklistHostsOk returns a tuple with the HttpBlacklistHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpBlacklistHosts

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetHttpBlacklistHosts(v string)`

SetHttpBlacklistHosts sets HttpBlacklistHosts field to given value.

### HasHttpBlacklistHosts

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasHttpBlacklistHosts() bool`

HasHttpBlacklistHosts returns a boolean if a field has been set.

### GetHttpApprovelistHosts

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetHttpApprovelistHosts() string`

GetHttpApprovelistHosts returns the HttpApprovelistHosts field if non-nil, zero value otherwise.

### GetHttpApprovelistHostsOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetHttpApprovelistHostsOk() (*string, bool)`

GetHttpApprovelistHostsOk returns a tuple with the HttpApprovelistHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpApprovelistHosts

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetHttpApprovelistHosts(v string)`

SetHttpApprovelistHosts sets HttpApprovelistHosts field to given value.

### HasHttpApprovelistHosts

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasHttpApprovelistHosts() bool`

HasHttpApprovelistHosts returns a boolean if a field has been set.

### GetNoAgent

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetAgentSSLVerify

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetAgentSSLVerify() bool`

GetAgentSSLVerify returns the AgentSSLVerify field if non-nil, zero value otherwise.

### GetAgentSSLVerifyOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetAgentSSLVerifyOk() (*bool, bool)`

GetAgentSSLVerifyOk returns a tuple with the AgentSSLVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSSLVerify

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetAgentSSLVerify(v bool)`

SetAgentSSLVerify sets AgentSSLVerify field to given value.

### HasAgentSSLVerify

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasAgentSSLVerify() bool`

HasAgentSSLVerify returns a boolean if a field has been set.

### GetDisableSSHPasswordAuth

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDisableSSHPasswordAuth() bool`

GetDisableSSHPasswordAuth returns the DisableSSHPasswordAuth field if non-nil, zero value otherwise.

### GetDisableSSHPasswordAuthOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDisableSSHPasswordAuthOk() (*bool, bool)`

GetDisableSSHPasswordAuthOk returns a tuple with the DisableSSHPasswordAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSSHPasswordAuth

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetDisableSSHPasswordAuth(v bool)`

SetDisableSSHPasswordAuth sets DisableSSHPasswordAuth field to given value.

### HasDisableSSHPasswordAuth

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasDisableSSHPasswordAuth() bool`

HasDisableSSHPasswordAuth returns a boolean if a field has been set.

### GetDefaultLocale

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDefaultLocale() string`

GetDefaultLocale returns the DefaultLocale field if non-nil, zero value otherwise.

### GetDefaultLocaleOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDefaultLocaleOk() (*string, bool)`

GetDefaultLocaleOk returns a tuple with the DefaultLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocale

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetDefaultLocale(v string)`

SetDefaultLocale sets DefaultLocale field to given value.

### HasDefaultLocale

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasDefaultLocale() bool`

HasDefaultLocale returns a boolean if a field has been set.

### GetDefaultVdiGateway

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDefaultVdiGateway() int64`

GetDefaultVdiGateway returns the DefaultVdiGateway field if non-nil, zero value otherwise.

### GetDefaultVdiGatewayOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetDefaultVdiGatewayOk() (*int64, bool)`

GetDefaultVdiGatewayOk returns a tuple with the DefaultVdiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVdiGateway

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetDefaultVdiGateway(v int64)`

SetDefaultVdiGateway sets DefaultVdiGateway field to given value.

### HasDefaultVdiGateway

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasDefaultVdiGateway() bool`

HasDefaultVdiGateway returns a boolean if a field has been set.

### GetMaxOptionListSize

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetMaxOptionListSize() int64`

GetMaxOptionListSize returns the MaxOptionListSize field if non-nil, zero value otherwise.

### GetMaxOptionListSizeOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetMaxOptionListSizeOk() (*int64, bool)`

GetMaxOptionListSizeOk returns a tuple with the MaxOptionListSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOptionListSize

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetMaxOptionListSize(v int64)`

SetMaxOptionListSize sets MaxOptionListSize field to given value.

### HasMaxOptionListSize

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasMaxOptionListSize() bool`

HasMaxOptionListSize returns a boolean if a field has been set.

### GetExchangeUrl

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetExchangeUrl() string`

GetExchangeUrl returns the ExchangeUrl field if non-nil, zero value otherwise.

### GetExchangeUrlOk

`func (o *UpdateApplianceSettingsRequestApplianceSettings) GetExchangeUrlOk() (*string, bool)`

GetExchangeUrlOk returns a tuple with the ExchangeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeUrl

`func (o *UpdateApplianceSettingsRequestApplianceSettings) SetExchangeUrl(v string)`

SetExchangeUrl sets ExchangeUrl field to given value.

### HasExchangeUrl

`func (o *UpdateApplianceSettingsRequestApplianceSettings) HasExchangeUrl() bool`

HasExchangeUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


