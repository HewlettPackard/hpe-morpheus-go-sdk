# UpdateInvoicesRequestInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]map[string]interface{}** | This adds or updates the specified Metadata tags and removes any tags not specified. Array of objects having a name and value.  | [optional] 
**AddTags** | Pointer to **[]map[string]interface{}** | Add or update value of Metadata tags. Array of objects having a name and value.  | [optional] 
**RemoveTags** | Pointer to **[]map[string]interface{}** | This removes the specified Metadata tags matching name and optionally value (if provided). Array of objects having a name and value.  | [optional] 

## Methods

### NewUpdateInvoicesRequestInvoice

`func NewUpdateInvoicesRequestInvoice() *UpdateInvoicesRequestInvoice`

NewUpdateInvoicesRequestInvoice instantiates a new UpdateInvoicesRequestInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInvoicesRequestInvoiceWithDefaults

`func NewUpdateInvoicesRequestInvoiceWithDefaults() *UpdateInvoicesRequestInvoice`

NewUpdateInvoicesRequestInvoiceWithDefaults instantiates a new UpdateInvoicesRequestInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *UpdateInvoicesRequestInvoice) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateInvoicesRequestInvoice) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateInvoicesRequestInvoice) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateInvoicesRequestInvoice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAddTags

`func (o *UpdateInvoicesRequestInvoice) GetAddTags() []map[string]interface{}`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *UpdateInvoicesRequestInvoice) GetAddTagsOk() (*[]map[string]interface{}, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *UpdateInvoicesRequestInvoice) SetAddTags(v []map[string]interface{})`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *UpdateInvoicesRequestInvoice) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *UpdateInvoicesRequestInvoice) GetRemoveTags() []map[string]interface{}`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *UpdateInvoicesRequestInvoice) GetRemoveTagsOk() (*[]map[string]interface{}, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *UpdateInvoicesRequestInvoice) SetRemoveTags(v []map[string]interface{})`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *UpdateInvoicesRequestInvoice) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


