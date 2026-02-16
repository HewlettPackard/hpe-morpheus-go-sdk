# GetRole200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**GetRole200ResponseRole**](GetRole200ResponseRole.md) |  | [optional] 
**FeaturePermissions** | Pointer to [**[]GetRole200ResponseFeaturePermissionsInner**](GetRole200ResponseFeaturePermissionsInner.md) |  | [optional] 
**GlobalSiteAccess** | Pointer to **string** |  | [optional] 
**Sites** | Pointer to [**[]GetRole200ResponseSitesInner**](GetRole200ResponseSitesInner.md) |  | [optional] 
**GlobalZoneAccess** | Pointer to **string** |  | [optional] 
**Zones** | Pointer to [**[]GetRole200ResponseZonesInner**](GetRole200ResponseZonesInner.md) |  | [optional] 
**GlobalInstanceTypeAccess** | Pointer to **string** |  | [optional] 
**InstanceTypePermissions** | Pointer to [**[]GetRole200ResponseInstanceTypePermissionsInner**](GetRole200ResponseInstanceTypePermissionsInner.md) |  | [optional] 
**GlobalAppTemplateAccess** | Pointer to **string** |  | [optional] 
**AppTemplatePermissions** | Pointer to [**[]GetRole200ResponseAppTemplatePermissionsInner**](GetRole200ResponseAppTemplatePermissionsInner.md) |  | [optional] 
**GlobalCatalogItemTypeAccess** | Pointer to **string** |  | [optional] 
**CatalogItemTypePermissions** | Pointer to [**[]GetRole200ResponseCatalogItemTypePermissionsInner**](GetRole200ResponseCatalogItemTypePermissionsInner.md) |  | [optional] 
**GlobalPersonaAccess** | Pointer to **string** |  | [optional] 
**PersonaPermissions** | Pointer to [**[]GetRole200ResponsePersonaPermissionsInner**](GetRole200ResponsePersonaPermissionsInner.md) |  | [optional] 
**GlobalVdiPoolAccess** | Pointer to **string** |  | [optional] 
**VdiPoolPermissions** | Pointer to [**[]GetRole200ResponseVdiPoolPermissionsInner**](GetRole200ResponseVdiPoolPermissionsInner.md) |  | [optional] 
**GlobalReportTypeAccess** | Pointer to **string** |  | [optional] 
**ReportTypePermissions** | Pointer to [**[]GetRole200ResponseReportTypePermissionsInner**](GetRole200ResponseReportTypePermissionsInner.md) |  | [optional] 
**GlobalTaskAccess** | Pointer to **string** |  | [optional] 
**TaskPermissions** | Pointer to [**[]GetRole200ResponseTaskPermissionsInner**](GetRole200ResponseTaskPermissionsInner.md) |  | [optional] 
**GlobalTaskSetAccess** | Pointer to **string** |  | [optional] 
**TaskSetPermissions** | Pointer to [**[]GetRole200ResponseTaskSetPermissionsInner**](GetRole200ResponseTaskSetPermissionsInner.md) |  | [optional] 

## Methods

### NewGetRole200Response

`func NewGetRole200Response() *GetRole200Response`

NewGetRole200Response instantiates a new GetRole200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRole200ResponseWithDefaults

`func NewGetRole200ResponseWithDefaults() *GetRole200Response`

NewGetRole200ResponseWithDefaults instantiates a new GetRole200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *GetRole200Response) GetRole() GetRole200ResponseRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetRole200Response) GetRoleOk() (*GetRole200ResponseRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetRole200Response) SetRole(v GetRole200ResponseRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetRole200Response) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetFeaturePermissions

`func (o *GetRole200Response) GetFeaturePermissions() []GetRole200ResponseFeaturePermissionsInner`

GetFeaturePermissions returns the FeaturePermissions field if non-nil, zero value otherwise.

### GetFeaturePermissionsOk

`func (o *GetRole200Response) GetFeaturePermissionsOk() (*[]GetRole200ResponseFeaturePermissionsInner, bool)`

GetFeaturePermissionsOk returns a tuple with the FeaturePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePermissions

