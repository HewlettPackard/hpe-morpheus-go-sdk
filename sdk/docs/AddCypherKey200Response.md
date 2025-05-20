# AddCypherKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**LeaseDuration** | Pointer to **int64** |  | [optional] 
**Cypher** | Pointer to [**ListCypherKeys200ResponseAllOfCyphersInner**](ListCypherKeys200ResponseAllOfCyphersInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddCypherKey200Response

`func NewAddCypherKey200Response() *AddCypherKey200Response`

NewAddCypherKey200Response instantiates a new AddCypherKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCypherKey200ResponseWithDefaults

`func NewAddCypherKey200ResponseWithDefaults() *AddCypherKey200Response`

NewAddCypherKey200ResponseWithDefaults instantiates a new AddCypherKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AddCypherKey200Response) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddCypherKey200Response) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddCypherKey200Response) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *AddCypherKey200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *AddCypherKey200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCypherKey200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCypherKey200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddCypherKey200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLeaseDuration

`func (o *AddCypherKey200Response) GetLeaseDuration() int64`

GetLeaseDuration returns the LeaseDuration field if non-nil, zero value otherwise.

### GetLeaseDurationOk

`func (o *AddCypherKey200Response) GetLeaseDurationOk() (*int64, bool)`

GetLeaseDurationOk returns a tuple with the LeaseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDuration

`func (o *AddCypherKey200Response) SetLeaseDuration(v int64)`

SetLeaseDuration sets LeaseDuration field to given value.

### HasLeaseDuration

`func (o *AddCypherKey200Response) HasLeaseDuration() bool`

HasLeaseDuration returns a boolean if a field has been set.

### GetCypher

`func (o *AddCypherKey200Response) GetCypher() ListCypherKeys200ResponseAllOfCyphersInner`

GetCypher returns the Cypher field if non-nil, zero value otherwise.

### GetCypherOk

`func (o *AddCypherKey200Response) GetCypherOk() (*ListCypherKeys200ResponseAllOfCyphersInner, bool)`

GetCypherOk returns a tuple with the Cypher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCypher

`func (o *AddCypherKey200Response) SetCypher(v ListCypherKeys200ResponseAllOfCyphersInner)`

SetCypher sets Cypher field to given value.

### HasCypher

`func (o *AddCypherKey200Response) HasCypher() bool`

HasCypher returns a boolean if a field has been set.

### GetSuccess

`func (o *AddCypherKey200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddCypherKey200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddCypherKey200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddCypherKey200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


