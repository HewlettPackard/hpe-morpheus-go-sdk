# BlueprintCreateSuccessConfigOneOf4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**Config** | Pointer to [**BlueprintCreateSuccessConfigOneOf4Config**](BlueprintCreateSuccessConfigOneOf4Config.md) |  | [optional] 
**Visibility** | Pointer to **string** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | Pointer to **map[string]interface{}** | Resource Permission Block | [optional] 
**Owner** | Pointer to **map[string]interface{}** | Owner | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | Tenant | [optional] 

## Methods

### NewBlueprintCreateSuccessConfigOneOf4

`func NewBlueprintCreateSuccessConfigOneOf4() *BlueprintCreateSuccessConfigOneOf4`

NewBlueprintCreateSuccessConfigOneOf4 instantiates a new BlueprintCreateSuccessConfigOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *BlueprintCreateSuccessConfigOneOf4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintCreateSuccessConfigOneOf4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintCreateSuccessConfigOneOf4) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintCreateSuccessConfigOneOf4) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *BlueprintCreateSuccessConfigOneOf4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintCreateSuccessConfigOneOf4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintCreateSuccessConfigOneOf4) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintCreateSuccessConfigOneOf4) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *BlueprintCreateSuccessConfigOneOf4) GetConfig() BlueprintCreateSuccessConfigOneOf4Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BlueprintCreateSuccessConfigOneOf4) GetConfigOk() (*BlueprintCreateSuccessConfigOneOf4Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BlueprintCreateSuccessConfigOneOf4) SetConfig(v BlueprintCreateSuccessConfigOneOf4Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BlueprintCreateSuccessConfigOneOf4) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *BlueprintCreateSuccessConfigOneOf4) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BlueprintCreateSuccessConfigOneOf4) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BlueprintCreateSuccessConfigOneOf4) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BlueprintCreateSuccessConfigOneOf4) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *BlueprintCreateSuccessConfigOneOf4) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *BlueprintCreateSuccessConfigOneOf4) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *BlueprintCreateSuccessConfigOneOf4) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *BlueprintCreateSuccessConfigOneOf4) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *BlueprintCreateSuccessConfigOneOf4) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BlueprintCreateSuccessConfigOneOf4) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BlueprintCreateSuccessConfigOneOf4) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BlueprintCreateSuccessConfigOneOf4) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *BlueprintCreateSuccessConfigOneOf4) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BlueprintCreateSuccessConfigOneOf4) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BlueprintCreateSuccessConfigOneOf4) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BlueprintCreateSuccessConfigOneOf4) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


