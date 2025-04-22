# AddRoles200Response

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
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddRoles200Response

`func NewAddRoles200Response() *AddRoles200Response`

NewAddRoles200Response instantiates a new AddRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRoles200ResponseWithDefaults

`func NewAddRoles200ResponseWithDefaults() *AddRoles200Response`

NewAddRoles200ResponseWithDefaults instantiates a new AddRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *AddRoles200Response) GetRole() ListRoles200ResponseAllOfRolesInner`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddRoles200Response) GetRoleOk() (*ListRoles200ResponseAllOfRolesInner, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddRoles200Response) SetRole(v ListRoles200ResponseAllOfRolesInner)`

SetRole sets Role field to given value.

### HasRole

`func (o *AddRoles200Response) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetFeaturePermissions

`func (o *AddRoles200Response) GetFeaturePermissions() []AddRoles200ResponseAllOfFeaturePermissionsInner`

GetFeaturePermissions returns the FeaturePermissions field if non-nil, zero value otherwise.

### GetFeaturePermissionsOk

`func (o *AddRoles200Response) GetFeaturePermissionsOk() (*[]AddRoles200ResponseAllOfFeaturePermissionsInner, bool)`

GetFeaturePermissionsOk returns a tuple with the FeaturePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePermissions

`func (o *AddRoles200Response) SetFeaturePermissions(v []AddRoles200ResponseAllOfFeaturePermissionsInner)`

SetFeaturePermissions sets FeaturePermissions field to given value.

### HasFeaturePermissions

`func (o *AddRoles200Response) HasFeaturePermissions() bool`

HasFeaturePermissions returns a boolean if a field has been set.

### GetGlobalSiteAccess

`func (o *AddRoles200Response) GetGlobalSiteAccess() string`

GetGlobalSiteAccess returns the GlobalSiteAccess field if non-nil, zero value otherwise.

### GetGlobalSiteAccessOk

`func (o *AddRoles200Response) GetGlobalSiteAccessOk() (*string, bool)`

GetGlobalSiteAccessOk returns a tuple with the GlobalSiteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSiteAccess

`func (o *AddRoles200Response) SetGlobalSiteAccess(v string)`

SetGlobalSiteAccess sets GlobalSiteAccess field to given value.

### HasGlobalSiteAccess

`func (o *AddRoles200Response) HasGlobalSiteAccess() bool`

HasGlobalSiteAccess returns a boolean if a field has been set.

### GetSites

`func (o *AddRoles200Response) GetSites() []AddRoles200ResponseAllOfSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *AddRoles200Response) GetSitesOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *AddRoles200Response) SetSites(v []AddRoles200ResponseAllOfSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *AddRoles200Response) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetGlobalZoneAccess

`func (o *AddRoles200Response) GetGlobalZoneAccess() string`

GetGlobalZoneAccess returns the GlobalZoneAccess field if non-nil, zero value otherwise.

### GetGlobalZoneAccessOk

`func (o *AddRoles200Response) GetGlobalZoneAccessOk() (*string, bool)`

GetGlobalZoneAccessOk returns a tuple with the GlobalZoneAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalZoneAccess

`func (o *AddRoles200Response) SetGlobalZoneAccess(v string)`

SetGlobalZoneAccess sets GlobalZoneAccess field to given value.

### HasGlobalZoneAccess

`func (o *AddRoles200Response) HasGlobalZoneAccess() bool`

HasGlobalZoneAccess returns a boolean if a field has been set.

### GetZones

