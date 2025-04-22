# ApplianceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**ApplianceId** | Pointer to **string** |  | [optional] 
**ApplianceUrl** | Pointer to **string** |  | [optional] 
**InternalApplianceUrl** | Pointer to **string** |  | [optional] 
**CorsAllowed** | Pointer to **string** |  | [optional] 
**RegistrationEnabled** | Pointer to **bool** |  | [optional] 
**DefaultRoleId** | Pointer to **string** |  | [optional] 
**DefaultUserRoleId** | Pointer to **string** |  | [optional] 
**DockerPrivilegedMode** | Pointer to **bool** |  | [optional] 
**ExpirePwdDays** | Pointer to **string** |  | [optional] 
**DisableAfterAttempts** | Pointer to **string** |  | [optional] 
**DisableAfterDaysInactive** | Pointer to **string** |  | [optional] 
**WarnUserDaysBefore** | Pointer to **string** |  | [optional] 
**SmtpMailFrom** | Pointer to **string** |  | [optional] 
**SmtpServer** | Pointer to **string** |  | [optional] 
**SmtpPort** | Pointer to **string** |  | [optional] 
**SmtpSSL** | Pointer to **bool** |  | [optional] 
**SmtpTLS** | Pointer to **bool** |  | [optional] 
**SmtpUser** | Pointer to **string** |  | [optional] 
**SmtpPassword** | Pointer to **string** |  | [optional] 
**SmtpPasswordHash** | Pointer to **string** |  | [optional] 
**ProxyHost** | Pointer to **string** |  | [optional] 
**ProxyPort** | Pointer to **string** |  | [optional] 
**ProxyUser** | Pointer to **string** |  | [optional] 
**ProxyPassword** | Pointer to **string** |  | [optional] 
**ProxyPasswordHash** | Pointer to **string** |  | [optional] 
**ProxyDomain** | Pointer to **string** |  | [optional] 
**ProxyWorkstation** | Pointer to **string** |  | [optional] 
**CurrencyProvider** | Pointer to **string** |  | [optional] 
**CurrencyKey** | Pointer to **string** |  | [optional] 
**EnabledZoneTypes** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**StatsRetainmentPeriod** | Pointer to **int64** |  | [optional] 

## Methods

### NewApplianceSettings

`func NewApplianceSettings() *ApplianceSettings`

NewApplianceSettings instantiates a new ApplianceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceSettingsWithDefaults

`func NewApplianceSettingsWithDefaults() *ApplianceSettings`

NewApplianceSettingsWithDefaults instantiates a new ApplianceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ApplianceSettings) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApplianceSettings) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApplianceSettings) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApplianceSettings) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetApplianceId

`func (o *ApplianceSettings) GetApplianceId() string`

GetApplianceId returns the ApplianceId field if non-nil, zero value otherwise.

### GetApplianceIdOk

`func (o *ApplianceSettings) GetApplianceIdOk() (*string, bool)`

GetApplianceIdOk returns a tuple with the ApplianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceId

`func (o *ApplianceSettings) SetApplianceId(v string)`

SetApplianceId sets ApplianceId field to given value.

### HasApplianceId

`func (o *ApplianceSettings) HasApplianceId() bool`

HasApplianceId returns a boolean if a field has been set.

### GetApplianceUrl

