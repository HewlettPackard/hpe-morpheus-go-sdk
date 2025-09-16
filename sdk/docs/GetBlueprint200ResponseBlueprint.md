# GetBlueprint200ResponseBlueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**ResourcePermission** | Pointer to **map[string]interface{}** |  | [optional] 
**Owner** | Pointer to [**ListActivity200ResponseAllOfActivityInnerActivityInnerUser**](ListActivity200ResponseAllOfActivityInnerActivityInnerUser.md) |  | [optional] 
**Tenant** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewGetBlueprint200ResponseBlueprint

`func NewGetBlueprint200ResponseBlueprint() *GetBlueprint200ResponseBlueprint`

NewGetBlueprint200ResponseBlueprint instantiates a new GetBlueprint200ResponseBlueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBlueprint200ResponseBlueprintWithDefaults

`func NewGetBlueprint200ResponseBlueprintWithDefaults() *GetBlueprint200ResponseBlueprint`

NewGetBlueprint200ResponseBlueprintWithDefaults instantiates a new GetBlueprint200ResponseBlueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetBlueprint200ResponseBlueprint) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBlueprint200ResponseBlueprint) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBlueprint200ResponseBlueprint) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetBlueprint200ResponseBlueprint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetBlueprint200ResponseBlueprint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetBlueprint200ResponseBlueprint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetBlueprint200ResponseBlueprint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetBlueprint200ResponseBlueprint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetBlueprint200ResponseBlueprint) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetBlueprint200ResponseBlueprint) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetBlueprint200ResponseBlueprint) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetBlueprint200ResponseBlueprint) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *GetBlueprint200ResponseBlueprint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetBlueprint200ResponseBlueprint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetBlueprint200ResponseBlueprint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetBlueprint200ResponseBlueprint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *GetBlueprint200ResponseBlueprint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetBlueprint200ResponseBlueprint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetBlueprint200ResponseBlueprint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetBlueprint200ResponseBlueprint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetBlueprint200ResponseBlueprint) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetBlueprint200ResponseBlueprint) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *GetBlueprint200ResponseBlueprint) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetBlueprint200ResponseBlueprint) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetBlueprint200ResponseBlueprint) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetBlueprint200ResponseBlueprint) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetBlueprint200ResponseBlueprint) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetBlueprint200ResponseBlueprint) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetConfig

`func (o *GetBlueprint200ResponseBlueprint) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetBlueprint200ResponseBlueprint) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetBlueprint200ResponseBlueprint) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetBlueprint200ResponseBlueprint) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *GetBlueprint200ResponseBlueprint) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetBlueprint200ResponseBlueprint) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetBlueprint200ResponseBlueprint) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetBlueprint200ResponseBlueprint) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *GetBlueprint200ResponseBlueprint) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *GetBlueprint200ResponseBlueprint) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *GetBlueprint200ResponseBlueprint) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *GetBlueprint200ResponseBlueprint) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### SetResourcePermissionNil

`func (o *GetBlueprint200ResponseBlueprint) SetResourcePermissionNil(b bool)`

 SetResourcePermissionNil sets the value for ResourcePermission to be an explicit nil

### UnsetResourcePermission
`func (o *GetBlueprint200ResponseBlueprint) UnsetResourcePermission()`

UnsetResourcePermission ensures that no value is present for ResourcePermission, not even an explicit nil
### GetOwner

`func (o *GetBlueprint200ResponseBlueprint) GetOwner() ListActivity200ResponseAllOfActivityInnerActivityInnerUser`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetBlueprint200ResponseBlueprint) GetOwnerOk() (*ListActivity200ResponseAllOfActivityInnerActivityInnerUser, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetBlueprint200ResponseBlueprint) SetOwner(v ListActivity200ResponseAllOfActivityInnerActivityInnerUser)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetBlueprint200ResponseBlueprint) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *GetBlueprint200ResponseBlueprint) GetTenant() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *GetBlueprint200ResponseBlueprint) GetTenantOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *GetBlueprint200ResponseBlueprint) SetTenant(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *GetBlueprint200ResponseBlueprint) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


