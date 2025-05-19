# ListEmailTemplates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTemplates** | Pointer to [**[]ListEmailTemplates200ResponseAllOfEmailTemplatesInner**](ListEmailTemplates200ResponseAllOfEmailTemplatesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListEmailTemplates200Response

`func NewListEmailTemplates200Response() *ListEmailTemplates200Response`

NewListEmailTemplates200Response instantiates a new ListEmailTemplates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEmailTemplates200ResponseWithDefaults

`func NewListEmailTemplates200ResponseWithDefaults() *ListEmailTemplates200Response`

NewListEmailTemplates200ResponseWithDefaults instantiates a new ListEmailTemplates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailTemplates

`func (o *ListEmailTemplates200Response) GetEmailTemplates() []ListEmailTemplates200ResponseAllOfEmailTemplatesInner`

GetEmailTemplates returns the EmailTemplates field if non-nil, zero value otherwise.

### GetEmailTemplatesOk

`func (o *ListEmailTemplates200Response) GetEmailTemplatesOk() (*[]ListEmailTemplates200ResponseAllOfEmailTemplatesInner, bool)`

GetEmailTemplatesOk returns a tuple with the EmailTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplates

`func (o *ListEmailTemplates200Response) SetEmailTemplates(v []ListEmailTemplates200ResponseAllOfEmailTemplatesInner)`

SetEmailTemplates sets EmailTemplates field to given value.

### HasEmailTemplates

`func (o *ListEmailTemplates200Response) HasEmailTemplates() bool`

HasEmailTemplates returns a boolean if a field has been set.

### GetMeta

`func (o *ListEmailTemplates200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListEmailTemplates200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListEmailTemplates200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListEmailTemplates200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


