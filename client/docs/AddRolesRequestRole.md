# AddRolesRequestRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | **string** | Authority (Name) | 
**Description** | Pointer to **string** | Description | [optional] 
**LandingUrl** | Pointer to **string** | An optional override for the default landing page after login for a user. | [optional] 
**RoleType** | Pointer to **string** | Role type | [optional] [default to "user"]
**BaseRoleId** | Pointer to **int64** | Base Role ID. Create the new role with the same permissions and access levels that the specified base role has. If this is not passed, the role is create without any permissions. | [optional] 
**Multitenant** | Pointer to **bool** | Multitenant roles are copied to all tenant accounts and kept in sync until a sub-tenant user modifies their copy of the role. *Only available to master tenant* | [optional] [default to false]
**MultitenantLocked** | Pointer to **bool** | Multitenant Locked, prevents sub-tenant users from modifying their copy of multienant roles. *Only available to master tenant* | [optional] [default to false]
**DefaultPersona** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]AddRolesRequestRolePermissionsInner**](AddRolesRequestRolePermissionsInner.md) | Set the access level for the specified permissions. | [optional] 
**GlobalSiteAccess** | Pointer to **string** | Set the default access level for for groups (sites). Only applies to user roles. | [optional] 
**Sites** | Pointer to [**[]AddRolesRequestRoleSitesInner**](AddRolesRequestRoleSitesInner.md) | Set the access level for the specified groups (sites). Only applies to user roles. | [optional] 
**GlobalZoneAccess** | Pointer to **string** | Set the default access level for for clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**Zones** | Pointer to [**[]AddRolesRequestRoleZonesInner**](AddRolesRequestRoleZonesInner.md) | Set the access level for the specified clouds (zones). Only applies to base account (tenant) roles. | [optional] 
**GlobalInstanceTypeAccess** | Pointer to **string** | Set the default access level for for instance types | [optional] 
**InstanceTypes** | Pointer to [**[]AddRolesRequestRoleInstanceTypesInner**](AddRolesRequestRoleInstanceTypesInner.md) | Set the access level for the specified instance types | [optional] 
**GlobalAppTemplateAccess** | Pointer to **string** | Set the default access level for blueprints | [optional] 
**AppTemplates** | Pointer to [**[]AddRolesRequestRoleAppTemplatesInner**](AddRolesRequestRoleAppTemplatesInner.md) | Set the access level for the specified blueprints (appTemplates) | [optional] 
**GlobalCatalogItemTypeAccess** | Pointer to **string** | Set the default access level for catalog item types | [optional] 
**CatalogItemTypes** | Pointer to [**[]AddRolesRequestRoleCatalogItemTypesInner**](AddRolesRequestRoleCatalogItemTypesInner.md) | Set the access level for the specified catalog item types | [optional] 
**GlobalPersonaAccess** | Pointer to **string** | Set the default access level for personas | [optional] 
**Personas** | Pointer to [**[]AddRolesRequestRolePersonasInner**](AddRolesRequestRolePersonasInner.md) | Set the access level for the specified personas | [optional] 
**GlobalVdiPoolAccess** | Pointer to **string** | Set the default access level for VDI pools | [optional] 
**VdiPools** | Pointer to [**[]AddRolesRequestRoleVdiPoolsInner**](AddRolesRequestRoleVdiPoolsInner.md) | Set the access level for the specified VDI pools | [optional] 
**GlobalReportTypeAccess** | Pointer to **string** | Set the default access level for report types | [optional] 
**ReportTypes** | Pointer to [**[]AddRolesRequestRoleReportTypesInner**](AddRolesRequestRoleReportTypesInner.md) | Set the access level for the specified report types | [optional] 
**GlobalTaskAccess** | Pointer to **string** | Set the default access level for tasks | [optional] 
**Tasks** | Pointer to [**[]AddRolesRequestRoleTasksInner**](AddRolesRequestRoleTasksInner.md) | Set the access level for the specified tasks | [optional] 
**GlobalTaskSetAccess** | Pointer to **string** | Set the default access level for workflows (taskSets) | [optional] 
**TaskSets** | Pointer to [**[]AddRolesRequestRoleTaskSetsInner**](AddRolesRequestRoleTaskSetsInner.md) | Set the access level for the specified workflows (taskSets) | [optional] 
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

### GetPermissions

