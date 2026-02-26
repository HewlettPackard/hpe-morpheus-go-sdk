# AddRolesRequestRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | **string** | Authority (Name) | 
**Description** | Pointer to **NullableString** | Description | [optional] 
**LandingUrl** | Pointer to **NullableString** | An optional override for the default landing page after login for a user. | [optional] 
**RoleType** | Pointer to **string** | Role type | [optional] [default to "user"]
**BaseRoleId** | Pointer to **int64** | Base Role ID. Create the new role with the same permissions and access levels that the specified base role has. If this is not passed, the role is create without any permissions. | [optional] 
**Multitenant** | Pointer to **bool** | Multitenant roles are copied to all tenant accounts and kept in sync until a sub-tenant user modifies their copy of the role. *Only available to master tenant, only applies to user roles* | [optional] [default to false]
**MultitenantLocked** | Pointer to **bool** | Multitenant Locked, prevents sub-tenant users from modifying their copy of multitenant roles. *Only available to master tenant, only applies to user roles* | [optional] [default to false]
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
**GlobalClusterTypeAccess** | Pointer to **string** | Set the default access level for cluster types | [optional] 
**ClusterTypePermissions** | Pointer to [**[]AddRolesRequestRoleClusterTypePermissionsInner**](AddRolesRequestRoleClusterTypePermissionsInner.md) | Set the access level for the specified cluster types | [optional] 
**Owner** | Pointer to **int64** | Set the role owner (tenant) by ID. *Only available to master tenant* | [optional] 

## Methods

### NewAddRolesRequestRole

`func NewAddRolesRequestRole(authority string, ) *AddRolesRequestRole`

NewAddRolesRequestRole instantiates a new AddRolesRequestRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRolesRequestRoleWithDefaults

`func NewAddRolesRequestRoleWithDefaults() *AddRolesRequestRole`

NewAddRolesRequestRoleWithDefaults instantiates a new AddRolesRequestRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *AddRolesRequestRole) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *AddRolesRequestRole) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *AddRolesRequestRole) SetAuthority(v string)`

SetAuthority sets Authority field to given value.


### GetDescription

`func (o *AddRolesRequestRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddRolesRequestRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddRolesRequestRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddRolesRequestRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddRolesRequestRole) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddRolesRequestRole) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLandingUrl

`func (o *AddRolesRequestRole) GetLandingUrl() string`

GetLandingUrl returns the LandingUrl field if non-nil, zero value otherwise.

### GetLandingUrlOk

`func (o *AddRolesRequestRole) GetLandingUrlOk() (*string, bool)`

GetLandingUrlOk returns a tuple with the LandingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrl

`func (o *AddRolesRequestRole) SetLandingUrl(v string)`

SetLandingUrl sets LandingUrl field to given value.

### HasLandingUrl

`func (o *AddRolesRequestRole) HasLandingUrl() bool`

HasLandingUrl returns a boolean if a field has been set.

### SetLandingUrlNil

`func (o *AddRolesRequestRole) SetLandingUrlNil(b bool)`

 SetLandingUrlNil sets the value for LandingUrl to be an explicit nil

### UnsetLandingUrl
`func (o *AddRolesRequestRole) UnsetLandingUrl()`

UnsetLandingUrl ensures that no value is present for LandingUrl, not even an explicit nil
### GetRoleType

