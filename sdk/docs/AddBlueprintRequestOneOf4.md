# AddBlueprintRequestOneOf4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Tiers** | **map[string]interface{}** | Tier definitions - Create in UI to view a baseline for object | 

## Methods

### NewAddBlueprintRequestOneOf4

`func NewAddBlueprintRequestOneOf4(name string, type_ string, tiers map[string]interface{}, ) *AddBlueprintRequestOneOf4`

NewAddBlueprintRequestOneOf4 instantiates a new AddBlueprintRequestOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprintRequestOneOf4WithDefaults

`func NewAddBlueprintRequestOneOf4WithDefaults() *AddBlueprintRequestOneOf4`

NewAddBlueprintRequestOneOf4WithDefaults instantiates a new AddBlueprintRequestOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddBlueprintRequestOneOf4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBlueprintRequestOneOf4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBlueprintRequestOneOf4) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddBlueprintRequestOneOf4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddBlueprintRequestOneOf4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddBlueprintRequestOneOf4) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *AddBlueprintRequestOneOf4) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddBlueprintRequestOneOf4) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddBlueprintRequestOneOf4) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddBlueprintRequestOneOf4) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTiers

`func (o *AddBlueprintRequestOneOf4) GetTiers() map[string]interface{}`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *AddBlueprintRequestOneOf4) GetTiersOk() (*map[string]interface{}, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *AddBlueprintRequestOneOf4) SetTiers(v map[string]interface{})`

SetTiers sets Tiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