`func (o *ApplianceSettings) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ApplianceSettings) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ApplianceSettings) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ApplianceSettings) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetInternalApplianceUrl

`func (o *ApplianceSettings) GetInternalApplianceUrl() string`

GetInternalApplianceUrl returns the InternalApplianceUrl field if non-nil, zero value otherwise.

### GetInternalApplianceUrlOk

`func (o *ApplianceSettings) GetInternalApplianceUrlOk() (*string, bool)`

GetInternalApplianceUrlOk returns a tuple with the InternalApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalApplianceUrl

`func (o *ApplianceSettings) SetInternalApplianceUrl(v string)`

SetInternalApplianceUrl sets InternalApplianceUrl field to given value.

### HasInternalApplianceUrl

`func (o *ApplianceSettings) HasInternalApplianceUrl() bool`

HasInternalApplianceUrl returns a boolean if a field has been set.

### GetCorsAllowed

`func (o *ApplianceSettings) GetCorsAllowed() string`

GetCorsAllowed returns the CorsAllowed field if non-nil, zero value otherwise.

### GetCorsAllowedOk

`func (o *ApplianceSettings) GetCorsAllowedOk() (*string, bool)`

GetCorsAllowedOk returns a tuple with the CorsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowed

`func (o *ApplianceSettings) SetCorsAllowed(v string)`

SetCorsAllowed sets CorsAllowed field to given value.

### HasCorsAllowed

`func (o *ApplianceSettings) HasCorsAllowed() bool`

HasCorsAllowed returns a boolean if a field has been set.

### GetRegistrationEnabled

`func (o *ApplianceSettings) GetRegistrationEnabled() bool`

GetRegistrationEnabled returns the RegistrationEnabled field if non-nil, zero value otherwise.

### GetRegistrationEnabledOk

`func (o *ApplianceSettings) GetRegistrationEnabledOk() (*bool, bool)`

GetRegistrationEnabledOk returns a tuple with the RegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEnabled

`func (o *ApplianceSettings) SetRegistrationEnabled(v bool)`

SetRegistrationEnabled sets RegistrationEnabled field to given value.

### HasRegistrationEnabled

`func (o *ApplianceSettings) HasRegistrationEnabled() bool`

HasRegistrationEnabled returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *ApplianceSettings) GetDefaultRoleId() string`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *ApplianceSettings) GetDefaultRoleIdOk() (*string, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *ApplianceSettings) SetDefaultRoleId(v string)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *ApplianceSettings) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetDefaultUserRoleId

`func (o *ApplianceSettings) GetDefaultUserRoleId() string`

GetDefaultUserRoleId returns the DefaultUserRoleId field if non-nil, zero value otherwise.

### GetDefaultUserRoleIdOk

`func (o *ApplianceSettings) GetDefaultUserRoleIdOk() (*string, bool)`

GetDefaultUserRoleIdOk returns a tuple with the DefaultUserRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserRoleId

`func (o *ApplianceSettings) SetDefaultUserRoleId(v string)`

SetDefaultUserRoleId sets DefaultUserRoleId field to given value.

### HasDefaultUserRoleId

`func (o *ApplianceSettings) HasDefaultUserRoleId() bool`

HasDefaultUserRoleId returns a boolean if a field has been set.

### GetDockerPrivilegedMode

`func (o *ApplianceSettings) GetDockerPrivilegedMode() bool`

GetDockerPrivilegedMode returns the DockerPrivilegedMode field if non-nil, zero value otherwise.

### GetDockerPrivilegedModeOk

`func (o *ApplianceSettings) GetDockerPrivilegedModeOk() (*bool, bool)`

GetDockerPrivilegedModeOk returns a tuple with the DockerPrivilegedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerPrivilegedMode

`func (o *ApplianceSettings) SetDockerPrivilegedMode(v bool)`

SetDockerPrivilegedMode sets DockerPrivilegedMode field to given value.

### HasDockerPrivilegedMode

`func (o *ApplianceSettings) HasDockerPrivilegedMode() bool`

HasDockerPrivilegedMode returns a boolean if a field has been set.

### GetExpirePwdDays

`func (o *ApplianceSettings) GetExpirePwdDays() string`

GetExpirePwdDays returns the ExpirePwdDays field if non-nil, zero value otherwise.

### GetExpirePwdDaysOk

`func (o *ApplianceSettings) GetExpirePwdDaysOk() (*string, bool)`

GetExpirePwdDaysOk returns a tuple with the ExpirePwdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirePwdDays

`func (o *ApplianceSettings) SetExpirePwdDays(v string)`

SetExpirePwdDays sets ExpirePwdDays field to given value.

### HasExpirePwdDays

`func (o *ApplianceSettings) HasExpirePwdDays() bool`

HasExpirePwdDays returns a boolean if a field has been set.

### GetDisableAfterAttempts

`func (o *ApplianceSettings) GetDisableAfterAttempts() string`

GetDisableAfterAttempts returns the DisableAfterAttempts field if non-nil, zero value otherwise.

### GetDisableAfterAttemptsOk

`func (o *ApplianceSettings) GetDisableAfterAttemptsOk() (*string, bool)`