`func (o *AddRoles200Response) GetZones() []AddRoles200ResponseAllOfSitesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AddRoles200Response) GetZonesOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AddRoles200Response) SetZones(v []AddRoles200ResponseAllOfSitesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AddRoles200Response) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetGlobalInstanceTypeAccess

`func (o *AddRoles200Response) GetGlobalInstanceTypeAccess() string`

GetGlobalInstanceTypeAccess returns the GlobalInstanceTypeAccess field if non-nil, zero value otherwise.

### GetGlobalInstanceTypeAccessOk

`func (o *AddRoles200Response) GetGlobalInstanceTypeAccessOk() (*string, bool)`

GetGlobalInstanceTypeAccessOk returns a tuple with the GlobalInstanceTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalInstanceTypeAccess

`func (o *AddRoles200Response) SetGlobalInstanceTypeAccess(v string)`

SetGlobalInstanceTypeAccess sets GlobalInstanceTypeAccess field to given value.

### HasGlobalInstanceTypeAccess

`func (o *AddRoles200Response) HasGlobalInstanceTypeAccess() bool`

HasGlobalInstanceTypeAccess returns a boolean if a field has been set.

### GetInstanceTypePermissions

`func (o *AddRoles200Response) GetInstanceTypePermissions() []AddRoles200ResponseAllOfInstanceTypePermissionsInner`

GetInstanceTypePermissions returns the InstanceTypePermissions field if non-nil, zero value otherwise.

### GetInstanceTypePermissionsOk

`func (o *AddRoles200Response) GetInstanceTypePermissionsOk() (*[]AddRoles200ResponseAllOfInstanceTypePermissionsInner, bool)`

GetInstanceTypePermissionsOk returns a tuple with the InstanceTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypePermissions

`func (o *AddRoles200Response) SetInstanceTypePermissions(v []AddRoles200ResponseAllOfInstanceTypePermissionsInner)`

SetInstanceTypePermissions sets InstanceTypePermissions field to given value.

### HasInstanceTypePermissions

`func (o *AddRoles200Response) HasInstanceTypePermissions() bool`

HasInstanceTypePermissions returns a boolean if a field has been set.

### GetGlobalAppTemplateAccess

`func (o *AddRoles200Response) GetGlobalAppTemplateAccess() string`

GetGlobalAppTemplateAccess returns the GlobalAppTemplateAccess field if non-nil, zero value otherwise.

### GetGlobalAppTemplateAccessOk

`func (o *AddRoles200Response) GetGlobalAppTemplateAccessOk() (*string, bool)`

GetGlobalAppTemplateAccessOk returns a tuple with the GlobalAppTemplateAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAppTemplateAccess

`func (o *AddRoles200Response) SetGlobalAppTemplateAccess(v string)`

SetGlobalAppTemplateAccess sets GlobalAppTemplateAccess field to given value.

### HasGlobalAppTemplateAccess

`func (o *AddRoles200Response) HasGlobalAppTemplateAccess() bool`

HasGlobalAppTemplateAccess returns a boolean if a field has been set.

### GetAppTemplatePermissions

`func (o *AddRoles200Response) GetAppTemplatePermissions() []AddRoles200ResponseAllOfAppTemplatePermissionsInner`

GetAppTemplatePermissions returns the AppTemplatePermissions field if non-nil, zero value otherwise.

### GetAppTemplatePermissionsOk

`func (o *AddRoles200Response) GetAppTemplatePermissionsOk() (*[]AddRoles200ResponseAllOfAppTemplatePermissionsInner, bool)`

GetAppTemplatePermissionsOk returns a tuple with the AppTemplatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplatePermissions

`func (o *AddRoles200Response) SetAppTemplatePermissions(v []AddRoles200ResponseAllOfAppTemplatePermissionsInner)`

SetAppTemplatePermissions sets AppTemplatePermissions field to given value.

### HasAppTemplatePermissions

`func (o *AddRoles200Response) HasAppTemplatePermissions() bool`

HasAppTemplatePermissions returns a boolean if a field has been set.

### GetGlobalCatalogItemTypeAccess

`func (o *AddRoles200Response) GetGlobalCatalogItemTypeAccess() string`

GetGlobalCatalogItemTypeAccess returns the GlobalCatalogItemTypeAccess field if non-nil, zero value otherwise.

### GetGlobalCatalogItemTypeAccessOk

`func (o *AddRoles200Response) GetGlobalCatalogItemTypeAccessOk() (*string, bool)`

GetGlobalCatalogItemTypeAccessOk returns a tuple with the GlobalCatalogItemTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCatalogItemTypeAccess

`func (o *AddRoles200Response) SetGlobalCatalogItemTypeAccess(v string)`

SetGlobalCatalogItemTypeAccess sets GlobalCatalogItemTypeAccess field to given value.

### HasGlobalCatalogItemTypeAccess

`func (o *AddRoles200Response) HasGlobalCatalogItemTypeAccess() bool`

HasGlobalCatalogItemTypeAccess returns a boolean if a field has been set.

### GetCatalogItemTypePermissions

`func (o *AddRoles200Response) GetCatalogItemTypePermissions() []AddRoles200ResponseAllOfSitesInner`

GetCatalogItemTypePermissions returns the CatalogItemTypePermissions field if non-nil, zero value otherwise.

### GetCatalogItemTypePermissionsOk

`func (o *AddRoles200Response) GetCatalogItemTypePermissionsOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetCatalogItemTypePermissionsOk returns a tuple with the CatalogItemTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypePermissions

`func (o *AddRoles200Response) SetCatalogItemTypePermissions(v []AddRoles200ResponseAllOfSitesInner)`

