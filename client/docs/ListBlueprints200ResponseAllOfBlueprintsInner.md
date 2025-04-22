# ListBlueprints200ResponseAllOfBlueprintsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**ResourcePermission** | Pointer to **map[string]interface{}** |  | [optional] 
**Owner** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**Tenant** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewListBlueprints200ResponseAllOfBlueprintsInner

`func NewListBlueprints200ResponseAllOfBlueprintsInner() *ListBlueprints200ResponseAllOfBlueprintsInner`

NewListBlueprints200ResponseAllOfBlueprintsInner instantiates a new ListBlueprints200ResponseAllOfBlueprintsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBlueprints200ResponseAllOfBlueprintsInnerWithDefaults

`func NewListBlueprints200ResponseAllOfBlueprintsInnerWithDefaults() *ListBlueprints200ResponseAllOfBlueprintsInner`

NewListBlueprints200ResponseAllOfBlueprintsInnerWithDefaults instantiates a new ListBlueprints200ResponseAllOfBlueprintsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetTenant() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) GetTenantOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) SetTenant(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ListBlueprints200ResponseAllOfBlueprintsInner) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


