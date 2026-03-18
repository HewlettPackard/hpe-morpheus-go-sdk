# UpdateContactsRequestContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name scoped to your account for the contact | [optional] 
**EmailAddress** | Pointer to **string** | Email notification address | [optional] 
**SmsAddress** | Pointer to **string** | SMS notification address | [optional] 
**SlackHook** | Pointer to **string** | Slack Hook | [optional] 

## Methods

### NewUpdateContactsRequestContact

`func NewUpdateContactsRequestContact() *UpdateContactsRequestContact`

NewUpdateContactsRequestContact instantiates a new UpdateContactsRequestContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContactsRequestContactWithDefaults

`func NewUpdateContactsRequestContactWithDefaults() *UpdateContactsRequestContact`

NewUpdateContactsRequestContactWithDefaults instantiates a new UpdateContactsRequestContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateContactsRequestContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateContactsRequestContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateContactsRequestContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateContactsRequestContact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *UpdateContactsRequestContact) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UpdateContactsRequestContact) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UpdateContactsRequestContact) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *UpdateContactsRequestContact) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetSmsAddress

`func (o *UpdateContactsRequestContact) GetSmsAddress() string`

GetSmsAddress returns the SmsAddress field if non-nil, zero value otherwise.

### GetSmsAddressOk

`func (o *UpdateContactsRequestContact) GetSmsAddressOk() (*string, bool)`

GetSmsAddressOk returns a tuple with the SmsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsAddress

`func (o *UpdateContactsRequestContact) SetSmsAddress(v string)`

SetSmsAddress sets SmsAddress field to given value.

### HasSmsAddress

`func (o *UpdateContactsRequestContact) HasSmsAddress() bool`

HasSmsAddress returns a boolean if a field has been set.

### GetSlackHook

`func (o *UpdateContactsRequestContact) GetSlackHook() string`

GetSlackHook returns the SlackHook field if non-nil, zero value otherwise.

### GetSlackHookOk

`func (o *UpdateContactsRequestContact) GetSlackHookOk() (*string, bool)`

GetSlackHookOk returns a tuple with the SlackHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackHook

`func (o *UpdateContactsRequestContact) SetSlackHook(v string)`

SetSlackHook sets SlackHook field to given value.

### HasSlackHook

`func (o *UpdateContactsRequestContact) HasSlackHook() bool`

HasSlackHook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