SetCatalogItemTypePermissions sets CatalogItemTypePermissions field to given value.

### HasCatalogItemTypePermissions

`func (o *AddRoles200Response) HasCatalogItemTypePermissions() bool`

HasCatalogItemTypePermissions returns a boolean if a field has been set.

### GetGlobalPersonaAccess

`func (o *AddRoles200Response) GetGlobalPersonaAccess() string`

GetGlobalPersonaAccess returns the GlobalPersonaAccess field if non-nil, zero value otherwise.

### GetGlobalPersonaAccessOk

`func (o *AddRoles200Response) GetGlobalPersonaAccessOk() (*string, bool)`

GetGlobalPersonaAccessOk returns a tuple with the GlobalPersonaAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPersonaAccess

`func (o *AddRoles200Response) SetGlobalPersonaAccess(v string)`

SetGlobalPersonaAccess sets GlobalPersonaAccess field to given value.

### HasGlobalPersonaAccess

`func (o *AddRoles200Response) HasGlobalPersonaAccess() bool`

HasGlobalPersonaAccess returns a boolean if a field has been set.

### GetPersonaPermissions

`func (o *AddRoles200Response) GetPersonaPermissions() []AddRoles200ResponseAllOfInstanceTypePermissionsInner`

GetPersonaPermissions returns the PersonaPermissions field if non-nil, zero value otherwise.

### GetPersonaPermissionsOk

`func (o *AddRoles200Response) GetPersonaPermissionsOk() (*[]AddRoles200ResponseAllOfInstanceTypePermissionsInner, bool)`

GetPersonaPermissionsOk returns a tuple with the PersonaPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaPermissions

`func (o *AddRoles200Response) SetPersonaPermissions(v []AddRoles200ResponseAllOfInstanceTypePermissionsInner)`

SetPersonaPermissions sets PersonaPermissions field to given value.

### HasPersonaPermissions

`func (o *AddRoles200Response) HasPersonaPermissions() bool`

HasPersonaPermissions returns a boolean if a field has been set.

### GetGlobalVdiPoolAccess

`func (o *AddRoles200Response) GetGlobalVdiPoolAccess() string`

GetGlobalVdiPoolAccess returns the GlobalVdiPoolAccess field if non-nil, zero value otherwise.

### GetGlobalVdiPoolAccessOk

`func (o *AddRoles200Response) GetGlobalVdiPoolAccessOk() (*string, bool)`

GetGlobalVdiPoolAccessOk returns a tuple with the GlobalVdiPoolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalVdiPoolAccess

`func (o *AddRoles200Response) SetGlobalVdiPoolAccess(v string)`

SetGlobalVdiPoolAccess sets GlobalVdiPoolAccess field to given value.

### HasGlobalVdiPoolAccess

`func (o *AddRoles200Response) HasGlobalVdiPoolAccess() bool`

HasGlobalVdiPoolAccess returns a boolean if a field has been set.

### GetVdiPoolPermissions

`func (o *AddRoles200Response) GetVdiPoolPermissions() []AddRoles200ResponseAllOfSitesInner`

GetVdiPoolPermissions returns the VdiPoolPermissions field if non-nil, zero value otherwise.

### GetVdiPoolPermissionsOk

