# TagsPolicyTypeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strict** | **bool** | Strict enforcement | 
**Key** | Pointer to **string** | Tag key to enforce | [optional] 
**ValueListId** | Pointer to **string** | ID of value from value list (optional) | [optional] 
**Value** | Pointer to **string** | Tag value (optional, can be left empty for any value) | [optional] 

## Methods

### NewTagsPolicyTypeConfiguration

`func NewTagsPolicyTypeConfiguration(strict bool, ) *TagsPolicyTypeConfiguration`

NewTagsPolicyTypeConfiguration instantiates a new TagsPolicyTypeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsPolicyTypeConfigurationWithDefaults

`func NewTagsPolicyTypeConfigurationWithDefaults() *TagsPolicyTypeConfiguration`

NewTagsPolicyTypeConfigurationWithDefaults instantiates a new TagsPolicyTypeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrict

`func (o *TagsPolicyTypeConfiguration) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *TagsPolicyTypeConfiguration) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *TagsPolicyTypeConfiguration) SetStrict(v bool)`

SetStrict sets Strict field to given value.


### GetKey

`func (o *TagsPolicyTypeConfiguration) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TagsPolicyTypeConfiguration) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TagsPolicyTypeConfiguration) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TagsPolicyTypeConfiguration) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValueListId

`func (o *TagsPolicyTypeConfiguration) GetValueListId() string`

GetValueListId returns the ValueListId field if non-nil, zero value otherwise.

### GetValueListIdOk

`func (o *TagsPolicyTypeConfiguration) GetValueListIdOk() (*string, bool)`

GetValueListIdOk returns a tuple with the ValueListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueListId

`func (o *TagsPolicyTypeConfiguration) SetValueListId(v string)`

SetValueListId sets ValueListId field to given value.

### HasValueListId

`func (o *TagsPolicyTypeConfiguration) HasValueListId() bool`

HasValueListId returns a boolean if a field has been set.

### GetValue

`func (o *TagsPolicyTypeConfiguration) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TagsPolicyTypeConfiguration) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TagsPolicyTypeConfiguration) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *TagsPolicyTypeConfiguration) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