`func (o *GetRole200Response) SetFeaturePermissions(v []GetRole200ResponseFeaturePermissionsInner)`

SetFeaturePermissions sets FeaturePermissions field to given value.

### HasFeaturePermissions

`func (o *GetRole200Response) HasFeaturePermissions() bool`

HasFeaturePermissions returns a boolean if a field has been set.

### GetGlobalSiteAccess

`func (o *GetRole200Response) GetGlobalSiteAccess() string`

GetGlobalSiteAccess returns the GlobalSiteAccess field if non-nil, zero value otherwise.

### GetGlobalSiteAccessOk

`func (o *GetRole200Response) GetGlobalSiteAccessOk() (*string, bool)`

GetGlobalSiteAccessOk returns a tuple with the GlobalSiteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSiteAccess

`func (o *GetRole200Response) SetGlobalSiteAccess(v string)`

SetGlobalSiteAccess sets GlobalSiteAccess field to given value.

### HasGlobalSiteAccess

`func (o *GetRole200Response) HasGlobalSiteAccess() bool`

HasGlobalSiteAccess returns a boolean if a field has been set.

### GetSites

`func (o *GetRole200Response) GetSites() []GetRole200ResponseSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *GetRole200Response) GetSitesOk() (*[]GetRole200ResponseSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *GetRole200Response) SetSites(v []GetRole200ResponseSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *GetRole200Response) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetGlobalZoneAccess

`func (o *GetRole200Response) GetGlobalZoneAccess() string`

GetGlobalZoneAccess returns the GlobalZoneAccess field if non-nil, zero value otherwise.

### GetGlobalZoneAccessOk

`func (o *GetRole200Response) GetGlobalZoneAccessOk() (*string, bool)`

GetGlobalZoneAccessOk returns a tuple with the GlobalZoneAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalZoneAccess

`func (o *GetRole200Response) SetGlobalZoneAccess(v string)`

SetGlobalZoneAccess sets GlobalZoneAccess field to given value.

### HasGlobalZoneAccess

`func (o *GetRole200Response) HasGlobalZoneAccess() bool`

HasGlobalZoneAccess returns a boolean if a field has been set.

### GetZones

`func (o *GetRole200Response) GetZones() []GetRole200ResponseZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *GetRole200Response) GetZonesOk() (*[]GetRole200ResponseZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *GetRole200Response) SetZones(v []GetRole200ResponseZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *GetRole200Response) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetGlobalInstanceTypeAccess

`func (o *GetRole200Response) GetGlobalInstanceTypeAccess() string`

GetGlobalInstanceTypeAccess returns the GlobalInstanceTypeAccess field if non-nil, zero value otherwise.

### GetGlobalInstanceTypeAccessOk

`func (o *GetRole200Response) GetGlobalInstanceTypeAccessOk() (*string, bool)`

GetGlobalInstanceTypeAccessOk returns a tuple with the GlobalInstanceTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalInstanceTypeAccess

`func (o *GetRole200Response) SetGlobalInstanceTypeAccess(v string)`

SetGlobalInstanceTypeAccess sets GlobalInstanceTypeAccess field to given value.

### HasGlobalInstanceTypeAccess

`func (o *GetRole200Response) HasGlobalInstanceTypeAccess() bool`

HasGlobalInstanceTypeAccess returns a boolean if a field has been set.

### GetInstanceTypePermissions

`func (o *GetRole200Response) GetInstanceTypePermissions() []GetRole200ResponseInstanceTypePermissionsInner`

GetInstanceTypePermissions returns the InstanceTypePermissions field if non-nil, zero value otherwise.

### GetInstanceTypePermissionsOk

`func (o *GetRole200Response) GetInstanceTypePermissionsOk() (*[]GetRole200ResponseInstanceTypePermissionsInner, bool)`

GetInstanceTypePermissionsOk returns a tuple with the InstanceTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypePermissions

`func (o *GetRole200Response) SetInstanceTypePermissions(v []GetRole200ResponseInstanceTypePermissionsInner)`

SetInstanceTypePermissions sets InstanceTypePermissions field to given value.

### HasInstanceTypePermissions

`func (o *GetRole200Response) HasInstanceTypePermissions() bool`

