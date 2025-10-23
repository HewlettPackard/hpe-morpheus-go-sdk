# AddContactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | [**AddContactsRequestContact**](AddContactsRequestContact.md) |  | 

## Methods

### NewAddContactsRequest

`func NewAddContactsRequest(contact AddContactsRequestContact, ) *AddContactsRequest`

NewAddContactsRequest instantiates a new AddContactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddContactsRequestWithDefaults

`func NewAddContactsRequestWithDefaults() *AddContactsRequest`

NewAddContactsRequestWithDefaults instantiates a new AddContactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *AddContactsRequest) GetContact() AddContactsRequestContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AddContactsRequest) GetContactOk() (*AddContactsRequestContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AddContactsRequest) SetContact(v AddContactsRequestContact)`

SetContact sets Contact field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