GetDisableAfterAttemptsOk returns a tuple with the DisableAfterAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAfterAttempts

`func (o *ApplianceSettings) SetDisableAfterAttempts(v string)`

SetDisableAfterAttempts sets DisableAfterAttempts field to given value.

### HasDisableAfterAttempts

`func (o *ApplianceSettings) HasDisableAfterAttempts() bool`

HasDisableAfterAttempts returns a boolean if a field has been set.

### GetDisableAfterDaysInactive

`func (o *ApplianceSettings) GetDisableAfterDaysInactive() string`

GetDisableAfterDaysInactive returns the DisableAfterDaysInactive field if non-nil, zero value otherwise.

### GetDisableAfterDaysInactiveOk

`func (o *ApplianceSettings) GetDisableAfterDaysInactiveOk() (*string, bool)`

GetDisableAfterDaysInactiveOk returns a tuple with the DisableAfterDaysInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAfterDaysInactive

`func (o *ApplianceSettings) SetDisableAfterDaysInactive(v string)`

SetDisableAfterDaysInactive sets DisableAfterDaysInactive field to given value.

### HasDisableAfterDaysInactive

`func (o *ApplianceSettings) HasDisableAfterDaysInactive() bool`

HasDisableAfterDaysInactive returns a boolean if a field has been set.

### GetWarnUserDaysBefore

`func (o *ApplianceSettings) GetWarnUserDaysBefore() string`

GetWarnUserDaysBefore returns the WarnUserDaysBefore field if non-nil, zero value otherwise.

### GetWarnUserDaysBeforeOk

`func (o *ApplianceSettings) GetWarnUserDaysBeforeOk() (*string, bool)`

GetWarnUserDaysBeforeOk returns a tuple with the WarnUserDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnUserDaysBefore

`func (o *ApplianceSettings) SetWarnUserDaysBefore(v string)`

SetWarnUserDaysBefore sets WarnUserDaysBefore field to given value.

### HasWarnUserDaysBefore

`func (o *ApplianceSettings) HasWarnUserDaysBefore() bool`

HasWarnUserDaysBefore returns a boolean if a field has been set.

### GetSmtpMailFrom

`func (o *ApplianceSettings) GetSmtpMailFrom() string`

GetSmtpMailFrom returns the SmtpMailFrom field if non-nil, zero value otherwise.

### GetSmtpMailFromOk

`func (o *ApplianceSettings) GetSmtpMailFromOk() (*string, bool)`

GetSmtpMailFromOk returns a tuple with the SmtpMailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpMailFrom

`func (o *ApplianceSettings) SetSmtpMailFrom(v string)`

SetSmtpMailFrom sets SmtpMailFrom field to given value.

### HasSmtpMailFrom

`func (o *ApplianceSettings) HasSmtpMailFrom() bool`

HasSmtpMailFrom returns a boolean if a field has been set.

### GetSmtpServer