HasInstanceTypePermissions returns a boolean if a field has been set.

### GetGlobalAppTemplateAccess

`func (o *GetRole200Response) GetGlobalAppTemplateAccess() string`

GetGlobalAppTemplateAccess returns the GlobalAppTemplateAccess field if non-nil, zero value otherwise.

### GetGlobalAppTemplateAccessOk

`func (o *GetRole200Response) GetGlobalAppTemplateAccessOk() (*string, bool)`

GetGlobalAppTemplateAccessOk returns a tuple with the GlobalAppTemplateAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAppTemplateAccess

`func (o *GetRole200Response) SetGlobalAppTemplateAccess(v string)`

SetGlobalAppTemplateAccess sets GlobalAppTemplateAccess field to given value.

### HasGlobalAppTemplateAccess

`func (o *GetRole200Response) HasGlobalAppTemplateAccess() bool`

HasGlobalAppTemplateAccess returns a boolean if a field has been set.

### GetAppTemplatePermissions

`func (o *GetRole200Response) GetAppTemplatePermissions() []GetRole200ResponseAppTemplatePermissionsInner`

GetAppTemplatePermissions returns the AppTemplatePermissions field if non-nil, zero value otherwise.

### GetAppTemplatePermissionsOk

`func (o *GetRole200Response) GetAppTemplatePermissionsOk() (*[]GetRole200ResponseAppTemplatePermissionsInner, bool)`

GetAppTemplatePermissionsOk returns a tuple with the AppTemplatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplatePermissions

`func (o *GetRole200Response) SetAppTemplatePermissions(v []GetRole200ResponseAppTemplatePermissionsInner)`

SetAppTemplatePermissions sets AppTemplatePermissions field to given value.

### HasAppTemplatePermissions

`func (o *GetRole200Response) HasAppTemplatePermissions() bool`

HasAppTemplatePermissions returns a boolean if a field has been set.

### GetGlobalCatalogItemTypeAccess

`func (o *GetRole200Response) GetGlobalCatalogItemTypeAccess() string`

GetGlobalCatalogItemTypeAccess returns the GlobalCatalogItemTypeAccess field if non-nil, zero value otherwise.

### GetGlobalCatalogItemTypeAccessOk

`func (o *GetRole200Response) GetGlobalCatalogItemTypeAccessOk() (*string, bool)`

GetGlobalCatalogItemTypeAccessOk returns a tuple with the GlobalCatalogItemTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCatalogItemTypeAccess

`func (o *GetRole200Response) SetGlobalCatalogItemTypeAccess(v string)`

SetGlobalCatalogItemTypeAccess sets GlobalCatalogItemTypeAccess field to given value.

### HasGlobalCatalogItemTypeAccess

`func (o *GetRole200Response) HasGlobalCatalogItemTypeAccess() bool`

HasGlobalCatalogItemTypeAccess returns a boolean if a field has been set.

### GetCatalogItemTypePermissions

`func (o *GetRole200Response) GetCatalogItemTypePermissions() []GetRole200ResponseCatalogItemTypePermissionsInner`

GetCatalogItemTypePermissions returns the CatalogItemTypePermissions field if non-nil, zero value otherwise.

### GetCatalogItemTypePermissionsOk

`func (o *GetRole200Response) GetCatalogItemTypePermissionsOk() (*[]GetRole200ResponseCatalogItemTypePermissionsInner, bool)`

GetCatalogItemTypePermissionsOk returns a tuple with the CatalogItemTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypePermissions

`func (o *GetRole200Response) SetCatalogItemTypePermissions(v []GetRole200ResponseCatalogItemTypePermissionsInner)`

SetCatalogItemTypePermissions sets CatalogItemTypePermissions field to given value.

### HasCatalogItemTypePermissions

`func (o *GetRole200Response) HasCatalogItemTypePermissions() bool`

HasCatalogItemTypePermissions returns a boolean if a field has been set.

### GetGlobalPersonaAccess

`func (o *GetRole200Response) GetGlobalPersonaAccess() string`

GetGlobalPersonaAccess returns the GlobalPersonaAccess field if non-nil, zero value otherwise.

