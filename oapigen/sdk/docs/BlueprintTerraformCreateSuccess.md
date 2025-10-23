# BlueprintTerraformCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**Terraform** | Pointer to [**AddBlueprintRequestOneOf5Terraform**](AddBlueprintRequestOneOf5Terraform.md) |  | [optional] 
**Config** | Pointer to [**AddBlueprintRequestOneOf5Config**](AddBlueprintRequestOneOf5Config.md) |  | [optional] 
**Visibility** | Pointer to **string** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | Pointer to **map[string]interface{}** | Resource Permission Block | [optional] 
**Owner** | Pointer to **map[string]interface{}** | Owner | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | Tenant | [optional] 

## Methods

### NewBlueprintTerraformCreateSuccess

`func NewBlueprintTerraformCreateSuccess() *BlueprintTerraformCreateSuccess`

NewBlueprintTerraformCreateSuccess instantiates a new BlueprintTerraformCreateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintTerraformCreateSuccessWithDefaults

`func NewBlueprintTerraformCreateSuccessWithDefaults() *BlueprintTerraformCreateSuccess`

NewBlueprintTerraformCreateSuccessWithDefaults instantiates a new BlueprintTerraformCreateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintTerraformCreateSuccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintTerraformCreateSuccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintTerraformCreateSuccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintTerraformCreateSuccess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *BlueprintTerraformCreateSuccess) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintTerraformCreateSuccess) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintTerraformCreateSuccess) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintTerraformCreateSuccess) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *BlueprintTerraformCreateSuccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintTerraformCreateSuccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintTerraformCreateSuccess) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintTerraformCreateSuccess) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTerraform

`func (o *BlueprintTerraformCreateSuccess) GetTerraform() AddBlueprintRequestOneOf5Terraform`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *BlueprintTerraformCreateSuccess) GetTerraformOk() (*AddBlueprintRequestOneOf5Terraform, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *BlueprintTerraformCreateSuccess) SetTerraform(v AddBlueprintRequestOneOf5Terraform)`

SetTerraform sets Terraform field to given value.

### HasTerraform

`func (o *BlueprintTerraformCreateSuccess) HasTerraform() bool`

HasTerraform returns a boolean if a field has been set.

### GetConfig

`func (o *BlueprintTerraformCreateSuccess) GetConfig() AddBlueprintRequestOneOf5Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BlueprintTerraformCreateSuccess) GetConfigOk() (*AddBlueprintRequestOneOf5Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BlueprintTerraformCreateSuccess) SetConfig(v AddBlueprintRequestOneOf5Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BlueprintTerraformCreateSuccess) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *BlueprintTerraformCreateSuccess) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BlueprintTerraformCreateSuccess) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BlueprintTerraformCreateSuccess) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BlueprintTerraformCreateSuccess) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *BlueprintTerraformCreateSuccess) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *BlueprintTerraformCreateSuccess) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *BlueprintTerraformCreateSuccess) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *BlueprintTerraformCreateSuccess) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *BlueprintTerraformCreateSuccess) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BlueprintTerraformCreateSuccess) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BlueprintTerraformCreateSuccess) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BlueprintTerraformCreateSuccess) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *BlueprintTerraformCreateSuccess) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BlueprintTerraformCreateSuccess) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BlueprintTerraformCreateSuccess) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BlueprintTerraformCreateSuccess) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


