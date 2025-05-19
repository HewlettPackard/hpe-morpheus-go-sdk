# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**ListRoles200ResponseAllOfRolesInner**](ListRoles200ResponseAllOfRolesInner.md) |  | [optional] 
**FeaturePermissions** | Pointer to [**[]AddRoles200ResponseAllOfFeaturePermissionsInner**](AddRoles200ResponseAllOfFeaturePermissionsInner.md) |  | [optional] 
**GlobalSiteAccess** | Pointer to **string** |  | [optional] 
**Sites** | Pointer to [**[]AddRoles200ResponseAllOfSitesInner**](AddRoles200ResponseAllOfSitesInner.md) |  | [optional] 
**GlobalZoneAccess** | Pointer to **string** |  | [optional] 
**Zones** | Pointer to [**[]AddRoles200ResponseAllOfSitesInner**](AddRoles200ResponseAllOfSitesInner.md) |  | [optional] 
**GlobalInstanceTypeAccess** | Pointer to **string** |  | [optional] 
**InstanceTypePermissions** | Pointer to [**[]AddRoles200ResponseAllOfInstanceTypePermissionsInner**](AddRoles200ResponseAllOfInstanceTypePermissionsInner.md) |  | [optional] 
**GlobalAppTemplateAccess** | Pointer to **string** |  | [optional] 
**AppTemplatePermissions** | Pointer to [**[]AddRoles200ResponseAllOfAppTemplatePermissionsInner**](AddRoles200ResponseAllOfAppTemplatePermissionsInner.md) |  | [optional] 
**GlobalCatalogItemTypeAccess** | Pointer to **string** |  | [optional] 
**CatalogItemTypePermissions** | Pointer to [**[]AddRoles200ResponseAllOfSitesInner**](AddRoles200ResponseAllOfSitesInner.md) |  | [optional] 
**GlobalPersonaAccess** | Pointer to **string** |  | [optional] 
**PersonaPermissions** | Pointer to [**[]AddRoles200ResponseAllOfInstanceTypePermissionsInner**](AddRoles200ResponseAllOfInstanceTypePermissionsInner.md) |  | [optional] 
**GlobalVdiPoolAccess** | Pointer to **string** |  | [optional] 
**VdiPoolPermissions** | Pointer to [**[]AddRoles200ResponseAllOfSitesInner**](AddRoles200ResponseAllOfSitesInner.md) |  | [optional] 
**GlobalReportTypeAccess** | Pointer to **string** |  | [optional] 
**ReportTypePermissions** | Pointer to [**[]AddRoles200ResponseAllOfInstanceTypePermissionsInner**](AddRoles200ResponseAllOfInstanceTypePermissionsInner.md) |  | [optional] 
**GlobalTaskAccess** | Pointer to **string** |  | [optional] 
**TaskPermissions** | Pointer to [**[]AddRoles200ResponseAllOfAppTemplatePermissionsInner**](AddRoles200ResponseAllOfAppTemplatePermissionsInner.md) |  | [optional] 
**GlobalTaskSetAccess** | Pointer to **string** |  | [optional] 
**TaskSetPermissions** | Pointer to [**[]AddRoles200ResponseAllOfAppTemplatePermissionsInner**](AddRoles200ResponseAllOfAppTemplatePermissionsInner.md) |  | [optional] 

## Methods

### NewRole

`func NewRole() *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *Role) GetRole() ListRoles200ResponseAllOfRolesInner`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Role) GetRoleOk() (*ListRoles200ResponseAllOfRolesInner, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Role) SetRole(v ListRoles200ResponseAllOfRolesInner)`

SetRole sets Role field to given value.

### HasRole

`func (o *Role) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetFeaturePermissions

`func (o *Role) GetFeaturePermissions() []AddRoles200ResponseAllOfFeaturePermissionsInner`

GetFeaturePermissions returns the FeaturePermissions field if non-nil, zero value otherwise.

### GetFeaturePermissionsOk