`func (o *AddRolesRequestRole) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *AddRolesRequestRole) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *AddRolesRequestRole) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *AddRolesRequestRole) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetBaseRoleId

`func (o *AddRolesRequestRole) GetBaseRoleId() int64`

GetBaseRoleId returns the BaseRoleId field if non-nil, zero value otherwise.

### GetBaseRoleIdOk

`func (o *AddRolesRequestRole) GetBaseRoleIdOk() (*int64, bool)`

GetBaseRoleIdOk returns a tuple with the BaseRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRoleId

`func (o *AddRolesRequestRole) SetBaseRoleId(v int64)`

SetBaseRoleId sets BaseRoleId field to given value.

### HasBaseRoleId

`func (o *AddRolesRequestRole) HasBaseRoleId() bool`

HasBaseRoleId returns a boolean if a field has been set.

### GetMultitenant

`func (o *AddRolesRequestRole) GetMultitenant() bool`

GetMultitenant returns the Multitenant field if non-nil, zero value otherwise.

### GetMultitenantOk

`func (o *AddRolesRequestRole) GetMultitenantOk() (*bool, bool)`

GetMultitenantOk returns a tuple with the Multitenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenant

`func (o *AddRolesRequestRole) SetMultitenant(v bool)`

SetMultitenant sets Multitenant field to given value.

### HasMultitenant

`func (o *AddRolesRequestRole) HasMultitenant() bool`

HasMultitenant returns a boolean if a field has been set.

### GetMultitenantLocked

`func (o *AddRolesRequestRole) GetMultitenantLocked() bool`

GetMultitenantLocked returns the MultitenantLocked field if non-nil, zero value otherwise.

### GetMultitenantLockedOk

`func (o *AddRolesRequestRole) GetMultitenantLockedOk() (*bool, bool)`

GetMultitenantLockedOk returns a tuple with the MultitenantLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenantLocked

`func (o *AddRolesRequestRole) SetMultitenantLocked(v bool)`

SetMultitenantLocked sets MultitenantLocked field to given value.

### HasMultitenantLocked

`func (o *AddRolesRequestRole) HasMultitenantLocked() bool`

HasMultitenantLocked returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *AddRolesRequestRole) GetDefaultPersona() string`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *AddRolesRequestRole) GetDefaultPersonaOk() (*string, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *AddRolesRequestRole) SetDefaultPersona(v string)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *AddRolesRequestRole) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### SetDefaultPersonaNil

`func (o *AddRolesRequestRole) SetDefaultPersonaNil(b bool)`

 SetDefaultPersonaNil sets the value for DefaultPersona to be an explicit nil

### UnsetDefaultPersona
`func (o *AddRolesRequestRole) UnsetDefaultPersona()`

UnsetDefaultPersona ensures that no value is present for DefaultPersona, not even an explicit nil
### GetFeaturePermissions

`func (o *AddRolesRequestRole) GetFeaturePermissions() []AddRolesRequestRoleFeaturePermissionsInner`

GetFeaturePermissions returns the FeaturePermissions field if non-nil, zero value otherwise.

### GetFeaturePermissionsOk

`func (o *AddRolesRequestRole) GetFeaturePermissionsOk() (*[]AddRolesRequestRoleFeaturePermissionsInner, bool)`

GetFeaturePermissionsOk returns a tuple with the FeaturePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePermissions

`func (o *AddRolesRequestRole) SetFeaturePermissions(v []AddRolesRequestRoleFeaturePermissionsInner)`

SetFeaturePermissions sets FeaturePermissions field to given value.

### HasFeaturePermissions

`func (o *AddRolesRequestRole) HasFeaturePermissions() bool`

HasFeaturePermissions returns a boolean if a field has been set.

### GetGlobalSiteAccess

`func (o *AddRolesRequestRole) GetGlobalSiteAccess() string`

GetGlobalSiteAccess returns the GlobalSiteAccess field if non-nil, zero value otherwise.

### GetGlobalSiteAccessOk

`func (o *AddRolesRequestRole) GetGlobalSiteAccessOk() (*string, bool)`

GetGlobalSiteAccessOk returns a tuple with the GlobalSiteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSiteAccess

`func (o *AddRolesRequestRole) SetGlobalSiteAccess(v string)`

SetGlobalSiteAccess sets GlobalSiteAccess field to given value.

### HasGlobalSiteAccess

`func (o *AddRolesRequestRole) HasGlobalSiteAccess() bool`

HasGlobalSiteAccess returns a boolean if a field has been set.

### GetSites

`func (o *AddRolesRequestRole) GetSites() []AddRolesRequestRoleSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *AddRolesRequestRole) GetSitesOk() (*[]AddRolesRequestRoleSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *AddRolesRequestRole) SetSites(v []AddRolesRequestRoleSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *AddRolesRequestRole) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetGlobalZoneAccess

`func (o *AddRolesRequestRole) GetGlobalZoneAccess() string`

GetGlobalZoneAccess returns the GlobalZoneAccess field if non-nil, zero value otherwise.

### GetGlobalZoneAccessOk

`func (o *AddRolesRequestRole) GetGlobalZoneAccessOk() (*string, bool)`

GetGlobalZoneAccessOk returns a tuple with the GlobalZoneAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalZoneAccess

`func (o *AddRolesRequestRole) SetGlobalZoneAccess(v string)`

SetGlobalZoneAccess sets GlobalZoneAccess field to given value.

### HasGlobalZoneAccess

`func (o *AddRolesRequestRole) HasGlobalZoneAccess() bool`

HasGlobalZoneAccess returns a boolean if a field has been set.

### GetZones

`func (o *AddRolesRequestRole) GetZones() []AddRolesRequestRoleZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AddRolesRequestRole) GetZonesOk() (*[]AddRolesRequestRoleZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AddRolesRequestRole) SetZones(v []AddRolesRequestRoleZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AddRolesRequestRole) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetGlobalInstanceTypeAccess

