# ListApplianceSettings200ResponseApplianceSettings

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

### NewListApplianceSettings200ResponseApplianceSettings

`func NewListApplianceSettings200ResponseApplianceSettings() *ListApplianceSettings200ResponseApplianceSettings`

NewListApplianceSettings200ResponseApplianceSettings instantiates a new ListApplianceSettings200ResponseApplianceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplianceSettings200ResponseApplianceSettingsWithDefaults

`func NewListApplianceSettings200ResponseApplianceSettingsWithDefaults() *ListApplianceSettings200ResponseApplianceSettings`

NewListApplianceSettings200ResponseApplianceSettingsWithDefaults instantiates a new ListApplianceSettings200ResponseApplianceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetApplianceId

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetApplianceId() string`

GetApplianceId returns the ApplianceId field if non-nil, zero value otherwise.

### GetApplianceIdOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetApplianceIdOk() (*string, bool)`

GetApplianceIdOk returns a tuple with the ApplianceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceId

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetApplianceId(v string)`

SetApplianceId sets ApplianceId field to given value.

### HasApplianceId

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasApplianceId() bool`

HasApplianceId returns a boolean if a field has been set.

### GetApplianceUrl

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetInternalApplianceUrl

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetInternalApplianceUrl() string`

GetInternalApplianceUrl returns the InternalApplianceUrl field if non-nil, zero value otherwise.

### GetInternalApplianceUrlOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetInternalApplianceUrlOk() (*string, bool)`

GetInternalApplianceUrlOk returns a tuple with the InternalApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalApplianceUrl

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetInternalApplianceUrl(v string)`

SetInternalApplianceUrl sets InternalApplianceUrl field to given value.

### HasInternalApplianceUrl

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasInternalApplianceUrl() bool`

HasInternalApplianceUrl returns a boolean if a field has been set.

### GetCorsAllowed

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetCorsAllowed() string`

GetCorsAllowed returns the CorsAllowed field if non-nil, zero value otherwise.

### GetCorsAllowedOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetCorsAllowedOk() (*string, bool)`

GetCorsAllowedOk returns a tuple with the CorsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowed

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetCorsAllowed(v string)`

SetCorsAllowed sets CorsAllowed field to given value.

### HasCorsAllowed

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasCorsAllowed() bool`

HasCorsAllowed returns a boolean if a field has been set.

### GetRegistrationEnabled

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetRegistrationEnabled() bool`

GetRegistrationEnabled returns the RegistrationEnabled field if non-nil, zero value otherwise.

### GetRegistrationEnabledOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetRegistrationEnabledOk() (*bool, bool)`

GetRegistrationEnabledOk returns a tuple with the RegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEnabled

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetRegistrationEnabled(v bool)`

SetRegistrationEnabled sets RegistrationEnabled field to given value.

### HasRegistrationEnabled

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasRegistrationEnabled() bool`

HasRegistrationEnabled returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetDefaultRoleId() string`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetDefaultRoleIdOk() (*string, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetDefaultRoleId(v string)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetDefaultUserRoleId

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetDefaultUserRoleId() string`

GetDefaultUserRoleId returns the DefaultUserRoleId field if non-nil, zero value otherwise.

### GetDefaultUserRoleIdOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetDefaultUserRoleIdOk() (*string, bool)`

GetDefaultUserRoleIdOk returns a tuple with the DefaultUserRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserRoleId

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetDefaultUserRoleId(v string)`

SetDefaultUserRoleId sets DefaultUserRoleId field to given value.

### HasDefaultUserRoleId

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasDefaultUserRoleId() bool`

HasDefaultUserRoleId returns a boolean if a field has been set.

### GetDockerPrivilegedMode

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetDockerPrivilegedMode() bool`

GetDockerPrivilegedMode returns the DockerPrivilegedMode field if non-nil, zero value otherwise.

### GetDockerPrivilegedModeOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetDockerPrivilegedModeOk() (*bool, bool)`

GetDockerPrivilegedModeOk returns a tuple with the DockerPrivilegedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerPrivilegedMode

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetDockerPrivilegedMode(v bool)`

SetDockerPrivilegedMode sets DockerPrivilegedMode field to given value.

### HasDockerPrivilegedMode

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasDockerPrivilegedMode() bool`

HasDockerPrivilegedMode returns a boolean if a field has been set.

### GetExpirePwdDays

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetExpirePwdDays() string`

GetExpirePwdDays returns the ExpirePwdDays field if non-nil, zero value otherwise.

### GetExpirePwdDaysOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetExpirePwdDaysOk() (*string, bool)`

GetExpirePwdDaysOk returns a tuple with the ExpirePwdDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirePwdDays

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetExpirePwdDays(v string)`

SetExpirePwdDays sets ExpirePwdDays field to given value.

### HasExpirePwdDays

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasExpirePwdDays() bool`

HasExpirePwdDays returns a boolean if a field has been set.

### GetDisableAfterAttempts

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetDisableAfterAttempts() string`

GetDisableAfterAttempts returns the DisableAfterAttempts field if non-nil, zero value otherwise.

### GetDisableAfterAttemptsOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetDisableAfterAttemptsOk() (*string, bool)`

