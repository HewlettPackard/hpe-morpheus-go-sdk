# UpdateRoleRequestRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to **string** | Authority (Name) | [optional] 
**Description** | Pointer to **NullableString** | Description | [optional] 
**LandingUrl** | Pointer to **NullableString** | An optional override for the default landing page after login for a user. | [optional] 
**Multitenant** | Pointer to **bool** | Multitenant roles are copied to all tenant accounts and kept in sync until a sub-tenant user modifies their copy of the role. *Only available to master tenant, only applies to user roles* | [optional] 
**MultitenantLocked** | Pointer to **bool** | Multitenant Locked, prevents sub-tenant users from modifying their copy of multitenant roles. *Only available to master tenant, only applies to user roles* | [optional] 
**DefaultPersona** | Pointer to **NullableString** |  | [optional] 
**FeaturePermissions** | Pointer to [**[]AddRolesRequestRoleFeaturePermissionsInner**](AddRolesRequestRoleFeaturePermissionsInner.md) | Set the access level for the specified permissions. | [optional] 
**GlobalSiteAccess** | Pointer to **string** | Set the default access level for for groups (sites). Only applies to user roles. | [optional] 
**Sites** | Pointer to [**[]AddRolesRequestRoleSitesInner**](AddRolesRequestRoleSitesInner.md) | Set the access level for the specified groups (sites). Only applies to user roles. | [optional] 
**GlobalZoneAccess** | Pointer to **string** | Set the default access level for for clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**Zones** | Pointer to [**[]AddRolesRequestRoleZonesInner**](AddRolesRequestRoleZonesInner.md) | Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**GlobalInstanceTypeAccess** | Pointer to **string** | Set the default access level for for instance types | [optional] 
**InstanceTypePermissions** | Pointer to [**[]AddRolesRequestRoleInstanceTypePermissionsInner**](AddRolesRequestRoleInstanceTypePermissionsInner.md) | Set the access level for the specified instance types | [optional] 
**GlobalAppTemplateAccess** | Pointer to **string** | Set the default access level for blueprints | [optional] 
**AppTemplatePermissions** | Pointer to [**[]AddRolesRequestRoleAppTemplatePermissionsInner**](AddRolesRequestRoleAppTemplatePermissionsInner.md) | Set the access level for the specified blueprints (appTemplates) | [optional] 
**GlobalCatalogItemTypeAccess** | Pointer to **string** | Set the default access level for catalog item types | [optional] 
**CatalogItemTypePermissions** | Pointer to [**[]AddRolesRequestRoleCatalogItemTypePermissionsInner**](AddRolesRequestRoleCatalogItemTypePermissionsInner.md) | Set the access level for the specified catalog item types | [optional] 
**GlobalPersonaAccess** | Pointer to **string** | Set the default access level for personas | [optional] 
**PersonaPermissions** | Pointer to [**[]AddRolesRequestRolePersonaPermissionsInner**](AddRolesRequestRolePersonaPermissionsInner.md) | Set the access level for the specified personas | [optional] 
**GlobalVdiPoolAccess** | Pointer to **string** | Set the default access level for VDI pools | [optional] 
**VdiPoolPermissions** | Pointer to [**[]AddRolesRequestRoleVdiPoolPermissionsInner**](AddRolesRequestRoleVdiPoolPermissionsInner.md) | Set the access level for the specified VDI pools | [optional] 
**GlobalReportTypeAccess** | Pointer to **string** | Set the default access level for report types | [optional] 
**ReportTypePermissions** | Pointer to [**[]AddRolesRequestRoleReportTypePermissionsInner**](AddRolesRequestRoleReportTypePermissionsInner.md) | Set the access level for the specified report types | [optional] 
**GlobalTaskAccess** | Pointer to **string** | Set the default access level for tasks | [optional] 
**TaskPermissions** | Pointer to [**[]AddRolesRequestRoleTaskPermissionsInner**](AddRolesRequestRoleTaskPermissionsInner.md) | Set the access level for the specified tasks | [optional] 
**GlobalTaskSetAccess** | Pointer to **string** | Set the default access level for workflows (taskSets) | [optional] 
**TaskSetPermissions** | Pointer to [**[]AddRolesRequestRoleTaskSetPermissionsInner**](AddRolesRequestRoleTaskSetPermissionsInner.md) | Set the access level for the specified workflows (taskSets) | [optional] 

## Methods

### NewUpdateRoleRequestRole