`func (o *Role) GetFeaturePermissionsOk() (*[]AddRoles200ResponseAllOfFeaturePermissionsInner, bool)`

GetFeaturePermissionsOk returns a tuple with the FeaturePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePermissions

`func (o *Role) SetFeaturePermissions(v []AddRoles200ResponseAllOfFeaturePermissionsInner)`

SetFeaturePermissions sets FeaturePermissions field to given value.

### HasFeaturePermissions

`func (o *Role) HasFeaturePermissions() bool`

HasFeaturePermissions returns a boolean if a field has been set.

### GetGlobalSiteAccess

`func (o *Role) GetGlobalSiteAccess() string`

GetGlobalSiteAccess returns the GlobalSiteAccess field if non-nil, zero value otherwise.

### GetGlobalSiteAccessOk

`func (o *Role) GetGlobalSiteAccessOk() (*string, bool)`

GetGlobalSiteAccessOk returns a tuple with the GlobalSiteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSiteAccess

`func (o *Role) SetGlobalSiteAccess(v string)`

SetGlobalSiteAccess sets GlobalSiteAccess field to given value.

### HasGlobalSiteAccess

`func (o *Role) HasGlobalSiteAccess() bool`

HasGlobalSiteAccess returns a boolean if a field has been set.

### GetSites

`func (o *Role) GetSites() []AddRoles200ResponseAllOfSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *Role) GetSitesOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *Role) SetSites(v []AddRoles200ResponseAllOfSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *Role) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetGlobalZoneAccess

`func (o *Role) GetGlobalZoneAccess() string`

GetGlobalZoneAccess returns the GlobalZoneAccess field if non-nil, zero value otherwise.

### GetGlobalZoneAccessOk

`func (o *Role) GetGlobalZoneAccessOk() (*string, bool)`

GetGlobalZoneAccessOk returns a tuple with the GlobalZoneAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalZoneAccess

`func (o *Role) SetGlobalZoneAccess(v string)`

SetGlobalZoneAccess sets GlobalZoneAccess field to given value.

### HasGlobalZoneAccess

`func (o *Role) HasGlobalZoneAccess() bool`

HasGlobalZoneAccess returns a boolean if a field has been set.

### GetZones

`func (o *Role) GetZones() []AddRoles200ResponseAllOfSitesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *Role) GetZonesOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *Role) SetZones(v []AddRoles200ResponseAllOfSitesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *Role) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetGlobalInstanceTypeAccess

`func (o *Role) GetGlobalInstanceTypeAccess() string`

GetGlobalInstanceTypeAccess returns the GlobalInstanceTypeAccess field if non-nil, zero value otherwise.

### GetGlobalInstanceTypeAccessOk

`func (o *Role) GetGlobalInstanceTypeAccessOk() (*string, bool)`

GetGlobalInstanceTypeAccessOk returns a tuple with the GlobalInstanceTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalInstanceTypeAccess

`func (o *Role) SetGlobalInstanceTypeAccess(v string)`

SetGlobalInstanceTypeAccess sets GlobalInstanceTypeAccess field to given value.

### HasGlobalInstanceTypeAccess

`func (o *Role) HasGlobalInstanceTypeAccess() bool`

HasGlobalInstanceTypeAccess returns a boolean if a field has been set.

### GetInstanceTypePermissions

`func (o *Role) GetInstanceTypePermissions() []AddRoles200ResponseAllOfInstanceTypePermissionsInner`

GetInstanceTypePermissions returns the InstanceTypePermissions field if non-nil, zero value otherwise.

### GetInstanceTypePermissionsOk

`func (o *Role) GetInstanceTypePermissionsOk() (*[]AddRoles200ResponseAllOfInstanceTypePermissionsInner, bool)`

GetInstanceTypePermissionsOk returns a tuple with the InstanceTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypePermissions

`func (o *Role) SetInstanceTypePermissions(v []AddRoles200ResponseAllOfInstanceTypePermissionsInner)`