`func (o *ApplianceSettings) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *ApplianceSettings) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *ApplianceSettings) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *ApplianceSettings) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetSmtpPort

`func (o *ApplianceSettings) GetSmtpPort() string`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *ApplianceSettings) GetSmtpPortOk() (*string, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *ApplianceSettings) SetSmtpPort(v string)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *ApplianceSettings) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpSSL

`func (o *ApplianceSettings) GetSmtpSSL() bool`

GetSmtpSSL returns the SmtpSSL field if non-nil, zero value otherwise.

### GetSmtpSSLOk

`func (o *ApplianceSettings) GetSmtpSSLOk() (*bool, bool)`

GetSmtpSSLOk returns a tuple with the SmtpSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSSL

`func (o *ApplianceSettings) SetSmtpSSL(v bool)`

SetSmtpSSL sets SmtpSSL field to given value.

### HasSmtpSSL

`func (o *ApplianceSettings) HasSmtpSSL() bool`

HasSmtpSSL returns a boolean if a field has been set.

### GetSmtpTLS

`func (o *ApplianceSettings) GetSmtpTLS() bool`

GetSmtpTLS returns the SmtpTLS field if non-nil, zero value otherwise.

### GetSmtpTLSOk

`func (o *ApplianceSettings) GetSmtpTLSOk() (*bool, bool)`

GetSmtpTLSOk returns a tuple with the SmtpTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpTLS

`func (o *ApplianceSettings) SetSmtpTLS(v bool)`

SetSmtpTLS sets SmtpTLS field to given value.

### HasSmtpTLS

`func (o *ApplianceSettings) HasSmtpTLS() bool`

HasSmtpTLS returns a boolean if a field has been set.

### GetSmtpUser

`func (o *ApplianceSettings) GetSmtpUser() string`

GetSmtpUser returns the SmtpUser field if non-nil, zero value otherwise.

### GetSmtpUserOk

`func (o *ApplianceSettings) GetSmtpUserOk() (*string, bool)`

GetSmtpUserOk returns a tuple with the SmtpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpUser

`func (o *ApplianceSettings) SetSmtpUser(v string)`

SetSmtpUser sets SmtpUser field to given value.

### HasSmtpUser

`func (o *ApplianceSettings) HasSmtpUser() bool`

HasSmtpUser returns a boolean if a field has been set.

### GetSmtpPassword

`func (o *ApplianceSettings) GetSmtpPassword() string`

GetSmtpPassword returns the SmtpPassword field if non-nil, zero value otherwise.

### GetSmtpPasswordOk

`func (o *ApplianceSettings) GetSmtpPasswordOk() (*string, bool)`

GetSmtpPasswordOk returns a tuple with the SmtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPassword

`func (o *ApplianceSettings) SetSmtpPassword(v string)`

SetSmtpPassword sets SmtpPassword field to given value.

### HasSmtpPassword

`func (o *ApplianceSettings) HasSmtpPassword() bool`

HasSmtpPassword returns a boolean if a field has been set.

### GetSmtpPasswordHash

`func (o *ApplianceSettings) GetSmtpPasswordHash() string`

GetSmtpPasswordHash returns the SmtpPasswordHash field if non-nil, zero value otherwise.

### GetSmtpPasswordHashOk

`func (o *ApplianceSettings) GetSmtpPasswordHashOk() (*string, bool)`

GetSmtpPasswordHashOk returns a tuple with the SmtpPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPasswordHash

`func (o *ApplianceSettings) SetSmtpPasswordHash(v string)`

SetSmtpPasswordHash sets SmtpPasswordHash field to given value.

### HasSmtpPasswordHash

`func (o *ApplianceSettings) HasSmtpPasswordHash() bool`

HasSmtpPasswordHash returns a boolean if a field has been set.

### GetProxyHost

`func (o *ApplianceSettings) GetProxyHost() string`

GetProxyHost returns the ProxyHost field if non-nil, zero value otherwise.

### GetProxyHostOk

`func (o *ApplianceSettings) GetProxyHostOk() (*string, bool)`

GetProxyHostOk returns a tuple with the ProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHost

`func (o *ApplianceSettings) SetProxyHost(v string)`

SetProxyHost sets ProxyHost field to given value.

### HasProxyHost

`func (o *ApplianceSettings) HasProxyHost() bool`

HasProxyHost returns a boolean if a field has been set.

### GetProxyPort

`func (o *ApplianceSettings) GetProxyPort() string`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *ApplianceSettings) GetProxyPortOk() (*string, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *ApplianceSettings) SetProxyPort(v string)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *ApplianceSettings) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyUser

`func (o *ApplianceSettings) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *ApplianceSettings) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *ApplianceSettings) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *ApplianceSettings) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### GetProxyPassword

