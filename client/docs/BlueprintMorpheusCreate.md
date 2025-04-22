# BlueprintMorpheusCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Tiers** | **map[string]interface{}** | Tier definitions - Create in UI to view a baseline for object | 

## Methods

### NewBlueprintMorpheusCreate

`func NewBlueprintMorpheusCreate(name string, type_ string, tiers map[string]interface{}, ) *BlueprintMorpheusCreate`

NewBlueprintMorpheusCreate instantiates a new BlueprintMorpheusCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintMorpheusCreateWithDefaults

`func NewBlueprintMorpheusCreateWithDefaults() *BlueprintMorpheusCreate`

NewBlueprintMorpheusCreateWithDefaults instantiates a new BlueprintMorpheusCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintMorpheusCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintMorpheusCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintMorpheusCreate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *BlueprintMorpheusCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintMorpheusCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintMorpheusCreate) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *BlueprintMorpheusCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *BlueprintMorpheusCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *BlueprintMorpheusCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *BlueprintMorpheusCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTiers

`func (o *BlueprintMorpheusCreate) GetTiers() map[string]interface{}`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *BlueprintMorpheusCreate) GetTiersOk() (*map[string]interface{}, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *BlueprintMorpheusCreate) SetTiers(v map[string]interface{})`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


