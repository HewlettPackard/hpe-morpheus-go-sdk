# UpdateRole200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**UpdateRole200ResponseAllOfRole**](UpdateRole200ResponseAllOfRole.md) |  | [optional] 
**FeaturePermissions** | Pointer to [**[]UpdateRole200ResponseAllOfFeaturePermissionsInner**](UpdateRole200ResponseAllOfFeaturePermissionsInner.md) |  | [optional] 
**GlobalSiteAccess** | Pointer to **string** |  | [optional] 
**Sites** | Pointer to [**[]UpdateRole200ResponseAllOfSitesInner**](UpdateRole200ResponseAllOfSitesInner.md) |  | [optional] 
**GlobalZoneAccess** | Pointer to **string** |  | [optional] 
**Zones** | Pointer to [**[]UpdateRole200ResponseAllOfZonesInner**](UpdateRole200ResponseAllOfZonesInner.md) |  | [optional] 
**GlobalInstanceTypeAccess** | Pointer to **string** |  | [optional] 
**InstanceTypePermissions** | Pointer to [**[]UpdateRole200ResponseAllOfInstanceTypePermissionsInner**](UpdateRole200ResponseAllOfInstanceTypePermissionsInner.md) |  | [optional] 
**GlobalAppTemplateAccess** | Pointer to **string** |  | [optional] 
**AppTemplatePermissions** | Pointer to [**[]UpdateRole200ResponseAllOfAppTemplatePermissionsInner**](UpdateRole200ResponseAllOfAppTemplatePermissionsInner.md) |  | [optional] 
**GlobalCatalogItemTypeAccess** | Pointer to **string** |  | [optional] 
**CatalogItemTypePermissions** | Pointer to [**[]UpdateRole200ResponseAllOfCatalogItemTypePermissionsInner**](UpdateRole200ResponseAllOfCatalogItemTypePermissionsInner.md) |  | [optional] 
**GlobalPersonaAccess** | Pointer to **string** |  | [optional] 
**PersonaPermissions** | Pointer to [**[]UpdateRole200ResponseAllOfPersonaPermissionsInner**](UpdateRole200ResponseAllOfPersonaPermissionsInner.md) |  | [optional] 
**GlobalVdiPoolAccess** | Pointer to **string** |  | [optional] 
**VdiPoolPermissions** | Pointer to [**[]UpdateRole200ResponseAllOfVdiPoolPermissionsInner**](UpdateRole200ResponseAllOfVdiPoolPermissionsInner.md) |  | [optional] 
**GlobalReportTypeAccess** | Pointer to **string** |  | [optional] 
**ReportTypePermissions** | Pointer to [**[]UpdateRole200ResponseAllOfReportTypePermissionsInner**](UpdateRole200ResponseAllOfReportTypePermissionsInner.md) |  | [optional] 
**GlobalTaskAccess** | Pointer to **string** |  | [optional] 
**TaskPermissions** | Pointer to [**[]UpdateRole200ResponseAllOfTaskPermissionsInner**](UpdateRole200ResponseAllOfTaskPermissionsInner.md) |  | [optional] 
**GlobalTaskSetAccess** | Pointer to **string** |  | [optional] 
**TaskSetPermissions** | Pointer to [**[]UpdateRole200ResponseAllOfTaskSetPermissionsInner**](UpdateRole200ResponseAllOfTaskSetPermissionsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateRole200Response

`func NewUpdateRole200Response() *UpdateRole200Response`

NewUpdateRole200Response instantiates a new UpdateRole200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRole200ResponseWithDefaults

`func NewUpdateRole200ResponseWithDefaults() *UpdateRole200Response`

NewUpdateRole200ResponseWithDefaults instantiates a new UpdateRole200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *UpdateRole200Response) GetRole() UpdateRole200ResponseAllOfRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateRole200Response) GetRoleOk() (*UpdateRole200ResponseAllOfRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateRole200Response) SetRole(v UpdateRole200ResponseAllOfRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateRole200Response) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetFeaturePermissions

`func (o *UpdateRole200Response) GetFeaturePermissions() []UpdateRole200ResponseAllOfFeaturePermissionsInner`

GetFeaturePermissions returns the FeaturePermissions field if non-nil, zero value otherwise.

### GetFeaturePermissionsOk

`func (o *UpdateRole200Response) GetFeaturePermissionsOk() (*[]UpdateRole200ResponseAllOfFeaturePermissionsInner, bool)`

