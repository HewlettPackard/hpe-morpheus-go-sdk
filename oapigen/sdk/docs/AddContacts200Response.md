# AddContacts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**ListContacts200ResponseAllOfContactsInner**](ListContacts200ResponseAllOfContactsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddContacts200Response

`func NewAddContacts200Response() *AddContacts200Response`

NewAddContacts200Response instantiates a new AddContacts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddContacts200ResponseWithDefaults

`func NewAddContacts200ResponseWithDefaults() *AddContacts200Response`

NewAddContacts200ResponseWithDefaults instantiates a new AddContacts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *AddContacts200Response) GetContact() ListContacts200ResponseAllOfContactsInner`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AddContacts200Response) GetContactOk() (*ListContacts200ResponseAllOfContactsInner, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AddContacts200Response) SetContact(v ListContacts200ResponseAllOfContactsInner)`

SetContact sets Contact field to given value.

### HasContact

`func (o *AddContacts200Response) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetSuccess

`func (o *AddContacts200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddContacts200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddContacts200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddContacts200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


