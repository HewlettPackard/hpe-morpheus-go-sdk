# ListCloudDatastores200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastores** | Pointer to [**[]ListCloudDatastores200ResponseAllOfDatastoresInner**](ListCloudDatastores200ResponseAllOfDatastoresInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCloudDatastores200Response

`func NewListCloudDatastores200Response() *ListCloudDatastores200Response`

NewListCloudDatastores200Response instantiates a new ListCloudDatastores200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudDatastores200ResponseWithDefaults

`func NewListCloudDatastores200ResponseWithDefaults() *ListCloudDatastores200Response`

NewListCloudDatastores200ResponseWithDefaults instantiates a new ListCloudDatastores200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastores

`func (o *ListCloudDatastores200Response) GetDatastores() []ListCloudDatastores200ResponseAllOfDatastoresInner`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *ListCloudDatastores200Response) GetDatastoresOk() (*[]ListCloudDatastores200ResponseAllOfDatastoresInner, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *ListCloudDatastores200Response) SetDatastores(v []ListCloudDatastores200ResponseAllOfDatastoresInner)`

SetDatastores sets Datastores field to given value.

### HasDatastores

`func (o *ListCloudDatastores200Response) HasDatastores() bool`

HasDatastores returns a boolean if a field has been set.

### GetMeta

`func (o *ListCloudDatastores200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCloudDatastores200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCloudDatastores200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCloudDatastores200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


