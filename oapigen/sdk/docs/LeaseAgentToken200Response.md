# LeaseAgentToken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Indicates the request succeeded. | 
**Url** | **string** | Websocket endpoint URL for the token scope. | 
**Headers** | [**LeaseAgentToken200ResponseHeaders**](LeaseAgentToken200ResponseHeaders.md) |  | 
**Data** | **map[string]interface{}** | Additional lease payload returned by the server. | 

## Methods

### NewLeaseAgentToken200Response

`func NewLeaseAgentToken200Response(success bool, url string, headers LeaseAgentToken200ResponseHeaders, data map[string]interface{}, ) *LeaseAgentToken200Response`

NewLeaseAgentToken200Response instantiates a new LeaseAgentToken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaseAgentToken200ResponseWithDefaults

`func NewLeaseAgentToken200ResponseWithDefaults() *LeaseAgentToken200Response`

NewLeaseAgentToken200ResponseWithDefaults instantiates a new LeaseAgentToken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *LeaseAgentToken200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *LeaseAgentToken200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *LeaseAgentToken200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetUrl

`func (o *LeaseAgentToken200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LeaseAgentToken200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LeaseAgentToken200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHeaders

`func (o *LeaseAgentToken200Response) GetHeaders() LeaseAgentToken200ResponseHeaders`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *LeaseAgentToken200Response) GetHeadersOk() (*LeaseAgentToken200ResponseHeaders, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *LeaseAgentToken200Response) SetHeaders(v LeaseAgentToken200ResponseHeaders)`

SetHeaders sets Headers field to given value.


### GetData

`func (o *LeaseAgentToken200Response) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LeaseAgentToken200Response) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LeaseAgentToken200Response) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