`func (o *AddRolesRequestRole) GetGlobalInstanceTypeAccess() string`

GetGlobalInstanceTypeAccess returns the GlobalInstanceTypeAccess field if non-nil, zero value otherwise.

### GetGlobalInstanceTypeAccessOk

`func (o *AddRolesRequestRole) GetGlobalInstanceTypeAccessOk() (*string, bool)`

GetGlobalInstanceTypeAccessOk returns a tuple with the GlobalInstanceTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalInstanceTypeAccess

`func (o *AddRolesRequestRole) SetGlobalInstanceTypeAccess(v string)`

SetGlobalInstanceTypeAccess sets GlobalInstanceTypeAccess field to given value.

### HasGlobalInstanceTypeAccess

`func (o *AddRolesRequestRole) HasGlobalInstanceTypeAccess() bool`

HasGlobalInstanceTypeAccess returns a boolean if a field has been set.

### GetInstanceTypePermissions

`func (o *AddRolesRequestRole) GetInstanceTypePermissions() []AddRolesRequestRoleInstanceTypePermissionsInner`

GetInstanceTypePermissions returns the InstanceTypePermissions field if non-nil, zero value otherwise.

### GetInstanceTypePermissionsOk

`func (o *AddRolesRequestRole) GetInstanceTypePermissionsOk() (*[]AddRolesRequestRoleInstanceTypePermissionsInner, bool)`

GetInstanceTypePermissionsOk returns a tuple with the InstanceTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypePermissions

`func (o *AddRolesRequestRole) SetInstanceTypePermissions(v []AddRolesRequestRoleInstanceTypePermissionsInner)`

SetInstanceTypePermissions sets InstanceTypePermissions field to given value.

### HasInstanceTypePermissions

`func (o *AddRolesRequestRole) HasInstanceTypePermissions() bool`

HasInstanceTypePermissions returns a boolean if a field has been set.

### GetGlobalAppTemplateAccess

`func (o *AddRolesRequestRole) GetGlobalAppTemplateAccess() string`

GetGlobalAppTemplateAccess returns the GlobalAppTemplateAccess field if non-nil, zero value otherwise.

### GetGlobalAppTemplateAccessOk

`func (o *AddRolesRequestRole) GetGlobalAppTemplateAccessOk() (*string, bool)`

GetGlobalAppTemplateAccessOk returns a tuple with the GlobalAppTemplateAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAppTemplateAccess

`func (o *AddRolesRequestRole) SetGlobalAppTemplateAccess(v string)`

SetGlobalAppTemplateAccess sets GlobalAppTemplateAccess field to given value.

### HasGlobalAppTemplateAccess

`func (o *AddRolesRequestRole) HasGlobalAppTemplateAccess() bool`

HasGlobalAppTemplateAccess returns a boolean if a field has been set.

### GetAppTemplatePermissions

`func (o *AddRolesRequestRole) GetAppTemplatePermissions() []AddRolesRequestRoleAppTemplatePermissionsInner`

GetAppTemplatePermissions returns the AppTemplatePermissions field if non-nil, zero value otherwise.

### GetAppTemplatePermissionsOk

`func (o *AddRolesRequestRole) GetAppTemplatePermissionsOk() (*[]AddRolesRequestRoleAppTemplatePermissionsInner, bool)`

GetAppTemplatePermissionsOk returns a tuple with the AppTemplatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplatePermissions

`func (o *AddRolesRequestRole) SetAppTemplatePermissions(v []AddRolesRequestRoleAppTemplatePermissionsInner)`

SetAppTemplatePermissions sets AppTemplatePermissions field to given value.

### HasAppTemplatePermissions

`func (o *AddRolesRequestRole) HasAppTemplatePermissions() bool`

HasAppTemplatePermissions returns a boolean if a field has been set.

