# BlueprintMorpheusCreateSuccessConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**Tiers** | Pointer to **map[string]interface{}** | Tier definitions - Create in UI to view a baseline for object | [optional] 

## Methods

### NewBlueprintMorpheusCreateSuccessConfig

`func NewBlueprintMorpheusCreateSuccessConfig() *BlueprintMorpheusCreateSuccessConfig`

NewBlueprintMorpheusCreateSuccessConfig instantiates a new BlueprintMorpheusCreateSuccessConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintMorpheusCreateSuccessConfigWithDefaults

`func NewBlueprintMorpheusCreateSuccessConfigWithDefaults() *BlueprintMorpheusCreateSuccessConfig`

NewBlueprintMorpheusCreateSuccessConfigWithDefaults instantiates a new BlueprintMorpheusCreateSuccessConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *BlueprintMorpheusCreateSuccessConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintMorpheusCreateSuccessConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintMorpheusCreateSuccessConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintMorpheusCreateSuccessConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *BlueprintMorpheusCreateSuccessConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintMorpheusCreateSuccessConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintMorpheusCreateSuccessConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintMorpheusCreateSuccessConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *BlueprintMorpheusCreateSuccessConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintMorpheusCreateSuccessConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintMorpheusCreateSuccessConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintMorpheusCreateSuccessConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTiers

`func (o *BlueprintMorpheusCreateSuccessConfig) GetTiers() map[string]interface{}`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *BlueprintMorpheusCreateSuccessConfig) GetTiersOk() (*map[string]interface{}, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *BlueprintMorpheusCreateSuccessConfig) SetTiers(v map[string]interface{})`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *BlueprintMorpheusCreateSuccessConfig) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