GetFeaturePermissionsOk returns a tuple with the FeaturePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturePermissions

`func (o *UpdateRole200Response) SetFeaturePermissions(v []UpdateRole200ResponseAllOfFeaturePermissionsInner)`

SetFeaturePermissions sets FeaturePermissions field to given value.

### HasFeaturePermissions

`func (o *UpdateRole200Response) HasFeaturePermissions() bool`

HasFeaturePermissions returns a boolean if a field has been set.

### GetGlobalSiteAccess

`func (o *UpdateRole200Response) GetGlobalSiteAccess() string`

GetGlobalSiteAccess returns the GlobalSiteAccess field if non-nil, zero value otherwise.

### GetGlobalSiteAccessOk

`func (o *UpdateRole200Response) GetGlobalSiteAccessOk() (*string, bool)`

GetGlobalSiteAccessOk returns a tuple with the GlobalSiteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSiteAccess

`func (o *UpdateRole200Response) SetGlobalSiteAccess(v string)`

SetGlobalSiteAccess sets GlobalSiteAccess field to given value.

### HasGlobalSiteAccess

`func (o *UpdateRole200Response) HasGlobalSiteAccess() bool`

HasGlobalSiteAccess returns a boolean if a field has been set.

### GetSites

`func (o *UpdateRole200Response) GetSites() []UpdateRole200ResponseAllOfSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UpdateRole200Response) GetSitesOk() (*[]UpdateRole200ResponseAllOfSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UpdateRole200Response) SetSites(v []UpdateRole200ResponseAllOfSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *UpdateRole200Response) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetGlobalZoneAccess

`func (o *UpdateRole200Response) GetGlobalZoneAccess() string`

GetGlobalZoneAccess returns the GlobalZoneAccess field if non-nil, zero value otherwise.

### GetGlobalZoneAccessOk

`func (o *UpdateRole200Response) GetGlobalZoneAccessOk() (*string, bool)`

GetGlobalZoneAccessOk returns a tuple with the GlobalZoneAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalZoneAccess

`func (o *UpdateRole200Response) SetGlobalZoneAccess(v string)`

SetGlobalZoneAccess sets GlobalZoneAccess field to given value.

### HasGlobalZoneAccess

`func (o *UpdateRole200Response) HasGlobalZoneAccess() bool`

HasGlobalZoneAccess returns a boolean if a field has been set.

### GetZones

`func (o *UpdateRole200Response) GetZones() []UpdateRole200ResponseAllOfZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *UpdateRole200Response) GetZonesOk() (*[]UpdateRole200ResponseAllOfZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *UpdateRole200Response) SetZones(v []UpdateRole200ResponseAllOfZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *UpdateRole200Response) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetGlobalInstanceTypeAccess

`func (o *UpdateRole200Response) GetGlobalInstanceTypeAccess() string`

GetGlobalInstanceTypeAccess returns the GlobalInstanceTypeAccess field if non-nil, zero value otherwise.

### GetGlobalInstanceTypeAccessOk

`func (o *UpdateRole200Response) GetGlobalInstanceTypeAccessOk() (*string, bool)`

GetGlobalInstanceTypeAccessOk returns a tuple with the GlobalInstanceTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalInstanceTypeAccess

`func (o *UpdateRole200Response) SetGlobalInstanceTypeAccess(v string)`

SetGlobalInstanceTypeAccess sets GlobalInstanceTypeAccess field to given value.

### HasGlobalInstanceTypeAccess

`func (o *UpdateRole200Response) HasGlobalInstanceTypeAccess() bool`

HasGlobalInstanceTypeAccess returns a boolean if a field has been set.

### GetInstanceTypePermissions

`func (o *UpdateRole200Response) GetInstanceTypePermissions() []UpdateRole200ResponseAllOfInstanceTypePermissionsInner`

GetInstanceTypePermissions returns the InstanceTypePermissions field if non-nil, zero value otherwise.

### GetInstanceTypePermissionsOk

`func (o *UpdateRole200Response) GetInstanceTypePermissionsOk() (*[]UpdateRole200ResponseAllOfInstanceTypePermissionsInner, bool)`

GetInstanceTypePermissionsOk returns a tuple with the InstanceTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypePermissions

`func (o *UpdateRole200Response) SetInstanceTypePermissions(v []UpdateRole200ResponseAllOfInstanceTypePermissionsInner)`