GetDisableAfterAttemptsOk returns a tuple with the DisableAfterAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAfterAttempts

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetDisableAfterAttempts(v string)`

SetDisableAfterAttempts sets DisableAfterAttempts field to given value.

### HasDisableAfterAttempts

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasDisableAfterAttempts() bool`

HasDisableAfterAttempts returns a boolean if a field has been set.

### GetDisableAfterDaysInactive

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetDisableAfterDaysInactive() string`

GetDisableAfterDaysInactive returns the DisableAfterDaysInactive field if non-nil, zero value otherwise.

### GetDisableAfterDaysInactiveOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetDisableAfterDaysInactiveOk() (*string, bool)`

GetDisableAfterDaysInactiveOk returns a tuple with the DisableAfterDaysInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAfterDaysInactive

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetDisableAfterDaysInactive(v string)`

SetDisableAfterDaysInactive sets DisableAfterDaysInactive field to given value.

### HasDisableAfterDaysInactive

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasDisableAfterDaysInactive() bool`

HasDisableAfterDaysInactive returns a boolean if a field has been set.

### GetWarnUserDaysBefore

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetWarnUserDaysBefore() string`

GetWarnUserDaysBefore returns the WarnUserDaysBefore field if non-nil, zero value otherwise.

### GetWarnUserDaysBeforeOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetWarnUserDaysBeforeOk() (*string, bool)`

GetWarnUserDaysBeforeOk returns a tuple with the WarnUserDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnUserDaysBefore

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetWarnUserDaysBefore(v string)`

SetWarnUserDaysBefore sets WarnUserDaysBefore field to given value.

### HasWarnUserDaysBefore

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasWarnUserDaysBefore() bool`

HasWarnUserDaysBefore returns a boolean if a field has been set.

### GetSmtpMailFrom

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpMailFrom() string`

GetSmtpMailFrom returns the SmtpMailFrom field if non-nil, zero value otherwise.

### GetSmtpMailFromOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpMailFromOk() (*string, bool)`

GetSmtpMailFromOk returns a tuple with the SmtpMailFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpMailFrom

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetSmtpMailFrom(v string)`

SetSmtpMailFrom sets SmtpMailFrom field to given value.

### HasSmtpMailFrom

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasSmtpMailFrom() bool`

HasSmtpMailFrom returns a boolean if a field has been set.

### GetSmtpServer

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetSmtpPort

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpPort() string`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpPortOk() (*string, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetSmtpPort(v string)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpSSL

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpSSL() bool`

GetSmtpSSL returns the SmtpSSL field if non-nil, zero value otherwise.

### GetSmtpSSLOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpSSLOk() (*bool, bool)`

GetSmtpSSLOk returns a tuple with the SmtpSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSSL

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetSmtpSSL(v bool)`

SetSmtpSSL sets SmtpSSL field to given value.

### HasSmtpSSL

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasSmtpSSL() bool`

HasSmtpSSL returns a boolean if a field has been set.

### GetSmtpTLS

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpTLS() bool`

GetSmtpTLS returns the SmtpTLS field if non-nil, zero value otherwise.

### GetSmtpTLSOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpTLSOk() (*bool, bool)`

GetSmtpTLSOk returns a tuple with the SmtpTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpTLS

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetSmtpTLS(v bool)`

SetSmtpTLS sets SmtpTLS field to given value.

### HasSmtpTLS

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasSmtpTLS() bool`

HasSmtpTLS returns a boolean if a field has been set.

### GetSmtpUser

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpUser() string`

GetSmtpUser returns the SmtpUser field if non-nil, zero value otherwise.

### GetSmtpUserOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpUserOk() (*string, bool)`

GetSmtpUserOk returns a tuple with the SmtpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpUser

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetSmtpUser(v string)`

SetSmtpUser sets SmtpUser field to given value.

### HasSmtpUser

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasSmtpUser() bool`

HasSmtpUser returns a boolean if a field has been set.

### GetSmtpPassword

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpPassword() string`

GetSmtpPassword returns the SmtpPassword field if non-nil, zero value otherwise.

### GetSmtpPasswordOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpPasswordOk() (*string, bool)`

GetSmtpPasswordOk returns a tuple with the SmtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPassword

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetSmtpPassword(v string)`

SetSmtpPassword sets SmtpPassword field to given value.

### HasSmtpPassword

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasSmtpPassword() bool`

HasSmtpPassword returns a boolean if a field has been set.

### GetSmtpPasswordHash

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpPasswordHash() string`

GetSmtpPasswordHash returns the SmtpPasswordHash field if non-nil, zero value otherwise.

### GetSmtpPasswordHashOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetSmtpPasswordHashOk() (*string, bool)`

GetSmtpPasswordHashOk returns a tuple with the SmtpPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPasswordHash

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetSmtpPasswordHash(v string)`

SetSmtpPasswordHash sets SmtpPasswordHash field to given value.

### HasSmtpPasswordHash

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasSmtpPasswordHash() bool`

HasSmtpPasswordHash returns a boolean if a field has been set.

