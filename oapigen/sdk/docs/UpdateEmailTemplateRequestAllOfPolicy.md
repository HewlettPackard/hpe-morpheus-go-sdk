# UpdateEmailTemplateRequestAllOfPolicy

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

### NewUpdateEmailTemplateRequestAllOfPolicy

`func NewUpdateEmailTemplateRequestAllOfPolicy() *UpdateEmailTemplateRequestAllOfPolicy`

NewUpdateEmailTemplateRequestAllOfPolicy instantiates a new UpdateEmailTemplateRequestAllOfPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEmailTemplateRequestAllOfPolicyWithDefaults

`func NewUpdateEmailTemplateRequestAllOfPolicyWithDefaults() *UpdateEmailTemplateRequestAllOfPolicy`

NewUpdateEmailTemplateRequestAllOfPolicyWithDefaults instantiates a new UpdateEmailTemplateRequestAllOfPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateEmailTemplateRequestAllOfPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateEmailTemplateRequestAllOfPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEmailTemplateRequestAllOfPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateEmailTemplateRequestAllOfPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateEmailTemplateRequestAllOfPolicy) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateEmailTemplateRequestAllOfPolicy) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetOwner() AddEmailTemplateRequestEmailTemplateOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetOwnerOk() (*AddEmailTemplateRequestEmailTemplateOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateEmailTemplateRequestAllOfPolicy) SetOwner(v AddEmailTemplateRequestEmailTemplateOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateEmailTemplateRequestAllOfPolicy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccounts

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetAccounts() []AddEmailTemplateRequestEmailTemplateAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetAccountsOk() (*[]AddEmailTemplateRequestEmailTemplateAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdateEmailTemplateRequestAllOfPolicy) SetAccounts(v []AddEmailTemplateRequestEmailTemplateAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UpdateEmailTemplateRequestAllOfPolicy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *UpdateEmailTemplateRequestAllOfPolicy) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *UpdateEmailTemplateRequestAllOfPolicy) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil
### GetTemplate

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UpdateEmailTemplateRequestAllOfPolicy) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UpdateEmailTemplateRequestAllOfPolicy) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *UpdateEmailTemplateRequestAllOfPolicy) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