SetInstanceTypePermissions sets InstanceTypePermissions field to given value.

### HasInstanceTypePermissions

`func (o *Role) HasInstanceTypePermissions() bool`

HasInstanceTypePermissions returns a boolean if a field has been set.

### GetGlobalAppTemplateAccess

`func (o *Role) GetGlobalAppTemplateAccess() string`

GetGlobalAppTemplateAccess returns the GlobalAppTemplateAccess field if non-nil, zero value otherwise.

### GetGlobalAppTemplateAccessOk

`func (o *Role) GetGlobalAppTemplateAccessOk() (*string, bool)`

GetGlobalAppTemplateAccessOk returns a tuple with the GlobalAppTemplateAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAppTemplateAccess

`func (o *Role) SetGlobalAppTemplateAccess(v string)`

SetGlobalAppTemplateAccess sets GlobalAppTemplateAccess field to given value.

### HasGlobalAppTemplateAccess

`func (o *Role) HasGlobalAppTemplateAccess() bool`

HasGlobalAppTemplateAccess returns a boolean if a field has been set.

### GetAppTemplatePermissions

`func (o *Role) GetAppTemplatePermissions() []AddRoles200ResponseAllOfAppTemplatePermissionsInner`

GetAppTemplatePermissions returns the AppTemplatePermissions field if non-nil, zero value otherwise.

### GetAppTemplatePermissionsOk

`func (o *Role) GetAppTemplatePermissionsOk() (*[]AddRoles200ResponseAllOfAppTemplatePermissionsInner, bool)`

GetAppTemplatePermissionsOk returns a tuple with the AppTemplatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplatePermissions

`func (o *Role) SetAppTemplatePermissions(v []AddRoles200ResponseAllOfAppTemplatePermissionsInner)`

SetAppTemplatePermissions sets AppTemplatePermissions field to given value.

### HasAppTemplatePermissions

`func (o *Role) HasAppTemplatePermissions() bool`

HasAppTemplatePermissions returns a boolean if a field has been set.

### GetGlobalCatalogItemTypeAccess

`func (o *Role) GetGlobalCatalogItemTypeAccess() string`

GetGlobalCatalogItemTypeAccess returns the GlobalCatalogItemTypeAccess field if non-nil, zero value otherwise.

### GetGlobalCatalogItemTypeAccessOk

`func (o *Role) GetGlobalCatalogItemTypeAccessOk() (*string, bool)`

GetGlobalCatalogItemTypeAccessOk returns a tuple with the GlobalCatalogItemTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCatalogItemTypeAccess

`func (o *Role) SetGlobalCatalogItemTypeAccess(v string)`

SetGlobalCatalogItemTypeAccess sets GlobalCatalogItemTypeAccess field to given value.

### HasGlobalCatalogItemTypeAccess

`func (o *Role) HasGlobalCatalogItemTypeAccess() bool`

HasGlobalCatalogItemTypeAccess returns a boolean if a field has been set.

### GetCatalogItemTypePermissions

`func (o *Role) GetCatalogItemTypePermissions() []AddRoles200ResponseAllOfSitesInner`

GetCatalogItemTypePermissions returns the CatalogItemTypePermissions field if non-nil, zero value otherwise.

### GetCatalogItemTypePermissionsOk

