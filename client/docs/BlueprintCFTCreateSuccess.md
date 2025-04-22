# BlueprintCFTCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**CloudFormation** | Pointer to [**AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation**](AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation.md) |  | [optional] 
**Visibility** | Pointer to **string** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | Pointer to **map[string]interface{}** | Resource Permission Block | [optional] 
**Owner** | Pointer to **map[string]interface{}** | Owner | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | Tenant | [optional] 

## Methods

### NewBlueprintCFTCreateSuccess

`func NewBlueprintCFTCreateSuccess() *BlueprintCFTCreateSuccess`

NewBlueprintCFTCreateSuccess instantiates a new BlueprintCFTCreateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCFTCreateSuccessWithDefaults

`func NewBlueprintCFTCreateSuccessWithDefaults() *BlueprintCFTCreateSuccess`

NewBlueprintCFTCreateSuccessWithDefaults instantiates a new BlueprintCFTCreateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintCFTCreateSuccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintCFTCreateSuccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintCFTCreateSuccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintCFTCreateSuccess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *BlueprintCFTCreateSuccess) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintCFTCreateSuccess) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintCFTCreateSuccess) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintCFTCreateSuccess) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *BlueprintCFTCreateSuccess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintCFTCreateSuccess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintCFTCreateSuccess) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintCFTCreateSuccess) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCloudFormation

`func (o *BlueprintCFTCreateSuccess) GetCloudFormation() AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation`

GetCloudFormation returns the CloudFormation field if non-nil, zero value otherwise.

### GetCloudFormationOk

`func (o *BlueprintCFTCreateSuccess) GetCloudFormationOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation, bool)`

GetCloudFormationOk returns a tuple with the CloudFormation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormation

`func (o *BlueprintCFTCreateSuccess) SetCloudFormation(v AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation)`

SetCloudFormation sets CloudFormation field to given value.

### HasCloudFormation

`func (o *BlueprintCFTCreateSuccess) HasCloudFormation() bool`

HasCloudFormation returns a boolean if a field has been set.

### GetVisibility

`func (o *BlueprintCFTCreateSuccess) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BlueprintCFTCreateSuccess) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BlueprintCFTCreateSuccess) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BlueprintCFTCreateSuccess) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *BlueprintCFTCreateSuccess) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *BlueprintCFTCreateSuccess) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *BlueprintCFTCreateSuccess) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *BlueprintCFTCreateSuccess) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *BlueprintCFTCreateSuccess) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BlueprintCFTCreateSuccess) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BlueprintCFTCreateSuccess) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BlueprintCFTCreateSuccess) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *BlueprintCFTCreateSuccess) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BlueprintCFTCreateSuccess) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BlueprintCFTCreateSuccess) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BlueprintCFTCreateSuccess) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


