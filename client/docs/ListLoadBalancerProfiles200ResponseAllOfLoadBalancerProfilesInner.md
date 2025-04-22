# ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **string** |  | [optional] 
**ServiceTypeDisplay** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ProxyType** | Pointer to **string** |  | [optional] 
**RedirectRewrite** | Pointer to **string** |  | [optional] 
**PersistenceType** | Pointer to **string** |  | [optional] 
**SslEnabled** | Pointer to **string** |  | [optional] 
**SslCert** | Pointer to **string** |  | [optional] 
**AccountCertificate** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**InsertXforwardedFor** | Pointer to **bool** |  | [optional] 
**PersistenceCookieName** | Pointer to **string** |  | [optional] 
**PersistenceExpiresIn** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner

`func NewListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner() *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner`

NewListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner instantiates a new ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInnerWithDefaults

`func NewListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInnerWithDefaults() *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner`

NewListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInnerWithDefaults instantiates a new ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetLoadBalancer() ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetLoadBalancerOk() (*ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetLoadBalancer(v ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetServiceType

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetServiceTypeDisplay

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetServiceTypeDisplay() string`

GetServiceTypeDisplay returns the ServiceTypeDisplay field if non-nil, zero value otherwise.

### GetServiceTypeDisplayOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetServiceTypeDisplayOk() (*string, bool)`

GetServiceTypeDisplayOk returns a tuple with the ServiceTypeDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypeDisplay

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetServiceTypeDisplay(v string)`

SetServiceTypeDisplay sets ServiceTypeDisplay field to given value.

### HasServiceTypeDisplay

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasServiceTypeDisplay() bool`

HasServiceTypeDisplay returns a boolean if a field has been set.

### GetVisibility

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProxyType

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetRedirectRewrite

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetRedirectRewrite() string`

GetRedirectRewrite returns the RedirectRewrite field if non-nil, zero value otherwise.

### GetRedirectRewriteOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetRedirectRewriteOk() (*string, bool)`

GetRedirectRewriteOk returns a tuple with the RedirectRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectRewrite

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetRedirectRewrite(v string)`

SetRedirectRewrite sets RedirectRewrite field to given value.

### HasRedirectRewrite

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasRedirectRewrite() bool`

HasRedirectRewrite returns a boolean if a field has been set.

### GetPersistenceType

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetPersistenceType() string`

GetPersistenceType returns the PersistenceType field if non-nil, zero value otherwise.

### GetPersistenceTypeOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetPersistenceTypeOk() (*string, bool)`

GetPersistenceTypeOk returns a tuple with the PersistenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceType

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetPersistenceType(v string)`

SetPersistenceType sets PersistenceType field to given value.

### HasPersistenceType

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasPersistenceType() bool`

HasPersistenceType returns a boolean if a field has been set.

### GetSslEnabled

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetSslEnabled() string`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetSslEnabledOk() (*string, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetSslEnabled(v string)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### GetSslCert

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetSslCert() string`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetSslCertOk() (*string, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetSslCert(v string)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### GetAccountCertificate

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetAccountCertificate() string`

GetAccountCertificate returns the AccountCertificate field if non-nil, zero value otherwise.

### GetAccountCertificateOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetAccountCertificateOk() (*string, bool)`

GetAccountCertificateOk returns a tuple with the AccountCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCertificate

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetAccountCertificate(v string)`

SetAccountCertificate sets AccountCertificate field to given value.

### HasAccountCertificate

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasAccountCertificate() bool`

HasAccountCertificate returns a boolean if a field has been set.

### GetEnabled

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetInsertXforwardedFor

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetInsertXforwardedFor() bool`

GetInsertXforwardedFor returns the InsertXforwardedFor field if non-nil, zero value otherwise.

### GetInsertXforwardedForOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetInsertXforwardedForOk() (*bool, bool)`

GetInsertXforwardedForOk returns a tuple with the InsertXforwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertXforwardedFor

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetInsertXforwardedFor(v bool)`

SetInsertXforwardedFor sets InsertXforwardedFor field to given value.

### HasInsertXforwardedFor

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasInsertXforwardedFor() bool`

HasInsertXforwardedFor returns a boolean if a field has been set.

### GetPersistenceCookieName

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetPersistenceCookieName() string`

GetPersistenceCookieName returns the PersistenceCookieName field if non-nil, zero value otherwise.

### GetPersistenceCookieNameOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetPersistenceCookieNameOk() (*string, bool)`

GetPersistenceCookieNameOk returns a tuple with the PersistenceCookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceCookieName

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetPersistenceCookieName(v string)`

SetPersistenceCookieName sets PersistenceCookieName field to given value.

### HasPersistenceCookieName

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasPersistenceCookieName() bool`

HasPersistenceCookieName returns a boolean if a field has been set.

### GetPersistenceExpiresIn

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetPersistenceExpiresIn() string`

GetPersistenceExpiresIn returns the PersistenceExpiresIn field if non-nil, zero value otherwise.

### GetPersistenceExpiresInOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetPersistenceExpiresInOk() (*string, bool)`

GetPersistenceExpiresInOk returns a tuple with the PersistenceExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceExpiresIn

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetPersistenceExpiresIn(v string)`

SetPersistenceExpiresIn sets PersistenceExpiresIn field to given value.

### HasPersistenceExpiresIn

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasPersistenceExpiresIn() bool`

HasPersistenceExpiresIn returns a boolean if a field has been set.

### GetEditable

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetConfig

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