`func (o *Role) GetCatalogItemTypePermissionsOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetCatalogItemTypePermissionsOk returns a tuple with the CatalogItemTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypePermissions

`func (o *Role) SetCatalogItemTypePermissions(v []AddRoles200ResponseAllOfSitesInner)`

SetCatalogItemTypePermissions sets CatalogItemTypePermissions field to given value.

### HasCatalogItemTypePermissions

`func (o *Role) HasCatalogItemTypePermissions() bool`

HasCatalogItemTypePermissions returns a boolean if a field has been set.

### GetGlobalPersonaAccess

`func (o *Role) GetGlobalPersonaAccess() string`

GetGlobalPersonaAccess returns the GlobalPersonaAccess field if non-nil, zero value otherwise.

### GetGlobalPersonaAccessOk

`func (o *Role) GetGlobalPersonaAccessOk() (*string, bool)`

GetGlobalPersonaAccessOk returns a tuple with the GlobalPersonaAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPersonaAccess

`func (o *Role) SetGlobalPersonaAccess(v string)`

SetGlobalPersonaAccess sets GlobalPersonaAccess field to given value.

### HasGlobalPersonaAccess

`func (o *Role) HasGlobalPersonaAccess() bool`

HasGlobalPersonaAccess returns a boolean if a field has been set.

### GetPersonaPermissions

`func (o *Role) GetPersonaPermissions() []AddRoles200ResponseAllOfInstanceTypePermissionsInner`

GetPersonaPermissions returns the PersonaPermissions field if non-nil, zero value otherwise.

### GetPersonaPermissionsOk

`func (o *Role) GetPersonaPermissionsOk() (*[]AddRoles200ResponseAllOfInstanceTypePermissionsInner, bool)`

GetPersonaPermissionsOk returns a tuple with the PersonaPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaPermissions

`func (o *Role) SetPersonaPermissions(v []AddRoles200ResponseAllOfInstanceTypePermissionsInner)`

SetPersonaPermissions sets PersonaPermissions field to given value.

### HasPersonaPermissions

`func (o *Role) HasPersonaPermissions() bool`

HasPersonaPermissions returns a boolean if a field has been set.

### GetGlobalVdiPoolAccess

`func (o *Role) GetGlobalVdiPoolAccess() string`

GetGlobalVdiPoolAccess returns the GlobalVdiPoolAccess field if non-nil, zero value otherwise.

### GetGlobalVdiPoolAccessOk

`func (o *Role) GetGlobalVdiPoolAccessOk() (*string, bool)`

GetGlobalVdiPoolAccessOk returns a tuple with the GlobalVdiPoolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalVdiPoolAccess

`func (o *Role) SetGlobalVdiPoolAccess(v string)`

SetGlobalVdiPoolAccess sets GlobalVdiPoolAccess field to given value.

### HasGlobalVdiPoolAccess

`func (o *Role) HasGlobalVdiPoolAccess() bool`

HasGlobalVdiPoolAccess returns a boolean if a field has been set.

### GetVdiPoolPermissions

`func (o *Role) GetVdiPoolPermissions() []AddRoles200ResponseAllOfSitesInner`

GetVdiPoolPermissions returns the VdiPoolPermissions field if non-nil, zero value otherwise.

### GetVdiPoolPermissionsOk

`func (o *Role) GetVdiPoolPermissionsOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetVdiPoolPermissionsOk returns a tuple with the VdiPoolPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPoolPermissions

`func (o *Role) SetVdiPoolPermissions(v []AddRoles200ResponseAllOfSitesInner)`

SetVdiPoolPermissions sets VdiPoolPermissions field to given value.

### HasVdiPoolPermissions

`func (o *Role) HasVdiPoolPermissions() bool`

HasVdiPoolPermissions returns a boolean if a field has been set.

### GetGlobalReportTypeAccess

`func (o *Role) GetGlobalReportTypeAccess() string`

GetGlobalReportTypeAccess returns the GlobalReportTypeAccess field if non-nil, zero value otherwise.

### GetGlobalReportTypeAccessOk

`func (o *Role) GetGlobalReportTypeAccessOk() (*string, bool)`

GetGlobalReportTypeAccessOk returns a tuple with the GlobalReportTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalReportTypeAccess

`func (o *Role) SetGlobalReportTypeAccess(v string)`

SetGlobalReportTypeAccess sets GlobalReportTypeAccess field to given value.

### HasGlobalReportTypeAccess

`func (o *Role) HasGlobalReportTypeAccess() bool`

HasGlobalReportTypeAccess returns a boolean if a field has been set.

### GetReportTypePermissions

