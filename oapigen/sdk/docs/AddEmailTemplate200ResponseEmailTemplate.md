# AddEmailTemplate200ResponseEmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | The name of the email template. This is set by morpheus.  | [optional] 
**Code** | Pointer to **string** | A unique code for the email template. This code is used to reference the email template and as a reference of the templates type.  | [optional] 
**Owner** | Pointer to [**AddEmailTemplate200ResponseEmailTemplateOwner**](AddEmailTemplate200ResponseEmailTemplateOwner.md) |  | [optional] 
**Accounts** | Pointer to [**[]AddEmailTemplate200ResponseEmailTemplateAccountsInner**](AddEmailTemplate200ResponseEmailTemplateAccountsInner.md) |  | [optional] 
**Template** | Pointer to **string** | The email template. This is the actual email template that is sent to the user. This uses handlebars notation (not javascript)  | [optional] 

## Methods

### NewAddEmailTemplate200ResponseEmailTemplate

`func NewAddEmailTemplate200ResponseEmailTemplate() *AddEmailTemplate200ResponseEmailTemplate`

NewAddEmailTemplate200ResponseEmailTemplate instantiates a new AddEmailTemplate200ResponseEmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEmailTemplate200ResponseEmailTemplateWithDefaults

`func NewAddEmailTemplate200ResponseEmailTemplateWithDefaults() *AddEmailTemplate200ResponseEmailTemplate`

NewAddEmailTemplate200ResponseEmailTemplateWithDefaults instantiates a new AddEmailTemplate200ResponseEmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddEmailTemplate200ResponseEmailTemplate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddEmailTemplate200ResponseEmailTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddEmailTemplate200ResponseEmailTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddEmailTemplate200ResponseEmailTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddEmailTemplate200ResponseEmailTemplate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddEmailTemplate200ResponseEmailTemplate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetOwner

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetOwner() AddEmailTemplate200ResponseEmailTemplateOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetOwnerOk() (*AddEmailTemplate200ResponseEmailTemplateOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddEmailTemplate200ResponseEmailTemplate) SetOwner(v AddEmailTemplate200ResponseEmailTemplateOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddEmailTemplate200ResponseEmailTemplate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccounts

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetAccounts() []AddEmailTemplate200ResponseEmailTemplateAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetAccountsOk() (*[]AddEmailTemplate200ResponseEmailTemplateAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AddEmailTemplate200ResponseEmailTemplate) SetAccounts(v []AddEmailTemplate200ResponseEmailTemplateAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AddEmailTemplate200ResponseEmailTemplate) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *AddEmailTemplate200ResponseEmailTemplate) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *AddEmailTemplate200ResponseEmailTemplate) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil
### GetTemplate

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AddEmailTemplate200ResponseEmailTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AddEmailTemplate200ResponseEmailTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *AddEmailTemplate200ResponseEmailTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


