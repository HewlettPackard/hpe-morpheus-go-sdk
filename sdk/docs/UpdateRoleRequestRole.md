# UpdateRoleRequestRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to **string** | Authority (Name) | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**LandingUrl** | Pointer to **string** | An optional override for the default landing page after login for a user. | [optional] 
**DefaultPersona** | Pointer to **string** | Set the default persona by code. | [optional] 
**Permissions** | Pointer to [**[]AddRolesRequestRoleFeaturePermissionsInner**](AddRolesRequestRoleFeaturePermissionsInner.md) | Set the access level for the specified permissions. | [optional] 
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
**Personas** | Pointer to [**[]UpdateRoleRequestRolePersonasInner**](UpdateRoleRequestRolePersonasInner.md) | Set the access level for the specified personas | [optional] 
**GlobalVdiPoolAccess** | Pointer to **string** | Set the default access level for VDI pools | [optional] 
**VdiPools** | Pointer to [**[]AddRolesRequestRoleVdiPoolsInner**](AddRolesRequestRoleVdiPoolsInner.md) | Set the access level for the specified VDI pools | [optional] 
**GlobalReportTypeAccess** | Pointer to **string** | Set the default access level for report types | [optional] 
**ReportTypes** | Pointer to [**[]AddRolesRequestRoleReportTypesInner**](AddRolesRequestRoleReportTypesInner.md) | Set the access level for the specified report types | [optional] 
**GlobalTaskAccess** | Pointer to **string** | Set the default access level for tasks | [optional] 
**Tasks** | Pointer to [**[]AddRolesRequestRoleTasksInner**](AddRolesRequestRoleTasksInner.md) | Set the access level for the specified tasks | [optional] 
**GlobalTaskSetAccess** | Pointer to **string** | Set the default access level for workflows (taskSets) | [optional] 
**TaskSets** | Pointer to [**[]AddRolesRequestRoleTaskSetsInner**](AddRolesRequestRoleTaskSetsInner.md) | Set the access level for the specified workflows (taskSets) | [optional] 

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

### GetPermissions

