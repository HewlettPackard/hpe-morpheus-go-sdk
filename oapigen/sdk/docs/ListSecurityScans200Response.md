# ListSecurityScans200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityScans** | Pointer to [**[]ListSecurityScans200ResponseAllOfSecurityScansInner**](ListSecurityScans200ResponseAllOfSecurityScansInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListSecurityScans200Response

`func NewListSecurityScans200Response() *ListSecurityScans200Response`

NewListSecurityScans200Response instantiates a new ListSecurityScans200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityScans200ResponseWithDefaults

`func NewListSecurityScans200ResponseWithDefaults() *ListSecurityScans200Response`

NewListSecurityScans200ResponseWithDefaults instantiates a new ListSecurityScans200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityScans

`func (o *ListSecurityScans200Response) GetSecurityScans() []ListSecurityScans200ResponseAllOfSecurityScansInner`

GetSecurityScans returns the SecurityScans field if non-nil, zero value otherwise.

### GetSecurityScansOk

`func (o *ListSecurityScans200Response) GetSecurityScansOk() (*[]ListSecurityScans200ResponseAllOfSecurityScansInner, bool)`

GetSecurityScansOk returns a tuple with the SecurityScans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityScans

`func (o *ListSecurityScans200Response) SetSecurityScans(v []ListSecurityScans200ResponseAllOfSecurityScansInner)`

SetSecurityScans sets SecurityScans field to given value.

### HasSecurityScans

`func (o *ListSecurityScans200Response) HasSecurityScans() bool`

HasSecurityScans returns a boolean if a field has been set.

### GetMeta

`func (o *ListSecurityScans200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSecurityScans200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSecurityScans200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSecurityScans200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


