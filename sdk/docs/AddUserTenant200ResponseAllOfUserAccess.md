# AddUserTenant200ResponseAllOfUserAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]AddUserTenant200ResponseAllOfUserAccessFeaturesInner**](AddUserTenant200ResponseAllOfUserAccessFeaturesInner.md) |  | [optional] 
**Zones** | Pointer to [**[]AddRoles200ResponseAllOfSitesInner**](AddRoles200ResponseAllOfSitesInner.md) |  | [optional] 
**Sites** | Pointer to [**[]AddRoles200ResponseAllOfSitesInner**](AddRoles200ResponseAllOfSitesInner.md) |  | [optional] 
**InstanceTypes** | Pointer to [**[]AddRoles200ResponseAllOfInstanceTypePermissionsInner**](AddRoles200ResponseAllOfInstanceTypePermissionsInner.md) |  | [optional] 
**AppTemplates** | Pointer to [**[]AddRoles200ResponseAllOfSitesInner**](AddRoles200ResponseAllOfSitesInner.md) |  | [optional] 
**CatalogItemTypes** | Pointer to [**[]AddRoles200ResponseAllOfSitesInner**](AddRoles200ResponseAllOfSitesInner.md) |  | [optional] 
**Personas** | Pointer to [**[]AddRoles200ResponseAllOfInstanceTypePermissionsInner**](AddRoles200ResponseAllOfInstanceTypePermissionsInner.md) |  | [optional] 
**VdiPools** | Pointer to [**[]AddRoles200ResponseAllOfSitesInner**](AddRoles200ResponseAllOfSitesInner.md) |  | [optional] 
**ReportTypes** | Pointer to [**[]AddRoles200ResponseAllOfInstanceTypePermissionsInner**](AddRoles200ResponseAllOfInstanceTypePermissionsInner.md) |  | [optional] 
**Tasks** | Pointer to [**[]AddRoles200ResponseAllOfAppTemplatePermissionsInner**](AddRoles200ResponseAllOfAppTemplatePermissionsInner.md) |  | [optional] 
**TaskSets** | Pointer to [**[]AddRoles200ResponseAllOfAppTemplatePermissionsInner**](AddRoles200ResponseAllOfAppTemplatePermissionsInner.md) |  | [optional] 

## Methods

### NewAddUserTenant200ResponseAllOfUserAccess

`func NewAddUserTenant200ResponseAllOfUserAccess() *AddUserTenant200ResponseAllOfUserAccess`

NewAddUserTenant200ResponseAllOfUserAccess instantiates a new AddUserTenant200ResponseAllOfUserAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserTenant200ResponseAllOfUserAccessWithDefaults

`func NewAddUserTenant200ResponseAllOfUserAccessWithDefaults() *AddUserTenant200ResponseAllOfUserAccess`