`func (o *Role) GetReportTypePermissions() []AddRoles200ResponseAllOfInstanceTypePermissionsInner`

GetReportTypePermissions returns the ReportTypePermissions field if non-nil, zero value otherwise.

### GetReportTypePermissionsOk

`func (o *Role) GetReportTypePermissionsOk() (*[]AddRoles200ResponseAllOfInstanceTypePermissionsInner, bool)`

GetReportTypePermissionsOk returns a tuple with the ReportTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypePermissions

`func (o *Role) SetReportTypePermissions(v []AddRoles200ResponseAllOfInstanceTypePermissionsInner)`

SetReportTypePermissions sets ReportTypePermissions field to given value.

### HasReportTypePermissions

`func (o *Role) HasReportTypePermissions() bool`

HasReportTypePermissions returns a boolean if a field has been set.

### GetGlobalTaskAccess

`func (o *Role) GetGlobalTaskAccess() string`

GetGlobalTaskAccess returns the GlobalTaskAccess field if non-nil, zero value otherwise.

### GetGlobalTaskAccessOk

`func (o *Role) GetGlobalTaskAccessOk() (*string, bool)`

GetGlobalTaskAccessOk returns a tuple with the GlobalTaskAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskAccess

`func (o *Role) SetGlobalTaskAccess(v string)`

SetGlobalTaskAccess sets GlobalTaskAccess field to given value.

### HasGlobalTaskAccess

`func (o *Role) HasGlobalTaskAccess() bool`

HasGlobalTaskAccess returns a boolean if a field has been set.

### GetTaskPermissions

`func (o *Role) GetTaskPermissions() []AddRoles200ResponseAllOfAppTemplatePermissionsInner`

GetTaskPermissions returns the TaskPermissions field if non-nil, zero value otherwise.

### GetTaskPermissionsOk

`func (o *Role) GetTaskPermissionsOk() (*[]AddRoles200ResponseAllOfAppTemplatePermissionsInner, bool)`

GetTaskPermissionsOk returns a tuple with the TaskPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPermissions

`func (o *Role) SetTaskPermissions(v []AddRoles200ResponseAllOfAppTemplatePermissionsInner)`

SetTaskPermissions sets TaskPermissions field to given value.

### HasTaskPermissions

`func (o *Role) HasTaskPermissions() bool`

HasTaskPermissions returns a boolean if a field has been set.

### GetGlobalTaskSetAccess

`func (o *Role) GetGlobalTaskSetAccess() string`

GetGlobalTaskSetAccess returns the GlobalTaskSetAccess field if non-nil, zero value otherwise.

### GetGlobalTaskSetAccessOk

`func (o *Role) GetGlobalTaskSetAccessOk() (*string, bool)`

GetGlobalTaskSetAccessOk returns a tuple with the GlobalTaskSetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskSetAccess

`func (o *Role) SetGlobalTaskSetAccess(v string)`

SetGlobalTaskSetAccess sets GlobalTaskSetAccess field to given value.

### HasGlobalTaskSetAccess

`func (o *Role) HasGlobalTaskSetAccess() bool`

HasGlobalTaskSetAccess returns a boolean if a field has been set.

### GetTaskSetPermissions

`func (o *Role) GetTaskSetPermissions() []AddRoles200ResponseAllOfAppTemplatePermissionsInner`

GetTaskSetPermissions returns the TaskSetPermissions field if non-nil, zero value otherwise.

### GetTaskSetPermissionsOk

`func (o *Role) GetTaskSetPermissionsOk() (*[]AddRoles200ResponseAllOfAppTemplatePermissionsInner, bool)`

GetTaskSetPermissionsOk returns a tuple with the TaskSetPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetPermissions

`func (o *Role) SetTaskSetPermissions(v []AddRoles200ResponseAllOfAppTemplatePermissionsInner)`

SetTaskSetPermissions sets TaskSetPermissions field to given value.

### HasTaskSetPermissions

`func (o *Role) HasTaskSetPermissions() bool`

HasTaskSetPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


