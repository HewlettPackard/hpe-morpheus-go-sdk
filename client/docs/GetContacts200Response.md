# GetContacts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**ListContacts200ResponseAllOfContactsInner**](ListContacts200ResponseAllOfContactsInner.md) |  | [optional] 

## Methods

### NewGetContacts200Response

`func NewGetContacts200Response() *GetContacts200Response`

NewGetContacts200Response instantiates a new GetContacts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContacts200ResponseWithDefaults

`func NewGetContacts200ResponseWithDefaults() *GetContacts200Response`

NewGetContacts200ResponseWithDefaults instantiates a new GetContacts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *GetContacts200Response) GetContact() ListContacts200ResponseAllOfContactsInner`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *GetContacts200Response) GetContactOk() (*ListContacts200ResponseAllOfContactsInner, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *GetContacts200Response) SetContact(v ListContacts200ResponseAllOfContactsInner)`

SetContact sets Contact field to given value.

### HasContact

`func (o *GetContacts200Response) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


