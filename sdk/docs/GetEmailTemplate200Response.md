# GetEmailTemplate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTemplate** | Pointer to [**ListEmailTemplates200ResponseAllOfEmailTemplatesInner**](ListEmailTemplates200ResponseAllOfEmailTemplatesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetEmailTemplate200Response

`func NewGetEmailTemplate200Response() *GetEmailTemplate200Response`

NewGetEmailTemplate200Response instantiates a new GetEmailTemplate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmailTemplate200ResponseWithDefaults

`func NewGetEmailTemplate200ResponseWithDefaults() *GetEmailTemplate200Response`

NewGetEmailTemplate200ResponseWithDefaults instantiates a new GetEmailTemplate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailTemplate

`func (o *GetEmailTemplate200Response) GetEmailTemplate() ListEmailTemplates200ResponseAllOfEmailTemplatesInner`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *GetEmailTemplate200Response) GetEmailTemplateOk() (*ListEmailTemplates200ResponseAllOfEmailTemplatesInner, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *GetEmailTemplate200Response) SetEmailTemplate(v ListEmailTemplates200ResponseAllOfEmailTemplatesInner)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *GetEmailTemplate200Response) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### GetSuccess

`func (o *GetEmailTemplate200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetEmailTemplate200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetEmailTemplate200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetEmailTemplate200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


