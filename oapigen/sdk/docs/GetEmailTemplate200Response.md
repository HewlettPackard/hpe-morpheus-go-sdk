# GetEmailTemplate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailTemplate** | Pointer to [**AddEmailTemplate200ResponseEmailTemplate**](AddEmailTemplate200ResponseEmailTemplate.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetEmailTemplate200Response

`func NewGetEmailTemplate200Response() *GetEmailTemplate200Response`

NewGetEmailTemplate200Response instantiates a new GetEmailTemplate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetEmailTemplate

`func (o *GetEmailTemplate200Response) GetEmailTemplate() AddEmailTemplate200ResponseEmailTemplate`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *GetEmailTemplate200Response) GetEmailTemplateOk() (*AddEmailTemplate200ResponseEmailTemplate, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *GetEmailTemplate200Response) SetEmailTemplate(v AddEmailTemplate200ResponseEmailTemplate)`

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


