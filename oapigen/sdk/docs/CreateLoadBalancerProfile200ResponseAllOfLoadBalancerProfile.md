# CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileLoadBalancer**](CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileLoadBalancer.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**ServiceTypeDisplay** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ProxyType** | Pointer to **NullableString** |  | [optional] 
**RedirectRewrite** | Pointer to **NullableString** |  | [optional] 
**PersistenceType** | Pointer to **NullableString** |  | [optional] 
**SslEnabled** | Pointer to **NullableString** |  | [optional] 
**SslCert** | Pointer to **NullableString** |  | [optional] 
**AccountCertificate** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**RedirectUrl** | Pointer to **NullableString** |  | [optional] 
**InsertXforwardedFor** | Pointer to **bool** |  | [optional] 
**PersistenceCookieName** | Pointer to **NullableString** |  | [optional] 
**PersistenceExpiresIn** | Pointer to **NullableString** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile

`func NewCreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile() *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile`

NewCreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile instantiates a new CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileWithDefaults

`func NewCreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileWithDefaults() *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile`

NewCreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileWithDefaults instantiates a new CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetLoadBalancer() CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetLoadBalancerOk() (*CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetLoadBalancer(v CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetServiceType

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetServiceTypeDisplay

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetServiceTypeDisplay() string`

GetServiceTypeDisplay returns the ServiceTypeDisplay field if non-nil, zero value otherwise.

### GetServiceTypeDisplayOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetServiceTypeDisplayOk() (*string, bool)`

GetServiceTypeDisplayOk returns a tuple with the ServiceTypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypeDisplay

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetServiceTypeDisplay(v string)`

SetServiceTypeDisplay sets ServiceTypeDisplay field to given value.

### HasServiceTypeDisplay

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasServiceTypeDisplay() bool`

HasServiceTypeDisplay returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProxyType

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### SetProxyTypeNil

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetProxyTypeNil(b bool)`

 SetProxyTypeNil sets the value for ProxyType to be an explicit nil

### UnsetProxyType
`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetProxyType()`

UnsetProxyType ensures that no value is present for ProxyType, not even an explicit nil
### GetRedirectRewrite

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetRedirectRewrite() string`

GetRedirectRewrite returns the RedirectRewrite field if non-nil, zero value otherwise.

### GetRedirectRewriteOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetRedirectRewriteOk() (*string, bool)`

GetRedirectRewriteOk returns a tuple with the RedirectRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectRewrite

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetRedirectRewrite(v string)`

SetRedirectRewrite sets RedirectRewrite field to given value.

### HasRedirectRewrite

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasRedirectRewrite() bool`

HasRedirectRewrite returns a boolean if a field has been set.

### SetRedirectRewriteNil

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetRedirectRewriteNil(b bool)`

 SetRedirectRewriteNil sets the value for RedirectRewrite to be an explicit nil

### UnsetRedirectRewrite
`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetRedirectRewrite()`

UnsetRedirectRewrite ensures that no value is present for RedirectRewrite, not even an explicit nil
### GetPersistenceType

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceType() string`

GetPersistenceType returns the PersistenceType field if non-nil, zero value otherwise.

### GetPersistenceTypeOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceTypeOk() (*string, bool)`

GetPersistenceTypeOk returns a tuple with the PersistenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceType

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceType(v string)`

SetPersistenceType sets PersistenceType field to given value.

### HasPersistenceType

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasPersistenceType() bool`

HasPersistenceType returns a boolean if a field has been set.

### SetPersistenceTypeNil

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceTypeNil(b bool)`

 SetPersistenceTypeNil sets the value for PersistenceType to be an explicit nil

### UnsetPersistenceType
`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetPersistenceType()`

UnsetPersistenceType ensures that no value is present for PersistenceType, not even an explicit nil
### GetSslEnabled

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetSslEnabled() string`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetSslEnabledOk() (*string, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetSslEnabled(v string)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### SetSslEnabledNil

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetSslEnabledNil(b bool)`

 SetSslEnabledNil sets the value for SslEnabled to be an explicit nil

### UnsetSslEnabled
`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetSslEnabled()`

UnsetSslEnabled ensures that no value is present for SslEnabled, not even an explicit nil
### GetSslCert

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetAccountCertificate

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetAccountCertificate() string`

GetAccountCertificate returns the AccountCertificate field if non-nil, zero value otherwise.

### GetAccountCertificateOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetAccountCertificateOk() (*string, bool)`

GetAccountCertificateOk returns a tuple with the AccountCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCertificate

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetAccountCertificate(v string)`

SetAccountCertificate sets AccountCertificate field to given value.

### HasAccountCertificate

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasAccountCertificate() bool`

HasAccountCertificate returns a boolean if a field has been set.

### SetAccountCertificateNil

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetAccountCertificateNil(b bool)`

 SetAccountCertificateNil sets the value for AccountCertificate to be an explicit nil

### UnsetAccountCertificate
`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetAccountCertificate()`

UnsetAccountCertificate ensures that no value is present for AccountCertificate, not even an explicit nil
### GetEnabled

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetInsertXforwardedFor

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetInsertXforwardedFor() bool`

GetInsertXforwardedFor returns the InsertXforwardedFor field if non-nil, zero value otherwise.

### GetInsertXforwardedForOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetInsertXforwardedForOk() (*bool, bool)`

GetInsertXforwardedForOk returns a tuple with the InsertXforwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertXforwardedFor

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetInsertXforwardedFor(v bool)`

SetInsertXforwardedFor sets InsertXforwardedFor field to given value.

### HasInsertXforwardedFor

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasInsertXforwardedFor() bool`

HasInsertXforwardedFor returns a boolean if a field has been set.

### GetPersistenceCookieName

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceCookieName() string`

GetPersistenceCookieName returns the PersistenceCookieName field if non-nil, zero value otherwise.

### GetPersistenceCookieNameOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceCookieNameOk() (*string, bool)`

GetPersistenceCookieNameOk returns a tuple with the PersistenceCookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceCookieName

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceCookieName(v string)`

SetPersistenceCookieName sets PersistenceCookieName field to given value.

### HasPersistenceCookieName

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasPersistenceCookieName() bool`

HasPersistenceCookieName returns a boolean if a field has been set.

### SetPersistenceCookieNameNil

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceCookieNameNil(b bool)`

 SetPersistenceCookieNameNil sets the value for PersistenceCookieName to be an explicit nil

### UnsetPersistenceCookieName
`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetPersistenceCookieName()`

UnsetPersistenceCookieName ensures that no value is present for PersistenceCookieName, not even an explicit nil
### GetPersistenceExpiresIn

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceExpiresIn() string`

GetPersistenceExpiresIn returns the PersistenceExpiresIn field if non-nil, zero value otherwise.

### GetPersistenceExpiresInOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceExpiresInOk() (*string, bool)`

GetPersistenceExpiresInOk returns a tuple with the PersistenceExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceExpiresIn

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceExpiresIn(v string)`

SetPersistenceExpiresIn sets PersistenceExpiresIn field to given value.

### HasPersistenceExpiresIn

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasPersistenceExpiresIn() bool`

HasPersistenceExpiresIn returns a boolean if a field has been set.

### SetPersistenceExpiresInNil

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceExpiresInNil(b bool)`

 SetPersistenceExpiresInNil sets the value for PersistenceExpiresIn to be an explicit nil

### UnsetPersistenceExpiresIn
`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetPersistenceExpiresIn()`

UnsetPersistenceExpiresIn ensures that no value is present for PersistenceExpiresIn, not even an explicit nil
### GetEditable

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetConfig

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CreateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


