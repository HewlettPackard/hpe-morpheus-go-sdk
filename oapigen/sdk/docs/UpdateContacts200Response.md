# UpdateContacts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to [**UpdateContacts200ResponseAllOfContact**](UpdateContacts200ResponseAllOfContact.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateContacts200Response

`func NewUpdateContacts200Response() *UpdateContacts200Response`

NewUpdateContacts200Response instantiates a new UpdateContacts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetContact

`func (o *UpdateContacts200Response) GetContact() UpdateContacts200ResponseAllOfContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *UpdateContacts200Response) GetContactOk() (*UpdateContacts200ResponseAllOfContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *UpdateContacts200Response) SetContact(v UpdateContacts200ResponseAllOfContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *UpdateContacts200Response) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateContacts200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateContacts200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateContacts200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateContacts200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


