# BlueprintARMCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**Arm** | Pointer to [**AddBlueprintRequestOneOfArm**](AddBlueprintRequestOneOfArm.md) |  | [optional] 
**Visibility** | Pointer to **string** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | Pointer to **map[string]interface{}** | Resource Permission Block | [optional] 
**Owner** | Pointer to **map[string]interface{}** | Owner | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | Tenant | [optional] 

## Methods

### NewBlueprintARMCreateSuccess

`func NewBlueprintARMCreateSuccess() *BlueprintARMCreateSuccess`

NewBlueprintARMCreateSuccess instantiates a new BlueprintARMCreateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintARMCreateSuccessWithDefaults

`func NewBlueprintARMCreateSuccessWithDefaults() *BlueprintARMCreateSuccess`

NewBlueprintARMCreateSuccessWithDefaults instantiates a new BlueprintARMCreateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintARMCreateSuccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintARMCreateSuccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintARMCreateSuccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintARMCreateSuccess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *BlueprintARMCreateSuccess) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintARMCreateSuccess) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintARMCreateSuccess) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintARMCreateSuccess) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *BlueprintARMCreateSuccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintARMCreateSuccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintARMCreateSuccess) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintARMCreateSuccess) HasType() bool`

HasType returns a boolean if a field has been set.

### GetArm

`func (o *BlueprintARMCreateSuccess) GetArm() AddBlueprintRequestOneOfArm`

GetArm returns the Arm field if non-nil, zero value otherwise.

### GetArmOk

`func (o *BlueprintARMCreateSuccess) GetArmOk() (*AddBlueprintRequestOneOfArm, bool)`

GetArmOk returns a tuple with the Arm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArm

`func (o *BlueprintARMCreateSuccess) SetArm(v AddBlueprintRequestOneOfArm)`

SetArm sets Arm field to given value.

### HasArm

`func (o *BlueprintARMCreateSuccess) HasArm() bool`

HasArm returns a boolean if a field has been set.

### GetVisibility

`func (o *BlueprintARMCreateSuccess) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BlueprintARMCreateSuccess) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BlueprintARMCreateSuccess) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BlueprintARMCreateSuccess) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *BlueprintARMCreateSuccess) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *BlueprintARMCreateSuccess) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *BlueprintARMCreateSuccess) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *BlueprintARMCreateSuccess) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *BlueprintARMCreateSuccess) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BlueprintARMCreateSuccess) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BlueprintARMCreateSuccess) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BlueprintARMCreateSuccess) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *BlueprintARMCreateSuccess) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BlueprintARMCreateSuccess) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BlueprintARMCreateSuccess) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BlueprintARMCreateSuccess) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


