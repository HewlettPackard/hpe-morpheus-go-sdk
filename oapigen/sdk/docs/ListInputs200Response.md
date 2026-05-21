# ListInputs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionTypes** | Pointer to [**[]ListInputs200ResponseAllOfOptionTypesInner**](ListInputs200ResponseAllOfOptionTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListInputs200Response

`func NewListInputs200Response() *ListInputs200Response`

NewListInputs200Response instantiates a new ListInputs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInputs200ResponseWithDefaults

`func NewListInputs200ResponseWithDefaults() *ListInputs200Response`

NewListInputs200ResponseWithDefaults instantiates a new ListInputs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionTypes

`func (o *ListInputs200Response) GetOptionTypes() []ListInputs200ResponseAllOfOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListInputs200Response) GetOptionTypesOk() (*[]ListInputs200ResponseAllOfOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListInputs200Response) SetOptionTypes(v []ListInputs200ResponseAllOfOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListInputs200Response) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListInputs200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInputs200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInputs200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInputs200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


