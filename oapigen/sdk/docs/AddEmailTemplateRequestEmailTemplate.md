# AddEmailTemplateRequestEmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | The name of the email template. This is set by morpheus.  | [optional] 
**Code** | Pointer to **string** | A unique code for the email template. This code is used to reference the email template and as a reference of the templates type.  | [optional] 
**Owner** | Pointer to [**AddEmailTemplateRequestEmailTemplateOwner**](AddEmailTemplateRequestEmailTemplateOwner.md) |  | [optional] 
**Accounts** | Pointer to [**[]AddEmailTemplateRequestEmailTemplateAccountsInner**](AddEmailTemplateRequestEmailTemplateAccountsInner.md) |  | [optional] 
**Template** | Pointer to **string** | The email template. This is the actual email template that is sent to the user. This uses handlebars notation (not javascript)  | [optional] 

## Methods

### NewAddEmailTemplateRequestEmailTemplate

`func NewAddEmailTemplateRequestEmailTemplate() *AddEmailTemplateRequestEmailTemplate`

NewAddEmailTemplateRequestEmailTemplate instantiates a new AddEmailTemplateRequestEmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddEmailTemplateRequestEmailTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddEmailTemplateRequestEmailTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddEmailTemplateRequestEmailTemplate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddEmailTemplateRequestEmailTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddEmailTemplateRequestEmailTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddEmailTemplateRequestEmailTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddEmailTemplateRequestEmailTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddEmailTemplateRequestEmailTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddEmailTemplateRequestEmailTemplate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddEmailTemplateRequestEmailTemplate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddEmailTemplateRequestEmailTemplate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddEmailTemplateRequestEmailTemplate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetOwner

`func (o *AddEmailTemplateRequestEmailTemplate) GetOwner() AddEmailTemplateRequestEmailTemplateOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddEmailTemplateRequestEmailTemplate) GetOwnerOk() (*AddEmailTemplateRequestEmailTemplateOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddEmailTemplateRequestEmailTemplate) SetOwner(v AddEmailTemplateRequestEmailTemplateOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddEmailTemplateRequestEmailTemplate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccounts

`func (o *AddEmailTemplateRequestEmailTemplate) GetAccounts() []AddEmailTemplateRequestEmailTemplateAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AddEmailTemplateRequestEmailTemplate) GetAccountsOk() (*[]AddEmailTemplateRequestEmailTemplateAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AddEmailTemplateRequestEmailTemplate) SetAccounts(v []AddEmailTemplateRequestEmailTemplateAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AddEmailTemplateRequestEmailTemplate) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *AddEmailTemplateRequestEmailTemplate) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *AddEmailTemplateRequestEmailTemplate) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil
### GetTemplate

`func (o *AddEmailTemplateRequestEmailTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AddEmailTemplateRequestEmailTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AddEmailTemplateRequestEmailTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *AddEmailTemplateRequestEmailTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