`func NewUpdateRoleRequestRole() *UpdateRoleRequestRole`

NewUpdateRoleRequestRole instantiates a new UpdateRoleRequestRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleRequestRoleWithDefaults

`func NewUpdateRoleRequestRoleWithDefaults() *UpdateRoleRequestRole`

NewUpdateRoleRequestRoleWithDefaults instantiates a new UpdateRoleRequestRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *UpdateRoleRequestRole) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *UpdateRoleRequestRole) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *UpdateRoleRequestRole) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *UpdateRoleRequestRole) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateRoleRequestRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRoleRequestRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRoleRequestRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateRoleRequestRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateRoleRequestRole) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateRoleRequestRole) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLandingUrl

`func (o *UpdateRoleRequestRole) GetLandingUrl() string`

GetLandingUrl returns the LandingUrl field if non-nil, zero value otherwise.

### GetLandingUrlOk

`func (o *UpdateRoleRequestRole) GetLandingUrlOk() (*string, bool)`

GetLandingUrlOk returns a tuple with the LandingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrl

`func (o *UpdateRoleRequestRole) SetLandingUrl(v string)`

SetLandingUrl sets LandingUrl field to given value.

### HasLandingUrl

`func (o *UpdateRoleRequestRole) HasLandingUrl() bool`

HasLandingUrl returns a boolean if a field has been set.

### SetLandingUrlNil

`func (o *UpdateRoleRequestRole) SetLandingUrlNil(b bool)`

 SetLandingUrlNil sets the value for LandingUrl to be an explicit nil

### UnsetLandingUrl
`func (o *UpdateRoleRequestRole) UnsetLandingUrl()`

UnsetLandingUrl ensures that no value is present for LandingUrl, not even an explicit nil
### GetMultitenant

`func (o *UpdateRoleRequestRole) GetMultitenant() bool`

GetMultitenant returns the Multitenant field if non-nil, zero value otherwise.

### GetMultitenantOk

`func (o *UpdateRoleRequestRole) GetMultitenantOk() (*bool, bool)`

GetMultitenantOk returns a tuple with the Multitenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenant

`func (o *UpdateRoleRequestRole) SetMultitenant(v bool)`

SetMultitenant sets Multitenant field to given value.

### HasMultitenant

`func (o *UpdateRoleRequestRole) HasMultitenant() bool`

HasMultitenant returns a boolean if a field has been set.

### GetMultitenantLocked

`func (o *UpdateRoleRequestRole) GetMultitenantLocked() bool`

GetMultitenantLocked returns the MultitenantLocked field if non-nil, zero value otherwise.

### GetMultitenantLockedOk

`func (o *UpdateRoleRequestRole) GetMultitenantLockedOk() (*bool, bool)`

GetMultitenantLockedOk returns a tuple with the MultitenantLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenantLocked

`func (o *UpdateRoleRequestRole) SetMultitenantLocked(v bool)`

SetMultitenantLocked sets MultitenantLocked field to given value.

### HasMultitenantLocked

`func (o *UpdateRoleRequestRole) HasMultitenantLocked() bool`

