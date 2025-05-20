# EmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | The name of the email template. This is set by morpheus.  | [optional] 
**Code** | Pointer to **string** | A unique code for the email template. This code is used to reference the email template and as a reference of the templates type.  | [optional] 
**Owner** | Pointer to [**ListEmailTemplates200ResponseAllOfEmailTemplatesInnerOwner**](ListEmailTemplates200ResponseAllOfEmailTemplatesInnerOwner.md) |  | [optional] 
**Accounts** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Template** | Pointer to **string** | The email template. This is the actual email template that is sent to the user. This uses handlebars notation (not javascript)  | [optional] 

## Methods

### NewEmailTemplate

`func NewEmailTemplate() *EmailTemplate`

NewEmailTemplate instantiates a new EmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateWithDefaults

`func NewEmailTemplateWithDefaults() *EmailTemplate`

NewEmailTemplateWithDefaults instantiates a new EmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailTemplate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EmailTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EmailTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *EmailTemplate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EmailTemplate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EmailTemplate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EmailTemplate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetOwner

`func (o *EmailTemplate) GetOwner() ListEmailTemplates200ResponseAllOfEmailTemplatesInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *EmailTemplate) GetOwnerOk() (*ListEmailTemplates200ResponseAllOfEmailTemplatesInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *EmailTemplate) SetOwner(v ListEmailTemplates200ResponseAllOfEmailTemplatesInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *EmailTemplate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccounts

`func (o *EmailTemplate) GetAccounts() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *EmailTemplate) GetAccountsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *EmailTemplate) SetAccounts(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *EmailTemplate) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetTemplate

`func (o *EmailTemplate) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EmailTemplate) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EmailTemplate) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EmailTemplate) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