### GetGlobalCatalogItemTypeAccess

`func (o *AddRolesRequestRole) GetGlobalCatalogItemTypeAccess() string`

GetGlobalCatalogItemTypeAccess returns the GlobalCatalogItemTypeAccess field if non-nil, zero value otherwise.

### GetGlobalCatalogItemTypeAccessOk

`func (o *AddRolesRequestRole) GetGlobalCatalogItemTypeAccessOk() (*string, bool)`

GetGlobalCatalogItemTypeAccessOk returns a tuple with the GlobalCatalogItemTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCatalogItemTypeAccess

`func (o *AddRolesRequestRole) SetGlobalCatalogItemTypeAccess(v string)`

SetGlobalCatalogItemTypeAccess sets GlobalCatalogItemTypeAccess field to given value.

### HasGlobalCatalogItemTypeAccess

`func (o *AddRolesRequestRole) HasGlobalCatalogItemTypeAccess() bool`

HasGlobalCatalogItemTypeAccess returns a boolean if a field has been set.

### GetCatalogItemTypePermissions

`func (o *AddRolesRequestRole) GetCatalogItemTypePermissions() []AddRolesRequestRoleCatalogItemTypePermissionsInner`

GetCatalogItemTypePermissions returns the CatalogItemTypePermissions field if non-nil, zero value otherwise.

### GetCatalogItemTypePermissionsOk

`func (o *AddRolesRequestRole) GetCatalogItemTypePermissionsOk() (*[]AddRolesRequestRoleCatalogItemTypePermissionsInner, bool)`

GetCatalogItemTypePermissionsOk returns a tuple with the CatalogItemTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypePermissions

`func (o *AddRolesRequestRole) SetCatalogItemTypePermissions(v []AddRolesRequestRoleCatalogItemTypePermissionsInner)`

SetCatalogItemTypePermissions sets CatalogItemTypePermissions field to given value.

### HasCatalogItemTypePermissions

`func (o *AddRolesRequestRole) HasCatalogItemTypePermissions() bool`

HasCatalogItemTypePermissions returns a boolean if a field has been set.

### GetGlobalPersonaAccess

`func (o *AddRolesRequestRole) GetGlobalPersonaAccess() string`

GetGlobalPersonaAccess returns the GlobalPersonaAccess field if non-nil, zero value otherwise.

### GetGlobalPersonaAccessOk

`func (o *AddRolesRequestRole) GetGlobalPersonaAccessOk() (*string, bool)`

GetGlobalPersonaAccessOk returns a tuple with the GlobalPersonaAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPersonaAccess

`func (o *AddRolesRequestRole) SetGlobalPersonaAccess(v string)`

SetGlobalPersonaAccess sets GlobalPersonaAccess field to given value.

### HasGlobalPersonaAccess

`func (o *AddRolesRequestRole) HasGlobalPersonaAccess() bool`

HasGlobalPersonaAccess returns a boolean if a field has been set.

### GetPersonaPermissions

`func (o *AddRolesRequestRole) GetPersonaPermissions() []AddRolesRequestRolePersonaPermissionsInner`

GetPersonaPermissions returns the PersonaPermissions field if non-nil, zero value otherwise.

### GetPersonaPermissionsOk

`func (o *AddRolesRequestRole) GetPersonaPermissionsOk() (*[]AddRolesRequestRolePersonaPermissionsInner, bool)`

GetPersonaPermissionsOk returns a tuple with the PersonaPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaPermissions

`func (o *AddRolesRequestRole) SetPersonaPermissions(v []AddRolesRequestRolePersonaPermissionsInner)`

SetPersonaPermissions sets PersonaPermissions field to given value.

### HasPersonaPermissions

`func (o *AddRolesRequestRole) HasPersonaPermissions() bool`

HasPersonaPermissions returns a boolean if a field has been set.

### GetGlobalVdiPoolAccess

`func (o *AddRolesRequestRole) GetGlobalVdiPoolAccess() string`

GetGlobalVdiPoolAccess returns the GlobalVdiPoolAccess field if non-nil, zero value otherwise.

### GetGlobalVdiPoolAccessOk

`func (o *AddRolesRequestRole) GetGlobalVdiPoolAccessOk() (*string, bool)`

GetGlobalVdiPoolAccessOk returns a tuple with the GlobalVdiPoolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalVdiPoolAccess

`func (o *AddRolesRequestRole) SetGlobalVdiPoolAccess(v string)`

SetGlobalVdiPoolAccess sets GlobalVdiPoolAccess field to given value.

### HasGlobalVdiPoolAccess

`func (o *AddRolesRequestRole) HasGlobalVdiPoolAccess() bool`

HasGlobalVdiPoolAccess returns a boolean if a field has been set.

### GetVdiPoolPermissions

`func (o *AddRolesRequestRole) GetVdiPoolPermissions() []AddRolesRequestRoleVdiPoolPermissionsInner`

GetVdiPoolPermissions returns the VdiPoolPermissions field if non-nil, zero value otherwise.

### GetVdiPoolPermissionsOk

`func (o *AddRolesRequestRole) GetVdiPoolPermissionsOk() (*[]AddRolesRequestRoleVdiPoolPermissionsInner, bool)`

GetVdiPoolPermissionsOk returns a tuple with the VdiPoolPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPoolPermissions

`func (o *AddRolesRequestRole) SetVdiPoolPermissions(v []AddRolesRequestRoleVdiPoolPermissionsInner)`

SetVdiPoolPermissions sets VdiPoolPermissions field to given value.

### HasVdiPoolPermissions

`func (o *AddRolesRequestRole) HasVdiPoolPermissions() bool`

HasVdiPoolPermissions returns a boolean if a field has been set.

### GetGlobalReportTypeAccess

`func (o *AddRolesRequestRole) GetGlobalReportTypeAccess() string`

GetGlobalReportTypeAccess returns the GlobalReportTypeAccess field if non-nil, zero value otherwise.

### GetGlobalReportTypeAccessOk

`func (o *AddRolesRequestRole) GetGlobalReportTypeAccessOk() (*string, bool)`

GetGlobalReportTypeAccessOk returns a tuple with the GlobalReportTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalReportTypeAccess

`func (o *AddRolesRequestRole) SetGlobalReportTypeAccess(v string)`

SetGlobalReportTypeAccess sets GlobalReportTypeAccess field to given value.

### HasGlobalReportTypeAccess

`func (o *AddRolesRequestRole) HasGlobalReportTypeAccess() bool`

HasGlobalReportTypeAccess returns a boolean if a field has been set.

### GetReportTypePermissions

`func (o *AddRolesRequestRole) GetReportTypePermissions() []AddRolesRequestRoleReportTypePermissionsInner`

GetReportTypePermissions returns the ReportTypePermissions field if non-nil, zero value otherwise.

### GetReportTypePermissionsOk

`func (o *AddRolesRequestRole) GetReportTypePermissionsOk() (*[]AddRolesRequestRoleReportTypePermissionsInner, bool)`

GetReportTypePermissionsOk returns a tuple with the ReportTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypePermissions

`func (o *AddRolesRequestRole) SetReportTypePermissions(v []AddRolesRequestRoleReportTypePermissionsInner)`

SetReportTypePermissions sets ReportTypePermissions field to given value.

### HasReportTypePermissions

`func (o *AddRolesRequestRole) HasReportTypePermissions() bool`

HasReportTypePermissions returns a boolean if a field has been set.

### GetGlobalTaskAccess

`func (o *AddRolesRequestRole) GetGlobalTaskAccess() string`

GetGlobalTaskAccess returns the GlobalTaskAccess field if non-nil, zero value otherwise.

### GetGlobalTaskAccessOk

`func (o *AddRolesRequestRole) GetGlobalTaskAccessOk() (*string, bool)`

GetGlobalTaskAccessOk returns a tuple with the GlobalTaskAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskAccess

`func (o *AddRolesRequestRole) SetGlobalTaskAccess(v string)`

SetGlobalTaskAccess sets GlobalTaskAccess field to given value.

### HasGlobalTaskAccess

`func (o *AddRolesRequestRole) HasGlobalTaskAccess() bool`

HasGlobalTaskAccess returns a boolean if a field has been set.

### GetTaskPermissions

`func (o *AddRolesRequestRole) GetTaskPermissions() []AddRolesRequestRoleTaskPermissionsInner`

GetTaskPermissions returns the TaskPermissions field if non-nil, zero value otherwise.

### GetTaskPermissionsOk

`func (o *AddRolesRequestRole) GetTaskPermissionsOk() (*[]AddRolesRequestRoleTaskPermissionsInner, bool)`