HasMultitenantLocked returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *UpdateRoleRequestRole) GetDefaultPersona() string`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *UpdateRoleRequestRole) GetDefaultPersonaOk() (*string, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *UpdateRoleRequestRole) SetDefaultPersona(v string)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *UpdateRoleRequestRole) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### SetDefaultPersonaNil

`func (o *UpdateRoleRequestRole) SetDefaultPersonaNil(b bool)`

 SetDefaultPersonaNil sets the value for DefaultPersona to be an explicit nil

### UnsetDefaultPersona
`func (o *UpdateRoleRequestRole) UnsetDefaultPersona()`

UnsetDefaultPersona ensures that no value is present for DefaultPersona, not even an explicit nil
### GetFeaturePermissions

`func (o *UpdateRoleRequestRole) GetFeaturePermissions() []AddRolesRequestRoleFeaturePermissionsInner`

GetFeaturePermissions returns the FeaturePermissions field if non-nil, zero value otherwise.

### GetFeaturePermissionsOk

`func (o *UpdateRoleRequestRole) GetFeaturePermissionsOk() (*[]AddRolesRequestRoleFeaturePermissionsInner, bool)`

GetFeaturePermissionsOk returns a tuple with the FeaturePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePermissions

`func (o *UpdateRoleRequestRole) SetFeaturePermissions(v []AddRolesRequestRoleFeaturePermissionsInner)`

SetFeaturePermissions sets FeaturePermissions field to given value.

### HasFeaturePermissions

`func (o *UpdateRoleRequestRole) HasFeaturePermissions() bool`

HasFeaturePermissions returns a boolean if a field has been set.

### GetGlobalSiteAccess

`func (o *UpdateRoleRequestRole) GetGlobalSiteAccess() string`

GetGlobalSiteAccess returns the GlobalSiteAccess field if non-nil, zero value otherwise.

### GetGlobalSiteAccessOk

`func (o *UpdateRoleRequestRole) GetGlobalSiteAccessOk() (*string, bool)`

GetGlobalSiteAccessOk returns a tuple with the GlobalSiteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSiteAccess

`func (o *UpdateRoleRequestRole) SetGlobalSiteAccess(v string)`

SetGlobalSiteAccess sets GlobalSiteAccess field to given value.

### HasGlobalSiteAccess

`func (o *UpdateRoleRequestRole) HasGlobalSiteAccess() bool`

HasGlobalSiteAccess returns a boolean if a field has been set.

### GetSites

`func (o *UpdateRoleRequestRole) GetSites() []AddRolesRequestRoleSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateRoleRequestRole) GetSitesOk() (*[]AddRolesRequestRoleSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateRoleRequestRole) SetSites(v []AddRolesRequestRoleSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateRoleRequestRole) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetGlobalZoneAccess

`func (o *UpdateRoleRequestRole) GetGlobalZoneAccess() string`

GetGlobalZoneAccess returns the GlobalZoneAccess field if non-nil, zero value otherwise.

### GetGlobalZoneAccessOk

`func (o *UpdateRoleRequestRole) GetGlobalZoneAccessOk() (*string, bool)`

GetGlobalZoneAccessOk returns a tuple with the GlobalZoneAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalZoneAccess

`func (o *UpdateRoleRequestRole) SetGlobalZoneAccess(v string)`

SetGlobalZoneAccess sets GlobalZoneAccess field to given value.

### HasGlobalZoneAccess

`func (o *UpdateRoleRequestRole) HasGlobalZoneAccess() bool`

HasGlobalZoneAccess returns a boolean if a field has been set.

### GetZones

