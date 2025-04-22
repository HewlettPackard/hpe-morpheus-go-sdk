# ListContacts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | Pointer to [**[]ListContacts200ResponseAllOfContactsInner**](ListContacts200ResponseAllOfContactsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListContacts200Response

`func NewListContacts200Response() *ListContacts200Response`

NewListContacts200Response instantiates a new ListContacts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListContacts200ResponseWithDefaults

`func NewListContacts200ResponseWithDefaults() *ListContacts200Response`

NewListContacts200ResponseWithDefaults instantiates a new ListContacts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *ListContacts200Response) GetContacts() []ListContacts200ResponseAllOfContactsInner`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ListContacts200Response) GetContactsOk() (*[]ListContacts200ResponseAllOfContactsInner, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ListContacts200Response) SetContacts(v []ListContacts200ResponseAllOfContactsInner)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ListContacts200Response) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetMeta

`func (o *ListContacts200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListContacts200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListContacts200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListContacts200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


