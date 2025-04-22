# BlueprintHelmCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**Helm** | Pointer to [**AddBlueprintRequestOneOf2Helm**](AddBlueprintRequestOneOf2Helm.md) |  | [optional] 
**Visibility** | Pointer to **string** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | Pointer to **map[string]interface{}** | Resource Permission Block | [optional] 
**Owner** | Pointer to **map[string]interface{}** | Owner | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | Tenant | [optional] 

## Methods

### NewBlueprintHelmCreateSuccess

`func NewBlueprintHelmCreateSuccess() *BlueprintHelmCreateSuccess`

NewBlueprintHelmCreateSuccess instantiates a new BlueprintHelmCreateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintHelmCreateSuccessWithDefaults

`func NewBlueprintHelmCreateSuccessWithDefaults() *BlueprintHelmCreateSuccess`

NewBlueprintHelmCreateSuccessWithDefaults instantiates a new BlueprintHelmCreateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintHelmCreateSuccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintHelmCreateSuccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintHelmCreateSuccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintHelmCreateSuccess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *BlueprintHelmCreateSuccess) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintHelmCreateSuccess) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintHelmCreateSuccess) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintHelmCreateSuccess) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *BlueprintHelmCreateSuccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintHelmCreateSuccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintHelmCreateSuccess) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintHelmCreateSuccess) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHelm

`func (o *BlueprintHelmCreateSuccess) GetHelm() AddBlueprintRequestOneOf2Helm`

GetHelm returns the Helm field if non-nil, zero value otherwise.

### GetHelmOk

`func (o *BlueprintHelmCreateSuccess) GetHelmOk() (*AddBlueprintRequestOneOf2Helm, bool)`

GetHelmOk returns a tuple with the Helm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelm

`func (o *BlueprintHelmCreateSuccess) SetHelm(v AddBlueprintRequestOneOf2Helm)`

SetHelm sets Helm field to given value.

### HasHelm

`func (o *BlueprintHelmCreateSuccess) HasHelm() bool`

HasHelm returns a boolean if a field has been set.

### GetVisibility

`func (o *BlueprintHelmCreateSuccess) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BlueprintHelmCreateSuccess) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BlueprintHelmCreateSuccess) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BlueprintHelmCreateSuccess) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *BlueprintHelmCreateSuccess) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *BlueprintHelmCreateSuccess) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *BlueprintHelmCreateSuccess) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *BlueprintHelmCreateSuccess) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *BlueprintHelmCreateSuccess) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BlueprintHelmCreateSuccess) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BlueprintHelmCreateSuccess) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BlueprintHelmCreateSuccess) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *BlueprintHelmCreateSuccess) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BlueprintHelmCreateSuccess) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BlueprintHelmCreateSuccess) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BlueprintHelmCreateSuccess) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