SetInstanceTypePermissions sets InstanceTypePermissions field to given value.

### HasInstanceTypePermissions

`func (o *UpdateRole200Response) HasInstanceTypePermissions() bool`

HasInstanceTypePermissions returns a boolean if a field has been set.

### GetGlobalAppTemplateAccess

`func (o *UpdateRole200Response) GetGlobalAppTemplateAccess() string`

GetGlobalAppTemplateAccess returns the GlobalAppTemplateAccess field if non-nil, zero value otherwise.

### GetGlobalAppTemplateAccessOk

`func (o *UpdateRole200Response) GetGlobalAppTemplateAccessOk() (*string, bool)`

GetGlobalAppTemplateAccessOk returns a tuple with the GlobalAppTemplateAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAppTemplateAccess

`func (o *UpdateRole200Response) SetGlobalAppTemplateAccess(v string)`

SetGlobalAppTemplateAccess sets GlobalAppTemplateAccess field to given value.

### HasGlobalAppTemplateAccess

`func (o *UpdateRole200Response) HasGlobalAppTemplateAccess() bool`

HasGlobalAppTemplateAccess returns a boolean if a field has been set.

### GetAppTemplatePermissions

`func (o *UpdateRole200Response) GetAppTemplatePermissions() []UpdateRole200ResponseAllOfAppTemplatePermissionsInner`

GetAppTemplatePermissions returns the AppTemplatePermissions field if non-nil, zero value otherwise.

### GetAppTemplatePermissionsOk

`func (o *UpdateRole200Response) GetAppTemplatePermissionsOk() (*[]UpdateRole200ResponseAllOfAppTemplatePermissionsInner, bool)`

GetAppTemplatePermissionsOk returns a tuple with the AppTemplatePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplatePermissions

`func (o *UpdateRole200Response) SetAppTemplatePermissions(v []UpdateRole200ResponseAllOfAppTemplatePermissionsInner)`

SetAppTemplatePermissions sets AppTemplatePermissions field to given value.

### HasAppTemplatePermissions

`func (o *UpdateRole200Response) HasAppTemplatePermissions() bool`

HasAppTemplatePermissions returns a boolean if a field has been set.

### GetGlobalCatalogItemTypeAccess

`func (o *UpdateRole200Response) GetGlobalCatalogItemTypeAccess() string`

GetGlobalCatalogItemTypeAccess returns the GlobalCatalogItemTypeAccess field if non-nil, zero value otherwise.

### GetGlobalCatalogItemTypeAccessOk

`func (o *UpdateRole200Response) GetGlobalCatalogItemTypeAccessOk() (*string, bool)`

GetGlobalCatalogItemTypeAccessOk returns a tuple with the GlobalCatalogItemTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCatalogItemTypeAccess

`func (o *UpdateRole200Response) SetGlobalCatalogItemTypeAccess(v string)`

SetGlobalCatalogItemTypeAccess sets GlobalCatalogItemTypeAccess field to given value.

### HasGlobalCatalogItemTypeAccess

`func (o *UpdateRole200Response) HasGlobalCatalogItemTypeAccess() bool`

HasGlobalCatalogItemTypeAccess returns a boolean if a field has been set.

### GetCatalogItemTypePermissions

`func (o *UpdateRole200Response) GetCatalogItemTypePermissions() []UpdateRole200ResponseAllOfCatalogItemTypePermissionsInner`

GetCatalogItemTypePermissions returns the CatalogItemTypePermissions field if non-nil, zero value otherwise.

### GetCatalogItemTypePermissionsOk

`func (o *UpdateRole200Response) GetCatalogItemTypePermissionsOk() (*[]UpdateRole200ResponseAllOfCatalogItemTypePermissionsInner, bool)`

GetCatalogItemTypePermissionsOk returns a tuple with the CatalogItemTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypePermissions

`func (o *UpdateRole200Response) SetCatalogItemTypePermissions(v []UpdateRole200ResponseAllOfCatalogItemTypePermissionsInner)`

SetCatalogItemTypePermissions sets CatalogItemTypePermissions field to given value.

### HasCatalogItemTypePermissions

`func (o *UpdateRole200Response) HasCatalogItemTypePermissions() bool`

HasCatalogItemTypePermissions returns a boolean if a field has been set.

### GetGlobalPersonaAccess

`func (o *UpdateRole200Response) GetGlobalPersonaAccess() string`

