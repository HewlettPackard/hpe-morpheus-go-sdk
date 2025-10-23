# ListCloudFolders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folders** | Pointer to [**[]ListCloudFolders200ResponseAllOfFoldersInner**](ListCloudFolders200ResponseAllOfFoldersInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCloudFolders200Response

`func NewListCloudFolders200Response() *ListCloudFolders200Response`

NewListCloudFolders200Response instantiates a new ListCloudFolders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudFolders200ResponseWithDefaults

`func NewListCloudFolders200ResponseWithDefaults() *ListCloudFolders200Response`

NewListCloudFolders200ResponseWithDefaults instantiates a new ListCloudFolders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolders

`func (o *ListCloudFolders200Response) GetFolders() []ListCloudFolders200ResponseAllOfFoldersInner`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *ListCloudFolders200Response) GetFoldersOk() (*[]ListCloudFolders200ResponseAllOfFoldersInner, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *ListCloudFolders200Response) SetFolders(v []ListCloudFolders200ResponseAllOfFoldersInner)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *ListCloudFolders200Response) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetMeta

`func (o *ListCloudFolders200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCloudFolders200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCloudFolders200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCloudFolders200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


