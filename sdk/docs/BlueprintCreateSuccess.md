# BlueprintCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Blueprint ID | [optional] 
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Description** | Pointer to **string** | A description for the blueprint | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Category** | Pointer to **string** | Category | [optional] 
**Config** | Pointer to [**AddBlueprint200ResponseAllOfBlueprintConfig**](AddBlueprint200ResponseAllOfBlueprintConfig.md) |  | [optional] 

## Methods

### NewBlueprintCreateSuccess

`func NewBlueprintCreateSuccess() *BlueprintCreateSuccess`

NewBlueprintCreateSuccess instantiates a new BlueprintCreateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCreateSuccessWithDefaults

`func NewBlueprintCreateSuccessWithDefaults() *BlueprintCreateSuccess`

NewBlueprintCreateSuccessWithDefaults instantiates a new BlueprintCreateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlueprintCreateSuccess) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlueprintCreateSuccess) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlueprintCreateSuccess) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BlueprintCreateSuccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BlueprintCreateSuccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintCreateSuccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintCreateSuccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintCreateSuccess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BlueprintCreateSuccess) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintCreateSuccess) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintCreateSuccess) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintCreateSuccess) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *BlueprintCreateSuccess) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BlueprintCreateSuccess) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BlueprintCreateSuccess) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BlueprintCreateSuccess) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCategory

`func (o *BlueprintCreateSuccess) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BlueprintCreateSuccess) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BlueprintCreateSuccess) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BlueprintCreateSuccess) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *BlueprintCreateSuccess) GetConfig() AddBlueprint200ResponseAllOfBlueprintConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BlueprintCreateSuccess) GetConfigOk() (*AddBlueprint200ResponseAllOfBlueprintConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BlueprintCreateSuccess) SetConfig(v AddBlueprint200ResponseAllOfBlueprintConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BlueprintCreateSuccess) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