### GetGlobalPersonaAccessOk

`func (o *GetRole200Response) GetGlobalPersonaAccessOk() (*string, bool)`

GetGlobalPersonaAccessOk returns a tuple with the GlobalPersonaAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPersonaAccess

`func (o *GetRole200Response) SetGlobalPersonaAccess(v string)`

SetGlobalPersonaAccess sets GlobalPersonaAccess field to given value.

### HasGlobalPersonaAccess

`func (o *GetRole200Response) HasGlobalPersonaAccess() bool`

HasGlobalPersonaAccess returns a boolean if a field has been set.

### GetPersonaPermissions

`func (o *GetRole200Response) GetPersonaPermissions() []GetRole200ResponsePersonaPermissionsInner`

GetPersonaPermissions returns the PersonaPermissions field if non-nil, zero value otherwise.

### GetPersonaPermissionsOk

`func (o *GetRole200Response) GetPersonaPermissionsOk() (*[]GetRole200ResponsePersonaPermissionsInner, bool)`

GetPersonaPermissionsOk returns a tuple with the PersonaPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaPermissions

`func (o *GetRole200Response) SetPersonaPermissions(v []GetRole200ResponsePersonaPermissionsInner)`

SetPersonaPermissions sets PersonaPermissions field to given value.

### HasPersonaPermissions

`func (o *GetRole200Response) HasPersonaPermissions() bool`

HasPersonaPermissions returns a boolean if a field has been set.

### GetGlobalVdiPoolAccess

`func (o *GetRole200Response) GetGlobalVdiPoolAccess() string`

GetGlobalVdiPoolAccess returns the GlobalVdiPoolAccess field if non-nil, zero value otherwise.

### GetGlobalVdiPoolAccessOk

`func (o *GetRole200Response) GetGlobalVdiPoolAccessOk() (*string, bool)`

GetGlobalVdiPoolAccessOk returns a tuple with the GlobalVdiPoolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalVdiPoolAccess

`func (o *GetRole200Response) SetGlobalVdiPoolAccess(v string)`

SetGlobalVdiPoolAccess sets GlobalVdiPoolAccess field to given value.

### HasGlobalVdiPoolAccess

`func (o *GetRole200Response) HasGlobalVdiPoolAccess() bool`

HasGlobalVdiPoolAccess returns a boolean if a field has been set.

### GetVdiPoolPermissions

`func (o *GetRole200Response) GetVdiPoolPermissions() []GetRole200ResponseVdiPoolPermissionsInner`

GetVdiPoolPermissions returns the VdiPoolPermissions field if non-nil, zero value otherwise.

### GetVdiPoolPermissionsOk

`func (o *GetRole200Response) GetVdiPoolPermissionsOk() (*[]GetRole200ResponseVdiPoolPermissionsInner, bool)`

GetVdiPoolPermissionsOk returns a tuple with the VdiPoolPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPoolPermissions

`func (o *GetRole200Response) SetVdiPoolPermissions(v []GetRole200ResponseVdiPoolPermissionsInner)`

SetVdiPoolPermissions sets VdiPoolPermissions field to given value.

### HasVdiPoolPermissions

`func (o *GetRole200Response) HasVdiPoolPermissions() bool`

HasVdiPoolPermissions returns a boolean if a field has been set.

### GetGlobalReportTypeAccess

`func (o *GetRole200Response) GetGlobalReportTypeAccess() string`

GetGlobalReportTypeAccess returns the GlobalReportTypeAccess field if non-nil, zero value otherwise.

### GetGlobalReportTypeAccessOk

`func (o *GetRole200Response) GetGlobalReportTypeAccessOk() (*string, bool)`

GetGlobalReportTypeAccessOk returns a tuple with the GlobalReportTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalReportTypeAccess

`func (o *GetRole200Response) SetGlobalReportTypeAccess(v string)`

SetGlobalReportTypeAccess sets GlobalReportTypeAccess field to given value.

### HasGlobalReportTypeAccess

`func (o *GetRole200Response) HasGlobalReportTypeAccess() bool`

HasGlobalReportTypeAccess returns a boolean if a field has been set.