`func (o *ApplianceSettings) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *ApplianceSettings) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *ApplianceSettings) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *ApplianceSettings) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetProxyPasswordHash

`func (o *ApplianceSettings) GetProxyPasswordHash() string`

GetProxyPasswordHash returns the ProxyPasswordHash field if non-nil, zero value otherwise.

### GetProxyPasswordHashOk

`func (o *ApplianceSettings) GetProxyPasswordHashOk() (*string, bool)`

GetProxyPasswordHashOk returns a tuple with the ProxyPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPasswordHash

`func (o *ApplianceSettings) SetProxyPasswordHash(v string)`

SetProxyPasswordHash sets ProxyPasswordHash field to given value.

### HasProxyPasswordHash

`func (o *ApplianceSettings) HasProxyPasswordHash() bool`

HasProxyPasswordHash returns a boolean if a field has been set.

### GetProxyDomain

`func (o *ApplianceSettings) GetProxyDomain() string`

GetProxyDomain returns the ProxyDomain field if non-nil, zero value otherwise.

### GetProxyDomainOk

`func (o *ApplianceSettings) GetProxyDomainOk() (*string, bool)`

GetProxyDomainOk returns a tuple with the ProxyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyDomain

`func (o *ApplianceSettings) SetProxyDomain(v string)`

SetProxyDomain sets ProxyDomain field to given value.

### HasProxyDomain

`func (o *ApplianceSettings) HasProxyDomain() bool`

HasProxyDomain returns a boolean if a field has been set.

### GetProxyWorkstation

`func (o *ApplianceSettings) GetProxyWorkstation() string`

GetProxyWorkstation returns the ProxyWorkstation field if non-nil, zero value otherwise.

### GetProxyWorkstationOk

`func (o *ApplianceSettings) GetProxyWorkstationOk() (*string, bool)`

GetProxyWorkstationOk returns a tuple with the ProxyWorkstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyWorkstation

`func (o *ApplianceSettings) SetProxyWorkstation(v string)`

SetProxyWorkstation sets ProxyWorkstation field to given value.

### HasProxyWorkstation

`func (o *ApplianceSettings) HasProxyWorkstation() bool`

HasProxyWorkstation returns a boolean if a field has been set.

### GetCurrencyProvider

`func (o *ApplianceSettings) GetCurrencyProvider() string`

GetCurrencyProvider returns the CurrencyProvider field if non-nil, zero value otherwise.

### GetCurrencyProviderOk

`func (o *ApplianceSettings) GetCurrencyProviderOk() (*string, bool)`

GetCurrencyProviderOk returns a tuple with the CurrencyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyProvider

`func (o *ApplianceSettings) SetCurrencyProvider(v string)`

SetCurrencyProvider sets CurrencyProvider field to given value.

### HasCurrencyProvider

`func (o *ApplianceSettings) HasCurrencyProvider() bool`

HasCurrencyProvider returns a boolean if a field has been set.

### GetCurrencyKey

`func (o *ApplianceSettings) GetCurrencyKey() string`

GetCurrencyKey returns the CurrencyKey field if non-nil, zero value otherwise.

### GetCurrencyKeyOk

`func (o *ApplianceSettings) GetCurrencyKeyOk() (*string, bool)`

GetCurrencyKeyOk returns a tuple with the CurrencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyKey

`func (o *ApplianceSettings) SetCurrencyKey(v string)`

SetCurrencyKey sets CurrencyKey field to given value.

### HasCurrencyKey

`func (o *ApplianceSettings) HasCurrencyKey() bool`

HasCurrencyKey returns a boolean if a field has been set.

### GetEnabledZoneTypes

`func (o *ApplianceSettings) GetEnabledZoneTypes() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetEnabledZoneTypes returns the EnabledZoneTypes field if non-nil, zero value otherwise.

### GetEnabledZoneTypesOk

`func (o *ApplianceSettings) GetEnabledZoneTypesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetEnabledZoneTypesOk returns a tuple with the EnabledZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledZoneTypes

`func (o *ApplianceSettings) SetEnabledZoneTypes(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetEnabledZoneTypes sets EnabledZoneTypes field to given value.

### HasEnabledZoneTypes

`func (o *ApplianceSettings) HasEnabledZoneTypes() bool`

HasEnabledZoneTypes returns a boolean if a field has been set.

### GetStatsRetainmentPeriod

`func (o *ApplianceSettings) GetStatsRetainmentPeriod() int64`

GetStatsRetainmentPeriod returns the StatsRetainmentPeriod field if non-nil, zero value otherwise.

### GetStatsRetainmentPeriodOk

`func (o *ApplianceSettings) GetStatsRetainmentPeriodOk() (*int64, bool)`

GetStatsRetainmentPeriodOk returns a tuple with the StatsRetainmentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsRetainmentPeriod

`func (o *ApplianceSettings) SetStatsRetainmentPeriod(v int64)`

SetStatsRetainmentPeriod sets StatsRetainmentPeriod field to given value.

### HasStatsRetainmentPeriod

`func (o *ApplianceSettings) HasStatsRetainmentPeriod() bool`

HasStatsRetainmentPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


