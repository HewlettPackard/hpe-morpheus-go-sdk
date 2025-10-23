# BlueprintMorpheusCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**Config** | Pointer to [**AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config**](AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config.md) |  | [optional] 
**Visibility** | Pointer to **string** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | Pointer to **map[string]interface{}** | Resource Permission Block | [optional] 
**Owner** | Pointer to **map[string]interface{}** | Owner | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | Tenant | [optional] 

## Methods

### NewBlueprintMorpheusCreateSuccess

`func NewBlueprintMorpheusCreateSuccess() *BlueprintMorpheusCreateSuccess`

NewBlueprintMorpheusCreateSuccess instantiates a new BlueprintMorpheusCreateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintMorpheusCreateSuccessWithDefaults

`func NewBlueprintMorpheusCreateSuccessWithDefaults() *BlueprintMorpheusCreateSuccess`

NewBlueprintMorpheusCreateSuccessWithDefaults instantiates a new BlueprintMorpheusCreateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintMorpheusCreateSuccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintMorpheusCreateSuccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintMorpheusCreateSuccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintMorpheusCreateSuccess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *BlueprintMorpheusCreateSuccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintMorpheusCreateSuccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintMorpheusCreateSuccess) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintMorpheusCreateSuccess) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *BlueprintMorpheusCreateSuccess) GetConfig() AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BlueprintMorpheusCreateSuccess) GetConfigOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BlueprintMorpheusCreateSuccess) SetConfig(v AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BlueprintMorpheusCreateSuccess) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *BlueprintMorpheusCreateSuccess) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BlueprintMorpheusCreateSuccess) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BlueprintMorpheusCreateSuccess) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BlueprintMorpheusCreateSuccess) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *BlueprintMorpheusCreateSuccess) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *BlueprintMorpheusCreateSuccess) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *BlueprintMorpheusCreateSuccess) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *BlueprintMorpheusCreateSuccess) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *BlueprintMorpheusCreateSuccess) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BlueprintMorpheusCreateSuccess) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BlueprintMorpheusCreateSuccess) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BlueprintMorpheusCreateSuccess) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *BlueprintMorpheusCreateSuccess) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BlueprintMorpheusCreateSuccess) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BlueprintMorpheusCreateSuccess) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BlueprintMorpheusCreateSuccess) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


