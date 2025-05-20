# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SmsAddress** | Pointer to **string** |  | [optional] 
**SlackHook** | Pointer to **string** |  | [optional] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contact) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contact) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contact) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Contact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmailAddress

`func (o *Contact) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Contact) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Contact) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *Contact) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetName

`func (o *Contact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Contact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Contact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Contact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSmsAddress

`func (o *Contact) GetSmsAddress() string`

GetSmsAddress returns the SmsAddress field if non-nil, zero value otherwise.

### GetSmsAddressOk

`func (o *Contact) GetSmsAddressOk() (*string, bool)`

GetSmsAddressOk returns a tuple with the SmsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsAddress

`func (o *Contact) SetSmsAddress(v string)`

SetSmsAddress sets SmsAddress field to given value.

### HasSmsAddress

`func (o *Contact) HasSmsAddress() bool`

HasSmsAddress returns a boolean if a field has been set.

### GetSlackHook

`func (o *Contact) GetSlackHook() string`

GetSlackHook returns the SlackHook field if non-nil, zero value otherwise.

### GetSlackHookOk

`func (o *Contact) GetSlackHookOk() (*string, bool)`

GetSlackHookOk returns a tuple with the SlackHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackHook

`func (o *Contact) SetSlackHook(v string)`

SetSlackHook sets SlackHook field to given value.

### HasSlackHook

`func (o *Contact) HasSlackHook() bool`

HasSlackHook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


