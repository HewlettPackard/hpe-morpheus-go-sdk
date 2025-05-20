# ResizeInstance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**UpdateInstance200ResponseAllOfOneOfInstance**](UpdateInstance200ResponseAllOfOneOfInstance.md) |  | 
**ZoneId** | **int64** | The Cloud ID to provision the instance onto. | 
**Success** | Pointer to **bool** |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewResizeInstance200Response

`func NewResizeInstance200Response(instance UpdateInstance200ResponseAllOfOneOfInstance, zoneId int64, ) *ResizeInstance200Response`

NewResizeInstance200Response instantiates a new ResizeInstance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeInstance200ResponseWithDefaults

`func NewResizeInstance200ResponseWithDefaults() *ResizeInstance200Response`

NewResizeInstance200ResponseWithDefaults instantiates a new ResizeInstance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *ResizeInstance200Response) GetInstance() UpdateInstance200ResponseAllOfOneOfInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ResizeInstance200Response) GetInstanceOk() (*UpdateInstance200ResponseAllOfOneOfInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ResizeInstance200Response) SetInstance(v UpdateInstance200ResponseAllOfOneOfInstance)`

SetInstance sets Instance field to given value.


### GetZoneId

`func (o *ResizeInstance200Response) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ResizeInstance200Response) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ResizeInstance200Response) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.


### GetSuccess

`func (o *ResizeInstance200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResizeInstance200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResizeInstance200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResizeInstance200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrors

`func (o *ResizeInstance200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ResizeInstance200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ResizeInstance200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ResizeInstance200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


