# Blueprint

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

### NewBlueprint

`func NewBlueprint() *Blueprint`

NewBlueprint instantiates a new Blueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintWithDefaults

`func NewBlueprintWithDefaults() *Blueprint`

NewBlueprintWithDefaults instantiates a new Blueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Blueprint) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Blueprint) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Blueprint) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Blueprint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Blueprint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Blueprint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Blueprint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Blueprint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *Blueprint) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Blueprint) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Blueprint) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Blueprint) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *Blueprint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Blueprint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Blueprint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Blueprint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *Blueprint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Blueprint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Blueprint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Blueprint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *Blueprint) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Blueprint) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Blueprint) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Blueprint) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *Blueprint) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Blueprint) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Blueprint) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Blueprint) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *Blueprint) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Blueprint) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Blueprint) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Blueprint) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *Blueprint) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *Blueprint) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *Blueprint) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *Blueprint) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *Blueprint) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Blueprint) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Blueprint) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Blueprint) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *Blueprint) GetTenant() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Blueprint) GetTenantOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Blueprint) SetTenant(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Blueprint) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


