# AddInstance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**AddInstance200ResponseAllOfOneOfInstance**](AddInstance200ResponseAllOfOneOfInstance.md) |  | 
**ZoneId** | **int64** | The Cloud ID to provision the instance onto. | 
**Success** | Pointer to **bool** |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAddInstance200Response

`func NewAddInstance200Response(instance AddInstance200ResponseAllOfOneOfInstance, zoneId int64, ) *AddInstance200Response`

NewAddInstance200Response instantiates a new AddInstance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstance200ResponseWithDefaults

`func NewAddInstance200ResponseWithDefaults() *AddInstance200Response`

NewAddInstance200ResponseWithDefaults instantiates a new AddInstance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *AddInstance200Response) GetInstance() AddInstance200ResponseAllOfOneOfInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddInstance200Response) GetInstanceOk() (*AddInstance200ResponseAllOfOneOfInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddInstance200Response) SetInstance(v AddInstance200ResponseAllOfOneOfInstance)`

SetInstance sets Instance field to given value.


### GetZoneId

`func (o *AddInstance200Response) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddInstance200Response) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddInstance200Response) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.


### GetSuccess

`func (o *AddInstance200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddInstance200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddInstance200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddInstance200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrors

`func (o *AddInstance200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AddInstance200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AddInstance200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AddInstance200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