GetGlobalPersonaAccess returns the GlobalPersonaAccess field if non-nil, zero value otherwise.

### GetGlobalPersonaAccessOk

`func (o *UpdateRole200Response) GetGlobalPersonaAccessOk() (*string, bool)`

GetGlobalPersonaAccessOk returns a tuple with the GlobalPersonaAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPersonaAccess

`func (o *UpdateRole200Response) SetGlobalPersonaAccess(v string)`

SetGlobalPersonaAccess sets GlobalPersonaAccess field to given value.

### HasGlobalPersonaAccess

`func (o *UpdateRole200Response) HasGlobalPersonaAccess() bool`

HasGlobalPersonaAccess returns a boolean if a field has been set.

### GetPersonaPermissions

`func (o *UpdateRole200Response) GetPersonaPermissions() []UpdateRole200ResponseAllOfPersonaPermissionsInner`

GetPersonaPermissions returns the PersonaPermissions field if non-nil, zero value otherwise.

### GetPersonaPermissionsOk

`func (o *UpdateRole200Response) GetPersonaPermissionsOk() (*[]UpdateRole200ResponseAllOfPersonaPermissionsInner, bool)`

GetPersonaPermissionsOk returns a tuple with the PersonaPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaPermissions

`func (o *UpdateRole200Response) SetPersonaPermissions(v []UpdateRole200ResponseAllOfPersonaPermissionsInner)`

SetPersonaPermissions sets PersonaPermissions field to given value.

### HasPersonaPermissions

`func (o *UpdateRole200Response) HasPersonaPermissions() bool`

HasPersonaPermissions returns a boolean if a field has been set.

### GetGlobalVdiPoolAccess

`func (o *UpdateRole200Response) GetGlobalVdiPoolAccess() string`

GetGlobalVdiPoolAccess returns the GlobalVdiPoolAccess field if non-nil, zero value otherwise.

### GetGlobalVdiPoolAccessOk

`func (o *UpdateRole200Response) GetGlobalVdiPoolAccessOk() (*string, bool)`

GetGlobalVdiPoolAccessOk returns a tuple with the GlobalVdiPoolAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalVdiPoolAccess

`func (o *UpdateRole200Response) SetGlobalVdiPoolAccess(v string)`

SetGlobalVdiPoolAccess sets GlobalVdiPoolAccess field to given value.

### HasGlobalVdiPoolAccess

`func (o *UpdateRole200Response) HasGlobalVdiPoolAccess() bool`

HasGlobalVdiPoolAccess returns a boolean if a field has been set.

### GetVdiPoolPermissions

`func (o *UpdateRole200Response) GetVdiPoolPermissions() []UpdateRole200ResponseAllOfVdiPoolPermissionsInner`

GetVdiPoolPermissions returns the VdiPoolPermissions field if non-nil, zero value otherwise.

### GetVdiPoolPermissionsOk

`func (o *UpdateRole200Response) GetVdiPoolPermissionsOk() (*[]UpdateRole200ResponseAllOfVdiPoolPermissionsInner, bool)`

GetVdiPoolPermissionsOk returns a tuple with the VdiPoolPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPoolPermissions

`func (o *UpdateRole200Response) SetVdiPoolPermissions(v []UpdateRole200ResponseAllOfVdiPoolPermissionsInner)`

SetVdiPoolPermissions sets VdiPoolPermissions field to given value.

### HasVdiPoolPermissions

`func (o *UpdateRole200Response) HasVdiPoolPermissions() bool`

HasVdiPoolPermissions returns a boolean if a field has been set.

### GetGlobalReportTypeAccess

`func (o *UpdateRole200Response) GetGlobalReportTypeAccess() string`

GetGlobalReportTypeAccess returns the GlobalReportTypeAccess field if non-nil, zero value otherwise.

### GetGlobalReportTypeAccessOk

`func (o *UpdateRole200Response) GetGlobalReportTypeAccessOk() (*string, bool)`

GetGlobalReportTypeAccessOk returns a tuple with the GlobalReportTypeAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalReportTypeAccess

`func (o *UpdateRole200Response) SetGlobalReportTypeAccess(v string)`

SetGlobalReportTypeAccess sets GlobalReportTypeAccess field to given value.

### HasGlobalReportTypeAccess

`func (o *UpdateRole200Response) HasGlobalReportTypeAccess() bool`

HasGlobalReportTypeAccess returns a boolean if a field has been set.

### GetReportTypePermissions