`func (o *UpdateRoleRequestRole) GetPermissions() []AddRolesRequestRoleFeaturePermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateRoleRequestRole) GetPermissionsOk() (*[]AddRolesRequestRoleFeaturePermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateRoleRequestRole) SetPermissions(v []AddRolesRequestRoleFeaturePermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateRoleRequestRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

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

### GetInstanceTypes

`func (o *UpdateRoleRequestRole) GetInstanceTypes() []AddRolesRequestRoleInstanceTypesInner`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *UpdateRoleRequestRole) GetInstanceTypesOk() (*[]AddRolesRequestRoleInstanceTypesInner, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *UpdateRoleRequestRole) SetInstanceTypes(v []AddRolesRequestRoleInstanceTypesInner)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *UpdateRoleRequestRole) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

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

### GetAppTemplates

`func (o *UpdateRoleRequestRole) GetAppTemplates() []AddRolesRequestRoleAppTemplatesInner`

GetAppTemplates returns the AppTemplates field if non-nil, zero value otherwise.

### GetAppTemplatesOk

`func (o *UpdateRoleRequestRole) GetAppTemplatesOk() (*[]AddRolesRequestRoleAppTemplatesInner, bool)`

GetAppTemplatesOk returns a tuple with the AppTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplates

`func (o *UpdateRoleRequestRole) SetAppTemplates(v []AddRolesRequestRoleAppTemplatesInner)`

SetAppTemplates sets AppTemplates field to given value.

### HasAppTemplates

`func (o *UpdateRoleRequestRole) HasAppTemplates() bool`

HasAppTemplates returns a boolean if a field has been set.

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

### GetCatalogItemTypes

`func (o *UpdateRoleRequestRole) GetCatalogItemTypes() []AddRolesRequestRoleCatalogItemTypesInner`

GetCatalogItemTypes returns the CatalogItemTypes field if non-nil, zero value otherwise.

### GetCatalogItemTypesOk

`func (o *UpdateRoleRequestRole) GetCatalogItemTypesOk() (*[]AddRolesRequestRoleCatalogItemTypesInner, bool)`

GetCatalogItemTypesOk returns a tuple with the CatalogItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypes

`func (o *UpdateRoleRequestRole) SetCatalogItemTypes(v []AddRolesRequestRoleCatalogItemTypesInner)`

SetCatalogItemTypes sets CatalogItemTypes field to given value.

### HasCatalogItemTypes

`func (o *UpdateRoleRequestRole) HasCatalogItemTypes() bool`

HasCatalogItemTypes returns a boolean if a field has been set.

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

### GetPersonas

`func (o *UpdateRoleRequestRole) GetPersonas() []UpdateRoleRequestRolePersonasInner`

GetPersonas returns the Personas field if non-nil, zero value otherwise.

### GetPersonasOk

`func (o *UpdateRoleRequestRole) GetPersonasOk() (*[]UpdateRoleRequestRolePersonasInner, bool)`

GetPersonasOk returns a tuple with the Personas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonas

`func (o *UpdateRoleRequestRole) SetPersonas(v []UpdateRoleRequestRolePersonasInner)`

SetPersonas sets Personas field to given value.

### HasPersonas

`func (o *UpdateRoleRequestRole) HasPersonas() bool`

HasPersonas returns a boolean if a field has been set.

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

### GetVdiPools

`func (o *UpdateRoleRequestRole) GetVdiPools() []AddRolesRequestRoleVdiPoolsInner`

GetVdiPools returns the VdiPools field if non-nil, zero value otherwise.

### GetVdiPoolsOk

`func (o *UpdateRoleRequestRole) GetVdiPoolsOk() (*[]AddRolesRequestRoleVdiPoolsInner, bool)`

GetVdiPoolsOk returns a tuple with the VdiPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPools

`func (o *UpdateRoleRequestRole) SetVdiPools(v []AddRolesRequestRoleVdiPoolsInner)`

SetVdiPools sets VdiPools field to given value.

### HasVdiPools

`func (o *UpdateRoleRequestRole) HasVdiPools() bool`

HasVdiPools returns a boolean if a field has been set.

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

### GetReportTypes

`func (o *UpdateRoleRequestRole) GetReportTypes() []AddRolesRequestRoleReportTypesInner`

GetReportTypes returns the ReportTypes field if non-nil, zero value otherwise.

### GetReportTypesOk

`func (o *UpdateRoleRequestRole) GetReportTypesOk() (*[]AddRolesRequestRoleReportTypesInner, bool)`

GetReportTypesOk returns a tuple with the ReportTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypes

`func (o *UpdateRoleRequestRole) SetReportTypes(v []AddRolesRequestRoleReportTypesInner)`

SetReportTypes sets ReportTypes field to given value.

### HasReportTypes

`func (o *UpdateRoleRequestRole) HasReportTypes() bool`

HasReportTypes returns a boolean if a field has been set.

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

### GetTasks

`func (o *UpdateRoleRequestRole) GetTasks() []AddRolesRequestRoleTasksInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *UpdateRoleRequestRole) GetTasksOk() (*[]AddRolesRequestRoleTasksInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *UpdateRoleRequestRole) SetTasks(v []AddRolesRequestRoleTasksInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *UpdateRoleRequestRole) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

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

### GetTaskSets

`func (o *UpdateRoleRequestRole) GetTaskSets() []AddRolesRequestRoleTaskSetsInner`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *UpdateRoleRequestRole) GetTaskSetsOk() (*[]AddRolesRequestRoleTaskSetsInner, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *UpdateRoleRequestRole) SetTaskSets(v []AddRolesRequestRoleTaskSetsInner)`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *UpdateRoleRequestRole) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


