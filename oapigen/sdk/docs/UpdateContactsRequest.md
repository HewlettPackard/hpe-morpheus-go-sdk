# UpdateContactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | [**UpdateContactsRequestContact**](UpdateContactsRequestContact.md) |  | 

## Methods

### NewUpdateContactsRequest

`func NewUpdateContactsRequest(contact UpdateContactsRequestContact, ) *UpdateContactsRequest`

NewUpdateContactsRequest instantiates a new UpdateContactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContactsRequestWithDefaults

`func NewUpdateContactsRequestWithDefaults() *UpdateContactsRequest`

NewUpdateContactsRequestWithDefaults instantiates a new UpdateContactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *UpdateContactsRequest) GetContact() UpdateContactsRequestContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *UpdateContactsRequest) GetContactOk() (*UpdateContactsRequestContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *UpdateContactsRequest) SetContact(v UpdateContactsRequestContact)`

SetContact sets Contact field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