`func (o *AddRoles200Response) GetVdiPoolPermissionsOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetVdiPoolPermissionsOk returns a tuple with the VdiPoolPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPoolPermissions

`func (o *AddRoles200Response) SetVdiPoolPermissions(v []AddRoles200ResponseAllOfSitesInner)`

SetVdiPoolPermissions sets VdiPoolPermissions field to given value.

### HasVdiPoolPermissions

`func (o *AddRoles200Response) HasVdiPoolPermissions() bool`

HasVdiPoolPermissions returns a boolean if a field has been set.

### GetGlobalReportTypeAccess

`func (o *AddRoles200Response) GetGlobalReportTypeAccess() string`

GetGlobalReportTypeAccess returns the GlobalReportTypeAccess field if non-nil, zero value otherwise.

### GetGlobalReportTypeAccessOk

`func (o *AddRoles200Response) GetGlobalReportTypeAccessOk() (*string, bool)`

GetGlobalReportTypeAccessOk returns a tuple with the GlobalReportTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalReportTypeAccess

`func (o *AddRoles200Response) SetGlobalReportTypeAccess(v string)`

SetGlobalReportTypeAccess sets GlobalReportTypeAccess field to given value.

### HasGlobalReportTypeAccess

`func (o *AddRoles200Response) HasGlobalReportTypeAccess() bool`

HasGlobalReportTypeAccess returns a boolean if a field has been set.

### GetReportTypePermissions

`func (o *AddRoles200Response) GetReportTypePermissions() []AddRoles200ResponseAllOfInstanceTypePermissionsInner`

GetReportTypePermissions returns the ReportTypePermissions field if non-nil, zero value otherwise.

### GetReportTypePermissionsOk

`func (o *AddRoles200Response) GetReportTypePermissionsOk() (*[]AddRoles200ResponseAllOfInstanceTypePermissionsInner, bool)`

GetReportTypePermissionsOk returns a tuple with the ReportTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypePermissions

`func (o *AddRoles200Response) SetReportTypePermissions(v []AddRoles200ResponseAllOfInstanceTypePermissionsInner)`

SetReportTypePermissions sets ReportTypePermissions field to given value.

### HasReportTypePermissions

`func (o *AddRoles200Response) HasReportTypePermissions() bool`

HasReportTypePermissions returns a boolean if a field has been set.

### GetGlobalTaskAccess

`func (o *AddRoles200Response) GetGlobalTaskAccess() string`

GetGlobalTaskAccess returns the GlobalTaskAccess field if non-nil, zero value otherwise.

### GetGlobalTaskAccessOk

`func (o *AddRoles200Response) GetGlobalTaskAccessOk() (*string, bool)`

GetGlobalTaskAccessOk returns a tuple with the GlobalTaskAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskAccess

`func (o *AddRoles200Response) SetGlobalTaskAccess(v string)`

SetGlobalTaskAccess sets GlobalTaskAccess field to given value.

### HasGlobalTaskAccess

`func (o *AddRoles200Response) HasGlobalTaskAccess() bool`

HasGlobalTaskAccess returns a boolean if a field has been set.

### GetTaskPermissions

`func (o *AddRoles200Response) GetTaskPermissions() []AddRoles200ResponseAllOfAppTemplatePermissionsInner`

GetTaskPermissions returns the TaskPermissions field if non-nil, zero value otherwise.

### GetTaskPermissionsOk

`func (o *AddRoles200Response) GetTaskPermissionsOk() (*[]AddRoles200ResponseAllOfAppTemplatePermissionsInner, bool)`

GetTaskPermissionsOk returns a tuple with the TaskPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPermissions

`func (o *AddRoles200Response) SetTaskPermissions(v []AddRoles200ResponseAllOfAppTemplatePermissionsInner)`

SetTaskPermissions sets TaskPermissions field to given value.

### HasTaskPermissions

`func (o *AddRoles200Response) HasTaskPermissions() bool`

HasTaskPermissions returns a boolean if a field has been set.

### GetGlobalTaskSetAccess

`func (o *AddRoles200Response) GetGlobalTaskSetAccess() string`

GetGlobalTaskSetAccess returns the GlobalTaskSetAccess field if non-nil, zero value otherwise.

### GetGlobalTaskSetAccessOk

`func (o *AddRoles200Response) GetGlobalTaskSetAccessOk() (*string, bool)`

GetGlobalTaskSetAccessOk returns a tuple with the GlobalTaskSetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskSetAccess

`func (o *AddRoles200Response) SetGlobalTaskSetAccess(v string)`

SetGlobalTaskSetAccess sets GlobalTaskSetAccess field to given value.

### HasGlobalTaskSetAccess

`func (o *AddRoles200Response) HasGlobalTaskSetAccess() bool`

HasGlobalTaskSetAccess returns a boolean if a field has been set.

### GetTaskSetPermissions

`func (o *AddRoles200Response) GetTaskSetPermissions() []AddRoles200ResponseAllOfAppTemplatePermissionsInner`

GetTaskSetPermissions returns the TaskSetPermissions field if non-nil, zero value otherwise.

### GetTaskSetPermissionsOk

`func (o *AddRoles200Response) GetTaskSetPermissionsOk() (*[]AddRoles200ResponseAllOfAppTemplatePermissionsInner, bool)`

GetTaskSetPermissionsOk returns a tuple with the TaskSetPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetPermissions

`func (o *AddRoles200Response) SetTaskSetPermissions(v []AddRoles200ResponseAllOfAppTemplatePermissionsInner)`

SetTaskSetPermissions sets TaskSetPermissions field to given value.

### HasTaskSetPermissions

`func (o *AddRoles200Response) HasTaskSetPermissions() bool`

HasTaskSetPermissions returns a boolean if a field has been set.

### GetSuccess

`func (o *AddRoles200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddRoles200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddRoles200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddRoles200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


