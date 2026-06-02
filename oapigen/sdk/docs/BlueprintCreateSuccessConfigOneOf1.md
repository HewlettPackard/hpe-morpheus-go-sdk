# BlueprintCreateSuccessConfigOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**CloudFormation** | Pointer to [**BlueprintCreateSuccessConfigOneOf1CloudFormation**](BlueprintCreateSuccessConfigOneOf1CloudFormation.md) |  | [optional] 
**Visibility** | Pointer to **string** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | Pointer to **map[string]interface{}** | Resource Permission Block | [optional] 
**Owner** | Pointer to **map[string]interface{}** | Owner | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | Tenant | [optional] 

## Methods

### NewBlueprintCreateSuccessConfigOneOf1

`func NewBlueprintCreateSuccessConfigOneOf1() *BlueprintCreateSuccessConfigOneOf1`

NewBlueprintCreateSuccessConfigOneOf1 instantiates a new BlueprintCreateSuccessConfigOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *BlueprintCreateSuccessConfigOneOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintCreateSuccessConfigOneOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintCreateSuccessConfigOneOf1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintCreateSuccessConfigOneOf1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *BlueprintCreateSuccessConfigOneOf1) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintCreateSuccessConfigOneOf1) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintCreateSuccessConfigOneOf1) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintCreateSuccessConfigOneOf1) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *BlueprintCreateSuccessConfigOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintCreateSuccessConfigOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintCreateSuccessConfigOneOf1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintCreateSuccessConfigOneOf1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCloudFormation

`func (o *BlueprintCreateSuccessConfigOneOf1) GetCloudFormation() BlueprintCreateSuccessConfigOneOf1CloudFormation`

GetCloudFormation returns the CloudFormation field if non-nil, zero value otherwise.

### GetCloudFormationOk

`func (o *BlueprintCreateSuccessConfigOneOf1) GetCloudFormationOk() (*BlueprintCreateSuccessConfigOneOf1CloudFormation, bool)`

GetCloudFormationOk returns a tuple with the CloudFormation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormation

`func (o *BlueprintCreateSuccessConfigOneOf1) SetCloudFormation(v BlueprintCreateSuccessConfigOneOf1CloudFormation)`

SetCloudFormation sets CloudFormation field to given value.

### HasCloudFormation

`func (o *BlueprintCreateSuccessConfigOneOf1) HasCloudFormation() bool`

HasCloudFormation returns a boolean if a field has been set.

### GetVisibility

`func (o *BlueprintCreateSuccessConfigOneOf1) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BlueprintCreateSuccessConfigOneOf1) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BlueprintCreateSuccessConfigOneOf1) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BlueprintCreateSuccessConfigOneOf1) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *BlueprintCreateSuccessConfigOneOf1) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *BlueprintCreateSuccessConfigOneOf1) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *BlueprintCreateSuccessConfigOneOf1) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *BlueprintCreateSuccessConfigOneOf1) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *BlueprintCreateSuccessConfigOneOf1) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BlueprintCreateSuccessConfigOneOf1) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BlueprintCreateSuccessConfigOneOf1) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BlueprintCreateSuccessConfigOneOf1) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *BlueprintCreateSuccessConfigOneOf1) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BlueprintCreateSuccessConfigOneOf1) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BlueprintCreateSuccessConfigOneOf1) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BlueprintCreateSuccessConfigOneOf1) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


