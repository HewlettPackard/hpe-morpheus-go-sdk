# ListChecks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to [**[]ListChecks200ResponseAllOfChecksInner**](ListChecks200ResponseAllOfChecksInner.md) |  | [optional] 
**Meta** | Pointer to [**ListAlerts200ResponseAllOfMeta**](ListAlerts200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListChecks200Response

`func NewListChecks200Response() *ListChecks200Response`

NewListChecks200Response instantiates a new ListChecks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListChecks200ResponseWithDefaults

`func NewListChecks200ResponseWithDefaults() *ListChecks200Response`

NewListChecks200ResponseWithDefaults instantiates a new ListChecks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *ListChecks200Response) GetChecks() []ListChecks200ResponseAllOfChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ListChecks200Response) GetChecksOk() (*[]ListChecks200ResponseAllOfChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ListChecks200Response) SetChecks(v []ListChecks200ResponseAllOfChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ListChecks200Response) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetMeta

`func (o *ListChecks200Response) GetMeta() ListAlerts200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListChecks200Response) GetMetaOk() (*ListAlerts200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListChecks200Response) SetMeta(v ListAlerts200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListChecks200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