### GetReportTypePermissions

`func (o *GetRole200Response) GetReportTypePermissions() []GetRole200ResponseReportTypePermissionsInner`

GetReportTypePermissions returns the ReportTypePermissions field if non-nil, zero value otherwise.

### GetReportTypePermissionsOk

`func (o *GetRole200Response) GetReportTypePermissionsOk() (*[]GetRole200ResponseReportTypePermissionsInner, bool)`

GetReportTypePermissionsOk returns a tuple with the ReportTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypePermissions

`func (o *GetRole200Response) SetReportTypePermissions(v []GetRole200ResponseReportTypePermissionsInner)`

SetReportTypePermissions sets ReportTypePermissions field to given value.

### HasReportTypePermissions

`func (o *GetRole200Response) HasReportTypePermissions() bool`

HasReportTypePermissions returns a boolean if a field has been set.

### GetGlobalTaskAccess

`func (o *GetRole200Response) GetGlobalTaskAccess() string`

GetGlobalTaskAccess returns the GlobalTaskAccess field if non-nil, zero value otherwise.

### GetGlobalTaskAccessOk

`func (o *GetRole200Response) GetGlobalTaskAccessOk() (*string, bool)`

GetGlobalTaskAccessOk returns a tuple with the GlobalTaskAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskAccess

`func (o *GetRole200Response) SetGlobalTaskAccess(v string)`

SetGlobalTaskAccess sets GlobalTaskAccess field to given value.

### HasGlobalTaskAccess

`func (o *GetRole200Response) HasGlobalTaskAccess() bool`

HasGlobalTaskAccess returns a boolean if a field has been set.

### GetTaskPermissions

`func (o *GetRole200Response) GetTaskPermissions() []GetRole200ResponseTaskPermissionsInner`

GetTaskPermissions returns the TaskPermissions field if non-nil, zero value otherwise.

### GetTaskPermissionsOk

`func (o *GetRole200Response) GetTaskPermissionsOk() (*[]GetRole200ResponseTaskPermissionsInner, bool)`

GetTaskPermissionsOk returns a tuple with the TaskPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPermissions

`func (o *GetRole200Response) SetTaskPermissions(v []GetRole200ResponseTaskPermissionsInner)`

SetTaskPermissions sets TaskPermissions field to given value.

### HasTaskPermissions

`func (o *GetRole200Response) HasTaskPermissions() bool`

HasTaskPermissions returns a boolean if a field has been set.

### GetGlobalTaskSetAccess

`func (o *GetRole200Response) GetGlobalTaskSetAccess() string`

GetGlobalTaskSetAccess returns the GlobalTaskSetAccess field if non-nil, zero value otherwise.

### GetGlobalTaskSetAccessOk

`func (o *GetRole200Response) GetGlobalTaskSetAccessOk() (*string, bool)`

GetGlobalTaskSetAccessOk returns a tuple with the GlobalTaskSetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskSetAccess

`func (o *GetRole200Response) SetGlobalTaskSetAccess(v string)`

SetGlobalTaskSetAccess sets GlobalTaskSetAccess field to given value.

### HasGlobalTaskSetAccess

`func (o *GetRole200Response) HasGlobalTaskSetAccess() bool`

HasGlobalTaskSetAccess returns a boolean if a field has been set.

### GetTaskSetPermissions

`func (o *GetRole200Response) GetTaskSetPermissions() []GetRole200ResponseTaskSetPermissionsInner`

GetTaskSetPermissions returns the TaskSetPermissions field if non-nil, zero value otherwise.

### GetTaskSetPermissionsOk

`func (o *GetRole200Response) GetTaskSetPermissionsOk() (*[]GetRole200ResponseTaskSetPermissionsInner, bool)`

GetTaskSetPermissionsOk returns a tuple with the TaskSetPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetPermissions

`func (o *GetRole200Response) SetTaskSetPermissions(v []GetRole200ResponseTaskSetPermissionsInner)`

SetTaskSetPermissions sets TaskSetPermissions field to given value.

### HasTaskSetPermissions

`func (o *GetRole200Response) HasTaskSetPermissions() bool`

HasTaskSetPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


