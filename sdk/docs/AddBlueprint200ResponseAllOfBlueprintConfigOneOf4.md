# AddBlueprint200ResponseAllOfBlueprintConfigOneOf4

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

### NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf4

`func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf4() *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4`

NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf4 instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf4WithDefaults

`func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf4WithDefaults() *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4`

NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf4WithDefaults instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetConfig() AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetConfigOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetConfig(v AddBlueprint200ResponseAllOfBlueprintConfigOneOf4Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf4) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