`func (o *UpdateRoleRequestRole) GetZones() []AddRolesRequestRoleZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *UpdateRoleRequestRole) GetZonesOk() (*[]AddRolesRequestRoleZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *UpdateRoleRequestRole) SetZones(v []AddRolesRequestRoleZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *UpdateRoleRequestRole) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetGlobalInstanceTypeAccess

`func (o *UpdateRoleRequestRole) GetGlobalInstanceTypeAccess() string`

GetGlobalInstanceTypeAccess returns the GlobalInstanceTypeAccess field if non-nil, zero value otherwise.

### GetGlobalInstanceTypeAccessOk

`func (o *UpdateRoleRequestRole) GetGlobalInstanceTypeAccessOk() (*string, bool)`

GetGlobalInstanceTypeAccessOk returns a tuple with the GlobalInstanceTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalInstanceTypeAccess

`func (o *UpdateRoleRequestRole) SetGlobalInstanceTypeAccess(v string)`

SetGlobalInstanceTypeAccess sets GlobalInstanceTypeAccess field to given value.

### HasGlobalInstanceTypeAccess

`func (o *UpdateRoleRequestRole) HasGlobalInstanceTypeAccess() bool`

HasGlobalInstanceTypeAccess returns a boolean if a field has been set.

### GetInstanceTypePermissions

`func (o *UpdateRoleRequestRole) GetInstanceTypePermissions() []AddRolesRequestRoleInstanceTypePermissionsInner`

GetInstanceTypePermissions returns the InstanceTypePermissions field if non-nil, zero value otherwise.

### GetInstanceTypePermissionsOk

`func (o *UpdateRoleRequestRole) GetInstanceTypePermissionsOk() (*[]AddRolesRequestRoleInstanceTypePermissionsInner, bool)`

GetInstanceTypePermissionsOk returns a tuple with the InstanceTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypePermissions

`func (o *UpdateRoleRequestRole) SetInstanceTypePermissions(v []AddRolesRequestRoleInstanceTypePermissionsInner)`

SetInstanceTypePermissions sets InstanceTypePermissions field to given value.

### HasInstanceTypePermissions

`func (o *UpdateRoleRequestRole) HasInstanceTypePermissions() bool`

HasInstanceTypePermissions returns a boolean if a field has been set.

### GetGlobalAppTemplateAccess

`func (o *UpdateRoleRequestRole) GetGlobalAppTemplateAccess() string`

GetGlobalAppTemplateAccess returns the GlobalAppTemplateAccess field if non-nil, zero value otherwise.

### GetGlobalAppTemplateAccessOk

`func (o *UpdateRoleRequestRole) GetGlobalAppTemplateAccessOk() (*string, bool)`

GetGlobalAppTemplateAccessOk returns a tuple with the GlobalAppTemplateAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAppTemplateAccess

`func (o *UpdateRoleRequestRole) SetGlobalAppTemplateAccess(v string)`

SetGlobalAppTemplateAccess sets GlobalAppTemplateAccess field to given value.

### HasGlobalAppTemplateAccess

`func (o *UpdateRoleRequestRole) HasGlobalAppTemplateAccess() bool`

HasGlobalAppTemplateAccess returns a boolean if a field has been set.

### GetAppTemplatePermissions

`func (o *UpdateRoleRequestRole) GetAppTemplatePermissions() []AddRolesRequestRoleAppTemplatePermissionsInner`

GetAppTemplatePermissions returns the AppTemplatePermissions field if non-nil, zero value otherwise.

### GetAppTemplatePermissionsOk

`func (o *UpdateRoleRequestRole) GetAppTemplatePermissionsOk() (*[]AddRolesRequestRoleAppTemplatePermissionsInner, bool)`

GetAppTemplatePermissionsOk returns a tuple with the AppTemplatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplatePermissions

`func (o *UpdateRoleRequestRole) SetAppTemplatePermissions(v []AddRolesRequestRoleAppTemplatePermissionsInner)`

SetAppTemplatePermissions sets AppTemplatePermissions field to given value.

### HasAppTemplatePermissions

`func (o *UpdateRoleRequestRole) HasAppTemplatePermissions() bool`

HasAppTemplatePermissions returns a boolean if a field has been set.

### GetGlobalCatalogItemTypeAccess

`func (o *UpdateRoleRequestRole) GetGlobalCatalogItemTypeAccess() string`

GetGlobalCatalogItemTypeAccess returns the GlobalCatalogItemTypeAccess field if non-nil, zero value otherwise.

### GetGlobalCatalogItemTypeAccessOk

`func (o *UpdateRoleRequestRole) GetGlobalCatalogItemTypeAccessOk() (*string, bool)`

GetGlobalCatalogItemTypeAccessOk returns a tuple with the GlobalCatalogItemTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCatalogItemTypeAccess

`func (o *UpdateRoleRequestRole) SetGlobalCatalogItemTypeAccess(v string)`

SetGlobalCatalogItemTypeAccess sets GlobalCatalogItemTypeAccess field to given value.

### HasGlobalCatalogItemTypeAccess

`func (o *UpdateRoleRequestRole) HasGlobalCatalogItemTypeAccess() bool`

HasGlobalCatalogItemTypeAccess returns a boolean if a field has been set.

### GetCatalogItemTypePermissions

`func (o *UpdateRoleRequestRole) GetCatalogItemTypePermissions() []AddRolesRequestRoleCatalogItemTypePermissionsInner`

GetCatalogItemTypePermissions returns the CatalogItemTypePermissions field if non-nil, zero value otherwise.

### GetCatalogItemTypePermissionsOk

`func (o *UpdateRoleRequestRole) GetCatalogItemTypePermissionsOk() (*[]AddRolesRequestRoleCatalogItemTypePermissionsInner, bool)`

GetCatalogItemTypePermissionsOk returns a tuple with the CatalogItemTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypePermissions

`func (o *UpdateRoleRequestRole) SetCatalogItemTypePermissions(v []AddRolesRequestRoleCatalogItemTypePermissionsInner)`

SetCatalogItemTypePermissions sets CatalogItemTypePermissions field to given value.

### HasCatalogItemTypePermissions

`func (o *UpdateRoleRequestRole) HasCatalogItemTypePermissions() bool`

HasCatalogItemTypePermissions returns a boolean if a field has been set.

### GetGlobalPersonaAccess

`func (o *UpdateRoleRequestRole) GetGlobalPersonaAccess() string`

GetGlobalPersonaAccess returns the GlobalPersonaAccess field if non-nil, zero value otherwise.

### GetGlobalPersonaAccessOk

`func (o *UpdateRoleRequestRole) GetGlobalPersonaAccessOk() (*string, bool)`

GetGlobalPersonaAccessOk returns a tuple with the GlobalPersonaAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPersonaAccess

`func (o *UpdateRoleRequestRole) SetGlobalPersonaAccess(v string)`

SetGlobalPersonaAccess sets GlobalPersonaAccess field to given value.

### HasGlobalPersonaAccess

`func (o *UpdateRoleRequestRole) HasGlobalPersonaAccess() bool`

HasGlobalPersonaAccess returns a boolean if a field has been set.

### GetPersonaPermissions

`func (o *UpdateRoleRequestRole) GetPersonaPermissions() []AddRolesRequestRolePersonaPermissionsInner`

GetPersonaPermissions returns the PersonaPermissions field if non-nil, zero value otherwise.

### GetPersonaPermissionsOk

`func (o *UpdateRoleRequestRole) GetPersonaPermissionsOk() (*[]AddRolesRequestRolePersonaPermissionsInner, bool)`

GetPersonaPermissionsOk returns a tuple with the PersonaPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaPermissions

`func (o *UpdateRoleRequestRole) SetPersonaPermissions(v []AddRolesRequestRolePersonaPermissionsInner)`

SetPersonaPermissions sets PersonaPermissions field to given value.

### HasPersonaPermissions

`func (o *UpdateRoleRequestRole) HasPersonaPermissions() bool`

HasPersonaPermissions returns a boolean if a field has been set.

### GetGlobalVdiPoolAccess

`func (o *UpdateRoleRequestRole) GetGlobalVdiPoolAccess() string`

GetGlobalVdiPoolAccess returns the GlobalVdiPoolAccess field if non-nil, zero value otherwise.

### GetGlobalVdiPoolAccessOk

`func (o *UpdateRoleRequestRole) GetGlobalVdiPoolAccessOk() (*string, bool)`

GetGlobalVdiPoolAccessOk returns a tuple with the GlobalVdiPoolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalVdiPoolAccess

`func (o *UpdateRoleRequestRole) SetGlobalVdiPoolAccess(v string)`

SetGlobalVdiPoolAccess sets GlobalVdiPoolAccess field to given value.

### HasGlobalVdiPoolAccess

`func (o *UpdateRoleRequestRole) HasGlobalVdiPoolAccess() bool`

HasGlobalVdiPoolAccess returns a boolean if a field has been set.

### GetVdiPoolPermissions

`func (o *UpdateRoleRequestRole) GetVdiPoolPermissions() []AddRolesRequestRoleVdiPoolPermissionsInner`

GetVdiPoolPermissions returns the VdiPoolPermissions field if non-nil, zero value otherwise.

### GetVdiPoolPermissionsOk

`func (o *UpdateRoleRequestRole) GetVdiPoolPermissionsOk() (*[]AddRolesRequestRoleVdiPoolPermissionsInner, bool)`

GetVdiPoolPermissionsOk returns a tuple with the VdiPoolPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPoolPermissions

`func (o *UpdateRoleRequestRole) SetVdiPoolPermissions(v []AddRolesRequestRoleVdiPoolPermissionsInner)`

SetVdiPoolPermissions sets VdiPoolPermissions field to given value.

### HasVdiPoolPermissions

`func (o *UpdateRoleRequestRole) HasVdiPoolPermissions() bool`

HasVdiPoolPermissions returns a boolean if a field has been set.

### GetGlobalReportTypeAccess

`func (o *UpdateRoleRequestRole) GetGlobalReportTypeAccess() string`

GetGlobalReportTypeAccess returns the GlobalReportTypeAccess field if non-nil, zero value otherwise.

### GetGlobalReportTypeAccessOk

`func (o *UpdateRoleRequestRole) GetGlobalReportTypeAccessOk() (*string, bool)`

GetGlobalReportTypeAccessOk returns a tuple with the GlobalReportTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalReportTypeAccess

`func (o *UpdateRoleRequestRole) SetGlobalReportTypeAccess(v string)`

SetGlobalReportTypeAccess sets GlobalReportTypeAccess field to given value.

### HasGlobalReportTypeAccess

`func (o *UpdateRoleRequestRole) HasGlobalReportTypeAccess() bool`

HasGlobalReportTypeAccess returns a boolean if a field has been set.

### GetReportTypePermissions

`func (o *UpdateRoleRequestRole) GetReportTypePermissions() []AddRolesRequestRoleReportTypePermissionsInner`

GetReportTypePermissions returns the ReportTypePermissions field if non-nil, zero value otherwise.

### GetReportTypePermissionsOk

`func (o *UpdateRoleRequestRole) GetReportTypePermissionsOk() (*[]AddRolesRequestRoleReportTypePermissionsInner, bool)`

GetReportTypePermissionsOk returns a tuple with the ReportTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypePermissions

`func (o *UpdateRoleRequestRole) SetReportTypePermissions(v []AddRolesRequestRoleReportTypePermissionsInner)`

SetReportTypePermissions sets ReportTypePermissions field to given value.

### HasReportTypePermissions

`func (o *UpdateRoleRequestRole) HasReportTypePermissions() bool`

HasReportTypePermissions returns a boolean if a field has been set.

### GetGlobalTaskAccess

`func (o *UpdateRoleRequestRole) GetGlobalTaskAccess() string`

GetGlobalTaskAccess returns the GlobalTaskAccess field if non-nil, zero value otherwise.

### GetGlobalTaskAccessOk

`func (o *UpdateRoleRequestRole) GetGlobalTaskAccessOk() (*string, bool)`

GetGlobalTaskAccessOk returns a tuple with the GlobalTaskAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskAccess

`func (o *UpdateRoleRequestRole) SetGlobalTaskAccess(v string)`

SetGlobalTaskAccess sets GlobalTaskAccess field to given value.

### HasGlobalTaskAccess

`func (o *UpdateRoleRequestRole) HasGlobalTaskAccess() bool`

HasGlobalTaskAccess returns a boolean if a field has been set.

### GetTaskPermissions

`func (o *UpdateRoleRequestRole) GetTaskPermissions() []AddRolesRequestRoleTaskPermissionsInner`

GetTaskPermissions returns the TaskPermissions field if non-nil, zero value otherwise.

### GetTaskPermissionsOk

`func (o *UpdateRoleRequestRole) GetTaskPermissionsOk() (*[]AddRolesRequestRoleTaskPermissionsInner, bool)`

GetTaskPermissionsOk returns a tuple with the TaskPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPermissions

`func (o *UpdateRoleRequestRole) SetTaskPermissions(v []AddRolesRequestRoleTaskPermissionsInner)`

SetTaskPermissions sets TaskPermissions field to given value.

### HasTaskPermissions

`func (o *UpdateRoleRequestRole) HasTaskPermissions() bool`

HasTaskPermissions returns a boolean if a field has been set.

### GetGlobalTaskSetAccess

`func (o *UpdateRoleRequestRole) GetGlobalTaskSetAccess() string`

GetGlobalTaskSetAccess returns the GlobalTaskSetAccess field if non-nil, zero value otherwise.

### GetGlobalTaskSetAccessOk

`func (o *UpdateRoleRequestRole) GetGlobalTaskSetAccessOk() (*string, bool)`

GetGlobalTaskSetAccessOk returns a tuple with the GlobalTaskSetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskSetAccess

`func (o *UpdateRoleRequestRole) SetGlobalTaskSetAccess(v string)`

SetGlobalTaskSetAccess sets GlobalTaskSetAccess field to given value.

### HasGlobalTaskSetAccess

`func (o *UpdateRoleRequestRole) HasGlobalTaskSetAccess() bool`

HasGlobalTaskSetAccess returns a boolean if a field has been set.

### GetTaskSetPermissions

`func (o *UpdateRoleRequestRole) GetTaskSetPermissions() []AddRolesRequestRoleTaskSetPermissionsInner`

GetTaskSetPermissions returns the TaskSetPermissions field if non-nil, zero value otherwise.

### GetTaskSetPermissionsOk

`func (o *UpdateRoleRequestRole) GetTaskSetPermissionsOk() (*[]AddRolesRequestRoleTaskSetPermissionsInner, bool)`

GetTaskSetPermissionsOk returns a tuple with the TaskSetPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetPermissions

`func (o *UpdateRoleRequestRole) SetTaskSetPermissions(v []AddRolesRequestRoleTaskSetPermissionsInner)`

SetTaskSetPermissions sets TaskSetPermissions field to given value.

### HasTaskSetPermissions

`func (o *UpdateRoleRequestRole) HasTaskSetPermissions() bool`

HasTaskSetPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


