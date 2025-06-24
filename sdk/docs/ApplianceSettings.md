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
**DefaultUserRoleId** | Pointer to **NullableString** |  | [optional] 
**DockerPrivilegedMode** | Pointer to **bool** |  | [optional] 
**ExpirePwdDays** | Pointer to **NullableString** |  | [optional] 
**DisableAfterAttempts** | Pointer to **NullableString** |  | [optional] 
**DisableAfterDaysInactive** | Pointer to **NullableString** |  | [optional] 
**WarnUserDaysBefore** | Pointer to **NullableString** |  | [optional] 
**SmtpMailFrom** | Pointer to **NullableString** |  | [optional] 
**SmtpServer** | Pointer to **NullableString** |  | [optional] 
**SmtpPort** | Pointer to **NullableString** |  | [optional] 
**SmtpSSL** | Pointer to **bool** |  | [optional] 
**SmtpTLS** | Pointer to **bool** |  | [optional] 
**SmtpUser** | Pointer to **NullableString** |  | [optional] 
**SmtpPassword** | Pointer to **NullableString** |  | [optional] 
**SmtpPasswordHash** | Pointer to **NullableString** |  | [optional] 
**ProxyHost** | Pointer to **NullableString** |  | [optional] 
**ProxyPort** | Pointer to **NullableString** |  | [optional] 
**ProxyUser** | Pointer to **NullableString** |  | [optional] 
**ProxyPassword** | Pointer to **NullableString** |  | [optional] 
**ProxyPasswordHash** | Pointer to **NullableString** |  | [optional] 
**ProxyDomain** | Pointer to **NullableString** |  | [optional] 
**ProxyWorkstation** | Pointer to **NullableString** |  | [optional] 
**CurrencyProvider** | Pointer to **NullableString** |  | [optional] 
**CurrencyKey** | Pointer to **NullableString** |  | [optional] 
**EnabledZoneTypes** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
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

### SetDefaultUserRoleIdNil

`func (o *ApplianceSettings) SetDefaultUserRoleIdNil(b bool)`

 SetDefaultUserRoleIdNil sets the value for DefaultUserRoleId to be an explicit nil

### UnsetDefaultUserRoleId
`func (o *ApplianceSettings) UnsetDefaultUserRoleId()`

UnsetDefaultUserRoleId ensures that no value is present for DefaultUserRoleId, not even an explicit nil
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

### SetExpirePwdDaysNil

`func (o *ApplianceSettings) SetExpirePwdDaysNil(b bool)`

 SetExpirePwdDaysNil sets the value for ExpirePwdDays to be an explicit nil

### UnsetExpirePwdDays
`func (o *ApplianceSettings) UnsetExpirePwdDays()`

UnsetExpirePwdDays ensures that no value is present for ExpirePwdDays, not even an explicit nil
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

### SetDisableAfterAttemptsNil

`func (o *ApplianceSettings) SetDisableAfterAttemptsNil(b bool)`

 SetDisableAfterAttemptsNil sets the value for DisableAfterAttempts to be an explicit nil

### UnsetDisableAfterAttempts
`func (o *ApplianceSettings) UnsetDisableAfterAttempts()`

UnsetDisableAfterAttempts ensures that no value is present for DisableAfterAttempts, not even an explicit nil
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

### SetDisableAfterDaysInactiveNil

`func (o *ApplianceSettings) SetDisableAfterDaysInactiveNil(b bool)`

 SetDisableAfterDaysInactiveNil sets the value for DisableAfterDaysInactive to be an explicit nil

### UnsetDisableAfterDaysInactive
`func (o *ApplianceSettings) UnsetDisableAfterDaysInactive()`

UnsetDisableAfterDaysInactive ensures that no value is present for DisableAfterDaysInactive, not even an explicit nil
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

### SetWarnUserDaysBeforeNil

`func (o *ApplianceSettings) SetWarnUserDaysBeforeNil(b bool)`

 SetWarnUserDaysBeforeNil sets the value for WarnUserDaysBefore to be an explicit nil

### UnsetWarnUserDaysBefore
`func (o *ApplianceSettings) UnsetWarnUserDaysBefore()`

UnsetWarnUserDaysBefore ensures that no value is present for WarnUserDaysBefore, not even an explicit nil
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

### SetSmtpMailFromNil

`func (o *ApplianceSettings) SetSmtpMailFromNil(b bool)`

 SetSmtpMailFromNil sets the value for SmtpMailFrom to be an explicit nil

### UnsetSmtpMailFrom
`func (o *ApplianceSettings) UnsetSmtpMailFrom()`

UnsetSmtpMailFrom ensures that no value is present for SmtpMailFrom, not even an explicit nil
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

### SetSmtpServerNil

`func (o *ApplianceSettings) SetSmtpServerNil(b bool)`

 SetSmtpServerNil sets the value for SmtpServer to be an explicit nil

### UnsetSmtpServer
`func (o *ApplianceSettings) UnsetSmtpServer()`

UnsetSmtpServer ensures that no value is present for SmtpServer, not even an explicit nil
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

### SetSmtpPortNil

`func (o *ApplianceSettings) SetSmtpPortNil(b bool)`

 SetSmtpPortNil sets the value for SmtpPort to be an explicit nil

### UnsetSmtpPort
`func (o *ApplianceSettings) UnsetSmtpPort()`

UnsetSmtpPort ensures that no value is present for SmtpPort, not even an explicit nil
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

### SetSmtpUserNil

`func (o *ApplianceSettings) SetSmtpUserNil(b bool)`

 SetSmtpUserNil sets the value for SmtpUser to be an explicit nil

### UnsetSmtpUser
`func (o *ApplianceSettings) UnsetSmtpUser()`

