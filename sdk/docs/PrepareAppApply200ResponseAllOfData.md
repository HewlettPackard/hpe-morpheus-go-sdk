# PrepareAppApply200ResponseAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AutoValidate** | Pointer to **bool** |  | [optional] 
**Terraform** | Pointer to [**PrepareAppApply200ResponseAllOfDataTerraform**](PrepareAppApply200ResponseAllOfDataTerraform.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**BlueprintName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TemplateId** | Pointer to **int64** |  | [optional] 
**BlueprintId** | Pointer to **int64** |  | [optional] 
**Group** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewPrepareAppApply200ResponseAllOfData

`func NewPrepareAppApply200ResponseAllOfData() *PrepareAppApply200ResponseAllOfData`

NewPrepareAppApply200ResponseAllOfData instantiates a new PrepareAppApply200ResponseAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrepareAppApply200ResponseAllOfDataWithDefaults

`func NewPrepareAppApply200ResponseAllOfDataWithDefaults() *PrepareAppApply200ResponseAllOfData`

NewPrepareAppApply200ResponseAllOfDataWithDefaults instantiates a new PrepareAppApply200ResponseAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *PrepareAppApply200ResponseAllOfData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PrepareAppApply200ResponseAllOfData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PrepareAppApply200ResponseAllOfData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PrepareAppApply200ResponseAllOfData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *PrepareAppApply200ResponseAllOfData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrepareAppApply200ResponseAllOfData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrepareAppApply200ResponseAllOfData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrepareAppApply200ResponseAllOfData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAutoValidate

`func (o *PrepareAppApply200ResponseAllOfData) GetAutoValidate() bool`

GetAutoValidate returns the AutoValidate field if non-nil, zero value otherwise.

### GetAutoValidateOk

`func (o *PrepareAppApply200ResponseAllOfData) GetAutoValidateOk() (*bool, bool)`

GetAutoValidateOk returns a tuple with the AutoValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoValidate

`func (o *PrepareAppApply200ResponseAllOfData) SetAutoValidate(v bool)`

SetAutoValidate sets AutoValidate field to given value.

### HasAutoValidate

`func (o *PrepareAppApply200ResponseAllOfData) HasAutoValidate() bool`

HasAutoValidate returns a boolean if a field has been set.

### GetTerraform

`func (o *PrepareAppApply200ResponseAllOfData) GetTerraform() PrepareAppApply200ResponseAllOfDataTerraform`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *PrepareAppApply200ResponseAllOfData) GetTerraformOk() (*PrepareAppApply200ResponseAllOfDataTerraform, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *PrepareAppApply200ResponseAllOfData) SetTerraform(v PrepareAppApply200ResponseAllOfDataTerraform)`

SetTerraform sets Terraform field to given value.

### HasTerraform

`func (o *PrepareAppApply200ResponseAllOfData) HasTerraform() bool`

HasTerraform returns a boolean if a field has been set.

### GetType

`func (o *PrepareAppApply200ResponseAllOfData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrepareAppApply200ResponseAllOfData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrepareAppApply200ResponseAllOfData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PrepareAppApply200ResponseAllOfData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *PrepareAppApply200ResponseAllOfData) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PrepareAppApply200ResponseAllOfData) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PrepareAppApply200ResponseAllOfData) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PrepareAppApply200ResponseAllOfData) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetBlueprintName

`func (o *PrepareAppApply200ResponseAllOfData) GetBlueprintName() string`

GetBlueprintName returns the BlueprintName field if non-nil, zero value otherwise.

### GetBlueprintNameOk

`func (o *PrepareAppApply200ResponseAllOfData) GetBlueprintNameOk() (*string, bool)`

GetBlueprintNameOk returns a tuple with the BlueprintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintName

`func (o *PrepareAppApply200ResponseAllOfData) SetBlueprintName(v string)`

SetBlueprintName sets BlueprintName field to given value.

### HasBlueprintName

`func (o *PrepareAppApply200ResponseAllOfData) HasBlueprintName() bool`

HasBlueprintName returns a boolean if a field has been set.

### GetDescription

`func (o *PrepareAppApply200ResponseAllOfData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrepareAppApply200ResponseAllOfData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrepareAppApply200ResponseAllOfData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrepareAppApply200ResponseAllOfData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplateId

`func (o *PrepareAppApply200ResponseAllOfData) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *PrepareAppApply200ResponseAllOfData) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *PrepareAppApply200ResponseAllOfData) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *PrepareAppApply200ResponseAllOfData) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetBlueprintId

`func (o *PrepareAppApply200ResponseAllOfData) GetBlueprintId() int64`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *PrepareAppApply200ResponseAllOfData) GetBlueprintIdOk() (*int64, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *PrepareAppApply200ResponseAllOfData) SetBlueprintId(v int64)`

SetBlueprintId sets BlueprintId field to given value.

### HasBlueprintId

`func (o *PrepareAppApply200ResponseAllOfData) HasBlueprintId() bool`

HasBlueprintId returns a boolean if a field has been set.

### GetGroup

`func (o *PrepareAppApply200ResponseAllOfData) GetGroup() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PrepareAppApply200ResponseAllOfData) GetGroupOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PrepareAppApply200ResponseAllOfData) SetGroup(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PrepareAppApply200ResponseAllOfData) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


