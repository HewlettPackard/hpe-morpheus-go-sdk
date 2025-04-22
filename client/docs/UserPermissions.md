# UserPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Zones** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Sites** | Pointer to **[]map[string]interface{}** |  | [optional] 
**InstanceTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**AppTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CatalogItemTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Personas** | Pointer to **[]map[string]interface{}** |  | [optional] 
**VdiPools** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ReportTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tasks** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TaskSets** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewUserPermissions

`func NewUserPermissions() *UserPermissions`

NewUserPermissions instantiates a new UserPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPermissionsWithDefaults

`func NewUserPermissionsWithDefaults() *UserPermissions`

NewUserPermissionsWithDefaults instantiates a new UserPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *UserPermissions) GetFeatures() []map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UserPermissions) GetFeaturesOk() (*[]map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *UserPermissions) SetFeatures(v []map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *UserPermissions) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetZones

`func (o *UserPermissions) GetZones() []map[string]interface{}`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *UserPermissions) GetZonesOk() (*[]map[string]interface{}, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *UserPermissions) SetZones(v []map[string]interface{})`

SetZones sets Zones field to given value.

### HasZones

`func (o *UserPermissions) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetSites

`func (o *UserPermissions) GetSites() []map[string]interface{}`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *UserPermissions) GetSitesOk() (*[]map[string]interface{}, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *UserPermissions) SetSites(v []map[string]interface{})`

SetSites sets Sites field to given value.

### HasSites

`func (o *UserPermissions) HasSites() bool`

HasSites returns a boolean if a field has been set.

### GetInstanceTypes

`func (o *UserPermissions) GetInstanceTypes() []map[string]interface{}`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *UserPermissions) GetInstanceTypesOk() (*[]map[string]interface{}, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *UserPermissions) SetInstanceTypes(v []map[string]interface{})`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *UserPermissions) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetAppTemplates

`func (o *UserPermissions) GetAppTemplates() []map[string]interface{}`

GetAppTemplates returns the AppTemplates field if non-nil, zero value otherwise.

### GetAppTemplatesOk

`func (o *UserPermissions) GetAppTemplatesOk() (*[]map[string]interface{}, bool)`

GetAppTemplatesOk returns a tuple with the AppTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppTemplates

`func (o *UserPermissions) SetAppTemplates(v []map[string]interface{})`

SetAppTemplates sets AppTemplates field to given value.

### HasAppTemplates

`func (o *UserPermissions) HasAppTemplates() bool`

HasAppTemplates returns a boolean if a field has been set.

### GetCatalogItemTypes

`func (o *UserPermissions) GetCatalogItemTypes() []map[string]interface{}`

GetCatalogItemTypes returns the CatalogItemTypes field if non-nil, zero value otherwise.

### GetCatalogItemTypesOk

`func (o *UserPermissions) GetCatalogItemTypesOk() (*[]map[string]interface{}, bool)`

GetCatalogItemTypesOk returns a tuple with the CatalogItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypes

`func (o *UserPermissions) SetCatalogItemTypes(v []map[string]interface{})`

SetCatalogItemTypes sets CatalogItemTypes field to given value.

### HasCatalogItemTypes

`func (o *UserPermissions) HasCatalogItemTypes() bool`

HasCatalogItemTypes returns a boolean if a field has been set.

### GetPersonas

`func (o *UserPermissions) GetPersonas() []map[string]interface{}`

GetPersonas returns the Personas field if non-nil, zero value otherwise.

### GetPersonasOk

`func (o *UserPermissions) GetPersonasOk() (*[]map[string]interface{}, bool)`

GetPersonasOk returns a tuple with the Personas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonas

`func (o *UserPermissions) SetPersonas(v []map[string]interface{})`

SetPersonas sets Personas field to given value.

### HasPersonas

`func (o *UserPermissions) HasPersonas() bool`

HasPersonas returns a boolean if a field has been set.

### GetVdiPools

`func (o *UserPermissions) GetVdiPools() []map[string]interface{}`

GetVdiPools returns the VdiPools field if non-nil, zero value otherwise.

### GetVdiPoolsOk

`func (o *UserPermissions) GetVdiPoolsOk() (*[]map[string]interface{}, bool)`

GetVdiPoolsOk returns a tuple with the VdiPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPools

`func (o *UserPermissions) SetVdiPools(v []map[string]interface{})`

SetVdiPools sets VdiPools field to given value.

### HasVdiPools

`func (o *UserPermissions) HasVdiPools() bool`

HasVdiPools returns a boolean if a field has been set.

### GetReportTypes

`func (o *UserPermissions) GetReportTypes() []map[string]interface{}`

GetReportTypes returns the ReportTypes field if non-nil, zero value otherwise.

### GetReportTypesOk

`func (o *UserPermissions) GetReportTypesOk() (*[]map[string]interface{}, bool)`

GetReportTypesOk returns a tuple with the ReportTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTypes

`func (o *UserPermissions) SetReportTypes(v []map[string]interface{})`

SetReportTypes sets ReportTypes field to given value.

### HasReportTypes

`func (o *UserPermissions) HasReportTypes() bool`

HasReportTypes returns a boolean if a field has been set.

### GetTasks

`func (o *UserPermissions) GetTasks() []map[string]interface{}`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *UserPermissions) GetTasksOk() (*[]map[string]interface{}, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *UserPermissions) SetTasks(v []map[string]interface{})`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *UserPermissions) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTaskSets

`func (o *UserPermissions) GetTaskSets() []map[string]interface{}`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *UserPermissions) GetTaskSetsOk() (*[]map[string]interface{}, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *UserPermissions) SetTaskSets(v []map[string]interface{})`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *UserPermissions) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


