# UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileLoadBalancer**](UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileLoadBalancer.md) |  | [optional] 
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

### NewUpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile

`func NewUpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile() *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile`

NewUpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile instantiates a new UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetLoadBalancer() UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetLoadBalancerOk() (*UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetLoadBalancer(v UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfileLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetServiceType

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetServiceTypeDisplay

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetServiceTypeDisplay() string`

GetServiceTypeDisplay returns the ServiceTypeDisplay field if non-nil, zero value otherwise.

### GetServiceTypeDisplayOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetServiceTypeDisplayOk() (*string, bool)`

GetServiceTypeDisplayOk returns a tuple with the ServiceTypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypeDisplay

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetServiceTypeDisplay(v string)`

SetServiceTypeDisplay sets ServiceTypeDisplay field to given value.

### HasServiceTypeDisplay

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasServiceTypeDisplay() bool`

HasServiceTypeDisplay returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProxyType

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### SetProxyTypeNil

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetProxyTypeNil(b bool)`

 SetProxyTypeNil sets the value for ProxyType to be an explicit nil

### UnsetProxyType
`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetProxyType()`

UnsetProxyType ensures that no value is present for ProxyType, not even an explicit nil
### GetRedirectRewrite

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetRedirectRewrite() string`

GetRedirectRewrite returns the RedirectRewrite field if non-nil, zero value otherwise.

### GetRedirectRewriteOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetRedirectRewriteOk() (*string, bool)`

GetRedirectRewriteOk returns a tuple with the RedirectRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectRewrite

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetRedirectRewrite(v string)`

SetRedirectRewrite sets RedirectRewrite field to given value.

### HasRedirectRewrite

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasRedirectRewrite() bool`

HasRedirectRewrite returns a boolean if a field has been set.

### SetRedirectRewriteNil

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetRedirectRewriteNil(b bool)`

 SetRedirectRewriteNil sets the value for RedirectRewrite to be an explicit nil

### UnsetRedirectRewrite
`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetRedirectRewrite()`

UnsetRedirectRewrite ensures that no value is present for RedirectRewrite, not even an explicit nil
### GetPersistenceType

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceType() string`

GetPersistenceType returns the PersistenceType field if non-nil, zero value otherwise.

### GetPersistenceTypeOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceTypeOk() (*string, bool)`

GetPersistenceTypeOk returns a tuple with the PersistenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceType

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceType(v string)`

SetPersistenceType sets PersistenceType field to given value.

### HasPersistenceType

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasPersistenceType() bool`

HasPersistenceType returns a boolean if a field has been set.

### SetPersistenceTypeNil

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceTypeNil(b bool)`

 SetPersistenceTypeNil sets the value for PersistenceType to be an explicit nil

### UnsetPersistenceType
`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetPersistenceType()`

UnsetPersistenceType ensures that no value is present for PersistenceType, not even an explicit nil
### GetSslEnabled

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetSslEnabled() string`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetSslEnabledOk() (*string, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetSslEnabled(v string)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### SetSslEnabledNil

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetSslEnabledNil(b bool)`

 SetSslEnabledNil sets the value for SslEnabled to be an explicit nil

### UnsetSslEnabled
`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetSslEnabled()`

UnsetSslEnabled ensures that no value is present for SslEnabled, not even an explicit nil
### GetSslCert

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetAccountCertificate

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetAccountCertificate() string`

GetAccountCertificate returns the AccountCertificate field if non-nil, zero value otherwise.

### GetAccountCertificateOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetAccountCertificateOk() (*string, bool)`

GetAccountCertificateOk returns a tuple with the AccountCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCertificate

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetAccountCertificate(v string)`

SetAccountCertificate sets AccountCertificate field to given value.

### HasAccountCertificate

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasAccountCertificate() bool`

HasAccountCertificate returns a boolean if a field has been set.

### SetAccountCertificateNil

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetAccountCertificateNil(b bool)`

 SetAccountCertificateNil sets the value for AccountCertificate to be an explicit nil

### UnsetAccountCertificate
`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetAccountCertificate()`

UnsetAccountCertificate ensures that no value is present for AccountCertificate, not even an explicit nil
### GetEnabled

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetInsertXforwardedFor

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetInsertXforwardedFor() bool`

GetInsertXforwardedFor returns the InsertXforwardedFor field if non-nil, zero value otherwise.

### GetInsertXforwardedForOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetInsertXforwardedForOk() (*bool, bool)`

GetInsertXforwardedForOk returns a tuple with the InsertXforwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertXforwardedFor

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetInsertXforwardedFor(v bool)`

SetInsertXforwardedFor sets InsertXforwardedFor field to given value.

### HasInsertXforwardedFor

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasInsertXforwardedFor() bool`

HasInsertXforwardedFor returns a boolean if a field has been set.

### GetPersistenceCookieName

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceCookieName() string`

GetPersistenceCookieName returns the PersistenceCookieName field if non-nil, zero value otherwise.

### GetPersistenceCookieNameOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceCookieNameOk() (*string, bool)`

GetPersistenceCookieNameOk returns a tuple with the PersistenceCookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceCookieName

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceCookieName(v string)`

SetPersistenceCookieName sets PersistenceCookieName field to given value.

### HasPersistenceCookieName

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasPersistenceCookieName() bool`

HasPersistenceCookieName returns a boolean if a field has been set.

### SetPersistenceCookieNameNil

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceCookieNameNil(b bool)`

 SetPersistenceCookieNameNil sets the value for PersistenceCookieName to be an explicit nil

### UnsetPersistenceCookieName
`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetPersistenceCookieName()`

UnsetPersistenceCookieName ensures that no value is present for PersistenceCookieName, not even an explicit nil
### GetPersistenceExpiresIn

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceExpiresIn() string`

GetPersistenceExpiresIn returns the PersistenceExpiresIn field if non-nil, zero value otherwise.

### GetPersistenceExpiresInOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetPersistenceExpiresInOk() (*string, bool)`

GetPersistenceExpiresInOk returns a tuple with the PersistenceExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceExpiresIn

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceExpiresIn(v string)`

SetPersistenceExpiresIn sets PersistenceExpiresIn field to given value.

### HasPersistenceExpiresIn

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasPersistenceExpiresIn() bool`

HasPersistenceExpiresIn returns a boolean if a field has been set.

### SetPersistenceExpiresInNil

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetPersistenceExpiresInNil(b bool)`

 SetPersistenceExpiresInNil sets the value for PersistenceExpiresIn to be an explicit nil

### UnsetPersistenceExpiresIn
`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetPersistenceExpiresIn()`

UnsetPersistenceExpiresIn ensures that no value is present for PersistenceExpiresIn, not even an explicit nil
### GetEditable

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateLoadBalancerProfile200ResponseAllOfLoadBalancerProfile) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


