# GetLoadBalancerProfile200ResponseLoadBalancerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**GetLoadBalancerProfile200ResponseLoadBalancerProfileLoadBalancer**](GetLoadBalancerProfile200ResponseLoadBalancerProfileLoadBalancer.md) |  | [optional] 
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

### NewGetLoadBalancerProfile200ResponseLoadBalancerProfile

`func NewGetLoadBalancerProfile200ResponseLoadBalancerProfile() *GetLoadBalancerProfile200ResponseLoadBalancerProfile`

NewGetLoadBalancerProfile200ResponseLoadBalancerProfile instantiates a new GetLoadBalancerProfile200ResponseLoadBalancerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoadBalancerProfile200ResponseLoadBalancerProfileWithDefaults

`func NewGetLoadBalancerProfile200ResponseLoadBalancerProfileWithDefaults() *GetLoadBalancerProfile200ResponseLoadBalancerProfile`

NewGetLoadBalancerProfile200ResponseLoadBalancerProfileWithDefaults instantiates a new GetLoadBalancerProfile200ResponseLoadBalancerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetLoadBalancer() GetLoadBalancerProfile200ResponseLoadBalancerProfileLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetLoadBalancerOk() (*GetLoadBalancerProfile200ResponseLoadBalancerProfileLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetLoadBalancer(v GetLoadBalancerProfile200ResponseLoadBalancerProfileLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetServiceType

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetServiceTypeDisplay

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetServiceTypeDisplay() string`

GetServiceTypeDisplay returns the ServiceTypeDisplay field if non-nil, zero value otherwise.

### GetServiceTypeDisplayOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetServiceTypeDisplayOk() (*string, bool)`

GetServiceTypeDisplayOk returns a tuple with the ServiceTypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypeDisplay

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetServiceTypeDisplay(v string)`

SetServiceTypeDisplay sets ServiceTypeDisplay field to given value.

### HasServiceTypeDisplay

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasServiceTypeDisplay() bool`

HasServiceTypeDisplay returns a boolean if a field has been set.

### GetVisibility

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProxyType

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### SetProxyTypeNil

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetProxyTypeNil(b bool)`

 SetProxyTypeNil sets the value for ProxyType to be an explicit nil

### UnsetProxyType
`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) UnsetProxyType()`

UnsetProxyType ensures that no value is present for ProxyType, not even an explicit nil
### GetRedirectRewrite

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetRedirectRewrite() string`

GetRedirectRewrite returns the RedirectRewrite field if non-nil, zero value otherwise.

### GetRedirectRewriteOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetRedirectRewriteOk() (*string, bool)`

GetRedirectRewriteOk returns a tuple with the RedirectRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectRewrite

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetRedirectRewrite(v string)`

SetRedirectRewrite sets RedirectRewrite field to given value.

### HasRedirectRewrite

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasRedirectRewrite() bool`

HasRedirectRewrite returns a boolean if a field has been set.

### SetRedirectRewriteNil

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetRedirectRewriteNil(b bool)`

 SetRedirectRewriteNil sets the value for RedirectRewrite to be an explicit nil

### UnsetRedirectRewrite
`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) UnsetRedirectRewrite()`

UnsetRedirectRewrite ensures that no value is present for RedirectRewrite, not even an explicit nil
### GetPersistenceType

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetPersistenceType() string`

GetPersistenceType returns the PersistenceType field if non-nil, zero value otherwise.

### GetPersistenceTypeOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetPersistenceTypeOk() (*string, bool)`

GetPersistenceTypeOk returns a tuple with the PersistenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceType

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetPersistenceType(v string)`

SetPersistenceType sets PersistenceType field to given value.

### HasPersistenceType

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasPersistenceType() bool`

HasPersistenceType returns a boolean if a field has been set.

### SetPersistenceTypeNil

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetPersistenceTypeNil(b bool)`

 SetPersistenceTypeNil sets the value for PersistenceType to be an explicit nil

### UnsetPersistenceType
`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) UnsetPersistenceType()`

UnsetPersistenceType ensures that no value is present for PersistenceType, not even an explicit nil
### GetSslEnabled

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetSslEnabled() string`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetSslEnabledOk() (*string, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetSslEnabled(v string)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### SetSslEnabledNil

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetSslEnabledNil(b bool)`

 SetSslEnabledNil sets the value for SslEnabled to be an explicit nil

### UnsetSslEnabled
`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) UnsetSslEnabled()`

UnsetSslEnabled ensures that no value is present for SslEnabled, not even an explicit nil
### GetSslCert

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### SetSslCertNil

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetSslCertNil(b bool)`

 SetSslCertNil sets the value for SslCert to be an explicit nil