UnsetSmtpUser ensures that no value is present for SmtpUser, not even an explicit nil
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

### SetSmtpPasswordNil

`func (o *ApplianceSettings) SetSmtpPasswordNil(b bool)`

 SetSmtpPasswordNil sets the value for SmtpPassword to be an explicit nil

### UnsetSmtpPassword
`func (o *ApplianceSettings) UnsetSmtpPassword()`

UnsetSmtpPassword ensures that no value is present for SmtpPassword, not even an explicit nil
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

### SetSmtpPasswordHashNil

`func (o *ApplianceSettings) SetSmtpPasswordHashNil(b bool)`

 SetSmtpPasswordHashNil sets the value for SmtpPasswordHash to be an explicit nil

### UnsetSmtpPasswordHash
`func (o *ApplianceSettings) UnsetSmtpPasswordHash()`

UnsetSmtpPasswordHash ensures that no value is present for SmtpPasswordHash, not even an explicit nil
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

### SetProxyHostNil

`func (o *ApplianceSettings) SetProxyHostNil(b bool)`

 SetProxyHostNil sets the value for ProxyHost to be an explicit nil

### UnsetProxyHost
`func (o *ApplianceSettings) UnsetProxyHost()`

UnsetProxyHost ensures that no value is present for ProxyHost, not even an explicit nil
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

### SetProxyPortNil

`func (o *ApplianceSettings) SetProxyPortNil(b bool)`

 SetProxyPortNil sets the value for ProxyPort to be an explicit nil

### UnsetProxyPort
`func (o *ApplianceSettings) UnsetProxyPort()`

UnsetProxyPort ensures that no value is present for ProxyPort, not even an explicit nil
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

### SetProxyUserNil

`func (o *ApplianceSettings) SetProxyUserNil(b bool)`

 SetProxyUserNil sets the value for ProxyUser to be an explicit nil

### UnsetProxyUser
`func (o *ApplianceSettings) UnsetProxyUser()`

UnsetProxyUser ensures that no value is present for ProxyUser, not even an explicit nil
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

### SetProxyPasswordNil

`func (o *ApplianceSettings) SetProxyPasswordNil(b bool)`

 SetProxyPasswordNil sets the value for ProxyPassword to be an explicit nil

### UnsetProxyPassword
`func (o *ApplianceSettings) UnsetProxyPassword()`

UnsetProxyPassword ensures that no value is present for ProxyPassword, not even an explicit nil
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

### SetProxyPasswordHashNil

`func (o *ApplianceSettings) SetProxyPasswordHashNil(b bool)`

 SetProxyPasswordHashNil sets the value for ProxyPasswordHash to be an explicit nil

### UnsetProxyPasswordHash
`func (o *ApplianceSettings) UnsetProxyPasswordHash()`

UnsetProxyPasswordHash ensures that no value is present for ProxyPasswordHash, not even an explicit nil
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

### SetProxyDomainNil

`func (o *ApplianceSettings) SetProxyDomainNil(b bool)`

 SetProxyDomainNil sets the value for ProxyDomain to be an explicit nil

### UnsetProxyDomain
`func (o *ApplianceSettings) UnsetProxyDomain()`

UnsetProxyDomain ensures that no value is present for ProxyDomain, not even an explicit nil
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

### SetProxyWorkstationNil

`func (o *ApplianceSettings) SetProxyWorkstationNil(b bool)`

 SetProxyWorkstationNil sets the value for ProxyWorkstation to be an explicit nil

### UnsetProxyWorkstation
`func (o *ApplianceSettings) UnsetProxyWorkstation()`

UnsetProxyWorkstation ensures that no value is present for ProxyWorkstation, not even an explicit nil
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

### SetCurrencyProviderNil

`func (o *ApplianceSettings) SetCurrencyProviderNil(b bool)`

 SetCurrencyProviderNil sets the value for CurrencyProvider to be an explicit nil

### UnsetCurrencyProvider
`func (o *ApplianceSettings) UnsetCurrencyProvider()`

UnsetCurrencyProvider ensures that no value is present for CurrencyProvider, not even an explicit nil
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

### SetCurrencyKeyNil

`func (o *ApplianceSettings) SetCurrencyKeyNil(b bool)`

 SetCurrencyKeyNil sets the value for CurrencyKey to be an explicit nil

### UnsetCurrencyKey
`func (o *ApplianceSettings) UnsetCurrencyKey()`

UnsetCurrencyKey ensures that no value is present for CurrencyKey, not even an explicit nil
### GetEnabledZoneTypes

`func (o *ApplianceSettings) GetEnabledZoneTypes() []GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetEnabledZoneTypes returns the EnabledZoneTypes field if non-nil, zero value otherwise.

### GetEnabledZoneTypesOk

`func (o *ApplianceSettings) GetEnabledZoneTypesOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetEnabledZoneTypesOk returns a tuple with the EnabledZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledZoneTypes

`func (o *ApplianceSettings) SetEnabledZoneTypes(v []GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetEnabledZoneTypes sets EnabledZoneTypes field to given value.

### HasEnabledZoneTypes

`func (o *ApplianceSettings) HasEnabledZoneTypes() bool`

HasEnabledZoneTypes returns a boolean if a field has been set.

### SetEnabledZoneTypesNil

`func (o *ApplianceSettings) SetEnabledZoneTypesNil(b bool)`

 SetEnabledZoneTypesNil sets the value for EnabledZoneTypes to be an explicit nil

### UnsetEnabledZoneTypes
`func (o *ApplianceSettings) UnsetEnabledZoneTypes()`

UnsetEnabledZoneTypes ensures that no value is present for EnabledZoneTypes, not even an explicit nil
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