NewAddUserTenant200ResponseAllOfUserAccessWithDefaults instantiates a new AddUserTenant200ResponseAllOfUserAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetFeatures() []AddUserTenant200ResponseAllOfUserAccessFeaturesInner`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetFeaturesOk() (*[]AddUserTenant200ResponseAllOfUserAccessFeaturesInner, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetFeatures(v []AddUserTenant200ResponseAllOfUserAccessFeaturesInner)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetZones

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetZones() []AddRoles200ResponseAllOfSitesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetZonesOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetZones(v []AddRoles200ResponseAllOfSitesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetSites

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetSites() []AddRoles200ResponseAllOfSitesInner`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetSitesOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetSites(v []AddRoles200ResponseAllOfSitesInner)`

SetSites sets Sites field to given value.

### HasSites

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetInstanceTypes() []AddRoles200ResponseAllOfInstanceTypePermissionsInner`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetInstanceTypesOk() (*[]AddRoles200ResponseAllOfInstanceTypePermissionsInner, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetInstanceTypes(v []AddRoles200ResponseAllOfInstanceTypePermissionsInner)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetAppTemplates

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetAppTemplates() []AddRoles200ResponseAllOfSitesInner`

GetAppTemplates returns the AppTemplates field if non-nil, zero value otherwise.

### GetAppTemplatesOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetAppTemplatesOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetAppTemplatesOk returns a tuple with the AppTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplates

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetAppTemplates(v []AddRoles200ResponseAllOfSitesInner)`

SetAppTemplates sets AppTemplates field to given value.

### HasAppTemplates

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasAppTemplates() bool`

HasAppTemplates returns a boolean if a field has been set.

### GetCatalogItemTypes

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetCatalogItemTypes() []AddRoles200ResponseAllOfSitesInner`

GetCatalogItemTypes returns the CatalogItemTypes field if non-nil, zero value otherwise.

### GetCatalogItemTypesOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetCatalogItemTypesOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetCatalogItemTypesOk returns a tuple with the CatalogItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypes

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetCatalogItemTypes(v []AddRoles200ResponseAllOfSitesInner)`

SetCatalogItemTypes sets CatalogItemTypes field to given value.

### HasCatalogItemTypes

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasCatalogItemTypes() bool`

HasCatalogItemTypes returns a boolean if a field has been set.

### GetPersonas

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetPersonas() []AddRoles200ResponseAllOfInstanceTypePermissionsInner`

GetPersonas returns the Personas field if non-nil, zero value otherwise.

### GetPersonasOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetPersonasOk() (*[]AddRoles200ResponseAllOfInstanceTypePermissionsInner, bool)`

GetPersonasOk returns a tuple with the Personas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonas

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetPersonas(v []AddRoles200ResponseAllOfInstanceTypePermissionsInner)`

SetPersonas sets Personas field to given value.

### HasPersonas

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasPersonas() bool`

HasPersonas returns a boolean if a field has been set.

### GetVdiPools

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetVdiPools() []AddRoles200ResponseAllOfSitesInner`

GetVdiPools returns the VdiPools field if non-nil, zero value otherwise.

### GetVdiPoolsOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetVdiPoolsOk() (*[]AddRoles200ResponseAllOfSitesInner, bool)`

GetVdiPoolsOk returns a tuple with the VdiPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPools

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetVdiPools(v []AddRoles200ResponseAllOfSitesInner)`

SetVdiPools sets VdiPools field to given value.

### HasVdiPools

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasVdiPools() bool`

HasVdiPools returns a boolean if a field has been set.

### GetReportTypes

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetReportTypes() []AddRoles200ResponseAllOfInstanceTypePermissionsInner`

GetReportTypes returns the ReportTypes field if non-nil, zero value otherwise.

### GetReportTypesOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetReportTypesOk() (*[]AddRoles200ResponseAllOfInstanceTypePermissionsInner, bool)`

GetReportTypesOk returns a tuple with the ReportTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypes

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetReportTypes(v []AddRoles200ResponseAllOfInstanceTypePermissionsInner)`

SetReportTypes sets ReportTypes field to given value.

### HasReportTypes

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasReportTypes() bool`

HasReportTypes returns a boolean if a field has been set.

### GetTasks

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetTasks() []AddRoles200ResponseAllOfAppTemplatePermissionsInner`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetTasksOk() (*[]AddRoles200ResponseAllOfAppTemplatePermissionsInner, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetTasks(v []AddRoles200ResponseAllOfAppTemplatePermissionsInner)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTaskSets

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetTaskSets() []AddRoles200ResponseAllOfAppTemplatePermissionsInner`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *AddUserTenant200ResponseAllOfUserAccess) GetTaskSetsOk() (*[]AddRoles200ResponseAllOfAppTemplatePermissionsInner, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *AddUserTenant200ResponseAllOfUserAccess) SetTaskSets(v []AddRoles200ResponseAllOfAppTemplatePermissionsInner)`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *AddUserTenant200ResponseAllOfUserAccess) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


