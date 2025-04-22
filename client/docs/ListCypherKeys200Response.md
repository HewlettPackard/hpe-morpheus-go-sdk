# ListCypherKeys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ListCypherKeys200ResponseAllOfData**](ListCypherKeys200ResponseAllOfData.md) |  | [optional] 
**Cyphers** | Pointer to [**[]ListCypherKeys200ResponseAllOfCyphersInner**](ListCypherKeys200ResponseAllOfCyphersInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewListCypherKeys200Response

`func NewListCypherKeys200Response() *ListCypherKeys200Response`

NewListCypherKeys200Response instantiates a new ListCypherKeys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCypherKeys200ResponseWithDefaults

`func NewListCypherKeys200ResponseWithDefaults() *ListCypherKeys200Response`

NewListCypherKeys200ResponseWithDefaults instantiates a new ListCypherKeys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCypherKeys200Response) GetData() ListCypherKeys200ResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCypherKeys200Response) GetDataOk() (*ListCypherKeys200ResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCypherKeys200Response) SetData(v ListCypherKeys200ResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *ListCypherKeys200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetCyphers

`func (o *ListCypherKeys200Response) GetCyphers() []ListCypherKeys200ResponseAllOfCyphersInner`

GetCyphers returns the Cyphers field if non-nil, zero value otherwise.

### GetCyphersOk

`func (o *ListCypherKeys200Response) GetCyphersOk() (*[]ListCypherKeys200ResponseAllOfCyphersInner, bool)`

GetCyphersOk returns a tuple with the Cyphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCyphers

`func (o *ListCypherKeys200Response) SetCyphers(v []ListCypherKeys200ResponseAllOfCyphersInner)`

SetCyphers sets Cyphers field to given value.

### HasCyphers

`func (o *ListCypherKeys200Response) HasCyphers() bool`

HasCyphers returns a boolean if a field has been set.

### GetMeta

`func (o *ListCypherKeys200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCypherKeys200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCypherKeys200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCypherKeys200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSuccess

`func (o *ListCypherKeys200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListCypherKeys200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListCypherKeys200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListCypherKeys200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