`func (o *AddRolesRequestRole) GetPermissions() []AddRolesRequestRolePermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AddRolesRequestRole) GetPermissionsOk() (*[]AddRolesRequestRolePermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AddRolesRequestRole) SetPermissions(v []AddRolesRequestRolePermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AddRolesRequestRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

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

### GetInstanceTypes

`func (o *AddRolesRequestRole) GetInstanceTypes() []AddRolesRequestRoleInstanceTypesInner`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *AddRolesRequestRole) GetInstanceTypesOk() (*[]AddRolesRequestRoleInstanceTypesInner, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *AddRolesRequestRole) SetInstanceTypes(v []AddRolesRequestRoleInstanceTypesInner)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *AddRolesRequestRole) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

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

### GetAppTemplates

`func (o *AddRolesRequestRole) GetAppTemplates() []AddRolesRequestRoleAppTemplatesInner`

GetAppTemplates returns the AppTemplates field if non-nil, zero value otherwise.

### GetAppTemplatesOk

`func (o *AddRolesRequestRole) GetAppTemplatesOk() (*[]AddRolesRequestRoleAppTemplatesInner, bool)`

GetAppTemplatesOk returns a tuple with the AppTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplates

`func (o *AddRolesRequestRole) SetAppTemplates(v []AddRolesRequestRoleAppTemplatesInner)`

SetAppTemplates sets AppTemplates field to given value.

### HasAppTemplates

`func (o *AddRolesRequestRole) HasAppTemplates() bool`

HasAppTemplates returns a boolean if a field has been set.

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

### GetCatalogItemTypes

`func (o *AddRolesRequestRole) GetCatalogItemTypes() []AddRolesRequestRoleCatalogItemTypesInner`

GetCatalogItemTypes returns the CatalogItemTypes field if non-nil, zero value otherwise.

### GetCatalogItemTypesOk

`func (o *AddRolesRequestRole) GetCatalogItemTypesOk() (*[]AddRolesRequestRoleCatalogItemTypesInner, bool)`

GetCatalogItemTypesOk returns a tuple with the CatalogItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypes

`func (o *AddRolesRequestRole) SetCatalogItemTypes(v []AddRolesRequestRoleCatalogItemTypesInner)`

SetCatalogItemTypes sets CatalogItemTypes field to given value.

### HasCatalogItemTypes

`func (o *AddRolesRequestRole) HasCatalogItemTypes() bool`

HasCatalogItemTypes returns a boolean if a field has been set.

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

### GetPersonas

`func (o *AddRolesRequestRole) GetPersonas() []AddRolesRequestRolePersonasInner`

GetPersonas returns the Personas field if non-nil, zero value otherwise.

### GetPersonasOk

`func (o *AddRolesRequestRole) GetPersonasOk() (*[]AddRolesRequestRolePersonasInner, bool)`

GetPersonasOk returns a tuple with the Personas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonas

`func (o *AddRolesRequestRole) SetPersonas(v []AddRolesRequestRolePersonasInner)`

SetPersonas sets Personas field to given value.

### HasPersonas

`func (o *AddRolesRequestRole) HasPersonas() bool`

HasPersonas returns a boolean if a field has been set.

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

### GetVdiPools

`func (o *AddRolesRequestRole) GetVdiPools() []AddRolesRequestRoleVdiPoolsInner`

GetVdiPools returns the VdiPools field if non-nil, zero value otherwise.

### GetVdiPoolsOk

`func (o *AddRolesRequestRole) GetVdiPoolsOk() (*[]AddRolesRequestRoleVdiPoolsInner, bool)`

GetVdiPoolsOk returns a tuple with the VdiPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPools

`func (o *AddRolesRequestRole) SetVdiPools(v []AddRolesRequestRoleVdiPoolsInner)`

SetVdiPools sets VdiPools field to given value.

### HasVdiPools

`func (o *AddRolesRequestRole) HasVdiPools() bool`

HasVdiPools returns a boolean if a field has been set.

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

### GetReportTypes

`func (o *AddRolesRequestRole) GetReportTypes() []AddRolesRequestRoleReportTypesInner`

GetReportTypes returns the ReportTypes field if non-nil, zero value otherwise.

### GetReportTypesOk

`func (o *AddRolesRequestRole) GetReportTypesOk() (*[]AddRolesRequestRoleReportTypesInner, bool)`

GetReportTypesOk returns a tuple with the ReportTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypes

`func (o *AddRolesRequestRole) SetReportTypes(v []AddRolesRequestRoleReportTypesInner)`

SetReportTypes sets ReportTypes field to given value.

### HasReportTypes

`func (o *AddRolesRequestRole) HasReportTypes() bool`

HasReportTypes returns a boolean if a field has been set.

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

### GetTasks

`func (o *AddRolesRequestRole) GetTasks() []AddRolesRequestRoleTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *AddRolesRequestRole) GetTasksOk() (*[]AddRolesRequestRoleTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *AddRolesRequestRole) SetTasks(v []AddRolesRequestRoleTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *AddRolesRequestRole) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

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

### GetTaskSets

`func (o *AddRolesRequestRole) GetTaskSets() []AddRolesRequestRoleTaskSetsInner`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *AddRolesRequestRole) GetTaskSetsOk() (*[]AddRolesRequestRoleTaskSetsInner, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *AddRolesRequestRole) SetTaskSets(v []AddRolesRequestRoleTaskSetsInner)`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *AddRolesRequestRole) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

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


