# GetCypherKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**GetCypherKey200ResponseAllOfData**](GetCypherKey200ResponseAllOfData.md) |  | [optional] 
**Type** | Pointer to **string** | Type of data that was written to the key | [optional] 
**LeaseDuration** | Pointer to **int32** | Lease duration in seconds, 0 means no expiry. | [optional] 
**Cypher** | Pointer to [**ListCypherKeys200ResponseAllOfCyphersInner**](ListCypherKeys200ResponseAllOfCyphersInner.md) |  | [optional] 

## Methods

### NewGetCypherKey200Response

`func NewGetCypherKey200Response() *GetCypherKey200Response`

NewGetCypherKey200Response instantiates a new GetCypherKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCypherKey200ResponseWithDefaults

`func NewGetCypherKey200ResponseWithDefaults() *GetCypherKey200Response`

NewGetCypherKey200ResponseWithDefaults instantiates a new GetCypherKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetCypherKey200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetCypherKey200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetCypherKey200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetCypherKey200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *GetCypherKey200Response) GetData() GetCypherKey200ResponseAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCypherKey200Response) GetDataOk() (*GetCypherKey200ResponseAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCypherKey200Response) SetData(v GetCypherKey200ResponseAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *GetCypherKey200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *GetCypherKey200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCypherKey200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCypherKey200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCypherKey200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLeaseDuration

`func (o *GetCypherKey200Response) GetLeaseDuration() int32`

GetLeaseDuration returns the LeaseDuration field if non-nil, zero value otherwise.

### GetLeaseDurationOk

`func (o *GetCypherKey200Response) GetLeaseDurationOk() (*int32, bool)`

GetLeaseDurationOk returns a tuple with the LeaseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDuration

`func (o *GetCypherKey200Response) SetLeaseDuration(v int32)`

SetLeaseDuration sets LeaseDuration field to given value.

### HasLeaseDuration

`func (o *GetCypherKey200Response) HasLeaseDuration() bool`

HasLeaseDuration returns a boolean if a field has been set.

### GetCypher

`func (o *GetCypherKey200Response) GetCypher() ListCypherKeys200ResponseAllOfCyphersInner`

GetCypher returns the Cypher field if non-nil, zero value otherwise.

### GetCypherOk

`func (o *GetCypherKey200Response) GetCypherOk() (*ListCypherKeys200ResponseAllOfCyphersInner, bool)`

GetCypherOk returns a tuple with the Cypher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCypher

`func (o *GetCypherKey200Response) SetCypher(v ListCypherKeys200ResponseAllOfCyphersInner)`

SetCypher sets Cypher field to given value.

### HasCypher

`func (o *GetCypherKey200Response) HasCypher() bool`

HasCypher returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