`func (o *UpdateRole200Response) GetReportTypePermissions() []UpdateRole200ResponseAllOfReportTypePermissionsInner`

GetReportTypePermissions returns the ReportTypePermissions field if non-nil, zero value otherwise.

### GetReportTypePermissionsOk

`func (o *UpdateRole200Response) GetReportTypePermissionsOk() (*[]UpdateRole200ResponseAllOfReportTypePermissionsInner, bool)`

GetReportTypePermissionsOk returns a tuple with the ReportTypePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypePermissions

`func (o *UpdateRole200Response) SetReportTypePermissions(v []UpdateRole200ResponseAllOfReportTypePermissionsInner)`

SetReportTypePermissions sets ReportTypePermissions field to given value.

### HasReportTypePermissions

`func (o *UpdateRole200Response) HasReportTypePermissions() bool`

HasReportTypePermissions returns a boolean if a field has been set.

### GetGlobalTaskAccess

`func (o *UpdateRole200Response) GetGlobalTaskAccess() string`

GetGlobalTaskAccess returns the GlobalTaskAccess field if non-nil, zero value otherwise.

### GetGlobalTaskAccessOk

`func (o *UpdateRole200Response) GetGlobalTaskAccessOk() (*string, bool)`

GetGlobalTaskAccessOk returns a tuple with the GlobalTaskAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskAccess

`func (o *UpdateRole200Response) SetGlobalTaskAccess(v string)`

SetGlobalTaskAccess sets GlobalTaskAccess field to given value.

### HasGlobalTaskAccess

`func (o *UpdateRole200Response) HasGlobalTaskAccess() bool`

HasGlobalTaskAccess returns a boolean if a field has been set.

### GetTaskPermissions

`func (o *UpdateRole200Response) GetTaskPermissions() []UpdateRole200ResponseAllOfTaskPermissionsInner`

GetTaskPermissions returns the TaskPermissions field if non-nil, zero value otherwise.

### GetTaskPermissionsOk

`func (o *UpdateRole200Response) GetTaskPermissionsOk() (*[]UpdateRole200ResponseAllOfTaskPermissionsInner, bool)`

GetTaskPermissionsOk returns a tuple with the TaskPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPermissions

`func (o *UpdateRole200Response) SetTaskPermissions(v []UpdateRole200ResponseAllOfTaskPermissionsInner)`

SetTaskPermissions sets TaskPermissions field to given value.

### HasTaskPermissions

`func (o *UpdateRole200Response) HasTaskPermissions() bool`

HasTaskPermissions returns a boolean if a field has been set.

### GetGlobalTaskSetAccess

`func (o *UpdateRole200Response) GetGlobalTaskSetAccess() string`

GetGlobalTaskSetAccess returns the GlobalTaskSetAccess field if non-nil, zero value otherwise.

### GetGlobalTaskSetAccessOk

`func (o *UpdateRole200Response) GetGlobalTaskSetAccessOk() (*string, bool)`

GetGlobalTaskSetAccessOk returns a tuple with the GlobalTaskSetAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTaskSetAccess

`func (o *UpdateRole200Response) SetGlobalTaskSetAccess(v string)`

SetGlobalTaskSetAccess sets GlobalTaskSetAccess field to given value.

### HasGlobalTaskSetAccess

`func (o *UpdateRole200Response) HasGlobalTaskSetAccess() bool`

HasGlobalTaskSetAccess returns a boolean if a field has been set.

### GetTaskSetPermissions

`func (o *UpdateRole200Response) GetTaskSetPermissions() []UpdateRole200ResponseAllOfTaskSetPermissionsInner`

GetTaskSetPermissions returns the TaskSetPermissions field if non-nil, zero value otherwise.

### GetTaskSetPermissionsOk

`func (o *UpdateRole200Response) GetTaskSetPermissionsOk() (*[]UpdateRole200ResponseAllOfTaskSetPermissionsInner, bool)`

GetTaskSetPermissionsOk returns a tuple with the TaskSetPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetPermissions

`func (o *UpdateRole200Response) SetTaskSetPermissions(v []UpdateRole200ResponseAllOfTaskSetPermissionsInner)`

SetTaskSetPermissions sets TaskSetPermissions field to given value.

### HasTaskSetPermissions

`func (o *UpdateRole200Response) HasTaskSetPermissions() bool`

HasTaskSetPermissions returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateRole200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateRole200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateRole200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateRole200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


