# UpdateWikiApp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**GetWikiApp200ResponsePage**](GetWikiApp200ResponsePage.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateWikiApp200Response

`func NewUpdateWikiApp200Response() *UpdateWikiApp200Response`

NewUpdateWikiApp200Response instantiates a new UpdateWikiApp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWikiApp200ResponseWithDefaults

`func NewUpdateWikiApp200ResponseWithDefaults() *UpdateWikiApp200Response`

NewUpdateWikiApp200ResponseWithDefaults instantiates a new UpdateWikiApp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *UpdateWikiApp200Response) GetPage() GetWikiApp200ResponsePage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UpdateWikiApp200Response) GetPageOk() (*GetWikiApp200ResponsePage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UpdateWikiApp200Response) SetPage(v GetWikiApp200ResponsePage)`

SetPage sets Page field to given value.

### HasPage

`func (o *UpdateWikiApp200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateWikiApp200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateWikiApp200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateWikiApp200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateWikiApp200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


