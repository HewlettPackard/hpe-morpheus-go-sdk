# ListEmailTemplates200ResponseAllOfEmailTemplatesInner

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

### NewListEmailTemplates200ResponseAllOfEmailTemplatesInner

`func NewListEmailTemplates200ResponseAllOfEmailTemplatesInner() *ListEmailTemplates200ResponseAllOfEmailTemplatesInner`

NewListEmailTemplates200ResponseAllOfEmailTemplatesInner instantiates a new ListEmailTemplates200ResponseAllOfEmailTemplatesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEmailTemplates200ResponseAllOfEmailTemplatesInnerWithDefaults

`func NewListEmailTemplates200ResponseAllOfEmailTemplatesInnerWithDefaults() *ListEmailTemplates200ResponseAllOfEmailTemplatesInner`

NewListEmailTemplates200ResponseAllOfEmailTemplatesInnerWithDefaults instantiates a new ListEmailTemplates200ResponseAllOfEmailTemplatesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetOwner

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetOwner() ListEmailTemplates200ResponseAllOfEmailTemplatesInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetOwnerOk() (*ListEmailTemplates200ResponseAllOfEmailTemplatesInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) SetOwner(v ListEmailTemplates200ResponseAllOfEmailTemplatesInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccounts

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetAccounts() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetAccountsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) SetAccounts(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetTemplate

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ListEmailTemplates200ResponseAllOfEmailTemplatesInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