### GetProxyHost

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyHost() string`

GetProxyHost returns the ProxyHost field if non-nil, zero value otherwise.

### GetProxyHostOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyHostOk() (*string, bool)`

GetProxyHostOk returns a tuple with the ProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHost

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetProxyHost(v string)`

SetProxyHost sets ProxyHost field to given value.

### HasProxyHost

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasProxyHost() bool`

HasProxyHost returns a boolean if a field has been set.

### GetProxyPort

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyPort() string`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyPortOk() (*string, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetProxyPort(v string)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyUser

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyUser() string`

GetProxyUser returns the ProxyUser field if non-nil, zero value otherwise.

### GetProxyUserOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyUserOk() (*string, bool)`

GetProxyUserOk returns a tuple with the ProxyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUser

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetProxyUser(v string)`

SetProxyUser sets ProxyUser field to given value.

### HasProxyUser

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasProxyUser() bool`

HasProxyUser returns a boolean if a field has been set.

### GetProxyPassword

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetProxyPasswordHash

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyPasswordHash() string`

GetProxyPasswordHash returns the ProxyPasswordHash field if non-nil, zero value otherwise.

### GetProxyPasswordHashOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyPasswordHashOk() (*string, bool)`

GetProxyPasswordHashOk returns a tuple with the ProxyPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPasswordHash

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetProxyPasswordHash(v string)`

SetProxyPasswordHash sets ProxyPasswordHash field to given value.

### HasProxyPasswordHash

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasProxyPasswordHash() bool`

HasProxyPasswordHash returns a boolean if a field has been set.

### GetProxyDomain

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyDomain() string`

GetProxyDomain returns the ProxyDomain field if non-nil, zero value otherwise.

### GetProxyDomainOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyDomainOk() (*string, bool)`

GetProxyDomainOk returns a tuple with the ProxyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyDomain

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetProxyDomain(v string)`

SetProxyDomain sets ProxyDomain field to given value.

### HasProxyDomain

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasProxyDomain() bool`

HasProxyDomain returns a boolean if a field has been set.

### GetProxyWorkstation

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyWorkstation() string`

GetProxyWorkstation returns the ProxyWorkstation field if non-nil, zero value otherwise.

### GetProxyWorkstationOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetProxyWorkstationOk() (*string, bool)`

GetProxyWorkstationOk returns a tuple with the ProxyWorkstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyWorkstation

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetProxyWorkstation(v string)`

SetProxyWorkstation sets ProxyWorkstation field to given value.

### HasProxyWorkstation

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasProxyWorkstation() bool`

HasProxyWorkstation returns a boolean if a field has been set.

### GetCurrencyProvider

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetCurrencyProvider() string`

GetCurrencyProvider returns the CurrencyProvider field if non-nil, zero value otherwise.

### GetCurrencyProviderOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetCurrencyProviderOk() (*string, bool)`

GetCurrencyProviderOk returns a tuple with the CurrencyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyProvider

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetCurrencyProvider(v string)`

SetCurrencyProvider sets CurrencyProvider field to given value.

### HasCurrencyProvider

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasCurrencyProvider() bool`

HasCurrencyProvider returns a boolean if a field has been set.

### GetCurrencyKey

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetCurrencyKey() string`

GetCurrencyKey returns the CurrencyKey field if non-nil, zero value otherwise.

### GetCurrencyKeyOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetCurrencyKeyOk() (*string, bool)`

GetCurrencyKeyOk returns a tuple with the CurrencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyKey

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetCurrencyKey(v string)`

SetCurrencyKey sets CurrencyKey field to given value.

### HasCurrencyKey

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasCurrencyKey() bool`

HasCurrencyKey returns a boolean if a field has been set.

### GetEnabledZoneTypes

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetEnabledZoneTypes() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetEnabledZoneTypes returns the EnabledZoneTypes field if non-nil, zero value otherwise.

### GetEnabledZoneTypesOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetEnabledZoneTypesOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetEnabledZoneTypesOk returns a tuple with the EnabledZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledZoneTypes

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetEnabledZoneTypes(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetEnabledZoneTypes sets EnabledZoneTypes field to given value.

### HasEnabledZoneTypes

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasEnabledZoneTypes() bool`

HasEnabledZoneTypes returns a boolean if a field has been set.

### GetStatsRetainmentPeriod

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetStatsRetainmentPeriod() int64`

GetStatsRetainmentPeriod returns the StatsRetainmentPeriod field if non-nil, zero value otherwise.

### GetStatsRetainmentPeriodOk

`func (o *ListApplianceSettings200ResponseApplianceSettings) GetStatsRetainmentPeriodOk() (*int64, bool)`

GetStatsRetainmentPeriodOk returns a tuple with the StatsRetainmentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsRetainmentPeriod

`func (o *ListApplianceSettings200ResponseApplianceSettings) SetStatsRetainmentPeriod(v int64)`

SetStatsRetainmentPeriod sets StatsRetainmentPeriod field to given value.

### HasStatsRetainmentPeriod

`func (o *ListApplianceSettings200ResponseApplianceSettings) HasStatsRetainmentPeriod() bool`

HasStatsRetainmentPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