### UnsetSslCert
`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) UnsetSslCert()`

UnsetSslCert ensures that no value is present for SslCert, not even an explicit nil
### GetAccountCertificate

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetAccountCertificate() string`

GetAccountCertificate returns the AccountCertificate field if non-nil, zero value otherwise.

### GetAccountCertificateOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetAccountCertificateOk() (*string, bool)`

GetAccountCertificateOk returns a tuple with the AccountCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCertificate

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetAccountCertificate(v string)`

SetAccountCertificate sets AccountCertificate field to given value.

### HasAccountCertificate

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasAccountCertificate() bool`

HasAccountCertificate returns a boolean if a field has been set.

### SetAccountCertificateNil

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetAccountCertificateNil(b bool)`

 SetAccountCertificateNil sets the value for AccountCertificate to be an explicit nil

### UnsetAccountCertificate
`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) UnsetAccountCertificate()`

UnsetAccountCertificate ensures that no value is present for AccountCertificate, not even an explicit nil
### GetEnabled

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetInsertXforwardedFor

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetInsertXforwardedFor() bool`

GetInsertXforwardedFor returns the InsertXforwardedFor field if non-nil, zero value otherwise.

### GetInsertXforwardedForOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetInsertXforwardedForOk() (*bool, bool)`

GetInsertXforwardedForOk returns a tuple with the InsertXforwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertXforwardedFor

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetInsertXforwardedFor(v bool)`

SetInsertXforwardedFor sets InsertXforwardedFor field to given value.

### HasInsertXforwardedFor

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasInsertXforwardedFor() bool`

HasInsertXforwardedFor returns a boolean if a field has been set.

### GetPersistenceCookieName

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetPersistenceCookieName() string`

GetPersistenceCookieName returns the PersistenceCookieName field if non-nil, zero value otherwise.

### GetPersistenceCookieNameOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetPersistenceCookieNameOk() (*string, bool)`

GetPersistenceCookieNameOk returns a tuple with the PersistenceCookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceCookieName

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetPersistenceCookieName(v string)`

SetPersistenceCookieName sets PersistenceCookieName field to given value.

### HasPersistenceCookieName

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasPersistenceCookieName() bool`

HasPersistenceCookieName returns a boolean if a field has been set.

### SetPersistenceCookieNameNil

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetPersistenceCookieNameNil(b bool)`

 SetPersistenceCookieNameNil sets the value for PersistenceCookieName to be an explicit nil

### UnsetPersistenceCookieName
`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) UnsetPersistenceCookieName()`

UnsetPersistenceCookieName ensures that no value is present for PersistenceCookieName, not even an explicit nil
### GetPersistenceExpiresIn

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetPersistenceExpiresIn() string`

GetPersistenceExpiresIn returns the PersistenceExpiresIn field if non-nil, zero value otherwise.

### GetPersistenceExpiresInOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetPersistenceExpiresInOk() (*string, bool)`

GetPersistenceExpiresInOk returns a tuple with the PersistenceExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceExpiresIn

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetPersistenceExpiresIn(v string)`

SetPersistenceExpiresIn sets PersistenceExpiresIn field to given value.

### HasPersistenceExpiresIn

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasPersistenceExpiresIn() bool`

HasPersistenceExpiresIn returns a boolean if a field has been set.

### SetPersistenceExpiresInNil

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetPersistenceExpiresInNil(b bool)`

 SetPersistenceExpiresInNil sets the value for PersistenceExpiresIn to be an explicit nil

### UnsetPersistenceExpiresIn
`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) UnsetPersistenceExpiresIn()`

UnsetPersistenceExpiresIn ensures that no value is present for PersistenceExpiresIn, not even an explicit nil
### GetEditable

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetConfig

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetLoadBalancerProfile200ResponseLoadBalancerProfile) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