GetTaskPermissionsOk returns a tuple with the TaskPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPermissions

`func (o *AddRolesRequestRole) SetTaskPermissions(v []AddRolesRequestRoleTaskPermissionsInner)`

SetTaskPermissions sets TaskPermissions field to given value.

### HasTaskPermissions

`func (o *AddRolesRequestRole) HasTaskPermissions() bool`

HasTaskPermissions returns a boolean if a field has been set.

### GetGlobalTaskSetAccess

`func (o *AddRolesRequestRole) GetGlobalTaskSetAccess() string`

GetGlobalTaskSetAccess returns the GlobalTaskSetAccess field if non-nil, zero value otherwise.

### GetGlobalTaskSetAccessOk

`func (o *AddRolesRequestRole) GetGlobalTaskSetAccessOk() (*string, bool)`

GetGlobalTaskSetAccessOk returns a tuple with the GlobalTaskSetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskSetAccess

`func (o *AddRolesRequestRole) SetGlobalTaskSetAccess(v string)`

SetGlobalTaskSetAccess sets GlobalTaskSetAccess field to given value.

### HasGlobalTaskSetAccess

`func (o *AddRolesRequestRole) HasGlobalTaskSetAccess() bool`

HasGlobalTaskSetAccess returns a boolean if a field has been set.

### GetTaskSetPermissions

`func (o *AddRolesRequestRole) GetTaskSetPermissions() []AddRolesRequestRoleTaskSetPermissionsInner`

GetTaskSetPermissions returns the TaskSetPermissions field if non-nil, zero value otherwise.

### GetTaskSetPermissionsOk

`func (o *AddRolesRequestRole) GetTaskSetPermissionsOk() (*[]AddRolesRequestRoleTaskSetPermissionsInner, bool)`

GetTaskSetPermissionsOk returns a tuple with the TaskSetPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetPermissions

`func (o *AddRolesRequestRole) SetTaskSetPermissions(v []AddRolesRequestRoleTaskSetPermissionsInner)`

SetTaskSetPermissions sets TaskSetPermissions field to given value.

### HasTaskSetPermissions

`func (o *AddRolesRequestRole) HasTaskSetPermissions() bool`

HasTaskSetPermissions returns a boolean if a field has been set.

### GetGlobalClusterTypeAccess

`func (o *AddRolesRequestRole) GetGlobalClusterTypeAccess() string`

GetGlobalClusterTypeAccess returns the GlobalClusterTypeAccess field if non-nil, zero value otherwise.

### GetGlobalClusterTypeAccessOk

`func (o *AddRolesRequestRole) GetGlobalClusterTypeAccessOk() (*string, bool)`

GetGlobalClusterTypeAccessOk returns a tuple with the GlobalClusterTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalClusterTypeAccess

`func (o *AddRolesRequestRole) SetGlobalClusterTypeAccess(v string)`

SetGlobalClusterTypeAccess sets GlobalClusterTypeAccess field to given value.

### HasGlobalClusterTypeAccess

`func (o *AddRolesRequestRole) HasGlobalClusterTypeAccess() bool`

HasGlobalClusterTypeAccess returns a boolean if a field has been set.

### GetClusterTypePermissions

`func (o *AddRolesRequestRole) GetClusterTypePermissions() []AddRolesRequestRoleClusterTypePermissionsInner`

GetClusterTypePermissions returns the ClusterTypePermissions field if non-nil, zero value otherwise.

### GetClusterTypePermissionsOk

`func (o *AddRolesRequestRole) GetClusterTypePermissionsOk() (*[]AddRolesRequestRoleClusterTypePermissionsInner, bool)`

GetClusterTypePermissionsOk returns a tuple with the ClusterTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterTypePermissions

`func (o *AddRolesRequestRole) SetClusterTypePermissions(v []AddRolesRequestRoleClusterTypePermissionsInner)`

SetClusterTypePermissions sets ClusterTypePermissions field to given value.

### HasClusterTypePermissions

`func (o *AddRolesRequestRole) HasClusterTypePermissions() bool`

HasClusterTypePermissions returns a boolean if a field has been set.

### GetOwner

`func (o *AddRolesRequestRole) GetOwner() int64`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddRolesRequestRole) GetOwnerOk() (*int64, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddRolesRequestRole) SetOwner(v int64)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddRolesRequestRole) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


