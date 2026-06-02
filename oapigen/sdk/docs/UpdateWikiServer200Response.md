# UpdateWikiServer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**UpdateWikiCluster200ResponseAllOfPage**](UpdateWikiCluster200ResponseAllOfPage.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateWikiServer200Response

`func NewUpdateWikiServer200Response() *UpdateWikiServer200Response`

NewUpdateWikiServer200Response instantiates a new UpdateWikiServer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetPage

`func (o *UpdateWikiServer200Response) GetPage() UpdateWikiCluster200ResponseAllOfPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UpdateWikiServer200Response) GetPageOk() (*UpdateWikiCluster200ResponseAllOfPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UpdateWikiServer200Response) SetPage(v UpdateWikiCluster200ResponseAllOfPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *UpdateWikiServer200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateWikiServer200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateWikiServer200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateWikiServer200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateWikiServer200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


