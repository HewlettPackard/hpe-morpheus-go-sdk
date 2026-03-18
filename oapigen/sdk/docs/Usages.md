# Usages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to [**[]UsagesActivityInner**](UsagesActivityInner.md) |  | [optional] 
**Meta** | Pointer to [**UsagesMeta**](UsagesMeta.md) |  | [optional] 

## Methods

### NewUsages

`func NewUsages() *Usages`

NewUsages instantiates a new Usages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagesWithDefaults

`func NewUsagesWithDefaults() *Usages`

NewUsagesWithDefaults instantiates a new Usages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *Usages) GetActivity() []UsagesActivityInner`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *Usages) GetActivityOk() (*[]UsagesActivityInner, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *Usages) SetActivity(v []UsagesActivityInner)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *Usages) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetMeta

`func (o *Usages) GetMeta() UsagesMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Usages) GetMetaOk() (*UsagesMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Usages) SetMeta(v UsagesMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Usages) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


