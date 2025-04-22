# AddAppsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | Pointer to **int64** |  | [optional] 
**BlueprintId** | [**AddAppsRequestBlueprintId**](AddAppsRequestBlueprintId.md) |  | 
**Name** | **string** | A unique name for the app | 
**Description** | Pointer to **string** | Description | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Group** | Pointer to [**AddAppsRequestGroup**](AddAppsRequestGroup.md) |  | [optional] 
**DefaultCloud** | Pointer to [**AddAppsRequestDefaultCloud**](AddAppsRequestDefaultCloud.md) |  | [optional] 
**Environment** | Pointer to **string** | Environment code (appContext) | [optional] 
**Tiers** | Pointer to **map[string]interface{}** | Configuration of app elements | [optional] 

## Methods

### NewAddAppsRequest

`func NewAddAppsRequest(blueprintId AddAppsRequestBlueprintId, name string, ) *AddAppsRequest`

NewAddAppsRequest instantiates a new AddAppsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAppsRequestWithDefaults

`func NewAddAppsRequestWithDefaults() *AddAppsRequest`

NewAddAppsRequestWithDefaults instantiates a new AddAppsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *AddAppsRequest) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *AddAppsRequest) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *AddAppsRequest) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *AddAppsRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetBlueprintId

`func (o *AddAppsRequest) GetBlueprintId() AddAppsRequestBlueprintId`

GetBlueprintId returns the BlueprintId field if non-nil, zero value otherwise.

### GetBlueprintIdOk

`func (o *AddAppsRequest) GetBlueprintIdOk() (*AddAppsRequestBlueprintId, bool)`

GetBlueprintIdOk returns a tuple with the BlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintId

`func (o *AddAppsRequest) SetBlueprintId(v AddAppsRequestBlueprintId)`

SetBlueprintId sets BlueprintId field to given value.


### GetName

`func (o *AddAppsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddAppsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddAppsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddAppsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAppsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAppsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAppsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AddAppsRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddAppsRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddAppsRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddAppsRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetGroup

`func (o *AddAppsRequest) GetGroup() AddAppsRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AddAppsRequest) GetGroupOk() (*AddAppsRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AddAppsRequest) SetGroup(v AddAppsRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AddAppsRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetDefaultCloud

`func (o *AddAppsRequest) GetDefaultCloud() AddAppsRequestDefaultCloud`

GetDefaultCloud returns the DefaultCloud field if non-nil, zero value otherwise.

### GetDefaultCloudOk

`func (o *AddAppsRequest) GetDefaultCloudOk() (*AddAppsRequestDefaultCloud, bool)`

GetDefaultCloudOk returns a tuple with the DefaultCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCloud

`func (o *AddAppsRequest) SetDefaultCloud(v AddAppsRequestDefaultCloud)`

SetDefaultCloud sets DefaultCloud field to given value.

### HasDefaultCloud

`func (o *AddAppsRequest) HasDefaultCloud() bool`

HasDefaultCloud returns a boolean if a field has been set.

### GetEnvironment

`func (o *AddAppsRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AddAppsRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AddAppsRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AddAppsRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetTiers

`func (o *AddAppsRequest) GetTiers() map[string]interface{}`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *AddAppsRequest) GetTiersOk() (*map[string]interface{}, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *AddAppsRequest) SetTiers(v map[string]interface{})`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *AddAppsRequest) HasTiers() bool`

HasTiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


